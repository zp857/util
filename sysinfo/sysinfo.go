package sysinfo

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"time"
)

func GetCpuPercent() float64 {
	percent, _ := cpu.Percent(time.Second, false)
	return percent[0]
}

func GetMemPercent() float64 {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.UsedPercent
}

func GetDiskPercent() float64 {
	usage, err := disk.Usage("/")
	if err != nil {
		fmt.Println("Failed to get disk usage:", err)
		return 0
	}

	percentage := float64(usage.Used) / float64(usage.Total) * 100
	return percentage
}
