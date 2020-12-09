package excelsearch

import (
	"fmt"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/fatih/color"
)

/*
[
	{
		sheetName: "Sheet1",
		founds: {
			cellName: "A1",
			found: "検索文字列"
		}
	}
]
*/
type Xlsx struct {
	BookName string
	Sheets   []sheet
}

type sheet struct {
	SheetName string
	Founds    []found
}

type found struct {
	CellName string
	Found    string
}

func Grep(sep string, xlsxPath string) (Xlsx, error) {
	if len(xlsxPath) == 0 {
		return Xlsx{}, nil
	}

	// ファイルオープン
	f, err := excelize.OpenFile(xlsxPath)
	if err != nil {
		fmt.Println("an error occurred while open xlsx file.")
		return Xlsx{}, nil
	}

	// 各シート毎に全セルを検索
	sheets, err2 := searchXlsx(f, sep)
	if err2 != nil {
		fmt.Println(err2)
	}

	var book Xlsx
	if len(sheets) > 0 {
		book = Xlsx{
			BookName: xlsxPath,
			Sheets:   sheets,
		}
	}
	return book, nil
}

func searchXlsx(f *excelize.File, sep string) ([]sheet, error) {
	var sheets []sheet
	for _, sheetName := range f.GetSheetList() {
		// シートの全セルを取得
		cols, err := f.GetCols(sheetName)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		// 文字列検索
		var founds []found
		for colNum, col := range cols {
			for rowNum, rowCell := range col {
				if len(rowCell) != 0 {
					cellName, err := excelize.CoordinatesToCellName(colNum+1, rowNum+1)
					if err != nil {
						fmt.Println(err)
						continue
					}

					foundWord := ""
					if s := strings.TrimSpace(rowCell); s != "" {
						foundWord = search(s, sep)
					}

					if foundWord != "" {
						findResult := found{
							CellName: cellName,
							Found:    foundWord,
						}
						founds = append(founds, findResult)
					}
				}
			}
		}

		if len(founds) > 0 {
			s := sheet{
				SheetName: sheetName,
				Founds:    founds,
			}
			sheets = append(sheets, s)
		}
	}
	return sheets, nil
}

func search(s string, sep string) string {
	if res := strings.Index(s, sep); res > -1 {
		// 一致箇所のみハイライトする
		highlight := color.New(color.FgRed).Add(color.Bold).Sprint(s[res : res+len(sep)])
		before := string(s[:res])
		after := string(s[res+len(sep):])
		return fmt.Sprintf("%s%s%s\n", before, highlight, after)
		// fmt.Println(foundWord)
	}
	return ""
}
