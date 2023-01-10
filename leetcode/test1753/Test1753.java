package leetcode.test1753;

import java.util.Arrays;
import java.util.Comparator;
import java.util.PriorityQueue;
import java.util.Queue;

public class Test1753 {

    public int maximumScore(int a, int b, int c) {
        Queue<Integer> pq = new PriorityQueue<>((o1, o2) -> o2 - o1);
        pq.addAll(Arrays.asList(a, b, c));
        var res = 0;
        while (!pq.isEmpty() && pq.size() >= 2) {
            var v1 = pq.poll();
            var v2 = pq.poll();
            var v3 = pq.peek();
            int dif = 0;
            if (v3 == null) {
                dif = v2;
            } else {
                dif = Math.min(v1.equals(v3)? 1 : v1-v3, v2);
            }
            v1 -= dif;
            v2 -= dif;
            res += dif;
            if (v1 > 0) { pq.offer(v1); }
            if (v2 > 0) { pq.offer(v2); }
        }
        return res;
    }

}
