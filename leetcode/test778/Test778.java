package leetcode.test778;

import java.util.Comparator;
import java.util.PriorityQueue;
import java.util.Queue;

public class Test778 {

    int[][] dirs = new int[][]{{1, 0}, {0, 1}, {-1, 0}, {0, -1}};
    public int swimInWater(int[][] grid) {
        var n = grid.length;
        Queue<int[]> pq = new PriorityQueue<>(Comparator.comparingInt(o -> o[0]));
        boolean[][] visited = new boolean[n][n];
        pq.offer(new int[]{grid[0][0], 0, 0});
        visited[0][0] = true;
        for (int i = 0;i < n*n;i++) {
            while (!pq.isEmpty() && pq.peek()[0]<=i) {
                var cur = pq.poll();
                if (cur[1]==n-1 && cur[2]==n-1) {
                    return i;
                }
                for (int j = 0;j < dirs.length;j++) {
                    int nx = cur[1]+dirs[j][0], ny = cur[2]+dirs[j][1];
                    if (nx<0 || nx>=n || ny<0 || ny>=n || visited[nx][ny]) {
                        continue;
                    }
                    visited[nx][ny] = true;
                    pq.offer(new int[]{grid[nx][ny], nx, ny});
                }
            }
        }
        return n*n;
    }
}
