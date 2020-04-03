package dapr

import (
	"bufio"
	"fmt"
	"os"
)

func CmdPause(msg string) {
	fmt.Print(msg)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func GetCmdArg() []string {
	var args_v []string
	for _, args := range os.Args {
		//fmt.Println("参数" + strconv.Itoa(idx) + ":", args)
		args_v = append(args_v, args)
	}
	return args_v
}
