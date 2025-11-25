package main

/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
type CountPlus1 []int

func (cp1 CountPlus1) Set(i, count int) {
	cp1[i] = count + 1
}
func (cp1 CountPlus1) Get(i int) int {
	return cp1[i] - 1
}
func (cp1 CountPlus1) IsValid(i int) bool {
	return !(cp1[i] == 0)
}

func coinChange(coins []int, amount int) int {
	cp1 := make(CountPlus1, amount+1)
	cp1.Set(0, 0)
	for _, coin := range coins {
		if coin < amount {
			cp1.Set(coin, 1)
		} else if coin == amount {
			return 1
		}
	}
	for subamount := 1; subamount <= amount; subamount++ {
		minCount := subamount
		feasible := false
		for _, coin := range coins {
			if coin <= subamount {
				leftamount := subamount - coin
				if !cp1.IsValid(leftamount) {
					continue
				}
				feasibleCnt := cp1.Get(leftamount)
				if feasibleCnt < minCount {
					minCount = feasibleCnt
					feasible = true
				}
			}
		}
		if feasible {
			cp1.Set(subamount, minCount+1)
		}
	}
	return cp1.Get(amount)
}

// @lc code=end
