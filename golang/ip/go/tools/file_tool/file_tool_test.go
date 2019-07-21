package file_tool

import (
	"fmt"
	"testing"
)

func TestGetFileRawList(t *testing.T) {
	if list, err := GetFileRawList("../a.log.2"); err != nil {
		t.Error(err)
	} else if len(list) != 1008811 {
		t.Error("raw count not right")
	} else {
		t.Log("succeed")
	}
}

func TestWriteMapToExcelBySort(t *testing.T) {
	fileName := "sortfile"
	dataMp := map[string]int{
		"电信":  3,
		"电信1": 9,
		"电信4": 7,
	}

	titles := []string{"运营商", "数量"}
	if err := WriteMapToExcelBySort(fileName, titles, dataMp); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}
}
