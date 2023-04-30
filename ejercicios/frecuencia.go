package ejercicios

import (
	"guia6/dictionary"
)

// recibe una cadena de caracteres y calcula la frecuencia de aparición de cada letra. Por ejemplo, en un main:
// s := "hola como estas"
// fmt.Print(ejercicios.Frecuencia(s).StringLn())
// debe imprimir {a: 2, c: 1, e: 1, h: 1, l: 1, m: 1, o: 3, s: 2, t: 1,}
// solo alfabeto inglés mayúsculas y minúsculas
func Frecuencia(s string) *dictionary.Dictionary[string, int] {
	//d *dictionary.Dictionary[string, int] esto está mal porque es un puntero a nil, falta una asignación
	d := dictionary.NewDictionary[string, int]()

	for i := 0; i < len(s); i++ {
		if ((s)[i] >= 97 && (s)[i] <= 122) ||
			((s)[i] >= 65 && (s)[i] <= 90) {

			d.Put(string(s[i]), d.Get(string(s[i]))+1)
		}
	}

	return &d //devuelve el valor asociado al puntero
}
