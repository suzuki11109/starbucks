package starbucks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet2ShortsFromOrderMSize(t *testing.T) {
	order := PlaceOrder("M", "Latte", "Hot")

	coffee := order.getCoffee()

	assert.Equal(t, coffee.shorts, 2)
}

func TestGet3ShortsFromOrderLSize(t *testing.T) {
	order := PlaceOrder("L", "Latte", "Hot")

	coffee := order.getCoffee()

	assert.Equal(t, coffee.shorts, 3)
}

func TestGetCoffeeTypeLatteFromOrderLatte(t *testing.T) {
	order := PlaceOrder("L", "Latte", "Hot")

	coffee := order.getCoffee()

	assert.Equal(t, coffee.coffeeType, "Latte")
}

func TestGetTemperatureHotFromOrderHot(t *testing.T) {
	order := PlaceOrder("L", "Latte", "Hot")

	coffee := order.getCoffee()

	assert.Equal(t, coffee.temperature, "Hot")
}

func TestGetPrice100FromOrderMSize(t *testing.T) {
	order := PlaceOrder("M", "Latte", "Hot")

	assert.Equal(t, order.getPrice(), 100)
}

func TestGetPrice120FromOrderLSize(t *testing.T) {
	order := PlaceOrder("L", "Latte", "Hot")

	assert.Equal(t, order.getPrice(), 120)
}

func TestGetPricePlus10FromOrderIced(t *testing.T) {
	order := PlaceOrder("M", "Latte", "Iced")

	assert.Equal(t, order.getPrice(), 110)
}

func TestGetPricePlus10FromOrderCappuccino(t *testing.T) {
	order := PlaceOrder("M", "Cappuccino", "Hot")

	assert.Equal(t, order.getPrice(), 110)
}
