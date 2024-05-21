abstract class Contacto {
    protected Persona persona;

    public Contacto(Persona persona) {
        this.persona = persona;
    }

    public Persona getPersona() {
        return persona;
    }

    @Override
    public abstract String toString();
}