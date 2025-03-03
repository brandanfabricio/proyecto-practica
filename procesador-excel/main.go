package main

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

// mejora
/*
funciona
tengo que hacer un struc
mejorar la obtencion de dato y lectura de archivo
obcion de exportar en arra u json

*/

type SharedStrings struct {
	XMLName xml.Name `xml:"sst"`
	Strings []Text   `xml:"si"`
}

type Text struct {
	Text string `xml:"t"`
}

// Estructuras para mapear sheet1.xml
type Worksheet struct {
	XMLName   xml.Name  `xml:"worksheet"`
	SheetData SheetData `xml:"sheetData"`
}

type SheetData struct {
	Rows []Row `xml:"row"`
}

type Row struct {
	Cells []Cell `xml:"c"`
}

type Cell struct {
	Ref   string `xml:"r,attr"` // Posición (A1, B1, etc.)
	Type  string `xml:"t,attr"` // Tipo (s = shared string)
	Value string `xml:"v"`      // Índice o valor
}

func main() {

	file, err := os.Open("./tes2.xlsx")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	start, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}

	zipData := make([]byte, start.Size())
	_, err = file.Read(zipData)

	if err != nil && err != io.EOF {
		fmt.Println("Error al leer el archivo ZIP:", err)
		return
	}

	reader, err := zip.NewReader(bytes.NewReader(zipData), start.Size())
	if err != nil {
		fmt.Println("Error al abrir el ZIP en memoria:", err)
		return
	}

	sharedStringsMap := make(map[int]string)

	// obtengo los valores que no en la planilla principal
	for _, file := range reader.File {
		if file.Name == "xl/sharedStrings.xml" {
			// fmt.Println(index, " Archivo dentro del ZIP: ", file.Name)
			rc, err := file.Open()

			if err != nil {
				fmt.Println("Error al abrir archivo dentro del ZIP:", err)
				continue
			}
			// Leer el contenido del archivo
			var sst SharedStrings
			// le el contendi y genera la decodificacion con xml me retorna *
			decoder := xml.NewDecoder(rc)
			// uso la decodficacion  generado y aplico en metodo para formar en un lenjuate entendible
			if err := decoder.Decode(&sst); err != nil {
				fmt.Println("Error al parsear XML:", err)
				return
			}
			for i, si := range sst.Strings {
				// clearSpace := strings.TrimSpace(si.Text)
				// format := strings.ReplaceAll(clearSpace, " ", "_")
				// format = strings.ReplaceAll(format, ".", "-")
				// fmt.Printf("%v i -> %s \n", i, format)
				sharedStringsMap[i] = strings.TrimSpace(si.Text)
			}
		}

	}

	var data [][]string

	// data := make([][]string,)
	for index, file := range reader.File {
		if file.Name == "xl/worksheets/sheet1.xml" {
			fmt.Println(index, " ", file.Name)
			rc, err := file.Open()

			if err != nil {
				fmt.Println("Error al abrir archivo dentro del ZIP:", err)
			}
			var dataSheet Worksheet
			decoder := xml.NewDecoder(rc)
			if err := decoder.Decode(&dataSheet); err != nil {
				fmt.Println("Error al parsear XML:", err)
				return
			}

			// Inicializar la matriz `data`
			// fmt.Println(len(dataSheet.SheetData.Rows))
			data = make([][]string, len(dataSheet.SheetData.Rows))
			for index, row := range dataSheet.SheetData.Rows {
				// Asegurar que la fila tenga espacio suficiente
				// fmt.Println(len(row.Cells))
				data[index] = make([]string, len(row.Cells))

				for indexCell, cell := range row.Cells {

					// fmt.Println(cel.Ref)
					// fmt.Println(cel.Type)
					// fmt.Println(cel.Value)

					// ################## Headers
					if index == 0 && cell.Type == "s" {

						indexV, err := strconv.Atoi(cell.Value)
						if err == nil {
							value := strings.Trim(string(cell.Value), "")
							// fmt.Printf(" %s -> %s -> %s\n", cell.Ref, sharedStringsMap[index], value)
							dataStri := sharedStringsMap[indexV]
							if dataStri == "" {
								// fmt.Println("-> ", value)
								// data = append(data, value)
								data[index][indexCell] = value
							} else {
								// fmt.Println("<- ", dataStri)
								// fmt.Println("<- ", dataStri)

								data[index][indexCell] = dataStri
								// data[0]
								// data = append(data, dataStri)

							}

						}
					}

					// ################## celdas
					if index > 0 {
						indexV, err := strconv.Atoi(cell.Value)
						if err == nil {
							value := strings.Trim(string(cell.Value), "")
							// fmt.Printf(" %s -> %s -> %s\n", cell.Ref, sharedStringsMap[index], value)

							dataStri := sharedStringsMap[indexV]
							// if indexCell[2]
							if indexCell == 1 {

								date, err := strconv.ParseFloat(value, 64)
								if err != nil {
									fmt.Println("Error convertir")
								}
								valueDate := excelDateToTime(float64(date))
								value = valueDate.Format("2006-01-02")
							}
							if dataStri == "" {
								// fmt.Println("-> ", value)
								// data = append(data, value)
								data[index][indexCell] = value

							} else {
								// fmt.Println("<- ", dataStri)
								// data = append(data, dataStri)
								data[index][indexCell] = dataStri

							}

						}
					}

				}
				// if index == 3 {
				// 	// // return
				// 	break
				// }
			}
		}
	}
	// for _, row := range data {

	// }

	file, err = os.Create("Prueba3.txt")

	if err != nil {
		fmt.Println("Error creando el archivo:", err)
		return
	}
	defer file.Close() // Asegurar que el archivo se cierre al finalizar

	// Escribir cada fila de la matriz en una línea del archivo
	for _, row := range data {
		line := strings.Join(row, "\t") // Unir los elementos con tabulaciones
		_, err := file.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error escribiendo en el archivo:", err)
			return
		}
	}

	fmt.Println("Matriz guardada correctamente en")

	// fmt.Println("######################")
	// fmt.Println(len(data))
	// fmt.Println(data)

	// fmt.Println("##########")
	// fmt.Println(sharedStringsMap)
	// fmt.Println("##########")

	// for _, s := range sharedStringsMap {

	// }

	// fmt.Println("############Json##########")
	// fmt.Println("######################")
	// fmt.Println()
	// jsonData, err := json.MarshalIndent(dataExcel, "", "  ")
	// if err != nil {
	// 	fmt.Println("Error al convertir a JSON:", err)
	// 	return
	// }

	// // Imprimir JSON

	// test, err := os.Create("test2.json")

	// if err != nil {
	// 	fmt.Println("error creacion ")
	// }
	// test.WriteString(string(jsonData))

	// for index, file := range reader.File {

	// 	if index == 4 {
	// 		rc, err := file.Open()

	// 		if err != nil {
	// 			fmt.Println("Error al abrir archivo dentro del ZIP:", err)
	// 			continue
	// 		}

	// 		var ws Worksheet

	// 		decoder := xml.NewDecoder(rc)

	// 		if err := decoder.Decode(&ws); err != nil {
	// 			fmt.Println("Error al parsear XML:", err)
	// 			return
	// 		}

	// 		for _, row := range ws.SheetData.Rows {
	// 			for _, cell := range row.Cells {
	// 				if cell.Type == "s" {
	// 					index, err := strconv.Atoi(cell.Value)
	// 					if err == nil {
	// 						value := strings.Trim(string(cell.Value), "")

	// 						fmt.Printf("Celda %s -> %s -> %s\n", cell.Ref, sharedStringsMap[index], value)
	// 					} else {

	// 						fmt.Printf("Celda %s -> %s\n", cell.Ref, cell.Value)
	// 					}
	// 				}
	// 			}
	// 		}

	// 	}
	// }

}
func excelDateToTime(serial float64) time.Time {
	// Excel cuenta desde 1900-01-01 pero incluye erróneamente 1900-02-29 como día válido (año bisiesto inexistente)
	baseDate := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC) // Se resta 1 día por corrección
	return baseDate.AddDate(0, 0, int(serial))
}

// func crearOb(){

// 		// dataExcel := map[string]interface{}
// 		var dataExcel []map[string]interface{}

// 		for index, file := range reader.File {
// 			if file.Name == "xl/worksheets/sheet1.xml" {
// 				fmt.Printf("%d -> %v \n", index, file.Name)
// 				rc, err := file.Open()
// 				if err != nil {
// 					fmt.Println("Error al abrir archivo dentro del ZIP:", err)
// 				}
// 				var dataSheet Worksheet
// 				decoder := xml.NewDecoder(rc)
// 				if err := decoder.Decode(&dataSheet); err != nil {
// 					fmt.Println("Error al parsear XML:", err)
// 					return
// 				}
// 				// var headers []string
// 				var headers2 map[string]interface{} = make(map[string]interface{})
// 				for indexRow, row := range dataSheet.SheetData.Rows {

// 					if indexRow == 0 {

// 						for _, cell := range row.Cells {
// 							if cell.Type == "s" {
// 								indexExce, err := strconv.Atoi(cell.Value)
// 								if err == nil {
// 									clearSpace := strings.TrimSpace(sharedStringsMap[indexExce])

// 									format := strings.ReplaceAll(strings.ReplaceAll(clearSpace, "/", "_"), " ", "_")

// 									format = strings.ReplaceAll(format, ".", "_")
// 									// headers = append(headers, format)
// 									// header := make(map[string]interface{})
// 									value := string(rune(cell.Ref[0]))
// 									// header[value] = format
// 									// headers2 = append(headers2, header)
// 									headers2[value] = format
// 								}

// 							}

// 						}
// 					}
// 					dataCel := make(map[string]interface{})
// 					// mal la lectura de las cabexxeras
// 					// if indexRow == 4 {
// 					// 	fmt.Println(headers2)
// 					// 	fmt.Println(row)
// 					for _, cell := range row.Cells {
// 						indexExce, err := strconv.Atoi(cell.Value)
// 						if err == nil {
// 							// fmt.Println(index, "-> ", sharedStringsMap[indexExce])
// 							value := string(rune(cell.Ref[0]))
// 							valueH, ok := headers2[value].(string)
// 							if ok {
// 								dataCel[valueH] = sharedStringsMap[indexExce]

// 							}

// 							//
// 						}

// 					}
// 					// }

// 					if indexRow > 0 {

// 						dataExcel = append(dataExcel, dataCel)
// 					}

// 				}
// 			}

// 		}

// }
