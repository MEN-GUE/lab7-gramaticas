package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"lab7/internal/gramatica"
	"lab7/internal/validador"
)

func procesarArchivo(ruta string) {
	fmt.Printf("\n==================================================\n")
	fmt.Printf("Procesando archivo: %s\n", ruta)
	fmt.Printf("==================================================\n")

	archivo, err := os.Open(ruta)
	if err != nil {
		fmt.Printf("Error al cargar el archivo de texto: %s\n", ruta)
		return
	}
	defer archivo.Close()

	g := gramatica.NuevaGramatica()
	escaner := bufio.NewScanner(archivo)

	for escaner.Scan() {
		linea := strings.TrimSpace(escaner.Text())
		if linea == "" {
			continue
		}
		validador.ValidarLinea(linea)
		g.AgregarProduccion(linea)
	}

	fmt.Println("[1] Gramática original validada y cargada:")
	g.Imprimir()
	
	fmt.Println("\n[2] Ejecutando algoritmo de eliminación de producciones-ε...")
	gLimpia := gramatica.EliminarEpsilon(g)

	fmt.Println("\n[3] Resultado: Gramática sin producciones-ε:")
	gLimpia.Imprimir()
}

func main() {
	procesarArchivo("gramatica1.txt")
	procesarArchivo("gramatica2.txt")
}