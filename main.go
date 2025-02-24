package main

import (
	"fmt"

	"github.com/kyraslab/go-soal-logic/logic_01"
	"github.com/kyraslab/go-soal-logic/logic_02"
	"github.com/kyraslab/go-soal-logic/utils"
)

func main() {
	fmt.Println("LOGIC-01")

	fmt.Println("Soal No. 1:")
	fmt.Println(utils.PrinterOneDimensionalSlice(logic_01.SoalNo01(10)))
	fmt.Println("Soal No. 2:")
	fmt.Println(utils.PrinterOneDimensionalSlice(logic_01.SoalNo02(10)))
	fmt.Println("Soal No. 3:")
	fmt.Println(utils.PrinterOneDimensionalSlice(logic_01.SoalNo03(10)))
	fmt.Println("Soal No. 4:")
	fmt.Println(utils.PrinterOneDimensionalSlice(logic_01.SoalNo04(10)))
	fmt.Println("Soal No. 5:")
	fmt.Println(utils.PrinterOneDimensionalSlice(logic_01.SoalNo05(10)))
	fmt.Println("Soal No. 6:")
	fmt.Println(utils.PrinterOneDimensionalSlice(logic_01.SoalNo06(10)))
	fmt.Println("Soal No. 7:")
	fmt.Println(utils.PrinterOneDimensionalSlice(logic_01.SoalNo07(10)))
	fmt.Println(utils.PrinterOneDimensionalSlice(logic_01.SoalNo07(11)))
	fmt.Println("Soal No. 8:")
	fmt.Println(utils.PrinterOneDimensionalSlice(logic_01.SoalNo08(10)))
	fmt.Println(utils.PrinterOneDimensionalSlice(logic_01.SoalNo08(11)))
	fmt.Println("Soal No. 9:")
	fmt.Println(utils.PrinterOneDimensionalSlice(logic_01.SoalNo09(10)))
	fmt.Println(utils.PrinterOneDimensionalSlice(logic_01.SoalNo09(11)))
	fmt.Println("Soal No. 10:")
	fmt.Println(utils.PrinterOneDimensionalSlice(logic_01.SoalNo10(10)))
	fmt.Println("Soal No. 11:")
	fmt.Println(utils.PrinterOneDimensionalSlice(logic_01.SoalNo11(10)))
	fmt.Println("Soal No. 12:")
	fmt.Println(utils.PrinterOneDimensionalSlice(logic_01.SoalNo12(12)))

	fmt.Println("LOGIC-02")
	fmt.Println("Soal No. 1:")
	fmt.Println(utils.PrinterTwoDimensionalSlice(logic_02.SoalNo01(9)))
	fmt.Println("Soal No. 2:")
	fmt.Println(utils.PrinterTwoDimensionalSlice(logic_02.SoalNo02(9)))
	fmt.Println("Soal No. 3:")
	fmt.Println(utils.PrinterTwoDimensionalSlice(logic_02.SoalNo03(9)))
}
