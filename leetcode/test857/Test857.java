package leetcode.test857;

import java.util.Arrays;
import java.util.Comparator;
import java.util.PriorityQueue;
import java.util.Queue;

// q1*w2 <= w1*q2
public class Test857 {

    public double mincostToHireWorkers(int[] quality, int[] wage, int k) {
        int n = quality.length;
        int[][] arr = new int[n][2];
        for (int i = 0;i < n;i++) {
            arr[i][0] = wage[i];
            arr[i][1] = quality[i];
        }
        Arrays.sort(arr, (o1, o2) -> o1[0] * o2[1] - o1[1] * o2[0]);
        Queue<Integer> pq = new PriorityQueue<>((o1, o2) -> o2 - o1);
        double res = Double.MAX_VALUE, sum = 0;
        for (int i = 0;i < n;i++) {
            int[] cur = arr[i];
            sum -= i-k >= 0 ? pq.remove() : 0;
            sum += cur[1];
            pq.offer(cur[1]);
            double fac = (double) cur[0] / (double) cur[1];
            if (pq.size() == k) {
                res = Math.min(res, sum*fac);
            }
        }
        return res;
    }

}
