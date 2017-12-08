package registers

import "testing"

func TestExecute(t *testing.T) {

	cases := []struct {
		in   string
		want int
	}{
		{`b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`, 1},
	}
	for _, c := range cases {
		got := Execute(c.in)
		if got != c.want {
			t.Errorf("Expected Execute(%v) to return %d got %d", c.in, c.want, got)
		}
	}
}

func TestParseInstructionLine(t *testing.T) {
	cases := []struct {
		in   string
		want Instruction
	}{
		{"gif dec -533 if q <= -7", Instruction{"dec", "gif", -533, "<=", "q", -7}},
		{"q dec 2 if qt < 77", Instruction{"dec", "q", 2, "<", "qt", 77}},
		{"a inc 0 if qt > 1234567", Instruction{"inc", "a", 0, ">", "qt", 1234567}},
		{"asdfg inc -0 if qtert >= 2", Instruction{"inc", "asdfg", 0, ">=", "qtert", 2}},
		{"b inc 1234567 if qt == -0", Instruction{"inc", "b", 1234567, "==", "qt", 0}},
		{"ht inc -2 if qt != 0", Instruction{"inc", "ht", -2, "!=", "qt", 0}},
	}
	for _, c := range cases {
		got := ParseInstructionLine(c.in)
		if got.Operator != c.want.Operator || got.RegisterName != c.want.RegisterName || got.Value != c.want.Value || got.ConditionOperator != c.want.ConditionOperator || got.ConditionRegisterName != c.want.ConditionRegisterName || got.ConditionValue != c.want.ConditionValue {
			t.Errorf("Expected ParseInstructionLine(\"%s\") to return %s, %s, %d, %s, %s, %d, got %s, %s, %d, %s, %s, %d", c.in, c.want.Operator, c.want.RegisterName, c.want.Value, c.want.ConditionOperator, c.want.ConditionRegisterName, c.want.ConditionValue, got.Operator, got.RegisterName, got.Value, got.ConditionOperator, got.ConditionRegisterName, got.ConditionValue)
		}
	}
}
