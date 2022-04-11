package main

import (
	"fmt"
	"strings"
)

func AnalizadorComando(comando string) {

	lineacomando := "" // donde se guarda el primer comando
	contador := 0      //contador general para recorrer el comando
	comandosep := strings.Split(comando, "")

	//comprube si viene vacio el comando
	if strings.Compare(comandosep[0], "\n") == 0 {
		fmt.Println("Error -> Comando vacio")
	} else {
		//simula un while

		for strings.Compare(comandosep[contador], "*") != 0 { // si no viene vacio -> \n

			if strings.Compare(comandosep[contador], " ") == 0 { // si viene espacio
				break
			} else {
				lineacomando += strings.ToLower(comandosep[contador]) // va concatenando cada char del comando
				contador++

			}

		}
	}

	if strings.Compare(lineacomando, "exec") == 0 {
		//fmt.Println("Entro " + lineacomando)
		AnalisisExec(comando)
	} else if strings.Compare(lineacomando, "mkdisk") == 0 {
		AnalisiMkdisk(comando)
	} else if strings.Compare(lineacomando, "rmdisk") == 0 {
		fmt.Println("Entro " + lineacomando)
	} else if strings.Compare(lineacomando, "fdisk") == 0 {
		fmt.Println("Entro " + lineacomando)
	} else if strings.Compare(lineacomando, "mount") == 0 {
		fmt.Println("Entro " + lineacomando)
	} else if strings.Compare(lineacomando, "rep") == 0 {
		fmt.Println("Entro " + lineacomando)
	} else if strings.Compare(lineacomando, "pause") == 0 {
		fmt.Println("Entro " + lineacomando)
	} else {
		fmt.Println("Error -> Comando invalido")
	}
}
