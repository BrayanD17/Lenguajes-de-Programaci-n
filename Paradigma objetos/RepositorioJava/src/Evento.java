abstract class Evento {
    protected String descripcion;

    public Evento(String descripcion) {
        this.descripcion = descripcion;
    }

    public String getDescripcion() {
        return descripcion;
    }

    @Override
    public abstract String toString();
}
