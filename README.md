# DOCUMENTACIÓN DEL PROYECTO [cite: 1]

**Laboratorio de Software 3**  
**Proyecto:** ActividadLabortorioIII  
**Integrantes:** Cristian Camilo Colorado, Juan Jose Patiño Perez, Juan Manuel Ramos  
**Docente:** Troyano  
**Lenguaje:** Go 1.21

## 1. Descripción del proyecto
El proyecto ActividadLaboratorioIII corresponde a una actividad de Laboratorio de Software 3 orientada al cálculo de la definitiva de estudiantes y a la aplicación del principio de separación estructural. El programa solicita tres notas, cada una con una ponderación de 30 %, 30 % y 40 %, respectivamente.
El código desarrollado hasta el momento incluye la captura y validación de las tres notas. La especificación de la actividad establece además que, una vez procesadas una o más definitivas y el usuario solicite detener el programa, se debe mostrar un resumen estadístico.

## 2. Objetivo
Desarrollar un programa en Go que permita ingresar y validar las notas de los estudiantes, calcular su definitiva mediante las ponderaciones establecidas y, al finalizar el proceso, presentar un resumen estadístico de las definitivas procesadas, manteniendo una separación clara entre entrada, procesamiento, salida y flujo general.

## 3. Requisitos y funcionalidades
* Solicitar tres notas por estudiante.
* Aplicar las ponderaciones de 30 %, 30 % y 40 %.
* Aceptar únicamente valores numéricos entre 0.0 y 5.0.
* Rechazar valores no numéricos y valores fuera del rango permitido.
* Solicitar nuevamente la nota cuando la entrada sea inválida.
* Permitir continuar ingresando estudiantes.
* Permitir detener el programa cuando el usuario lo solicite.
* Mostrar un resumen estadístico al finalizar.

## 4. Historia de usuario
Como instructor del curso, quiero calcular la definitiva de un estudiante usando tres notas ponderadas al 30 %, 30 % y 40 %, para determinar el resultado final de acuerdo con la política de calificación del curso.

## 5. Criterios de aceptación
Los criterios siguientes se toman de la especificación proporcionada por el docente.

| ID | Criterio | Resultado esperado |
|---|---|---|
| CA-01 | Rechazar notas inválidas | Si el valor no es numérico, es menor que 0.0 o mayor que 5.0, el sistema debe rechazarlo y solicitarlo nuevamente. |
| CA-02 | Detener y mostrar resumen | Cuando el instructor solicite detener el programa después de haber calculado una o más definitivas, debe terminar el bucle de entrada y mostrar el resumen estadístico. |
| CA-03 | Contenido del resumen | Debe mostrar total de estudiantes, definitiva mínima, definitiva máxima, promedio de definitiva, aprobados y reprobados. |

## 6. Requisitos técnicos

### 6.1 Dependencias
El código recibido no utiliza librerías externas. Utiliza únicamente paquetes de la biblioteca estándar de Go.

| Paquete | Uso |
|---|---|
| `bufio` | Lectura de datos desde la consola mediante Scanner. |
| `fmt` | Mensajes, impresión y formateo de resultados. |
| `os` | Acceso a la entrada estándar mediante `os.Stdin`. |
| `strconv` | Conversión de texto a valores `float64`. |
| `strings` | Limpieza de texto y normalización de respuestas. |

### 6.2 Versiones
* **Módulo Go:** ActividadLabortorioIII
* **Versión declarada en go.mod:** Go 1.21

El proyecto se encuentra configurado como un módulo de Go mediante el archivo `go.mod`.

## 7. Instalación / Setup
1. Instalar Go.
2. Abrir Visual Studio Code.
3. Abrir la carpeta `ActividadLabortorioIII`.
4. Verificar que estén `README.md`, `go.mod` y `main.go`.
5. Abrir una terminal dentro de la carpeta del proyecto.
6. Comprobar la instalación con `go version`.
7. Ejecutar con `go run main.go`.

## 8. Uso básico
1. Ejecutar el programa.
2. Ingresar la nota 1 (30 %).
3. Ingresar la nota 2 (30 %).
4. Ingresar la nota 3 (40 %).
5. Si una nota es inválida, corregirla cuando el programa la solicite nuevamente.
6. El programa muestra las notas capturadas.
7. Indicar `s/si/sí` para continuar o `n/no` para finalizar.
8. Al finalizar, el sistema debe mostrar el resumen estadístico definido en los criterios de aceptación.

## 9. Estructura actual del proyecto
La estructura observada en el repositorio es:

```text
ActividadLabortorioIII/
├── README.md
├── go.mod
└── main.go
```

| Archivo | Responsabilidad |
|---|---|
| `README.md` | Contiene actualmente el título/identificación básica del proyecto. |
| `go.mod` | Define el módulo ActividadLabortorioIII y declara Go 1.21. |
| `main.go` | Contiene el código desarrollado hasta el momento para captura y validación de las tres notas. |

## 10. Funciones implementadas en main.go

| Función | Responsabilidad |
|---|---|
| `EsNotaValida` | Comprueba que una nota esté entre 0.0 y 5.0. |
| `ConvertirYValidarNota` | Limpia la entrada, la convierte a `float64` y aplica la validación. |
| `SolicitarNotaValida` | Solicita la nota al usuario y repite la solicitud cuando la entrada es inválida. |
| `SolicitarTresNotas` | Coordina la captura de las tres notas y sus ponderaciones. |
| `DeseaContinuar` | Pregunta si el usuario desea procesar otro estudiante. |
| `main` | Inicializa el Scanner y coordina el ciclo general de entrada. |

## 11. Separación estructural
La documentación del docente define la separación estructural como la organización de un programa en la que diferentes responsabilidades se colocan en diferentes partes del código. También indica que un programa estructuralmente separado evita mezclar entrada, procesamiento, salida, flujo de alto nivel y detalles de implementación.

En el código del equipo se observa esta idea mediante funciones con responsabilidades específicas. La validación está separada de la solicitud de datos y la decisión de continuar está separada del resto del procesamiento.

### 11.1 Aplicación en el código
* **`EsNotaValida`** concentra la regla de negocio del rango válido.
* **`ConvertirYValidarNota`** se ocupa de convertir la entrada textual y validarla.
* **`SolicitarNotaValida`** maneja la interacción y los reintentos.
* **`SolicitarTresNotas`** coordina la captura de las tres notas.
* **`DeseaContinuar`** maneja específicamente la decisión de continuar o detenerse.
* **`main`** funciona como punto de coordinación general.

### 11.2 Relación con el ejemplo del profesor
El código de referencia del profesor lleva la separación estructural a un nivel más desacoplado: separa la entrada mediante `FloatReader`, la presentación mediante `ResultPresenter`, la lógica de negocio mediante funciones puras y la coordinación mediante `RunFinalGradeCalculator`. El código del equipo aplica el mismo principio general mediante funciones separadas, aunque actualmente no utiliza esas interfaces.

Por tanto, la separación estructural no significa que el código tenga que ser idéntico al ejemplo del profesor; significa que las responsabilidades deben estar organizadas y no concentradas innecesariamente en una sola función.

## 12. Cálculo de la definitiva
La historia de usuario establece tres notas con ponderaciones de 30 %, 30 % y 40 %. La definitiva debe calcularse de la siguiente manera:

**Definitiva = (Nota 1 × 0.30) + (Nota 2 × 0.30) + (Nota 3 × 0.40)**

*Ejemplo:* para las notas 4.0, 3.5 y 4.5, la definitiva es 4.05.

*Nota de implementación:* el `main.go` recibido hasta este momento captura y valida las tres notas, pero todavía no contiene la función que realiza este cálculo. Esta función debe integrarse posteriormente por el integrante encargado del procesamiento de la definitiva.

## 13. Resumen estadístico requerido
De acuerdo con el criterio de aceptación de la actividad, cuando el usuario seleccione la opción de detener el programa, el sistema debe terminar el bucle de entrada y mostrar un resumen estadístico.

| Dato | Descripción |
|---|---|
| **Cantidad total de estudiantes** | Número total de estudiantes procesados. |
| **Nota final mínima** | Menor definitiva obtenida. |
| **Nota final máxima** | Mayor definitiva obtenida. |
| **Promedio de nota final** | Promedio de todas las definitivas calculadas. |
| **Aprobados** | Número de estudiantes que alcanzaron la condición de aprobación. |
| **Reprobados** | Número de estudiantes que no alcanzaron la condición de aprobación. |

El código de captura recibido todavía no implementa este resumen. Esta sección documenta el requisito funcional que debe cumplir la integración final, sin afirmar que ya está implementado.

## 14. Pruebas

| ID | Caso | Entrada | Resultado esperado |
|---|---|---|---|
| **CP-01** | Nota válida | `4.0` | La nota es aceptada. |
| **CP-02** | Límite inferior | `0.0` | La nota es aceptada. |
| **CP-03** | Límite superior | `5.0` | La nota es aceptada. |
| **CP-04** | Menor al rango | `-1.0` | Se rechaza y se solicita nuevamente. |
| **CP-05** | Mayor al rango | `6.0` | Se rechaza y se solicita nuevamente. |
| **CP-06** | Dato no numérico | `abc` | Se rechaza y se solicita nuevamente. |
| **CP-07** | Tres notas válidas | `4.0`, `3.5`, `4.5` | Se capturan correctamente. |
| **CP-08** | Continuar | `s` | Se inicia la captura de otro estudiante. |
| **CP-09** | Continuar | `sí` | Se inicia la captura de otro estudiante. |
| **CP-10** | Finalizar | `n` | Termina el bucle y muestra el resumen cuando la integración esté completa. |
| **CP-11** | Finalizar | `no` | Termina el bucle y muestra el resumen cuando la integración esté completa. |
| **CP-12** | Cálculo | `4.0`, `3.5`, `4.5` | La definitiva esperada es 4.05. |

## 15. Evidencias de ejecución
* Ingreso correcto de las tres notas.
* Rechazo de un dato no numérico.
* Rechazo de una nota menor que 0.0 o mayor que 5.0.
* Opción para continuar con otro estudiante.
* Finalización del programa.
* Cálculo de la definitiva.
* Resumen estadístico final.

*(Las capturas reales deben incorporarse a esta sección antes de entregar el documento).*

## 16. Trazabilidad

| Requisito | Criterio | Prueba | Evidencia |
|---|---|---|---|
| Validación de notas | CA-01 | CP-01 a CP-06 | Capturas de entradas válidas e inválidas |
| Procesamiento y finalización | CA-02 | CP-08 a CP-11 | Captura de continuación/finalización |
| Definitiva | Historia de usuario | CP-12 | Captura del cálculo |
| Resumen estadístico | CA-02 / CA-03 | Pruebas del resumen | Captura del resumen final |
