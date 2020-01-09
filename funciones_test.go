package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSeparaDatos(t *testing.T) {
	lista := [9][2]string{{"11A05AB398765UJ102N2300", "23"},
		{"1nA05AB398765UJ102N2300", "Error en convertir el largo a numerico: 1n"},
		{"11N05AB398765UJ102N2300", "Tipo de dato definido como Numerico y es Alfanumerico"},
		{"11C05AB398765UJ102N2300", "Tipo dato debe ser N o A, no C"},
		{"11A05AB398765UJ10nN2300", "Error en convertir el largo a numerico: 0n"},
		{"11A05AB398765UJ102C2300", "Tipo dato debe ser N o A, no C"},
		{"11A05AB398765UJ102N230c", "Tipo de dato definido como Numerico y es Alfanumerico"},
		{"", "No se introdujo el dato a leer"},
		{"11A05AB398765UJ102N230001N661", "23"},
	}
	for _, dato := range lista {
		mapa, err := separaDatos([]byte(dato[0]))
		if err == "" {
			assert.Contains(t, mapa, dato[1])
		} else {
			assert.Equal(t, err, dato[1])
		}

	}
}
