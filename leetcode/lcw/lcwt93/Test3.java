package leetcode.lcw.lcwt93;

import com.sun.source.tree.Tree;

import java.util.Arrays;
import java.util.TreeSet;

public class Test3 {

    public int maxJump(int[] stones) {
        var n = stones.length;
        int l = 0, r = stones[n-1];
        TreeSet<Integer> treeSet = new TreeSet<>();
        for (int a : stones) {
            treeSet.add(a);
        }
        while (l <= r) {
            var mid = l + (r-l)/2;
            if (isValid((TreeSet<Integer>) treeSet.clone(), mid)[0] == 0) {
                l = mid + 1;
            } else {
                if (mid == 0 || isValid((TreeSet<Integer>) treeSet.clone(), mid-1)[0] == 0) {
                    return isValid((TreeSet<Integer>) treeSet.clone(), mid)[1];
                }
                r = mid - 1;
            }
        }
        return 0;
    }

    private int[] isValid(TreeSet<Integer> treeSet, int t) {
        var mv = treeSet.last();
        int cur = 0, res = 0;
        while (cur != mv) {
            var nxt = cur + t;
            var v = treeSet.floor(nxt);
            if (v == null) {
                return new int[]{0, 0};
            }
            treeSet.remove(v);
            res = Math.max(res, v - cur);
            cur = v;
        }
        cur = mv;
        while (cur != 0) {
            var nxt = cur - t;
            var v = treeSet.ceiling(nxt);
            if (v == null || v > cur) {
                return new int[]{0, 0};
            }
            treeSet.remove(v);
            res = Math.max(res, cur - v);
            cur = v;
        }
        return new int[]{1, res};
    }

    public static void main(String[] args) {
        int[] arr = new int[]{0, 3};
        Test3 test3 = new Test3();
        System.out.println(test3.maxJump(arr));
    }
}
