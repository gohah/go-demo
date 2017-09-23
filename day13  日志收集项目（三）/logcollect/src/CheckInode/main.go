package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"
	"time"
)

func main2() {
	fileName := flag.String("file", "", "filename")

	flag.Parse()

	if fileName == nil || *fileName == "" {
		panic("fileName is empty")
	}
	fmt.Printf("fileName is %v\n", *fileName)
	file, err := os.Open(*fileName)
	if err != nil {
		panic("file open err")
	}
	statOri, err := file.Stat()

	stat, ok := statOri.Sys().(*syscall.Stat_t)
	if !ok {
		fmt.Printf("not a syscall.Stat_t")
		return
	}
	fmt.Printf("open file stat = %#v\n ", stat)
	fmt.Printf("open file stat.Ino = %#v\n", stat.Ino)
	go sysStat(*fileName)
	readOffset(file)

}

func main() {
	fileName := flag.String("file", "", "filename")

	flag.Parse()

	if fileName == nil || *fileName == "" {
		panic("fileName is empty")
	}
	fmt.Printf("fileName is %v\n", *fileName)

	readFileEOF(*fileName)
}

func sysStat(filePath string) {
	for true {
		time.Sleep(time.Second)
		fileInfo, _ := os.Stat(filePath)

		stat, ok := fileInfo.Sys().(*syscall.Stat_t)
		if !ok {
			fmt.Printf("not a syscall.Stat_t")
		}
		fmt.Printf("syscall stat = %#v\n", stat)
		fmt.Printf("syscall stat.Ino = %#v\n", stat.Ino)
		fmt.Printf("syscall stat.Size = %#v\n", stat.Size)
	}
}

func readOffset(file *os.File) {
	go readFile(file)
	for true {
		offset, err := file.Seek(0, os.SEEK_CUR)

		if err != nil {
			fmt.Println("seek err: %v", err)
		}
		fmt.Printf("cur offset is %v\n", offset)
		time.Sleep(time.Second * 10)
	}

}

func readFile(file *os.File) {
	buffer := make([]byte, 10)
	for true {
		n, err := file.Read(buffer)
		if err != nil {
			fmt.Println("read err: %v", err)
		}
		fmt.Printf("read n:%v, b:%v\n", n, string(buffer))
		time.Sleep(time.Second)
	}

}

func readFileEOF(fileName string) {
	file, _ := os.Open(fileName)
	buffer := make([]byte, 512)
	fileInfo, _ := os.Stat(fileName)

	stat, _ := fileInfo.Sys().(*syscall.Stat_t)
	for true {
		n, err := file.Read(buffer)
		if err != nil {
			fmt.Println("read err: %v", err)

		} else {
			fmt.Printf("read n:%v, b:%v\n", n, string(buffer))
		}
		offset, _ := file.Seek(0, os.SEEK_CUR)
		fmt.Printf("curOffset: %v, size: %v\n", offset, stat.Size)
		time.Sleep(time.Second)
	}
}

func readFileEOF2(fileName string) {
	file, _ := os.Open(fileName)
	buffer := make([]byte, 512)

	fileInfo, _ := os.Stat(fileName)

	stat, _ := fileInfo.Sys().(*syscall.Stat_t)
	file.Seek(stat.Size-stat.Size+1, 0)
	for true {
		n, err := file.Read(buffer)
		if err != nil {
			fmt.Println("read err: %v", err)
		} else {
			fmt.Printf("read n:%v, b:%v\n", n, string(buffer))
		}
		time.Sleep(time.Second)
	}
}
