package tailf

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime/debug"
	"strings"
	"sync"
	"time"
	"util"
	"github.com/gobwas/glob"
	"github.com/hpcloud/tail"
	"gopkg.in/tomb.v1"
)

type TailObj struct {
	Files      []*TailFileObj
	tomb       tomb.Tomb
	lines      chan *tail.Line
	Topic      string
	pattern    string
	prefixType int
	splitType  int
	lock       sync.Mutex
}

type TailFileObj struct {
	Filename   string
	pathmd5    string
	tails      *tail.Tail
	Offset     int64
	isStop     bool
	RecordTime int64
}

const (
	BUFFER_LIMIT             = 1024 * 1024 // 1M
	LINE_SEPARATOR           = "\t"
	RESERVERD_LINE_SEPARATOR = "\n"
)
const (
	ONE_LINE_TYPE = iota
	SPACE_LINE_TYPE
	TIME_LINE_TYPE
)
const (
	NONE_PREFIX = iota
	HOST_PREFIX
	IP_PREFIX
)

var (
	OFFSET_DIR string
	ErrMsgCh   chan error
)

func init() {
	file, _ := exec.LookPath(os.Args[0])
	dir, _ := path.Split(file)
	OFFSET_DIR = fmt.Sprintf("%s/../%s", dir, "offset")

	if !Exist(OFFSET_DIR) {
		err := os.Mkdir(fmt.Sprintf("%s/../%s", dir, "offset"), 0755)
		if err != nil {
			fmt.Println("mkdir %v err: %v", OFFSET_DIR, err)
		}
	}
	ErrMsgCh = make(chan error, 1024)
}

func Exist(dirName string) bool {
	_, err := os.Stat(dirName)
	return err == nil || os.IsExist(err)
}

func NewTail(pattern string, splitType int, prefixType int, topic string) (obj *TailObj, err error) {
	obj = &TailObj{
		splitType:  splitType,
		prefixType: prefixType,
		Topic:      topic,
	}
	obj.pattern = pattern
	obj.lines = make(chan *tail.Line)
	go obj.monitorFile(pattern)
	return obj, nil
}

//if offset > size, p.Offset = -1
func (p *TailFileObj) readOffset(topic string) {
	//when log file is collected by two topics
	b := md5.Sum([]byte(topic + "#" + p.Filename))
	p.pathmd5 = string(b[:])
	data, err := ioutil.ReadFile(fmt.Sprintf("%s/%x.oft", OFFSET_DIR, p.pathmd5))
	if err != nil {
		p.Offset = -1
		return
	}
	var obj TailFileObj
	if err := json.Unmarshal(data, &obj); err != nil {
		p.Offset = -1
		return
	}
	if obj.Offset == 0 && obj.RecordTime < time.Now().Add(-time.Hour*24).Unix() { //never read file from start_offset
		p.Offset = -1
		return
	}
	newFileInfo, err := os.Stat(p.Filename)
	if err != nil {
		p.Offset = -1
		return
	}
	if newFileInfo.Size() < obj.Offset  {
		p.Offset = 0
		return
	}
	p.Offset = obj.Offset
}
func (p *TailFileObj) writeOffset() error {
	p.RecordTime = time.Now().Unix()
	b, err := json.Marshal(p)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("%s/%x.oft", OFFSET_DIR, p.pathmd5), b, 0644)
	if err != nil {
		return fmt.Errorf("path:%s, err:%s", fmt.Sprintf("%s/%x.oft", OFFSET_DIR, p.pathmd5), err.Error())
	}
	return nil
}
func (p *TailObj) monitorFile(filePattern string) {
	defer func() {
		if e := recover(); e != nil {
			ErrMsgCh <- fmt.Errorf("err:%s stack:%s", e, string(debug.Stack()))
		}
	}()
	pattern := filepath.Base(filePattern)
	if !((strings.Contains(pattern, "*") == true) ||
		(strings.Contains(pattern, "{") && strings.Contains(pattern, "}")) ||
		(strings.Contains(pattern, "[") && strings.Contains(pattern, "]"))) {
		if err := p.addFileTail(filePattern); err != nil {
			ErrMsgCh <- err
		}
		return
	}
	dir := filepath.Dir(filePattern)
	reg, err := glob.Compile(pattern)
	if err != nil {
		ErrMsgCh <- err
		return
	}
	var succ bool
	for {
		if err := p.tomb.Err(); err != tomb.ErrStillAlive {
			ErrMsgCh <- err
			return
		}
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			ErrMsgCh <- err
			time.Sleep(5 * time.Second)
		}
		for _, f := range files {
			if f.IsDir() {
				continue
			}
			if reg.Match(f.Name()) == false {
				continue
			}
			succ = true
			p.lock.Lock()
			for _, F := range p.Files {
				if filepath.Base(F.Filename) == f.Name() {
					succ = false
					break
				}
			}
			if succ {
				p.addFileTail(fmt.Sprintf("%s/%s", dir, f.Name()))
			}
			p.lock.Unlock()
		}

		time.Sleep(10 * time.Second)
	}
}
func (p *TailObj) GetPrefix() string {
	switch p.prefixType {
	case HOST_PREFIX:
		host, err := os.Hostname()
		if err != nil {
			return fmt.Sprintf("UNKNOWN_HOST:%s ", err)
		}
		return fmt.Sprintf("%s ", host)
	case IP_PREFIX:
		return util.GetLocalIP()
	}
	return "127.0.0.1 "
}
func (p *TailObj) collectLine(s *TailFileObj) {
	defer func() {
		if e := recover(); e != nil {
			ErrMsgCh <- fmt.Errorf("collect lines err:%s stack:%s", e, string(debug.Stack()))
		}
	}()
	var msg *tail.Line
	var ok bool
	var buffer bytes.Buffer
	if p.splitType != NONE_PREFIX {
		buffer.WriteString(p.GetPrefix())
	}
	for !s.isStop {
		msg, ok = <-s.tails.Lines
		if !ok {
			if buffer.Len() > 0 {
				p.lines <- tail.NewLine(buffer.String())
				buffer.Reset()
			}
			ErrMsgCh <- fmt.Errorf("tail file close reopen, filename:%s", s.tails.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		if p.splitType == ONE_LINE_TYPE {
			if p.prefixType != NONE_PREFIX {
				buffer.WriteString(msg.Text)
				msg.Text = buffer.String()
				buffer.Reset()
				buffer.WriteString(p.GetPrefix())
			}
			p.lines <- msg
		} else if p.splitType == SPACE_LINE_TYPE {
			size := len(msg.Text)
			text := msg.Text
			if size == 0 || buffer.Len() > BUFFER_LIMIT {
				if buffer.Len() != 0 {
					msg.Text = buffer.String()
					p.lines <- msg
					buffer.Reset()
					if p.prefixType != NONE_PREFIX {
						buffer.WriteString(p.GetPrefix())
					}
				}
			}
			if size > 0 {
				buffer.WriteString(text)
				buffer.WriteString(LINE_SEPARATOR)
			}

		} else if p.splitType == TIME_LINE_TYPE {
			size := len(msg.Text)
			text := msg.Text
			if phpTimeFormat([]byte(msg.Text)) || time.Now().After(msg.Time.Add(time.Minute*3)) || buffer.Len() > BUFFER_LIMIT {
				if buffer.Len() != 0 {
					prefix := ""
					if p.prefixType != NONE_PREFIX {
						prefix = p.GetPrefix()
					}
					msg.Text = buffer.String()
					if len(msg.Text) > len(prefix) {
						p.lines <- msg
					}
					buffer.Reset()
					buffer.WriteString(prefix)
				}
			}
			if size > 0 {
				buffer.WriteString(text)
				buffer.WriteString(LINE_SEPARATOR)
			}
		}
	}
}

func (p *TailObj) addFileTail(filename string) (err error) {
	obj := &TailFileObj{
		Filename: filename,
		isStop:   false,
	}
	obj.readOffset(p.Topic)
	var info *tail.SeekInfo
	if obj.Offset == -1 {
		info = &tail.SeekInfo{Offset: 0, Whence: 2}
	} else {
		info = &tail.SeekInfo{Offset: obj.Offset, Whence: 0}
	}
	obj.tails, err = tail.TailFile(filename, tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  info,
		MustExist: false,
		Poll:      true,
	})
	if err != nil {
		return err
	}
	p.Files = append(p.Files, obj)
	go p.collectLine(obj)
	//	go p.checkValidFile(obj)
	return nil
}

/*
//this may do some delete work..., if crash
func (t *TailObj) checkValidFile(p *TailFileObj) {
	defer func() {
		if e := recover(); e != nil {
			ErrMsgCh <- fmt.Errorf("err:%s stack:%s", e, string(debug.Stack()))
		}
		if p.isStop == false {
			p.isStop = true
			t.forceReopen(p)
		}
	}()
	var lastOffset int64 = 0
loop:
	for !p.isStop {
		time.Sleep(time.Minute)
		offset := p.Offset
		if p.isStop {
			return
		}
		tailFile := p.tails.CurFile()
		if tailFile == nil {
			lastOffset = offset
			continue
		}

		oldFileInfo, err := tailFile.Stat()
		if err != nil {
			lastOffset = offset
			continue
		}
		newFileInfo, err := os.Stat(p.Filename)
		if err != nil {
			if offset > oldFileInfo.Size() {
				ErrMsgCh <- fmt.Errorf("stop tailf , filename: %v not found, offset larger than oldfile.Size", p.Filename)
				break loop
			}
			lastOffset = offset
			continue
		}

		newStat, ok := newFileInfo.Sys().(*syscall.Stat_t)
		if !ok {
			lastOffset = offset
			continue
		}
		oldStat, ok := oldFileInfo.Sys().(*syscall.Stat_t)
		if !ok {
			lastOffset = offset
			continue
		}

		//如果inode变化了， 同时当前的offset 等于在读文件的offset， 重新打开文件, 保证读完旧文件
		if oldStat.Ino != newStat.Ino && offset == lastOffset {
			ErrMsgCh <- fmt.Errorf("stop tailf ,  filename: %v inode change, offset larger than oldfile.Size", p.Filename)
			break loop
		}
		lastOffset = offset
	}
}
*/
func (p *TailObj) GetOneLine() chan *tail.Line {
	return p.lines
}

func (p *TailObj) FlushOffset() (err error) {
	for _, f := range p.Files {
		f.Offset, err = f.tails.Tell()
		if err != nil {
			return fmt.Errorf("file:%s get current offset error:%v, topic:%s", f.Filename, err, p.Topic)
		}
		if err = f.writeOffset(); err != nil {
			return fmt.Errorf("file:%s write offset error:%v, topic:%s", f.Filename, err, p.Topic)
		}
	}
	return nil
}

func (p *TailFileObj) ForceStop() {
	p.isStop = true
	p.Offset, _ = p.tails.Tell()
	p.writeOffset()
	p.tails.Stop()
}
func (p *TailObj) ForceStop() {
	for _, f := range p.Files {
		f.ForceStop()
	}
	p.tomb.Kill(nil)
}

func (p *TailFileObj) Stop() (err error) {
	p.isStop = true
	p.Offset, err = p.tails.Tell()
	if err != nil {
		return
	}
	if err = p.writeOffset(); err != nil {
		return
	}
	return p.tails.Stop()
}

func (p *TailObj) Stop() (err error) {
	for _, f := range p.Files {
		if err = f.Stop(); err != nil {
			return
		}
	}
	p.tomb.Kill(nil)
	return
}

func (p *TailObj) forceReopen(t *TailFileObj) {
	// Always reopen truncated files (Follow is true)
	t.Stop()
	p.lock.Lock()
	defer p.lock.Unlock()
	for i, obj := range p.Files {
		if obj == t {
			p.Files = append((p.Files)[:i], (p.Files)[i+1:]...)
			break
		}
	}
}
