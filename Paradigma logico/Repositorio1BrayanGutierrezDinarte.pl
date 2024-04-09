/* 1)	Construya una función que se llame sub_conjunto. Esta función recibe dos listas y debe retornar True 
cuando el primer argumento es subconjunto completo del segundo y #f en caso contrario. Por ejemplo: 
sub_conjunto([],[1,2,3,4,5]).
True
sub_conjunto([1,2,3],[1,2,3,4,5]).
True
sub_conjunto([1,2,6],[1,2,3,4,5]).
False
*/
miembro(N,[N|_]).
miembro(N,[_|Tail]):- miembro(N,Tail),!.

sub_conjunto([], _).
sub_conjunto([X|Xs], Ys) :-
    miembro(X, Ys),
    sub_conjunto(Xs, Ys).

/* 2)	Implemente la función aplanar. Esta función recibe una lista con múltiples listas anidadas como elementos y 
devuelve una lista con los mismos elementos de manera lineal (sin listas). Ej: 
aplanar([1,2,[3,[4,5],[6,7]]],X).
True
X=[1,2,3,4,5,6,7].
*/
aplanar([], []).
aplanar([X|Xs], Flat) :-
    is_list(X),
    aplanar(X, FlatX),
    aplanar(Xs, FlatXs),
    append(FlatX, FlatXs, Flat).
aplanar([X|Xs], [X|FlatXs]) :-
    \+ is_list(X),
    aplanar(Xs, FlatXs).

/* 3)	La distancia de Hamming entre dos cadenas es el número de posiciones 
en que los correspondientes elementos son distintos. Por ejemplo, 
la distancia de Hamming entre “roma” y “loba” es 2 
(porque hay 2 posiciones en las que los elementos correspondientes son distintos: la posición 1 y la 3). 
Nótese en los ejemplos, que el tamaño de las cadenas puede no ser igual y 
por ende debe tomarse como referencia para la comparación el tamaño de palabra menor)
Ej:	
distanciaH("romano","comino",X).
X = 2 
distanciaH("romano","camino",X).
X = 3 
distanciaH("roma","comino",X).
X = 2 
distanciaH("romano","ron",X).
X = 1 
distanciaH("romano","cama",X).
X = 2
*/
% Predicado para calcular la distancia de Hamming entre dos cadenas

distanciaH(Cadena1, Cadena2, Distancia) :-
    atom_chars(Cadena1, Lista1),           
    atom_chars(Cadena2, Lista2),           
    min_length(Lista1, Lista2, MinLength),
    hamming_distance(Lista1, Lista2, MinLength, Distancia).

min_length(Lista1, Lista2, MinLength) :-
    length(Lista1, Length1),          
    length(Lista2, Length2),          
    MinLength is min(Length1, Length2).

hamming_distance(_, _, 0, 0) :- !.     
hamming_distance([X|Resto1], [X|Resto2], N, Distancia) :-
    hamming_distance(Resto1, Resto2, N, Distancia). 
hamming_distance([_|Resto1], [_|Resto2], N, Distancia) :-
    N > 0,                     
    NewN is N - 1,             
    hamming_distance(Resto1, Resto2, NewN, NewDistancia),
    Distancia is NewDistancia + 1.

hamming_distance(_, _, _, 1).


/* 4)	A partir del siguiente conocimiento relacionado con la carta de un restaurante, 
escriba los hechos y las reglas necesarias para responder a las preguntas que se le hacen.
Los elementos que interesan para el problema de menús de restaurantes son principalmente los 
platos que se pueden consumir y para esto una eventual clasificación puede ser la siguiente:
*/
%comida(nombre, categoria, calorias).
% Tipos de platos
tipo_plato(entrada, [guacamole, ensalada, consome, tostadas_caprese]).
tipo_plato(carne, [filete_de_cerdo, pollo_al_horno, carne_en_salsa]).
tipo_plato(pescado, [tilapia, salmon, trucha]).
tipo_plato(postre, [flan, nueces_con_miel, naranja_confitada, flan_de_coco]).

% Calorías por plato
calorias(guacamole, 200).
calorias(ensalada, 150).
calorias(consome, 300).
calorias(tostadas_caprese, 250).
calorias(filete_de_cerdo, 400).
calorias(pollo_al_horno, 280).
calorias(carne_en_salsa, 320).
calorias(tilapia, 160).
calorias(salmon, 300).
calorias(trucha, 225).
calorias(flan, 200).
calorias(nueces_con_miel, 500).
calorias(naranja_confitada, 450).
calorias(flan_de_coco, 375).

suma_calorias([], 0).
suma_calorias([Plato|Platos], TotalCalorias) :-
    calorias(Plato, CaloriasPlato),
    suma_calorias(Platos, CaloriasRestantes),
    TotalCalorias is CaloriasPlato + CaloriasRestantes.

pertenece(Elem, [Elem|_]).
pertenece(Elem, [_|T]) :- pertenece(Elem, T).

comida_completa_sin_repeticiones(ComidasAnteriores, MaxCalorias, ComidaCompleta) :-
    tipo_plato(entrada, Entradas),
    tipo_plato(carne, Carnes),
    tipo_plato(pescado, Pescados),
    tipo_plato(postre, Postres),
    pertenece(Entrada, Entradas),
    pertenece(Carne, Carnes),
    pertenece(Pescado, Pescados),
    pertenece(Postre, Postres),
    not(pertenece(Entrada, ComidasAnteriores)),
    not(pertenece(Carne, ComidasAnteriores)),
    not(pertenece(Pescado, ComidasAnteriores)),
    not(pertenece(Postre, ComidasAnteriores)),
    suma_calorias([Entrada, Carne, Pescado, Postre], TotalCalorias),
    TotalCalorias =< MaxCalorias,
    ComidaCompleta = [Entrada, Carne, Pescado, Postre].

consultar_comidas(N, MaxCalorias, Combinaciones) :-
    findall(Combinacion, (between(1, N, _), comida_completa_sin_repeticiones([], MaxCalorias, Combinacion)), Combinaciones).

% El codigo se ejecuta asi: consultar_comidas(5, 1000, X). siendo 5 las comidas, 1000 las calorias y X las comidas resultantes.