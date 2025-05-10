package errorhandling

import (
	"errors"
	"fmt"
	"math/rand"
)

var randomNumber = rand.Intn(101) // 0 ile 99 arasında üretir
func StartGame() {
	estimate := -1
	fmt.Println("0-100 arası bir sayı tuttum tahmin et !")
	for {
		fmt.Scanln(&estimate)
		_, err := checkValue(estimate)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Bildiniz :)")
			break
		}
	}

}
func checkValue(val interface{}) (bool, error) {
	num, ok := val.(int)

	if !ok {
		return false, errors.New("Lütfen tam sayı giriniz !")
	} else if num < 0 || num > 100 {
		return false, errors.New("0-100 arası bir sayı girmelisiniz !")
	}

	if randomNumber > num {
		return false, errors.New("Daha büyük bir sayi giriniz :")
	}
	if randomNumber < num {
		return false, errors.New("Daha küçük bir sayi giriniz :")
	}

	return true, nil

}
