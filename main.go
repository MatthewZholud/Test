package main

import (
	"fmt"
)


type shoe struct {
	name  string
}

type Product interface {
	doStuff()
}

func getProduct(shoeBrand string) (Product, error) {
	if shoeBrand == "Nike" {
		return newNike(), nil
	}
	if shoeBrand == "Adidas" {
		return newAdidas(), nil
	}
	return nil, fmt.Errorf("wrong shoe type passed")
}

type Nike struct{
	shoe
}

func (n Nike) doStuff()  {
	fmt.Println("hello from", n.name)
}

func newNike() Nike {
	return Nike{
		shoe: shoe{
			name:  "Nike",
		},
	}
}

type Adidas struct{
	shoe
}

func (a Adidas) doStuff()  {
	fmt.Println("hello from", a.name)
}

func newAdidas() Adidas {
	return Adidas{
		shoe: shoe{
			name:  "Adidas",
		},
	}
}

func main() {
	nike, _ := getProduct("Nike")
	nike.doStuff()

	adidas, _ := getProduct("Adidas")
	adidas.doStuff()
}
