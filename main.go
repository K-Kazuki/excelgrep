package main

import (
	"fmt"
	"sync"

	"github.com/K-Kazuki/excel_grep/excelsearch"
)

func main() {
	fmt.Println("START")

	files, err := excelsearch.Find("")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(files)

	word := "えー"

	if len(files) > 0 {
		res := make(chan excelsearch.Book)
		wg := new(sync.WaitGroup)
		for _, f := range files {
			wg.Add(1)
			go func(f string) {
				defer wg.Done()
				result, err := excelsearch.Grep(word, f)
				if err != nil {
					fmt.Println(err)
				}
				if len(result.Sheets) > 0 {
					res <- result
				}
			}(f)
		}

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
	}

	fmt.Println("END")
}
