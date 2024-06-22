package mimetype

import "github.com/bayashi/qq/dictionary"

// https://fastapi.metacpan.org/source/MARKOV/MIME-Types-2.26/lib/MIME/types.db
// curl https://fastapi.metacpan.org/source/MARKOV/MIME-Types-2.26/lib/MIME/types.db| perl -an -e 'chomp $_; if($_ eq ""){ exit; } if ($_ ne "") { my($k, $v)=split(/;/, $_); print qq|\t"$k": {\n\t\tSubject: "$v",\n\t\tDescription: "",\n\t},\n|; }'
var Ref = map[string]dictionary.Element{
	"#": {
		Subject:     "application/x-mathcad",
		Description: "",
	},
	"%": {
		Subject:     "application/x-trash",
		Description: "",
	},
	"0fs": {
		Subject:     "application/vnd.erofs",
		Description: "",
	},
	"123": {
		Subject:     "application/vnd.lotus-1-2-3",
		Description: "",
	},
	"1": {
		Subject:     "application/vnd.ieee.1905",
		Description: "",
	},
	"1clr": {
		Subject:     "application/clr",
		Description: "",
	},
	"1km": {
		Subject:     "application/vnd.1000minds.decision-model+xml",
		Description: "",
	},
	"210": {
		Subject:     "application/p21",
		Description: "",
	},
	"323": {
		Subject:     "text/x-h323",
		Description: "",
	},
	"32x": {
		Subject:     "application/x-genesis-rom",
		Description: "",
	},
	"3": {
		Subject:     "application/yaml",
		Description: "",
	},
	"3MF": {
		Subject:     "application/vnd.ms-3mfdocument",
		Description: "",
	},
	"3dm": {
		Subject:     "model/vnd.flatland.3dml",
		Description: "",
	},
	"3dmf": {
		Subject:     "x-world/x-3dmf",
		Description: "",
	},
	"3dml": {
		Subject:     "model/vnd.flatland.3dml",
		Description: "",
	},
	"3ds": {
		Subject:     "image/x-3ds",
		Description: "",
	},
	"3g2": {
		Subject:     "video/3gpp2",
		Description: "",
	},
	"3ga": {
		Subject:     "video/3gpp",
		Description: "",
	},
	"3gp2": {
		Subject:     "video/3gpp2",
		Description: "",
	},
	"3gp": {
		Subject:     "video/3gpp",
		Description: "",
	},
	"3gpp2": {
		Subject:     "video/3gpp2",
		Description: "",
	},
	"3gpp": {
		Subject:     "video/3gpp",
		Description: "",
	},
	"3mf": {
		Subject:     "model/3mf",
		Description: "",
	},
	"3tz": {
		Subject:     "application/vnd.maxar.archive.3tz+zip",
		Description: "",
	},
	"602": {
		Subject:     "application/x-t602",
		Description: "",
	},
	"669": {
		Subject:     "audio/x-mod",
		Description: "",
	},
	"726": {
		Subject:     "audio/32kadpcm",
		Description: "",
	},
	"7z": {
		Subject:     "application/x-7z-compressed",
		Description: "",
	},
	"AIT": {
		Subject:     "application/vnd.dvb.ait",
		Description: "",
	},
	"AMR": {
		Subject:     "audio/amr",
		Description: "",
	},
	"AWB": {
		Subject:     "audio/amr-wb",
		Description: "",
	},
	"BDM": {
		Subject:     "application/vnd.syncml.dm+wbxml",
		Description: "",
	},
	"BED": {
		Subject:     "application/vnd.realvnc.bed",
		Description: "",
	},
	"BLEND": {
		Subject:     "application/x-blender",
		Description: "",
	},
	"BMI": {
		Subject:     "application/vnd.bmi",
		Description: "",
	},
	"BOX": {
		Subject:     "application/vnd.previewsystems.box",
		Description: "",
	},
	"C": {
		Subject:     "text/x-c++src",
		Description: "",
	},
	"CAB": {
		Subject:     "application/vnd.ubisoft.webplayer",
		Description: "",
	},
	"CER": {
		Subject:     "application/pkix-cert",
		Description: "",
	},
	"CHRT": {
		Subject:     "application/vnd.kde.kchart",
		Description: "",
	},
	"CQL": {
		Subject:     "text/cql",
		Description: "",
	},
	"CRL": {
		Subject:     "application/pkix-crl",
		Description: "",
	},
	"CSV": {
		Subject:     "text/csv",
		Description: "",
	},
	"DAF": {
		Subject:     "application/vnd.mobius.daf",
		Description: "",
	},
	"DIS": {
		Subject:     "application/vnd.mobius.dis",
		Description: "",
	},
	"EDM": {
		Subject:     "application/vnd.novadigm.edm",
		Description: "",
	},
	"EDX": {
		Subject:     "application/vnd.novadigm.edx",
		Description: "",
	},
	"ENW": {
		Subject:     "audio/evrcnw",
		Description: "",
	},
	"EP": {
		Subject:     "application/vnd.bluetooth.ep.oob",
		Description: "",
	},
	"EVB": {
		Subject:     "audio/evrcb",
		Description: "",
	},
	"EVC": {
		Subject:     "audio/evrc",
		Description: "",
	},
	"EVW": {
		Subject:     "audio/evrcwb",
		Description: "",
	},
	"EXT": {
		Subject:     "application/vnd.novadigm.ext",
		Description: "",
	},
	"FLW": {
		Subject:     "application/vnd.kde.kivio",
		Description: "",
	},
	"FNC": {
		Subject:     "application/vnd.frogans.fnc",
		Description: "",
	},
	"FO": {
		Subject:     "application/vnd.software602.filler.form+xml",
		Description: "",
	},
	"FTC": {
		Subject:     "application/vnd.fluxtime.clip",
		Description: "",
	},
	"HPUB": {
		Subject:     "application/prs.hpub+zip",
		Description: "",
	},
	"J2C": {
		Subject:     "image/j2c",
		Description: "",
	},
	"J2K": {
		Subject:     "image/j2c",
		Description: "",
	},
	"KARBON": {
		Subject:     "application/vnd.kde.karbon",
		Description: "",
	},
	"KFO": {
		Subject:     "application/vnd.kde.kformula",
		Description: "",
	},
	"KON": {
		Subject:     "application/vnd.kde.kontour",
		Description: "",
	},
	"KPR": {
		Subject:     "application/vnd.kde.kpresenter",
		Description: "",
	},
	"KPT": {
		Subject:     "application/vnd.kde.kpresenter",
		Description: "",
	},
	"KSP": {
		Subject:     "application/vnd.kde.kspread",
		Description: "",
	},
	"KWD": {
		Subject:     "application/vnd.kde.kword",
		Description: "",
	},
	"KWT": {
		Subject:     "application/vnd.kde.kword",
		Description: "",
	},
	"L16": {
		Subject:     "audio/l16",
		Description: "",
	},
	"LBC": {
		Subject:     "audio/ilbc",
		Description: "",
	},
	"LE": {
		Subject:     "application/vnd.bluetooth.le.oob",
		Description: "",
	},
	"LTF": {
		Subject:     "application/vnd.frogans.ltf",
		Description: "",
	},
	"MBK": {
		Subject:     "application/vnd.mobius.mbk",
		Description: "",
	},
	"MC1": {
		Subject:     "application/vnd.medcalcdata",
		Description: "",
	},
	"MCD": {
		Subject:     "application/vnd.mcd",
		Description: "",
	},
	"MQY": {
		Subject:     "application/vnd.mobius.mqy",
		Description: "",
	},
	"MSL": {
		Subject:     "application/vnd.mobius.msl",
		Description: "",
	},
	"NND": {
		Subject:     "application/vnd.noblenet-directory",
		Description: "",
	},
	"NNS": {
		Subject:     "application/vnd.noblenet-sealer",
		Description: "",
	},
	"NNW": {
		Subject:     "application/vnd.noblenet-web",
		Description: "",
	},
	"ORQ": {
		Subject:     "application/ocsp-request",
		Description: "",
	},
	"ORS": {
		Subject:     "application/ocsp-response",
		Description: "",
	},
	"P2P": {
		Subject:     "application/vnd.wfa.p2p",
		Description: "",
	},
	"PAR2": {
		Subject:     "application/x-par2",
		Description: "",
	},
	"PFR": {
		Subject:     "application/font-tdpfr",
		Description: "",
	},
	"PGB": {
		Subject:     "image/vnd.globalgraphics.pgb",
		Description: "",
	},
	"PKI": {
		Subject:     "application/pkixcmp",
		Description: "",
	},
	"PL": {
		Subject:     "application/x-perl",
		Description: "",
	},
	"PLC": {
		Subject:     "application/vnd.mobius.plc",
		Description: "",
	},
	"PPD": {
		Subject:     "application/vnd.cups-ppd",
		Description: "",
	},
	"QBO": {
		Subject:     "application/vnd.intu.qbo",
		Description: "",
	},
	"QCP": {
		Subject:     "audio/evrc-qcp",
		Description: "",
	},
	"QFX": {
		Subject:     "application/vnd.intu.qfx",
		Description: "",
	},
	"SAR": {
		Subject:     "application/vnd.sar",
		Description: "",
	},
	"SCQ": {
		Subject:     "application/scvp-cv-request",
		Description: "",
	},
	"SCS": {
		Subject:     "application/scvp-cv-response",
		Description: "",
	},
	"SMV": {
		Subject:     "audio/smv",
		Description: "",
	},
	"SPP": {
		Subject:     "application/scvp-vp-response",
		Description: "",
	},
	"SPQ": {
		Subject:     "application/scvp-vp-request",
		Description: "",
	},
	"SSE": {
		Subject:     "application/vnd.kodak-descriptor",
		Description: "",
	},
	"T38": {
		Subject:     "image/t38",
		Description: "",
	},
	"TFX": {
		Subject:     "image/tiff-fx",
		Description: "",
	},
	"TIF": {
		Subject:     "image/tiff",
		Description: "",
	},
	"TXF": {
		Subject:     "application/vnd.mobius.txf",
		Description: "",
	},
	"Text": {
		Subject:     "application/vnd.cncf.helm.chart.provenance.v1.prov",
		Description: "",
	},
	"UTZ": {
		Subject:     "application/vnd.uiq.theme",
		Description: "",
	},
	"VBOX": {
		Subject:     "application/vnd.previewsystems.box",
		Description: "",
	},
	"VES": {
		Subject:     "application/vnd.ves.encrypted",
		Description: "",
	},
	"VFK": {
		Subject:     "text/vnd.exchangeable",
		Description: "",
	},
	"VPM": {
		Subject:     "multipart/voice-message",
		Description: "",
	},
	"VWX": {
		Subject:     "application/vnd.vectorworks",
		Description: "",
	},
	"WAV": {
		Subject:     "audio/l16",
		Description: "",
	},
	"WSC": {
		Subject:     "application/vnd.wfa.wsc",
		Description: "",
	},
	"XAR": {
		Subject:     "application/vnd.xara",
		Description: "",
	},
	"XDM": {
		Subject:     "application/vnd.syncml.dm+xml",
		Description: "",
	},
	"XML": {
		Subject:     "application/vnd.infotech.project+xml",
		Description: "",
	},
	"XOP": {
		Subject:     "application/xop+xml",
		Description: "",
	},
	"XPR": {
		Subject:     "application/vnd.is-xpr",
		Description: "",
	},
	"XPW": {
		Subject:     "application/vnd.intercon.formnet",
		Description: "",
	},
	"XPX": {
		Subject:     "application/vnd.intercon.formnet",
		Description: "",
	},
	"XSM": {
		Subject:     "application/vnd.syncml+xml",
		Description: "",
	},
	"X_B": {
		Subject:     "model/vnd.parasolid.transmit.binary",
		Description: "",
	},
	"X_T": {
		Subject:     "model/vnd.parasolid.transmit.text",
		Description: "",
	},
	"Z": {
		Subject:     "application/x-compress",
		Description: "",
	},
	"ZFO": {
		Subject:     "application/vnd.software602.filler.form-xml-zip",
		Description: "",
	},
	"ZMM": {
		Subject:     "application/vnd.handheld-entertainment+xml",
		Description: "",
	},
	"a2l": {
		Subject:     "application/a2l",
		Description: "",
	},
	"a": {
		Subject:     "text/vnd.a",
		Description: "",
	},
	"aa3": {
		Subject:     "audio/atrac-advanced-lossless",
		Description: "",
	},
	"aab": {
		Subject:     "application/x-authorware-bin",
		Description: "",
	},
	"aac": {
		Subject:     "audio/aac",
		Description: "",
	},
	"aal": {
		Subject:     "audio/atrac-advanced-lossless",
		Description: "",
	},
	"aam": {
		Subject:     "application/x-authorware-map",
		Description: "",
	},
	"aas": {
		Subject:     "application/x-authorware-seg",
		Description: "",
	},
	"abc": {
		Subject:     "text/vnd.abc",
		Description: "",
	},
	"abw": {
		Subject:     "application/x-abiword",
		Description: "",
	},
	"ac2": {
		Subject:     "application/vnd.banana-accounting",
		Description: "",
	},
	"ac3": {
		Subject:     "audio/ac3",
		Description: "",
	},
	"ac": {
		Subject:     "application/pkix-attr-cert",
		Description: "",
	},
	"acc": {
		Subject:     "application/vnd.americandynamics.acc",
		Description: "",
	},
	"ace": {
		Subject:     "application/x-ace-compressed",
		Description: "",
	},
	"acgi": {
		Subject:     "text/html",
		Description: "",
	},
	"acu": {
		Subject:     "application/vnd.acucobol",
		Description: "",
	},
	"acutc": {
		Subject:     "application/vnd.acucorp",
		Description: "",
	},
	"adb": {
		Subject:     "text/x-adasrc",
		Description: "",
	},
	"adp": {
		Subject:     "audio/x-adpcm",
		Description: "",
	},
	"ads": {
		Subject:     "text/x-adasrc",
		Description: "",
	},
	"adts": {
		Subject:     "audio/aac",
		Description: "",
	},
	"aep": {
		Subject:     "application/vnd.audiograph",
		Description: "",
	},
	"afl": {
		Subject:     "video/x-animaflex",
		Description: "",
	},
	"afm": {
		Subject:     "application/x-font-type1",
		Description: "",
	},
	"afp": {
		Subject:     "application/vnd.afpc.modca",
		Description: "",
	},
	"ag": {
		Subject:     "image/x-applix-graphics",
		Description: "",
	},
	"agb": {
		Subject:     "application/x-gba-rom",
		Description: "",
	},
	"age": {
		Subject:     "application/vnd.age",
		Description: "",
	},
	"ahead": {
		Subject:     "application/vnd.ahead.space",
		Description: "",
	},
	"ai": {
		Subject:     "application/postscript",
		Description: "",
	},
	"aif": {
		Subject:     "audio/x-aiff",
		Description: "",
	},
	"aifc": {
		Subject:     "audio/x-aiff",
		Description: "",
	},
	"aiff": {
		Subject:     "audio/x-aiff",
		Description: "",
	},
	"aiffc": {
		Subject:     "audio/x-aifc",
		Description: "",
	},
	"aim": {
		Subject:     "application/x-aim",
		Description: "",
	},
	"aion": {
		Subject:     "application/vnd.veritone.aion+json",
		Description: "",
	},
	"aip": {
		Subject:     "text/x-audiosoft-intra",
		Description: "",
	},
	"air": {
		Subject:     "application/vnd.adobe.air-application-installer-package+zip",
		Description: "",
	},
	"ait": {
		Subject:     "application/vnd.dvb.ait",
		Description: "",
	},
	"al": {
		Subject:     "application/x-perl",
		Description: "",
	},
	"alc": {
		Subject:     "x-chemical/x-alchemy",
		Description: "",
	},
	"alz": {
		Subject:     "application/x-alz",
		Description: "",
	},
	"ami": {
		Subject:     "application/vnd.amiga.ami",
		Description: "",
	},
	"aml": {
		Subject:     "application/aml",
		Description: "",
	},
	"amlx": {
		Subject:     "application/automationml-amlx+zip",
		Description: "",
	},
	"amr": {
		Subject:     "audio/amr",
		Description: "",
	},
	"amz": {
		Subject:     "audio/x-amzxml",
		Description: "",
	},
	"ani": {
		Subject:     "application/x-navi-animation",
		Description: "",
	},
	"anx": {
		Subject:     "application/x-annodex",
		Description: "",
	},
	"any": {
		Subject:     "application/vnd.mitsubishi.misty-guard.trustweb",
		Description: "",
	},
	"aos": {
		Subject:     "application/x-nokia-9000-communicator-add-on-software",
		Description: "",
	},
	"ape": {
		Subject:     "audio/x-ape",
		Description: "",
	},
	"apk": {
		Subject:     "application/vnd.android.package-archive",
		Description: "",
	},
	"apkg": {
		Subject:     "application/vnd.anki",
		Description: "",
	},
	"apng": {
		Subject:     "image/apng",
		Description: "",
	},
	"appcache": {
		Subject:     "text/cache-manifest",
		Description: "",
	},
	"application": {
		Subject:     "application/x-ms-application",
		Description: "",
	},
	"apr": {
		Subject:     "application/vnd.lotus-approach",
		Description: "",
	},
	"aps": {
		Subject:     "application/x-mime",
		Description: "",
	},
	"apxml": {
		Subject:     "application/auth-policy+xml",
		Description: "",
	},
	"ar": {
		Subject:     "application/x-archive",
		Description: "",
	},
	"arc": {
		Subject:     "application/x-freearc",
		Description: "",
	},
	"arj": {
		Subject:     "application/x-arj",
		Description: "",
	},
	"arrow": {
		Subject:     "application/vnd.apache.arrow.file",
		Description: "",
	},
	"arrows": {
		Subject:     "application/vnd.apache.arrow.stream",
		Description: "",
	},
	"art": {
		Subject:     "image/x-jg",
		Description: "",
	},
	"artisan": {
		Subject:     "application/vnd.artisan+json",
		Description: "",
	},
	"arw": {
		Subject:     "image/x-sony-arw",
		Description: "",
	},
	"as": {
		Subject:     "application/x-applix-spreadsheet",
		Description: "",
	},
	"asc": {
		Subject:     "application/pgp-keys",
		Description: "",
	},
	"ascii": {
		Subject:     "text/vnd.ascii-art",
		Description: "",
	},
	"asf": {
		Subject:     "application/vnd.ms-asf",
		Description: "",
	},
	"asice": {
		Subject:     "application/vnd.etsi.asic-e+zip",
		Description: "",
	},
	"asics": {
		Subject:     "application/vnd.etsi.asic-s+zip",
		Description: "",
	},
	"asm": {
		Subject:     "text/x-asm",
		Description: "",
	},
	"asn": {
		Subject:     "x-chemical/x-ncbi-asn1",
		Description: "",
	},
	"aso": {
		Subject:     "application/vnd.accpac.simply.aso",
		Description: "",
	},
	"asp": {
		Subject:     "text/x-asp",
		Description: "",
	},
	"ass": {
		Subject:     "audio/aac",
		Description: "",
	},
	"asx": {
		Subject:     "video/x-ms-asf",
		Description: "",
	},
	"at3": {
		Subject:     "audio/atrac3",
		Description: "",
	},
	"atc": {
		Subject:     "application/vnd.acucorp",
		Description: "",
	},
	"atf": {
		Subject:     "application/atf",
		Description: "",
	},
	"atfx": {
		Subject:     "application/atfx",
		Description: "",
	},
	"atom": {
		Subject:     "application/atom+xml",
		Description: "",
	},
	"atomcat": {
		Subject:     "application/atomcat+xml",
		Description: "",
	},
	"atomdeleted": {
		Subject:     "application/atomdeleted+xml",
		Description: "",
	},
	"atomsrv": {
		Subject:     "application/x-atomserv+xml",
		Description: "",
	},
	"atomsvc": {
		Subject:     "application/atomsvc+xml",
		Description: "",
	},
	"atx": {
		Subject:     "application/vnd.antix.game-component",
		Description: "",
	},
	"atxml": {
		Subject:     "application/atxml",
		Description: "",
	},
	"au": {
		Subject:     "audio/basic",
		Description: "",
	},
	"auc": {
		Subject:     "application/tamp-apex-update-confirm",
		Description: "",
	},
	"avci": {
		Subject:     "image/avci",
		Description: "",
	},
	"avcs": {
		Subject:     "image/avcs",
		Description: "",
	},
	"avf": {
		Subject:     "video/x-msvideo",
		Description: "",
	},
	"avi": {
		Subject:     "video/x-msvideo",
		Description: "",
	},
	"avif": {
		Subject:     "image/avif",
		Description: "",
	},
	"avs": {
		Subject:     "video/x-avs-video",
		Description: "",
	},
	"aw": {
		Subject:     "application/x-applixware",
		Description: "",
	},
	"awb": {
		Subject:     "audio/amr-wb",
		Description: "",
	},
	"awk": {
		Subject:     "application/x-awk",
		Description: "",
	},
	"axa": {
		Subject:     "audio/x-annodex",
		Description: "",
	},
	"axv": {
		Subject:     "video/x-annodex",
		Description: "",
	},
	"axx": {
		Subject:     "application/vnd.xecrets-encrypted",
		Description: "",
	},
	"azf": {
		Subject:     "application/vnd.airzip.filesecure.azf",
		Description: "",
	},
	"azs": {
		Subject:     "application/vnd.airzip.filesecure.azs",
		Description: "",
	},
	"azv": {
		Subject:     "image/vnd.airzip.accelerator.azv",
		Description: "",
	},
	"azw3": {
		Subject:     "application/vnd.amazon.mobi8-ebook",
		Description: "",
	},
	"azw": {
		Subject:     "application/vnd.amazon.ebook",
		Description: "",
	},
	"b16": {
		Subject:     "image/vnd.pco.b16",
		Description: "",
	},
	"b": {
		Subject:     "x-chemical/x-molconn-z",
		Description: "",
	},
	"bak": {
		Subject:     "application/x-trash",
		Description: "",
	},
	"bar": {
		Subject:     "application/vnd.qualcomm.brew-app-res",
		Description: "",
	},
	"bary": {
		Subject:     "model/vnd.bary",
		Description: "",
	},
	"bat": {
		Subject:     "application/x-msdos-program",
		Description: "",
	},
	"bck": {
		Subject:     "application/x-vmsbackup",
		Description: "",
	},
	"bcpio": {
		Subject:     "application/x-bcpio",
		Description: "",
	},
	"bdf": {
		Subject:     "application/x-font-bdf",
		Description: "",
	},
	"bdm": {
		Subject:     "application/vnd.syncml.dm+wbxml",
		Description: "",
	},
	"bdmv": {
		Subject:     "video/mp2t",
		Description: "",
	},
	"bdo": {
		Subject:     "application/vnd.nato.bindingdataobject+xml",
		Description: "",
	},
	"bdob": {
		Subject:     "application/vnd.nato.bindingdataobject+cbor",
		Description: "",
	},
	"bdoj": {
		Subject:     "application/vnd.nato.bindingdataobject+json",
		Description: "",
	},
	"bed": {
		Subject:     "application/vnd.realvnc.bed",
		Description: "",
	},
	"bh2": {
		Subject:     "application/vnd.fujitsu.oasysprs",
		Description: "",
	},
	"bib": {
		Subject:     "text/x-bibtex",
		Description: "",
	},
	"bik": {
		Subject:     "video/vnd.radgamettools.bink",
		Description: "",
	},
	"bin": {
		Subject:     "applicatoin/octet-stream",
		Description: "",
	},
	"bk2": {
		Subject:     "video/vnd.radgamettools.bink",
		Description: "",
	},
	"bkm": {
		Subject:     "application/vnd.nervana",
		Description: "",
	},
	"blb": {
		Subject:     "application/x-blorb",
		Description: "",
	},
	"bleep": {
		Subject:     "application/x-bleeper",
		Description: "",
	},
	"blend": {
		Subject:     "application/x-blender",
		Description: "",
	},
	"blender": {
		Subject:     "application/x-blender",
		Description: "",
	},
	"blorb": {
		Subject:     "application/x-blorb",
		Description: "",
	},
	"bm": {
		Subject:     "image/bmp",
		Description: "",
	},
	"bmed": {
		Subject:     "multipart/vnd.bint.med-plus",
		Description: "",
	},
	"bmi": {
		Subject:     "application/vnd.bmi",
		Description: "",
	},
	"bmml": {
		Subject:     "application/vnd.balsamiq.bmml+xml",
		Description: "",
	},
	"bmp": {
		Subject:     "image/bmp",
		Description: "",
	},
	"bmpr": {
		Subject:     "application/vnd.balsamiq.bmpr",
		Description: "",
	},
	"boo": {
		Subject:     "text/x-boo",
		Description: "",
	},
	"book": {
		Subject:     "application/x-maker",
		Description: "",
	},
	"box": {
		Subject:     "application/vnd.previewsystems.box",
		Description: "",
	},
	"boz": {
		Subject:     "application/x-bzip2",
		Description: "",
	},
	"bpd": {
		Subject:     "application/vnd.fints",
		Description: "",
	},
	"bpk": {
		Subject:     "application/octet-stream",
		Description: "",
	},
	"brf": {
		Subject:     "text/plain",
		Description: "",
	},
	"bsd": {
		Subject:     "x-chemical/x-crossfire",
		Description: "",
	},
	"bsh": {
		Subject:     "application/x-bsh",
		Description: "",
	},
	"bsp": {
		Subject:     "model/vnd.valve.source.compiled-map",
		Description: "",
	},
	"btf": {
		Subject:     "image/prs.btif",
		Description: "",
	},
	"btif": {
		Subject:     "image/prs.btif",
		Description: "",
	},
	"buffer": {
		Subject:     "application/octet-stream",
		Description: "",
	},
	"bz2": {
		Subject:     "application/x-bzip2",
		Description: "",
	},
	"bz3": {
		Subject:     "application/vnd.bzip3",
		Description: "",
	},
	"bz": {
		Subject:     "application/x-bzip",
		Description: "",
	},
	"c++": {
		Subject:     "text/x-c++src",
		Description: "",
	},
	"c11amc": {
		Subject:     "application/vnd.cluetrust.cartomobile-config",
		Description: "",
	},
	"c11amz": {
		Subject:     "application/vnd.cluetrust.cartomobile-config-pkg",
		Description: "",
	},
	"c2pa": {
		Subject:     "application/c2pa",
		Description: "",
	},
	"c3d": {
		Subject:     "x-chemical/x-chem3d",
		Description: "",
	},
	"c3ex": {
		Subject:     "application/cccex",
		Description: "",
	},
	"c4d": {
		Subject:     "application/vnd.clonk.c4group",
		Description: "",
	},
	"c4f": {
		Subject:     "application/vnd.clonk.c4group",
		Description: "",
	},
	"c4g": {
		Subject:     "application/vnd.clonk.c4group",
		Description: "",
	},
	"c4p": {
		Subject:     "application/vnd.clonk.c4group",
		Description: "",
	},
	"c4u": {
		Subject:     "application/vnd.clonk.c4group",
		Description: "",
	},
	"c9r": {
		Subject:     "application/vnd.cryptomator.encrypted",
		Description: "",
	},
	"c9s": {
		Subject:     "application/vnd.cryptomator.encrypted",
		Description: "",
	},
	"c": {
		Subject:     "text/x-csrc",
		Description: "",
	},
	"cab": {
		Subject:     "application/vnd.ms-cab-compressed",
		Description: "",
	},
	"cac": {
		Subject:     "x-chemical/x-cache",
		Description: "",
	},
	"cache": {
		Subject:     "x-chemical/x-cache",
		Description: "",
	},
	"caf": {
		Subject:     "audio/x-caf",
		Description: "",
	},
	"cap": {
		Subject:     "application/vnd.tcpdump.pcap",
		Description: "",
	},
	"car": {
		Subject:     "application/vnd.ipld.car",
		Description: "",
	},
	"carjson": {
		Subject:     "application/vnd.eu.kasparian.car+json",
		Description: "",
	},
	"cascii": {
		Subject:     "x-chemical/x-cactvs-binary",
		Description: "",
	},
	"cat": {
		Subject:     "application/vnd.ms-pki.seccat",
		Description: "",
	},
	"cb7": {
		Subject:     "application/x-cbr",
		Description: "",
	},
	"cba": {
		Subject:     "application/x-cbr",
		Description: "",
	},
	"cbin": {
		Subject:     "x-chemical/x-cactvs-binary",
		Description: "",
	},
	"cbl": {
		Subject:     "text/x-cobol",
		Description: "",
	},
	"cbor": {
		Subject:     "application/cbor",
		Description: "",
	},
	"cbr": {
		Subject:     "application/vnd.comicbook-rar",
		Description: "",
	},
	"cbt": {
		Subject:     "application/x-cbr",
		Description: "",
	},
	"cbz": {
		Subject:     "application/vnd.comicbook+zip",
		Description: "",
	},
	"cc": {
		Subject:     "text/x-c++src",
		Description: "",
	},
	"ccad": {
		Subject:     "application/x-clariscad",
		Description: "",
	},
	"ccc": {
		Subject:     "text/vnd.net2phone.commcenter.command",
		Description: "",
	},
	"ccmp": {
		Subject:     "application/ccmp+xml",
		Description: "",
	},
	"ccmx": {
		Subject:     "application/x-ccmx",
		Description: "",
	},
	"cco": {
		Subject:     "application/x-cocoa",
		Description: "",
	},
	"cct": {
		Subject:     "application/x-director",
		Description: "",
	},
	"ccxml": {
		Subject:     "application/ccxml+xml",
		Description: "",
	},
	"cda": {
		Subject:     "application/x-cdf",
		Description: "",
	},
	"cdbcmsg": {
		Subject:     "application/vnd.contact.cmsg",
		Description: "",
	},
	"cdf": {
		Subject:     "application/x-cdf",
		Description: "",
	},
	"cdfx": {
		Subject:     "application/cdfx+xml",
		Description: "",
	},
	"cdkey": {
		Subject:     "application/vnd.mediastation.cdkey",
		Description: "",
	},
	"cdmia": {
		Subject:     "application/cdmi-capability",
		Description: "",
	},
	"cdmic": {
		Subject:     "application/cdmi-container",
		Description: "",
	},
	"cdmid": {
		Subject:     "application/cdmi-domain",
		Description: "",
	},
	"cdmio": {
		Subject:     "application/cdmi-object",
		Description: "",
	},
	"cdmiq": {
		Subject:     "application/cdmi-queue",
		Description: "",
	},
	"cdr": {
		Subject:     "image/x-coreldraw",
		Description: "",
	},
	"cdt": {
		Subject:     "image/x-coreldrawtemplate",
		Description: "",
	},
	"cdx": {
		Subject:     "x-chemical/x-cdx",
		Description: "",
	},
	"cdxml": {
		Subject:     "application/vnd.chemdraw+xml",
		Description: "",
	},
	"cdy": {
		Subject:     "application/vnd.cinderella",
		Description: "",
	},
	"cea": {
		Subject:     "application/cea",
		Description: "",
	},
	"cef": {
		Subject:     "x-chemical/x-cxf",
		Description: "",
	},
	"cellml": {
		Subject:     "application/cellml+xml",
		Description: "",
	},
	"cer": {
		Subject:     "application/pkix-cert",
		Description: "",
	},
	"cert": {
		Subject:     "application/x-x509-ca-cert",
		Description: "",
	},
	"cfs": {
		Subject:     "application/x-cfs-compressed",
		Description: "",
	},
	"cgb": {
		Subject:     "application/x-gameboy-rom",
		Description: "",
	},
	"cgm": {
		Subject:     "image/cgm",
		Description: "",
	},
	"cha": {
		Subject:     "application/x-chat",
		Description: "",
	},
	"chat": {
		Subject:     "application/x-chat",
		Description: "",
	},
	"chm": {
		Subject:     "application/vnd.ms-htmlhelp",
		Description: "",
	},
	"chrt": {
		Subject:     "application/vnd.kde.kchart",
		Description: "",
	},
	"cif": {
		Subject:     "application/vnd.multiad.creator.cif",
		Description: "",
	},
	"cii": {
		Subject:     "application/vnd.anser-web-certificate-issue-initiation",
		Description: "",
	},
	"cil": {
		Subject:     "application/vnd.ms-artgalry",
		Description: "",
	},
	"cjs": {
		Subject:     "application/node",
		Description: "",
	},
	"cl": {
		Subject:     "application/simple-filter+xml",
		Description: "",
	},
	"cla": {
		Subject:     "application/vnd.claymore",
		Description: "",
	},
	"class": {
		Subject:     "application/vnd.dvb.dvbj",
		Description: "",
	},
	"cld": {
		Subject:     "model/vnd.cld",
		Description: "",
	},
	"clkk": {
		Subject:     "application/vnd.crick.clicker.keyboard",
		Description: "",
	},
	"clkp": {
		Subject:     "application/vnd.crick.clicker.palette",
		Description: "",
	},
	"clkt": {
		Subject:     "application/vnd.crick.clicker.template",
		Description: "",
	},
	"clkw": {
		Subject:     "application/vnd.crick.clicker.wordbank",
		Description: "",
	},
	"clkx": {
		Subject:     "application/vnd.crick.clicker",
		Description: "",
	},
	"clp": {
		Subject:     "application/x-msclip",
		Description: "",
	},
	"clpi": {
		Subject:     "video/mp2t",
		Description: "",
	},
	"cls": {
		Subject:     "text/x-tex",
		Description: "",
	},
	"clue": {
		Subject:     "application/clue_info+xml",
		Description: "",
	},
	"cmake": {
		Subject:     "text/x-cmake",
		Description: "",
	},
	"cmc": {
		Subject:     "application/vnd.cosmocaller",
		Description: "",
	},
	"cmdf": {
		Subject:     "x-chemical/x-cmdf",
		Description: "",
	},
	"cml": {
		Subject:     "x-chemical/x-cml",
		Description: "",
	},
	"cmp": {
		Subject:     "application/vnd.yellowriver-custom-menu",
		Description: "",
	},
	"cmsc": {
		Subject:     "application/cms",
		Description: "",
	},
	"cmx": {
		Subject:     "image/x-cmx",
		Description: "",
	},
	"cnd": {
		Subject:     "text/jcr-cnd",
		Description: "",
	},
	"cob": {
		Subject:     "text/x-cobol",
		Description: "",
	},
	"cod": {
		Subject:     "application/vnd.rim.cod",
		Description: "",
	},
	"coffee": {
		Subject:     "application/vnd.coffeescript",
		Description: "",
	},
	"com": {
		Subject:     "application/x-msdos-program",
		Description: "",
	},
	"conf": {
		Subject:     "text/plain",
		Description: "",
	},
	"coswid": {
		Subject:     "application/swid+cbor",
		Description: "",
	},
	"cpa": {
		Subject:     "x-chemical/x-compass",
		Description: "",
	},
	"cpi": {
		Subject:     "video/mp2t",
		Description: "",
	},
	"cpio": {
		Subject:     "application/x-cpio",
		Description: "",
	},
	"cpkg": {
		Subject:     "application/vnd.xmpie.cpkg",
		Description: "",
	},
	"cpl": {
		Subject:     "application/cpl+xml",
		Description: "",
	},
	"cpp": {
		Subject:     "text/x-c++src",
		Description: "",
	},
	"cpt": {
		Subject:     "audio/vnd.dts",
		Description: "",
	},
	"cr2": {
		Subject:     "image/x-canon-cr2",
		Description: "",
	},
	"crd": {
		Subject:     "application/x-mscardfile",
		Description: "",
	},
	"crdownload": {
		Subject:     "application/x-partial-download",
		Description: "",
	},
	"crl": {
		Subject:     "application/pkix-crl",
		Description: "",
	},
	"crt": {
		Subject:     "application/x-x509-ca-cert",
		Description: "",
	},
	"crtr": {
		Subject:     "application/vnd.multiad.creator",
		Description: "",
	},
	"crw": {
		Subject:     "image/x-canon-crw",
		Description: "",
	},
	"crx": {
		Subject:     "application/x-chrome-extension",
		Description: "",
	},
	"cryptomator": {
		Subject:     "application/vnd.cryptomator.vault",
		Description: "",
	},
	"cryptonote": {
		Subject:     "application/vnd.rig.cryptonote",
		Description: "",
	},
	"cs": {
		Subject:     "text/x-csharp",
		Description: "",
	},
	"csf": {
		Subject:     "x-chemical/x-cache-csf",
		Description: "",
	},
	"csh": {
		Subject:     "application/x-csh",
		Description: "",
	},
	"csl": {
		Subject:     "application/vnd.citationstyles.style+xml",
		Description: "",
	},
	"csm": {
		Subject:     "application/x-cu-seeme",
		Description: "",
	},
	"csml": {
		Subject:     "x-chemical/x-csml",
		Description: "",
	},
	"csp": {
		Subject:     "application/vnd.commonspace",
		Description: "",
	},
	"csrattrs": {
		Subject:     "application/csrattrs",
		Description: "",
	},
	"css": {
		Subject:     "text/css",
		Description: "",
	},
	"cst": {
		Subject:     "application/vnd.commonspace",
		Description: "",
	},
	"csv": {
		Subject:     "text/csv",
		Description: "",
	},
	"csvs": {
		Subject:     "text/csv-schema",
		Description: "",
	},
	"ctab": {
		Subject:     "x-chemical/x-cactvs-binary",
		Description: "",
	},
	"ctx": {
		Subject:     "x-chemical/x-ctx",
		Description: "",
	},
	"cu": {
		Subject:     "application/x-cu-seeme",
		Description: "",
	},
	"cub": {
		Subject:     "x-chemical/x-gaussian-cube",
		Description: "",
	},
	"cuc": {
		Subject:     "application/tamp-community-update-confirm",
		Description: "",
	},
	"cue": {
		Subject:     "application/x-cue",
		Description: "",
	},
	"cur": {
		Subject:     "image/x-win-bitmap",
		Description: "",
	},
	"curl": {
		Subject:     "application/vnd.curl",
		Description: "",
	},
	"cw": {
		Subject:     "application/prs.cww",
		Description: "",
	},
	"cwl.json": {
		Subject:     "application/cwl+json",
		Description: "",
	},
	"cwl": {
		Subject:     "application/cwl",
		Description: "",
	},
	"cww": {
		Subject:     "application/prs.cww",
		Description: "",
	},
	"cxf": {
		Subject:     "x-chemical/x-cxf",
		Description: "",
	},
	"cxt": {
		Subject:     "application/x-director",
		Description: "",
	},
	"cxx": {
		Subject:     "text/x-c++src",
		Description: "",
	},
	"d": {
		Subject:     "text/x-dsrc",
		Description: "",
	},
	"dae": {
		Subject:     "model/vnd.collada+xml",
		Description: "",
	},
	"daf": {
		Subject:     "application/vnd.mobius.daf",
		Description: "",
	},
	"dar": {
		Subject:     "application/x-dar",
		Description: "",
	},
	"dart": {
		Subject:     "application/vnd.dart",
		Description: "",
	},
	"dat": {
		Subject:     "application/octet-stream",
		Description: "",
	},
	"dataless": {
		Subject:     "application/vnd.fdsn.mseed",
		Description: "",
	},
	"davmount": {
		Subject:     "application/davmount+xml",
		Description: "",
	},
	"dbf": {
		Subject:     "application/vnd.dbf",
		Description: "",
	},
	"dbk": {
		Subject:     "application/x-docbook+xml",
		Description: "",
	},
	"dc": {
		Subject:     "application/x-dc-rom",
		Description: "",
	},
	"dcd": {
		Subject:     "application/dcd",
		Description: "",
	},
	"dcl": {
		Subject:     "text/x-dcl",
		Description: "",
	},
	"dcm": {
		Subject:     "application/dicom",
		Description: "",
	},
	"dcr": {
		Subject:     "application/x-director",
		Description: "",
	},
	"dcurl": {
		Subject:     "text/vnd.curl.dcurl",
		Description: "",
	},
	"dd2": {
		Subject:     "application/vnd.oma.dd2+xml",
		Description: "",
	},
	"ddd": {
		Subject:     "application/vnd.fujixerox.ddd",
		Description: "",
	},
	"ddf": {
		Subject:     "application/vnd.syncml.dmddf+wbxml",
		Description: "",
	},
	"dds": {
		Subject:     "image/x-dds",
		Description: "",
	},
	"deb": {
		Subject:     "application/vnd.debian.binary-package",
		Description: "",
	},
	"deepv": {
		Subject:     "application/x-deepv",
		Description: "",
	},
	"def": {
		Subject:     "text/plain",
		Description: "",
	},
	"deploy": {
		Subject:     "application/octet-stream",
		Description: "",
	},
	"der": {
		Subject:     "application/x-x509-ca-cert",
		Description: "",
	},
	"desktop": {
		Subject:     "application/x-desktop",
		Description: "",
	},
	"dfac": {
		Subject:     "application/vnd.dreamfactory",
		Description: "",
	},
	"dgc": {
		Subject:     "application/x-dgc-compressed",
		Description: "",
	},
	"dgn": {
		Subject:     "image/vnd.dgn",
		Description: "",
	},
	"di": {
		Subject:     "text/x-dsrc",
		Description: "",
	},
	"dia": {
		Subject:     "application/x-dia-diagram",
		Description: "",
	},
	"dib": {
		Subject:     "image/bmp",
		Description: "",
	},
	"dic": {
		Subject:     "text/x-c",
		Description: "",
	},
	"dif": {
		Subject:     "video/dv",
		Description: "",
	},
	"diff": {
		Subject:     "text/x-diff",
		Description: "",
	},
	"dii": {
		Subject:     "application/dii",
		Description: "",
	},
	"dim": {
		Subject:     "application/vnd.fastcopy-disk-image",
		Description: "",
	},
	"dir": {
		Subject:     "application/x-director",
		Description: "",
	},
	"dis": {
		Subject:     "application/vnd.mobius.dis",
		Description: "",
	},
	"disclosed": {
		Subject:     "application/vnd.sss-ntf",
		Description: "",
	},
	"disposition-notification": {
		Subject:     "message/disposition-notification",
		Description: "",
	},
	"dist": {
		Subject:     "application/vnd.apple.installer+xml",
		Description: "",
	},
	"distz": {
		Subject:     "application/vnd.apple.installer+xml",
		Description: "",
	},
	"dit": {
		Subject:     "application/dit",
		Description: "",
	},
	"dive": {
		Subject:     "application/vnd.patentdive",
		Description: "",
	},
	"divx": {
		Subject:     "video/x-msvideo",
		Description: "",
	},
	"djv": {
		Subject:     "image/vnd.djvu",
		Description: "",
	},
	"djvu": {
		Subject:     "image/vnd.djvu",
		Description: "",
	},
	"dl": {
		Subject:     "application/vnd.datalog",
		Description: "",
	},
	"dll": {
		Subject:     "application/vnd.microsoft.portable-executable",
		Description: "",
	},
	"dls": {
		Subject:     "audio/dls",
		Description: "",
	},
	"dmg": {
		Subject:     "application/x-apple-diskimage",
		Description: "",
	},
	"dmp": {
		Subject:     "application/vnd.tcpdump.pcap",
		Description: "",
	},
	"dms": {
		Subject:     "text/vnd.dmclientscript",
		Description: "",
	},
	"dna": {
		Subject:     "application/vnd.dna",
		Description: "",
	},
	"dng": {
		Subject:     "image/x-adobe-dng",
		Description: "",
	},
	"doc": {
		Subject:     "application/msword",
		Description: "",
	},
	"docbook": {
		Subject:     "application/x-docbook+xml",
		Description: "",
	},
	"docjson": {
		Subject:     "application/vnd.document+json",
		Description: "",
	},
	"docm": {
		Subject:     "application/vnd.ms-word.document.macroenabled.12",
		Description: "",
	},
	"docx": {
		Subject:     "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		Description: "",
	},
	"dor": {
		Subject:     "model/vnd.gs-gdl",
		Description: "",
	},
	"dot": {
		Subject:     "text/vnd.graphviz",
		Description: "",
	},
	"dotm": {
		Subject:     "application/vnd.ms-word.template.macroenabled.12",
		Description: "",
	},
	"dotx": {
		Subject:     "application/vnd.openxmlformats-officedocument.wordprocessingml.template",
		Description: "",
	},
	"dp": {
		Subject:     "application/vnd.osgi.dp",
		Description: "",
	},
	"dpg": {
		Subject:     "application/vnd.dpgraph",
		Description: "",
	},
	"dpgraph": {
		Subject:     "application/vnd.dpgraph",
		Description: "",
	},
	"dpkg": {
		Subject:     "application/vnd.xmpie.dpkg",
		Description: "",
	},
	"dpx": {
		Subject:     "image/dpx",
		Description: "",
	},
	"dra": {
		Subject:     "audio/vnd.dra",
		Description: "",
	},
	"drle": {
		Subject:     "image/dicom-rle",
		Description: "",
	},
	"drw": {
		Subject:     "application/x-drafting",
		Description: "",
	},
	"dsc": {
		Subject:     "text/prs.lines.tag",
		Description: "",
	},
	"dsl": {
		Subject:     "text/x-dsl",
		Description: "",
	},
	"dsm": {
		Subject:     "application/vnd.desmume.movie",
		Description: "",
	},
	"dssc": {
		Subject:     "application/dssc+der",
		Description: "",
	},
	"dtb": {
		Subject:     "application/x-dtbook+xml",
		Description: "",
	},
	"dtd": {
		Subject:     "application/xml-dtd",
		Description: "",
	},
	"dts": {
		Subject:     "audio/vnd.dts",
		Description: "",
	},
	"dtshd": {
		Subject:     "audio/vnd.dts.hd",
		Description: "",
	},
	"dtx": {
		Subject:     "text/x-tex",
		Description: "",
	},
	"dump": {
		Subject:     "application/octet-stream",
		Description: "",
	},
	"dv": {
		Subject:     "video/dv",
		Description: "",
	},
	"dvb": {
		Subject:     "audio/vnd.dvb.file",
		Description: "",
	},
	"dvc": {
		Subject:     "application/dvcs",
		Description: "",
	},
	"dvi": {
		Subject:     "application/x-dvi",
		Description: "",
	},
	"dwd": {
		Subject:     "application/atsc-dwd+xml",
		Description: "",
	},
	"dwf": {
		Subject:     "model/vnd.dwf",
		Description: "",
	},
	"dwg": {
		Subject:     "image/vnd.dwg",
		Description: "",
	},
	"dx": {
		Subject:     "x-chemical/x-jcamp-dx",
		Description: "",
	},
	"dxf": {
		Subject:     "image/vnd.dxf",
		Description: "",
	},
	"dxp": {
		Subject:     "application/vnd.spotfire.dxp",
		Description: "",
	},
	"dxr": {
		Subject:     "application/vnd.dxr",
		Description: "",
	},
	"dzr": {
		Subject:     "application/vnd.dzr",
		Description: "",
	},
	"e": {
		Subject:     "text/x-eiffel",
		Description: "",
	},
	"eFIF": {
		Subject:     "application/vnd.picsel",
		Description: "",
	},
	"ebuild": {
		Subject:     "application/vnd.gentoo.ebuild",
		Description: "",
	},
	"ecelp4800": {
		Subject:     "audio/vnd.nuera.ecelp4800",
		Description: "",
	},
	"ecelp7470": {
		Subject:     "audio/vnd.nuera.ecelp7470",
		Description: "",
	},
	"ecelp9600": {
		Subject:     "audio/vnd.nuera.ecelp9600",
		Description: "",
	},
	"ecig": {
		Subject:     "application/vnd.evolv.ecig.settings",
		Description: "",
	},
	"ecigprofile": {
		Subject:     "application/vnd.evolv.ecig.profile",
		Description: "",
	},
	"ecigtheme": {
		Subject:     "application/vnd.evolv.ecig.theme",
		Description: "",
	},
	"eclass": {
		Subject:     "application/vnd.gentoo.eclass",
		Description: "",
	},
	"ecma": {
		Subject:     "application/ecmascript",
		Description: "",
	},
	"edm": {
		Subject:     "application/vnd.novadigm.edm",
		Description: "",
	},
	"edx": {
		Subject:     "application/vnd.novadigm.edx",
		Description: "",
	},
	"efi": {
		Subject:     "application/efi",
		Description: "",
	},
	"efif": {
		Subject:     "application/vnd.picsel",
		Description: "",
	},
	"egon": {
		Subject:     "application/x-egon",
		Description: "",
	},
	"ei6": {
		Subject:     "application/vnd.pg.osasli",
		Description: "",
	},
	"eif": {
		Subject:     "text/x-eiffel",
		Description: "",
	},
	"el": {
		Subject:     "text/x-script.elisp",
		Description: "",
	},
	"elc": {
		Subject:     "application/x-bytecode.elisp",
		Description: "",
	},
	"eln": {
		Subject:     "application/vnd.eln+zip",
		Description: "",
	},
	"emb": {
		Subject:     "x-chemical/x-embl-dl-nucleotide",
		Description: "",
	},
	"embl": {
		Subject:     "x-chemical/x-embl-dl-nucleotide",
		Description: "",
	},
	"emf": {
		Subject:     "image/emf",
		Description: "",
	},
	"eml": {
		Subject:     "message/rfc822",
		Description: "",
	},
	"emm": {
		Subject:     "application/vnd.ibm.electronic-media",
		Description: "",
	},
	"emma": {
		Subject:     "application/emma+xml",
		Description: "",
	},
	"emotionml": {
		Subject:     "application/emotionml+xml",
		Description: "",
	},
	"emp": {
		Subject:     "application/vnd.emusic-emusic_package",
		Description: "",
	},
	"emz": {
		Subject:     "application/x-msmetafile",
		Description: "",
	},
	"ent": {
		Subject:     "application/vnd.nervana",
		Description: "",
	},
	"entity": {
		Subject:     "application/vnd.nervana",
		Description: "",
	},
	"env": {
		Subject:     "application/x-envoy",
		Description: "",
	},
	"enw": {
		Subject:     "audio/evrcnw",
		Description: "",
	},
	"eol": {
		Subject:     "audio/vnd.digital-winds",
		Description: "",
	},
	"eot": {
		Subject:     "application/vnd.ms-fontobject",
		Description: "",
	},
	"eps2": {
		Subject:     "application/postscript",
		Description: "",
	},
	"eps3": {
		Subject:     "application/postscript",
		Description: "",
	},
	"eps": {
		Subject:     "application/postscript",
		Description: "",
	},
	"epsf": {
		Subject:     "application/postscript",
		Description: "",
	},
	"epsi": {
		Subject:     "application/postscript",
		Description: "",
	},
	"epub": {
		Subject:     "application/epub+zip",
		Description: "",
	},
	"erf": {
		Subject:     "image/x-epson-erf",
		Description: "",
	},
	"erl": {
		Subject:     "text/x-erlang",
		Description: "",
	},
	"erofs": {
		Subject:     "application/vnd.erofs",
		Description: "",
	},
	"es3": {
		Subject:     "application/vnd.eszigno3+xml",
		Description: "",
	},
	"es": {
		Subject:     "application/ecmascript",
		Description: "",
	},
	"esa": {
		Subject:     "application/vnd.osgi.subsystem",
		Description: "",
	},
	"esf": {
		Subject:     "application/vnd.epson.esf",
		Description: "",
	},
	"espass": {
		Subject:     "application/vnd.espass-espass+zip",
		Description: "",
	},
	"et3": {
		Subject:     "application/vnd.eszigno3+xml",
		Description: "",
	},
	"etheme": {
		Subject:     "application/x-e-theme",
		Description: "",
	},
	"etx": {
		Subject:     "text/x-setext",
		Description: "",
	},
	"eva": {
		Subject:     "application/x-eva",
		Description: "",
	},
	"evb": {
		Subject:     "audio/evrcb",
		Description: "",
	},
	"evc": {
		Subject:     "audio/evrc",
		Description: "",
	},
	"event-stream": {
		Subject:     "text/x-event-stream",
		Description: "",
	},
	"evw": {
		Subject:     "audio/evrcwb",
		Description: "",
	},
	"evy": {
		Subject:     "application/x-envoy",
		Description: "",
	},
	"exe": {
		Subject:     "application/vnd.microsoft.portable-executable",
		Description: "",
	},
	"exi": {
		Subject:     "application/exi",
		Description: "",
	},
	"exp": {
		Subject:     "application/express",
		Description: "",
	},
	"exr": {
		Subject:     "image/aces",
		Description: "",
	},
	"ext": {
		Subject:     "application/vnd.novadigm.ext",
		Description: "",
	},
	"ez2": {
		Subject:     "application/vnd.ezpix-album",
		Description: "",
	},
	"ez3": {
		Subject:     "application/vnd.ezpix-package",
		Description: "",
	},
	"ez": {
		Subject:     "application/andrew-inset",
		Description: "",
	},
	"f4a": {
		Subject:     "audio/mp4",
		Description: "",
	},
	"f4b": {
		Subject:     "audio/mp4",
		Description: "",
	},
	"f4p": {
		Subject:     "video/mp4",
		Description: "",
	},
	"f4v": {
		Subject:     "video/mp4",
		Description: "",
	},
	"f77": {
		Subject:     "text/x-fortran",
		Description: "",
	},
	"f90": {
		Subject:     "text/x-fortran",
		Description: "",
	},
	"f95": {
		Subject:     "text/x-fortran",
		Description: "",
	},
	"f": {
		Subject:     "text/x-fortran",
		Description: "",
	},
	"fb2": {
		Subject:     "application/x-fictionbook+xml",
		Description: "",
	},
	"fb": {
		Subject:     "application/x-maker",
		Description: "",
	},
	"fbdoc": {
		Subject:     "application/x-maker",
		Description: "",
	},
	"fbs": {
		Subject:     "image/vnd.fastbidsheet",
		Description: "",
	},
	"fcdt": {
		Subject:     "application/vnd.adobe.formscentral.fcdt",
		Description: "",
	},
	"fch": {
		Subject:     "x-chemical/x-gaussian-checkpoint",
		Description: "",
	},
	"fchk": {
		Subject:     "x-chemical/x-gaussian-checkpoint",
		Description: "",
	},
	"fcs": {
		Subject:     "application/vnd.isac.fcs",
		Description: "",
	},
	"fdf": {
		Subject:     "application/fdf",
		Description: "",
	},
	"fdt": {
		Subject:     "application/fdt+xml",
		Description: "",
	},
	"fe_launch": {
		Subject:     "application/vnd.denovo.fcselayout-link",
		Description: "",
	},
	"fg5": {
		Subject:     "application/vnd.fujitsu.oasysgp",
		Description: "",
	},
	"fgd": {
		Subject:     "application/x-director",
		Description: "",
	},
	"fh4": {
		Subject:     "image/x-freehand",
		Description: "",
	},
	"fh5": {
		Subject:     "image/x-freehand",
		Description: "",
	},
	"fh7": {
		Subject:     "image/x-freehand",
		Description: "",
	},
	"fh": {
		Subject:     "image/x-freehand",
		Description: "",
	},
	"fhc": {
		Subject:     "image/x-freehand",
		Description: "",
	},
	"fif": {
		Subject:     "application/x-fractals",
		Description: "",
	},
	"fig": {
		Subject:     "application/x-xfig",
		Description: "",
	},
	"fits": {
		Subject:     "application/fits",
		Description: "",
	},
	"fl": {
		Subject:     "application/x-fluid",
		Description: "",
	},
	"fla": {
		Subject:     "application/vnd.dtg.local.flash",
		Description: "",
	},
	"flac": {
		Subject:     "audio/x-flac",
		Description: "",
	},
	"flb": {
		Subject:     "application/vnd.ficlab.flb+zip",
		Description: "",
	},
	"flc": {
		Subject:     "video/x-flic",
		Description: "",
	},
	"fli": {
		Subject:     "video/x-fli",
		Description: "",
	},
	"flo": {
		Subject:     "application/vnd.micrografx.flo",
		Description: "",
	},
	"flt": {
		Subject:     "text/vnd.ficlab.flt",
		Description: "",
	},
	"flv": {
		Subject:     "video/x-flv",
		Description: "",
	},
	"flw": {
		Subject:     "application/vnd.kde.kivio",
		Description: "",
	},
	"flx": {
		Subject:     "text/vnd.fmi.flexstor",
		Description: "",
	},
	"fly": {
		Subject:     "text/vnd.fly",
		Description: "",
	},
	"fm": {
		Subject:     "application/x-maker",
		Description: "",
	},
	"fmf": {
		Subject:     "video/x-atomic3d-feature",
		Description: "",
	},
	"fnc": {
		Subject:     "application/vnd.frogans.fnc",
		Description: "",
	},
	"fo": {
		Subject:     "text/x-xslfo",
		Description: "",
	},
	"fodg": {
		Subject:     "application/vnd.oasis.opendocument.graphics-flat-xml",
		Description: "",
	},
	"fodp": {
		Subject:     "application/vnd.oasis.opendocument.presentation-flat-xml",
		Description: "",
	},
	"fods": {
		Subject:     "application/vnd.oasis.opendocument.spreadsheet-flat-xml",
		Description: "",
	},
	"fodt": {
		Subject:     "application/vnd.oasis.opendocument.text-flat-xml",
		Description: "",
	},
	"for": {
		Subject:     "text/x-fortran",
		Description: "",
	},
	"fpx": {
		Subject:     "application/vnd.netfpx",
		Description: "",
	},
	"frame": {
		Subject:     "application/x-maker",
		Description: "",
	},
	"frl": {
		Subject:     "application/x-freeloader",
		Description: "",
	},
	"frm": {
		Subject:     "application/vnd.ufdl",
		Description: "",
	},
	"fsc": {
		Subject:     "application/vnd.fsc.weblaunch",
		Description: "",
	},
	"fst": {
		Subject:     "image/vnd.fst",
		Description: "",
	},
	"ftc": {
		Subject:     "application/vnd.fluxtime.clip",
		Description: "",
	},
	"fti": {
		Subject:     "application/vnd.anser-web-funds-transfer-initiation",
		Description: "",
	},
	"funk": {
		Subject:     "audio/x-make",
		Description: "",
	},
	"fvt": {
		Subject:     "video/vnd.fvt",
		Description: "",
	},
	"fxm": {
		Subject:     "video/x-javafx",
		Description: "",
	},
	"fxp": {
		Subject:     "application/vnd.adobe.fxp",
		Description: "",
	},
	"fxpl": {
		Subject:     "application/vnd.adobe.fxp",
		Description: "",
	},
	"fzs": {
		Subject:     "application/vnd.fuzzysheet",
		Description: "",
	},
	"g2w": {
		Subject:     "application/vnd.geoplan",
		Description: "",
	},
	"g3": {
		Subject:     "application/vnd.geocube+xml",
		Description: "",
	},
	"g3w": {
		Subject:     "application/vnd.geospace",
		Description: "",
	},
	"g": {
		Subject:     "text/plain",
		Description: "",
	},
	"gac": {
		Subject:     "application/vnd.groove-account",
		Description: "",
	},
	"gal": {
		Subject:     "x-chemical/x-gaussian-log",
		Description: "",
	},
	"gam": {
		Subject:     "x-chemical/x-gamess-input",
		Description: "",
	},
	"gamin": {
		Subject:     "x-chemical/x-gamess-input",
		Description: "",
	},
	"gau": {
		Subject:     "x-chemical/x-gaussian-input",
		Description: "",
	},
	"gb": {
		Subject:     "application/x-gameboy-rom",
		Description: "",
	},
	"gba": {
		Subject:     "application/x-gba-rom",
		Description: "",
	},
	"gbc": {
		Subject:     "application/x-gameboy-rom",
		Description: "",
	},
	"gbr": {
		Subject:     "application/rpki-ghostbusters",
		Description: "",
	},
	"gca": {
		Subject:     "application/x-gca-compressed",
		Description: "",
	},
	"gcd": {
		Subject:     "text/x-pcs-gcd",
		Description: "",
	},
	"gcf": {
		Subject:     "application/x-graphing-calculator",
		Description: "",
	},
	"gcg": {
		Subject:     "x-chemical/x-gcg8-sequence",
		Description: "",
	},
	"gcrd": {
		Subject:     "text/vcard",
		Description: "",
	},
	"gdl": {
		Subject:     "model/vnd.gs-gdl",
		Description: "",
	},
	"gdz": {
		Subject:     "application/vnd.familysearch.gedcom+zip",
		Description: "",
	},
	"ged": {
		Subject:     "text/vnd.familysearch.gedcom",
		Description: "",
	},
	"gedcom": {
		Subject:     "application/x-gedcom",
		Description: "",
	},
	"gem": {
		Subject:     "application/x-tar",
		Description: "",
	},
	"gen": {
		Subject:     "x-chemical/x-genbank",
		Description: "",
	},
	"genozip": {
		Subject:     "application/vnd.genozip",
		Description: "",
	},
	"geo": {
		Subject:     "application/vnd.dynageo",
		Description: "",
	},
	"geojson": {
		Subject:     "application/geo+json",
		Description: "",
	},
	"gex": {
		Subject:     "application/vnd.geometry-explorer",
		Description: "",
	},
	"gf": {
		Subject:     "application/x-tex-gf",
		Description: "",
	},
	"gff3": {
		Subject:     "text/gff3",
		Description: "",
	},
	"gg": {
		Subject:     "application/x-sms-rom",
		Description: "",
	},
	"ggb": {
		Subject:     "application/vnd.geogebra.file",
		Description: "",
	},
	"ggs": {
		Subject:     "application/vnd.geogebra.slides",
		Description: "",
	},
	"ggt": {
		Subject:     "application/vnd.geogebra.tool",
		Description: "",
	},
	"ghf": {
		Subject:     "application/vnd.groove-help",
		Description: "",
	},
	"gif": {
		Subject:     "image/gif",
		Description: "",
	},
	"gim": {
		Subject:     "application/vnd.groove-identity-message",
		Description: "",
	},
	"gjc": {
		Subject:     "x-chemical/x-gaussian-input",
		Description: "",
	},
	"gjf": {
		Subject:     "x-chemical/x-gaussian-input",
		Description: "",
	},
	"gl": {
		Subject:     "video/x-gl",
		Description: "",
	},
	"glade": {
		Subject:     "application/x-glade",
		Description: "",
	},
	"glb": {
		Subject:     "model/gltf-binary",
		Description: "",
	},
	"glbin": {
		Subject:     "application/gltf-buffer",
		Description: "",
	},
	"glbuf": {
		Subject:     "application/gltf-buffer",
		Description: "",
	},
	"gltf": {
		Subject:     "model/gltf+json",
		Description: "",
	},
	"gml": {
		Subject:     "application/gml+xml",
		Description: "",
	},
	"gmo": {
		Subject:     "application/x-gettext-translation",
		Description: "",
	},
	"gmx": {
		Subject:     "application/vnd.gmx",
		Description: "",
	},
	"gnc": {
		Subject:     "application/x-gnucash",
		Description: "",
	},
	"gnd": {
		Subject:     "application/x-gnunet-directory",
		Description: "",
	},
	"gnucash": {
		Subject:     "application/x-gnucash",
		Description: "",
	},
	"gnumeric": {
		Subject:     "application/x-gnumeric",
		Description: "",
	},
	"gnuplot": {
		Subject:     "application/x-gnuplot",
		Description: "",
	},
	"go": {
		Subject:     "text/x-go",
		Description: "",
	},
	"gp": {
		Subject:     "application/x-gnuplot",
		Description: "",
	},
	"gpg": {
		Subject:     "application/pgp-encrypted",
		Description: "",
	},
	"gph": {
		Subject:     "application/vnd.flographit",
		Description: "",
	},
	"gpkg.tar": {
		Subject:     "application/vnd.gentoo.gpkg",
		Description: "",
	},
	"gpkg": {
		Subject:     "application/geopackage+sqlite3",
		Description: "",
	},
	"gplt": {
		Subject:     "application/x-gnuplot",
		Description: "",
	},
	"gpt": {
		Subject:     "x-chemical/x-mopac-graph",
		Description: "",
	},
	"gpx": {
		Subject:     "application/x-gpx+xml",
		Description: "",
	},
	"gqf": {
		Subject:     "application/vnd.grafeq",
		Description: "",
	},
	"gqs": {
		Subject:     "application/vnd.grafeq",
		Description: "",
	},
	"gra": {
		Subject:     "application/x-graphite",
		Description: "",
	},
	"gram": {
		Subject:     "application/srgs",
		Description: "",
	},
	"gramps": {
		Subject:     "application/x-gramps-xml",
		Description: "",
	},
	"grd": {
		Subject:     "application/vnd.gentics.grd+json",
		Description: "",
	},
	"gre": {
		Subject:     "application/vnd.geometry-explorer",
		Description: "",
	},
	"grv": {
		Subject:     "application/vnd.groove-injector",
		Description: "",
	},
	"grxml": {
		Subject:     "application/srgs+xml",
		Description: "",
	},
	"gs": {
		Subject:     "text/x-genie",
		Description: "",
	},
	"gsd": {
		Subject:     "audio/x-gsm",
		Description: "",
	},
	"gsf": {
		Subject:     "application/x-font",
		Description: "",
	},
	"gsheet": {
		Subject:     "application/urc-grpsheet+xml",
		Description: "",
	},
	"gsm": {
		Subject:     "model/vnd.gs-gdl",
		Description: "",
	},
	"gsp": {
		Subject:     "application/x-gsp",
		Description: "",
	},
	"gss": {
		Subject:     "application/x-gss",
		Description: "",
	},
	"gtar": {
		Subject:     "application/x-gtar",
		Description: "",
	},
	"gtm": {
		Subject:     "application/vnd.groove-tool-message",
		Description: "",
	},
	"gtw": {
		Subject:     "model/vnd.gtw",
		Description: "",
	},
	"gv": {
		Subject:     "text/vnd.graphviz",
		Description: "",
	},
	"gvp": {
		Subject:     "text/x-google-video-pointer",
		Description: "",
	},
	"gxf": {
		Subject:     "application/x-gxf",
		Description: "",
	},
	"gxt": {
		Subject:     "application/vnd.geonext",
		Description: "",
	},
	"gz": {
		Subject:     "application/gzip",
		Description: "",
	},
	"gzip": {
		Subject:     "application/x-gzip",
		Description: "",
	},
	"h++": {
		Subject:     "text/x-c++hdr",
		Description: "",
	},
	"h261": {
		Subject:     "video/h261",
		Description: "",
	},
	"h263": {
		Subject:     "video/h263",
		Description: "",
	},
	"h264": {
		Subject:     "video/h264",
		Description: "",
	},
	"h4": {
		Subject:     "application/x-hdf",
		Description: "",
	},
	"h5": {
		Subject:     "application/x-hdf",
		Description: "",
	},
	"h": {
		Subject:     "text/x-chdr",
		Description: "",
	},
	"hal": {
		Subject:     "application/vnd.hal+xml",
		Description: "",
	},
	"hans": {
		Subject:     "text/vnd.hans",
		Description: "",
	},
	"hbc": {
		Subject:     "application/vnd.fints",
		Description: "",
	},
	"hbci": {
		Subject:     "application/vnd.fints",
		Description: "",
	},
	"hdf4": {
		Subject:     "application/x-hdf",
		Description: "",
	},
	"hdf5": {
		Subject:     "application/x-hdf",
		Description: "",
	},
	"hdf": {
		Subject:     "application/x-hdf",
		Description: "",
	},
	"hdr": {
		Subject:     "image/vnd.radiance",
		Description: "",
	},
	"hdt": {
		Subject:     "application/vnd.hdt",
		Description: "",
	},
	"heic": {
		Subject:     "image/heic",
		Description: "",
	},
	"heics": {
		Subject:     "image/heic-sequence",
		Description: "",
	},
	"heif": {
		Subject:     "image/heif",
		Description: "",
	},
	"heifs": {
		Subject:     "image/avif",
		Description: "",
	},
	"hej2": {
		Subject:     "image/hej2k",
		Description: "",
	},
	"held": {
		Subject:     "application/atsc-held+xml",
		Description: "",
	},
	"heldxml": {
		Subject:     "application/held+xml",
		Description: "",
	},
	"help": {
		Subject:     "application/x-helpfile",
		Description: "",
	},
	"hep": {
		Subject:     "application/x-hep",
		Description: "",
	},
	"hgl": {
		Subject:     "text/vnd.hgl",
		Description: "",
	},
	"hh": {
		Subject:     "text/x-c++hdr",
		Description: "",
	},
	"hif": {
		Subject:     "image/avif",
		Description: "",
	},
	"hin": {
		Subject:     "x-chemical/x-hin",
		Description: "",
	},
	"hlb": {
		Subject:     "text/x-script",
		Description: "",
	},
	"hlp": {
		Subject:     "application/x-winhlp",
		Description: "",
	},
	"hp": {
		Subject:     "text/x-c++hdr",
		Description: "",
	},
	"hpg": {
		Subject:     "application/vnd.hp-hpgl",
		Description: "",
	},
	"hpgl": {
		Subject:     "application/vnd.hp-hpgl",
		Description: "",
	},
	"hpi": {
		Subject:     "application/vnd.hp-hpid",
		Description: "",
	},
	"hpid": {
		Subject:     "application/vnd.hp-hpid",
		Description: "",
	},
	"hpp": {
		Subject:     "text/x-c++hdr",
		Description: "",
	},
	"hps": {
		Subject:     "application/vnd.hp-hps",
		Description: "",
	},
	"hqx": {
		Subject:     "application/mac-binhex40",
		Description: "",
	},
	"hs": {
		Subject:     "text/x-haskell",
		Description: "",
	},
	"hsj2": {
		Subject:     "image/hsj2",
		Description: "",
	},
	"hsl": {
		Subject:     "application/vnd.hsl",
		Description: "",
	},
	"hta": {
		Subject:     "application/x-hta",
		Description: "",
	},
	"htc": {
		Subject:     "text/x-component",
		Description: "",
	},
	"htke": {
		Subject:     "application/vnd.kenameaapp",
		Description: "",
	},
	"htm": {
		Subject:     "text/html",
		Description: "",
	},
	"html": {
		Subject:     "text/html",
		Description: "",
	},
	"htmls": {
		Subject:     "text/html",
		Description: "",
	},
	"htmlx": {
		Subject:     "text/html",
		Description: "",
	},
	"htt": {
		Subject:     "text/x-webviewhtml",
		Description: "",
	},
	"htx": {
		Subject:     "text/html",
		Description: "",
	},
	"hvd": {
		Subject:     "application/vnd.yamaha.hv-dic",
		Description: "",
	},
	"hvp": {
		Subject:     "application/vnd.yamaha.hv-voice",
		Description: "",
	},
	"hvs": {
		Subject:     "application/vnd.yamaha.hv-script",
		Description: "",
	},
	"hwp": {
		Subject:     "application/x-hwp",
		Description: "",
	},
	"hwt": {
		Subject:     "application/x-hwt",
		Description: "",
	},
	"hxx": {
		Subject:     "text/x-c++hdr",
		Description: "",
	},
	"i2g": {
		Subject:     "application/vnd.intergeo",
		Description: "",
	},
	"ic1": {
		Subject:     "application/vnd.commerce-battelle",
		Description: "",
	},
	"ica": {
		Subject:     "application/vnd.commerce-battelle",
		Description: "",
	},
	"icb": {
		Subject:     "image/x-tga",
		Description: "",
	},
	"icc": {
		Subject:     "application/vnd.commerce-battelle",
		Description: "",
	},
	"icd": {
		Subject:     "application/vnd.commerce-battelle",
		Description: "",
	},
	"ice": {
		Subject:     "x-conference/x-cooltalk",
		Description: "",
	},
	"icf": {
		Subject:     "application/vnd.commerce-battelle",
		Description: "",
	},
	"icm": {
		Subject:     "application/vnd.iccprofile",
		Description: "",
	},
	"icns": {
		Subject:     "image/x-icns",
		Description: "",
	},
	"ico": {
		Subject:     "image/vnd.microsoft.icon",
		Description: "",
	},
	"ics": {
		Subject:     "text/calendar",
		Description: "",
	},
	"icz": {
		Subject:     "text/calendar",
		Description: "",
	},
	"idc": {
		Subject:     "text/plain",
		Description: "",
	},
	"idl": {
		Subject:     "text/x-idl",
		Description: "",
	},
	"ief": {
		Subject:     "image/ief",
		Description: "",
	},
	"iefs": {
		Subject:     "image/ief",
		Description: "",
	},
	"iepd": {
		Subject:     "application/vnd.nato.openxmlformats-package.iepd+zip",
		Description: "",
	},
	"ifb": {
		Subject:     "text/calendar",
		Description: "",
	},
	"ifc": {
		Subject:     "application/p21",
		Description: "",
	},
	"iff": {
		Subject:     "image/x-ilbm",
		Description: "",
	},
	"ifm": {
		Subject:     "application/vnd.shana.informed.formdata",
		Description: "",
	},
	"iges": {
		Subject:     "model/iges",
		Description: "",
	},
	"igl": {
		Subject:     "application/vnd.igloader",
		Description: "",
	},
	"igm": {
		Subject:     "application/vnd.insors.igm",
		Description: "",
	},
	"ign": {
		Subject:     "application/vnd.coreos.ignition+json",
		Description: "",
	},
	"ignition": {
		Subject:     "application/vnd.coreos.ignition+json",
		Description: "",
	},
	"igs": {
		Subject:     "model/iges",
		Description: "",
	},
	"igx": {
		Subject:     "application/vnd.micrografx.igx",
		Description: "",
	},
	"iif": {
		Subject:     "application/vnd.shana.informed.interchange",
		Description: "",
	},
	"iii": {
		Subject:     "application/x-iphone",
		Description: "",
	},
	"ilbm": {
		Subject:     "image/x-ilbm",
		Description: "",
	},
	"ima": {
		Subject:     "application/x-ima",
		Description: "",
	},
	"imagemap": {
		Subject:     "application/x-imagemap",
		Description: "",
	},
	"imap": {
		Subject:     "application/x-imagemap",
		Description: "",
	},
	"ime": {
		Subject:     "text/x-imelody",
		Description: "",
	},
	"imf": {
		Subject:     "application/vnd.imagemeter.folder+zip",
		Description: "",
	},
	"img": {
		Subject:     "application/vnd.efi.img",
		Description: "",
	},
	"imgcal": {
		Subject:     "application/vnd.3lightssoftware.imagescal",
		Description: "",
	},
	"imi": {
		Subject:     "application/vnd.imagemeter.image+zip",
		Description: "",
	},
	"imp": {
		Subject:     "application/vnd.accpac.simply.imp",
		Description: "",
	},
	"ims": {
		Subject:     "application/vnd.ms-ims",
		Description: "",
	},
	"imscc": {
		Subject:     "application/vnd.ims.imsccv1p1",
		Description: "",
	},
	"imy": {
		Subject:     "text/x-imelody",
		Description: "",
	},
	"in": {
		Subject:     "text/plain",
		Description: "",
	},
	"inf": {
		Subject:     "application/x-inf",
		Description: "",
	},
	"info": {
		Subject:     "application/x-info",
		Description: "",
	},
	"ini": {
		Subject:     "text/plain",
		Description: "",
	},
	"ink": {
		Subject:     "application/inkml+xml",
		Description: "",
	},
	"inkml+xml": {
		Subject:     "application/inkml+xml",
		Description: "",
	},
	"inkml": {
		Subject:     "application/inkml+xml",
		Description: "",
	},
	"inp": {
		Subject:     "x-chemical/x-gamess-input",
		Description: "",
	},
	"ins": {
		Subject:     "application/x-internet-signup",
		Description: "",
	},
	"install": {
		Subject:     "application/x-install-instructions",
		Description: "",
	},
	"iota": {
		Subject:     "application/vnd.astraea-software.iota",
		Description: "",
	},
	"ip": {
		Subject:     "application/x-ip2",
		Description: "",
	},
	"ipfix": {
		Subject:     "application/ipfix",
		Description: "",
	},
	"ipk": {
		Subject:     "application/vnd.shana.informed.package",
		Description: "",
	},
	"ipns-record": {
		Subject:     "application/vnd.ipfs.ipns-record",
		Description: "",
	},
	"iptables": {
		Subject:     "text/x-iptables",
		Description: "",
	},
	"irm": {
		Subject:     "application/vnd.ibm.rights-management",
		Description: "",
	},
	"irp": {
		Subject:     "application/vnd.irepository.package+xml",
		Description: "",
	},
	"ism": {
		Subject:     "model/vnd.gs-gdl",
		Description: "",
	},
	"iso9660": {
		Subject:     "application/x-cd-image",
		Description: "",
	},
	"iso": {
		Subject:     "application/vnd.efi.iso",
		Description: "",
	},
	"isp": {
		Subject:     "application/x-internet-signup",
		Description: "",
	},
	"ist": {
		Subject:     "x-chemical/x-isostar",
		Description: "",
	},
	"istc": {
		Subject:     "application/vnd.veryant.thin",
		Description: "",
	},
	"istr": {
		Subject:     "x-chemical/x-isostar",
		Description: "",
	},
	"isu": {
		Subject:     "video/x-isvideo",
		Description: "",
	},
	"isws": {
		Subject:     "application/vnd.veryant.thin",
		Description: "",
	},
	"it87": {
		Subject:     "application/x-it87",
		Description: "",
	},
	"it": {
		Subject:     "audio/x-it",
		Description: "",
	},
	"itp": {
		Subject:     "application/vnd.shana.informed.formtemplate",
		Description: "",
	},
	"its": {
		Subject:     "application/its+xml",
		Description: "",
	},
	"iv": {
		Subject:     "application/x-inventor",
		Description: "",
	},
	"ivp": {
		Subject:     "application/vnd.immervision-ivp",
		Description: "",
	},
	"ivr": {
		Subject:     "x-i-world/x-i-vrml",
		Description: "",
	},
	"ivu": {
		Subject:     "application/vnd.immervision-ivu",
		Description: "",
	},
	"ivy": {
		Subject:     "application/x-livescreen",
		Description: "",
	},
	"j2c": {
		Subject:     "image/j2c",
		Description: "",
	},
	"j2k": {
		Subject:     "image/j2c",
		Description: "",
	},
	"jad": {
		Subject:     "text/vnd.sun.j2me.app-descriptor",
		Description: "",
	},
	"jam": {
		Subject:     "application/vnd.jam",
		Description: "",
	},
	"jar": {
		Subject:     "application/java-archive",
		Description: "",
	},
	"jav": {
		Subject:     "text/x-java-source",
		Description: "",
	},
	"java": {
		Subject:     "text/x-java",
		Description: "",
	},
	"jceks": {
		Subject:     "application/x-java-jce-keystore",
		Description: "",
	},
	"jcm": {
		Subject:     "application/x-java-commerce",
		Description: "",
	},
	"jdx": {
		Subject:     "x-chemical/x-jcamp-dx",
		Description: "",
	},
	"jfif-tbnl": {
		Subject:     "image/jpeg",
		Description: "",
	},
	"jfif": {
		Subject:     "image/jpeg",
		Description: "",
	},
	"jhc": {
		Subject:     "image/jphc",
		Description: "",
	},
	"jisp": {
		Subject:     "application/vnd.jisp",
		Description: "",
	},
	"jks": {
		Subject:     "application/x-java-keystore",
		Description: "",
	},
	"jls": {
		Subject:     "image/jls",
		Description: "",
	},
	"jlt": {
		Subject:     "application/vnd.hp-jlyt",
		Description: "",
	},
	"jmz": {
		Subject:     "application/x-jmol",
		Description: "",
	},
	"jng": {
		Subject:     "image/x-jng",
		Description: "",
	},
	"jnlp": {
		Subject:     "application/x-java-jnlp-file",
		Description: "",
	},
	"joda": {
		Subject:     "application/vnd.joost.joda-archive",
		Description: "",
	},
	"jp2": {
		Subject:     "image/jp2",
		Description: "",
	},
	"jpe": {
		Subject:     "image/jpeg",
		Description: "",
	},
	"jpeg": {
		Subject:     "image/jpeg",
		Description: "",
	},
	"jpf": {
		Subject:     "image/jpx",
		Description: "",
	},
	"jpg2": {
		Subject:     "image/jp2",
		Description: "",
	},
	"jpg": {
		Subject:     "image/jpeg",
		Description: "",
	},
	"jpgm": {
		Subject:     "image/jpm",
		Description: "",
	},
	"jpgv": {
		Subject:     "video/jpeg",
		Description: "",
	},
	"jph": {
		Subject:     "image/jph",
		Description: "",
	},
	"jpm": {
		Subject:     "image/jpm",
		Description: "",
	},
	"jpr": {
		Subject:     "application/x-jbuilder-project",
		Description: "",
	},
	"jps": {
		Subject:     "image/x-jps",
		Description: "",
	},
	"jpx": {
		Subject:     "image/jpx",
		Description: "",
	},
	"jrd": {
		Subject:     "application/jrd+json",
		Description: "",
	},
	"js": {
		Subject:     "application/javascript",
		Description: "",
	},
	"jsm": {
		Subject:     "application/javascript",
		Description: "",
	},
	"json-patch": {
		Subject:     "application/json-patch+json",
		Description: "",
	},
	"json": {
		Subject:     "application/json",
		Description: "",
	},
	"jsonld": {
		Subject:     "application/ld+json",
		Description: "",
	},
	"jsonml": {
		Subject:     "application/x-jsonml+json",
		Description: "",
	},
	"jsontd": {
		Subject:     "application/td+json",
		Description: "",
	},
	"jsontm": {
		Subject:     "application/tm+json",
		Description: "",
	},
	"jt": {
		Subject:     "model/jt",
		Description: "",
	},
	"jtd": {
		Subject:     "text/vnd.esmertec.theme-descriptor",
		Description: "",
	},
	"jut": {
		Subject:     "image/x-jutvision",
		Description: "",
	},
	"jxr": {
		Subject:     "image/jxr",
		Description: "",
	},
	"jxra": {
		Subject:     "image/jxra",
		Description: "",
	},
	"jxrs": {
		Subject:     "image/jxrs",
		Description: "",
	},
	"jxs": {
		Subject:     "image/jxs",
		Description: "",
	},
	"jxsc": {
		Subject:     "image/jxsc",
		Description: "",
	},
	"jxsi": {
		Subject:     "image/jxsi",
		Description: "",
	},
	"jxss": {
		Subject:     "image/jxss",
		Description: "",
	},
	"k25": {
		Subject:     "image/x-kodak-k25",
		Description: "",
	},
	"kar": {
		Subject:     "audio/x-midi",
		Description: "",
	},
	"karbon": {
		Subject:     "application/vnd.kde.karbon",
		Description: "",
	},
	"kcm": {
		Subject:     "application/vnd.nervana",
		Description: "",
	},
	"kdc": {
		Subject:     "image/x-kodak-kdc",
		Description: "",
	},
	"kdelnk": {
		Subject:     "application/x-desktop",
		Description: "",
	},
	"kexi": {
		Subject:     "application/x-kexiproject-sqlite2",
		Description: "",
	},
	"kexic": {
		Subject:     "application/x-kexi-connectiondata",
		Description: "",
	},
	"kexis": {
		Subject:     "application/x-kexiproject-shortcut",
		Description: "",
	},
	"key": {
		Subject:     "application/vnd.apple.keynote",
		Description: "",
	},
	"kfo": {
		Subject:     "application/vnd.kde.kformula",
		Description: "",
	},
	"kia": {
		Subject:     "application/vnd.kidspiration",
		Description: "",
	},
	"kil": {
		Subject:     "application/x-killustrator",
		Description: "",
	},
	"kin": {
		Subject:     "x-chemical/x-kinemage",
		Description: "",
	},
	"kino": {
		Subject:     "application/smil+xml",
		Description: "",
	},
	"kml": {
		Subject:     "application/vnd.google-earth.kml+xml",
		Description: "",
	},
	"kmz": {
		Subject:     "application/vnd.google-earth.kmz",
		Description: "",
	},
	"kne": {
		Subject:     "application/vnd.kinar",
		Description: "",
	},
	"knp": {
		Subject:     "application/vnd.kinar",
		Description: "",
	},
	"kom": {
		Subject:     "application/vnd.fints",
		Description: "",
	},
	"kon": {
		Subject:     "application/vnd.kde.kontour",
		Description: "",
	},
	"koz": {
		Subject:     "audio/vnd.audiokoz",
		Description: "",
	},
	"kpm": {
		Subject:     "application/x-kpovmodeler",
		Description: "",
	},
	"kpr": {
		Subject:     "application/vnd.kde.kpresenter",
		Description: "",
	},
	"kpt": {
		Subject:     "application/vnd.kde.kpresenter",
		Description: "",
	},
	"kpxx": {
		Subject:     "application/vnd.ds-keypoint",
		Description: "",
	},
	"kra": {
		Subject:     "application/x-krita",
		Description: "",
	},
	"ks": {
		Subject:     "application/x-java-keystore",
		Description: "",
	},
	"ksh": {
		Subject:     "application/x-ksh",
		Description: "",
	},
	"ksp": {
		Subject:     "application/vnd.kde.kspread",
		Description: "",
	},
	"ktr": {
		Subject:     "application/vnd.kahootz",
		Description: "",
	},
	"ktx2": {
		Subject:     "image/ktx2",
		Description: "",
	},
	"ktx": {
		Subject:     "image/ktx",
		Description: "",
	},
	"ktz": {
		Subject:     "application/vnd.kahootz",
		Description: "",
	},
	"kud": {
		Subject:     "application/x-kugar",
		Description: "",
	},
	"kwd": {
		Subject:     "application/vnd.kde.kword",
		Description: "",
	},
	"kwt": {
		Subject:     "application/vnd.kde.kword",
		Description: "",
	},
	"l16": {
		Subject:     "audio/l16",
		Description: "",
	},
	"la": {
		Subject:     "audio/x-nspaudio",
		Description: "",
	},
	"lam": {
		Subject:     "audio/x-liveaudio",
		Description: "",
	},
	"lasjson": {
		Subject:     "application/vnd.las.las+json",
		Description: "",
	},
	"lasxml": {
		Subject:     "application/vnd.las.las+xml",
		Description: "",
	},
	"latex": {
		Subject:     "application/x-latex",
		Description: "",
	},
	"lbc": {
		Subject:     "audio/ilbc",
		Description: "",
	},
	"lbd": {
		Subject:     "application/vnd.llamagraphics.life-balance.desktop",
		Description: "",
	},
	"lbe": {
		Subject:     "application/vnd.llamagraphics.life-balance.exchange+xml",
		Description: "",
	},
	"lbm": {
		Subject:     "image/x-ilbm",
		Description: "",
	},
	"lca": {
		Subject:     "application/vnd.logipipe.circuit+zip",
		Description: "",
	},
	"lcs": {
		Subject:     "application/vnd.logipipe.circuit+zip",
		Description: "",
	},
	"ldif": {
		Subject:     "text/x-ldif",
		Description: "",
	},
	"les": {
		Subject:     "application/vnd.hhe.lesson-player",
		Description: "",
	},
	"lgr": {
		Subject:     "application/lgr+xml",
		Description: "",
	},
	"lha": {
		Subject:     "application/x-lha",
		Description: "",
	},
	"lhs": {
		Subject:     "text/x-literate-haskell",
		Description: "",
	},
	"lhx": {
		Subject:     "application/octet-stream",
		Description: "",
	},
	"lhz": {
		Subject:     "application/x-lhz",
		Description: "",
	},
	"lhzd": {
		Subject:     "application/vnd.belightsoft.lhzd+zip",
		Description: "",
	},
	"lhzl": {
		Subject:     "application/vnd.belightsoft.lhzl+zip",
		Description: "",
	},
	"lin": {
		Subject:     "application/x-bbolin",
		Description: "",
	},
	"line": {
		Subject:     "application/vnd.nebumind.line",
		Description: "",
	},
	"link66": {
		Subject:     "application/vnd.route66.link66+xml",
		Description: "",
	},
	"list3820": {
		Subject:     "application/vnd.afpc.modca",
		Description: "",
	},
	"list": {
		Subject:     "text/plain",
		Description: "",
	},
	"listafp": {
		Subject:     "application/vnd.afpc.modca",
		Description: "",
	},
	"lma": {
		Subject:     "audio/x-nspaudio",
		Description: "",
	},
	"lmp": {
		Subject:     "model/vnd.gs-gdl",
		Description: "",
	},
	"lnk": {
		Subject:     "application/x-ms-shortcut",
		Description: "",
	},
	"loas": {
		Subject:     "audio/aac",
		Description: "",
	},
	"log": {
		Subject:     "text/x-log",
		Description: "",
	},
	"loom": {
		Subject:     "application/vnd.loom",
		Description: "",
	},
	"lostsyncxml": {
		Subject:     "application/lostsync+xml",
		Description: "",
	},
	"lostxml": {
		Subject:     "application/lost+xml",
		Description: "",
	},
	"lrf": {
		Subject:     "application/octet-stream",
		Description: "",
	},
	"lrm": {
		Subject:     "application/vnd.ms-lrm",
		Description: "",
	},
	"lrv": {
		Subject:     "video/mp4",
		Description: "",
	},
	"lrz": {
		Subject:     "application/x-lrzip",
		Description: "",
	},
	"lsf": {
		Subject:     "video/x-la-asf",
		Description: "",
	},
	"lsp": {
		Subject:     "application/x-lisp",
		Description: "",
	},
	"lst": {
		Subject:     "text/plain",
		Description: "",
	},
	"lsx": {
		Subject:     "video/x-la-asf",
		Description: "",
	},
	"ltf": {
		Subject:     "application/vnd.frogans.ltf",
		Description: "",
	},
	"ltx": {
		Subject:     "text/x-tex",
		Description: "",
	},
	"lua": {
		Subject:     "text/x-lua",
		Description: "",
	},
	"luac": {
		Subject:     "application/x-lua-bytecode",
		Description: "",
	},
	"lvp": {
		Subject:     "audio/vnd.lucent.voice",
		Description: "",
	},
	"lwo": {
		Subject:     "image/x-lwo",
		Description: "",
	},
	"lwob": {
		Subject:     "image/x-lwo",
		Description: "",
	},
	"lwp": {
		Subject:     "application/vnd.lotus-wordpro",
		Description: "",
	},
	"lws": {
		Subject:     "image/x-lws",
		Description: "",
	},
	"lxf": {
		Subject:     "application/lxf",
		Description: "",
	},
	"ly": {
		Subject:     "text/x-lilypond",
		Description: "",
	},
	"lyx": {
		Subject:     "application/x-lyx",
		Description: "",
	},
	"lz4": {
		Subject:     "application/x-lz4",
		Description: "",
	},
	"lz": {
		Subject:     "application/x-lzip",
		Description: "",
	},
	"lzh": {
		Subject:     "application/x-lzh",
		Description: "",
	},
	"lzma": {
		Subject:     "application/x-lzma",
		Description: "",
	},
	"lzo": {
		Subject:     "application/x-lzop",
		Description: "",
	},
	"lzx": {
		Subject:     "application/x-lzx",
		Description: "",
	},
	"m13": {
		Subject:     "application/x-msmediaview",
		Description: "",
	},
	"m14": {
		Subject:     "application/x-msmediaview",
		Description: "",
	},
	"m15": {
		Subject:     "audio/x-mod",
		Description: "",
	},
	"m1u": {
		Subject:     "video/vnd.mpegurl",
		Description: "",
	},
	"m1v": {
		Subject:     "video/mpeg",
		Description: "",
	},
	"m21": {
		Subject:     "application/mp21",
		Description: "",
	},
	"m2a": {
		Subject:     "audio/mpeg",
		Description: "",
	},
	"m2t": {
		Subject:     "video/mp2t",
		Description: "",
	},
	"m2ts": {
		Subject:     "video/mp2t",
		Description: "",
	},
	"m2v": {
		Subject:     "video/mpeg",
		Description: "",
	},
	"m3a": {
		Subject:     "audio/mpeg",
		Description: "",
	},
	"m3g": {
		Subject:     "application/x-m3g",
		Description: "",
	},
	"m3u8": {
		Subject:     "application/vnd.apple.mpegurl",
		Description: "",
	},
	"m3u": {
		Subject:     "application/vnd.apple.mpegurl",
		Description: "",
	},
	"m4": {
		Subject:     "application/x-m4",
		Description: "",
	},
	"m4a": {
		Subject:     "audio/mpeg",
		Description: "",
	},
	"m4b": {
		Subject:     "audio/x-m4b",
		Description: "",
	},
	"m4p": {
		Subject:     "application/mp4",
		Description: "",
	},
	"m4s": {
		Subject:     "video/iso.segment",
		Description: "",
	},
	"m4u": {
		Subject:     "video/vnd.mpegurl",
		Description: "",
	},
	"m4v": {
		Subject:     "video/x-m4v",
		Description: "",
	},
	"m": {
		Subject:     "application/vnd.wolfram.mathematica.package",
		Description: "",
	},
	"ma": {
		Subject:     "application/mathematica",
		Description: "",
	},
	"mab": {
		Subject:     "application/x-markaby",
		Description: "",
	},
	"mads": {
		Subject:     "application/mads+xml",
		Description: "",
	},
	"maei": {
		Subject:     "application/mmt-aei+xml",
		Description: "",
	},
	"mag": {
		Subject:     "application/vnd.ecowin.chart",
		Description: "",
	},
	"mak": {
		Subject:     "text/x-makefile",
		Description: "",
	},
	"maker": {
		Subject:     "application/x-maker",
		Description: "",
	},
	"man": {
		Subject:     "application/x-troff-man",
		Description: "",
	},
	"manifest": {
		Subject:     "text/cache-manifest",
		Description: "",
	},
	"map": {
		Subject:     "application/x-navimap",
		Description: "",
	},
	"mar": {
		Subject:     "application/octet-stream",
		Description: "",
	},
	"markdown": {
		Subject:     "text/markdown",
		Description: "",
	},
	"mathml": {
		Subject:     "application/mathml+xml",
		Description: "",
	},
	"mb": {
		Subject:     "application/mathematica",
		Description: "",
	},
	"mbd": {
		Subject:     "application/x-mbedlet",
		Description: "",
	},
	"mbk": {
		Subject:     "application/vnd.mobius.mbk",
		Description: "",
	},
	"mbox": {
		Subject:     "application/mbox",
		Description: "",
	},
	"mbsdf": {
		Subject:     "application/vnd.mdl-mbsdf",
		Description: "",
	},
	"mc$": {
		Subject:     "application/x-magic-cap-package-1.0",
		Description: "",
	},
	"mc1": {
		Subject:     "application/vnd.medcalcdata",
		Description: "",
	},
	"mc2": {
		Subject:     "text/vnd.senx.warpscript",
		Description: "",
	},
	"mcd": {
		Subject:     "application/vnd.mcd",
		Description: "",
	},
	"mcf": {
		Subject:     "image/x-vasa",
		Description: "",
	},
	"mcif": {
		Subject:     "x-chemical/x-mmcif",
		Description: "",
	},
	"mcm": {
		Subject:     "x-chemical/x-macmolecule",
		Description: "",
	},
	"mcp": {
		Subject:     "application/x-netmc",
		Description: "",
	},
	"mcurl": {
		Subject:     "text/vnd.curl.mcurl",
		Description: "",
	},
	"md": {
		Subject:     "text/markdown",
		Description: "",
	},
	"mda": {
		Subject:     "application/x-msaccess",
		Description: "",
	},
	"mdb": {
		Subject:     "application/x-msaccess",
		Description: "",
	},
	"mdc": {
		Subject:     "application/vnd.marlin.drm.mdcf",
		Description: "",
	},
	"mde": {
		Subject:     "application/x-msaccess",
		Description: "",
	},
	"mdf": {
		Subject:     "application/x-msaccess",
		Description: "",
	},
	"mdi": {
		Subject:     "image/vnd.ms-modi",
		Description: "",
	},
	"mdl": {
		Subject:     "application/vnd.mdl",
		Description: "",
	},
	"mdp": {
		Subject:     "application/dash+xml",
		Description: "",
	},
	"mdx": {
		Subject:     "application/x-genesis-rom",
		Description: "",
	},
	"me": {
		Subject:     "application/x-troff-me",
		Description: "",
	},
	"med": {
		Subject:     "audio/x-mod",
		Description: "",
	},
	"mermaid": {
		Subject:     "application/vnd.mermaid",
		Description: "",
	},
	"mesh": {
		Subject:     "model/mesh",
		Description: "",
	},
	"meta4": {
		Subject:     "application/metalink4+xml",
		Description: "",
	},
	"metalink": {
		Subject:     "application/x-metalink+xml",
		Description: "",
	},
	"mets": {
		Subject:     "application/mets+xml",
		Description: "",
	},
	"mf4": {
		Subject:     "application/mf4",
		Description: "",
	},
	"mfm": {
		Subject:     "application/vnd.mfmp",
		Description: "",
	},
	"mft": {
		Subject:     "application/rpki-manifest",
		Description: "",
	},
	"mgp": {
		Subject:     "application/vnd.osgeo.mapguide.package",
		Description: "",
	},
	"mgz": {
		Subject:     "application/vnd.proteus.magazine",
		Description: "",
	},
	"mhas": {
		Subject:     "audio/mhas",
		Description: "",
	},
	"mht": {
		Subject:     "message/rfc822",
		Description: "",
	},
	"mhtml": {
		Subject:     "message/rfc822",
		Description: "",
	},
	"mid": {
		Subject:     "audio/sp-midi",
		Description: "",
	},
	"midi": {
		Subject:     "audio/x-midi",
		Description: "",
	},
	"mie": {
		Subject:     "application/x-mie",
		Description: "",
	},
	"mif": {
		Subject:     "application/vnd.mif",
		Description: "",
	},
	"mime": {
		Subject:     "message/rfc822",
		Description: "",
	},
	"minipsf": {
		Subject:     "audio/x-minipsf",
		Description: "",
	},
	"miz": {
		Subject:     "text/mizar",
		Description: "",
	},
	"mj2": {
		Subject:     "video/mj2",
		Description: "",
	},
	"mjf": {
		Subject:     "audio/x-vnd.audioexplosion.mjuicemediafile",
		Description: "",
	},
	"mjp2": {
		Subject:     "video/mj2",
		Description: "",
	},
	"mjpg": {
		Subject:     "video/x-motion-jpeg",
		Description: "",
	},
	"mjs": {
		Subject:     "text/ecmascript",
		Description: "",
	},
	"mk3d": {
		Subject:     "video/matroska-3d",
		Description: "",
	},
	"mk": {
		Subject:     "text/x-makefile",
		Description: "",
	},
	"mka": {
		Subject:     "audio/matroska",
		Description: "",
	},
	"mkd": {
		Subject:     "text/x-markdown",
		Description: "",
	},
	"mks": {
		Subject:     "video/x-matroska",
		Description: "",
	},
	"mkv": {
		Subject:     "video/matroska",
		Description: "",
	},
	"ml2": {
		Subject:     "application/vnd.sybyl.mol2",
		Description: "",
	},
	"ml": {
		Subject:     "text/x-ocaml",
		Description: "",
	},
	"mli": {
		Subject:     "text/x-ocaml",
		Description: "",
	},
	"mlp": {
		Subject:     "audio/vnd.dolby.mlp",
		Description: "",
	},
	"mm": {
		Subject:     "application/x-freemind",
		Description: "",
	},
	"mmd": {
		Subject:     "application/vnd.chipnuts.karaoke-mmd",
		Description: "",
	},
	"mmdb": {
		Subject:     "application/vnd.maxmind.maxmind-db",
		Description: "",
	},
	"mme": {
		Subject:     "application/x-base64",
		Description: "",
	},
	"mmf": {
		Subject:     "application/vnd.smaf",
		Description: "",
	},
	"mml": {
		Subject:     "text/x-mathml",
		Description: "",
	},
	"mmod": {
		Subject:     "x-chemical/x-macromodel-input",
		Description: "",
	},
	"mmr": {
		Subject:     "image/vnd.fujixerox.edmics-mmr",
		Description: "",
	},
	"mng": {
		Subject:     "video/x-mng",
		Description: "",
	},
	"mny": {
		Subject:     "application/x-msmoney",
		Description: "",
	},
	"mo3": {
		Subject:     "audio/x-mo3",
		Description: "",
	},
	"mo": {
		Subject:     "application/x-gettext-translation",
		Description: "",
	},
	"mobi": {
		Subject:     "application/x-mobipocket-ebook",
		Description: "",
	},
	"moc": {
		Subject:     "text/x-moc",
		Description: "",
	},
	"mod": {
		Subject:     "application/xml-dtd",
		Description: "",
	},
	"model-inter": {
		Subject:     "application/vnd.vd-study",
		Description: "",
	},
	"modl": {
		Subject:     "application/vnd.modl",
		Description: "",
	},
	"mods": {
		Subject:     "application/mods+xml",
		Description: "",
	},
	"mof": {
		Subject:     "text/x-mof",
		Description: "",
	},
	"mol2": {
		Subject:     "application/vnd.sybyl.mol2",
		Description: "",
	},
	"mol": {
		Subject:     "x-chemical/x-mdl-molfile",
		Description: "",
	},
	"moml": {
		Subject:     "model/vnd.moml+xml",
		Description: "",
	},
	"moo": {
		Subject:     "x-chemical/x-mopac-out",
		Description: "",
	},
	"moov": {
		Subject:     "video/quicktime",
		Description: "",
	},
	"mop": {
		Subject:     "x-chemical/x-mopac-input",
		Description: "",
	},
	"mopcrt": {
		Subject:     "x-chemical/x-mopac-input",
		Description: "",
	},
	"mov": {
		Subject:     "video/quicktime",
		Description: "",
	},
	"movie": {
		Subject:     "video/x-sgi-movie",
		Description: "",
	},
	"mp+": {
		Subject:     "audio/x-musepack",
		Description: "",
	},
	"mp1": {
		Subject:     "audio/mpeg",
		Description: "",
	},
	"mp21": {
		Subject:     "application/mp21",
		Description: "",
	},
	"mp2": {
		Subject:     "audio/mpeg",
		Description: "",
	},
	"mp2a": {
		Subject:     "audio/mpeg",
		Description: "",
	},
	"mp3": {
		Subject:     "audio/mpeg",
		Description: "",
	},
	"mp4": {
		Subject:     "video/mp4",
		Description: "",
	},
	"mp4a": {
		Subject:     "audio/mp4",
		Description: "",
	},
	"mp4s": {
		Subject:     "application/mp4",
		Description: "",
	},
	"mp4v": {
		Subject:     "video/mp4",
		Description: "",
	},
	"mpa": {
		Subject:     "audio/mpeg",
		Description: "",
	},
	"mpc": {
		Subject:     "application/vnd.mophun.certificate",
		Description: "",
	},
	"mpd": {
		Subject:     "application/dash+xml",
		Description: "",
	},
	"mpdd": {
		Subject:     "application/dashdelta",
		Description: "",
	},
	"mpe": {
		Subject:     "video/mpeg",
		Description: "",
	},
	"mpeg": {
		Subject:     "video/mpeg",
		Description: "",
	},
	"mpega": {
		Subject:     "audio/mpeg",
		Description: "",
	},
	"mpf": {
		Subject:     "application/media-policy-dataset+xml",
		Description: "",
	},
	"mpg4": {
		Subject:     "application/mp4",
		Description: "",
	},
	"mpg": {
		Subject:     "video/mpeg",
		Description: "",
	},
	"mpga": {
		Subject:     "audio/mpeg",
		Description: "",
	},
	"mpkg": {
		Subject:     "application/vnd.apple.installer+xml",
		Description: "",
	},
	"mpl": {
		Subject:     "video/mp2t",
		Description: "",
	},
	"mpls": {
		Subject:     "video/mp2t",
		Description: "",
	},
	"mpm": {
		Subject:     "application/vnd.blueice.multipass",
		Description: "",
	},
	"mpn": {
		Subject:     "application/vnd.mophun.application",
		Description: "",
	},
	"mpp": {
		Subject:     "application/dash-patch+xml",
		Description: "",
	},
	"mpt": {
		Subject:     "application/vnd.ms-project",
		Description: "",
	},
	"mpv": {
		Subject:     "video/x-matroska",
		Description: "",
	},
	"mpw": {
		Subject:     "application/vnd.exstream-empower+zip",
		Description: "",
	},
	"mpx": {
		Subject:     "application/x-project",
		Description: "",
	},
	"mpy": {
		Subject:     "application/vnd.ibm.minipay",
		Description: "",
	},
	"mqy": {
		Subject:     "application/vnd.mobius.mqy",
		Description: "",
	},
	"mrc": {
		Subject:     "application/marc",
		Description: "",
	},
	"mrcx": {
		Subject:     "application/marcxml+xml",
		Description: "",
	},
	"mrl": {
		Subject:     "text/x-mrml",
		Description: "",
	},
	"mrml": {
		Subject:     "text/x-mrml",
		Description: "",
	},
	"mrw": {
		Subject:     "image/x-minolta-mrw",
		Description: "",
	},
	"ms": {
		Subject:     "application/x-troff-ms",
		Description: "",
	},
	"msa": {
		Subject:     "application/vnd.msa-disk-image",
		Description: "",
	},
	"mscml": {
		Subject:     "application/mediaservercontrol+xml",
		Description: "",
	},
	"msd": {
		Subject:     "application/vnd.fdsn.mseed",
		Description: "",
	},
	"mseed": {
		Subject:     "application/vnd.fdsn.mseed",
		Description: "",
	},
	"mseq": {
		Subject:     "application/vnd.mseq",
		Description: "",
	},
	"msf": {
		Subject:     "application/vnd.epson.msf",
		Description: "",
	},
	"msh": {
		Subject:     "model/mesh",
		Description: "",
	},
	"msi": {
		Subject:     "application/x-msi",
		Description: "",
	},
	"msl": {
		Subject:     "application/vnd.mobius.msl",
		Description: "",
	},
	"msm": {
		Subject:     "model/vnd.gs-gdl",
		Description: "",
	},
	"msod": {
		Subject:     "image/x-msod",
		Description: "",
	},
	"msty": {
		Subject:     "application/vnd.muvee.style",
		Description: "",
	},
	"msx": {
		Subject:     "application/x-msx-rom",
		Description: "",
	},
	"mtl": {
		Subject:     "model/mtl",
		Description: "",
	},
	"mtm": {
		Subject:     "audio/x-mod",
		Description: "",
	},
	"mts": {
		Subject:     "model/vnd.mts",
		Description: "",
	},
	"multitrack": {
		Subject:     "audio/vnd.presonus.multitrack",
		Description: "",
	},
	"mup": {
		Subject:     "text/x-mup",
		Description: "",
	},
	"mus": {
		Subject:     "application/vnd.musician",
		Description: "",
	},
	"musd": {
		Subject:     "application/mmt-usd+xml",
		Description: "",
	},
	"musicxml": {
		Subject:     "application/vnd.recordare.musicxml+xml",
		Description: "",
	},
	"mv": {
		Subject:     "video/x-sgi-movie",
		Description: "",
	},
	"mvb": {
		Subject:     "x-chemical/x-mopac-vib",
		Description: "",
	},
	"mvt": {
		Subject:     "application/vnd.mapbox-vector-tile",
		Description: "",
	},
	"mwc": {
		Subject:     "application/vnd.dpgraph",
		Description: "",
	},
	"mwf": {
		Subject:     "application/vnd.mfer",
		Description: "",
	},
	"mxf": {
		Subject:     "application/mxf",
		Description: "",
	},
	"mxi": {
		Subject:     "application/vnd.vd-study",
		Description: "",
	},
	"mxl": {
		Subject:     "application/vnd.recordare.musicxml",
		Description: "",
	},
	"mxmf": {
		Subject:     "audio/mobile-xmf",
		Description: "",
	},
	"mxml": {
		Subject:     "application/xhtml-voice+xml",
		Description: "",
	},
	"mxs": {
		Subject:     "application/vnd.triscape.mxs",
		Description: "",
	},
	"mxu": {
		Subject:     "video/vnd.mpegurl",
		Description: "",
	},
	"my": {
		Subject:     "audio/x-make",
		Description: "",
	},
	"mzz": {
		Subject:     "application/x-vnd.audioexplosion.mzz",
		Description: "",
	},
	"n-gage": {
		Subject:     "application/vnd.nokia.n-gage.symbian.install",
		Description: "",
	},
	"n3": {
		Subject:     "text/n3",
		Description: "",
	},
	"n64": {
		Subject:     "application/x-n64-rom",
		Description: "",
	},
	"nap": {
		Subject:     "image/naplps",
		Description: "",
	},
	"naplps": {
		Subject:     "image/naplps",
		Description: "",
	},
	"nb": {
		Subject:     "application/mathematica",
		Description: "",
	},
	"nbp": {
		Subject:     "application/vnd.wolfram.player",
		Description: "",
	},
	"nc": {
		Subject:     "application/x-netcdf",
		Description: "",
	},
	"ncm": {
		Subject:     "application/vnd.nokia.configuration-message",
		Description: "",
	},
	"ncx": {
		Subject:     "application/x-dtbncx+xml",
		Description: "",
	},
	"nds": {
		Subject:     "application/vnd.nintendo.nitro.rom",
		Description: "",
	},
	"nebul": {
		Subject:     "application/vnd.nebumind.line",
		Description: "",
	},
	"nef": {
		Subject:     "image/x-nikon-nef",
		Description: "",
	},
	"nes": {
		Subject:     "application/x-nes-rom",
		Description: "",
	},
	"nez": {
		Subject:     "application/x-nes-rom",
		Description: "",
	},
	"nfo": {
		Subject:     "text/x-nfo",
		Description: "",
	},
	"ngdat": {
		Subject:     "application/vnd.nokia.n-gage.data",
		Description: "",
	},
	"nif": {
		Subject:     "image/x-niff",
		Description: "",
	},
	"niff": {
		Subject:     "image/x-niff",
		Description: "",
	},
	"nim": {
		Subject:     "video/vnd.nokia.interleaved-multimedia",
		Description: "",
	},
	"nimn": {
		Subject:     "application/vnd.nimn",
		Description: "",
	},
	"nitf": {
		Subject:     "application/vnd.nitf",
		Description: "",
	},
	"nix": {
		Subject:     "application/x-mix-transfer",
		Description: "",
	},
	"nlu": {
		Subject:     "application/vnd.neurolanguage.nlu",
		Description: "",
	},
	"nml": {
		Subject:     "application/vnd.enliven",
		Description: "",
	},
	"nnd": {
		Subject:     "application/vnd.noblenet-directory",
		Description: "",
	},
	"nns": {
		Subject:     "application/vnd.noblenet-sealer",
		Description: "",
	},
	"nnw": {
		Subject:     "application/vnd.noblenet-web",
		Description: "",
	},
	"not": {
		Subject:     "text/x-mup",
		Description: "",
	},
	"notebook": {
		Subject:     "application/vnd.smart.notebook",
		Description: "",
	},
	"npx": {
		Subject:     "image/vnd.net-fpx",
		Description: "",
	},
	"nq": {
		Subject:     "application/n-quads",
		Description: "",
	},
	"nsc": {
		Subject:     "application/x-conference",
		Description: "",
	},
	"nsf": {
		Subject:     "application/vnd.lotus-notes",
		Description: "",
	},
	"nsv": {
		Subject:     "video/x-nsv",
		Description: "",
	},
	"nt": {
		Subject:     "application/n-triples",
		Description: "",
	},
	"ntf": {
		Subject:     "application/vnd.nitf",
		Description: "",
	},
	"numbers": {
		Subject:     "application/vnd.apple.numbers",
		Description: "",
	},
	"nvd": {
		Subject:     "application/x-navidoc",
		Description: "",
	},
	"nwc": {
		Subject:     "application/x-nwc",
		Description: "",
	},
	"nws": {
		Subject:     "message/rfc822",
		Description: "",
	},
	"nzb": {
		Subject:     "application/x-nzb",
		Description: "",
	},
	"o": {
		Subject:     "application/x-object",
		Description: "",
	},
	"oa2": {
		Subject:     "application/vnd.fujitsu.oasys2",
		Description: "",
	},
	"oa3": {
		Subject:     "application/vnd.fujitsu.oasys3",
		Description: "",
	},
	"oas": {
		Subject:     "application/vnd.fujitsu.oasys",
		Description: "",
	},
	"ob": {
		Subject:     "application/vnd.1ob",
		Description: "",
	},
	"obd": {
		Subject:     "application/x-msbinder",
		Description: "",
	},
	"obg": {
		Subject:     "application/vnd.openblox.game-binary",
		Description: "",
	},
	"obgx": {
		Subject:     "application/vnd.openblox.game+xml",
		Description: "",
	},
	"obj": {
		Subject:     "model/obj",
		Description: "",
	},
	"ocl": {
		Subject:     "text/x-ocl",
		Description: "",
	},
	"oda": {
		Subject:     "application/oda",
		Description: "",
	},
	"odb": {
		Subject:     "application/vnd.oasis.opendocument.base",
		Description: "",
	},
	"odc": {
		Subject:     "application/vnd.oasis.opendocument.chart",
		Description: "",
	},
	"odf": {
		Subject:     "application/vnd.oasis.opendocument.formula",
		Description: "",
	},
	"odft": {
		Subject:     "application/vnd.oasis.opendocument.formula-template",
		Description: "",
	},
	"odg": {
		Subject:     "application/vnd.oasis.opendocument.graphics",
		Description: "",
	},
	"odi": {
		Subject:     "application/vnd.oasis.opendocument.image",
		Description: "",
	},
	"odm": {
		Subject:     "application/vnd.oasis.opendocument.text-master",
		Description: "",
	},
	"odp": {
		Subject:     "application/vnd.oasis.opendocument.presentation",
		Description: "",
	},
	"ods": {
		Subject:     "application/vnd.oasis.opendocument.spreadsheet",
		Description: "",
	},
	"odt": {
		Subject:     "application/vnd.oasis.opendocument.text",
		Description: "",
	},
	"odx": {
		Subject:     "application/odx",
		Description: "",
	},
	"oeb": {
		Subject:     "application/vnd.openeye.oeb",
		Description: "",
	},
	"oga": {
		Subject:     "audio/ogg",
		Description: "",
	},
	"ogex": {
		Subject:     "model/vnd.opengex",
		Description: "",
	},
	"ogg": {
		Subject:     "audio/ogg",
		Description: "",
	},
	"ogm": {
		Subject:     "video/x-ogm+ogg",
		Description: "",
	},
	"ogv": {
		Subject:     "video/ogg",
		Description: "",
	},
	"ogx": {
		Subject:     "application/ogg",
		Description: "",
	},
	"old": {
		Subject:     "application/x-trash",
		Description: "",
	},
	"oleo": {
		Subject:     "application/x-oleo",
		Description: "",
	},
	"omc": {
		Subject:     "application/x-omc",
		Description: "",
	},
	"omcd": {
		Subject:     "application/x-omcdatamaker",
		Description: "",
	},
	"omcr": {
		Subject:     "application/x-omcregerator",
		Description: "",
	},
	"omdoc": {
		Subject:     "application/x-omdoc+xml",
		Description: "",
	},
	"omg": {
		Subject:     "audio/atrac-advanced-lossless",
		Description: "",
	},
	"onepkg": {
		Subject:     "application/x-onenote",
		Description: "",
	},
	"onetmp": {
		Subject:     "application/x-onenote",
		Description: "",
	},
	"onetoc2": {
		Subject:     "application/x-onenote",
		Description: "",
	},
	"onetoc": {
		Subject:     "application/x-onenote",
		Description: "",
	},
	"ooc": {
		Subject:     "text/x-ooc",
		Description: "",
	},
	"opf": {
		Subject:     "application/oebps-package+xml",
		Description: "",
	},
	"opml": {
		Subject:     "text/x-opml",
		Description: "",
	},
	"oprc": {
		Subject:     "application/vnd.palm",
		Description: "",
	},
	"opus": {
		Subject:     "audio/ogg",
		Description: "",
	},
	"or3": {
		Subject:     "application/vnd.lotus-organizer",
		Description: "",
	},
	"ora": {
		Subject:     "image/x-openraster",
		Description: "",
	},
	"orf": {
		Subject:     "image/x-olympus-orf",
		Description: "",
	},
	"org": {
		Subject:     "application/vnd.lotus-organizer",
		Description: "",
	},
	"orq": {
		Subject:     "application/ocsp-request",
		Description: "",
	},
	"ors": {
		Subject:     "application/ocsp-response",
		Description: "",
	},
	"osf": {
		Subject:     "application/vnd.yamaha.openscoreformat",
		Description: "",
	},
	"osfpvg": {
		Subject:     "application/vnd.yamaha.openscoreformat.osfpvg+xml",
		Description: "",
	},
	"osm": {
		Subject:     "application/vnd.openstreetmap.data+xml",
		Description: "",
	},
	"ota": {
		Subject:     "application/vnd.android.ota",
		Description: "",
	},
	"otc": {
		Subject:     "application/vnd.oasis.opendocument.chart-template",
		Description: "",
	},
	"otf": {
		Subject:     "font/otf",
		Description: "",
	},
	"otg": {
		Subject:     "application/vnd.oasis.opendocument.graphics-template",
		Description: "",
	},
	"oth": {
		Subject:     "application/vnd.oasis.opendocument.text-web",
		Description: "",
	},
	"oti": {
		Subject:     "application/vnd.oasis.opendocument.image-template",
		Description: "",
	},
	"otm": {
		Subject:     "application/vnd.oasis.opendocument.text-master-template",
		Description: "",
	},
	"otp": {
		Subject:     "application/vnd.oasis.opendocument.presentation-template",
		Description: "",
	},
	"ots": {
		Subject:     "application/vnd.oasis.opendocument.spreadsheet-template",
		Description: "",
	},
	"ott": {
		Subject:     "application/vnd.oasis.opendocument.text-template",
		Description: "",
	},
	"ovl": {
		Subject:     "application/vnd.afpc.modca-overlay",
		Description: "",
	},
	"owl": {
		Subject:     "application/vnd.biopax.rdf+xml",
		Description: "",
	},
	"owx": {
		Subject:     "application/x-owl+xml",
		Description: "",
	},
	"oxlicg": {
		Subject:     "application/vnd.oxli.countgraph",
		Description: "",
	},
	"oxps": {
		Subject:     "application/oxps",
		Description: "",
	},
	"oxt": {
		Subject:     "application/vnd.openofficeorg.extension",
		Description: "",
	},
	"oza": {
		Subject:     "application/x-oz-application",
		Description: "",
	},
	"p10": {
		Subject:     "application/pkcs10",
		Description: "",
	},
	"p12": {
		Subject:     "application/pkcs12",
		Description: "",
	},
	"p21": {
		Subject:     "application/p21",
		Description: "",
	},
	"p65": {
		Subject:     "application/x-pagemaker",
		Description: "",
	},
	"p7a": {
		Subject:     "application/x-pkcs7-signature",
		Description: "",
	},
	"p7b": {
		Subject:     "application/x-pkcs7-certificates",
		Description: "",
	},
	"p7c": {
		Subject:     "application/pkcs7-mime",
		Description: "",
	},
	"p7m": {
		Subject:     "application/pkcs7-mime",
		Description: "",
	},
	"p7r": {
		Subject:     "application/x-pkcs7-certreqresp",
		Description: "",
	},
	"p7s": {
		Subject:     "application/pkcs7-signature",
		Description: "",
	},
	"p8": {
		Subject:     "application/pkcs8",
		Description: "",
	},
	"p8e": {
		Subject:     "application/pkcs8-encrypted",
		Description: "",
	},
	"p": {
		Subject:     "text/x-pascal",
		Description: "",
	},
	"pac": {
		Subject:     "application/x-ns-proxy-autoconfig",
		Description: "",
	},
	"pack": {
		Subject:     "application/x-java-pack200",
		Description: "",
	},
	"package": {
		Subject:     "application/vnd.autopackage",
		Description: "",
	},
	"pages": {
		Subject:     "application/vnd.apple.pages",
		Description: "",
	},
	"pak": {
		Subject:     "application/x-pak",
		Description: "",
	},
	"par2": {
		Subject:     "application/x-par2",
		Description: "",
	},
	"part": {
		Subject:     "application/x-pro_eng",
		Description: "",
	},
	"pas": {
		Subject:     "text/x-pascal",
		Description: "",
	},
	"pat": {
		Subject:     "image/x-coreldrawpattern",
		Description: "",
	},
	"patch": {
		Subject:     "text/x-diff",
		Description: "",
	},
	"paw": {
		Subject:     "application/vnd.pawaafile",
		Description: "",
	},
	"pbd": {
		Subject:     "application/vnd.powerbuilder6",
		Description: "",
	},
	"pbm": {
		Subject:     "image/x-portable-bitmap",
		Description: "",
	},
	"pcap": {
		Subject:     "application/vnd.tcpdump.pcap",
		Description: "",
	},
	"pcd": {
		Subject:     "image/x-photo-cd",
		Description: "",
	},
	"pce": {
		Subject:     "application/x-pc-engine-rom",
		Description: "",
	},
	"pcf.z": {
		Subject:     "application/x-font",
		Description: "",
	},
	"pcf": {
		Subject:     "application/x-font",
		Description: "",
	},
	"pcl": {
		Subject:     "application/vnd.hp-pcl",
		Description: "",
	},
	"pclxl": {
		Subject:     "application/vnd.hp-pclxl",
		Description: "",
	},
	"pct": {
		Subject:     "image/x-pict",
		Description: "",
	},
	"pcurl": {
		Subject:     "application/vnd.curl.pcurl",
		Description: "",
	},
	"pcx": {
		Subject:     "image/vnd.zbrush.pcx",
		Description: "",
	},
	"pdb": {
		Subject:     "application/vnd.palm",
		Description: "",
	},
	"pdc": {
		Subject:     "application/x-aportisdoc",
		Description: "",
	},
	"pdf": {
		Subject:     "application/pdf",
		Description: "",
	},
	"pdx": {
		Subject:     "application/pdx",
		Description: "",
	},
	"pef": {
		Subject:     "image/x-pentax-pef",
		Description: "",
	},
	"pem": {
		Subject:     "application/pem-certificate-chain",
		Description: "",
	},
	"perl": {
		Subject:     "application/x-perl",
		Description: "",
	},
	"pfa": {
		Subject:     "application/x-font",
		Description: "",
	},
	"pfb": {
		Subject:     "application/x-font",
		Description: "",
	},
	"pfm": {
		Subject:     "application/x-font-type1",
		Description: "",
	},
	"pfr": {
		Subject:     "application/vnd.dvb.pfr",
		Description: "",
	},
	"pfunk": {
		Subject:     "audio/x-make",
		Description: "",
	},
	"pfx": {
		Subject:     "application/pkcs12",
		Description: "",
	},
	"pgb": {
		Subject:     "image/vnd.globalgraphics.pgb",
		Description: "",
	},
	"pgm": {
		Subject:     "image/x-portable-graymap",
		Description: "",
	},
	"pgn": {
		Subject:     "application/vnd.chess-pgn",
		Description: "",
	},
	"pgp": {
		Subject:     "application/pgp-signature",
		Description: "",
	},
	"php3": {
		Subject:     "application/x-httpd-php3",
		Description: "",
	},
	"php3p": {
		Subject:     "application/x-httpd-php3-preprocessed",
		Description: "",
	},
	"php4": {
		Subject:     "application/x-httpd-php4",
		Description: "",
	},
	"php5": {
		Subject:     "application/x-httpd-php5",
		Description: "",
	},
	"php": {
		Subject:     "application/x-httpd-php",
		Description: "",
	},
	"phps": {
		Subject:     "application/x-httpd-php-source",
		Description: "",
	},
	"pht": {
		Subject:     "application/x-httpd-php",
		Description: "",
	},
	"phtml": {
		Subject:     "application/x-httpd-php",
		Description: "",
	},
	"pic": {
		Subject:     "image/vnd.radiance",
		Description: "",
	},
	"pict1": {
		Subject:     "image/x-pict",
		Description: "",
	},
	"pict2": {
		Subject:     "image/x-pict",
		Description: "",
	},
	"pict": {
		Subject:     "image/x-pict",
		Description: "",
	},
	"pil": {
		Subject:     "application/vnd.piaccess.application-licence",
		Description: "",
	},
	"pk": {
		Subject:     "application/x-tex-pk",
		Description: "",
	},
	"pkd": {
		Subject:     "application/vnd.fints",
		Description: "",
	},
	"pkg": {
		Subject:     "application/vnd.apple.installer+xml",
		Description: "",
	},
	"pki": {
		Subject:     "application/pkixcmp",
		Description: "",
	},
	"pkipath": {
		Subject:     "application/pkix-pkipath",
		Description: "",
	},
	"pko": {
		Subject:     "application/vnd.ms-pki.pko",
		Description: "",
	},
	"pkr": {
		Subject:     "application/pgp-keys",
		Description: "",
	},
	"pl": {
		Subject:     "application/x-perl",
		Description: "",
	},
	"pla": {
		Subject:     "audio/x-iriver-pla",
		Description: "",
	},
	"plb": {
		Subject:     "application/vnd.3gpp.pic-bw-large",
		Description: "",
	},
	"plc": {
		Subject:     "application/vnd.mobius.plc",
		Description: "",
	},
	"plf": {
		Subject:     "application/vnd.pocketlearn",
		Description: "",
	},
	"plj": {
		Subject:     "audio/vnd.everad.plj",
		Description: "",
	},
	"pln": {
		Subject:     "application/x-planperfect",
		Description: "",
	},
	"plp": {
		Subject:     "application/vnd.panoply",
		Description: "",
	},
	"pls": {
		Subject:     "audio/x-scpls",
		Description: "",
	},
	"plt": {
		Subject:     "application/vnd.hp-hpgl",
		Description: "",
	},
	"plx": {
		Subject:     "application/x-pixclscript",
		Description: "",
	},
	"pm4": {
		Subject:     "application/x-pagemaker",
		Description: "",
	},
	"pm5": {
		Subject:     "application/x-pagemaker",
		Description: "",
	},
	"pm6": {
		Subject:     "application/x-pagemaker",
		Description: "",
	},
	"pm": {
		Subject:     "application/x-pagemaker",
		Description: "",
	},
	"pmd": {
		Subject:     "application/x-pagemaker",
		Description: "",
	},
	"pml": {
		Subject:     "application/vnd.ctc-posml",
		Description: "",
	},
	"png": {
		Subject:     "image/png",
		Description: "",
	},
	"pnm": {
		Subject:     "image/x-portable-anymap",
		Description: "",
	},
	"pntg": {
		Subject:     "image/x-macpaint",
		Description: "",
	},
	"po": {
		Subject:     "text/x-gettext-translation",
		Description: "",
	},
	"pod": {
		Subject:     "application/x-perl",
		Description: "",
	},
	"por": {
		Subject:     "application/x-spss-por",
		Description: "",
	},
	"portpkg": {
		Subject:     "application/vnd.macports.portpkg",
		Description: "",
	},
	"pot": {
		Subject:     "application/vnd.ms-powerpoint",
		Description: "",
	},
	"potm": {
		Subject:     "application/vnd.ms-powerpoint.template.macroenabled.12",
		Description: "",
	},
	"potx": {
		Subject:     "application/vnd.openxmlformats-officedocument.presentationml.template",
		Description: "",
	},
	"pov": {
		Subject:     "model/x-pov",
		Description: "",
	},
	"ppa": {
		Subject:     "application/vnd.ms-powerpoint",
		Description: "",
	},
	"ppam": {
		Subject:     "application/vnd.ms-powerpoint.addin.macroenabled.12",
		Description: "",
	},
	"ppd.gz": {
		Subject:     "application/vnd.cups-ppd",
		Description: "",
	},
	"ppd": {
		Subject:     "application/vnd.cups-ppd",
		Description: "",
	},
	"ppkg": {
		Subject:     "application/vnd.xmpie.ppkg",
		Description: "",
	},
	"ppm": {
		Subject:     "image/x-portable-pixmap",
		Description: "",
	},
	"pps": {
		Subject:     "application/vnd.ms-powerpoint",
		Description: "",
	},
	"ppsm": {
		Subject:     "application/vnd.ms-powerpoint.slideshow.macroenabled.12",
		Description: "",
	},
	"ppsx": {
		Subject:     "application/vnd.openxmlformats-officedocument.presentationml.slideshow",
		Description: "",
	},
	"ppt": {
		Subject:     "application/vnd.ms-powerpoint",
		Description: "",
	},
	"pptm": {
		Subject:     "application/vnd.ms-powerpoint.presentation.macroenabled.12",
		Description: "",
	},
	"ppttc": {
		Subject:     "application/vnd.think-cell.ppttc+json",
		Description: "",
	},
	"pptx": {
		Subject:     "application/vnd.openxmlformats-officedocument.presentationml.presentation",
		Description: "",
	},
	"ppz": {
		Subject:     "application/x-mspowerpoint",
		Description: "",
	},
	"pqa": {
		Subject:     "application/vnd.palm",
		Description: "",
	},
	"prc": {
		Subject:     "model/prc",
		Description: "",
	},
	"pre": {
		Subject:     "application/vnd.lotus-freelance",
		Description: "",
	},
	"preminet": {
		Subject:     "application/vnd.preminet",
		Description: "",
	},
	"prf": {
		Subject:     "application/x-pics-rules",
		Description: "",
	},
	"provx": {
		Subject:     "application/provenance+xml",
		Description: "",
	},
	"prt": {
		Subject:     "x-chemical/x-ncbi-asn1-ascii",
		Description: "",
	},
	"prz": {
		Subject:     "application/vnd.lotus-freelance",
		Description: "",
	},
	"ps-z": {
		Subject:     "application/postscript",
		Description: "",
	},
	"ps": {
		Subject:     "application/postscript",
		Description: "",
	},
	"psb": {
		Subject:     "application/vnd.3gpp.pic-bw-small",
		Description: "",
	},
	"psd": {
		Subject:     "image/vnd.adobe.photoshop",
		Description: "",
	},
	"pseg3820": {
		Subject:     "application/vnd.afpc.modca",
		Description: "",
	},
	"psf": {
		Subject:     "application/x-font-linux-psf",
		Description: "",
	},
	"psflib": {
		Subject:     "audio/x-psflib",
		Description: "",
	},
	"psfs": {
		Subject:     "application/vnd.psfs",
		Description: "",
	},
	"psg": {
		Subject:     "application/vnd.afpc.modca-pagesegment",
		Description: "",
	},
	"psid": {
		Subject:     "audio/prs.sid",
		Description: "",
	},
	"pskcxml": {
		Subject:     "application/pskc+xml",
		Description: "",
	},
	"psw": {
		Subject:     "application/x-pocket-word",
		Description: "",
	},
	"pt5": {
		Subject:     "application/x-pagemaker",
		Description: "",
	},
	"pt": {
		Subject:     "application/vnd.snesdev-page-table",
		Description: "",
	},
	"pti": {
		Subject:     "application/vnd.pvi.ptid1",
		Description: "",
	},
	"ptid": {
		Subject:     "application/vnd.pvi.ptid1",
		Description: "",
	},
	"ptrom": {
		Subject:     "application/vnd.snesdev-page-table",
		Description: "",
	},
	"pub": {
		Subject:     "application/vnd.exstream-package",
		Description: "",
	},
	"pvb": {
		Subject:     "application/vnd.3gpp.pic-bw-var",
		Description: "",
	},
	"pvu": {
		Subject:     "x-paleovu/x-pv",
		Description: "",
	},
	"pw": {
		Subject:     "application/x-pw",
		Description: "",
	},
	"pwn": {
		Subject:     "application/vnd.3m.post-it-notes",
		Description: "",
	},
	"pwz": {
		Subject:     "application/vnd.ms-powerpoint",
		Description: "",
	},
	"py": {
		Subject:     "application/x-python",
		Description: "",
	},
	"pya": {
		Subject:     "audio/vnd.ms-playready.media.pya",
		Description: "",
	},
	"pyc": {
		Subject:     "application/x-python-code",
		Description: "",
	},
	"pyo": {
		Subject:     "model/vnd.pytha.pyox",
		Description: "",
	},
	"pyox": {
		Subject:     "model/vnd.pytha.pyox",
		Description: "",
	},
	"pyv": {
		Subject:     "video/vnd.ms-playready.media.pyv",
		Description: "",
	},
	"pyx": {
		Subject:     "text/x-python",
		Description: "",
	},
	"qam": {
		Subject:     "application/vnd.epson.quickanime",
		Description: "",
	},
	"qbo": {
		Subject:     "application/vnd.intu.qbo",
		Description: "",
	},
	"qca": {
		Subject:     "application/vnd.ericsson.quickcall",
		Description: "",
	},
	"qcall": {
		Subject:     "application/vnd.ericsson.quickcall",
		Description: "",
	},
	"qcp": {
		Subject:     "audio/evrc-qcp",
		Description: "",
	},
	"qd3": {
		Subject:     "x-world/x-3dmf",
		Description: "",
	},
	"qd3d": {
		Subject:     "x-world/x-3dmf",
		Description: "",
	},
	"qfx": {
		Subject:     "application/vnd.intu.qfx",
		Description: "",
	},
	"qgs": {
		Subject:     "application/x-qgis",
		Description: "",
	},
	"qif": {
		Subject:     "image/x-quicktime",
		Description: "",
	},
	"qml": {
		Subject:     "text/x-qml",
		Description: "",
	},
	"qmlproject": {
		Subject:     "text/x-qml",
		Description: "",
	},
	"qmltypes": {
		Subject:     "text/x-qml",
		Description: "",
	},
	"qp": {
		Subject:     "application/x-qpress",
		Description: "",
	},
	"qps": {
		Subject:     "application/vnd.publishare-delta-tree",
		Description: "",
	},
	"qt": {
		Subject:     "video/quicktime",
		Description: "",
	},
	"qtc": {
		Subject:     "video/x-qtc",
		Description: "",
	},
	"qti": {
		Subject:     "image/x-quicktime",
		Description: "",
	},
	"qtif": {
		Subject:     "image/x-quicktime",
		Description: "",
	},
	"qtl": {
		Subject:     "application/x-quicktimeplayer",
		Description: "",
	},
	"qtvr": {
		Subject:     "video/quicktime",
		Description: "",
	},
	"quiz": {
		Subject:     "application/vnd.quobject-quoxdocument",
		Description: "",
	},
	"quox": {
		Subject:     "application/vnd.quobject-quoxdocument",
		Description: "",
	},
	"qvd": {
		Subject:     "application/vnd.theqvd",
		Description: "",
	},
	"qwd": {
		Subject:     "application/vnd.quark.quarkxpress",
		Description: "",
	},
	"qwt": {
		Subject:     "application/vnd.quark.quarkxpress",
		Description: "",
	},
	"qxb": {
		Subject:     "application/vnd.quark.quarkxpress",
		Description: "",
	},
	"qxd": {
		Subject:     "application/vnd.quark.quarkxpress",
		Description: "",
	},
	"qxl": {
		Subject:     "application/vnd.quark.quarkxpress",
		Description: "",
	},
	"qxt": {
		Subject:     "application/vnd.quark.quarkxpress",
		Description: "",
	},
	"ra": {
		Subject:     "audio/x-pn-realaudio",
		Description: "",
	},
	"raf": {
		Subject:     "image/x-fuji-raf",
		Description: "",
	},
	"ram": {
		Subject:     "audio/x-pn-realaudio",
		Description: "",
	},
	"rapd": {
		Subject:     "application/route-apd+xml",
		Description: "",
	},
	"rar": {
		Subject:     "application/vnd.rar",
		Description: "",
	},
	"ras": {
		Subject:     "image/x-cmu-raster",
		Description: "",
	},
	"rast": {
		Subject:     "image/x-cmu-raster",
		Description: "",
	},
	"raw-disk-image": {
		Subject:     "application/x-raw-disk-image",
		Description: "",
	},
	"raw": {
		Subject:     "image/x-panasonic-raw",
		Description: "",
	},
	"rax": {
		Subject:     "audio/vnd.rn-realaudio",
		Description: "",
	},
	"rb": {
		Subject:     "application/x-ruby",
		Description: "",
	},
	"rbw": {
		Subject:     "application/x-ruby",
		Description: "",
	},
	"rcprofile": {
		Subject:     "application/vnd.ipunplugged.rcprofile",
		Description: "",
	},
	"rct": {
		Subject:     "application/prs.nprend",
		Description: "",
	},
	"rd": {
		Subject:     "x-chemical/x-mdl-rdfile",
		Description: "",
	},
	"rdf-crypt": {
		Subject:     "application/prs.rdf-xml-crypt",
		Description: "",
	},
	"rdf": {
		Subject:     "application/rdf+xml",
		Description: "",
	},
	"rdfs": {
		Subject:     "application/rdf+xml",
		Description: "",
	},
	"rdz": {
		Subject:     "application/vnd.data-vision.rdz",
		Description: "",
	},
	"reg": {
		Subject:     "text/x-ms-regedit",
		Description: "",
	},
	"rej": {
		Subject:     "text/x-reject",
		Description: "",
	},
	"relo": {
		Subject:     "application/p2p-overlay+xml",
		Description: "",
	},
	"reload": {
		Subject:     "application/vnd.resilient.logic",
		Description: "",
	},
	"rep": {
		Subject:     "application/vnd.businessobjects",
		Description: "",
	},
	"req": {
		Subject:     "application/vnd.nervana",
		Description: "",
	},
	"request": {
		Subject:     "application/vnd.nervana",
		Description: "",
	},
	"res": {
		Subject:     "application/x-dtbresource+xml",
		Description: "",
	},
	"rexx": {
		Subject:     "text/x-script.rexx",
		Description: "",
	},
	"rf": {
		Subject:     "image/vnd.rn-realflash",
		Description: "",
	},
	"rfcxml": {
		Subject:     "application/rfc+xml",
		Description: "",
	},
	"rgb": {
		Subject:     "image/x-rgb",
		Description: "",
	},
	"rgbe": {
		Subject:     "image/vnd.radiance",
		Description: "",
	},
	"rhtml": {
		Subject:     "application/x-html+ruby",
		Description: "",
	},
	"rif": {
		Subject:     "application/reginfo+xml",
		Description: "",
	},
	"rip": {
		Subject:     "audio/vnd.rip",
		Description: "",
	},
	"ris": {
		Subject:     "application/x-research-info-systems",
		Description: "",
	},
	"rl": {
		Subject:     "application/resource-lists+xml",
		Description: "",
	},
	"rlc": {
		Subject:     "image/vnd.fujixerox.edmics-rlc",
		Description: "",
	},
	"rld": {
		Subject:     "application/resource-lists-diff+xml",
		Description: "",
	},
	"rle": {
		Subject:     "image/x-rle",
		Description: "",
	},
	"rlm": {
		Subject:     "application/vnd.resilient.logic",
		Description: "",
	},
	"rm": {
		Subject:     "audio/vnd.hns.audio",
		Description: "",
	},
	"rmi": {
		Subject:     "audio/x-midi",
		Description: "",
	},
	"rmj": {
		Subject:     "application/vnd.rn-realmedia",
		Description: "",
	},
	"rmm": {
		Subject:     "audio/x-pn-realaudio",
		Description: "",
	},
	"rmp": {
		Subject:     "audio/x-pn-realaudio-plugin",
		Description: "",
	},
	"rms": {
		Subject:     "application/vnd.jcp.javame.midlet-rms",
		Description: "",
	},
	"rmvb": {
		Subject:     "application/vnd.rn-realmedia-vbr",
		Description: "",
	},
	"rmx": {
		Subject:     "application/vnd.rn-realmedia",
		Description: "",
	},
	"rnc": {
		Subject:     "application/relax-ng-compact-syntax",
		Description: "",
	},
	"rnd": {
		Subject:     "application/prs.nprend",
		Description: "",
	},
	"rng": {
		Subject:     "application/x-ringing-tones",
		Description: "",
	},
	"rnx": {
		Subject:     "application/vnd.rn-realplayer",
		Description: "",
	},
	"roa": {
		Subject:     "application/rpki-roa",
		Description: "",
	},
	"roff": {
		Subject:     "application/x-troff",
		Description: "",
	},
	"ros": {
		Subject:     "x-chemical/x-rosdal",
		Description: "",
	},
	"rp9": {
		Subject:     "application/vnd.cloanto.rp9",
		Description: "",
	},
	"rp": {
		Subject:     "application/vnd.relpipe",
		Description: "",
	},
	"rpm": {
		Subject:     "application/x-redhat-package-manager",
		Description: "",
	},
	"rpss": {
		Subject:     "application/vnd.nokia.radio-presets",
		Description: "",
	},
	"rpst": {
		Subject:     "application/vnd.nokia.radio-preset",
		Description: "",
	},
	"rq": {
		Subject:     "application/sparql-query",
		Description: "",
	},
	"rs": {
		Subject:     "application/rls-services+xml",
		Description: "",
	},
	"rsat": {
		Subject:     "application/atsc-rsat+xml",
		Description: "",
	},
	"rsd": {
		Subject:     "application/x-rsd+xml",
		Description: "",
	},
	"rsheet": {
		Subject:     "application/urc-ressheet+xml",
		Description: "",
	},
	"rsm": {
		Subject:     "model/vnd.gs-gdl",
		Description: "",
	},
	"rss": {
		Subject:     "application/x-rss+xml",
		Description: "",
	},
	"rst": {
		Subject:     "text/prs.fallenstein.rst",
		Description: "",
	},
	"rt": {
		Subject:     "text/richtext",
		Description: "",
	},
	"rtf": {
		Subject:     "application/rtf",
		Description: "",
	},
	"rtx": {
		Subject:     "text/richtext",
		Description: "",
	},
	"rusd": {
		Subject:     "application/route-usd+xml",
		Description: "",
	},
	"rv": {
		Subject:     "video/vnd.rn-realvideo",
		Description: "",
	},
	"rvx": {
		Subject:     "video/vnd.rn-realvideo",
		Description: "",
	},
	"rw2": {
		Subject:     "image/x-panasonic-raw2",
		Description: "",
	},
	"rxn": {
		Subject:     "x-chemical/x-mdl-rxnfile",
		Description: "",
	},
	"rxr": {
		Subject:     "application/vnd.medicalholodeck.recordxr",
		Description: "",
	},
	"s11": {
		Subject:     "video/vnd.sealed.mpeg1",
		Description: "",
	},
	"s14": {
		Subject:     "video/vnd.sealed.mpeg4",
		Description: "",
	},
	"s1a": {
		Subject:     "application/vnd.sealedmedia.softseal.pdf",
		Description: "",
	},
	"s1e": {
		Subject:     "application/vnd.sealed.xls",
		Description: "",
	},
	"s1g": {
		Subject:     "image/vnd.sealedmedia.softseal.gif",
		Description: "",
	},
	"s1h": {
		Subject:     "application/vnd.sealedmedia.softseal.html",
		Description: "",
	},
	"s1j": {
		Subject:     "image/vnd.sealedmedia.softseal.jpg",
		Description: "",
	},
	"s1m": {
		Subject:     "audio/vnd.sealedmedia.softseal.mpeg",
		Description: "",
	},
	"s1n": {
		Subject:     "image/vnd.sealed.png",
		Description: "",
	},
	"s1p": {
		Subject:     "application/vnd.sealed.ppt",
		Description: "",
	},
	"s1q": {
		Subject:     "video/vnd.sealedmedia.softseal.mov",
		Description: "",
	},
	"s1w": {
		Subject:     "application/vnd.sealed.doc",
		Description: "",
	},
	"s3df": {
		Subject:     "application/vnd.sealed.3df",
		Description: "",
	},
	"s3m": {
		Subject:     "audio/x-s3m",
		Description: "",
	},
	"s": {
		Subject:     "text/x-asm",
		Description: "",
	},
	"sac": {
		Subject:     "application/tamp-sequence-adjust-confirm",
		Description: "",
	},
	"saf": {
		Subject:     "application/vnd.yamaha.smaf-audio",
		Description: "",
	},
	"sam": {
		Subject:     "application/x-amipro",
		Description: "",
	},
	"sami": {
		Subject:     "application/x-sami",
		Description: "",
	},
	"sarif-external-properties": {
		Subject:     "application/sarif-external-properties+json",
		Description: "",
	},
	"sarif.json": {
		Subject:     "application/sarif+json",
		Description: "",
	},
	"sarif": {
		Subject:     "application/sarif+json",
		Description: "",
	},
	"sav": {
		Subject:     "application/x-spss",
		Description: "",
	},
	"saveme": {
		Subject:     "application/octet-stream",
		Description: "",
	},
	"sbk": {
		Subject:     "application/x-tbook",
		Description: "",
	},
	"sbml": {
		Subject:     "application/sbml+xml",
		Description: "",
	},
	"sbs": {
		Subject:     "application/x-spss",
		Description: "",
	},
	"sc": {
		Subject:     "application/vnd.ibm.secure-container",
		Description: "",
	},
	"scala": {
		Subject:     "text/x-scala",
		Description: "",
	},
	"scd.gz": {
		Subject:     "application/vnd.scribus",
		Description: "",
	},
	"scd": {
		Subject:     "application/vnd.scribus",
		Description: "",
	},
	"sce": {
		Subject:     "application/vnd.etsi.asic-e+zip",
		Description: "",
	},
	"scim": {
		Subject:     "application/scim+json",
		Description: "",
	},
	"scl": {
		Subject:     "application/vnd.sycle+xml",
		Description: "",
	},
	"scld": {
		Subject:     "application/vnd.doremir.scorecloud-binary-document",
		Description: "",
	},
	"scm": {
		Subject:     "application/scim+json",
		Description: "",
	},
	"scq": {
		Subject:     "application/scvp-cv-request",
		Description: "",
	},
	"scr": {
		Subject:     "application/x-silverlight",
		Description: "",
	},
	"scs": {
		Subject:     "application/vnd.etsi.asic-s+zip",
		Description: "",
	},
	"scsf": {
		Subject:     "application/vnd.sealed.csf",
		Description: "",
	},
	"sct": {
		Subject:     "text/x-scriptlet",
		Description: "",
	},
	"scurl": {
		Subject:     "text/vnd.curl.scurl",
		Description: "",
	},
	"sd2": {
		Subject:     "audio/x-sd2",
		Description: "",
	},
	"sd": {
		Subject:     "x-chemical/x-mdl-sdfile",
		Description: "",
	},
	"sda": {
		Subject:     "application/vnd.stardivision.draw",
		Description: "",
	},
	"sdc": {
		Subject:     "application/vnd.stardivision.calc",
		Description: "",
	},
	"sdd": {
		Subject:     "application/vnd.stardivision.impress",
		Description: "",
	},
	"sdf": {
		Subject:     "application/vnd.kinar",
		Description: "",
	},
	"sdkd": {
		Subject:     "application/vnd.solent.sdkm+xml",
		Description: "",
	},
	"sdkm": {
		Subject:     "application/vnd.solent.sdkm+xml",
		Description: "",
	},
	"sdml": {
		Subject:     "text/plain",
		Description: "",
	},
	"sdo": {
		Subject:     "application/vnd.sealed.doc",
		Description: "",
	},
	"sdoc": {
		Subject:     "application/vnd.sealed.doc",
		Description: "",
	},
	"sdp": {
		Subject:     "application/sdp",
		Description: "",
	},
	"sdr": {
		Subject:     "application/x-sounder",
		Description: "",
	},
	"sds": {
		Subject:     "application/vnd.stardivision.chart",
		Description: "",
	},
	"sdw": {
		Subject:     "application/vnd.stardivision.writer",
		Description: "",
	},
	"sea": {
		Subject:     "application/x-sea",
		Description: "",
	},
	"see": {
		Subject:     "application/vnd.seemail",
		Description: "",
	},
	"seed": {
		Subject:     "application/vnd.fdsn.mseed",
		Description: "",
	},
	"sem": {
		Subject:     "application/vnd.sealed.eml",
		Description: "",
	},
	"sema": {
		Subject:     "application/vnd.sema",
		Description: "",
	},
	"semd": {
		Subject:     "application/vnd.semd",
		Description: "",
	},
	"semf": {
		Subject:     "application/vnd.semf",
		Description: "",
	},
	"seml": {
		Subject:     "application/vnd.sealed.eml",
		Description: "",
	},
	"senml-etchc": {
		Subject:     "application/senml-etch+cbor",
		Description: "",
	},
	"senml-etchj": {
		Subject:     "application/senml-etch+json",
		Description: "",
	},
	"senml": {
		Subject:     "application/senml+json",
		Description: "",
	},
	"senmlc": {
		Subject:     "application/senml+cbor",
		Description: "",
	},
	"senmle": {
		Subject:     "application/senml-exi",
		Description: "",
	},
	"senmlx": {
		Subject:     "application/senml+xml",
		Description: "",
	},
	"sensml": {
		Subject:     "application/sensml+json",
		Description: "",
	},
	"sensmlc": {
		Subject:     "application/sensml+cbor",
		Description: "",
	},
	"sensmle": {
		Subject:     "application/sensml-exi",
		Description: "",
	},
	"sensmlx": {
		Subject:     "application/sensml+xml",
		Description: "",
	},
	"ser": {
		Subject:     "application/x-java-serialized-object",
		Description: "",
	},
	"set": {
		Subject:     "application/x-set",
		Description: "",
	},
	"setpay": {
		Subject:     "application/set-payment-initiation",
		Description: "",
	},
	"setreg": {
		Subject:     "application/set-registration-initiation",
		Description: "",
	},
	"sfc": {
		Subject:     "application/vnd.nintendo.snes.rom",
		Description: "",
	},
	"sfd-hdstx": {
		Subject:     "application/vnd.hydrostatix.sof-data",
		Description: "",
	},
	"sfd": {
		Subject:     "application/vnd.font-fontforge-sfd",
		Description: "",
	},
	"sfs": {
		Subject:     "application/vnd.spotfire.sfs",
		Description: "",
	},
	"sfv": {
		Subject:     "text/x-sfv",
		Description: "",
	},
	"sg": {
		Subject:     "application/x-sms-rom",
		Description: "",
	},
	"sgb": {
		Subject:     "application/x-gameboy-rom",
		Description: "",
	},
	"sgf": {
		Subject:     "application/x-go-sgf",
		Description: "",
	},
	"sgi": {
		Subject:     "image/vnd.sealedmedia.softseal.gif",
		Description: "",
	},
	"sgif": {
		Subject:     "image/vnd.sealedmedia.softseal.gif",
		Description: "",
	},
	"sgl": {
		Subject:     "application/vnd.stardivision.writer-global",
		Description: "",
	},
	"sgm": {
		Subject:     "text/sgml",
		Description: "",
	},
	"sgml": {
		Subject:     "text/sgml",
		Description: "",
	},
	"sh": {
		Subject:     "application/x-sh",
		Description: "",
	},
	"shaclc": {
		Subject:     "text/shaclc",
		Description: "",
	},
	"shape": {
		Subject:     "application/x-dia-shape",
		Description: "",
	},
	"shar": {
		Subject:     "application/x-shar",
		Description: "",
	},
	"shc": {
		Subject:     "text/shaclc",
		Description: "",
	},
	"shex": {
		Subject:     "text/shex",
		Description: "",
	},
	"shf": {
		Subject:     "application/shf+xml",
		Description: "",
	},
	"shn": {
		Subject:     "application/x-shorten",
		Description: "",
	},
	"shp": {
		Subject:     "application/vnd.shp",
		Description: "",
	},
	"shtml": {
		Subject:     "text/html",
		Description: "",
	},
	"shx": {
		Subject:     "application/vnd.shx",
		Description: "",
	},
	"si": {
		Subject:     "text/vnd.wap.si",
		Description: "",
	},
	"siag": {
		Subject:     "application/x-siag",
		Description: "",
	},
	"sic": {
		Subject:     "application/vnd.wap.sic",
		Description: "",
	},
	"sid": {
		Subject:     "application/yang-sid+json",
		Description: "",
	},
	"sieve": {
		Subject:     "application/sieve",
		Description: "",
	},
	"sig": {
		Subject:     "application/pgp-signature",
		Description: "",
	},
	"sik": {
		Subject:     "application/x-trash",
		Description: "",
	},
	"sil": {
		Subject:     "audio/x-silk",
		Description: "",
	},
	"silo": {
		Subject:     "model/mesh",
		Description: "",
	},
	"sipa": {
		Subject:     "application/vnd.smintio.portals.archive",
		Description: "",
	},
	"sis": {
		Subject:     "application/vnd.symbian.install",
		Description: "",
	},
	"sisx": {
		Subject:     "x-epoc/x-sisx-app",
		Description: "",
	},
	"sit": {
		Subject:     "application/x-stuffit",
		Description: "",
	},
	"sitx": {
		Subject:     "application/x-stuffit",
		Description: "",
	},
	"siv": {
		Subject:     "application/sieve",
		Description: "",
	},
	"sjp": {
		Subject:     "image/vnd.sealedmedia.softseal.jpg",
		Description: "",
	},
	"sjpg": {
		Subject:     "image/vnd.sealedmedia.softseal.jpg",
		Description: "",
	},
	"sk1": {
		Subject:     "image/x-skencil",
		Description: "",
	},
	"sk": {
		Subject:     "image/x-skencil",
		Description: "",
	},
	"skd": {
		Subject:     "application/x-koan",
		Description: "",
	},
	"skm": {
		Subject:     "application/x-koan",
		Description: "",
	},
	"skp": {
		Subject:     "application/x-koan",
		Description: "",
	},
	"skr": {
		Subject:     "application/pgp-keys",
		Description: "",
	},
	"skt": {
		Subject:     "application/x-koan",
		Description: "",
	},
	"sl": {
		Subject:     "text/vnd.wap.sl",
		Description: "",
	},
	"sla.gz": {
		Subject:     "application/vnd.scribus",
		Description: "",
	},
	"sla": {
		Subject:     "application/vnd.scribus",
		Description: "",
	},
	"slaz": {
		Subject:     "application/vnd.scribus",
		Description: "",
	},
	"slc": {
		Subject:     "application/vnd.wap.slc",
		Description: "",
	},
	"sldm": {
		Subject:     "application/vnd.ms-powerpoint.slide.macroenabled.12",
		Description: "",
	},
	"sldx": {
		Subject:     "application/vnd.openxmlformats-officedocument.presentationml.slide",
		Description: "",
	},
	"slk": {
		Subject:     "text/x-spreadsheet",
		Description: "",
	},
	"sls": {
		Subject:     "application/route-s-tsid+xml",
		Description: "",
	},
	"slt": {
		Subject:     "application/vnd.epson.salt",
		Description: "",
	},
	"sm": {
		Subject:     "application/vnd.stepmania.stepchart",
		Description: "",
	},
	"smaf": {
		Subject:     "application/x-smaf",
		Description: "",
	},
	"smc": {
		Subject:     "application/vnd.nintendo.snes.rom",
		Description: "",
	},
	"smd": {
		Subject:     "application/vnd.stardivision.mail",
		Description: "",
	},
	"smf": {
		Subject:     "application/vnd.stardivision.math",
		Description: "",
	},
	"smh": {
		Subject:     "application/vnd.sealed.mht",
		Description: "",
	},
	"smht": {
		Subject:     "application/vnd.sealed.mht",
		Description: "",
	},
	"smi": {
		Subject:     "application/smil",
		Description: "",
	},
	"smil": {
		Subject:     "application/smil",
		Description: "",
	},
	"smk": {
		Subject:     "video/vnd.radgamettools.smacker",
		Description: "",
	},
	"sml": {
		Subject:     "application/smil",
		Description: "",
	},
	"smo": {
		Subject:     "video/vnd.sealedmedia.softseal.mov",
		Description: "",
	},
	"smov": {
		Subject:     "video/vnd.sealedmedia.softseal.mov",
		Description: "",
	},
	"smp3": {
		Subject:     "audio/vnd.sealedmedia.softseal.mpeg",
		Description: "",
	},
	"smp": {
		Subject:     "audio/vnd.sealedmedia.softseal.mpeg",
		Description: "",
	},
	"smpg": {
		Subject:     "video/vnd.sealed.mpeg1",
		Description: "",
	},
	"sms": {
		Subject:     "application/vnd.3gpp.sms",
		Description: "",
	},
	"smv": {
		Subject:     "audio/smv",
		Description: "",
	},
	"smzip": {
		Subject:     "application/vnd.stepmania.package",
		Description: "",
	},
	"snd": {
		Subject:     "audio/basic",
		Description: "",
	},
	"snf": {
		Subject:     "application/x-font-snf",
		Description: "",
	},
	"so": {
		Subject:     "application/x-sharedlib",
		Description: "",
	},
	"soa": {
		Subject:     "text/dns",
		Description: "",
	},
	"soc": {
		Subject:     "application/sgml-open-catalog",
		Description: "",
	},
	"sofa": {
		Subject:     "audio/sofa",
		Description: "",
	},
	"sol": {
		Subject:     "application/x-solids",
		Description: "",
	},
	"sos": {
		Subject:     "text/vnd.sosi",
		Description: "",
	},
	"spc": {
		Subject:     "x-chemical/x-galactic-spc",
		Description: "",
	},
	"spd": {
		Subject:     "application/vnd.sealedmedia.softseal.pdf",
		Description: "",
	},
	"spdf": {
		Subject:     "application/vnd.sealedmedia.softseal.pdf",
		Description: "",
	},
	"spdx.json": {
		Subject:     "application/spdx+json",
		Description: "",
	},
	"spdx": {
		Subject:     "text/spdx",
		Description: "",
	},
	"spec": {
		Subject:     "text/x-rpm-spec",
		Description: "",
	},
	"spf": {
		Subject:     "application/vnd.yamaha.smaf-phrase",
		Description: "",
	},
	"spl": {
		Subject:     "application/x-futuresplash",
		Description: "",
	},
	"spm": {
		Subject:     "application/x-source-rpm",
		Description: "",
	},
	"spn": {
		Subject:     "image/vnd.sealed.png",
		Description: "",
	},
	"spng": {
		Subject:     "image/vnd.sealed.png",
		Description: "",
	},
	"spo": {
		Subject:     "text/vnd.in3d.spot",
		Description: "",
	},
	"spot": {
		Subject:     "text/vnd.in3d.spot",
		Description: "",
	},
	"spp": {
		Subject:     "application/vnd.sealed.ppt",
		Description: "",
	},
	"sppt": {
		Subject:     "application/vnd.sealed.ppt",
		Description: "",
	},
	"spq": {
		Subject:     "application/scvp-vp-request",
		Description: "",
	},
	"spr": {
		Subject:     "application/x-sprite",
		Description: "",
	},
	"sprite": {
		Subject:     "application/x-sprite",
		Description: "",
	},
	"sps": {
		Subject:     "application/x-spss",
		Description: "",
	},
	"spx": {
		Subject:     "audio/ogg",
		Description: "",
	},
	"sql": {
		Subject:     "application/sql",
		Description: "",
	},
	"sr2": {
		Subject:     "image/x-sony-sr2",
		Description: "",
	},
	"sr": {
		Subject:     "application/vnd.sigrok.session",
		Description: "",
	},
	"src": {
		Subject:     "application/x-wais-source",
		Description: "",
	},
	"srf": {
		Subject:     "image/x-sony-srf",
		Description: "",
	},
	"srt": {
		Subject:     "application/x-subrip",
		Description: "",
	},
	"sru": {
		Subject:     "application/sru+xml",
		Description: "",
	},
	"srx": {
		Subject:     "application/sparql-results+xml",
		Description: "",
	},
	"ss": {
		Subject:     "text/x-scheme",
		Description: "",
	},
	"ssa": {
		Subject:     "text/x-ssa",
		Description: "",
	},
	"ssdl": {
		Subject:     "application/x-ssdl+xml",
		Description: "",
	},
	"sse": {
		Subject:     "application/vnd.kodak-descriptor",
		Description: "",
	},
	"ssf": {
		Subject:     "application/vnd.epson.ssf",
		Description: "",
	},
	"ssi": {
		Subject:     "text/x-server-parsed-html",
		Description: "",
	},
	"ssm": {
		Subject:     "application/x-streamingmedia",
		Description: "",
	},
	"ssml": {
		Subject:     "application/ssml+xml",
		Description: "",
	},
	"sst": {
		Subject:     "application/vnd.ms-pki.certstore",
		Description: "",
	},
	"ssv": {
		Subject:     "application/vnd.crypto-shade-file",
		Description: "",
	},
	"ssvc": {
		Subject:     "application/vnd.crypto-shade-file",
		Description: "",
	},
	"ssw": {
		Subject:     "video/vnd.sealed.swf",
		Description: "",
	},
	"sswf": {
		Subject:     "video/vnd.sealed.swf",
		Description: "",
	},
	"st": {
		Subject:     "application/vnd.sailingtracker.track",
		Description: "",
	},
	"stc": {
		Subject:     "application/vnd.sun.xml.calc.template",
		Description: "",
	},
	"std": {
		Subject:     "application/vnd.sun.xml.draw.template",
		Description: "",
	},
	"step": {
		Subject:     "model/step",
		Description: "",
	},
	"stf": {
		Subject:     "application/vnd.wt.stf",
		Description: "",
	},
	"sti": {
		Subject:     "application/vnd.sun.xml.impress.template",
		Description: "",
	},
	"stif": {
		Subject:     "application/vnd.sealed.tiff",
		Description: "",
	},
	"stix": {
		Subject:     "application/stix+json",
		Description: "",
	},
	"stk": {
		Subject:     "application/hyperstudio",
		Description: "",
	},
	"stl": {
		Subject:     "model/stl",
		Description: "",
	},
	"stm": {
		Subject:     "application/vnd.sealedmedia.softseal.html",
		Description: "",
	},
	"stml": {
		Subject:     "application/vnd.sealedmedia.softseal.html",
		Description: "",
	},
	"stp": {
		Subject:     "application/p21",
		Description: "",
	},
	"stpnc": {
		Subject:     "application/p21",
		Description: "",
	},
	"stpx": {
		Subject:     "model/step+xml",
		Description: "",
	},
	"stpxz": {
		Subject:     "model/step-xml+zip",
		Description: "",
	},
	"stpz": {
		Subject:     "application/p21+zip",
		Description: "",
	},
	"str": {
		Subject:     "application/vnd.pg.format",
		Description: "",
	},
	"study-inter": {
		Subject:     "application/vnd.vd-study",
		Description: "",
	},
	"stw": {
		Subject:     "application/vnd.sun.xml.writer.template",
		Description: "",
	},
	"sty": {
		Subject:     "text/x-tex",
		Description: "",
	},
	"sub": {
		Subject:     "image/vnd.dvb.subtitle",
		Description: "",
	},
	"sun": {
		Subject:     "image/x-sun-raster",
		Description: "",
	},
	"sus": {
		Subject:     "application/vnd.sus-calendar",
		Description: "",
	},
	"susp": {
		Subject:     "application/vnd.sus-calendar",
		Description: "",
	},
	"sv4cpio": {
		Subject:     "application/x-sv4cpio",
		Description: "",
	},
	"sv4crc": {
		Subject:     "application/x-sv4crc",
		Description: "",
	},
	"sv": {
		Subject:     "text/x-svsrc",
		Description: "",
	},
	"svc": {
		Subject:     "application/vnd.dvb.service",
		Description: "",
	},
	"svd": {
		Subject:     "application/vnd.svd",
		Description: "",
	},
	"svf": {
		Subject:     "image/vnd.dwg",
		Description: "",
	},
	"svg": {
		Subject:     "image/svg+xml",
		Description: "",
	},
	"svgz": {
		Subject:     "image/svg+xml",
		Description: "",
	},
	"svh": {
		Subject:     "text/x-svhdr",
		Description: "",
	},
	"svr": {
		Subject:     "application/x-world",
		Description: "",
	},
	"sw": {
		Subject:     "x-chemical/x-swissprot",
		Description: "",
	},
	"swa": {
		Subject:     "application/x-director",
		Description: "",
	},
	"swf": {
		Subject:     "application/vnd.adobe.flash.movie",
		Description: "",
	},
	"swfl": {
		Subject:     "application/x-shockwave-flash",
		Description: "",
	},
	"swi": {
		Subject:     "application/vnd.arastra.swi",
		Description: "",
	},
	"swidtag": {
		Subject:     "application/swid+xml",
		Description: "",
	},
	"swm": {
		Subject:     "application/x-ms-wim",
		Description: "",
	},
	"sxc": {
		Subject:     "application/vnd.sun.xml.calc",
		Description: "",
	},
	"sxd": {
		Subject:     "application/vnd.sun.xml.draw",
		Description: "",
	},
	"sxg": {
		Subject:     "application/vnd.sun.xml.writer.global",
		Description: "",
	},
	"sxi": {
		Subject:     "application/vnd.vd-study",
		Description: "",
	},
	"sxl": {
		Subject:     "application/vnd.sealed.xls",
		Description: "",
	},
	"sxls": {
		Subject:     "application/vnd.sealed.xls",
		Description: "",
	},
	"sxm": {
		Subject:     "application/vnd.sun.xml.math",
		Description: "",
	},
	"sxw": {
		Subject:     "application/vnd.sun.xml.writer",
		Description: "",
	},
	"sy2": {
		Subject:     "application/vnd.sybyl.mol2",
		Description: "",
	},
	"syft.json": {
		Subject:     "application/vnd.syft+json",
		Description: "",
	},
	"sylk": {
		Subject:     "text/x-spreadsheet",
		Description: "",
	},
	"t2t": {
		Subject:     "text/x-txt2tags",
		Description: "",
	},
	"t3": {
		Subject:     "application/x-t3vm-image",
		Description: "",
	},
	"t": {
		Subject:     "application/x-troff",
		Description: "",
	},
	"tag": {
		Subject:     "text/prs.lines.tag",
		Description: "",
	},
	"taglet": {
		Subject:     "application/vnd.mynfc",
		Description: "",
	},
	"talk": {
		Subject:     "text/x-speech",
		Description: "",
	},
	"tam": {
		Subject:     "application/vnd.onepager",
		Description: "",
	},
	"tamp": {
		Subject:     "application/vnd.onepagertamp",
		Description: "",
	},
	"tamx": {
		Subject:     "application/vnd.onepagertamx",
		Description: "",
	},
	"tao": {
		Subject:     "application/vnd.tao.intent-module-archive",
		Description: "",
	},
	"tap": {
		Subject:     "image/vnd.tencent.tap",
		Description: "",
	},
	"tar": {
		Subject:     "application/x-tar",
		Description: "",
	},
	"tat": {
		Subject:     "application/vnd.onepagertat",
		Description: "",
	},
	"tatp": {
		Subject:     "application/vnd.onepagertatp",
		Description: "",
	},
	"tatx": {
		Subject:     "application/vnd.onepagertatx",
		Description: "",
	},
	"tau": {
		Subject:     "application/tamp-apex-update",
		Description: "",
	},
	"taz": {
		Subject:     "application/x-gtar",
		Description: "",
	},
	"tb2": {
		Subject:     "application/x-bzip-compressed-tar",
		Description: "",
	},
	"tbk": {
		Subject:     "application/x-toolbook",
		Description: "",
	},
	"tbz2": {
		Subject:     "application/x-gtar",
		Description: "",
	},
	"tbz": {
		Subject:     "application/x-gtar",
		Description: "",
	},
	"tcap": {
		Subject:     "application/vnd.3gpp2.tcap",
		Description: "",
	},
	"tcl": {
		Subject:     "application/x-tcl",
		Description: "",
	},
	"tcsh": {
		Subject:     "text/x-script.tcsh",
		Description: "",
	},
	"tcu": {
		Subject:     "application/tamp-community-update",
		Description: "",
	},
	"td": {
		Subject:     "application/urc-targetdesc+xml",
		Description: "",
	},
	"teacher": {
		Subject:     "application/vnd.smart.teacher",
		Description: "",
	},
	"tei": {
		Subject:     "application/tei+xml",
		Description: "",
	},
	"teicorpus": {
		Subject:     "application/tei+xml",
		Description: "",
	},
	"ter": {
		Subject:     "application/tamp-error",
		Description: "",
	},
	"tex": {
		Subject:     "application/x-tex",
		Description: "",
	},
	"texi": {
		Subject:     "text/prs.texi",
		Description: "",
	},
	"texinfo": {
		Subject:     "application/x-texinfo",
		Description: "",
	},
	"text": {
		Subject:     "application/x-plain",
		Description: "",
	},
	"tfi": {
		Subject:     "application/thraud+xml",
		Description: "",
	},
	"tfm": {
		Subject:     "application/x-tex-tfm",
		Description: "",
	},
	"tga": {
		Subject:     "image/x-targa",
		Description: "",
	},
	"tgf": {
		Subject:     "x-chemical/x-mdl-tgf",
		Description: "",
	},
	"tgz": {
		Subject:     "application/x-gtar",
		Description: "",
	},
	"theme": {
		Subject:     "application/x-theme",
		Description: "",
	},
	"themepack": {
		Subject:     "application/x-windows-themepack",
		Description: "",
	},
	"thmx": {
		Subject:     "application/vnd.ms-officetheme",
		Description: "",
	},
	"tif": {
		Subject:     "image/tiff",
		Description: "",
	},
	"tiff": {
		Subject:     "image/tiff",
		Description: "",
	},
	"tk": {
		Subject:     "text/x-tcl",
		Description: "",
	},
	"tlclient": {
		Subject:     "application/vnd.cendio.thinlinc.clientconf",
		Description: "",
	},
	"tlrz": {
		Subject:     "application/x-lrzip-compressed-tar",
		Description: "",
	},
	"tlz": {
		Subject:     "application/x-lzma-compressed-tar",
		Description: "",
	},
	"tm.json": {
		Subject:     "application/tm+json",
		Description: "",
	},
	"tm.jsonld": {
		Subject:     "application/tm+json",
		Description: "",
	},
	"tm": {
		Subject:     "text/x-texmacs",
		Description: "",
	},
	"tmo": {
		Subject:     "application/vnd.tmobile-livetv",
		Description: "",
	},
	"tnef": {
		Subject:     "application/vnd.ms-tnef",
		Description: "",
	},
	"tnf": {
		Subject:     "application/vnd.ms-tnef",
		Description: "",
	},
	"toc": {
		Subject:     "application/x-cdrdao-toc",
		Description: "",
	},
	"torrent": {
		Subject:     "application/x-bittorrent",
		Description: "",
	},
	"tpic": {
		Subject:     "image/x-tga",
		Description: "",
	},
	"tpl": {
		Subject:     "application/vnd.groove-tool-template",
		Description: "",
	},
	"tpt": {
		Subject:     "application/vnd.trid.tpt",
		Description: "",
	},
	"tr": {
		Subject:     "application/x-troff",
		Description: "",
	},
	"tra": {
		Subject:     "application/vnd.trueapp",
		Description: "",
	},
	"tree": {
		Subject:     "application/vnd.rainstor.data",
		Description: "",
	},
	"trig": {
		Subject:     "application/trig",
		Description: "",
	},
	"trm": {
		Subject:     "application/x-msterminal",
		Description: "",
	},
	"ts": {
		Subject:     "text/vnd.trolltech.linguist",
		Description: "",
	},
	"tsa": {
		Subject:     "application/tamp-sequence-adjust",
		Description: "",
	},
	"tsd": {
		Subject:     "application/timestamped-data",
		Description: "",
	},
	"tsi": {
		Subject:     "audio/x-tsp-audio",
		Description: "",
	},
	"tsp": {
		Subject:     "application/x-dsptype",
		Description: "",
	},
	"tsq": {
		Subject:     "application/tamp-status-query",
		Description: "",
	},
	"tsr": {
		Subject:     "application/tamp-status-response",
		Description: "",
	},
	"tst": {
		Subject:     "application/vnd.etsi.timestamp-token",
		Description: "",
	},
	"tsv": {
		Subject:     "text/tab-separated-values",
		Description: "",
	},
	"tta": {
		Subject:     "audio/x-tta",
		Description: "",
	},
	"ttc": {
		Subject:     "font/collection",
		Description: "",
	},
	"ttf": {
		Subject:     "font/ttf",
		Description: "",
	},
	"ttl": {
		Subject:     "text/turtle",
		Description: "",
	},
	"ttml": {
		Subject:     "application/ttml+xml",
		Description: "",
	},
	"ttx": {
		Subject:     "application/x-font-ttx",
		Description: "",
	},
	"tuc": {
		Subject:     "application/tamp-update-confirm",
		Description: "",
	},
	"tur": {
		Subject:     "application/tamp-update",
		Description: "",
	},
	"turbot": {
		Subject:     "image/x-florian",
		Description: "",
	},
	"twd": {
		Subject:     "application/vnd.simtech-mindmapper",
		Description: "",
	},
	"twds": {
		Subject:     "application/vnd.simtech-mindmapper",
		Description: "",
	},
	"txd": {
		Subject:     "application/vnd.genomatix.tuxedo",
		Description: "",
	},
	"txf": {
		Subject:     "application/vnd.mobius.txf",
		Description: "",
	},
	"txt": {
		Subject:     "text/plain",
		Description: "",
	},
	"txz": {
		Subject:     "application/x-xz-compressed-tar",
		Description: "",
	},
	"tzo": {
		Subject:     "application/x-tzo",
		Description: "",
	},
	"u32": {
		Subject:     "application/x-authorware-bin",
		Description: "",
	},
	"u3d": {
		Subject:     "model/u3d",
		Description: "",
	},
	"u8dsn": {
		Subject:     "message/global-delivery-status",
		Description: "",
	},
	"u8mdn": {
		Subject:     "message/global-disposition-notification",
		Description: "",
	},
	"u8msg": {
		Subject:     "message/global",
		Description: "",
	},
	"udeb": {
		Subject:     "application/vnd.debian.binary-package",
		Description: "",
	},
	"ufd": {
		Subject:     "application/vnd.ufdl",
		Description: "",
	},
	"ufdl": {
		Subject:     "application/vnd.ufdl",
		Description: "",
	},
	"ufraw": {
		Subject:     "application/x-ufraw",
		Description: "",
	},
	"ui": {
		Subject:     "application/x-designer",
		Description: "",
	},
	"uil": {
		Subject:     "text/x-uil",
		Description: "",
	},
	"uis": {
		Subject:     "application/urc-uisocketdesc+xml",
		Description: "",
	},
	"uls": {
		Subject:     "text/x-iuls",
		Description: "",
	},
	"ult": {
		Subject:     "audio/x-mod",
		Description: "",
	},
	"ulx": {
		Subject:     "application/x-glulx",
		Description: "",
	},
	"umj": {
		Subject:     "application/vnd.umajin",
		Description: "",
	},
	"unf": {
		Subject:     "application/x-nes-rom",
		Description: "",
	},
	"uni": {
		Subject:     "text/uri-list",
		Description: "",
	},
	"unif": {
		Subject:     "application/x-nes-rom",
		Description: "",
	},
	"unis": {
		Subject:     "text/uri-list",
		Description: "",
	},
	"unityweb": {
		Subject:     "application/vnd.unity",
		Description: "",
	},
	"unv": {
		Subject:     "application/x-i-deas",
		Description: "",
	},
	"uo": {
		Subject:     "application/vnd.uoml+xml",
		Description: "",
	},
	"uoml": {
		Subject:     "application/vnd.uoml+xml",
		Description: "",
	},
	"upa": {
		Subject:     "application/vnd.fints",
		Description: "",
	},
	"uri": {
		Subject:     "text/uri-list",
		Description: "",
	},
	"uric": {
		Subject:     "text/vnd.si.uricatalogue",
		Description: "",
	},
	"urim": {
		Subject:     "application/vnd.uri-map",
		Description: "",
	},
	"urimap": {
		Subject:     "application/vnd.uri-map",
		Description: "",
	},
	"uris": {
		Subject:     "text/uri-list",
		Description: "",
	},
	"url": {
		Subject:     "application/x-mswinurl",
		Description: "",
	},
	"urls": {
		Subject:     "text/uri-list",
		Description: "",
	},
	"usda": {
		Subject:     "model/vnd.usda",
		Description: "",
	},
	"usdz": {
		Subject:     "model/vnd.usdz+zip",
		Description: "",
	},
	"ustar": {
		Subject:     "application/x-ustar",
		Description: "",
	},
	"utz": {
		Subject:     "application/vnd.uiq.theme",
		Description: "",
	},
	"uu": {
		Subject:     "text/x-uuencode",
		Description: "",
	},
	"uue": {
		Subject:     "text/x-uuencode",
		Description: "",
	},
	"uva": {
		Subject:     "audio/vnd.dece.audio",
		Description: "",
	},
	"uvd": {
		Subject:     "application/vnd.dece.data",
		Description: "",
	},
	"uvf": {
		Subject:     "application/vnd.dece.data",
		Description: "",
	},
	"uvg": {
		Subject:     "image/vnd.dece.graphic",
		Description: "",
	},
	"uvh": {
		Subject:     "video/vnd.dece.hd",
		Description: "",
	},
	"uvi": {
		Subject:     "image/vnd.dece.graphic",
		Description: "",
	},
	"uvm": {
		Subject:     "video/vnd.dece.mobile",
		Description: "",
	},
	"uvp": {
		Subject:     "video/vnd.dece.pd",
		Description: "",
	},
	"uvs": {
		Subject:     "video/vnd.dece.sd",
		Description: "",
	},
	"uvt": {
		Subject:     "application/vnd.dece.ttml+xml",
		Description: "",
	},
	"uvu": {
		Subject:     "video/vnd.dece.mp4",
		Description: "",
	},
	"uvv": {
		Subject:     "video/vnd.dece.video",
		Description: "",
	},
	"uvva": {
		Subject:     "audio/vnd.dece.audio",
		Description: "",
	},
	"uvvd": {
		Subject:     "application/vnd.dece.data",
		Description: "",
	},
	"uvvf": {
		Subject:     "application/vnd.dece.data",
		Description: "",
	},
	"uvvg": {
		Subject:     "image/vnd.dece.graphic",
		Description: "",
	},
	"uvvh": {
		Subject:     "video/vnd.dece.hd",
		Description: "",
	},
	"uvvi": {
		Subject:     "image/vnd.dece.graphic",
		Description: "",
	},
	"uvvm": {
		Subject:     "video/vnd.dece.mobile",
		Description: "",
	},
	"uvvp": {
		Subject:     "video/vnd.dece.pd",
		Description: "",
	},
	"uvvs": {
		Subject:     "video/vnd.dece.sd",
		Description: "",
	},
	"uvvt": {
		Subject:     "application/vnd.dece.ttml+xml",
		Description: "",
	},
	"uvvu": {
		Subject:     "video/vnd.dece.mp4",
		Description: "",
	},
	"uvvv": {
		Subject:     "video/vnd.dece.video",
		Description: "",
	},
	"uvvx": {
		Subject:     "application/vnd.dece.unspecified",
		Description: "",
	},
	"uvvz": {
		Subject:     "application/vnd.dece.zip",
		Description: "",
	},
	"uvx": {
		Subject:     "application/vnd.dece.unspecified",
		Description: "",
	},
	"uvz": {
		Subject:     "application/vnd.dece.zip",
		Description: "",
	},
	"v64": {
		Subject:     "application/x-n64-rom",
		Description: "",
	},
	"v": {
		Subject:     "text/x-verilog",
		Description: "",
	},
	"val": {
		Subject:     "x-chemical/x-ncbi-asn1-binary",
		Description: "",
	},
	"vala": {
		Subject:     "text/x-vala",
		Description: "",
	},
	"vapi": {
		Subject:     "text/x-vala",
		Description: "",
	},
	"vbk": {
		Subject:     "audio/vnd.nortel.vbk",
		Description: "",
	},
	"vcard": {
		Subject:     "text/vcard",
		Description: "",
	},
	"vcd": {
		Subject:     "application/x-cdlink",
		Description: "",
	},
	"vcf.bz2": {
		Subject:     "application/prs.vcfbzip2",
		Description: "",
	},
	"vcf": {
		Subject:     "text/vcard",
		Description: "",
	},
	"vcfbzip2": {
		Subject:     "application/prs.vcfbzip2",
		Description: "",
	},
	"vcg": {
		Subject:     "application/vnd.groove-vcard",
		Description: "",
	},
	"vcj": {
		Subject:     "application/voucher-cms+json",
		Description: "",
	},
	"vcs": {
		Subject:     "text/x-vcalendar",
		Description: "",
	},
	"vct": {
		Subject:     "text/vcard",
		Description: "",
	},
	"vcx": {
		Subject:     "application/vnd.vcx",
		Description: "",
	},
	"vd": {
		Subject:     "application/vividence.scriptfile",
		Description: "",
	},
	"vda": {
		Subject:     "application/x-vda",
		Description: "",
	},
	"vdo": {
		Subject:     "video/x-vdo",
		Description: "",
	},
	"vds": {
		Subject:     "model/vnd.sap.vds",
		Description: "",
	},
	"vew": {
		Subject:     "application/x-groupwise",
		Description: "",
	},
	"vfr": {
		Subject:     "application/vnd.tml",
		Description: "",
	},
	"vhd": {
		Subject:     "text/x-vhdl",
		Description: "",
	},
	"vhdl": {
		Subject:     "text/x-vhdl",
		Description: "",
	},
	"viaframe": {
		Subject:     "application/vnd.tml",
		Description: "",
	},
	"vis": {
		Subject:     "application/vnd.informix-visionary",
		Description: "",
	},
	"viv": {
		Subject:     "video/vnd.vivo",
		Description: "",
	},
	"vivo": {
		Subject:     "video/vnd.vivo",
		Description: "",
	},
	"vlc": {
		Subject:     "audio/x-mpegurl",
		Description: "",
	},
	"vmd": {
		Subject:     "x-chemical/x-vmd",
		Description: "",
	},
	"vmf": {
		Subject:     "application/x-vocaltec-media-file",
		Description: "",
	},
	"vms": {
		Subject:     "x-chemical/x-vamas-iso14976",
		Description: "",
	},
	"vmt": {
		Subject:     "application/vnd.valve.source.material",
		Description: "",
	},
	"vob": {
		Subject:     "video/x-ms-vob",
		Description: "",
	},
	"voc": {
		Subject:     "audio/x-voc",
		Description: "",
	},
	"vor": {
		Subject:     "application/vnd.stardivision.writer",
		Description: "",
	},
	"vos": {
		Subject:     "video/x-vosaic",
		Description: "",
	},
	"vox": {
		Subject:     "application/x-authorware-bin",
		Description: "",
	},
	"vqe": {
		Subject:     "audio/x-twinvq-plugin",
		Description: "",
	},
	"vqf": {
		Subject:     "audio/x-twinvq",
		Description: "",
	},
	"vql": {
		Subject:     "audio/x-twinvq-plugin",
		Description: "",
	},
	"vrm": {
		Subject:     "x-world/x-vrml",
		Description: "",
	},
	"vrml": {
		Subject:     "model/vrml",
		Description: "",
	},
	"vrt": {
		Subject:     "x-world/x-vrt",
		Description: "",
	},
	"vsc": {
		Subject:     "application/vnd.vidsoft.vidconference",
		Description: "",
	},
	"vsd": {
		Subject:     "application/vnd.visio",
		Description: "",
	},
	"vsf": {
		Subject:     "application/vividence.scriptfile",
		Description: "",
	},
	"vss": {
		Subject:     "application/vnd.visio",
		Description: "",
	},
	"vst": {
		Subject:     "application/vnd.visio",
		Description: "",
	},
	"vsw": {
		Subject:     "application/vnd.visio",
		Description: "",
	},
	"vtd": {
		Subject:     "application/vividence.scriptfile",
		Description: "",
	},
	"vtf": {
		Subject:     "image/vnd.valve.source.texture",
		Description: "",
	},
	"vtnstd": {
		Subject:     "application/vnd.veritone.aion+json",
		Description: "",
	},
	"vtt": {
		Subject:     "text/vtt",
		Description: "",
	},
	"vtu": {
		Subject:     "model/vnd.vtu",
		Description: "",
	},
	"vxml": {
		Subject:     "application/voicexml+xml",
		Description: "",
	},
	"w3d": {
		Subject:     "application/x-director",
		Description: "",
	},
	"w60": {
		Subject:     "application/x-wordperfect6.0",
		Description: "",
	},
	"w61": {
		Subject:     "application/x-wordperfect6.1",
		Description: "",
	},
	"w6w": {
		Subject:     "application/msword",
		Description: "",
	},
	"wad": {
		Subject:     "application/x-doom",
		Description: "",
	},
	"wadl": {
		Subject:     "application/vnd.sun.wadl+xml",
		Description: "",
	},
	"wafl": {
		Subject:     "application/vnd.wasmflow.wafl",
		Description: "",
	},
	"wasm": {
		Subject:     "application/wasm",
		Description: "",
	},
	"wav": {
		Subject:     "audio/vnd.dts",
		Description: "",
	},
	"wax": {
		Subject:     "audio/x-ms-wax",
		Description: "",
	},
	"wb1": {
		Subject:     "application/x-qpro",
		Description: "",
	},
	"wb2": {
		Subject:     "application/x-quattropro",
		Description: "",
	},
	"wb3": {
		Subject:     "application/x-quattropro",
		Description: "",
	},
	"wbmp": {
		Subject:     "image/vnd.wap.wbmp",
		Description: "",
	},
	"wbs": {
		Subject:     "application/vnd.criticaltools.wbs+xml",
		Description: "",
	},
	"wbxml": {
		Subject:     "application/vnd.wap.sic",
		Description: "",
	},
	"wcm": {
		Subject:     "application/vnd.ms-works",
		Description: "",
	},
	"wdb": {
		Subject:     "application/vnd.ms-works",
		Description: "",
	},
	"wdp": {
		Subject:     "image/vnd.ms-photo",
		Description: "",
	},
	"web": {
		Subject:     "application/vnd.xara",
		Description: "",
	},
	"weba": {
		Subject:     "audio/webm",
		Description: "",
	},
	"webapp": {
		Subject:     "application/x-web-app-manifest+json",
		Description: "",
	},
	"webm": {
		Subject:     "video/webm",
		Description: "",
	},
	"webmanifest": {
		Subject:     "application/manifest+json",
		Description: "",
	},
	"webp": {
		Subject:     "image/webp",
		Description: "",
	},
	"wg": {
		Subject:     "application/vnd.pmi.widget",
		Description: "",
	},
	"wgsl": {
		Subject:     "text/wgsl",
		Description: "",
	},
	"wgt": {
		Subject:     "application/widget",
		Description: "",
	},
	"wif": {
		Subject:     "application/watcherinfo+xml",
		Description: "",
	},
	"wim": {
		Subject:     "application/x-ms-wim",
		Description: "",
	},
	"win": {
		Subject:     "model/vnd.gs-gdl",
		Description: "",
	},
	"wiz": {
		Subject:     "application/msword",
		Description: "",
	},
	"wk1": {
		Subject:     "application/x-123",
		Description: "",
	},
	"wk3": {
		Subject:     "application/vnd.lotus-1-2-3",
		Description: "",
	},
	"wk4": {
		Subject:     "application/vnd.lotus-1-2-3",
		Description: "",
	},
	"wk": {
		Subject:     "application/x-123",
		Description: "",
	},
	"wkdownload": {
		Subject:     "application/x-partial-download",
		Description: "",
	},
	"wks": {
		Subject:     "application/vnd.lotus-1-2-3",
		Description: "",
	},
	"wlnk": {
		Subject:     "application/link-format",
		Description: "",
	},
	"wm": {
		Subject:     "video/x-ms-wm",
		Description: "",
	},
	"wma": {
		Subject:     "audio/x-ms-wma",
		Description: "",
	},
	"wmc": {
		Subject:     "application/vnd.wmc",
		Description: "",
	},
	"wmd": {
		Subject:     "application/x-ms-wmd",
		Description: "",
	},
	"wmf": {
		Subject:     "image/wmf",
		Description: "",
	},
	"wml": {
		Subject:     "text/vnd.wap.wml",
		Description: "",
	},
	"wmlc": {
		Subject:     "application/vnd.wap.wmlc",
		Description: "",
	},
	"wmls": {
		Subject:     "text/vnd.wap.wmlscript",
		Description: "",
	},
	"wmlsc": {
		Subject:     "application/vnd.wap.wmlscriptc",
		Description: "",
	},
	"wmv": {
		Subject:     "video/x-ms-wmv",
		Description: "",
	},
	"wmx": {
		Subject:     "video/x-ms-wmx",
		Description: "",
	},
	"wmz": {
		Subject:     "application/x-ms-wmz",
		Description: "",
	},
	"woff2": {
		Subject:     "font/woff2",
		Description: "",
	},
	"woff": {
		Subject:     "font/woff",
		Description: "",
	},
	"word": {
		Subject:     "application/msword",
		Description: "",
	},
	"wp4": {
		Subject:     "application/vnd.wordperfect",
		Description: "",
	},
	"wp5": {
		Subject:     "application/vnd.wordperfect5.1",
		Description: "",
	},
	"wp6": {
		Subject:     "application/x-wordperfect6.1",
		Description: "",
	},
	"wp": {
		Subject:     "application/wordperfect5.1",
		Description: "",
	},
	"wpd": {
		Subject:     "application/vnd.wordperfect",
		Description: "",
	},
	"wpg": {
		Subject:     "application/x-wpg",
		Description: "",
	},
	"wpl": {
		Subject:     "application/vnd.ms-wpl",
		Description: "",
	},
	"wpp": {
		Subject:     "application/vnd.wordperfect",
		Description: "",
	},
	"wps": {
		Subject:     "application/vnd.ms-works",
		Description: "",
	},
	"wq1": {
		Subject:     "application/x-lotus",
		Description: "",
	},
	"wqd": {
		Subject:     "application/vnd.wqd",
		Description: "",
	},
	"wrd": {
		Subject:     "application/x-msword",
		Description: "",
	},
	"wri": {
		Subject:     "application/x-mswrite",
		Description: "",
	},
	"wrl": {
		Subject:     "model/vrml",
		Description: "",
	},
	"wrz": {
		Subject:     "model/vrml",
		Description: "",
	},
	"wsc": {
		Subject:     "text/x-scriptlet",
		Description: "",
	},
	"wsdl": {
		Subject:     "application/wsdl+xml",
		Description: "",
	},
	"wsgi": {
		Subject:     "text/x-python",
		Description: "",
	},
	"wspolicy": {
		Subject:     "application/wspolicy+xml",
		Description: "",
	},
	"wsrc": {
		Subject:     "application/x-wais-source",
		Description: "",
	},
	"wtb": {
		Subject:     "application/vnd.webturbo",
		Description: "",
	},
	"wtk": {
		Subject:     "application/x-wintalk",
		Description: "",
	},
	"wv": {
		Subject:     "application/vnd.wv.csp+wbxml",
		Description: "",
	},
	"wvc": {
		Subject:     "audio/x-wavpack-correction",
		Description: "",
	},
	"wvp": {
		Subject:     "audio/x-wavpack",
		Description: "",
	},
	"wvx": {
		Subject:     "video/x-ms-wvx",
		Description: "",
	},
	"wwf": {
		Subject:     "application/x-wwf",
		Description: "",
	},
	"wz": {
		Subject:     "application/x-wingz",
		Description: "",
	},
	"x-png": {
		Subject:     "image/png",
		Description: "",
	},
	"x32": {
		Subject:     "application/x-authorware-bin",
		Description: "",
	},
	"x3d": {
		Subject:     "application/vnd.hzn-3d-crossword",
		Description: "",
	},
	"x3db": {
		Subject:     "model/x3d+fastinfoset",
		Description: "",
	},
	"x3dbz": {
		Subject:     "model/x-x3d+binary",
		Description: "",
	},
	"x3dv": {
		Subject:     "model/x3d-vrml",
		Description: "",
	},
	"x3dvz": {
		Subject:     "model/x-x3d+vrml",
		Description: "",
	},
	"x3dz": {
		Subject:     "model/x3d+xml",
		Description: "",
	},
	"x3f": {
		Subject:     "image/x-sigma-x3f",
		Description: "",
	},
	"x_b": {
		Subject:     "model/vnd.parasolid.transmit.binary",
		Description: "",
	},
	"x_t": {
		Subject:     "model/vnd.parasolid.transmit.text",
		Description: "",
	},
	"xac": {
		Subject:     "application/x-gnucash",
		Description: "",
	},
	"xaml": {
		Subject:     "application/x-xaml+xml",
		Description: "",
	},
	"xap": {
		Subject:     "application/x-silverlight-app",
		Description: "",
	},
	"xar": {
		Subject:     "application/vnd.xara",
		Description: "",
	},
	"xav": {
		Subject:     "application/xcap-att+xml",
		Description: "",
	},
	"xbap": {
		Subject:     "application/x-ms-xbap",
		Description: "",
	},
	"xbd": {
		Subject:     "application/vnd.fujifilm.fb.docuworks.binder",
		Description: "",
	},
	"xbel": {
		Subject:     "application/x-xbel",
		Description: "",
	},
	"xbl": {
		Subject:     "application/xml",
		Description: "",
	},
	"xbm": {
		Subject:     "image/x-xbitmap",
		Description: "",
	},
	"xca": {
		Subject:     "application/xcap-caps+xml",
		Description: "",
	},
	"xcf": {
		Subject:     "application/x-xcf",
		Description: "",
	},
	"xcs": {
		Subject:     "application/calendar+xml",
		Description: "",
	},
	"xct": {
		Subject:     "application/vnd.fujifilm.fb.docuworks.container",
		Description: "",
	},
	"xdcf": {
		Subject:     "application/vnd.gov.sk.xmldatacontainer+xml",
		Description: "",
	},
	"xdd": {
		Subject:     "application/bacnet-xdd+zip",
		Description: "",
	},
	"xdf": {
		Subject:     "application/mrb-consumer+xml",
		Description: "",
	},
	"xdm": {
		Subject:     "application/vnd.syncml.dm+xml",
		Description: "",
	},
	"xdp": {
		Subject:     "application/vnd.adobe.xdp+xml",
		Description: "",
	},
	"xdr": {
		Subject:     "video/x-amt-demorun",
		Description: "",
	},
	"xdssc": {
		Subject:     "application/dssc+xml",
		Description: "",
	},
	"xdw": {
		Subject:     "application/vnd.fujifilm.fb.docuworks",
		Description: "",
	},
	"xel": {
		Subject:     "application/xcap-el+xml",
		Description: "",
	},
	"xenc": {
		Subject:     "application/xenc+xml",
		Description: "",
	},
	"xer": {
		Subject:     "application/patch-ops-error+xml",
		Description: "",
	},
	"xfd": {
		Subject:     "application/vnd.xfdl",
		Description: "",
	},
	"xfdf": {
		Subject:     "application/xfdf",
		Description: "",
	},
	"xfdl": {
		Subject:     "application/vnd.xfdl",
		Description: "",
	},
	"xgz": {
		Subject:     "x-xgl/x-drawing",
		Description: "",
	},
	"xhe": {
		Subject:     "audio/usac",
		Description: "",
	},
	"xht": {
		Subject:     "application/xhtml+xml",
		Description: "",
	},
	"xhtm": {
		Subject:     "application/vnd.pwg-xhtml-print+xml",
		Description: "",
	},
	"xhtml": {
		Subject:     "application/vnd.pwg-xhtml-print+xml",
		Description: "",
	},
	"xhvml": {
		Subject:     "application/xhtml-voice+xml",
		Description: "",
	},
	"xi": {
		Subject:     "audio/x-xi",
		Description: "",
	},
	"xif": {
		Subject:     "image/vnd.xiff",
		Description: "",
	},
	"xl": {
		Subject:     "application/x-excel",
		Description: "",
	},
	"xla": {
		Subject:     "application/vnd.ms-excel",
		Description: "",
	},
	"xlam": {
		Subject:     "application/vnd.ms-excel.addin.macroenabled.12",
		Description: "",
	},
	"xlb": {
		Subject:     "application/vnd.ms-excel",
		Description: "",
	},
	"xlc": {
		Subject:     "application/vnd.ms-excel",
		Description: "",
	},
	"xld": {
		Subject:     "application/x-excel",
		Description: "",
	},
	"xlf": {
		Subject:     "application/xliff+xml",
		Description: "",
	},
	"xliff": {
		Subject:     "application/x-xliff",
		Description: "",
	},
	"xlim": {
		Subject:     "application/vnd.xmpie.xlim",
		Description: "",
	},
	"xlk": {
		Subject:     "application/x-excel",
		Description: "",
	},
	"xll": {
		Subject:     "application/x-excel",
		Description: "",
	},
	"xlm": {
		Subject:     "application/vnd.ms-excel",
		Description: "",
	},
	"xlr": {
		Subject:     "application/vnd.ms-works",
		Description: "",
	},
	"xls": {
		Subject:     "application/vnd.ms-excel",
		Description: "",
	},
	"xlsb": {
		Subject:     "application/vnd.ms-excel.sheet.binary.macroenabled.12",
		Description: "",
	},
	"xlsm": {
		Subject:     "application/vnd.ms-excel.sheet.macroenabled.12",
		Description: "",
	},
	"xlsx": {
		Subject:     "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		Description: "",
	},
	"xlt": {
		Subject:     "application/vnd.ms-excel",
		Description: "",
	},
	"xltm": {
		Subject:     "application/vnd.ms-excel.template.macroenabled.12",
		Description: "",
	},
	"xltx": {
		Subject:     "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		Description: "",
	},
	"xlv": {
		Subject:     "application/x-excel",
		Description: "",
	},
	"xlw": {
		Subject:     "application/vnd.ms-excel",
		Description: "",
	},
	"xm": {
		Subject:     "audio/x-xm",
		Description: "",
	},
	"xmf": {
		Subject:     "audio/x-xmf",
		Description: "",
	},
	"xmi": {
		Subject:     "text/x-xmi",
		Description: "",
	},
	"xml": {
		Subject:     "application/xml",
		Description: "",
	},
	"xmls": {
		Subject:     "application/dskpp+xml",
		Description: "",
	},
	"xmt_bin": {
		Subject:     "model/vnd.parasolid.transmit.binary",
		Description: "",
	},
	"xmt_txt": {
		Subject:     "model/vnd.parasolid.transmit.text",
		Description: "",
	},
	"xmz": {
		Subject:     "x-xgl/x-movie",
		Description: "",
	},
	"xns": {
		Subject:     "application/xcap-ns+xml",
		Description: "",
	},
	"xo": {
		Subject:     "application/vnd.olpc-sugar",
		Description: "",
	},
	"xodp": {
		Subject:     "application/vnd.collabio.xodocuments.presentation",
		Description: "",
	},
	"xods": {
		Subject:     "application/vnd.collabio.xodocuments.spreadsheet",
		Description: "",
	},
	"xodt": {
		Subject:     "application/vnd.collabio.xodocuments.document",
		Description: "",
	},
	"xop": {
		Subject:     "application/xop+xml",
		Description: "",
	},
	"xotp": {
		Subject:     "application/vnd.collabio.xodocuments.presentation-template",
		Description: "",
	},
	"xots": {
		Subject:     "application/vnd.collabio.xodocuments.spreadsheet-template",
		Description: "",
	},
	"xott": {
		Subject:     "application/vnd.collabio.xodocuments.document-template",
		Description: "",
	},
	"xpak": {
		Subject:     "application/vnd.gentoo.xpak",
		Description: "",
	},
	"xpdl": {
		Subject:     "application/xml",
		Description: "",
	},
	"xpi": {
		Subject:     "application/x-xpinstall",
		Description: "",
	},
	"xpix": {
		Subject:     "application/x-vnd.ls-xpix",
		Description: "",
	},
	"xpl": {
		Subject:     "application/x-xproc+xml",
		Description: "",
	},
	"xpm": {
		Subject:     "image/x-xpixmap",
		Description: "",
	},
	"xpr": {
		Subject:     "application/vnd.is-xpr",
		Description: "",
	},
	"xps": {
		Subject:     "application/vnd.ms-xpsdocument",
		Description: "",
	},
	"xpw": {
		Subject:     "application/vnd.intercon.formnet",
		Description: "",
	},
	"xpx": {
		Subject:     "application/vnd.intercon.formnet",
		Description: "",
	},
	"xsd": {
		Subject:     "application/xml",
		Description: "",
	},
	"xsf": {
		Subject:     "application/prs.xsf+xml",
		Description: "",
	},
	"xsl": {
		Subject:     "application/xml",
		Description: "",
	},
	"xslfo": {
		Subject:     "text/x-xslfo",
		Description: "",
	},
	"xslt": {
		Subject:     "application/xslt+xml",
		Description: "",
	},
	"xsm": {
		Subject:     "application/vnd.syncml+xml",
		Description: "",
	},
	"xspf": {
		Subject:     "application/x-xspf+xml",
		Description: "",
	},
	"xsr": {
		Subject:     "video/x-amt-showrun",
		Description: "",
	},
	"xtel": {
		Subject:     "x-chemical/x-xtel",
		Description: "",
	},
	"xul": {
		Subject:     "application/vnd.mozilla.aul+xml",
		Description: "",
	},
	"xvm": {
		Subject:     "application/xhtml-voice+xml",
		Description: "",
	},
	"xvml": {
		Subject:     "application/xhtml-voice+xml",
		Description: "",
	},
	"xwd": {
		Subject:     "image/x-xwindowdump",
		Description: "",
	},
	"xyz": {
		Subject:     "x-chemical/x-xyz",
		Description: "",
	},
	"xyze": {
		Subject:     "image/vnd.radiance",
		Description: "",
	},
	"xz": {
		Subject:     "application/x-xz",
		Description: "",
	},
	"yaml": {
		Subject:     "application/vnd.oai.workflows",
		Description: "",
	},
	"yang": {
		Subject:     "application/yang",
		Description: "",
	},
	"yin": {
		Subject:     "application/yin+xml",
		Description: "",
	},
	"yme": {
		Subject:     "application/vnd.yaoweme",
		Description: "",
	},
	"yml": {
		Subject:     "application/vnd.oai.workflows+yaml",
		Description: "",
	},
	"yt": {
		Subject:     "video/vnd.youtube.yt",
		Description: "",
	},
	"z1": {
		Subject:     "application/x-zmachine",
		Description: "",
	},
	"z2": {
		Subject:     "application/x-zmachine",
		Description: "",
	},
	"z3": {
		Subject:     "application/x-zmachine",
		Description: "",
	},
	"z4": {
		Subject:     "application/x-zmachine",
		Description: "",
	},
	"z5": {
		Subject:     "application/x-zmachine",
		Description: "",
	},
	"z64": {
		Subject:     "application/x-n64-rom",
		Description: "",
	},
	"z6": {
		Subject:     "application/x-zmachine",
		Description: "",
	},
	"z7": {
		Subject:     "application/x-zmachine",
		Description: "",
	},
	"z8": {
		Subject:     "application/x-zmachine",
		Description: "",
	},
	"z": {
		Subject:     "application/x-compress",
		Description: "",
	},
	"zabw": {
		Subject:     "application/x-abiword",
		Description: "",
	},
	"zaz": {
		Subject:     "application/vnd.zzazz.deck+xml",
		Description: "",
	},
	"zfc": {
		Subject:     "application/vnd.filmit.zfc",
		Description: "",
	},
	"zip": {
		Subject:     "application/zip",
		Description: "",
	},
	"zir": {
		Subject:     "application/vnd.zul",
		Description: "",
	},
	"zirz": {
		Subject:     "application/vnd.zul",
		Description: "",
	},
	"zmm": {
		Subject:     "application/vnd.handheld-entertainment+xml",
		Description: "",
	},
	"zmt": {
		Subject:     "x-chemical/x-mopac-input",
		Description: "",
	},
	"zone": {
		Subject:     "text/dns",
		Description: "",
	},
	"zoo": {
		Subject:     "application/x-zoo",
		Description: "",
	},
	"zsav": {
		Subject:     "application/x-spss-sav",
		Description: "",
	},
	"zsh": {
		Subject:     "text/x-script.zsh",
		Description: "",
	},
	"zst": {
		Subject:     "application/zstd",
		Description: "",
	},
	"zz": {
		Subject:     "application/zlib",
		Description: "",
	},
	"~": {
		Subject:     "application/x-trash",
		Description: "",
	},
}

func GetRef() map[string]dictionary.Element {
	return Ref
}
