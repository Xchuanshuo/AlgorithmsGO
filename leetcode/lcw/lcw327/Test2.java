package leetcode.lcw.lcw327;

import java.util.PriorityQueue;
import java.util.Queue;

public class Test2 {

    public long maxKelements(int[] nums, int k) {
        Queue<Integer> pq = new PriorityQueue<>((o1, o2) -> o2 - o1);
        for (int num : nums) {
            pq.offer(num);
        }
        long res = 0;
        while (!pq.isEmpty() && k > 0) {
            k--;
            var cur = pq.poll();
            res += cur;
            pq.offer((int)(Math.ceil((double)cur / 3)));
        }
        return res;
    }
}


//输入：
//        "ab"
//        "abcc"
//        输出：
//        true
//        预期：
//        false