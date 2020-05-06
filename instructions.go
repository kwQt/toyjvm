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
}

func (inst *Ldc) fetchOperand(r *bytesReader) {

}

func (inst *Ldc) exec(fr *frame) {

}

type Getstatic struct {
}

func (inst *Getstatic) fetchOperand(r *bytesReader) {

}

func (inst *Getstatic) exec(fr *frame) {

}

type Invokevirtual struct {
}

func (inst *Invokevirtual) fetchOperand(r *bytesReader) {

}

func (inst *Invokevirtual) exec(fr *frame) {

}

type Return struct {
}

func (inst *Return) fetchOperand(r *bytesReader) {

}

func (inst *Return) exec(fr *frame) {

}
