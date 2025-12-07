# Task Tracker CLI

Una herramienta de línea de comandos para gestionar tareas.

## Instalación

```bash
git clone https://github.com/DevLumuz/Task-Tracker-CLI.git
cd Task-Tracker-CLI
go install ./cmd/task-cli
```

> Asegúrate de tener `$HOME/go/bin` en tu PATH:
> ```bash
> export PATH=$PATH:$HOME/go/bin
> ```

## Uso

```bash
task-cli --add "Mi tarea"
task-cli --mark-in-progress 1
task-cli --mark-done 1
task-cli --list
```
