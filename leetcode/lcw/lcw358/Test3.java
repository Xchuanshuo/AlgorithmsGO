package leetcode.lcw.lcw358;

import java.util.List;
import java.util.TreeSet;

public class Test3 {

    public int minAbsoluteDifference(List<Integer> nums, int x) {
        var tree = new TreeSet<Integer>();
        var res = Math.abs(nums.get(0) - nums.get(nums.size()-1));
        var idx = 0;
        for (int i = x;i < nums.size();i++) {
            tree.add(nums.get(idx++));
            int cur = nums.get(i);
            var l = tree.floor(cur);
            if (l != null) {
                res = Math.min(res, Math.abs(l - cur));
            }
            var r = tree.ceiling(cur);
            if (r != null) {
                res = Math.min(res, Math.abs(r - cur));
            }

        }
        return res;
    }
}
