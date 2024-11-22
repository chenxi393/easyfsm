package easyfsm

type Transition struct {
	Event string
	Src   State
	Dst   []State
}

type State struct {
	StateId   int
	StateName string
	// Type      int8   // 状态类型 1 开始 2 结束
	CondExpr string // 状态转移条件
}

type BizDesc struct {
	BizName     string
	Transitions []Transition
}
