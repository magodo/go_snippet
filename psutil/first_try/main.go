package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

var args struct {
	path string
}

func init() {
	flag.StringVar(&args.path, "p", "/tmp", "path of some mount point")
	flag.Parse()
}

type HostCapacity struct {
	disk_total_space         uint64
	disk_free_space          uint64
	io_read_byte_per_second  uint64
	io_write_byte_per_second uint64
	cpu_idle                 float64
	mem_available            uint64
}

const cumulateSecond int = 1

func CollectHostCapacity(path string) (capacity HostCapacity, err error) {

	/* Get disk statistics */
	usageStat, err := disk.Usage(path)
	if err != nil {
		err = errors.Wrapf(err, "disk.Usage() on %s failed", path)
		return
	}
	capacity.disk_total_space = usageStat.Total
	capacity.disk_free_space = usageStat.Free

	/* Get memory statistics */
	vmemStatPtr, err := mem.VirtualMemory()
	if err != nil {
		err = errors.Wrap(err, "mem.VirtualMemory() failed")
		return
	}
	capacity.mem_available = vmemStatPtr.Available
	fmt.Println(vmemStatPtr)

	/* Get I/O statistics */
	partition, err := getDeviceFromPath(path)
	if err != nil {
		err = errors.Wrapf(err, "findPartitionOfPath() on %s failed", path)
		return
	}

	ioCounterStatMap, err := disk.IOCounters(partition)
	if err != nil {
		err = errors.Wrapf(err, "disk.IOCounters() on %s failed", partition)
		return
	}
	var iocounterStart disk.IOCountersStat
	for _, iocounterStart = range ioCounterStatMap {
		// we just want to get the only element in map
		break
	}

	// sleep for cumulate
	time.Sleep(time.Duration(cumulateSecond) * time.Second)

	ioCounterStatMap, err = disk.IOCounters(partition)
	if err != nil {
		err = errors.Wrapf(err, "disk.IOCounters() on %s failed", partition)
		return
	}
	var iocounterEnd disk.IOCountersStat
	for _, iocounterEnd = range ioCounterStatMap {
		// we just want to get the only element in map
		break
	}

	capacity.io_read_byte_per_second = (iocounterEnd.ReadBytes - iocounterStart.ReadBytes) / uint64(cumulateSecond)
	capacity.io_write_byte_per_second = (iocounterEnd.WriteBytes - iocounterStart.WriteBytes) / uint64(cumulateSecond)

	/* Get cpu statistics */
	cpuTimeStats, err := cpu.Times(false)
	if err != nil {
		err = errors.Wrap(err, "cpu.Times() failed")
		return
	}
	cpustat := cpuTimeStats[0]
	capacity.cpu_idle = cpustat.Idle / cpustat.Total()
	fmt.Println(cpustat, cpustat.Total())

	return
}

//getDeviceFromPath find the partition (e.g. /dev/sda1 if its mount point contains the `path`)
func getDeviceFromPath(path string) (partition string, err error) {
	partitionStats, err := disk.Partitions(true)
	if err != nil {
		err = errors.Wrap(err, "disk.Partitions() failed")
		return
	}

	mntPointLen := 4096 // PATH_MAX on linux
	for _, pstat := range partitionStats {
		if strings.HasPrefix(path, pstat.Mountpoint) {
			if len(pstat.Mountpoint) < mntPointLen {
				// NOTE: the "most-matched" mount point always has less redundent suffix path
				partition = pstat.Device
				mntPointLen = len(pstat.Mountpoint)
			}
		}
	}
	if partition == "" {
		err = fmt.Errorf("No partition found containing path: %s", path)
	}
	return
}

func main() {
	capacity, err := CollectHostCapacity(args.path)
	if err != nil {
		fmt.Println(err)
		return
	}
	spew.Dump(capacity)
}
