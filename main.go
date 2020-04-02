package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

// cpu info
func getCpuInfo() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err:%v", err)
	}
	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}
	// CPU使用率
	percent, _ := cpu.Percent(time.Second, true)
	// fmt.Println("type:", reflect.TypeOf(percent))
	println(percent)
	var sum float64 = 0
	for _, num := range percent {
		sum += num
	}
	fmt.Println("sum:", int(sum))

	fmt.Printf("cpu percent:%v\n", percent)
}

func main() {
	v, _ := mem.VirtualMemory()
	getCpuInfo()
	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
	// convert to JSON. String() is also implemented
	fmt.Println(v)
}
