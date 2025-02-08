package main

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
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

	file, err := os.Open("rep.xlsx")

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
	// sharedStringsMap := make(map[int]string)

	// for index, file := range reader.File {

	// 	if index == 7 {
	// 		fmt.Println(index, " Archivo dentro del ZIP: ", file.Name)
	// 		rc, err := file.Open()

	// 		if err != nil {
	// 			fmt.Println("Error al abrir archivo dentro del ZIP:", err)
	// 			continue
	// 		}
	// 		// Leer el contenido del archivo
	// 		var sst SharedStrings
	// 		// le el contendi y genera la decodificacion con xml me retorna *
	// 		decoder := xml.NewDecoder(rc)
	// 		// uso la decodficacion  generado y aplico en metodo para formar en un lenjuate entendible
	// 		if err := decoder.Decode(&sst); err != nil {
	// 			fmt.Println("Error al parsear XML:", err)
	// 			return
	// 		}
	// 		for i, si := range sst.Strings {
	// 			// clearSpace := strings.TrimSpace(si.Text)
	// 			// format := strings.ReplaceAll(clearSpace, " ", "_")
	// 			// format = strings.ReplaceAll(format, ".", "-")
	// 			// fmt.Printf("%v i -> %s \n", i, format)
	// 			sharedStringsMap[i] = strings.TrimSpace(si.Text)
	// 		}
	// 	}

	// }

	// for _, s := range sharedStringsMap {

	// 	fmt.Println(s)
	// }

	for index, file := range reader.File {

		if index == 4 {
			fmt.Printf("%d -> %v \n", index, file.Name)
			rc, err := file.Open()

			if err != nil {
				fmt.Println("Error al abrir archivo dentro del ZIP:", err)
			}

			var dateSheet Worksheet

			decoder := xml.NewDecoder(rc)

			if err := decoder.Decode(&dateSheet); err != nil {
				fmt.Println("Error al parsear XML:", err)
				return

			}

		}

	}

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
