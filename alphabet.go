package goarabic

// Harf holds the Arabic character with its different representation forms (glyphs).
type Harf struct {
	Unicode, Isolated, Beginning, Medium, Final rune
}

// Vowels (Tashkeel) characters.
var (
	FATHA    = '\u064e'
	FATHATAN = '\u064b'
	DAMMA    = '\u064f'
	DAMMATAN = '\u064c'
	KASRA    = '\u0650'
	KASRATAN = '\u064d'
	SHADDA   = '\u0651'
	SUKUN    = '\u0652'
)

// Arabic Alphabet using the new Harf type.
var (
	ALEF_HAMZA_ABOVE = Harf{ // أ
		Unicode:   '\u0623',
		Isolated:  '\ufe83',
		Beginning: '\u0623',
		Medium:    '\ufe84',
		Final:     '\ufe84',
	}

	ALEF = Harf{ // ا
		Unicode:   '\u0627',
		Isolated:  '\ufe8d',
		Beginning: '\u0627',
		Medium:    '\ufe8e',
		Final:     '\ufe8e',
	}

	ALEF_MADDA_ABOVE = Harf{ // آ
		Unicode:   '\u0622',
		Isolated:  '\ufe81',
		Beginning: '\u0622',
		Medium:    '\ufe82',
		Final:     '\ufe82',
	}

	HAMZA = Harf{ // ء
		Unicode:   '\u0621',
		Isolated:  '\ufe80',
		Beginning: '\u0621',
		Medium:    '\u0621',
		Final:     '\u0621',
	}

	WAW_HAMZA_ABOVE = Harf{ // ؤ
		Unicode:   '\u0624',
		Isolated:  '\ufe85',
		Beginning: '\u0624',
		Medium:    '\ufe86',
		Final:     '\ufe86',
	}

	ALEF_HAMZA_BELOW = Harf{ // أ
		Unicode:   '\u0625',
		Isolated:  '\ufe87',
		Beginning: '\u0625',
		Medium:    '\ufe88',
		Final:     '\ufe88',
	}

	YEH_HAMZA_ABOVE = Harf{ // ئ
		Unicode:   '\u0626',
		Isolated:  '\ufe89',
		Beginning: '\ufe8b',
		Medium:    '\ufe8c',
		Final:     '\ufe8a',
	}

	BEH = Harf{ // ب
		Unicode:   '\u0628',
		Isolated:  '\ufe8f',
		Beginning: '\ufe91',
		Medium:    '\ufe92',
		Final:     '\ufe90',
	}

	PEH = Harf{ // پ
		Unicode:   '\u067e',
		Isolated:  '\ufb56',
		Beginning: '\ufb58',
		Medium:    '\ufb59',
		Final:     '\ufb57',
	}

	TEH = Harf{ // ت
		Unicode:   '\u062A',
		Isolated:  '\ufe95',
		Beginning: '\ufe97',
		Medium:    '\ufe98',
		Final:     '\ufe96',
	}

	TEH_MARBUTA = Harf{ // ة
		Unicode:   '\u0629',
		Isolated:  '\ufe93',
		Beginning: '\u0629',
		Medium:    '\u0629',
		Final:     '\ufe94',
	}

	THEH = Harf{ // ث
		Unicode:   '\u062b',
		Isolated:  '\ufe99',
		Beginning: '\ufe9b',
		Medium:    '\ufe9c',
		Final:     '\ufe9a',
	}

	JEEM = Harf{ // ج
		Unicode:   '\u062c',
		Isolated:  '\ufe9d',
		Beginning: '\ufe9f',
		Medium:    '\ufea0',
		Final:     '\ufe9e',
	}

	TCHEH = Harf{ // چ
		Unicode:   '\u0686',
		Isolated:  '\ufb7a',
		Beginning: '\ufb7c',
		Medium:    '\ufb7d',
		Final:     '\ufb7b',
	}

	HAH = Harf{ // ح
		Unicode:   '\u062d',
		Isolated:  '\ufea1',
		Beginning: '\ufea3',
		Medium:    '\ufea4',
		Final:     '\ufea2',
	}

	KHAH = Harf{ // خ
		Unicode:   '\u062e',
		Isolated:  '\ufea5',
		Beginning: '\ufea7',
		Medium:    '\ufea8',
		Final:     '\ufea6',
	}

	DAL = Harf{ // د
		Unicode:   '\u062f',
		Isolated:  '\ufea9',
		Beginning: '\u062f',
		Medium:    '\ufeaa',
		Final:     '\ufeaa',
	}

	THAL = Harf{ // ذ
		Unicode:   '\u0630',
		Isolated:  '\ufeab',
		Beginning: '\u0630',
		Medium:    '\ufeac',
		Final:     '\ufeac',
	}

	REH = Harf{ // ر
		Unicode:   '\u0631',
		Isolated:  '\ufead',
		Beginning: '\u0631',
		Medium:    '\ufeae',
		Final:     '\ufeae',
	}

	JEH = Harf{
		Unicode:   '\u0698',
		Isolated:  '\ufb8a',
		Beginning: '\u0698',
		Medium:    '\ufb8b',
		Final:     '\ufb8b',
	}

	ZAIN = Harf{ // ز
		Unicode:   '\u0632',
		Isolated:  '\ufeaf',
		Beginning: '\u0632',
		Medium:    '\ufeb0',
		Final:     '\ufeb0',
	}

	SEEN = Harf{ // س
		Unicode:   '\u0633',
		Isolated:  '\ufeb1',
		Beginning: '\ufeb3',
		Medium:    '\ufeb4',
		Final:     '\ufeb2',
	}

	SHEEN = Harf{ // ش
		Unicode:   '\u0634',
		Isolated:  '\ufeb5',
		Beginning: '\ufeb7',
		Medium:    '\ufeb8',
		Final:     '\ufeb6',
	}

	SAD = Harf{ // ص
		Unicode:   '\u0635',
		Isolated:  '\ufeb9',
		Beginning: '\ufebb',
		Medium:    '\ufebc',
		Final:     '\ufeba',
	}

	DAD = Harf{ // ض
		Unicode:   '\u0636',
		Isolated:  '\ufebd',
		Beginning: '\ufebf',
		Medium:    '\ufec0',
		Final:     '\ufebe',
	}

	TAH = Harf{ // ط
		Unicode:   '\u0637',
		Isolated:  '\ufec1',
		Beginning: '\ufec3',
		Medium:    '\ufec4',
		Final:     '\ufec2',
	}

	ZAH = Harf{ // ظ
		Unicode:   '\u0638',
		Isolated:  '\ufec5',
		Beginning: '\ufec7',
		Medium:    '\ufec8',
		Final:     '\ufec6',
	}

	AIN = Harf{ // ع
		Unicode:   '\u0639',
		Isolated:  '\ufec9',
		Beginning: '\ufecb',
		Medium:    '\ufecc',
		Final:     '\ufeca',
	}

	GHAIN = Harf{ // غ
		Unicode:   '\u063a',
		Isolated:  '\ufecd',
		Beginning: '\ufecf',
		Medium:    '\ufed0',
		Final:     '\ufece',
	}

	FEH = Harf{ // ف
		Unicode:   '\u0641',
		Isolated:  '\ufed1',
		Beginning: '\ufed3',
		Medium:    '\ufed4',
		Final:     '\ufed2',
	}

	QAF = Harf{ // ق
		Unicode:   '\u0642',
		Isolated:  '\ufed5',
		Beginning: '\ufed7',
		Medium:    '\ufed8',
		Final:     '\ufed6',
	}

	KAF = Harf{ // ك
		Unicode:   '\u0643',
		Isolated:  '\ufed9',
		Beginning: '\ufedb',
		Medium:    '\ufedc',
		Final:     '\ufeda',
	}

	KEHEH = Harf{ // ک
		Unicode:   '\u06a9',
		Isolated:  '\ufb8e',
		Beginning: '\ufb90',
		Medium:    '\ufb91',
		Final:     '\ufb8f',
	}

	GAF = Harf{ // گ
		Unicode:   '\u06af',
		Isolated:  '\ufb92',
		Beginning: '\ufb94',
		Medium:    '\ufb95',
		Final:     '\ufb93',
	}

	LAM = Harf{ // ل
		Unicode:   '\u0644',
		Isolated:  '\ufedd',
		Beginning: '\ufedf',
		Medium:    '\ufee0',
		Final:     '\ufede',
	}

	MEEM = Harf{ // م
		Unicode:   '\u0645',
		Isolated:  '\ufee1',
		Beginning: '\ufee3',
		Medium:    '\ufee4',
		Final:     '\ufee2',
	}

	NOON = Harf{ // ن
		Unicode:   '\u0646',
		Isolated:  '\ufee5',
		Beginning: '\ufee7',
		Medium:    '\ufee8',
		Final:     '\ufee6',
	}

	HEH = Harf{ // ه
		Unicode:   '\u0647',
		Isolated:  '\ufee9',
		Beginning: '\ufeeb',
		Medium:    '\ufeec',
		Final:     '\ufeea',
	}

	WAW = Harf{ // و
		Unicode:   '\u0648',
		Isolated:  '\ufeed',
		Beginning: '\u0648',
		Medium:    '\ufeee',
		Final:     '\ufeee',
	}

	YEH = Harf{ // ی
		Unicode:   '\u06cc',
		Isolated:  '\ufbfc',
		Beginning: '\ufbfe',
		Medium:    '\ufbff',
		Final:     '\ufbfd',
	}

	ARABICYEH = Harf{ // ي
		Unicode:   '\u064a',
		Isolated:  '\ufef1',
		Beginning: '\ufef3',
		Medium:    '\ufef4',
		Final:     '\ufef2',
	}

	ALEF_MAKSURA = Harf{ // ى
		Unicode:   '\u0649',
		Isolated:  '\ufeef',
		Beginning: '\u0649',
		Medium:    '\ufef0',
		Final:     '\ufef0',
	}

	TATWEEL = Harf{ // ـ
		Unicode:   '\u0640',
		Isolated:  '\u0640',
		Beginning: '\u0640',
		Medium:    '\u0640',
		Final:     '\u0640',
	}

	LAM_ALEF = Harf{ // لا
		Unicode:   '\ufefb',
		Isolated:  '\ufefb',
		Beginning: '\ufefb',
		Medium:    '\ufefc',
		Final:     '\ufefc',
	}

	LAM_ALEF_HAMZA_ABOVE = Harf{ // ﻷ
		Unicode:   '\ufef7',
		Isolated:  '\ufef7',
		Beginning: '\ufef7',
		Medium:    '\ufef8',
		Final:     '\ufef8',
	}
)

var runeAlphabetMap = map[rune]Harf{
	// أ
	'\u0623': ALEF_HAMZA_ABOVE,
	'\ufe83': ALEF_HAMZA_ABOVE,
	'\ufe84': ALEF_HAMZA_ABOVE,

	// ا
	'\u0627': ALEF,
	'\ufe8d': ALEF,
	'\ufe8e': ALEF,

	// آ
	'\u0622': ALEF_MADDA_ABOVE,
	'\ufe81': ALEF_MADDA_ABOVE,
	'\ufe82': ALEF_MADDA_ABOVE,

	// ء
	'\u0621': HAMZA,
	'\ufe80': HAMZA,

	// ؤ
	'\u0624': WAW_HAMZA_ABOVE,
	'\ufe85': WAW_HAMZA_ABOVE,
	'\ufe86': WAW_HAMZA_ABOVE,

	// أ
	'\u0625': ALEF_HAMZA_BELOW,
	'\ufe87': ALEF_HAMZA_BELOW,
	'\ufe88': ALEF_HAMZA_BELOW,

	// ئ
	'\u0626': YEH_HAMZA_ABOVE,
	'\ufe89': YEH_HAMZA_ABOVE,
	'\ufe8b': YEH_HAMZA_ABOVE,
	'\ufe8c': YEH_HAMZA_ABOVE,
	'\ufe8a': YEH_HAMZA_ABOVE,

	// ب
	'\u0628': BEH,
	'\ufe8f': BEH,
	'\ufe91': BEH,
	'\ufe92': BEH,
	'\ufe90': BEH,

	// پ
	'\u067e': PEH,
	'\ufb56': PEH,
	'\ufb58': PEH,
	'\ufb59': PEH,
	'\ufb57': PEH,

	// ت
	'\u062A': TEH,
	'\ufe95': TEH,
	'\ufe97': TEH,
	'\ufe98': TEH,
	'\ufe96': TEH,

	// ة
	'\u0629': TEH_MARBUTA,
	'\ufe93': TEH_MARBUTA,
	'\ufe94': TEH_MARBUTA,

	// ث
	'\u062b': THEH,
	'\ufe99': THEH,
	'\ufe9b': THEH,
	'\ufe9c': THEH,
	'\ufe9a': THEH,

	// ج
	'\u062c': JEEM,
	'\ufe9d': JEEM,
	'\ufe9f': JEEM,
	'\ufea0': JEEM,
	'\ufe9e': JEEM,

	// چ
	'\u0686': TCHEH,
	'\ufb7a': TCHEH,
	'\ufb7c': TCHEH,
	'\ufb7d': TCHEH,
	'\ufb7b': TCHEH,

	// ح
	'\u062d': HAH,
	'\ufea1': HAH,
	'\ufea3': HAH,
	'\ufea4': HAH,
	'\ufea2': HAH,

	// خ
	'\u062e': KHAH,
	'\ufea5': KHAH,
	'\ufea7': KHAH,
	'\ufea8': KHAH,
	'\ufea6': KHAH,

	// د
	'\u062f': DAL,
	'\ufea9': DAL,
	'\ufeaa': DAL,

	// ذ
	'\u0630': THAL,
	'\ufeab': THAL,
	'\ufeac': THAL,

	// ر
	'\u0631': REH,
	'\ufead': REH,
	'\ufeae': REH,

	// ژ
	'\u0698': JEH,
	'\ufb8a': JEH,
	'\ufb8b': JEH,

	// ز
	'\u0632': ZAIN,
	'\ufeaf': ZAIN,
	'\ufeb0': ZAIN,

	// س
	'\u0633': SEEN,
	'\ufeb1': SEEN,
	'\ufeb3': SEEN,
	'\ufeb4': SEEN,
	'\ufeb2': SEEN,

	// ش
	'\u0634': SHEEN,
	'\ufeb5': SHEEN,
	'\ufeb7': SHEEN,
	'\ufeb8': SHEEN,
	'\ufeb6': SHEEN,

	// ص
	'\u0635': SAD,
	'\ufeb9': SAD,
	'\ufebb': SAD,
	'\ufebc': SAD,
	'\ufeba': SAD,

	// ض
	'\u0636': DAD,
	'\ufebd': DAD,
	'\ufebf': DAD,
	'\ufec0': DAD,
	'\ufebe': DAD,

	// ط
	'\u0637': TAH,
	'\ufec1': TAH,
	'\ufec3': TAH,
	'\ufec4': TAH,
	'\ufec2': TAH,

	// ظ
	'\u0638': ZAH,
	'\ufec5': ZAH,
	'\ufec7': ZAH,
	'\ufec8': ZAH,
	'\ufec6': ZAH,

	// ع
	'\u0639': AIN,
	'\ufec9': AIN,
	'\ufecb': AIN,
	'\ufecc': AIN,
	'\ufeca': AIN,

	// غ
	'\u063a': GHAIN,
	'\ufecd': GHAIN,
	'\ufecf': GHAIN,
	'\ufed0': GHAIN,
	'\ufece': GHAIN,

	// ف
	'\u0641': FEH,
	'\ufed1': FEH,
	'\ufed3': FEH,
	'\ufed4': FEH,
	'\ufed2': FEH,

	// ق
	'\u0642': QAF,
	'\ufed5': QAF,
	'\ufed7': QAF,
	'\ufed8': QAF,
	'\ufed6': QAF,

	// ك
	'\u0643': KAF,
	'\ufed9': KAF,
	'\ufedb': KAF,
	'\ufedc': KAF,
	'\ufeda': KAF,

	// ک
	'\u06a9': KEHEH,
	'\ufb8e': KEHEH,
	'\ufb90': KEHEH,
	'\ufb91': KEHEH,
	'\ufb8f': KEHEH,

	// گ
	'\u06af': GAF,
	'\ufb92': GAF,
	'\ufb94': GAF,
	'\ufb95': GAF,
	'\ufb93': GAF,

	// ل
	'\u0644': LAM,
	'\ufedd': LAM,
	'\ufedf': LAM,
	'\ufee0': LAM,
	'\ufede': LAM,

	// م
	'\u0645': MEEM,
	'\ufee1': MEEM,
	'\ufee3': MEEM,
	'\ufee4': MEEM,
	'\ufee2': MEEM,

	// ن
	'\u0646': NOON,
	'\ufee5': NOON,
	'\ufee7': NOON,
	'\ufee8': NOON,
	'\ufee6': NOON,

	// ه
	'\u0647': HEH,
	'\ufee9': HEH,
	'\ufeeb': HEH,
	'\ufeec': HEH,
	'\ufeea': HEH,

	// و
	'\u0648': WAW,
	'\ufeed': WAW,
	'\ufeee': WAW,

	// ی
	'\u06cc': YEH,
	'\ufbfc': YEH,
	'\ufbfe': YEH,
	'\ufbff': YEH,
	'\ufbfd': YEH,

	// ي
	'\u064a': ARABICYEH,
	'\ufef1': ARABICYEH,
	'\ufef3': ARABICYEH,
	'\ufef4': ARABICYEH,
	'\ufef2': ARABICYEH,

	// ى
	'\u0649': ALEF_MAKSURA,
	'\ufeef': ALEF_MAKSURA,
	'\ufef0': ALEF_MAKSURA,

	// ـ
	'\u0640': TATWEEL,

	// لا
	'\ufefb': LAM_ALEF,
	'\ufefc': LAM_ALEF,

	// ﻷ
	'\ufef7': LAM_ALEF_HAMZA_ABOVE,
	'\ufef8': LAM_ALEF_HAMZA_ABOVE,
}

const tatweelRune = '\u0640'

// use map for faster lookups.
var tashkeel = map[rune]struct{}{
	FATHA:    {},
	FATHATAN: {},
	DAMMA:    {},
	DAMMATAN: {},
	KASRA:    {},
	KASRATAN: {},
	SHADDA:   {},
	SUKUN:    {},
}

// use map for faster lookups.
// var special_char = map[rune]bool{"": true, ' ': true, '?': true,
//	'؟': true, '.': true, KASRATAN: true,
//	SHADDA: true, SUKUN: true}

// use map for faster lookups.
var runeBeginningAfter = map[rune]struct{}{
	// أ
	'\u0623': {},
	'\ufe83': {},
	'\ufe84': {},

	// آ
	'\u0622': {},
	'\ufe81': {},
	'\ufe82': {},

	// ا
	'\u0627': {},
	'\ufe8d': {},
	'\ufe8e': {},

	// ء
	'\u0621': {},
	'\ufe80': {},

	// ؤ
	'\u0624': {},
	'\ufe85': {},
	'\ufe86': {},

	// أ
	'\u0625': {},
	'\ufe87': {},
	'\ufe88': {},

	// ة
	'\u0629': {},
	'\ufe93': {},
	'\ufe94': {},

	// د
	'\u062f': {},
	'\ufea9': {},
	'\ufeaa': {},

	// ذ
	'\u0630': {},
	'\ufeab': {},
	'\ufeac': {},

	// ر
	'\u0631': {},
	'\ufead': {},
	'\ufeae': {},

	// ز
	'\u0632': {},
	'\ufeaf': {},
	'\ufeb0': {},

	// ى
	'\u0649': {},
	'\ufeef': {},
	'\ufef0': {},

	// لا
	'\ufefb': {},
	'\ufefc': {},
}
