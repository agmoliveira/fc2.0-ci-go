package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(15, 15)

	if total != 30 {
		t.Errorf(" Resultado da soma invalidado esperado valor: Resultado %d, esperado %d", total, 30)

	}
}
