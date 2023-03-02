package leetcode.lcw.lcw331;

import java.util.PriorityQueue;
import java.util.Queue;

public class Test1 {

    public long pickGifts(int[] gifts, int k) {
        Queue<Integer> pq = new PriorityQueue<>((o1, o2) -> o2 - o1);
        for (int g : gifts) {
            pq.offer(g);
        }
        for (int i = 0;i < k;i++) {
            if (pq.isEmpty()) {
                break;
            }
            var cur = pq.poll();
            pq.offer((int) Math.sqrt(cur));
        }
        long res = 0;
        while (!pq.isEmpty()) {
            res += pq.poll();
        }
        return res;
    }
}
