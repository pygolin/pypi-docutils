package urischemes

import (
	πg "github.com/pygolin/runtime"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/site-packages/docutils/utils/urischemes.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ß__doc__ := πg.InternStr("__doc__")
		ßabout := πg.InternStr("about")
		ßacap := πg.InternStr("acap")
		ßaddbook := πg.InternStr("addbook")
		ßafp := πg.InternStr("afp")
		ßafs := πg.InternStr("afs")
		ßaim := πg.InternStr("aim")
		ßcallto := πg.InternStr("callto")
		ßcastanet := πg.InternStr("castanet")
		ßchttp := πg.InternStr("chttp")
		ßcid := πg.InternStr("cid")
		ßcrid := πg.InternStr("crid")
		ßdata := πg.InternStr("data")
		ßdav := πg.InternStr("dav")
		ßdict := πg.InternStr("dict")
		ßdns := πg.InternStr("dns")
		ßeid := πg.InternStr("eid")
		ßfax := πg.InternStr("fax")
		ßfeed := πg.InternStr("feed")
		ßfile := πg.InternStr("file")
		ßfinger := πg.InternStr("finger")
		ßfreenet := πg.InternStr("freenet")
		ßftp := πg.InternStr("ftp")
		ßgo := πg.InternStr("go")
		ßgopher := πg.InternStr("gopher")
		ßh323 := πg.InternStr("h323")
		ßh324 := πg.InternStr("h324")
		ßhdl := πg.InternStr("hdl")
		ßhnews := πg.InternStr("hnews")
		ßhttp := πg.InternStr("http")
		ßhttps := πg.InternStr("https")
		ßhydra := πg.InternStr("hydra")
		ßiioploc := πg.InternStr("iioploc")
		ßilu := πg.InternStr("ilu")
		ßim := πg.InternStr("im")
		ßimap := πg.InternStr("imap")
		ßinfo := πg.InternStr("info")
		ßior := πg.InternStr("ior")
		ßipp := πg.InternStr("ipp")
		ßirc := πg.InternStr("irc")
		ßiseek := πg.InternStr("iseek")
		ßjar := πg.InternStr("jar")
		ßjavascript := πg.InternStr("javascript")
		ßjdbc := πg.InternStr("jdbc")
		ßldap := πg.InternStr("ldap")
		ßlifn := πg.InternStr("lifn")
		ßlivescript := πg.InternStr("livescript")
		ßlrq := πg.InternStr("lrq")
		ßmailbox := πg.InternStr("mailbox")
		ßmailserver := πg.InternStr("mailserver")
		ßmailto := πg.InternStr("mailto")
		ßmd5 := πg.InternStr("md5")
		ßmid := πg.InternStr("mid")
		ßmocha := πg.InternStr("mocha")
		ßmodem := πg.InternStr("modem")
		ßmtqp := πg.InternStr("mtqp")
		ßmupdate := πg.InternStr("mupdate")
		ßnews := πg.InternStr("news")
		ßnfs := πg.InternStr("nfs")
		ßnntp := πg.InternStr("nntp")
		ßopaquelocktoken := πg.InternStr("opaquelocktoken")
		ßphone := πg.InternStr("phone")
		ßpop := πg.InternStr("pop")
		ßpop3 := πg.InternStr("pop3")
		ßpres := πg.InternStr("pres")
		ßprinter := πg.InternStr("printer")
		ßprospero := πg.InternStr("prospero")
		ßrdar := πg.InternStr("rdar")
		ßres := πg.InternStr("res")
		ßrtsp := πg.InternStr("rtsp")
		ßrvp := πg.InternStr("rvp")
		ßrwhois := πg.InternStr("rwhois")
		ßrx := πg.InternStr("rx")
		ßschemes := πg.InternStr("schemes")
		ßsdp := πg.InternStr("sdp")
		ßservice := πg.InternStr("service")
		ßshttp := πg.InternStr("shttp")
		ßsip := πg.InternStr("sip")
		ßsips := πg.InternStr("sips")
		ßsmb := πg.InternStr("smb")
		ßsnews := πg.InternStr("snews")
		ßsnmp := πg.InternStr("snmp")
		ßssh := πg.InternStr("ssh")
		ßt120 := πg.InternStr("t120")
		ßtag := πg.InternStr("tag")
		ßtcp := πg.InternStr("tcp")
		ßtel := πg.InternStr("tel")
		ßtelephone := πg.InternStr("telephone")
		ßtelnet := πg.InternStr("telnet")
		ßtftp := πg.InternStr("tftp")
		ßtip := πg.InternStr("tip")
		ßtn3270 := πg.InternStr("tn3270")
		ßtv := πg.InternStr("tv")
		ßurn := πg.InternStr("urn")
		ßuuid := πg.InternStr("uuid")
		ßvemmi := πg.InternStr("vemmi")
		ßvideotex := πg.InternStr("videotex")
		ßwais := πg.InternStr("wais")
		ßwhodp := πg.InternStr("whodp")
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
			// line 5: """
			πF.SetLineno(5)
			// line 5: """
			πF.SetLineno(5)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n`schemes` is a dictionary with lowercase URI addressing schemes as\nkeys and descriptions as values. It was compiled from the index at\nhttp://www.iana.org/assignments/uri-schemes (revised 2005-11-28)\nand an older list at http://www.w3.org/Addressing/schemes.html.\n").ToObject()); πE != nil {
				continue
			}
			// line 14: schemes = {
			πF.SetLineno(14)
			πTemp001 = πg.NewDict()
			if πE = πTemp001.SetItem(πF, ßabout.ToObject(), πg.NewStr("provides information on Navigator").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßacap.ToObject(), πg.NewStr("Application Configuration Access Protocol; RFC 2244").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßaddbook.ToObject(), πg.NewStr("To add vCard entries to Communicator's Address Book").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßafp.ToObject(), πg.NewStr("Apple Filing Protocol").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßafs.ToObject(), πg.NewStr("Andrew File System global file names").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßaim.ToObject(), πg.NewStr("AOL Instant Messenger").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcallto.ToObject(), πg.NewStr("for NetMeeting links").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcastanet.ToObject(), πg.NewStr("Castanet Tuner URLs for Netcaster").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßchttp.ToObject(), πg.NewStr("cached HTTP supported by RealPlayer").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcid.ToObject(), πg.NewStr("content identifier; RFC 2392").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßcrid.ToObject(), πg.NewStr("TV-Anytime Content Reference Identifier; RFC 4078").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdata.ToObject(), πg.NewStr("allows inclusion of small data items as \"immediate\" data; RFC 2397").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdav.ToObject(), πg.NewStr("Distributed Authoring and Versioning Protocol; RFC 2518").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdict.ToObject(), πg.NewStr("dictionary service protocol; RFC 2229").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßdns.ToObject(), πg.NewStr("Domain Name System resources").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßeid.ToObject(), πg.NewStr("External ID; non-URL data; general escape mechanism to allow access to information for applications that are too specialized to justify their own schemes").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßfax.ToObject(), πg.NewStr("a connection to a terminal that can handle telefaxes (facsimiles); RFC 2806").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßfeed.ToObject(), πg.NewStr("NetNewsWire feed").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßfile.ToObject(), πg.NewStr("Host-specific file names; RFC 1738").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßfinger.ToObject(), ß.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßfreenet.ToObject(), ß.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßftp.ToObject(), πg.NewStr("File Transfer Protocol; RFC 1738").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßgo.ToObject(), πg.NewStr("go; RFC 3368").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßgopher.ToObject(), πg.NewStr("The Gopher Protocol").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("gsm-sms").ToObject(), πg.NewStr("Global System for Mobile Communications Short Message Service").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßh323.ToObject(), πg.NewStr("video (audiovisual) communication on local area networks; RFC 3508").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßh324.ToObject(), πg.NewStr("video and audio communications over low bitrate connections such as POTS modem connections").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßhdl.ToObject(), πg.NewStr("CNRI handle system").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßhnews.ToObject(), πg.NewStr("an HTTP-tunneling variant of the NNTP news protocol").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßhttp.ToObject(), πg.NewStr("Hypertext Transfer Protocol; RFC 2616").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßhttps.ToObject(), πg.NewStr("HTTP over SSL; RFC 2818").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßhydra.ToObject(), πg.NewStr("SubEthaEdit URI.  See http://www.codingmonkeys.de/subethaedit.").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßiioploc.ToObject(), πg.NewStr("Internet Inter-ORB Protocol Location?").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßilu.ToObject(), πg.NewStr("Inter-Language Unification").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßim.ToObject(), πg.NewStr("Instant Messaging; RFC 3860").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßimap.ToObject(), πg.NewStr("Internet Message Access Protocol; RFC 2192").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßinfo.ToObject(), πg.NewStr("Information Assets with Identifiers in Public Namespaces").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßior.ToObject(), πg.NewStr("CORBA interoperable object reference").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßipp.ToObject(), πg.NewStr("Internet Printing Protocol; RFC 3510").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßirc.ToObject(), πg.NewStr("Internet Relay Chat").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("iris.beep").ToObject(), πg.NewStr("iris.beep; RFC 3983").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßiseek.ToObject(), πg.NewStr("See www.ambrosiasw.com;  a little util for OS X.").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßjar.ToObject(), πg.NewStr("Java archive").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßjavascript.ToObject(), πg.NewStr("JavaScript code; evaluates the expression after the colon").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßjdbc.ToObject(), πg.NewStr("JDBC connection URI.").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßldap.ToObject(), πg.NewStr("Lightweight Directory Access Protocol").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlifn.ToObject(), ß.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlivescript.ToObject(), ß.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßlrq.ToObject(), ß.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßmailbox.ToObject(), πg.NewStr("Mail folder access").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßmailserver.ToObject(), πg.NewStr("Access to data available from mail servers").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßmailto.ToObject(), πg.NewStr("Electronic mail address; RFC 2368").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßmd5.ToObject(), ß.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßmid.ToObject(), πg.NewStr("message identifier; RFC 2392").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßmocha.ToObject(), ß.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßmodem.ToObject(), πg.NewStr("a connection to a terminal that can handle incoming data calls; RFC 2806").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßmtqp.ToObject(), πg.NewStr("Message Tracking Query Protocol; RFC 3887").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßmupdate.ToObject(), πg.NewStr("Mailbox Update (MUPDATE) Protocol; RFC 3656").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnews.ToObject(), πg.NewStr("USENET news; RFC 1738").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnfs.ToObject(), πg.NewStr("Network File System protocol; RFC 2224").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßnntp.ToObject(), πg.NewStr("USENET news using NNTP access; RFC 1738").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßopaquelocktoken.ToObject(), πg.NewStr("RFC 2518").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßphone.ToObject(), ß.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßpop.ToObject(), πg.NewStr("Post Office Protocol; RFC 2384").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßpop3.ToObject(), πg.NewStr("Post Office Protocol v3").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßpres.ToObject(), πg.NewStr("Presence; RFC 3859").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßprinter.ToObject(), ß.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßprospero.ToObject(), πg.NewStr("Prospero Directory Service; RFC 4157").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrdar.ToObject(), πg.NewStr("URLs found in Darwin source (http://www.opensource.apple.com/darwinsource/).").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßres.ToObject(), ß.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrtsp.ToObject(), πg.NewStr("real time streaming protocol; RFC 2326").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrvp.ToObject(), ß.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrwhois.ToObject(), ß.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßrx.ToObject(), πg.NewStr("Remote Execution").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsdp.ToObject(), ß.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßservice.ToObject(), πg.NewStr("service location; RFC 2609").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßshttp.ToObject(), πg.NewStr("secure hypertext transfer protocol").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsip.ToObject(), πg.NewStr("Session Initiation Protocol; RFC 3261").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsips.ToObject(), πg.NewStr("secure session intitiaion protocol; RFC 3261").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsmb.ToObject(), πg.NewStr("SAMBA filesystems.").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsnews.ToObject(), πg.NewStr("For NNTP postings via SSL").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßsnmp.ToObject(), πg.NewStr("Simple Network Management Protocol; RFC 4088").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("soap.beep").ToObject(), πg.NewStr("RFC 3288").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("soap.beeps").ToObject(), πg.NewStr("RFC 3288").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßssh.ToObject(), πg.NewStr("Reference to interactive sessions via ssh.").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßt120.ToObject(), πg.NewStr("real time data conferencing (audiographics)").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtag.ToObject(), πg.NewStr("RFC 4151").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtcp.ToObject(), ß.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtel.ToObject(), πg.NewStr("a connection to a terminal that handles normal voice telephone calls, a voice mailbox or another voice messaging system or a service that can be operated using DTMF tones; RFC 3966.").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtelephone.ToObject(), ßtelephone.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtelnet.ToObject(), πg.NewStr("Reference to interactive sessions; RFC 4248").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtftp.ToObject(), πg.NewStr("Trivial File Transfer Protocol; RFC 3617").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtip.ToObject(), πg.NewStr("Transaction Internet Protocol; RFC 2371").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtn3270.ToObject(), πg.NewStr("Interactive 3270 emulation sessions").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßtv.ToObject(), ß.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßurn.ToObject(), πg.NewStr("Uniform Resource Name; RFC 2141").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßuuid.ToObject(), ß.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvemmi.ToObject(), πg.NewStr("versatile multimedia interface; RFC 2122").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßvideotex.ToObject(), ß.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("view-source").ToObject(), πg.NewStr("displays HTML code that was generated with JavaScript").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßwais.ToObject(), πg.NewStr("Wide Area Information Servers; RFC 4156").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, ßwhodp.ToObject(), ß.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("whois++").ToObject(), πg.NewStr("Distributed directory service.").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("x-man-page").ToObject(), πg.NewStr("Opens man page in Terminal.app on OS X (see macosxhints.com)").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("xmlrpc.beep").ToObject(), πg.NewStr("RFC 3529").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("xmlrpc.beeps").ToObject(), πg.NewStr("RFC 3529").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("z39.50r").ToObject(), πg.NewStr("Z39.50 Retrieval; RFC 2056").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp001.SetItem(πF, πg.NewStr("z39.50s").ToObject(), πg.NewStr("Z39.50 Session; RFC 2056").ToObject()); πE != nil {
				continue
			}
			πTemp002 = πTemp001.ToObject()
			if πE = πF.Globals().SetItem(πF, ßschemes.ToObject(), πTemp002); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("urischemes", Code)
}
