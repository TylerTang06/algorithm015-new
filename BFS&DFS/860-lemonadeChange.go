package week04

func lemonadeChange(bills []int) bool {
	if bills == nil || len(bills) == 0 {
		return true
	}

	count5, count10 := 0, 0
	for _, bill := range bills {
		if bill == 5 {
			count5++
		}
		if bill == 10 {
			if count5 >= 1 {
				count5--
				count10++
			} else {
				return false
			}
		}
		if bill == 20 {
			if count10 >= 1 && count5 >= 1 {
				count5--
				count10--
			} else if count10 == 0 && count5 >= 3 {
				count5 -= 3
			} else {
				return false
			}
		}
	}

	return true
}
