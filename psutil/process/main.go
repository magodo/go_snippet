package main

import (
	"fmt"

	"github.com/shirou/gopsutil/process"
)

func main() {
	procs, _ := process.Processes()
	for _, proc := range procs {
		cmd, _ := proc.Cmdline()
		fmt.Printf("%10d -> %10s\n", proc.Pid, cmd)
	}
}
