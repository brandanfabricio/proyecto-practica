package main

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

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

	file, err := os.Open("./orden.xlsx")

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

	for index, file := range reader.File {

		if file.Name == "xl/sharedStrings.xml" {
			fmt.Println(index, " Archivo dentro del ZIP: ", file.Name)
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
	// for _, s := range sharedStringsMap {

	// }

	// dataExcel := map[string]interface{}
	var dataExcel []map[string]interface{}

	for index, file := range reader.File {

		if file.Name == "xl/worksheets/sheet1.xml" {
			fmt.Printf("%d -> %v \n", index, file.Name)
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
			// var headers []string
			var headers2 map[string]interface{} = make(map[string]interface{})
			for indexRow, row := range dataSheet.SheetData.Rows {

				if indexRow == 0 {

					for _, cell := range row.Cells {
						if cell.Type == "s" {
							indexExce, err := strconv.Atoi(cell.Value)
							if err == nil {
								clearSpace := strings.TrimSpace(sharedStringsMap[indexExce])

								format := strings.ReplaceAll(strings.ReplaceAll(clearSpace, "/", "_"), " ", "_")

								format = strings.ReplaceAll(format, ".", "_")
								// headers = append(headers, format)
								// header := make(map[string]interface{})
								value := string(rune(cell.Ref[0]))
								// header[value] = format
								// headers2 = append(headers2, header)
								headers2[value] = format
							}

						}

					}
				}
				dataCel := make(map[string]interface{})
				// mal la lectura de las cabexxeras
				if indexRow == 4 {
					fmt.Println(headers2)
					fmt.Println(row)

					for _, cell := range row.Cells {
						indexExce, err := strconv.Atoi(cell.Value)
						if err == nil {
							// fmt.Println(index, "-> ", sharedStringsMap[indexExce])
							value := string(rune(cell.Ref[0]))
							valueH, ok := headers2[value].(string)
							if ok {
								dataCel[valueH] = sharedStringsMap[indexExce]

							}

							//
						}

					}
				}

				if indexRow > 0 {

					dataExcel = append(dataExcel, dataCel)
				}

			}
		}

	}

	fmt.Println()
	jsonData, err := json.MarshalIndent(dataExcel, "", "  ")
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}

	// Imprimir JSON

	test, err := os.Create("test2.json")

	if err != nil {
		fmt.Println("error creacion ")
	}
	test.WriteString(string(jsonData))

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
