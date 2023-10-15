package test858

/**
 * @Description https://leetcode.cn/problems/mirror-reflection
 * idea: 反射光线对称，到达墙上边界时反射到墙左边离上边界的距离实际等同于镜子延长后在
		左边隔下边界的距离，所以可以假设镜子在原来的基础上【无限延长】
		1.光线到达角落，需满足光线走过的纵向距离能整除p, 而实际反射次数又为整数，
		 所以最早走到角落的位置为lcm(p,q)。即p、q的最小公倍数
		2.分情况讨论。若lcm/p为偶数次，实际是光走到上边界后折返，此时能到达的下边界顶点只能是0
					若为奇数次，出口肯定为上边界，此时顶点可能为2或1。根据光线与墙的接触次数
					即可判断。即lcm/q，若为偶数次，则出口为2，奇数则为1。
		光线有没有可能到达左下角? 没可能，题目保证了一定会遇到一个接收器，而左下角为光源
 **/

func mirrorReflection(p int, q int) int {
	var lcm = p * q / gcd(p, q)
	var a = lcm / p
	var b = lcm / q
	if a%2 == 0 {
		return 0
	}
	if b%2 == 0 {
		return 2
	}
	return 1
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
