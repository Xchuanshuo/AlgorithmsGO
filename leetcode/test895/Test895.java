package leetcode.test895;

import java.util.*;

public class Test895 {

    class FreqStack {

        Map<Integer, T> map;
        TreeSet<T> set;
        int time = 0;

        public FreqStack() {
            set = new TreeSet<>(((o1, o2) -> {
                if (o1.cnt == o2.cnt) {
                    return o1.q.peekLast() - o2.q.peekLast();
                }
                return o1.cnt - o2.cnt;
            }));
        }

        public void push(int val) {
            time++;
            var cur = map.getOrDefault(val, new T(val));
            set.remove(cur);
            cur.q.offer(time);
            cur.cnt++;
            set.add(cur);
        }

        public int pop() {
            var maxT = set.pollLast();
            maxT.cnt--;
            maxT.q.pollLast();
            if (maxT.cnt > 0) {
                set.add(maxT);
            }
            return maxT.val;
        }
    }

    class T {
        int val, cnt;
        LinkedList<Integer> q;
        public T(int v) {
            this.val = v;
            this.cnt = 0;
            this.q = new LinkedList<>();
        }
    }
}
