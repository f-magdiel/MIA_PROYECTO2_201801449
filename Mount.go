package main

import (
	"fmt"
	"strings"
)

func AnalisisMount(comando string) {
	contador := 0
	var linecomand [100]string
	newcomando := strings.Split(comando, "")
	lineacomandos := ""
	// sre realizar una copia del array para mejor manejo
	copy(linecomand[:], newcomando[:])
	//banderas
	flag_path := false
	flag_name := false

	//valores
	valor_path := ""
	valor_name := ""

	//simula un while
	for strings.Compare(linecomand[contador], "\n") != 0 && strings.Compare(linecomand[contador], "#") != 0 {
		//validacion de caracters por interrupcion
		if strings.Compare(linecomand[contador], " ") == 0 {
			contador++
			lineacomandos = ""
		} else if strings.Compare(linecomand[contador], "=") == 0 {
			lineacomandos += strings.ToLower(linecomand[contador])
			contador++
		} else {
			lineacomandos += strings.ToLower(linecomand[contador])
			contador++
		}

		//validacion de valores de comandos
		if strings.Compare(lineacomandos, "mount") == 0 {
			fmt.Println("Encontro : " + lineacomandos)
			lineacomandos = ""
			contador++
		} else if strings.Compare(lineacomandos, "-paht=") == 0 {
			fmt.Println("Encontro : " + lineacomandos)
			lineacomandos = ""
			flag_path = true
			for strings.Compare(linecomand[contador], "\n") != 0 {
				if strings.Compare(linecomand[contador], "\"") == 0 { // si viene con comilla doble
					contador++
					//simula un while
					for strings.Compare(linecomand[contador], "\n") != 0 {
						if strings.Compare(linecomand[contador], "\"") == 0 { //finaliza path
							contador++
							break
						} else {
							valor_path += linecomand[contador]
							contador++
						}
					}
				} else {
					if strings.Compare(linecomand[contador], " ") == 0 || strings.Compare(linecomand[contador], "\n") == 0 {
						contador++
						break
					} else {
						valor_path += linecomand[contador]
						contador++
					}
				}
			}
			fmt.Println("Valor: " + valor_path)
		} else if strings.Compare(lineacomandos, "-name=") == 0 {
			fmt.Println("Encontro : " + lineacomandos)
			lineacomandos = ""
			flag_name = true
			//simula un while
			for strings.Compare(linecomand[contador], "\n") != 0 {
				if strings.Compare(linecomand[contador], " ") == 0 || strings.Compare(linecomand[contador], "\n") == 0 {
					contador++
					break
				} else {
					valor_name += linecomand[contador]
					contador++
				}
			}
			fmt.Println("Valor: " + valor_name)
		}

	}

	//validacion de valores
	if flag_name == true && flag_path == true {
		//se manda a montar
	} else {
		fmt.Println("Error -> Los parametros no son validos")
	}
}
