package leetcode.lcw.lcw323;

import java.util.Arrays;
import java.util.Comparator;
import java.util.PriorityQueue;
import java.util.Queue;
import java.util.stream.IntStream;

/**
 * @Description https://leetcode.cn/problems/maximum-number-of-points-from-grid-queries/
 * idea: 优先队列分层遍历，限制条件为qv
 *
 **/

public class Test4 {

    int[][] dirs = {{1, 0}, {0, 1}, {-1, 0}, {0, -1}};

    public int[] maxPoints(int[][] grid, int[] queries) {
        var k = queries.length;
        var ids = IntStream.range(0, k).boxed().toArray(Integer[]::new);
        Arrays.sort(ids, Comparator.comparingInt(o -> queries[o]));
        int m = grid.length, n = grid[0].length;
        Queue<int[]> pq = new PriorityQueue<>(Comparator.comparingInt(o -> o[0]));
        var visited = new boolean[m][n];
        pq.offer(new int[]{grid[0][0], 0, 0});
        visited[0][0] = true;
        var res = new int[k];
        var cnt = 0;
        for (int id : ids) {
            var qv = queries[id];
            while (!pq.isEmpty() && pq.peek()[0] < qv) {
                cnt++;
                var cur = pq.poll();
                int x = cur[1], y = cur[2];
                for (int i = 0;i < dirs.length;i++) {
                    int nx = x+dirs[i][0], ny = y+dirs[i][1];
                    if (nx<0 || nx>=m || ny<0 || ny>=n || visited[nx][ny]) {
                        continue;
                    }
                    visited[nx][ny] = true;
                    pq.offer(new int[]{grid[nx][ny], nx, ny});
                }
            }
            res[id] = cnt;
        }
        return res;
    }
}
