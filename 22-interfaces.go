package main

import "fmt"

type FiguraGeometrica interface {
	area() int
}

type Quadrado struct {
	Lado int
}

func (q Quadrado) area() int {
	return q.Lado * 2
}

type Retangulo struct {
	Comprimento, Largura int
}

func (r Retangulo) area() int {
	return r.Comprimento * r.Largura
}

func CalculaArea(f FiguraGeometrica) {
	fmt.Println(f.area())
}

func main() {
	quadrado := Quadrado{Lado: 5}
	retangulo := Retangulo{Comprimento: 5, Largura: 10}
	CalculaArea(quadrado)
	CalculaArea(retangulo)
}
