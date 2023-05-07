import os
import uncompyle2
import marshal
import traceback

# Set the path of the directory you want to iterate through
dir_path = "./extracted"
decompiled_path = "./decompiled"

# Iterate through the directory and print the name of each file
for file_name in os.listdir(dir_path):
    source_path = os.path.join(dir_path, file_name)
    file = marshal.loads(open(source_path, "rb").read()[8:])
    file_path = os.path.join(decompiled_path, file.co_filename)
    try:
        os.makedirs(os.path.dirname(file_path))
    except:
        pass
    try:
        with open(file_path, "wb") as f:
            f.write('# file: ' + source_path + '\n')
            uncompyle2.uncompyle_file(source_path, f)
    except:
        traceback.print_exc()
        print "decomp failed", file_name, file.co_filename
    # decoded = ""
    # with open(file_path, 'r') as file:
    #     escaped_string = file.read()
    #     decoded = escaped_string.decode('string_escape')
    # with open(file_path, 'w') as file:
    #     file.write(decoded)

