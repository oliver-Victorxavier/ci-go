package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)
	if total != 30 {
		t.Errorf("Resultado inválido: Resultado %d. Esperado: %d", total, 30)
	}
}

func TestSoma2(t *testing.T) {
	total := Soma(10, 10)
	if total != 20 {
		t.Errorf("Resultado inválido: Resultado %d. Esperado: %d", total, 20)
	}
}
