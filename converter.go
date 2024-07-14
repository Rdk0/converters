package main

import (
	"fmt"
	"math"
)

func main(){
	var ic50, mw, deltG, nonHatoms, le  float64
	//var formula string
    for {
	fmt.Println("press control-C to exit")
    fmt.Println("input values separated by space for IC50 (nM) and molecular mass (Da)")
	_, err := fmt.Scanf("%f %f", &ic50, &mw)
		if err != nil {
			fmt.Printf("an error occuried %v: ", err)
		}
        deltG = -0.59574*math.Log(ic50/1e9)
	nonHatoms = mw/13.286
	le = deltG/nonHatoms
	fmt.Printf("for IC50 = %0.1f nM, deltaG = -%0.2f kcal/mol \n\n", ic50, deltG)
	fmt.Printf("estimated number of heavy atoms = %0.0f and LE = %0.2f", nonHatoms,le )
	fmt.Println("\n******************************************************")
}
}
