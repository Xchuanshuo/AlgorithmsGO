package leetcode.test1383;

import java.util.Arrays;
import java.util.PriorityQueue;
import java.util.Queue;

/**
 * @Description <a href="https://leetcode.cn/problems/maximum-performance-of-a-team/">...</a>
 * idea: 求序列 且需要看窗口内的最小值， 有两个变量 所以将1个排序后固定 另一个使用堆动态处理
 *         1.将效率值从大到小排序，保证处理当前时一定是最小的时间
 *         2.对于速度值维护一个小顶堆，保证选取人数>k时优先将较小的speed弹出
 *       类似题: test857
 **/

public class test1383 {

    public int maxPerformance(int n, int[] speed, int[] efficiency, int k) {
        int[][] arr = new int[n][2];
        for (int i = 0;i < n;i++) {
            arr[i][0] = speed[i];
            arr[i][1] = efficiency[i];
        }
        Arrays.sort(arr, (o1, o2) -> o2[1] - o1[1]);
        Queue<Integer> pq = new PriorityQueue<>();
        int M = (int)1e9 + 7;
        long res = 0, sum = 0;
        for (int i = 0;i < n;i++) {
            int[] cur = arr[i];
            if (i >= k) {
                // if ((sum - pq.peek() + cur[0])*cur[1] < res) {
                //     continue;
                // }
                sum -= pq.poll();
            }
            sum += cur[0];
            pq.offer(cur[0]);
            res = Math.max(res, sum * cur[1]);
        }
        return (int) (res%M);
    }

}
