package shell

import (
	"github.com/docker/docker/pkg/broadcaster"
	"github.com/docker/docker/pkg/ioutils"
	"github.com/kr/pty"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

type Process struct {
	logs    *Logger
	Command *exec.Cmd
	pty     *os.File
	stdout  *broadcaster.Unbuffered
}

func StartProcess(cmd *exec.Cmd) (*Process, error) {
	f, err := pty.Start(cmd)
	if err != nil {
		return nil, err
	}
	log.Printf("Making process %d", cmd.Process.Pid)
	process := &Process{
		logs:    NewLogger(),
		Command: cmd,
		pty:     f,
		stdout:  new(broadcaster.Unbuffered),
	}
	process.stdout.Add(process.logs)
	go func() {
		io.Copy(process.stdout, process.pty)
		process.pty.Close()
	}()
	return process, nil
}

// stdoutPipe creates a new io.ReadCloser with an empty bytes pipe.
// It adds this new out pipe to the Stdout broadcaster.
// This will block stdout if unconsumed.
func (process *Process) stdoutPipe() io.ReadCloser {
	bytesPipe := ioutils.NewBytesPipe()
	process.stdout.Add(bytesPipe)
	return bytesPipe
}

func (process *Process) Attach(stdin io.ReadCloser, stdout io.Writer, logs bool) {
	go func() {
		io.Copy(process.pty, stdin)
		stdin.Close()
	}()

	var reader io.ReadCloser
	if logs {
		reader = ioutil.NopCloser(io.MultiReader(process.logs.Reader(), process.stdoutPipe()))
	} else {
		reader = process.stdoutPipe()
	}
	go func() {
		io.Copy(stdout, reader)
		reader.Close()
	}()
}
