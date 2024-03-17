package adress

import (
	"strings"
)

func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	tipoEndereco := strings.Split(endereco, " ")[0]
	enderecoFormatado := strings.ToLower(tipoEndereco)

	for _, tipo := range tiposValidos {
		if enderecoFormatado == tipo {
			return enderecoFormatado
		}
	}
	return "false"
}
