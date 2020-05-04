package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no class file path")
		return
	}
	path := os.Args[1]

	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	cf := classFile{}
	cf.parseClassFile(data)
}

func (cf *classFile) parseClassFile(data []byte) {
	reader := bytesReader{data, 0, nil}
	cf.magic = reader.readMagic()
	cf.minorVersion = reader.readUnit16()
	cf.majorVersion = reader.readUnit16()
	cf.constantPoolCount = reader.readUnit16()
	cf.constantPool = reader.readConstantPool(int(cf.constantPoolCount))
	reader.cp = cf.constantPool
	cf.accessFlags = reader.readUnit16()
	cf.thisClass = reader.readUnit16()
	cf.superClass = reader.readUnit16()
	cf.interfacesCount = reader.readUnit16()
	cf.interfaces = reader.readInterfaces(int(cf.interfacesCount))
	cf.fieldsCount = reader.readUnit16()
	cf.fields = reader.readFields(int(cf.fieldsCount))
	cf.methodsCount = reader.readUnit16()
	cf.methods = reader.readMethods(int(cf.methodsCount))
	cf.attributesCount = reader.readUnit16()
	cf.attributes = reader.readAttributes(int(cf.attributesCount))
}

type bytesReader struct {
	data   []byte
	curIdx int
	cp     []constantInfo
}

func (r *bytesReader) readMagic() uint32 {
	magic := r.readUnit32()
	if magic != 0xCAFEBABE {
		panic("Loaded File is not Java Class File")
	}
	return magic
}

func (r *bytesReader) readConstantPool(count int) []constantInfo {
	cpTable := make([]constantInfo, count)
	for i := 1; i < count; i++ {
		tag := r.readUnit8()
		cpTable[i] = r.readConstantInfo(tag)
	}
	return cpTable
}

func (r *bytesReader) readConstantInfo(tag uint8) constantInfo {
	var info interface{}

	switch tag {
	case CONSTANT_Utf8:
		length := r.readUnit16()
		bytes := r.readBytes(int(length))
		info = CONSTANT_Utf8_info{length, bytes}
	case CONSTANT_Integer:
	case CONSTANT_Float:
	case CONSTANT_Long:
	case CONSTANT_Double:
	case CONSTANT_Class:
		nameIndex := r.readUnit16()
		info = CONSTANT_Class_info{nameIndex}
	case CONSTANT_String:
		stringIndex := r.readUnit16()
		info = CONSTANT_String_info{stringIndex}
	case CONSTANT_Fieldref:
		classIndex := r.readUnit16()
		nameAndTypeIndex := r.readUnit16()
		info = CONSTANT_Fieldref_info{classIndex, nameAndTypeIndex}
	case CONSTANT_Methodref:
		classIndex := r.readUnit16()
		nameAndTypeIndex := r.readUnit16()
		info = CONSTANT_Methodref_info{classIndex, nameAndTypeIndex}
	case CONSTANT_InterfaceMethodref:
	case CONSTANT_NameAndType:
		nameIndex := r.readUnit16()
		descriptorIndex := r.readUnit16()
		info = CONSTANT_NameAndType_info{nameIndex, descriptorIndex}
	case CONSTANT_MethodHandle:
	case CONSTANT_MethodType:
	case CONSTANT_Dynamic:
	case CONSTANT_InvokeDynamic:
	case CONSTANT_Module:
	case CONSTANT_Package:
	}

	return constantInfo{tag, info}
}

func (r *bytesReader) readInterfaces(count int) []uint16 {
	intefaces := make([]uint16, count)
	for i := 0; i < count; i++ {
		intefaces[i] = r.readUnit16()
	}
	return intefaces
}

func (r *bytesReader) readFields(count int) []fieldsInfo {
	fields := make([]fieldsInfo, count)
	for i := 0; i < count; i++ {
		accessFlags := r.readUnit16()
		nameIndex := r.readUnit16()
		descriptorIndex := r.readUnit16()
		attributesCount := r.readUnit16()
		attributes := r.readAttributes(int(attributesCount))
		fields[i] = fieldsInfo{accessFlags, nameIndex, descriptorIndex, attributesCount, attributes}
	}
	return fields
}

func (r *bytesReader) readMethods(count int) []methodsInfo {
	methods := make([]methodsInfo, count)
	for i := 0; i < count; i++ {
		accessFlags := r.readUnit16()
		nameIndex := r.readUnit16()
		descriptorIndex := r.readUnit16()
		attributesCount := r.readUnit16()
		attributes := r.readAttributes(int(attributesCount))
		methods[i] = methodsInfo{accessFlags, nameIndex, descriptorIndex, attributesCount, attributes}
	}
	return methods
}

func (r *bytesReader) readAttributes(count int) []attributeInfo {
	attributes := make([]attributeInfo, count)
	for i := 0; i < count; i++ {
		attributes[i] = r.readAttributeInfo()
	}
	return attributes
}

func (r *bytesReader) readAttributeInfo() attributeInfo {
	attributeNameIndex := r.readUnit16()
	attributeLength := r.readUnit32()
	var info interface{}
	name := getUTF8(r.cp, attributeNameIndex)
	switch name {
	case Code:
		info = r.readCodeAttribute()
	case LineNumberTable:
		info = r.readLinuNumberTableAttribute()
	case SourceFile:
		info = sourceFileAttribute{r.readUnit16()}
	}
	return attributeInfo{attributeNameIndex, attributeLength, info}
}

func (r *bytesReader) readCodeAttribute() codeAttribute {
	maxStack := r.readUnit16()
	maxLocal := r.readUnit16()
	codeLength := r.readUnit32()
	code := make([]uint8, codeLength)
	for i := 0; i < int(codeLength); i++ {
		code[i] = r.readUnit8()
	}
	exceptionTableLength := r.readUnit16()
	table := make([]exceptionTable, exceptionTableLength)
	for i := 0; i < int(exceptionTableLength); i++ {
		startPC := r.readUnit16()
		endPC := r.readUnit16()
		headerPC := r.readUnit16()
		catchType := r.readUnit16()
		table[i] = exceptionTable{startPC, endPC, headerPC, catchType}
	}

	attributesCount := r.readUnit16()
	attributes := r.readAttributes(int(attributesCount))

	return codeAttribute{maxStack, maxLocal, codeLength, code, exceptionTableLength, table, attributesCount, attributes}
}

func (r *bytesReader) readLinuNumberTableAttribute() lineNumberTableAttribute {
	lineNumberTableLength := r.readUnit16()
	table := make([]lineNumberTable, lineNumberTableLength)
	for i := 0; i < int(lineNumberTableLength); i++ {
		startPC := r.readUnit16()
		lineNumber := r.readUnit16()
		table[i] = lineNumberTable{startPC, lineNumber}
	}
	return lineNumberTableAttribute{lineNumberTableLength, table}
}

func getUTF8(cp []constantInfo, cpIndex uint16) string {
	if cpIndex == 0 {
		return ""
	}
	utf8Info, ok := cp[cpIndex].info.(CONSTANT_Utf8_info)
	if !ok {
		panic("can not get Constant_Utf8_info")
	}
	return string(utf8Info.bytes)
}

func (r *bytesReader) readBytes(num int) []byte {
	result := r.data[r.curIdx : r.curIdx+num]
	r.curIdx += num
	return result
}

func (r *bytesReader) readUnit8() uint8 {
	result := r.data[r.curIdx]
	r.curIdx++
	return result
}

func (r *bytesReader) readUnit16() uint16 {
	result := binary.BigEndian.Uint16(r.data[r.curIdx : r.curIdx+2])
	r.curIdx += 2
	return result
}

func (r *bytesReader) readUnit32() uint32 {
	result := binary.BigEndian.Uint32(r.data[r.curIdx : r.curIdx+4])
	r.curIdx += 4
	return result
}
