package easyfsm

import (
	"testing"
)

func testInit() {
	InitOrderSt := NewState(0, "InitOrderSt", WithType(StartState))
	PendingPaymentSt := NewState(10, "PendingPaymentSt")
	OrderProcessingSt := NewState(20, "OrderProcessingSt")
	ShippedSt := NewState(30, "ShippedSt")
	InTransitSt := NewState(40, "InTransitSt")
	DeliveredSt := NewState(50, "DeliveredSt", WithType(EndState))
	RefundSt := NewState(60, "RefundSt", WithType(EndState))
	RetrunProcessingSt := NewState(70, "ReturnProcessingSt")
	RefundCompletedSt := NewState(102, "RefundCompletedSt", WithType(EndState))
	CanceledSt := NewState(-1, "CanceledSt", WithType(CancelState))
	Init(map[string]BizDesc{
		"order_test": {
			Transitions: []Transition{
				{
					Event: "OrderCreateEt",
					Src: []State{
						InitOrderSt,
					},
					Dst: []DstState{
						{
							State: PendingPaymentSt,
						},
					},
				},
				{
					Event: "PaymentSuccessEt",
					Src: []State{
						PendingPaymentSt,
					},
					Dst: []DstState{
						{
							State: OrderProcessingSt,
						},
					},
				},
				{
					Event: "OrderPackedEt",
					Src: []State{
						OrderProcessingSt,
					},
					Dst: []DstState{
						{
							State: ShippedSt,
						},
					},
				},
				{
					Event: "OrderShippedEt",
					Src: []State{
						ShippedSt,
					},
					Dst: []DstState{
						{
							State: InTransitSt,
						},
					},
				},
				{
					Event: "OrderDeliveredEt",
					Src: []State{
						InTransitSt,
					},
					Dst: []DstState{
						{
							State:    DeliveredSt,
							CondExpr: "customer_received == 1",
						},
						{
							State:    RefundSt,
							CondExpr: "customer_received == 0",
						},
					},
				},
				{
					Event: "OrderReturnEt",
					Src: []State{
						DeliveredSt,
					},
					Dst: []DstState{
						{
							State: RetrunProcessingSt,
						},
					},
				},
				{
					Event: "ReturnCompleteEt",
					Src: []State{
						RetrunProcessingSt,
					},
					Dst: []DstState{
						{
							State: RefundCompletedSt,
						},
					},
				},

				{
					Event: "OrderCancelEt",
					Src: []State{
						PendingPaymentSt,
						OrderProcessingSt,
					},
					Dst: []DstState{
						{
							State: CanceledSt,
						},
					},
				},
			},
		},
	})
}

func TestFsm(t *testing.T) {
	testInit()
	fsm, err := NewFsm("order_test", 0)
	if err != nil {
		t.Error(err)
	}
	err = fsm.Tran("OrderCreateEt", nil)
	if err != nil {
		t.Error(err)
	}
	if fsm.curState != 10 {
		t.Error("curState is not 10")
	}
	err = fsm.Tran("PaymentSuccessEt", nil)
	if err != nil {
		t.Error(err)
	}
	if fsm.curState != 20 {
		t.Error("curState is not 20")
	}
	err = fsm.Tran("OrderPackedEt", nil)
	if err != nil {
		t.Error(err)
	}
	if fsm.curState != 30 {
		t.Error("curState is not 30")
	}
	err = fsm.Tran("OrderShippedEt", nil)
	if err != nil {
		t.Error(err)
	}
	if fsm.curState != 40 {
		t.Error("curState is not 40")
	}
	para := make(map[string]interface{})
	para["customer_received"] = 1
	err = fsm.Tran("OrderDeliveredEt", para)
	if err != nil {
		t.Error(err)
	}
	if fsm.curState != 50 {
		t.Error("curState is not 50")
	}
}

func TestConditionTran(t *testing.T) {
	testInit()
	fsm, err := NewFsm("order_test", 40)
	if err != nil {
		t.Error(err)
	}
	para := make(map[string]interface{})
	para["customer_received"] = 0
	err = fsm.Tran("OrderDeliveredEt", para)
	if err != nil {
		t.Error(err)
	}
	if fsm.curState != 60 {
		t.Error("curState is not 60")
	}
}

func TestCanTran(t *testing.T) {
	testInit()
	fsm, err := NewFsm("order_test", 40)
	if err != nil {
		t.Error(err)
	}
	if fsm.CanTran("ReturnCompleteEt") {
		t.Error("can tran")
	}
}
