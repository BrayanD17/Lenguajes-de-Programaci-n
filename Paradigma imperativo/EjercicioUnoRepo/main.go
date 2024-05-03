package main

import (
	"fmt"
)

/*
1)	Cree una lista de clientes con información previamente insertada desde la función main para clientes de un
determinado negocio. Cada cliente debe contener la siguiente información:

	type infoCliente struct {
	   nombre string
	   correo string
	   edad   int32
	}

Se debe crear una función “agregarCliente” con la cual hacer las siguientes 10 inserciones de clientes en la lista:

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

	// Imprimir la lista de clientes
	fmt.Println("Lista de clientes:")
	for _, cliente := range listaClientes {
		fmt.Printf("Nombre: %s, Correo: %s, Edad: %d\n", cliente.nombre, cliente.correo, cliente.edad)
	}
}
