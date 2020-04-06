package delta_server

import (
	"fmt"
	"lfcomlibGO/dapr"
	"lfcomlibGO/infra"
	"os"
	"time"
)

var (
	TimeGap      int
	DurationMin  int
	FileName     string
	DurationSecX int
)

func starter() {
	fmt.Println("DeltaGo Ver1.0")
}

func SetTimeGap() int {
	dapr.GetCmdArg()
	if len(dapr.GetCmdArg()) < 2 {
		println("Please input time gap(Seconds):")
		fmt.Scan(&TimeGap)
		return TimeGap
	} else {
		TimeGap := dapr.StringToInt(dapr.GetCmdArg()[1])
		return TimeGap
	}
}

func SetCsvFile() string {
	head := []string{"Time", "CPU", "Memory"}
	FileName = "Data_"
	FileName += dapr.TimeNowFormat_YYYY_MM_DD()
	FileName += "_"
	FileName += dapr.TimeStamp()
	FileName += ".csv"
	dapr.WriteCSV(FileName, head)
	return FileName
}

func SetDuration() int {
	dapr.GetCmdArg()
	if len(dapr.GetCmdArg()) < 3 {
		println("Please input time Duration(Minutes):")
		fmt.Scan(&DurationMin)
		return DurationMin * 60
	} else {
		DurationMin := dapr.StringToInt(dapr.GetCmdArg()[2])
		return DurationMin * 60
	}
}

func Run() {
	starter()
	TimeGap := SetTimeGap()
	FileName := SetCsvFile()
	DurationSec := SetDuration()
	for {
		CpuP := infra.GetCpuPercent()
		MemoP := infra.GetMemoInfo()
		//row := [][]string{{dapr.IntToString(CpuP), dapr.IntToString(MemoP)}}
		row := []string{dapr.TimeNowFormat_YYYY_MM_DD_HH_MM_SS(), dapr.IntToString(CpuP), dapr.IntToString(MemoP)}
		dapr.WriteCSV(FileName, row)
		fmt.Printf("Time:%s CPU:%s Memo:%s\n", row[0], row[1], row[2])
		time.Sleep(time.Duration(TimeGap-1) * time.Second)
		DurationSecX += TimeGap
		if DurationSecX >= DurationSec {
			os.Exit(0)
		}
	}
}
