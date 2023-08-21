package leetcode.lcw.lcw357;

// 选长度为k的序列，不关心顺序 1.按profit排序 后取最后k个的值，并用set保存当前的type总数
//                        2.考虑加入前面[0,n-k-1]的元素，剔除后续的元素。
//                         若当前type与后k个重复，加入后不会使优雅度更加大，所以无需加入；
//                         若当前type与后k个不重复，且加入后，后k个需要剔除一个【不重复】的，优雅度也不会比之前更加大，不考虑加入；
//                         所以只在当前type与后k个不重复，且后k个有重复的情况下，才可能使值变大，此时考虑加入。
//                         对于剔除操作，总是选择剔除重复元素中较小的，能让结果更优

import java.util.*;

public class Test6932 {

    public long findMaximumElegance(int[][] items, int k) {
        var n = items.length;
        var sum = 0L;
        Arrays.sort(items, (o1, o2) -> o1[0] - o2[0]);
        var set = new HashSet<Integer>();
        var st = new Stack<Integer>();
        for (int i = n-1;i >= n-k;i--) {
            int v = items[i][0], t = items[i][1];
            sum += v;
            if (!set.contains(t)) {
                set.add(t);
            } else {
                st.push(v);
            }
        }
        var res = sum + (long) set.size() *set.size();
        for (int i = n-k-1;i >= 0;i--) {
            int v = items[i][0], t = items[i][1];
            // 添加新的type后与原来总数相同 得出的优雅度一定不会比之前大
            if (set.contains(t) || st.isEmpty()) {
                continue;
            }
            set.add(t);
            sum += v - st.pop();
            res = Math.max(res, sum + (long) set.size() *set.size());
        }
        return res;
    }
}
