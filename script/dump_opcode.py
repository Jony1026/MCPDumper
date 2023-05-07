import opcode

def dump_opcodes():
    for opname, opnum in sorted(opcode.opmap.items(), key=lambda x: x[1]):
        print("opcodes[\"{}\"] = {}".format(opname.replace('+', '_'), opnum))

if __name__ == "__main__":
    dump_opcodes()