import marshal
import dis
import sys

def print_code_objects(code_obj):
    # print(f"{' ' * indent}Code object: {code_obj}")
    # print(f"{' ' * indent}Filename: {code_obj.co_filename}")
    # print(f"{' ' * indent}Disassembled bytecode:")
    dis.dis(code_obj)
    # print()

    for const in code_obj.co_consts:
        if isinstance(const, type(code_obj)):
            print_code_objects(const)

# file = open("./repaired/1.pyc", "rb").read()
# file = open("./extracted/bfa60943.pyc", "rb").read()[8:]
# file = open("./script/redirect.pyc", "rb").read()[8:]
# file = open("./repaired/redirect-fix.pyc", "rb").read()[8:]
m = marshal.loads(file)

print '# ', m.co_filename
print_code_objects(m)