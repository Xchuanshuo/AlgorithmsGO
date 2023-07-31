package leetcode.test1499;

import java.util.PriorityQueue;

public class Test1499 {

    // yi-xi + yj+xj
    public int findMaxValueOfEquation(int[][] points, int k) {
        var pq = new PriorityQueue<int[]>((o1, o2) -> o2[0] - o1[0]);
        var res = Integer.MIN_VALUE;
        for (int[] p : points) {
            int x = p[0], y = p[1];
            while (!pq.isEmpty() && Math.abs(x - pq.peek()[1]) > k) {
                pq.poll();
            }
            if (!pq.isEmpty()) {
                res = Math.max(res, x + y + pq.peek()[0]);
            }
            pq.offer(new int[]{y-x, x});
        }
        return res;
    }

}
