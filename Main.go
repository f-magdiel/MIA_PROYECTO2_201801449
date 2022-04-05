package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	salir := true
	for salir {
		fmt.Println("-----------------------------------------")
		fmt.Println("|  PROYECTO 2                           |")
		fmt.Println("|  MANEJO E IMPLEMETANCION DE ARCHIVOS  |")
		fmt.Println("|  FRANCISCO MAGDIEL ASICONA MATEO      |")
		fmt.Println("|  PRIMER SEMESTRE 2022                 |")
		fmt.Println("-----------------------------------------")
		fmt.Print("COMANDO >> ")
		lector := bufio.NewReader(os.Stdin)   // para leer la cadena de string
		comando, _ := lector.ReadString('\n') // lee cada caracter
		if strings.Compare(comando, "salir\n") == 0 || strings.Compare(comando, "SALIR\n") == 0 {
			salir = false
			fmt.Println("Se detuvo el programa...")
		} else {
			AnalizadorComando(comando)
		}
	}

}
