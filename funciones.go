package main

import (
	//"fmt"
	"strconv"
)

const larTipo int = 5
const larLargo int = 2

func separaDatos(entrada []byte) (map[string]string, string) {

	mapa := make(map[string]string)
	var errores string = ""
	var datoEntrada = string(entrada)

	if datoEntrada == "" {
		errores = "No se introdujo el dato a leer"
		return mapa, errores
	}

	var largoTotal int = len(datoEntrada)

	tipo, largoNum, valor, errores := rescataValores(0, datoEntrada)

	if errores != "" {
		return mapa, errores
	}

	//imprimeLog(tipo, largoNum, valor, 0)
	errores = validaTipo(tipo, valor)
	if errores != "" {
		return mapa, errores
	}

	mapa[tipo[1:3]] = valor

	if (larTipo + largoNum) < largoTotal {
		var totalLargo int = larTipo + largoNum
		for i := 1; i < largoTotal; i++ {
			tipo, largoNum, valor, errores := rescataValores(totalLargo, datoEntrada)

			if errores != "" {
				return mapa, errores
			}

			//imprimeLog(tipo, largoNum, valor, i)

			errores = validaTipo(tipo, valor)
			if errores != "" {
				return mapa, errores
			}

			mapa[tipo[1:3]] = valor

			if (totalLargo + larTipo + largoNum) < largoTotal {
				totalLargo = totalLargo + larTipo + largoNum
			} else {
				i = largoTotal
			}
		}
	}

	return mapa, errores
}

func rescataValores(largoAcomulado int, datoEntrada string) (string, int, string, string) {
	var tipo string = ""
	var valor string = ""
	var errores string = ""

	largoNum, err := strconv.Atoi(datoEntrada[largoAcomulado : largoAcomulado+larLargo])

	if err != nil {
		errores = "Error en convertir el largo a numerico: " + datoEntrada[largoAcomulado:largoAcomulado+larLargo]
		return tipo, 0, valor, errores
	}

	tipo = datoEntrada[largoAcomulado+larLargo : largoAcomulado+larTipo]
	valor = datoEntrada[largoAcomulado+larTipo : largoAcomulado+larTipo+largoNum]

	return tipo, largoNum, valor, errores

}

func validaTipo(tipo string, valor string) string {
	tipoDato := tipo[:1]
	if tipoDato != "N" &&
		tipoDato != "n" &&
		tipoDato != "A" &&
		tipoDato != "a" {
		return "Tipo dato debe ser N o A, no " + string(tipoDato)
	}

	if tipoDato == "N" || tipoDato == "n" {
		_, err := strconv.Atoi(valor)
		if err != nil {
			return "Tipo de dato definido como Numerico y es Alfanumerico"
		}
	}
	return ""
}

/*func imprimeLog(tipo string, largo int, valor string, vuelta int) {
	fmt.Println("Tipo:", tipo)
	fmt.Println("Largo:", largo)
	fmt.Println("Valor:", valor)
	fmt.Println("Vuelta:", vuelta)
	fmt.Println("======================")
}*/
