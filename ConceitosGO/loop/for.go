package loop

import (
	"fmt"
	"time"
)

func For_Loop_Base() int {
	i := 0 // valor

	for i < 10 { // condicao
		i++ // faz isso enquanto a condicao for verdade
		// fmt.Println(i)
	}

	return i
}

func Loop_Infinito() string {
	for {
		time.Sleep(time.Second * 2)
		fmt.Println("oi")
	}
}

func Loop_For_Contador_Na_Mesma_Linha() {
	// na mesma linha : valor; condicao; mudarValor-Pra_mudar_a_condicao

	for j := 0; j < 10; j++ {
		time.Sleep(time.Second)
		fmt.Println(j)
	}

}

func Loop_Em_Array_Com_Range() string{
	arrayNomes := []string{"Rei", "Guga", "Leo"}
	vazio := ""

	for _, valores := range arrayNomes {
		fmt.Println(valores)
		vazio += fmt.Sprint(" -- ", valores)
	}

	return vazio

}

/*
Loop_For
Conceito: Enquanto a condicao for verdadeira faça algo, obs: ele causa efeito colateral porque age diretamente na variavel mudando seu valor.

Funcionalidade: O que quero que se repita ou interaja a cada volta do loop coloco dentro do loop.

for_na_mesma_linha: declara apos a clausula for : oValor; condicao; mudarValor-Pra_mudar_a_condicao

iterar_com_range_em_array : range quer dizer variedades, ele vai devolver duas variaveis em justa posicao que podemos escolher o nome e declara-las após a clausula for - dizendo que essas 2 variaveis SERÃO range doArrayAlvo { fazer algo com as variaveis } , obs: senao quisermos usar uma das variaveis que ele devolve substitua por anderline _., se for iterar sobre string que é uma colecao de palavras o resultado vira em codigo ASC para te-lo como texto envolva-o no metodo string( resultado )

Recursos: dentro do for nao consigo da return senao ele para o loop entao para ter este mesmo efeito -> posso devolver a variavel modificada como retorno da funcao.


*/
