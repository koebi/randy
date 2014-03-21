package main

import (
    "fmt"
    "os"
    "encoding/csv"
)

var csvPath string = "./csv/magic_items_full.csv"

func main() {
    itemData, err := os.Open(csvPath)
    defer itemData.Close()
    if err != nil {
        panic(err)
    }
    reader := csv.NewReader(itemData)
    reader.Comma = ','

    //read the first line of the CSV, which contains the Table Headers
    //into itemSpec
    itemSpec, err := reader.Read()
    if err != nil {
        panic(err)
    }

    //read the rest of the CSV-File into the twodimensional array
    //items
    items, err := reader.ReadAll()
    if err != nil {
        panic(err)
    }
    
    //find out which index corresponds to which value
    var valIndex =make(map[string]int)
    for j := range itemSpec {
        valInd[itemSpec[j]]=j
    }

    

}




