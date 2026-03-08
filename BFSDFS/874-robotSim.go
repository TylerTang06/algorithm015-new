package week04

func robotSim(commands []int, obstacles [][]int) int {
	if commands == nil || len(commands) == 0 {
		return 0
	}

	// use map to store position of obstacles so that use O(1) time to judge whether meet obstacles
	myMap := make(map[[2]int]bool, len(obstacles))
	for _, arr := range obstacles {
		myMap[[2]int{arr[0], arr[1]}] = true
	}

	// good skill to record the direction
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	currX, currY, curDir, res := 0, 0, 0, 0
	for _, val := range commands {
		if val == -1 {
			curDir = (curDir + 1) % 4
		}
		if val == -2 {
			curDir = (curDir + 3) % 4
		}
		if val >= 1 && val <= 9 {
			for i := 0; i < val; i++ {
				if !myMap[[2]int{currX + dx[curDir], currY + dy[curDir]}] {
					currX += dx[curDir]
					currY += dy[curDir]
					// get max result
					if res < currX*currX+currY*currY {
						res = currX*currX + currY*currY
					}
				} else {
					break
				}
			}
		}
	}

	return res
}
