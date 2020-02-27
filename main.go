package main

import (
	"fmt"

	"github.com/rnazmo/sandbox__go_stringer_20200227/painkiller"
)

func main() {
	fmt.Println(painkiller.Paracetamol)          // Paracetamol
	fmt.Println(painkiller.Paracetamol.String()) // Paracetamol
}
