package main

/// (c) 2015, Bernhard Tittelbach, xro@realraum.at

/// IDEA: wait for input on stdin
/// start a timer "-t <seconds>" and
/// run given command "$*" after no
/// more stdinput is received for x seconds

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

var delay_s time.Duration
var showhelp bool

func init() {
	flag.DurationVar(&delay_s, "delay", time.Second*10, "after which events have stopped coming, execute cmd. e.g.: 2m")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n%s [-d <duration>] <cmd>\nWill execute cmd after <duration> has elapsed without any new input on stdin\n\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(1)
	}
}

func runcmd(tC <-chan time.Time) {
	var procAttr os.ProcAttr
	procAttr.Files = []*os.File{nil, os.Stdout, os.Stderr}
	execpath, err := exec.LookPath(flag.Args()[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	args := make([]string, 1)
	args[0] = execpath
	args = append(args, flag.Args()[1:]...)
	for range tC {
		fmt.Printf("Running cmd: %s\n", strings.Join(args, " "))
		_, err := os.StartProcess(execpath, args, &procAttr)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Listening to stdin. Will run cmd after %.1fs without new input\n", delay_s.Seconds())
	timer := time.NewTimer(0)
	timer.Stop()
	buffer := make([]byte, 1)
	go runcmd(timer.C)
	for {
		_, err := reader.Read(buffer)
		if err != nil {
			break
		}
		timer.Reset(delay_s)
	}
}
