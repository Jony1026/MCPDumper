import uncompyle2

with open("./repaired/4.py", "wb") as fileobj:
    uncompyle2.uncompyle_file("./extracted/b1f41873.pyc", fileobj)
    # uncompyle2.uncompyle_file("./repaired/1.pyc", fileobj)