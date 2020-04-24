package main

func lemonadeChange(bills []int) bool {
	moneys := [3]int{}
	for _, x := range bills {
		if x == 5 {
			moneys[0]++
		} else if x == 10 {
			if moneys[0] > 0 {
				moneys[0]--
				moneys[1]++
			} else {
				return false
			}
		} else {
			if moneys[1] > 0 && moneys[0] > 0 {
				moneys[0]--
				moneys[1]--
				moneys[2]++
			} else if moneys[0] >= 3 {
				moneys[0] -= 3
				moneys[2]++
			} else {
				return false
			}
		}
	}
	return true
}
