package House

type IBuilder interface {
	SetNumOfWindow()
	SetNumOfGate()
	GetHome() *Home
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
