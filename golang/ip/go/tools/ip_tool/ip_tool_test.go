package ip_tool

import "testing"
import "encoding/json"

func TestGetIpInfoMap(t *testing.T) {
	{
		ipList := []string{}
		infoMap, err := GetIpInfoMap(ipList)
		if err != nil || len(infoMap) != 0 {
			t.Error("wrong")
		} else {
			t.Log("success")
		}
	}

	{

		ipList := []string{"106.113.59.104", "112.14.69.127", "124.77.245.10", "124.77.245.10"}
		infoMap, err := GetIpInfoMap(ipList)
		if err != nil || len(infoMap) != 3 {
			t.Error("wrong")
		} else {
			t.Log("success")
			if bs, err := json.Marshal(infoMap); err != nil {
				t.Error("json Marshal err:", err)
			} else {
				t.Log(string(bs))
			}
		}
	}
}
