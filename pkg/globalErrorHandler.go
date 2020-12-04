package pkg

import "fmt"

func Go(fun func()) {
	go func() {
		defer func() {
			err := recover()
			fmt.Println("panic: ", err)
		}()

		fun()
	}()
}
