package leetcode.test1882;

import java.util.Comparator;
import java.util.PriorityQueue;

public class Test1882 {

    public int[] assignTasks(int[] servers, int[] tasks) {
        var free = new PriorityQueue<Integer>((o1, o2) -> {
            if (servers[o1] == servers[o2]) {
                return o1 - o2;
            }
            return servers[o1]-servers[o2];
        });
        var busy = new PriorityQueue<int[]>(Comparator.comparingInt(o -> o[1]));
        for (int i = 0;i < servers.length;i++) {
            free.offer(i);
        }
        var res = new int[tasks.length];
        var cur = 0;
        for (int i = 0;i < tasks.length;i++) {
            cur = Math.max(cur, i);
            if (free.isEmpty()) {
                var p = busy.poll();
                cur = Math.max(cur, p[1]);
                free.offer(p[0]);
            }
            while (!busy.isEmpty() && busy.peek()[1] <= cur) {
                free.offer(busy.poll()[0]);
            }
            var top = free.poll();
            res[i] = top;
            busy.offer(new int[]{top, cur + tasks[i]});
        }
        return res;
    }
}
