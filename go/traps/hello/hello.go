package hello

import "fmt"

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func main() {
	fmt.Printf("hello, world %s %q\n", "Sunday", "Sunday")

	months := [...]string{
		1: "January",
		2: "February",
		3: "March",
		4: "April",
		5: "May",
		6: "June",
		7: "July",
		8: "August",
		9: "September",
		10: "October",
		11: "November",
		12: "December"}

	fmt.Printf("Months: %v", months)

	spring := months[1:4]
	fmt.Printf("Spring: %v, list-format %v\n", spring, k(spring))

	var m = make(map[string]int)
	m[k(spring)]++
	fmt.Printf("m: %v", m)
}
