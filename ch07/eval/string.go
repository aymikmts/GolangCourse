package eval

import (
	"fmt"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%g", l)
}

func (u unary) String() string {
	return fmt.Sprintf("(%c%s)", u.op, u.x)
}

func (b binary) String() string {
	return fmt.Sprintf("(%s %c %s)", b.x, b.op, b.y)
}

func (c call) String() string {
	var ret string
	ret = fmt.Sprintf("%s(", c.fn)
	for i, arg := range c.args {
		if i > 0 {
			ret += ", "
		}
		ret += fmt.Sprintf("%s", arg)
	}
	ret += ")"
	return ret
}
