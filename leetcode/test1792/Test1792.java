package leetcode.test1792;

import java.util.Arrays;
import java.util.PriorityQueue;
import java.util.Queue;

public class Test1792 {

    public double maxAverageRatio(int[][] cs, int stu) {
        Queue<int[]> pq = new PriorityQueue<>((o1, o2) -> {
           double p1 = o1[0]+1.0, p2 = o2[0]+1.0;
           double t1 = o1[1]+1.0, t2 = o2[1]+1.0;
           return (int) (p2/t2 - p1/t1);
        });
        for (int[] c : cs) {
            pq.offer(c);
        }
        while (stu > 0 && !pq.isEmpty()) {
            var cur = pq.poll();
            cur[0] += 1;
            cur[1] += 1;
            pq.offer(cur);
            stu--;
        }
        var n = cs.length;
        var res = 0.0;
        while (!pq.isEmpty()) {
            var cur = pq.poll();
            res += (double)cur[1]/(double)(cur[0])*1;
        }
        return res / n;
    }
}
