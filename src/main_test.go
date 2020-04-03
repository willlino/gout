package main

import "testing"


func TestSoma( t *testing.T ){
	result := soma(5, 5)

	if result != 10 {
		t.Errorf("soma falhou, valor 10 esperado, mas o valor está %v", result)
	}
}