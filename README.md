## Midproc

"midproc" provides functionality for running commands with optional arguments as background processes. This package allows Go programs to create detached processes which can live and exit independently of their parents. This package is intended to be utilized to create an intermediate process runner (*). Such a runner is used to create a 2nd-degree child (grandchild) process, then it is discarded, which results in a detached child process.

(*) See [midprocrunner](https://github.com/foresthoffman/midprocrunner) for a intermediate process runner.

### Installation

Run `go get github.com/foresthoffman/midproc`

### Importing

Import this package by including `github.com/foresthoffman/midproc` in your import block.

e.g.

```Go
package main

import(
	...
	"github.com/foresthoffman/midproc"
)
```

### Usage

Here's a simple example of the syntax:

```Go
package main

import(
    "fmt"
    "github.com/foresthoffman/midproc"
)

func main() {
    pid, err := midproc.Run("sleep", "30") // Creates a detached child process
                                           // that will run the "sleep" command
                                           // with a duration of 30 seconds,
                                           // after which it exits.
    if nil != err {
        panic(err)
    }
    fmt.Println(pid) // Outputs the process ID of the "sleep" command.
}
```

_That's all, enjoy!_
