package errorhandling

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var randomNumber = rand.Intn(101) // 0 ile 99 arasında üretir
func StartGame() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("0-100 arası bir sayı tuttum tahmin et !")
	for {
		fmt.Print("Tahminin: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		_, err := checkValue(input)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Bildiniz :)")
			break
		}
	}

}
func checkValue(val interface{}) (bool, error) {
	if value, ok := val.(string); ok {
		num, err := strconv.Atoi(value)
		if err != nil {
			return false, &invalidValue{"Lütfen tam sayı giriniz !"}
		} else if num < 0 || num > 100 {
			return false, &outOfRangeException{num, "0-100 arası sayı giriniz !"}
		}

		if randomNumber > num {
			return false, &invalidValue{"Daha büyük bir sayi giriniz !"}
		}
		if randomNumber < num {
			return false, &invalidValue{"Daha küçük bir sayi giriniz :"}
		}
	}

	return true, nil

}
