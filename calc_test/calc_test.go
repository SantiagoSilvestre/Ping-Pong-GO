package main

import "testing"

// padrão triple A A - AAA
// A - Arrange (Preparar)
// A - Act (Rodar)
// A - Assert (verificar as asserções)
func TestShouldSumCorrect(t *testing.T) { // ShouldSumCorrect
	// arrange
	test := soma(3, 2, 1)

	// act
	result := 6

	// assert
	if test != result {
		t.Error("Valor esperado: ", result, "Valor retornado: ", test)
	}
}

func TestShouldSumIncorrect(t *testing.T) { // ShouldSumIncorrect
	test := soma(3, 2, 1)
	result := 7
	if test != result {
		t.Error("Valor esperado:", result, "Valor retornado: ", test)
	}

}

func TestShouldMultiCorrect(t *testing.T) {
	test := multiplica(10, 10)
	result := 100
	if test != result {
		t.Error("Valor esperado:", result, "Valor retorando: ", test)
	}
}

func TestShouldMultIncorrect(t *testing.T) {
	test := multiplica(10, 10)
	result := 22115
	if test != result {
		t.Error("Valor esperado:", result, "Valor retorando: ", test)
	}
}

func TestShoudDividerCorrect(t *testing.T) {
	test := divide(10, 10) // arrange
	result := 1            //act
	if test != result {    // assert
		t.Error("Valor esperado:", result, "Valor retorando: ", test)
	}
}

func TestShoudDividerIncorrect(t *testing.T) {
	test := divide(10, 10) // arrange
	result := 2            //act
	if test != result {    // assert
		t.Error("Valor esperado:", result, "Valor retorando: ", test)
	}
}

func TestShoudSubCorrect(t *testing.T) {
	test := subtracao(10, 5) // arrange
	result := 5              //act
	if test != result {      // assert
		t.Error("Valor esperado:", result, "Valor retorando: ", test)
	}
}

func TestShoudSubIncorrect(t *testing.T) {
	test := subtracao(10, 10) // arrange
	result := 10              //act
	if test != result {       // assert
		t.Error("Valor esperado:", result, "Valor retorando: ", test)
	}
}
