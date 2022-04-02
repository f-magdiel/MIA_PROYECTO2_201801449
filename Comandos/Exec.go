package Comandos

import (
	"fmt"
	"strings"
)

func AnalisisExec(comando string) {
	//variables que contendrán data
	path_exec := ""
	lineacomando := ""
	newcomando := strings.Split(comando, "") // se aplica split para separar cada caracter
	contador := 0
	//simula un while
	for strings.Compare(newcomando[contador], "\n") != 0 { // buscar a exec
		if strings.Compare(newcomando[contador], " ") == 0 {
			contador++
			lineacomando = "" // se resetea
			break
		} else if strings.Compare(newcomando[contador], "=") == 0 {
			lineacomando += strings.ToLower(newcomando[contador])
			contador++
		} else {
			lineacomando += strings.ToLower(newcomando[contador])
			contador++
		}

		if strings.Compare(lineacomando, "exec") == 0 {
			fmt.Println("Encontro : " + lineacomando)
			lineacomando = ""
			contador++
		} else if strings.Compare(lineacomando, "-path=") == 0 {
			fmt.Println("Encontro: " + lineacomando)
			lineacomando = ""
			//simula in while
			for strings.Compare(newcomando[contador], "\n") != 0 {
				if strings.Compare(newcomando[contador], "\"") == 0 { // si viene con comilla doble
					contador++
					//simula un while
					for strings.Compare(newcomando[contador], "\n") != 0 {
						if strings.Compare(newcomando[contador], "\"") == 0 { //finaliza path
							contador++
							break
						} else {
							path_exec += newcomando[contador]
							contador++
						}
					}
				} else {
					if strings.Compare(newcomando[contador], " ") == 0 || strings.Compare(newcomando[contador], "\n") == 0 {
						contador++
						break
					} else {
						path_exec += newcomando[contador]
						contador++
					}
				}
			}
			fmt.Println("Valor: " + path_exec)
		}
	}
}
