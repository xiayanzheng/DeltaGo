package dapr

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func WriteCSV(filePath string, data []string) bool {
	//创建一个新文件
	newFileName := filePath
	//这样打开，每次都会清空文件内容
	//nfs, err := os.Create(newFileName)

	//这样可以追加写
	nfs, err := os.OpenFile(newFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("can not create file, err is %+v", err)
	}
	defer nfs.Close()
	nfs.Seek(0, io.SeekEnd)
	w := csv.NewWriter(nfs)
	//设置属性
	w.Comma = ','
	w.UseCRLF = true
	row := data
	err = w.Write(row)
	if err != nil {
		log.Fatalf("can not write, err is %+v", err)
		return false
	}
	//这里必须刷新，才能将数据写入文件。
	w.Flush()
	return true
}
