package infra

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"time"
)

func GetMemoInfo() int {
	v, _ := mem.VirtualMemory()

	//fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
	// convert to JSON. String() is also implemented
	return int(v.UsedPercent)
}

// cpu info
func GetCpuInfo() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err:%v", err)
	}
	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}
}

func GetHostName() string {
	n, _ := host.Info()
	return n.Hostname
}

func GetCpuPercent() int {
	percent, _ := cpu.Percent(time.Second, false)
	// fmt.Println("type:", reflect.TypeOf(percent))
	var sum float64 = 0
	for _, num := range percent {
		sum += num
	}
	//fmt.Printf("cpu percent:%v\n", percent)
	return int(sum)
}
