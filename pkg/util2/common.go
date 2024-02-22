package util2

import (
	"fmt"
	"io"
	"log"
	"ops_client/pkg/util"
	"os"
	"strings"
	"time"
)

func GetRootPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// return dir
	return strings.Replace(dir, "\\", "/", -1)
}

func CommonLog(service string, msg string) bool {
	var dirPath, file string
	if service == "" {
		dirPath = GetRootPath() + "/logs" + "/common"
		file = dirPath + "/" + "common" + time.Now().Format("01") + ".log"
	} else {
		dirPath = GetRootPath() + "/logs/" + service
		file = dirPath + "/" + service + time.Now().Format("01") + ".log"
	}

	if !util.IsDir(dirPath) {
		if err := os.Mkdir(dirPath, 0775); err != nil {
			tmpBool := tmpLogWrite(time.Now().Local().Format("2006-01-02 15:04:05") + "mkdir failed！ " + err.Error())
			if !tmpBool {
				panic(fmt.Errorf("临时日志文件写入失败"))
			}
		}
	}
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		tmpBool := tmpLogWrite(time.Now().Local().Format("2006-01-02 15:04:05") + "打开日志文件失败！ " + err.Error())
		if !tmpBool {
			panic(fmt.Errorf("临时日志文件写入失败"))
		}
	}
	log.SetOutput(logFile)
	log.SetPrefix("[" + service + "]" + "[" + time.Now().Local().Format("2006-01-02 15:04:05") + "] ")
	log.Println(msg)
	return true
}

func tmpLogWrite(msg string) bool {
	filePath := GetRootPath() + "/logs/tmp.log"

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("临时文件打开失败")
		return false
	}

	defer file.Close()
	// 创建一个写入器用作追加
	writer := io.MultiWriter(file)
	if _, err := io.WriteString(writer, msg+"\n"); err != nil {
		return false
	}
	return true
}
