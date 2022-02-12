package acme

import (
	"os/exec"
	"io"
	"bufio"
)



func SendCommand(comm Command, Payload string) error {
	cmd := exec.Command(comm.BaseCommand, string(comm.Rwl), comm.Path)
	pipe, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	go func() {
		defer pipe.Close()
		io.WriteString(pipe, Payload)
	}()
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func ReadCommand(comm Command) (r *bufio.Reader, e error) {
	cmd := exec.Command(comm.BaseCommand, string(comm.Rwl), comm.Path)
	cmdStdOut, err := cmd.StdoutPipe()
	if err != nil {
		return r, err
	}
	cmd.Start()
	r = bufio.NewReader(cmdStdOut)
	return r, nil
}