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
	case opIconst0:
		return &Iconst0{}

	case opIconst1:
		return &Iconst1{}

	case opIconst2:
		return &Iconst2{}

	case opIconst3:
		return &Iconst3{}

	case opIconst4:
		return &Iconst4{}

	case opIconst5:
		return &Iconst5{}

	case opBipush:
		return &Bipush{}

	case opLdc:
		return &Ldc{}

	case opIload1:
		return &Iload1{}

	case opIload2:
		return &Iload2{}

	case opIstore1:
		return &Istore1{}

	case opIstore2:
		return &Istore2{}

	case opIadd:
		return &Iadd{}

	case opIsub:
		return &Isub{}

	case opImul:
		return &Imul{}

	case opIdiv:
		return &Idiv{}

	case opGetStatic:
		return &Getstatic{}

	case opInvokevirtual:
		return &Invokevirtual{}

	case opReturn:
		return &Return{}
	}

	return nil
}

type Iconst0 struct {
}

func (inst *Iconst0) fetchOperand(r *bytesReader) {
}

func (inst *Iconst0) exec(fr *frame) {
	fr.operandStack.push(0)
}

type Iconst1 struct {
}

func (inst *Iconst1) fetchOperand(r *bytesReader) {
}

func (inst *Iconst1) exec(fr *frame) {
	fr.operandStack.push(1)
}

type Iconst2 struct {
}

func (inst *Iconst2) fetchOperand(r *bytesReader) {
}

func (inst *Iconst2) exec(fr *frame) {
	fr.operandStack.push(2)
}

type Iconst3 struct {
}

func (inst *Iconst3) fetchOperand(r *bytesReader) {
}

func (inst *Iconst3) exec(fr *frame) {
	fr.operandStack.push(3)
}

type Iconst4 struct {
}

func (inst *Iconst4) fetchOperand(r *bytesReader) {
}

func (inst *Iconst4) exec(fr *frame) {
	fr.operandStack.push(4)
}

type Iconst5 struct {
}

func (inst *Iconst5) fetchOperand(r *bytesReader) {
}

func (inst *Iconst5) exec(fr *frame) {
	fr.operandStack.push(5)
}

type Bipush struct {
	operand uint8
}

func (inst *Bipush) fetchOperand(r *bytesReader) {
	inst.operand = r.readUnit8()
}

func (inst *Bipush) exec(fr *frame) {
	fr.operandStack.push(int(inst.operand))
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

type Iload1 struct {
}

func (inst *Iload1) fetchOperand(r *bytesReader) {
}

func (inst *Iload1) exec(fr *frame) {
	fr.operandStack.push(fr.localVars[1])
}

type Iload2 struct {
}

func (inst *Iload2) fetchOperand(r *bytesReader) {
}

func (inst *Iload2) exec(fr *frame) {
	fr.operandStack.push(fr.localVars[2])
}

type Istore1 struct {
}

func (inst *Istore1) fetchOperand(r *bytesReader) {
}

func (inst *Istore1) exec(fr *frame) {
	value := fr.operandStack.pop().(int)
	fr.localVars[1] = value
}

type Istore2 struct {
}

func (inst *Istore2) fetchOperand(r *bytesReader) {
}

func (inst *Istore2) exec(fr *frame) {
	value := fr.operandStack.pop()
	fr.localVars[2] = value
}

type Iadd struct {
}

func (inst *Iadd) fetchOperand(r *bytesReader) {
}

func (inst *Iadd) exec(fr *frame) {
	val1 := fr.operandStack.pop().(int)
	val2 := fr.operandStack.pop().(int)
	fr.operandStack.push(val2 + val1)
}

type Isub struct {
}

func (inst *Isub) fetchOperand(r *bytesReader) {
}

func (inst *Isub) exec(fr *frame) {
	val1 := fr.operandStack.pop().(int)
	val2 := fr.operandStack.pop().(int)
	fr.operandStack.push(val2 - val1)
}

type Imul struct {
}

func (inst *Imul) fetchOperand(r *bytesReader) {
}

func (inst *Imul) exec(fr *frame) {
	val1 := fr.operandStack.pop().(int)
	val2 := fr.operandStack.pop().(int)
	fr.operandStack.push(val2 * val1)
}

type Idiv struct {
}

func (inst *Idiv) fetchOperand(r *bytesReader) {
}

func (inst *Idiv) exec(fr *frame) {
	val1 := fr.operandStack.pop().(int)
	val2 := fr.operandStack.pop().(int)
	fr.operandStack.push(val2 / val1)
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
