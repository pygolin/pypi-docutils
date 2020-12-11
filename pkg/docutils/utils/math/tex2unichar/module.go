package tex2unichar

import (
	πg "github.com/pygolin/runtime"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/utils/math/tex2unichar.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ßAC := πg.InternStr("AC")
		ßAPLcomment := πg.InternStr("APLcomment")
		ßAPLdownarrowbox := πg.InternStr("APLdownarrowbox")
		ßAPLinput := πg.InternStr("APLinput")
		ßAPLinv := πg.InternStr("APLinv")
		ßAPLleftarrowbox := πg.InternStr("APLleftarrowbox")
		ßAPLlog := πg.InternStr("APLlog")
		ßAPLrightarrowbox := πg.InternStr("APLrightarrowbox")
		ßAPLuparrowbox := πg.InternStr("APLuparrowbox")
		ßAries := πg.InternStr("Aries")
		ßBbbk := πg.InternStr("Bbbk")
		ßBumpeq := πg.InternStr("Bumpeq")
		ßCIRCLE := πg.InternStr("CIRCLE")
		ßCap := πg.InternStr("Cap")
		ßCheckedBox := πg.InternStr("CheckedBox")
		ßCircle := πg.InternStr("Circle")
		ßCup := πg.InternStr("Cup")
		ßDelta := πg.InternStr("Delta")
		ßDiamond := πg.InternStr("Diamond")
		ßDoteq := πg.InternStr("Doteq")
		ßDownarrow := πg.InternStr("Downarrow")
		ßFinv := πg.InternStr("Finv")
		ßGame := πg.InternStr("Game")
		ßGamma := πg.InternStr("Gamma")
		ßGemini := πg.InternStr("Gemini")
		ßIm := πg.InternStr("Im")
		ßJoin := πg.InternStr("Join")
		ßJupiter := πg.InternStr("Jupiter")
		ßLEFTCIRCLE := πg.InternStr("LEFTCIRCLE")
		ßLEFTcircle := πg.InternStr("LEFTcircle")
		ßLHD := πg.InternStr("LHD")
		ßLambda := πg.InternStr("Lambda")
		ßLbag := πg.InternStr("Lbag")
		ßLeftarrow := πg.InternStr("Leftarrow")
		ßLeftrightarrow := πg.InternStr("Leftrightarrow")
		ßLeo := πg.InternStr("Leo")
		ßLibra := πg.InternStr("Libra")
		ßLleftarrow := πg.InternStr("Lleftarrow")
		ßLongleftarrow := πg.InternStr("Longleftarrow")
		ßLongleftrightarrow := πg.InternStr("Longleftrightarrow")
		ßLongmapsfrom := πg.InternStr("Longmapsfrom")
		ßLongmapsto := πg.InternStr("Longmapsto")
		ßLongrightarrow := πg.InternStr("Longrightarrow")
		ßLsh := πg.InternStr("Lsh")
		ßMapsfrom := πg.InternStr("Mapsfrom")
		ßMapsto := πg.InternStr("Mapsto")
		ßMars := πg.InternStr("Mars")
		ßMercury := πg.InternStr("Mercury")
		ßNeptune := πg.InternStr("Neptune")
		ßOmega := πg.InternStr("Omega")
		ßPhi := πg.InternStr("Phi")
		ßPi := πg.InternStr("Pi")
		ßPluto := πg.InternStr("Pluto")
		ßPsi := πg.InternStr("Psi")
		ßRHD := πg.InternStr("RHD")
		ßRIGHTCIRCLE := πg.InternStr("RIGHTCIRCLE")
		ßRIGHTcircle := πg.InternStr("RIGHTcircle")
		ßRbag := πg.InternStr("Rbag")
		ßRe := πg.InternStr("Re")
		ßRightarrow := πg.InternStr("Rightarrow")
		ßRrightarrow := πg.InternStr("Rrightarrow")
		ßRsh := πg.InternStr("Rsh")
		ßSaturn := πg.InternStr("Saturn")
		ßScorpio := πg.InternStr("Scorpio")
		ßSigma := πg.InternStr("Sigma")
		ßSquare := πg.InternStr("Square")
		ßSubset := πg.InternStr("Subset")
		ßSun := πg.InternStr("Sun")
		ßSupset := πg.InternStr("Supset")
		ßTaurus := πg.InternStr("Taurus")
		ßTheta := πg.InternStr("Theta")
		ßUparrow := πg.InternStr("Uparrow")
		ßUpdownarrow := πg.InternStr("Updownarrow")
		ßUpsilon := πg.InternStr("Upsilon")
		ßUranus := πg.InternStr("Uranus")
		ßVDash := πg.InternStr("VDash")
		ßVdash := πg.InternStr("Vdash")
		ßVenus := πg.InternStr("Venus")
		ßVert := πg.InternStr("Vert")
		ßVvdash := πg.InternStr("Vvdash")
		ßXBox := πg.InternStr("XBox")
		ßXi := πg.InternStr("Xi")
		ßYup := πg.InternStr("Yup")
		ß_ := πg.InternStr("_")
		ßacute := πg.InternStr("acute")
		ßaleph := πg.InternStr("aleph")
		ßalpha := πg.InternStr("alpha")
		ßamalg := πg.InternStr("amalg")
		ßangle := πg.InternStr("angle")
		ßapprge := πg.InternStr("apprge")
		ßapprle := πg.InternStr("apprle")
		ßapprox := πg.InternStr("approx")
		ßapproxeq := πg.InternStr("approxeq")
		ßaquarius := πg.InternStr("aquarius")
		ßaries := πg.InternStr("aries")
		ßast := πg.InternStr("ast")
		ßasymp := πg.InternStr("asymp")
		ßbackepsilon := πg.InternStr("backepsilon")
		ßbackprime := πg.InternStr("backprime")
		ßbacksim := πg.InternStr("backsim")
		ßbacksimeq := πg.InternStr("backsimeq")
		ßbackslash := πg.InternStr("backslash")
		ßbar := πg.InternStr("bar")
		ßbarin := πg.InternStr("barin")
		ßbarleftharpoon := πg.InternStr("barleftharpoon")
		ßbarrightharpoon := πg.InternStr("barrightharpoon")
		ßbarwedge := πg.InternStr("barwedge")
		ßbecause := πg.InternStr("because")
		ßbeta := πg.InternStr("beta")
		ßbeth := πg.InternStr("beth")
		ßbetween := πg.InternStr("between")
		ßbigcap := πg.InternStr("bigcap")
		ßbigcup := πg.InternStr("bigcup")
		ßbiginterleave := πg.InternStr("biginterleave")
		ßbigodot := πg.InternStr("bigodot")
		ßbigoplus := πg.InternStr("bigoplus")
		ßbigotimes := πg.InternStr("bigotimes")
		ßbigsqcup := πg.InternStr("bigsqcup")
		ßbigstar := πg.InternStr("bigstar")
		ßbigtriangledown := πg.InternStr("bigtriangledown")
		ßbigtriangleup := πg.InternStr("bigtriangleup")
		ßbiguplus := πg.InternStr("biguplus")
		ßbigvee := πg.InternStr("bigvee")
		ßbigwedge := πg.InternStr("bigwedge")
		ßbinampersand := πg.InternStr("binampersand")
		ßbindnasrepma := πg.InternStr("bindnasrepma")
		ßblacklozenge := πg.InternStr("blacklozenge")
		ßblacksmiley := πg.InternStr("blacksmiley")
		ßblacksquare := πg.InternStr("blacksquare")
		ßblacktriangledown := πg.InternStr("blacktriangledown")
		ßblacktriangleleft := πg.InternStr("blacktriangleleft")
		ßblacktriangleright := πg.InternStr("blacktriangleright")
		ßblacktriangleup := πg.InternStr("blacktriangleup")
		ßbot := πg.InternStr("bot")
		ßbowtie := πg.InternStr("bowtie")
		ßboxast := πg.InternStr("boxast")
		ßboxbar := πg.InternStr("boxbar")
		ßboxbox := πg.InternStr("boxbox")
		ßboxbslash := πg.InternStr("boxbslash")
		ßboxcircle := πg.InternStr("boxcircle")
		ßboxdot := πg.InternStr("boxdot")
		ßboxminus := πg.InternStr("boxminus")
		ßboxplus := πg.InternStr("boxplus")
		ßboxslash := πg.InternStr("boxslash")
		ßboxtimes := πg.InternStr("boxtimes")
		ßboy := πg.InternStr("boy")
		ßbreve := πg.InternStr("breve")
		ßbullet := πg.InternStr("bullet")
		ßbumpeq := πg.InternStr("bumpeq")
		ßcancer := πg.InternStr("cancer")
		ßcap := πg.InternStr("cap")
		ßcapricornus := πg.InternStr("capricornus")
		ßcdot := πg.InternStr("cdot")
		ßcdots := πg.InternStr("cdots")
		ßcent := πg.InternStr("cent")
		ßcenterdot := πg.InternStr("centerdot")
		ßcheck := πg.InternStr("check")
		ßcheckmark := πg.InternStr("checkmark")
		ßchi := πg.InternStr("chi")
		ßcirc := πg.InternStr("circ")
		ßcirceq := πg.InternStr("circeq")
		ßcirclearrowleft := πg.InternStr("circlearrowleft")
		ßcirclearrowright := πg.InternStr("circlearrowright")
		ßcircledR := πg.InternStr("circledR")
		ßcircledast := πg.InternStr("circledast")
		ßcircledcirc := πg.InternStr("circledcirc")
		ßcircleddash := πg.InternStr("circleddash")
		ßclubsuit := πg.InternStr("clubsuit")
		ßcoloneq := πg.InternStr("coloneq")
		ßcomplement := πg.InternStr("complement")
		ßcong := πg.InternStr("cong")
		ßcoprod := πg.InternStr("coprod")
		ßcorresponds := πg.InternStr("corresponds")
		ßcup := πg.InternStr("cup")
		ßcurlyeqprec := πg.InternStr("curlyeqprec")
		ßcurlyeqsucc := πg.InternStr("curlyeqsucc")
		ßcurlyvee := πg.InternStr("curlyvee")
		ßcurlywedge := πg.InternStr("curlywedge")
		ßcurvearrowleft := πg.InternStr("curvearrowleft")
		ßcurvearrowright := πg.InternStr("curvearrowright")
		ßdagger := πg.InternStr("dagger")
		ßdaleth := πg.InternStr("daleth")
		ßdasharrow := πg.InternStr("dasharrow")
		ßdashleftarrow := πg.InternStr("dashleftarrow")
		ßdashrightarrow := πg.InternStr("dashrightarrow")
		ßdashv := πg.InternStr("dashv")
		ßddagger := πg.InternStr("ddagger")
		ßddddot := πg.InternStr("ddddot")
		ßdddot := πg.InternStr("dddot")
		ßddot := πg.InternStr("ddot")
		ßddots := πg.InternStr("ddots")
		ßdelta := πg.InternStr("delta")
		ßdiameter := πg.InternStr("diameter")
		ßdiamond := πg.InternStr("diamond")
		ßdiamondsuit := πg.InternStr("diamondsuit")
		ßdigamma := πg.InternStr("digamma")
		ßdiv := πg.InternStr("div")
		ßdivideontimes := πg.InternStr("divideontimes")
		ßdlsh := πg.InternStr("dlsh")
		ßdot := πg.InternStr("dot")
		ßdoteq := πg.InternStr("doteq")
		ßdoteqdot := πg.InternStr("doteqdot")
		ßdotplus := πg.InternStr("dotplus")
		ßdoublebarwedge := πg.InternStr("doublebarwedge")
		ßdownarrow := πg.InternStr("downarrow")
		ßdowndownarrows := πg.InternStr("downdownarrows")
		ßdowndownharpoons := πg.InternStr("downdownharpoons")
		ßdownharpoonleft := πg.InternStr("downharpoonleft")
		ßdownharpoonright := πg.InternStr("downharpoonright")
		ßdownuparrows := πg.InternStr("downuparrows")
		ßdownupharpoons := πg.InternStr("downupharpoons")
		ßdrsh := πg.InternStr("drsh")
		ßearth := πg.InternStr("earth")
		ßell := πg.InternStr("ell")
		ßepsilon := πg.InternStr("epsilon")
		ßeqcirc := πg.InternStr("eqcirc")
		ßeqcolon := πg.InternStr("eqcolon")
		ßeqsim := πg.InternStr("eqsim")
		ßeqslantgtr := πg.InternStr("eqslantgtr")
		ßeqslantless := πg.InternStr("eqslantless")
		ßequiv := πg.InternStr("equiv")
		ßeta := πg.InternStr("eta")
		ßeth := πg.InternStr("eth")
		ßexists := πg.InternStr("exists")
		ßfallingdotseq := πg.InternStr("fallingdotseq")
		ßfatsemi := πg.InternStr("fatsemi")
		ßfemale := πg.InternStr("female")
		ßfint := πg.InternStr("fint")
		ßflat := πg.InternStr("flat")
		ßforall := πg.InternStr("forall")
		ßfourth := πg.InternStr("fourth")
		ßfrown := πg.InternStr("frown")
		ßfrownie := πg.InternStr("frownie")
		ßgamma := πg.InternStr("gamma")
		ßge := πg.InternStr("ge")
		ßgemini := πg.InternStr("gemini")
		ßgeq := πg.InternStr("geq")
		ßgeqq := πg.InternStr("geqq")
		ßgeqslant := πg.InternStr("geqslant")
		ßgets := πg.InternStr("gets")
		ßgg := πg.InternStr("gg")
		ßggcurly := πg.InternStr("ggcurly")
		ßggg := πg.InternStr("ggg")
		ßgimel := πg.InternStr("gimel")
		ßgirl := πg.InternStr("girl")
		ßgnapprox := πg.InternStr("gnapprox")
		ßgneq := πg.InternStr("gneq")
		ßgneqq := πg.InternStr("gneqq")
		ßgnsim := πg.InternStr("gnsim")
		ßgrave := πg.InternStr("grave")
		ßgtrapprox := πg.InternStr("gtrapprox")
		ßgtrdot := πg.InternStr("gtrdot")
		ßgtreqless := πg.InternStr("gtreqless")
		ßgtreqqless := πg.InternStr("gtreqqless")
		ßgtrless := πg.InternStr("gtrless")
		ßgtrsim := πg.InternStr("gtrsim")
		ßhash := πg.InternStr("hash")
		ßhat := πg.InternStr("hat")
		ßhbar := πg.InternStr("hbar")
		ßheartsuit := πg.InternStr("heartsuit")
		ßhookleftarrow := πg.InternStr("hookleftarrow")
		ßhookrightarrow := πg.InternStr("hookrightarrow")
		ßhslash := πg.InternStr("hslash")
		ßiddots := πg.InternStr("iddots")
		ßiiiint := πg.InternStr("iiiint")
		ßiiint := πg.InternStr("iiint")
		ßiint := πg.InternStr("iint")
		ßimath := πg.InternStr("imath")
		ßimpliedby := πg.InternStr("impliedby")
		ßimplies := πg.InternStr("implies")
		ßin := πg.InternStr("in")
		ßinfty := πg.InternStr("infty")
		ßint := πg.InternStr("int")
		ßintercal := πg.InternStr("intercal")
		ßinterleave := πg.InternStr("interleave")
		ßinvneg := πg.InternStr("invneg")
		ßiota := πg.InternStr("iota")
		ßjmath := πg.InternStr("jmath")
		ßjupiter := πg.InternStr("jupiter")
		ßkappa := πg.InternStr("kappa")
		ßlambda := πg.InternStr("lambda")
		ßland := πg.InternStr("land")
		ßlangle := πg.InternStr("langle")
		ßlbag := πg.InternStr("lbag")
		ßlbrace := πg.InternStr("lbrace")
		ßlbrack := πg.InternStr("lbrack")
		ßlceil := πg.InternStr("lceil")
		ßldots := πg.InternStr("ldots")
		ßle := πg.InternStr("le")
		ßleftarrow := πg.InternStr("leftarrow")
		ßleftarrowtail := πg.InternStr("leftarrowtail")
		ßleftarrowtriangle := πg.InternStr("leftarrowtriangle")
		ßleftbarharpoon := πg.InternStr("leftbarharpoon")
		ßleftharpoondown := πg.InternStr("leftharpoondown")
		ßleftharpoonup := πg.InternStr("leftharpoonup")
		ßleftleftarrows := πg.InternStr("leftleftarrows")
		ßleftleftharpoons := πg.InternStr("leftleftharpoons")
		ßleftmoon := πg.InternStr("leftmoon")
		ßleftrightarrow := πg.InternStr("leftrightarrow")
		ßleftrightarrows := πg.InternStr("leftrightarrows")
		ßleftrightarrowtriangle := πg.InternStr("leftrightarrowtriangle")
		ßleftrightharpoon := πg.InternStr("leftrightharpoon")
		ßleftrightharpoons := πg.InternStr("leftrightharpoons")
		ßleftrightsquigarrow := πg.InternStr("leftrightsquigarrow")
		ßleftslice := πg.InternStr("leftslice")
		ßleftsquigarrow := πg.InternStr("leftsquigarrow")
		ßleftthreetimes := πg.InternStr("leftthreetimes")
		ßleftturn := πg.InternStr("leftturn")
		ßleo := πg.InternStr("leo")
		ßleq := πg.InternStr("leq")
		ßleqq := πg.InternStr("leqq")
		ßleqslant := πg.InternStr("leqslant")
		ßlessapprox := πg.InternStr("lessapprox")
		ßlessdot := πg.InternStr("lessdot")
		ßlesseqgtr := πg.InternStr("lesseqgtr")
		ßlesseqqgtr := πg.InternStr("lesseqqgtr")
		ßlessgtr := πg.InternStr("lessgtr")
		ßlesssim := πg.InternStr("lesssim")
		ßlfloor := πg.InternStr("lfloor")
		ßlgroup := πg.InternStr("lgroup")
		ßlhd := πg.InternStr("lhd")
		ßlibra := πg.InternStr("libra")
		ßlightning := πg.InternStr("lightning")
		ßll := πg.InternStr("ll")
		ßllbracket := πg.InternStr("llbracket")
		ßllcorner := πg.InternStr("llcorner")
		ßllcurly := πg.InternStr("llcurly")
		ßlll := πg.InternStr("lll")
		ßllparenthesis := πg.InternStr("llparenthesis")
		ßlnapprox := πg.InternStr("lnapprox")
		ßlneq := πg.InternStr("lneq")
		ßlneqq := πg.InternStr("lneqq")
		ßlnot := πg.InternStr("lnot")
		ßlnsim := πg.InternStr("lnsim")
		ßlongleftarrow := πg.InternStr("longleftarrow")
		ßlongleftrightarrow := πg.InternStr("longleftrightarrow")
		ßlongmapsfrom := πg.InternStr("longmapsfrom")
		ßlongmapsto := πg.InternStr("longmapsto")
		ßlongrightarrow := πg.InternStr("longrightarrow")
		ßlooparrowleft := πg.InternStr("looparrowleft")
		ßlooparrowright := πg.InternStr("looparrowright")
		ßlor := πg.InternStr("lor")
		ßlozenge := πg.InternStr("lozenge")
		ßlrcorner := πg.InternStr("lrcorner")
		ßltimes := πg.InternStr("ltimes")
		ßmale := πg.InternStr("male")
		ßmaltese := πg.InternStr("maltese")
		ßmapsfrom := πg.InternStr("mapsfrom")
		ßmapsto := πg.InternStr("mapsto")
		ßmathaccent := πg.InternStr("mathaccent")
		ßmathalpha := πg.InternStr("mathalpha")
		ßmathbin := πg.InternStr("mathbin")
		ßmathclose := πg.InternStr("mathclose")
		ßmathdollar := πg.InternStr("mathdollar")
		ßmathfence := πg.InternStr("mathfence")
		ßmathop := πg.InternStr("mathop")
		ßmathopen := πg.InternStr("mathopen")
		ßmathord := πg.InternStr("mathord")
		ßmathover := πg.InternStr("mathover")
		ßmathradical := πg.InternStr("mathradical")
		ßmathrel := πg.InternStr("mathrel")
		ßmathring := πg.InternStr("mathring")
		ßmathunder := πg.InternStr("mathunder")
		ßmeasuredangle := πg.InternStr("measuredangle")
		ßmedspace := πg.InternStr("medspace")
		ßmercury := πg.InternStr("mercury")
		ßmho := πg.InternStr("mho")
		ßmid := πg.InternStr("mid")
		ßmodels := πg.InternStr("models")
		ßmp := πg.InternStr("mp")
		ßmu := πg.InternStr("mu")
		ßmultimap := πg.InternStr("multimap")
		ßnLeftarrow := πg.InternStr("nLeftarrow")
		ßnLeftrightarrow := πg.InternStr("nLeftrightarrow")
		ßnRightarrow := πg.InternStr("nRightarrow")
		ßnVDash := πg.InternStr("nVDash")
		ßnVdash := πg.InternStr("nVdash")
		ßnabla := πg.InternStr("nabla")
		ßnatural := πg.InternStr("natural")
		ßncong := πg.InternStr("ncong")
		ßne := πg.InternStr("ne")
		ßnearrow := πg.InternStr("nearrow")
		ßneg := πg.InternStr("neg")
		ßneptune := πg.InternStr("neptune")
		ßneq := πg.InternStr("neq")
		ßnexists := πg.InternStr("nexists")
		ßngeq := πg.InternStr("ngeq")
		ßngtr := πg.InternStr("ngtr")
		ßni := πg.InternStr("ni")
		ßnleftarrow := πg.InternStr("nleftarrow")
		ßnleftrightarrow := πg.InternStr("nleftrightarrow")
		ßnleq := πg.InternStr("nleq")
		ßnless := πg.InternStr("nless")
		ßnmid := πg.InternStr("nmid")
		ßnot := πg.InternStr("not")
		ßnotasymp := πg.InternStr("notasymp")
		ßnotbackslash := πg.InternStr("notbackslash")
		ßnotin := πg.InternStr("notin")
		ßnotowner := πg.InternStr("notowner")
		ßnotslash := πg.InternStr("notslash")
		ßnparallel := πg.InternStr("nparallel")
		ßnprec := πg.InternStr("nprec")
		ßnpreceq := πg.InternStr("npreceq")
		ßnrightarrow := πg.InternStr("nrightarrow")
		ßnsim := πg.InternStr("nsim")
		ßnsubseteq := πg.InternStr("nsubseteq")
		ßnsucc := πg.InternStr("nsucc")
		ßnsucceq := πg.InternStr("nsucceq")
		ßnsupseteq := πg.InternStr("nsupseteq")
		ßntriangleleft := πg.InternStr("ntriangleleft")
		ßntrianglelefteq := πg.InternStr("ntrianglelefteq")
		ßntriangleright := πg.InternStr("ntriangleright")
		ßntrianglerighteq := πg.InternStr("ntrianglerighteq")
		ßnu := πg.InternStr("nu")
		ßnvDash := πg.InternStr("nvDash")
		ßnvdash := πg.InternStr("nvdash")
		ßnwarrow := πg.InternStr("nwarrow")
		ßodot := πg.InternStr("odot")
		ßoiint := πg.InternStr("oiint")
		ßoint := πg.InternStr("oint")
		ßointctrclockwise := πg.InternStr("ointctrclockwise")
		ßomega := πg.InternStr("omega")
		ßominus := πg.InternStr("ominus")
		ßoplus := πg.InternStr("oplus")
		ßoslash := πg.InternStr("oslash")
		ßotimes := πg.InternStr("otimes")
		ßoverbrace := πg.InternStr("overbrace")
		ßoverleftarrow := πg.InternStr("overleftarrow")
		ßoverleftrightarrow := πg.InternStr("overleftrightarrow")
		ßoverline := πg.InternStr("overline")
		ßoverrightarrow := πg.InternStr("overrightarrow")
		ßowns := πg.InternStr("owns")
		ßparallel := πg.InternStr("parallel")
		ßpartial := πg.InternStr("partial")
		ßperp := πg.InternStr("perp")
		ßphi := πg.InternStr("phi")
		ßpi := πg.InternStr("pi")
		ßpisces := πg.InternStr("pisces")
		ßpitchfork := πg.InternStr("pitchfork")
		ßpluto := πg.InternStr("pluto")
		ßpm := πg.InternStr("pm")
		ßpounds := πg.InternStr("pounds")
		ßprec := πg.InternStr("prec")
		ßprecapprox := πg.InternStr("precapprox")
		ßpreccurlyeq := πg.InternStr("preccurlyeq")
		ßpreceq := πg.InternStr("preceq")
		ßprecnapprox := πg.InternStr("precnapprox")
		ßprecnsim := πg.InternStr("precnsim")
		ßprecsim := πg.InternStr("precsim")
		ßprime := πg.InternStr("prime")
		ßprod := πg.InternStr("prod")
		ßpropto := πg.InternStr("propto")
		ßpsi := πg.InternStr("psi")
		ßquad := πg.InternStr("quad")
		ßquarternote := πg.InternStr("quarternote")
		ßrangle := πg.InternStr("rangle")
		ßrbag := πg.InternStr("rbag")
		ßrbrace := πg.InternStr("rbrace")
		ßrbrack := πg.InternStr("rbrack")
		ßrceil := πg.InternStr("rceil")
		ßrestriction := πg.InternStr("restriction")
		ßrfloor := πg.InternStr("rfloor")
		ßrgroup := πg.InternStr("rgroup")
		ßrhd := πg.InternStr("rhd")
		ßrho := πg.InternStr("rho")
		ßrightarrow := πg.InternStr("rightarrow")
		ßrightarrowtail := πg.InternStr("rightarrowtail")
		ßrightarrowtriangle := πg.InternStr("rightarrowtriangle")
		ßrightbarharpoon := πg.InternStr("rightbarharpoon")
		ßrightharpoondown := πg.InternStr("rightharpoondown")
		ßrightharpoonup := πg.InternStr("rightharpoonup")
		ßrightleftarrows := πg.InternStr("rightleftarrows")
		ßrightleftharpoon := πg.InternStr("rightleftharpoon")
		ßrightleftharpoons := πg.InternStr("rightleftharpoons")
		ßrightmoon := πg.InternStr("rightmoon")
		ßrightrightarrows := πg.InternStr("rightrightarrows")
		ßrightrightharpoons := πg.InternStr("rightrightharpoons")
		ßrightslice := πg.InternStr("rightslice")
		ßrightsquigarrow := πg.InternStr("rightsquigarrow")
		ßrightthreetimes := πg.InternStr("rightthreetimes")
		ßrightturn := πg.InternStr("rightturn")
		ßrisingdotseq := πg.InternStr("risingdotseq")
		ßrrbracket := πg.InternStr("rrbracket")
		ßrrparenthesis := πg.InternStr("rrparenthesis")
		ßrtimes := πg.InternStr("rtimes")
		ßsagittarius := πg.InternStr("sagittarius")
		ßsaturn := πg.InternStr("saturn")
		ßscorpio := πg.InternStr("scorpio")
		ßsearrow := πg.InternStr("searrow")
		ßsecond := πg.InternStr("second")
		ßsetminus := πg.InternStr("setminus")
		ßsharp := πg.InternStr("sharp")
		ßsigma := πg.InternStr("sigma")
		ßsim := πg.InternStr("sim")
		ßsimeq := πg.InternStr("simeq")
		ßslash := πg.InternStr("slash")
		ßsmallfrown := πg.InternStr("smallfrown")
		ßsmallsetminus := πg.InternStr("smallsetminus")
		ßsmallsmile := πg.InternStr("smallsmile")
		ßsmalltriangledown := πg.InternStr("smalltriangledown")
		ßsmalltriangleleft := πg.InternStr("smalltriangleleft")
		ßsmalltriangleright := πg.InternStr("smalltriangleright")
		ßsmalltriangleup := πg.InternStr("smalltriangleup")
		ßsmile := πg.InternStr("smile")
		ßsmiley := πg.InternStr("smiley")
		ßspace := πg.InternStr("space")
		ßspadesuit := πg.InternStr("spadesuit")
		ßspddot := πg.InternStr("spddot")
		ßsphat := πg.InternStr("sphat")
		ßsphericalangle := πg.InternStr("sphericalangle")
		ßsptilde := πg.InternStr("sptilde")
		ßsqcap := πg.InternStr("sqcap")
		ßsqcup := πg.InternStr("sqcup")
		ßsqint := πg.InternStr("sqint")
		ßsqrt := πg.InternStr("sqrt")
		ßsqsubset := πg.InternStr("sqsubset")
		ßsqsubseteq := πg.InternStr("sqsubseteq")
		ßsqsupset := πg.InternStr("sqsupset")
		ßsqsupseteq := πg.InternStr("sqsupseteq")
		ßsquare := πg.InternStr("square")
		ßsslash := πg.InternStr("sslash")
		ßstar := πg.InternStr("star")
		ßsubset := πg.InternStr("subset")
		ßsubseteq := πg.InternStr("subseteq")
		ßsubseteqq := πg.InternStr("subseteqq")
		ßsubsetneq := πg.InternStr("subsetneq")
		ßsubsetneqq := πg.InternStr("subsetneqq")
		ßsucc := πg.InternStr("succ")
		ßsuccapprox := πg.InternStr("succapprox")
		ßsucccurlyeq := πg.InternStr("succcurlyeq")
		ßsucceq := πg.InternStr("succeq")
		ßsuccnapprox := πg.InternStr("succnapprox")
		ßsuccnsim := πg.InternStr("succnsim")
		ßsuccsim := πg.InternStr("succsim")
		ßsum := πg.InternStr("sum")
		ßsun := πg.InternStr("sun")
		ßsupset := πg.InternStr("supset")
		ßsupseteq := πg.InternStr("supseteq")
		ßsupseteqq := πg.InternStr("supseteqq")
		ßsupsetneq := πg.InternStr("supsetneq")
		ßsupsetneqq := πg.InternStr("supsetneqq")
		ßswarrow := πg.InternStr("swarrow")
		ßtalloblong := πg.InternStr("talloblong")
		ßtau := πg.InternStr("tau")
		ßtaurus := πg.InternStr("taurus")
		ßtherefore := πg.InternStr("therefore")
		ßtheta := πg.InternStr("theta")
		ßthird := πg.InternStr("third")
		ßtilde := πg.InternStr("tilde")
		ßtimes := πg.InternStr("times")
		ßto := πg.InternStr("to")
		ßtop := πg.InternStr("top")
		ßtriangle := πg.InternStr("triangle")
		ßtriangledown := πg.InternStr("triangledown")
		ßtriangleleft := πg.InternStr("triangleleft")
		ßtrianglelefteq := πg.InternStr("trianglelefteq")
		ßtriangleq := πg.InternStr("triangleq")
		ßtriangleright := πg.InternStr("triangleright")
		ßtrianglerighteq := πg.InternStr("trianglerighteq")
		ßtwoheadleftarrow := πg.InternStr("twoheadleftarrow")
		ßtwoheadrightarrow := πg.InternStr("twoheadrightarrow")
		ßtwonotes := πg.InternStr("twonotes")
		ßulcorner := πg.InternStr("ulcorner")
		ßunderbar := πg.InternStr("underbar")
		ßunderbrace := πg.InternStr("underbrace")
		ßunderleftarrow := πg.InternStr("underleftarrow")
		ßunderline := πg.InternStr("underline")
		ßunderrightarrow := πg.InternStr("underrightarrow")
		ßuparrow := πg.InternStr("uparrow")
		ßupdownarrow := πg.InternStr("updownarrow")
		ßupdownarrows := πg.InternStr("updownarrows")
		ßupdownharpoons := πg.InternStr("updownharpoons")
		ßupharpoonleft := πg.InternStr("upharpoonleft")
		ßupharpoonright := πg.InternStr("upharpoonright")
		ßuplus := πg.InternStr("uplus")
		ßupsilon := πg.InternStr("upsilon")
		ßupuparrows := πg.InternStr("upuparrows")
		ßupupharpoons := πg.InternStr("upupharpoons")
		ßuranus := πg.InternStr("uranus")
		ßurcorner := πg.InternStr("urcorner")
		ßvDash := πg.InternStr("vDash")
		ßvarDelta := πg.InternStr("varDelta")
		ßvarEarth := πg.InternStr("varEarth")
		ßvarGamma := πg.InternStr("varGamma")
		ßvarLambda := πg.InternStr("varLambda")
		ßvarOmega := πg.InternStr("varOmega")
		ßvarPhi := πg.InternStr("varPhi")
		ßvarPi := πg.InternStr("varPi")
		ßvarPsi := πg.InternStr("varPsi")
		ßvarSigma := πg.InternStr("varSigma")
		ßvarTheta := πg.InternStr("varTheta")
		ßvarUpsilon := πg.InternStr("varUpsilon")
		ßvarXi := πg.InternStr("varXi")
		ßvarepsilon := πg.InternStr("varepsilon")
		ßvarkappa := πg.InternStr("varkappa")
		ßvarnothing := πg.InternStr("varnothing")
		ßvarointclockwise := πg.InternStr("varointclockwise")
		ßvarphi := πg.InternStr("varphi")
		ßvarpi := πg.InternStr("varpi")
		ßvarpropto := πg.InternStr("varpropto")
		ßvarrho := πg.InternStr("varrho")
		ßvarsigma := πg.InternStr("varsigma")
		ßvartheta := πg.InternStr("vartheta")
		ßvartriangle := πg.InternStr("vartriangle")
		ßvartriangleleft := πg.InternStr("vartriangleleft")
		ßvartriangleright := πg.InternStr("vartriangleright")
		ßvdash := πg.InternStr("vdash")
		ßvdots := πg.InternStr("vdots")
		ßvec := πg.InternStr("vec")
		ßvee := πg.InternStr("vee")
		ßveebar := πg.InternStr("veebar")
		ßvert := πg.InternStr("vert")
		ßvirgo := πg.InternStr("virgo")
		ßwasylozenge := πg.InternStr("wasylozenge")
		ßwasytherefore := πg.InternStr("wasytherefore")
		ßwedge := πg.InternStr("wedge")
		ßwidehat := πg.InternStr("widehat")
		ßwideparen := πg.InternStr("wideparen")
		ßwidetilde := πg.InternStr("widetilde")
		ßwp := πg.InternStr("wp")
		ßwr := πg.InternStr("wr")
		ßxi := πg.InternStr("xi")
		ßyen := πg.InternStr("yen")
		ßzeta := πg.InternStr("zeta")
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
			// line 9: mathaccent = {
			πF.SetLineno(9)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, ßacute.ToObject(), πg.NewUnicode("\xcc\x81").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbar.ToObject(), πg.NewUnicode("\xcc\x84").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbreve.ToObject(), πg.NewUnicode("\xcc\x86").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcheck.ToObject(), πg.NewUnicode("\xcc\x8c").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßddddot.ToObject(), πg.NewUnicode("\xe2\x83\x9c").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdddot.ToObject(), πg.NewUnicode("\xe2\x83\x9b").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßddot.ToObject(), πg.NewUnicode("\xcc\x88").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdot.ToObject(), πg.NewUnicode("\xcc\x87").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßgrave.ToObject(), πg.NewUnicode("\xcc\x80").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßhat.ToObject(), πg.NewUnicode("\xcc\x82").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßmathring.ToObject(), πg.NewUnicode("\xcc\x8a").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnot.ToObject(), πg.NewUnicode("\xcc\xb8").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßoverleftarrow.ToObject(), πg.NewUnicode("\xe2\x83\x96").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßoverleftrightarrow.ToObject(), πg.NewUnicode("\xe2\x83\xa1").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßoverline.ToObject(), πg.NewUnicode("\xcc\x85").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßoverrightarrow.ToObject(), πg.NewUnicode("\xe2\x83\x97").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtilde.ToObject(), πg.NewUnicode("\xcc\x83").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßunderbar.ToObject(), πg.NewUnicode("\xcc\xb1").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßunderleftarrow.ToObject(), πg.NewUnicode("\xe2\x83\xae").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßunderline.ToObject(), πg.NewUnicode("\xcc\xb2").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßunderrightarrow.ToObject(), πg.NewUnicode("\xe2\x83\xaf").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvec.ToObject(), πg.NewUnicode("\xe2\x83\x97").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßwidehat.ToObject(), πg.NewUnicode("\xcc\x82").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßwidetilde.ToObject(), πg.NewUnicode("\xcc\x83").ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßmathaccent.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 35: mathalpha = {
			πF.SetLineno(35)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, ßBbbk.ToObject(), πg.NewUnicode("\xf0\x9d\x95\x9c").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßDelta.ToObject(), πg.NewUnicode("\xce\x94").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßGamma.ToObject(), πg.NewUnicode("\xce\x93").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßIm.ToObject(), πg.NewUnicode("\xe2\x84\x91").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßLambda.ToObject(), πg.NewUnicode("\xce\x9b").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßOmega.ToObject(), πg.NewUnicode("\xce\xa9").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßPhi.ToObject(), πg.NewUnicode("\xce\xa6").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßPi.ToObject(), πg.NewUnicode("\xce\xa0").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßPsi.ToObject(), πg.NewUnicode("\xce\xa8").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßRe.ToObject(), πg.NewUnicode("\xe2\x84\x9c").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßSigma.ToObject(), πg.NewUnicode("\xce\xa3").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßTheta.ToObject(), πg.NewUnicode("\xce\x98").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßUpsilon.ToObject(), πg.NewUnicode("\xce\xa5").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßXi.ToObject(), πg.NewUnicode("\xce\x9e").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßaleph.ToObject(), πg.NewUnicode("\xe2\x84\xb5").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßalpha.ToObject(), πg.NewUnicode("\xce\xb1").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbeta.ToObject(), πg.NewUnicode("\xce\xb2").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbeth.ToObject(), πg.NewUnicode("\xe2\x84\xb6").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßchi.ToObject(), πg.NewUnicode("\xcf\x87").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdaleth.ToObject(), πg.NewUnicode("\xe2\x84\xb8").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdelta.ToObject(), πg.NewUnicode("\xce\xb4").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdigamma.ToObject(), πg.NewUnicode("\xcf\x9c").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßell.ToObject(), πg.NewUnicode("\xe2\x84\x93").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßepsilon.ToObject(), πg.NewUnicode("\xcf\xb5").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßeta.ToObject(), πg.NewUnicode("\xce\xb7").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßeth.ToObject(), πg.NewUnicode("\xc3\xb0").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßgamma.ToObject(), πg.NewUnicode("\xce\xb3").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßgimel.ToObject(), πg.NewUnicode("\xe2\x84\xb7").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßhbar.ToObject(), πg.NewUnicode("\xe2\x84\x8f").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßhslash.ToObject(), πg.NewUnicode("\xe2\x84\x8f").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßimath.ToObject(), πg.NewUnicode("\xc4\xb1").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßiota.ToObject(), πg.NewUnicode("\xce\xb9").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßjmath.ToObject(), πg.NewUnicode("\xc8\xb7").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßkappa.ToObject(), πg.NewUnicode("\xce\xba").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlambda.ToObject(), πg.NewUnicode("\xce\xbb").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßmu.ToObject(), πg.NewUnicode("\xce\xbc").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnu.ToObject(), πg.NewUnicode("\xce\xbd").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßomega.ToObject(), πg.NewUnicode("\xcf\x89").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßphi.ToObject(), πg.NewUnicode("\xcf\x95").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßpi.ToObject(), πg.NewUnicode("\xcf\x80").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßpsi.ToObject(), πg.NewUnicode("\xcf\x88").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrho.ToObject(), πg.NewUnicode("\xcf\x81").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsigma.ToObject(), πg.NewUnicode("\xcf\x83").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtau.ToObject(), πg.NewUnicode("\xcf\x84").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtheta.ToObject(), πg.NewUnicode("\xce\xb8").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßupsilon.ToObject(), πg.NewUnicode("\xcf\x85").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvarDelta.ToObject(), πg.NewUnicode("\xf0\x9d\x9b\xa5").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvarGamma.ToObject(), πg.NewUnicode("\xf0\x9d\x9b\xa4").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvarLambda.ToObject(), πg.NewUnicode("\xf0\x9d\x9b\xac").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvarOmega.ToObject(), πg.NewUnicode("\xf0\x9d\x9b\xba").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvarPhi.ToObject(), πg.NewUnicode("\xf0\x9d\x9b\xb7").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvarPi.ToObject(), πg.NewUnicode("\xf0\x9d\x9b\xb1").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvarPsi.ToObject(), πg.NewUnicode("\xf0\x9d\x9b\xb9").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvarSigma.ToObject(), πg.NewUnicode("\xf0\x9d\x9b\xb4").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvarTheta.ToObject(), πg.NewUnicode("\xf0\x9d\x9b\xa9").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvarUpsilon.ToObject(), πg.NewUnicode("\xf0\x9d\x9b\xb6").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvarXi.ToObject(), πg.NewUnicode("\xf0\x9d\x9b\xaf").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvarepsilon.ToObject(), πg.NewUnicode("\xce\xb5").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvarkappa.ToObject(), πg.NewUnicode("\xf0\x9d\x9c\x98").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvarphi.ToObject(), πg.NewUnicode("\xcf\x86").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvarpi.ToObject(), πg.NewUnicode("\xcf\x96").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvarrho.ToObject(), πg.NewUnicode("\xcf\xb1").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvarsigma.ToObject(), πg.NewUnicode("\xcf\x82").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvartheta.ToObject(), πg.NewUnicode("\xcf\x91").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßwp.ToObject(), πg.NewUnicode("\xe2\x84\x98").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßxi.ToObject(), πg.NewUnicode("\xce\xbe").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßzeta.ToObject(), πg.NewUnicode("\xce\xb6").ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßmathalpha.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 104: mathbin = {
			πF.SetLineno(104)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, ßCap.ToObject(), πg.NewUnicode("\xe2\x8b\x92").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßCircle.ToObject(), πg.NewUnicode("\xe2\x97\x8b").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßCup.ToObject(), πg.NewUnicode("\xe2\x8b\x93").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßLHD.ToObject(), πg.NewUnicode("\xe2\x97\x80").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßRHD.ToObject(), πg.NewUnicode("\xe2\x96\xb6").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßamalg.ToObject(), πg.NewUnicode("\xe2\xa8\xbf").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßast.ToObject(), πg.NewUnicode("\xe2\x88\x97").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbarwedge.ToObject(), πg.NewUnicode("\xe2\x8a\xbc").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbigtriangledown.ToObject(), πg.NewUnicode("\xe2\x96\xbd").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbigtriangleup.ToObject(), πg.NewUnicode("\xe2\x96\xb3").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbindnasrepma.ToObject(), πg.NewUnicode("\xe2\x85\x8b").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßblacklozenge.ToObject(), πg.NewUnicode("\xe2\xa7\xab").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßblacktriangledown.ToObject(), πg.NewUnicode("\xe2\x96\xbe").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßblacktriangleleft.ToObject(), πg.NewUnicode("\xe2\x97\x82").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßblacktriangleright.ToObject(), πg.NewUnicode("\xe2\x96\xb8").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßblacktriangleup.ToObject(), πg.NewUnicode("\xe2\x96\xb4").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßboxast.ToObject(), πg.NewUnicode("\xe2\xa7\x86").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßboxbar.ToObject(), πg.NewUnicode("\xe2\x97\xab").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßboxbox.ToObject(), πg.NewUnicode("\xe2\xa7\x88").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßboxbslash.ToObject(), πg.NewUnicode("\xe2\xa7\x85").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßboxcircle.ToObject(), πg.NewUnicode("\xe2\xa7\x87").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßboxdot.ToObject(), πg.NewUnicode("\xe2\x8a\xa1").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßboxminus.ToObject(), πg.NewUnicode("\xe2\x8a\x9f").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßboxplus.ToObject(), πg.NewUnicode("\xe2\x8a\x9e").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßboxslash.ToObject(), πg.NewUnicode("\xe2\xa7\x84").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßboxtimes.ToObject(), πg.NewUnicode("\xe2\x8a\xa0").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbullet.ToObject(), πg.NewUnicode("\xe2\x88\x99").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcap.ToObject(), πg.NewUnicode("\xe2\x88\xa9").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcdot.ToObject(), πg.NewUnicode("\xe2\x8b\x85").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcirc.ToObject(), πg.NewUnicode("\xe2\x88\x98").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcircledast.ToObject(), πg.NewUnicode("\xe2\x8a\x9b").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcircledcirc.ToObject(), πg.NewUnicode("\xe2\x8a\x9a").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcircleddash.ToObject(), πg.NewUnicode("\xe2\x8a\x9d").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcup.ToObject(), πg.NewUnicode("\xe2\x88\xaa").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcurlyvee.ToObject(), πg.NewUnicode("\xe2\x8b\x8e").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcurlywedge.ToObject(), πg.NewUnicode("\xe2\x8b\x8f").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdagger.ToObject(), πg.NewUnicode("\xe2\x80\xa0").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßddagger.ToObject(), πg.NewUnicode("\xe2\x80\xa1").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdiamond.ToObject(), πg.NewUnicode("\xe2\x8b\x84").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdiv.ToObject(), πg.NewUnicode("\xc3\xb7").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdivideontimes.ToObject(), πg.NewUnicode("\xe2\x8b\x87").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdotplus.ToObject(), πg.NewUnicode("\xe2\x88\x94").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdoublebarwedge.ToObject(), πg.NewUnicode("\xe2\xa9\x9e").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßintercal.ToObject(), πg.NewUnicode("\xe2\x8a\xba").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßinterleave.ToObject(), πg.NewUnicode("\xe2\xab\xb4").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßland.ToObject(), πg.NewUnicode("\xe2\x88\xa7").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßleftthreetimes.ToObject(), πg.NewUnicode("\xe2\x8b\x8b").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlhd.ToObject(), πg.NewUnicode("\xe2\x97\x81").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlor.ToObject(), πg.NewUnicode("\xe2\x88\xa8").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßltimes.ToObject(), πg.NewUnicode("\xe2\x8b\x89").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßmp.ToObject(), πg.NewUnicode("\xe2\x88\x93").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßodot.ToObject(), πg.NewUnicode("\xe2\x8a\x99").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßominus.ToObject(), πg.NewUnicode("\xe2\x8a\x96").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßoplus.ToObject(), πg.NewUnicode("\xe2\x8a\x95").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßoslash.ToObject(), πg.NewUnicode("\xe2\x8a\x98").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßotimes.ToObject(), πg.NewUnicode("\xe2\x8a\x97").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßpm.ToObject(), πg.NewUnicode("\xc2\xb1").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrhd.ToObject(), πg.NewUnicode("\xe2\x96\xb7").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrightthreetimes.ToObject(), πg.NewUnicode("\xe2\x8b\x8c").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrtimes.ToObject(), πg.NewUnicode("\xe2\x8b\x8a").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsetminus.ToObject(), πg.NewUnicode("\xe2\xa7\xb5").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßslash.ToObject(), πg.NewUnicode("\xe2\x88\x95").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsmallsetminus.ToObject(), πg.NewUnicode("\xe2\x88\x96").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsmalltriangledown.ToObject(), πg.NewUnicode("\xe2\x96\xbf").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsmalltriangleleft.ToObject(), πg.NewUnicode("\xe2\x97\x83").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsmalltriangleright.ToObject(), πg.NewUnicode("\xe2\x96\xb9").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsmalltriangleup.ToObject(), πg.NewUnicode("\xe2\x96\xb5").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsqcap.ToObject(), πg.NewUnicode("\xe2\x8a\x93").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsqcup.ToObject(), πg.NewUnicode("\xe2\x8a\x94").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsslash.ToObject(), πg.NewUnicode("\xe2\xab\xbd").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßstar.ToObject(), πg.NewUnicode("\xe2\x8b\x86").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtalloblong.ToObject(), πg.NewUnicode("\xe2\xab\xbe").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtimes.ToObject(), πg.NewUnicode("\xc3\x97").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtriangle.ToObject(), πg.NewUnicode("\xe2\x96\xb3").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtriangledown.ToObject(), πg.NewUnicode("\xe2\x96\xbf").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtriangleleft.ToObject(), πg.NewUnicode("\xe2\x97\x83").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtriangleright.ToObject(), πg.NewUnicode("\xe2\x96\xb9").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßuplus.ToObject(), πg.NewUnicode("\xe2\x8a\x8e").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvartriangle.ToObject(), πg.NewUnicode("\xe2\x96\xb3").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvee.ToObject(), πg.NewUnicode("\xe2\x88\xa8").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßveebar.ToObject(), πg.NewUnicode("\xe2\x8a\xbb").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßwedge.ToObject(), πg.NewUnicode("\xe2\x88\xa7").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßwr.ToObject(), πg.NewUnicode("\xe2\x89\x80").ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßmathbin.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 189: mathclose = {
			πF.SetLineno(189)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, ßRbag.ToObject(), πg.NewUnicode("\xe2\x9f\x86").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlrcorner.ToObject(), πg.NewUnicode("\xe2\x8c\x9f").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrangle.ToObject(), πg.NewUnicode("\xe2\x9f\xa9").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrbag.ToObject(), πg.NewUnicode("\xe2\x9f\x86").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrbrace.ToObject(), πg.NewUnicode("}").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrbrack.ToObject(), πg.NewUnicode("]").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrceil.ToObject(), πg.NewUnicode("\xe2\x8c\x89").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrfloor.ToObject(), πg.NewUnicode("\xe2\x8c\x8b").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrgroup.ToObject(), πg.NewUnicode("\xe2\x9f\xaf").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrrbracket.ToObject(), πg.NewUnicode("\xe2\x9f\xa7").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrrparenthesis.ToObject(), πg.NewUnicode("\xe2\xa6\x88").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßurcorner.ToObject(), πg.NewUnicode("\xe2\x8c\x9d").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("}").ToObject(), πg.NewUnicode("}").ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßmathclose.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 204: mathfence = {
			πF.SetLineno(204)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, ßVert.ToObject(), πg.NewUnicode("\xe2\x80\x96").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvert.ToObject(), πg.NewUnicode("|").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("|").ToObject(), πg.NewUnicode("\xe2\x80\x96").ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßmathfence.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 209: mathop = {
			πF.SetLineno(209)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, ßJoin.ToObject(), πg.NewUnicode("\xe2\xa8\x9d").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbigcap.ToObject(), πg.NewUnicode("\xe2\x8b\x82").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbigcup.ToObject(), πg.NewUnicode("\xe2\x8b\x83").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbiginterleave.ToObject(), πg.NewUnicode("\xe2\xab\xbc").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbigodot.ToObject(), πg.NewUnicode("\xe2\xa8\x80").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbigoplus.ToObject(), πg.NewUnicode("\xe2\xa8\x81").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbigotimes.ToObject(), πg.NewUnicode("\xe2\xa8\x82").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbigsqcup.ToObject(), πg.NewUnicode("\xe2\xa8\x86").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbiguplus.ToObject(), πg.NewUnicode("\xe2\xa8\x84").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbigvee.ToObject(), πg.NewUnicode("\xe2\x8b\x81").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbigwedge.ToObject(), πg.NewUnicode("\xe2\x8b\x80").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcoprod.ToObject(), πg.NewUnicode("\xe2\x88\x90").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßfatsemi.ToObject(), πg.NewUnicode("\xe2\xa8\x9f").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßfint.ToObject(), πg.NewUnicode("\xe2\xa8\x8f").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßiiiint.ToObject(), πg.NewUnicode("\xe2\xa8\x8c").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßiiint.ToObject(), πg.NewUnicode("\xe2\x88\xad").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßiint.ToObject(), πg.NewUnicode("\xe2\x88\xac").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßint.ToObject(), πg.NewUnicode("\xe2\x88\xab").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßoiint.ToObject(), πg.NewUnicode("\xe2\x88\xaf").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßoint.ToObject(), πg.NewUnicode("\xe2\x88\xae").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßointctrclockwise.ToObject(), πg.NewUnicode("\xe2\x88\xb3").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßprod.ToObject(), πg.NewUnicode("\xe2\x88\x8f").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsqint.ToObject(), πg.NewUnicode("\xe2\xa8\x96").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsum.ToObject(), πg.NewUnicode("\xe2\x88\x91").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvarointclockwise.ToObject(), πg.NewUnicode("\xe2\x88\xb2").ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßmathop.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 236: mathopen = {
			πF.SetLineno(236)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, ßLbag.ToObject(), πg.NewUnicode("\xe2\x9f\x85").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlangle.ToObject(), πg.NewUnicode("\xe2\x9f\xa8").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlbag.ToObject(), πg.NewUnicode("\xe2\x9f\x85").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlbrace.ToObject(), πg.NewUnicode("{").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlbrack.ToObject(), πg.NewUnicode("[").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlceil.ToObject(), πg.NewUnicode("\xe2\x8c\x88").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlfloor.ToObject(), πg.NewUnicode("\xe2\x8c\x8a").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlgroup.ToObject(), πg.NewUnicode("\xe2\x9f\xae").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßllbracket.ToObject(), πg.NewUnicode("\xe2\x9f\xa6").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßllcorner.ToObject(), πg.NewUnicode("\xe2\x8c\x9e").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßllparenthesis.ToObject(), πg.NewUnicode("\xe2\xa6\x87").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßulcorner.ToObject(), πg.NewUnicode("\xe2\x8c\x9c").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("{").ToObject(), πg.NewUnicode("{").ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßmathopen.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 251: mathord = {
			πF.SetLineno(251)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, πg.NewStr("#").ToObject(), πg.NewUnicode("#").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("$").ToObject(), πg.NewUnicode("$").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("%").ToObject(), πg.NewUnicode("%").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("&").ToObject(), πg.NewUnicode("&").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßAC.ToObject(), πg.NewUnicode("\xe2\x88\xbf").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßAPLcomment.ToObject(), πg.NewUnicode("\xe2\x8d\x9d").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßAPLdownarrowbox.ToObject(), πg.NewUnicode("\xe2\x8d\x97").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßAPLinput.ToObject(), πg.NewUnicode("\xe2\x8d\x9e").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßAPLinv.ToObject(), πg.NewUnicode("\xe2\x8c\xb9").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßAPLleftarrowbox.ToObject(), πg.NewUnicode("\xe2\x8d\x87").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßAPLlog.ToObject(), πg.NewUnicode("\xe2\x8d\x9f").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßAPLrightarrowbox.ToObject(), πg.NewUnicode("\xe2\x8d\x88").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßAPLuparrowbox.ToObject(), πg.NewUnicode("\xe2\x8d\x90").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßAries.ToObject(), πg.NewUnicode("\xe2\x99\x88").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßCIRCLE.ToObject(), πg.NewUnicode("\xe2\x97\x8f").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßCheckedBox.ToObject(), πg.NewUnicode("\xe2\x98\x91").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßDiamond.ToObject(), πg.NewUnicode("\xe2\x97\x87").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßFinv.ToObject(), πg.NewUnicode("\xe2\x84\xb2").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßGame.ToObject(), πg.NewUnicode("\xe2\x85\x81").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßGemini.ToObject(), πg.NewUnicode("\xe2\x99\x8a").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßJupiter.ToObject(), πg.NewUnicode("\xe2\x99\x83").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßLEFTCIRCLE.ToObject(), πg.NewUnicode("\xe2\x97\x96").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßLEFTcircle.ToObject(), πg.NewUnicode("\xe2\x97\x90").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßLeo.ToObject(), πg.NewUnicode("\xe2\x99\x8c").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßLibra.ToObject(), πg.NewUnicode("\xe2\x99\x8e").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßMars.ToObject(), πg.NewUnicode("\xe2\x99\x82").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßMercury.ToObject(), πg.NewUnicode("\xe2\x98\xbf").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßNeptune.ToObject(), πg.NewUnicode("\xe2\x99\x86").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßPluto.ToObject(), πg.NewUnicode("\xe2\x99\x87").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßRIGHTCIRCLE.ToObject(), πg.NewUnicode("\xe2\x97\x97").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßRIGHTcircle.ToObject(), πg.NewUnicode("\xe2\x97\x91").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßSaturn.ToObject(), πg.NewUnicode("\xe2\x99\x84").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßScorpio.ToObject(), πg.NewUnicode("\xe2\x99\x8f").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßSquare.ToObject(), πg.NewUnicode("\xe2\x98\x90").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßSun.ToObject(), πg.NewUnicode("\xe2\x98\x89").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßTaurus.ToObject(), πg.NewUnicode("\xe2\x99\x89").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßUranus.ToObject(), πg.NewUnicode("\xe2\x99\x85").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßVenus.ToObject(), πg.NewUnicode("\xe2\x99\x80").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßXBox.ToObject(), πg.NewUnicode("\xe2\x98\x92").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßYup.ToObject(), πg.NewUnicode("\xe2\x85\x84").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ß_.ToObject(), πg.NewUnicode("_").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßangle.ToObject(), πg.NewUnicode("\xe2\x88\xa0").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßaquarius.ToObject(), πg.NewUnicode("\xe2\x99\x92").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßaries.ToObject(), πg.NewUnicode("\xe2\x99\x88").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßast.ToObject(), πg.NewUnicode("*").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbackepsilon.ToObject(), πg.NewUnicode("\xcf\xb6").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbackprime.ToObject(), πg.NewUnicode("\xe2\x80\xb5").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbackslash.ToObject(), πg.NewUnicode("\\").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbecause.ToObject(), πg.NewUnicode("\xe2\x88\xb5").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbigstar.ToObject(), πg.NewUnicode("\xe2\x98\x85").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbinampersand.ToObject(), πg.NewUnicode("&").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßblacklozenge.ToObject(), πg.NewUnicode("\xe2\xac\xa7").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßblacksmiley.ToObject(), πg.NewUnicode("\xe2\x98\xbb").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßblacksquare.ToObject(), πg.NewUnicode("\xe2\x97\xbc").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbot.ToObject(), πg.NewUnicode("\xe2\x8a\xa5").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßboy.ToObject(), πg.NewUnicode("\xe2\x99\x82").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcancer.ToObject(), πg.NewUnicode("\xe2\x99\x8b").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcapricornus.ToObject(), πg.NewUnicode("\xe2\x99\x91").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcdots.ToObject(), πg.NewUnicode("\xe2\x8b\xaf").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcent.ToObject(), πg.NewUnicode("\xc2\xa2").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcenterdot.ToObject(), πg.NewUnicode("\xe2\xac\x9d").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcheckmark.ToObject(), πg.NewUnicode("\xe2\x9c\x93").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcirclearrowleft.ToObject(), πg.NewUnicode("\xe2\x86\xba").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcirclearrowright.ToObject(), πg.NewUnicode("\xe2\x86\xbb").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcircledR.ToObject(), πg.NewUnicode("\xc2\xae").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcircledcirc.ToObject(), πg.NewUnicode("\xe2\x97\x8e").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßclubsuit.ToObject(), πg.NewUnicode("\xe2\x99\xa3").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcomplement.ToObject(), πg.NewUnicode("\xe2\x88\x81").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdasharrow.ToObject(), πg.NewUnicode("\xe2\x87\xa2").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdashleftarrow.ToObject(), πg.NewUnicode("\xe2\x87\xa0").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdashrightarrow.ToObject(), πg.NewUnicode("\xe2\x87\xa2").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdiameter.ToObject(), πg.NewUnicode("\xe2\x8c\x80").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdiamondsuit.ToObject(), πg.NewUnicode("\xe2\x99\xa2").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßearth.ToObject(), πg.NewUnicode("\xe2\x99\x81").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßexists.ToObject(), πg.NewUnicode("\xe2\x88\x83").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßfemale.ToObject(), πg.NewUnicode("\xe2\x99\x80").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßflat.ToObject(), πg.NewUnicode("\xe2\x99\xad").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßforall.ToObject(), πg.NewUnicode("\xe2\x88\x80").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßfourth.ToObject(), πg.NewUnicode("\xe2\x81\x97").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßfrownie.ToObject(), πg.NewUnicode("\xe2\x98\xb9").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßgemini.ToObject(), πg.NewUnicode("\xe2\x99\x8a").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßgirl.ToObject(), πg.NewUnicode("\xe2\x99\x80").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßheartsuit.ToObject(), πg.NewUnicode("\xe2\x99\xa1").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßinfty.ToObject(), πg.NewUnicode("\xe2\x88\x9e").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßinvneg.ToObject(), πg.NewUnicode("\xe2\x8c\x90").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßjupiter.ToObject(), πg.NewUnicode("\xe2\x99\x83").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßldots.ToObject(), πg.NewUnicode("\xe2\x80\xa6").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßleftmoon.ToObject(), πg.NewUnicode("\xe2\x98\xbe").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßleftturn.ToObject(), πg.NewUnicode("\xe2\x86\xba").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßleo.ToObject(), πg.NewUnicode("\xe2\x99\x8c").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlibra.ToObject(), πg.NewUnicode("\xe2\x99\x8e").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlnot.ToObject(), πg.NewUnicode("\xc2\xac").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlozenge.ToObject(), πg.NewUnicode("\xe2\x97\x8a").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßmale.ToObject(), πg.NewUnicode("\xe2\x99\x82").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßmaltese.ToObject(), πg.NewUnicode("\xe2\x9c\xa0").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßmathdollar.ToObject(), πg.NewUnicode("$").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßmeasuredangle.ToObject(), πg.NewUnicode("\xe2\x88\xa1").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßmercury.ToObject(), πg.NewUnicode("\xe2\x98\xbf").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßmho.ToObject(), πg.NewUnicode("\xe2\x84\xa7").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnabla.ToObject(), πg.NewUnicode("\xe2\x88\x87").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnatural.ToObject(), πg.NewUnicode("\xe2\x99\xae").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßneg.ToObject(), πg.NewUnicode("\xc2\xac").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßneptune.ToObject(), πg.NewUnicode("\xe2\x99\x86").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnexists.ToObject(), πg.NewUnicode("\xe2\x88\x84").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnotbackslash.ToObject(), πg.NewUnicode("\xe2\x8d\x80").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßpartial.ToObject(), πg.NewUnicode("\xe2\x88\x82").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßpisces.ToObject(), πg.NewUnicode("\xe2\x99\x93").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßpluto.ToObject(), πg.NewUnicode("\xe2\x99\x87").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßpounds.ToObject(), πg.NewUnicode("\xc2\xa3").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßprime.ToObject(), πg.NewUnicode("\xe2\x80\xb2").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßquarternote.ToObject(), πg.NewUnicode("\xe2\x99\xa9").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrightmoon.ToObject(), πg.NewUnicode("\xe2\x98\xbd").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrightturn.ToObject(), πg.NewUnicode("\xe2\x86\xbb").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsagittarius.ToObject(), πg.NewUnicode("\xe2\x99\x90").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsaturn.ToObject(), πg.NewUnicode("\xe2\x99\x84").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßscorpio.ToObject(), πg.NewUnicode("\xe2\x99\x8f").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsecond.ToObject(), πg.NewUnicode("\xe2\x80\xb3").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsharp.ToObject(), πg.NewUnicode("\xe2\x99\xaf").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsim.ToObject(), πg.NewUnicode("~").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßslash.ToObject(), πg.NewUnicode("/").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsmiley.ToObject(), πg.NewUnicode("\xe2\x98\xba").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßspadesuit.ToObject(), πg.NewUnicode("\xe2\x99\xa0").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßspddot.ToObject(), πg.NewUnicode("\xc2\xa8").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsphat.ToObject(), πg.NewUnicode("^").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsphericalangle.ToObject(), πg.NewUnicode("\xe2\x88\xa2").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsptilde.ToObject(), πg.NewUnicode("~").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsquare.ToObject(), πg.NewUnicode("\xe2\x97\xbb").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsun.ToObject(), πg.NewUnicode("\xe2\x98\xbc").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtaurus.ToObject(), πg.NewUnicode("\xe2\x99\x89").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtherefore.ToObject(), πg.NewUnicode("\xe2\x88\xb4").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßthird.ToObject(), πg.NewUnicode("\xe2\x80\xb4").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtop.ToObject(), πg.NewUnicode("\xe2\x8a\xa4").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtriangleleft.ToObject(), πg.NewUnicode("\xe2\x97\x85").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtriangleright.ToObject(), πg.NewUnicode("\xe2\x96\xbb").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtwonotes.ToObject(), πg.NewUnicode("\xe2\x99\xab").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßuranus.ToObject(), πg.NewUnicode("\xe2\x99\x85").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvarEarth.ToObject(), πg.NewUnicode("\xe2\x99\x81").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvarnothing.ToObject(), πg.NewUnicode("\xe2\x88\x85").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvirgo.ToObject(), πg.NewUnicode("\xe2\x99\x8d").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßwasylozenge.ToObject(), πg.NewUnicode("\xe2\x8c\x91").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßwasytherefore.ToObject(), πg.NewUnicode("\xe2\x88\xb4").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßyen.ToObject(), πg.NewUnicode("\xc2\xa5").ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßmathord.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 395: mathover = {
			πF.SetLineno(395)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, ßoverbrace.ToObject(), πg.NewUnicode("\xe2\x8f\x9e").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßwideparen.ToObject(), πg.NewUnicode("\xe2\x8f\x9c").ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßmathover.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 399: mathradical = {
			πF.SetLineno(399)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, ßsqrt.ToObject(), πg.NewUnicode("\xe2\x88\x9a").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("sqrt[3]").ToObject(), πg.NewUnicode("\xe2\x88\x9b").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("sqrt[4]").ToObject(), πg.NewUnicode("\xe2\x88\x9c").ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßmathradical.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 404: mathrel = {
			πF.SetLineno(404)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, ßBumpeq.ToObject(), πg.NewUnicode("\xe2\x89\x8e").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßDoteq.ToObject(), πg.NewUnicode("\xe2\x89\x91").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßDownarrow.ToObject(), πg.NewUnicode("\xe2\x87\x93").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßLeftarrow.ToObject(), πg.NewUnicode("\xe2\x87\x90").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßLeftrightarrow.ToObject(), πg.NewUnicode("\xe2\x87\x94").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßLleftarrow.ToObject(), πg.NewUnicode("\xe2\x87\x9a").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßLongleftarrow.ToObject(), πg.NewUnicode("\xe2\x9f\xb8").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßLongleftrightarrow.ToObject(), πg.NewUnicode("\xe2\x9f\xba").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßLongmapsfrom.ToObject(), πg.NewUnicode("\xe2\x9f\xbd").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßLongmapsto.ToObject(), πg.NewUnicode("\xe2\x9f\xbe").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßLongrightarrow.ToObject(), πg.NewUnicode("\xe2\x9f\xb9").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßLsh.ToObject(), πg.NewUnicode("\xe2\x86\xb0").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßMapsfrom.ToObject(), πg.NewUnicode("\xe2\xa4\x86").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßMapsto.ToObject(), πg.NewUnicode("\xe2\xa4\x87").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßRightarrow.ToObject(), πg.NewUnicode("\xe2\x87\x92").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßRrightarrow.ToObject(), πg.NewUnicode("\xe2\x87\x9b").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßRsh.ToObject(), πg.NewUnicode("\xe2\x86\xb1").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßSubset.ToObject(), πg.NewUnicode("\xe2\x8b\x90").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßSupset.ToObject(), πg.NewUnicode("\xe2\x8b\x91").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßUparrow.ToObject(), πg.NewUnicode("\xe2\x87\x91").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßUpdownarrow.ToObject(), πg.NewUnicode("\xe2\x87\x95").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßVDash.ToObject(), πg.NewUnicode("\xe2\x8a\xab").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßVdash.ToObject(), πg.NewUnicode("\xe2\x8a\xa9").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßVvdash.ToObject(), πg.NewUnicode("\xe2\x8a\xaa").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßapprge.ToObject(), πg.NewUnicode("\xe2\x89\xb3").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßapprle.ToObject(), πg.NewUnicode("\xe2\x89\xb2").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßapprox.ToObject(), πg.NewUnicode("\xe2\x89\x88").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßapproxeq.ToObject(), πg.NewUnicode("\xe2\x89\x8a").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßasymp.ToObject(), πg.NewUnicode("\xe2\x89\x8d").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbacksim.ToObject(), πg.NewUnicode("\xe2\x88\xbd").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbacksimeq.ToObject(), πg.NewUnicode("\xe2\x8b\x8d").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbarin.ToObject(), πg.NewUnicode("\xe2\x8b\xb6").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbarleftharpoon.ToObject(), πg.NewUnicode("\xe2\xa5\xab").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbarrightharpoon.ToObject(), πg.NewUnicode("\xe2\xa5\xad").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbetween.ToObject(), πg.NewUnicode("\xe2\x89\xac").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbowtie.ToObject(), πg.NewUnicode("\xe2\x8b\x88").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßbumpeq.ToObject(), πg.NewUnicode("\xe2\x89\x8f").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcirceq.ToObject(), πg.NewUnicode("\xe2\x89\x97").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcoloneq.ToObject(), πg.NewUnicode("\xe2\x89\x94").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcong.ToObject(), πg.NewUnicode("\xe2\x89\x85").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcorresponds.ToObject(), πg.NewUnicode("\xe2\x89\x99").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcurlyeqprec.ToObject(), πg.NewUnicode("\xe2\x8b\x9e").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcurlyeqsucc.ToObject(), πg.NewUnicode("\xe2\x8b\x9f").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcurvearrowleft.ToObject(), πg.NewUnicode("\xe2\x86\xb6").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcurvearrowright.ToObject(), πg.NewUnicode("\xe2\x86\xb7").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdashv.ToObject(), πg.NewUnicode("\xe2\x8a\xa3").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßddots.ToObject(), πg.NewUnicode("\xe2\x8b\xb1").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdlsh.ToObject(), πg.NewUnicode("\xe2\x86\xb2").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdoteq.ToObject(), πg.NewUnicode("\xe2\x89\x90").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdoteqdot.ToObject(), πg.NewUnicode("\xe2\x89\x91").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdownarrow.ToObject(), πg.NewUnicode("\xe2\x86\x93").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdowndownarrows.ToObject(), πg.NewUnicode("\xe2\x87\x8a").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdowndownharpoons.ToObject(), πg.NewUnicode("\xe2\xa5\xa5").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdownharpoonleft.ToObject(), πg.NewUnicode("\xe2\x87\x83").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdownharpoonright.ToObject(), πg.NewUnicode("\xe2\x87\x82").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdownuparrows.ToObject(), πg.NewUnicode("\xe2\x87\xb5").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdownupharpoons.ToObject(), πg.NewUnicode("\xe2\xa5\xaf").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdrsh.ToObject(), πg.NewUnicode("\xe2\x86\xb3").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßeqcirc.ToObject(), πg.NewUnicode("\xe2\x89\x96").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßeqcolon.ToObject(), πg.NewUnicode("\xe2\x89\x95").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßeqsim.ToObject(), πg.NewUnicode("\xe2\x89\x82").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßeqslantgtr.ToObject(), πg.NewUnicode("\xe2\xaa\x96").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßeqslantless.ToObject(), πg.NewUnicode("\xe2\xaa\x95").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßequiv.ToObject(), πg.NewUnicode("\xe2\x89\xa1").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßfallingdotseq.ToObject(), πg.NewUnicode("\xe2\x89\x92").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßfrown.ToObject(), πg.NewUnicode("\xe2\x8c\xa2").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßge.ToObject(), πg.NewUnicode("\xe2\x89\xa5").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßgeq.ToObject(), πg.NewUnicode("\xe2\x89\xa5").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßgeqq.ToObject(), πg.NewUnicode("\xe2\x89\xa7").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßgeqslant.ToObject(), πg.NewUnicode("\xe2\xa9\xbe").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßgets.ToObject(), πg.NewUnicode("\xe2\x86\x90").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßgg.ToObject(), πg.NewUnicode("\xe2\x89\xab").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßggcurly.ToObject(), πg.NewUnicode("\xe2\xaa\xbc").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßggg.ToObject(), πg.NewUnicode("\xe2\x8b\x99").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßgnapprox.ToObject(), πg.NewUnicode("\xe2\xaa\x8a").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßgneq.ToObject(), πg.NewUnicode("\xe2\xaa\x88").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßgneqq.ToObject(), πg.NewUnicode("\xe2\x89\xa9").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßgnsim.ToObject(), πg.NewUnicode("\xe2\x8b\xa7").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßgtrapprox.ToObject(), πg.NewUnicode("\xe2\xaa\x86").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßgtrdot.ToObject(), πg.NewUnicode("\xe2\x8b\x97").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßgtreqless.ToObject(), πg.NewUnicode("\xe2\x8b\x9b").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßgtreqqless.ToObject(), πg.NewUnicode("\xe2\xaa\x8c").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßgtrless.ToObject(), πg.NewUnicode("\xe2\x89\xb7").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßgtrsim.ToObject(), πg.NewUnicode("\xe2\x89\xb3").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßhash.ToObject(), πg.NewUnicode("\xe2\x8b\x95").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßhookleftarrow.ToObject(), πg.NewUnicode("\xe2\x86\xa9").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßhookrightarrow.ToObject(), πg.NewUnicode("\xe2\x86\xaa").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßiddots.ToObject(), πg.NewUnicode("\xe2\x8b\xb0").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßimpliedby.ToObject(), πg.NewUnicode("\xe2\x9f\xb8").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßimplies.ToObject(), πg.NewUnicode("\xe2\x9f\xb9").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßin.ToObject(), πg.NewUnicode("\xe2\x88\x88").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßle.ToObject(), πg.NewUnicode("\xe2\x89\xa4").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßleftarrow.ToObject(), πg.NewUnicode("\xe2\x86\x90").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßleftarrowtail.ToObject(), πg.NewUnicode("\xe2\x86\xa2").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßleftarrowtriangle.ToObject(), πg.NewUnicode("\xe2\x87\xbd").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßleftbarharpoon.ToObject(), πg.NewUnicode("\xe2\xa5\xaa").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßleftharpoondown.ToObject(), πg.NewUnicode("\xe2\x86\xbd").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßleftharpoonup.ToObject(), πg.NewUnicode("\xe2\x86\xbc").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßleftleftarrows.ToObject(), πg.NewUnicode("\xe2\x87\x87").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßleftleftharpoons.ToObject(), πg.NewUnicode("\xe2\xa5\xa2").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßleftrightarrow.ToObject(), πg.NewUnicode("\xe2\x86\x94").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßleftrightarrows.ToObject(), πg.NewUnicode("\xe2\x87\x86").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßleftrightarrowtriangle.ToObject(), πg.NewUnicode("\xe2\x87\xbf").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßleftrightharpoon.ToObject(), πg.NewUnicode("\xe2\xa5\x8a").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßleftrightharpoons.ToObject(), πg.NewUnicode("\xe2\x87\x8b").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßleftrightsquigarrow.ToObject(), πg.NewUnicode("\xe2\x86\xad").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßleftslice.ToObject(), πg.NewUnicode("\xe2\xaa\xa6").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßleftsquigarrow.ToObject(), πg.NewUnicode("\xe2\x87\x9c").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßleq.ToObject(), πg.NewUnicode("\xe2\x89\xa4").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßleqq.ToObject(), πg.NewUnicode("\xe2\x89\xa6").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßleqslant.ToObject(), πg.NewUnicode("\xe2\xa9\xbd").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlessapprox.ToObject(), πg.NewUnicode("\xe2\xaa\x85").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlessdot.ToObject(), πg.NewUnicode("\xe2\x8b\x96").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlesseqgtr.ToObject(), πg.NewUnicode("\xe2\x8b\x9a").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlesseqqgtr.ToObject(), πg.NewUnicode("\xe2\xaa\x8b").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlessgtr.ToObject(), πg.NewUnicode("\xe2\x89\xb6").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlesssim.ToObject(), πg.NewUnicode("\xe2\x89\xb2").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlightning.ToObject(), πg.NewUnicode("\xe2\x86\xaf").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßll.ToObject(), πg.NewUnicode("\xe2\x89\xaa").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßllcurly.ToObject(), πg.NewUnicode("\xe2\xaa\xbb").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlll.ToObject(), πg.NewUnicode("\xe2\x8b\x98").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlnapprox.ToObject(), πg.NewUnicode("\xe2\xaa\x89").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlneq.ToObject(), πg.NewUnicode("\xe2\xaa\x87").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlneqq.ToObject(), πg.NewUnicode("\xe2\x89\xa8").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlnsim.ToObject(), πg.NewUnicode("\xe2\x8b\xa6").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlongleftarrow.ToObject(), πg.NewUnicode("\xe2\x9f\xb5").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlongleftrightarrow.ToObject(), πg.NewUnicode("\xe2\x9f\xb7").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlongmapsfrom.ToObject(), πg.NewUnicode("\xe2\x9f\xbb").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlongmapsto.ToObject(), πg.NewUnicode("\xe2\x9f\xbc").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlongrightarrow.ToObject(), πg.NewUnicode("\xe2\x9f\xb6").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlooparrowleft.ToObject(), πg.NewUnicode("\xe2\x86\xab").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlooparrowright.ToObject(), πg.NewUnicode("\xe2\x86\xac").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßmapsfrom.ToObject(), πg.NewUnicode("\xe2\x86\xa4").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßmapsto.ToObject(), πg.NewUnicode("\xe2\x86\xa6").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßmid.ToObject(), πg.NewUnicode("\xe2\x88\xa3").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßmodels.ToObject(), πg.NewUnicode("\xe2\x8a\xa7").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßmultimap.ToObject(), πg.NewUnicode("\xe2\x8a\xb8").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnLeftarrow.ToObject(), πg.NewUnicode("\xe2\x87\x8d").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnLeftrightarrow.ToObject(), πg.NewUnicode("\xe2\x87\x8e").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnRightarrow.ToObject(), πg.NewUnicode("\xe2\x87\x8f").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnVDash.ToObject(), πg.NewUnicode("\xe2\x8a\xaf").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnVdash.ToObject(), πg.NewUnicode("\xe2\x8a\xae").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßncong.ToObject(), πg.NewUnicode("\xe2\x89\x87").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßne.ToObject(), πg.NewUnicode("\xe2\x89\xa0").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnearrow.ToObject(), πg.NewUnicode("\xe2\x86\x97").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßneq.ToObject(), πg.NewUnicode("\xe2\x89\xa0").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßngeq.ToObject(), πg.NewUnicode("\xe2\x89\xb1").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßngtr.ToObject(), πg.NewUnicode("\xe2\x89\xaf").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßni.ToObject(), πg.NewUnicode("\xe2\x88\x8b").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnleftarrow.ToObject(), πg.NewUnicode("\xe2\x86\x9a").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnleftrightarrow.ToObject(), πg.NewUnicode("\xe2\x86\xae").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnleq.ToObject(), πg.NewUnicode("\xe2\x89\xb0").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnless.ToObject(), πg.NewUnicode("\xe2\x89\xae").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnmid.ToObject(), πg.NewUnicode("\xe2\x88\xa4").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnotasymp.ToObject(), πg.NewUnicode("\xe2\x89\xad").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnotin.ToObject(), πg.NewUnicode("\xe2\x88\x89").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnotowner.ToObject(), πg.NewUnicode("\xe2\x88\x8c").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnotslash.ToObject(), πg.NewUnicode("\xe2\x8c\xbf").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnparallel.ToObject(), πg.NewUnicode("\xe2\x88\xa6").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnprec.ToObject(), πg.NewUnicode("\xe2\x8a\x80").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnpreceq.ToObject(), πg.NewUnicode("\xe2\x8b\xa0").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnrightarrow.ToObject(), πg.NewUnicode("\xe2\x86\x9b").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnsim.ToObject(), πg.NewUnicode("\xe2\x89\x81").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnsubseteq.ToObject(), πg.NewUnicode("\xe2\x8a\x88").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnsucc.ToObject(), πg.NewUnicode("\xe2\x8a\x81").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnsucceq.ToObject(), πg.NewUnicode("\xe2\x8b\xa1").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnsupseteq.ToObject(), πg.NewUnicode("\xe2\x8a\x89").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßntriangleleft.ToObject(), πg.NewUnicode("\xe2\x8b\xaa").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßntrianglelefteq.ToObject(), πg.NewUnicode("\xe2\x8b\xac").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßntriangleright.ToObject(), πg.NewUnicode("\xe2\x8b\xab").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßntrianglerighteq.ToObject(), πg.NewUnicode("\xe2\x8b\xad").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnvDash.ToObject(), πg.NewUnicode("\xe2\x8a\xad").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnvdash.ToObject(), πg.NewUnicode("\xe2\x8a\xac").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnwarrow.ToObject(), πg.NewUnicode("\xe2\x86\x96").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßowns.ToObject(), πg.NewUnicode("\xe2\x88\x8b").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßparallel.ToObject(), πg.NewUnicode("\xe2\x88\xa5").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßperp.ToObject(), πg.NewUnicode("\xe2\x9f\x82").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßpitchfork.ToObject(), πg.NewUnicode("\xe2\x8b\x94").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßprec.ToObject(), πg.NewUnicode("\xe2\x89\xba").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßprecapprox.ToObject(), πg.NewUnicode("\xe2\xaa\xb7").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßpreccurlyeq.ToObject(), πg.NewUnicode("\xe2\x89\xbc").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßpreceq.ToObject(), πg.NewUnicode("\xe2\xaa\xaf").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßprecnapprox.ToObject(), πg.NewUnicode("\xe2\xaa\xb9").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßprecnsim.ToObject(), πg.NewUnicode("\xe2\x8b\xa8").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßprecsim.ToObject(), πg.NewUnicode("\xe2\x89\xbe").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßpropto.ToObject(), πg.NewUnicode("\xe2\x88\x9d").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrestriction.ToObject(), πg.NewUnicode("\xe2\x86\xbe").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrightarrow.ToObject(), πg.NewUnicode("\xe2\x86\x92").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrightarrowtail.ToObject(), πg.NewUnicode("\xe2\x86\xa3").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrightarrowtriangle.ToObject(), πg.NewUnicode("\xe2\x87\xbe").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrightbarharpoon.ToObject(), πg.NewUnicode("\xe2\xa5\xac").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrightharpoondown.ToObject(), πg.NewUnicode("\xe2\x87\x81").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrightharpoonup.ToObject(), πg.NewUnicode("\xe2\x87\x80").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrightleftarrows.ToObject(), πg.NewUnicode("\xe2\x87\x84").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrightleftharpoon.ToObject(), πg.NewUnicode("\xe2\xa5\x8b").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrightleftharpoons.ToObject(), πg.NewUnicode("\xe2\x87\x8c").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrightrightarrows.ToObject(), πg.NewUnicode("\xe2\x87\x89").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrightrightharpoons.ToObject(), πg.NewUnicode("\xe2\xa5\xa4").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrightslice.ToObject(), πg.NewUnicode("\xe2\xaa\xa7").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrightsquigarrow.ToObject(), πg.NewUnicode("\xe2\x87\x9d").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrisingdotseq.ToObject(), πg.NewUnicode("\xe2\x89\x93").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsearrow.ToObject(), πg.NewUnicode("\xe2\x86\x98").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsim.ToObject(), πg.NewUnicode("\xe2\x88\xbc").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsimeq.ToObject(), πg.NewUnicode("\xe2\x89\x83").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsmallfrown.ToObject(), πg.NewUnicode("\xe2\x8c\xa2").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsmallsmile.ToObject(), πg.NewUnicode("\xe2\x8c\xa3").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsmile.ToObject(), πg.NewUnicode("\xe2\x8c\xa3").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsqsubset.ToObject(), πg.NewUnicode("\xe2\x8a\x8f").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsqsubseteq.ToObject(), πg.NewUnicode("\xe2\x8a\x91").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsqsupset.ToObject(), πg.NewUnicode("\xe2\x8a\x90").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsqsupseteq.ToObject(), πg.NewUnicode("\xe2\x8a\x92").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsubset.ToObject(), πg.NewUnicode("\xe2\x8a\x82").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsubseteq.ToObject(), πg.NewUnicode("\xe2\x8a\x86").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsubseteqq.ToObject(), πg.NewUnicode("\xe2\xab\x85").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsubsetneq.ToObject(), πg.NewUnicode("\xe2\x8a\x8a").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsubsetneqq.ToObject(), πg.NewUnicode("\xe2\xab\x8b").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsucc.ToObject(), πg.NewUnicode("\xe2\x89\xbb").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsuccapprox.ToObject(), πg.NewUnicode("\xe2\xaa\xb8").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsucccurlyeq.ToObject(), πg.NewUnicode("\xe2\x89\xbd").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsucceq.ToObject(), πg.NewUnicode("\xe2\xaa\xb0").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsuccnapprox.ToObject(), πg.NewUnicode("\xe2\xaa\xba").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsuccnsim.ToObject(), πg.NewUnicode("\xe2\x8b\xa9").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsuccsim.ToObject(), πg.NewUnicode("\xe2\x89\xbf").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsupset.ToObject(), πg.NewUnicode("\xe2\x8a\x83").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsupseteq.ToObject(), πg.NewUnicode("\xe2\x8a\x87").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsupseteqq.ToObject(), πg.NewUnicode("\xe2\xab\x86").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsupsetneq.ToObject(), πg.NewUnicode("\xe2\x8a\x8b").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsupsetneqq.ToObject(), πg.NewUnicode("\xe2\xab\x8c").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßswarrow.ToObject(), πg.NewUnicode("\xe2\x86\x99").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßto.ToObject(), πg.NewUnicode("\xe2\x86\x92").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtrianglelefteq.ToObject(), πg.NewUnicode("\xe2\x8a\xb4").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtriangleq.ToObject(), πg.NewUnicode("\xe2\x89\x9c").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtrianglerighteq.ToObject(), πg.NewUnicode("\xe2\x8a\xb5").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtwoheadleftarrow.ToObject(), πg.NewUnicode("\xe2\x86\x9e").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtwoheadrightarrow.ToObject(), πg.NewUnicode("\xe2\x86\xa0").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßuparrow.ToObject(), πg.NewUnicode("\xe2\x86\x91").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßupdownarrow.ToObject(), πg.NewUnicode("\xe2\x86\x95").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßupdownarrows.ToObject(), πg.NewUnicode("\xe2\x87\x85").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßupdownharpoons.ToObject(), πg.NewUnicode("\xe2\xa5\xae").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßupharpoonleft.ToObject(), πg.NewUnicode("\xe2\x86\xbf").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßupharpoonright.ToObject(), πg.NewUnicode("\xe2\x86\xbe").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßupuparrows.ToObject(), πg.NewUnicode("\xe2\x87\x88").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßupupharpoons.ToObject(), πg.NewUnicode("\xe2\xa5\xa3").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvDash.ToObject(), πg.NewUnicode("\xe2\x8a\xa8").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvarpropto.ToObject(), πg.NewUnicode("\xe2\x88\x9d").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvartriangleleft.ToObject(), πg.NewUnicode("\xe2\x8a\xb2").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvartriangleright.ToObject(), πg.NewUnicode("\xe2\x8a\xb3").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvdash.ToObject(), πg.NewUnicode("\xe2\x8a\xa2").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvdots.ToObject(), πg.NewUnicode("\xe2\x8b\xae").ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßmathrel.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 655: mathunder = {
			πF.SetLineno(655)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, ßunderbrace.ToObject(), πg.NewUnicode("\xe2\x8f\x9f").ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßmathunder.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 658: space = {
			πF.SetLineno(658)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, πg.NewStr(":").ToObject(), πg.NewUnicode("\xe2\x81\x9f").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßmedspace.ToObject(), πg.NewUnicode("\xe2\x81\x9f").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßquad.ToObject(), πg.NewUnicode("\xe2\x80\x81").ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßspace.ToObject(), πTemp002); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("tex2unichar", Code)
}
