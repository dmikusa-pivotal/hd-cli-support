package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type TicketEntry struct {
	Description string
	Body        []byte
}

type TicketLog struct {
	Name   string
	Writer io.Writer
}

func NewTicketLog() (tl TicketLog) {
	file, _ := ioutil.TempFile(os.TempDir(), "cf-support-")
	tl.Name = file.Name()
	tl.Writer = file
	return
}

func (tl TicketLog) Append(te TicketEntry) {
	io.WriteString(tl.Writer, fmt.Sprintf("## %s\n", te.Description))
	io.WriteString(tl.Writer, "```\n")
	tl.Writer.Write(te.Body)
	io.WriteString(tl.Writer, "```\n")
}
