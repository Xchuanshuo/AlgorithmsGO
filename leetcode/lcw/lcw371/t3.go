package lcw371

var INF = int(1e9)

func minOperations(nums1 []int, nums2 []int) int {
	var n = len(nums1)
	if n == 1 {
		return 0
	}
	var dfs func(idx int, rmx1, rmx2 int) int
	dfs = func(idx int, rmx1, rmx2 int) int {
		if idx == 0 {
			if nums1[0] <= rmx1 && nums2[0] <= rmx2 {
				return 0
			} else if nums2[0] <= rmx1 && nums1[0] <= rmx2 {
				return 1
			}
			return INF
		}
		if (nums1[idx] > rmx1 || nums2[idx] > rmx2) &&
			(nums2[idx] > rmx1 || nums1[idx] > rmx2) {
			return INF
		}
		if (nums1[idx] <= rmx1 && nums2[idx] <= rmx2) &&
			(nums2[idx] <= rmx1 && nums1[idx] <= rmx2) {
			return min(dfs(idx-1, max(rmx1, nums1[idx]), max(rmx2, nums2[idx])),
				1+dfs(idx-1, max(rmx1, nums2[idx]), max(rmx2, nums1[idx])))
		}
		if nums1[idx] <= rmx1 && nums2[idx] <= rmx2 {
			return dfs(idx-1, max(rmx1, nums1[idx]), max(rmx2, nums2[idx]))
		}
		return 1 + dfs(idx-1, max(rmx1, nums2[idx]), max(rmx2, nums1[idx]))
	}
	var res = min(dfs(n-2, nums1[n-1], nums2[n-1]), 1+dfs(n-2, nums2[n-1], nums1[n-1]))
	if res >= INF {
		return -1
	}
	return res
}