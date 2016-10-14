package starbucks

type Order struct {
	coffeeType  CoffeeType
	size        Size
	temperature string
}

func (o Order) getPrice() int {
	var price int
	price += o.size.getPrice()
	price += o.coffeeType.getPrice()

	return price
}

func (o Order) getCoffee() Coffee {
	return Coffee{
		shorts:      o.size.getShorts(),
		coffeeType:  o.coffeeType.name(),
		temperature: o.temperature,
	}
}

type Coffee struct {
	coffeeType  string
	shorts      int
	temperature string
}

func PlaceOrder(size string, coffeeType string, temperature string) Order {
	return Order{
		size:        getSize(size),
		coffeeType:  getCoffeeType(coffeeType),
		temperature: temperature,
	}
}
