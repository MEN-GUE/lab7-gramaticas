package validador

import (
	"log"
	"regexp"
)

func ValidarLinea(linea string) {
	patron := `^[A-Z]\s*->\s*[a-zA-Z0-9ε]+(\s*\|\s*[a-zA-Z0-9ε]+)*$`
	re := regexp.MustCompile(patron)
	
	if !re.MatchString(linea) {
		log.Fatalf("Error de sintaxis detectado en la producción: '%s'. La ejecución se ha detenido.", linea)
	}
}