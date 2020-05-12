package main

const (
	opLdc           = 0x12
	opGetStatic     = 0xb2
	opInvokevirtual = 0xb6
	opReturn        = 0xb1
)

type instruction interface {
	fetchOperand(r *bytesReader)
	exec(fr *frame)
}

func decode(op uint8) instruction {
	switch op {
	case opLdc:
		return &Ldc{}

	case opGetStatic:
		return &Getstatic{}

	case opInvokevirtual:
		return &Invokevirtual{}

	case opReturn:
		return &Return{}
	}

	return nil
}

type Ldc struct {
	operand uint8
}

func (inst *Ldc) fetchOperand(r *bytesReader) {
	inst.operand = r.readUnit8()
}

func (inst *Ldc) exec(fr *frame) {
	// need to consider Class and Method
	index := getString(fr.constantPool, inst.operand).stringIndex
	value := getUTF8(fr.constantPool, index)
	fr.operandStack.push(value)
}

type Getstatic struct {
	operand uint16
}

func (inst *Getstatic) fetchOperand(r *bytesReader) {
	inst.operand = r.readUnit16()
}

func (inst *Getstatic) exec(fr *frame) {
	ref := getFieldRef(fr.constantPool, inst.operand)
	classRef := getClass(fr.constantPool, ref.classIndex)
	className := getUTF8(fr.constantPool, classRef.nameIndex)
	fieldName := getUTF8(fr.constantPool, getNameAndType(fr.constantPool, ref.nameAndTypeIndex).nameIndex)
	fr.operandStack.push(callerInfo{className, fieldName})
}

type Invokevirtual struct {
	operand uint16
}

func (inst *Invokevirtual) fetchOperand(r *bytesReader) {
	inst.operand = r.readUnit16()
}

func (inst *Invokevirtual) exec(fr *frame) {
	methodInfo := getMethodRef(fr.constantPool, inst.operand)
	nameAndType := getNameAndType(fr.constantPool, methodInfo.nameAndTypeIndex)

	methodClass := getClass(fr.constantPool, methodInfo.classIndex)
	methodClassName := getUTF8(fr.constantPool, methodClass.nameIndex)
	methodName := getUTF8(fr.constantPool, nameAndType.nameIndex)

	arguments := fr.operandStack.pop()

	caller := fr.operandStack.pop().(callerInfo)
	callee := calleeInfo{methodClassName, methodName}
	callInstanceMethod(caller, callee, arguments)
}

type Return struct {
}

func (inst *Return) fetchOperand(r *bytesReader) {
}

func (inst *Return) exec(fr *frame) {
}
