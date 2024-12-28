# Task Trace Go

Lea este README en otros idiomas:
- [Português](README.pt-BR.md)
- [English](README.md)
- [Français](README.fr.md)

Task Tracker es un proyecto utilizado para rastrear y gestionar tus tareas. En este proyecto, crearás una interfaz de línea de comandos simple (CLI) para rastrear lo que necesitas hacer, lo que ya has hecho y en lo que estás trabajando actualmente. Este proyecto te ayudará a practicar tus habilidades de programación, incluyendo trabajar con el sistema de archivos, manejar entradas del usuario y construir una aplicación CLI sencilla.

## Roadmap.sh

https://roadmap.sh/projects/task-tracker

## Requisitos

La aplicación debe ejecutarse desde la línea de comandos, aceptar acciones e inputs del usuario como argumentos, y almacenar las tareas en un archivo JSON. El usuario debe ser capaz de:

- Agregar, actualizar y eliminar tareas
- Marcar una tarea como en progreso o completada
- Listar todas las tareas
- Listar todas las tareas completadas
- Listar todas las tareas pendientes
- Listar todas las tareas en progreso

Aquí hay algunas restricciones para guiar la implementación:

- Puedes usar cualquier lenguaje de programación para construir este proyecto.
- Usa argumentos posicionales en la línea de comandos para aceptar entradas del usuario.
- Usa un archivo JSON para almacenar las tareas en el directorio actual.
- El archivo JSON debe crearse si no existe.
- Usa el módulo nativo del sistema de archivos de tu lenguaje de programación para interactuar con el archivo JSON.
- No uses bibliotecas o frameworks externos para construir este proyecto.
- Asegúrate de manejar errores y casos extremos de manera adecuada.

## Ejemplo

A continuación se muestra la lista de comandos y su uso:

```bash
# Agregar una nueva tarea
task-cli add "Comprar comestibles"
# Salida: Tarea agregada con éxito (ID: 1)

# Actualizar y eliminar tareas
task-cli update 1 "Comprar comestibles y cocinar la cena"
task-cli delete 1

# Marcar una tarea como en progreso o completada
task-cli mark-in-progress 1
task-cli mark-done 1

# Listar todas las tareas
task-cli list

# Listar tareas por estado
task-cli list done
task-cli list todo
task-cli list in-progress
```

# Task Tracker CLI

## Propiedades de una Tarea

Cada tarea debe tener las siguientes propiedades:

- **id**: Un identificador único para la tarea.
- **description**: Una breve descripción de la tarea.
- **status**: El estado de la tarea (`todo`, `in-progress`, `done`).
- **createdAt**: La fecha y hora en que se creó la tarea.
- **updatedAt**: La fecha y hora de la última actualización de la tarea.

Asegúrate de agregar estas propiedades al archivo JSON al agregar una nueva tarea y actualízalas al modificar una tarea.

---

## Cómo Empezar

Aquí tienes algunos pasos para ayudarte a empezar con el proyecto **Task Tracker CLI**:

### Configura Tu Entorno de Desarrollo

1. Elige un lenguaje de programación con el que te sientas cómodo (por ejemplo, Python, JavaScript, etc.).
2. Asegúrate de tener instalado un editor de código o IDE (por ejemplo, VSCode, PyCharm).

### Inicialización del Proyecto

1. Crea un nuevo directorio para tu **Task Tracker CLI**.
2. Inicializa un sistema de control de versiones (por ejemplo, Git) para gestionar tu proyecto.

### Implementación de Funcionalidades

1. Comienza creando una estructura básica de CLI para manejar las entradas del usuario.
2. Implementa cada funcionalidad una por una, asegurándote de probar a fondo antes de pasar a la siguiente:
   - Primero, implementa la funcionalidad de **agregar tareas**.
   - Luego, implementa **listar tareas**.
   - Posteriormente, implementa **actualizar tareas** y **marcar tareas como en progreso o completadas**.

### Pruebas y Depuración

1. Prueba cada funcionalidad individualmente para asegurarte de que funcionan como se espera.
2. Verifica el archivo JSON para confirmar que las tareas se están almacenando correctamente.
3. Depura cualquier problema que surja durante el desarrollo.

### Finalización del Proyecto

1. Asegúrate de que todas las funcionalidades estén implementadas y probadas.
2. Limpia tu código y agrega comentarios donde sea necesario.
3. Escribe un buen archivo **README** explicando cómo usar tu **Task Tracker CLI**.

---

## Resultado Final

Al final de este proyecto, habrás desarrollado una herramienta práctica que puede ayudarte a ti o a otros a gestionar tareas de manera eficiente. Este proyecto sienta una base sólida para proyectos de programación más avanzados y aplicaciones del mundo real.

**¡Feliz codificación!**
