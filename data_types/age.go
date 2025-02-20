package data_types

type Age int

func Category(a Age) string {
	if a < 18 {
		return "Child"
	} else if a < 63 {
		return "Adult"
	}
	return "Pensioner"
}
