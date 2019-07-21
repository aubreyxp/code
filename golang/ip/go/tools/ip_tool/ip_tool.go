package ip_tool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	OperatorMp = map[string]int{
		"电信":   1,
		"联通":   1,
		"铁通":   1,
		"长城":   1,
		"移动":   1,
		"鹏博士":  1,
		"广电网":  1,
		"华数":   1,
		"软银":   1,
		"CZ88": 1,
	}

	AddressMp = map[string]int{
		"北京":  1,
		"天津":  1,
		"上海":  1,
		"重庆":  1,
		"河北":  1,
		"山西":  1,
		"辽宁":  1,
		"吉林":  1,
		"黑龙江": 1,
		"江苏":  1,
		"浙江":  1,
		"安徽":  1,
		"福建":  1,
		"江西":  1,
		"山东":  1,
		"河南":  1,
		"湖北":  1,
		"湖南":  1,
		"广东":  1,
		"海南":  1,
		"四川":  1,
		"贵州":  1,
		"云南":  1,
		"陕西":  1,
		"甘肃":  1,
		"青海":  1,
		"台湾":  1,
		"内蒙":  1,
		"广西":  1,
		"西藏":  1,
		"宁夏":  1,
		"新疆":  1,
		"香港":  1,
		"澳门":  1,
	}
)

func getArea(area string) string {
	for k, _ := range OperatorMp {
		if strings.Contains(area, k) {
			return k
		}
	}

	return "其它"
}

func getAdress(address string) string {
	for k, _ := range AddressMp {
		if strings.Contains(address, k) {
			return k
		}
	}

	return "其它"
}

type IpInfo struct {
	Ip       string `json:"ip"`
	Operator string `json:"area"`
	Address  string `json:"country"`
	Count    int
}

func GetIpInfoMap(ipList []string) (infoMap map[string]*IpInfo, err error) {
	ipInfoMap := map[string]*IpInfo{}
	for _, ip := range ipList {
		if _, ok := ipInfoMap[ip]; ok {
			ipInfoMap[ip].Count++
			continue
		}

		if ipInfo, err := getIpInfo(ip); err != nil {
			fmt.Println("GetIpInfoList Error:", err)
			return ipInfoMap, err
		} else {
			ipInfo.Count++
			ipInfoMap[ip] = ipInfo
		}
	}

	return ipInfoMap, nil
}

func getIpInfo(ip string) (*IpInfo, error) {
	// http://127.0.0.1:2060/?ip=114.231.164.21
	ii := &IpInfo{}
	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:2060/?ip=%s", ip))
	if err != nil {
		return ii, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ii, err
	}

	ret := map[string]*IpInfo{}
	if err := json.Unmarshal(body, &ret); err != nil {
		return ii, err
	}
	//fmt.Println("body:", string(body))
	//retJson, err := json.Marshal(ret)
	//fmt.Printf("ret:%+v %v \n", string(retJson), err)
	return ret[ip], nil
}
