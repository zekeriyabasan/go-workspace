package conditionals

import "fmt"

func GetCustomerBalance(cutomerNumber int) float64 {
	balances := map[int]float64{
		183671: 500.00,
		183881: 1200.75,
		183761: 800.20,
	}
	bakiye, existCustomer := balances[cutomerNumber]
	if !existCustomer {
		fmt.Printf("Kullanıcı ID %v için bakiye bulunamadı", cutomerNumber)
		return 0
	}

	return bakiye

}

func CreateAccounting(req float64, customerNumber int) {

	if GetCustomerBalance(customerNumber) > req {
		fmt.Print("Paranız hazırlanıyor ...")
	} else {
		fmt.Print("Bakiye yetersiz!")
	}

}
