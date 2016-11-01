package goarabic

// Harf holds the Arabic character with its different representation forms (glyphs).
type Harf struct {
	Unicode, Isolated, Beggining, Medium, Final rune
}

// Vowels (Tashkeel) characters.
var (
	FATHA    rune = '\u064e'
	FATHATAN rune = '\u064b'
	DAMMA    rune = '\u064f'
	DAMMATAN rune = '\u064c'
	KASRA    rune = '\u0650'
	KASRATAN rune = '\u064d'
	SHADDA   rune = '\u0651'
	SUKUN    rune = '\u0652'
)

// Arabic Alphabet using the new Harf type.
var (
	ALEF_HAMZA_ABOVE = Harf{ // أ
		Unicode:   '\u0623',
		Isolated:  '\ufe83',
		Beggining: '\u0623',
		Medium:    '\ufe84',
		Final:     '\ufe84'}

	ALEF = Harf{ // ا
		Unicode:   '\u0627',
		Isolated:  '\ufe8d',
		Beggining: '\u0627',
		Medium:    '\ufe8e',
		Final:     '\ufe8e'}

	ALEF_MADDA_ABOVE = Harf{ // آ
		Unicode:   '\u0622',
		Isolated:  '\ufe81',
		Beggining: '\u0622',
		Medium:    '\ufe82',
		Final:     '\ufe82'}

	HAMZA = Harf{ // ء
		Unicode:   '\u0621',
		Isolated:  '\ufe80',
		Beggining: '\u0621',
		Medium:    '\u0621',
		Final:     '\u0621'}

	WAW_HAMZA_ABOVE = Harf{ // ؤ
		Unicode:   '\u0624',
		Isolated:  '\ufe85',
		Beggining: '\u0624',
		Medium:    '\ufe86',
		Final:     '\ufe86'}

	ALEF_HAMZA_BELOW = Harf{ // أ
		Unicode:   '\u0625',
		Isolated:  '\ufe87',
		Beggining: '\u0625',
		Medium:    '\ufe88',
		Final:     '\ufe88'}

	YEH_HAMZA_ABOVE = Harf{ // ئ
		Unicode:   '\u0626',
		Isolated:  '\ufe89',
		Beggining: '\ufe8b',
		Medium:    '\ufe8c',
		Final:     '\ufe8a'}

	BEH = Harf{ // ب
		Unicode:   '\u0628',
		Isolated:  '\ufe8f',
		Beggining: '\ufe91',
		Medium:    '\ufe92',
		Final:     '\ufe90'}

	PEH = Harf{ // پ
		Unicode:   '\u067e',
		Isolated:  '\ufb56',
		Beggining: '\ufb58',
		Medium:    '\ufb59',
		Final:     '\ufb57'}

	TEH = Harf{ // ت
		Unicode:   '\u062A',
		Isolated:  '\ufe95',
		Beggining: '\ufe97',
		Medium:    '\ufe98',
		Final:     '\ufe96'}

	TEH_MARBUTA = Harf{ // ة
		Unicode:   '\u0629',
		Isolated:  '\ufe93',
		Beggining: '\u0629',
		Medium:    '\u0629',
		Final:     '\ufe94'}

	THEH = Harf{ // ث
		Unicode:   '\u062b',
		Isolated:  '\ufe99',
		Beggining: '\ufe9b',
		Medium:    '\ufe9c',
		Final:     '\ufe9a'}

	JEEM = Harf{ // ج
		Unicode:   '\u062c',
		Isolated:  '\ufe9d',
		Beggining: '\ufe9f',
		Medium:    '\ufea0',
		Final:     '\ufe9e'}

	TCHEH = Harf{ // چ
		Unicode:   '\u0686',
		Isolated:  '\ufb7a',
		Beggining: '\ufb7c',
		Medium:    '\ufb7d',
		Final:     '\ufb7b'}

	HAH = Harf{ // ح
		Unicode:   '\u062d',
		Isolated:  '\ufea1',
		Beggining: '\ufea3',
		Medium:    '\ufea4',
		Final:     '\ufea2'}

	KHAH = Harf{ // خ
		Unicode:   '\u062e',
		Isolated:  '\ufea5',
		Beggining: '\ufea7',
		Medium:    '\ufea8',
		Final:     '\ufea6'}

	DAL = Harf{ // د
		Unicode:   '\u062f',
		Isolated:  '\ufea9',
		Beggining: '\u062f',
		Medium:    '\ufeaa',
		Final:     '\ufeaa'}

	THAL = Harf{ // ذ
		Unicode:   '\u0630',
		Isolated:  '\ufeab',
		Beggining: '\u0630',
		Medium:    '\ufeac',
		Final:     '\ufeac'}

	REH = Harf{ // ر
		Unicode:   '\u0631',
		Isolated:  '\ufead',
		Beggining: '\u0631',
		Medium:    '\ufeae',
		Final:     '\ufeae'}

	JEH = Harf{
		Unicode:   '\u0698',
		Isolated:  '\ufb8a',
		Beggining: '\u0698',
		Medium:    '\ufb8b',
		Final:     '\ufb8b',
	}

	ZAIN = Harf{ // ز
		Unicode:   '\u0632',
		Isolated:  '\ufeaf',
		Beggining: '\u0632',
		Medium:    '\ufeb0',
		Final:     '\ufeb0'}

	SEEN = Harf{ // س
		Unicode:   '\u0633',
		Isolated:  '\ufeb1',
		Beggining: '\ufeb3',
		Medium:    '\ufeb4',
		Final:     '\ufeb2'}

	SHEEN = Harf{ // ش
		Unicode:   '\u0634',
		Isolated:  '\ufeb5',
		Beggining: '\ufeb7',
		Medium:    '\ufeb8',
		Final:     '\ufeb6'}

	SAD = Harf{ // ص
		Unicode:   '\u0635',
		Isolated:  '\ufeb9',
		Beggining: '\ufebb',
		Medium:    '\ufebc',
		Final:     '\ufeba'}

	DAD = Harf{ // ض
		Unicode:   '\u0636',
		Isolated:  '\ufebd',
		Beggining: '\ufebf',
		Medium:    '\ufec0',
		Final:     '\ufebe'}

	TAH = Harf{ // ط
		Unicode:   '\u0637',
		Isolated:  '\ufec1',
		Beggining: '\ufec3',
		Medium:    '\ufec4',
		Final:     '\ufec2'}

	ZAH = Harf{ // ظ
		Unicode:   '\u0638',
		Isolated:  '\ufec5',
		Beggining: '\ufec7',
		Medium:    '\ufec8',
		Final:     '\ufec6'}

	AIN = Harf{ // ع
		Unicode:   '\u0639',
		Isolated:  '\ufec9',
		Beggining: '\ufecb',
		Medium:    '\ufecc',
		Final:     '\ufeca'}

	GHAIN = Harf{ // غ
		Unicode:   '\u063a',
		Isolated:  '\ufecd',
		Beggining: '\ufecf',
		Medium:    '\ufed0',
		Final:     '\ufece'}

	FEH = Harf{ // ف
		Unicode:   '\u0641',
		Isolated:  '\ufed1',
		Beggining: '\ufed3',
		Medium:    '\ufed4',
		Final:     '\ufed2'}

	QAF = Harf{ // ق
		Unicode:   '\u0642',
		Isolated:  '\ufed5',
		Beggining: '\ufed7',
		Medium:    '\ufed8',
		Final:     '\ufed6'}

	KAF = Harf{ // ك
		Unicode:   '\u0643',
		Isolated:  '\ufed9',
		Beggining: '\ufedb',
		Medium:    '\ufedc',
		Final:     '\ufeda'}

	KEHEH = Harf{ // ک
		Unicode:   '\u06a9',
		Isolated:  '\ufb8e',
		Beggining: '\ufb90',
		Medium:    '\ufb91',
		Final:     '\ufb8f',
	}

	GAF = Harf{ // گ
		Unicode:   '\u06af',
		Isolated:  '\ufb92',
		Beggining: '\ufb94',
		Medium:    '\ufb95',
		Final:     '\ufb93'}

	LAM = Harf{ // ل
		Unicode:   '\u0644',
		Isolated:  '\ufedd',
		Beggining: '\ufedf',
		Medium:    '\ufee0',
		Final:     '\ufede'}

	MEEM = Harf{ // م
		Unicode:   '\u0645',
		Isolated:  '\ufee1',
		Beggining: '\ufee3',
		Medium:    '\ufee4',
		Final:     '\ufee2'}

	NOON = Harf{ // ن
		Unicode:   '\u0646',
		Isolated:  '\ufee5',
		Beggining: '\ufee7',
		Medium:    '\ufee8',
		Final:     '\ufee6'}

	HEH = Harf{ // ه
		Unicode:   '\u0647',
		Isolated:  '\ufee9',
		Beggining: '\ufeeb',
		Medium:    '\ufeec',
		Final:     '\ufeea'}

	WAW = Harf{ // و
		Unicode:   '\u0648',
		Isolated:  '\ufeed',
		Beggining: '\u0648',
		Medium:    '\ufeee',
		Final:     '\ufeee'}

	YEH = Harf{ // ی
		Unicode:   '\u06cc',
		Isolated:  '\ufbfc',
		Beggining: '\ufbfe',
		Medium:    '\ufbff',
		Final:     '\ufbfd'}

	ARABICYEH = Harf{ // ي
		Unicode:   '\u064a',
		Isolated:  '\ufef1',
		Beggining: '\ufef3',
		Medium:    '\ufef4',
		Final:     '\ufef2'}

	ALEF_MAKSURA = Harf{ // ى
		Unicode:   '\u0649',
		Isolated:  '\ufeef',
		Beggining: '\u0649',
		Medium:    '\ufef0',
		Final:     '\ufef0'}

	TATWEEL = Harf{ // ـ
		Unicode:   '\u0640',
		Isolated:  '\u0640',
		Beggining: '\u0640',
		Medium:    '\u0640',
		Final:     '\u0640'}

	LAM_ALEF = Harf{ // لا
		Unicode:   '\ufefb',
		Isolated:  '\ufefb',
		Beggining: '\ufefb',
		Medium:    '\ufefc',
		Final:     '\ufefc'}

	LAM_ALEF_HAMZA_ABOVE = Harf{ // ﻷ
		Unicode:   '\ufef7',
		Isolated:  '\ufef7',
		Beggining: '\ufef7',
		Medium:    '\ufef8',
		Final:     '\ufef8'}
)

var alphabet = []Harf{
	ALEF_HAMZA_ABOVE,
	ALEF,
	ALEF_MADDA_ABOVE,
	HAMZA,
	WAW_HAMZA_ABOVE,
	ALEF_HAMZA_BELOW,
	YEH_HAMZA_ABOVE,
	BEH,
	PEH,
	TEH,
	TEH_MARBUTA,
	THEH,
	JEEM,
	TCHEH,
	HAH,
	KHAH,
	DAL,
	THAL,
	REH,
	JEH,
	ZAIN,
	SEEN,
	SHEEN,
	SAD,
	DAD,
	TAH,
	ZAH,
	AIN,
	GHAIN,
	FEH,
	QAF,
	KAF,
	KEHEH,
	GAF,
	LAM,
	MEEM,
	NOON,
	HEH,
	WAW,
	YEH,
	ARABICYEH,
	ALEF_MAKSURA,
	TATWEEL,
	LAM_ALEF,
	LAM_ALEF_HAMZA_ABOVE,
}

// use map for faster lookups.
var tashkeel = map[rune]bool{FATHA: true, FATHATAN: true, DAMMA: true,
	DAMMATAN: true, KASRA: true, KASRATAN: true,
	SHADDA: true, SUKUN: true}

// use map for faster lookups.
// var special_char = map[rune]bool{"": true, ' ': true, '?': true,
//	'؟': true, '.': true, KASRATAN: true,
//	SHADDA: true, SUKUN: true}

// use map for faster lookups.
var beggining_after = map[Harf]bool{
	ALEF_HAMZA_ABOVE: true,
	ALEF_MADDA_ABOVE: true,
	ALEF:             true,
	HAMZA:            true,
	WAW_HAMZA_ABOVE:  true,
	ALEF_HAMZA_BELOW: true,
	TEH_MARBUTA:      true,
	DAL:              true,
	THAL:             true,
	REH:              true,
	ZAIN:             true,
	WAW:              true,
	ALEF_MAKSURA:     true}
