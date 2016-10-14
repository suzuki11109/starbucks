package starbucks

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

type Temperature interface {
	getTemperature() string
	getPrice() int
}

type Hot struct {
}

func (h Hot) getTemperature() string {
	return "Hot"
}

func (h Hot) getPrice() int {
	return 0
}

type Iced struct {
}

func (i Iced) getTemperature() string {
	return "Iced"
}

func (i Iced) getPrice() int {
	return 10
}

func createTemperature(temperature string) Temperature {
	switch temperature {
	case "Hot":
		return Hot{}
	case "Iced":
		return Iced{}
	}
	return Hot{}
}

func PlaceOrder(size string, coffeeType string, temperature string) Order {
	return Order{
		size:        getSize(size),
		coffeeType:  getCoffeeType(coffeeType),
		temperature: createTemperature(temperature),
	}
}
