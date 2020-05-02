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
