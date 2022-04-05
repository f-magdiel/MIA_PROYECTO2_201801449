package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func LeerArchivo(path string) {
	//abrir el archivo
	file, err := os.Open(path)

	//si hay error al abrir el achivo
	if err != nil {
		log.Fatalf("Error al abrir archivo: %s", err)
	}
	filescanner := bufio.NewScanner(file)

	for filescanner.Scan() {
		linea := string(filescanner.Text())
		if len(linea) != 0 {
			newline := strings.Split(linea, "")
			if strings.Compare(newline[0], "#") != 0 { // si es un comentario se omite
				AnalizadorComando(linea)
			}
		}

	}
	//error mientras se lee el archivo
	if err := filescanner.Err(); err != nil {
		log.Fatalf("Error en el archivo %s", err)
	}
	// se cierra el archivo
	file.Close()
}
