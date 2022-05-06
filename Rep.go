package main

/*
import (
	"fmt"
	"strconv"
	"strings"
)
*/
/* func AnalsisRep(comando string) {
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
		} else if strings.Compare(linecomand[contador], "-id=") == 0 {
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
	generaReporte(flag_name, flag_id, flag_path, valor_name, valor_id, valor_path)
}

func generaReporte(f_name bool, f_id bool, f_path bool, _name string, _id string, _path string) {
	fmt.Println("Se genera reporte disk......")
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
					//que tipo de particiones tiene el disco
					size_totaldisk := arraydisk[i].size
					size_p1 := 0
					size_p2 := 0

					//for k := 0; k < 4; k++ {
					fmt.Println(i)
					fmt.Println(j)
					if strings.Compare(arraydisk[i].Part[j].tipo, "p") == 0 {
						fmt.Println("Se genera primaria")
						//contenido se llena
						contenido += "<td colspan='1' rowspan='1'>\n"
						contenido += "<table color='orange' border='1' cellborder='1' cellpadding='10' cellspacing='0'>\n"
						contenido += "<tr><td>Primaria <br/> "

						size_p1 = int(arraydisk[i].Part[j].size)
						op := (size_p1 * 100) / int(size_totaldisk)
						suma += size_p1

						contenido += strconv.FormatInt(int64(op), 10)
						contenido += " % </td></tr>\n"
						contenido += "</table>\n"
						contenido += "</td>\n"
						size_p1 = 0
					} else if strings.Compare(arraydisk[i].Part[j].tipo, "e") == 0 {
						fmt.Println("Se genera primaria")
						//conteido
						contenido += "<td colspan='1' rowspan='1'>\n"
						contenido += "<table color='red' border='1' cellborder='1' cellpadding='10' cellspacing='0'>\n"
						contenido += "<tr> "

						var reslog int64
						size_p2 = int(arraydisk[i].Part[j].size)
						suma += size_p2
						var sizebr int64

						for l := 0; l < 24; l++ {
							contenido += " <td>EBR</td><td>Logica1 <br/> "
							if arraydisk[i].Logic[l].size != 0 {
								sizebr = arraydisk[i].Logic[l].size
								reslog += int64(sizebr)
								res := sizebr * 100 / int64(size_p2)
								contenido += strconv.FormatInt(res, 10)
								contenido += "% </td> "
							}
						}

						if reslog < sizebr {
							var por int64 = sizebr - int64(reslog)
							var resul int64 = por * 100 / int64(size_p2)
							contenido += "<td>Libre <br/> "
							contenido += strconv.FormatInt(resul, 10)
							contenido += " % </td>"
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
							contenido += "<td colspan='1' rowspan='1'>\n"
							contenido += "<table color='orange' border='1' cellborder='1' cellpadding='10' cellspacing='0'>\n"
							contenido += "<tr><td>Libre <br/> "
							contenido += strconv.FormatInt(resfree, 10)
							contenido += "% </td></tr>\n"
							contenido += "</table>\n"
							contenido += "</td>\n"
						}
						suma = 0
					}
					//}
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
	fmt.Println(contenido)
	fmt.Println("Fin Grafica......................")
}
*/
