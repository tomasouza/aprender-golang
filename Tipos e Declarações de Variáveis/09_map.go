package main

import "fmt"

func main() {
	m := map[string]int{} // declarando um mapeamento de string para int e atribuindo a variavel
	m2 := map[string]int{ // declarando um mapeamento de string para int e atribuindo a variavel
		"a": 1,
		"b": 2,
		"c": 3,
	}
	m3 := make(map[string]int) // declarando um mapeamento de string para int e atribuindo a variavel

	fmt.Println(m, m2, m3)
	m["hello"] = 20
	fmt.Println(m["hello"])
	fmt.Println(len(m), len(m2), len(m3))

	m["present"] = 0
	fmt.Println(m["present"], m["not present"])
	//virgula ok idioma
	val, ok := m["present"] // a primeira variavel ganha o valor a segunda se ele esta presente no mapa
	fmt.Println(val, ok)

	val, ok = m["not present"]
	fmt.Println(val, ok)

	//deletando de mapas
	delete(m, "present")
	val, ok = m["present"] // como deletamos o ok agora exibira false
	fmt.Println(val, ok)

	delete(m, "not present")
	val, ok = m["not present"]
	fmt.Println(val, ok)
}
