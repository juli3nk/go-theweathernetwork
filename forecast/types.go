package forecast

type CurrentArr []Current

type Current struct {
	LblUpdateTime      string `json:"lbl_updatetime"`       // Updated on
	UpdateTime         string `json:"updatetime"`           // Fri Dec 7 8:35 AM
	UpdateTimeStampGMT string `json:"updatetime_stamp_gmt"` // 1544189700000
	WXCondition        string `json:"wxcondition"`          // A few clouds
	Icon               string `json:"icon"`                 // 2
	Inic               string `json:"inic"`
	Temperature        string `json:"temperature"`      // -9
	FeelsLike          string `json:"feels_like"`       // -15
	TemperatureUnit    string `json:"temperature_unit"` // C
	PlaceCode          string `json:"placecode"`        // CAQC0363
}

type TheWeatherNetwork struct {
	Obs          Obs          `json:"obs"`
	Sterm        Sterm        `json:"sterm"`
	Lterm        Lterm        `json:"lterm"`
	SevenDays    SevenDays    `json:"sevendays"`
	FourteenDays FourteenDays `json:"fourteendays"`
	Reports      Reports      `json:"reports"`
	SWO          SWO          `json:"swo"`
	Bug          Bug          `json:"bug"`
	// NightSky
	// DaySky
	Hourly Hourly `json:"hourly"`
	Type   string `json:"type"`
	Code   string `json:"code"`
}

type Obs struct {
	LabelUpdateTime    string `json:"lbl_updatetime"`       // "Émis le :"
	UpdateTime         string `json:"updatetime"`           // "lun. déc. 5 9:09 AM"
	UpdateTimeStampGMT string `json:"updatetime_stamp_gmt"` // "1480946940000"
	Icon               string `json:"icon"`                 // "16"
	Background         string `json:"background"`           // "snow"
	ImageURL           string `json:"image_url"`            // "//s1.twnmm.com/images/fr_ca/"
	KnotUnit           string `json:"knot_unit"`            // "nd"
	WindSpeedKnot      int    `json:"windSpeed_knot"`       // 4
	WindGustSpeedKnot  int    `json:"windGustSpeed_knot"`   // 6
	SunriseTime        string `json:"sunrise_time"`         // "07:19"
	SunriseGMT         int    `json:"sunrise_gmt"`          // 1480940340
	SunsetTime         string `json:"sunset_time"`          // "16:09"
	SunsetGMT          int    `json:"sunset_gmt"`           // 1480972140
	SunriseSunsetFlag  bool   `json:"sunrisesunset_flag"`   // false
	AltTPC             string `json:"altpc"`                // null
	PollenLevelIcon    string `json:"pollen_level_icon"`    // "noval"
	CeilingFIcon       string `json:"ceiling_f_icon"`       // "low"
	UVLevelIcon        string `json:"uv_level_icon"`        // "noval"
	WindDirectionIcon  string `json:"windDirection_icon"`   // "ne"
	HumidityIcon       string `json:"humidity_icon"`        // "high"
	VisibilityIcon     string `json:"visibility_icon"`      // "low"
	PressureIcon       string `json:"pressure_icon"`        // "medium-high"
	WX                 string `json:"wx"`                   // "S"
	WXC                string `json:"wxc"`                  // "Neige"
	WXCA               string `json:"wxca"`                 // "snow"
	WXSP               string `json:"wxsp"`                 // "PRECIP"
	P                  string `json:"p"`                    // "101.3"
	PT                 string `json:"pt"`                   // "1"
	PU                 string `json:"pu"`                   // "kPa"
	WD                 string `json:"wd"`                   // "N.-E."
	WK                 string `json:"wk"`                   // "7"
	W                  string `json:"w"`                    // "7"
	WU                 string `json:"wu"`                   // "km/h"
	WG                 string `json:"wg"`                   // "11"
	WGU                string `json:"wgu"`                  // "km/h"
	T                  string `json:"t"`                    // "-4"
	F                  string `json:"f"`                    // "-7"
	H                  string `json:"h"`                    // "100"
	TC                 string `json:"tc"`                   // "-4"
	FC                 string `json:"fc"`                   // "-7"
	PK                 string `json:"pk"`                   // "101.3"
	VK                 string `json:"vk"`                   // "1.2"
	V                  string `json:"v"`                    // "1.2"
	VU                 string `json:"vu"`                   // "km"
	CE                 string `json:"ce"`                   // "1100"
	CEU                string `json:"ceu"`                  // "pi"
	TU                 string `json:"tu"`                   // "C"
}

type Period struct {
	N                  string `json:"n"`                   // "1"
	Icon               string `json:"icon"`                // "16"
	SkyTenths          string `json:"sky_tenths"`          // "10"
	FeelsLikeNightUnit string `json:"feelsLikeNight_unit"` // C
	DewptUnit          string `json:"dewpt_unit"`          // "C"
	Hour               string `json:"hour"`                // "11 am"
	CDate              string `json:"cdate"`               // "Monday, December 5"
	DaynameAlt         string `json:"dayname_alt"`         // "Mon"
	RainValue          string `json:"rain_value"`          // -
	SnowValue          string `json:"snow_value"`          // -
	RainBarHeight      int    `json:"rain_bar_height"`     // 0
	SnowBarHeight      int    `json:"snow_bar_height"`     // 0
	ShowRainUnit       bool   `json:"showrainunit"`        // false
	ShowNowUnit        bool   `json:"showsnowunit"`        // true
	PopClass           string `json:"pop_class"`           // "pop7"
	CloudCoverage      int    `json:"cloud_coverage"`      // 100
	CCClass            string `json:"cc_class"`            // "cc10"
	TSG                string `json:"tsg"`                 // "1480953600000"
	TSL                int    `json:"tsl"`                 // 1480935600000
	TC                 int    `json:"tc"`                  // -3
	FC                 int    `json:"fc"`                  // -6
	PP                 string `json:"pp"`                  // "70"
	WD                 string `json:"wd"`                  // "E"
	WK                 string `json:"wk"`                  // "8"
	WGK                string `json:"wgk"`                 // "12"
	R                  string `json:"r"`                   // "-"
	S                  string `json:"s"`                   // "<1 cm"
	RR                 string `json:"rr"`                  // "0"
	SR                 string `json:"sr"`                  // "<1"
	TU                 string `json:"tu"`                  // "C"
	FU                 string `json:"fu"`                  // "C"
	TMU                string `json:"tmu"`                 // "C"
	TMAU               string `json:"tmau"`                // "C"
	T                  string `json:"t"`                   // "-3"
	F                  string `json:"f"`                   // "-6"
	W                  string `json:"w"`                   // "8"
	WU                 string `json:"wu"`                  // "km/h"
	WG                 string `json:"wg"`                  // "12"
	WGU                string `json:"wgu"`                 // "km/h"
	RU                 string `json:"ru"`                  // "mm"
	SU                 string `json:"su"`                  // "cm"
	SD                 string `json:"sd"`                  // "Mon Dec 5"
	DN                 string `json:"dn"`                  // "Mon"
	MS                 string `json:"ms"`                  // "Dec"
	IT                 string `json:"it"`                  // "Scattered flurries"
	II                 string `json:"ii"`                  // "chart-sun"
	IC                 string `json:"ic"`                  // "sunny"
	WX                 string `json:"wx"`                  // "V-"
	B                  string `json:"b"`                   // "default"
}

type Sterm struct {
	Periods              []Period `json:"periods"`
	PopFlag              int      `json:"popflag"`      // 1
	GustFlag             int      `json:"gustflag"`     // 1
	WindGustFlag         int      `json:"windgustflag"` // 1
	WindUnit             string   `json:"wind_unit"`    // km/h
	RainSummary          string   `json:"rainsummary"`
	SnowSummary          string   `json:"snow_summary"` // Snow: 5-10 cm
	RainSummaryShort     string   `json:"rainsummary_short"`
	SnowSummaryShort     string   `json:"snowsummary_short"`      // 5-10 cm
	PrecipOutlookSummary string   `json:"precip_outlook_summary"` // 11:00am Thu to 10:00pm Fri
	IsRain               string   `json:"isRain"`
	IsSnow               string   `json:"isSnow"`
	NoRainSnow           string   `json:"noRainSnow"`
	WXCondLongFlag       int      `json:"wx_cond_long_flag"` // 0
	RF                   int      `json:"rf"`                // 0
	SF                   int      `json:"sf"`                // 1
	RA                   int      `json:"ra"`                // 0
	SA                   float64  `json:"sa"`                // 6.699999999999999
	SS                   string   `json:"ss"`                // Snow: 5-10 cm
}

type Lterm struct {
	TempMaxDay5         int     `json:"tempmax_day5"`            // -9
	SatTemp             string  `json:"sat_temp"`                // -11
	SatCond             string  `json:"sat_cond"`                // snow
	SunTemp             string  `json:"sun_temp"`                // -2
	SunCond             string  `json:"sun_cond"`                // snow
	WXCondLongFlag      int     `json:"wx_cond_long_flag"`       // 0
	WXCondLongFlagNight int     `json:"wx_cond_long_flag_night"` // 0
	RA                  float64 `json:"ra"`                      // 20.599999999999998
	SA                  float64 `json:"sa"`                      // 1.4000000000000001
}

type SevenDays struct {
	Periods                      []Period `json:"periods"`
	ShowRain                     bool     `json:"show_rain"`                         // true
	ShowSnow                     bool     `json:"show_snow"`                         // true
	WindUnit                     string   `json:"wind_unit"`                         // km/h
	PelmMeasurement              string   `json:"pelm_measurement"`                  // metric
	WXCondLongFlag               int      `json:"wx_cond_long_flag"`                 // 0
	WXCondLongFlagNight          int      `json:"wx_cond_long_flag_night"`           // 0
	UseNewPrecipRangesDataAccess bool     `json:"use_new_precip_ranges_dataaccess"`  // false
	UseNewPrecipRangesStaticData bool     `json:"use_new_precip_ranges_static_data"` // false
	RF                           int      `json:"rf"`                                // null
	SF                           int      `json:"sf"`                                // null
	O                            string   `json:"o"`                                 // forecasts: 7 day: caqc0363
}

type FourteenDays struct {
	Periods        []Period `json:"periods"`
	HistoricHigh   string   `json:"historic_high"`     // -2
	HistoricLow    string   `json:"historic_low"`      // -11
	TimePeriod     string   `json:"time_period"`       // Dec 8 - Dec 21
	TempUnit       string   `json:"tempUnit"`          // C
	WindUnit       string   `json:"wind_unit"`         // km/h
	WXCondLongFlag bool     `json:"wx_cond_long_flag"` // false
	ShowRainFlag   bool     `json:"show_rain_flag"`    // true
	ShowSnowFlag   bool     `json:"show_snow_flag"`    // true
	O              string   `json:"o"`                 // forecasts: 14 day: caqc0363
}

type Reports struct {
	PlaceCode     string `json:"placecode"` // caqc0363
	AQExists      bool   `json:"aq_exists"` // true
	AQIndex       string `json:"aq_index"`  // 11
	AQLevel       string `json:"aq_level"`
	UVExists      bool   `json:"uv_exists"`     // true
	LastUV        string `json:"last_uv"`       // 0
	UVIndex       string `json:"uv_index"`      // 0
	UVLabel       string `json:"uv_label"`      // Low
	PollenExists  bool   `json:"pollen_exists"` // false
	PollenIndex   string `json:"pollen_index"`
	PollenName    string `json:"pollen_name"`
	PollenADLevel string `json:"pollen_adlevel"`
	FLULevel      string `json:"flu_level"` // 2
	Alternate     bool   `json:"alternate"` // false
}

type SWO struct {
	SWOLevel string `json:"swo_level"` // 2
	SWOType  string `json:"swo_type"`  // snow
}

type Bug struct {
	BugBlackFly string `json:"bug_black_fly"` // 5
	BugMosquito string `json:"bug_mosquito"`  // 5
	BugDeerFly  string `json:"bug_deer_fly"`  // 5
}

type Hourly struct {
	Periods              []Period `json:"periods"`
	MaxDate              string   `json:"maxDate"`
	MinDate              string   `json:"minDate"`
	HistoricHigh         string   `json:"historic_high"`
	HistoricLow          string   `json:"historic_low"`
	PrecipOutlookSummary string   `json:"precip_outlook_summary"` // 11:00am Thu to 10:00pm Fri
	RainSummaryShort     string   `json:"rainsummary_short"`
	SnowSummaryShort     string   `json:"snowsummary_short"` // 5-10 cm
	WXCondLongFlag       bool     `json:"wx_cond_long_flag"` // false
	RF                   bool     `json:"rf"`
	SF                   bool     `json:"sf"`
	O                    string   `json:"o"`
}

type WeatherText struct {
	Cards []Card `json:"cards"`
}

type Card struct {
	Card           string    `json:"card"`       // afternoon
	Icon           string    `json:"icon"`       // 16
	TempC          string    `json:"temp_c"`     // 0
	TempF          int       `json:"temp_f"`     // 32
	Label          string    `json:"label"`      // This Afternoon
	TimeRange      string    `json:"time_range"` // 12pm – 6pm
	Period1        PeriodBis `json:"period1"`
	Period2        PeriodBis `json:"period2"`
	PhraseMetric   string    `json:"phrase_metric"`   // Cloudy with a few flurries in the afternoon. POP 40%. Snow: <1 cm.
	PhraseImperial string    `json:"phrase_imperial"` // Cloudy with a few flurries in the afternoon. POP 40%. Snow: <0.5 in.
}

type PeriodBis struct {
	TemperatureC  int     `json:"temperature_c"`  // 0
	WindDirection string  `json:"wind_direction"` // SW
	WindSpeed     string  `json:"wind_speed"`     // 18
	WXCode        string  `json:"wxcode"`         // V-
	POP           string  `json:"pop"`            // 40
	PeriodName    string  `json:"period_name"`    // afternoon
	Snow          float64 `json:"snow"`           // 0.4
	Rain          int     `json:"rain"`           // 0
}
