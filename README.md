# Bank-operations
Сalculates banking transactions in Go (Golang).

# Functions

**Deposit:**
This function is used to calculate the minimum and maximum deposit amount depending on the currency. The calculation is made for one month. The final result for UZS is displayed in tiyin (tiyin is a unit of currency of Uzbekistan, equal to 1⁄100 of a som), and for USD in cents.

**Syntax:**
```go
func Calculate(amount int, currency string) (min int, max int) {
...
}
```



**Transer:**
This function calculates the total amount when transferring funds from one card to another. The total amount also includes an additional bonus amount (percentage of the amount) depending on the currency. The final result for UZS is displayed in tiyin (tiyin is a unit of currency of Uzbekistan, equal to 1⁄100 of a som), and for USD in cents.

**Syntax**
```go
func Total(amount int, currency string) (totalSum int) {
...
}
```



# Installation

1. Clone this code into a directory:
 ```sh
 git clone https://github.com/SardorMS/bank-operations
 ```

2. Create a module named for example - "bank-operations":
 
 ```sh
 go mod init bank-operations
 ``` 
 
3. To RUN Example test use:
```sh
go test -v ./...
```


# Test Results

```
=== RUN   ExampleCalculate
--- PASS: ExampleCalculate (0.00s)
PASS
ok      bank-operations/pkg/bank/deposit   0.114s
=== RUN   ExampleTotal
--- PASS: ExampleTotal (0.00s)
PASS
ok      bank-operations/pkg/bank/transfer  0.020s
```
