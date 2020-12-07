package main

import "fmt"

func main() {
	x := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	x[`rabbit`] = []string{`parrot`, `running`, `love`}

	if _, ok := x[`bond_james`]; ok {
		delete(x, `bond_james`)
	}

	for k, v := range x {
		fmt.Println("record :: ", k)
		for i, val := range v {
			fmt.Printf("\t№ %v\tval :: %v\n", i, val)
		}
	}
}
