package week07

import "container/list"

func findCircleNum(M [][]int) int {
	if M == nil || len(M) == 0 || len(M[0]) == 0 {
		return 0
	}

	res := 0
	// BFS
	myQue := list.New()
	visited := make([]int, len(M))
	for i := 0; i < len(M); i++ {
		if visited[i] == 0 {
			myQue.PushBack(i)
			visited[i] = 1
		} else {
			continue
		}

		for myQue.Len() > 0 {
			k := myQue.Front().Value.(int)
			myQue.Remove(myQue.Front())
			for j := 0; j < len(M); j++ {
				if M[k][j] == 1 && visited[j] == 0 {
					visited[j] = 1
					myQue.PushBack(j)
				}
			}
		}
		res++
	}

	return res
}
