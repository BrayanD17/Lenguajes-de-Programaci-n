public class Main {
    public static void main(String[] args) {
        // Usando Singleton (Lazy)
        /*
        * Singleton
        * El patrón Singleton es útil para la clase Agenda porque asegura que solo haya una instancia de la agenda en
        * la aplicación, evitando inconsistencias y problemas de sincronización. Elegimos una implementación Lazy
        * Singleton para retrasar la creación de la instancia hasta que sea necesaria, lo que puede ser útil en
        * términos de eficiencia si la instancia de Agenda no es utilizada inmediatamente.
        */
        Agenda agenda = Agenda.getInstancia();

        Persona persona1 = new Persona("Juan", "Perez", "juan.perez@example.com");
        Persona persona2 = new Persona("Maria", "Gonzalez", "maria.gonzalez@example.com");

        // Usando Abstract Factory para crear contactos y eventos
        /*
        * Abstract Factory
        * El patrón Abstract Factory es adecuado para la creación de contactos y eventos, permitiendo una mayor
        * flexibilidad y extensibilidad en la creación de diferentes tipos de contactos y eventos sin acoplar el
        * código cliente a las clases concretas. Esto facilita la incorporación de nuevos tipos de contactos y
        * eventos en el futuro.
        */
        AgendaFactory contactoFactory = new ContactoFactory();
        Contacto contacto1 = contactoFactory.crearContacto(persona1);
        Evento evento1 = contactoFactory.crearEvento("Reunión de equipo");

        AgendaFactory familiarFactory = new FamiliarFactory("Hermana");
        Contacto contacto2 = familiarFactory.crearContacto(persona2);
        Evento evento2 = familiarFactory.crearEvento("Conferencia de Java");

        agenda.añadirElemento(contacto1);
        agenda.añadirElemento(contacto2);
        agenda.añadirElemento(evento1);
        agenda.añadirElemento(evento2);

        System.out.println("Todos los elementos de la agenda:");
        agenda.imprimirElementos();

        System.out.println("\nContactos en la agenda:");
        agenda.filtrarContactos().forEach(System.out::println);

        System.out.println("\nEventos en la agenda:");
        agenda.filtrarEventos().forEach(System.out::println);
        /*
        * Conclusión
        * Esta implementación en Java sigue un enfoque 100% orientado a objetos y utiliza patrones de diseño para
        * mejorar la flexibilidad y mantenibilidad del código. La clase Agenda es un Singleton, y usamos
        * Abstract Factory para la creación de contactos y eventos. La solución incluye todos los requisitos del
        * problema, mostrando cómo manejar agendas de contactos y eventos de manera eficiente.
        */
    }
}
