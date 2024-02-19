package leetcode.test2336;

import java.util.*;

public class Test2336 {

    class SmallestInfiniteSet {

        Queue<Integer> pq;
        Set<Integer> set;

        public SmallestInfiniteSet() {
            this.pq = new PriorityQueue<>();
            this.set = new HashSet<>();
            for (int i = 1; i <= 1000;i++) {
                pq.offer(i);
                set.add(i);
            }
        }

        public int popSmallest() {
            var cur = pq.poll();
            set.remove(cur);
            return cur;
        }

        public void addBack(int num) {
            if (set.contains(num)) {
                return;
            }
            set.add(num);
            pq.offer(num);
        }
    }

}
