package utils

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/viper"
)

func ExecFZF(context string) string {
	cmdOptions := viper.GetString("fzf_options")
	cmdOutput := &bytes.Buffer{}

	command := fmt.Sprintf("echo -e '%s' | fzf %s", context, cmdOptions)

	c := exec.Command("bash", "-c", command)
	c.Stdout = cmdOutput
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin

	err := c.Run()
	if err != nil && err.Error() != "exit status 130" {
		os.Exit(1)
	}

	return strings.Trim(cmdOutput.String(), "\n")
}

func Exec(command string, args ...string) ([]byte, error) {
	var stderr, stdout bytes.Buffer
	cmd := exec.Command(command, args...)

	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	if stderr.String() != "" {
		return nil, errors.New(stderr.String())
	}

	return stdout.Bytes(), nil
}

func ExecStdout(command string, args ...string) error {
	c := exec.Command(command, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin

	err := c.Run()
	if err != nil {
		return err
	}

	return nil
}
