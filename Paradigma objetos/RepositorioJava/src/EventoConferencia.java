class EventoConferencia extends Evento {
    private String lugar;

    public EventoConferencia(String descripcion, String lugar) {
        super(descripcion);
        this.lugar = lugar;
    }

    @Override
    public String toString() {
        return "EventoConferencia{" +
                "descripcion='" + descripcion + '\'' +
                ", lugar='" + lugar + '\'' +
                '}';
    }
}