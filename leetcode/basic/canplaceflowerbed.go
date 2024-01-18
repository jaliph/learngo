package basic

// https://leetcode.com/problems/can-place-flowers
func CanPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
			if n == 0 {
				return true
			}
		}
	}

	return false
}
