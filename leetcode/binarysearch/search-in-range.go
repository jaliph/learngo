//https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
package binarysearch

func SearchRange(nums []int, target int) []int {
    if len(nums) == 0 {
        return []int{-1, -1}
    }
    l, r := 0, len(nums) - 1
    var mid int

    for l < r {
        mid = l + ((r - l) / 2)
        if target > nums[mid] {
            l = mid + 1
        } else {
            r = mid
        }
    }
    if nums[l] != target {
        return []int{-1, -1}
    }
    res := []int{l}

    l, r = 0, len(nums) - 1
     for l <= r {
        mid = l + ((r - l) / 2)
        if target == nums[mid] {
            l = mid + 1
        } else if target > nums[mid] {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    return append(res, r)
}
