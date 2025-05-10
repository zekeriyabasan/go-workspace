package deferstatement

import "fmt"

func A() {
	fmt.Println("A çalıştı.")
}

func B() {
	// bi kutu gibi düşün func çalıştığında deferlar içine atılır ve fonksiyon hata da alsa en son deferler son giren ilk çıkar mantığı ile çalışır. örneğin logmada kullanabiliriz.
	defer Deferly_Function()

	A()
	fmt.Println("B çalıştı.")
}

func Deferly_Function() {
	fmt.Println("Defer li function çalıştı")
}
