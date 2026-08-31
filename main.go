package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// --- LÓGICA DE NEGOCIO Y VALIDACIÓN ---

// EsNotaValida verifica las reglas de negocio (0.0 <= nota <= 5.0)
func EsNotaValida(nota float64) bool {
	return nota >= 0.0 && nota <= 5.0
}

// ConvertirYValidarNota intenta convertir un texto a número y aplica el rango
func ConvertirYValidarNota(entradaTexto string) (float64, bool) {
	// Limpia espacios y saltos de línea
	entradaLimpia := strings.TrimSpace(entradaTexto)

	// Valida que sea un dato numérico
	nota, err := strconv.ParseFloat(entradaLimpia, 64)
	if err != nil {
		return 0, false // No es numérico
	}

	// Valida que esté en el rango de 0.0 a 5.0
	if !EsNotaValida(nota) {
		return 0, false // Fuera de rango
	}

	return nota, true // Nota válida
}

// --- INTERACCIÓN CON LA CONSOLA (I/O) ---

// SolicitarNotaValida se encarga del ciclo de reintento en la terminal
func SolicitarNotaValida(scanner *bufio.Scanner, mensajePrompt string) float64 {
	for {
		fmt.Print(mensajePrompt)
		if !scanner.Scan() {
			continue
		}

		textoIngresado := scanner.Text()
		nota, valida := ConvertirYValidarNota(textoIngresado)

		if valida {
			return nota
		}

		fmt.Println(" Error: Debe ingresar un valor numérico entre 0.0 y 5.0. Intente de nuevo.")
	}
}

// SolicitarTresNotas coordina la lectura de las 3 notas del estudiante[cite: 1]
func SolicitarTresNotas(scanner *bufio.Scanner) (float64, float64, float64) {
	fmt.Println("\n--- Ingreso de Notas del Estudiante ---")
	nota1 := SolicitarNotaValida(scanner, "Ingrese la nota 1 (30%): ")
	nota2 := SolicitarNotaValida(scanner, "Ingrese la nota 2 (30%): ")
	nota3 := SolicitarNotaValida(scanner, "Ingrese la nota 3 (40%): ")

	return nota1, nota2, nota3
}

// DeseaContinuar consulta al usuario si desea procesar otro estudiante o detenerse[cite: 1]
func DeseaContinuar(scanner *bufio.Scanner) bool {
	for {
		fmt.Print("\n¿Desea ingresar otro estudiante? (s/n): ")
		if scanner.Scan() {
			respuesta := strings.ToLower(strings.TrimSpace(scanner.Text()))
			if respuesta == "s" || respuesta == "si" || respuesta == "sí" {
				return true
			}
			if respuesta == "n" || respuesta == "no" {
				return false
			}
		}
		fmt.Println(" Opción inválida. Por favor ingrese 's' para sí o 'n' para no.")
	}
}

// --- PRUEBA / MAIN ---

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Bucle principal del programa según la opción de continuar
	for {
		// 1. Llama a tu función para solicitar las 3 notas validada
		n1, n2, n3 := SolicitarTresNotas(scanner)

		// Solo para demostrar que capturó bien los valores:
		fmt.Printf("Notas capturadas correctamente: [%.2f, %.2f, %.2f]\n", n1, n2, n3)

		// Aquí tu compañero(a) usará estas notas para calcular la nota final...

		// 2. Pregunta si desea procesar otro estudiante o detener el programa[cite: 1]
		if !DeseaContinuar(scanner) {
			fmt.Println("\n¡Proceso finalizado por el usuario!")
			break
		}
	}
}
