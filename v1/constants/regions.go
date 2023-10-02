package constants

const (
	Argentina                           string = "es-ar"
	Australia                           string = "en-au"
	Austria                             string = "de-at"
	BahrainEnglish                      string = "en-bh"
	Bahrain                             string = "ar-bh"
	BelgiumFrench                       string = "fr-be"
	Belgium                             string = "nl-be"
	Bolivia                             string = "es-bo"
	Brasil                              string = "pt-br"
	BulgariaEnglish                     string = "en-bg"
	Bulgaria                            string = "bg-bg"
	Canada                              string = "en-ca"
	CanadaFrench                        string = "fr-ca"
	Chile                               string = "es-cl"
	MainlandChina                       string = "zh-hans-cn"
	Colombia                            string = "es-co"
	CostaRica                           string = "es-cr"
	CroatiaEnglish                      string = "en-hr"
	Croatia                             string = "hr-hr"
	Cyprus                              string = "en-cy"
	CzechRepublicEnglish                string = "en-cz"
	CzechRepublic                       string = "cs-cz"
	Denmark                             string = "da-dk"
	DenmarkEnglish                      string = "en-dk"
	Ecuador                             string = "es-ec"
	ElSalvador                          string = "es-sv"
	FinlandEnglish                      string = "en-fi"
	Finland                             string = "fi-fi"
	France                              string = "fr-fr"
	Germany                             string = "de-de"
	GreeceEnglish                       string = "en-gr"
	Greece                              string = "el-gr"
	Guatemala                           string = "es-gt"
	Honduras                            string = "es-hn"
	HongKongEnglish                     string = "en-hk"
	HongKong                            string = "zh-hans-hk"
	HongKongTraditionalChinese          string = "zh-hant-hk"
	HungaryEnglish                      string = "en-hu"
	Hungary                             string = "hu-hu"
	IcelandEnglish                      string = "en-is"
	India                               string = "en-in"
	IndonesiaEnglish                    string = "en-id"
	Ireland                             string = "en-ie"
	IsraelEnglish                       string = "en-il"
	Israel                              string = "he-il"
	Italy                               string = "it-it"
	Japan                               string = "ja-jp"
	Korea                               string = "ko-kr"
	Kuwait                              string = "ar-kw"
	KuwaitEnglish                       string = "en-kw"
	Lebanon                             string = "ar-lb"
	LebanonEnglish                      string = "en-lb"
	Luxembourg                          string = "de-lu"
	LuxembourgFrench                    string = "fr-lu"
	MalaysiaEnglish                     string = "en-my"
	Malta                               string = "en-mt"
	Mexico                              string = "es-mx"
	Nederland                           string = "nl-nl"
	NewZealand                          string = "en-nz"
	Nicaragua                           string = "es-ni"
	NorwayEnglish                       string = "en-no"
	Norway                              string = "no-no"
	Oman                                string = "ar-om"
	OmanEnglish                         string = "en-om"
	Panama                              string = "es-pa"
	Paraguay                            string = "es-py"
	Peru                                string = "es-pe"
	PhilippinesEnglish                  string = "en-ph"
	PolandEnglish                       string = "en-pl"
	Poland                              string = "pl-pl"
	Portugal                            string = "pt-pt"
	Qatar                               string = "ar-qa"
	QatarEnglish                        string = "en-qa"
	RomaniaEnglish                      string = "en-ro"
	Romania                             string = "ro-ro"
	Russia                              string = "ru-ru"
	SaudiArabia                         string = "ar-sa"
	SaudiArabiaEnglish                  string = "en-sa"
	Serbia                              string = "sr-rs"
	SingaporeEnglish                    string = "en-sg"
	SloveniaEnglish                     string = "en-si"
	Slovenia                            string = "sl-si"
	SlovakiaEnglish                     string = "en-sk"
	Slovakia                            string = "sk-sk"
	SouthAfrica                         string = "en-za"
	Spain                               string = "es-es"
	SwedenEnglish                       string = "en-se"
	Sweden                              string = "sv-se"
	Switzerland                         string = "de-ch"
	SwitzerlandFrench                   string = "fr-ch"
	SwitzerlandItaliano                 string = "it-ch"
	TaiwanEnglish                       string = "en-tw"
	Taiwan                              string = "zh-hant-tw"
	ThailandEnglish                     string = "en-th"
	Thailand                            string = "th-th"
	TurkeyEnglish                       string = "en-tr"
	Turkey                              string = "tr-tr"
	UkraineRussian                      string = "ru-ua"
	Ukraine                             string = "uk-ua"
	UnitedArabEmiratesMiddleEast        string = "ar-ae"
	UnitedArabEmiratesMiddleEastEnglish string = "en-ae"
	UnitedStates                        string = "en-us"
	UnitedKingdom                       string = "en-gb"
	Uruguay                             string = "es-uy"
	VietnamEnglish                      string = "en-vn"
)

func GetRegionsList() map[string]string {
	regions := map[string]string{
		Argentina: "Argentina",
		Australia: "Australia",
		Austria: "Austria",
		BahrainEnglish: "Bahrain English",
		Bahrain: "Bahrain",
		BelgiumFrench: "Belgium French",
		Belgium: "Belgium",
		Bolivia: "Bolivia",
		Brasil: "Brasil",
		BulgariaEnglish: "Bulgaria English",
		Bulgaria: "Bulgaria",
		Canada: "Canada",
		CanadaFrench: "Canada French",
		Chile: "Chile",
		MainlandChina: "Mainland China",
		Colombia: "Colombia",
		CostaRica: "Costa Rica",
		CroatiaEnglish: "Croatia English",
		Croatia: "Croatia",
		Cyprus: "Cyprus",
		CzechRepublicEnglish: "Czech Republic English",
		CzechRepublic: "Czech Republic",
		Denmark: "Denmark",
		DenmarkEnglish: "Denmark English",
		Ecuador: "Ecuador",
		ElSalvador: "El Salvador",
		FinlandEnglish: "Finland English",
		Finland: "Finland",
		France: "France",
		Germany: "Germany",
		GreeceEnglish: "Greece English",
		Greece: "Greece",
		Guatemala: "Guatemala",
		Honduras: "Honduras",
		HongKongEnglish: "Hong Kong English",
		HongKong: "Hong Kong",
		HongKongTraditionalChinese: "Hong Kong Traditional Chinese",
		HungaryEnglish: "Hungary English",
		Hungary: "Hungary",
		IcelandEnglish: "Iceland English",
		India: "India",
		IndonesiaEnglish: "Indonesia English",
		Ireland: "Ireland",
		IsraelEnglish: "Israel English",
		Israel: "Israel",
		Italy: "Italy",
		Japan: "Japan",
		Korea: "Korea",
		Kuwait: "Kuwait",
		KuwaitEnglish: "Kuwait English",
		Lebanon: "Lebanon",
		LebanonEnglish: "Lebanon English",
		Luxembourg: "Luxembourg",
		LuxembourgFrench: "Luxembourg French",
		MalaysiaEnglish: "Malaysia English",
		Malta: "Malta",
		Mexico: "Mexico",
		Nederland: "Nederland",
		NewZealand: "New Zealand",
		Nicaragua: "Nicaragua",
		NorwayEnglish: "Norway English",
		Norway: "Norway",
		Oman: "Oman",
		OmanEnglish: "Oman English",
		Panama: "Panama",
		Paraguay: "Paraguay",
		Peru: "Peru",
		PhilippinesEnglish: "Philippines English",
		PolandEnglish: "Poland English",
		Poland: "Poland",
		Portugal: "Portugal",
		Qatar: "Qatar",
		QatarEnglish: "Qatar English",
		RomaniaEnglish: "Romania English",
		Romania: "Romania",
		Russia: "Russia",
		SaudiArabia: "Saudi Arabia",
		SaudiArabiaEnglish: "Saudi Arabia English",
		Serbia: "Serbia",
		SingaporeEnglish: "Singapore English",
		SloveniaEnglish: "Slovenia English",
		Slovenia: "Slovenia",
		SlovakiaEnglish: "Slovakia English",
		Slovakia: "Slovakia",
		SouthAfrica: "South Africa",
		Spain: "Spain",
		SwedenEnglish: "Sweden English",
		Sweden: "Sweden",
		Switzerland: "Switzerland",
		SwitzerlandFrench: "Switzerland French",
		SwitzerlandItaliano: "Switzerland Italiano",
		TaiwanEnglish: "Taiwan English",
		Taiwan: "Taiwan",
		ThailandEnglish: "Thailand English",
		Thailand: "Thailand",
		TurkeyEnglish: "Turkey English",
		Turkey: "Turkey",
		UkraineRussian: "Ukraine Russian",
		Ukraine: "Ukraine",
		UnitedArabEmiratesMiddleEast: "United Arab Emirates Middle East",
		UnitedArabEmiratesMiddleEastEnglish: "United Arab Emirates Middle East English",
		UnitedStates: "United States",
		UnitedKingdom: "United Kingdom",
		Uruguay: "Uruguay",
		VietnamEnglish: "Vietnam English",
	}
	return regions
}
