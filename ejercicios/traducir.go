package ejercicios

import (
	"guia6/dictionary"
)

func Traducir(texto string, dic dictionary.Dictionary[string, string]) string {

	palabra := ""
	cadena := "{"

	for i := 0; i < len(texto); i++ {

		if ((texto)[i] >= 97 && (texto)[i] <= 122) ||
			((texto)[i] >= 65 && (texto)[i] <= 90) {

			palabra += string(texto[i])

			if i == len(texto)-1 {
				cadena += palabra + ": " + dic.Get(palabra) + ", "
			}

		} else {
			if dic.Contains(palabra) {
				cadena += palabra + ": " + dic.Get(palabra) + ", "
				palabra = ""
			} else {
				if len(palabra) > 0 {
					cadena += palabra + ": *nil, "
					palabra = ""
				}
			}

		}

	}

	cadena = cadena[:len(cadena)-2]
	cadena += "}"
	return cadena

}
