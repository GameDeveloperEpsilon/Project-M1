package main


import "fmt"

func main() {
	fmt.Println("Hello World!")
	
	var var0 float32
	var var1 float32

	var0 = 5.5
	var1 = 72.1
	fmt.Printf("%.1f + %.1f = %.1f\n", var0, var1, var0 + var1)
	writeIntroduction("anonymous")
	fmt.Println(add3Nums(5, 20, 13))
	fmt.Println("Goodbye World!")
}

func writeIntroduction(name string) {
	fmt.Printf("Hello, my name is %s.\n", name)
}

func add3Nums(num0, num1, num2 int32) int32 {
	return num0 + num1 + num2
}