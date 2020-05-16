package main

type frameStack []frame

type frame struct {
	frameStack   frameStack
	prev         *frame
	operandStack operandStack
	localVars    []interface{}
	constantPool []constantInfo
	code         *codeAttribute
}

type operandStack []interface{}

func (fr *frame) execute() {
	code := fr.code.code
	reader := bytesReader{code, 0, nil}
	for {
		if reader.curIdx == len(reader.data) {
			break
		}
		op := reader.readUnit8()
		inst := decode(op)
		inst.fetchOperand(&reader)
		inst.exec(fr)
	}
}

func (stack *operandStack) pop() interface{} {
	top := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return top
}

func (stack *operandStack) push(operand interface{}) {
	*stack = append(*stack, operand)
}

func (stack *frameStack) push(fr *frame) {
	*stack = append(*stack, *fr)
}
