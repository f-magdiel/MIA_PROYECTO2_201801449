package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func AnalisisRmdisk(comando string) {

	var lineacomand [100]string
	newcomando := strings.Split(comando, "")
	lineacomando := ""

	//copy
	copy(lineacomand[:], newcomando[:])

	//banderas
	flag_path := false

	//valor
	valor_path := ""

	contador := 0

	//simula un while
	for strings.Compare(lineacomand[contador], "\n") != 0 && strings.Compare(lineacomand[contador], "#") != 0 {
		//validacion de caracteres para interrupcion

		if strings.Compare(lineacomand[contador], " ") == 0 {
			contador++
			lineacomando = ""
		} else if strings.Compare(lineacomand[contador], "=") == 0 {
			lineacomando += strings.ToLower(lineacomand[contador])
			contador++
		} else {
			lineacomando += strings.ToLower(lineacomand[contador])
			contador++
		}

		//validacion de comandos y sus valores
		if strings.Compare(lineacomando, "rmdisk") == 0 {
			fmt.Println("Encontro: " + lineacomando)
			lineacomando = ""
			contador++
		} else if strings.Compare(lineacomando, "-path=") == 0 {
			fmt.Println("Encontro: " + lineacomando)
			lineacomando = ""
			flag_path = true

			//simula in while
			if strings.Compare(lineacomand[contador], "\"") == 0 {

				contador++
				//simula otro while
				for strings.Compare(lineacomand[contador], "\n") != 0 {
					if strings.Compare(lineacomand[contador], "\"") == 0 {
						break
					} else {
						valor_path += lineacomand[contador]
						contador++
					}
				}
			} else {
				//simula un while
				for strings.Compare(lineacomand[contador], "\n") != 0 {
					if strings.Compare(lineacomand[contador], " ") == 0 || strings.Compare(lineacomand[contador], "\n") == 0 {
						break
					} else {
						valor_path += lineacomand[contador]
						contador++
					}
				}
			}

			fmt.Println("Valor: " + valor_path)
		}
	}

	//proceso para eliminar disk
	flag_available := validacionDisk(valor_path) // para verificar si existe el disco a eliminar
	if flag_path == true {

		if flag_available == true {
			fmt.Println("Se eliminará este disco " + valor_path)
			fmt.Print("Desea eliminarlo (S/N): ")
			lec := bufio.NewReader(os.Stdin)       // para leer la cadena de string
			instruccion, _ := lec.ReadString('\n') // lee cada caracter
			if strings.Compare(instruccion, "S\n") == 0 || strings.Compare(instruccion, "s\n") == 0 {
				//funcion para eliminar el disco
				deleteDisk(valor_path)
				fmt.Println("¡Disco eliminado correctamente!")
			}
		} else {
			fmt.Println("Error -> El disco no existe")
		}
	} else {
		fmt.Println("Error -> No está el path ")
	}
}

func deleteDisk(path string) {
	err := os.Remove(path)
	if err != nil {
		log.Fatal("Error al eliminar disco", err)
	}
}

func validacionDisk(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true

}
