package transfer

//Total - count the total amount for the transfer
func Total(amount int, currency string) (totalSum int) {

	bonus := PercentFor(currency)
	totalSum = amount + (amount * bonus / 1000)
	return totalSum
}

//PercentFor - returns the percent % on the deposit by currency
func PercentFor(currency string) int {

	if currency == "UZS" {
		return 7 
	}

	if currency == "USD" {
		return 10
	}
	return 0
}
