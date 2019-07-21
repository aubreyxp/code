package viewer

import (
	"fmt"
	"ip/tools/ip_tool"
	"testing"
)

func TestIpUnionViewer(t *testing.T) {

	list := []ip_tool.IpInfo{}
	for i := 1; i <= 1000; i++ {
		list = append(list, ip_tool.IpInfo{
			Operator: "电信",
			Address:  "北京",
			Count:    i,
		})
	}

	iuv := &IpUniqueViewer{}
	fmt.Println("union opertor:", iuv.GetOperatorInfo(list))
	fmt.Println("union address:", iuv.GetAddressInfo(list))
	//type IpAllViewer struct {
	iav := &IpAllViewer{}
	fmt.Println("all opertor:", iav.GetOperatorInfo(list))
	fmt.Println("all address:", iav.GetAddressInfo(list))
}
