#|
1) Amplie al ejercicio visto en clase, permitiendo crear listas que representen facturas de ventas y
creando las funciones necesarias para calcular a partir de dichas facturas (nuevas listas): 
a.   El impuesto total a cancelar de la factura.
Solo aquellos productos que sobrepasen un monto mayor definido de argumento en la función pagan impuesto y
dicho impuesto es siempre el 13% sobre el monto.
b.   El monto total de la factura sin impuesto, sin importar ningún umbral

Para el ejercicio, cree quemadas al menos dos facturas a su gusto.
Escriba al menos 3 casos de prueba de sus funciones

|#
;(define factura#(nombre, cantidad del producto, precio unitario, monto a pagar) 
(define factura1 '(("Arroz" 8 1800 14400)
                         ("Frijoles" 5 1200 6000)
                         ("Azúcar" 6 1000 6000)
                         ("Café" 2 2800 5800)
                         ("Leche" 9 1200 10800)))

(define factura2 '(("Aguacate" 3 1500 4500)
                         ("Tomate" 4 900 3600)
                         ("Queso" 2 1800 3600)
                         ("Lechuga" 3 400 1200)
                         ("Pepino" 2 500 1000)))
(define factura3 '(("Audifonos JBL" 3 32000 96000)
                         ("Teclado LogiTec" 1 110000 110000)
                         ("Mouse Red Dragon" 2 28000 56000)))
#|
Punto a) El impuesto total a cancelar de la factura.
Solo aquellos productos que sobrepasen un monto mayor
definido de argumento en la función pagan impuesto y
dicho impuesto es siempre el 13% sobre el monto.
|#
(define (impuesto fact montoMayor)
  (cond
    ((null? fact) (list 0))
    ((> (car (reverse (car fact))) montoMayor)
     (cons (* (car (reverse (car fact))) 0.13)
           (impuesto (cdr fact) montoMayor)))
    (else
     '(0))))

(displayln "Ejercicio 1 punto a) ")
(displayln (apply + (impuesto factura1 5900)))
(displayln (apply + (impuesto factura2 1100)))
(displayln (apply + (impuesto factura3 30000)))
#|
Punto b)El monto total de la factura sin impuesto, sin importar ningún umbral
|#
(define (costoProductos fact)
  (define (ultimo-elemento sublista)
    (car (reverse sublista)))
  (map ultimo-elemento fact))

(displayln "Ejercicio 1 punto b) ")
(displayln (apply + (costoProductos factura1)))
(displayln (apply + (costoProductos factura2)))
(displayln (apply + (costoProductos factura3)))
#|
2) Haciendo uso de la función filter, implemente una función que, a partir de una lista de cadenas de parámetro,
filtre aquellas que contengan una subcadena que el usuario indique en otro argumento. Ej

sub_cadenas "la" ("la casa", "el perro", "pintando la cerca")

("la casa" "pintando la cerca")

sub_cadenas '(1 2 4) '((1 2 3 4 5 11) (12 22 56 9) (2 5 1 2 43 4))
((1 2 3 4 5 11) (2 5 1 2 43 4))
|#
;Codigo haciendo uso de numeros en lugar de letras
(define (contiene-sublista? sublista lista)
  (cond
    ((null? sublista) #t)
    ((null? lista) #f)
    ((equal? (car lista) (car sublista))
     (contiene-sublista? (cdr sublista) (cdr lista)))
    (else (contiene-sublista? sublista (cdr lista)))))

(define (filtrar-contiene-sublista sublista lista)
  (filter (lambda (x) (contiene-sublista? sublista x)) lista))
(displayln "Ejercicio #2: ")
(displayln (filtrar-contiene-sublista '(1 2 4) '((1 2 3 4 5 11) (12 22 56 9) (2 5 1 2 43 4) (1001 1 1233 2 4432 1298 4))))
;Codigo haciendo uso de letras en lugar de numeros
(define (string-contains? haystack needle)
  (not (equal? #f (regexp-match (regexp (regexp-quote needle)) haystack))))

(define (sub-cadenas subcadena lista)
  (filter (lambda (cadena) (string-contains? cadena subcadena)) lista))

(define lista-cadenas '("la casa, " "el perro, " "pintando la cerca, " "comiendo granos, " "lavando la ropa, " "lavando el carro."))
(displayln (sub-cadenas "la" lista-cadenas))


