package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// --- LÓGICA DE NEGOCIO ---

// ResumenEstadistico agrupa los datos estadísticos solicitados.
type ResumenEstadistico struct {
	TotalEstudiantes   int
	NotaMinima         float64
	NotaMaxima         float64
	PromedioNotas      float64
	CantidadAprobados  int
	CantidadReprobados int
}

// CalcularNotaFinal aplica los pesos del 30%, 30% y 40% a las notas.
func CalcularNotaFinal(nota1, nota2, nota3 float64) float64 {
	return (nota1 * 0.30) + (nota2 * 0.30) + (nota3 * 0.40)
}

// EstaAprobado determina si un estudiante aprobó (nota >= 3.0).
func EstaAprobado(notaFinal float64) bool {
	return notaFinal >= 3.0
}

// CalcularEstadisticas procesa el slice de notas finales y retorna el resumen.
func CalcularEstadisticas(notasFinales []float64) ResumenEstadistico {
	if len(notasFinales) == 0 {
		return ResumenEstadistico{}
	}

	estadisticas := ResumenEstadistico{
		TotalEstudiantes: len(notasFinales),
		NotaMinima:       notasFinales[0],
		NotaMaxima:       notasFinales[0],
	}

	var suma float64

	for _, nota := range notasFinales {
		if nota < estadisticas.NotaMinima {
			estadisticas.NotaMinima = nota
		}
		if nota > estadisticas.NotaMaxima {
			estadisticas.NotaMaxima = nota
		}

		suma += nota

		if EstaAprobado(nota) {
			estadisticas.CantidadAprobados++
		} else {
			estadisticas.CantidadReprobados++
		}
	}

	estadisticas.PromedioNotas = suma / float64(estadisticas.TotalEstudiantes)
	return estadisticas
}

// EsNotaValida verifica las reglas de negocio (0.0 <= nota <= 5.0).
func EsNotaValida(nota float64) bool {
	return nota >= 0.0 && nota <= 5.0
}

// ConvertirYValidarNota intenta convertir un texto a número y aplica el rango.
func ConvertirYValidarNota(entradaTexto string) (float64, bool) {
	entradaLimpia := strings.TrimSpace(entradaTexto)

	nota, err := strconv.ParseFloat(entradaLimpia, 64)
	if err != nil {
		return 0, false
	}

	if !EsNotaValida(nota) {
		return 0, false
	}

	return nota, true
}

// --- INTERACCIÓN CON LA CONSOLA ---

// SolicitarNotaValida se encarga del ciclo de reintento en la terminal.
func SolicitarNotaValida(lector *bufio.Scanner, mensajePrompt string) float64 {
	for {
		fmt.Print(mensajePrompt)
		if !lector.Scan() {
			fmt.Println("\nNo se pudo leer la entrada.")
			continue
		}

		nota, valida := ConvertirYValidarNota(lector.Text())
		if valida {
			return nota
		}

		fmt.Println("Error: Debe ingresar un valor numérico entre 0.0 y 5.0.")
	}
}

// SolicitarTresNotas coordina la lectura de las 3 notas del estudiante.
func SolicitarTresNotas(lector *bufio.Scanner) (float64, float64, float64) {
	fmt.Println("\n--- Ingreso de notas del estudiante ---")
	nota1 := SolicitarNotaValida(lector, "Ingrese la nota 1 (30%): ")
	nota2 := SolicitarNotaValida(lector, "Ingrese la nota 2 (30%): ")
	nota3 := SolicitarNotaValida(lector, "Ingrese la nota 3 (40%): ")

	return nota1, nota2, nota3
}

// DeseaContinuar consulta al usuario si desea procesar otro estudiante o detenerse.
func DeseaContinuar(lector *bufio.Scanner) bool {
	for {
		fmt.Print("\n¿Desea ingresar otro estudiante? (s/n): ")
		if lector.Scan() {
			respuesta := strings.ToLower(strings.TrimSpace(lector.Text()))
			if respuesta == "s" || respuesta == "si" || respuesta == "sí" {
				return true
			}
			if respuesta == "n" || respuesta == "no" {
				return false
			}
		}
		fmt.Println("Opción inválida. Por favor ingrese 's' para sí o 'n' para no.")
	}
}

// ImprimirEstadisticas muestra el resumen final por consola.
func ImprimirEstadisticas(estadisticas ResumenEstadistico) {
	fmt.Println("\n--- Resumen general ---")
	fmt.Printf("Total de estudiantes: %d\n", estadisticas.TotalEstudiantes)
	fmt.Printf("Nota mínima: %.2f\n", estadisticas.NotaMinima)
	fmt.Printf("Nota máxima: %.2f\n", estadisticas.NotaMaxima)
	fmt.Printf("Promedio de notas: %.2f\n", estadisticas.PromedioNotas)
	fmt.Printf("Aprobados: %d\n", estadisticas.CantidadAprobados)
	fmt.Printf("Reprobados: %d\n", estadisticas.CantidadReprobados)
}

// --- MAIN ---

func main() {
	lector := bufio.NewScanner(os.Stdin)
	notasFinales := make([]float64, 0)

	fmt.Println("=== Cálculo de notas finales ===")

	for {
		nota1, nota2, nota3 := SolicitarTresNotas(lector)
		notaFinal := CalcularNotaFinal(nota1, nota2, nota3)
		notasFinales = append(notasFinales, notaFinal)

		fmt.Printf("Nota final: %.2f\n", notaFinal)
		if EstaAprobado(notaFinal) {
			fmt.Println("Estado: Aprobado")
		} else {
			fmt.Println("Estado: Reprobado")
		}

		if !DeseaContinuar(lector) {
			break
		}
	}

	estadisticas := CalcularEstadisticas(notasFinales)
	ImprimirEstadisticas(estadisticas)
}
