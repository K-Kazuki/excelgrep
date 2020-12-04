package main

import (
	"fmt"

	"github.com/K-Kazuki/excel_grep/excelsearch"
)

func main() {
	fmt.Println("START")

	result, err := excelsearch.Grep("sample_files/sample.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

	fmt.Println("END")
}
