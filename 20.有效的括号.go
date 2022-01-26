/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */
package main

// @lc code=start
type Stack struct {
	stack  []rune
	length int
}

func (s *Stack) Push(r rune) {
	s.stack = append(s.stack, r)
	s.length++
}
func (s *Stack) Pop() rune {
	if s.length < 1 {
		return 0
	}
	r := s.stack[s.length-1]
	s.stack = s.stack[0 : s.length-1]
	s.length--
	return r
}
func isValid(s string) bool {
	stack := Stack{}
	for _, r := range s {
		switch r {
		case '(', '{', '[':
			stack.Push(r)
		case ')':
			top := stack.Pop()
			if top != '(' {
				return false
			}
		case '}':
			top := stack.Pop()
			if top != '{' {
				return false
			}
		case ']':
			top := stack.Pop()
			if top != '[' {
				return false
			}
		default:
			return false
		}
	}

	return stack.length == 0
}

// @lc code=end
