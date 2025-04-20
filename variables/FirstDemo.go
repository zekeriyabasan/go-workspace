package variables

import "fmt"

func FirstDemo() {
	fmt.Println("Hi!")
	var hi string = "Hi Zek !"

	for i := 0; i < 5; i++ {
		fmt.Println(hi)

	}

	var floatNum float32 = 304.44

	fmt.Printf("float sayı : %T\n", floatNum)

	sayitipsiz := "tipi tipi"
	fmt.Printf("tipi : %T\n", sayitipsiz)

	var isSame bool

	var name1 string = "Zekeriya"
	var name2 string = "Mustafa"

	isSame = name1 == name2

	if isSame {
		print(name1)
	} else {
		print(name1 + "!=" + name2)
	}
}
