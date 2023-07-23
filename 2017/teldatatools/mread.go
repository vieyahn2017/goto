package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/axgle/mahonia"
	"io/ioutil"
	//"strconv"
	"strings"
)

//读取WorkACSV文件数据
func readWorkA(file string) (workas WorkAS, err error) {
	//	workas = make([]WorkA, 0)
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("(ErrorCode:0001) error is " + err.Error())
		return
	}
	decode := mahonia.NewDecoder("gbk")
	if decode == nil {
		fmt.Println("(ErrorCode:0002) error is " + err.Error())
		return workas, errors.New("tmahonia.NewDecoder")
	}
	rcsv := csv.NewReader(strings.NewReader(decode.ConvertString(string(buf))))
	//允许双引号字符，否则报错 (ErrorCode:0003) error is line 29752, column 15: extraneous " in field
	rcsv.LazyQuotes = true
	//检查字段个数
	rcsv.FieldsPerRecord = 21

	ret, err := rcsv.ReadAll()
	if err != nil {
		fmt.Println("(ErrorCode:0003) error is " + err.Error())
		return
	}

	for i, v := range ret {
		if i == 0 {
			continue
		}
		var worka WorkA
		worka.WID = v[1]
		worka.WName = v[2]
		worka.OID = v[3]
		worka.Org = v[4]
		worka.PID = v[5]
		worka.Partment = v[6]
		worka.OrgName = v[7]
		worka.WIDClass = v[8]
		worka.ChanClass = v[9]
		worka.WIDClassB = v[10]
		worka.AreaID = v[11]
		worka.Area = v[12]
		worka.WIDStatus = v[13]
		worka.WCode = v[14]
		worka.WNameA = v[15]
		worka.WJob = v[16]
		worka.WSource = v[17]
		worka.WType = v[18]
		worka.WIsAdd = v[19]
		worka.WStatus = v[20]
		workas = append(workas, worka)
	}
	return
}

//读取WorkBCSV文件数据
func readWorkB(file string) (workbs WorkBS, err error) {
	//workbs = make([]WorkB, 0)
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("(ErrorCode:0001) error is " + err.Error())
		return
	}
	decode := mahonia.NewDecoder("gbk")
	if decode == nil {
		fmt.Println("(ErrorCode:0002) error is " + err.Error())
		return workbs, errors.New("tmahonia.NewDecoder")
	}
	rcsv := csv.NewReader(strings.NewReader(decode.ConvertString(string(buf))))
	//允许双引号字符，否则报错 (ErrorCode:0003) error is line 29752, column 15: extraneous " in field
	rcsv.LazyQuotes = true
	//检查字段个数
	rcsv.FieldsPerRecord = 13

	ret, err := rcsv.ReadAll()
	if err != nil {
		fmt.Println("(ErrorCode:0003) error is " + err.Error())
		return
	}

	for i, v := range ret {
		if i == 0 {
			continue
		}
		var workb WorkB
		workb.WCode = v[1]
		workb.WTreeID = v[2]
		workb.WTreeOrg = v[3]
		workb.WNameB = v[4]
		workb.WCard = v[5]
		workb.WPhone = v[6]
		workb.WEmail = v[7]
		workb.WJobB = v[8]
		workb.WSourceB = v[9]
		workb.WTypeB = v[10]
		workb.WIsAddB = v[11]
		workb.WStatusB = v[12]
		workbs = append(workbs, workb)
	}
	return
}
