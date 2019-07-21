package file_tool

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"

	"ip/internal/common"

	"github.com/tealeg/xlsx"
)

func GetFileRawList(fileName string) (list []string, err error) {
	fi, errf := os.Open(fileName)
	if errf != nil {
		fmt.Printf("Error: %s\n", errf)
		return list, errf
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		ip, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		list = append(list, string(ip))
	}

	return list, nil
}

// sort by value, up to low
func WriteMapToExcelBySort(fileName string, titles []string, dataMp map[string]int) error {
	if fileName == "" {
		return errors.New("fileName empty")
	}

	if len(titles) != 2 {
		return errors.New("titles count wrong")
	}

	if len(dataMp) == 0 {
		return errors.New("data empty")
	}

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		return err
	}
	row = sheet.AddRow()
	for _, title := range titles {
		cell = row.AddCell()
		cell.Value = title
	}

	sortList := common.SortMapByValue(dataMp)
	fmt.Printf("%+v", sortList)

	for _, data := range sortList {
		row = sheet.AddRow()

		cell = row.AddCell()
		cell.Value = data.Key

		cell = row.AddCell()
		cell.Value = fmt.Sprintf("%d", data.Value)
	}

	err = file.Save(fmt.Sprintf("%s.xlsx", fileName))
	if err != nil {
		return err
	}

	return nil
}
