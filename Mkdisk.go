package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func AnalisiMkdisk(comando string) {

	var linecomand [100]string
	newcomando := strings.Split(comando, "")
	lineacomando := ""

	// se realiza una copia del array para manejo
	copy(linecomand[:], newcomando[:])
	fmt.Println(linecomand)
	//banderas
	flag_size := false
	flag_fit := false
	flag_unit := false
	flag_path := false

	//valores
	valor_size := ""
	valor_fit := "ff"
	valor_unit := "m"
	valor_path := ""

	contador := 0
	//simula un while
	for strings.Compare(linecomand[contador], "*") != 0 || strings.Compare(linecomand[contador], "#") != 0 {

		//validacion de caracteres para interrupcion
		if strings.Compare(linecomand[contador], " ") == 0 {
			contador++
			lineacomando = ""
		} else if strings.Compare(linecomand[contador], "=") == 0 {
			lineacomando += strings.ToLower(linecomand[contador])
			contador++
		} else if strings.Compare(linecomand[contador], "*") == 0 {
			break
		} else if strings.Compare(linecomand[contador], "#") == 0 {
			break
		} else {
			lineacomando += strings.ToLower(linecomand[contador])
			contador++
		}

		//validacion de valores de comando
		if strings.Compare(lineacomando, "mkdisk") == 0 {
			fmt.Println("Encontro: " + lineacomando)
			lineacomando = ""
			contador++

		} else if strings.Compare(lineacomando, "-size=") == 0 {
			fmt.Println("Encontro: " + lineacomando)
			lineacomando = ""
			flag_size = false
			//simula un while
			for strings.Compare(linecomand[contador], "*") != 0 {
				if strings.Compare(linecomand[contador], " ") == 0 || strings.Compare(linecomand[contador], "\n") == 0 {
					contador++
					break
				} else {
					valor_size += linecomand[contador]
					contador++
				}
			}
			fmt.Println("Valor: " + valor_size)

		} else if strings.Compare(lineacomando, "-fit") == 0 {
			fmt.Println("Encontro: " + lineacomando)
			lineacomando = ""
			flag_fit = true
			//simula un while
			for strings.Compare(linecomand[contador], "*") != 0 {
				if strings.Compare(linecomand[contador], " ") == 0 || strings.Compare(linecomand[contador], "\n") == 0 {
					contador++
					break
				} else {
					valor_fit += strings.ToLower(linecomand[contador])
					contador++
				}
			}
			fmt.Println("Valor: " + valor_fit)

		} else if strings.Compare(lineacomando, "-unit=") == 0 {
			fmt.Println("Encontro: " + lineacomando)
			lineacomando = ""
			flag_unit = true
			//directo
			valor_unit = strings.ToLower(linecomand[contador])
			contador++
			fmt.Println("Valor: " + valor_unit)

		} else if strings.Compare(lineacomando, "-path=") == 0 {
			fmt.Println("Encontro: " + lineacomando)
			lineacomando = ""
			flag_path = true
			//simula un while
			for strings.Compare(linecomand[contador], "*") != 0 {
				if strings.Compare(linecomand[contador], "\"") == 0 { // si viene con comilla doble
					contador++
					//simula un while
					for strings.Compare(linecomand[contador], "*") != 0 {
						if strings.Compare(linecomand[contador], "\"") == 0 { //finaliza path
							contador++
							break
						} else {
							valor_path += linecomand[contador]
							contador++
						}
					}
				} else {
					if strings.Compare(linecomand[contador], " ") == 0 || strings.Compare(linecomand[contador], "*") == 0 {
						contador++
						break
					} else {
						valor_path += linecomand[contador]
						contador++
					}
				}
			}
			fmt.Println("Valor : " + valor_path)

		}
	}
	//---------------PROCESO CREACION DE DISCOS-------------------------
	fmt.Println("Inicio de proceso")
	contadorDiagonal := 0
	for _, ele := range valor_path { // se cuenta cuantas diagonales hay para directorio
		if strings.Compare(string(ele), "/") == 0 {
			contadorDiagonal++
		}
	}

	valor_directorio := ""
	auxContador := 0
	for _, ele := range valor_path { // se obtiene solo directorio
		if strings.Compare(string(ele), "/") == 0 {
			valor_directorio += string(ele)
			auxContador++
			if contadorDiagonal == auxContador {
				break
			}
		}
	}

	flag_directorio := validacionDirectorio(valor_directorio) // funcion que valida si existe el directorio
	flag_disco := validacionArchivo(valor_path)               //funcion que valida si existe el disco

	//validacion directorio
	if flag_directorio == true { // existe el directorio
		fmt.Println("¡Existe Directorio!")
		if flag_disco == true {
			fmt.Println("Error -> ¡El disco ya existe con ese nombre!")
		} else {
			fmt.Println("¡El disco no existe, se procede con la creación!")
			crearDisco(flag_size, flag_unit, flag_path, flag_fit, valor_size, valor_path, valor_unit, valor_fit)
		}
	} else { // no existe el directorio
		fmt.Println("¡Directorio no existe!")
		merr := os.Mkdir(valor_directorio, 0755)
		if merr != nil {
			log.Fatal(merr)
		}
		fmt.Println("¡Directorio creado exitosamente!")
		crearDisco(flag_size, flag_unit, flag_path, flag_fit, valor_size, valor_path, valor_unit, valor_fit)
	}

}

func validacionDirectorio(directorio string) bool {
	if _, err := os.Stat(directorio); !os.IsNotExist(err) {
		return true
	} else {
		err = os.Mkdir(directorio, 0755)
		if err != nil {
			return false
			log.Fatal(err)
		}
		return true
	}

}

func validacionArchivo(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false

}

func crearDisco(_flag_size bool, _flag_unit bool, _flag_path bool, _flag_fit bool, size string, path string, unit string, fit string) {
	disco := mbr{} // se crea la estructura de mbr
	var tamano int64

	if _flag_size == true {
		_tamano, err := strconv.Atoi(size)
		if err != nil {
			log.Fatal(err)
		}
		if _tamano > 0 {
			fmt.Println("¡Tamaño disco valido!")
			tamano = int64(_tamano)
		} else {
			fmt.Println("Error -> ¡Tamaño de disco no valido!")
		}
	}

	if _flag_fit == true {
		if strings.Compare(fit, "bf") == 0 {
			copy(disco.dsk_fit[:], fit)
		} else if strings.Compare(fit, "ff") == 0 {
			copy(disco.dsk_fit[:], fit)
		} else if strings.Compare(fit, "wf") == 0 {
			copy(disco.dsk_fit[:], fit)
		} else {
			fmt.Println("Error -> ¡Valor invalido de fit!")
		}
	} else {
		copy(disco.dsk_fit[:], "ff") // si no es especificado es ff
	}

	if _flag_unit == true {
		if strings.Compare(unit, "k") == 0 { // si es kilobytes
			disco.mbr_tamano = tamano * 1024
		} else if strings.Compare(unit, "m") == 0 { // si es megabytes
			disco.mbr_tamano = tamano * 1024 * 1024
		} else {
			fmt.Println("Error -> ¡Valor de unit invalido!")
		}
	} else {
		disco.mbr_tamano = tamano * 1024 * 1024 // si no es especificado es megabytes
	}

	if _flag_path == true {
		disco.mbr_dsk_signature = int64(rand.Intn(100))
		disco.mbr_tamano = tamano

		//iniciando valores de particion
		for i := 0; i < 4; i++ {
			disco.partition[i].part_status = '0'
			disco.partition[i].part_type = '-'
			disco.partition[i].part_start = 0
			disco.partition[i].part_size = 0

		}
		file, err := os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
		var vacio int8 = 0
		var binario bytes.Buffer
		binary.Write(&binario, binary.BigEndian, vacio)

		file.Seek(disco.mbr_tamano, 0)
		binary.Write(&binario, binary.BigEndian, vacio)
		fecha := time.Now()
		fmt.Println(fecha)
		file.Close()
	}

}
