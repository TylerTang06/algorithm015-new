package week07

// two directions BFS
func minMutation(start string, end string, bank []string) int {
	if start == "" || end == "" || bank == nil || len(bank) == 0 {
		return -1
	}

	myMap := make(map[string]bool, len(bank))
	for _, b := range bank {
		myMap[b] = true
	}
	if _, ok := myMap[end]; !ok {
		return -1
	}
	delete(myMap, start)

	arr := []byte{'A', 'C', 'G', 'T'}
	visited := make(map[string]bool, len(bank))
	visited[start] = true
	beginMap, endMap := map[string]bool{start: true}, map[string]bool{end: true}
	step := 1
	for len(beginMap) > 0 && len(endMap) > 0 {
		if len(beginMap) > len(endMap) {
			beginMap, endMap = endMap, beginMap
		}

		curMap := map[string]bool{}
		for s := range beginMap {
			for i := 0; i < len(s); i++ {
				for j := 0; j < 4; j++ {
					if s[i] == arr[j] {
						continue
					}

					newS := s[:i] + string(arr[j]) + s[i+1:]
					if _, ok := myMap[newS]; ok {
						if _, yes := endMap[newS]; yes {
							return step
						}

						if _, yes := visited[newS]; !yes {
							visited[newS] = true
							curMap[newS] = true
						}
					}
				}
			}
		}
		step++
		beginMap = curMap
	}

	return -1
}
