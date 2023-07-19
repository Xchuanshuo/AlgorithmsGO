package leetcode.test2532;

import java.util.Arrays;
import java.util.Comparator;
import java.util.PriorityQueue;

public class Test2532 {

    public int findCrossingTime(int n, int k, int[][] time) {
        Arrays.sort(time, (o1, o2) -> o1[0] + o1[2] - (o2[0] + o2[2]));
        var pickR = new PriorityQueue<int[]>(Comparator.comparingInt(o -> o[1]));
        var putL = new PriorityQueue<int[]>(Comparator.comparingInt(o -> o[1]));
        var waitL = new PriorityQueue<int[]>((o1, o2) -> o2[0] - o1[0]);
        var waitR = new PriorityQueue<int[]>((o1, o2) -> o2[0] - o1[0]);
        for (int i = 0;i < k;i++) {
            waitL.offer(new int[]{i, 0});
        }
        var cur = 0;
        while (n > 0) {
            while (!pickR.isEmpty() && pickR.peek()[1] <= cur) {
                waitR.offer(pickR.poll());
            }
            while (!putL.isEmpty() && putL.peek()[1] <= cur) {
                waitL.offer(putL.poll());
            }
            if (!waitR.isEmpty()) {
                var p = waitR.poll();
                cur += time[p[0]][2];
                var nxt = cur + time[p[0]][3];
                putL.offer(new int[]{p[0], nxt});
            } else if (!waitL.isEmpty()) {
                var p = waitL.poll();
                cur += time[p[0]][0];
                var nxt = cur + time[p[0]][1];
                pickR.offer(new int[]{p[0], nxt});
                n--;
            } else if (!pickR.isEmpty() && !putL.isEmpty()) {
                cur = Math.min(pickR.peek()[1], putL.peek()[1]);
            } else if (!pickR.isEmpty()) {
                cur = pickR.peek()[1];
            } else if (!putL.isEmpty()) {
                cur = putL.peek()[1];
            }
        }
        while (!pickR.isEmpty()) {
            var p = pickR.poll();
            cur = Math.max(cur, p[1]) + time[p[0]][2];
        }
        return cur;
    }
}
