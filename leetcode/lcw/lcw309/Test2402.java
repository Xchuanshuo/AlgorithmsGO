package leetcode.lcw.lcw309;

import java.util.*;

public class Test2402 {

    public int mostBooked(int n, int[][] meetings) {
        TreeSet<Integer> set = new TreeSet<>();
        for (int i = 0;i < n;i++) {
            set.add(i);
        }
        Queue<Pair> pq = new PriorityQueue<>(Comparator.comparingLong(o -> o.endTime));
        Arrays.sort(meetings, Comparator.comparingInt((int[] a) -> a[0]));
        Map<Integer, Integer> cntMap = new HashMap<>();
        long curTime = 0;
        for (int i = 0;i < meetings.length;i++) {
            int[] cur = meetings[i];
            curTime = Math.max(curTime, cur[0]);
            if (set.size() == 0 || (!pq.isEmpty() && pq.peek().endTime <= curTime)) {
                curTime = Math.max(curTime, pq.peek().endTime);
                while (!pq.isEmpty() && pq.peek().endTime <= curTime) {
                    Pair pair = pq.remove();
                    set.add(pair.id);
                    curTime = Math.max(curTime, pair.endTime);
                }
            }
            int id = set.pollFirst();
            cntMap.put(id, cntMap.getOrDefault(id, 0) + 1);
            pq.offer(new Pair(id, curTime + cur[1] - cur[0]));
        }
        int resID = 0;
        for (int i = 0;i < n;i++) {
            if (cntMap.getOrDefault(i, 0) > cntMap.getOrDefault(resID, 0)) {
                resID = i;
            }
        }
        return resID;
    }


    class Pair {
        int id;
        long endTime;
        Pair(int id, long endTime) {
            this.id = id;
            this.endTime = endTime;
        }
    }

    public static void main(String[] args) {
        int n = 4;
        int[][] arr = new int[][]{{18,19},{3,12},{17,19},{2,13}, {7,10}};
        Test2402 t = new Test2402();
        int res = t.mostBooked(n, arr);
        System.out.println(res);
    }
}
