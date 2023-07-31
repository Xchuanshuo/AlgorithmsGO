package leetcode.test1354;

import java.util.PriorityQueue;

/**
 * @Description <a href="https://leetcode.cn/problems/construct-target-array-with-multiple-sums/">...</a>
 * idea: 最大的数一定是最后构造的，考虑逆向构造
 **/

public class Test1354 {

    public boolean isPossible(int[] target) {
        var pq = new PriorityQueue<Integer>((o1, o2) -> o2 - o1);
        var sum = 0L;
        for (int t : target) {
            pq.offer(t);
            sum += t;
        }
        while (pq.peek() != 1) {
            var max = pq.poll();
            var r = sum - max;
            if (r > max - 1 || r == 0) {
                return false;
            }
            if (r != 1 && max%r == 0) {
                return false;
            }
            var nxt = (int) (max%r);
            sum = r + nxt;
            pq.offer(nxt);
        }
        return true;
    }
}
