package main

import (
	"fmt"
	"reflect"
	"sort"
)

type infoCliente struct {
	nombre string
	correo string
	edad   int32
}

type listaClientes []infoCliente

func (lc *listaClientes) agregarCliente(nombre, correo string, edad int32) {
	*lc = append(*lc, infoCliente{nombre: nombre, correo: correo, edad: edad})
}

/*
5)	En el formato que el estudiante prefiera, escriba una función llamada “correosOrdenadosAlfabeticos”
que devuelva una lista nueva con todos los correos de clientes ordenados alfabéticamente.
Se debe probar la función llamándola en el main.
*/
// Función para obtener los correos ordenados alfabéticamente
func correosOrdenadosAlfabeticos(clientes listaClientes) []string {
	// Función para filtrar los correos electrónicos de la lista de clientes
	filtrarCorreos := func(any) bool {
		return true
	}

	// Filtrar los correos electrónicos de la lista de clientes
	correos := filter3(clientes, filtrarCorreos)

	// Extraer los correos electrónicos de la lista filtrada
	var correosStrings []string
	for _, cliente := range correos {
		correosStrings = append(correosStrings, cliente.(infoCliente).correo)
	}

	// Ordenar alfabéticamente los correos electrónicos
	sort.Strings(correosStrings)

	return correosStrings
}

// Función de filtro genérico
func filter3(list any, f func(any) bool) []any {
	result := make([]any, 0)
	switch reflect.TypeOf(list).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(list)
		for i := 0; i < s.Len(); i++ {
			if f(s.Index(i).Interface()) {
				result = append(result, s.Index(i).Interface())
			}
		}
	}
	return result
}

func main() {
	var listaClientes listaClientes
	// Agregar clientes a la lista usando la función agregarCliente
	listaClientes.agregarCliente("Oscar Viquez", "oviquez@tec.ac.cr", 44)
	listaClientes.agregarCliente("Pedro Perez", "elsegundo@gmail.com", 30)
	listaClientes.agregarCliente("Maria Lopez", "mlopez@hotmail.com", 18)
	listaClientes.agregarCliente("Juan Rodriguez", "jrodriguez@gmail.com", 25)
	listaClientes.agregarCliente("Luisa Gonzalez", "muyseguro@tec.ac.cr", 67)
	listaClientes.agregarCliente("Marco Rojas", "loquesea@hotmail.com", 47)
	listaClientes.agregarCliente("Marta Saborio", "msaborio@gmail.com", 33)
	listaClientes.agregarCliente("Camila Segura", "csegura@ice.co.cr", 19)
	listaClientes.agregarCliente("Fernando Rojas", "frojas@estado.gov", 27)
	listaClientes.agregarCliente("Rosa Ramirez", "risuenna@gmail.com", 50)

	// Obtener la lista de correos ordenados alfabéticamente
	correosOrdenados := correosOrdenadosAlfabeticos(listaClientes)

	// Imprimir la lista de correos ordenados
	fmt.Println("Correos ordenados alfabéticamente:")
	for _, correo := range correosOrdenados {
		fmt.Println(correo)
	}

}
