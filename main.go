package main

import (
	productDdal "everymore-go/project/dal"
)

func main() {
	// variables.FirstDemo()
	// conditionals.CreateAccounting(12314.33, 41414)
	// conditionals.Demo1()
	// loops.PredicateGame()
	// loops.WriteFriendshipNumber(17296, 1846)
	// slices.Demo1()
	// functions.HiFunction("Zekeriya BAŞAN")
	// functions.Add(134, 400)
	// added, removed, divided := functions.MultiReturnAddRemoveDivide(56, 7)

	// fmt.Println("toplam: ", added)
	// fmt.Println("fark: ", removed)
	// fmt.Println("bölüm: ", divided)

	// _, difference, _ := functions.MultiReturnAddRemoveDivide(56, 7)
	// fmt.Println("Sadece farkı aldık :", difference)

	// fmt.Println("Toplam: ", functions.AddVariadic(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	// numbers := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}
	// fmt.Println("Array de verebiliriz > Toplam: ", functions.AddVariadic(numbers...))

	// maps.FetchDictionaryValue("blue")
	// maps.FetchDictionaryValue("asdasdasd")

	// maps.FetchNumberText(1)
	// maps.FetchNumberText(999)

	// maps.AddANumber(999, "dokuz yüz doksan dokuz")
	// maps.FetchNumberText(999)

	// maps.DeleteANumber(1000)
	// maps.DeleteANumber(999)

	// maps.UpdateANumber(1000, "BİN")
	// maps.UpdateANumber(1, "BİR")

	// for_range.AddOddNumbers(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	// for_range.WriteOddNumbers()

	// colors := map[string]string{"blue": "mavi", "yellow": "sarı", "pink": "pembe"}
	// for_range.WriteToConsoleMap(colors, " - ")

	// number := 20

	// pointers.Demo1(&number)
	// fmt.Println(number)

	// numbers := []int{1, 2, 3, 4, 5}
	// pointers.Demo2(numbers)
	// fmt.Println("MAIN : ", numbers)

	// structs.Demo1()
	// var z = structs.Customer{Id: 1, Name: "Zekeriya", Surname: "BAŞAN", Age: 27}
	// var m = structs.Customer{Id: 2, Name: "Mustafa Kemal", Surname: "BAŞAN", Age: 27}
	// var c = structs.Customer{Id: 3, Name: "Çağdaş", Surname: "BAŞAN", Age: 27}
	// var cc = structs.Customer{Id: 4, Name: "Çağrı", Surname: "BAŞAN", Age: 27}

	// structs.AddACustomer(z)
	// structs.AddACustomer(m)
	// structs.AddACustomer(c)
	// structs.AddACustomer(cc)

	// fmt.Println(structs.GetCustomers())

	// m.Age = 31
	// c.Age = 34
	// cc.Age = 38

	// structs.UpdateACustomer(m)
	// structs.UpdateACustomer(c)
	// structs.UpdateACustomer(cc)

	// fmt.Println(structs.GetCustomers())

	// goroutines.WriteColors([]string{"mavi", "sarı", "yeşil", "kırmızı", "mor", "siyah", "gri"}) // sync
	// goroutines.WriteNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})                                   //sync

	// go goroutines.WriteColors([]string{"mavi", "sarı", "yeşil", "kırmızı", "mor", "siyah", "gri"}) // async
	// go goroutines.WriteNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})                                //async

	// time.Sleep(15 * time.Second)

	// sumOddNumberCn := make(chan int) // bu kanallar sayasinde async işlemin tamamalndığını garanti ediyoruz artı değer alıyoruz.
	// writedCn := make(chan bool)

	// go channels.WriteColors(writedCn, []string{"mavi", "sarı", "yeşil", "kırmızı", "mor", "siyah", "gri"}) // async
	// go channels.SumOddNumber(sumOddNumberCn, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 99})                      //async

	// isWritedColors, totalOddNumers := <-writedCn, <-sumOddNumberCn

	// fmt.Printf("Renkler yazıldı mı? : %v\nToplam: %d\n", isWritedColors, totalOddNumers)

	// interfaces.Demo1() // interdacelerde aynı metodu imlemente eder gibi yazıyoruz aracı bir fun ile structlar üz kullanıyoruz
	// interfaces.Demo2()

	// deferstatement.B()

	// deferstatement.Calculate()

	// deferstatement.Demo3()

	// errorhandling.OpenFile("./error_handling/zeko.txt")
	// errorhandling.OpenFile("./invalidPAth/zeko.txt")

	// interfaces.Demo3()

	// errorhandling.StartGame()

	// stringfunctions.Demo1()

	// restful.Demo1()
	// restful.AddTodo()

	productDdal.GetAll()
}
