package main
import "testing"

func TestSum(t * testing.T) {
	total := sum(15, 15)
	if total != 30 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: %d", total, 30)
	}
}

func TestSub(t * testing.T) {
	total := sub(15, 10)
	if total != 5 {
		t.Errorf("Resultado da subtração é inválido: Resultado %d. Esperado: %d", total, 5)
	}
}

func TestTimes(t * testing.T) {
	total := times(5, 5)
	if total != 25 {
		t.Errorf("Resultado da multiplicação é inválido: Resultado %d. Esperado: %d", total, 25)
	}
}

func TestSumA(t * testing.T) {
	total := sumA(15, 10)
	if total != 40 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: %d", total, 40)
	}
}

func TestSumB(t * testing.T) {
	total := sumB(15, 10)
	if total != 35 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: %d", total, 35)
	}
}
