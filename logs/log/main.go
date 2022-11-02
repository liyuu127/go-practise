package main

import (
	"log"
	"os"
)

// std
// Ldate：当前时区的日期，格式是：2009/01/23。
// Ltime：当前时区的时间，格式是：01:23:23，精确到秒。
// Lmicroseconds：当前时区的时间，格式是：01:23:23.862600，精确到微妙。
// Llongfile：全文件名和行号。
// Lshortfile：当前文件名和行号，会覆盖Llongfile。
// LUTC：使用UTC而非本地时区。
// Lmsgprefix：将“前缀”从行的开头移至消息之前。
// LstdFlags：标准Logger的默认值（Ldate、Ltime）。
var logger = log.New(os.Stderr, "[Debug] ", log.Llongfile)

func main() {
	logFile, err := os.Create("./log.log")
	logger.SetOutput(logFile)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("create file log.log failed")
	}
	logger.SetOutput(os.Stdout)
	logger.Print("call Print: line1")
	logger.Println("call Println: line2")

	// 修改日志配置
	logger.SetPrefix("[Info] ")
	logger.SetFlags(log.Ldate)
	logger.SetOutput(os.Stdout)
	logger.Print("Info check stdout")
}
