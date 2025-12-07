package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"tasks/internal/service"
)

func main() {
	acctions := map[string]int{"--add": 1, "--remove": 2, "--update": 3, "--get": 4, "--list": 5}

	if len(os.Args) > 2 {
		// obtener el servicio
		s := service.NewTaskService()
		action, exists := acctions[os.Args[1]]
		arg2 := os.Args[2]
		if exists && arg2 != "" {
			switch action {
			case 1:
				if err := s.Add(arg2); err != nil {
					log.Fatal(err)
				}
				fmt.Println("La tarea se guardo exitosamente.")
			case 2:
				if err := s.Remove(arg2); err != nil {
					log.Fatal("error eliminando tarea")
				}
				fmt.Printf("se elimino la tarea %s", arg2)
			case 3:
				task_message := os.Args[3]
				if task_message == "" {
					fmt.Println("la tarea debe contener un mensaje valido")
				}
				if err := s.Update(arg2, task_message); err != nil {
					log.Fatal(err)
				}
				fmt.Println("actulizacion  realizada")
			case 4:
				task, err := s.Get(arg2)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("tarea encontrada  %v", task)
			case 5:
				n, err := strconv.Atoi(arg2)
				if err != nil {
					log.Fatal("error: el segundo argumento no es válido")
					return
				}
				for _, v := range s.List(n) {
					if v.Id != "" {
						fmt.Printf("Tarae ID: %s, Mensaje: %s \n", v.Id, v.Message)
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
