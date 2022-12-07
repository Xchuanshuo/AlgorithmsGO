package leetcode.lcw.lcw322;

import java.util.*;

public class Test3 {

    public int minScore(int n, int[][] roads) {
        Map<Integer, List<int[]>> g = new HashMap<>();
        for (int[] e : roads) {
            if (!g.containsKey(e[0])) {
                g.put(e[0], new ArrayList<>());
            }
            if (!g.containsKey(e[1])) {
                g.put(e[1], new ArrayList<>());
            }
            g.get(e[0]).add(new int[]{e[1], e[2]});
            g.get(e[1]).add(new int[]{e[0], e[2]});
        }
        var dist = new int[n + 1];
        Arrays.fill(dist, Integer.MAX_VALUE);
        Queue<int[]> pq = new PriorityQueue<>(Comparator.comparingInt(o -> o[1]));
        pq.offer(new int[]{1, Integer.MAX_VALUE});
        while (!pq.isEmpty()) {
            int[] cur = pq.remove();
            // 需要考虑前面加入的节点 被后续更新的更加小, 此时该方案应该舍弃
            if (dist[cur[0]] < cur[1] || g.get(cur[0]) == null) continue;
            for (int[] nxt : g.get(cur[0])) {
                if (dist[nxt[0]] > Math.min(cur[1], nxt[1])) {
                    dist[nxt[0]] = Math.min(cur[1], nxt[1]);
                    pq.offer(new int[]{nxt[0], dist[nxt[0]]});
                }
            }
        }
        return dist[n];
    }

}
