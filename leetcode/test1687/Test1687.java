package leetcode.test1687;

import java.util.*;

/**
 * @Description <a href="https://leetcode.cn/problems/delivering-boxes-from-storage-to-ports/">...</a>
 * idea: dp + 优先队列
 * 1.设dp[i]表示前i个箱子送到相应码头的最小花费
 * 2.dp[i]= dp[j] + sums[i]-sums[j+1]+2  可选区间为[j+1..i]
 * 枚举j取值区间即可，但此时时间复杂度为O(N^2)
 * 注意点: 因为区间数
 * 优化: 观察sums[i]+2为当前状态计算的【固定值】，在前面【所有】不超过maxB和maxW的区间，
 * dp[j]-sums[j+1]中取最小值即可
 * 实现: 使用优先队列 / 单调队列维护最小值
 **/

public class Test1687 {

    public int boxDelivering(int[][] b, int ports, int maxB, int maxW) {
        var n = b.length;
        var dp = new int[n + 1];
        Arrays.fill(dp, Integer.MAX_VALUE);
        Queue<int[]> pq = new PriorityQueue<>(Comparator.comparingInt(o -> o[2]));
        pq.offer(new int[]{0, 0, 0}); // 箱子索引, 容量, dp[j]值
        int sum1 = 0, sum2 = 0;
        for (int i = 1; i <= n; i++) {
            sum2 += (i - 2) >= 0 ? (b[i - 1][0] == b[i - 2][0] ? 0 : 1) : 0;
            pq.offer(new int[]{i - 1, sum1, dp[i - 1] - sum2});
            sum1 += b[i - 1][1];
            while (!pq.isEmpty() && (i - pq.peek()[0] > maxB || sum1 - pq.peek()[1] > maxW)) {
                pq.poll();
            }
            dp[i] = Math.min(dp[i], pq.peek()[2] + sum2 + 2);
        }
        return dp[n];
    }

    public int boxDelivering1(int[][] b, int ports, int maxB, int maxW) {
        var n = b.length;
        var dp = new int[n + 1];
        Arrays.fill(dp, Integer.MAX_VALUE);
        Deque<int[]> q = new LinkedList<>();
        q.offer(new int[]{0, 0, 0}); // 箱子索引, 容量, dp[j]值
        int sum1 = 0, sum2 = 0;
        for (int i = 1; i <= n; i++) {
            sum2 += (i - 2) >= 0 ? (b[i - 1][0] == b[i - 2][0] ? 0 : 1) : 0;
            var cur = new int[]{i - 1, sum1, dp[i - 1] - sum2};
            sum1 += b[i - 1][1];
            // 1.维护单调性
            while (!q.isEmpty() && q.peekLast()[2] > cur[2]) {
                q.pollLast();
            }
            q.offerLast(cur);
            // 2.维护合法的窗口
            while (!q.isEmpty() && (i - q.peekFirst()[0] > maxB || sum1 - q.peekFirst()[1] > maxW)) {
                q.pollFirst();
            }
            dp[i] = Math.min(dp[i], q.peekFirst()[2] + sum2 + 2);
        }
        return dp[n];
    }

}
