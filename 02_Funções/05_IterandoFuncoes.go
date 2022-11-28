package main

import (
	"fmt"
)

func main() {
	/* declarando um slice de funcoes
	Verifique que a assinatura precisa ser a mesma para que consigamos colocar mais funcoes no slice
	*/
	funcs := []func() (string, error){
		funcaoUm, funcaoDois, funcaoTres, funcaoQuatro,
	}
	ops := []string{
		"funcaoUm", "funcaoDois", "funcaoTres", "funcaoQuatro",
	}
	//iterando sobre o slice de funcoes e executando uma a uma
	for i, f := range funcs {
		v, err := f()
		if err != nil {
			fmt.Println(ops[i], "failed:", err)
		} else {
			fmt.Println(v)
		}
	}
}

func funcaoUm() (string, error) {
	return "Funcao Um Invocada", nil

}
func funcaoDois() (string, error) {
	return "Funcao Dois Invocada", nil
}
func funcaoTres() (string, error) {
	return "Funcao Tres Invocada", nil
}
func funcaoQuatro() (string, error) {
	return "Funcao Quatro Invocada", nil
}
