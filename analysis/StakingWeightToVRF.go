package analysis

import (
	"fmt"
	"github.com/tealeg/xlsx/v3"
	"math"
	"strconv"
	"time"
)

// VRF probability analysis
func genPowerSheet(sh *xlsx.Sheet, aes AnalystEntitys, power float64) {
	// add header
	r := sh.AddRow()
	addrName := r.AddCell()
	addrName.Value = "address"
	weightName := r.AddCell()
	weightName.Value = "weight"

	// add new data
	for _, ae := range aes {
		row := sh.AddRow()
		addr := row.AddCell()
		addr.Value = ae.Address
		weight := row.AddCell()
		newValue := math.Pow(float64(ae.Weight), power)
		weight.Value = strconv.Itoa(int(newValue))
	}
}

// generate weight sheets
func genWeightSheets(sheet *xlsx.Sheet, resultName string) {
	powers := []float64{0.5,0.6,0.7,0.8,0.9}
	aes := GetAnalysisData(sheet)
	file := xlsx.NewFile()
	for _, p := range powers {
		sh,err := file.AddSheet("power"+fmt.Sprintf("%.2f", p))
		if(err != nil) {
			fmt.Println("AddSheet failed","err", err)
		}
		genPowerSheet(sh, aes, p)
	}
	err := file.Save(resultName)
	if err != nil {
		fmt.Println("Save file failed", "err", err)
	}
}

func AnalysisVRF() {
	orgFile := "C:\\code\\go\\src\\github.com\\DataAnalysis\\data\\mainnet_launch_simulation.xlsx"
	sheetName := "simulation"

	now  := time.Now()
	dstFile := "C:\\code\\go\\src\\github.com\\DataAnalysis\\data\\mainnet_launch_simulation_"+strconv.FormatInt(now.Unix(), 10)+".xlsx"
	wb, err := xlsx.OpenFile(orgFile)
	if err != nil {
		fmt.Println(err)
	}

	sh, ok := wb.Sheet[sheetName]
	if(!ok) {
		fmt.Println("sheet name not exist", "sheetName",sheetName)
	}
	genWeightSheets(sh, dstFile)
}