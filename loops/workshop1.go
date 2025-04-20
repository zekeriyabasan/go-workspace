package loops

import (
	"fmt"
	"math/rand"
)

func PredicateGame() {
	createdNum := rand.Intn(100)
	estimate := -1

	fmt.Println("0-100 arası bir sayı tuttum hadi tahmin et bakalım :)")
	for {
		fmt.Scanln(&estimate)

		if createdNum < estimate {
			fmt.Println("Daha küçük bir sayi !")

		} else if createdNum > estimate {
			fmt.Println("Daha büyük bir sayi !")
		} else {
			fmt.Printf("Sanırım bildin tuttuğum sayi : %v tahmin : %v", createdNum, estimate)
		}

	}

}
