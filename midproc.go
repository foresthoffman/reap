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
)

func generateBgCmd(name string, arg ...string) string {
	var argList string
	for _, val := range arg {
		argList += val + " "
	}

	cmdString := fmt.Sprintf(
		"nohup %s %s < /dev/null &>/dev/null & echo $! | awk '/[0-9]+$/{ print $0 }'",
		name,
		argList,
	)

	return cmdString
}

func Run(name string, arg ...string) {
	sleepCmd := exec.Command(generateBgCmd(name, arg...))
	err := sleepCmd.Start()
	if nil != err {
		fmt.Printf("midproc process failed: %v\n", err)
	} else {
		fmt.Println("success!")
	}

	oBytes, err := sleepCmd.Output()
	if nil != err {
		fmt.Printf("midproc process couldn't collect output: %v\n", err)
	} else {
		fmt.Printf("midproc output: %s\n", oBytes)
	}
}
