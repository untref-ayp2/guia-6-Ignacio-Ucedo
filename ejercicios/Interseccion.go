package ejercicios

import (
	"guia6/dictionary"
	"guia6/linkedlist"
)

// toma 2 slices y devuelve una nueva lista que es el resultado de la interseccion entre los mismos.
// esto es los strings que están en s1 como en s2
func Interseccion(s1 []string, s2 []string) linkedlist.LinkedList[string] {
	d := dictionary.NewDictionary[string, int]()

	s3 := append(s1, s2...)
	list := linkedlist.NewLinkedList[string]()

	for i := 0; i < len(s3); i++ {
		d.Put(s3[i], d.Get(s3[i])+1)

		if d.Get(s3[i]) > 1 && !list.Contains(s3[i]) {
			list.InsertAt(s3[i], list.Size()) // O(1)
		}
	}

	return *list

}
