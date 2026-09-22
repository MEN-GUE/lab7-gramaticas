package gramatica

import (
	"fmt"
	"strings"
)

func EncontrarAnulables(g *Gramatica) map[string]bool {
	anulables := make(map[string]bool)
	cambio := true

	for cambio {
		cambio = false
		for cabeza, cuerpos := range g.Producciones {
			if anulables[cabeza] {
				continue
			}
			for _, cuerpo := range cuerpos {
				if cuerpo == "ε" {
					anulables[cabeza] = true
					cambio = true
					break
				}
				
				todosAnulables := true
				for _, simbolo := range cuerpo {
					if !anulables[string(simbolo)] {
						todosAnulables = false
						break
					}
				}
				
				if todosAnulables {
					anulables[cabeza] = true
					cambio = true
					break
				}
			}
		}
	}
	return anulables
}

func generarCombinaciones(cuerpo string, anulables map[string]bool) []string {
	var indicesAnulables []int
	for i, simbolo := range cuerpo {
		if anulables[string(simbolo)] {
			indicesAnulables = append(indicesAnulables, i)
		}
	}

	m := len(indicesAnulables)
	if m == 0 {
		if cuerpo != "ε" {
			return []string{cuerpo}
		}
		return []string{}
	}

	resultadosMap := make(map[string]bool)
	totalCombinaciones := 1 << m

	for i := 0; i < totalCombinaciones; i++ {
		var nuevaCadena strings.Builder
		for j, simbolo := range cuerpo {
			esAnulable := false
			bitActivo := false

			for k, idxAnulable := range indicesAnulables {
				if j == idxAnulable {
					esAnulable = true
					if (i & (1 << k)) != 0 {
						bitActivo = true
					}
					break
				}
			}

			if !esAnulable || bitActivo {
				nuevaCadena.WriteRune(simbolo)
			}
		}

		resultado := nuevaCadena.String()
		if resultado != "" && resultado != "ε" {
			resultadosMap[resultado] = true
		}
	}

	var combinaciones []string
	for res := range resultadosMap {
		combinaciones = append(combinaciones, res)
	}
	return combinaciones
}

func EliminarEpsilon(g *Gramatica) *Gramatica {
	anulables := EncontrarAnulables(g)
	
	fmt.Printf("Símbolos anulables detectados: ")
	for key := range anulables {
		fmt.Printf("%s ", key)
	}
	fmt.Println()

	nuevaGramatica := NuevaGramatica()

	for _, cabeza := range g.NoTerminales {
		nuevaGramatica.NoTerminales = append(nuevaGramatica.NoTerminales, cabeza)
		cuerpos := g.Producciones[cabeza]
		cuerposUnicos := make(map[string]bool)

		for _, cuerpo := range cuerpos {
			if cuerpo == "ε" {
				continue
			}
			combinaciones := generarCombinaciones(cuerpo, anulables)
			for _, comb := range combinaciones {
				cuerposUnicos[comb] = true
			}
		}

		for cuerpo := range cuerposUnicos {
			nuevaGramatica.Producciones[cabeza] = append(nuevaGramatica.Producciones[cabeza], cuerpo)
		}
	}

	return nuevaGramatica
}

func (g *Gramatica) Imprimir() {
	for _, cabeza := range g.NoTerminales {
		cuerpos := g.Producciones[cabeza]
		if len(cuerpos) > 0 {
			fmt.Printf("%s -> %s\n", cabeza, strings.Join(cuerpos, " | "))
		}
	}
}