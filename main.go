package main

import "fmt"

func main() {
	var pontoDeEmbulicaoKelvin float32 = 373
	var pontoDeEmbulicaoCelsios = conversaoKelvinCelsios(pontoDeEmbulicaoKelvin)
	fmt.Println("Valor do ponto de ebulicao da agua em celscio: ", pontoDeEmbulicaoCelsios)
}

func conversaoKelvinCelsios(valorKelvin float32) float32 {
	var celsios float32 = 0.0
	celsios = valorKelvin - 273
	return celsios
}
