package transfer

import (
	"fmt"
)

func ExampleTotal() {

	fmt.Println(Total(0, "UZS"))
	fmt.Println(Total(0, "USD"))
	fmt.Println(Total(500_000_00, "UZS"))
	fmt.Println(Total(500_00, "USD"))
	fmt.Println(Total(1_000_000_00, "UZS"))
	fmt.Println(Total(1000_00, "USD"))

	//Output:
	//0
	//0
	//50350000
	//50500
	//100700000
	//101000

}
