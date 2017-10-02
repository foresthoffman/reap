/**
 * midproc.go
 *
 * Copyright (c) 2017 Forest Hoffman. All Rights Reserved.
 * License: MIT License (see the included LICENSE file) or download at
 *     https://raw.githubusercontent.com/foresthoffman/midproc/master/LICENSE
 */

package midproc

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

// Generates a command that will create a process to run the provided command (with its
// arguments) in the background. Immediately after the process is started, the PID will be
// collected and sent to standard output. Since this command is prefixed by the "nohup" command
// and the Stdin and Stdout streams for the background-process have been directed to the bit-
// bucket (/dev/null), only the PID will be output.
func generateBgCmd(name string, arg ...string) string {
	cmdString := fmt.Sprintf(
		"nohup %s %s < /dev/null &>/dev/null & echo $! | awk '/[0-9]+$/{ print $0 }'",
		name,
		strings.Join(arg, " "),
	)

	return cmdString
}

// Runs a specified command with the provided arguments. The child process that is generated will
// be put in the background and its PID will be returned. An invalid PID or a script-error will
// result in a non-nil error being returned.
//
// In the case of a script-error, the error output to Stderror will be captured and returned along
// with the error from the exec.Cmd.Run() call.
func Run(name string, arg ...string) (string, error) {
	var out bytes.Buffer
	var stderr bytes.Buffer
	sleepCmd := exec.Command("/bin/bash", "-c", generateBgCmd(name, arg...))
	sleepCmd.Stdout = &out
	sleepCmd.Stderr = &stderr

	err := sleepCmd.Run()
	if nil != err {
		return "0", errors.New(fmt.Sprintf("%v: %v", err, stderr.String()))
	}

	// remove the newline character that gets appended to the PID via the "echo" command
	outStr := strings.Replace(out.String(), "\n", "", 1)
	if "0" == outStr {
		return "0", errors.New("Invalid PID of 0")
	}

	return outStr, nil
}
