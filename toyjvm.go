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

type bytesReader struct {
	data   []byte
	curIdx int
}

func (cf *classFile) parseClassFile(data []byte) {
	reader := bytesReader{data, 0}
	cf.magic = reader.readMagic()
	cf.minorVersion = reader.readUnit16()
	cf.majorVersion = reader.readUnit16()
	cf.constantPoolCount = reader.readUnit16()
	cf.constantPool = reader.readConstantPool(int(cf.constantPoolCount))
}

func (r *bytesReader) readMagic() uint32 {
	magic := r.readUnit32()
	if magic != 0xCAFEBABE {
		panic("Loaded File is not Java Class File")
	}
	return magic
}

func (r *bytesReader) readConstantPool(cpCount int) []constantInfo {
	constantPool := []constantInfo{}

	for i := 0; i < cpCount-1; i++ {
		tag := r.readUnit8()
		next := r.readConstantInfo(tag)
		constantPool = append(constantPool, next)
	}

	return constantPool
}

func (r *bytesReader) readConstantInfo(tag uint8) constantInfo {
	var info interface{}

	switch tag {
	case CONSTANT_Utf8:
	case CONSTANT_Integer:
	case CONSTANT_Float:
	case CONSTANT_Long:
	case CONSTANT_Double:
	case CONSTANT_Class:
	case CONSTANT_String:
	case CONSTANT_Fieldref:
	case CONSTANT_Methodref:
	case CONSTANT_InterfaceMethodref:
	case CONSTANT_NameAndType:
	case CONSTANT_MethodHandle:
	case CONSTANT_MethodType:
	case CONSTANT_Dynamic:
	case CONSTANT_InvokeDynamic:
	case CONSTANT_Module:
	case CONSTANT_Package:
	}

	return constantInfo{tag, info}
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
