class Solution:

    def __init__(self):
        pass

    def minOperations(self, nums1, nums2):
        INF = 1000000000
        n = len(nums1)
        if n == 1:
            return 0
        @cache
        def dfs(idx, rmx1, rmx2):
            if idx == 0:
                if nums1[0] <= rmx1 and nums2[0] <= rmx2:
                    return 0
                elif nums2[0] <= rmx1 and nums1[0] <= rmx2:
                    return 1
                return INF

            if (nums1[idx] > rmx1 or nums2[idx] > rmx2) and (nums2[idx] > rmx1 or nums1[idx] > rmx2):
                return INF
            if (nums1[idx] <= rmx1 and nums2[idx] <= rmx2) and (nums2[idx] <= rmx1 and nums1[idx] <= rmx2):
                return min(dfs(idx-1, max(rmx1, nums1[idx]), max(rmx2, nums2[idx])), 1 + dfs(idx-1, max(rmx1, nums2[idx]), max(rmx2, nums1[idx])))
            if nums1[idx] <= rmx1 and nums2[idx] <= rmx2:
                return dfs(idx-1, max(rmx1, nums1[idx]), max(rmx2, nums2[idx]))
            return 1 + dfs(idx-1, max(rmx1, nums2[idx]), max(rmx2, nums1[idx]))

        res = min(dfs(n-2, nums1[-1], nums2[-1]), 1 + dfs(n-2, nums2[-1], nums1[-1]))
        if res >= INF:
            return -1
        return res