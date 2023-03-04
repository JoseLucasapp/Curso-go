package main

import "fmt"

type item struct {
	prodID int
	qtde   int
	preco  float64
}

type pedido struct {
	userId int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0

	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}

	return total
}

func main() {
	pedido1 := pedido{
		userId: 1,
		itens: []item{
			item{1, 2, 12.10},
			item{2, 1, 23.49},
			item{11, 100, 3.13},
		},
	}

	pedido2 := pedido{
		userId: 2,
		itens: []item{
			{prodID: 1, qtde: 1, preco: 12.10},
			{prodID: 2, qtde: 3, preco: 23.49},
			{prodID: 11, qtde: 2, preco: 3.13},
		},
	}

	fmt.Printf("O preço do pedido é R$ %.2f\n", pedido1.valorTotal())
	fmt.Printf("O preço do pedido 2 é R$ %.2f", pedido2.valorTotal())
}
