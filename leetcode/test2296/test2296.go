package test2296

// TextEditor 双栈模拟 左移: s1栈顶入s2; 右移: s2栈顶入s1
type TextEditor struct {
	s1, s2 []byte
}

func Constructor() TextEditor {
	return TextEditor{}
}

func (this *TextEditor) AddText(text string) {
	this.s1 = append(this.s1, text...)
}

func (this *TextEditor) DeleteText(k int) int {
	var cnt = 0
	for len(this.s1) > 0 && k > 0 {
		this.s1 = this.s1[0 : len(this.s1)-1]
		k--
		cnt++
	}
	return cnt
}

func (this *TextEditor) text() string {
	return string(this.s1[max(len(this.s1)-10, 0):])
}

func (this *TextEditor) CursorLeft(k int) string {
	for len(this.s1) > 0 && k > 0 {
		this.s2 = append(this.s2, this.s1[len(this.s1)-1])
		this.s1 = this.s1[0 : len(this.s1)-1]
		k--
	}
	return this.text()
}

func (this *TextEditor) CursorRight(k int) string {
	for len(this.s2) > 0 && k > 0 {
		this.s1 = append(this.s1, this.s2[len(this.s2)-1])
		this.s2 = this.s2[0 : len(this.s2)-1]
		k--
	}
	return this.text()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}