package leetcode.test630;

import java.util.Arrays;
import java.util.Comparator;
import java.util.PriorityQueue;
import java.util.Queue;

public class Test630 {

    public int scheduleCourse(int[][] cs) {
        Arrays.sort(cs, Comparator.comparingInt(o -> o[1]));
        var pq = new PriorityQueue<Integer>((o1, o2) -> o2 - o1);
        var times = 0;
        for (int[] c : cs) {
            if (times + c[0] <= c[1]) {
                pq.offer(c[0]);
                times += c[0];
            } else if (!pq.isEmpty() && pq.peek() > c[0]) {
                // 用较少上课时间的课程替换掉时间较长的课程
                times -= pq.poll() - c[0];
                pq.offer(c[0]);
            }
        }
        return pq.size();
    }

}
