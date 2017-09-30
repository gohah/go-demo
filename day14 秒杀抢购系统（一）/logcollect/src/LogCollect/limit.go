package main

import (
	"runtime"
	"time"

	"oldboy/src/kafka"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/process"
)

var (
	totalCpuUsage float32
	procCpuUsage  float32
	sampleTime    int32 = 1000 //ms
	limitType     int   = SLIENCE_TYPE
)

const (
	SMART_TYPE = iota
	FAST_TYPE
	SLIENCE_TYPE
)

func SetSampleTime(ms int32) {
	if ms >= 200 && ms <= 5000 {
		sampleTime = ms
	}
}

func sampleProc(pid int32) {
	proc, err := process.NewProcess(pid)
	if err != nil {
		return
	}
	pcpu, _ := proc.Times()
	cost := pcpu.User + pcpu.System
	for {
		time.Sleep(time.Duration(sampleTime) * time.Millisecond)
		pcpu, _ = proc.Times()
		procCpuUsage = float32(pcpu.User+pcpu.System-cost) / (float32(sampleTime) / 1000.0)
		cost = pcpu.User + pcpu.System
	}
}
func sampleTotalCpu() {
	cpu_num := int32(runtime.NumCPU())
	tcpu, _ := cpu.Times(false)
	cost := tcpu[0].User + tcpu[0].System
	for {
		time.Sleep(time.Duration(sampleTime) * time.Millisecond)
		tcpu, _ := cpu.Times(false)
		totalCpuUsage = (float32(tcpu[0].User + tcpu[0].System - cost)) / (float32(sampleTime*cpu_num) / 1000)
		cost = tcpu[0].User + tcpu[0].System
	}
}
func call_smart() {
	if totalCpuUsage < 0.7 {
		if kafka.IsFull() {
			kafka.LimitIncr(1000)
			taskLog.Notice("smart, total cpu is low:%f queue is full,count:%d, now:%d", totalCpuUsage, kafka.GetCount(), kafka.GetLimit())
		}
	} else {
		if kafka.IsFull() == false {
			kafka.LimitDecr(1000)
			taskLog.Notice("smart, total cpu is high:%f queue is not full,count:%d, now:%d", totalCpuUsage, kafka.GetCount(), kafka.GetLimit())
		} else if procCpuUsage > 0.5 {
			kafka.LimitDecr(1000)
			taskLog.Notice("smart, total cpu is high:%f queue is full, proc_cpu:%f,count:%d, now:%d", totalCpuUsage, procCpuUsage, kafka.GetCount(), kafka.GetLimit())
		}
	}
}
func call_fast() {
	if kafka.IsFull() {
		kafka.LimitIncr(1000)
		taskLog.Notice("kafka queue is full,count:%d incr limit, now:%d", kafka.GetCount(), kafka.GetLimit())
	}
}

func call_slience() {
	if procCpuUsage > 0.2 {
		kafka.LimitDecr(1000)
		taskLog.Notice("proc_cpu:%f decr limit, now:%d", procCpuUsage, kafka.GetLimit())
	}
}
func resourceLimit() {
	for {
		switch limitType {
		case SMART_TYPE:
			call_smart()
		case FAST_TYPE:
			call_fast()
		case SLIENCE_TYPE:
			call_slience()
		}
		time.Sleep(2 * time.Second)
	}
}
