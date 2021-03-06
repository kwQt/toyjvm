package main

type classFile struct {
	magic             uint32
	minorVersion      uint16
	majorVersion      uint16
	constantPoolCount uint16
	constantPool      []constantInfo
	accessFlags       uint16
	thisClass         uint16
	superClass        uint16
	interfacesCount   uint16
	interfaces        []uint16
	fieldsCount       uint16
	fields            []fieldsInfo
	methodsCount      uint16
	methods           []methodsInfo
	attributesCount   uint16
	attributes        []attributeInfo
}

type constantInfo struct {
	tag  uint8
	info interface{}
}

const (
	CONSTANT_Utf8               = 1
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_Class              = 7
	CONSTANT_String             = 8
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_NameAndType        = 12
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_Dynamic            = 17
	CONSTANT_InvokeDynamic      = 18
	CONSTANT_Module             = 19
	CONSTANT_Package            = 20
)

type CONSTANT_Utf8_info struct {
	length uint16
	bytes  []byte
}

type CONSTANT_Class_info struct {
	nameIndex uint16
}

type CONSTANT_String_info struct {
	stringIndex uint16
}

type CONSTANT_Fieldref_info struct {
	classIndex       uint16
	nameAndTypeIndex uint16
}

type CONSTANT_Methodref_info struct {
	classIndex       uint16
	nameAndTypeIndex uint16
}

type CONSTANT_NameAndType_info struct {
	nameIndex       uint16
	descriptorIndex uint16
}

type fieldsInfo struct {
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributesCount uint16
	attributes      []attributeInfo
}

type methodsInfo struct {
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributesCount uint16
	attributes      []attributeInfo
}

type attributeInfo struct {
	attributeNameIndex uint16
	attributeLength    uint32
	info               interface{}
}

type codeAttribute struct {
	maxStack             uint16
	maxLocals            uint16
	codeLength           uint32
	code                 []uint8
	exceptionTableLength uint16
	exceptionTable       []exceptionTable
	attributesCount      uint16
	attributes           []attributeInfo
}

type exceptionTable struct {
	startPC   uint16
	endPC     uint16
	headerPC  uint16
	catchType uint16
}

type lineNumberTableAttribute struct {
	lineNumberTableLength uint16
	lineNumberTable       []lineNumberTable
}

type lineNumberTable struct {
	startPC    uint16
	lineNumber uint16
}

type sourceFileAttribute struct {
	sourceFileIndex uint16
}

const (
	Code            = "Code"
	LineNumberTable = "LineNumberTable"
	SourceFile      = "SourceFile"
)
