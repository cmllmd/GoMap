package main

import "fmt"

func main() {
	var horario string
	m := make(map[string]string)
	m["Dia"] = "Bom dia"
	m["Tarde"] = "Boa tarde"
	m["Noite"] = "Boa noite"

	fmt.Println("DIA/TARDE/NOITE")
	fmt.Scanln(&horario)

	switch horario {
	case "DIA":
		fmt.Println(m["Dia"])
	case "TARDE":
		fmt.Println(m["Tarde"])
	case "NOITE":
		fmt.Println(m["Noite"])

	}

}
