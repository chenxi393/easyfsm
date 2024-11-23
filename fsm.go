package easyfsm

import (
	"fmt"
)

func NewFsm(biz string, curState int) (*Fsm, error) {
	if _, ok := globalBizTrans[biz]; !ok {
		return nil, fmt.Errorf("biz %s not exist", biz)
	}
	fsm := &Fsm{
		bizName:  biz,
		curState: curState,
	}
	fsm.transMap = globalBizTrans[biz]
	return fsm, nil
}

// true 意味着可以转移到下一个状态
func (f *Fsm) CanTran(event string) bool {
	if _, ok := f.transMap[SrcEventKey{
		Event:    event,
		SrcState: f.curState,
	}]; !ok {
		return false
	}
	return true
}

// 状态转移
func (f *Fsm) Tran(event string, paras map[string]interface{}) error {
	if !f.CanTran(event) {
		return fmt.Errorf("%d can not tran by %s", f.curState, event)
	}
	transition := f.transMap[SrcEventKey{
		Event:    event,
		SrcState: f.curState,
	}]
	// 执行状态转移
	for _, desState := range transition.Dst {
		if desState.CondExpr == "" {
			f.curState = desState.StateId
			return nil
		}
		// 执行条件表达式
		result, err := globalExpression[desState.CondExpr].Evaluate(paras)
		if err != nil {
			return err
		}
		resultBool, ok := result.(bool)
		if !ok {
			return fmt.Errorf("invalid expression %s", desState.CondExpr)
		}
		if resultBool {
			f.curState = desState.StateId
			return nil
		}
	}
	return fmt.Errorf("can not tran by %s", event)
}

func (f *Fsm) CurState() int {
	return f.curState
}

// 状态机
type Fsm struct {
	bizName  string // 业务之间状态转移是隔离的 需要提前初始化
	curState int
	transMap map[SrcEventKey]Transition // 状态转移map
}
