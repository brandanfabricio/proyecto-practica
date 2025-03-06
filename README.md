# proyecto-practica
**Ruta de Aprendizaje en Go (De Fácil a Avanzado)**

## **Nivel 1: Fundamentos de Go**
**Objetivo:** Dominar la sintaxis, estructuras de datos y manejo de archivos.

1. **Calculadora CLI**
   - Uso de `fmt.Scan` y `os.Args`.
   - Implementar operaciones matemáticas básicas (+, -, *, /).

2. **Generador de contraseñas seguras**
   - Uso de `crypto/rand`.
   - Evitar repeticiones con `map[string]bool`.
   
3. **Lector de archivos CSV/Excel sin librerías**
   - Leer archivos con `os.Open` y `bufio.Reader`.
   - Parseo de datos usando `strings.Split`.

4. **Conversor de texto a mayúsculas/minúsculas**
   - Lectura de entrada del usuario.
   - Uso de `strings.ToUpper` y `strings.ToLower`.

## **Nivel 2: Algoritmos y Estructuras de Datos**
**Objetivo:** Implementar estructuras de datos clásicas y optimizar operaciones.

5. **Sistema de colas y pilas**
   - Implementar `queue` y `stack` con slices y listas enlazadas.
   
6. **Ordenamiento y búsqueda**
   - QuickSort, MergeSort y Búsqueda Binaria.
   
7. **Gestor de tareas CLI con persistencia en archivos**
   - Uso de `os.WriteFile`, `json`, y slices dinámicos.
   
8. **Diccionario de sinónimos en memoria**
   - Implementar `map[string][]string` para almacenamiento.
   - Persistencia opcional en JSON o CSV.

## **Nivel 3: Networking y Servidores Desde Cero**
**Objetivo:** Aprender sobre redes, sockets y protocolos sin `net/http`.

9. **Cliente-servidor TCP desde cero**
   - Implementación manual con `net.Listen` y `net.Conn`.
   
10. **Servidor HTTP desde cero sin `net/http`**
    - Parseo de peticiones HTTP manualmente.
    - Uso de `bufio.Reader` para leer requests.

11. **Cliente y servidor de WebSockets manual**
    - Implementación sin `gorilla/websocket`.
    - Manejo de frames y handshakes.

12. **Proxy HTTP simple**
    - Redirigir tráfico de un servidor a otro.
    - Implementación con `net.Dial` y `io.Copy`.

## **Nivel 4: Concurrencia y Sistemas Distribuidos**
**Objetivo:** Dominar `goroutines`, `channels`, `sync.Mutex`, y arquitecturas paralelas.

13. **Worker Pool para tareas concurrentes**
    - Implementación con `channel` y `sync.WaitGroup`.
    
14. **Chat en tiempo real con TCP y WebSockets**
    - Comunicación entre múltiples clientes.
    
15. **Scraper sin librerías de terceros**
    - Extraer HTML con `golang.org/x/net/html`.

16. **Implementación de un mini "Kafka"**
    - Cola de mensajes distribuida con persistencia en disco.

17. **Cómputo distribuido en Go**
    - Ejecutar tareas en múltiples nodos.
    - Comunicación con `rpc` y `gRPC`.

## **Nivel 5: Construyendo Software Avanzado**
**Objetivo:** Diseñar sistemas complejos aplicando arquitectura avanzada.

18. **Mini intérprete de lenguaje con AST**
    - Construcción de parser y ejecución de expresiones matemáticas.

19. **Base de datos simple tipo SQLite en Go**
    - Implementación de estructuras B-Tree y manejo de páginas de datos.

20. **Motor de plantillas HTML en Go**
    - Parseo y renderización dinámica de HTML.
    
21. **Compilador simple en Go**
    - Generación de bytecode básico para una máquina virtual personalizada.

22. **Sistema de archivos distribuido (tipo IPFS)**
    - Almacenamiento de archivos en múltiples nodos.

---
Este roadmap cubre desde lo más básico hasta temas avanzados como servidores, concurrencia, redes y diseño de sistemas distribuidos. ¿Quieres enfocarte en algún proyecto en particular? 🚀




Dado tu enfoque en usar solo la librería estándar de Go, aquí tienes algunos proyectos desafiantes para practicar y fortalecer tu conocimiento:

Proyectos para Practicar en Go (Solo con la Librería Estándar)
1. Intérprete de un Lenguaje Simple
Implementa un intérprete para un lenguaje básico (similar a Lisp o una calculadora de expresiones).
Usa un autómata finito o un parser recursivo para analizar expresiones matemáticas.
2. Servidor HTTP desde Cero
Implementa un servidor HTTP sin usar net/http, manejando conexiones con net.TCPListener.
Soporta peticiones GET y POST, parsea encabezados y devuelve respuestas.
3. Procesamiento de Archivos de Texto Masivos
Lee y analiza archivos de varios GB línea por línea con bufio.Scanner.
Implementa un contador de palabras o frecuencias de términos.
4. Compresor y Descompresor de Archivos
Crea una implementación simple de compresión/descompresión usando Huffman o Run-Length Encoding.
5. Simulador de Autómatas Finitos
Diseña un autómata finito determinista (DFA) para validar patrones en cadenas de texto (ejemplo: validar números, correos, etc.).
Extiéndelo a un autómata no determinista (NFA) para una mini implementación de expresiones regulares.
6. Parser HTML Manual
Implementa un parser de HTML con un autómata o una estructura de árbol.
Extrae texto de etiquetas <p>, <h1> y atributos de <a>.
7. Generador de Código QR
Implementa el algoritmo de generación de QR sin usar librerías externas.
Usa matrices de bits para representar los patrones.
8. Base de Datos en Memoria
Crea una base de datos simple que almacene pares clave-valor.
Agrega soporte para persistencia usando archivos binarios o JSON.
9. Interpretador de CSV en Stream
Procesa archivos CSV línea por línea sin cargar todo en memoria.
Implementa filtros y agrupaciones simples.
10. Mini Motor de Búsqueda
Indexa archivos de texto y permite búsquedas por palabras clave.
Usa estructuras de datos eficientes como tries o tablas hash.
Si quieres empezar con alguno en particular, dime y te ayudo a guiar el diseño. 🚀