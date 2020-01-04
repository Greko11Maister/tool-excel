package filtros

import (
	"fmt"
	"strconv"
)

func FindForCol(X [][]string, valueSearch string, indixCol int) {
	var terminado int
	var encontrado bool
	for i := 0; i < len(X); i++ {
		//fmt.Println("Find: ",X[i][indixCol] , valueSearch, X[i][indixCol] == valueSearch)
		if X[i][indixCol] == valueSearch {
			terminado = 0
			fmt.Println("En el metodo Find: ", X[i][indixCol], " en la fila: ", i)
			encontrado = true

		} else {
			terminado++
		}
	}
	if terminado > 0 && !encontrado {
		fmt.Println("No se encontro Find: ", valueSearch, indixCol)
	}

}

func Compare_Files(X [][]string, Y [][]string, indexColF1 int, indexColF2 int) {
	var found bool
	fmt.Println(len(X), len(Y))
	for i := 0; i < len(X); i++ {
		for e := 0; e < len(Y); e++ {
			if X[i][indexColF1] == Y[e][indexColF2] {
				found = true
				fmt.Println("Se encontro la concidencia: File 1 row: (", i, ") ", X[i][indexColF1], " == File 2 row (", e, ")", Y[e][indexColF2])
			}
		}
		if !found {
			fmt.Println("No se encontro File 1 row: (", i, ") ", X[i][indexColF1])
		}

	}

}

func IndicesFilesForSearch(X [][]string) (a string) {
	var Options string
	for i := 0; i < len(X); i++ {
		for e := 0; e < len(X[i]); e++ {
			if i == 0 {
				Options = Options + " | (" + strconv.Itoa(e) + ") " + X[i][e]
			}
		}
	}

	return Options
}
