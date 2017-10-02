/**
 * doc.go
 *
 * Copyright (c) 2017 Forest Hoffman. All Rights Reserved.
 * License: MIT License (see the included LICENSE file) or download at
 *     https://raw.githubusercontent.com/foresthoffman/midproc/master/LICENSE
 */

/*
"midproc" provides functionality for running commands with optional arguments as background processes. This package allows Go programs to create detached processes which can live and exit independently of their parents. This package is intended to be utilized to create an intermediate process runner (*). Such a runner is used to create a 2nd-degree child (grandchild) process, then it is discarded, which results in a detached child process.

Here's an example:

    - Parent // the primary process that needs to spawn detached processes
      - Child A // one of potentially many child processes that need to be detached

In a situation such as above, sending a SIGKILL or SIGTERM signal to the Child A process will cause the child process to enter a zombie state. Meaning that it has been commanded to shutdown, and consumes no resources, but its reference in the process table remains. The zombie is only created when the Parent process has not yet exited. This becomes an issue if the Parent is any process where uptime is paramount, such as a server. It's a catch 22, shutting down the server will clean up the zombie processes, but it would also defeat the purpose of having the child processes in the first place. At that point, using goroutines would be more efficient.

In order for the Child processes to properly detach from the Parent, they must be orphaned. This means that they will have to belong to some process, which will exit and leave them without a Parent process. Then, the detached processes will only exit when they have either completed their task, or have been sent a termination signal.

Here's what the solution would look like:

    - Parent
      - Intermediate Parent // the intermediate process, whose job is to spawn
                            // an individual child process and then exit
        - Child A

When the Intermediate Parent exits, the Child will be orphaned resulting a process hierarchy like this:

    - Parent
    - Child A

Now, both the Parent and Child A can behave independently of each other. This way the Parent, a server for example, can be restarted without total loss of functionality from the Child process.

This is of course assuming that the host operating system is *nix based, as this package makes use of the "os/exec" package, which states the following in its documentation:

    "Note that the examples in [os/exec] assume a Unix system.
    They may not run on Windows,..." - in go1.9

(*) See https://github.com/foresthoffman/midprocrunner for a intermediate process runner.

Usage:

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
*/
package midproc
