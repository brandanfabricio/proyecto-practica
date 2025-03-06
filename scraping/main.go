package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {

	resp, err := http.Get("https://tunovelaligeras.com/novelas/a-will-eternal-tnl/")

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	size := len(string(body))
	fmt.Println(size)

	index := strings.Index(string(body), `<div class="summary_image"`)

	end := strings.Index(string(body[index:]), `</div>`)

	sizediv := len("</div>")

	fmt.Println(index)
	fmt.Println(end)
	divImg := string(body[index : index+end+sizediv])

	imgIndex := strings.Index(string(divImg), "src")

	fmt.Println(string(divImg[imgIndex]))

	// file, err := os.Create("test1.html")

	// if err != nil {
	// 	fmt.Println("ERROR crear ")
	// }
	// _, err = file.WriteString(string(body))

	// if err != nil {
	// 	fmt.Println("ERROR escribir ")
	// }

	// fmt.Println(string(body))
	// doc, err := html.Parse(body.Body)
	// if err != nil {
	// 	fmt.Println("Error al parsear HTML:", err)
	// 	return
	// }
	// fmt.Println(resp.Body)

}
