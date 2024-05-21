class ContactoFactory implements AgendaFactory {
    @Override
    public Contacto crearContacto(Persona persona) {
        return new ContactoSimple(persona); // o cualquier otra subclase de Contacto
    }

    @Override
    public Evento crearEvento(String descripcion) {
        return new EventoSimple(descripcion); // o cualquier otra subclase de Evento
    }
}