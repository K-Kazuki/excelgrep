package excelsearch

import (
	"fmt"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func Grep(xlsxPath string) (map[string]string, error) {
	if len(xlsxPath) == 0 {
		return nil, nil
	}

	f, err := excelize.OpenFile(xlsxPath)
	if err != nil {
		fmt.Println("an error occurred while open xlsx file.")
		return nil, nil
	}

	cols, err := f.GetCols("Sheet1")
	if err != nil {
		fmt.Println(err.Error())
	}

	result := make(map[string]string)
	for colNum, col := range cols {
		for rowNum, rowCell := range col {

			if len(rowCell) != 0 {
				cellName, err := excelize.CoordinatesToCellName(colNum+1, rowNum+1)
				if err != nil {
					fmt.Println(err)
					continue
				}

				s := fmt.Sprintf("%s: %s", cellName, strings.TrimSpace(rowCell))
				result[cellName] = s
			}
		}
	}

	return result, nil
}
