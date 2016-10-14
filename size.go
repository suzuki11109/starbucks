package starbucks

type Size interface {
	getShorts() int
	getPrice() int
}

type MSize struct {
}

func (m MSize) getPrice() int {
	return 100
}

func (m MSize) getShorts() int {
	return 2
}

type LSize struct {
}

func (l LSize) getPrice() int {
	return 120
}

func (l LSize) getShorts() int {
	return 3
}

func getSize(size string) Size {
	switch size {
	case "M":
		return MSize{}
	case "L":
		return LSize{}
	}
	return MSize{}
}
