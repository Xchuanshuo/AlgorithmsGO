package leetcode.test901;

import java.util.Stack;

public class Test901 {

    class StockSpanner {

        Stack<int[]> s;
        int idx = 0;

        public StockSpanner() {
            s = new Stack<>();
            s.push(new int[]{-1, Integer.MAX_VALUE});
        }

        public int next(int price) {
            while (!s.isEmpty() && price >= s.peek()[1]) {
                s.pop();
            }
            var cur = idx - s.peek()[0];
            s.push(new int[]{idx, price});
            idx++;
            return cur;
        }
    }

}
