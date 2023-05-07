package pymarshal

import (
	"bytes"
	"crypto/rc4"
	"encoding/binary"
	"errors"
	"fmt"
)

const (
	TYPE_NETEASE_CODE       byte = 0x61
	TYPE_NETEASE_CODE_NEW   byte = 0x6f
	TYPE_NETEASE_STRING_1   byte = 0x08
	TYPE_NETEASE_INTERNED_1 byte = 0x0f
	TYPE_NETEASE_UNICODE_1  byte = 0x0e
	TYPE_NETEASE_STRING_2   byte = 0x17
	TYPE_NETEASE_INTERNED_2 byte = 0x1a
	TYPE_NETEASE_UNICODE_2  byte = 0x1d
)

var rc4Cipher *rc4.Cipher

func initCipher() {
	var err error

	rc4Cipher, err = rc4.NewCipher([]byte{0xA7, 0xD, 0x37, 0x7A})
	if err != nil {
		panic(err)
	}
}

func repairNetEaseString(reader *bytes.Reader, writer *bytes.Buffer, write bool) error {
	var n int32
	err := binary.Read(reader, binary.LittleEndian, &n)
	if err != nil {
		return err
	}

	err = binary.Write(writer, binary.LittleEndian, n)
	if err != nil {
		return err
	}

	data := make([]byte, n)

	_, err = reader.Read(data)
	if err != nil {
		return err
	}

	decrypt := make([]byte, n)
	initCipher()
	rc4Cipher.XORKeyStream(decrypt, data)

	_, err = writer.Write(decrypt)
	if err != nil {
		return err
	}

	if !write {
		strings = append(strings, decrypt)
	}

	return nil
}

func repairNetEaseStringNew(reader *bytes.Reader, writer *bytes.Buffer, write bool) error {
	var n int32
	err := binary.Read(reader, binary.LittleEndian, &n)
	if err != nil {
		return err
	}

	err = binary.Write(writer, binary.LittleEndian, n)
	if err != nil {
		return err
	}

	data := make([]byte, n)

	_, err = reader.Read(data)
	if err != nil {
		return err
	}

	for i := 0; i < len(data); i++ {
		data[i] ^= 0x8d
	}

	_, err = writer.Write(data)
	if err != nil {
		return err
	}

	if !write {
		strings = append(strings, data)
	}

	return nil
}

func repairNetEaseCode1(reader *bytes.Reader, writer *bytes.Buffer) error {
	locals := make([]byte, 4)
	_, err := reader.Read(locals)
	if err != nil {
		return err
	}

	flags := make([]byte, 4)
	_, err = reader.Read(flags)
	if err != nil {
		return err
	}

	var consts bytes.Buffer
	err = RepairMarshal(reader, &consts)
	if err != nil {
		return err
	}

	stacksize := make([]byte, 4)
	_, err = reader.Read(stacksize)
	if err != nil {
		return err
	}

	var varnames bytes.Buffer
	err = RepairMarshal(reader, &varnames)
	if err != nil {
		return err
	}

	argcount := make([]byte, 4)
	_, err = reader.Read(argcount)
	if err != nil {
		return err
	}

	var freevars bytes.Buffer
	err = RepairMarshal(reader, &freevars)
	if err != nil {
		return err
	}

	var names bytes.Buffer
	err = RepairMarshal(reader, &names)
	if err != nil {
		return err
	}

	var cellvars bytes.Buffer
	err = RepairMarshal(reader, &cellvars)
	if err != nil {
		return err
	}

	var name bytes.Buffer
	err = RepairMarshal(reader, &name)
	if err != nil {
		return err
	}

	var code bytes.Buffer
	err = RepairMarshal(reader, &code)
	if err != nil {
		return err
	}

	firstlineno := make([]byte, 4)
	_, err = reader.Read(firstlineno)
	if err != nil {
		return err
	}

	var lnotab bytes.Buffer
	err = RepairMarshal(reader, &lnotab)
	if err != nil {
		return err
	}

	var opcodeType uint32
	err = binary.Read(reader, binary.LittleEndian, &opcodeType)
	if err != nil {
		return err
	}
	fmt.Printf("opcode: %x %d\n", opcodeType, opcodeType)
	var ncode bytes.Buffer
	err = repairOpcode(bytes.NewReader(code.Bytes()), &ncode, opcodeType)
	if err != nil {
		return err
	}

	var filename bytes.Buffer
	err = RepairMarshal(reader, &filename)
	if err != nil {
		return err
	}

	_, err = writer.Write(argcount)
	if err != nil {
		return err
	}
	_, err = writer.Write(locals)
	if err != nil {
		return err
	}
	_, err = writer.Write(stacksize)
	if err != nil {
		return err
	}
	_, err = writer.Write(flags)
	if err != nil {
		return err
	}
	_, err = writer.Write(ncode.Bytes())
	if err != nil {
		return err
	}
	_, err = writer.Write(consts.Bytes())
	if err != nil {
		return err
	}
	_, err = writer.Write(names.Bytes())
	if err != nil {
		return err
	}
	_, err = writer.Write(varnames.Bytes())
	if err != nil {
		return err
	}
	_, err = writer.Write(freevars.Bytes())
	if err != nil {
		return err
	}
	_, err = writer.Write(cellvars.Bytes())
	if err != nil {
		return err
	}
	_, err = writer.Write(filename.Bytes())
	if err != nil {
		return err
	}
	_, err = writer.Write(name.Bytes())
	if err != nil {
		return err
	}
	_, err = writer.Write(firstlineno)
	if err != nil {
		return err
	}
	_, err = writer.Write(lnotab.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func repairNetEaseCode(reader *bytes.Reader, writer *bytes.Buffer) error {
	var lnotab bytes.Buffer
	err := RepairMarshal(reader, &lnotab)
	if err != nil {
		return err
	}

	var varnames bytes.Buffer
	err = RepairMarshal(reader, &varnames)
	if err != nil {
		return err
	}

	flags := make([]byte, 4)
	_, err = reader.Read(flags)
	if err != nil {
		return err
	}

	var freevars bytes.Buffer
	err = RepairMarshal(reader, &freevars)
	if err != nil {
		return err
	}

	var cellvars bytes.Buffer
	err = RepairMarshal(reader, &cellvars)
	if err != nil {
		return err
	}

	var filename bytes.Buffer
	err = RepairMarshal(reader, &filename)
	if err != nil {
		return err
	}

	stacksize := make([]byte, 4)
	_, err = reader.Read(stacksize)
	if err != nil {
		return err
	}

	firstlineno := make([]byte, 4)
	_, err = reader.Read(firstlineno)
	if err != nil {
		return err
	}

	var consts bytes.Buffer
	if depth == 1 {
		// clean up constant pool
		if type1, _ := reader.ReadByte(); type1 == TYPE_TUPLE {
			var n int32
			err := binary.Read(reader, binary.LittleEndian, &n)
			if err != nil {
				return err
			}

			data := make([]bytes.Buffer, n)
			codeIndex := int32(-1)

			for i := int32(0); i < n; i++ {
				err = RepairMarshal(reader, &data[i])
				if err != nil {
					return err
				}

				if data[i].Bytes()[0] == TYPE_CODE {
					codeIndex = i
				}
			}

			consts.WriteByte(TYPE_TUPLE)
			err = binary.Write(&consts, binary.LittleEndian, n)
			if err != nil {
				return err
			}

			for i := 0; i < len(data); i++ {
				if i == int(codeIndex) {
					consts.Write([]byte{TYPE_NONE})
				} else {
					consts.Write(data[i].Bytes())
				}
			}
		} else {
			panic(fmt.Errorf("invalid constant pool type: %x", type1))
		}
	} else {
		err = RepairMarshal(reader, &consts)
		if err != nil {
			return err
		}
	}

	argcount := make([]byte, 4)
	_, err = reader.Read(argcount)
	if err != nil {
		return err
	}

	var code bytes.Buffer
	err = RepairMarshal(reader, &code)
	if err != nil {
		return err
	}

	locals := make([]byte, 4)
	_, err = reader.Read(locals)
	if err != nil {
		return err
	}

	var name bytes.Buffer
	err = RepairMarshal(reader, &name)
	if err != nil {
		return err
	}

	var names bytes.Buffer
	err = RepairMarshal(reader, &names)
	if err != nil {
		return err
	}

	var opcodeType uint32
	err = binary.Read(reader, binary.LittleEndian, &opcodeType)
	if err != nil {
		return err
	}
	fmt.Printf("操作码类型: %x\n", opcodeType)
	if opcodeMap, ok := neteaseOpcodes[opcodeType]; ok {
		data := code.Bytes()
		index := 5
		for index < len(data) {
			q := opcodeMap[data[index]]
			if q != "" {
				p, ok := opcodes[q]
				if ok {
					data[index] = p
					if p >= 90 {
						index += 2
					}
				} else {
					println("未知的index:", data[index], q)
				}
			} else {
				println("未知的index:", data[index])
			}
			index++
		}
		if depth == 1 && len(data) > 10 {
			data = append(data[:len(data)-3*3-1], data[len(data)-1])
			var buf bytes.Buffer
			var n int32
			err := binary.Read(bytes.NewReader(data[1:5]), binary.LittleEndian, &n)
			if err != nil {
				return err
			}

			err = binary.Write(&buf, binary.LittleEndian, n-3*3)
			if err != nil {
				return err
			}

			num := buf.Bytes()
			data[1] = num[0]
			data[2] = num[1]
			data[3] = num[2]
			data[4] = num[3]
		}
		code = *bytes.NewBuffer(data)
	} else {
		fmt.Printf("No opcode mapping found %x\n", opcodeType)
	}

	_, err = writer.Write(argcount)
	if err != nil {
		return err
	}
	_, err = writer.Write(locals)
	if err != nil {
		return err
	}
	_, err = writer.Write(stacksize)
	if err != nil {
		return err
	}
	_, err = writer.Write(flags)
	if err != nil {
		return err
	}
	_, err = writer.Write(code.Bytes())
	if err != nil {
		return err
	}
	_, err = writer.Write(consts.Bytes())
	if err != nil {
		return err
	}
	_, err = writer.Write(names.Bytes())
	if err != nil {
		return err
	}
	_, err = writer.Write(varnames.Bytes())
	if err != nil {
		return err
	}
	_, err = writer.Write(freevars.Bytes())
	if err != nil {
		return err
	}
	_, err = writer.Write(cellvars.Bytes())
	if err != nil {
		return err
	}
	_, err = writer.Write(filename.Bytes())
	if err != nil {
		return err
	}
	_, err = writer.Write(name.Bytes())
	if err != nil {
		return err
	}
	_, err = writer.Write(firstlineno)
	if err != nil {
		return err
	}
	_, err = writer.Write(lnotab.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func repairOpcode(reader *bytes.Reader, writer *bytes.Buffer, id uint32) error {
	err := copyBytes(reader, writer, 1)
	if err != nil {
		return err
	}

	var n int32
	err = binary.Read(reader, binary.LittleEndian, &n)
	if err != nil {
		return err
	}

	err = binary.Write(writer, binary.LittleEndian, n)
	if err != nil {
		return err
	}

	data := make([]byte, n)

	_, err = reader.Read(data)
	if err != nil {
		return err
	}

	if opcodeMap, ok := neteaseOpcodes[id]; ok {
		index := 0
		for index < len(data) {
			q := opcodeMap[data[index]]
			if q != "" {
				p, ok := opcodes[q]
				if ok {
					data[index] = p
					if p >= 90 {
						index += 2
					}
				} else {
					println("未知的位置:", data[index], q)
				}
			} else {
				println("未知的位置:", data[index])
			}
			index++
		}

		_, err = writer.Write(data)
		return err
	} else {
		return errors.New("no opcode map found")
	}
}
