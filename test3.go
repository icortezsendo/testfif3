package main

import (
	"fmt"
)

func main() {
	b := []byte("11A05AB398765UJ102N2300")
	mapa, err := separaDatos(b)

	if err != "" {
		fmt.Println(err)
	} else {
		fmt.Println(mapa)
	}
}
