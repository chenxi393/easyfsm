# easyfsm
English | [简体中文](README_CN.md) 

## Project Overview
This is a simple state machine implemented in Go that supports state transitions based on conditional expressions. The state machine can dynamically switch between multiple states based on different events and conditions. It is suitable for scenarios that require state management based on specific conditions, such as workflow engines, game state management, user permission control, and more.

## Features
* **Condition-Driven State Transitions**: Supports controlling state transitions through conditional expressions.
* **Flexible State Definitions**: Allows the definition of an arbitrary number of states and events, supporting various state transition logic.
* **Easy to Extend**: The code structure is clear, making it easy to add new features and states.

## Example Code
Here is a simple example of the state machine implementation:
```go
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
```

## Contribution
Contributions are welcome! Please submit pull requests or create issues for bugs or suggestions.