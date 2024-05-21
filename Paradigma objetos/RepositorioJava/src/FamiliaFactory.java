class FamiliarFactory implements AgendaFactory {
    private String relacion;

    public FamiliarFactory(String relacion) {
        this.relacion = relacion;
    }

    @Override
    public Contacto crearContacto(Persona persona) {
        return new ContactoFamiliar(persona, relacion);
    }

    @Override
    public Evento crearEvento(String descripcion) {
        return new EventoConferencia(descripcion, "Casa"); // ejemplo de evento específico
    }
}