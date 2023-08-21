package leetcode.test1851;

import java.util.Arrays;
import java.util.Comparator;
import java.util.PriorityQueue;

/**
 * @Description <a href="https://leetcode.cn/problems/minimum-interval-to-include-each-query/">...</a>
 * idea: 关键点: 1.查询排序，对于当前查询的位置p，所有左边界<p的区间入对，入对后右区间<p的对后续查询也无效，出队
 *              2.队列的元素取最小长度的
 **/

public class Test1851 {

    public int[] minInterval(int[][] invs, int[] queries) {
        Arrays.sort(invs, Comparator.comparingInt(o -> o[0]));
        var qs = new int[queries.length][2];
        for (int i = 0;i < queries.length;i++) {
            qs[i] = new int[]{queries[i], i};
        }
        Arrays.sort(qs, Comparator.comparingInt(o -> o[0]));
        var pq = new PriorityQueue<int[]>(Comparator.comparingInt(o -> (o[1] - o[0])));
        var res = new int[qs.length];
        var i = 0;
        for (var q : qs) {
            var id = q[1];
            while (i < invs.length && invs[i][0] <= q[0]) {
                pq.offer(new int[]{invs[i][0], invs[i][1]});
                i++;
            }
            while (!pq.isEmpty() && pq.peek()[1] < q[0]) {
                pq.poll();
            }
            if (pq.isEmpty()) {
                res[id] = -1;
            } else {
                res[id] = pq.peek()[1] - pq.peek()[0]+1;
            }
        }
        return res;
    }
}
