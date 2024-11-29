package easyfsm

import (
	"errors"
	"fmt"

	"github.com/Knetic/govaluate"
)

var (
	// globalBizDesc  = make(map[string]BizDesc)
	globalBizTrans   = make(map[string]map[SrcEventKey]Transition)
	globalExpression = make(map[string]*govaluate.EvaluableExpression)
)

// 程序启动时 初始化状态机
func Init(bizDescMap map[string]BizDesc) error {
	// globalBizDesc = bizDescMap
	for bizName, bizDesc := range bizDescMap {
		transMap := make(map[SrcEventKey]Transition)
		for _, transition := range bizDesc.Transitions {
			for _, srcState := range transition.Src {
				if _, ok := transMap[SrcEventKey{
					SrcState: srcState.StateId,
					Event:    transition.Event,
				}]; ok {
					return fmt.Errorf("duplicate transition %s", transition.Event)
				}
				transMap[SrcEventKey{
					SrcState: srcState.StateId,
					Event:    transition.Event,
				}] = transition
			}
		}
		globalBizTrans[bizName] = transMap
	}
	for _, transMap := range globalBizTrans {
		for _, transition := range transMap {
			for _, desState := range transition.Dst {
				if desState.CondExpr == "" {
					continue
				}
				expression, err := govaluate.NewEvaluableExpression(desState.CondExpr)
				if err != nil {
					return errors.Join(err, fmt.Errorf("invalid expression %s", desState.CondExpr))
				}
				globalExpression[desState.CondExpr] = expression
			}
		}
	}
	// 是否需要校验状态机是否正确？
	return nil
}
