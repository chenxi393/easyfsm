package easyfsm

type Transition struct {
	Event string
	Src   []State
	Dst   []DstState
}

type DstState struct {
	State
	CondExpr string // 状态转移条件
}

type BizDesc struct {
	Transitions []Transition
}
