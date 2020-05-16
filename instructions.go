package main

const (
	opIconst0       = 0x3
	opIconst1       = 0x4
	opIconst2       = 0x5
	opIconst3       = 0x6
	opIconst4       = 0x7
	opIconst5       = 0x8
	opBipush        = 0x10
	opLdc           = 0x12
	opIload         = 0x15
	opIload0        = 0x1a
	opIload1        = 0x1b
	opIload2        = 0x1c
	opIload3        = 0x1d
	opIstore        = 0x36
	opIstore0       = 0x3b
	opIstore1       = 0x3c
	opIstore2       = 0x3d
	opIstore3       = 0x3e
	opIadd          = 0x60
	opIsub          = 0x64
	opImul          = 0x68
	opIdiv          = 0x6c
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
	// TODO: Consider Class and Method
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
