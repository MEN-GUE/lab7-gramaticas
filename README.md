# Laboratorio 7 - Simplificación de Gramáticas

Este repositorio contiene la implementación en Go para el Ejercicio 2 del Laboratorio 7. El programa se encarga de cargar gramáticas libres de contexto desde archivos de texto, validar su sintaxis y aplicar el algoritmo correspondiente para eliminar las producciones vacías (épsilon).

## Autor
Juan Menéndez

## Video Demostrativo
En el siguiente enlace puede visualizar la ejecución del programa, el procesamiento de las gramáticas y la demostración de la validación de errores sintácticos:

[Ver demostración en YouTube](https://youtu.be/53HjLEh2d6k)

## Instrucciones de Ejecución

1. Tener Go instalado en su sistema.
2. Abra una terminal y ubíquese en la raíz del proyecto.
3. Verifique que los archivos de entrada (`gramatica1.txt` y `gramatica2.txt`) se encuentren ubicados en esta misma carpeta raíz.
4. Para iniciar el programa, ejecute el siguiente comando:

```bash
go run ./cmd/simplificador
``