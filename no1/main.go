// Naufal Humam R.P - 2502010585
// [elipsis]
// elipsis merupakan  parameter yang dapat mengambil nol atau lebih argumen
// dengan tipe yang sama (parameter variadik)

package main

import "fmt"

func PrintData(datas ...string) {
	for idx, value := range datas {
		fmt.Printf("Data dari id: %d adalah %s\n", idx, value)
	}
}

func main() {
	fruit := []string{"pisang", "mangga", "jeruk", "semangka"}
	PrintData(fruit...)
}
