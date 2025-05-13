package valid_anagram

func IsAnagram(s string, t string) bool {
	myMap1 := make(map[rune]int);
	myMap2 := make(map[rune]int);
	for _, value := range s {
		myMap1[value]++;
	}

	for _, value := range t {
		myMap2[value]++;
	}

	if len(myMap1) != len(myMap2) {
		return false;
	}

	for key, value := range myMap1 {
		if myMap2[key] != value {
			return false;
		}
	}

	return true;
}
