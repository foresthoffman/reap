#!/bin/bash
#
# midproc.sh
# 
# Runs the specified command with any number of arguments as a background process and prints the
# process ID (PID) of the background-process to standard output.
#
# Copyright (c) Forest Hoffman, 2017.
# All Rights Reserved.
##

# Holds the initial command.
cmd="${1}"

# Will hold a concatenated string of arguments for the provided command.
args=""

# Holds the initial number of arguments.
argLen=$#

# Loops through all arguments but the first and concatenates them together.
while [ $# -gt 0 ]; do
	if [ $# -lt $argLen ]; then
		args="${args} ${1}"
	fi

	# Moves the next argument's value into the place of the first (e.g. $1).
	shift
done

# Runs the provided command as a background process and filters the PID of the process in the
# output string.
nohup $cmd $args < /dev/null &>/dev/null &

# Collects the PID from the background-process. The PID is output by the "&" character after the
# command above.
echo $! | awk '/[0-9]+$/{ print $0 }'
