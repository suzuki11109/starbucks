package starbucks

type CoffeeType interface {
	getPrice() int
	name() string
}

type Latte struct {
}

func (c Latte) name() string {
	return "Latte"
}

func (c Latte) getPrice() int {
	return 0
}

type Cappuccino struct {
}

func (c Cappuccino) name() string {
	return "Cappuccino"
}

func (c Cappuccino) getPrice() int {
	return 10
}

type Mocha struct {
}

func (m Mocha) name() string {
	return "Mocha"
}

func (m Mocha) getPrice() int {
	return 20
}

func getCoffeeType(coffeeType string) CoffeeType {
	switch coffeeType {
	case "Cappuccino":
		return Cappuccino{}
	case "Latte":
		return Latte{}
	case "Mocha":
		return Mocha{}
	}

	return Latte{}
}
