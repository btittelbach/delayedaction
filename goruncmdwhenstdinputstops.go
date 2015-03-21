package main

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

var delay_s int64
var showhelp bool

func init() {
	flag.Int64Var(&delay_s, "delay", 10, "in seconds after which events have stopped coming, execute cmd")
	flag.BoolVar(&showhelp, "help", false, "show help")
	flag.Parse()
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
	fmt.Printf("Listening to stdin. Will run cmd after %ds without new input\n", delay_s)
	timer := time.NewTimer(0)
	timer.Stop()
	buffer := make([]byte, 1)
	go runcmd(timer.C)
	for {
		_, err := reader.Read(buffer)
		if err != nil {
			break
		}
		timer.Reset(time.Second * time.Duration(delay_s))
	}
}
