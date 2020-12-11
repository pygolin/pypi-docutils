package __main__

import (
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/lxml"
	_ "github.com/pygolin/stdlib/pkg/os"
	_ "github.com/pygolin/stdlib/pkg/shutil"
	_ "github.com/pygolin/stdlib/pkg/sys"
	_ "github.com/pygolin/stdlib/pkg/tempfile"
	_ "github.com/pygolin/stdlib/pkg/zipfile"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils-0.16.data/scripts/rst2odt_prepstyles.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ßNAMESPACES := πg.InternStr("NAMESPACES")
		ßZIP_DEFLATED := πg.InternStr("ZIP_DEFLATED")
		ßZipFile := πg.InternStr("ZipFile")
		ß__doc__ := πg.InternStr("__doc__")
		ß__main__ := πg.InternStr("__main__")
		ß__name__ := πg.InternStr("__name__")
		ßargv := πg.InternStr("argv")
		ßattrib := πg.InternStr("attrib")
		ßclose := πg.InternStr("close")
		ßetree := πg.InternStr("etree")
		ßexit := πg.InternStr("exit")
		ßfdopen := πg.InternStr("fdopen")
		ßfilename := πg.InternStr("filename")
		ßfo := πg.InternStr("fo")
		ßfromstring := πg.InternStr("fromstring")
		ßinfolist := πg.InternStr("infolist")
		ßlen := πg.InternStr("len")
		ßmain := πg.InternStr("main")
		ßmkstemp := πg.InternStr("mkstemp")
		ßmove := πg.InternStr("move")
		ßos := πg.InternStr("os")
		ßprepstyle := πg.InternStr("prepstyle")
		ßprint := πg.InternStr("print")
		ßread := πg.InternStr("read")
		ßshutil := πg.InternStr("shutil")
		ßstartswith := πg.InternStr("startswith")
		ßstderr := πg.InternStr("stderr")
		ßstyle := πg.InternStr("style")
		ßsys := πg.InternStr("sys")
		ßtostring := πg.InternStr("tostring")
		ßw := πg.InternStr("w")
		ßwritestr := πg.InternStr("writestr")
		ßxpath := πg.InternStr("xpath")
		ßzipfile := πg.InternStr("zipfile")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 *πg.Dict
		_ = πTemp004
		var πTemp005 []πg.Param
		_ = πTemp005
		var πTemp006 *πg.Object
		_ = πTemp006
		var πTemp007 *πg.Object
		_ = πTemp007
		var πTemp008 bool
		_ = πTemp008
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default:
				panic("unexpected function state")
			}
			// line 7: """
			πF.SetLineno(7)
			// line 7: """
			πF.SetLineno(7)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nFix a word-processor-generated styles.odt for odtwriter use: Drop page size\nspecifications from styles.xml in STYLE_FILE.odt.\n").ToObject()); πE != nil {
				continue
			}
			// line 14: from __future__ import print_function
			πF.SetLineno(14)
			// line 16: from lxml import etree
			πF.SetLineno(16)
			if πTemp002, πE = πg.ImportModule(πF, "lxml"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßetree); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßetree.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 17: import sys
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 18: import zipfile
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "zipfile"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßzipfile.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 19: from tempfile import mkstemp
			πF.SetLineno(19)
			if πTemp002, πE = πg.ImportModule(πF, "tempfile"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßmkstemp); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmkstemp.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 20: import shutil
			πF.SetLineno(20)
			if πTemp002, πE = πg.ImportModule(πF, "shutil"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßshutil.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 21: import os
			πF.SetLineno(21)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 23: NAMESPACES = {
			πF.SetLineno(23)
			πTemp004 = πg.NewDict()
			if πE = πTemp004.SetItem(πF, ßstyle.ToObject(), πg.NewStr("urn:oasis:names:tc:opendocument:xmlns:style:1.0").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ßfo.ToObject(), πg.NewStr("urn:oasis:names:tc:opendocument:xmlns:xsl-fo-compatible:1.0").ToObject()); πE != nil {
				continue
			}
			πTemp001 = πTemp004.ToObject()
			if πE = πF.Globals().SetItem(πF, ßNAMESPACES.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 29: def prepstyle(filename):
			πF.SetLineno(29)
			πTemp005 = make([]πg.Param, 1)
			πTemp005[0] = πg.Param{Name: "filename", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("prepstyle", "/usr/lib/python2.7/site-packages/docutils-0.16.data/scripts/rst2odt_prepstyles.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfilename *πg.Object = πArgs[0]
				_ = µfilename
				var µzin *πg.Object = πg.UnboundLocal
				_ = µzin
				var µstyles *πg.Object = πg.UnboundLocal
				_ = µstyles
				var µroot *πg.Object = πg.UnboundLocal
				_ = µroot
				var µel *πg.Object = πg.UnboundLocal
				_ = µel
				var µattr *πg.Object = πg.UnboundLocal
				_ = µattr
				var µtempname *πg.Object = πg.UnboundLocal
				_ = µtempname
				var µzout *πg.Object = πg.UnboundLocal
				_ = µzout
				var µitem *πg.Object = πg.UnboundLocal
				_ = µitem
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 πg.KWArgs
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 []*πg.Object
				_ = πTemp012
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1:
						goto Label1
					case 2:
						goto Label2
					case 4:
						goto Label4
					case 5:
						goto Label5
					case 9:
						goto Label9
					case 10:
						goto Label10
					default:
						panic("unexpected function state")
					}
					// line 31: zin = zipfile.ZipFile(filename)
					πF.SetLineno(31)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp001[0] = µfilename
					if πTemp002, πE = πg.ResolveGlobal(πF, ßzipfile); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßZipFile, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µzin = πTemp002
					// line 32: styles = zin.read("styles.xml")
					πF.SetLineno(32)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("styles.xml").ToObject()
					if πE = πg.CheckLocal(πF, µzin, "zin"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µzin, ßread, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µstyles = πTemp003
					// line 34: root = etree.fromstring(styles)
					πF.SetLineno(34)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µstyles, "styles"); πE != nil {
						continue
					}
					πTemp001[0] = µstyles
					if πTemp002, πE = πg.ResolveGlobal(πF, ßetree); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßfromstring, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µroot = πTemp002
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("//style:page-layout-properties").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNAMESPACES); πE != nil {
						continue
					}
					πTemp004 = πg.KWArgs{
						{"namespaces", πTemp003},
					}
					if πE = πg.CheckLocal(πF, µroot, "root"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µroot, ßxpath, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, πTemp004); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp006 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label3
					}
					if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp007 = !isStop
					} else {
						πTemp007 = true
						µel = πTemp003
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(1)
					if πE = πg.CheckLocal(πF, µel, "el"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µel, ßattrib, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Iter(πF, πTemp005); πE != nil {
						continue
					}
					πF.PushCheckpoint(5)
					πTemp007 = false
				Label4:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp007 {
						πF.PopCheckpoint()
						goto Label6
					}
					if πTemp005, πE = πg.Next(πF, πTemp003); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp008 = !isStop
					} else {
						πTemp008 = true
						µattr = πTemp005
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(4)
					πTemp001 = πF.MakeArgs(1)
					πTemp009 = ßfo.ToObject()
					if πTemp011, πE = πg.ResolveGlobal(πF, ßNAMESPACES); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, πTemp011, πTemp009); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Mod(πF, πg.NewStr("{%s}").ToObject(), πTemp010); πE != nil {
						continue
					}
					πTemp001[0] = πTemp005
					if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µattr, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp008, πE = πg.IsTrue(πF, πTemp009); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label7
					}
					goto Label8
					// line 38: if attr.startswith("{%s}" % NAMESPACES["fo"]):
					πF.SetLineno(38)
				Label7:
					// line 39: del el.attrib[attr]
					πF.SetLineno(39)
					if πE = πg.CheckLocal(πF, µel, "el"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µel, ßattrib, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
						continue
					}
					πTemp009 = µattr
					if πE = πg.DelItem(πF, πTemp005, πTemp009); πE != nil {
						continue
					}
					goto Label8
				Label8:
					continue
				Label5:
					if πE != nil || πR != nil {
						continue
					}
				Label6:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 41: tempname = mkstemp()
					πF.SetLineno(41)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmkstemp); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µtempname = πTemp003
					// line 42: zout = zipfile.ZipFile(os.fdopen(tempname[0], "w"), "w",
					πF.SetLineno(42)
					πTemp001 = πF.MakeArgs(3)
					πTemp012 = πF.MakeArgs(2)
					πTemp002 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µtempname, "tempname"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µtempname, πTemp002); πE != nil {
						continue
					}
					πTemp012[0] = πTemp003
					πTemp012[1] = ßw.ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßfdopen, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp012, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp012)
					πTemp001[0] = πTemp002
					πTemp001[1] = ßw.ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßzipfile); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßZIP_DEFLATED, nil); πE != nil {
						continue
					}
					πTemp001[2] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßzipfile); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßZipFile, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µzout = πTemp002
					if πE = πg.CheckLocal(πF, µzin, "zin"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µzin, ßinfolist, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
						continue
					}
					πF.PushCheckpoint(10)
					πTemp006 = false
				Label9:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label11
					}
					if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp007 = !isStop
					} else {
						πTemp007 = true
						µitem = πTemp003
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(9)
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µitem, ßfilename, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp005, πg.NewStr("styles.xml").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label12
					}
					goto Label13
					// line 46: if item.filename == "styles.xml":
					πF.SetLineno(46)
				Label12:
					// line 47: zout.writestr(item, etree.tostring(root))
					πF.SetLineno(47)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πTemp001[0] = µitem
					πTemp012 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µroot, "root"); πE != nil {
						continue
					}
					πTemp012[0] = µroot
					if πTemp003, πE = πg.ResolveGlobal(πF, ßetree); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßtostring, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp005.Call(πF, πTemp012, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp012)
					πTemp001[1] = πTemp003
					if πE = πg.CheckLocal(πF, µzout, "zout"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µzout, ßwritestr, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label14
				Label13:
					// line 49: zout.writestr(item, zin.read(item.filename))
					πF.SetLineno(49)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πTemp001[0] = µitem
					πTemp012 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µitem, ßfilename, nil); πE != nil {
						continue
					}
					πTemp012[0] = πTemp003
					if πE = πg.CheckLocal(πF, µzin, "zin"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µzin, ßread, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp012, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp012)
					πTemp001[1] = πTemp005
					if πE = πg.CheckLocal(πF, µzout, "zout"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µzout, ßwritestr, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label14
				Label14:
					continue
				Label10:
					if πE != nil || πR != nil {
						continue
					}
				Label11:
					// line 51: zout.close()
					πF.SetLineno(51)
					if πE = πg.CheckLocal(πF, µzout, "zout"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µzout, ßclose, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 52: zin.close()
					πF.SetLineno(52)
					if πE = πg.CheckLocal(πF, µzin, "zin"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µzin, ßclose, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 53: shutil.move(tempname[1], filename)
					πF.SetLineno(53)
					πTemp001 = πF.MakeArgs(2)
					πTemp002 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µtempname, "tempname"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µtempname, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp001[1] = µfilename
					if πTemp002, πE = πg.ResolveGlobal(πF, ßshutil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmove, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßprepstyle.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 56: def main():
			πF.SetLineno(56)
			πTemp005 = make([]πg.Param, 0)
			πTemp003 = πg.NewFunction(πg.NewCode("main", "/usr/lib/python2.7/site-packages/docutils-0.16.data/scripts/rst2odt_prepstyles.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargs *πg.Object = πg.UnboundLocal
				_ = µargs
				var µfilename *πg.Object = πg.UnboundLocal
				_ = µfilename
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 πg.KWArgs
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 57: args = sys.argv[1:]
					πF.SetLineno(57)
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßargv, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
						continue
					}
					µargs = πTemp002
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					πTemp005[0] = µargs
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp001, πE = πg.NE(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label1
					}
					goto Label2
					// line 58: if len(args) != 1:
					πF.SetLineno(58)
				Label1:
					// line 59: print(__doc__, file=sys.stderr)
					πF.SetLineno(59)
					πTemp005 = πF.MakeArgs(1)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß__doc__); πE != nil {
						continue
					}
					πTemp005[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstderr, nil); πE != nil {
						continue
					}
					πTemp007 = πg.KWArgs{
						{"file", πTemp002},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, πTemp007); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 60: print("Usage: %s STYLE_FILE.odt\n" % sys.argv[0], file=sys.stderr)
					πF.SetLineno(60)
					πTemp005 = πF.MakeArgs(1)
					πTemp002 = πg.NewInt(0).ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, πTemp004, ßargv, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp008, πTemp002); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("Usage: %s STYLE_FILE.odt\n").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp005[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstderr, nil); πE != nil {
						continue
					}
					πTemp007 = πg.KWArgs{
						{"file", πTemp002},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, πTemp007); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 61: sys.exit(1)
					πF.SetLineno(61)
					πTemp005 = πF.MakeArgs(1)
					πTemp005[0] = πg.NewInt(1).ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßexit, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					goto Label2
				Label2:
					// line 62: filename = args[0]
					πF.SetLineno(62)
					πTemp001 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
						continue
					}
					µfilename = πTemp002
					// line 63: prepstyle(filename)
					πF.SetLineno(63)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
						continue
					}
					πTemp005[0] = µfilename
					if πTemp001, πE = πg.ResolveGlobal(πF, ßprepstyle); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßmain.ToObject(), πTemp003); πE != nil {
				continue
			}
			if πTemp007, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp006, πE = πg.Eq(πF, πTemp007, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp008, πE = πg.IsTrue(πF, πTemp006); πE != nil {
				continue
			}
			if πTemp008 {
				goto Label1
			}
			goto Label2
			// line 66: if __name__ == '__main__':
			πF.SetLineno(66)
		Label1:
			// line 67: main()
			πF.SetLineno(67)
			if πTemp006, πE = πg.ResolveGlobal(πF, ßmain); πE != nil {
				continue
			}
			if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("__main__", Code)
}
