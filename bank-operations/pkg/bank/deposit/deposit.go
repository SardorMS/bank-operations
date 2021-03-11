package deposit

//Calculate - counts minimum and maximum deposit parameters
func Calculate(amount int, currency string) (min int, max int) {
	minPercent, maxPercent := PercentFor(currency)

	min = amount * minPercent / 100 / 12
	max = amount * maxPercent / 100 / 12

	return min, max
}

//PercentFor - returns the percent on the deposit by currency
func PercentFor(currency string) (min int, max int) {
	if currency == "UZS" {
		return 5, 7
	}

	if currency == "USD" {
		return 2, 3
	}

	return 0, 0
}
