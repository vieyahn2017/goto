package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "teldatatools"
	app.Usage = "Auto analysis Tel Data！！！"
	app.Author = "PanYingYun"
	app.Email = "panyingyun@gmail.com/panyy@wasu.com"
	app.Version = "1.0.0"
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "inputa,ia", Value: "", Usage: "teldatatools -ia XXX.csv"},
		cli.StringFlag{Name: "inputb,ib", Value: "", Usage: "teldatatools -ib XXX.csv"},
		cli.StringFlag{Name: "output,o", Value: "", Usage: "teldatatools -ib XXX.csv"},
	}

	app.Action = func(c *cli.Context) {
		infilea := c.String("inputa")
		infileb := c.String("inputb")
		outfile := c.String("output")

		fmt.Println("infilea = " + infilea)
		fmt.Println("infileb = " + infileb)
		fmt.Println("outfile = " + outfile)

		if len(infilea) != 0 {
			fmt.Println(<-TelDataCmp(infilea, infileb, outfile))
		} else {
			fmt.Println("输入数据错误！")
			fmt.Println("Usage: teldatatools -ia XXX.csv -ib XXX.csv")
		}

	}

	app.Run(os.Args)
}

//合并两个表名
//输入:infile 点播数据文件名
func TelDataCmp(infilea string, infileb string, outfile string) <-chan string {
	c := make(chan string)
	go func() {
		//获取文件名前缀

		//从文件A读取并预处理数据
		start := time.Now()
		workas, _ := readWorkA(infilea)
		end := time.Now()
		sort.Sort(workas)
		fmt.Printf("workas size = %v\n", len(workas))
		fmt.Printf("workas 数据读取耗时 %v(秒)\n", end.Sub(start).Seconds())

		//从文件B读取并预处理数据
		start = time.Now()
		workbs, _ := readWorkB(infileb)
		end = time.Now()
		sort.Sort(workbs)
		fmt.Printf("workbs size = %v\n", len(workbs))
		fmt.Printf("workbs 数据读取耗时 %v(秒)\n", end.Sub(start).Seconds())

		start = time.Now()
		//从B表不全A表数据
		workcs := convertWorkC(workas, workbs)
		//检查C表中重复数据，并且输出行号
		checkWorkC(workcs)
		//输出C表格
		writeWorkC(workcs, outfile)
		end = time.Now()
		fmt.Printf("合并耗时 %v(秒)\n", end.Sub(start).Seconds())
		c <- "OK, 合并结束!!!"
	}()
	return c
}

func checkWorkA(workas WorkAS) {
	sort.Sort(workas)
	lenght := len(workas)
	for i, _ := range workas {
		if i+1 < lenght {
			if strings.Compare(workas[i].WCode, workas[i+1].WCode) == 0 {
				fmt.Printf("发现重复数据: \n %v  \n %v \n", workas[i], workas[i+1])
			}
		}
	}
}

func checkWorkB(workbs WorkBS) {
	sort.Sort(workbs)
	lenght := len(workbs)
	for i, _ := range workbs {
		if i+1 < lenght {
			if strings.Compare(workbs[i].WCode, workbs[i+1].WCode) == 0 {
				fmt.Printf("发现重复数据: \n %v  \n %v \n", workbs[i], workbs[i+1])
			}
		}
	}
}

func checkWorkC(workcs WorkCS) {
	sort.Sort(workcs)
	lenght := len(workcs)
	for i, _ := range workcs {
		if i+1 < lenght {
			if strings.Compare(workcs[i].WCode, workcs[i+1].WCode) == 0 {
				fmt.Printf("第[%v]行发现重复数据\n", i+2)
			}
		}
	}
}

func convertWorkC(workas WorkAS, workbs WorkBS) (workcs WorkCS) {
	workcs = make(WorkCS, 0)
	for _, worka := range workas {
		var workc WorkC
		workc.WCode = worka.WCode

		workc.WID = worka.WID
		workc.WName = worka.WName
		workc.OID = worka.OID
		workc.Org = worka.Org
		workc.PID = worka.PID
		workc.Partment = worka.Partment
		workc.OrgName = worka.OrgName

		workc.WIDClass = worka.WIDClass
		workc.ChanClass = worka.ChanClass
		workc.WIDClassB = worka.WIDClassB
		workc.AreaID = worka.AreaID
		workc.Area = worka.Area
		workc.WIDStatus = worka.WIDStatus

		workc.WNameA = worka.WNameA
		workc.WJob = worka.WJob
		workc.WSource = worka.WSource
		workc.WType = worka.WType
		workc.WIsAdd = worka.WIsAdd
		workc.WStatus = worka.WStatus

		workb := findWorkAByWCode(workbs, workc.WCode)

		workc.WTreeID = workb.WTreeID
		workc.WTreeOrg = workb.WTreeOrg
		workc.WNameB = workb.WNameB
		workc.WCard = workb.WCard
		workc.WPhone = workb.WPhone
		workc.WEmail = workb.WEmail
		workc.WJobB = workb.WJobB
		workc.WSourceB = workb.WSourceB
		workc.WTypeB = workb.WTypeB
		workc.WIsAddB = workb.WIsAddB
		workc.WStatusB = workb.WStatusB

		workcs = append(workcs, workc)
	}
	return
}

func findWorkAByWCode(workbs WorkBS, WCode string) (workb WorkB) {
	for _, v := range workbs {
		if strings.Compare(v.WCode, WCode) == 0 {
			workb = v
			break
		}
	}
	return workb
}
