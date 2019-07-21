package main

import (
	"fmt"
	"ip/internal/app/ip_statistic/viewer"
	"ip/tools/file_tool"
	"ip/tools/ip_tool"
)

func main() {
	var (
		ipList   []string
		err      error
		fileName string
		titles   []string
		iv       viewer.IpViewer
		infMap   map[string]*ip_tool.IpInfo
	)

	if ipList, err = file_tool.GetFileRawList("a.log.2"); err != nil {
		fmt.Println("Error:", err)
		return
	}

	infMap, err = ip_tool.GetIpInfoMap(ipList)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	iv = &viewer.IpUniqueViewer{}

	// 1 去重ip后的省份统计
	unionAddressMap := iv.GetAddressInfo(infMap)
	fileName = "IP省份统计(去重复)"
	titles = []string{"省份", "数量"}
	if err = file_tool.WriteMapToExcelBySort(fileName, titles, unionAddressMap); err != nil {
		fmt.Println("Error:", err)
	}

	// 2 去重ip后的运营商统计
	unionOpertorMap := iv.GetOperatorInfo(infMap)
	fileName = "IP运营商统计(去重复)"
	titles = []string{"运营商", "数量"}
	if err = file_tool.WriteMapToExcelBySort(fileName, titles, unionOpertorMap); err != nil {
		fmt.Println("Error:", err)
	}

	iv = &viewer.IpAllViewer{}

	// 3 不去重的城市统计
	allAddressMap := iv.GetAddressInfo(infMap)
	fileName = "IP省份统计"
	titles = []string{"省份", "数量"}
	if err = file_tool.WriteMapToExcelBySort(fileName, titles, allAddressMap); err != nil {
		fmt.Println("Error:", err)
	}

	// 4 不去重的运营商统计
	allOpertorMap := iv.GetOperatorInfo(infMap)
	fileName = "IP运营商统计"
	titles = []string{"运营商", "数量"}
	if err = file_tool.WriteMapToExcelBySort(fileName, titles, allOpertorMap); err != nil {
		fmt.Println("Error:", err)
	}
}
