package viewer

import (
	"ip/tools/ip_tool"
)

type IpViewer interface {
	GetOperatorInfo(ipInfoMap map[string]*ip_tool.IpInfo) map[string]int
	GetAddressInfo(ipInfoMap map[string]*ip_tool.IpInfo) map[string]int
}

type IpUniqueViewer struct {
}

func (iuv *IpUniqueViewer) GetOperatorInfo(ipInfoMap map[string]*ip_tool.IpInfo) map[string]int {
	statistics := map[string]int{}
	for _, info := range ipInfoMap {
		statistics[info.Operator]++
	}
	return statistics
}

func (iuv *IpUniqueViewer) GetAddressInfo(ipInfoMap map[string]*ip_tool.IpInfo) map[string]int {
	statistics := map[string]int{}
	for _, info := range ipInfoMap {
		statistics[info.Address]++
	}
	return statistics
}

type IpAllViewer struct {
}

func (iav *IpAllViewer) GetAddressInfo(ipInfoMap map[string]*ip_tool.IpInfo) map[string]int {
	statistics := map[string]int{}
	for _, info := range ipInfoMap {
		statistics[info.Address] += info.Count
	}
	return statistics
}

func (iav *IpAllViewer) GetOperatorInfo(ipInfoMap map[string]*ip_tool.IpInfo) map[string]int {
	statistics := map[string]int{}
	for _, info := range ipInfoMap {
		statistics[info.Operator] += info.Count
	}
	return statistics
}
