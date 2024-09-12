package main

import "fmt"

type Igeneric interface {
	int | float64
}

type IGeneric_em_Estruturas[T Igeneric | string | any] struct {
	Data T
}

func SomaGenerica[T Igeneric](i T) T {
	var res T = i + 1
	return res
}

func GenericEmEstruturas() {
	numeros := IGeneric_em_Estruturas[int]{Data: 10}
	textos := IGeneric_em_Estruturas[string]{Data: "Foo"}

	fmt.Println(numeros.Data) // [1 2 3]
	fmt.Println(textos.Data)  // [1 2 3]

	// passing string as type parameter
	// modelStr := Model[string]{Data: []string{"a", "b", "c"}}
	// fmt.Println(modelStr.Data)
}

func main() {
	// fmt.Println(SomaGenerica(22.22))
	GenericEmEstruturas()
}

/*
Aṕlicabilidade: Usar Generics quando voce precisa dizer que vai ser um tipo ou outro.
Beneficio: evitar fazer um monte de tipos ou funcoes para satisfazer cada caso de uso.
Declaracao_de_estrutura: fazer uma interface nomea-la e dentro dela em unica linha colocar que tipos essa interface suportará pra uso.
Declaracao_no_uso: em funcoes antes dos parenteses de params colocar colchetes dizendo o que vai ser o tipo que vou usar com uma variavel de Letra Maiuscula e do lado dela a Nomeada interface que diz os tipos que poderei usar.

Generics em estruturas , para que uma das propriedades seja de um tipo ou de outro
voce tipa os tipos possiveis dentro de colchetes apos o nome da estrutura, e na prop voce da o tipo de valor com a variavel que escolheu pra ser o tipo generico de valor ex: type Model[T INumber | string | any] struct{ Data T }
no uso em funcoes: em funcoes por enquanto defino qual o tipo somente chamando a generica estrutura dentro da funcao e ali defino qual o tipo a ser usado entre os ja marcados. #todo: ainda nao to consequindo definir o tipo na declaracao da funcao repassando variavel pra outras variaveis genericas, mas da pra usar assim tipando dentro da funcao.

*/
