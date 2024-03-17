// Teste de Unidade
// (go test -v) para um retorno mais detalhado dos testes
// (go test --cover)  para obter a % coberta pelas testes
// (go test --coverprofile tests.txt) cria um txt documentando os testes
// (go tool cover --func=tests.txt) mostra a cobertura dos testes referentes a função testada
// (go tool cover --html=tests.txt) cria um html temporário mostrando a coberturea de testes spbre a função testada
package adress

import (
	"strings"
	"testing"
)

// TestTipoDeEndereco testa apenas uma possibilidade de entrada
func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()

	enderecoTeste := "estrada da pedra"
	enderecoTesteFormatado := strings.Split(enderecoTeste, " ")[0]

	if TipoDeEndereco(enderecoTeste) != enderecoTesteFormatado {
		t.Errorf("Tipo '%s' não atende aos requisitos.", enderecoTesteFormatado)
	}
}

type cenario struct {
	recebido string
	esperado string
}

// TestCenariosTipoDeEndereco mapeia todos os casos de sucesso possíveis,
// garantindo que não haverá erros ao passar um tipo bálido
func TestCenariosTipoDeEndereco(t *testing.T) {
	t.Parallel()

	cenariosPossiveis := []cenario{
		{"Rua dos Ingleses", "rua"},
		{"RUA dos Ingleses", "rua"},
		{"R ua dos Ingleses", "false"},
		{"estrada do Amor", "estrada"},
		{"EStRAdA do Amor", "estrada"},
		{"e s t r a d a do Amor", "false"},
		{"AVENIDA da Paz", "avenida"},
		{"Avenida da Paz", "avenida"},
		{"A.V.E.N.I.D.A da Paz", "false"},
		{"RoDovIA Brasileira", "rodovia"},
		{"ROdovIA Brasileira", "rodovia"},
		{"Ro-Do-vI-A Brasileira", "false"},
		{"Praça do 90", "false"},
	}

	for _, teste := range cenariosPossiveis {
		if TipoDeEndereco(teste.recebido) != teste.esperado {
			recebidoFormatado := strings.Split(teste.recebido, " ")[0]
			t.Errorf("Tipo '%s' não atende aos requisitos.", recebidoFormatado)
		}
	}

}
