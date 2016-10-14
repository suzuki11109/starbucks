package starbucks

func PlaceOrder(size string, coffeeType string, temperature string) Order {
	return Order{
		size:        getSize(size),
		coffeeType:  getCoffeeType(coffeeType),
		temperature: createTemperature(temperature),
	}
}

type Order struct {
	coffeeType  CoffeeType
	size        Size
	temperature Temperature
}

func (o Order) getPrice() int {
	var price int
	price += o.size.getPrice()
	price += o.coffeeType.getPrice()
	price += o.temperature.getPrice()

	return price
}

func (o Order) getCoffee() Coffee {
	return Coffee{
		shorts:      o.size.getShorts(),
		coffeeType:  o.coffeeType.name(),
		temperature: o.temperature.getTemperature(),
	}
}

type Coffee struct {
	coffeeType  string
	shorts      int
	temperature string
}
