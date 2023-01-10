package leetcode.lcw.lcw327;

/**
 * @Description https://leetcode.cn/problems/time-to-cross-a-bridge/
 * idea: 优先队列模拟 + 分情况讨论
 *      设置4个队列 put、pick, 分别表示放置货物到新仓库、从旧仓库搬运货物工人队列
 *                waitL、waitR, 分别表示桥左边，右边等待工人队列
 *      对于过桥，时间是串行，而搬运和放置货物不同工人时间可以并行
 *      记录当前时间cur，直到把全部货物搬完
 *      对于put、和pick小于当前时间的扔入对应等待队列即可
 *      1.若当前有工人搬运货物后待过桥，优先让工人过桥
 *      2.若当前有剩余货物和剩余工人, 则让工人过桥搬货物
 *      3.均不满足说明工人出于忙碌，考虑等待最早完成时间的工人释放
 *
 *      关键: 1.保证取数据是当前处于就绪状态 2.没有就绪状态时, 选取一个最新到达的事件
 **/

import java.util.Arrays;
import java.util.Comparator;
import java.util.PriorityQueue;

public class Test4 {

    public int findCrossingTime(int n, int k, int[][] time) {
        // 提前固定工人顺序
        Arrays.sort(time, Comparator.comparingInt(o -> o[0] + o[2]));
        var pick = new PriorityQueue<int[]>(Comparator.comparingInt(o -> o[1]));
        var put = new PriorityQueue<int[]>(Comparator.comparingInt(o -> o[1]));
        var waitL = new PriorityQueue<int[]>((o1, o2) -> o2[0] - o1[0]);
        var waitR = new PriorityQueue<int[]>((o1, o2) -> o2[0] - o1[0]);
        for (int i = 0;i < time.length;i++) {
            waitL.offer(new int[]{i, 0});
        }
        var cur = 0;
        while (n > 0) {
            while (!pick.isEmpty() && pick.peek()[1] <= cur) waitR.offer(pick.poll());
            while (!put.isEmpty() && put.peek()[1] <= cur) waitL.offer(put.poll());
            if (!waitR.isEmpty() && waitR.peek()[1] <= cur) {
                var p = waitR.poll();
                cur += time[p[0]][2];
                var nxt = cur + time[p[0]][3];
                var worker = new int[]{p[0], nxt};
                put.offer(worker);
            } else if (!waitL.isEmpty() && waitL.peek()[1] <= cur) {
                var p = waitL.poll();
                cur += time[p[0]][0];
                var nxt = cur + time[p[0]][1];
                var worker = new int[]{p[0], nxt};
                pick.offer(worker);
                n--;
            } else if (!pick.isEmpty() && !put.isEmpty()) {
                cur = Math.min(pick.peek()[1], put.peek()[1]);
            } else if (!pick.isEmpty()) {
                cur = pick.peek()[1];
            } else if (!put.isEmpty()){
                cur = put.peek()[1];
            }
        }
        // 需要将前面的工人过完桥的时间与当前工人pick的时间取较大值
        while (!pick.isEmpty()) {
            var p = pick.poll();
            cur = Math.max(cur, p[1]) + time[p[0]][2];
        }
        return cur;
    }
}
