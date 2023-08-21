package leetcode.test1847;

// 排序 + 二分(有序集合)

import java.util.*;

public class Test1847 {

    public int[] closestRoom(int[][] rooms, int[][] queries) {
        var list = new LinkedList<int[]>();
        for (var r : rooms) {
            list.add(new int[]{r[0], r[1], -1});
        }
        for (int i = 0;i < queries.length;i++) {
            list.add(new int[]{queries[i][0], queries[i][1], i});
        }
        list.sort((o1, o2) -> {
            if (o1[1] == o2[1]) {
                return o1[2] - o2[2];
            }
            return o2[1] - o1[1];
        });
        var res = new int[queries.length];
        var tree = new TreeSet<Integer>();
        for (int[] cur : list) {
            int id = cur[0], t = cur[2];
            if (t < 0) {
                tree.add(id);
                continue;
            }
            int tid = -1, val = Integer.MAX_VALUE;
            var r = tree.ceiling(id);
            if (r != null) {
                val = Math.min(val, Math.abs(id - r));
                tid = r;
            }
            var l = tree.floor(id);
            if (l != null && Math.abs(id-l) <= val) {
                tid = l;
            }
            res[t] = tid;
        }
        return res;
    }
}
