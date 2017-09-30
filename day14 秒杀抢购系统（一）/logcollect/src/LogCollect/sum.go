package main

import (
	//"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"math/rand"
)

/*
* the nginx_sum_log file path is /home/work/logs/nginx/old/sum_log_20161130
* the app_sum_log file path is /home/work/applogs/old/sum_log_20161130
 */

const (
	NGX_FILE_PREFIX = "/home/work/logs/nginx/old/sum_log_"
	APP_FILE_PREFIX = "/home/work/logs/applogs/old/sum_log_"
)

func parseTime(t time.Time) string {
	return t.Format("20060102")
}

func RunSumLog() {
	for {
		now := time.Now()
		minutes := now.Minute()
		sleepTime := 10 - minutes
		if minutes > 10 {
			sleepTime = 70 - minutes
		}
		<-time.After(time.Duration(sleepTime) * time.Minute + time.Duration(rand.Intn(10)) * time.Second)
		sumLog2Kafka(now.Add(-time.Minute * 11))
		time.Sleep(time.Minute * 50)
	}
}

/*
func RunSumLog() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		now := time.Now()
		hours := now.Hour()
		sleepTime := 4 - hours
		if hours > 4 {
			sleepTime = 28 - hours
		}
		<-time.After(time.Duration(sleepTime)*time.Hour + time.Duration(rand.Intn(1800))*time.Second)
		sumLog2Kafka(now.Add(-time.Hour * 24))
		time.Sleep(time.Hour * 23)
	}
}*/

func sumLog2Kafka(t time.Time) {
	dateSuffix := parseTime(t)
	ngxFile := NGX_FILE_PREFIX + dateSuffix
	appFile := APP_FILE_PREFIX + dateSuffix
	//kafka can send 10000+ rows per time, so do not need to split data.
	data, err := readFile(ngxFile)
	if err != nil {
		taskLog.Fatal("SumLog2Kafka readFile failed, %v", err)
	} else {
		if len(data) > 0 {
			sendFile(data, "nginx")
			data = nil
		}
	}

	data, err = readFile(appFile)
	if err != nil {
		taskLog.Fatal("SumLog2Kafka readFile failed, %v", err)
	} else {
		if len(data) > 0 {
			sendFile(data, "app")
			data = nil
		}
	}

}

func readFile(filePath string) ([]byte, error) {

	_, err := os.Stat(filePath)
	if err != nil {
		//fmt.Printf("readFile, path:%v, err:%v\n", filePath, err)
		//return nil, fmt.Errorf("stat file failed, file: %v, err: %v", filePath, err)
		return []byte{}, nil
	}

	dataBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		//fmt.Printf("readFile, path:%v, err:%v\n", filePath, err)
		return nil, fmt.Errorf("readFile failed, file: %v, err: %v", filePath, err)
	}
	return dataBytes, nil
}

func sendFile(data []byte, sumType string) {
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	taskLog.Monitor(LogMap{"sum_type": sumType, "sum_log_len": len(data), "add_time": timeStr})
	taskLog.Monitor(LogMap{"sum_type": sumType, "sum_log": string(data), "add_time": timeStr})
}
