package loops

import "fmt"

func WriteFriendshipNumber(sayi1 int, sayi2 int) {
	sayi1Toplam := CalculateAllDvider(sayi1)
	sayi2Toplam := CalculateAllDvider(sayi2)

	if sayi1Toplam == sayi2 && sayi2Toplam == sayi1 {
		print("Siz Kardeş sayılarsınıııız !")
	}

}

func CalculateAllDvider(sayi int) int {
	result := 0
	for i := 1; i < sayi; i++ {
		if sayi%i == 0 {
			result += i
		}

	}
	fmt.Println(result)
	return result

}
