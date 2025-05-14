package group_anagrams

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

func GroupAnagrams(strs []string) [][]string {
	anagrams := make(map[int][]string);
	usedWords := make(map[int]bool);
	for index1, str1 := range strs {
		for index2, str2 := range strs {
			if IsAnagram(str1, str2) && index1 != index2 && !usedWords[index2] {
				if array, exists := anagrams[index1]; exists {
					array = append(array, str2);
					anagrams[index1] = array;
				} else {
					anagrams[index1] = []string{str1,str2};
				}
				usedWords[index2] = true;
				usedWords[index1] = true;
			}
		}

		if !usedWords[index1] {
			anagrams[index1] = []string{str1};
		}
	}

	result := make([][]string, 0, len(anagrams));
	for _, value := range anagrams {
		result = append(result, value);
	}

	return result;
}
