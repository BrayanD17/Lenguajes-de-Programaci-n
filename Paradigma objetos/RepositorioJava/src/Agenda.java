import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

class Agenda {
    private List<Object> elementos;
    private static Agenda instancia;

    // Constructor privado
    private Agenda() {
        this.elementos = new ArrayList<>();
    }

    // Eager Singleton
    private static final Agenda instanciaEager = new Agenda();

    // Lazy Singleton
    public static Agenda getInstancia() {
        if (instancia == null) {
            instancia = new Agenda();
        }
        return instancia;
    }

    // Eager Singleton getter
    public static Agenda getInstanciaEager() {
        return instanciaEager;
    }

    public void añadirElemento(Object elemento) {
        elementos.add(elemento);
    }

    public void eliminarElemento(Object elemento) {
        elementos.remove(elemento);
    }

    public void modificarElemento(int index, Object nuevoElemento) {
        if (index >= 0 && index < elementos.size()) {
            elementos.set(index, nuevoElemento);
        }
    }

    public void imprimirElementos() {
        for (Object elemento : elementos) {
            System.out.println(elemento);
        }
    }

    public List<Object> filtrarContactos() {
        return elementos.stream().filter(e -> e instanceof Contacto).collect(Collectors.toList());
    }

    public List<Object> filtrarEventos() {
        return elementos.stream().filter(e -> e instanceof Evento).collect(Collectors.toList());
    }
}
