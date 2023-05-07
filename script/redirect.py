# coding=utf-8
# Embedded file name: redirect.py


def init_rotor():
    asdf_dn = 'j2h56ogodh3sk'
    asdf_dt = '=dziaq;'
    asdf_df = '|`o=5v7!"-276'
    asdf_tm = asdf_dn * 4 + (asdf_dt + asdf_dn + asdf_df) * 5 + '$' + '@' + asdf_dt * 7 + asdf_df * 2 + '&' + ']' + '`'
    import rotor
    return rotor.newrotor(asdf_tm)


def file_exist(file):
    try:
        f = open(file)
        f.close()
        return True
    except IOError:
        return False


class McpImporter(object):
    rot = init_rotor()
    Ext = '.mcs'
    Py = '.py'

    def __init__(self, path):
        self._path = path
        self._isMcp = None

    def find_module(self, fullname, path = None):
        import fop
        if path is None:
            path = self._path
        fullname = fullname.replace('.', '/')
        pkg_name = ''.join((fullname, '/__init__', McpImporter.Ext))
        py_pkg_name = ''.join((fullname, '/__init__', McpImporter.Py))
        if fop.find_file(pkg_name, path):
            self._isMcp = True
            return self
        elif file_exist(''.join((path, py_pkg_name))):
            self._isMcp = False
            return self
        extFullname = fullname + McpImporter.Ext
        pyFullname = fullname + McpImporter.Py
        if fop.find_file(extFullname, path):
            self._isMcp = True
            return self
        elif file_exist(''.join((path, pyFullname))):
            self._isMcp = False
            return self
        else:
            return

    def load_module(self, fullname):
        import fop, marshal, zlib
        is_pkg = True
        mod_path = fullname.replace('.', '/') + '/__init__'
        mod_name = fullname
        if not self._isMcp:
            if not file_exist(''.join((self._path, mod_path + McpImporter.Py))):
                is_pkg = False
                mod_path = fullname.replace('.', '/')
                mod_name = fullname
        elif not fop.find_file(mod_path + McpImporter.Ext, self._path):
            is_pkg = False
            mod_path = fullname.replace('.', '/')
            mod_name = fullname
        if not self._isMcp:
            with open(''.join((self._path, mod_path + McpImporter.Py)), 'r') as f:
                data = f.read()
                code = compile(data, mod_name, 'exec')
        else:
            data = fop.get_file(mod_path + McpImporter.Ext, self._path)
            data = McpImporter.rot.decrypt(data)
            data = zlib.decompress(data)
            data = self._reverse_data(data)
            code = marshal.loads(data)
        path = None
        if is_pkg:
            path = [self._path]
        mod = fop.new_module(mod_name, code, path)
        if not mod:
            print 'import mod error', mod
        return mod

    def _reverse_data(self, data):
        L = list(data)
        L = map(lambda ch: chr(ord(ch) ^ 156), L[:130]) + L[130:]
        L.reverse()
        return ''.join(L)


import sys
del sys.path_hooks[:]
sys.path_hooks.append(McpImporter)
sys.path_importer_cache.clear()
