package tailf

import (
	_ "fmt"
)

//php slow log can be splited by space_line
//[xxxx-xxxx-xxxx xx:xx:xx regard as a valid datetime
func phpTimeFormat(log []byte) bool {
	stage := 0
	partNum := 0
	partNumLimit := 4
	size := len(log)
	if size <= 12 { //[s0-s1-s2 s3:s4:s5] --> s5
		return false
	}
	if log[0] != '[' {
		return false
	}
	i := 1
loop:
	for i < size {
		//fmt.Printf("stage :%v, char:%v \n", stage, string(log[i]))
		switch log[i] {
		case '-':
			if stage >= 2 {
				return false
			}
			stage++
			partNum = 0
		case ':':
			if stage >= 5 || stage < 2 {
				return false
			}
			stage++
			partNum = 0
		case ' ':
			if stage == 5 {
				break loop
			}
			if stage > 2 {
				return false
			}
			stage++
			partNumLimit = 2
			partNum = 0
		case ']':
			break loop
		default:
			partNum++
			if stage != 5 && partNum > partNumLimit {
				return false
			}
		}
		i++
	}
	if stage == 5 && i > 12 {
		return true
	}
	return false
}
