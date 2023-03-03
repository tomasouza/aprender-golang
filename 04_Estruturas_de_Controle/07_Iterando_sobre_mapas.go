package main

import "fmt"

func main() {

	m := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}

	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}

	//The order of the keys and values varies; some runs may be identical. This is actually a security feature. In earlier Go versions, the iteration order for keys in a map was usually (but not always) the same if you inserted the same items into a map. This caused two problems:
	//People would write code that assumed that the order was fixed, and this would break at weird times.
	//If maps always hash items to the exact same values, and you know that a server is storing some user data in a map, you can actually slow down a server with an attack called Hash DoS by sending it specially crafted data where all of the keys hash to the same bucket.
	//To prevent both of these problems, the Go team made two changes to the map implementation. First, they modified the hash algorithm for maps to include a random number that’s generated every time a map variable is created. Next, they made the order of a for-range iteration over a map vary a bit each time the map is looped over. These two changes make it far harder to implement a Hash DoS attack.

}
