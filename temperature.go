package starbucks

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
