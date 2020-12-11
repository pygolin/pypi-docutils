package latex2mathml

import (
	_ "github.com/pygolin/pypi-docutils"
	πg "github.com/pygolin/runtime"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ßA := πg.InternStr("A")
		ßB := πg.InternStr("B")
		ßC := πg.InternStr("C")
		ßD := πg.InternStr("D")
		ßDelta := πg.InternStr("Delta")
		ßE := πg.InternStr("E")
		ßF := πg.InternStr("F")
		ßFalse := πg.InternStr("False")
		ßG := πg.InternStr("G")
		ßGamma := πg.InternStr("Gamma")
		ßGreek := πg.InternStr("Greek")
		ßH := πg.InternStr("H")
		ßI := πg.InternStr("I")
		ßJ := πg.InternStr("J")
		ßK := πg.InternStr("K")
		ßL := πg.InternStr("L")
		ßLambda := πg.InternStr("Lambda")
		ßM := πg.InternStr("M")
		ßN := πg.InternStr("N")
		ßNone := πg.InternStr("None")
		ßO := πg.InternStr("O")
		ßOmega := πg.InternStr("Omega")
		ßP := πg.InternStr("P")
		ßPhi := πg.InternStr("Phi")
		ßPi := πg.InternStr("Pi")
		ßPr := πg.InternStr("Pr")
		ßPsi := πg.InternStr("Psi")
		ßQ := πg.InternStr("Q")
		ßR := πg.InternStr("R")
		ßS := πg.InternStr("S")
		ßSigma := πg.InternStr("Sigma")
		ßSyntaxError := πg.InternStr("SyntaxError")
		ßT := πg.InternStr("T")
		ßTheta := πg.InternStr("Theta")
		ßTrue := πg.InternStr("True")
		ßU := πg.InternStr("U")
		ßUpsilon := πg.InternStr("Upsilon")
		ßV := πg.InternStr("V")
		ßW := πg.InternStr("W")
		ßX := πg.InternStr("X")
		ßXi := πg.InternStr("Xi")
		ßY := πg.InternStr("Y")
		ßZ := πg.InternStr("Z")
		ß_ := πg.InternStr("_")
		ß__class__ := πg.InternStr("__class__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__repr__ := πg.InternStr("__repr__")
		ßa := πg.InternStr("a")
		ßacute := πg.InternStr("acute")
		ßappend := πg.InternStr("append")
		ßarccos := πg.InternStr("arccos")
		ßarcsin := πg.InternStr("arcsin")
		ßarctan := πg.InternStr("arctan")
		ßarg := πg.InternStr("arg")
		ßattrs := πg.InternStr("attrs")
		ßb := πg.InternStr("b")
		ßbar := πg.InternStr("bar")
		ßbegin := πg.InternStr("begin")
		ßbold := πg.InternStr("bold")
		ßbreve := πg.InternStr("breve")
		ßc := πg.InternStr("c")
		ßcheck := πg.InternStr("check")
		ßchildren := πg.InternStr("children")
		ßclose := πg.InternStr("close")
		ßclosepar := πg.InternStr("closepar")
		ßcolon := πg.InternStr("colon")
		ßcos := πg.InternStr("cos")
		ßcosh := πg.InternStr("cosh")
		ßcot := πg.InternStr("cot")
		ßcoth := πg.InternStr("coth")
		ßcsc := πg.InternStr("csc")
		ßd := πg.InternStr("d")
		ßdata := πg.InternStr("data")
		ßdddot := πg.InternStr("dddot")
		ßddot := πg.InternStr("ddot")
		ßdeg := πg.InternStr("deg")
		ßdelete_child := πg.InternStr("delete_child")
		ßdet := πg.InternStr("det")
		ßdim := πg.InternStr("dim")
		ßdot := πg.InternStr("dot")
		ße := πg.InternStr("e")
		ßend := πg.InternStr("end")
		ßexp := πg.InternStr("exp")
		ßextend := πg.InternStr("extend")
		ßf := πg.InternStr("f")
		ßfind := πg.InternStr("find")
		ßfrac := πg.InternStr("frac")
		ßfull := πg.InternStr("full")
		ßfunctions := πg.InternStr("functions")
		ßg := πg.InternStr("g")
		ßgcd := πg.InternStr("gcd")
		ßget := πg.InternStr("get")
		ßgrave := πg.InternStr("grave")
		ßh := πg.InternStr("h")
		ßhandle_keyword := πg.InternStr("handle_keyword")
		ßhasattr := πg.InternStr("hasattr")
		ßhat := πg.InternStr("hat")
		ßhom := πg.InternStr("hom")
		ßi := πg.InternStr("i")
		ßinf := πg.InternStr("inf")
		ßinjlim := πg.InternStr("injlim")
		ßinline := πg.InternStr("inline")
		ßint := πg.InternStr("int")
		ßisalpha := πg.InternStr("isalpha")
		ßisdigit := πg.InternStr("isdigit")
		ßisinstance := πg.InternStr("isinstance")
		ßisupper := πg.InternStr("isupper")
		ßitems := πg.InternStr("items")
		ßj := πg.InternStr("j")
		ßjoin := πg.InternStr("join")
		ßk := πg.InternStr("k")
		ßker := πg.InternStr("ker")
		ßl := πg.InternStr("l")
		ßleft := πg.InternStr("left")
		ßlen := πg.InternStr("len")
		ßletters := πg.InternStr("letters")
		ßlg := πg.InternStr("lg")
		ßlim := πg.InternStr("lim")
		ßliminf := πg.InternStr("liminf")
		ßlimsup := πg.InternStr("limsup")
		ßlist := πg.InternStr("list")
		ßln := πg.InternStr("ln")
		ßlog := πg.InternStr("log")
		ßm := πg.InternStr("m")
		ßmath := πg.InternStr("math")
		ßmathalpha := πg.InternStr("mathalpha")
		ßmathbb := πg.InternStr("mathbb")
		ßmathbf := πg.InternStr("mathbf")
		ßmathbin := πg.InternStr("mathbin")
		ßmathcal := πg.InternStr("mathcal")
		ßmathclose := πg.InternStr("mathclose")
		ßmathfence := πg.InternStr("mathfence")
		ßmathop := πg.InternStr("mathop")
		ßmathopen := πg.InternStr("mathopen")
		ßmathord := πg.InternStr("mathord")
		ßmathrel := πg.InternStr("mathrel")
		ßmathring := πg.InternStr("mathring")
		ßmathrm := πg.InternStr("mathrm")
		ßmathscr := πg.InternStr("mathscr")
		ßmax := πg.InternStr("max")
		ßmfenced := πg.InternStr("mfenced")
		ßmfrac := πg.InternStr("mfrac")
		ßmi := πg.InternStr("mi")
		ßmin := πg.InternStr("min")
		ßmn := πg.InternStr("mn")
		ßmo := πg.InternStr("mo")
		ßmover := πg.InternStr("mover")
		ßmroot := πg.InternStr("mroot")
		ßmrow := πg.InternStr("mrow")
		ßmspace := πg.InternStr("mspace")
		ßmsqrt := πg.InternStr("msqrt")
		ßmstyle := πg.InternStr("mstyle")
		ßmsub := πg.InternStr("msub")
		ßmsubsup := πg.InternStr("msubsup")
		ßmsup := πg.InternStr("msup")
		ßmtable := πg.InternStr("mtable")
		ßmtd := πg.InternStr("mtd")
		ßmtext := πg.InternStr("mtext")
		ßmtr := πg.InternStr("mtr")
		ßmunder := πg.InternStr("munder")
		ßmunderover := πg.InternStr("munderover")
		ßmx := πg.InternStr("mx")
		ßn := πg.InternStr("n")
		ßnchildren := πg.InternStr("nchildren")
		ßnegatables := πg.InternStr("negatables")
		ßnot := πg.InternStr("not")
		ßo := πg.InternStr("o")
		ßobject := πg.InternStr("object")
		ßoint := πg.InternStr("oint")
		ßopenpar := πg.InternStr("openpar")
		ßover := πg.InternStr("over")
		ßoverleftrightarrow := πg.InternStr("overleftrightarrow")
		ßp := πg.InternStr("p")
		ßparent := πg.InternStr("parent")
		ßparse_latex_math := πg.InternStr("parse_latex_math")
		ßprod := πg.InternStr("prod")
		ßprojlim := πg.InternStr("projlim")
		ßq := πg.InternStr("q")
		ßr := πg.InternStr("r")
		ßrepr := πg.InternStr("repr")
		ßreverse := πg.InternStr("reverse")
		ßreversed := πg.InternStr("reversed")
		ßright := πg.InternStr("right")
		ßs := πg.InternStr("s")
		ßsec := πg.InternStr("sec")
		ßsin := πg.InternStr("sin")
		ßsinh := πg.InternStr("sinh")
		ßspecial := πg.InternStr("special")
		ßsplit := πg.InternStr("split")
		ßsqrt := πg.InternStr("sqrt")
		ßstartswith := πg.InternStr("startswith")
		ßsum := πg.InternStr("sum")
		ßsumintprod := πg.InternStr("sumintprod")
		ßsup := πg.InternStr("sup")
		ßt := πg.InternStr("t")
		ßtan := πg.InternStr("tan")
		ßtanh := πg.InternStr("tanh")
		ßtex2mathml := πg.InternStr("tex2mathml")
		ßtex2unichar := πg.InternStr("tex2unichar")
		ßtext := πg.InternStr("text")
		ßtilde := πg.InternStr("tilde")
		ßtranslation := πg.InternStr("translation")
		ßu := πg.InternStr("u")
		ßupdate := πg.InternStr("update")
		ßv := πg.InternStr("v")
		ßvarinjlim := πg.InternStr("varinjlim")
		ßvarliminf := πg.InternStr("varliminf")
		ßvarlimsup := πg.InternStr("varlimsup")
		ßvarprojlim := πg.InternStr("varprojlim")
		ßvec := πg.InternStr("vec")
		ßw := πg.InternStr("w")
		ßx := πg.InternStr("x")
		ßxml := πg.InternStr("xml")
		ßxml_body := πg.InternStr("xml_body")
		ßxml_end := πg.InternStr("xml_end")
		ßxml_start := πg.InternStr("xml_start")
		ßy := πg.InternStr("y")
		ßz := πg.InternStr("z")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Dict
		_ = πTemp003
		var πTemp004 *πg.Object
		_ = πTemp004
		var πTemp005 []πg.Param
		_ = πTemp005
		var πTemp006 *πg.Object
		_ = πTemp006
		var πTemp007 *πg.Object
		_ = πTemp007
		var πTemp008 *πg.Object
		_ = πTemp008
		var πTemp009 *πg.Object
		_ = πTemp009
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default:
				panic("unexpected function state")
			}
			// line 18: """Convert LaTex math code into presentational MathML"""
			πF.SetLineno(18)
			// line 18: """Convert LaTex math code into presentational MathML"""
			πF.SetLineno(18)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Convert LaTex math code into presentational MathML").ToObject()); πE != nil {
				continue
			}
			// line 22: import docutils.utils.math.tex2unichar as tex2unichar
			πF.SetLineno(22)
			if πTemp002, πE = πg.ImportModule(πF, "docutils.utils.math.tex2unichar"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[3]
			if πE = πF.Globals().SetItem(πF, ßtex2unichar.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 25: over = {'acute':    u'\u00B4', # u'\u0301',
			πF.SetLineno(25)
			πTemp003 = πg.NewDict()
			if πE = πTemp003.SetItem(πF, ßacute.ToObject(), πg.NewUnicode("\xc2\xb4").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßbar.ToObject(), πg.NewUnicode("\xc2\xaf").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßbreve.ToObject(), πg.NewUnicode("\xcb\x98").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßcheck.ToObject(), πg.NewUnicode("\xcb\x87").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßdot.ToObject(), πg.NewUnicode("\xcb\x99").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßddot.ToObject(), πg.NewUnicode("\xc2\xa8").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßdddot.ToObject(), πg.NewUnicode("\xe2\x83\x9b").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßgrave.ToObject(), πg.NewUnicode("`").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßhat.ToObject(), πg.NewUnicode("^").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßmathring.ToObject(), πg.NewUnicode("\xcb\x9a").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßoverleftrightarrow.ToObject(), πg.NewUnicode("\xe2\x83\xa1").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßtilde.ToObject(), πg.NewUnicode("\xcb\x9c").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßvec.ToObject(), πg.NewUnicode("\xe2\x83\x97").ToObject()); πE != nil {
				continue
			}
			πTemp001 = πTemp003.ToObject()
			if πE = πF.Globals().SetItem(πF, ßover.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 40: Greek = { # Capital Greek letters: (upright in TeX style)
			πF.SetLineno(40)
			πTemp003 = πg.NewDict()
			if πE = πTemp003.SetItem(πF, ßPhi.ToObject(), πg.NewUnicode("\xce\xa6").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßXi.ToObject(), πg.NewUnicode("\xce\x9e").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßSigma.ToObject(), πg.NewUnicode("\xce\xa3").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßPsi.ToObject(), πg.NewUnicode("\xce\xa8").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßDelta.ToObject(), πg.NewUnicode("\xce\x94").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßTheta.ToObject(), πg.NewUnicode("\xce\x98").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßUpsilon.ToObject(), πg.NewUnicode("\xcf\x92").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßPi.ToObject(), πg.NewUnicode("\xce\xa0").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßOmega.ToObject(), πg.NewUnicode("\xce\xa9").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßGamma.ToObject(), πg.NewUnicode("\xce\x93").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßLambda.ToObject(), πg.NewUnicode("\xce\x9b").ToObject()); πE != nil {
				continue
			}
			πTemp001 = πTemp003.ToObject()
			if πE = πF.Globals().SetItem(πF, ßGreek.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 46: letters = tex2unichar.mathalpha
			πF.SetLineno(46)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßtex2unichar); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßmathalpha, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßletters.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 48: special = tex2unichar.mathbin         # Binary symbols
			πF.SetLineno(48)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßtex2unichar); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßmathbin, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßspecial.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 49: special.update(tex2unichar.mathrel)   # Relation symbols, arrow symbols
			πF.SetLineno(49)
			πTemp002 = πF.MakeArgs(1)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßtex2unichar); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßmathrel, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			if πTemp001, πE = πg.ResolveGlobal(πF, ßspecial); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßupdate, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 50: special.update(tex2unichar.mathord)   # Miscellaneous symbols
			πF.SetLineno(50)
			πTemp002 = πF.MakeArgs(1)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßtex2unichar); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßmathord, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			if πTemp001, πE = πg.ResolveGlobal(πF, ßspecial); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßupdate, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 51: special.update(tex2unichar.mathop)    # Variable-sized symbols
			πF.SetLineno(51)
			πTemp002 = πF.MakeArgs(1)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßtex2unichar); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßmathop, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			if πTemp001, πE = πg.ResolveGlobal(πF, ßspecial); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßupdate, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 52: special.update(tex2unichar.mathopen)  # Braces
			πF.SetLineno(52)
			πTemp002 = πF.MakeArgs(1)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßtex2unichar); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßmathopen, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			if πTemp001, πE = πg.ResolveGlobal(πF, ßspecial); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßupdate, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 53: special.update(tex2unichar.mathclose) # Braces
			πF.SetLineno(53)
			πTemp002 = πF.MakeArgs(1)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßtex2unichar); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßmathclose, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			if πTemp001, πE = πg.ResolveGlobal(πF, ßspecial); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßupdate, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 54: special.update(tex2unichar.mathfence)
			πF.SetLineno(54)
			πTemp002 = πF.MakeArgs(1)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßtex2unichar); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßmathfence, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			if πTemp001, πE = πg.ResolveGlobal(πF, ßspecial); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßupdate, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 56: sumintprod = ''.join([special[symbol] for symbol in
			πF.SetLineno(56)
			πTemp002 = πF.MakeArgs(1)
			πTemp005 = make([]πg.Param, 0)
			πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsymbol *πg.Object = πg.UnboundLocal
				_ = µsymbol
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 bool
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						case 1:
							goto Label1
						case 2:
							goto Label2
						case 4:
							goto Label4
						default:
							panic("unexpected function state")
						}
						πTemp002 = make([]*πg.Object, 4)
						πTemp002[0] = ßsum.ToObject()
						πTemp002[1] = ßint.ToObject()
						πTemp002[2] = ßoint.ToObject()
						πTemp002[3] = ßprod.ToObject()
						πTemp003 = πg.NewList(πTemp002...).ToObject()
						if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
							continue
						}
						πF.PushCheckpoint(2)
						πTemp004 = false
					Label1:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp004 {
							πF.PopCheckpoint()
							goto Label3
						}
						if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
							isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
							if exc != nil {
								πE = exc
							} else if isStop {
								πE = nil
								πF.RestoreExc(nil, nil)
							}
							πTemp005 = !isStop
						} else {
							πTemp005 = true
							µsymbol = πTemp003
						}
						if πE != nil || !πTemp005 {
							continue
						}
						πF.PushCheckpoint(1)
						// line 56: sumintprod = ''.join([special[symbol] for symbol in
						πF.SetLineno(56)
						if πE = πg.CheckLocal(πF, µsymbol, "symbol"); πE != nil {
							continue
						}
						πTemp003 = µsymbol
						if πTemp007, πE = πg.ResolveGlobal(πF, ßspecial); πE != nil {
							continue
						}
						if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp003); πE != nil {
							continue
						}
						πF.PushCheckpoint(4)
						return πTemp006, nil
					Label4:
						πTemp003 = πSent
						continue
					Label2:
						if πE != nil || πR != nil {
							continue
						}
					Label3:
					}
					return nil, πE
				}).ToObject(), nil
			}), πF.Globals()).ToObject()
			if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp006}, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp001
			if πTemp001, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
				continue
			}
			if πTemp006, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßsumintprod.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 59: functions = ['arccos', 'arcsin', 'arctan', 'arg', 'cos',  'cosh',
			πF.SetLineno(59)
			πTemp002 = make([]*πg.Object, 38)
			πTemp002[0] = ßarccos.ToObject()
			πTemp002[1] = ßarcsin.ToObject()
			πTemp002[2] = ßarctan.ToObject()
			πTemp002[3] = ßarg.ToObject()
			πTemp002[4] = ßcos.ToObject()
			πTemp002[5] = ßcosh.ToObject()
			πTemp002[6] = ßcot.ToObject()
			πTemp002[7] = ßcoth.ToObject()
			πTemp002[8] = ßcsc.ToObject()
			πTemp002[9] = ßdeg.ToObject()
			πTemp002[10] = ßdet.ToObject()
			πTemp002[11] = ßdim.ToObject()
			πTemp002[12] = ßexp.ToObject()
			πTemp002[13] = ßgcd.ToObject()
			πTemp002[14] = ßhom.ToObject()
			πTemp002[15] = ßinf.ToObject()
			πTemp002[16] = ßker.ToObject()
			πTemp002[17] = ßlg.ToObject()
			πTemp002[18] = ßlim.ToObject()
			πTemp002[19] = ßliminf.ToObject()
			πTemp002[20] = ßlimsup.ToObject()
			πTemp002[21] = ßln.ToObject()
			πTemp002[22] = ßlog.ToObject()
			πTemp002[23] = ßmax.ToObject()
			πTemp002[24] = ßmin.ToObject()
			πTemp002[25] = ßPr.ToObject()
			πTemp002[26] = ßsec.ToObject()
			πTemp002[27] = ßsin.ToObject()
			πTemp002[28] = ßsinh.ToObject()
			πTemp002[29] = ßsup.ToObject()
			πTemp002[30] = ßtan.ToObject()
			πTemp002[31] = ßtanh.ToObject()
			πTemp002[32] = ßinjlim.ToObject()
			πTemp002[33] = ßvarinjlim.ToObject()
			πTemp002[34] = ßvarlimsup.ToObject()
			πTemp002[35] = ßprojlim.ToObject()
			πTemp002[36] = ßvarliminf.ToObject()
			πTemp002[37] = ßvarprojlim.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßfunctions.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 69: mathbb = {
			πF.SetLineno(69)
			πTemp003 = πg.NewDict()
			if πE = πTemp003.SetItem(πF, ßA.ToObject(), πg.NewUnicode("\xf0\x9d\x94\xb8").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßB.ToObject(), πg.NewUnicode("\xf0\x9d\x94\xb9").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßC.ToObject(), πg.NewUnicode("\xe2\x84\x82").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßD.ToObject(), πg.NewUnicode("\xf0\x9d\x94\xbb").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßE.ToObject(), πg.NewUnicode("\xf0\x9d\x94\xbc").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßF.ToObject(), πg.NewUnicode("\xf0\x9d\x94\xbd").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßG.ToObject(), πg.NewUnicode("\xf0\x9d\x94\xbe").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßH.ToObject(), πg.NewUnicode("\xe2\x84\x8d").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßI.ToObject(), πg.NewUnicode("\xf0\x9d\x95\x80").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßJ.ToObject(), πg.NewUnicode("\xf0\x9d\x95\x81").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßK.ToObject(), πg.NewUnicode("\xf0\x9d\x95\x82").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßL.ToObject(), πg.NewUnicode("\xf0\x9d\x95\x83").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßM.ToObject(), πg.NewUnicode("\xf0\x9d\x95\x84").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßN.ToObject(), πg.NewUnicode("\xe2\x84\x95").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßO.ToObject(), πg.NewUnicode("\xf0\x9d\x95\x86").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßP.ToObject(), πg.NewUnicode("\xe2\x84\x99").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßQ.ToObject(), πg.NewUnicode("\xe2\x84\x9a").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßR.ToObject(), πg.NewUnicode("\xe2\x84\x9d").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßS.ToObject(), πg.NewUnicode("\xf0\x9d\x95\x8a").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßT.ToObject(), πg.NewUnicode("\xf0\x9d\x95\x8b").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßU.ToObject(), πg.NewUnicode("\xf0\x9d\x95\x8c").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßV.ToObject(), πg.NewUnicode("\xf0\x9d\x95\x8d").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßW.ToObject(), πg.NewUnicode("\xf0\x9d\x95\x8e").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßX.ToObject(), πg.NewUnicode("\xf0\x9d\x95\x8f").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßY.ToObject(), πg.NewUnicode("\xf0\x9d\x95\x90").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßZ.ToObject(), πg.NewUnicode("\xe2\x84\xa4").ToObject()); πE != nil {
				continue
			}
			πTemp001 = πTemp003.ToObject()
			if πE = πF.Globals().SetItem(πF, ßmathbb.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 98: mathscr = {
			πF.SetLineno(98)
			πTemp003 = πg.NewDict()
			if πE = πTemp003.SetItem(πF, ßA.ToObject(), πg.NewUnicode("\xf0\x9d\x92\x9c").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßB.ToObject(), πg.NewUnicode("\xe2\x84\xac").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßC.ToObject(), πg.NewUnicode("\xf0\x9d\x92\x9e").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßD.ToObject(), πg.NewUnicode("\xf0\x9d\x92\x9f").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßE.ToObject(), πg.NewUnicode("\xe2\x84\xb0").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßF.ToObject(), πg.NewUnicode("\xe2\x84\xb1").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßG.ToObject(), πg.NewUnicode("\xf0\x9d\x92\xa2").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßH.ToObject(), πg.NewUnicode("\xe2\x84\x8b").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßI.ToObject(), πg.NewUnicode("\xe2\x84\x90").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßJ.ToObject(), πg.NewUnicode("\xf0\x9d\x92\xa5").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßK.ToObject(), πg.NewUnicode("\xf0\x9d\x92\xa6").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßL.ToObject(), πg.NewUnicode("\xe2\x84\x92").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßM.ToObject(), πg.NewUnicode("\xe2\x84\xb3").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßN.ToObject(), πg.NewUnicode("\xf0\x9d\x92\xa9").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßO.ToObject(), πg.NewUnicode("\xf0\x9d\x92\xaa").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßP.ToObject(), πg.NewUnicode("\xf0\x9d\x92\xab").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßQ.ToObject(), πg.NewUnicode("\xf0\x9d\x92\xac").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßR.ToObject(), πg.NewUnicode("\xe2\x84\x9b").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßS.ToObject(), πg.NewUnicode("\xf0\x9d\x92\xae").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßT.ToObject(), πg.NewUnicode("\xf0\x9d\x92\xaf").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßU.ToObject(), πg.NewUnicode("\xf0\x9d\x92\xb0").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßV.ToObject(), πg.NewUnicode("\xf0\x9d\x92\xb1").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßW.ToObject(), πg.NewUnicode("\xf0\x9d\x92\xb2").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßX.ToObject(), πg.NewUnicode("\xf0\x9d\x92\xb3").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßY.ToObject(), πg.NewUnicode("\xf0\x9d\x92\xb4").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßZ.ToObject(), πg.NewUnicode("\xf0\x9d\x92\xb5").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßa.ToObject(), πg.NewUnicode("\xf0\x9d\x92\xb6").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßb.ToObject(), πg.NewUnicode("\xf0\x9d\x92\xb7").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßc.ToObject(), πg.NewUnicode("\xf0\x9d\x92\xb8").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßd.ToObject(), πg.NewUnicode("\xf0\x9d\x92\xb9").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ße.ToObject(), πg.NewUnicode("\xe2\x84\xaf").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßf.ToObject(), πg.NewUnicode("\xf0\x9d\x92\xbb").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßg.ToObject(), πg.NewUnicode("\xe2\x84\x8a").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßh.ToObject(), πg.NewUnicode("\xf0\x9d\x92\xbd").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßi.ToObject(), πg.NewUnicode("\xf0\x9d\x92\xbe").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßj.ToObject(), πg.NewUnicode("\xf0\x9d\x92\xbf").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßk.ToObject(), πg.NewUnicode("\xf0\x9d\x93\x80").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßl.ToObject(), πg.NewUnicode("\xf0\x9d\x93\x81").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßm.ToObject(), πg.NewUnicode("\xf0\x9d\x93\x82").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßn.ToObject(), πg.NewUnicode("\xf0\x9d\x93\x83").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßo.ToObject(), πg.NewUnicode("\xe2\x84\xb4").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßp.ToObject(), πg.NewUnicode("\xf0\x9d\x93\x85").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßq.ToObject(), πg.NewUnicode("\xf0\x9d\x93\x86").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßr.ToObject(), πg.NewUnicode("\xf0\x9d\x93\x87").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßs.ToObject(), πg.NewUnicode("\xf0\x9d\x93\x88").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßt.ToObject(), πg.NewUnicode("\xf0\x9d\x93\x89").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßu.ToObject(), πg.NewUnicode("\xf0\x9d\x93\x8a").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßv.ToObject(), πg.NewUnicode("\xf0\x9d\x93\x8b").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßw.ToObject(), πg.NewUnicode("\xf0\x9d\x93\x8c").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßx.ToObject(), πg.NewUnicode("\xf0\x9d\x93\x8d").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßy.ToObject(), πg.NewUnicode("\xf0\x9d\x93\x8e").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ßz.ToObject(), πg.NewUnicode("\xf0\x9d\x93\x8f").ToObject()); πE != nil {
				continue
			}
			πTemp001 = πTemp003.ToObject()
			if πE = πF.Globals().SetItem(πF, ßmathscr.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 153: negatables = {'=': u'\u2260',
			πF.SetLineno(153)
			πTemp003 = πg.NewDict()
			if πE = πTemp003.SetItem(πF, πg.NewStr("=").ToObject(), πg.NewUnicode("\xe2\x89\xa0").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, πg.NewStr("\\in").ToObject(), πg.NewUnicode("\xe2\x88\x89").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, πg.NewStr("\\equiv").ToObject(), πg.NewUnicode("\xe2\x89\xa2").ToObject()); πE != nil {
				continue
			}
			πTemp001 = πTemp003.ToObject()
			if πE = πF.Globals().SetItem(πF, ßnegatables.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 158: class math(object):
			πF.SetLineno(158)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("math", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 159: """Base class for MathML elements."""
					πF.SetLineno(159)
					// line 159: """Base class for MathML elements."""
					πF.SetLineno(159)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Base class for MathML elements.").ToObject()); πE != nil {
						continue
					}
					// line 161: nchildren = 1000000
					πF.SetLineno(161)
					if πE = πClass.SetItem(πF, ßnchildren.ToObject(), πg.NewInt(1000000).ToObject()); πE != nil {
						continue
					}
					// line 162: """Required number of children"""
					πF.SetLineno(162)
					// line 164: def __init__(self, children=None, inline=None):
					πF.SetLineno(164)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "children", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "inline", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µchildren *πg.Object = πArgs[1]
						_ = µchildren
						var µinline *πg.Object = πArgs[2]
						_ = µinline
						var µchild *πg.Object = πg.UnboundLocal
						_ = µchild
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 6:
								goto Label6
							case 7:
								goto Label7
							default:
								panic("unexpected function state")
							}
							// line 165: """math([children]) -> MathML element
							πF.SetLineno(165)
							// line 169: self.children = []
							πF.SetLineno(169)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßchildren, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µchildren != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 170: if children is not None:
							πF.SetLineno(170)
						Label1:
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							πTemp001[0] = µchildren
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 171: if isinstance(children, list):
							πF.SetLineno(171)
						Label3:
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µchildren); πE != nil {
								continue
							}
							πF.PushCheckpoint(7)
							πTemp004 = false
						Label6:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label8
							}
							if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp005 = !isStop
							} else {
								πTemp005 = true
								µchild = πTemp003
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(6)
							// line 173: self.append(child)
							πF.SetLineno(173)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
								continue
							}
							πTemp001[0] = µchild
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
							goto Label5
						Label4:
							// line 176: self.append(children)
							πF.SetLineno(176)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							πTemp001[0] = µchildren
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label5
						Label5:
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µinline, "inline"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µinline != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label9
							}
							goto Label10
							// line 178: if inline is not None:
							πF.SetLineno(178)
						Label9:
							// line 179: self.inline = inline
							πF.SetLineno(179)
							if πE = πg.CheckLocal(πF, µinline, "inline"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µinline); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßinline, πTemp002); πE != nil {
								continue
							}
							goto Label10
						Label10:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 165: """math([children]) -> MathML element
					πF.SetLineno(165)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("math([children]) -> MathML element\n\n        children can be one child or a list of children.").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__init__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 181: def __repr__(self):
					πF.SetLineno(181)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__repr__", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 []πg.Param
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							πTemp001[1] = ßchildren.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 182: if hasattr(self, 'children'):
							πF.SetLineno(182)
						Label1:
							// line 183: return self.__class__.__name__ + '(%s)' % \
							πF.SetLineno(183)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ß__name__, nil); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							πTemp008 = make([]πg.Param, 0)
							πTemp007 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µchild *πg.Object = πg.UnboundLocal
								_ = µchild
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 *πg.Object
								_ = πTemp002
								var πTemp003 bool
								_ = πTemp003
								var πTemp004 bool
								_ = πTemp004
								var πTemp005 []*πg.Object
								_ = πTemp005
								var πTemp006 *πg.Object
								_ = πTemp006
								var πR *πg.Object
								_ = πR
								var πE *πg.BaseException
								_ = πE
								return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
									for ; πF.State() >= 0; πF.PopCheckpoint() {
										switch πF.State() {
										case 0:
										case 1:
											goto Label1
										case 2:
											goto Label2
										case 4:
											goto Label4
										default:
											panic("unexpected function state")
										}
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
											continue
										}
										πF.PushCheckpoint(2)
										πTemp003 = false
									Label1:
										if πE != nil || πR != nil {
											continue
										}
										if πTemp003 {
											πF.PopCheckpoint()
											goto Label3
										}
										if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
											isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
											if exc != nil {
												πE = exc
											} else if isStop {
												πE = nil
												πF.RestoreExc(nil, nil)
											}
											πTemp004 = !isStop
										} else {
											πTemp004 = true
											µchild = πTemp002
										}
										if πE != nil || !πTemp004 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 184: ','.join([repr(child) for child in self.children])
										πF.SetLineno(184)
										πTemp005 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
											continue
										}
										πTemp005[0] = µchild
										if πTemp002, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
											continue
										}
										if πTemp006, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp005)
										πF.PushCheckpoint(4)
										return πTemp006, nil
									Label4:
										πTemp002 = πSent
										continue
									Label2:
										if πE != nil || πR != nil {
											continue
										}
									Label3:
									}
									return nil, πE
								}).ToObject(), nil
							}), πF.Globals()).ToObject()
							if πTemp009, πE = πTemp007.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ListType.Call(πF, πg.Args{πTemp009}, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp006
							if πTemp006, πE = πg.GetAttr(πF, πg.NewStr(",").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("(%s)").ToObject(), πTemp009); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp005, πTemp003); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label3
						Label2:
							// line 186: return self.__class__.__name__
							πF.SetLineno(186)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__name__, nil); πE != nil {
								continue
							}
							πR = πTemp003
							continue
							goto Label3
						Label3:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 188: def full(self):
					πF.SetLineno(188)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("full", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
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
							// line 189: """Room for more children?"""
							πF.SetLineno(189)
							// line 191: return len(self.children) >= self.nchildren
							πF.SetLineno(191)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßnchildren, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GE(πF, πTemp004, πTemp003); πE != nil {
								continue
							}
							πR = πTemp001
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßfull.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 189: """Room for more children?"""
					πF.SetLineno(189)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Room for more children?").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßfull); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 193: def append(self, child):
					πF.SetLineno(193)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "child", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("append", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µchild *πg.Object = πArgs[1]
						_ = µchild
						var µnode *πg.Object = πg.UnboundLocal
						_ = µnode
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
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
							default:
								panic("unexpected function state")
							}
							// line 194: """append(child) -> element
							πF.SetLineno(194)
							// line 199: assert not self.full()
							πF.SetLineno(199)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfull, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
								continue
							}
							// line 200: self.children.append(child)
							πF.SetLineno(200)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
								continue
							}
							πTemp005[0] = µchild
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 201: child.parent = self
							πF.SetLineno(201)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µself); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µchild, ßparent, πTemp001); πE != nil {
								continue
							}
							// line 202: node = self
							πF.SetLineno(202)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							µnode = µself
							// line 203: while node.full():
							πF.SetLineno(203)
							πF.PushCheckpoint(2)
							πTemp004 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µnode, ßfull, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 204: node = node.parent
							πF.SetLineno(204)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µnode, ßparent, nil); πE != nil {
								continue
							}
							µnode = πTemp001
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 205: return node
							πF.SetLineno(205)
							if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
								continue
							}
							πR = µnode
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßappend.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 194: """append(child) -> element
					πF.SetLineno(194)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("append(child) -> element\n\n        Appends child and returns self if self is not full or first\n        non-full parent.").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßappend); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 207: def delete_child(self):
					πF.SetLineno(207)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("delete_child", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µchild *πg.Object = πg.UnboundLocal
						_ = µchild
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
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
							// line 208: """delete_child() -> child
							πF.SetLineno(208)
							// line 212: child = self.children[-1]
							πF.SetLineno(212)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							µchild = πTemp002
							// line 213: del self.children[-1]
							πF.SetLineno(213)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πE = πg.DelItem(πF, πTemp001, πTemp002); πE != nil {
								continue
							}
							// line 214: return child
							πF.SetLineno(214)
							if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
								continue
							}
							πR = µchild
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdelete_child.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 208: """delete_child() -> child
					πF.SetLineno(208)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("delete_child() -> child\n\n        Delete last child and return it.").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßdelete_child); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 216: def close(self):
					πF.SetLineno(216)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("close", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µparent *πg.Object = πg.UnboundLocal
						_ = µparent
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
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
							default:
								panic("unexpected function state")
							}
							// line 217: """close() -> parent
							πF.SetLineno(217)
							// line 221: parent = self.parent
							πF.SetLineno(221)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßparent, nil); πE != nil {
								continue
							}
							µparent = πTemp001
							// line 222: while parent.full():
							πF.SetLineno(222)
							πF.PushCheckpoint(2)
							πTemp002 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp002 {
								πF.PopCheckpoint()
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µparent, ßfull, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 223: parent = parent.parent
							πF.SetLineno(223)
							if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µparent, ßparent, nil); πE != nil {
								continue
							}
							µparent = πTemp001
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 224: return parent
							πF.SetLineno(224)
							if πE = πg.CheckLocal(πF, µparent, "parent"); πE != nil {
								continue
							}
							πR = µparent
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßclose.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 217: """close() -> parent
					πF.SetLineno(217)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("close() -> parent\n\n        Close element and return first non-full element.").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßclose); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 226: def xml(self):
					πF.SetLineno(226)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("xml", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
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
							// line 227: """xml() -> xml-string"""
							πF.SetLineno(227)
							// line 229: return self.xml_start() + self.xml_body() + self.xml_end()
							πF.SetLineno(229)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßxml_start, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßxml_body, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßxml_end, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, πTemp004); πE != nil {
								continue
							}
							πR = πTemp001
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßxml.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 227: """xml() -> xml-string"""
					πF.SetLineno(227)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("xml() -> xml-string").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßxml); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
						continue
					}
					// line 231: def xml_start(self):
					πF.SetLineno(231)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("xml_start", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µxmlns *πg.Object = πg.UnboundLocal
						_ = µxmlns
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
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
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							πTemp002[1] = ßinline.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 232: if not hasattr(self, 'inline'):
							πF.SetLineno(232)
						Label1:
							// line 233: return ['<%s>' % self.__class__.__name__]
							πF.SetLineno(233)
							πTemp002 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß__name__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("<%s>").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							πTemp001 = πg.NewList(πTemp002...).ToObject()
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 234: xmlns = 'http://www.w3.org/1998/Math/MathML'
							πF.SetLineno(234)
							µxmlns = πg.NewStr("http://www.w3.org/1998/Math/MathML").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßinline, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							goto Label4
							// line 235: if self.inline:
							πF.SetLineno(235)
						Label3:
							// line 236: return ['<math xmlns="%s">' % xmlns]
							πF.SetLineno(236)
							πTemp002 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µxmlns, "xmlns"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("<math xmlns=\"%s\">").ToObject(), µxmlns); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							πTemp001 = πg.NewList(πTemp002...).ToObject()
							πR = πTemp001
							continue
							goto Label5
						Label4:
							// line 238: return ['<math xmlns="%s" mode="display">' % xmlns]
							πF.SetLineno(238)
							πTemp002 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µxmlns, "xmlns"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("<math xmlns=\"%s\" mode=\"display\">").ToObject(), µxmlns); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							πTemp001 = πg.NewList(πTemp002...).ToObject()
							πR = πTemp001
							continue
							goto Label5
						Label5:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßxml_start.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 240: def xml_end(self):
					πF.SetLineno(240)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("xml_end", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
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
							// line 241: return ['</%s>' % self.__class__.__name__]
							πF.SetLineno(241)
							πTemp001 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß__name__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("</%s>").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							πR = πTemp002
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßxml_end.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 243: def xml_body(self):
					πF.SetLineno(243)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("xml_body", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µxml *πg.Object = πg.UnboundLocal
						_ = µxml
						var µchild *πg.Object = πg.UnboundLocal
						_ = µchild
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
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
							default:
								panic("unexpected function state")
							}
							// line 244: xml = []
							πF.SetLineno(244)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µxml = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp004 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
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
								πTemp005 = !isStop
							} else {
								πTemp005 = true
								µchild = πTemp003
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 246: xml.extend(child.xml())
							πF.SetLineno(246)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µchild, ßxml, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp006
							if πE = πg.CheckLocal(πF, µxml, "xml"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µxml, ßextend, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 247: return xml
							πF.SetLineno(247)
							if πE = πg.CheckLocal(πF, µxml, "xml"); πE != nil {
								continue
							}
							πR = µxml
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßxml_body.ToObject(), πTemp011); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("math").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmath.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 249: class mrow(math):
			πF.SetLineno(249)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("mrow", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 250: def xml_start(self):
					πF.SetLineno(250)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("xml_start", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
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
							// line 251: return ['\n<%s>' % self.__class__.__name__]
							πF.SetLineno(251)
							πTemp001 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß__name__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("\n<%s>").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							πR = πTemp002
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßxml_start.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("mrow").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmrow.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 253: class mtable(math):
			πF.SetLineno(253)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("mtable", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 254: def xml_start(self):
					πF.SetLineno(254)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("xml_start", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
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
							// line 255: return ['\n<%s>' % self.__class__.__name__]
							πF.SetLineno(255)
							πTemp001 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß__name__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("\n<%s>").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							πR = πTemp002
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßxml_start.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("mtable").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmtable.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 257: class mtr(mrow): pass
			πF.SetLineno(257)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßmrow); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("mtr", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 257: class mtr(mrow): pass
					πF.SetLineno(257)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("mtr").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmtr.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 258: class mtd(mrow): pass
			πF.SetLineno(258)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßmrow); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("mtd", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 258: class mtd(mrow): pass
					πF.SetLineno(258)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("mtd").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmtd.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 260: class mx(math):
			πF.SetLineno(260)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("mx", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 261: """Base class for mo, mi, and mn"""
					πF.SetLineno(261)
					// line 261: """Base class for mo, mi, and mn"""
					πF.SetLineno(261)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Base class for mo, mi, and mn").ToObject()); πE != nil {
						continue
					}
					// line 263: nchildren = 0
					πF.SetLineno(263)
					if πE = πClass.SetItem(πF, ßnchildren.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					// line 264: def __init__(self, data):
					πF.SetLineno(264)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "data", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdata *πg.Object = πArgs[1]
						_ = µdata
						var πTemp001 *πg.Object
						_ = πTemp001
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
							// line 265: self.data = data
							πF.SetLineno(265)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µdata); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdata, πTemp001); πE != nil {
								continue
							}
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 267: def xml_body(self):
					πF.SetLineno(267)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("xml_body", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
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
							// line 268: return [self.data]
							πF.SetLineno(268)
							πTemp001 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							πR = πTemp002
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßxml_body.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("mx").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmx.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 270: class mo(mx):
			πF.SetLineno(270)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßmx); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("mo", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				var πTemp001 *πg.Dict
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []πg.Param
				_ = πTemp003
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 271: translation = {'<': '&lt;', '>': '&gt;'}
					πF.SetLineno(271)
					πTemp001 = πg.NewDict()
					if πE = πTemp001.SetItem(πF, πg.NewStr("<").ToObject(), πg.NewStr("&lt;").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr(">").ToObject(), πg.NewStr("&gt;").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp001.ToObject()
					if πE = πClass.SetItem(πF, ßtranslation.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 272: def xml_body(self):
					πF.SetLineno(272)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("xml_body", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
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
							// line 273: return [self.translation.get(self.data, self.data)]
							πF.SetLineno(273)
							πTemp001 = make([]*πg.Object, 1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtranslation, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßget, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp003
							πTemp003 = πg.NewList(πTemp001...).ToObject()
							πR = πTemp003
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßxml_body.ToObject(), πTemp002); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("mo").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmo.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 275: class mi(mx): pass
			πF.SetLineno(275)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßmx); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("mi", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 275: class mi(mx): pass
					πF.SetLineno(275)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("mi").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmi.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 276: class mn(mx): pass
			πF.SetLineno(276)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßmx); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("mn", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 276: class mn(mx): pass
					πF.SetLineno(276)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("mn").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmn.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 278: class msub(math):
			πF.SetLineno(278)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("msub", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 279: nchildren = 2
					πF.SetLineno(279)
					if πE = πClass.SetItem(πF, ßnchildren.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("msub").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmsub.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 281: class msup(math):
			πF.SetLineno(281)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("msup", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 282: nchildren = 2
					πF.SetLineno(282)
					if πE = πClass.SetItem(πF, ßnchildren.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("msup").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmsup.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 284: class msqrt(math):
			πF.SetLineno(284)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("msqrt", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 285: nchildren = 1
					πF.SetLineno(285)
					if πE = πClass.SetItem(πF, ßnchildren.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("msqrt").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmsqrt.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 287: class mroot(math):
			πF.SetLineno(287)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("mroot", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 288: nchildren = 2
					πF.SetLineno(288)
					if πE = πClass.SetItem(πF, ßnchildren.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("mroot").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmroot.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 290: class mfrac(math):
			πF.SetLineno(290)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("mfrac", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 291: nchildren = 2
					πF.SetLineno(291)
					if πE = πClass.SetItem(πF, ßnchildren.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("mfrac").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmfrac.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 293: class msubsup(math):
			πF.SetLineno(293)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("msubsup", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 294: nchildren = 3
					πF.SetLineno(294)
					if πE = πClass.SetItem(πF, ßnchildren.ToObject(), πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					// line 295: def __init__(self, children=None, reversed=False):
					πF.SetLineno(295)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "children", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "reversed", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µchildren *πg.Object = πArgs[1]
						_ = µchildren
						var µreversed *πg.Object = πArgs[2]
						_ = µreversed
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
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
							// line 296: self.reversed = reversed
							πF.SetLineno(296)
							if πE = πg.CheckLocal(πF, µreversed, "reversed"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µreversed); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßreversed, πTemp001); πE != nil {
								continue
							}
							// line 297: math.__init__(self, children)
							πF.SetLineno(297)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							πTemp002[1] = µchildren
							if πTemp001, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 299: def xml(self):
					πF.SetLineno(299)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("xml", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreversed, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 300: if self.reversed:
							πF.SetLineno(300)
						Label1:
							// line 302: self.children[1:3] = [self.children[2], self.children[1]]
							πF.SetLineno(302)
							πTemp003 = make([]*πg.Object, 2)
							πTemp001 = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							πTemp001 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.NewInt(3).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, πTemp005, πTemp006, πTemp004); πE != nil {
								continue
							}
							// line 303: self.reversed = False
							πF.SetLineno(303)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßreversed, πTemp004); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 304: return math.xml(self)
							πF.SetLineno(304)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßxml, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πR = πTemp001
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßxml.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("msubsup").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmsubsup.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 306: class mfenced(math):
			πF.SetLineno(306)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("mfenced", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				var πTemp001 *πg.Dict
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []πg.Param
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 307: translation = {'\\{': '{', '\\langle': u'\u2329',
					πF.SetLineno(307)
					πTemp001 = πg.NewDict()
					if πE = πTemp001.SetItem(πF, πg.NewStr("\\{").ToObject(), πg.NewStr("{").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("\\langle").ToObject(), πg.NewUnicode("\xe2\x8c\xa9").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("\\}").ToObject(), πg.NewStr("}").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("\\rangle").ToObject(), πg.NewUnicode("\xe2\x8c\xaa").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr(".").ToObject(), ß.ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp001.ToObject()
					if πE = πClass.SetItem(πF, ßtranslation.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 310: def __init__(self, par):
					πF.SetLineno(310)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "par", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µpar *πg.Object = πArgs[1]
						_ = µpar
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
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
							// line 311: self.openpar = par
							πF.SetLineno(311)
							if πE = πg.CheckLocal(πF, µpar, "par"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µpar); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßopenpar, πTemp001); πE != nil {
								continue
							}
							// line 312: math.__init__(self)
							πF.SetLineno(312)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 314: def xml_start(self):
					πF.SetLineno(314)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("xml_start", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µopen *πg.Object = πg.UnboundLocal
						_ = µopen
						var µclose *πg.Object = πg.UnboundLocal
						_ = µclose
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
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
							// line 315: open = self.translation.get(self.openpar, self.openpar)
							πF.SetLineno(315)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßopenpar, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßopenpar, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtranslation, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µopen = πTemp002
							// line 316: close = self.translation.get(self.closepar, self.closepar)
							πF.SetLineno(316)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßclosepar, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßclosepar, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtranslation, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µclose = πTemp002
							// line 317: return ['<mfenced open="%s" close="%s">' % (open, close)]
							πF.SetLineno(317)
							πTemp001 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µopen, "open"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µclose, "close"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(µopen, µclose).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("<mfenced open=\"%s\" close=\"%s\">").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							πR = πTemp002
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßxml_start.ToObject(), πTemp004); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("mfenced").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmfenced.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 319: class mspace(math):
			πF.SetLineno(319)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("mspace", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 320: nchildren = 0
					πF.SetLineno(320)
					if πE = πClass.SetItem(πF, ßnchildren.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("mspace").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmspace.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 322: class mstyle(math):
			πF.SetLineno(322)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("mstyle", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 323: def __init__(self, children=None, nchildren=None, **kwargs):
					πF.SetLineno(323)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "children", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "nchildren", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp002, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µchildren *πg.Object = πArgs[1]
						_ = µchildren
						var µnchildren *πg.Object = πArgs[2]
						_ = µnchildren
						var µkwargs *πg.Object = πArgs[3]
						_ = µkwargs
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
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
							if πE = πg.CheckLocal(πF, µnchildren, "nchildren"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µnchildren != πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 324: if nchildren is not None:
							πF.SetLineno(324)
						Label1:
							// line 325: self.nchildren = nchildren
							πF.SetLineno(325)
							if πE = πg.CheckLocal(πF, µnchildren, "nchildren"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µnchildren); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßnchildren, πTemp001); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 326: math.__init__(self, children)
							πF.SetLineno(326)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							πTemp004[1] = µchildren
							if πTemp001, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 327: self.attrs = kwargs
							πF.SetLineno(327)
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µkwargs); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßattrs, πTemp001); πE != nil {
								continue
							}
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 329: def xml_start(self):
					πF.SetLineno(329)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("xml_start", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 []πg.Param
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
							// line 330: return ['<mstyle '] + ['%s="%s"' % item
							πF.SetLineno(330)
							πTemp003 = make([]*πg.Object, 1)
							πTemp003[0] = πg.NewStr("<mstyle ").ToObject()
							πTemp004 = πg.NewList(πTemp003...).ToObject()
							πTemp007 = make([]πg.Param, 0)
							πTemp006 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µitem *πg.Object = πg.UnboundLocal
								_ = µitem
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 *πg.Object
								_ = πTemp002
								var πTemp003 *πg.Object
								_ = πTemp003
								var πTemp004 bool
								_ = πTemp004
								var πTemp005 bool
								_ = πTemp005
								var πR *πg.Object
								_ = πR
								var πE *πg.BaseException
								_ = πE
								return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
									for ; πF.State() >= 0; πF.PopCheckpoint() {
										switch πF.State() {
										case 0:
										case 1:
											goto Label1
										case 2:
											goto Label2
										case 4:
											goto Label4
										default:
											panic("unexpected function state")
										}
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, µself, ßattrs, nil); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßitems, nil); πE != nil {
											continue
										}
										if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
											continue
										}
										πF.PushCheckpoint(2)
										πTemp004 = false
									Label1:
										if πE != nil || πR != nil {
											continue
										}
										if πTemp004 {
											πF.PopCheckpoint()
											goto Label3
										}
										if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
											isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
											if exc != nil {
												πE = exc
											} else if isStop {
												πE = nil
												πF.RestoreExc(nil, nil)
											}
											πTemp005 = !isStop
										} else {
											πTemp005 = true
											µitem = πTemp002
										}
										if πE != nil || !πTemp005 {
											continue
										}
										πF.PushCheckpoint(1)
										// line 330: return ['<mstyle '] + ['%s="%s"' % item
										πF.SetLineno(330)
										if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s=\"%s\"").ToObject(), µitem); πE != nil {
											continue
										}
										πF.PushCheckpoint(4)
										return πTemp002, nil
									Label4:
										πTemp003 = πSent
										continue
									Label2:
										if πE != nil || πR != nil {
											continue
										}
									Label3:
									}
									return nil, πE
								}).ToObject(), nil
							}), πF.Globals()).ToObject()
							if πTemp008, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ListType.Call(πF, πg.Args{πTemp008}, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πTemp003 = make([]*πg.Object, 1)
							πTemp003[0] = πg.NewStr(">").ToObject()
							πTemp004 = πg.NewList(πTemp003...).ToObject()
							if πTemp001, πE = πg.Add(πF, πTemp002, πTemp004); πE != nil {
								continue
							}
							πR = πTemp001
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßxml_start.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("mstyle").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmstyle.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 333: class mover(math):
			πF.SetLineno(333)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("mover", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 334: nchildren = 2
					πF.SetLineno(334)
					if πE = πClass.SetItem(πF, ßnchildren.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					// line 335: def __init__(self, children=None, reversed=False):
					πF.SetLineno(335)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "children", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "reversed", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µchildren *πg.Object = πArgs[1]
						_ = µchildren
						var µreversed *πg.Object = πArgs[2]
						_ = µreversed
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
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
							// line 336: self.reversed = reversed
							πF.SetLineno(336)
							if πE = πg.CheckLocal(πF, µreversed, "reversed"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µreversed); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßreversed, πTemp001); πE != nil {
								continue
							}
							// line 337: math.__init__(self, children)
							πF.SetLineno(337)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							πTemp002[1] = µchildren
							if πTemp001, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 339: def xml(self):
					πF.SetLineno(339)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("xml", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreversed, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 340: if self.reversed:
							πF.SetLineno(340)
						Label1:
							// line 341: self.children.reverse()
							πF.SetLineno(341)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßchildren, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßreverse, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 342: self.reversed = False
							πF.SetLineno(342)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßreversed, πTemp003); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 343: return math.xml(self)
							πF.SetLineno(343)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßxml, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πR = πTemp001
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßxml.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("mover").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmover.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 345: class munder(math):
			πF.SetLineno(345)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("munder", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 346: nchildren = 2
					πF.SetLineno(346)
					if πE = πClass.SetItem(πF, ßnchildren.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("munder").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmunder.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 348: class munderover(math):
			πF.SetLineno(348)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("munderover", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 349: nchildren = 3
					πF.SetLineno(349)
					if πE = πClass.SetItem(πF, ßnchildren.ToObject(), πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					// line 350: def __init__(self, children=None):
					πF.SetLineno(350)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "children", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µchildren *πg.Object = πArgs[1]
						_ = µchildren
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
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
							// line 351: math.__init__(self, children)
							πF.SetLineno(351)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µchildren, "children"); πE != nil {
								continue
							}
							πTemp001[1] = µchildren
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("munderover").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmunderover.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 353: class mtext(math):
			πF.SetLineno(353)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("mtext", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 354: nchildren = 0
					πF.SetLineno(354)
					if πE = πClass.SetItem(πF, ßnchildren.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					// line 355: def __init__(self, text):
					πF.SetLineno(355)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "text", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtext *πg.Object = πArgs[1]
						_ = µtext
						var πTemp001 *πg.Object
						_ = πTemp001
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
							// line 356: self.text = text
							πF.SetLineno(356)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µtext); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtext, πTemp001); πE != nil {
								continue
							}
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 358: def xml_body(self):
					πF.SetLineno(358)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("xml_body", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
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
							// line 359: return [self.text]
							πF.SetLineno(359)
							πTemp001 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtext, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							πR = πTemp002
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßxml_body.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("mtext").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmtext.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 361: def parse_latex_math(string, inline=True):
			πF.SetLineno(361)
			πTemp005 = make([]πg.Param, 2)
			πTemp005[0] = πg.Param{Name: "string", Def: nil}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp005[1] = πg.Param{Name: "inline", Def: πTemp006}
			πTemp001 = πg.NewFunction(πg.NewCode("parse_latex_math", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µstring *πg.Object = πArgs[0]
				_ = µstring
				var µinline *πg.Object = πArgs[1]
				_ = µinline
				var µnode *πg.Object = πg.UnboundLocal
				_ = µnode
				var µtree *πg.Object = πg.UnboundLocal
				_ = µtree
				var µn *πg.Object = πg.UnboundLocal
				_ = µn
				var µc *πg.Object = πg.UnboundLocal
				_ = µc
				var µskip *πg.Object = πg.UnboundLocal
				_ = µskip
				var µc2 *πg.Object = πg.UnboundLocal
				_ = µc2
				var µi *πg.Object = πg.UnboundLocal
				_ = µi
				var µname *πg.Object = πg.UnboundLocal
				_ = µname
				var µentry *πg.Object = πg.UnboundLocal
				_ = µentry
				var µrow *πg.Object = πg.UnboundLocal
				_ = µrow
				var µchild *πg.Object = πg.UnboundLocal
				_ = µchild
				var µsub *πg.Object = πg.UnboundLocal
				_ = µsub
				var µsup *πg.Object = πg.UnboundLocal
				_ = µsup
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 πg.KWArgs
				_ = πTemp005
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 bool
				_ = πTemp010
				var πTemp011 bool
				_ = πTemp011
				var πTemp012 *πg.Object
				_ = πTemp012
				var πTemp013 *πg.Object
				_ = πTemp013
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 4:
						goto Label4
					case 5:
						goto Label5
					case 30:
						goto Label30
					case 29:
						goto Label29
					default:
						panic("unexpected function state")
					}
					// line 362: """parse_latex_math(string [,inline]) -> MathML-tree
					πF.SetLineno(362)
					// line 370: string = ' '.join(string.split())
					πF.SetLineno(370)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µstring, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.GetAttr(πF, πg.NewStr(" ").ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µstring = πTemp003
					if πE = πg.CheckLocal(πF, µinline, "inline"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µinline); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 372: if inline:
					πF.SetLineno(372)
				Label1:
					// line 373: node = mrow()
					πF.SetLineno(373)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmrow); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µnode = πTemp003
					// line 374: tree = math(node, inline=True)
					πF.SetLineno(374)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					πTemp001[0] = µnode
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp005 = πg.KWArgs{
						{"inline", πTemp002},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp005); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtree = πTemp003
					goto Label3
				Label2:
					// line 376: node = mtd()
					πF.SetLineno(376)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmtd); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µnode = πTemp003
					// line 377: tree = math(mtable(mtr(node)), inline=False)
					πF.SetLineno(377)
					πTemp001 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(1)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					πTemp007[0] = µnode
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmtr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp006[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmtable); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πTemp005 = πg.KWArgs{
						{"inline", πTemp002},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmath); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp005); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtree = πTemp003
					goto Label3
				Label3:
					// line 379: while len(string) > 0:
					πF.SetLineno(379)
					πF.PushCheckpoint(5)
					πTemp004 = false
				Label4:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label6
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					πTemp001[0] = µstring
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.GT(πF, πTemp009, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(4)
					// line 380: n = len(string)
					πF.SetLineno(380)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					πTemp001[0] = µstring
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µn = πTemp003
					// line 381: c = string[0]
					πF.SetLineno(381)
					πTemp002 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µstring, πTemp002); πE != nil {
						continue
					}
					µc = πTemp003
					// line 382: skip = 1  # number of characters consumed
					πF.SetLineno(382)
					µskip = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GT(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label7
					}
					goto Label8
					// line 383: if n > 1:
					πF.SetLineno(383)
				Label7:
					// line 384: c2 = string[1]
					πF.SetLineno(384)
					πTemp002 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µstring, πTemp002); πE != nil {
						continue
					}
					µc2 = πTemp003
					goto Label9
				Label8:
					// line 386: c2 = ''
					πF.SetLineno(386)
					µc2 = ß.ToObject()
					goto Label9
				Label9:
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µc, πg.NewStr(" ").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label10
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µc, πg.NewStr("\\").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label11
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µc, ßisalpha, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label12
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µc, ßisdigit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label13
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Contains(πF, πg.NewStr("+-*/=()[]|<>,.!?':;@").ToObject(), µc); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp008).ToObject()
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label14
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µc, ß_.ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label15
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µc, πg.NewStr("^").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label16
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µc, πg.NewStr("{").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label17
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µc, πg.NewStr("}").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label18
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µc, πg.NewStr("&").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label19
					}
					goto Label20
					// line 387: if c == ' ':
					πF.SetLineno(387)
				Label10:
					// line 388: pass
					πF.SetLineno(388)
					goto Label21
					// line 389: elif c == '\\':
					πF.SetLineno(389)
				Label11:
					if πE = πg.CheckLocal(πF, µc2, "c2"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Contains(πF, πg.NewStr("{}").ToObject(), µc2); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp008).ToObject()
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label22
					}
					if πE = πg.CheckLocal(πF, µc2, "c2"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µc2, πg.NewStr(" ").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label23
					}
					if πE = πg.CheckLocal(πF, µc2, "c2"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µc2, πg.NewStr(",").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label24
					}
					if πE = πg.CheckLocal(πF, µc2, "c2"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µc2, ßisalpha, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label25
					}
					if πE = πg.CheckLocal(πF, µc2, "c2"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µc2, πg.NewStr("\\").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label26
					}
					goto Label27
					// line 390: if c2 in '{}':
					πF.SetLineno(390)
				Label22:
					// line 391: node = node.append(mo(c2))
					πF.SetLineno(391)
					πTemp001 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µc2, "c2"); πE != nil {
						continue
					}
					πTemp006[0] = µc2
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmo); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µnode, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µnode = πTemp003
					// line 392: skip = 2
					πF.SetLineno(392)
					µskip = πg.NewInt(2).ToObject()
					goto Label28
					// line 393: elif c2 == ' ':
					πF.SetLineno(393)
				Label23:
					// line 394: node = node.append(mspace())
					πF.SetLineno(394)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmspace); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µnode, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µnode = πTemp003
					// line 395: skip = 2
					πF.SetLineno(395)
					µskip = πg.NewInt(2).ToObject()
					goto Label28
					// line 396: elif c2 == ',': # TODO: small space
					πF.SetLineno(396)
				Label24:
					// line 397: node = node.append(mspace())
					πF.SetLineno(397)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmspace); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µnode, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µnode = πTemp003
					// line 398: skip = 2
					πF.SetLineno(398)
					µskip = πg.NewInt(2).ToObject()
					goto Label28
					// line 399: elif c2.isalpha():
					πF.SetLineno(399)
				Label25:
					// line 401: i = 2
					πF.SetLineno(401)
					µi = πg.NewInt(2).ToObject()
					// line 402: while i < n and string[i].isalpha():
					πF.SetLineno(402)
					πF.PushCheckpoint(30)
					πTemp008 = false
				Label29:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp008 {
						πF.PopCheckpoint()
						goto Label31
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LT(πF, µi, µn); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp011, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp011 {
						goto Label32
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp003 = µi
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetItem(πF, µstring, πTemp003); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp009, ßisalpha, nil); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002 = πTemp009
				Label32:
					if πTemp010, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp010 {
						continue
					}
					πF.PushCheckpoint(29)
					// line 403: i += 1
					πF.SetLineno(403)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µi = πTemp002
					continue
				Label30:
					if πE != nil || πR != nil {
						continue
					}
				Label31:
					// line 404: name = string[1:i]
					πF.SetLineno(404)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), µi, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µstring, πTemp002); πE != nil {
						continue
					}
					µname = πTemp003
					// line 405: node, skip = handle_keyword(name, node, string[i:])
					πF.SetLineno(405)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001[0] = µname
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					πTemp001[1] = µnode
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µi, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µstring, πTemp002); πE != nil {
						continue
					}
					πTemp001[2] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßhandle_keyword); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp009}}}, πTemp003); πE != nil {
						continue
					}
					µnode = πTemp002
					µskip = πTemp009
					// line 406: skip += i
					πF.SetLineno(406)
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, µskip, µi); πE != nil {
						continue
					}
					µskip = πTemp002
					goto Label28
					// line 407: elif c2 == '\\':
					πF.SetLineno(407)
				Label26:
					// line 409: entry = mtd()
					πF.SetLineno(409)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmtd); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µentry = πTemp003
					// line 410: row = mtr(entry)
					πF.SetLineno(410)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µentry, "entry"); πE != nil {
						continue
					}
					πTemp001[0] = µentry
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmtr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µrow = πTemp003
					// line 411: node.close().close().append(row)
					πF.SetLineno(411)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µrow, "row"); πE != nil {
						continue
					}
					πTemp001[0] = µrow
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µnode, ßclose, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßclose, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 412: node = entry
					πF.SetLineno(412)
					if πE = πg.CheckLocal(πF, µentry, "entry"); πE != nil {
						continue
					}
					µnode = µentry
					// line 413: skip = 2
					πF.SetLineno(413)
					µskip = πg.NewInt(2).ToObject()
					goto Label28
				Label27:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µc2, "c2"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µc, µc2).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewUnicode("Syntax error: \"%s%s\"").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßSyntaxError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 415: raise SyntaxError(u'Syntax error: "%s%s"' % (c, c2))
					πF.SetLineno(415)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label28
				Label28:
					goto Label21
					// line 416: elif c.isalpha():
					πF.SetLineno(416)
				Label12:
					// line 417: node = node.append(mi(c))
					πF.SetLineno(417)
					πTemp001 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					πTemp006[0] = µc
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmi); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µnode, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µnode = πTemp003
					goto Label21
					// line 418: elif c.isdigit():
					πF.SetLineno(418)
				Label13:
					// line 419: node = node.append(mn(c))
					πF.SetLineno(419)
					πTemp001 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					πTemp006[0] = µc
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmn); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µnode, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µnode = πTemp003
					goto Label21
					// line 420: elif c in "+-*/=()[]|<>,.!?':;@":
					πF.SetLineno(420)
				Label14:
					// line 421: node = node.append(mo(c))
					πF.SetLineno(421)
					πTemp001 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					πTemp006[0] = µc
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmo); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µnode, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µnode = πTemp003
					goto Label21
					// line 422: elif c == '_':
					πF.SetLineno(422)
				Label15:
					// line 423: child = node.delete_child()
					πF.SetLineno(423)
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µnode, ßdelete_child, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µchild = πTemp003
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
						continue
					}
					πTemp001[0] = µchild
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmsup); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label33
					}
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
						continue
					}
					πTemp001[0] = µchild
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmo); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp002 = πTemp009
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp008 {
						goto Label34
					}
					if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, µchild, ßdata, nil); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveGlobal(πF, ßsumintprod); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Contains(πF, πTemp012, πTemp009); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp010).ToObject()
					πTemp002 = πTemp003
				Label34:
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label35
					}
					goto Label36
					// line 424: if isinstance(child, msup):
					πF.SetLineno(424)
				Label33:
					// line 425: sub = msubsup(child.children, reversed=True)
					πF.SetLineno(425)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µchild, ßchildren, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp005 = πg.KWArgs{
						{"reversed", πTemp002},
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmsubsup); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp005); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µsub = πTemp003
					goto Label37
					// line 426: elif isinstance(child, mo) and child.data in sumintprod:
					πF.SetLineno(426)
				Label35:
					// line 427: sub = munder(child)
					πF.SetLineno(427)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
						continue
					}
					πTemp001[0] = µchild
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmunder); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µsub = πTemp003
					goto Label37
				Label36:
					// line 429: sub = msub(child)
					πF.SetLineno(429)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
						continue
					}
					πTemp001[0] = µchild
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmsub); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µsub = πTemp003
					goto Label37
				Label37:
					// line 430: node.append(sub)
					πF.SetLineno(430)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsub, "sub"); πE != nil {
						continue
					}
					πTemp001[0] = µsub
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µnode, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 431: node = sub
					πF.SetLineno(431)
					if πE = πg.CheckLocal(πF, µsub, "sub"); πE != nil {
						continue
					}
					µnode = µsub
					goto Label21
					// line 432: elif c == '^':
					πF.SetLineno(432)
				Label16:
					// line 433: child = node.delete_child()
					πF.SetLineno(433)
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µnode, ßdelete_child, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µchild = πTemp003
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
						continue
					}
					πTemp001[0] = µchild
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmsub); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label38
					}
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
						continue
					}
					πTemp001[0] = µchild
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmo); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp002 = πTemp009
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp008 {
						goto Label39
					}
					if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, µchild, ßdata, nil); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveGlobal(πF, ßsumintprod); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Contains(πF, πTemp012, πTemp009); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp010).ToObject()
					πTemp002 = πTemp003
				Label39:
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label40
					}
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
						continue
					}
					πTemp001[0] = µchild
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmunder); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp002 = πTemp009
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp008 {
						goto Label41
					}
					πTemp009 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetAttr(πF, µchild, ßchildren, nil); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetItem(πF, πTemp013, πTemp009); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, πTemp012, ßdata, nil); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveGlobal(πF, ßsumintprod); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Contains(πF, πTemp012, πTemp009); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp010).ToObject()
					πTemp002 = πTemp003
				Label41:
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label42
					}
					goto Label43
					// line 434: if isinstance(child, msub):
					πF.SetLineno(434)
				Label38:
					// line 435: sup = msubsup(child.children)
					πF.SetLineno(435)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µchild, ßchildren, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmsubsup); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µsup = πTemp003
					goto Label44
					// line 436: elif isinstance(child, mo) and child.data in sumintprod:
					πF.SetLineno(436)
				Label40:
					// line 437: sup = mover(child)
					πF.SetLineno(437)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
						continue
					}
					πTemp001[0] = µchild
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmover); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µsup = πTemp003
					goto Label44
					// line 438: elif (isinstance(child, munder) and
					πF.SetLineno(438)
				Label42:
					// line 440: sup = munderover(child.children)
					πF.SetLineno(440)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µchild, ßchildren, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmunderover); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µsup = πTemp003
					goto Label44
				Label43:
					// line 442: sup = msup(child)
					πF.SetLineno(442)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µchild, "child"); πE != nil {
						continue
					}
					πTemp001[0] = µchild
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmsup); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µsup = πTemp003
					goto Label44
				Label44:
					// line 443: node.append(sup)
					πF.SetLineno(443)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsup, "sup"); πE != nil {
						continue
					}
					πTemp001[0] = µsup
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µnode, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 444: node = sup
					πF.SetLineno(444)
					if πE = πg.CheckLocal(πF, µsup, "sup"); πE != nil {
						continue
					}
					µnode = µsup
					goto Label21
					// line 445: elif c == '{':
					πF.SetLineno(445)
				Label17:
					// line 446: row = mrow()
					πF.SetLineno(446)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmrow); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µrow = πTemp003
					// line 447: node.append(row)
					πF.SetLineno(447)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µrow, "row"); πE != nil {
						continue
					}
					πTemp001[0] = µrow
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µnode, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 448: node = row
					πF.SetLineno(448)
					if πE = πg.CheckLocal(πF, µrow, "row"); πE != nil {
						continue
					}
					µnode = µrow
					goto Label21
					// line 449: elif c == '}':
					πF.SetLineno(449)
				Label18:
					// line 450: node = node.close()
					πF.SetLineno(450)
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µnode, ßclose, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µnode = πTemp003
					goto Label21
					// line 451: elif c == '&':
					πF.SetLineno(451)
				Label19:
					// line 452: entry = mtd()
					πF.SetLineno(452)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßmtd); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µentry = πTemp003
					// line 453: node.close().append(entry)
					πF.SetLineno(453)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µentry, "entry"); πE != nil {
						continue
					}
					πTemp001[0] = µentry
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µnode, ßclose, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 454: node = entry
					πF.SetLineno(454)
					if πE = πg.CheckLocal(πF, µentry, "entry"); πE != nil {
						continue
					}
					µnode = µentry
					goto Label21
				Label20:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewUnicode("Illegal character: \"%s\"").ToObject(), µc); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßSyntaxError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 456: raise SyntaxError(u'Illegal character: "%s"' % c)
					πF.SetLineno(456)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label21
				Label21:
					// line 457: string = string[skip:]
					πF.SetLineno(457)
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µskip, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µstring, πTemp002); πE != nil {
						continue
					}
					µstring = πTemp003
					continue
				Label5:
					if πE != nil || πR != nil {
						continue
					}
				Label6:
					// line 458: return tree
					πF.SetLineno(458)
					if πE = πg.CheckLocal(πF, µtree, "tree"); πE != nil {
						continue
					}
					πR = µtree
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßparse_latex_math.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 362: """parse_latex_math(string [,inline]) -> MathML-tree
			πF.SetLineno(362)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("parse_latex_math(string [,inline]) -> MathML-tree\n\n    Returns a MathML-tree parsed from string.  inline=True is for\n    inline math and inline=False is for displayed math.\n\n    tree is the whole tree and node is the current element.").ToObject()); πE != nil {
				continue
			}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßparse_latex_math); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
				continue
			}
			// line 461: def handle_keyword(name, node, string):
			πF.SetLineno(461)
			πTemp005 = make([]πg.Param, 3)
			πTemp005[0] = πg.Param{Name: "name", Def: nil}
			πTemp005[1] = πg.Param{Name: "node", Def: nil}
			πTemp005[2] = πg.Param{Name: "string", Def: nil}
			πTemp006 = πg.NewFunction(πg.NewCode("handle_keyword", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µname *πg.Object = πArgs[0]
				_ = µname
				var µnode *πg.Object = πArgs[1]
				_ = µnode
				var µstring *πg.Object = πArgs[2]
				_ = µstring
				var µskip *πg.Object = πg.UnboundLocal
				_ = µskip
				var µentry *πg.Object = πg.UnboundLocal
				_ = µentry
				var µtable *πg.Object = πg.UnboundLocal
				_ = µtable
				var µi *πg.Object = πg.UnboundLocal
				_ = µi
				var µsqrt *πg.Object = πg.UnboundLocal
				_ = µsqrt
				var µfrac *πg.Object = πg.UnboundLocal
				_ = µfrac
				var µpar *πg.Object = πg.UnboundLocal
				_ = µpar
				var µfenced *πg.Object = πg.UnboundLocal
				_ = µfenced
				var µrow *πg.Object = πg.UnboundLocal
				_ = µrow
				var µoperator *πg.Object = πg.UnboundLocal
				_ = µoperator
				var µstyle *πg.Object = πg.UnboundLocal
				_ = µstyle
				var µovr *πg.Object = πg.UnboundLocal
				_ = µovr
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 πg.KWArgs
				_ = πTemp009
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 32:
						goto Label32
					case 36:
						goto Label36
					case 37:
						goto Label37
					case 41:
						goto Label41
					case 42:
						goto Label42
					case 31:
						goto Label31
					default:
						panic("unexpected function state")
					}
					// line 462: skip = 0
					πF.SetLineno(462)
					µskip = πg.NewInt(0).ToObject()
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					πTemp004[0] = µstring
					if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp003, πE = πg.GT(πF, πTemp006, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					πTemp005 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µstring, πTemp005); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp006, πg.NewStr(" ").ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
				Label1:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label2
					}
					goto Label3
					// line 463: if len(string) > 0 and string[0] == ' ':
					πF.SetLineno(463)
				Label2:
					// line 464: string = string[1:]
					πF.SetLineno(464)
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µstring, πTemp001); πE != nil {
						continue
					}
					µstring = πTemp003
					// line 465: skip = 1
					πF.SetLineno(465)
					µskip = πg.NewInt(1).ToObject()
					goto Label3
				Label3:
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µname, ßbegin.ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label4
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µname, ßend.ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label5
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(ßtext.ToObject(), ßmathrm.ToObject()).ToObject()
					if πTemp002, πE = πg.Contains(πF, πTemp003, µname); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label6
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µname, ßsqrt.ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label7
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µname, ßfrac.ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label8
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µname, ßleft.ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label9
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µname, ßright.ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label10
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µname, ßnot.ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label11
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µname, ßmathbf.ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label12
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µname, ßmathbb.ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label13
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(ßmathscr.ToObject(), ßmathcal.ToObject()).ToObject()
					if πTemp002, πE = πg.Contains(πF, πTemp003, µname); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label14
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µname, ßcolon.ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label15
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßGreek); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Contains(πF, πTemp003, µname); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label16
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßletters); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Contains(πF, πTemp003, µname); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label17
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßspecial); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Contains(πF, πTemp003, µname); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label18
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßfunctions); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Contains(πF, πTemp003, µname); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label19
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßover); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Contains(πF, πTemp003, µname); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label20
					}
					goto Label21
					// line 466: if name == 'begin':
					πF.SetLineno(466)
				Label4:
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("{matrix}").ToObject()
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µstring, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp002, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label23
					}
					goto Label24
					// line 467: if not string.startswith('{matrix}'):
					πF.SetLineno(467)
				Label23:
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewUnicode("Environment not supported! Supported environment: \"matrix\".").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßSyntaxError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 468: raise SyntaxError(u'Environment not supported! '
					πF.SetLineno(468)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label24
				Label24:
					// line 470: skip += 8
					πF.SetLineno(470)
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAdd(πF, µskip, πg.NewInt(8).ToObject()); πE != nil {
						continue
					}
					µskip = πTemp001
					// line 471: entry = mtd()
					πF.SetLineno(471)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmtd); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µentry = πTemp003
					// line 472: table = mtable(mtr(entry))
					πF.SetLineno(472)
					πTemp004 = πF.MakeArgs(1)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µentry, "entry"); πE != nil {
						continue
					}
					πTemp007[0] = µentry
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmtr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp004[0] = πTemp003
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmtable); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µtable = πTemp003
					// line 473: node.append(table)
					πF.SetLineno(473)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtable, "table"); πE != nil {
						continue
					}
					πTemp004[0] = µtable
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µnode, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 474: node = entry
					πF.SetLineno(474)
					if πE = πg.CheckLocal(πF, µentry, "entry"); πE != nil {
						continue
					}
					µnode = µentry
					goto Label22
					// line 475: elif name == 'end':
					πF.SetLineno(475)
				Label5:
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("{matrix}").ToObject()
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µstring, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp002, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label25
					}
					goto Label26
					// line 476: if not string.startswith('{matrix}'):
					πF.SetLineno(476)
				Label25:
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewUnicode("Expected \"\\end{matrix}\"!").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßSyntaxError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 477: raise SyntaxError(u'Expected "\\end{matrix}"!')
					πF.SetLineno(477)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label26
				Label26:
					// line 478: skip += 8
					πF.SetLineno(478)
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAdd(πF, µskip, πg.NewInt(8).ToObject()); πE != nil {
						continue
					}
					µskip = πTemp001
					// line 479: node = node.close().close().close()
					πF.SetLineno(479)
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µnode, ßclose, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßclose, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßclose, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µnode = πTemp003
					goto Label22
					// line 480: elif name in ('text', 'mathrm'):
					πF.SetLineno(480)
				Label6:
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µstring, πTemp003); πE != nil {
						continue
					}
					if πTemp001, πE = πg.NE(πF, πTemp005, πg.NewStr("{").ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label27
					}
					goto Label28
					// line 481: if string[0] != '{':
					πF.SetLineno(481)
				Label27:
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewUnicode("Expected \"\\text{...}\"!").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßSyntaxError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 482: raise SyntaxError(u'Expected "\\text{...}"!')
					πF.SetLineno(482)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label28
				Label28:
					// line 483: i = string.find('}')
					πF.SetLineno(483)
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("}").ToObject()
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µstring, ßfind, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µi = πTemp003
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µi, πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label29
					}
					goto Label30
					// line 484: if i == -1:
					πF.SetLineno(484)
				Label29:
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewUnicode("Expected \"\\text{...}\"!").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßSyntaxError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 485: raise SyntaxError(u'Expected "\\text{...}"!')
					πF.SetLineno(485)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label30
				Label30:
					// line 486: node = node.append(mtext(string[1:i]))
					πF.SetLineno(486)
					πTemp004 = πF.MakeArgs(1)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), µi, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µstring, πTemp001); πE != nil {
						continue
					}
					πTemp007[0] = πTemp003
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmtext); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp004[0] = πTemp003
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µnode, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µnode = πTemp003
					// line 487: skip += i + 1
					πF.SetLineno(487)
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IAdd(πF, µskip, πTemp001); πE != nil {
						continue
					}
					µskip = πTemp003
					goto Label22
					// line 488: elif name == 'sqrt':
					πF.SetLineno(488)
				Label7:
					// line 489: sqrt = msqrt()
					πF.SetLineno(489)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmsqrt); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µsqrt = πTemp003
					// line 490: node.append(sqrt)
					πF.SetLineno(490)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsqrt, "sqrt"); πE != nil {
						continue
					}
					πTemp004[0] = µsqrt
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µnode, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 491: node = sqrt
					πF.SetLineno(491)
					if πE = πg.CheckLocal(πF, µsqrt, "sqrt"); πE != nil {
						continue
					}
					µnode = µsqrt
					goto Label22
					// line 492: elif name == 'frac':
					πF.SetLineno(492)
				Label8:
					// line 493: frac = mfrac()
					πF.SetLineno(493)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmfrac); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µfrac = πTemp003
					// line 494: node.append(frac)
					πF.SetLineno(494)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfrac, "frac"); πE != nil {
						continue
					}
					πTemp004[0] = µfrac
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µnode, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 495: node = frac
					πF.SetLineno(495)
					if πE = πg.CheckLocal(πF, µfrac, "frac"); πE != nil {
						continue
					}
					µnode = µfrac
					goto Label22
					// line 496: elif name == 'left':
					πF.SetLineno(496)
				Label9:
					πTemp004 = make([]*πg.Object, 6)
					πTemp004[0] = πg.NewStr("(").ToObject()
					πTemp004[1] = πg.NewStr("[").ToObject()
					πTemp004[2] = πg.NewStr("|").ToObject()
					πTemp004[3] = πg.NewStr("\\{").ToObject()
					πTemp004[4] = πg.NewStr("\\langle").ToObject()
					πTemp004[5] = πg.NewStr(".").ToObject()
					πTemp003 = πg.NewList(πTemp004...).ToObject()
					if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
						continue
					}
					πF.PushCheckpoint(32)
					πTemp002 = false
				Label31:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp002 {
						πF.PopCheckpoint()
						goto Label33
					}
					if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
						µpar = πTemp003
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(31)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpar, "par"); πE != nil {
						continue
					}
					πTemp004[0] = µpar
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µstring, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label34
					}
					goto Label35
					// line 498: if string.startswith(par):
					πF.SetLineno(498)
				Label34:
					// line 499: break
					πF.SetLineno(499)
					πTemp002 = true
					continue
					goto Label35
				Label35:
					continue
				Label32:
					if πE != nil || πR != nil {
						continue
					}
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewUnicode("Missing left-brace!").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßSyntaxError); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 501: raise SyntaxError(u'Missing left-brace!')
					πF.SetLineno(501)
					πE = πF.Raise(πTemp005, nil, nil)
					continue
				Label33:
					// line 502: fenced = mfenced(par)
					πF.SetLineno(502)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpar, "par"); πE != nil {
						continue
					}
					πTemp004[0] = µpar
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmfenced); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µfenced = πTemp003
					// line 503: node.append(fenced)
					πF.SetLineno(503)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfenced, "fenced"); πE != nil {
						continue
					}
					πTemp004[0] = µfenced
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µnode, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 504: row = mrow()
					πF.SetLineno(504)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmrow); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µrow = πTemp003
					// line 505: fenced.append(row)
					πF.SetLineno(505)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µrow, "row"); πE != nil {
						continue
					}
					πTemp004[0] = µrow
					if πE = πg.CheckLocal(πF, µfenced, "fenced"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µfenced, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 506: node = row
					πF.SetLineno(506)
					if πE = πg.CheckLocal(πF, µrow, "row"); πE != nil {
						continue
					}
					µnode = µrow
					// line 507: skip += len(par)
					πF.SetLineno(507)
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpar, "par"); πE != nil {
						continue
					}
					πTemp004[0] = µpar
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp001, πE = πg.IAdd(πF, µskip, πTemp003); πE != nil {
						continue
					}
					µskip = πTemp001
					goto Label22
					// line 508: elif name == 'right':
					πF.SetLineno(508)
				Label10:
					πTemp004 = make([]*πg.Object, 6)
					πTemp004[0] = πg.NewStr(")").ToObject()
					πTemp004[1] = πg.NewStr("]").ToObject()
					πTemp004[2] = πg.NewStr("|").ToObject()
					πTemp004[3] = πg.NewStr("\\}").ToObject()
					πTemp004[4] = πg.NewStr("\\rangle").ToObject()
					πTemp004[5] = πg.NewStr(".").ToObject()
					πTemp003 = πg.NewList(πTemp004...).ToObject()
					if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
						continue
					}
					πF.PushCheckpoint(37)
					πTemp002 = false
				Label36:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp002 {
						πF.PopCheckpoint()
						goto Label38
					}
					if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
						µpar = πTemp003
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(36)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpar, "par"); πE != nil {
						continue
					}
					πTemp004[0] = µpar
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µstring, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label39
					}
					goto Label40
					// line 510: if string.startswith(par):
					πF.SetLineno(510)
				Label39:
					// line 511: break
					πF.SetLineno(511)
					πTemp002 = true
					continue
					goto Label40
				Label40:
					continue
				Label37:
					if πE != nil || πR != nil {
						continue
					}
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewUnicode("Missing right-brace!").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßSyntaxError); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 513: raise SyntaxError(u'Missing right-brace!')
					πF.SetLineno(513)
					πE = πF.Raise(πTemp005, nil, nil)
					continue
				Label38:
					// line 514: node = node.close()
					πF.SetLineno(514)
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µnode, ßclose, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µnode = πTemp003
					// line 515: node.closepar = par
					πF.SetLineno(515)
					if πE = πg.CheckLocal(πF, µpar, "par"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µpar); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µnode, ßclosepar, πTemp001); πE != nil {
						continue
					}
					// line 516: node = node.close()
					πF.SetLineno(516)
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µnode, ßclose, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µnode = πTemp003
					// line 517: skip += len(par)
					πF.SetLineno(517)
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpar, "par"); πE != nil {
						continue
					}
					πTemp004[0] = µpar
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp001, πE = πg.IAdd(πF, µskip, πTemp003); πE != nil {
						continue
					}
					µskip = πTemp001
					goto Label22
					// line 518: elif name == 'not':
					πF.SetLineno(518)
				Label11:
					if πTemp003, πE = πg.ResolveGlobal(πF, ßnegatables); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
						continue
					}
					πF.PushCheckpoint(42)
					πTemp002 = false
				Label41:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp002 {
						πF.PopCheckpoint()
						goto Label43
					}
					if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
						µoperator = πTemp003
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(41)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µoperator, "operator"); πE != nil {
						continue
					}
					πTemp004[0] = µoperator
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µstring, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label44
					}
					goto Label45
					// line 520: if string.startswith(operator):
					πF.SetLineno(520)
				Label44:
					// line 521: break
					πF.SetLineno(521)
					πTemp002 = true
					continue
					goto Label45
				Label45:
					continue
				Label42:
					if πE != nil || πR != nil {
						continue
					}
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewUnicode("Expected something to negate: \"\\not ...\"!").ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßSyntaxError); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 523: raise SyntaxError(u'Expected something to negate: "\\not ..."!')
					πF.SetLineno(523)
					πE = πF.Raise(πTemp005, nil, nil)
					continue
				Label43:
					// line 524: node = node.append(mo(negatables[operator]))
					πF.SetLineno(524)
					πTemp004 = πF.MakeArgs(1)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µoperator, "operator"); πE != nil {
						continue
					}
					πTemp001 = µoperator
					if πTemp005, πE = πg.ResolveGlobal(πF, ßnegatables); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
						continue
					}
					πTemp007[0] = πTemp003
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmo); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp004[0] = πTemp003
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µnode, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µnode = πTemp003
					// line 525: skip += len(operator)
					πF.SetLineno(525)
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µoperator, "operator"); πE != nil {
						continue
					}
					πTemp004[0] = µoperator
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp001, πE = πg.IAdd(πF, µskip, πTemp003); πE != nil {
						continue
					}
					µskip = πTemp001
					goto Label22
					// line 526: elif name == 'mathbf':
					πF.SetLineno(526)
				Label12:
					// line 527: style = mstyle(nchildren=1, fontweight='bold')
					πF.SetLineno(527)
					πTemp009 = πg.KWArgs{
						{"nchildren", πg.NewInt(1).ToObject()},
						{"fontweight", ßbold.ToObject()},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmstyle); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, πTemp009); πE != nil {
						continue
					}
					µstyle = πTemp003
					// line 528: node.append(style)
					πF.SetLineno(528)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µstyle, "style"); πE != nil {
						continue
					}
					πTemp004[0] = µstyle
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µnode, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 529: node = style
					πF.SetLineno(529)
					if πE = πg.CheckLocal(πF, µstyle, "style"); πE != nil {
						continue
					}
					µnode = µstyle
					goto Label22
					// line 530: elif name == 'mathbb':
					πF.SetLineno(530)
				Label13:
					πTemp005 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µstring, πTemp005); πE != nil {
						continue
					}
					if πTemp003, πE = πg.NE(πF, πTemp006, πg.NewStr("{").ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label46
					}
					πTemp005 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µstring, πTemp005); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp006, ßisupper, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp008).ToObject()
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label46
					}
					πTemp005 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µstring, πTemp005); πE != nil {
						continue
					}
					if πTemp003, πE = πg.NE(πF, πTemp006, πg.NewStr("}").ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
				Label46:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label47
					}
					goto Label48
					// line 531: if string[0] != '{' or not string[1].isupper() or string[2] != '}':
					πF.SetLineno(531)
				Label47:
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewUnicode("Expected something like \"\\mathbb{A}\"!").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßSyntaxError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 532: raise SyntaxError(u'Expected something like "\\mathbb{A}"!')
					πF.SetLineno(532)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label48
				Label48:
					// line 533: node = node.append(mi(mathbb[string[1]]))
					πF.SetLineno(533)
					πTemp004 = πF.MakeArgs(1)
					πTemp007 = πF.MakeArgs(1)
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µstring, πTemp003); πE != nil {
						continue
					}
					πTemp001 = πTemp005
					if πTemp005, πE = πg.ResolveGlobal(πF, ßmathbb); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
						continue
					}
					πTemp007[0] = πTemp003
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmi); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp004[0] = πTemp003
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µnode, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µnode = πTemp003
					// line 534: skip += 3
					πF.SetLineno(534)
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAdd(πF, µskip, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					µskip = πTemp001
					goto Label22
					// line 535: elif name in ('mathscr', 'mathcal'):
					πF.SetLineno(535)
				Label14:
					πTemp005 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µstring, πTemp005); πE != nil {
						continue
					}
					if πTemp003, πE = πg.NE(πF, πTemp006, πg.NewStr("{").ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label49
					}
					πTemp005 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µstring, πTemp005); πE != nil {
						continue
					}
					if πTemp003, πE = πg.NE(πF, πTemp006, πg.NewStr("}").ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
				Label49:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label50
					}
					goto Label51
					// line 536: if string[0] != '{' or string[2] != '}':
					πF.SetLineno(536)
				Label50:
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewUnicode("Expected something like \"\\mathscr{A}\"!").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßSyntaxError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 537: raise SyntaxError(u'Expected something like "\\mathscr{A}"!')
					πF.SetLineno(537)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label51
				Label51:
					// line 538: node = node.append(mi(mathscr[string[1]]))
					πF.SetLineno(538)
					πTemp004 = πF.MakeArgs(1)
					πTemp007 = πF.MakeArgs(1)
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µstring, πTemp003); πE != nil {
						continue
					}
					πTemp001 = πTemp005
					if πTemp005, πE = πg.ResolveGlobal(πF, ßmathscr); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
						continue
					}
					πTemp007[0] = πTemp003
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmi); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp004[0] = πTemp003
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µnode, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µnode = πTemp003
					// line 539: skip += 3
					πF.SetLineno(539)
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAdd(πF, µskip, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					µskip = πTemp001
					goto Label22
					// line 540: elif name == 'colon': # "normal" colon, not binary operator
					πF.SetLineno(540)
				Label15:
					// line 541: node = node.append(mo(':')) # TODO: add ``lspace="0pt"``
					πF.SetLineno(541)
					πTemp004 = πF.MakeArgs(1)
					πTemp007 = πF.MakeArgs(1)
					πTemp007[0] = πg.NewStr(":").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmo); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp004[0] = πTemp003
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µnode, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µnode = πTemp003
					goto Label22
					// line 542: elif name in Greek:   # Greek capitals (upright in "TeX style")
					πF.SetLineno(542)
				Label16:
					// line 543: node = node.append(mo(Greek[name]))
					πF.SetLineno(543)
					πTemp004 = πF.MakeArgs(1)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001 = µname
					if πTemp005, πE = πg.ResolveGlobal(πF, ßGreek); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
						continue
					}
					πTemp007[0] = πTemp003
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmo); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp004[0] = πTemp003
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µnode, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µnode = πTemp003
					goto Label22
					// line 546: elif name in letters:
					πF.SetLineno(546)
				Label17:
					// line 547: node = node.append(mi(letters[name]))
					πF.SetLineno(547)
					πTemp004 = πF.MakeArgs(1)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001 = µname
					if πTemp005, πE = πg.ResolveGlobal(πF, ßletters); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
						continue
					}
					πTemp007[0] = πTemp003
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmi); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp004[0] = πTemp003
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µnode, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µnode = πTemp003
					goto Label22
					// line 548: elif name in special:
					πF.SetLineno(548)
				Label18:
					// line 549: node = node.append(mo(special[name]))
					πF.SetLineno(549)
					πTemp004 = πF.MakeArgs(1)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001 = µname
					if πTemp005, πE = πg.ResolveGlobal(πF, ßspecial); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
						continue
					}
					πTemp007[0] = πTemp003
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmo); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp004[0] = πTemp003
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µnode, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µnode = πTemp003
					goto Label22
					// line 550: elif name in functions:
					πF.SetLineno(550)
				Label19:
					// line 551: node = node.append(mo(name))
					πF.SetLineno(551)
					πTemp004 = πF.MakeArgs(1)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp007[0] = µname
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmo); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp004[0] = πTemp003
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µnode, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µnode = πTemp003
					goto Label22
					// line 552: elif name in over:
					πF.SetLineno(552)
				Label20:
					// line 553: ovr = mover(mo(over[name]), reversed=True)
					πF.SetLineno(553)
					πTemp004 = πF.MakeArgs(1)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001 = µname
					if πTemp005, πE = πg.ResolveGlobal(πF, ßover); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
						continue
					}
					πTemp007[0] = πTemp003
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmo); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp004[0] = πTemp003
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					πTemp009 = πg.KWArgs{
						{"reversed", πTemp001},
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmover); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, πTemp009); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µovr = πTemp003
					// line 554: node.append(ovr)
					πF.SetLineno(554)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µovr, "ovr"); πE != nil {
						continue
					}
					πTemp004[0] = µovr
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µnode, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 555: node = ovr
					πF.SetLineno(555)
					if πE = πg.CheckLocal(πF, µovr, "ovr"); πE != nil {
						continue
					}
					µnode = µovr
					goto Label22
				Label21:
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πg.NewUnicode("Unknown LaTeX command: ").ToObject(), µname); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßSyntaxError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 557: raise SyntaxError(u'Unknown LaTeX command: ' + name)
					πF.SetLineno(557)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label22
				Label22:
					// line 559: return node, skip
					πF.SetLineno(559)
					if πE = πg.CheckLocal(πF, µnode, "node"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(µnode, µskip).ToObject()
					πR = πTemp001
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßhandle_keyword.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 561: def tex2mathml(tex_math, inline=True):
			πF.SetLineno(561)
			πTemp005 = make([]πg.Param, 2)
			πTemp005[0] = πg.Param{Name: "tex_math", Def: nil}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp005[1] = πg.Param{Name: "inline", Def: πTemp008}
			πTemp007 = πg.NewFunction(πg.NewCode("tex2mathml", "/usr/lib/python2.7/site-packages/docutils/utils/math/latex2mathml.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtex_math *πg.Object = πArgs[0]
				_ = µtex_math
				var µinline *πg.Object = πArgs[1]
				_ = µinline
				var µmathml_tree *πg.Object = πg.UnboundLocal
				_ = µmathml_tree
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 πg.KWArgs
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
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
					// line 562: """Return string with MathML code corresponding to `tex_math`.
					πF.SetLineno(562)
					// line 567: mathml_tree = parse_latex_math(tex_math, inline=inline)
					πF.SetLineno(567)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µtex_math, "tex_math"); πE != nil {
						continue
					}
					πTemp001[0] = µtex_math
					if πE = πg.CheckLocal(πF, µinline, "inline"); πE != nil {
						continue
					}
					πTemp002 = πg.KWArgs{
						{"inline", µinline},
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßparse_latex_math); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, πTemp002); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µmathml_tree = πTemp004
					// line 568: return ''.join(mathml_tree.xml())
					πF.SetLineno(568)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmathml_tree, "mathml_tree"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µmathml_tree, ßxml, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp004
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßtex2mathml.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 562: """Return string with MathML code corresponding to `tex_math`.
			πF.SetLineno(562)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("Return string with MathML code corresponding to `tex_math`.\n\n    `inline`=True is for inline math and `inline`=False for displayed math.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßtex2mathml); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("latex2mathml", Code)
}
