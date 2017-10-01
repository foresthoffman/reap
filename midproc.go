/**
 * midproc.go
 *
 * Copyright (c) 2017 Forest Hoffman. All Rights Reserved.
 * License: MIT License (see the included LICENSE file) or download at
 *     https://raw.githubusercontent.com/foresthoffman/midproc/master/LICENSE
 */

package midproc

import (
	"fmt"
	"os/exec"
	"bytes"
	"errors"
	"strconv"
	"strings"
)

func generateBgCmd(name string, arg ...string) string {

	// This command string, when run, will create a process to run the provided command (with its
	// arguments) in the background. Immediately after the process is started, the PID will be
	// collected and sent to standard output. Since this command is prefixed by the "nohup" command
	// and the stdin and stdout streams for the background-process have been directed to the bit-
	// bucket (/dev/null), only the PID will be output.
	cmdString := fmt.Sprintf(
		"nohup %s %s < /dev/null &>/dev/null & echo $! | awk '/[0-9]+$/{ print $0 }'",
		name,
		strings.Join(arg, " "),
	)

	return cmdString
}

func Run(name string, arg ...string) (uint32, error) {
	var out bytes.Buffer
	var stderr bytes.Buffer
	sleepCmd := exec.Command("/bin/bash", "-c", generateBgCmd(name, arg...))
	sleepCmd.Stdout = &out
	sleepCmd.Stderr = &stderr

	err := sleepCmd.Run()
	if nil != err {
		return 0, errors.New(fmt.Sprintf("%v: %v", err, stderr.String()))
	}

	// remove the newline character that gets appended to the PID via the "echo" command
	outStr := strings.Replace(out.String(), "\n", "", 1)

	i, err := strconv.ParseInt(outStr, 10, 64)
	if nil != err {
		return 0, err
	}
	pid := uint32(i)
	if 0 == pid {
		return 0, errors.New("Invalid PID of 0")
	}

	return pid, nil
}
