package House

type IBuilder interface {
	setNumOfWindow(window int)
	setNumOfGate(gate int)
}

func GetBuilder(build string) IBuilder {
	switch build {
	case "simple":
		return newSimpleHome()
	case "modern":
		return newModernHome()
	default:
		panic("Builder didn't exist")
	}
}
