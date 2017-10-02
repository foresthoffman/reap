/**
 * midproc_test.go
 *
 * Copyright (c) 2017 Forest Hoffman. All Rights Reserved.
 * License: MIT License (see the included LICENSE file) or download at
 *     https://raw.githubusercontent.com/foresthoffman/midproc/master/LICENSE
 */

package midproc

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"
)

// Ensures that the Run() function creates new background processes.
func TestRun(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	var pid string
	cmd := "sleep"
	arg := "5"

	pid, err := Run(cmd, arg)
	cmdStr := fmt.Sprintf(
		"ps aux | awk '/[^0-9]+(%s)[^\\n]*0:00 (%s %s)$/{ print $2,$11,$12 }'",
		pid,
		cmd,
		arg,
	)

	// checks for a process with the specified name and PID
	psCmd := exec.Command(
		"/bin/bash",
		"-c",
		cmdStr,
	)
	psCmd.Stdout = &stdout
	psCmd.Stderr = &stderr
	err = psCmd.Run()
	if nil != err {
		t.Errorf("Failed to check for process status. err: %v, %s", err, stderr.String())
	} else if "" == stdout.String() {
		t.Error("Failed to check for process status. Process was not found.")
	}
}
