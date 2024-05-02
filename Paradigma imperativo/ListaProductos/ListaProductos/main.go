package main

import (
	"errors"
	"fmt"
	"strings"
)

type producto struct {
	nombre   string
	cantidad int
	precio   int
}
type listaProductos []producto

var lProductos listaProductos

const existenciaMinima int = 10 //la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	// modificar el código para que cuando se agregue un producto, si este ya se encuentra, incrementar la cantidad
	// de elementos del producto y eventualmente el precio si es que es diferente
	//Listo
	var existe = false
	for i := 0; i < len(*l); i++ {
		if strings.EqualFold(nombre, (*l)[i].nombre) {
			(*l)[i].cantidad += cantidad
			if (*l)[i].precio != precio {
				(*l)[i].precio = precio
			}
			existe = true
		}
	}
	if !existe {
		(*l) = append((*l), producto{nombre: nombre, cantidad: cantidad, precio: precio})
	}
}

// Hacer una función adicional que reciba una cantidad potencialmente infinita de productos previamente creados y los agregue a la lista
// Listo
func (l *listaProductos) agregarProductos(variantes ...producto) {
	for _, p := range variantes {
		*l = append(*l, p)
	}
}

// modifique la función para que esta vez solo retorne el producto como tal y que retorne otro valor adicional "err" conteniendo un 0 si no hay error y números mayores si existe algún
// tipo de error como por ejemplo que el producto no exista. Una vez implementada la nueva función, cambie los usos de la anterior función en funciones posteriores, realizando los cambios
// que el uso de la nueva función ameriten
// Listo
func (l *listaProductos) buscarProducto(nombre string) (producto, int) {
	var result producto
	var err int = 0
	for _, p := range *l {
		if p.nombre == nombre {
			result = p
			return result, err // Producto encontrado, devolver producto y error 0
		}
	}
	// Producto no encontrado, devolver un producto vacío y un error mayor que 0
	return producto{}, 1
}

func (l *listaProductos) venderProducto(nombre string, cant int) {
	//modificar para que cuando no haya existencia de cantidad de productos, el producto se elimine de "la lista" y notificar dicha acción
	//modifique posteriormente para que se ingrese de parámetro la cantidad de productos a vender
	//Listo
	prod, err := l.buscarProducto(nombre)
	if err == 0 {
		if cant > prod.cantidad {
			fmt.Println("Error: No hay suficiente cantidad de productos.")
			return
		}
		prod.cantidad -= cant
		if prod.cantidad == 0 {
			// Eliminar el producto de la lista si la cantidad llega a cero
			for i, p := range *l {
				if p.nombre == nombre {
					*l = append((*l)[:i], (*l)[i+1:]...)
					fmt.Println("Producto eliminado:", nombre)
					break
				}
			}
		}
	} else {
		fmt.Println("Error: El producto no se encuentra en la lista.")
	}
}

// haga una función para a partir del nombre del producto, se pueda modificar el precio del mismo Pero utilizando la función buscarProducto MODIFICADA PREVIAMENTE
// Listo
func (l *listaProductos) modificarPrecio(nombre string, nuevoPrecio int) error {
	for i := range *l {
		if strings.EqualFold((*l)[i].nombre, nombre) {
			(*l)[i].precio = nuevoPrecio
			return nil
		}
	}
	return errors.New("El producto no se encuentra en la lista")
}

func llenarDatos() {
	// Agregar productos de uno en uno
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("café", 12, 4500)
	lProductos.agregarProducto("arroz", 15, 2600)
	// Agregar varios productos
	lProductos.agregarProductos(
		producto{nombre: "atun", cantidad: 15, precio: 15000},
		producto{nombre: "arandanos", cantidad: 1, precio: 1400},
		producto{nombre: "tomate", cantidad: 8, precio: 3000},
		producto{nombre: "lechuga", cantidad: 4, precio: 3600},
		// Puedes agregar mas productos aquí
	)

}

func (l *listaProductos) listarProductosMínimos() listaProductos {
	// debe retornar una nueva lista con productos con existencia mínima
	//Listo
	// Filtra y retorna una nueva lista con productos que tienen existencia mínima
	var productosMinimos listaProductos
	for _, p := range *l {
		if p.cantidad <= existenciaMinima {
			productosMinimos = append(productosMinimos, p)
		}
	}
	return productosMinimos
}

func main() {
	llenarDatos()
	fmt.Println("Lista de productos antes de modificar el precio:")
	fmt.Println(lProductos)

	err := lProductos.modificarPrecio("atun", 2700)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Precio de 'atun' modificado correctamente.")
	}

	fmt.Println("\nLista de productos después de modificar el precio:")
	fmt.Println(lProductos)

	fmt.Println("Lista de productos mínimos:")
	productosMinimos := lProductos.listarProductosMínimos()
	for _, p := range productosMinimos {
		fmt.Printf("Nombre: %s, Cantidad: %d, Precio: %d\n", p.nombre, p.cantidad, p.precio)
	}
}
