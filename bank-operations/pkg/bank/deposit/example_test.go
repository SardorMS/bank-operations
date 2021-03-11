package deposit

import (
	"fmt"
)

func ExampleCalculate() {

	fmt.Println(Calculate(0, "UZS"))
	fmt.Println(Calculate(0, "USD"))

	fmt.Println(Calculate(1_000_000_00, "UZS"))
	fmt.Println(Calculate(50_000_00, "USD"))

	fmt.Println(Calculate(10_000_000_00, "UZS"))
	fmt.Println(Calculate(100_000_00, "USD"))

	//Output:
	//0 0
	//0 0
	//416666 583333
	//8333 12500
	//4166666 5833333
	//16666 25000

}
