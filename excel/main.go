package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
)

// Estructura para leer `sharedStrings.xml`
type SharedStrings struct {
	XMLName xml.Name  `xml:"sst"`
	Strings []SSValue `xml:"si>t"`
}

type SSValue struct {
	Text string `xml:",chardata"`
}

// Estructura para leer `sheet1.xml`
type Worksheet struct {
	XMLName   xml.Name `xml:"worksheet"`
	SheetData struct {
		Rows []Row `xml:"row"`
	} `xml:"sheetData"`
}

type Row struct {
	Cells []Cell `xml:"c"`
}

type Cell struct {
	Type  string `xml:"t,attr,omitempty"`
	Value string `xml:"v"`
}

func main() {
	// Leer `sharedStrings.xml`
	sharedStringsMap, err := leerSharedStrings("xl/sharedStrings.xml")
	if err != nil {
		fmt.Println("Error leyendo sharedStrings:", err)
		return
	}

	// Leer `sheet1.xml`
	ws, err := leerWorksheet("xl/worksheets/sheet1.xml")
	if err != nil {
		fmt.Println("Error leyendo sheet1:", err)
		return
	}

	// Convertir a JSON dinámico
	jsonData, err := convertirAJSON(ws, sharedStringsMap)
	if err != nil {
		fmt.Println("Error convirtiendo a JSON:", err)
		return
	}

	// Guardar JSON
	err = os.WriteFile("output.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error guardando JSON:", err)
		return
	}

	fmt.Println("Datos exportados a output.json")
}

// 📌 Función para leer `sharedStrings.xml` y mapear índices a valores
func leerSharedStrings(filename string) (map[int]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var sst SharedStrings
	err = xml.Unmarshal(data, &sst)
	if err != nil {
		return nil, err
	}

	sharedStrings := make(map[int]string)
	for i, s := range sst.Strings {
		sharedStrings[i] = s.Text
	}
	return sharedStrings, nil
}

// 📌 Función para leer `sheet1.xml`
func leerWorksheet(filename string) (*Worksheet, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var ws Worksheet
	err = xml.Unmarshal(data, &ws)
	if err != nil {
		return nil, err
	}
	return &ws, nil
}

// 📌 Convertir los datos a JSON dinámicamente
func convertirAJSON(ws *Worksheet, sharedStringsMap map[int]string) ([]byte, error) {
	if len(ws.SheetData.Rows) == 0 {
		return nil, fmt.Errorf("la hoja está vacía")
	}

	// 📌 Obtener cabeceras de la primera fila
	var cabeceras []string
	for _, cell := range ws.SheetData.Rows[0].Cells {
		if cell.Type == "s" {
			index, err := strconv.Atoi(cell.Value)
			if err == nil {
				cabeceras = append(cabeceras, sharedStringsMap[index])
			}
		} else {
			cabeceras = append(cabeceras, cell.Value)
		}
	}

	// 📌 Recorrer las filas y mapear los valores a las cabeceras
	var registros []map[string]interface{}
	for i, row := range ws.SheetData.Rows {
		if i == 0 {
			continue // Saltar la fila de cabeceras
		}

		registro := make(map[string]interface{})
		for j, cell := range row.Cells {
			if j >= len(cabeceras) {
				continue // Evitar desbordamiento si hay más columnas de lo esperado
			}

			var valor interface{}
			if cell.Type == "s" { // Si es shared string
				index, err := strconv.Atoi(cell.Value)
				if err == nil {
					valor = sharedStringsMap[index]
				} else {
					valor = cell.Value
				}
			} else {
				// Convertir valores numéricos si es posible
				if num, err := strconv.Atoi(cell.Value); err == nil {
					valor = num
				} else if numF, err := strconv.ParseFloat(cell.Value, 64); err == nil {
					valor = numF
				} else {
					valor = cell.Value
				}
			}

			// Asignar el valor al mapa con la cabecera correspondiente
			registro[cabeceras[j]] = valor
		}

		registros = append(registros, registro)
	}

	// Convertir a JSON
	return json.MarshalIndent(registros, "", "    ")
}
