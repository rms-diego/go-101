package interfaceShowDocument

// criando uma interface para implementar um método que vai retornar o cpf ou cnpj
type ShowDocument interface {
	showTaxId() string
}
