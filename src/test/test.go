package main

import "fmt"

func main() {

	for i := 1; i <= 218; i++ {
		fmt.Printf("Insert into class_old(data_id) values('")
		print(i)
		fmt.Println("');")
	}
}
