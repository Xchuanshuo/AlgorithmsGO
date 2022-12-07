package leetcode.lcw.lcw318;

import java.util.PriorityQueue;

public class Test3 {

    public long totalCost(int[] costs, int k, int c) {
        PriorityQueue<Integer> q1 = new PriorityQueue<>((o1, o2) -> {
            if (costs[o1] == costs[o2]) {
                return o1 - o2;
            }
            return costs[o1] - costs[o2];
        });
        PriorityQueue<Integer> q2 = new PriorityQueue<>((o1, o2) -> {
            if (costs[o1] == costs[o2]) {
                return o1 - o2;
            }
            return costs[o1] - costs[o2];
        });
        int n = costs.length;
        for (int i = 0;i < Math.min(c, n);i++) {
            q1.offer(i);
        }
        for (int i = n-1;i >= Math.max(c, n - c);i--) {
            q2.offer(i);
        }
        int i = c, j = n - c - 1;
        long res = 0;
        for (int t = 0;t < k;t++) {
            if (!q1.isEmpty() && !q2.isEmpty()) {
                if (costs[q2.peek()] < costs[q1.peek()]) {
                    res += costs[q2.poll()];
                    if (i <= j) {
                        q2.offer(j--);
                    }
                } else {
                    res += costs[q1.poll()];
                    if (i <= j) {
                        q1.offer(i++);
                    }
                }
            } else if (!q1.isEmpty()) {
                res += costs[q1.poll()];
                if (i <= j) {
                    q1.offer(i++);
                }
            } else if (!q2.isEmpty()) {
                res += costs[q2.poll()];
                if (i <= j) {
                    q2.offer(j--);
                }
            }
        }
        return res;
    }

    public static void main(String[] args) {
        int[] arr = new int[]{18,64,12,21,21,78,36,58,88,58,99,26,92,91,53,10,24,25,20,92,73,63,51,65,87,6,17,32,14,42,46,65,43,9,75};
        int k = 13, c = 23;
        Test3 test3 = new Test3();
        System.out.println(test3.totalCost(arr, k, c));
    }
}