package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime/pprof"
	"strings"
	"syscall"
	"util"

	"oldboy/logcollect/src/kafka"

	"oldboy/logcollect/src/zklib"

	"oldboy/xrsmonitor"
	"oldboy/xstatsd"

	"github.com/Unknwon/goconfig"
)

var (
	zkservers    []string
	kafkaServers []string
	taskLog      *Logger
)

func signal_handler() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGTERM)
	var file *os.File
	for {
		sig := <-c
		switch sig {
		case syscall.SIGUSR1:
			mem, _ := os.Create("profile_mem.dat")
			pprof.WriteHeapProfile(mem)
			if mem != nil {
				mem.Close()
			}
			file, _ := os.Create("profile_cpu.dat")
			pprof.StartCPUProfile(file)
		case syscall.SIGUSR2:
			pprof.StopCPUProfile()
			if file != nil {
				file.Close()
				file = nil
			}
		case syscall.SIGTERM:
			NodeShutDown()
		}
	}
}
func init() {
	go sampleProc(int32(os.Getpid()))
	go sampleTotalCpu()
	go resourceLimit()
	go signal_handler()
	confile := flag.String("conf", "config.ini", "agent config file")
	var value string
	flag.Parse()
	fmt.Println(*confile)
	cfg, err := goconfig.LoadConfigFile(*confile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	value, err = cfg.GetValue("kafka", "servers")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	kafkaServers = strings.Split(value, ",")

	value, err = cfg.GetValue("zookeeper", "servers")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	zkservers = strings.Split(value, ",")

	timeout, err := cfg.Int("zookeeper", "timeout")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	addr, err := cfg.GetValue("statsd", "addr")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	group_name, err := cfg.GetValue("statsd", "group_name")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	service_name, err := cfg.GetValue("statsd", "service_name")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	environ, err := cfg.GetValue("statsd", "environ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	testIpPrefix, err := cfg.GetValue("test", "ip_prefix")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	localIP := util.GetLocalIP()
	taskLog = NewLogger(localIP)
	if isTestEnv := util.IsIpPrefix(localIP, testIpPrefix, ";"); isTestEnv {
		fmt.Printf("environ[test]\n")
		environ, err = cfg.GetValue("test", "environ")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		value, err = cfg.GetValue("test", "kafka_servers")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		kafkaServers = strings.Split(value, ",")

		value, err = cfg.GetValue("test", "zk_servers")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		zkservers = strings.Split(value, ",")

		addr, err = cfg.GetValue("test", "statsd_servers")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	fmt.Printf("%v %v %v %v", addr, group_name, service_name, environ)
	xstatsd.Init(addr, group_name, service_name, environ)
	zklib.Connect(zkservers, timeout, nil)
	if err = kafka.Connect(kafkaServers); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	xrsmonitor.Init()
}

func main() {
	Gao()
}
