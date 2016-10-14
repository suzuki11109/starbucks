package starbucks

type Order struct {
	coffeeType CoffeeType
	size       Size
}

func (o Order) getPrice() int {
	var price int
	price += o.size.getPrice()
	price += o.coffeeType.getPrice()

	return price
}

func (o Order) getCoffee() Coffee {
	return Coffee{
		shorts:     o.size.getShorts(),
		coffeeType: o.coffeeType.name(),
	}
}

type Coffee struct {
	coffeeType string
	shorts     int
}

func PlaceOrder(size string, coffeeType string) Order {
	return Order{
		size:       getSize(size),
		coffeeType: getCoffeeType(coffeeType),
	}
}
