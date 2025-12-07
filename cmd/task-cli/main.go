package main

import (
	"fmt"
	"log"
	"os"
	"tasks/internal/service"
)

func main() {
	acctions := map[string]int{
		"--add": 1, "--remove": 2,
		"--update": 3, "--mark-done": 4,
		"--mark-in-progress": 4, "--list": 5,
	}

	if len(os.Args) > 2 {
		// obtener el servicio
		s := service.NewTaskService()
		action, exists := acctions[os.Args[1]]
		arg2 := os.Args[2]
		if exists && arg2 != "" {
			switch action {
			// add
			case 1:
				if err := s.Add(arg2); err != nil {
					log.Fatal(err)
				}
				fmt.Println("La tarea se guardo exitosamente.")
			// remove
			case 2:
				if err := s.Remove(arg2); err != nil {
					log.Fatal("error eliminando tarea")
				}
				fmt.Printf("se elimino la tarea %s", arg2)
			// update
			case 3:
				task_message := os.Args[3]
				if task_message == "" {
					fmt.Println("la tarea debe contener un mensaje valido")
				}
				if err := s.Update(arg2, task_message); err != nil {
					log.Fatal(err)
				}
				fmt.Println("actulizacion  realizada")
			// mark-done or
			case 4:
				id := os.Args[2]
				var err error
				if os.Args[1] == "--mark-done" {
					err = s.Mark(id, "done")
				} else {
					err = s.Mark(id, "in-progress")

				}
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("se marco la tarea exitosamente")

			// list
			case 5:
				status := os.Args[2]
				if status == "" {
					fmt.Println("el filtro debe ser valido")
				}

				for _, task := range s.List(status) {
					if task.Id != "" {
						fmt.Printf("%+v \n", task)
					}

				}
			default:
				fmt.Println("accion no encontrada")
			}

		}

	} else {
		fmt.Println("No se proporcionaron argumentos de línea de comando.")
		return
	}
}
