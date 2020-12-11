package unichar2tex

import (
	πg "github.com/pygolin/runtime"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/utils/math/unichar2tex.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ßuni2tex_table := πg.InternStr("uni2tex_table")
		var πTemp001 *πg.Dict
		_ = πTemp001
		var πTemp002 *πg.Object
		_ = πTemp002
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default:
				panic("unexpected function state")
			}
			// line 8: uni2tex_table = {
			πF.SetLineno(8)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, πg.NewInt(160).ToObject(), πg.NewUnicode("~").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(163).ToObject(), πg.NewUnicode("\\pounds ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(165).ToObject(), πg.NewUnicode("\\yen ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(172).ToObject(), πg.NewUnicode("\\neg ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(174).ToObject(), πg.NewUnicode("\\circledR ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(177).ToObject(), πg.NewUnicode("\\pm ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(215).ToObject(), πg.NewUnicode("\\times ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(240).ToObject(), πg.NewUnicode("\\eth ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(247).ToObject(), πg.NewUnicode("\\div ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(305).ToObject(), πg.NewUnicode("\\imath ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(567).ToObject(), πg.NewUnicode("\\jmath ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(915).ToObject(), πg.NewUnicode("\\Gamma ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(916).ToObject(), πg.NewUnicode("\\Delta ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(920).ToObject(), πg.NewUnicode("\\Theta ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(923).ToObject(), πg.NewUnicode("\\Lambda ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(926).ToObject(), πg.NewUnicode("\\Xi ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(928).ToObject(), πg.NewUnicode("\\Pi ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(931).ToObject(), πg.NewUnicode("\\Sigma ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(933).ToObject(), πg.NewUnicode("\\Upsilon ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(934).ToObject(), πg.NewUnicode("\\Phi ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(936).ToObject(), πg.NewUnicode("\\Psi ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(937).ToObject(), πg.NewUnicode("\\Omega ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(945).ToObject(), πg.NewUnicode("\\alpha ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(946).ToObject(), πg.NewUnicode("\\beta ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(947).ToObject(), πg.NewUnicode("\\gamma ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(948).ToObject(), πg.NewUnicode("\\delta ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(949).ToObject(), πg.NewUnicode("\\varepsilon ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(950).ToObject(), πg.NewUnicode("\\zeta ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(951).ToObject(), πg.NewUnicode("\\eta ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(952).ToObject(), πg.NewUnicode("\\theta ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(953).ToObject(), πg.NewUnicode("\\iota ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(954).ToObject(), πg.NewUnicode("\\kappa ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(955).ToObject(), πg.NewUnicode("\\lambda ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(956).ToObject(), πg.NewUnicode("\\mu ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(957).ToObject(), πg.NewUnicode("\\nu ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(958).ToObject(), πg.NewUnicode("\\xi ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(960).ToObject(), πg.NewUnicode("\\pi ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(961).ToObject(), πg.NewUnicode("\\rho ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(962).ToObject(), πg.NewUnicode("\\varsigma ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(963).ToObject(), πg.NewUnicode("\\sigma ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(964).ToObject(), πg.NewUnicode("\\tau ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(965).ToObject(), πg.NewUnicode("\\upsilon ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(966).ToObject(), πg.NewUnicode("\\varphi ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(967).ToObject(), πg.NewUnicode("\\chi ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(968).ToObject(), πg.NewUnicode("\\psi ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(969).ToObject(), πg.NewUnicode("\\omega ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(977).ToObject(), πg.NewUnicode("\\vartheta ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(981).ToObject(), πg.NewUnicode("\\phi ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(982).ToObject(), πg.NewUnicode("\\varpi ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(989).ToObject(), πg.NewUnicode("\\digamma ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(1014).ToObject(), πg.NewUnicode("\\backepsilon ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8193).ToObject(), πg.NewUnicode("\\quad ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8214).ToObject(), πg.NewUnicode("\\| ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8224).ToObject(), πg.NewUnicode("\\dagger ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8225).ToObject(), πg.NewUnicode("\\ddagger ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8230).ToObject(), πg.NewUnicode("\\ldots ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8242).ToObject(), πg.NewUnicode("\\prime ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8245).ToObject(), πg.NewUnicode("\\backprime ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8287).ToObject(), πg.NewUnicode("\\: ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8450).ToObject(), πg.NewUnicode("\\mathbb{C}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8459).ToObject(), πg.NewUnicode("\\mathcal{H}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8460).ToObject(), πg.NewUnicode("\\mathfrak{H}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8461).ToObject(), πg.NewUnicode("\\mathbb{H}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8463).ToObject(), πg.NewUnicode("\\hslash ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8464).ToObject(), πg.NewUnicode("\\mathcal{I}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8465).ToObject(), πg.NewUnicode("\\Im ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8466).ToObject(), πg.NewUnicode("\\mathcal{L}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8467).ToObject(), πg.NewUnicode("\\ell ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8469).ToObject(), πg.NewUnicode("\\mathbb{N}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8472).ToObject(), πg.NewUnicode("\\wp ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8473).ToObject(), πg.NewUnicode("\\mathbb{P}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8474).ToObject(), πg.NewUnicode("\\mathbb{Q}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8475).ToObject(), πg.NewUnicode("\\mathcal{R}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8476).ToObject(), πg.NewUnicode("\\Re ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8477).ToObject(), πg.NewUnicode("\\mathbb{R}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8484).ToObject(), πg.NewUnicode("\\mathbb{Z}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8487).ToObject(), πg.NewUnicode("\\mho ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8488).ToObject(), πg.NewUnicode("\\mathfrak{Z}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8492).ToObject(), πg.NewUnicode("\\mathcal{B}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8493).ToObject(), πg.NewUnicode("\\mathfrak{C}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8496).ToObject(), πg.NewUnicode("\\mathcal{E}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8497).ToObject(), πg.NewUnicode("\\mathcal{F}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8498).ToObject(), πg.NewUnicode("\\Finv ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8499).ToObject(), πg.NewUnicode("\\mathcal{M}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8501).ToObject(), πg.NewUnicode("\\aleph ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8502).ToObject(), πg.NewUnicode("\\beth ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8503).ToObject(), πg.NewUnicode("\\gimel ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8504).ToObject(), πg.NewUnicode("\\daleth ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8592).ToObject(), πg.NewUnicode("\\leftarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8593).ToObject(), πg.NewUnicode("\\uparrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8594).ToObject(), πg.NewUnicode("\\rightarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8595).ToObject(), πg.NewUnicode("\\downarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8596).ToObject(), πg.NewUnicode("\\leftrightarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8597).ToObject(), πg.NewUnicode("\\updownarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8598).ToObject(), πg.NewUnicode("\\nwarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8599).ToObject(), πg.NewUnicode("\\nearrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8600).ToObject(), πg.NewUnicode("\\searrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8601).ToObject(), πg.NewUnicode("\\swarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8602).ToObject(), πg.NewUnicode("\\nleftarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8603).ToObject(), πg.NewUnicode("\\nrightarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8606).ToObject(), πg.NewUnicode("\\twoheadleftarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8608).ToObject(), πg.NewUnicode("\\twoheadrightarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8610).ToObject(), πg.NewUnicode("\\leftarrowtail ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8611).ToObject(), πg.NewUnicode("\\rightarrowtail ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8614).ToObject(), πg.NewUnicode("\\mapsto ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8617).ToObject(), πg.NewUnicode("\\hookleftarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8618).ToObject(), πg.NewUnicode("\\hookrightarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8619).ToObject(), πg.NewUnicode("\\looparrowleft ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8620).ToObject(), πg.NewUnicode("\\looparrowright ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8621).ToObject(), πg.NewUnicode("\\leftrightsquigarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8622).ToObject(), πg.NewUnicode("\\nleftrightarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8624).ToObject(), πg.NewUnicode("\\Lsh ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8625).ToObject(), πg.NewUnicode("\\Rsh ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8630).ToObject(), πg.NewUnicode("\\curvearrowleft ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8631).ToObject(), πg.NewUnicode("\\curvearrowright ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8634).ToObject(), πg.NewUnicode("\\circlearrowleft ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8635).ToObject(), πg.NewUnicode("\\circlearrowright ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8636).ToObject(), πg.NewUnicode("\\leftharpoonup ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8637).ToObject(), πg.NewUnicode("\\leftharpoondown ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8638).ToObject(), πg.NewUnicode("\\upharpoonright ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8639).ToObject(), πg.NewUnicode("\\upharpoonleft ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8640).ToObject(), πg.NewUnicode("\\rightharpoonup ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8641).ToObject(), πg.NewUnicode("\\rightharpoondown ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8642).ToObject(), πg.NewUnicode("\\downharpoonright ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8643).ToObject(), πg.NewUnicode("\\downharpoonleft ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8644).ToObject(), πg.NewUnicode("\\rightleftarrows ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8646).ToObject(), πg.NewUnicode("\\leftrightarrows ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8647).ToObject(), πg.NewUnicode("\\leftleftarrows ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8648).ToObject(), πg.NewUnicode("\\upuparrows ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8649).ToObject(), πg.NewUnicode("\\rightrightarrows ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8650).ToObject(), πg.NewUnicode("\\downdownarrows ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8651).ToObject(), πg.NewUnicode("\\leftrightharpoons ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8652).ToObject(), πg.NewUnicode("\\rightleftharpoons ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8653).ToObject(), πg.NewUnicode("\\nLeftarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8654).ToObject(), πg.NewUnicode("\\nLeftrightarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8655).ToObject(), πg.NewUnicode("\\nRightarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8656).ToObject(), πg.NewUnicode("\\Leftarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8657).ToObject(), πg.NewUnicode("\\Uparrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8658).ToObject(), πg.NewUnicode("\\Rightarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8659).ToObject(), πg.NewUnicode("\\Downarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8660).ToObject(), πg.NewUnicode("\\Leftrightarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8661).ToObject(), πg.NewUnicode("\\Updownarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8666).ToObject(), πg.NewUnicode("\\Lleftarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8667).ToObject(), πg.NewUnicode("\\Rrightarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8669).ToObject(), πg.NewUnicode("\\rightsquigarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8672).ToObject(), πg.NewUnicode("\\dashleftarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8674).ToObject(), πg.NewUnicode("\\dashrightarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8704).ToObject(), πg.NewUnicode("\\forall ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8705).ToObject(), πg.NewUnicode("\\complement ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8706).ToObject(), πg.NewUnicode("\\partial ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8707).ToObject(), πg.NewUnicode("\\exists ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8708).ToObject(), πg.NewUnicode("\\nexists ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8709).ToObject(), πg.NewUnicode("\\varnothing ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8711).ToObject(), πg.NewUnicode("\\nabla ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8712).ToObject(), πg.NewUnicode("\\in ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8713).ToObject(), πg.NewUnicode("\\notin ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8715).ToObject(), πg.NewUnicode("\\ni ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8719).ToObject(), πg.NewUnicode("\\prod ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8720).ToObject(), πg.NewUnicode("\\coprod ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8721).ToObject(), πg.NewUnicode("\\sum ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8722).ToObject(), πg.NewUnicode("-").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8723).ToObject(), πg.NewUnicode("\\mp ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8724).ToObject(), πg.NewUnicode("\\dotplus ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8725).ToObject(), πg.NewUnicode("\\slash ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8726).ToObject(), πg.NewUnicode("\\smallsetminus ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8727).ToObject(), πg.NewUnicode("\\ast ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8728).ToObject(), πg.NewUnicode("\\circ ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8729).ToObject(), πg.NewUnicode("\\bullet ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8730).ToObject(), πg.NewUnicode("\\sqrt ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8731).ToObject(), πg.NewUnicode("\\sqrt[3] ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8732).ToObject(), πg.NewUnicode("\\sqrt[4] ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8733).ToObject(), πg.NewUnicode("\\propto ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8734).ToObject(), πg.NewUnicode("\\infty ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8736).ToObject(), πg.NewUnicode("\\angle ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8737).ToObject(), πg.NewUnicode("\\measuredangle ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8738).ToObject(), πg.NewUnicode("\\sphericalangle ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8739).ToObject(), πg.NewUnicode("\\mid ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8740).ToObject(), πg.NewUnicode("\\nmid ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8741).ToObject(), πg.NewUnicode("\\parallel ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8742).ToObject(), πg.NewUnicode("\\nparallel ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8743).ToObject(), πg.NewUnicode("\\wedge ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8744).ToObject(), πg.NewUnicode("\\vee ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8745).ToObject(), πg.NewUnicode("\\cap ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8746).ToObject(), πg.NewUnicode("\\cup ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8747).ToObject(), πg.NewUnicode("\\int ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8748).ToObject(), πg.NewUnicode("\\iint ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8749).ToObject(), πg.NewUnicode("\\iiint ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8750).ToObject(), πg.NewUnicode("\\oint ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8756).ToObject(), πg.NewUnicode("\\therefore ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8757).ToObject(), πg.NewUnicode("\\because ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8758).ToObject(), πg.NewUnicode(":").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8764).ToObject(), πg.NewUnicode("\\sim ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8765).ToObject(), πg.NewUnicode("\\backsim ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8768).ToObject(), πg.NewUnicode("\\wr ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8769).ToObject(), πg.NewUnicode("\\nsim ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8770).ToObject(), πg.NewUnicode("\\eqsim ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8771).ToObject(), πg.NewUnicode("\\simeq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8773).ToObject(), πg.NewUnicode("\\cong ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8775).ToObject(), πg.NewUnicode("\\ncong ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8776).ToObject(), πg.NewUnicode("\\approx ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8778).ToObject(), πg.NewUnicode("\\approxeq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8781).ToObject(), πg.NewUnicode("\\asymp ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8782).ToObject(), πg.NewUnicode("\\Bumpeq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8783).ToObject(), πg.NewUnicode("\\bumpeq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8784).ToObject(), πg.NewUnicode("\\doteq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8785).ToObject(), πg.NewUnicode("\\Doteq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8786).ToObject(), πg.NewUnicode("\\fallingdotseq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8787).ToObject(), πg.NewUnicode("\\risingdotseq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8790).ToObject(), πg.NewUnicode("\\eqcirc ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8791).ToObject(), πg.NewUnicode("\\circeq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8796).ToObject(), πg.NewUnicode("\\triangleq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8800).ToObject(), πg.NewUnicode("\\neq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8801).ToObject(), πg.NewUnicode("\\equiv ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8804).ToObject(), πg.NewUnicode("\\leq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8805).ToObject(), πg.NewUnicode("\\geq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8806).ToObject(), πg.NewUnicode("\\leqq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8807).ToObject(), πg.NewUnicode("\\geqq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8808).ToObject(), πg.NewUnicode("\\lneqq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8809).ToObject(), πg.NewUnicode("\\gneqq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8810).ToObject(), πg.NewUnicode("\\ll ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8811).ToObject(), πg.NewUnicode("\\gg ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8812).ToObject(), πg.NewUnicode("\\between ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8814).ToObject(), πg.NewUnicode("\\nless ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8815).ToObject(), πg.NewUnicode("\\ngtr ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8816).ToObject(), πg.NewUnicode("\\nleq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8817).ToObject(), πg.NewUnicode("\\ngeq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8818).ToObject(), πg.NewUnicode("\\lesssim ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8819).ToObject(), πg.NewUnicode("\\gtrsim ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8822).ToObject(), πg.NewUnicode("\\lessgtr ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8823).ToObject(), πg.NewUnicode("\\gtrless ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8826).ToObject(), πg.NewUnicode("\\prec ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8827).ToObject(), πg.NewUnicode("\\succ ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8828).ToObject(), πg.NewUnicode("\\preccurlyeq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8829).ToObject(), πg.NewUnicode("\\succcurlyeq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8830).ToObject(), πg.NewUnicode("\\precsim ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8831).ToObject(), πg.NewUnicode("\\succsim ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8832).ToObject(), πg.NewUnicode("\\nprec ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8833).ToObject(), πg.NewUnicode("\\nsucc ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8834).ToObject(), πg.NewUnicode("\\subset ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8835).ToObject(), πg.NewUnicode("\\supset ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8838).ToObject(), πg.NewUnicode("\\subseteq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8839).ToObject(), πg.NewUnicode("\\supseteq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8840).ToObject(), πg.NewUnicode("\\nsubseteq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8841).ToObject(), πg.NewUnicode("\\nsupseteq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8842).ToObject(), πg.NewUnicode("\\subsetneq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8843).ToObject(), πg.NewUnicode("\\supsetneq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8846).ToObject(), πg.NewUnicode("\\uplus ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8847).ToObject(), πg.NewUnicode("\\sqsubset ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8848).ToObject(), πg.NewUnicode("\\sqsupset ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8849).ToObject(), πg.NewUnicode("\\sqsubseteq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8850).ToObject(), πg.NewUnicode("\\sqsupseteq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8851).ToObject(), πg.NewUnicode("\\sqcap ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8852).ToObject(), πg.NewUnicode("\\sqcup ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8853).ToObject(), πg.NewUnicode("\\oplus ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8854).ToObject(), πg.NewUnicode("\\ominus ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8855).ToObject(), πg.NewUnicode("\\otimes ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8856).ToObject(), πg.NewUnicode("\\oslash ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8857).ToObject(), πg.NewUnicode("\\odot ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8858).ToObject(), πg.NewUnicode("\\circledcirc ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8859).ToObject(), πg.NewUnicode("\\circledast ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8861).ToObject(), πg.NewUnicode("\\circleddash ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8862).ToObject(), πg.NewUnicode("\\boxplus ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8863).ToObject(), πg.NewUnicode("\\boxminus ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8864).ToObject(), πg.NewUnicode("\\boxtimes ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8865).ToObject(), πg.NewUnicode("\\boxdot ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8866).ToObject(), πg.NewUnicode("\\vdash ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8867).ToObject(), πg.NewUnicode("\\dashv ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8868).ToObject(), πg.NewUnicode("\\top ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8869).ToObject(), πg.NewUnicode("\\bot ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8871).ToObject(), πg.NewUnicode("\\models ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8872).ToObject(), πg.NewUnicode("\\vDash ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8873).ToObject(), πg.NewUnicode("\\Vdash ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8874).ToObject(), πg.NewUnicode("\\Vvdash ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8876).ToObject(), πg.NewUnicode("\\nvdash ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8877).ToObject(), πg.NewUnicode("\\nvDash ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8878).ToObject(), πg.NewUnicode("\\nVdash ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8879).ToObject(), πg.NewUnicode("\\nVDash ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8882).ToObject(), πg.NewUnicode("\\vartriangleleft ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8883).ToObject(), πg.NewUnicode("\\vartriangleright ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8884).ToObject(), πg.NewUnicode("\\trianglelefteq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8885).ToObject(), πg.NewUnicode("\\trianglerighteq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8888).ToObject(), πg.NewUnicode("\\multimap ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8890).ToObject(), πg.NewUnicode("\\intercal ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8891).ToObject(), πg.NewUnicode("\\veebar ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8892).ToObject(), πg.NewUnicode("\\barwedge ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8896).ToObject(), πg.NewUnicode("\\bigwedge ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8897).ToObject(), πg.NewUnicode("\\bigvee ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8898).ToObject(), πg.NewUnicode("\\bigcap ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8899).ToObject(), πg.NewUnicode("\\bigcup ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8900).ToObject(), πg.NewUnicode("\\diamond ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8901).ToObject(), πg.NewUnicode("\\cdot ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8902).ToObject(), πg.NewUnicode("\\star ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8903).ToObject(), πg.NewUnicode("\\divideontimes ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8904).ToObject(), πg.NewUnicode("\\bowtie ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8905).ToObject(), πg.NewUnicode("\\ltimes ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8906).ToObject(), πg.NewUnicode("\\rtimes ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8907).ToObject(), πg.NewUnicode("\\leftthreetimes ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8908).ToObject(), πg.NewUnicode("\\rightthreetimes ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8909).ToObject(), πg.NewUnicode("\\backsimeq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8910).ToObject(), πg.NewUnicode("\\curlyvee ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8911).ToObject(), πg.NewUnicode("\\curlywedge ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8912).ToObject(), πg.NewUnicode("\\Subset ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8913).ToObject(), πg.NewUnicode("\\Supset ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8914).ToObject(), πg.NewUnicode("\\Cap ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8915).ToObject(), πg.NewUnicode("\\Cup ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8916).ToObject(), πg.NewUnicode("\\pitchfork ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8918).ToObject(), πg.NewUnicode("\\lessdot ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8919).ToObject(), πg.NewUnicode("\\gtrdot ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8920).ToObject(), πg.NewUnicode("\\lll ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8921).ToObject(), πg.NewUnicode("\\ggg ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8922).ToObject(), πg.NewUnicode("\\lesseqgtr ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8923).ToObject(), πg.NewUnicode("\\gtreqless ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8926).ToObject(), πg.NewUnicode("\\curlyeqprec ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8927).ToObject(), πg.NewUnicode("\\curlyeqsucc ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8928).ToObject(), πg.NewUnicode("\\npreceq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8929).ToObject(), πg.NewUnicode("\\nsucceq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8934).ToObject(), πg.NewUnicode("\\lnsim ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8935).ToObject(), πg.NewUnicode("\\gnsim ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8936).ToObject(), πg.NewUnicode("\\precnsim ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8937).ToObject(), πg.NewUnicode("\\succnsim ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8938).ToObject(), πg.NewUnicode("\\ntriangleleft ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8939).ToObject(), πg.NewUnicode("\\ntriangleright ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8940).ToObject(), πg.NewUnicode("\\ntrianglelefteq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8941).ToObject(), πg.NewUnicode("\\ntrianglerighteq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8942).ToObject(), πg.NewUnicode("\\vdots ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8943).ToObject(), πg.NewUnicode("\\cdots ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8945).ToObject(), πg.NewUnicode("\\ddots ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8968).ToObject(), πg.NewUnicode("\\lceil ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8969).ToObject(), πg.NewUnicode("\\rceil ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8970).ToObject(), πg.NewUnicode("\\lfloor ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8971).ToObject(), πg.NewUnicode("\\rfloor ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8988).ToObject(), πg.NewUnicode("\\ulcorner ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8989).ToObject(), πg.NewUnicode("\\urcorner ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8990).ToObject(), πg.NewUnicode("\\llcorner ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8991).ToObject(), πg.NewUnicode("\\lrcorner ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8994).ToObject(), πg.NewUnicode("\\frown ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(8995).ToObject(), πg.NewUnicode("\\smile ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(9182).ToObject(), πg.NewUnicode("\\overbrace ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(9183).ToObject(), πg.NewUnicode("\\underbrace ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(9651).ToObject(), πg.NewUnicode("\\bigtriangleup ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(9655).ToObject(), πg.NewUnicode("\\rhd ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(9661).ToObject(), πg.NewUnicode("\\bigtriangledown ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(9665).ToObject(), πg.NewUnicode("\\lhd ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(9671).ToObject(), πg.NewUnicode("\\Diamond ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(9674).ToObject(), πg.NewUnicode("\\lozenge ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(9723).ToObject(), πg.NewUnicode("\\square ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(9724).ToObject(), πg.NewUnicode("\\blacksquare ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(9733).ToObject(), πg.NewUnicode("\\bigstar ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(9824).ToObject(), πg.NewUnicode("\\spadesuit ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(9825).ToObject(), πg.NewUnicode("\\heartsuit ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(9826).ToObject(), πg.NewUnicode("\\diamondsuit ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(9827).ToObject(), πg.NewUnicode("\\clubsuit ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(9837).ToObject(), πg.NewUnicode("\\flat ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(9838).ToObject(), πg.NewUnicode("\\natural ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(9839).ToObject(), πg.NewUnicode("\\sharp ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10003).ToObject(), πg.NewUnicode("\\checkmark ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10016).ToObject(), πg.NewUnicode("\\maltese ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10178).ToObject(), πg.NewUnicode("\\perp ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10216).ToObject(), πg.NewUnicode("\\langle ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10217).ToObject(), πg.NewUnicode("\\rangle ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10222).ToObject(), πg.NewUnicode("\\lgroup ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10223).ToObject(), πg.NewUnicode("\\rgroup ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10229).ToObject(), πg.NewUnicode("\\longleftarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10230).ToObject(), πg.NewUnicode("\\longrightarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10231).ToObject(), πg.NewUnicode("\\longleftrightarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10232).ToObject(), πg.NewUnicode("\\Longleftarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10233).ToObject(), πg.NewUnicode("\\Longrightarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10234).ToObject(), πg.NewUnicode("\\Longleftrightarrow ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10236).ToObject(), πg.NewUnicode("\\longmapsto ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10731).ToObject(), πg.NewUnicode("\\blacklozenge ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10741).ToObject(), πg.NewUnicode("\\setminus ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10752).ToObject(), πg.NewUnicode("\\bigodot ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10753).ToObject(), πg.NewUnicode("\\bigoplus ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10754).ToObject(), πg.NewUnicode("\\bigotimes ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10756).ToObject(), πg.NewUnicode("\\biguplus ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10758).ToObject(), πg.NewUnicode("\\bigsqcup ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10764).ToObject(), πg.NewUnicode("\\iiiint ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10781).ToObject(), πg.NewUnicode("\\Join ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10815).ToObject(), πg.NewUnicode("\\amalg ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10846).ToObject(), πg.NewUnicode("\\doublebarwedge ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10877).ToObject(), πg.NewUnicode("\\leqslant ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10878).ToObject(), πg.NewUnicode("\\geqslant ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10885).ToObject(), πg.NewUnicode("\\lessapprox ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10886).ToObject(), πg.NewUnicode("\\gtrapprox ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10887).ToObject(), πg.NewUnicode("\\lneq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10888).ToObject(), πg.NewUnicode("\\gneq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10889).ToObject(), πg.NewUnicode("\\lnapprox ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10890).ToObject(), πg.NewUnicode("\\gnapprox ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10891).ToObject(), πg.NewUnicode("\\lesseqqgtr ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10892).ToObject(), πg.NewUnicode("\\gtreqqless ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10901).ToObject(), πg.NewUnicode("\\eqslantless ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10902).ToObject(), πg.NewUnicode("\\eqslantgtr ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10927).ToObject(), πg.NewUnicode("\\preceq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10928).ToObject(), πg.NewUnicode("\\succeq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10935).ToObject(), πg.NewUnicode("\\precapprox ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10936).ToObject(), πg.NewUnicode("\\succapprox ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10937).ToObject(), πg.NewUnicode("\\precnapprox ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10938).ToObject(), πg.NewUnicode("\\succnapprox ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10949).ToObject(), πg.NewUnicode("\\subseteqq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10950).ToObject(), πg.NewUnicode("\\supseteqq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10955).ToObject(), πg.NewUnicode("\\subsetneqq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(10956).ToObject(), πg.NewUnicode("\\supsetneqq ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119808).ToObject(), πg.NewUnicode("\\mathbf{A}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119809).ToObject(), πg.NewUnicode("\\mathbf{B}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119810).ToObject(), πg.NewUnicode("\\mathbf{C}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119811).ToObject(), πg.NewUnicode("\\mathbf{D}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119812).ToObject(), πg.NewUnicode("\\mathbf{E}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119813).ToObject(), πg.NewUnicode("\\mathbf{F}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119814).ToObject(), πg.NewUnicode("\\mathbf{G}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119815).ToObject(), πg.NewUnicode("\\mathbf{H}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119816).ToObject(), πg.NewUnicode("\\mathbf{I}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119817).ToObject(), πg.NewUnicode("\\mathbf{J}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119818).ToObject(), πg.NewUnicode("\\mathbf{K}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119819).ToObject(), πg.NewUnicode("\\mathbf{L}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119820).ToObject(), πg.NewUnicode("\\mathbf{M}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119821).ToObject(), πg.NewUnicode("\\mathbf{N}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119822).ToObject(), πg.NewUnicode("\\mathbf{O}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119823).ToObject(), πg.NewUnicode("\\mathbf{P}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119824).ToObject(), πg.NewUnicode("\\mathbf{Q}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119825).ToObject(), πg.NewUnicode("\\mathbf{R}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119826).ToObject(), πg.NewUnicode("\\mathbf{S}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119827).ToObject(), πg.NewUnicode("\\mathbf{T}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119828).ToObject(), πg.NewUnicode("\\mathbf{U}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119829).ToObject(), πg.NewUnicode("\\mathbf{V}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119830).ToObject(), πg.NewUnicode("\\mathbf{W}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119831).ToObject(), πg.NewUnicode("\\mathbf{X}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119832).ToObject(), πg.NewUnicode("\\mathbf{Y}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119833).ToObject(), πg.NewUnicode("\\mathbf{Z}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119834).ToObject(), πg.NewUnicode("\\mathbf{a}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119835).ToObject(), πg.NewUnicode("\\mathbf{b}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119836).ToObject(), πg.NewUnicode("\\mathbf{c}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119837).ToObject(), πg.NewUnicode("\\mathbf{d}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119838).ToObject(), πg.NewUnicode("\\mathbf{e}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119839).ToObject(), πg.NewUnicode("\\mathbf{f}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119840).ToObject(), πg.NewUnicode("\\mathbf{g}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119841).ToObject(), πg.NewUnicode("\\mathbf{h}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119842).ToObject(), πg.NewUnicode("\\mathbf{i}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119843).ToObject(), πg.NewUnicode("\\mathbf{j}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119844).ToObject(), πg.NewUnicode("\\mathbf{k}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119845).ToObject(), πg.NewUnicode("\\mathbf{l}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119846).ToObject(), πg.NewUnicode("\\mathbf{m}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119847).ToObject(), πg.NewUnicode("\\mathbf{n}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119848).ToObject(), πg.NewUnicode("\\mathbf{o}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119849).ToObject(), πg.NewUnicode("\\mathbf{p}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119850).ToObject(), πg.NewUnicode("\\mathbf{q}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119851).ToObject(), πg.NewUnicode("\\mathbf{r}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119852).ToObject(), πg.NewUnicode("\\mathbf{s}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119853).ToObject(), πg.NewUnicode("\\mathbf{t}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119854).ToObject(), πg.NewUnicode("\\mathbf{u}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119855).ToObject(), πg.NewUnicode("\\mathbf{v}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119856).ToObject(), πg.NewUnicode("\\mathbf{w}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119857).ToObject(), πg.NewUnicode("\\mathbf{x}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119858).ToObject(), πg.NewUnicode("\\mathbf{y}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119859).ToObject(), πg.NewUnicode("\\mathbf{z}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119860).ToObject(), πg.NewUnicode("A").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119861).ToObject(), πg.NewUnicode("B").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119862).ToObject(), πg.NewUnicode("C").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119863).ToObject(), πg.NewUnicode("D").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119864).ToObject(), πg.NewUnicode("E").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119865).ToObject(), πg.NewUnicode("F").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119866).ToObject(), πg.NewUnicode("G").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119867).ToObject(), πg.NewUnicode("H").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119868).ToObject(), πg.NewUnicode("I").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119869).ToObject(), πg.NewUnicode("J").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119870).ToObject(), πg.NewUnicode("K").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119871).ToObject(), πg.NewUnicode("L").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119872).ToObject(), πg.NewUnicode("M").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119873).ToObject(), πg.NewUnicode("N").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119874).ToObject(), πg.NewUnicode("O").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119875).ToObject(), πg.NewUnicode("P").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119876).ToObject(), πg.NewUnicode("Q").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119877).ToObject(), πg.NewUnicode("R").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119878).ToObject(), πg.NewUnicode("S").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119879).ToObject(), πg.NewUnicode("T").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119880).ToObject(), πg.NewUnicode("U").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119881).ToObject(), πg.NewUnicode("V").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119882).ToObject(), πg.NewUnicode("W").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119883).ToObject(), πg.NewUnicode("X").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119884).ToObject(), πg.NewUnicode("Y").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119885).ToObject(), πg.NewUnicode("Z").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119886).ToObject(), πg.NewUnicode("a").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119887).ToObject(), πg.NewUnicode("b").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119888).ToObject(), πg.NewUnicode("c").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119889).ToObject(), πg.NewUnicode("d").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119890).ToObject(), πg.NewUnicode("e").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119891).ToObject(), πg.NewUnicode("f").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119892).ToObject(), πg.NewUnicode("g").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119894).ToObject(), πg.NewUnicode("i").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119895).ToObject(), πg.NewUnicode("j").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119896).ToObject(), πg.NewUnicode("k").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119897).ToObject(), πg.NewUnicode("l").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119898).ToObject(), πg.NewUnicode("m").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119899).ToObject(), πg.NewUnicode("n").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119900).ToObject(), πg.NewUnicode("o").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119901).ToObject(), πg.NewUnicode("p").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119902).ToObject(), πg.NewUnicode("q").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119903).ToObject(), πg.NewUnicode("r").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119904).ToObject(), πg.NewUnicode("s").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119905).ToObject(), πg.NewUnicode("t").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119906).ToObject(), πg.NewUnicode("u").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119907).ToObject(), πg.NewUnicode("v").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119908).ToObject(), πg.NewUnicode("w").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119909).ToObject(), πg.NewUnicode("x").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119910).ToObject(), πg.NewUnicode("y").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119911).ToObject(), πg.NewUnicode("z").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119964).ToObject(), πg.NewUnicode("\\mathcal{A}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119966).ToObject(), πg.NewUnicode("\\mathcal{C}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119967).ToObject(), πg.NewUnicode("\\mathcal{D}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119970).ToObject(), πg.NewUnicode("\\mathcal{G}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119973).ToObject(), πg.NewUnicode("\\mathcal{J}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119974).ToObject(), πg.NewUnicode("\\mathcal{K}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119977).ToObject(), πg.NewUnicode("\\mathcal{N}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119978).ToObject(), πg.NewUnicode("\\mathcal{O}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119979).ToObject(), πg.NewUnicode("\\mathcal{P}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119980).ToObject(), πg.NewUnicode("\\mathcal{Q}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119982).ToObject(), πg.NewUnicode("\\mathcal{S}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119983).ToObject(), πg.NewUnicode("\\mathcal{T}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119984).ToObject(), πg.NewUnicode("\\mathcal{U}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119985).ToObject(), πg.NewUnicode("\\mathcal{V}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119986).ToObject(), πg.NewUnicode("\\mathcal{W}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119987).ToObject(), πg.NewUnicode("\\mathcal{X}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119988).ToObject(), πg.NewUnicode("\\mathcal{Y}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(119989).ToObject(), πg.NewUnicode("\\mathcal{Z}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120068).ToObject(), πg.NewUnicode("\\mathfrak{A}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120069).ToObject(), πg.NewUnicode("\\mathfrak{B}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120071).ToObject(), πg.NewUnicode("\\mathfrak{D}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120072).ToObject(), πg.NewUnicode("\\mathfrak{E}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120073).ToObject(), πg.NewUnicode("\\mathfrak{F}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120074).ToObject(), πg.NewUnicode("\\mathfrak{G}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120077).ToObject(), πg.NewUnicode("\\mathfrak{J}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120078).ToObject(), πg.NewUnicode("\\mathfrak{K}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120079).ToObject(), πg.NewUnicode("\\mathfrak{L}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120080).ToObject(), πg.NewUnicode("\\mathfrak{M}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120081).ToObject(), πg.NewUnicode("\\mathfrak{N}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120082).ToObject(), πg.NewUnicode("\\mathfrak{O}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120083).ToObject(), πg.NewUnicode("\\mathfrak{P}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120084).ToObject(), πg.NewUnicode("\\mathfrak{Q}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120086).ToObject(), πg.NewUnicode("\\mathfrak{S}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120087).ToObject(), πg.NewUnicode("\\mathfrak{T}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120088).ToObject(), πg.NewUnicode("\\mathfrak{U}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120089).ToObject(), πg.NewUnicode("\\mathfrak{V}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120090).ToObject(), πg.NewUnicode("\\mathfrak{W}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120091).ToObject(), πg.NewUnicode("\\mathfrak{X}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120092).ToObject(), πg.NewUnicode("\\mathfrak{Y}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120094).ToObject(), πg.NewUnicode("\\mathfrak{a}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120095).ToObject(), πg.NewUnicode("\\mathfrak{b}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120096).ToObject(), πg.NewUnicode("\\mathfrak{c}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120097).ToObject(), πg.NewUnicode("\\mathfrak{d}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120098).ToObject(), πg.NewUnicode("\\mathfrak{e}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120099).ToObject(), πg.NewUnicode("\\mathfrak{f}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120100).ToObject(), πg.NewUnicode("\\mathfrak{g}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120101).ToObject(), πg.NewUnicode("\\mathfrak{h}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120102).ToObject(), πg.NewUnicode("\\mathfrak{i}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120103).ToObject(), πg.NewUnicode("\\mathfrak{j}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120104).ToObject(), πg.NewUnicode("\\mathfrak{k}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120105).ToObject(), πg.NewUnicode("\\mathfrak{l}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120106).ToObject(), πg.NewUnicode("\\mathfrak{m}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120107).ToObject(), πg.NewUnicode("\\mathfrak{n}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120108).ToObject(), πg.NewUnicode("\\mathfrak{o}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120109).ToObject(), πg.NewUnicode("\\mathfrak{p}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120110).ToObject(), πg.NewUnicode("\\mathfrak{q}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120111).ToObject(), πg.NewUnicode("\\mathfrak{r}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120112).ToObject(), πg.NewUnicode("\\mathfrak{s}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120113).ToObject(), πg.NewUnicode("\\mathfrak{t}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120114).ToObject(), πg.NewUnicode("\\mathfrak{u}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120115).ToObject(), πg.NewUnicode("\\mathfrak{v}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120116).ToObject(), πg.NewUnicode("\\mathfrak{w}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120117).ToObject(), πg.NewUnicode("\\mathfrak{x}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120118).ToObject(), πg.NewUnicode("\\mathfrak{y}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120119).ToObject(), πg.NewUnicode("\\mathfrak{z}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120120).ToObject(), πg.NewUnicode("\\mathbb{A}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120121).ToObject(), πg.NewUnicode("\\mathbb{B}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120123).ToObject(), πg.NewUnicode("\\mathbb{D}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120124).ToObject(), πg.NewUnicode("\\mathbb{E}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120125).ToObject(), πg.NewUnicode("\\mathbb{F}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120126).ToObject(), πg.NewUnicode("\\mathbb{G}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120128).ToObject(), πg.NewUnicode("\\mathbb{I}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120129).ToObject(), πg.NewUnicode("\\mathbb{J}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120130).ToObject(), πg.NewUnicode("\\mathbb{K}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120131).ToObject(), πg.NewUnicode("\\mathbb{L}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120132).ToObject(), πg.NewUnicode("\\mathbb{M}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120134).ToObject(), πg.NewUnicode("\\mathbb{O}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120138).ToObject(), πg.NewUnicode("\\mathbb{S}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120139).ToObject(), πg.NewUnicode("\\mathbb{T}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120140).ToObject(), πg.NewUnicode("\\mathbb{U}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120141).ToObject(), πg.NewUnicode("\\mathbb{V}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120142).ToObject(), πg.NewUnicode("\\mathbb{W}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120143).ToObject(), πg.NewUnicode("\\mathbb{X}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120144).ToObject(), πg.NewUnicode("\\mathbb{Y}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120156).ToObject(), πg.NewUnicode("\\Bbbk ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120224).ToObject(), πg.NewUnicode("\\mathsf{A}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120225).ToObject(), πg.NewUnicode("\\mathsf{B}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120226).ToObject(), πg.NewUnicode("\\mathsf{C}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120227).ToObject(), πg.NewUnicode("\\mathsf{D}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120228).ToObject(), πg.NewUnicode("\\mathsf{E}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120229).ToObject(), πg.NewUnicode("\\mathsf{F}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120230).ToObject(), πg.NewUnicode("\\mathsf{G}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120231).ToObject(), πg.NewUnicode("\\mathsf{H}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120232).ToObject(), πg.NewUnicode("\\mathsf{I}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120233).ToObject(), πg.NewUnicode("\\mathsf{J}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120234).ToObject(), πg.NewUnicode("\\mathsf{K}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120235).ToObject(), πg.NewUnicode("\\mathsf{L}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120236).ToObject(), πg.NewUnicode("\\mathsf{M}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120237).ToObject(), πg.NewUnicode("\\mathsf{N}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120238).ToObject(), πg.NewUnicode("\\mathsf{O}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120239).ToObject(), πg.NewUnicode("\\mathsf{P}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120240).ToObject(), πg.NewUnicode("\\mathsf{Q}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120241).ToObject(), πg.NewUnicode("\\mathsf{R}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120242).ToObject(), πg.NewUnicode("\\mathsf{S}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120243).ToObject(), πg.NewUnicode("\\mathsf{T}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120244).ToObject(), πg.NewUnicode("\\mathsf{U}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120245).ToObject(), πg.NewUnicode("\\mathsf{V}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120246).ToObject(), πg.NewUnicode("\\mathsf{W}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120247).ToObject(), πg.NewUnicode("\\mathsf{X}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120248).ToObject(), πg.NewUnicode("\\mathsf{Y}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120249).ToObject(), πg.NewUnicode("\\mathsf{Z}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120250).ToObject(), πg.NewUnicode("\\mathsf{a}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120251).ToObject(), πg.NewUnicode("\\mathsf{b}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120252).ToObject(), πg.NewUnicode("\\mathsf{c}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120253).ToObject(), πg.NewUnicode("\\mathsf{d}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120254).ToObject(), πg.NewUnicode("\\mathsf{e}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120255).ToObject(), πg.NewUnicode("\\mathsf{f}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120256).ToObject(), πg.NewUnicode("\\mathsf{g}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120257).ToObject(), πg.NewUnicode("\\mathsf{h}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120258).ToObject(), πg.NewUnicode("\\mathsf{i}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120259).ToObject(), πg.NewUnicode("\\mathsf{j}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120260).ToObject(), πg.NewUnicode("\\mathsf{k}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120261).ToObject(), πg.NewUnicode("\\mathsf{l}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120262).ToObject(), πg.NewUnicode("\\mathsf{m}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120263).ToObject(), πg.NewUnicode("\\mathsf{n}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120264).ToObject(), πg.NewUnicode("\\mathsf{o}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120265).ToObject(), πg.NewUnicode("\\mathsf{p}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120266).ToObject(), πg.NewUnicode("\\mathsf{q}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120267).ToObject(), πg.NewUnicode("\\mathsf{r}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120268).ToObject(), πg.NewUnicode("\\mathsf{s}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120269).ToObject(), πg.NewUnicode("\\mathsf{t}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120270).ToObject(), πg.NewUnicode("\\mathsf{u}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120271).ToObject(), πg.NewUnicode("\\mathsf{v}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120272).ToObject(), πg.NewUnicode("\\mathsf{w}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120273).ToObject(), πg.NewUnicode("\\mathsf{x}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120274).ToObject(), πg.NewUnicode("\\mathsf{y}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120275).ToObject(), πg.NewUnicode("\\mathsf{z}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120432).ToObject(), πg.NewUnicode("\\mathtt{A}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120433).ToObject(), πg.NewUnicode("\\mathtt{B}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120434).ToObject(), πg.NewUnicode("\\mathtt{C}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120435).ToObject(), πg.NewUnicode("\\mathtt{D}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120436).ToObject(), πg.NewUnicode("\\mathtt{E}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120437).ToObject(), πg.NewUnicode("\\mathtt{F}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120438).ToObject(), πg.NewUnicode("\\mathtt{G}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120439).ToObject(), πg.NewUnicode("\\mathtt{H}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120440).ToObject(), πg.NewUnicode("\\mathtt{I}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120441).ToObject(), πg.NewUnicode("\\mathtt{J}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120442).ToObject(), πg.NewUnicode("\\mathtt{K}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120443).ToObject(), πg.NewUnicode("\\mathtt{L}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120444).ToObject(), πg.NewUnicode("\\mathtt{M}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120445).ToObject(), πg.NewUnicode("\\mathtt{N}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120446).ToObject(), πg.NewUnicode("\\mathtt{O}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120447).ToObject(), πg.NewUnicode("\\mathtt{P}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120448).ToObject(), πg.NewUnicode("\\mathtt{Q}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120449).ToObject(), πg.NewUnicode("\\mathtt{R}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120450).ToObject(), πg.NewUnicode("\\mathtt{S}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120451).ToObject(), πg.NewUnicode("\\mathtt{T}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120452).ToObject(), πg.NewUnicode("\\mathtt{U}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120453).ToObject(), πg.NewUnicode("\\mathtt{V}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120454).ToObject(), πg.NewUnicode("\\mathtt{W}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120455).ToObject(), πg.NewUnicode("\\mathtt{X}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120456).ToObject(), πg.NewUnicode("\\mathtt{Y}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120457).ToObject(), πg.NewUnicode("\\mathtt{Z}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120458).ToObject(), πg.NewUnicode("\\mathtt{a}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120459).ToObject(), πg.NewUnicode("\\mathtt{b}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120460).ToObject(), πg.NewUnicode("\\mathtt{c}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120461).ToObject(), πg.NewUnicode("\\mathtt{d}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120462).ToObject(), πg.NewUnicode("\\mathtt{e}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120463).ToObject(), πg.NewUnicode("\\mathtt{f}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120464).ToObject(), πg.NewUnicode("\\mathtt{g}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120465).ToObject(), πg.NewUnicode("\\mathtt{h}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120466).ToObject(), πg.NewUnicode("\\mathtt{i}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120467).ToObject(), πg.NewUnicode("\\mathtt{j}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120468).ToObject(), πg.NewUnicode("\\mathtt{k}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120469).ToObject(), πg.NewUnicode("\\mathtt{l}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120470).ToObject(), πg.NewUnicode("\\mathtt{m}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120471).ToObject(), πg.NewUnicode("\\mathtt{n}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120472).ToObject(), πg.NewUnicode("\\mathtt{o}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120473).ToObject(), πg.NewUnicode("\\mathtt{p}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120474).ToObject(), πg.NewUnicode("\\mathtt{q}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120475).ToObject(), πg.NewUnicode("\\mathtt{r}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120476).ToObject(), πg.NewUnicode("\\mathtt{s}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120477).ToObject(), πg.NewUnicode("\\mathtt{t}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120478).ToObject(), πg.NewUnicode("\\mathtt{u}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120479).ToObject(), πg.NewUnicode("\\mathtt{v}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120480).ToObject(), πg.NewUnicode("\\mathtt{w}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120481).ToObject(), πg.NewUnicode("\\mathtt{x}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120482).ToObject(), πg.NewUnicode("\\mathtt{y}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120483).ToObject(), πg.NewUnicode("\\mathtt{z}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120484).ToObject(), πg.NewUnicode("\\imath ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120485).ToObject(), πg.NewUnicode("\\jmath ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120490).ToObject(), πg.NewUnicode("\\mathbf{\\Gamma}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120491).ToObject(), πg.NewUnicode("\\mathbf{\\Delta}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120495).ToObject(), πg.NewUnicode("\\mathbf{\\Theta}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120498).ToObject(), πg.NewUnicode("\\mathbf{\\Lambda}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120501).ToObject(), πg.NewUnicode("\\mathbf{\\Xi}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120503).ToObject(), πg.NewUnicode("\\mathbf{\\Pi}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120506).ToObject(), πg.NewUnicode("\\mathbf{\\Sigma}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120508).ToObject(), πg.NewUnicode("\\mathbf{\\Upsilon}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120509).ToObject(), πg.NewUnicode("\\mathbf{\\Phi}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120511).ToObject(), πg.NewUnicode("\\mathbf{\\Psi}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120512).ToObject(), πg.NewUnicode("\\mathbf{\\Omega}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120548).ToObject(), πg.NewUnicode("\\mathit{\\Gamma}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120549).ToObject(), πg.NewUnicode("\\mathit{\\Delta}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120553).ToObject(), πg.NewUnicode("\\mathit{\\Theta}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120556).ToObject(), πg.NewUnicode("\\mathit{\\Lambda}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120559).ToObject(), πg.NewUnicode("\\mathit{\\Xi}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120561).ToObject(), πg.NewUnicode("\\mathit{\\Pi}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120564).ToObject(), πg.NewUnicode("\\mathit{\\Sigma}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120566).ToObject(), πg.NewUnicode("\\mathit{\\Upsilon}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120567).ToObject(), πg.NewUnicode("\\mathit{\\Phi}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120569).ToObject(), πg.NewUnicode("\\mathit{\\Psi}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120570).ToObject(), πg.NewUnicode("\\mathit{\\Omega}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120572).ToObject(), πg.NewUnicode("\\alpha ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120573).ToObject(), πg.NewUnicode("\\beta ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120574).ToObject(), πg.NewUnicode("\\gamma ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120575).ToObject(), πg.NewUnicode("\\delta ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120576).ToObject(), πg.NewUnicode("\\varepsilon ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120577).ToObject(), πg.NewUnicode("\\zeta ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120578).ToObject(), πg.NewUnicode("\\eta ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120579).ToObject(), πg.NewUnicode("\\theta ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120580).ToObject(), πg.NewUnicode("\\iota ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120581).ToObject(), πg.NewUnicode("\\kappa ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120582).ToObject(), πg.NewUnicode("\\lambda ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120583).ToObject(), πg.NewUnicode("\\mu ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120584).ToObject(), πg.NewUnicode("\\nu ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120585).ToObject(), πg.NewUnicode("\\xi ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120587).ToObject(), πg.NewUnicode("\\pi ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120588).ToObject(), πg.NewUnicode("\\rho ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120589).ToObject(), πg.NewUnicode("\\varsigma ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120590).ToObject(), πg.NewUnicode("\\sigma ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120591).ToObject(), πg.NewUnicode("\\tau ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120592).ToObject(), πg.NewUnicode("\\upsilon ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120593).ToObject(), πg.NewUnicode("\\varphi ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120594).ToObject(), πg.NewUnicode("\\chi ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120595).ToObject(), πg.NewUnicode("\\psi ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120596).ToObject(), πg.NewUnicode("\\omega ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120597).ToObject(), πg.NewUnicode("\\partial ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120598).ToObject(), πg.NewUnicode("\\epsilon ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120599).ToObject(), πg.NewUnicode("\\vartheta ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120600).ToObject(), πg.NewUnicode("\\varkappa ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120601).ToObject(), πg.NewUnicode("\\phi ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120602).ToObject(), πg.NewUnicode("\\varrho ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120603).ToObject(), πg.NewUnicode("\\varpi ").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120782).ToObject(), πg.NewUnicode("\\mathbf{0}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120783).ToObject(), πg.NewUnicode("\\mathbf{1}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120784).ToObject(), πg.NewUnicode("\\mathbf{2}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120785).ToObject(), πg.NewUnicode("\\mathbf{3}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120786).ToObject(), πg.NewUnicode("\\mathbf{4}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120787).ToObject(), πg.NewUnicode("\\mathbf{5}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120788).ToObject(), πg.NewUnicode("\\mathbf{6}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120789).ToObject(), πg.NewUnicode("\\mathbf{7}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120790).ToObject(), πg.NewUnicode("\\mathbf{8}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120791).ToObject(), πg.NewUnicode("\\mathbf{9}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120802).ToObject(), πg.NewUnicode("\\mathsf{0}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120803).ToObject(), πg.NewUnicode("\\mathsf{1}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120804).ToObject(), πg.NewUnicode("\\mathsf{2}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120805).ToObject(), πg.NewUnicode("\\mathsf{3}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120806).ToObject(), πg.NewUnicode("\\mathsf{4}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120807).ToObject(), πg.NewUnicode("\\mathsf{5}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120808).ToObject(), πg.NewUnicode("\\mathsf{6}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120809).ToObject(), πg.NewUnicode("\\mathsf{7}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120810).ToObject(), πg.NewUnicode("\\mathsf{8}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120811).ToObject(), πg.NewUnicode("\\mathsf{9}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120822).ToObject(), πg.NewUnicode("\\mathtt{0}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120823).ToObject(), πg.NewUnicode("\\mathtt{1}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120824).ToObject(), πg.NewUnicode("\\mathtt{2}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120825).ToObject(), πg.NewUnicode("\\mathtt{3}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120826).ToObject(), πg.NewUnicode("\\mathtt{4}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120827).ToObject(), πg.NewUnicode("\\mathtt{5}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120828).ToObject(), πg.NewUnicode("\\mathtt{6}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120829).ToObject(), πg.NewUnicode("\\mathtt{7}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120830).ToObject(), πg.NewUnicode("\\mathtt{8}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewInt(120831).ToObject(), πg.NewUnicode("\\mathtt{9}").ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßuni2tex_table.ToObject(), πTemp002); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("unichar2tex", Code)
}
