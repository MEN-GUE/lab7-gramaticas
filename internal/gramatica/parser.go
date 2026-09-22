package gramatica

import "strings"

type Gramatica struct {
	Producciones map[string][]string
	NoTerminales []string
}

func NuevaGramatica() *Gramatica {
	return &Gramatica{
		Producciones: make(map[string][]string),
		NoTerminales: []string{},
	}
}

func (g *Gramatica) AgregarProduccion(linea string) {
	partes := strings.Split(linea, "->")
	cabeza := strings.TrimSpace(partes[0])
	cuerpos := strings.Split(partes[1], "|")

	if _, existe := g.Producciones[cabeza]; !existe {
		g.NoTerminales = append(g.NoTerminales, cabeza)
	}

	for _, cuerpo := range cuerpos {
		g.Producciones[cabeza] = append(g.Producciones[cabeza], strings.TrimSpace(cuerpo))
	}
}