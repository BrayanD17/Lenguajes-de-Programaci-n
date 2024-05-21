class ContactoEmpresarial extends Contacto {
    private String empresa;

    public ContactoEmpresarial(Persona persona, String empresa) {
        super(persona);
        this.empresa = empresa;
    }

    @Override
    public String toString() {
        return "ContactoEmpresarial{" +
                "empresa='" + empresa + '\'' +
                ", persona=" + persona +
                '}';
    }
}
