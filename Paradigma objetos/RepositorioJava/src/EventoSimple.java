class EventoSimple extends Evento {
    public EventoSimple(String descripcion) {
        super(descripcion);
    }

    @Override
    public String toString() {
        return "EventoSimple{" +
                "descripcion='" + descripcion + '\'' +
                '}';
    }
}