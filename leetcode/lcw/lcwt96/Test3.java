package leetcode.lcw.lcwt96;

import java.util.*;

public class Test3 {

    public long maxScore(int[] nums1, int[] nums2, int k) {
        var n = nums1.length;
        Integer[] idx = new Integer[n];
        for (int i = 0;i < idx.length;i++) {
            idx[i] = i;
        }
        Arrays.sort(idx, (o1, o2) -> nums2[o2] - nums2[o1]);
        Queue<Integer> pq = new PriorityQueue<>();
        long sum = 0, res = 0;
        for (int i = 0;i < n;i++) {
            sum += nums1[idx[i]];
            pq.offer(nums1[idx[i]]);
            if (pq.size() > k) {
                sum -= pq.poll();
            }
            res = Math.max(res, pq.size() < k ? 0 : sum * nums2[idx[i]]);
        }
        return res;
    }
}