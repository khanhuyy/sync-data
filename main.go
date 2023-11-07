package main

import (
    "fmt"

    "github.com/xuri/excelize/v2"
)

func main() {
    f, err := excelize.OpenFile("update.xlsx")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer func() {
        // Close the spreadsheet.
        if err := f.Close(); err != nil {
            fmt.Println(err)
        }
    }()
    // Get value from cell by given worksheet name and cell reference.
    // cell, err := f.GetCellValue("Update 1 27.10", "B2")
    // if err != nil {
    //     fmt.Println(err)
    //     return
    // }
    // fmt.Println(cell)
    // Get all the rows in the Sheet1.
    rows, err := f.GetRows("Update 1 27.10")
    if err != nil {
        fmt.Println(err)
        return
    }
    for _, row := range rows {
        for _, colCell := range row {
            fmt.Println(colCell, "\t")
        }
    }
}