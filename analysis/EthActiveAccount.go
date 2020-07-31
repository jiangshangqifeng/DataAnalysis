package analysis

import (
	"fmt"
	xlsx_v3 "github.com/tealeg/xlsx/v3"
	"math"
	"sort"
)

// ethereum active account analysis

type AnalystEntity struct {
	Address string
	Weight  int
}

type AnalystEntitys []AnalystEntity

func (s AnalystEntitys) Len() int           { return len(s) }
func (s AnalystEntitys) Less(i, j int) bool { return s[i].Weight > s[j].Weight }
func (s AnalystEntitys) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func GetAnalysisData(sheet *xlsx_v3.Sheet) AnalystEntitys {
	var aes AnalystEntitys

	i := 0
	for {
		if i >= sheet.MaxRow {
			break
		}

		//skip header
		if 0 == i {
			i++
			continue
		}

		var row AnalystEntity
		// get the address
		addressCell, err := sheet.Cell(i, 0)
		if err != nil {
			panic(err)
		}
		row.Address = addressCell.Value
		// get the count
		countCell, err := sheet.Cell(i, 1)
		if err != nil {
			fmt.Println(err)
		}
		// we got a cell, but what's in it?
		row.Weight, err = countCell.Int()
		if(err != nil) {
			fmt.Println(err)
		}
		aes = append(aes, row)
		i++
	}
	return  aes
}

// sum aes[0....i]'s Weight
// i is from 0
func sumWeighti(i int, aes AnalystEntitys) int {
	sum := 0
	for j, ae := range aes {
		if j > i {
			break
		}
		sum += ae.Weight
	}
	return sum
}

// Estimate self-similar's h
// Σweight[0,i] = 1 - h
// N*h = i
func GetSelfSimilarH(aes AnalystEntitys) float64 {
	h, min := float64(0), float64(1)
	indexH := 0
	N := aes.Len()
	curWeighti := 0
	totalWeight := sumWeighti(N, aes)
	fmt.Println("N",N, "totalWeight",totalWeight)
	for i, ae := range aes {
		curWeighti += ae.Weight
		weight := float64(curWeighti) / float64(totalWeight)
		h = float64(i+1)/float64(N)
		value := float64(weight -1) + h
		if math.Abs(value) < min {
			min = math.Abs(value)
			indexH = i+1
		}
		if i < 50 {
			fmt.Println("weight",weight,"h",h,"value",value,"min",min,"indexH",indexH)
		}
	}
	fmt.Println("min",min,"indexH",indexH)
	return float64(indexH) / float64(N)
}

func AnalysisEthereum() {
	fmt.Println("Starting a test")
	file := "C:\\code\\go\\src\\github.com\\DataAnalysis\\data\\Ethereum_24h_ActiveAccount1.xlsx"
	sheet := "ActiveAccounts"

	wb, err := xlsx_v3.OpenFile(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("Sheets in this file:")
	for i, sh := range wb.Sheets {
		fmt.Println(i, sh.Name)
	}
	fmt.Println("----")

	sh, ok := wb.Sheet[sheet]
	if(!ok) {
		fmt.Println("sheet name not exist", "sheet",sheet)
	}
	aes := GetAnalysisData(sh)
	sort.Sort(aes)
	//for i, ae := range aes {
	//	fmt.Println("row",i, "ae.address", ae.Address, "ae.count",ae.Weight)
	//}
	h := GetSelfSimilarH(aes)
	fmt.Println("h",h)
}