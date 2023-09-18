package main

import "fmt"

// ----- Funções -----

func nomeDaFuncao(agr1, agr2 int) int {
	return agr1 + agr2
}

// funcçõe são formatadas como func nomeFunc (ARGUMENTOS TIPO) (VALORES A SEREM RETORNADOS TIPO DO RETORNO)
//naked return

func getCoords() (x, y int) {
	return // RETORNA X E Y
}

// ----- Struct -----
/* Estutrutura de dados com chave valor, como um objeto*/
type carro struct {
	Modelo        string
	Cor           string
	Peso          string
	Altura        string
	Potencia      string
	RodaDianteira Roda
	RodaTraseira  Roda
}
type Roda struct {
	Raio     int
	Material string
}

func main() {
	// Printando informações
	fmt.Println("Hello world")

	// Declarando Variáveis
	/*
			tipos de variáveis
		bool
		int int8 int16 int32 int64 -> inteiros com customização de bytes alocados
		uint .. .. .. .. .. .. .. -> inteiros positivos
		string
		float32 float644 -> decimais
		complex -> numeros complexos rarametente usados
		byte
		rune

		PRINCIPAIS:
		int
		uint
		float64
	//*/
	//// modo 1 -  com casting de tipo
	//var var1 string = "Isso é uma string. "
	//var var2 int = 3
	//fmt.Println("Variável 1 = ", var1, "\nVariável 2 = ", var2)
	//
	//// modo 2 - maneira mais dinamica com o compilador atribuindo automaticamente o tipo da variável
	//var3 := "Isso é uma string sem casting"
	//var4 := 4
	//fmt.Println("============================================")
	//fmt.Println("Variável 3 = ", var3, "\nVariável 4 = ", var4)
	//
	//fmt.Println("============================================")
	//// modo 3 - passando multiplas variáveis em uma linha de comando
	//var5, var6 := "String 5", 5
	//fmt.Println("Variável 5 = ", var5, "\nVariável 6 = ", var6)
	//fmt.Println("============================================")
	//
	////convertesão de variáveis é possível
	//numeroDecimal := 2.6
	//numeroInteiro := int(numeroDecimal)
	//fmt.Println("Numero decimal = ", numeroDecimal, "\nDecimal convertido para inteiro = ", numeroInteiro)
	//
	////Constantes não podem ser sobrescritas com outros valores
	//const segundosEmMinutso = 60
	//const minutosEmHora = 60
	//const segundosEmHoras = segundosEmMinutso * minutosEmHora
	//fmt.Println("Segundos em uma Hora", segundosEmHoras)

	// ----- Condicionais -----

	altura := 4
	if altura > 4 {
		fmt.Println("você é alto o suficiente")
	} else {
		fmt.Println("Você é baixo")
	}
	// é possível fazer uma condicional declarando a variável na linha do if, desta forma a variável fica fora do escopo do programa

	if comprimento := 6; comprimento > 5 {
		println("O comprimento é suficiente")
	}

	// ----- Struct -----
	// pra instanciar structs:
	meuCarro := carro{} //instancia com valores vazios
	meuCarro.RodaDianteira.Raio = 5

}
