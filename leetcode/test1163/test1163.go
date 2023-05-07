package test1163

/**
 * @Description https://leetcode.cn/problems/last-substring-in-lexicographical-order/
 * idea: 因为是求字典序最大的, 答案必为从i开始到原字符串结束的后缀 即s[i,n-1]
		 使用双指针, i,j分别记录答案字符串和待比较字符串， 若两个字符串前缀相等，最后选较长的即可
		 所以可以用k记录两个字符串当前比较到的位置
	     1.若s[i+k]=s[j+k], 说明前k个位置都相等, 继续看下一个位置即可 此时k+1
  		 2.若s[i+k]>s[j+k], 说明字符串s[i]>s[j], j+k后面无需再看，直接从j+k的后一个
		   位置重新与i进行比较即可, 此时 j=j+k+1, k=0
         3.若s[i+k]<s[j+k], 说明从i开始的字符串s[i,n-1], s[i+1,n-1]...必小于j开始的字符串s[j,n-1], s[j+1,n-1]..,
		   此时将i指针移动到i+k后重新开始比较即可，j指针同样需要变更, 即i=max(i+k+1, j), j=i+1
 **/

func lastSubstring(s string) string {
	var n = len(s)
	i, j := 0, 1
	for k := 0; i+k < n && j+k < n; {
		if s[i+k] == s[j+k] {
			k++
		} else if s[i+k] > s[j+k] {
			j = j + k + 1
			k = 0
		} else {
			i = max(i+k+1, j)
			j = i + 1
			k = 0
		}
	}
	return s[i:]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
