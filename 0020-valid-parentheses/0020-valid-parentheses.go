type Stack[]rune

func (s *Stack) Push(el rune) {
	(*s) = append(*s, el)
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}


func (s *Stack) Pop() (rune, bool) {
    if s.IsEmpty() {
        return 0, false
    }
	index := len(*s) - 1
	value := (*s)[index]
	(*s) = (*s)[:index]
	return value, true
}


func (s *Stack) Peek() (rune, bool) {
    if s.IsEmpty() {
        return 0, false
    }
	return (*s)[len(*s)-1], true
}

func isValid(s string) bool {
    if len(s) == 1 {
        return false
    }
    var runeStack Stack
    lookup := [3][2]rune{
        { rune('('), rune(')') },
        { rune('['), rune(']') },
        { rune('{'), rune('}') },
    }
    for _, char := range s {
        switch char {
            case lookup[0][1]:
                popped, ok := runeStack.Pop()
                if ok {
                    if popped != lookup[0][0] {
                        return false
                    }  
                } else {
                    return false
                }
            case lookup[1][1]:
                popped, ok := runeStack.Pop()
                if ok {
                    if popped != lookup[1][0] {
                        return false
                    }  
                } else {
                    return false
                }
            case lookup[2][1]:
                popped, ok := runeStack.Pop()
                if ok {
                    if popped != lookup[2][0] {
                        return false
                    }  
                } else {
                    return false
                }
            default:
                runeStack.Push(char);
        }
    }
    return runeStack.IsEmpty()
}