package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"strconv"
)

func GetNumsFromString(strRaw string)string{
	iBegin := 0
	iEnd := 0
	bIsInNum := false

	if len(strRaw) < 0{
		return ""
	}
	if strRaw[0] > '9' || strRaw[0] < '0' {
		return ""
	}	
	
	for i, iChar := range strRaw{
		if iChar >= '0' && iChar <= '9'{
			if bIsInNum == false{
				iBegin = i
				bIsInNum = true
			}else{
				iEnd = i + 1
			}
		}else
		{
			if bIsInNum == true{
				bIsInNum = false
				break
			}
		}
	}
	if iEnd >= len(strRaw){
		return strRaw[iBegin:]
	}
	return strRaw[iBegin:iEnd]
}

func main() {
	if len(os.Args) <= 1{
		fmt.Printf("Usage:./%s FileName1 FileNme2 ...\n", os.Args[0])
		return
	}
	fmt.Println(os.Args)	
	fileList := os.Args[1:]
	for _, strFileName := range fileList {
		_, err := os.Stat(strFileName)
		if err != nil{
			fmt.Println(strFileName, err.Error())
			return 
		}
	}
	mapIntBool := make(map[int64]bool)
	mapInvalidInfo := make(map[string]int64)
	for _, strFileName := range fileList {
		file, err := os.Open(strFileName)
		if err != nil{
			fmt.Println(strFileName, err.Error())
			return 
		}
		defer file.Close()
		buf := bufio.NewReader(file)
    for {
        strLine, errR := buf.ReadBytes('\n')
        if errR != nil {
            if errR == io.EOF {
                break
            }
            fmt.Println(strFileName, errR.Error())
        }
				strNum := GetNumsFromString(string(strLine))
				if strNum == ""{
					continue
				}
				uId,err := strconv.ParseInt(strNum, 10, 64)
				if err != nil{
					mapInvalidInfo[strFileName] ++
					continue;
				}
				mapIntBool[uId] = true
		}
		fmt.Println("SetInMap", strFileName)
	}
	
	fmt.Printf("TotalSize %#v\n", len(mapIntBool))
	for key, uValue := range mapInvalidInfo{
		fmt.Println("InvalideSize:", key, uValue)
	}
	for key, _:= range mapIntBool{
		fmt.Println(key)
	}
	fmt.Println("OK")
}
