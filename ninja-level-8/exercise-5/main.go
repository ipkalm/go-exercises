package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type users []user

func (u users) Len() int {
	return len(u)
}

func (u users) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

type byAge struct {
	users
}

func (u byAge) Less(i, j int) bool {
	return u.users[i].Age < u.users[j].Age
}

type byLast struct {
	users
}

func (u byLast) Less(i, j int) bool {
	return u.users[i].Last < u.users[j].Last
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	sort.Sort(byAge{users})
	printUsers(users)

	sort.Sort(byLast{users})
	printUsers(users)

	for _, user := range users {
		sort.Strings(user.Sayings)
	}

	printUsers(users)
}

func printUsers(u []user) {
	fmt.Println("-----")
	for _, user := range u {
		fmt.Printf("%v\t%v\t%v\n", user.First, user.Last, user.Age)
		for _, sayings := range user.Sayings {
			fmt.Printf("\t%v\n", sayings)
		}
	}
}
