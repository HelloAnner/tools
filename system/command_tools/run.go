package command_tools

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"
	"runtime"
)

// RunCommand 执行系统命令
func RunCommand(executableFile, workDir string, args ...string) error {
	var stdBuffer bytes.Buffer
	if runtime.GOOS != "windows" {
		executableFile = fmt.Sprintf("./%s", executableFile)
	}
	log.Printf("run %s %s on %s\n", executableFile, args, workDir)
	cmd := exec.Command(executableFile, args...)
	cmd.Dir = workDir
	mw := io.MultiWriter(log.Writer(), &stdBuffer)

	cmd.Stdout = mw
	cmd.Stderr = mw

	if err := cmd.Run(); err != nil {
		panic(err)
	}
	return nil
}
