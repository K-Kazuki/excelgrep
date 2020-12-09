package main

import (
	"fmt"
	"sync"

	"github.com/K-Kazuki/excel_grep/excelsearch"
)

func main() {
	fmt.Println("START")

	res := make(chan excelsearch.Xlsx)
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		defer wg.Done()
		result, err := excelsearch.Grep("えー", "sample_files/sample.xlsx")
		if err != nil {
			fmt.Println(err)
		}
		res <- result
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		result, err := excelsearch.Grep("長い文字", "sample_files/sample2.xlsx")
		if err != nil {
			fmt.Println(err)
		}
		res <- result
	}()

	go func() {
		wg.Wait()
		close(res)
	}()

	for r := range res {
		fmt.Printf("%s\n", r.BookName)
		for _, s := range r.Sheets {
			fmt.Printf("\t%s\n", s.SheetName)
			for _, f := range s.Founds {
				fmt.Printf("\t\t%s : %s\n", f.CellName, f.Found)
			}
		}
	}

	fmt.Println("END")
}
