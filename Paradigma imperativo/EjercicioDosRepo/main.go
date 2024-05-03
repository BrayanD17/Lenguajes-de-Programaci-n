package main

import (
	"fmt"
	"reflect"
	"strings"
)

/*
2)	Utilizando exclusivamente la función filter genérica hecha en clase y
punteros para referenciar la listas de entrada, realice la función llamada
“listaClientes_ApellidoEnCorreo” que devuelve la lista de Clientes cuyos correos contengan el
apellido de la persona. Se debe probar la función llamándola en el main.
*/
type infoCliente struct {
	nombre string
	correo string
	edad   int32
}

type listaClientes []infoCliente

func (lc *listaClientes) agregarCliente(nombre, correo string, edad int32) {
	*lc = append(*lc, infoCliente{nombre: nombre, correo: correo, edad: edad})
}
func filter3(list interface{}, f func(interface{}) bool) []interface{} {
	result := make([]interface{}, 0)
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

func (lc listaClientes) listaClientes_ApellidoEnCorreo() listaClientes {
	result := filter3(lc, func(c interface{}) bool {
		cliente := c.(infoCliente)
		apellidos := obtenerApellidos(cliente.nombre)
		for _, apellido := range apellidos {
			if contieneApellidoEnCorreo(cliente.correo, apellido) {
				return true
			}
		}
		return false
	})

	clientesFiltrados := make(listaClientes, len(result))
	for i, c := range result {
		clientesFiltrados[i] = c.(infoCliente)
	}
	return clientesFiltrados
}

func obtenerApellidos(nombre string) []string {
	// Dividir el nombre en palabras separadas por espacio
	palabras := strings.Fields(nombre)
	if len(palabras) > 1 {
		// El último elemento es considerado como el apellido
		return palabras[1:]
	}
	return nil
}

func contieneApellidoEnCorreo(correo string, apellido string) bool {
	// Convertir el apellido a minúscula
	apellido = strings.ToLower(apellido)

	// Verificar si el apellido está presente en el correo electrónico
	return strings.Contains(strings.ToLower(correo), apellido)
}

func main() {
	var listaClientes listaClientes
	// Agregar clientes a la lista usando la funcion agregarCliente
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

	clientesFiltrados := listaClientes.listaClientes_ApellidoEnCorreo()

	fmt.Println("Clientes cuyo correo contiene su apellido:")
	for _, cliente := range clientesFiltrados {
		fmt.Printf("Nombre: %s, Correo: %s, Edad: %d\n", cliente.nombre, cliente.correo, cliente.edad)
	}
}
