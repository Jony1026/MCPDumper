# -*- coding: utf-8 -*-
import codecs
import re
from cStringIO import StringIO

class StringBuilder:
     _file_str = None

     def __init__(self):
         self._file_str = StringIO()

     def append(self, str):
         self._file_str.write(str)

     def toString(self):
         return self._file_str.getvalue()

def __decode_unicode(match):
    return codecs.decode(r'\u{}'.format(match.group(1)), 'unicode-escape').encode("utf-8")

def better_repr(var):
    if type(var) == str:
        # var = var.replace('\\', '\\\\').replace('\n', '\\n')
        # if "'" in var and '"' not in var:
        #     return '"' + var + '"'
        # return "'" + var.replace("'", "\\'") + "'"
        try:
            escaped_string = repr(var.decode("utf-8", errors='replace'))[1:]
            return re.sub(r'\\u([0-9a-fA-F]{4})', lambda m: __decode_unicode(m), escaped_string)
        except:
            pass
    elif type(var) == tuple:
        if len(var) == 0:
            return "()"
        li = list(var)
        sb = StringBuilder()
        sb.append("(")
        for e in li:
            sb.append(better_repr(e))
            sb.append(", ")
        return sb.toString()[:-2] + ")"
    return repr(var)