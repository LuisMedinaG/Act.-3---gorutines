package main

import (
	"fmt"
	"time"
)

var show bool

func Proceso(id int, c chan int) {
	i := int(0)
	for {
		select {
		case <-c:
			return
		default:
			i = i + 1
			time.Sleep(time.Millisecond * 500)
			if show {
				fmt.Printf("\nid %d: %d", id, i)
			}
		}
	}
}

func getMenuOpt() (opt int) {
	fmt.Print(`
****** MENU ******

1. Agregar proceso
2. Mostar proceso
3. Eliminar proceso
4. Salir
Ingrese opcion: `)
	fmt.Scan(&opt)
	return opt
}

func main() {
	show = false

    id := 0
	quit := make(map[int]chan int)

	var quitId int

	for {
		opt := getMenuOpt()
		switch opt {
		case 1:
			c := make(chan int)
			quit[id] = c
			go Proceso(id, c)
			id++
		case 2:
			show = !show
		case 3:
			fmt.Scan(&quitId)
			quit[quitId] <- quitId
		case 4:
			fmt.Println("\nEXITING PROGRAM.")
			return
		default:
			fmt.Println("\n", opt, "is an invalid option.")
		}
	}
}
