package main

import (
	"bufio"
	"fmt"
	"os"
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

		for strings.Compare(comandosep[contador], "\n") != 0 { // si no viene vacio -> \n

			if strings.Compare(comandosep[contador], " ") == 0 { // si viene espacio
				break
			} else {
				lineacomando += strings.ToLower(comandosep[contador]) // va concatenando cada char del comando
				contador++

			}

		}
	}

	if strings.Compare(lineacomando, "exec") == 0 {
		AnalisisExec(comando)

	} else if strings.Compare(lineacomando, "mkdisk") == 0 {
		AnalisiMkdisk(comando)

	} else if strings.Compare(lineacomando, "rmdisk") == 0 {
		AnalisisRmdisk(comando)

	} else if strings.Compare(lineacomando, "fdisk") == 0 {
		AnalisisFdisk(comando)

	} else if strings.Compare(lineacomando, "mount") == 0 {
		AnalisisMount(comando)
		MonstrarMount()

	} else if strings.Compare(lineacomando, "rep") == 0 {
		AnalsisRep(comando)

	} else if strings.Compare(lineacomando, "pause") == 0 {
		//pause
		fmt.Println("Pause.............")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		fmt.Println(input)
	} else {
		fmt.Println("Error -> Comando invalido")
	}
}
