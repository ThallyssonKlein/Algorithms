package contains_duplicate

func ContainsDuplicate(nums []int) bool {
	myMap := make(map[int]bool);

	for _, value := range nums {
		if myMap[value] == true {
			return true;
		}
		myMap[value] = true;
	}

	return false;
} 