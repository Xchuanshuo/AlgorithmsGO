package leetcode.test855;

import java.util.*;

public class Test855 {

    class ExamRoom {

        TreeSet<int[]> tree;
        Map<Integer, int[]> start, end;
        int N;

        public ExamRoom(int n) {
            this.N = n;
            start = new HashMap<>();
            end = new HashMap<>();
            tree = new TreeSet<>((o1, o2) -> {
                int d1 = dis(o1), d2 = dis(o2);
                if (d1 == d2) {
                    return o1[0] - o2[0];
                }
                return d2 - d1;
            });
            add(new int[]{-1, N});
        }

        public int seat() {
            var cur = tree.first();
            var seat = 0;
            if (cur[0] == -1) {
                seat = 0;
            } else if (cur[1] == N) {
                seat = N-1;
            } else {
                seat = cur[0] + (cur[1]-cur[0])/2;
            }
            remove(cur);
            add(new int[]{cur[0], seat});
            add(new int[]{seat, cur[1]});
            return seat;
        }

        public void leave(int p) {
            var s = start.get(p);
            var e = end.get(p);
            remove(s);
            remove(e);
            add(new int[]{e[0], s[1]});
        }

        private void remove(int[] a) {
            tree.remove(a);
            start.remove(a[0]);
            end.remove(a[1]);
        }

        private void add(int[] a) {
            tree.add(a);
            start.put(a[0], a);
            end.put(a[1], a);
        }

        private int dis(int[] a) {
            int l = a[0], r = a[1];
            if (l == -1) return r;
            if (r == N) return N - 1  - l;
            return (r - l)/2;
        }
    }

}
