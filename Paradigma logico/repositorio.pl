%Implemente un predicado que, a partir de una lista de cadenas de parámetro, 
%filtre aquellas que contengan una subcadena que el usuario indique en otro argumento.
%sub_cadena('la',['hola la cocina','el mueble','la cama'],X).

sub_cadena(_, [],[]).
sub_cadena(P, [H|T], [H|Y]):-
   textoLista(P, H),
   sub_cadena(P, T, Y),!.
sub_cadena(P, [_|T], M):-
   sub_cadena(P, T, M),!.

textoLista(Palabra, Oracion):-
   split_string(Oracion," ","", L),
   miembro(Palabra, L),
   write([Oracion]),!.

miembro(N,[N|_]).
miembro(N,[_|Tail]):- miembro(N,Tail),!.

%2)	Construya una función que se llame sub_conjunto. 
%  Esta función recibe dos listas y debe retornar True cuando el primer argumento es 
%  subconjunto completo del segundo y #f en caso contrario
% sub_conjunto([],[1,2,3,4,5]).
% True
% sub_conjunto([1,2,3],[1,2,3,4,5]).
% True

sub_conjunto([], _):-!.
sub_conjunto([H|T], L):-
   miembro(H, L),
   sub_conjunto(T, L),!.

%   3)	Implemente la función aplanar. Esta función recibe una lista con múltiples listas anidadas como elementos y devuelve una lista con los mismos elementos de manera lineal (sin listas). Ej: 
%   aplanar([1,2,[3,[4,5],[6,7]]],X).
%   True
%   X=[1,2,3,4,5,6,7].
aplanar([], _).
aplanar([H|T], [H|L]):-
   integer(H),
   aplanar(T, L),!.
aplanar([[X|Xs]|Y], [X|L]):-
   integer(X),
   aplanar([Xs|Y], L),!.
aplanar([X|Y], L):-
   is_list(X),
   aplanar(X, L),
   aplanar(Y, L),!.

is_list_or_number(X) :-
   %integer(X),         % Verificar si X es un número entero
   is_list(X).      % Verificar si X no es una lista

%is_list_or_number(X) :-
%   \+ integer(X),      % Verificar si X no es un número entero
%   X = [_|_].          % Verificar si X es una lista
  
   