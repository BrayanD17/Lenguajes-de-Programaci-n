#|
1)Se tiene una cuenta de ahorros en la cual se deposita un capital C al inicio de un periodo de un año.
Por cada año que transcurra se recibe un interés I sobre el capital depositado. Es decir al terminar el año se cuenta con el capital inicial más interés ganado.
Construya una función (int Cap, I, N) que calcula el monto que se recibe al depositar un capital a un interés dado durante un cierto número de años.
Debe producir los resultados de acuerdo con el siguiente cuadro

Capital	Interés	Años	Resultado
2000	0.10	0	2000
2000	0.10	1	2200
2000	0.10	2	2420
2000	0.10	3	2662
|#

(define (calcularCapitalAnual Capital Interes CantidadAnos)
  (cond
    ((= CantidadAnos 0) Capital)
    (true(calcularCapitalAnual (+ Capital (* Capital Interes)) Interes (- CantidadAnos 1)))))


(displayln "1) Ejercicio 1 del repositorio")
(displayln (list "Capital" 2000))
(displayln (list "Interés" 0.10))
(displayln (list "Años" 0))
(displayln (list "Resultado" (calcularCapitalAnual 2000 0.10 0)))
(displayln "_____________________________________________________")
(displayln (list "Capital" 2000))
(displayln (list "Interés" 0.10))
(displayln (list "Años" 1))
(displayln (list "Resultado" (calcularCapitalAnual 2000 0.10 1)))
(displayln "_____________________________________________________")
(displayln (list "Capital" 2000))
(displayln (list "Interés" 0.10))
(displayln (list "Años" 2))
(displayln (list "Resultado" (calcularCapitalAnual 2000 0.10 2)))
(displayln "_____________________________________________________")
(displayln (list "Capital" 2000))
(displayln (list "Interés" 0.10))
(displayln (list "Años" 3))
(displayln (list "Resultado" (calcularCapitalAnual 2000 0.10 3)))
(displayln "_____________________________________________________")
#|
2)Construya una función que se llame merge. Esta función recibe dos listas ordenadas y
produce otra lista ordenada con todos los elementos de las primeras dos listas A continuación se presentan algunos ejemplos
> (merge '(1 2 3 4) '(5 6 7 8))
(1 2 3 4 5 6 7 8)
> (merge '(1 2 3) '(1 2 3 4))
(1 1 2 2 3 3 4)
|#
(define (merge List1 List2)
  (cond
    ((null? List1) List2)
    ((null? List2) List1)
    ((< (car List1) (car List2))
     (cons (car List1) (merge (cdr List1) List2)))
    ((cons (car List2) (merge List1 (cdr List2))))))
(displayln "2)Ejercicio 2 del repositorio")
(displayln (list "Listas '(1 2 3 4) y '(5 6 7 8) ordenadas: " (merge '(1 2 3 4) '(5 6 7 8))))
(displayln (list "Listas '(1 2 3) y '(1 2 3 4) ordenadas: " (merge '(1 2 3) '(1 2 3 4))))
#|
3)Construya una función que se llame sub-conjunto?
Esta función recibe dos argumentos y debe producir un valor #t
cuando el primer argumento es subconjunto del segundo y #f
en caso contrario. Por ejemplo: 
> (sub-conjunto? '() '(a b c d e f))
#t
> (sub-conjunto? '(a b c) '(a b c d e f))
#t
> (sub-conjunto? ‘(a b x) ‘(a b c d e f))
#f
|#
(define (contiene-sublista? sublista lista)
  (cond
    ((null? sublista) #t)
    ((null? lista) #f)
    ((equal? (car lista) (car sublista))
     (contiene-sublista? (cdr lista) (cdr sublista)))
    (else (contiene-sublista? (cdr lista) sublista))))
(displayln "3)Ejercicio 3 del repositorio")
(displayln (list "La lista (1 2 3) contiene los elementos (1 2)" (contiene-sublista? '(1 2) '(1 2 3))))
(displayln (list "La lista (a veces programo) contiene los elementos (a programo)" (contiene-sublista? '(a programo) '(a veces programo))))
(displayln (list "La lista (a b c d e f) contiene los elementos ()" (contiene-sublista? '() '(a b c d e f))))
#|
4)Utilice la función map para definir la función eliminar_elemento que recibe
un elemento E y una lista L y retorna la lista L sin el elemento E.
Si el elemento no existe, retorna la lista L original.

> (eliminar_elemento 3 '(1 2 3 4 5))
(1 2 4 5)

> (eliminar_elemento 0 '(1 2 3 4 5))
(1 2 3 4 5)

|#
(define (eliminar_elemento elemento lista)
  (car(map (lambda (x) (eliminar elemento lista))lista)))

(define (eliminar elem l)
  (cond
    ((empty? l) '())
    ((equal? elem (car l)) (eliminar elem (cdr l)))
    (else (cons (car l) (eliminar elem (cdr l))))))

(displayln "4)Ejercicio 4 del repositorio")
(displayln (list "Lista completa '(1 2 3 4 5 12 34 76 1122), sin el 5 : " (eliminar_elemento 5 '(1 2 3 4 5 12 34 76 1122))))