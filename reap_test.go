/**
 * reap_test.go
 *
 * Copyright (c) 2017 Forest Hoffman. All Rights Reserved.
 * License: MIT License (see the included LICENSE file) or download at
 *     https://raw.githubusercontent.com/foresthoffman/reap/master/LICENSE
 */

package reap

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"
)

// Ensures that a new background processe is created.
func TestExec(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := "sleep"
	arg := "5"

	pid, err := Exec(cmd, arg)
	if nil != err {
		t.Fatalf("Failed to run \"%s %s\", %s", cmd, arg, err.Error())
	}
	cmdStr := fmt.Sprintf("ps aux | awk '/%d.*0:00 %s %s/{ print $2,$11,$12 }'", pid, cmd, arg)

	// Checks for a process with the specified name and PID.
	psCmd := exec.Command("/bin/bash", "-c", cmdStr)
	psCmd.Stdout = &stdout
	psCmd.Stderr = &stderr
	err = psCmd.Run()
	if nil != err {
		t.Fatalf("Failed to get status, %s, %s", err.Error(), stderr.String())
	}
	if "" == stdout.String() {
		t.Fatal("Failed to get status, process not found.")
	}
}
