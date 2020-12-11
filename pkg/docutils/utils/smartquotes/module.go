package smartquotes

import (
	πg "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib/pkg/argparse"
	_ "github.com/pygolin/stdlib/pkg/itertools"
	_ "github.com/pygolin/stdlib/pkg/locale"
	_ "github.com/pygolin/stdlib/pkg/re"
	_ "github.com/pygolin/stdlib/pkg/sys"
	_ "github.com/pygolin/stdlib/pkg/unittest"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/utils/smartquotes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ß1 := πg.InternStr("1")
		ß2 := πg.InternStr("2")
		ß3 := πg.InternStr("3")
		ßArgumentParser := πg.InternStr("ArgumentParser")
		ßB := πg.InternStr("B")
		ßD := πg.InternStr("D")
		ßFalse := πg.InternStr("False")
		ßKeyError := πg.InternStr("KeyError")
		ßLC_ALL := πg.InternStr("LC_ALL")
		ßNone := πg.InternStr("None")
		ßTestCase := πg.InternStr("TestCase")
		ßTestLoader := πg.InternStr("TestLoader")
		ßTestSmartypantsAllAttributes := πg.InternStr("TestSmartypantsAllAttributes")
		ßTextTestRunner := πg.InternStr("TextTestRunner")
		ßTrue := πg.InternStr("True")
		ßUNICODE := πg.InternStr("UNICODE")
		ßVERBOSE := πg.InternStr("VERBOSE")
		ß_ := πg.InternStr("_")
		ß__doc__ := πg.InternStr("__doc__")
		ß__init__ := πg.InternStr("__init__")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_basetag := πg.InternStr("_basetag")
		ß_subtags := πg.InternStr("_subtags")
		ß_tag := πg.InternStr("_tag")
		ßaction := πg.InternStr("action")
		ßactionhelp := πg.InternStr("actionhelp")
		ßadd_argument := πg.InternStr("add_argument")
		ßaf := πg.InternStr("af")
		ßalternative_quotes := πg.InternStr("alternative_quotes")
		ßapostrophe := πg.InternStr("apostrophe")
		ßargparse := πg.InternStr("argparse")
		ßargs := πg.InternStr("args")
		ßassertEqual := πg.InternStr("assertEqual")
		ßb := πg.InternStr("b")
		ßbg := πg.InternStr("bg")
		ßca := πg.InternStr("ca")
		ßcombinations := πg.InternStr("combinations")
		ßcompile := πg.InternStr("compile")
		ßcpquote := πg.InternStr("cpquote")
		ßcs := πg.InternStr("cs")
		ßcsquote := πg.InternStr("csquote")
		ßd := πg.InternStr("d")
		ßda := πg.InternStr("da")
		ßde := πg.InternStr("de")
		ßdecode := πg.InternStr("decode")
		ßdefault_smartypants_attr := πg.InternStr("default_smartypants_attr")
		ßdefaultlanguage := πg.InternStr("defaultlanguage")
		ßdoc := πg.InternStr("doc")
		ße := πg.InternStr("e")
		ßeducateBackticks := πg.InternStr("educateBackticks")
		ßeducateDashes := πg.InternStr("educateDashes")
		ßeducateDashesOldSchool := πg.InternStr("educateDashesOldSchool")
		ßeducateDashesOldSchoolInverted := πg.InternStr("educateDashesOldSchoolInverted")
		ßeducateEllipses := πg.InternStr("educateEllipses")
		ßeducateQuotes := πg.InternStr("educateQuotes")
		ßeducateSingleBackticks := πg.InternStr("educateSingleBackticks")
		ßeducate_tokens := πg.InternStr("educate_tokens")
		ßel := πg.InternStr("el")
		ßellipsis := πg.InternStr("ellipsis")
		ßemdash := πg.InternStr("emdash")
		ßen := πg.InternStr("en")
		ßencode := πg.InternStr("encode")
		ßencoding := πg.InternStr("encoding")
		ßend := πg.InternStr("end")
		ßendash := πg.InternStr("endash")
		ßeo := πg.InternStr("eo")
		ßes := πg.InternStr("es")
		ßet := πg.InternStr("et")
		ßeu := πg.InternStr("eu")
		ßfi := πg.InternStr("fi")
		ßfr := πg.InternStr("fr")
		ßgetdefaultlocale := πg.InternStr("getdefaultlocale")
		ßgl := πg.InternStr("gl")
		ßgroup := πg.InternStr("group")
		ßhe := πg.InternStr("he")
		ßhr := πg.InternStr("hr")
		ßhsb := πg.InternStr("hsb")
		ßhu := πg.InternStr("hu")
		ßi := πg.InternStr("i")
		ßis := πg.InternStr("is")
		ßit := πg.InternStr("it")
		ßitertools := πg.InternStr("itertools")
		ßja := πg.InternStr("ja")
		ßjoin := πg.InternStr("join")
		ßkey := πg.InternStr("key")
		ßkeys := πg.InternStr("keys")
		ßko := πg.InternStr("ko")
		ßlanguage := πg.InternStr("language")
		ßlen := πg.InternStr("len")
		ßliteral := πg.InternStr("literal")
		ßloadTestsFromTestCase := πg.InternStr("loadTestsFromTestCase")
		ßlocale := πg.InternStr("locale")
		ßlower := πg.InternStr("lower")
		ßlt := πg.InternStr("lt")
		ßlv := πg.InternStr("lv")
		ßmk := πg.InternStr("mk")
		ßn := πg.InternStr("n")
		ßnb := πg.InternStr("nb")
		ßnl := πg.InternStr("nl")
		ßnn := πg.InternStr("nn")
		ßno := πg.InternStr("no")
		ßobject := πg.InternStr("object")
		ßopquote := πg.InternStr("opquote")
		ßoptions := πg.InternStr("options")
		ßosquote := πg.InternStr("osquote")
		ßparse_args := πg.InternStr("parse_args")
		ßparser := πg.InternStr("parser")
		ßpl := πg.InternStr("pl")
		ßpop := πg.InternStr("pop")
		ßprint := πg.InternStr("print")
		ßprocessEscapes := πg.InternStr("processEscapes")
		ßpt := πg.InternStr("pt")
		ßq := πg.InternStr("q")
		ßquotes := πg.InternStr("quotes")
		ßrange := πg.InternStr("range")
		ßre := πg.InternStr("re")
		ßread := πg.InternStr("read")
		ßreplace := πg.InternStr("replace")
		ßro := πg.InternStr("ro")
		ßru := πg.InternStr("ru")
		ßrun := πg.InternStr("run")
		ßsearch := πg.InternStr("search")
		ßsetlocale := πg.InternStr("setlocale")
		ßsh := πg.InternStr("sh")
		ßsk := πg.InternStr("sk")
		ßsl := πg.InternStr("sl")
		ßsmartchars := πg.InternStr("smartchars")
		ßsmartyPants := πg.InternStr("smartyPants")
		ßsorted := πg.InternStr("sorted")
		ßsplit := πg.InternStr("split")
		ßsq := πg.InternStr("sq")
		ßsr := πg.InternStr("sr")
		ßstartswith := πg.InternStr("startswith")
		ßstdin := πg.InternStr("stdin")
		ßstore_true := πg.InternStr("store_true")
		ßstupefyEntities := πg.InternStr("stupefyEntities")
		ßstylehelp := πg.InternStr("stylehelp")
		ßsub := πg.InternStr("sub")
		ßsuite := πg.InternStr("suite")
		ßsv := πg.InternStr("sv")
		ßsys := πg.InternStr("sys")
		ßtag := πg.InternStr("tag")
		ßtags := πg.InternStr("tags")
		ßtest := πg.InternStr("test")
		ßtest_dates := πg.InternStr("test_dates")
		ßtest_educated_quotes := πg.InternStr("test_educated_quotes")
		ßtest_html_tags := πg.InternStr("test_html_tags")
		ßtext := πg.InternStr("text")
		ßtokenize := πg.InternStr("tokenize")
		ßtr := πg.InternStr("tr")
		ßuk := πg.InternStr("uk")
		ßunittest := πg.InternStr("unittest")
		ßutf8 := πg.InternStr("utf8")
		ßw := πg.InternStr("w")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Dict
		_ = πTemp003
		var πTemp004 *πg.Object
		_ = πTemp004
		var πTemp005 *πg.Object
		_ = πTemp005
		var πTemp006 []πg.Param
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
		var πTemp012 *πg.Object
		_ = πTemp012
		var πTemp013 *πg.Object
		_ = πTemp013
		var πTemp014 *πg.Object
		_ = πTemp014
		var πTemp015 *πg.Object
		_ = πTemp015
		var πTemp016 *πg.Object
		_ = πTemp016
		var πTemp017 *πg.Object
		_ = πTemp017
		var πTemp018 bool
		_ = πTemp018
		var πTemp019 *πg.Object
		_ = πTemp019
		var πTemp020 *πg.Object
		_ = πTemp020
		var πTemp021 *πg.BaseException
		_ = πTemp021
		var πTemp022 *πg.Traceback
		_ = πTemp022
		var πTemp023 []*πg.Object
		_ = πTemp023
		var πTemp024 bool
		_ = πTemp024
		var πTemp025 *πg.Object
		_ = πTemp025
		var πTemp026 bool
		_ = πTemp026
		var πTemp027 *πg.Object
		_ = πTemp027
		var πTemp028 *πg.Object
		_ = πTemp028
		var πTemp029 πg.KWArgs
		_ = πTemp029
		var πTemp030 *πg.Object
		_ = πTemp030
		var πTemp031 *πg.Object
		_ = πTemp031
		var πTemp032 []*πg.Object
		_ = πTemp032
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 4:
				goto Label4
			case 6:
				goto Label6
			case 7:
				goto Label7
			case 9:
				goto Label9
			case 10:
				goto Label10
			case 23:
				goto Label23
			case 24:
				goto Label24
			default:
				panic("unexpected function state")
			}
			// line 19: r"""
			πF.SetLineno(19)
			// line 19: r"""
			πF.SetLineno(19)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n=========================\nSmart Quotes for Docutils\n=========================\n\nSynopsis\n========\n\n\"SmartyPants\" is a free web publishing plug-in for Movable Type, Blosxom, and\nBBEdit that easily translates plain ASCII punctuation characters into \"smart\"\ntypographic punctuation characters.\n\n``smartquotes.py`` is an adaption of \"SmartyPants\" to Docutils_.\n\n* Using Unicode instead of HTML entities for typographic punctuation\n  characters, it works for any output format that supports Unicode.\n* Supports `language specific quote characters`__.\n\n__ http://en.wikipedia.org/wiki/Non-English_usage_of_quotation_marks\n\n\nAuthors\n=======\n\n`John Gruber`_ did all of the hard work of writing this software in Perl for\n`Movable Type`_ and almost all of this useful documentation.  `Chad Miller`_\nported it to Python to use with Pyblosxom_.\nAdapted to Docutils_ by G\xc3\xbcnter Milde.\n\nAdditional Credits\n==================\n\nPortions of the SmartyPants original work are based on Brad Choate's nifty\nMTRegex plug-in.  `Brad Choate`_ also contributed a few bits of source code to\nthis plug-in.  Brad Choate is a fine hacker indeed.\n\n`Jeremy Hedley`_ and `Charles Wiltgen`_ deserve mention for exemplary beta\ntesting of the original SmartyPants.\n\n`Rael Dornfest`_ ported SmartyPants to Blosxom.\n\n.. _Brad Choate: http://bradchoate.com/\n.. _Jeremy Hedley: http://antipixel.com/\n.. _Charles Wiltgen: http://playbacktime.com/\n.. _Rael Dornfest: http://raelity.org/\n\n\nCopyright and License\n=====================\n\nSmartyPants_ license (3-Clause BSD license):\n\n  Copyright (c) 2003 John Gruber (http://daringfireball.net/)\n  All rights reserved.\n\n  Redistribution and use in source and binary forms, with or without\n  modification, are permitted provided that the following conditions are\n  met:\n\n  * Redistributions of source code must retain the above copyright\n    notice, this list of conditions and the following disclaimer.\n\n  * Redistributions in binary form must reproduce the above copyright\n    notice, this list of conditions and the following disclaimer in\n    the documentation and/or other materials provided with the\n    distribution.\n\n  * Neither the name \"SmartyPants\" nor the names of its contributors\n    may be used to endorse or promote products derived from this\n    software without specific prior written permission.\n\n  This software is provided by the copyright holders and contributors\n  \"as is\" and any express or implied warranties, including, but not\n  limited to, the implied warranties of merchantability and fitness for\n  a particular purpose are disclaimed. In no event shall the copyright\n  owner or contributors be liable for any direct, indirect, incidental,\n  special, exemplary, or consequential damages (including, but not\n  limited to, procurement of substitute goods or services; loss of use,\n  data, or profits; or business interruption) however caused and on any\n  theory of liability, whether in contract, strict liability, or tort\n  (including negligence or otherwise) arising in any way out of the use\n  of this software, even if advised of the possibility of such damage.\n\nsmartypants.py license (2-Clause BSD license):\n\n  smartypants.py is a derivative work of SmartyPants.\n\n  Redistribution and use in source and binary forms, with or without\n  modification, are permitted provided that the following conditions are\n  met:\n\n  * Redistributions of source code must retain the above copyright\n    notice, this list of conditions and the following disclaimer.\n\n  * Redistributions in binary form must reproduce the above copyright\n    notice, this list of conditions and the following disclaimer in\n    the documentation and/or other materials provided with the\n    distribution.\n\n  This software is provided by the copyright holders and contributors\n  \"as is\" and any express or implied warranties, including, but not\n  limited to, the implied warranties of merchantability and fitness for\n  a particular purpose are disclaimed. In no event shall the copyright\n  owner or contributors be liable for any direct, indirect, incidental,\n  special, exemplary, or consequential damages (including, but not\n  limited to, procurement of substitute goods or services; loss of use,\n  data, or profits; or business interruption) however caused and on any\n  theory of liability, whether in contract, strict liability, or tort\n  (including negligence or otherwise) arising in any way out of the use\n  of this software, even if advised of the possibility of such damage.\n\n.. _John Gruber: http://daringfireball.net/\n.. _Chad Miller: http://web.chad.org/\n\n.. _Pyblosxom: http://pyblosxom.bluesock.org/\n.. _SmartyPants: http://daringfireball.net/projects/smartypants/\n.. _Movable Type: http://www.movabletype.org/\n.. _2-Clause BSD license: http://www.spdx.org/licenses/BSD-2-Clause\n.. _Docutils: http://docutils.sf.net/\n\nDescription\n===========\n\nSmartyPants can perform the following transformations:\n\n- Straight quotes ( \" and ' ) into \"curly\" quote characters\n- Backticks-style quotes (\\`\\`like this'') into \"curly\" quote characters\n- Dashes (``--`` and ``---``) into en- and em-dash entities\n- Three consecutive dots (``...`` or ``. . .``) into an ellipsis entity\n\nThis means you can write, edit, and save your posts using plain old\nASCII straight quotes, plain dashes, and plain dots, but your published\nposts (and final HTML output) will appear with smart quotes, em-dashes,\nand proper ellipses.\n\nSmartyPants does not modify characters within ``<pre>``, ``<code>``, ``<kbd>``,\n``<math>`` or ``<script>`` tag blocks. Typically, these tags are used to\ndisplay text where smart quotes and other \"smart punctuation\" would not be\nappropriate, such as source code or example markup.\n\n\nBackslash Escapes\n=================\n\nIf you need to use literal straight quotes (or plain hyphens and periods),\n`smartquotes` accepts the following backslash escape sequences to force\nASCII-punctuation. Mind, that you need two backslashes as Docutils expands it,\ntoo.\n\n========  =========\nEscape    Character\n========  =========\n``\\\\``    \\\\\n``\\\\\"``   \\\\\"\n``\\\\'``   \\\\'\n``\\\\.``   \\\\.\n``\\\\-``   \\\\-\n``\\\\```   \\\\`\n========  =========\n\nThis is useful, for example, when you want to use straight quotes as\nfoot and inch marks: 6\\\\'2\\\\\" tall; a 17\\\\\" iMac.\n\n\nCaveats\n=======\n\nWhy You Might Not Want to Use Smart Quotes in Your Weblog\n---------------------------------------------------------\n\nFor one thing, you might not care.\n\nMost normal, mentally stable individuals do not take notice of proper\ntypographic punctuation. Many design and typography nerds, however, break\nout in a nasty rash when they encounter, say, a restaurant sign that uses\na straight apostrophe to spell \"Joe's\".\n\nIf you're the sort of person who just doesn't care, you might well want to\ncontinue not caring. Using straight quotes -- and sticking to the 7-bit\nASCII character set in general -- is certainly a simpler way to live.\n\nEven if you *do* care about accurate typography, you still might want to\nthink twice before educating the quote characters in your weblog. One side\neffect of publishing curly quote characters is that it makes your\nweblog a bit harder for others to quote from using copy-and-paste. What\nhappens is that when someone copies text from your blog, the copied text\ncontains the 8-bit curly quote characters (as well as the 8-bit characters\nfor em-dashes and ellipses, if you use these options). These characters\nare not standard across different text encoding methods, which is why they\nneed to be encoded as characters.\n\nPeople copying text from your weblog, however, may not notice that you're\nusing curly quotes, and they'll go ahead and paste the unencoded 8-bit\ncharacters copied from their browser into an email message or their own\nweblog. When pasted as raw \"smart quotes\", these characters are likely to\nget mangled beyond recognition.\n\nThat said, my own opinion is that any decent text editor or email client\nmakes it easy to stupefy smart quote characters into their 7-bit\nequivalents, and I don't consider it my problem if you're using an\nindecent text editor or email client.\n\n\nAlgorithmic Shortcomings\n------------------------\n\nOne situation in which quotes will get curled the wrong way is when\napostrophes are used at the start of leading contractions. For example::\n\n  'Twas the night before Christmas.\n\nIn the case above, SmartyPants will turn the apostrophe into an opening\nsingle-quote, when in fact it should be the `right single quotation mark`\ncharacter which is also \"the preferred character to use for apostrophe\"\n(Unicode). I don't think this problem can be solved in the general case --\nevery word processor I've tried gets this wrong as well. In such cases, it's\nbest to use the proper character for closing single-quotes (\xe2\x80\x99) by hand.\n\nIn English, the same character is used for apostrophe and  closing single\nquote (both plain and \"smart\" ones). For other locales (French, Italean,\nSwiss, ...) \"smart\" single closing quotes differ from the curly apostrophe.\n\n   .. class:: language-fr\n\n   Il dit : \"C'est 'super' !\"\n\nIf the apostrophe is used at the end of a word, it cannot be distinguished\nfrom a single quote by the algorithm. Therefore, a text like::\n\n   .. class:: language-de-CH\n\n   \"Er sagt: 'Ich fass' es nicht.'\"\n\nwill get a single closing guillemet instead of an apostrophe.\n\nThis can be prevented by use use of the curly apostrophe character (\xe2\x80\x99) in\nthe source::\n\n   -  \"Er sagt: 'Ich fass' es nicht.'\"\n   +  \"Er sagt: 'Ich fass\xe2\x80\x99 es nicht.'\"\n\n\nVersion History\n===============\n\n1.8.1   2017-10-25\n        - Use open quote after Unicode whitespace, ZWSP, and ZWNJ.\n        - Code cleanup.\n\n1.8:    2017-04-24\n        - Command line front-end.\n\n1.7.1:  2017-03-19\n        - Update and extend language-dependent quotes.\n        - Differentiate apostrophe from single quote.\n\n1.7:    2012-11-19\n        - Internationalization: language-dependent quotes.\n\n1.6.1:  2012-11-06\n        - Refactor code, code cleanup,\n        - `educate_tokens()` generator as interface for Docutils.\n\n1.6:    2010-08-26\n        - Adaption to Docutils:\n          - Use Unicode instead of HTML entities,\n          - Remove code special to pyblosxom.\n\n1.5_1.6: Fri, 27 Jul 2007 07:06:40 -0400\n        - Fixed bug where blocks of precious unalterable text was instead\n          interpreted.  Thanks to Le Roux and Dirk van Oosterbosch.\n\n1.5_1.5: Sat, 13 Aug 2005 15:50:24 -0400\n        - Fix bogus magical quotation when there is no hint that the\n          user wants it, e.g., in \"21st century\".  Thanks to Nathan Hamblen.\n        - Be smarter about quotes before terminating numbers in an en-dash'ed\n          range.\n\n1.5_1.4: Thu, 10 Feb 2005 20:24:36 -0500\n        - Fix a date-processing bug, as reported by jacob childress.\n        - Begin a test-suite for ensuring correct output.\n        - Removed import of \"string\", since I didn't really need it.\n          (This was my first every Python program.  Sue me!)\n\n1.5_1.3: Wed, 15 Sep 2004 18:25:58 -0400\n        - Abort processing if the flavour is in forbidden-list.  Default of\n          [ \"rss\" ]   (Idea of Wolfgang SCHNERRING.)\n        - Remove stray virgules from en-dashes.  Patch by Wolfgang SCHNERRING.\n\n1.5_1.2: Mon, 24 May 2004 08:14:54 -0400\n        - Some single quotes weren't replaced properly.  Diff-tesuji played\n          by Benjamin GEIGER.\n\n1.5_1.1: Sun, 14 Mar 2004 14:38:28 -0500\n        - Support upcoming pyblosxom 0.9 plugin verification feature.\n\n1.5_1.0: Tue, 09 Mar 2004 08:08:35 -0500\n        - Initial release\n").ToObject()); πE != nil {
				continue
			}
			// line 318: from __future__ import print_function
			πF.SetLineno(318)
			// line 320: options = r"""
			πF.SetLineno(320)
			if πE = πF.Globals().SetItem(πF, ßoptions.ToObject(), πg.NewStr("\nOptions\n=======\n\nNumeric values are the easiest way to configure SmartyPants' behavior:\n\n:0:     Suppress all transformations. (Do nothing.)\n\n:1:     Performs default SmartyPants transformations: quotes (including\n        \\`\\`backticks'' -style), em-dashes, and ellipses. \"``--``\" (dash dash)\n        is used to signify an em-dash; there is no support for en-dashes\n\n:2:     Same as smarty_pants=\"1\", except that it uses the old-school typewriter\n        shorthand for dashes:  \"``--``\" (dash dash) for en-dashes, \"``---``\"\n        (dash dash dash)\n        for em-dashes.\n\n:3:     Same as smarty_pants=\"2\", but inverts the shorthand for dashes:\n        \"``--``\" (dash dash) for em-dashes, and \"``---``\" (dash dash dash) for\n        en-dashes.\n\n:-1:    Stupefy mode. Reverses the SmartyPants transformation process, turning\n        the characters produced by SmartyPants into their ASCII equivalents.\n        E.g. the LEFT DOUBLE QUOTATION MARK (\xe2\x80\x9c) is turned into a simple\n        double-quote (\\\"), \"\xe2\x80\x94\" is turned into two dashes, etc.\n\n\nThe following single-character attribute values can be combined to toggle\nindividual transformations from within the smarty_pants attribute. For\nexample, ``\"1\"`` is equivalent to ``\"qBde\"``.\n\n:q:     Educates normal quote characters: (\") and (').\n\n:b:     Educates \\`\\`backticks'' -style double quotes.\n\n:B:     Educates \\`\\`backticks'' -style double quotes and \\`single' quotes.\n\n:d:     Educates em-dashes.\n\n:D:     Educates em-dashes and en-dashes, using old-school typewriter shorthand:\n        (dash dash) for en-dashes, (dash dash dash) for em-dashes.\n\n:i:     Educates em-dashes and en-dashes, using inverted old-school typewriter\n        shorthand: (dash dash) for em-dashes, (dash dash dash) for en-dashes.\n\n:e:     Educates ellipses.\n\n:w:     Translates any instance of ``&quot;`` into a normal double-quote character.\n        This should be of no interest to most people, but of particular interest\n        to anyone who writes their posts using Dreamweaver, as Dreamweaver\n        inexplicably uses this entity to represent a literal double-quote\n        character. SmartyPants only educates normal quotes, not entities (because\n        ordinarily, entities are used for the explicit purpose of representing the\n        specific character they represent). The \"w\" option must be used in\n        conjunction with one (or both) of the other quote options (\"q\" or \"b\").\n        Thus, if you wish to apply all SmartyPants transformations (quotes, en-\n        and em-dashes, and ellipses) and also translate ``&quot;`` entities into\n        regular quotes so SmartyPants can educate them, you should pass the\n        following to the smarty_pants attribute:\n").ToObject()); πE != nil {
				continue
			}
			// line 382: default_smartypants_attr = "1"
			πF.SetLineno(382)
			if πE = πF.Globals().SetItem(πF, ßdefault_smartypants_attr.ToObject(), ß1.ToObject()); πE != nil {
				continue
			}
			// line 385: import re, sys
			πF.SetLineno(385)
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 387: class smartchars(object):
			πF.SetLineno(387)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("smartchars", "/usr/lib/python2.7/site-packages/docutils/utils/smartquotes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 388: """Smart quotes and dashes
					πF.SetLineno(388)
					// line 388: """Smart quotes and dashes
					πF.SetLineno(388)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Smart quotes and dashes\n    ").ToObject()); πE != nil {
						continue
					}
					// line 391: endash   = u'–' # "&#8211;" EN DASH
					πF.SetLineno(391)
					if πE = πClass.SetItem(πF, ßendash.ToObject(), πg.NewUnicode("\xe2\x80\x93").ToObject()); πE != nil {
						continue
					}
					// line 392: emdash   = u'—' # "&#8212;" EM DASH
					πF.SetLineno(392)
					if πE = πClass.SetItem(πF, ßemdash.ToObject(), πg.NewUnicode("\xe2\x80\x94").ToObject()); πE != nil {
						continue
					}
					// line 393: ellipsis = u'…' # "&#8230;" HORIZONTAL ELLIPSIS
					πF.SetLineno(393)
					if πE = πClass.SetItem(πF, ßellipsis.ToObject(), πg.NewUnicode("\xe2\x80\xa6").ToObject()); πE != nil {
						continue
					}
					// line 394: apostrophe = u'’' # "&#8217;" RIGHT SINGLE QUOTATION MARK
					πF.SetLineno(394)
					if πE = πClass.SetItem(πF, ßapostrophe.ToObject(), πg.NewUnicode("\xe2\x80\x99").ToObject()); πE != nil {
						continue
					}
					// line 409: quotes = {'af':           u'“”‘’',
					πF.SetLineno(409)
					πTemp001 = πg.NewDict()
					if πE = πTemp001.SetItem(πF, ßaf.ToObject(), πg.NewUnicode("\xe2\x80\x9c\xe2\x80\x9d\xe2\x80\x98\xe2\x80\x99").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("af-x-altquot").ToObject(), πg.NewUnicode("\xe2\x80\x9e\xe2\x80\x9d\xe2\x80\x9a\xe2\x80\x99").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßbg.ToObject(), πg.NewUnicode("\xe2\x80\x9e\xe2\x80\x9c\xe2\x80\x9a\xe2\x80\x98").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßca.ToObject(), πg.NewUnicode("\xc2\xab\xc2\xbb\xe2\x80\x9c\xe2\x80\x9d").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("ca-x-altquot").ToObject(), πg.NewUnicode("\xe2\x80\x9c\xe2\x80\x9d\xe2\x80\x98\xe2\x80\x99").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßcs.ToObject(), πg.NewUnicode("\xe2\x80\x9e\xe2\x80\x9c\xe2\x80\x9a\xe2\x80\x98").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("cs-x-altquot").ToObject(), πg.NewUnicode("\xc2\xbb\xc2\xab\xe2\x80\xba\xe2\x80\xb9").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßda.ToObject(), πg.NewUnicode("\xc2\xbb\xc2\xab\xe2\x80\xba\xe2\x80\xb9").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("da-x-altquot").ToObject(), πg.NewUnicode("\xe2\x80\x9e\xe2\x80\x9c\xe2\x80\x9a\xe2\x80\x98").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßde.ToObject(), πg.NewUnicode("\xe2\x80\x9e\xe2\x80\x9c\xe2\x80\x9a\xe2\x80\x98").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("de-x-altquot").ToObject(), πg.NewUnicode("\xc2\xbb\xc2\xab\xe2\x80\xba\xe2\x80\xb9").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("de-ch").ToObject(), πg.NewUnicode("\xc2\xab\xc2\xbb\xe2\x80\xb9\xe2\x80\xba").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßel.ToObject(), πg.NewUnicode("\xc2\xab\xc2\xbb\xe2\x80\x9c\xe2\x80\x9d").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßen.ToObject(), πg.NewUnicode("\xe2\x80\x9c\xe2\x80\x9d\xe2\x80\x98\xe2\x80\x99").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("en-uk-x-altquot").ToObject(), πg.NewUnicode("\xe2\x80\x98\xe2\x80\x99\xe2\x80\x9c\xe2\x80\x9d").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßeo.ToObject(), πg.NewUnicode("\xe2\x80\x9c\xe2\x80\x9d\xe2\x80\x98\xe2\x80\x99").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßes.ToObject(), πg.NewUnicode("\xc2\xab\xc2\xbb\xe2\x80\x9c\xe2\x80\x9d").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("es-x-altquot").ToObject(), πg.NewUnicode("\xe2\x80\x9c\xe2\x80\x9d\xe2\x80\x98\xe2\x80\x99").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßet.ToObject(), πg.NewUnicode("\xe2\x80\x9e\xe2\x80\x9c\xe2\x80\x9a\xe2\x80\x98").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("et-x-altquot").ToObject(), πg.NewUnicode("\xc2\xab\xc2\xbb\xe2\x80\xb9\xe2\x80\xba").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßeu.ToObject(), πg.NewUnicode("\xc2\xab\xc2\xbb\xe2\x80\xb9\xe2\x80\xba").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßfi.ToObject(), πg.NewUnicode("\xe2\x80\x9d\xe2\x80\x9d\xe2\x80\x99\xe2\x80\x99").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("fi-x-altquot").ToObject(), πg.NewUnicode("\xc2\xbb\xc2\xbb\xe2\x80\xba\xe2\x80\xba").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple4(πg.NewUnicode("\xc2\xab\xc2\xa0").ToObject(), πg.NewUnicode("\xc2\xa0\xc2\xbb").ToObject(), πg.NewUnicode("\xe2\x80\x9c").ToObject(), πg.NewUnicode("\xe2\x80\x9d").ToObject()).ToObject()
					if πE = πTemp001.SetItem(πF, ßfr.ToObject(), πTemp002); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple4(πg.NewUnicode("\xc2\xab\xe2\x80\xaf").ToObject(), πg.NewUnicode("\xe2\x80\xaf\xc2\xbb").ToObject(), πg.NewUnicode("\xe2\x80\x9c").ToObject(), πg.NewUnicode("\xe2\x80\x9d").ToObject()).ToObject()
					if πE = πTemp001.SetItem(πF, πg.NewStr("fr-x-altquot").ToObject(), πTemp002); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("fr-ch").ToObject(), πg.NewUnicode("\xc2\xab\xc2\xbb\xe2\x80\xb9\xe2\x80\xba").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple4(πg.NewUnicode("\xc2\xab\xe2\x80\xaf").ToObject(), πg.NewUnicode("\xe2\x80\xaf\xc2\xbb").ToObject(), πg.NewUnicode("\xe2\x80\xb9\xe2\x80\xaf").ToObject(), πg.NewUnicode("\xe2\x80\xaf\xe2\x80\xba").ToObject()).ToObject()
					if πE = πTemp001.SetItem(πF, πg.NewStr("fr-ch-x-altquot").ToObject(), πTemp002); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßgl.ToObject(), πg.NewUnicode("\xc2\xab\xc2\xbb\xe2\x80\x9c\xe2\x80\x9d").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßhe.ToObject(), πg.NewUnicode("\xe2\x80\x9d\xe2\x80\x9c\xc2\xbb\xc2\xab").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("he-x-altquot").ToObject(), πg.NewUnicode("\xe2\x80\x9e\xe2\x80\x9d\xe2\x80\x9a\xe2\x80\x99").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßhr.ToObject(), πg.NewUnicode("\xe2\x80\x9e\xe2\x80\x9d\xe2\x80\x98\xe2\x80\x99").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("hr-x-altquot").ToObject(), πg.NewUnicode("\xc2\xbb\xc2\xab\xe2\x80\xba\xe2\x80\xb9").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßhsb.ToObject(), πg.NewUnicode("\xe2\x80\x9e\xe2\x80\x9c\xe2\x80\x9a\xe2\x80\x98").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("hsb-x-altquot").ToObject(), πg.NewUnicode("\xc2\xbb\xc2\xab\xe2\x80\xba\xe2\x80\xb9").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßhu.ToObject(), πg.NewUnicode("\xe2\x80\x9e\xe2\x80\x9d\xc2\xab\xc2\xbb").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßis.ToObject(), πg.NewUnicode("\xe2\x80\x9e\xe2\x80\x9c\xe2\x80\x9a\xe2\x80\x98").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßit.ToObject(), πg.NewUnicode("\xc2\xab\xc2\xbb\xe2\x80\x9c\xe2\x80\x9d").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("it-ch").ToObject(), πg.NewUnicode("\xc2\xab\xc2\xbb\xe2\x80\xb9\xe2\x80\xba").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("it-x-altquot").ToObject(), πg.NewUnicode("\xe2\x80\x9c\xe2\x80\x9d\xe2\x80\x98\xe2\x80\x99").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßja.ToObject(), πg.NewUnicode("\xe3\x80\x8c\xe3\x80\x8d\xe3\x80\x8e\xe3\x80\x8f").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßko.ToObject(), πg.NewUnicode("\xe3\x80\x8a\xe3\x80\x8b\xe3\x80\x88\xe3\x80\x89").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßlt.ToObject(), πg.NewUnicode("\xe2\x80\x9e\xe2\x80\x9c\xe2\x80\x9a\xe2\x80\x98").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßlv.ToObject(), πg.NewUnicode("\xe2\x80\x9e\xe2\x80\x9c\xe2\x80\x9a\xe2\x80\x98").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßmk.ToObject(), πg.NewUnicode("\xe2\x80\x9e\xe2\x80\x9c\xe2\x80\x9a\xe2\x80\x98").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßnl.ToObject(), πg.NewUnicode("\xe2\x80\x9c\xe2\x80\x9d\xe2\x80\x98\xe2\x80\x99").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("nl-x-altquot").ToObject(), πg.NewUnicode("\xe2\x80\x9e\xe2\x80\x9d\xe2\x80\x9a\xe2\x80\x99").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßnb.ToObject(), πg.NewUnicode("\xc2\xab\xc2\xbb\xe2\x80\x99\xe2\x80\x99").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßnn.ToObject(), πg.NewUnicode("\xc2\xab\xc2\xbb\xe2\x80\x99\xe2\x80\x99").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("nn-x-altquot").ToObject(), πg.NewUnicode("\xc2\xab\xc2\xbb\xe2\x80\x98\xe2\x80\x99").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßno.ToObject(), πg.NewUnicode("\xc2\xab\xc2\xbb\xe2\x80\x99\xe2\x80\x99").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("no-x-altquot").ToObject(), πg.NewUnicode("\xc2\xab\xc2\xbb\xe2\x80\x98\xe2\x80\x99").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßpl.ToObject(), πg.NewUnicode("\xe2\x80\x9e\xe2\x80\x9d\xc2\xab\xc2\xbb").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("pl-x-altquot").ToObject(), πg.NewUnicode("\xc2\xab\xc2\xbb\xe2\x80\x9a\xe2\x80\x99").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßpt.ToObject(), πg.NewUnicode("\xc2\xab\xc2\xbb\xe2\x80\x9c\xe2\x80\x9d").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("pt-br").ToObject(), πg.NewUnicode("\xe2\x80\x9c\xe2\x80\x9d\xe2\x80\x98\xe2\x80\x99").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßro.ToObject(), πg.NewUnicode("\xe2\x80\x9e\xe2\x80\x9d\xc2\xab\xc2\xbb").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßru.ToObject(), πg.NewUnicode("\xc2\xab\xc2\xbb\xe2\x80\x9e\xe2\x80\x9c").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßsh.ToObject(), πg.NewUnicode("\xe2\x80\x9e\xe2\x80\x9d\xe2\x80\x9a\xe2\x80\x99").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("sh-x-altquot").ToObject(), πg.NewUnicode("\xc2\xbb\xc2\xab\xe2\x80\xba\xe2\x80\xb9").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßsk.ToObject(), πg.NewUnicode("\xe2\x80\x9e\xe2\x80\x9c\xe2\x80\x9a\xe2\x80\x98").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("sk-x-altquot").ToObject(), πg.NewUnicode("\xc2\xbb\xc2\xab\xe2\x80\xba\xe2\x80\xb9").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßsl.ToObject(), πg.NewUnicode("\xe2\x80\x9e\xe2\x80\x9c\xe2\x80\x9a\xe2\x80\x98").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("sl-x-altquot").ToObject(), πg.NewUnicode("\xc2\xbb\xc2\xab\xe2\x80\xba\xe2\x80\xb9").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßsq.ToObject(), πg.NewUnicode("\xc2\xab\xc2\xbb\xe2\x80\xb9\xe2\x80\xba").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("sq-x-altquot").ToObject(), πg.NewUnicode("\xe2\x80\x9c\xe2\x80\x9e\xe2\x80\x98\xe2\x80\x9a").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßsr.ToObject(), πg.NewUnicode("\xe2\x80\x9e\xe2\x80\x9d\xe2\x80\x99\xe2\x80\x99").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("sr-x-altquot").ToObject(), πg.NewUnicode("\xc2\xbb\xc2\xab\xe2\x80\xba\xe2\x80\xb9").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßsv.ToObject(), πg.NewUnicode("\xe2\x80\x9d\xe2\x80\x9d\xe2\x80\x99\xe2\x80\x99").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("sv-x-altquot").ToObject(), πg.NewUnicode("\xc2\xbb\xc2\xbb\xe2\x80\xba\xe2\x80\xba").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßtr.ToObject(), πg.NewUnicode("\xe2\x80\x9c\xe2\x80\x9d\xe2\x80\x98\xe2\x80\x99").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("tr-x-altquot").ToObject(), πg.NewUnicode("\xc2\xab\xc2\xbb\xe2\x80\xb9\xe2\x80\xba").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, ßuk.ToObject(), πg.NewUnicode("\xc2\xab\xc2\xbb\xe2\x80\x9e\xe2\x80\x9c").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("uk-x-altquot").ToObject(), πg.NewUnicode("\xe2\x80\x9e\xe2\x80\x9c\xe2\x80\x9a\xe2\x80\x98").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("zh-cn").ToObject(), πg.NewUnicode("\xe2\x80\x9c\xe2\x80\x9d\xe2\x80\x98\xe2\x80\x99").ToObject()); πE != nil {
						continue
					}
					if πE = πTemp001.SetItem(πF, πg.NewStr("zh-tw").ToObject(), πg.NewUnicode("\xe3\x80\x8c\xe3\x80\x8d\xe3\x80\x8e\xe3\x80\x8f").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp001.ToObject()
					if πE = πClass.SetItem(πF, ßquotes.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 496: def __init__(self, language='en'):
					πF.SetLineno(496)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "language", Def: ßen.ToObject()}
					πTemp002 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/site-packages/docutils/utils/smartquotes.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µlanguage *πg.Object = πArgs[1]
						_ = µlanguage
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
						var πTemp006 *πg.BaseException
						_ = πTemp006
						var πTemp007 *πg.Traceback
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2:
								goto Label2
							default:
								panic("unexpected function state")
							}
							// line 497: self.language = language
							πF.SetLineno(497)
							if πE = πg.CheckLocal(πF, µlanguage, "language"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µlanguage); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlanguage, πTemp001); πE != nil {
								continue
							}
							// line 498: try:
							πF.SetLineno(498)
							πF.PushCheckpoint(2)
							// line 499: (self.opquote, self.cpquote,
							πF.SetLineno(499)
							if πE = πg.CheckLocal(πF, µlanguage, "language"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlanguage, ßlower, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßquotes, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßopquote, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcpquote, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßosquote, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcsquote, πTemp005); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp006, πTemp007 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label3
							}
							πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
							continue
							// line 501: except KeyError:
							πF.SetLineno(501)
						Label3:
							// line 502: self.opquote, self.cpquote, self.osquote, self.csquote = u'""\'\''
							πF.SetLineno(502)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πg.NewUnicode("\"\"''").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßopquote, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcpquote, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßosquote, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcsquote, πTemp004); πE != nil {
								continue
							}
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
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
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp004, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("smartchars").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßsmartchars.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 505: def smartyPants(text, attr=default_smartypants_attr, language='en'):
			πF.SetLineno(505)
			πTemp006 = make([]πg.Param, 3)
			πTemp006[0] = πg.Param{Name: "text", Def: nil}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßdefault_smartypants_attr); πE != nil {
				continue
			}
			πTemp006[1] = πg.Param{Name: "attr", Def: πTemp004}
			πTemp006[2] = πg.Param{Name: "language", Def: ßen.ToObject()}
			πTemp001 = πg.NewFunction(πg.NewCode("smartyPants", "/usr/lib/python2.7/site-packages/docutils/utils/smartquotes.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtext *πg.Object = πArgs[0]
				_ = µtext
				var µattr *πg.Object = πArgs[1]
				_ = µattr
				var µlanguage *πg.Object = πArgs[2]
				_ = µlanguage
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []πg.Param
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
					// line 506: """Main function for "traditional" use."""
					πF.SetLineno(506)
					// line 508: return "".join([t for t in educate_tokens(tokenize(text),
					πF.SetLineno(508)
					πTemp001 = πF.MakeArgs(1)
					πTemp004 = make([]πg.Param, 0)
					πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/utils/smartquotes.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µt *πg.Object = πg.UnboundLocal
						_ = µt
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
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
								πTemp002 = πF.MakeArgs(3)
								πTemp003 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
									continue
								}
								πTemp003[0] = µtext
								if πTemp004, πE = πg.ResolveGlobal(πF, ßtokenize); πE != nil {
									continue
								}
								if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp003)
								πTemp002[0] = πTemp005
								if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
									continue
								}
								πTemp002[1] = µattr
								if πE = πg.CheckLocal(πF, µlanguage, "language"); πE != nil {
									continue
								}
								πTemp002[2] = µlanguage
								if πTemp004, πE = πg.ResolveGlobal(πF, ßeducate_tokens); πE != nil {
									continue
								}
								if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
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
								if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
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
									µt = πTemp004
								}
								if πE != nil || !πTemp007 {
									continue
								}
								πF.PushCheckpoint(1)
								// line 508: return "".join([t for t in educate_tokens(tokenize(text),
								πF.SetLineno(508)
								if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
									continue
								}
								πF.PushCheckpoint(4)
								return µt, nil
							Label4:
								πTemp004 = πSent
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
					if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp005
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßsmartyPants.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 506: """Main function for "traditional" use."""
			πF.SetLineno(506)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Main function for \"traditional\" use.").ToObject()); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßsmartyPants); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
				continue
			}
			// line 512: def educate_tokens(text_tokens, attr=default_smartypants_attr, language='en'):
			πF.SetLineno(512)
			πTemp006 = make([]πg.Param, 3)
			πTemp006[0] = πg.Param{Name: "text_tokens", Def: nil}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßdefault_smartypants_attr); πE != nil {
				continue
			}
			πTemp006[1] = πg.Param{Name: "attr", Def: πTemp005}
			πTemp006[2] = πg.Param{Name: "language", Def: ßen.ToObject()}
			πTemp004 = πg.NewFunction(πg.NewCode("educate_tokens", "/usr/lib/python2.7/site-packages/docutils/utils/smartquotes.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtext_tokens *πg.Object = πArgs[0]
				_ = µtext_tokens
				var µattr *πg.Object = πArgs[1]
				_ = µattr
				var µlanguage *πg.Object = πArgs[2]
				_ = µlanguage
				var µconvert_quot *πg.Object = πg.UnboundLocal
				_ = µconvert_quot
				var µdo_dashes *πg.Object = πg.UnboundLocal
				_ = µdo_dashes
				var µdo_backticks *πg.Object = πg.UnboundLocal
				_ = µdo_backticks
				var µdo_quotes *πg.Object = πg.UnboundLocal
				_ = µdo_quotes
				var µdo_ellipses *πg.Object = πg.UnboundLocal
				_ = µdo_ellipses
				var µdo_stupefy *πg.Object = πg.UnboundLocal
				_ = µdo_stupefy
				var µprev_token_last_char *πg.Object = πg.UnboundLocal
				_ = µprev_token_last_char
				var µttype *πg.Object = πg.UnboundLocal
				_ = µttype
				var µtext *πg.Object = πg.UnboundLocal
				_ = µtext
				var µlast_char *πg.Object = πg.UnboundLocal
				_ = µlast_char
				var µcontext *πg.Object = πg.UnboundLocal
				_ = µcontext
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 []*πg.Object
				_ = πTemp008
				var πTemp009 []*πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 πg.KWArgs
				_ = πTemp011
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						case 24:
							goto Label24
						case 32:
							goto Label32
						case 49:
							goto Label49
						case 29:
							goto Label29
						case 23:
							goto Label23
						default:
							panic("unexpected function state")
						}
						// line 513: """Return iterator that "educates" the items of `text_tokens`.
						πF.SetLineno(513)
						// line 531: convert_quot = False  # translate &quot; entities into normal quotes?
						πF.SetLineno(531)
						if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
							continue
						}
						µconvert_quot = πTemp001
						// line 532: do_dashes = False
						πF.SetLineno(532)
						if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
							continue
						}
						µdo_dashes = πTemp001
						// line 533: do_backticks = False
						πF.SetLineno(533)
						if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
							continue
						}
						µdo_backticks = πTemp001
						// line 534: do_quotes = False
						πF.SetLineno(534)
						if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
							continue
						}
						µdo_quotes = πTemp001
						// line 535: do_ellipses = False
						πF.SetLineno(535)
						if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
							continue
						}
						µdo_ellipses = πTemp001
						// line 536: do_stupefy = False
						πF.SetLineno(536)
						if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
							continue
						}
						µdo_stupefy = πTemp001
						if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
							continue
						}
						if πTemp001, πE = πg.Eq(πF, µattr, ß1.ToObject()); πE != nil {
							continue
						}
						if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
							continue
						}
						if πTemp002 {
							goto Label1
						}
						if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
							continue
						}
						if πTemp001, πE = πg.Eq(πF, µattr, ß2.ToObject()); πE != nil {
							continue
						}
						if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
							continue
						}
						if πTemp002 {
							goto Label2
						}
						if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
							continue
						}
						if πTemp001, πE = πg.Eq(πF, µattr, ß3.ToObject()); πE != nil {
							continue
						}
						if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
							continue
						}
						if πTemp002 {
							goto Label3
						}
						if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
							continue
						}
						if πTemp001, πE = πg.Eq(πF, µattr, πg.NewStr("-1").ToObject()); πE != nil {
							continue
						}
						if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
							continue
						}
						if πTemp002 {
							goto Label4
						}
						goto Label5
						// line 539: if attr == "1": # Do everything, turn all options on.
						πF.SetLineno(539)
					Label1:
						// line 540: do_quotes    = True
						πF.SetLineno(540)
						if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
							continue
						}
						µdo_quotes = πTemp001
						// line 541: do_backticks = True
						πF.SetLineno(541)
						if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
							continue
						}
						µdo_backticks = πTemp001
						// line 542: do_dashes    = 1
						πF.SetLineno(542)
						µdo_dashes = πg.NewInt(1).ToObject()
						// line 543: do_ellipses  = True
						πF.SetLineno(543)
						if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
							continue
						}
						µdo_ellipses = πTemp001
						goto Label6
						// line 544: elif attr == "2":
						πF.SetLineno(544)
					Label2:
						// line 546: do_quotes    = True
						πF.SetLineno(546)
						if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
							continue
						}
						µdo_quotes = πTemp001
						// line 547: do_backticks = True
						πF.SetLineno(547)
						if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
							continue
						}
						µdo_backticks = πTemp001
						// line 548: do_dashes    = 2
						πF.SetLineno(548)
						µdo_dashes = πg.NewInt(2).ToObject()
						// line 549: do_ellipses  = True
						πF.SetLineno(549)
						if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
							continue
						}
						µdo_ellipses = πTemp001
						goto Label6
						// line 550: elif attr == "3":
						πF.SetLineno(550)
					Label3:
						// line 552: do_quotes    = True
						πF.SetLineno(552)
						if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
							continue
						}
						µdo_quotes = πTemp001
						// line 553: do_backticks = True
						πF.SetLineno(553)
						if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
							continue
						}
						µdo_backticks = πTemp001
						// line 554: do_dashes    = 3
						πF.SetLineno(554)
						µdo_dashes = πg.NewInt(3).ToObject()
						// line 555: do_ellipses  = True
						πF.SetLineno(555)
						if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
							continue
						}
						µdo_ellipses = πTemp001
						goto Label6
						// line 556: elif attr == "-1": # Special "stupefy" mode.
						πF.SetLineno(556)
					Label4:
						// line 557: do_stupefy   = True
						πF.SetLineno(557)
						if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
							continue
						}
						µdo_stupefy = πTemp001
						goto Label6
					Label5:
						if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.Contains(πF, µattr, ßq.ToObject()); πE != nil {
							continue
						}
						πTemp001 = πg.GetBool(πTemp002).ToObject()
						if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
							continue
						}
						if πTemp002 {
							goto Label7
						}
						goto Label8
						// line 559: if "q" in attr: do_quotes = True
						πF.SetLineno(559)
					Label7:
						// line 559: if "q" in attr: do_quotes = True
						πF.SetLineno(559)
						if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
							continue
						}
						µdo_quotes = πTemp001
						goto Label8
					Label8:
						if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.Contains(πF, µattr, ßb.ToObject()); πE != nil {
							continue
						}
						πTemp001 = πg.GetBool(πTemp002).ToObject()
						if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
							continue
						}
						if πTemp002 {
							goto Label9
						}
						goto Label10
						// line 560: if "b" in attr: do_backticks = True
						πF.SetLineno(560)
					Label9:
						// line 560: if "b" in attr: do_backticks = True
						πF.SetLineno(560)
						if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
							continue
						}
						µdo_backticks = πTemp001
						goto Label10
					Label10:
						if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.Contains(πF, µattr, ßB.ToObject()); πE != nil {
							continue
						}
						πTemp001 = πg.GetBool(πTemp002).ToObject()
						if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
							continue
						}
						if πTemp002 {
							goto Label11
						}
						goto Label12
						// line 561: if "B" in attr: do_backticks = 2
						πF.SetLineno(561)
					Label11:
						// line 561: if "B" in attr: do_backticks = 2
						πF.SetLineno(561)
						µdo_backticks = πg.NewInt(2).ToObject()
						goto Label12
					Label12:
						if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.Contains(πF, µattr, ßd.ToObject()); πE != nil {
							continue
						}
						πTemp001 = πg.GetBool(πTemp002).ToObject()
						if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
							continue
						}
						if πTemp002 {
							goto Label13
						}
						goto Label14
						// line 562: if "d" in attr: do_dashes = 1
						πF.SetLineno(562)
					Label13:
						// line 562: if "d" in attr: do_dashes = 1
						πF.SetLineno(562)
						µdo_dashes = πg.NewInt(1).ToObject()
						goto Label14
					Label14:
						if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.Contains(πF, µattr, ßD.ToObject()); πE != nil {
							continue
						}
						πTemp001 = πg.GetBool(πTemp002).ToObject()
						if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
							continue
						}
						if πTemp002 {
							goto Label15
						}
						goto Label16
						// line 563: if "D" in attr: do_dashes = 2
						πF.SetLineno(563)
					Label15:
						// line 563: if "D" in attr: do_dashes = 2
						πF.SetLineno(563)
						µdo_dashes = πg.NewInt(2).ToObject()
						goto Label16
					Label16:
						if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.Contains(πF, µattr, ßi.ToObject()); πE != nil {
							continue
						}
						πTemp001 = πg.GetBool(πTemp002).ToObject()
						if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
							continue
						}
						if πTemp002 {
							goto Label17
						}
						goto Label18
						// line 564: if "i" in attr: do_dashes = 3
						πF.SetLineno(564)
					Label17:
						// line 564: if "i" in attr: do_dashes = 3
						πF.SetLineno(564)
						µdo_dashes = πg.NewInt(3).ToObject()
						goto Label18
					Label18:
						if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.Contains(πF, µattr, ße.ToObject()); πE != nil {
							continue
						}
						πTemp001 = πg.GetBool(πTemp002).ToObject()
						if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
							continue
						}
						if πTemp002 {
							goto Label19
						}
						goto Label20
						// line 565: if "e" in attr: do_ellipses = True
						πF.SetLineno(565)
					Label19:
						// line 565: if "e" in attr: do_ellipses = True
						πF.SetLineno(565)
						if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
							continue
						}
						µdo_ellipses = πTemp001
						goto Label20
					Label20:
						if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.Contains(πF, µattr, ßw.ToObject()); πE != nil {
							continue
						}
						πTemp001 = πg.GetBool(πTemp002).ToObject()
						if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
							continue
						}
						if πTemp002 {
							goto Label21
						}
						goto Label22
						// line 566: if "w" in attr: convert_quot = True
						πF.SetLineno(566)
					Label21:
						// line 566: if "w" in attr: convert_quot = True
						πF.SetLineno(566)
						if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
							continue
						}
						µconvert_quot = πTemp001
						goto Label22
					Label22:
						goto Label6
					Label6:
						// line 568: prev_token_last_char = " "
						πF.SetLineno(568)
						µprev_token_last_char = πg.NewStr(" ").ToObject()
						if πE = πg.CheckLocal(πF, µtext_tokens, "text_tokens"); πE != nil {
							continue
						}
						if πTemp001, πE = πg.Iter(πF, µtext_tokens); πE != nil {
							continue
						}
						πF.PushCheckpoint(24)
						πTemp002 = false
					Label23:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp002 {
							πF.PopCheckpoint()
							goto Label25
						}
						if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
							isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
							if exc != nil {
								πE = exc
							} else if isStop {
								πE = nil
								πF.RestoreExc(nil, nil)
							}
							πTemp003 = !isStop
						} else {
							πTemp003 = true
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp004); πE != nil {
								continue
							}
							µttype = πTemp005
							µtext = πTemp006
						}
						if πE != nil || !πTemp003 {
							continue
						}
						πF.PushCheckpoint(23)
						if πE = πg.CheckLocal(πF, µttype, "ttype"); πE != nil {
							continue
						}
						if πTemp005, πE = πg.Eq(πF, µttype, ßtag.ToObject()); πE != nil {
							continue
						}
						πTemp004 = πTemp005
						if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
							continue
						}
						if πTemp003 {
							goto Label26
						}
						if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
							continue
						}
						if πTemp007, πE = πg.IsTrue(πF, µtext); πE != nil {
							continue
						}
						πTemp005 = πg.GetBool(!πTemp007).ToObject()
						πTemp004 = πTemp005
					Label26:
						if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
							continue
						}
						if πTemp003 {
							goto Label27
						}
						goto Label28
						// line 576: if ttype == 'tag' or not text:
						πF.SetLineno(576)
					Label27:
						// line 577: yield text
						πF.SetLineno(577)
						if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
							continue
						}
						πF.PushCheckpoint(29)
						return µtext, nil
					Label29:
						πTemp004 = πSent
						// line 578: continue
						πF.SetLineno(578)
						continue
						goto Label28
					Label28:
						if πE = πg.CheckLocal(πF, µttype, "ttype"); πE != nil {
							continue
						}
						if πTemp004, πE = πg.Eq(πF, µttype, ßliteral.ToObject()); πE != nil {
							continue
						}
						if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
							continue
						}
						if πTemp003 {
							goto Label30
						}
						goto Label31
						// line 581: if ttype == 'literal':
						πF.SetLineno(581)
					Label30:
						// line 582: prev_token_last_char = text[-1:]
						πF.SetLineno(582)
						if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
							continue
						}
						if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πTemp005, πg.None, πg.None}, nil); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
							continue
						}
						if πTemp005, πE = πg.GetItem(πF, µtext, πTemp004); πE != nil {
							continue
						}
						µprev_token_last_char = πTemp005
						// line 583: yield text
						πF.SetLineno(583)
						if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
							continue
						}
						πF.PushCheckpoint(32)
						return µtext, nil
					Label32:
						πTemp004 = πSent
						// line 584: continue
						πF.SetLineno(584)
						continue
						goto Label31
					Label31:
						// line 586: last_char = text[-1:] # Remember last char before processing.
						πF.SetLineno(586)
						if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
							continue
						}
						if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πTemp005, πg.None, πg.None}, nil); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
							continue
						}
						if πTemp005, πE = πg.GetItem(πF, µtext, πTemp004); πE != nil {
							continue
						}
						µlast_char = πTemp005
						// line 588: text = processEscapes(text)
						πF.SetLineno(588)
						πTemp008 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
							continue
						}
						πTemp008[0] = µtext
						if πTemp004, πE = πg.ResolveGlobal(πF, ßprocessEscapes); πE != nil {
							continue
						}
						if πTemp005, πE = πTemp004.Call(πF, πTemp008, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp008)
						µtext = πTemp005
						if πE = πg.CheckLocal(πF, µconvert_quot, "convert_quot"); πE != nil {
							continue
						}
						if πTemp003, πE = πg.IsTrue(πF, µconvert_quot); πE != nil {
							continue
						}
						if πTemp003 {
							goto Label33
						}
						goto Label34
						// line 590: if convert_quot:
						πF.SetLineno(590)
					Label33:
						// line 591: text = re.sub('&quot;', '"', text)
						πF.SetLineno(591)
						πTemp008 = πF.MakeArgs(3)
						πTemp008[0] = πg.NewStr("&quot;").ToObject()
						πTemp008[1] = πg.NewStr("\"").ToObject()
						if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
							continue
						}
						πTemp008[2] = µtext
						if πTemp004, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
							continue
						}
						if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßsub, nil); πE != nil {
							continue
						}
						if πTemp004, πE = πTemp005.Call(πF, πTemp008, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp008)
						µtext = πTemp004
						goto Label34
					Label34:
						if πE = πg.CheckLocal(πF, µdo_dashes, "do_dashes"); πE != nil {
							continue
						}
						if πTemp004, πE = πg.Eq(πF, µdo_dashes, πg.NewInt(1).ToObject()); πE != nil {
							continue
						}
						if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
							continue
						}
						if πTemp003 {
							goto Label35
						}
						if πE = πg.CheckLocal(πF, µdo_dashes, "do_dashes"); πE != nil {
							continue
						}
						if πTemp004, πE = πg.Eq(πF, µdo_dashes, πg.NewInt(2).ToObject()); πE != nil {
							continue
						}
						if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
							continue
						}
						if πTemp003 {
							goto Label36
						}
						if πE = πg.CheckLocal(πF, µdo_dashes, "do_dashes"); πE != nil {
							continue
						}
						if πTemp004, πE = πg.Eq(πF, µdo_dashes, πg.NewInt(3).ToObject()); πE != nil {
							continue
						}
						if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
							continue
						}
						if πTemp003 {
							goto Label37
						}
						goto Label38
						// line 593: if do_dashes == 1:
						πF.SetLineno(593)
					Label35:
						// line 594: text = educateDashes(text)
						πF.SetLineno(594)
						πTemp008 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
							continue
						}
						πTemp008[0] = µtext
						if πTemp004, πE = πg.ResolveGlobal(πF, ßeducateDashes); πE != nil {
							continue
						}
						if πTemp005, πE = πTemp004.Call(πF, πTemp008, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp008)
						µtext = πTemp005
						goto Label38
						// line 595: elif do_dashes == 2:
						πF.SetLineno(595)
					Label36:
						// line 596: text = educateDashesOldSchool(text)
						πF.SetLineno(596)
						πTemp008 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
							continue
						}
						πTemp008[0] = µtext
						if πTemp004, πE = πg.ResolveGlobal(πF, ßeducateDashesOldSchool); πE != nil {
							continue
						}
						if πTemp005, πE = πTemp004.Call(πF, πTemp008, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp008)
						µtext = πTemp005
						goto Label38
						// line 597: elif do_dashes == 3:
						πF.SetLineno(597)
					Label37:
						// line 598: text = educateDashesOldSchoolInverted(text)
						πF.SetLineno(598)
						πTemp008 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
							continue
						}
						πTemp008[0] = µtext
						if πTemp004, πE = πg.ResolveGlobal(πF, ßeducateDashesOldSchoolInverted); πE != nil {
							continue
						}
						if πTemp005, πE = πTemp004.Call(πF, πTemp008, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp008)
						µtext = πTemp005
						goto Label38
					Label38:
						if πE = πg.CheckLocal(πF, µdo_ellipses, "do_ellipses"); πE != nil {
							continue
						}
						if πTemp003, πE = πg.IsTrue(πF, µdo_ellipses); πE != nil {
							continue
						}
						if πTemp003 {
							goto Label39
						}
						goto Label40
						// line 600: if do_ellipses:
						πF.SetLineno(600)
					Label39:
						// line 601: text = educateEllipses(text)
						πF.SetLineno(601)
						πTemp008 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
							continue
						}
						πTemp008[0] = µtext
						if πTemp004, πE = πg.ResolveGlobal(πF, ßeducateEllipses); πE != nil {
							continue
						}
						if πTemp005, πE = πTemp004.Call(πF, πTemp008, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp008)
						µtext = πTemp005
						goto Label40
					Label40:
						if πE = πg.CheckLocal(πF, µdo_backticks, "do_backticks"); πE != nil {
							continue
						}
						if πTemp003, πE = πg.IsTrue(πF, µdo_backticks); πE != nil {
							continue
						}
						if πTemp003 {
							goto Label41
						}
						goto Label42
						// line 604: if do_backticks:
						πF.SetLineno(604)
					Label41:
						// line 605: text = educateBackticks(text, language)
						πF.SetLineno(605)
						πTemp008 = πF.MakeArgs(2)
						if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
							continue
						}
						πTemp008[0] = µtext
						if πE = πg.CheckLocal(πF, µlanguage, "language"); πE != nil {
							continue
						}
						πTemp008[1] = µlanguage
						if πTemp004, πE = πg.ResolveGlobal(πF, ßeducateBackticks); πE != nil {
							continue
						}
						if πTemp005, πE = πTemp004.Call(πF, πTemp008, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp008)
						µtext = πTemp005
						goto Label42
					Label42:
						if πE = πg.CheckLocal(πF, µdo_backticks, "do_backticks"); πE != nil {
							continue
						}
						if πTemp004, πE = πg.Eq(πF, µdo_backticks, πg.NewInt(2).ToObject()); πE != nil {
							continue
						}
						if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
							continue
						}
						if πTemp003 {
							goto Label43
						}
						goto Label44
						// line 607: if do_backticks == 2:
						πF.SetLineno(607)
					Label43:
						// line 608: text = educateSingleBackticks(text, language)
						πF.SetLineno(608)
						πTemp008 = πF.MakeArgs(2)
						if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
							continue
						}
						πTemp008[0] = µtext
						if πE = πg.CheckLocal(πF, µlanguage, "language"); πE != nil {
							continue
						}
						πTemp008[1] = µlanguage
						if πTemp004, πE = πg.ResolveGlobal(πF, ßeducateSingleBackticks); πE != nil {
							continue
						}
						if πTemp005, πE = πTemp004.Call(πF, πTemp008, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp008)
						µtext = πTemp005
						goto Label44
					Label44:
						if πE = πg.CheckLocal(πF, µdo_quotes, "do_quotes"); πE != nil {
							continue
						}
						if πTemp003, πE = πg.IsTrue(πF, µdo_quotes); πE != nil {
							continue
						}
						if πTemp003 {
							goto Label45
						}
						goto Label46
						// line 610: if do_quotes:
						πF.SetLineno(610)
					Label45:
						// line 613: context = prev_token_last_char.replace('"', ';').replace("'", ';')
						πF.SetLineno(613)
						πTemp008 = πF.MakeArgs(2)
						πTemp008[0] = πg.NewStr("'").ToObject()
						πTemp008[1] = πg.NewStr(";").ToObject()
						πTemp009 = πF.MakeArgs(2)
						πTemp009[0] = πg.NewStr("\"").ToObject()
						πTemp009[1] = πg.NewStr(";").ToObject()
						if πE = πg.CheckLocal(πF, µprev_token_last_char, "prev_token_last_char"); πE != nil {
							continue
						}
						if πTemp004, πE = πg.GetAttr(πF, µprev_token_last_char, ßreplace, nil); πE != nil {
							continue
						}
						if πTemp005, πE = πTemp004.Call(πF, πTemp009, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp009)
						if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßreplace, nil); πE != nil {
							continue
						}
						if πTemp005, πE = πTemp004.Call(πF, πTemp008, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp008)
						µcontext = πTemp005
						// line 614: text = educateQuotes(context+text, language)[1:]
						πF.SetLineno(614)
						if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
							continue
						}
						πTemp008 = πF.MakeArgs(2)
						if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
							continue
						}
						if πTemp006, πE = πg.Add(πF, µcontext, µtext); πE != nil {
							continue
						}
						πTemp008[0] = πTemp006
						if πE = πg.CheckLocal(πF, µlanguage, "language"); πE != nil {
							continue
						}
						πTemp008[1] = µlanguage
						if πTemp006, πE = πg.ResolveGlobal(πF, ßeducateQuotes); πE != nil {
							continue
						}
						if πTemp010, πE = πTemp006.Call(πF, πTemp008, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp008)
						if πTemp005, πE = πg.GetItem(πF, πTemp010, πTemp004); πE != nil {
							continue
						}
						µtext = πTemp005
						goto Label46
					Label46:
						if πE = πg.CheckLocal(πF, µdo_stupefy, "do_stupefy"); πE != nil {
							continue
						}
						if πTemp003, πE = πg.IsTrue(πF, µdo_stupefy); πE != nil {
							continue
						}
						if πTemp003 {
							goto Label47
						}
						goto Label48
						// line 616: if do_stupefy:
						πF.SetLineno(616)
					Label47:
						// line 617: text = stupefyEntities(text, language)
						πF.SetLineno(617)
						πTemp008 = πF.MakeArgs(2)
						if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
							continue
						}
						πTemp008[0] = µtext
						if πE = πg.CheckLocal(πF, µlanguage, "language"); πE != nil {
							continue
						}
						πTemp008[1] = µlanguage
						if πTemp004, πE = πg.ResolveGlobal(πF, ßstupefyEntities); πE != nil {
							continue
						}
						if πTemp005, πE = πTemp004.Call(πF, πTemp008, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp008)
						µtext = πTemp005
						goto Label48
					Label48:
						// line 620: prev_token_last_char = last_char
						πF.SetLineno(620)
						if πE = πg.CheckLocal(πF, µlast_char, "last_char"); πE != nil {
							continue
						}
						µprev_token_last_char = µlast_char
						// line 622: text = processEscapes(text, restore=True)
						πF.SetLineno(622)
						πTemp008 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
							continue
						}
						πTemp008[0] = µtext
						if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
							continue
						}
						πTemp011 = πg.KWArgs{
							{"restore", πTemp004},
						}
						if πTemp004, πE = πg.ResolveGlobal(πF, ßprocessEscapes); πE != nil {
							continue
						}
						if πTemp005, πE = πTemp004.Call(πF, πTemp008, πTemp011); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp008)
						µtext = πTemp005
						// line 624: yield text
						πF.SetLineno(624)
						if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
							continue
						}
						πF.PushCheckpoint(49)
						return µtext, nil
					Label49:
						πTemp004 = πSent
						continue
					Label24:
						if πE != nil || πR != nil {
							continue
						}
					Label25:
					}
					return nil, πE
				}).ToObject(), nil
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßeducate_tokens.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 513: """Return iterator that "educates" the items of `text_tokens`.
			πF.SetLineno(513)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Return iterator that \"educates\" the items of `text_tokens`.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßeducate_tokens); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp005); πE != nil {
				continue
			}
			// line 628: def educateQuotes(text, language='en'):
			πF.SetLineno(628)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "text", Def: nil}
			πTemp006[1] = πg.Param{Name: "language", Def: ßen.ToObject()}
			πTemp005 = πg.NewFunction(πg.NewCode("educateQuotes", "/usr/lib/python2.7/site-packages/docutils/utils/smartquotes.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtext *πg.Object = πArgs[0]
				_ = µtext
				var µlanguage *πg.Object = πArgs[1]
				_ = µlanguage
				var µsmart *πg.Object = πg.UnboundLocal
				_ = µsmart
				var µpunct_class *πg.Object = πg.UnboundLocal
				_ = µpunct_class
				var µclose_class *πg.Object = πg.UnboundLocal
				_ = µclose_class
				var µopen_class *πg.Object = πg.UnboundLocal
				_ = µopen_class
				var µdec_dashes *πg.Object = πg.UnboundLocal
				_ = µdec_dashes
				var µopening_single_quotes_regex *πg.Object = πg.UnboundLocal
				_ = µopening_single_quotes_regex
				var µapostrophe_regex *πg.Object = πg.UnboundLocal
				_ = µapostrophe_regex
				var µclosing_single_quotes_regex *πg.Object = πg.UnboundLocal
				_ = µclosing_single_quotes_regex
				var µopening_double_quotes_regex *πg.Object = πg.UnboundLocal
				_ = µopening_double_quotes_regex
				var µclosing_double_quotes_regex *πg.Object = πg.UnboundLocal
				_ = µclosing_double_quotes_regex
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
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
					default:
						panic("unexpected function state")
					}
					// line 629: """
					πF.SetLineno(629)
					// line 638: smart = smartchars(language)
					πF.SetLineno(638)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlanguage, "language"); πE != nil {
						continue
					}
					πTemp001[0] = µlanguage
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsmartchars); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µsmart = πTemp003
					// line 640: punct_class = r"""[!"#\$\%'()*+,-.\/:;<=>?\@\[\\\]\^_`{|}~]"""
					πF.SetLineno(640)
					µpunct_class = πg.NewStr("[!\"#\\$\\%'()*+,-.\\/:;<=>?\\@\\[\\\\\\]\\^_`{|}~]").ToObject()
					// line 641: close_class = r"""[^\ \t\r\n\[\{\(\-]"""
					πF.SetLineno(641)
					µclose_class = πg.NewStr("[^\\ \\t\\r\\n\\[\\{\\(\\-]").ToObject()
					// line 642: open_class = u'[\u200B\u200C]' # ZWSP, ZWNJ
					πF.SetLineno(642)
					µopen_class = πg.NewUnicode("[\xe2\x80\x8b\xe2\x80\x8c]").ToObject()
					// line 643: dec_dashes = r"""&#8211;|&#8212;"""
					πF.SetLineno(643)
					µdec_dashes = πg.NewStr("&#8211;|&#8212;").ToObject()
					// line 648: text = re.sub(r"^'(?=%s\\B)" % (punct_class,), smart.csquote, text)
					πF.SetLineno(648)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µpunct_class, "punct_class"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple1(µpunct_class).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("^'(?=%s\\\\B)").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µsmart, "smart"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsmart, ßcsquote, nil); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[2] = µtext
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsub, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp002
					// line 649: text = re.sub(r'^"(?=%s\\B)' % (punct_class,), smart.cpquote, text)
					πF.SetLineno(649)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µpunct_class, "punct_class"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple1(µpunct_class).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("^\"(?=%s\\\\B)").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µsmart, "smart"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsmart, ßcpquote, nil); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[2] = µtext
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsub, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp002
					// line 653: text = re.sub(r""""'(?=\w)""", smart.opquote+smart.osquote, text)
					πF.SetLineno(653)
					πTemp001 = πF.MakeArgs(3)
					πTemp001[0] = πg.NewStr("\"'(?=\\w)").ToObject()
					if πE = πg.CheckLocal(πF, µsmart, "smart"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µsmart, ßopquote, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsmart, "smart"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µsmart, ßosquote, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πTemp003, πTemp004); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[2] = µtext
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsub, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp002
					// line 654: text = re.sub(r"""'"(?=\w)""", smart.osquote+smart.opquote, text)
					πF.SetLineno(654)
					πTemp001 = πF.MakeArgs(3)
					πTemp001[0] = πg.NewStr("'\"(?=\\w)").ToObject()
					if πE = πg.CheckLocal(πF, µsmart, "smart"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µsmart, ßosquote, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsmart, "smart"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µsmart, ßopquote, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πTemp003, πTemp004); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[2] = µtext
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsub, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp002
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = ßen.ToObject()
					if πE = πg.CheckLocal(πF, µlanguage, "language"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µlanguage, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					goto Label2
					// line 657: if language.startswith('en'): # TODO similar cases in other languages?
					πF.SetLineno(657)
				Label1:
					// line 658: text = re.sub(r"'(?=\d{2}s)", smart.apostrophe, text)
					πF.SetLineno(658)
					πTemp001 = πF.MakeArgs(3)
					πTemp001[0] = πg.NewStr("'(?=\\d{2}s)").ToObject()
					if πE = πg.CheckLocal(πF, µsmart, "smart"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsmart, ßapostrophe, nil); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[2] = µtext
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsub, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp002
					goto Label2
				Label2:
					// line 661: opening_single_quotes_regex = re.compile(u"""
					πF.SetLineno(661)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µopen_class, "open_class"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdec_dashes, "dec_dashes"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µopen_class, µdec_dashes).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewUnicode("\n                    (# ?<=  # look behind fails: requires fixed-width pattern\n                            \\s         |   # a whitespace char, or\n                            %s          |   # another separating char, or\n                            &nbsp;      |   # a non-breaking space entity, or\n                            [\xe2\x80\x93 \xe2\x80\x94 ]        |   # literal dashes, or\n                            --          |   # dumb dashes, or\n                            &[mn]dash;  |   # dash entities (named or\n                            %s          |   # decimal or\n                            &\\#x201[34];    # hex)\n                    )\n                    '                 # the quote\n                    (?=\\w)           # followed by a word character\n                    ").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp003, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßVERBOSE, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßUNICODE, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Or(πF, πTemp004, πTemp006); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µopening_single_quotes_regex = πTemp002
					// line 676: text = opening_single_quotes_regex.sub(r'\1'+smart.osquote, text)
					πF.SetLineno(676)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µsmart, "smart"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µsmart, ßosquote, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πg.NewStr("\\1").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[1] = µtext
					if πE = πg.CheckLocal(πF, µopening_single_quotes_regex, "opening_single_quotes_regex"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µopening_single_quotes_regex, ßsub, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp003
					if πE = πg.CheckLocal(πF, µsmart, "smart"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µsmart, ßcsquote, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsmart, "smart"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µsmart, ßapostrophe, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.NE(πF, πTemp003, πTemp004); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label3
					}
					goto Label4
					// line 679: if smart.csquote != smart.apostrophe:
					πF.SetLineno(679)
				Label3:
					// line 680: apostrophe_regex = re.compile(r"(?<=(\w|\d))'(?=\w)", re.UNICODE)
					πF.SetLineno(680)
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = πg.NewStr("(?<=(\\w|\\d))'(?=\\w)").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßUNICODE, nil); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µapostrophe_regex = πTemp002
					// line 681: text = apostrophe_regex.sub(smart.apostrophe, text)
					πF.SetLineno(681)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µsmart, "smart"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsmart, ßapostrophe, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[1] = µtext
					if πE = πg.CheckLocal(πF, µapostrophe_regex, "apostrophe_regex"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µapostrophe_regex, ßsub, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp003
					goto Label4
				Label4:
					// line 685: closing_single_quotes_regex = re.compile(r"""
					πF.SetLineno(685)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µclose_class, "close_class"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("\n                    (?<=%s)\n                    '\n                    ").ToObject(), µclose_class); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßVERBOSE, nil); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µclosing_single_quotes_regex = πTemp002
					// line 689: text = closing_single_quotes_regex.sub(smart.csquote, text)
					πF.SetLineno(689)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µsmart, "smart"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsmart, ßcsquote, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[1] = µtext
					if πE = πg.CheckLocal(πF, µclosing_single_quotes_regex, "closing_single_quotes_regex"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µclosing_single_quotes_regex, ßsub, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp003
					// line 692: text = re.sub(r"""'""", smart.osquote, text)
					πF.SetLineno(692)
					πTemp001 = πF.MakeArgs(3)
					πTemp001[0] = πg.NewStr("'").ToObject()
					if πE = πg.CheckLocal(πF, µsmart, "smart"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsmart, ßosquote, nil); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[2] = µtext
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsub, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp002
					// line 695: opening_double_quotes_regex = re.compile(u"""
					πF.SetLineno(695)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µopen_class, "open_class"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdec_dashes, "dec_dashes"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µopen_class, µdec_dashes).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewUnicode("\n                    (\n                            \\s         |   # a whitespace char, or\n                            %s          |   # another separating char, or\n                            &nbsp;      |   # a non-breaking space entity, or\n                            [\xe2\x80\x93 \xe2\x80\x94 ]        |   # literal dashes, or\n                            --          |   # dumb dashes, or\n                            &[mn]dash;  |   # dash entities (named or\n                            %s          |   # decimal or\n                            &\\#x201[34];    # hex)\n                    )\n                    \"                 # the quote\n                    (?=\\w)            # followed by a word character\n                    ").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp003, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßVERBOSE, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßUNICODE, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Or(πF, πTemp004, πTemp006); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µopening_double_quotes_regex = πTemp002
					// line 710: text = opening_double_quotes_regex.sub(r'\1'+smart.opquote, text)
					πF.SetLineno(710)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µsmart, "smart"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µsmart, ßopquote, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πg.NewStr("\\1").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[1] = µtext
					if πE = πg.CheckLocal(πF, µopening_double_quotes_regex, "opening_double_quotes_regex"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µopening_double_quotes_regex, ßsub, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp003
					// line 713: closing_double_quotes_regex = re.compile(r"""
					πF.SetLineno(713)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µclose_class, "close_class"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple1(µclose_class).ToObject()
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("\n                    (\n                    (?<=%s)\" | # char indicating the quote should be closing\n                    \"(?=\\s)    # whitespace behind\n                    )\n                    ").ToObject(), πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp003, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßVERBOSE, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßUNICODE, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Or(πF, πTemp004, πTemp006); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µclosing_double_quotes_regex = πTemp002
					// line 719: text = closing_double_quotes_regex.sub(smart.cpquote, text)
					πF.SetLineno(719)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µsmart, "smart"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsmart, ßcpquote, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[1] = µtext
					if πE = πg.CheckLocal(πF, µclosing_double_quotes_regex, "closing_double_quotes_regex"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µclosing_double_quotes_regex, ßsub, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp003
					// line 722: text = re.sub(r'"', smart.opquote, text)
					πF.SetLineno(722)
					πTemp001 = πF.MakeArgs(3)
					πTemp001[0] = πg.NewStr("\"").ToObject()
					if πE = πg.CheckLocal(πF, µsmart, "smart"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsmart, ßopquote, nil); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[2] = µtext
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsub, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp002
					// line 724: return text
					πF.SetLineno(724)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πR = µtext
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßeducateQuotes.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 629: """
			πF.SetLineno(629)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("\n    Parameter:  - text string (unicode or bytes).\n                - language (`BCP 47` language tag.)\n    Returns:    The `text`, with \"educated\" curly quote characters.\n\n    Example input:  \"Isn't this fun?\"\n    Example output: \xe2\x80\x9cIsn\xe2\x80\x99t this fun?\xe2\x80\x9c;\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßeducateQuotes); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
				continue
			}
			// line 727: def educateBackticks(text, language='en'):
			πF.SetLineno(727)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "text", Def: nil}
			πTemp006[1] = πg.Param{Name: "language", Def: ßen.ToObject()}
			πTemp007 = πg.NewFunction(πg.NewCode("educateBackticks", "/usr/lib/python2.7/site-packages/docutils/utils/smartquotes.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtext *πg.Object = πArgs[0]
				_ = µtext
				var µlanguage *πg.Object = πArgs[1]
				_ = µlanguage
				var µsmart *πg.Object = πg.UnboundLocal
				_ = µsmart
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
					// line 728: """
					πF.SetLineno(728)
					// line 735: smart = smartchars(language)
					πF.SetLineno(735)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlanguage, "language"); πE != nil {
						continue
					}
					πTemp001[0] = µlanguage
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsmartchars); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µsmart = πTemp003
					// line 737: text = re.sub(r"""``""", smart.opquote, text)
					πF.SetLineno(737)
					πTemp001 = πF.MakeArgs(3)
					πTemp001[0] = πg.NewStr("``").ToObject()
					if πE = πg.CheckLocal(πF, µsmart, "smart"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsmart, ßopquote, nil); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[2] = µtext
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsub, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp002
					// line 738: text = re.sub(r"""''""", smart.cpquote, text)
					πF.SetLineno(738)
					πTemp001 = πF.MakeArgs(3)
					πTemp001[0] = πg.NewStr("''").ToObject()
					if πE = πg.CheckLocal(πF, µsmart, "smart"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsmart, ßcpquote, nil); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[2] = µtext
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsub, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp002
					// line 739: return text
					πF.SetLineno(739)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πR = µtext
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßeducateBackticks.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 728: """
			πF.SetLineno(728)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("\n    Parameter:  String (unicode or bytes).\n    Returns:    The `text`, with ``backticks'' -style double quotes\n                translated into HTML curly quote entities.\n    Example input:  ``Isn't this fun?''\n    Example output: \xe2\x80\x9cIsn't this fun?\xe2\x80\x9c;\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßeducateBackticks); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
				continue
			}
			// line 742: def educateSingleBackticks(text, language='en'):
			πF.SetLineno(742)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "text", Def: nil}
			πTemp006[1] = πg.Param{Name: "language", Def: ßen.ToObject()}
			πTemp008 = πg.NewFunction(πg.NewCode("educateSingleBackticks", "/usr/lib/python2.7/site-packages/docutils/utils/smartquotes.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtext *πg.Object = πArgs[0]
				_ = µtext
				var µlanguage *πg.Object = πArgs[1]
				_ = µlanguage
				var µsmart *πg.Object = πg.UnboundLocal
				_ = µsmart
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
					// line 743: """
					πF.SetLineno(743)
					// line 751: smart = smartchars(language)
					πF.SetLineno(751)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlanguage, "language"); πE != nil {
						continue
					}
					πTemp001[0] = µlanguage
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsmartchars); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µsmart = πTemp003
					// line 753: text = re.sub(r"""`""", smart.osquote, text)
					πF.SetLineno(753)
					πTemp001 = πF.MakeArgs(3)
					πTemp001[0] = πg.NewStr("`").ToObject()
					if πE = πg.CheckLocal(πF, µsmart, "smart"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsmart, ßosquote, nil); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[2] = µtext
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsub, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp002
					// line 754: text = re.sub(r"""'""", smart.csquote, text)
					πF.SetLineno(754)
					πTemp001 = πF.MakeArgs(3)
					πTemp001[0] = πg.NewStr("'").ToObject()
					if πE = πg.CheckLocal(πF, µsmart, "smart"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsmart, ßcsquote, nil); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[2] = µtext
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsub, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp002
					// line 755: return text
					πF.SetLineno(755)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πR = µtext
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßeducateSingleBackticks.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 743: """
			πF.SetLineno(743)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("\n    Parameter:  String (unicode or bytes).\n    Returns:    The `text`, with `backticks' -style single quotes\n                translated into HTML curly quote entities.\n\n    Example input:  `Isn't this fun?'\n    Example output: \xe2\x80\x98Isn\xe2\x80\x99t this fun?\xe2\x80\x99\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßeducateSingleBackticks); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
				continue
			}
			// line 758: def educateDashes(text):
			πF.SetLineno(758)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "text", Def: nil}
			πTemp009 = πg.NewFunction(πg.NewCode("educateDashes", "/usr/lib/python2.7/site-packages/docutils/utils/smartquotes.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtext *πg.Object = πArgs[0]
				_ = µtext
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
					// line 759: """
					πF.SetLineno(759)
					// line 765: text = re.sub(r"""---""", smartchars.endash, text) # en  (yes, backwards)
					πF.SetLineno(765)
					πTemp001 = πF.MakeArgs(3)
					πTemp001[0] = πg.NewStr("---").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsmartchars); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßendash, nil); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[2] = µtext
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsub, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp002
					// line 766: text = re.sub(r"""--""", smartchars.emdash, text) # em (yes, backwards)
					πF.SetLineno(766)
					πTemp001 = πF.MakeArgs(3)
					πTemp001[0] = πg.NewStr("--").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsmartchars); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßemdash, nil); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[2] = µtext
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsub, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp002
					// line 767: return text
					πF.SetLineno(767)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πR = µtext
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßeducateDashes.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 759: """
			πF.SetLineno(759)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("\n    Parameter:  String (unicode or bytes).\n    Returns:    The `text`, with each instance of \"--\" translated to\n                an em-dash character.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßeducateDashes); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
				continue
			}
			// line 770: def educateDashesOldSchool(text):
			πF.SetLineno(770)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "text", Def: nil}
			πTemp010 = πg.NewFunction(πg.NewCode("educateDashesOldSchool", "/usr/lib/python2.7/site-packages/docutils/utils/smartquotes.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtext *πg.Object = πArgs[0]
				_ = µtext
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
					// line 771: """
					πF.SetLineno(771)
					// line 778: text = re.sub(r"""---""", smartchars.emdash, text)
					πF.SetLineno(778)
					πTemp001 = πF.MakeArgs(3)
					πTemp001[0] = πg.NewStr("---").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsmartchars); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßemdash, nil); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[2] = µtext
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsub, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp002
					// line 779: text = re.sub(r"""--""", smartchars.endash, text)
					πF.SetLineno(779)
					πTemp001 = πF.MakeArgs(3)
					πTemp001[0] = πg.NewStr("--").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsmartchars); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßendash, nil); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[2] = µtext
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsub, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp002
					// line 780: return text
					πF.SetLineno(780)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πR = µtext
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßeducateDashesOldSchool.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 771: """
			πF.SetLineno(771)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("\n    Parameter:  String (unicode or bytes).\n    Returns:    The `text`, with each instance of \"--\" translated to\n                an en-dash character, and each \"---\" translated to\n                an em-dash character.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßeducateDashesOldSchool); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
				continue
			}
			// line 783: def educateDashesOldSchoolInverted(text):
			πF.SetLineno(783)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "text", Def: nil}
			πTemp011 = πg.NewFunction(πg.NewCode("educateDashesOldSchoolInverted", "/usr/lib/python2.7/site-packages/docutils/utils/smartquotes.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtext *πg.Object = πArgs[0]
				_ = µtext
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
					// line 784: """
					πF.SetLineno(784)
					// line 797: text = re.sub(r"""---""", smartchars.endash, text)    # em
					πF.SetLineno(797)
					πTemp001 = πF.MakeArgs(3)
					πTemp001[0] = πg.NewStr("---").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsmartchars); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßendash, nil); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[2] = µtext
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsub, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp002
					// line 798: text = re.sub(r"""--""", smartchars.emdash, text)    # en
					πF.SetLineno(798)
					πTemp001 = πF.MakeArgs(3)
					πTemp001[0] = πg.NewStr("--").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsmartchars); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßemdash, nil); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[2] = µtext
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsub, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp002
					// line 799: return text
					πF.SetLineno(799)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πR = µtext
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßeducateDashesOldSchoolInverted.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 784: """
			πF.SetLineno(784)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("\n    Parameter:  String (unicode or bytes).\n    Returns:    The `text`, with each instance of \"--\" translated to\n                an em-dash character, and each \"---\" translated to\n                an en-dash character. Two reasons why: First, unlike the\n                en- and em-dash syntax supported by\n                EducateDashesOldSchool(), it's compatible with existing\n                entries written before SmartyPants 1.1, back when \"--\" was\n                only used for em-dashes.  Second, em-dashes are more\n                common than en-dashes, and so it sort of makes sense that\n                the shortcut should be shorter to type. (Thanks to Aaron\n                Swartz for the idea.)\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßeducateDashesOldSchoolInverted); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
				continue
			}
			// line 803: def educateEllipses(text):
			πF.SetLineno(803)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "text", Def: nil}
			πTemp012 = πg.NewFunction(πg.NewCode("educateEllipses", "/usr/lib/python2.7/site-packages/docutils/utils/smartquotes.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtext *πg.Object = πArgs[0]
				_ = µtext
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
					// line 804: """
					πF.SetLineno(804)
					// line 813: text = re.sub(r"""\.\.\.""", smartchars.ellipsis, text)
					πF.SetLineno(813)
					πTemp001 = πF.MakeArgs(3)
					πTemp001[0] = πg.NewStr("\\.\\.\\.").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsmartchars); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßellipsis, nil); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[2] = µtext
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsub, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp002
					// line 814: text = re.sub(r"""\. \. \.""", smartchars.ellipsis, text)
					πF.SetLineno(814)
					πTemp001 = πF.MakeArgs(3)
					πTemp001[0] = πg.NewStr("\\. \\. \\.").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsmartchars); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßellipsis, nil); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[2] = µtext
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsub, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp002
					// line 815: return text
					πF.SetLineno(815)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πR = µtext
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßeducateEllipses.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 804: """
			πF.SetLineno(804)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πg.NewStr("\n    Parameter:  String (unicode or bytes).\n    Returns:    The `text`, with each instance of \"...\" translated to\n                an ellipsis character.\n\n    Example input:  Huh...?\n    Example output: Huh&#8230;?\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp014, πE = πg.ResolveGlobal(πF, ßeducateEllipses); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp014, ß__doc__, πTemp013); πE != nil {
				continue
			}
			// line 818: def stupefyEntities(text, language='en'):
			πF.SetLineno(818)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "text", Def: nil}
			πTemp006[1] = πg.Param{Name: "language", Def: ßen.ToObject()}
			πTemp013 = πg.NewFunction(πg.NewCode("stupefyEntities", "/usr/lib/python2.7/site-packages/docutils/utils/smartquotes.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtext *πg.Object = πArgs[0]
				_ = µtext
				var µlanguage *πg.Object = πArgs[1]
				_ = µlanguage
				var µsmart *πg.Object = πg.UnboundLocal
				_ = µsmart
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
					// line 819: """
					πF.SetLineno(819)
					// line 827: smart = smartchars(language)
					πF.SetLineno(827)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlanguage, "language"); πE != nil {
						continue
					}
					πTemp001[0] = µlanguage
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsmartchars); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µsmart = πTemp003
					// line 829: text = re.sub(smart.endash, "-", text)  # en-dash
					πF.SetLineno(829)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µsmart, "smart"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsmart, ßendash, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					πTemp001[1] = πg.NewStr("-").ToObject()
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[2] = µtext
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsub, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp002
					// line 830: text = re.sub(smart.emdash, "--", text) # em-dash
					πF.SetLineno(830)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µsmart, "smart"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsmart, ßemdash, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					πTemp001[1] = πg.NewStr("--").ToObject()
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[2] = µtext
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsub, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp002
					// line 832: text = re.sub(smart.osquote, "'", text)  # open single quote
					πF.SetLineno(832)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µsmart, "smart"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsmart, ßosquote, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					πTemp001[1] = πg.NewStr("'").ToObject()
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[2] = µtext
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsub, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp002
					// line 833: text = re.sub(smart.csquote, "'", text)  # close single quote
					πF.SetLineno(833)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µsmart, "smart"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsmart, ßcsquote, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					πTemp001[1] = πg.NewStr("'").ToObject()
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[2] = µtext
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsub, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp002
					// line 835: text = re.sub(smart.opquote, '"', text)  # open double quote
					πF.SetLineno(835)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µsmart, "smart"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsmart, ßopquote, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					πTemp001[1] = πg.NewStr("\"").ToObject()
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[2] = µtext
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsub, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp002
					// line 836: text = re.sub(smart.cpquote, '"', text)  # close double quote
					πF.SetLineno(836)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µsmart, "smart"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsmart, ßcpquote, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					πTemp001[1] = πg.NewStr("\"").ToObject()
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[2] = µtext
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsub, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp002
					// line 838: text = re.sub(smart.ellipsis, '...', text)# ellipsis
					πF.SetLineno(838)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µsmart, "smart"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsmart, ßellipsis, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					πTemp001[1] = πg.NewStr("...").ToObject()
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πTemp001[2] = µtext
					if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsub, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtext = πTemp002
					// line 840: return text
					πF.SetLineno(840)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πR = µtext
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßstupefyEntities.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 819: """
			πF.SetLineno(819)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πg.NewStr("\n    Parameter:  String (unicode or bytes).\n    Returns:    The `text`, with each SmartyPants character translated to\n                its ASCII counterpart.\n\n    Example input:  \xe2\x80\x9cHello \xe2\x80\x94 world.\xe2\x80\x9d\n    Example output: \"Hello -- world.\"\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßstupefyEntities); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp015, ß__doc__, πTemp014); πE != nil {
				continue
			}
			// line 843: def processEscapes(text, restore=False):
			πF.SetLineno(843)
			πTemp006 = make([]πg.Param, 2)
			πTemp006[0] = πg.Param{Name: "text", Def: nil}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp006[1] = πg.Param{Name: "restore", Def: πTemp015}
			πTemp014 = πg.NewFunction(πg.NewCode("processEscapes", "/usr/lib/python2.7/site-packages/docutils/utils/smartquotes.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtext *πg.Object = πArgs[0]
				_ = µtext
				var µrestore *πg.Object = πArgs[1]
				_ = µrestore
				var µreplacements *πg.Object = πg.UnboundLocal
				_ = µreplacements
				var µch *πg.Object = πg.UnboundLocal
				_ = µch
				var µrep *πg.Object = πg.UnboundLocal
				_ = µrep
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 bool
				_ = πTemp009
				var πTemp010 []*πg.Object
				_ = πTemp010
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 8:
						goto Label8
					case 4:
						goto Label4
					case 5:
						goto Label5
					case 7:
						goto Label7
					default:
						panic("unexpected function state")
					}
					// line 844: r"""
					πF.SetLineno(844)
					// line 859: replacements = ((r'\\', r'&#92;'),
					πF.SetLineno(859)
					πTemp002 = πg.NewTuple2(πg.NewStr("\\\\").ToObject(), πg.NewStr("&#92;").ToObject()).ToObject()
					πTemp003 = πg.NewTuple2(πg.NewStr("\\\"").ToObject(), πg.NewStr("&#34;").ToObject()).ToObject()
					πTemp004 = πg.NewTuple2(πg.NewStr("\\'").ToObject(), πg.NewStr("&#39;").ToObject()).ToObject()
					πTemp005 = πg.NewTuple2(πg.NewStr("\\.").ToObject(), πg.NewStr("&#46;").ToObject()).ToObject()
					πTemp006 = πg.NewTuple2(πg.NewStr("\\-").ToObject(), πg.NewStr("&#45;").ToObject()).ToObject()
					πTemp007 = πg.NewTuple2(πg.NewStr("\\`").ToObject(), πg.NewStr("&#96;").ToObject()).ToObject()
					πTemp001 = πg.NewTuple6(πTemp002, πTemp003, πTemp004, πTemp005, πTemp006, πTemp007).ToObject()
					µreplacements = πTemp001
					if πE = πg.CheckLocal(πF, µrestore, "restore"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, µrestore); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label1
					}
					goto Label2
					// line 865: if restore:
					πF.SetLineno(865)
				Label1:
					if πE = πg.CheckLocal(πF, µreplacements, "replacements"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µreplacements); πE != nil {
						continue
					}
					πF.PushCheckpoint(5)
					πTemp008 = false
				Label4:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp008 {
						πF.PopCheckpoint()
						goto Label6
					}
					if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp009 = !isStop
					} else {
						πTemp009 = true
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp002); πE != nil {
							continue
						}
						µch = πTemp003
						µrep = πTemp004
					}
					if πE != nil || !πTemp009 {
						continue
					}
					πF.PushCheckpoint(4)
					// line 867: text = text.replace(rep, ch[1])
					πF.SetLineno(867)
					πTemp010 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µrep, "rep"); πE != nil {
						continue
					}
					πTemp010[0] = µrep
					πTemp002 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µch, "ch"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µch, πTemp002); πE != nil {
						continue
					}
					πTemp010[1] = πTemp003
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µtext, ßreplace, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					µtext = πTemp003
					continue
				Label5:
					if πE != nil || πR != nil {
						continue
					}
				Label6:
					goto Label3
				Label2:
					if πE = πg.CheckLocal(πF, µreplacements, "replacements"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µreplacements); πE != nil {
						continue
					}
					πF.PushCheckpoint(8)
					πTemp008 = false
				Label7:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp008 {
						πF.PopCheckpoint()
						goto Label9
					}
					if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp009 = !isStop
					} else {
						πTemp009 = true
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp002); πE != nil {
							continue
						}
						µch = πTemp003
						µrep = πTemp004
					}
					if πE != nil || !πTemp009 {
						continue
					}
					πF.PushCheckpoint(7)
					// line 870: text = text.replace(ch, rep)
					πF.SetLineno(870)
					πTemp010 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µch, "ch"); πE != nil {
						continue
					}
					πTemp010[0] = µch
					if πE = πg.CheckLocal(πF, µrep, "rep"); πE != nil {
						continue
					}
					πTemp010[1] = µrep
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µtext, ßreplace, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					µtext = πTemp003
					continue
				Label8:
					if πE != nil || πR != nil {
						continue
					}
				Label9:
					goto Label3
				Label3:
					// line 872: return text
					πF.SetLineno(872)
					if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
						continue
					}
					πR = µtext
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßprocessEscapes.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 844: r"""
			πF.SetLineno(844)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp015}, πg.NewStr("\n    Parameter:  String (unicode or bytes).\n    Returns:    The `text`, with after processing the following backslash\n                escape sequences. This is useful if you want to force a \"dumb\"\n                quote or other character to appear.\n\n                Escape  Value\n                ------  -----\n                \\\\      &#92;\n                \\\"      &#34;\n                \\'      &#39;\n                \\.      &#46;\n                \\-      &#45;\n                \\`      &#96;\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßprocessEscapes); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp016, ß__doc__, πTemp015); πE != nil {
				continue
			}
			// line 875: def tokenize(text):
			πF.SetLineno(875)
			πTemp006 = make([]πg.Param, 1)
			πTemp006[0] = πg.Param{Name: "text", Def: nil}
			πTemp015 = πg.NewFunction(πg.NewCode("tokenize", "/usr/lib/python2.7/site-packages/docutils/utils/smartquotes.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtext *πg.Object = πArgs[0]
				_ = µtext
				var µpos *πg.Object = πg.UnboundLocal
				_ = µpos
				var µlength *πg.Object = πg.UnboundLocal
				_ = µlength
				var µdepth *πg.Object = πg.UnboundLocal
				_ = µdepth
				var µnested_tags *πg.Object = πg.UnboundLocal
				_ = µnested_tags
				var µtag_soup *πg.Object = πg.UnboundLocal
				_ = µtag_soup
				var µtoken_match *πg.Object = πg.UnboundLocal
				_ = µtoken_match
				var µprevious_end *πg.Object = πg.UnboundLocal
				_ = µprevious_end
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
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
						case 10:
							goto Label10
						case 6:
							goto Label6
						case 7:
							goto Label7
						default:
							panic("unexpected function state")
						}
						// line 876: """
						πF.SetLineno(876)
						// line 889: pos = 0
						πF.SetLineno(889)
						µpos = πg.NewInt(0).ToObject()
						// line 890: length = len(text)
						πF.SetLineno(890)
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
							continue
						}
						πTemp001[0] = µtext
						if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						µlength = πTemp003
						// line 893: depth = 6
						πF.SetLineno(893)
						µdepth = πg.NewInt(6).ToObject()
						// line 894: nested_tags = "|".join(['(?:<(?:[^<>]',] * depth) + (')*>)' * depth)
						πF.SetLineno(894)
						πTemp001 = πF.MakeArgs(1)
						πTemp004 = make([]*πg.Object, 1)
						πTemp004[0] = πg.NewStr("(?:<(?:[^<>]").ToObject()
						πTemp005 = πg.NewList(πTemp004...).ToObject()
						if πE = πg.CheckLocal(πF, µdepth, "depth"); πE != nil {
							continue
						}
						if πTemp003, πE = πg.Mul(πF, πTemp005, µdepth); πE != nil {
							continue
						}
						πTemp001[0] = πTemp003
						if πTemp003, πE = πg.GetAttr(πF, πg.NewStr("|").ToObject(), ßjoin, nil); πE != nil {
							continue
						}
						if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						if πE = πg.CheckLocal(πF, µdepth, "depth"); πE != nil {
							continue
						}
						if πTemp003, πE = πg.Mul(πF, πg.NewStr(")*>)").ToObject(), µdepth); πE != nil {
							continue
						}
						if πTemp002, πE = πg.Add(πF, πTemp005, πTemp003); πE != nil {
							continue
						}
						µnested_tags = πTemp002
						// line 898: tag_soup = re.compile(r"""([^<]*)(<[^>]*>)""")
						πF.SetLineno(898)
						πTemp001 = πF.MakeArgs(1)
						πTemp001[0] = πg.NewStr("([^<]*)(<[^>]*>)").ToObject()
						if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcompile, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						µtag_soup = πTemp002
						// line 900: token_match = tag_soup.search(text)
						πF.SetLineno(900)
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
							continue
						}
						πTemp001[0] = µtext
						if πE = πg.CheckLocal(πF, µtag_soup, "tag_soup"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µtag_soup, ßsearch, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						µtoken_match = πTemp003
						// line 902: previous_end = 0
						πF.SetLineno(902)
						µprevious_end = πg.NewInt(0).ToObject()
						// line 903: while token_match is not None:
						πF.SetLineno(903)
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
						if πE = πg.CheckLocal(πF, µtoken_match, "token_match"); πE != nil {
							continue
						}
						if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
							continue
						}
						πTemp002 = πg.GetBool(µtoken_match != πTemp003).ToObject()
						if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
							continue
						}
						if πE != nil || !πTemp007 {
							continue
						}
						πF.PushCheckpoint(1)
						πTemp001 = πF.MakeArgs(1)
						πTemp001[0] = πg.NewInt(1).ToObject()
						if πE = πg.CheckLocal(πF, µtoken_match, "token_match"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µtoken_match, ßgroup, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
							continue
						}
						if πTemp007 {
							goto Label4
						}
						goto Label5
						// line 904: if token_match.group(1):
						πF.SetLineno(904)
					Label4:
						// line 905: yield ('text', token_match.group(1))
						πF.SetLineno(905)
						πTemp001 = πF.MakeArgs(1)
						πTemp001[0] = πg.NewInt(1).ToObject()
						if πE = πg.CheckLocal(πF, µtoken_match, "token_match"); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, µtoken_match, ßgroup, nil); πE != nil {
							continue
						}
						if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						πTemp002 = πg.NewTuple2(ßtext.ToObject(), πTemp005).ToObject()
						πF.PushCheckpoint(6)
						return πTemp002, nil
					Label6:
						πTemp003 = πSent
						goto Label5
					Label5:
						// line 907: yield ('tag', token_match.group(2))
						πF.SetLineno(907)
						πTemp001 = πF.MakeArgs(1)
						πTemp001[0] = πg.NewInt(2).ToObject()
						if πE = πg.CheckLocal(πF, µtoken_match, "token_match"); πE != nil {
							continue
						}
						if πTemp005, πE = πg.GetAttr(πF, µtoken_match, ßgroup, nil); πE != nil {
							continue
						}
						if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						πTemp003 = πg.NewTuple2(ßtag.ToObject(), πTemp008).ToObject()
						πF.PushCheckpoint(7)
						return πTemp003, nil
					Label7:
						πTemp005 = πSent
						// line 909: previous_end = token_match.end()
						πF.SetLineno(909)
						if πE = πg.CheckLocal(πF, µtoken_match, "token_match"); πE != nil {
							continue
						}
						if πTemp005, πE = πg.GetAttr(πF, µtoken_match, ßend, nil); πE != nil {
							continue
						}
						if πTemp008, πE = πTemp005.Call(πF, nil, nil); πE != nil {
							continue
						}
						µprevious_end = πTemp008
						// line 910: token_match = tag_soup.search(text, token_match.end())
						πF.SetLineno(910)
						πTemp001 = πF.MakeArgs(2)
						if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
							continue
						}
						πTemp001[0] = µtext
						if πE = πg.CheckLocal(πF, µtoken_match, "token_match"); πE != nil {
							continue
						}
						if πTemp005, πE = πg.GetAttr(πF, µtoken_match, ßend, nil); πE != nil {
							continue
						}
						if πTemp008, πE = πTemp005.Call(πF, nil, nil); πE != nil {
							continue
						}
						πTemp001[1] = πTemp008
						if πE = πg.CheckLocal(πF, µtag_soup, "tag_soup"); πE != nil {
							continue
						}
						if πTemp005, πE = πg.GetAttr(πF, µtag_soup, ßsearch, nil); πE != nil {
							continue
						}
						if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						µtoken_match = πTemp008
						continue
					Label2:
						if πE != nil || πR != nil {
							continue
						}
					Label3:
						if πE = πg.CheckLocal(πF, µprevious_end, "previous_end"); πE != nil {
							continue
						}
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
							continue
						}
						πTemp001[0] = µtext
						if πTemp008, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
							continue
						}
						if πTemp009, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						if πTemp005, πE = πg.LT(πF, µprevious_end, πTemp009); πE != nil {
							continue
						}
						if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
							continue
						}
						if πTemp006 {
							goto Label8
						}
						goto Label9
						// line 912: if previous_end < len(text):
						πF.SetLineno(912)
					Label8:
						// line 913: yield ('text', text[previous_end:])
						πF.SetLineno(913)
						if πE = πg.CheckLocal(πF, µprevious_end, "previous_end"); πE != nil {
							continue
						}
						if πTemp008, πE = πg.SliceType.Call(πF, πg.Args{µprevious_end, πg.None, πg.None}, nil); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
							continue
						}
						if πTemp009, πE = πg.GetItem(πF, µtext, πTemp008); πE != nil {
							continue
						}
						πTemp005 = πg.NewTuple2(ßtext.ToObject(), πTemp009).ToObject()
						πF.PushCheckpoint(10)
						return πTemp005, nil
					Label10:
						πTemp008 = πSent
						goto Label9
					Label9:
					}
					return nil, πE
				}).ToObject(), nil
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßtokenize.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 876: """
			πF.SetLineno(876)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp016}, πg.NewStr("\n    Parameter:  String containing HTML markup.\n    Returns:    An iterator that yields the tokens comprising the input\n                string. Each token is either a tag (possibly with nested,\n                tags contained therein, such as <a href=\"<MTFoo>\">, or a\n                run of text between tags. Each yielded element is a\n                two-element tuple; the first is either 'tag' or 'text';\n                the second is the actual value.\n\n    Based on the _tokenize() subroutine from Brad Choate's MTRegex plugin.\n        <http://www.bradchoate.com/past/mtregex.php>\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp017, πE = πg.ResolveGlobal(πF, ßtokenize); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp017, ß__doc__, πTemp016); πE != nil {
				continue
			}
			if πTemp017, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp016, πE = πg.Eq(πF, πTemp017, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp018, πE = πg.IsTrue(πF, πTemp016); πE != nil {
				continue
			}
			if πTemp018 {
				goto Label1
			}
			goto Label2
			// line 917: if __name__ == "__main__":
			πF.SetLineno(917)
		Label1:
			// line 919: import itertools
			πF.SetLineno(919)
			if πTemp002, πE = πg.ImportModule(πF, "itertools"); πE != nil {
				continue
			}
			πTemp016 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßitertools.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 920: try:
			πF.SetLineno(920)
			πF.PushCheckpoint(4)
			// line 921: import locale # module missing in Jython
			πF.SetLineno(921)
			if πTemp002, πE = πg.ImportModule(πF, "locale"); πE != nil {
				continue
			}
			πTemp016 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßlocale.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 922: locale.setlocale(locale.LC_ALL, '') # set to user defaults
			πF.SetLineno(922)
			πTemp002 = πF.MakeArgs(2)
			if πTemp016, πE = πg.ResolveGlobal(πF, ßlocale); πE != nil {
				continue
			}
			if πTemp017, πE = πg.GetAttr(πF, πTemp016, ßLC_ALL, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp017
			πTemp002[1] = ß.ToObject()
			if πTemp016, πE = πg.ResolveGlobal(πF, ßlocale); πE != nil {
				continue
			}
			if πTemp017, πE = πg.GetAttr(πF, πTemp016, ßsetlocale, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp017.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 923: defaultlanguage = locale.getdefaultlocale()[0]
			πF.SetLineno(923)
			πTemp016 = πg.NewInt(0).ToObject()
			if πTemp019, πE = πg.ResolveGlobal(πF, ßlocale); πE != nil {
				continue
			}
			if πTemp020, πE = πg.GetAttr(πF, πTemp019, ßgetdefaultlocale, nil); πE != nil {
				continue
			}
			if πTemp019, πE = πTemp020.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp017, πE = πg.GetItem(πF, πTemp019, πTemp016); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdefaultlanguage.ToObject(), πTemp017); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label3
		Label4:
			if πE == nil {
				continue
			}
			πE = nil
			πTemp021, πTemp022 = πF.ExcInfo()
			goto Label5
			// line 924: except:
			πF.SetLineno(924)
		Label5:
			// line 925: defaultlanguage = 'en'
			πF.SetLineno(925)
			if πE = πF.Globals().SetItem(πF, ßdefaultlanguage.ToObject(), ßen.ToObject()); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label3
		Label3:
			// line 928: defaultlanguage = defaultlanguage.lower().replace('-', '_')
			πF.SetLineno(928)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = πg.NewStr("-").ToObject()
			πTemp002[1] = ß_.ToObject()
			if πTemp016, πE = πg.ResolveGlobal(πF, ßdefaultlanguage); πE != nil {
				continue
			}
			if πTemp017, πE = πg.GetAttr(πF, πTemp016, ßlower, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp017.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp017, πE = πg.GetAttr(πF, πTemp016, ßreplace, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp017.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßdefaultlanguage.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 930: defaultlanguage = re.sub(r'_([a-zA-Z0-9])_', r'_\1-', defaultlanguage)
			πF.SetLineno(930)
			πTemp002 = πF.MakeArgs(3)
			πTemp002[0] = πg.NewStr("_([a-zA-Z0-9])_").ToObject()
			πTemp002[1] = πg.NewStr("_\\1-").ToObject()
			if πTemp016, πE = πg.ResolveGlobal(πF, ßdefaultlanguage); πE != nil {
				continue
			}
			πTemp002[2] = πTemp016
			if πTemp016, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp017, πE = πg.GetAttr(πF, πTemp016, ßsub, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp017.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßdefaultlanguage.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 931: _subtags = [subtag for subtag in defaultlanguage.split('_')]
			πF.SetLineno(931)
			πTemp006 = make([]πg.Param, 0)
			πTemp017 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/site-packages/docutils/utils/smartquotes.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsubtag *πg.Object = πg.UnboundLocal
				_ = µsubtag
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
				var πTemp006 bool
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
						πTemp002 = πF.MakeArgs(1)
						πTemp002[0] = ß_.ToObject()
						if πTemp003, πE = πg.ResolveGlobal(πF, ßdefaultlanguage); πE != nil {
							continue
						}
						if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsplit, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp002)
						if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
							continue
						}
						πF.PushCheckpoint(2)
						πTemp005 = false
					Label1:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp005 {
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
							πTemp006 = !isStop
						} else {
							πTemp006 = true
							µsubtag = πTemp003
						}
						if πE != nil || !πTemp006 {
							continue
						}
						πF.PushCheckpoint(1)
						// line 931: _subtags = [subtag for subtag in defaultlanguage.split('_')]
						πF.SetLineno(931)
						if πE = πg.CheckLocal(πF, µsubtag, "subtag"); πE != nil {
							continue
						}
						πF.PushCheckpoint(4)
						return µsubtag, nil
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
			if πTemp019, πE = πTemp017.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πg.ListType.Call(πF, πg.Args{πTemp019}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_subtags.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 932: _basetag = _subtags.pop(0)
			πF.SetLineno(932)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewInt(0).ToObject()
			if πTemp016, πE = πg.ResolveGlobal(πF, ß_subtags); πE != nil {
				continue
			}
			if πTemp019, πE = πg.GetAttr(πF, πTemp016, ßpop, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp019.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ß_basetag.ToObject(), πTemp016); πE != nil {
				continue
			}
			πTemp002 = πF.MakeArgs(3)
			πTemp023 = πF.MakeArgs(1)
			if πTemp019, πE = πg.ResolveGlobal(πF, ß_subtags); πE != nil {
				continue
			}
			πTemp023[0] = πTemp019
			if πTemp019, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
				continue
			}
			if πTemp020, πE = πTemp019.Call(πF, πTemp023, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp023)
			πTemp002[0] = πTemp020
			πTemp002[1] = πg.NewInt(0).ToObject()
			if πTemp019, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			πTemp002[2] = πTemp019
			if πTemp019, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
				continue
			}
			if πTemp020, πE = πTemp019.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πTemp016, πE = πg.Iter(πF, πTemp020); πE != nil {
				continue
			}
			πF.PushCheckpoint(7)
			πTemp018 = false
		Label6:
			if πE != nil || πR != nil {
				continue
			}
			if πTemp018 {
				πF.PopCheckpoint()
				goto Label8
			}
			if πTemp019, πE = πg.Next(πF, πTemp016); πE != nil {
				isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
				if exc != nil {
					πE = exc
				} else if isStop {
					πE = nil
					πF.RestoreExc(nil, nil)
				}
				πTemp024 = !isStop
			} else {
				πTemp024 = true
				if πE = πF.Globals().SetItem(πF, ßn.ToObject(), πTemp019); πE != nil {
					continue
				}
			}
			if πE != nil || !πTemp024 {
				continue
			}
			πF.PushCheckpoint(6)
			πTemp002 = πF.MakeArgs(2)
			if πTemp020, πE = πg.ResolveGlobal(πF, ß_subtags); πE != nil {
				continue
			}
			πTemp002[0] = πTemp020
			if πTemp020, πE = πg.ResolveGlobal(πF, ßn); πE != nil {
				continue
			}
			πTemp002[1] = πTemp020
			if πTemp020, πE = πg.ResolveGlobal(πF, ßitertools); πE != nil {
				continue
			}
			if πTemp025, πE = πg.GetAttr(πF, πTemp020, ßcombinations, nil); πE != nil {
				continue
			}
			if πTemp020, πE = πTemp025.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πTemp019, πE = πg.Iter(πF, πTemp020); πE != nil {
				continue
			}
			πF.PushCheckpoint(10)
			πTemp024 = false
		Label9:
			if πE != nil || πR != nil {
				continue
			}
			if πTemp024 {
				πF.PopCheckpoint()
				goto Label11
			}
			if πTemp020, πE = πg.Next(πF, πTemp019); πE != nil {
				isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
				if exc != nil {
					πE = exc
				} else if isStop {
					πE = nil
					πF.RestoreExc(nil, nil)
				}
				πTemp026 = !isStop
			} else {
				πTemp026 = true
				if πE = πF.Globals().SetItem(πF, ßtags.ToObject(), πTemp020); πE != nil {
					continue
				}
			}
			if πE != nil || !πTemp026 {
				continue
			}
			πF.PushCheckpoint(9)
			// line 936: _tag = '-'.join((_basetag,)+tags)
			πF.SetLineno(936)
			πTemp002 = πF.MakeArgs(1)
			if πTemp027, πE = πg.ResolveGlobal(πF, ß_basetag); πE != nil {
				continue
			}
			πTemp025 = πg.NewTuple1(πTemp027).ToObject()
			if πTemp027, πE = πg.ResolveGlobal(πF, ßtags); πE != nil {
				continue
			}
			if πTemp020, πE = πg.Add(πF, πTemp025, πTemp027); πE != nil {
				continue
			}
			πTemp002[0] = πTemp020
			if πTemp020, πE = πg.GetAttr(πF, πg.NewStr("-").ToObject(), ßjoin, nil); πE != nil {
				continue
			}
			if πTemp025, πE = πTemp020.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ß_tag.ToObject(), πTemp025); πE != nil {
				continue
			}
			if πTemp025, πE = πg.ResolveGlobal(πF, ß_tag); πE != nil {
				continue
			}
			if πTemp027, πE = πg.ResolveGlobal(πF, ßsmartchars); πE != nil {
				continue
			}
			if πTemp028, πE = πg.GetAttr(πF, πTemp027, ßquotes, nil); πE != nil {
				continue
			}
			if πTemp026, πE = πg.Contains(πF, πTemp028, πTemp025); πE != nil {
				continue
			}
			πTemp020 = πg.GetBool(πTemp026).ToObject()
			if πTemp026, πE = πg.IsTrue(πF, πTemp020); πE != nil {
				continue
			}
			if πTemp026 {
				goto Label12
			}
			goto Label13
			// line 937: if _tag in smartchars.quotes:
			πF.SetLineno(937)
		Label12:
			// line 938: defaultlanguage = _tag
			πF.SetLineno(938)
			if πTemp020, πE = πg.ResolveGlobal(πF, ß_tag); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdefaultlanguage.ToObject(), πTemp020); πE != nil {
				continue
			}
			// line 939: break
			πF.SetLineno(939)
			πTemp024 = true
			continue
			goto Label13
		Label13:
			continue
		Label10:
			if πE != nil || πR != nil {
				continue
			}
			if πTemp025, πE = πg.ResolveGlobal(πF, ß_basetag); πE != nil {
				continue
			}
			if πTemp027, πE = πg.ResolveGlobal(πF, ßsmartchars); πE != nil {
				continue
			}
			if πTemp028, πE = πg.GetAttr(πF, πTemp027, ßquotes, nil); πE != nil {
				continue
			}
			if πTemp026, πE = πg.Contains(πF, πTemp028, πTemp025); πE != nil {
				continue
			}
			πTemp020 = πg.GetBool(πTemp026).ToObject()
			if πTemp026, πE = πg.IsTrue(πF, πTemp020); πE != nil {
				continue
			}
			if πTemp026 {
				goto Label14
			}
			goto Label15
			// line 941: if _basetag in smartchars.quotes:
			πF.SetLineno(941)
		Label14:
			// line 942: defaultlanguage = _basetag
			πF.SetLineno(942)
			if πTemp020, πE = πg.ResolveGlobal(πF, ß_basetag); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdefaultlanguage.ToObject(), πTemp020); πE != nil {
				continue
			}
			goto Label16
		Label15:
			// line 944: defaultlanguage = 'en'
			πF.SetLineno(944)
			if πE = πF.Globals().SetItem(πF, ßdefaultlanguage.ToObject(), ßen.ToObject()); πE != nil {
				continue
			}
			goto Label16
		Label16:
		Label11:
			continue
		Label7:
			if πE != nil || πR != nil {
				continue
			}
		Label8:
			// line 947: import argparse
			πF.SetLineno(947)
			if πTemp002, πE = πg.ImportModule(πF, "argparse"); πE != nil {
				continue
			}
			πTemp016 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßargparse.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 948: parser = argparse.ArgumentParser(
			πF.SetLineno(948)
			πTemp029 = πg.KWArgs{
				{"description", πg.NewStr("Filter stdin making ASCII punctuation \"smart\".").ToObject()},
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßargparse); πE != nil {
				continue
			}
			if πTemp019, πE = πg.GetAttr(πF, πTemp016, ßArgumentParser, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp019.Call(πF, nil, πTemp029); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßparser.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 951: parser.add_argument("-a", "--action", default="1",
			πF.SetLineno(951)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = πg.NewStr("-a").ToObject()
			πTemp002[1] = πg.NewStr("--action").ToObject()
			πTemp029 = πg.KWArgs{
				{"default", ß1.ToObject()},
				{"help", πg.NewStr("what to do with the input (see --actionhelp)").ToObject()},
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßparser); πE != nil {
				continue
			}
			if πTemp019, πE = πg.GetAttr(πF, πTemp016, ßadd_argument, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp019.Call(πF, πTemp002, πTemp029); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 953: parser.add_argument("-e", "--encoding", default="utf8",
			πF.SetLineno(953)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = πg.NewStr("-e").ToObject()
			πTemp002[1] = πg.NewStr("--encoding").ToObject()
			πTemp029 = πg.KWArgs{
				{"default", ßutf8.ToObject()},
				{"help", πg.NewStr("text encoding").ToObject()},
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßparser); πE != nil {
				continue
			}
			if πTemp019, πE = πg.GetAttr(πF, πTemp016, ßadd_argument, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp019.Call(πF, πTemp002, πTemp029); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 955: parser.add_argument("-l", "--language", default=defaultlanguage,
			πF.SetLineno(955)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = πg.NewStr("-l").ToObject()
			πTemp002[1] = πg.NewStr("--language").ToObject()
			if πTemp016, πE = πg.ResolveGlobal(πF, ßdefaultlanguage); πE != nil {
				continue
			}
			if πTemp020, πE = πg.ResolveGlobal(πF, ßdefaultlanguage); πE != nil {
				continue
			}
			if πTemp019, πE = πg.Mod(πF, πg.NewStr("text language (BCP47 tag), Default: %s").ToObject(), πTemp020); πE != nil {
				continue
			}
			πTemp029 = πg.KWArgs{
				{"default", πTemp016},
				{"help", πTemp019},
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßparser); πE != nil {
				continue
			}
			if πTemp019, πE = πg.GetAttr(πF, πTemp016, ßadd_argument, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp019.Call(πF, πTemp002, πTemp029); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 958: parser.add_argument("-q", "--alternative-quotes", action="store_true",
			πF.SetLineno(958)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = πg.NewStr("-q").ToObject()
			πTemp002[1] = πg.NewStr("--alternative-quotes").ToObject()
			πTemp029 = πg.KWArgs{
				{"action", ßstore_true.ToObject()},
				{"help", πg.NewStr("use alternative quote style").ToObject()},
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßparser); πE != nil {
				continue
			}
			if πTemp019, πE = πg.GetAttr(πF, πTemp016, ßadd_argument, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp019.Call(πF, πTemp002, πTemp029); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 960: parser.add_argument("--doc", action="store_true",
			πF.SetLineno(960)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("--doc").ToObject()
			πTemp029 = πg.KWArgs{
				{"action", ßstore_true.ToObject()},
				{"help", πg.NewStr("print documentation").ToObject()},
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßparser); πE != nil {
				continue
			}
			if πTemp019, πE = πg.GetAttr(πF, πTemp016, ßadd_argument, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp019.Call(πF, πTemp002, πTemp029); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 962: parser.add_argument("--actionhelp", action="store_true",
			πF.SetLineno(962)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("--actionhelp").ToObject()
			πTemp029 = πg.KWArgs{
				{"action", ßstore_true.ToObject()},
				{"help", πg.NewStr("list available actions").ToObject()},
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßparser); πE != nil {
				continue
			}
			if πTemp019, πE = πg.GetAttr(πF, πTemp016, ßadd_argument, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp019.Call(πF, πTemp002, πTemp029); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 964: parser.add_argument("--stylehelp", action="store_true",
			πF.SetLineno(964)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("--stylehelp").ToObject()
			πTemp029 = πg.KWArgs{
				{"action", ßstore_true.ToObject()},
				{"help", πg.NewStr("list available quote styles").ToObject()},
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßparser); πE != nil {
				continue
			}
			if πTemp019, πE = πg.GetAttr(πF, πTemp016, ßadd_argument, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp019.Call(πF, πTemp002, πTemp029); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 966: parser.add_argument("--test", action="store_true",
			πF.SetLineno(966)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("--test").ToObject()
			πTemp029 = πg.KWArgs{
				{"action", ßstore_true.ToObject()},
				{"help", πg.NewStr("perform short self-test").ToObject()},
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßparser); πE != nil {
				continue
			}
			if πTemp019, πE = πg.GetAttr(πF, πTemp016, ßadd_argument, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp019.Call(πF, πTemp002, πTemp029); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 968: args = parser.parse_args()
			πF.SetLineno(968)
			if πTemp016, πE = πg.ResolveGlobal(πF, ßparser); πE != nil {
				continue
			}
			if πTemp019, πE = πg.GetAttr(πF, πTemp016, ßparse_args, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp019.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßargs.ToObject(), πTemp016); πE != nil {
				continue
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßargs); πE != nil {
				continue
			}
			if πTemp019, πE = πg.GetAttr(πF, πTemp016, ßdoc, nil); πE != nil {
				continue
			}
			if πTemp018, πE = πg.IsTrue(πF, πTemp019); πE != nil {
				continue
			}
			if πTemp018 {
				goto Label17
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßargs); πE != nil {
				continue
			}
			if πTemp019, πE = πg.GetAttr(πF, πTemp016, ßactionhelp, nil); πE != nil {
				continue
			}
			if πTemp018, πE = πg.IsTrue(πF, πTemp019); πE != nil {
				continue
			}
			if πTemp018 {
				goto Label18
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßargs); πE != nil {
				continue
			}
			if πTemp019, πE = πg.GetAttr(πF, πTemp016, ßstylehelp, nil); πE != nil {
				continue
			}
			if πTemp018, πE = πg.IsTrue(πF, πTemp019); πE != nil {
				continue
			}
			if πTemp018 {
				goto Label19
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßargs); πE != nil {
				continue
			}
			if πTemp019, πE = πg.GetAttr(πF, πTemp016, ßtest, nil); πE != nil {
				continue
			}
			if πTemp018, πE = πg.IsTrue(πF, πTemp019); πE != nil {
				continue
			}
			if πTemp018 {
				goto Label20
			}
			goto Label21
			// line 970: if args.doc:
			πF.SetLineno(970)
		Label17:
			// line 971: print(__doc__)
			πF.SetLineno(971)
			πTemp002 = πF.MakeArgs(1)
			if πTemp016, πE = πg.ResolveGlobal(πF, ß__doc__); πE != nil {
				continue
			}
			πTemp002[0] = πTemp016
			if πTemp016, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
				continue
			}
			if πTemp019, πE = πTemp016.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			goto Label22
			// line 972: elif args.actionhelp:
			πF.SetLineno(972)
		Label18:
			// line 973: print(options)
			πF.SetLineno(973)
			πTemp002 = πF.MakeArgs(1)
			if πTemp016, πE = πg.ResolveGlobal(πF, ßoptions); πE != nil {
				continue
			}
			πTemp002[0] = πTemp016
			if πTemp016, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
				continue
			}
			if πTemp019, πE = πTemp016.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			goto Label22
			// line 974: elif args.stylehelp:
			πF.SetLineno(974)
		Label19:
			// line 975: print()
			πF.SetLineno(975)
			if πTemp016, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
				continue
			}
			if πTemp019, πE = πTemp016.Call(πF, nil, nil); πE != nil {
				continue
			}
			// line 976: print("Available styles (primary open/close, secondary open/close)")
			πF.SetLineno(976)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("Available styles (primary open/close, secondary open/close)").ToObject()
			if πTemp016, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
				continue
			}
			if πTemp019, πE = πTemp016.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 977: print("language tag   quotes")
			πF.SetLineno(977)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("language tag   quotes").ToObject()
			if πTemp016, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
				continue
			}
			if πTemp019, πE = πTemp016.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 978: print("============   ======")
			πF.SetLineno(978)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("============   ======").ToObject()
			if πTemp016, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
				continue
			}
			if πTemp019, πE = πTemp016.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			πTemp002 = πF.MakeArgs(1)
			if πTemp019, πE = πg.ResolveGlobal(πF, ßsmartchars); πE != nil {
				continue
			}
			if πTemp020, πE = πg.GetAttr(πF, πTemp019, ßquotes, nil); πE != nil {
				continue
			}
			if πTemp019, πE = πg.GetAttr(πF, πTemp020, ßkeys, nil); πE != nil {
				continue
			}
			if πTemp020, πE = πTemp019.Call(πF, nil, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp020
			if πTemp019, πE = πg.ResolveGlobal(πF, ßsorted); πE != nil {
				continue
			}
			if πTemp020, πE = πTemp019.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πTemp016, πE = πg.Iter(πF, πTemp020); πE != nil {
				continue
			}
			πF.PushCheckpoint(24)
			πTemp018 = false
		Label23:
			if πE != nil || πR != nil {
				continue
			}
			if πTemp018 {
				πF.PopCheckpoint()
				goto Label25
			}
			if πTemp019, πE = πg.Next(πF, πTemp016); πE != nil {
				isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
				if exc != nil {
					πE = exc
				} else if isStop {
					πE = nil
					πF.RestoreExc(nil, nil)
				}
				πTemp024 = !isStop
			} else {
				πTemp024 = true
				if πE = πF.Globals().SetItem(πF, ßkey.ToObject(), πTemp019); πE != nil {
					continue
				}
			}
			if πE != nil || !πTemp024 {
				continue
			}
			πF.PushCheckpoint(23)
			// line 980: print("%-14s %s" % (key, smartchars.quotes[key]))
			πF.SetLineno(980)
			πTemp002 = πF.MakeArgs(1)
			if πTemp025, πE = πg.ResolveGlobal(πF, ßkey); πE != nil {
				continue
			}
			if πTemp028, πE = πg.ResolveGlobal(πF, ßkey); πE != nil {
				continue
			}
			πTemp027 = πTemp028
			if πTemp030, πE = πg.ResolveGlobal(πF, ßsmartchars); πE != nil {
				continue
			}
			if πTemp031, πE = πg.GetAttr(πF, πTemp030, ßquotes, nil); πE != nil {
				continue
			}
			if πTemp028, πE = πg.GetItem(πF, πTemp031, πTemp027); πE != nil {
				continue
			}
			πTemp020 = πg.NewTuple2(πTemp025, πTemp028).ToObject()
			if πTemp019, πE = πg.Mod(πF, πg.NewStr("%-14s %s").ToObject(), πTemp020); πE != nil {
				continue
			}
			πTemp002[0] = πTemp019
			if πTemp019, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
				continue
			}
			if πTemp020, πE = πTemp019.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			continue
		Label24:
			if πE != nil || πR != nil {
				continue
			}
		Label25:
			goto Label22
			// line 981: elif args.test:
			πF.SetLineno(981)
		Label20:
			// line 983: import unittest
			πF.SetLineno(983)
			if πTemp002, πE = πg.ImportModule(πF, "unittest"); πE != nil {
				continue
			}
			πTemp016 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßunittest.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 985: class TestSmartypantsAllAttributes(unittest.TestCase):
			πF.SetLineno(985)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp020, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
				continue
			}
			if πTemp025, πE = πg.GetAttr(πF, πTemp020, ßTestCase, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp025
			πTemp003 = πg.NewDict()
			if πTemp016, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp016); πE != nil {
				continue
			}
			_, πE = πg.NewCode("TestSmartypantsAllAttributes", "/usr/lib/python2.7/site-packages/docutils/utils/smartquotes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 987: def test_dates(self):
					πF.SetLineno(987)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_dates", "/usr/lib/python2.7/site-packages/docutils/utils/smartquotes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 988: self.assertEqual(smartyPants("1440-80's"), u"1440-80’s")
							πF.SetLineno(988)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("1440-80's").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsmartyPants); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewUnicode("1440-80\xe2\x80\x99s").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 989: self.assertEqual(smartyPants("1440-'80s"), u"1440-’80s")
							πF.SetLineno(989)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("1440-'80s").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsmartyPants); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewUnicode("1440-\xe2\x80\x9980s").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 990: self.assertEqual(smartyPants("1440---'80s"), u"1440–’80s")
							πF.SetLineno(990)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("1440---'80s").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsmartyPants); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewUnicode("1440\xe2\x80\x93\xe2\x80\x9980s").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 991: self.assertEqual(smartyPants("1960's"), u"1960’s")
							πF.SetLineno(991)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("1960's").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsmartyPants); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewUnicode("1960\xe2\x80\x99s").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 992: self.assertEqual(smartyPants("one two '60s"), u"one two ’60s")
							πF.SetLineno(992)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("one two '60s").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsmartyPants); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewUnicode("one two \xe2\x80\x9960s").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 993: self.assertEqual(smartyPants("'60s"), u"’60s")
							πF.SetLineno(993)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("'60s").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsmartyPants); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewUnicode("\xe2\x80\x9960s").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_dates.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 995: def test_educated_quotes(self):
					πF.SetLineno(995)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_educated_quotes", "/usr/lib/python2.7/site-packages/docutils/utils/smartquotes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 996: self.assertEqual(smartyPants('"Isn\'t this fun?"'), u'“Isn’t this fun?”')
							πF.SetLineno(996)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("\"Isn't this fun?\"").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsmartyPants); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewUnicode("\xe2\x80\x9cIsn\xe2\x80\x99t this fun?\xe2\x80\x9d").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_educated_quotes.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 998: def test_html_tags(self):
					πF.SetLineno(998)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_html_tags", "/usr/lib/python2.7/site-packages/docutils/utils/smartquotes.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µtext *πg.Object = πg.UnboundLocal
						_ = µtext
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
							// line 999: text = '<a src="foo">more</a>'
							πF.SetLineno(999)
							µtext = πg.NewStr("<a src=\"foo\">more</a>").ToObject()
							// line 1000: self.assertEqual(smartyPants(text), text)
							πF.SetLineno(1000)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp002[0] = µtext
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsmartyPants); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp001[1] = µtext
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_html_tags.ToObject(), πTemp004); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp019, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp019 == nil {
				πTemp019 = πg.TypeType.ToObject()
			}
			if πTemp020, πE = πTemp019.Call(πF, []*πg.Object{πg.NewStr("TestSmartypantsAllAttributes").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTestSmartypantsAllAttributes.ToObject(), πTemp020); πE != nil {
				continue
			}
			// line 1002: suite = unittest.TestLoader().loadTestsFromTestCase(
			πF.SetLineno(1002)
			πTemp002 = πF.MakeArgs(1)
			if πTemp016, πE = πg.ResolveGlobal(πF, ßTestSmartypantsAllAttributes); πE != nil {
				continue
			}
			πTemp002[0] = πTemp016
			if πTemp016, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
				continue
			}
			if πTemp019, πE = πg.GetAttr(πF, πTemp016, ßTestLoader, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp019.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp019, πE = πg.GetAttr(πF, πTemp016, ßloadTestsFromTestCase, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp019.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßsuite.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 1004: unittest.TextTestRunner().run(suite)
			πF.SetLineno(1004)
			πTemp002 = πF.MakeArgs(1)
			if πTemp016, πE = πg.ResolveGlobal(πF, ßsuite); πE != nil {
				continue
			}
			πTemp002[0] = πTemp016
			if πTemp016, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
				continue
			}
			if πTemp019, πE = πg.GetAttr(πF, πTemp016, ßTextTestRunner, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp019.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp019, πE = πg.GetAttr(πF, πTemp016, ßrun, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp019.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			goto Label22
		Label21:
			if πTemp016, πE = πg.ResolveGlobal(πF, ßargs); πE != nil {
				continue
			}
			if πTemp019, πE = πg.GetAttr(πF, πTemp016, ßalternative_quotes, nil); πE != nil {
				continue
			}
			if πTemp018, πE = πg.IsTrue(πF, πTemp019); πE != nil {
				continue
			}
			if πTemp018 {
				goto Label26
			}
			goto Label27
			// line 1007: if args.alternative_quotes:
			πF.SetLineno(1007)
		Label26:
			if πTemp019, πE = πg.ResolveGlobal(πF, ßargs); πE != nil {
				continue
			}
			if πTemp020, πE = πg.GetAttr(πF, πTemp019, ßlanguage, nil); πE != nil {
				continue
			}
			if πTemp018, πE = πg.Contains(πF, πTemp020, πg.NewStr("-x-altquot").ToObject()); πE != nil {
				continue
			}
			πTemp016 = πg.GetBool(πTemp018).ToObject()
			if πTemp018, πE = πg.IsTrue(πF, πTemp016); πE != nil {
				continue
			}
			if πTemp018 {
				goto Label28
			}
			goto Label29
			// line 1008: if '-x-altquot' in args.language:
			πF.SetLineno(1008)
		Label28:
			// line 1009: args.language = args.language.replace('-x-altquot', '')
			πF.SetLineno(1009)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = πg.NewStr("-x-altquot").ToObject()
			πTemp002[1] = ß.ToObject()
			if πTemp016, πE = πg.ResolveGlobal(πF, ßargs); πE != nil {
				continue
			}
			if πTemp019, πE = πg.GetAttr(πF, πTemp016, ßlanguage, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πg.GetAttr(πF, πTemp019, ßreplace, nil); πE != nil {
				continue
			}
			if πTemp019, πE = πTemp016.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp016}, πTemp019); πE != nil {
				continue
			}
			if πTemp020, πE = πg.ResolveGlobal(πF, ßargs); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp020, ßlanguage, πTemp016); πE != nil {
				continue
			}
			goto Label30
		Label29:
			// line 1011: args.language += '-x-altquot'
			πF.SetLineno(1011)
			if πTemp016, πE = πg.ResolveGlobal(πF, ßargs); πE != nil {
				continue
			}
			if πTemp019, πE = πg.GetAttr(πF, πTemp016, ßlanguage, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πg.IAdd(πF, πTemp019, πg.NewStr("-x-altquot").ToObject()); πE != nil {
				continue
			}
			if πTemp020, πE = πg.ResolveGlobal(πF, ßargs); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp020, ßlanguage, πTemp016); πE != nil {
				continue
			}
			goto Label30
		Label30:
			goto Label27
		Label27:
			// line 1012: text = sys.stdin.read().decode(args.encoding)
			πF.SetLineno(1012)
			πTemp002 = πF.MakeArgs(1)
			if πTemp016, πE = πg.ResolveGlobal(πF, ßargs); πE != nil {
				continue
			}
			if πTemp019, πE = πg.GetAttr(πF, πTemp016, ßencoding, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp019
			if πTemp016, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp019, πE = πg.GetAttr(πF, πTemp016, ßstdin, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πg.GetAttr(πF, πTemp019, ßread, nil); πE != nil {
				continue
			}
			if πTemp019, πE = πTemp016.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πg.GetAttr(πF, πTemp019, ßdecode, nil); πE != nil {
				continue
			}
			if πTemp019, πE = πTemp016.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßtext.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 1013: print(smartyPants(text, attr=args.action,
			πF.SetLineno(1013)
			πTemp002 = πF.MakeArgs(1)
			πTemp023 = πF.MakeArgs(1)
			if πTemp016, πE = πg.ResolveGlobal(πF, ßargs); πE != nil {
				continue
			}
			if πTemp019, πE = πg.GetAttr(πF, πTemp016, ßencoding, nil); πE != nil {
				continue
			}
			πTemp023[0] = πTemp019
			πTemp032 = πF.MakeArgs(1)
			if πTemp016, πE = πg.ResolveGlobal(πF, ßtext); πE != nil {
				continue
			}
			πTemp032[0] = πTemp016
			if πTemp016, πE = πg.ResolveGlobal(πF, ßargs); πE != nil {
				continue
			}
			if πTemp019, πE = πg.GetAttr(πF, πTemp016, ßaction, nil); πE != nil {
				continue
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßargs); πE != nil {
				continue
			}
			if πTemp020, πE = πg.GetAttr(πF, πTemp016, ßlanguage, nil); πE != nil {
				continue
			}
			πTemp029 = πg.KWArgs{
				{"attr", πTemp019},
				{"language", πTemp020},
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßsmartyPants); πE != nil {
				continue
			}
			if πTemp019, πE = πTemp016.Call(πF, πTemp032, πTemp029); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp032)
			if πTemp016, πE = πg.GetAttr(πF, πTemp019, ßencode, nil); πE != nil {
				continue
			}
			if πTemp019, πE = πTemp016.Call(πF, πTemp023, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp023)
			πTemp002[0] = πTemp019
			if πTemp016, πE = πg.ResolveGlobal(πF, ßprint); πE != nil {
				continue
			}
			if πTemp019, πE = πTemp016.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			goto Label22
		Label22:
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("smartquotes", Code)
}
