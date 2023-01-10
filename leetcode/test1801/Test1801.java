package leetcode.test1801;

import java.util.Comparator;
import java.util.PriorityQueue;
import java.util.Queue;

/**
 * @Description <a href="https://leetcode.cn/problems/number-of-orders-in-the-backlog/">...</a>
 * idea: 1.buy订单，min_sell <= buy, del sell;
 *       2.sell订单，max_buy >= sell, del buy
 **/

public class Test1801 {

    public int getNumberOfBacklogOrders(int[][] orders) {
        Queue<int[]> sellPQ = new PriorityQueue<>(Comparator.comparingInt(o -> o[0]));
        Queue<int[]> buyPQ = new PriorityQueue<>((o1, o2) -> o2[0] - o1[0]);
        int Mod = (int) 1e9+ 7;
        for (int[] o : orders) {
            if (o[2] == 0) { // buy
                process(o, sellPQ, buyPQ, true);
            } else { // sell
                process(o, buyPQ, sellPQ, false);
            }
        }
        var res = 0;
        while (!sellPQ.isEmpty()) {
            res = (res + sellPQ.poll()[1]) % Mod;
        }
        while (!buyPQ.isEmpty()) {
            res = (res + buyPQ.poll()[1]) % Mod;
        }
        return res;
    }

    private void process(int[] o, Queue<int[]> pq1, Queue<int[]> pq2, boolean f) {
        while (!pq1.isEmpty() && o[1] > 0 &&
                (f ? o[0] >= pq1.peek()[0] : pq1.peek()[0] >= o[0]) ) {
            var cur = pq1.poll();
            var cnt = Math.min(o[1], cur[1]);
            cur[1] -= cnt;
            o[1] -= cnt;
            if (cur[1] > 0) {
                pq1.offer(cur);
            }
        }
        if (o[1] > 0) {
            pq2.offer(new int[]{o[0], o[1]});
        }
    }

    public static void main(String[] args) {
        int[][] arr = new int[][]{{19,28,0},{9,4,1},{25,15,1}};
        Test1801 t = new Test1801();
        var res =t.getNumberOfBacklogOrders(arr);
        System.out.println(res);
    }
}
