package main

import "fmt"

func main() {
	success := make(map[int]string)

	success[123456789] = "Maria"
	success[468486132] = "Pedro"
	success[484231487] = "Ana"
	fmt.Println(success)

	for cpf, name := range success {
		fmt.Printf("%s (CPF: %d)\n", name, cpf)
	}

	fmt.Println(success[468486132])
	delete(success, 468486132)
	fmt.Println(success[468486132])
}
