package leetcode.lcw.lcw334;

import java.util.Comparator;
import java.util.PriorityQueue;

public class T4 {

    int[][] d = new int[][]{{0, 1}, {1, 0}, {0, -1}, {-1, 0}};

    public int minimumTime(int[][] g) {
        if (g[0][1] > 1 && g[1][0] > 1) {
            return -1;
        }
        int m = g.length, n = g[0].length;
        var visited = new boolean[m][n];
        var pq = new PriorityQueue<int[]>(Comparator.comparingInt(o -> o[2]));
        pq.offer(new int[]{0, 0, 0});
        visited[0][0] = true;
        while (!pq.isEmpty()) {
            var cur = pq.poll();
            int x = cur[0], y = cur[1], time = cur[2];
            if (x == m-1 && y == n-1) {
                return time;
            }
            for (int i = 0;i < d.length;i++) {
                int nx = x+d[i][0], ny = y+d[i][1];
                if (nx < 0 || nx >= m || ny < 0 || ny >= n || visited[nx][ny]) {
                    continue;
                }
                if (g[nx][ny] <= time + 1) {
                    pq.offer(new int[]{nx, ny, time+1});
                } else {
                    var ext = (time+g[nx][ny])%2 == 1 ? 0 : 1;
                    pq.offer(new int[]{nx, ny, g[nx][ny]+ ext});
                }
                visited[nx][ny] = true;
            }
        }
        return -1;
    }
}
