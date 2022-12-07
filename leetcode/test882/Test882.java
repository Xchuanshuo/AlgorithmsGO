package leetcode.test882;

import java.util.*;

/**
 * @Description https://leetcode.cn/problems/reachable-nodes-in-subdivided-graph/
 * idea: 1.边权值为 cnt 单源最短路, dijkstra
 *       2.统计可达节点个数，关键，防止【重复】计算，
 *          对于[0,n-1]的点i，若maxMoves>=dist[0][i]，则该点可计入答案
 *          对于细分点，扫描每条边 从0到达该边两端点后，还能走多远，与cnt取min即可
 **/

public class Test882 {

    public int reachableNodes(int[][] edges, int maxMoves, int n) {
        Map<Integer, List<int[]>> g = new HashMap<>();
        for (int[] e : edges) {
            if (!g.containsKey(e[0])) {
                g.put(e[0], new ArrayList<>());
            }
            if (!g.containsKey(e[1])) {
                g.put(e[1], new ArrayList<>());
            }
            g.get(e[0]).add(new int[]{e[1], e[2]+1});
            g.get(e[1]).add(new int[]{e[0], e[2]+1});
        }
        var dist = new int[n];
        Arrays.fill(dist, Integer.MAX_VALUE);
        dist[0] = 0;
        Queue<int[]> pq = new PriorityQueue<>(Comparator.comparingInt(o -> o[1]));
        pq.offer(new int[]{0, 0});
        while (!pq.isEmpty()) {
            int[] cur = pq.remove();
            // 需要考虑前面加入的节点 被后续更新的更加小, 此时该方案应该舍弃
            if (dist[cur[0]] < cur[1] || g.get(cur[0]) == null) continue;
            for (int[] nxt : g.get(cur[0])) {
                if (dist[nxt[0]] > dist[cur[0]] + nxt[1]) {
                    dist[nxt[0]] = dist[cur[0]] + nxt[1];
                    pq.offer(new int[]{nxt[0], dist[nxt[0]]});
                }
            }
        }
        var res = 0;
        for (int i = 0;i < n;i++) {
            if (dist[i] <= maxMoves) {
                res++;
            }
        }
        for (int[] e : edges) {
            int p = e[0], q = e[1], w = e[2];
            var a = Math.max(maxMoves-dist[p], 0);
            var b = Math.max(maxMoves-dist[q], 0);
            res += Math.min(a + b, w);
        }
        return res;
    }
}
