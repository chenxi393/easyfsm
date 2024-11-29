package easyfsm

type StateType int8

const (
	StartState  StateType = 1 // 开始节点
	EndState    StateType = 2 // 结束节点
	CancelState StateType = 3 // 取消节点
)

type State struct {
	StateId   int
	StateName string
	Type      StateType // 使用 StateType 类型
}

type StateOption func(*State)

func WithType(stateType StateType) StateOption {
	return func(s *State) {
		s.Type = stateType
	}
}

func NewState(stateId int, stateName string, opts ...StateOption) State {
	s := State{
		StateId:   stateId,
		StateName: stateName,
	}
	for _, opt := range opts {
		opt(&s)
	}
	return s
}
