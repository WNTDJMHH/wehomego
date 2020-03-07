package main

import (
	"fmt"
	"flag"
	"os"
	"bufio"
	"io"
	"strconv"
)

var strFileA string
var strFileB string
var strFunc string

func main() {
	flag.StringVar(&strFunc, "f", "comm", "comm/diff/merge")
	flag.StringVar(&strFileA, "i", "", "FileNameA")
	flag.StringVar(&strFileB, "j", "", "FileNameB")
	flag.Parse();
	if strFileA == "" || strFileB == ""{
		fmt.Println("Fatal FileNameA or FileNameB NULL PASS")
		flag.PrintDefaults()
		return
	}
	fmt.Printf("Func:%v,FileA:%v,FileB:%v\n", strFunc, strFileA, strFileB);
	fileA, err := os.Open(strFileA)
	if err != nil{
		fmt.Println(strFileA, err.Error())
		return 
	}
	defer fileA.Close()
		
	fileB, err := os.Open(strFileB)
	if err != nil{
		fmt.Println(strFileB, err.Error())
		return 
	}
	defer fileB.Close()

	mapIntBool := make(map[int64]bool)
	{
		buf := bufio.NewReader(fileB)
    for {
        strLine, errR := buf.ReadBytes('\n')
        if errR != nil {
            if errR == io.EOF {
                break
            }
            fmt.Println(strFileA, errR.Error())
        }
				uId,err := strconv.ParseInt(string(strLine[:len(strLine) - 1]), 10, 64)
				if err != nil{
					fmt.Println(strFileB, "|Invalid|", err.Error())
					continue;
				}
				mapIntBool[uId] = true
		}
	}	
	
	{
		buf := bufio.NewReader(fileA)
    for {
        strLine, errR := buf.ReadBytes('\n')
        if errR != nil {
            if errR == io.EOF {
                break
            }
            fmt.Println(strFileB, errR.Error())
        }
				uId,err := strconv.ParseInt(string(strLine[:len(strLine) - 1]), 10, 64)
				if err != nil{
					fmt.Println(strFileA, "|Invalid|", err.Error())
					continue;
				}

				_,bFound := mapIntBool[uId]
				switch strFunc{
					case "diff":
						if  bFound == false{
							fmt.Println(uId)
						}
					case "merge":
							mapIntBool[uId] = true
					default:
							if bFound == true{
								fmt.Println(uId)
							};
				}
		}
	}
	if strFunc == "merge"{
		for key, _:= range mapIntBool{
			fmt.Println(key)
		}
	}	
	fmt.Println("OK")
}
