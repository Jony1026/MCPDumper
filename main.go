package main

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"os/exec"
	"serifu/mcpk"
	"serifu/pymarshal"
	"strings"
	"time"
)

var rotor *mcpk.RotorObj
var idRedirect = mcpk.StringID("redirect.mcs")

func main() {
	// opcodeMap := make(map[byte]string)
	// opcodeMap[17] = "INPLACE_FLOOR_DIVIDE"
	// opcodeMap[85] = "SLICE_0"
	// opcodeMap[28] = "STORE_SLICE_0"
	// opcodeMap[183] = "DELETE_ATTR"
	// opcodeMap[158] = "STORE_GLOBAL"
	// opcodeMap[224] = "DUP_TOPX"
	// opcodeMap[241] = "POP_JUMP_IF_FALSE"
	// opcodeMap[200] = "POP_JUMP_IF_TRUE"
	// opcodeMap[121] = "DELETE_FAST"
	// opcodeMap[2] = "BINARY_POWER"
	// opcodeMap[42] = "PRINT_ITEM"
	// opcodeMap[60] = "RETURN_VALUE"
	// opcodeMap[114] = "FOR_ITER"
	// opcodeMap[48] = "BINARY_TRUE_DIVIDE"
	// opcodeMap[30] = "STORE_SLICE_2"
	// opcodeMap[26] = "INPLACE_XOR"
	// opcodeMap[81] = "INPLACE_SUBTRACT"
	// opcodeMap[214] = "LOAD_CONST"
	// opcodeMap[96] = "SETUP_LOOP"
	// opcodeMap[105] = "MAKE_CLOSURE"
	// opcodeMap[58] = "PRINT_EXPR"
	// opcodeMap[212] = "LOAD_NAME"
	// opcodeMap[190] = "BUILD_TUPLE"
	// opcodeMap[223] = "IMPORT_FROM"
	// opcodeMap[178] = "UNPACK_SEQUENCE"
	// opcodeMap[231] = "LOAD_ATTR"
	// opcodeMap[211] = "LOAD_DEREF"
	// opcodeMap[45] = "ROT_TWO"
	// opcodeMap[4] = "BINARY_MULTIPLY"
	// opcodeMap[88] = "SLICE_3"
	// opcodeMap[57] = "NOP"
	// opcodeMap[51] = "BINARY_DIVIDE"
	// opcodeMap[246] = "LOAD_CLOSURE"
	// opcodeMap[10] = "BINARY_FLOOR_DIVIDE"
	// opcodeMap[18] = "INPLACE_POWER"
	// opcodeMap[59] = "GET_ITER"
	// opcodeMap[74] = "POP_BLOCK"
	// opcodeMap[126] = "CALL_FUNCTION"
	// opcodeMap[87] = "SLICE_2"
	// opcodeMap[63] = "BREAK_LOOP"
	// opcodeMap[3] = "LOAD_LOCALS"
	// opcodeMap[107] = "DELETE_GLOBAL"
	// opcodeMap[225] = "EXTENDED_ARG"
	// opcodeMap[86] = "SLICE_1"
	// opcodeMap[77] = "INPLACE_DIVIDE"
	// opcodeMap[25] = "YIELD_VALUE"
	// opcodeMap[228] = "STORE_NAME"
	// opcodeMap[222] = "STORE_FAST"
	// opcodeMap[67] = "UNARY_POSITIVE"
	// opcodeMap[108] = "SETUP_FINALLY"
	// opcodeMap[55] = "UNARY_NOT"
	// opcodeMap[20] = "BINARY_SUBTRACT"
	// opcodeMap[72] = "BINARY_AND"
	// opcodeMap[192] = "LOAD_FAST"
	// opcodeMap[229] = "BUILD_SLICE"
	// opcodeMap[78] = "BINARY_ADD"
	// opcodeMap[109] = "BUILD_LIST"
	// opcodeMap[49] = "STORE_MAP"
	// opcodeMap[0] = "DELETE_SUBSCR"
	// opcodeMap[47] = "END_FINALLY"
	// opcodeMap[50] = "POP_TOP"
	// opcodeMap[11] = "BINARY_SUBSCR"
	// opcodeMap[68] = "INPLACE_RSHIFT"
	// opcodeMap[8] = "IMPORT_STAR"
	// opcodeMap[219] = "RAISE_VARARGS"
	// opcodeMap[91] = "UNARY_INVERT"
	// opcodeMap[54] = "BINARY_RSHIFT"
	// opcodeMap[6] = "BINARY_XOR"
	// opcodeMap[13] = "DELETE_SLICE_1"
	// opcodeMap[82] = "INPLACE_ADD"
	// opcodeMap[22] = "BINARY_LSHIFT"
	// opcodeMap[80] = "BUILD_CLASS"
	// opcodeMap[254] = "JUMP_ABSOLUTE"
	// opcodeMap[29] = "STORE_SLICE_1"
	// opcodeMap[90] = "STORE_SUBSCR"
	// opcodeMap[110] = "LIST_APPEND"
	// opcodeMap[32] = "ROT_FOUR"
	// opcodeMap[43] = "INPLACE_OR"
	// opcodeMap[244] = "LOAD_GLOBAL"
	// opcodeMap[118] = "CONTINUE_LOOP"
	// opcodeMap[247] = "SETUP_EXCEPT"
	// opcodeMap[66] = "INPLACE_TRUE_DIVIDE"
	// opcodeMap[71] = "WITH_CLEANUP"
	// opcodeMap[185] = "IMPORT_NAME"
	// opcodeMap[136] = "JUMP_IF_FALSE_OR_POP"
	// opcodeMap[155] = "SET_ADD"
	// opcodeMap[9] = "UNARY_NEGATIVE"
	// opcodeMap[24] = "PRINT_NEWLINE_TO"
	// opcodeMap[53] = "PRINT_NEWLINE"
	// opcodeMap[169] = "DELETE_NAME"
	// opcodeMap[119] = "JUMP_FORWARD"
	// opcodeMap[44] = "UNARY_CONVERT"
	// opcodeMap[89] = "INPLACE_MODULO"
	// opcodeMap[16] = "BINARY_OR"
	// opcodeMap[172] = "BUILD_MAP"
	// opcodeMap[73] = "INPLACE_MULTIPLY"
	// opcodeMap[103] = "JUMP_IF_TRUE_OR_POP"
	// opcodeMap[148] = "STORE_DEREF"
	// opcodeMap[46] = "ROT_THREE"
	// opcodeMap[36] = "EXEC_STMT"
	// opcodeMap[199] = "BUILD_SET"
	// opcodeMap[101] = "COMPARE_OP"
	// opcodeMap[173] = "SETUP_WITH"
	// opcodeMap[83] = "DUP_TOP"
	// opcodeMap[14] = "DELETE_SLICE_2"
	// opcodeMap[152] = "CALL_FUNCTION_KW"
	// opcodeMap[95] = "MAP_ADD"
	// opcodeMap[15] = "DELETE_SLICE_3"
	// opcodeMap[41] = "INPLACE_LSHIFT"
	// opcodeMap[69] = "INPLACE_AND"
	// opcodeMap[215] = "MAKE_FUNCTION"
	// opcodeMap[151] = "CALL_FUNCTION_VAR"
	// opcodeMap[120] = "STORE_ATTR"
	// opcodeMap[52] = "BINARY_MODULO"
	// opcodeMap[12] = "DELETE_SLICE_0"
	// opcodeMap[31] = "STORE_SLICE_3"
	// opcodeMap[76] = "PRINT_ITEM_TO"
	// opcodeMap[153] = "CALL_FUNCTION_VAR_KW"

	// file, err := os.Open("./opcode/ca49b20f.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// scanner := bufio.NewScanner(file)
	// opmap := make(map[byte]byte)
	// var i byte = 0
	// for scanner.Scan() {
	// 	value, _ := strconv.ParseInt(scanner.Text(), 16, 64)
	// 	opmap[i] = byte(value)
	// 	i++
	// }
	// file.Close()

	// file, err = os.Open("./opcode/bc58dbd5.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// scanner = bufio.NewScanner(file)
	// i = 0
	// for scanner.Scan() {
	// 	value, _ := strconv.ParseInt(scanner.Text(), 16, 64)
	// 	name := opcodeMap[opmap[i]]
	// 	fmt.Printf("opcodeMap[%d]=\"%s\"\n", byte(value), name)
	// 	i++
	// }
	// file.Close()

	// return

	// dat1, err := os.ReadFile("./extracted/1.pyc")
	// if err != nil {
	// 	panic(err)
	// }

	// repaired, err := pymarshal.PycRepairer(dat1, false)
	// if err != nil {
	// 	panic(err)
	// }

	// os.WriteFile("repaired/1.pyc", repaired, 0777)

	// return

	
	exe_name := os.Args[0]

	if len(os.Args) != 2 {
		fmt.Println("Usage: ./" + exe_name +" input_file_path")
		return
	}

	input_file_path := os.Args[1]
	

	  // 文件夹名称
	  folderNames := []string{"decompiled", "extracted"}

	  // 遍历所有需要创建的文件夹
	  for _, folderName := range folderNames {
		  // 检查文件夹是否存在
		  if _, err := os.Stat(folderName); os.IsNotExist(err) {
			  // 文件夹不存在，创建新文件夹
			  err := os.Mkdir(folderName, 0755)
			  if err != nil {
				  // 如果创建文件夹失败，输出错误信息
				  panic(err)
			  }
			  fmt.Printf("初始化 %s\n", folderName)
		  } else {
			  // 文件夹已经存在
			  fmt.Printf("%s 已经存在，请您备份删除后再使用本程序\n", folderName)
			  os.Exit(0)
		  }
	}

	time.Sleep(3 * time.Second)

	asdf_dn := "j2h56ogodh3sk"
	asdf_dt := "=dziaq;"
	asdf_df := "|`o=5v7!\"-276"

	asdf_tm := strings.Repeat(asdf_dn, 4) + strings.Repeat(asdf_dt+asdf_dn+asdf_df, 5) + "$@" + strings.Repeat(asdf_dt, 7) + strings.Repeat(asdf_df, 2) + "&]`"

	rotor = mcpk.NewRotorObj(6, asdf_tm)

	dat, err := os.ReadFile(input_file_path)
	// dat, err := os.ReadFile("./vanilla1.mcp")
	if err != nil {
		panic(err)
	}

	unpack(dat)

	println("解密完成！开发者nb！")
	println("即将执行python来解密pyc文件，如果您需要手动解密请退出程序...")
	// err = RepairMarshal(readBuffer, &result)

	time.Sleep(5 * time.Second)


	fmt.Println("_  __     _      _     _                                    ____ ")
fmt.Println("| |/ /__ _| | ___(_) __| | ___  ___  ___ ___  _ __   ___    |  _ \\  ___  ___ ")
fmt.Println("| ' // _` | |/ _ \\ |/ _` |/ _ \\/ __|/ __/ _ \\| '_ \\ / _ \\   | | | |/ _ \\/ __| ")
fmt.Println("| . \\ (_| | |  __/ | (_| | (_) \\__ \\ (_| (_) | |_) |  __/   | |_| |  __/ (__ ")
fmt.Println("|_|\\_\\__,_|_|\\___|_|\\__,_|\\___/|___/\\___\\___/| .__/ \\___|___|____/ \\___|\\___|")
fmt.Println("                                             |_|       |_____|          ")
fmt.Println("           _       ")
fmt.Println("  ___   __| | ___ _ __ ")
fmt.Println(" / _ \\ / _` |/ _ \\ '__|")
fmt.Println("| (_) | (_| |  __/ |   ")
fmt.Println(" \\___/ \\__,_|\\___|_|   ")


	 // 检查script文件夹是否存在
	 if _, err := os.Stat("script"); os.IsNotExist(err) {
        println("script文件夹不存在")
    }

    // 执行batch_decompile.py脚本
    cmd := exec.Command("C:/Python27/python.exe", "script/batch_decompile.py")
	output, err := cmd.CombinedOutput()

    err = cmd.Run()
    if err != nil {
        println("无法执行batch_decompile.py脚本")
    }
	fmt.Println(string(output))

	// filepath.Walk("./target", func(path string, info os.FileInfo, err error) error {
	// 	if err != nil {
	// 		return err
	// 	}

	// 	if !info.IsDir() {
	// 		println(path)
	// 		data, err := ioutil.ReadFile(path)
	// 		if err != nil {
	// 			return fmt.Errorf("failed to read file %s: %v", path, err)
	// 		}

	// 		unpack(data)
	// 	}

	// 	return nil
	// })

	// mcpkfs, err := mcpk.NewMCPKFileSystem(dat)
	// if err != nil {
	// 	panic(err)
	// }
	// println("directories:", len(mcpkfs.DirectoryEntryList))
	// file, err := mcpkfs.Open("mod/server/system/sendSimpleMsgSysServer.mcs")
	// if err != nil {
	// 	panic(err)
	// }
	// println(string(file))

	// return

}

func unpack(dat []byte) {
	start := binary.LittleEndian.Uint32(dat[0x0c:])
	middle := binary.LittleEndian.Uint32(dat[0x10:])
	end := binary.LittleEndian.Uint32(dat[0x14:])
	fmt.Printf("%x, %x, %x\n", start, middle, end)

	directory_table := dat[start:middle]
	var dirId, dirOff, dirSize, v1, v2, v3, v4 uint32
	for i := 0; i < len(directory_table); i += 0xc {
		dirId = binary.LittleEndian.Uint32(directory_table[i:])
		dirOff = binary.LittleEndian.Uint32(directory_table[i+4:])
		dirSize = binary.LittleEndian.Uint32(directory_table[i+8:])
		index_table := dat[middle+dirOff : middle+dirOff+(dirSize*0x10)]
		for i := 0; i < len(index_table); i += 0x10 {
			v1 = binary.LittleEndian.Uint32(index_table[i:])
			v2 = binary.LittleEndian.Uint32(index_table[i+4:])
			v3 = binary.LittleEndian.Uint32(index_table[i+8:])
			v4 = binary.LittleEndian.Uint32(index_table[i+12:])
			// if v1 == idRedirect {
			data := decrypt(dat[end+v2:end+v2+v3], v1)

			println("size:", len(data))

			repaired, err := pymarshal.PycRepairer(data, v1 != idRedirect)
			if err != nil {
				panic(err)
			}

			os.WriteFile(fmt.Sprintf("extracted/%x_%x.pyc", dirId, v1), repaired, 0777)
			fmt.Printf("%x\t%x\t%x\t%x\n", v1, v2, v3, v4)

			
			// }
		}
	}
}

func decrypt(data []byte, fileId uint32) []byte {
	decrypted := make([]byte, len(data))
	copy(decrypted, data)

	// decrypt
	if fileId == idRedirect {
		xorData := []byte("MCPK")
		for i := 0; i < len(xorData); i++ {
			decrypted[i] ^= xorData[i]
		}
	} else {
		rotor.DRegion(decrypted, true)
	}

	// unzip
	var unzipped bytes.Buffer
	reader, err := zlib.NewReader(bytes.NewBuffer(decrypted))
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(&unzipped, reader)
	if err != io.EOF && err != nil {
		panic(err)
	}
	reader.Close()

	if fileId != idRedirect {
		decrypted = reverseData(unzipped.Bytes())
	} else {
		decrypted = unzipped.Bytes()
	}

	return decrypted
}

func reverseData(data []byte) []byte {
	length := len(data)

	for i := 0; i < 130 && i < length; i++ {
		data[i] = data[i] ^ 156
	}

	for i := 0; i < length/2; i++ {
		data[i], data[length-i-1] = data[length-i-1], data[i]
	}

	return data
}
