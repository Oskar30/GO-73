package main

import (
	"fmt"
	"math/rand"
	"module9/pkg/human"
)

func main() {
	people := make(map[string]human.Man)
	human.RandMan(people)

	suspects := make([]string, 0, len(people))

	for k, _ := range people {
		suspects = append(suspects, k)
	}

	suspects = suspects[:rand.Intn(len(suspects))]

	maxCrimes := 0
	if len(suspects) > 0 {

		for i := len(suspects) - 1; i >= 0; i-- {

			if people[suspects[i]].Crimes > maxCrimes {
				maxCrimes = people[suspects[i]].Crimes
			}
		}

		for k, v := range people {
			if v.Crimes == maxCrimes {
				fmt.Println(people[k].Stringer())
			}
		}
	} else {
		fmt.Println("В базе данных нет информации по запрошенным подозреваемым")
	}

}
