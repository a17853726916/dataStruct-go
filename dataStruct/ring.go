package dataStruct

type Ring struct {
	Next *Ring
	Prev *Ring
	Val  interface{}
}

// 初始化空循环链表,只有一个的时候都是指向自己
func (r *Ring) Init() *Ring {
	r.Next = r
	r.Prev = r
	return r
}

// 构建N个节点链表

func New(n int) *Ring {

	if n <= 0 {
		return nil
	}
	r := new(Ring)
	temp := r

	for i := 0; i < n; i++ {
		temp.Next = &Ring{Prev: temp}
		temp = temp.Next
	}
	temp.Next = r
	r.Prev = temp

	return r
}

// 获取下一个节点

func (r *Ring) Nextnode() *Ring {
	if r.Next == nil {
		return r.Init()
	}
	return r.Next
}

// 获取上一个节点
func (r *Ring) PrevNode() *Ring {
	if r.Next == nil {
		return r.Init()
	}
	return r.Prev
}

// 添加节点
func (r *Ring) Link(s *Ring) *Ring {
	n := r.Nextnode()

	if s != nil {
		p := s.PrevNode() //如果没有创建一个指向自身的节点
		r.Next = s
		s.Prev = r
		n.Prev = p
		p.Next = n

	}
	return n
}

// 获取第N个节点
func (r *Ring) Move(n int) *Ring {
	if r.Next == nil {
		return r.Init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.Prev
		}
	case n > 0:
		for ; n > 0; n-- {
			r = r.Next
		}
	}
	return r
}

// 删除
func (r *Ring) Unlink(n int) *Ring {
	if n < 0 {
		return nil
	}
	return r.Link(r.Move(n + 1))
}
