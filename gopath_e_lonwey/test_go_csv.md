"encoding/csv"


```go
type WorkA struct {
	WID       string
	WName     string
	OID       string
	Org       string
	PID       string
	Partment  string
	OrgName   string
	WIDClass  string
	ChanClass string
	WIDClassB string
	AreaID    string
	Area      string
	WIDStatus string
	WCode     string
	WNameA    string
	WJob      string
	WSource   string
	WType     string
	WIsAdd    string
	WStatus   string
}



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



//输出C表（合并出来的表）
func writeWorkC(workcs WorkCS, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Printf("Fail to Open file %v\n", outfile)
		return err
	}

	defer file.Close()
	encode := mahonia.NewEncoder("gbk")
	if encode == nil {
		return errors.New("tmahonia.NewEncoder error")
	}
	writer := csv.NewWriter(file)

	recordVodTHead := make([]string, 31)

	recordVodTHead[0] = encode.ConvertString("销售员编码")
	recordVodTHead[1] = encode.ConvertString("营销工号")
	recordVodTHead[2] = encode.ConvertString("人名")
	recordVodTHead[3] = encode.ConvertString("组织树ID")
	recordVodTHead[4] = encode.ConvertString("组织树")
	recordVodTHead[5] = encode.ConvertString("部门ID")
	recordVodTHead[6] = encode.ConvertString("营销部门")
	recordVodTHead[7] = encode.ConvertString("组织名称")
	recordVodTHead[8] = encode.ConvertString("营销工号分类")
	recordVodTHead[9] = encode.ConvertString("实体渠道类型")
	recordVodTHead[10] = encode.ConvertString("营销工号细分")
	recordVodTHead[11] = encode.ConvertString("营销区域ID")
	recordVodTHead[12] = encode.ConvertString("营销区域")
	recordVodTHead[13] = encode.ConvertString("营销工号状态")

	recordVodTHead[14] = encode.ConvertString("销售员名称A")
	recordVodTHead[15] = encode.ConvertString("销售人员岗位")
	recordVodTHead[16] = encode.ConvertString("销售人员来源")
	recordVodTHead[17] = encode.ConvertString("销售人员用工性质")
	recordVodTHead[18] = encode.ConvertString("是否新增受理人员")
	recordVodTHead[19] = encode.ConvertString("销售员状态")

	recordVodTHead[20] = encode.ConvertString("销售树ID")
	recordVodTHead[21] = encode.ConvertString("销售人员组织树")
	recordVodTHead[22] = encode.ConvertString("销售员名称")
	recordVodTHead[23] = encode.ConvertString("身份证")
	recordVodTHead[24] = encode.ConvertString("手机号")
	recordVodTHead[25] = encode.ConvertString("邮箱")
	recordVodTHead[26] = encode.ConvertString("销售人员岗位")

	recordVodTHead[27] = encode.ConvertString("销售人员来源")
	recordVodTHead[28] = encode.ConvertString("销售人员用工性质")
	recordVodTHead[29] = encode.ConvertString("是否新增受理人员")
	recordVodTHead[30] = encode.ConvertString("销售员状态")

	err = writer.Write(recordVodTHead)

	record := make([]string, 31)
	for _, v := range workcs {
		record[0] = encode.ConvertString(v.WCode)
		record[1] = encode.ConvertString(v.WID)
		record[2] = encode.ConvertString(v.WName)
		record[3] = encode.ConvertString(v.OID)
		record[4] = encode.ConvertString(v.Org)
		record[5] = encode.ConvertString(v.PID)
		record[6] = encode.ConvertString(v.Partment)
		record[7] = encode.ConvertString(v.OrgName)
		record[8] = encode.ConvertString(v.WIDClass)
		record[9] = encode.ConvertString(v.ChanClass)
		record[10] = encode.ConvertString(v.WIDClassB)
		record[11] = encode.ConvertString(v.AreaID)
		record[12] = encode.ConvertString(v.Area)
		record[13] = encode.ConvertString(v.WIDStatus)
		record[14] = encode.ConvertString(v.WNameA)
		record[15] = encode.ConvertString(v.WJob)
		record[16] = encode.ConvertString(v.WSource)
		record[17] = encode.ConvertString(v.WType)
		record[18] = encode.ConvertString(v.WIsAdd)
		record[19] = encode.ConvertString(v.WStatus)
		record[20] = encode.ConvertString(v.WTreeID)
		record[21] = encode.ConvertString(v.WTreeOrg)
		record[22] = encode.ConvertString(v.WNameB)
		record[23] = encode.ConvertString(v.WCard)
		record[24] = encode.ConvertString(v.WPhone)
		record[25] = encode.ConvertString(v.WEmail)
		record[26] = encode.ConvertString(v.WJobB)
		record[27] = encode.ConvertString(v.WSourceB)
		record[28] = encode.ConvertString(v.WTypeB)
		record[29] = encode.ConvertString(v.WIsAddB)
		record[30] = encode.ConvertString(v.WStatusB)

		err = writer.Write(record)
		if err != nil {
			fmt.Printf("Fail to write file %v\n", outfile)
			return err
		}
	}
	writer.Flush()
	return nil
}

```