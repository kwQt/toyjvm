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
}

func (r *bytesReader) readMagic() uint32 {
	magic := r.readUnit32()
	if magic != 0xCAFEBABE {
		panic("Loaded File is not Java Class File")
	}
	return magic
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
