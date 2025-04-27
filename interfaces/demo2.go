package interfaces

import (
	"fmt"
	"math/rand"
)

type product interface {
	createOrder() (string, float64)
}

type deliver struct {
	distance       float64
	hasFreeDeliver bool
	deliverNumber  string
	totalCost      float64
	quantity       int
}

type bike struct {
	brand        string
	discountRate float32
	price        float64
	weight       float64
	deliver      deliver
}

type phone struct {
	brand            string
	discountRate     float32
	price            float64
	commissionRate   float32
	isInsuredDeliver bool
	weight           float64
	deliver          deliver
}

func (b bike) createOrder() (string, float64) {
	totalCost := 0.0
	deliverCost := 0.0
	if !b.deliver.hasFreeDeliver {
		deliverCost += b.weight * b.deliver.distance * 0.35 * float64(b.deliver.quantity)
	}

	totalCost += b.price * float64(b.deliver.quantity)

	if b.discountRate > 0 {
		totalCost -= totalCost * float64(b.discountRate)
	}

	b.deliver.totalCost = totalCost + deliverCost

	b.deliver.deliverNumber = fmt.Sprintf("%s-%04d", b.brand, rand.Intn(10000))

	fmt.Printf("Ürün: %s, Adet:%d, Kargo Bedeli: %f, Toplam: %f\n", b.brand, b.deliver.quantity, deliverCost, b.deliver.totalCost)

	return b.deliver.deliverNumber, b.deliver.totalCost
}
func (p phone) createOrder() (string, float64) {
	totalCost := 0.0
	deliverCost := 0.0
	if !p.deliver.hasFreeDeliver {
		if p.isInsuredDeliver {
			deliverCost += p.weight * p.deliver.distance * 0.65 * float64(p.deliver.quantity)
		} else {
			deliverCost += p.weight * p.deliver.distance * 0.35 * float64(p.deliver.quantity)
		}
	}
	totalCost = totalCost * float64(p.deliver.quantity)

	if p.discountRate > 0 {
		totalCost -= totalCost * float64(p.discountRate)
	}

	p.deliver.totalCost = totalCost + totalCost*float64(p.commissionRate) + deliverCost

	p.deliver.deliverNumber = p.brand + "-" + string(rand.Intn(10000))

	fmt.Printf("Ürün: %s, Adet:%d, Kargo Bedeli: %f, Toplam: %f\n", p.brand, p.deliver.quantity, deliverCost, p.deliver.totalCost)
	return p.deliver.deliverNumber, p.deliver.totalCost
}

func order(i []product) {
	cost := 0.0
	total := 0.0
	for _, product := range i {
		_, cost = product.createOrder()
		total += cost
	}
	fmt.Printf("Sipariş Numarası - %d Toplam Ödeme : %f", rand.Intn(10000), total)
}

func Demo2() {
	b := bike{brand: "subrosa salvador", price: 509.234, weight: 11.5, discountRate: 0.15, deliver: deliver{distance: 562.32, hasFreeDeliver: false, quantity: 4}}
	p := phone{
		brand:            "Iphone 15 pro",
		price:            76.569,
		weight:           0.5,
		discountRate:     0.011,
		isInsuredDeliver: true,
		deliver:          deliver{distance: 999.65, hasFreeDeliver: false, quantity: 3},
	}
	order([]product{b, p})

}
