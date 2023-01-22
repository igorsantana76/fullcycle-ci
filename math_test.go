package main

import "testing"

func TestSomar(t *testing.T) {

	total := soma(15, 15)

	if total != 30 {
		t.Error("Resultado deve ser 30")
	}
}

func TestSubtrair(t *testing.T) {

	total := subtracao(15, 15)

	if total != 0 {
		t.Error("Resultado deve ser 0")
	}
}

func TestMultiplicar(t *testing.T) {

	total := multiplicacao(15, 15)

	if total != 225 {
		t.Error("Resultado deve ser 225")
	}
}

func TestDividir(t *testing.T) {

	total := divisao(15, 15)

	if total != 1 {
		t.Error("Resultado deve ser 1")
	}
}
