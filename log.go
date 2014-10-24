package main

import (
	"bytes"
	"fmt"
	"github.com/daviddengcn/go-colortext"
	"sync"
	"time"
)

type clogger struct {
	idx  int
	proc string
}

var colors = []ct.Color{
	ct.Green,
	ct.Cyan,
	ct.Magenta,
	ct.Yellow,
	ct.Blue,
	ct.Red,
}
var ci int

var mutex = new(sync.Mutex)

// write handler of logger.
func (l *clogger) Write(p []byte) (int, error) {
	buf := bytes.NewBuffer(p)
	wrote := 0
	for {
		line, err := buf.ReadBytes('\n')
		if len(line) > 0 {
			if line[len(line)-1] != '\n' {
				line = append(line, '\n')
			}
			now := time.Now().Format("15:04:05")
			format := fmt.Sprintf("%%s %%%ds | ", maxProcNameLength)
			s := string(line)

			mutex.Lock()
			ct.ChangeColor(colors[l.idx], false, ct.None, false)
			fmt.Printf(format, now, l.proc)
			ct.ResetColor()
			fmt.Print(s)
			mutex.Unlock()

			wrote += len(line)
		}
		if err != nil {
			break
		}
	}
	return wrote, nil
}

// create logger instance.
func createLogger(proc string) *clogger {
	l := &clogger{ci, proc}
	ci++
	if ci >= len(colors) {
		ci = 0
	}
	return l
}
