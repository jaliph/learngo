/**
 * https://leetcode.com/problems/find-in-mountain-array/
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */
package binarysearch

type MountainArray struct {
}

func (this *MountainArray) get(index int) int { return 0 }
func (this *MountainArray) length() int       { return 0 }

func FindInMountainArray(target int, mountainArr *MountainArray) int {
	l, r := 1, mountainArr.length()-2 // -2 because we know mountain has min 3 length
	var mid int
	for l <= r {
		mid = l + (r-l)/2

		m1, m2, m3 := mountainArr.get(mid-1), mountainArr.get(mid), mountainArr.get(mid+1)
		if m1 < m2 && m2 < m3 {
			l = mid + 1
		} else if m1 > m2 && m2 > m3 {
			r = mid - 1
		} else {
			break
		}
	}

	peak := mid

	// Search in the first portion
	l, r = 0, peak
	for l <= r {
		mid = l + (r-l)/2
		m := mountainArr.get(mid)
		if target == m {
			return mid
		} else if target > m {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	// Search in the second portion

	l, r = peak, mountainArr.length()-1
	for l <= r {
		mid = l + (r-l)/2
		m := mountainArr.get(mid)
		if target == m {
			return mid
		} else if target > m {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}
