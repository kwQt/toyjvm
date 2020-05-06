package main

type frameStack []frame

type frame struct {
	frameStack   *frameStack
	prev         *frame
	operandStack []operand
	localVars    []localVar
	constantPool []constantInfo
	code         *codeAttribute
}

type operand interface{}

type localVar interface{}

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

func (stack *frameStack) push(fr *frame) {
	*stack = append(*stack, *fr)
}
