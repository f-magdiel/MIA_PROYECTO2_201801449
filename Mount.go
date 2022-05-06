package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"unsafe"
)

var arraydisk [99]Disco //pendiente de montar y crear arreglos
var abecedario [26]string = [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
var numeros [50]string = [50]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50"}
var contadorDot = 0

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
		} else if strings.Compare(lineacomandos, "-path=") == 0 {
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
		fmt.Println("Se procede a montar....")
		flag_montaje := true

		// se monta pero anteriormente se habia montado, ahora solo es de activar con una S
		for i := 0; i < 99; i++ {
			for j := 0; j < 4; j++ {
				if arraydisk[i].size != 0 {
					if strings.Compare(arraydisk[i].Part[j].path, valor_path) == 0 && strings.Compare(arraydisk[i].Part[j].name, valor_name) == 0 {
						arraydisk[i].Part[j].mostrar = "S"
						flag_montaje = false
						break
					} else {
						for k := 0; k < 24; k++ {
							if strings.Compare(arraydisk[i].Logic[k].path, valor_path) == 0 && strings.Compare(arraydisk[i].Logic[k].name, valor_name) == 0 {
								arraydisk[i].Logic[k].mostrar = "S"
								flag_montaje = false
								break
							}
						}

					}
				}
			}
		}
		if flag_montaje == true { // se monta literal el disco por completo
			Montaje(valor_name, valor_path)
			//Solo para activar
			for i := 0; i < 99; i++ {
				for j := 0; j < 4; j++ {
					if arraydisk[i].size != 0 {
						if strings.Compare(arraydisk[i].Part[j].path, valor_path) == 0 && strings.Compare(arraydisk[i].Part[j].name, valor_name) == 0 {
							arraydisk[i].Part[j].mostrar = "S"

							break
						} else {
							for k := 0; k < 24; k++ {
								if strings.Compare(arraydisk[i].Logic[k].path, valor_path) == 0 && strings.Compare(arraydisk[i].Logic[k].name, valor_name) == 0 {
									arraydisk[i].Logic[k].mostrar = "S"

									break
								}
							}

						}
					}
				}
			}
		}

	} else {
		fmt.Println("Error -> Los parametros no son validos")
	}
}

func Montaje(name string, path string) {
	disk_available := validacionDiscoExiste(path)
	existenciaParticion := false

	if disk_available == true {
		fmt.Println("El disco si existe se procede a montar")

		//leer el disco con el path indicado
		mbr := Mbr{}
		var size_mbr int64 = int64(unsafe.Sizeof(mbr))
		file, err := os.Open(path)
		if err != nil {
			log.Fatal("Error al abrir disco", err)
		} else {
			file.Seek(0, 0)
			data := leerSiguienteBytesMount(file, size_mbr)
			buffer := bytes.NewBuffer(data)
			err = binary.Read(buffer, binary.BigEndian, &mbr)
			if err != nil {
				log.Fatal("Error al leer disco", err)
			}
		}
		//valida si se encuentra la particion en el mbr
		for i := 0; i < 4; i++ {
			if strings.Compare(string(mbr.Partition[i].Part_name[:len(name)]), name) == 0 {
				existenciaParticion = true
				break
			}
		}

		//para montar el disco en el array
		if existenciaParticion == true {
			fmt.Println("Existe la particion se procede a montar")
			// se busca espacio para agregar el disco
			for i := 0; i < 99; i++ {
				if arraydisk[i].size == int64(0) { //si halla espacio
					//asigo disco
					var auxcont int64 = 0
					arraydisk[i].path = path
					if arraydisk[i].id > 0 {
						auxcont = arraydisk[i-1].id //consulto el anterior
					} else {
						auxcont = arraydisk[i].id //es primero
						auxcont++
						arraydisk[i].id = auxcont
					}

					//arraydisk[i].id = auxcont
					arraydisk[i].size = mbr.Mbr_tamano

					//se valida que particiones están activas
					for j := 0; j < 4; j++ {
						//es particion primaria y está activa
						if strings.Compare(string(mbr.Partition[j].Part_status), "1") == 0 && strings.Compare(string(mbr.Partition[i].Part_type), "p") == 0 {
							//para buscar espacio en ram
							for l := 0; l < 4; l++ {
								if arraydisk[i].Part[l].start == 0 && arraydisk[i].Part[l].size == 0 { // esta vacio
									//generador id
									auxid := "49"
									auxid += numeros[i]
									auxid += abecedario[l]
									fmt.Println("ID generador " + auxid)
									//asigna al array -> ram
									arraydisk[i].Part[l].id = auxid
									auxid = ""
									arraydisk[i].Part[l].name = string(mbr.Partition[j].Part_name[:len(name)])
									arraydisk[i].Part[l].path = path
									arraydisk[i].Part[l].size = mbr.Partition[j].Part_size
									arraydisk[i].Part[l].start = mbr.Partition[j].Part_start
									arraydisk[i].Part[l].tipo = "p"
									arraydisk[i].Part[l].mostrar = "N"
									//file.Close()
									break
								}
							}
							//break
						} else if strings.Compare(string(mbr.Partition[j].Part_status), "1") == 0 && strings.Compare(string(mbr.Partition[j].Part_type), "e") == 0 {
							// es una extendida que puede contener o no logicas
							//busco espacio en la ram
							for l := 0; l < 4; l++ {
								if arraydisk[i].Part[l].start == 0 && arraydisk[i].Part[l].size == 0 {
									//generador id
									auxid := "49"
									auxid += numeros[i]
									auxid += abecedario[l]
									fmt.Println("Id generador ext " + auxid)
									//se asigna el disco
									arraydisk[i].Part[l].name = name
									arraydisk[i].Part[l].id = auxid
									arraydisk[i].Part[l].path = path
									arraydisk[i].Part[l].size = mbr.Partition[j].Part_size
									arraydisk[i].Part[l].start = mbr.Partition[j].Part_start
									arraydisk[i].Part[l].tipo = "e"
									arraydisk[i].Part[l].mostrar = "N"
									auxid = ""
									//se buscan las logicas
									ebr := Ebr{}
									sig := 0
									file.Seek(mbr.Partition[l].Part_start+int64(unsafe.Sizeof(ebr))+1, 0)
									data := leerSiguienteBytesMount(file, int64(unsafe.Sizeof(ebr)))
									buffer := bytes.NewBuffer(data)
									err := binary.Read(buffer, binary.BigEndian, &ebr)
									if err != nil {
										log.Fatal("Error al leer el ebr ", err)
									}
									sig = int(ebr.Part_next)
									//simula un while
									for sig != -1 {
										//se genera id
										/* auxid := "49"
										auxid += string(numeros[i])
										auxid += string(abecedario[l]) */
										//se asigna en la ram la particion
										for k := 0; k < 24; k++ {
											//busco espacio para la logica
											if arraydisk[i].Logic[k].size == 0 {
												//arraydisk[i].Logic[k].id = auxid //creo que es 24
												arraydisk[i].Logic[k].name = name
												arraydisk[i].Logic[k].path = path
												arraydisk[i].Logic[k].size = ebr.Part_size
												arraydisk[i].Logic[k].start = ebr.Part_start
												arraydisk[i].Logic[k].tipo = "l"
												arraydisk[i].Logic[k].mostrar = "N"
												break
											}

										}
										sig = int(ebr.Part_next)
										//se vuelve a leer
										file.Seek(int64(sig), 0)
										data := leerSiguienteBytesMount(file, int64(unsafe.Sizeof(ebr)))
										buff := bytes.NewBuffer(data)
										err := binary.Read(buff, binary.BigEndian, &ebr)
										if err != nil {
											log.Fatal("Error al leer ebr ", err)
										}

										sig = int(ebr.Part_next)
									}
									//file.Close()
									break
								}
							}
						}
					}

					break
				}

			}
			file.Close()
		} else {
			fmt.Println("Error -> Particion no se encuentra ")
			//file.Close() //cerrado
			//disco leido
		}
	} else {
		fmt.Println("Error -> El disco no existe para montar")
	}
}

func validacionDiscoExiste(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false

}

func MonstrarMount() {
	fmt.Println("======================= MONTAJE =======================")
	fmt.Println("PATH	|	  NOMBRE	|	ID		 |  TIPO")
	for i := 0; i < 99; i++ {
		if arraydisk[i].size != 0 {
			for j := 0; j < 4; j++ {
				if strings.Compare(arraydisk[i].Part[j].tipo, "p") == 0 { // es de tipo primaria
					if strings.Compare(arraydisk[i].Part[j].mostrar, "S") == 0 {
						fmt.Println(arraydisk[i].Part[j].path + "    |   " + arraydisk[i].Part[j].name + "   |   " + "#" + arraydisk[i].Part[j].id + "   |   " + arraydisk[i].Part[j].tipo)
					}

				} else if strings.Compare(arraydisk[i].Part[j].tipo, "e") == 0 { //es de tipo extendida
					//se buscarn en las logicas
					if strings.Compare(arraydisk[i].Part[j].mostrar, "S") == 0 {
						fmt.Println(arraydisk[i].Part[j].path + "    |   " + arraydisk[i].Part[j].name + "   |    " + "#" + arraydisk[i].Part[j].id + "   |   " + arraydisk[i].Part[j].tipo)
					}

					/* for n := 0; n < 24; n++ {
						if strings.Compare(arraydisk[i].Logic[n].tipo, "l") == 0 {
							fmt.Println(arraydisk[i].Logic[n].path + "    |   " + arraydisk[i].Logic[n].name + "    |   " + "#" + arraydisk[i].Logic[n].id + "   |   " + arraydisk[i].Logic[n].tipo)
						}
					} */
				}
			}
		}
	}
}
func leerSiguienteBytesMount(file *os.File, number int64) []byte {
	bytes := make([]byte, number)
	_, err := file.Read(bytes)
	if err != nil {

		log.Fatal("Error al abrir byte en mount ", err)
	}
	return bytes
}

///REPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPP

func AnalsisRep(comando string) {
	contador := 0
	var linecomand [100]string
	newcomando := strings.Split(comando, "")
	lineacomandos := ""

	copy(linecomand[:], newcomando[:])

	flag_path := false
	flag_name := false
	flag_id := false

	valor_path := ""
	valor_name := ""
	valor_id := ""

	//while
	for strings.Compare(linecomand[contador], "\n") != 0 && strings.Compare(linecomand[contador], "#") != 0 {

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

		//validaciaoens comandos y valores
		if strings.Compare(lineacomandos, "rep") == 0 {
			fmt.Println("Encontro : " + lineacomandos)
			lineacomandos = ""
			contador++
		} else if strings.Compare(lineacomandos, "-id=") == 0 {
			fmt.Println("Encontro : " + lineacomandos)
			lineacomandos = ""
			flag_id = true

			for strings.Compare(linecomand[contador], "\n") != 0 {
				if strings.Compare(linecomand[contador], " ") == 0 || strings.Compare(linecomand[contador], "\n") == 0 {
					contador++
					break
				} else {
					valor_id += linecomand[contador]
					contador++
				}

			}
			fmt.Println("Valor : " + valor_id)
		} else if strings.Compare(lineacomandos, "-path=") == 0 {
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
			fmt.Println("Valor : " + valor_path)
		} else if strings.Compare(lineacomandos, "-name=") == 0 {
			fmt.Println("Encontro : " + lineacomandos)
			lineacomandos = ""
			flag_name = true

			for strings.Compare(linecomand[contador], "\n") != 0 {
				if strings.Compare(linecomand[contador], " ") == 0 || strings.Compare(linecomand[contador], "\n") == 0 {
					contador++
					break
				} else {
					valor_name += linecomand[contador]
					contador++
				}
			}
			fmt.Println("Valor : " + valor_name)
		}
	}
	//se manda para generar reporte
	if strings.Compare(valor_name, "disk") == 0 {
		fmt.Println("Se genera reporte Disk")
		generaReporte(flag_name, flag_id, flag_path, valor_name, valor_id, valor_path)
	}

}

func generaReporte(f_name bool, f_id bool, f_path bool, _name string, _id string, _path string) {
	fmt.Println("Se genera reporte disk......")
	fmt.Println("ID ingrsado-------" + _id)
	var contenido string
	suma := 0
	//llenaod contenido
	contenido += "digraph {\n"
	contenido += "tbl [\n"
	contenido += "shape=plaintext\n"
	contenido += "label=<\n"
	contenido += "<table border='2' cellborder='0' color='blue' cellspacing='1'>\n"
	contenido += "<tr>\n"
	contenido += "<td colspan='1' rowspan='1'>\n"
	contenido += "<table color='orange' border='1' cellborder='1' cellpadding='10' cellspacing='0'>\n"
	contenido += "<tr><td>MBR</td></tr>\n"
	contenido += "</table>\n"
	contenido += "</td>\n"

	//busco el id primero
	for i := 0; i < 99; i++ {
		if arraydisk[i].size != 0 {
			for j := 0; j < 4; j++ {
				if strings.Compare(arraydisk[i].Part[j].id, _id) == 0 { // si encuentro el id
					fmt.Println("ID = " + arraydisk[i].Part[j].id)
					//que tipo de particiones tiene el disco
					size_totaldisk := arraydisk[i].size
					size_p1 := 0
					size_p2 := 0

					for k := 0; k < 4; k++ {
						fmt.Println(i)
						fmt.Println(k)
						if strings.Compare(arraydisk[i].Part[k].tipo, "p") == 0 {
							fmt.Println("Se genera primaria")
							//contenido se llena
							contenido += "<td colspan='1' rowspan='1'>\n"
							contenido += "<table color='orange' border='1' cellborder='1' cellpadding='10' cellspacing='0'>\n"
							contenido += "<tr><td>Primaria <br/> "

							size_p1 = int(arraydisk[i].Part[k].size)
							op := (size_p1 * 100) / int(size_totaldisk)
							res := float64(op)
							res = math.Round(res)
							suma += size_p1

							contenido += strconv.FormatFloat(res, 'f', 2, 64)
							contenido += " % del disco</td></tr>\n"
							contenido += "</table>\n"
							contenido += "</td>\n"
							size_p1 = 0
						} else if strings.Compare(arraydisk[i].Part[k].tipo, "e") == 0 {
							fmt.Println("Se genera enxtendida")
							//conteido
							contenido += "<td colspan='1' rowspan='1'>\n"
							contenido += "<table color='red' border='1' cellborder='1' cellpadding='10' cellspacing='0'>\n"
							contenido += "<tr> "

							var reslog int64
							size_p2 = int(arraydisk[i].Part[k].size)
							suma += size_p2
							var sizebr int64

							for l := 0; l < 24; l++ {
								if arraydisk[i].Logic[l].size != 0 {
									contenido += " <td>EBR</td><td> Logica <br/> "
									sizebr = arraydisk[i].Logic[l].size

									reslog += int64(sizebr)
									//res := float64((sizebr * 100) / int64(size_p2))
									op1 := float64(sizebr)
									op2 := float32(size_totaldisk)
									op3 := op1 * 100 / float64(op2)

									contenido += strconv.FormatFloat(op3, 'f', 2, 64)
									contenido += " % del disco </td> "
								}
							}
							fmt.Println("tamreslog ---- tamebr")
							fmt.Println(reslog)
							fmt.Println(size_p2)
							if reslog < int64(size_p2) {
								var por int64 = int64(size_p2) - int64(reslog)
								var resul int64 = por * 100 / int64(size_totaldisk)
								res := float64(resul)
								res = math.Round(res)
								contenido += "<td> Libre Logica <br/> "
								contenido += strconv.FormatFloat(res, 'f', 2, 64)
								contenido += " % del disco</td>"
							}
							contenido += "</tr>\n"
							contenido += "</table>\n"
							contenido += "</td>\n"

							size_p2 = 0
						} else { //es libre
							fmt.Println("Se genera libre")
							if suma < int(size_totaldisk) {
								var resfree int64 = size_totaldisk - int64(suma)
								resfree = resfree * 100 / size_totaldisk
								res := float64(resfree)
								res = math.Round(res)
								contenido += "<td colspan='1' rowspan='1'>\n"
								contenido += "<table color='orange' border='1' cellborder='1' cellpadding='10' cellspacing='0'>\n"
								contenido += "<tr><td>Libre <br/> "
								contenido += strconv.FormatFloat(res, 'f', 2, 64)
								contenido += "% del disco</td></tr>\n"
								contenido += "</table>\n"
								contenido += "</td>\n"
							}
							break
						}
					}
					//break
				}

			}
		} else {
			break
		}

	}
	contenido += " </tr>\n"
	contenido += " </table>\n"
	contenido += ">];\n"
	contenido += "}\n"
	fmt.Println("Grafica.......................")
	//fmt.Println(contenido)
	fmt.Println("Fin Grafica......................")
	//Se valida el directorio para guardar
	cont_Diagonal := 0
	for _, ele := range _path {
		if strings.Compare(string(ele), "/") == 0 {
			cont_Diagonal++
		}
	}

	nuevo_directorio := ""
	auxconta := 0
	for _, ele := range _path {
		if strings.Compare(string(ele), "/") == 0 {
			nuevo_directorio += string(ele)
			auxconta++
			if cont_Diagonal == auxconta {
				break
			}
		} else {
			nuevo_directorio += string(ele)
		}
	}
	fmt.Println("Direcotoro a crear=================" + nuevo_directorio)
	flag_bandera := validacionGeneradorDirectorio(nuevo_directorio)
	fmt.Println("Banderaaaaaaa")
	fmt.Println(flag_bandera)
	if flag_bandera == true { //ya existe
		//solo se genera el reporte
		contadorDot++
		auxdot := strconv.Itoa(contadorDot)
		f, err := os.Create("reporte" + auxdot + ".dot")
		if err != nil {
			fmt.Println(err)
			return
		}
		l, err := f.WriteString(contenido)
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
		fmt.Println(l, "bytes written successfully")
		err = f.Close()

		//para renderizar el .dot
		com, _ := exec.LookPath("dot")
		cmd, _ := exec.Command(com, "-Tjpg", "reporte"+auxdot+".dot").Output()
		mode := int(0777)
		ioutil.WriteFile(_path, cmd, os.FileMode(mode))

	} else { // no existe el directorio
		fmt.Println("Se crea el directorio")
		crearDirectorioRep(nuevo_directorio) // si se pudo
		//se genera reporte
		contadorDot++
		auxdot := strconv.Itoa(contadorDot)
		f, err := os.Create("reporte" + auxdot + ".dot")
		if err != nil {
			fmt.Println(err)
			return
		}
		l, err := f.WriteString(contenido)
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
		fmt.Println(l, "bytes written successfully")
		err = f.Close()

		//para renderizar el .dot
		com, _ := exec.LookPath("dot")
		cmd, _ := exec.Command(com, "-Tjpg", "reporte"+auxdot+".dot").Output()
		mode := int(0777)
		ioutil.WriteFile(_path, cmd, os.FileMode(mode))
	}
}

func validacionGeneradorDirectorio(directorio string) bool {
	if _, err := os.Stat(directorio); !os.IsNotExist(err) {
		return true
	} else {
		return false
	}
}

func crearDirectorioRep(directorio string) { //error al crear el directori

	errx := os.Mkdir(directorio, 0755)
	if errx != nil {
		fmt.Println("Error al crear directorio rep " + directorio)
		log.Fatal(errx)
	}
}
