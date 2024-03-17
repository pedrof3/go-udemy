package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) {
		t.Parallel()
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Errorf("Valor para área recebida não corresponde ao esperado")
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		t.Parallel()
		cir := Circulo{10}
		areaEsperada := math.Pi * math.Pow(cir.Raio, 2)
		areaRecebida := cir.Area()

		if areaEsperada != areaRecebida {
			t.Errorf("Valor para área recebida não corresponde ao esperado")
		}
	})
}
