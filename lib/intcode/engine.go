package intcode

import (
	"fmt"
	"sync"
)

// Engine encapsulates an intcode consumer
type Engine struct {
	Ip       int
	Code     []int
	Inputs   *SignalQueue
	Outputs  *SignalQueue
	Modes    []int
	InputVal int
	OpMap    map[int]func(*Engine) int
}

// NewEngine makes a new Engine with a copy of an intcode
func NewEngine(code []int) Engine {
	codeCopy := make([]int, len(code))
	copy(codeCopy, code)

	q := newSignalQueue()
	opMap := map[int]func(*Engine) int{
		1: add,
		2: mult,
		3: input,
		4: output,
		5: jumpIfTrue,
		6: jumpIfFalse,
		7: lessThan,
		8: equals,
	}
	e := Engine{Code: codeCopy, Inputs: &q, OpMap: opMap}
	return e
}

// ConnectOutput sets an output queue
func (e *Engine) ConnectOutput(outputs *SignalQueue) {
	e.Outputs = outputs
}

// EvaluateIntcode evaluates and runs a given intcode
func (e *Engine) EvaluateIntcode(wg *sync.WaitGroup) {
	defer wg.Done()
	for e.Ip < len(e.Code) {
		var ipIncr int
		var op int

		op, e.Modes = opcodeParse(e.Code[e.Ip])
		if op != 99 {
			opFunc := e.OpMap[op]
			if opFunc != nil {
				ipIncr = opFunc(e)
			} else {
				fmt.Println("Something went wrong, invalid opcode: ", e.Ip, op)
			}
		} else {
			break
		}
		e.Ip += ipIncr
	}
}