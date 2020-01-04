package main

import (
	"bufio"
	"fmt"
	"gitlab.com/tool-excel/filtros"
	"gitlab.com/tool-excel/helpers"
	"os"
	"strings"
)

func main() {
	var index int
	var indexMenu int

	fmt.Println("\n\t*****************************")
	fmt.Println("\n\t TOOL XLSX \n created by: Greko Maister")
	fmt.Println("\n\t*****************************")
	fmt.Println("\n\n MENU\n" +
		"(1) FIND VALUE IN ROW.\n" +
		"(2) COMPARE 2 FILES FOR COL")
	fmt.Println("\nIngrese la opcion del menu:")

	fmt.Scan(&indexMenu)

	switch indexMenu {
	case 1:
		fmt.Println("Ingrese el path del archivo 1:")
		file1 := bufio.NewReader(os.Stdin)
		file_1, _ := file1.ReadString('\n')
		fmt.Println("Cargando el archivo 1:", strings.TrimRight(file_1, "\n"))

		xlsx := helpers.LoadFile(file_1)

		fmt.Println("\t Nombre de la Hoja: ", xlsx.GetSheetName(1))
		rows := xlsx.GetRows(xlsx.GetSheetName(1)) //toma las filas

		indicesSearch := filtros.IndicesFilesForSearch(rows) //Retorna los indices
		fmt.Println("\tOpciones de Filtrado :\n", indicesSearch)
		fmt.Println("------------------------------------------")
		fmt.Println("\t Ingrese el indice de busqueda:")

		_, err := fmt.Scan(&index)
		if err != nil {
			fmt.Println(err)
			//os.Exit(1)
		}
		fmt.Println("Ingrese el Valor de busqueda:")
		valueSearch := bufio.NewReader(os.Stdin) //URJ845
		valueSearch_, _ := valueSearch.ReadString('\n')

		if len(valueSearch_) > 0 {
			valueSearch_ = strings.TrimRight(valueSearch_, "\n")
			filtros.FindForCol(rows, valueSearch_, index)
		}
	case 2: //Compara 2 Archivos
		var indexF1 int
		var indexF2 int
		fmt.Println("Ingrese el nombre del archivo 1:")
		file1 := bufio.NewReader(os.Stdin)
		file_1, _ := file1.ReadString('\n')
		fmt.Println("Cargando el archivo 1:", strings.TrimRight(file_1, "\n"))

		xlsxf1 := helpers.LoadFile(file_1)
		fmt.Println("Nombre de la Hoja", xlsxf1.GetSheetName(1))
		rowsf1 := xlsxf1.GetRows(xlsxf1.GetSheetName(1))

		fmt.Println("Ingrese el nombre del archivo 2:")
		file2 := bufio.NewReader(os.Stdin)
		file_2, _ := file2.ReadString('\n')
		fmt.Println("Cargando el archivo 2:", strings.TrimRight(file_2, "\n"))

		xlsxf2 := helpers.LoadFile(file_2)
		fmt.Println("Nombre de la Hoja", xlsxf2.GetSheetName(1))
		rowsf2 := xlsxf2.GetRows(xlsxf2.GetSheetName(1))

		indicesSearch := filtros.IndicesFilesForSearch(rowsf1)
		fmt.Println("Opciones de Filtrado : "+xlsxf1.GetSheetName(1)+" \n", indicesSearch)

		fmt.Println("\nIngrese el indice de busqueda del Archivo 1")
		_, _ = fmt.Scan("%d", &indexF1)
		indicesSearch2 := filtros.IndicesFilesForSearch(rowsf2)
		fmt.Println("Opciones de Filtrado : "+xlsxf2.GetSheetName(1)+" \n", indicesSearch2)
		fmt.Println("\nIngrese el indice de busqueda del Archivo 2")
		_, _ = fmt.Scan("%d", &indexF2)
		filtros.Compare_Files(rowsf1, rowsf2, indexF1, indexF2)
	default:

	}
}
