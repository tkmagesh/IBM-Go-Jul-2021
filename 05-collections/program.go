package main

import "fmt"

func main() {
	//Array
	//var nos [5]int
	//var nos [5]int = [5]int{3, 5, 1, 4, 2}
	var nos [5]int = [...]int{4, 3, 1, 2, 5}
	fmt.Println(nos)

	//iterating using the indexer
	for i := 0; i < len(nos); i++ {
		fmt.Println(nos[i])
	}

	//iterating using the range construct
	for idx, no := range nos {
		fmt.Printf("Value at [%d] is %d\n", idx, no)
	}

	// creating a copy of the array
	newNos := nos
	newNos[0] = 200
	fmt.Println(newNos, nos)

	//slice
	var names []string = []string{"Magesh", "Suresh"}
	fmt.Println(names, len(names))

	//adding new values
	names = append(names, "Ramesh")
	fmt.Println(names)

	names = append(names, "Ganesh", "Rajesh")
	fmt.Println(names)

	newNames := []string{"John", "Philip"}
	names = append(names, newNames...)
	fmt.Println(names)

	//slicing
	/*
		[lo : hi] => from lo to hi-1
		[lo : ] => from lo to the last item
		[:hi] => from 0 to hi-1
		[lo : lo] => empty slice
		[lo : lo+1] => item at index 'lo'
		[:] => copy of the slice
	*/
	fmt.Println("Slicing from ", names)
	fmt.Println(names[2:4])
	fmt.Println(names[2:])
	fmt.Println(names[:3])
	fmt.Println(names[3:3])
	fmt.Println(names[3:4])

	dupNames := names[:]
	dupNames[0] = "Ram"
	fmt.Println(dupNames, names)

	dupNames = append(dupNames, "Smith")
	fmt.Println(dupNames, names)

	data := make([]int, 10)
	data = append(data, 100)
	fmt.Println(data)

	//maps
	fmt.Println()
	fmt.Println("Maps")
	var cityRanks map[string]int = map[string]int{
		"Bengaluru": 5,
		"Udupi":     1,
		"Mangaluru": 3,
		"Mysuru":    2,
	}

	fmt.Println(cityRanks)
	fmt.Println("Rank of Mysuru => ", cityRanks["Mysuru"])

	//adding a new entry
	cityRanks["Chennai"] = 6
	fmt.Println(cityRanks)

	//checking if a key exists
	if rank, exists := cityRanks["Tumukuru"]; exists {
		fmt.Println("Rank of Tumukuru is => ", rank)
	} else {
		fmt.Println("Tumukuru is not ranked yet!")
	}

	//removing an item
	delete(cityRanks, "Chennai")
	fmt.Println("After removing chennai, cityRanks => ", cityRanks)

	//iterate through the map
	for key, value := range cityRanks {
		fmt.Printf("Value at [%s] is %d\n", key, value)
	}

}
