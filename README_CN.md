# 基于条件表达式的简单状态机 Go 实现
[English](README.md) | 中文

## 项目简介
这是一个用 Go 语言实现的简单状态机，支持基于条件表达式的状态转换。该状态机可以根据不同的事件和条件，动态地在多个状态之间进行切换。它适用于需要根据特定条件进行状态管理的场景，例如工作流引擎、游戏状态管理、用户权限控制等。
## 特性
* 条件驱动的状态转换：支持通过条件表达式来控制状态的转换。
* 灵活的状态定义：可以定义任意数量的状态和事件，支持多种状态转换逻辑。
* 易于扩展：代码结构清晰，易于添加新功能和状态。

## 示例代码
以下是一个简单的状态机实现示例：
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

## 贡献
欢迎贡献代码、报告问题或提出建议！请提交 Pull Request 或创建 Issue。