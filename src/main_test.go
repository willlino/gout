package main

import "testing"


func TestSoma( t *testing.T ){
	result := soma(5, 5)

	if result != 10 {
		t.Errorf("soma falhou, valor 10 esperado, mas o valor está %v", result)
	}

	result2 := soma(2, 3)

	if(result2) != 5 {
		t.Errorf("soma falhou, valor 5 esperado, mas o valor está %v", result2)
	}
}