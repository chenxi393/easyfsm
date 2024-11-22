package easyfsm

import (
	"testing"
)

func testInit() {
	Init(map[string]BizDesc{
		"test": {
			BizName: "test",
			Transitions: []Transition{
				{
					Event: "time_event",
					Src: State{
						StateId:   1,
						StateName: "上班中",
					},
					Dst: []State{
						{
							StateId:   1,
							StateName: "上班中",
							CondExpr:  "time == 10",
						},
						{
							StateId:   2,
							StateName: "午休中",
							CondExpr:  "time == 12",
						},
						{
							StateId:   3,
							StateName: "下班状态",
							CondExpr:  "time == 18",
						},
					},
				},
				{
					Event: "time_event",
					Src: State{
						StateId:   2,
						StateName: "午休中",
					},
					Dst: []State{
						{
							StateId:   1,
							StateName: "上班中",
							CondExpr:  "time == 14",
						},
					},
				},
			},
		},
	})
}

func TestFsm(t *testing.T) {
	testInit()
	fms := NewFsm("test", 1)
	params := make(map[string]interface{})
	params["time"] = 10
	err := fms.Tran("time_event", params)
	if err != nil {
		t.Error(err)
	}
	if fms.curState != 1 {
		t.Error("curState is not 1")
	}
	params["time"] = 12
	err = fms.Tran("time_event", params)
	if err != nil {
		t.Error(err)
	}
	if fms.curState != 2 {
		t.Error("curState is not 2")
	}
	params["time"] = 14
	err = fms.Tran("time_event", params)
	if err != nil {
		t.Error(err)
	}
	if fms.curState != 1 {
		t.Error("curState is not 1")
	}
	params["time"] = 18
	err = fms.Tran("time_event", params)
	if err != nil {
		t.Error(err)
	}
	if fms.curState != 3 {
		t.Error("curState is not 3")
	}
}

func TestCanNot(t *testing.T) {
	testInit()
	fms := NewFsm("test", 2)
	params := make(map[string]interface{})
	params["time"] = 10
	err := fms.Tran("time_event", params)
	if err != nil {
		t.Error(err)
	}
	if fms.curState != 1 {
		t.Error("curState is not 1")
	}
}
