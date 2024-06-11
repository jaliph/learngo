// https://leetcode.com/problems/minimum-replacements-to-sort-the-array/

package greedy

/*
Consider this array [3,9,3]

In the first step we make out prev=3
Now we iterate backwards, now our value becomes 9 which is more than our prev value, so we have to break this in such a way that the max element in this is less than or equal to prev element and min element must as greater as possible. Why min element must be greater becauase we are following greedy approach we can easily solve further for elements.
How many splits we can do is the question?

9 -> 3 3 3 //to satisfy the condition.

split-1:9->6 3
split-2:9->3 3 3

Total Splits = Math.ceil(9/3) - 1 //becuase you have splitted it twice.

    Next task is what is the min value we got after splitting

splits = 2
value   = 9/(2+1) = 3
Now prev = value

    If out current value is simply smaller than prev element then we can assign it to the prev = currElement and move forward

Complexity

    Time complexity:
    O(N)

    Space complexity:
    O(1)
*/

func MinimumReplacement(nums []int) int64 {
	prev := nums[len(nums)-1]

	Ceil := func(a, b int) int {
		if a%b == 0 {
			return a / b
		}
		return (a / b) + 1
	}

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var res int64 = 0
	for i := len(nums) - 2; i >= 0; i-- {
		if prev >= nums[i] {
			prev = nums[i]
		} else {
			parts := Ceil(nums[i], prev) - 1
			res += int64(parts)
			prev = Min(prev, nums[i]/(parts+1))
		}
	}

	return res
}
