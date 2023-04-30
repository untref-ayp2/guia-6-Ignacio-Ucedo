package ejercicios

import (
	"guia6/dictionary"
)

func InformacionSolicitada(entrada dictionary.Dictionary[string, []string]) *dictionary.Dictionary[string, []string] {

	salida := dictionary.NewDictionary[string, []string]()
	llaves := entrada.GetKeys()
	var arreglo []string

	for i := 0; i < entrada.Size(); i++ {
		arreglo = append(arreglo, llaves[i]) // agrega la la primera fecha a un arreglo

		for j := 0; j < len(entrada.Get(llaves[i])); j++ { // itero los alumnos que fueron ese dia

			if !salida.Contains(entrada.Get(llaves[i])[j]) { // si ese alumno todavia no aparecio en el nuevo informe

				salida.Put(entrada.Get(llaves[i])[j], arreglo) // lo pongo como llave en el informe
				// y le asigno la fecha que guarde en el arreglo
			} else { // si ya aparece en el informe (esta no es la primera iteración en i) ,
				// le tengo que guardar otra fecha más. Para eso le agrego a arreglo ( que tiene la fecha
				// que quiero agregarle al alumno), la fecha que ya tiene alumno
				// y restauro el nuevo informe en la llave del alumno con las fechas
				// actualizadas
				arreglo = append(arreglo, salida.Get(entrada.Get(llaves[i])[j])[j])
				salida.Put(entrada.Get(llaves[i])[j], arreglo)

			}
		}
		arreglo = make([]string, 0) // tras guardar una fecha limpio el arreglo para pasar a la siguiente fecha
		// con un arreglo vacío
	}

	return &salida // devuelvo el puntero de salida
}
