package main

import (
	"strings"
)

/**
  表1 数据结构(20个字段)
**/
type WorkA struct {
	WID       string //营销工号
	WName     string //人名
	OID       string //组织树ID
	Org       string //组织树
	PID       string //部门ID
	Partment  string //营销部门
	OrgName   string //组织名称
	WIDClass  string //营销工号分类
	ChanClass string //实体渠道类型
	WIDClassB string //营销工号细分
	AreaID    string //营销区域ID
	Area      string //营销区域
	WIDStatus string //营销工号状态
	WCode     string //销售员编码
	WNameA    string //销售员名称A
	WJob      string //销售人员岗位
	WSource   string //销售人员来源
	WType     string //销售人员用工性质
	WIsAdd    string //是否新增受理人员
	WStatus   string //销售员状态
}

type WorkAS []WorkA

//排序
func (slice WorkAS) Len() int {
	return len(slice)
}

func (slice WorkAS) Less(i, j int) bool {
	return strings.Compare(slice[i].WCode, slice[j].WCode) < 0
}

func (slice WorkAS) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

/**
  表2 数据结构(20个字段)
**/
type WorkB struct {
	WCode    string //销售员编码
	WTreeID  string //销售树ID
	WTreeOrg string //销售人员组织树
	WNameB   string //销售员名称B
	WCard    string //身份证
	WPhone   string //手机号
	WEmail   string //邮箱
	WJobB    string //销售人员岗位
	WSourceB string //销售人员来源
	WTypeB   string //销售人员用工性质
	WIsAddB  string //是否新增受理人员
	WStatusB string //销售员状态
}

//排序
type WorkBS []WorkB

func (slice WorkBS) Len() int {
	return len(slice)
}

func (slice WorkBS) Less(i, j int) bool {
	return strings.Compare(slice[i].WCode, slice[j].WCode) < 0
}

func (slice WorkBS) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

/**
  表3 数据结构(20个字段)
**/
type WorkC struct {
	WCode string //销售员编码

	WID       string //营销工号
	WName     string //人名
	OID       string //组织树ID
	Org       string //组织树
	PID       string //部门ID
	Partment  string //营销部门
	OrgName   string //组织名称
	WIDClass  string //营销工号分类
	ChanClass string //实体渠道类型
	WIDClassB string //营销工号细分
	AreaID    string //营销区域ID
	Area      string //营销区域
	WIDStatus string //营销工号状态

	WNameA  string //销售员名称A
	WJob    string //销售人员岗位
	WSource string //销售人员来源
	WType   string //销售人员用工性质
	WIsAdd  string //是否新增受理人员
	WStatus string //销售员状态

	//WCode    string //销售员编码
	WTreeID  string //销售树ID
	WTreeOrg string //销售人员组织树
	WNameB   string //销售员名称B
	WCard    string //身份证
	WPhone   string //手机号
	WEmail   string //邮箱
	WJobB    string //销售人员岗位
	WSourceB string //销售人员来源
	WTypeB   string //销售人员用工性质
	WIsAddB  string //是否新增受理人员
	WStatusB string //销售员状态

}

//排序
type WorkCS []WorkC

func (slice WorkCS) Len() int {
	return len(slice)
}

func (slice WorkCS) Less(i, j int) bool {
	return strings.Compare(slice[i].WCode, slice[j].WCode) < 0
}

func (slice WorkCS) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
