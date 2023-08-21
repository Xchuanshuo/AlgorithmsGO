package leetcode.test1851.s2;

import java.util.ArrayList;
import java.util.Comparator;
import java.util.PriorityQueue;

public class Test1851 {

    public int[] minInterval(int[][] intervals, int[] queries) {
        var list = new ArrayList<int[]>();
        for (int[] v : intervals) {
            list.add(new int[]{v[0], v[1], -1});
        }
        for (int i = 0;i < queries.length;i++) {
            list.add(new int[]{queries[i],queries[i], i});
        }
        list.sort((o1, o2) -> {
            if (o1[0] == o2[0]) {
                return o1[2] - o2[2];
            }
            return o1[0] - o2[0];
        });
        var pq = new PriorityQueue<int[]>(Comparator.comparingInt(o -> (o[1] - o[0])));
        var res = new int[queries.length];
        for (int[] v : list) {
            int l = v[0], t = v[2];
            if (t == -1) {
                pq.offer(v);
                continue;
            }
            res[t] = -1;
            while (!pq.isEmpty() && pq.peek()[1] < l) {
                pq.poll();
            }
            if (!pq.isEmpty()) {
                res[t] = pq.peek()[1] - pq.peek()[0] + 1;
            }
        }
        return res;
    }
}
