package main

import (
	"fmt"
	"reflect"
	"strings"
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
4)	Utilizando exclusivamente las funciones map y filter genéricas hechas en clase, realice la función
llamada “clientesSugerenciasCorreos” que devuelve una nueva lista con todos los clientes originales, pero para
aquellos clientes cuyas direcciones de correo no hacen referencia al nombre de la persona, cambiarlos por otro
que usted sugiera que sí contemple el nombre de la persona(formato libre propuesto por cada estudiante).
Se debe probar la función llamándola en el main.
*/


func clientesSugerenciasCorreos(clientes []infoCliente) []infoCliente {
	// Identificar clientes cuyos correos no contienen el apellido en el correo
	clientesSinApellido := filter3(clientes, func(cliente any) bool {
		if c, ok := cliente.(infoCliente); ok {
			return !strings.Contains(c.correo, strings.ToLower(strings.Split(c.nombre, " ")[len(strings.Split(c.nombre, " "))-1]))
		}
		return false
	})

	// Generar sugerencias de correo para clientes sin apellido en el correo
	clientesConSugerencia := map3(clientesSinApellido, func(cliente any) any {
		if c, ok := cliente.(infoCliente); ok {
			nombreParts := strings.Split(c.nombre, " ")
			sugerencia := fmt.Sprintf("%s.%s@gmail.com", strings.ToLower(nombreParts[0]), strings.ToLower(nombreParts[len(nombreParts)-1]))
			return infoCliente{
				nombre: c.nombre,
				correo: sugerencia,
				edad:   c.edad,
			}
		}
		return cliente
	})

	// Convertir a []infoCliente
	var resultado []infoCliente
	for _, cliente := range clientesConSugerencia {
		if c, ok := cliente.(infoCliente); ok {
			resultado = append(resultado, c)
		}
	}
	return resultado
}

// Función map3
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

// Función filter3
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

	// Aplicar la función clientesSugerenciasCorreos
	clientesConSugerencias := clientesSugerenciasCorreos(listaClientes)

	// Imprimir la lista de clientes con sugerencias de correos
	fmt.Println("Lista de clientes con sugerencias de correos:")
	for _, cliente := range clientesConSugerencias {
		fmt.Printf("Nombre: %s, Correo: %s, Edad: %d\n", cliente.nombre, cliente.correo, cliente.edad)
	}
}
