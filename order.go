package starbucks

func PlaceOrder(size string, coffeeType string, temperature string) Order {
	return Order{
		items: []PricableItem{
			getSize(size),
			getCoffeeType(coffeeType),
			createTemperature(temperature),
		},
	}
}

type Order struct {
	items []PricableItem
}

type PricableItem interface {
	getPrice() int
}

func (o Order) getPrice() int {
	var price int
	for _, v := range o.items {
		price += v.getPrice()
	}
	return price
}

func (o Order) getCoffee() Coffee {
	coffee := Coffee{}

	for _, v := range o.items {
		switch v.(type) {
		case Size:
			coffee.shorts = v.(Size).getShorts()
		case CoffeeType:
			coffee.coffeeType = v.(CoffeeType).name()
		case Temperature:
			coffee.temperature = v.(Temperature).getTemperature()
		}
	}

	return coffee
}

type Coffee struct {
	coffeeType  string
	shorts      int
	temperature string
}
