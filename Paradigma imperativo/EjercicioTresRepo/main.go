package main

import (
	"fmt"
	"reflect"
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
3)	Utilizando exclusivamente las funciones genéricas map/filter y la función reduce así como punteros para
referenciar la listas de entrada, realice la función llamada “cantidadCorreosCostaRica” que devuelve la
cantidad de clientes cuyos correos pertenecen a dominios de Costa Rica (terminados en “.cr”).
Se debe probar la función llamándola en el main.*/

// Map aplica la función f a cada elemento de la lista y devuelve una nueva lista con los resultados
func map3(list any, f func(any) any) []any {
	result := make([]any, 0)
	switch reflect.TypeOf(list).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(list)
		for i := 0; i < s.Len(); i++ {
			result = append(result, f(s.Index(i).Interface()))
		}
	}
	return result
}

// Filter devuelve una nueva lista que contiene solo los elementos de la lista original para los cuales la función f devuelve verdadero
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

// Reduce aplica una función f que combina todos los elementos de la lista en un solo valor
func reduce(list any, f func(any, any) any, initial any) any {
	switch reflect.TypeOf(list).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(list)
		result := initial
		for i := 0; i < s.Len(); i++ {
			result = f(result, s.Index(i).Interface())
		}
		return result
	default:
		return initial
	}
}

// Función para verificar si el correo termina en ".cr"
func esCorreoCR(email any) bool {
	if correo, ok := email.(string); ok {
		return len(correo) >= 3 && correo[len(correo)-3:] == ".cr"
	}
	return false
}

// Función para contar la cantidad de correos de Costa Rica
func cantidadCorreosCostaRica(clientes []infoCliente) int {
	// Filtrar los correos que terminan en ".cr"
	correosCR := filter3(clientes, func(cliente any) bool {
		if c, ok := cliente.(infoCliente); ok {
			return esCorreoCR(c.correo)
		}
		return false
	})

	// Mapear los correos filtrados a 1 (representando un cliente que cumple con el criterio)
	mapeado := map3(correosCR, func(any) any {
		return 1
	})

	// Reducir la lista mapeada sumando los valores
	total := reduce(mapeado, func(acc, val any) any {
		if a, ok := acc.(int); ok {
			if v, ok := val.(int); ok {
				return a + v
			}
		}
		return acc
	}, 0)

	return total.(int)
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

	cantidad := cantidadCorreosCostaRica(listaClientes)
	fmt.Printf("Cantidad de clientes con correos de Costa Rica: %d\n", cantidad)
}
