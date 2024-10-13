package stock // <- Notice The Package Name

var Stock = 0 // shorthand variable declaration is only allowed inside functions

func Init(initQuantity int) {
	Stock = initQuantity
}
