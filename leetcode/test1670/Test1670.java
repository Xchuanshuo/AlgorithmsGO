package leetcode.test1670;

import java.util.Deque;
import java.util.LinkedList;

/**
 * @Description https://leetcode.cn/problems/design-front-middle-back-queue/
 * idea:    1.使用两个双端队列
 *          2.封装一个函数adjust维护两个队列的diff
 *          3.每次操作后调用adjust即可
 **/

public class Test1670 {

    class FrontMiddleBackQueue {

        Deque<Integer> q1, q2;

        public FrontMiddleBackQueue() {
            this.q1 = new LinkedList<>();
            this.q2 = new LinkedList<>();
        }

        public void pushFront(int val) {
            q1.offerFirst(val);
            adjust();
        }

        public void pushMiddle(int val) {
            q1.offerLast(val);
            adjust();
        }

        public void pushBack(int val) {
            q2.offerLast(val);
            adjust();
        }

        private void adjust() {
            while (Math.abs(q1.size()-q2.size()) > 1) {
                if (q1.size() > q2.size()) {
                    var cur = q1.pollLast();
                    q2.offerFirst(cur);
                } else if (q1.size() < q2.size()) {
                    var cur = q2.pollFirst();
                    q1.offerLast(cur);
                }
            }
            if (q1.size() > q2.size()) {
                var cur = q1.pollLast();
                q2.offerFirst(cur);
            }
        }

        public int popFront() {
            if (q1.isEmpty() && q2.isEmpty()) {
                return -1;
            }
            var cur = 0;
            if (q1.isEmpty()) {
                cur = q2.pollFirst();
            } else {
                cur = q1.pollFirst();
            }
            adjust();
            return cur;
        }

        public int popMiddle() {
            if (q1.isEmpty() && q2.isEmpty()) {
                return -1;
            }
            var cur = 0;
            if (q1.size() < q2.size()) {
                cur = q2.pollFirst();
            } else {
                cur = q1.pollLast();
            }
            adjust();
            return cur;
        }

        public int popBack() {
            if (q1.isEmpty() && q2.isEmpty()) {
                return -1;
            }
            var cur = 0;
            if (q2.isEmpty()) {
                cur = q1.pollLast();
            } else {
                cur = q2.pollLast();
            }
            adjust();
            return cur;
        }
    }

}
