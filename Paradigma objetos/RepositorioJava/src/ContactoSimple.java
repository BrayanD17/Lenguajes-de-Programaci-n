class ContactoSimple extends Contacto {
    public ContactoSimple(Persona persona) {
        super(persona);
    }

    @Override
    public String toString() {
        return "ContactoSimple{" + persona + '}';
    }
}