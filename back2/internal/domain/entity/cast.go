package entity

import (
	"encoding/json"
	"time"
)

type Cast struct {
	ID              int64       `xorm:"pk autoincr"`
	GroupID         int64       `xorm:"notnull"`
	CastID          string      `xorm:"varchar(7) notnull unique"`
	LastName        string      `xorm:"varchar(30) notnull"`
	FirstName       string      `xorm:"varchar(30) notnull"`
	CastName        string      `xorm:"varchar(30) notnull"`
	Birthdate       time.Time   `xorm:"date" json:"birthdate" time_format:"2006-01-02"`
	PasswordHash    string      `xorm:"varchar(255) notnull"`
	LineID          string      `xorm:"varchar(30) unique notnull"`
	EffectiveDate   time.Time   `xorm:"date"`                            // 証明書有効期限
	TattooFlg       string      `xorm:"varchar(1) notnull default('0')"` // 刺青フラグ
	TattooArea      string      `xorm:"varchar(255) "`                   // 刺青部位
	Allergy         string      `xorm:"varchar(255) "`                   // アレルギー
	StretchMarksFlg string      `xorm:"varchar(1) notnull default('0')"` // ストレッチマークフラグ
	SmokingFlg      string      `xorm:"varchar(1) notnull default('0')"` // 喫煙フラグ
	ForeignerFlg    string      `xorm:"varchar(1) notnull default('0')"` // 外人フラグ
	WorkingFlg      string      `xorm:"varchar(1) notnull default('0')"` // 勤務フラグ
	CreatedAt       time.Time   `xorm:"created"`
	UpdatedAt       time.Time   `xorm:"updated"`
	IsDeleted       string      `xorm:"varchar(1) notnull default('0')"` //削除フラグ
	OptionFlags     OptionFlags `json:"option_flags"`
}

type OptionFlags uint32

const (
	OptionFlagsNone OptionFlags = 1 << iota
	Option1
	Option2
	Option3
	Option4
	Option5
	Option6
	Option7
	Option8
	Option9
	Option10
	Option11
	Option12
	Option13
	Option14
	Option15
	Option16
	Option17
	Option18
	Option19
	Option20
	Option21
	Option22
	Option23
	Option24
	Option25
	Option26
	Option27
	Option28
	Option29
)

var flagMap = map[OptionFlags]string{
	Option1:  "Dキス",
	Option2:  "全身リップ",
	Option3:  "生フェラ",
	Option4:  "アナル舐め",
	Option5:  "玉舐め",
	Option6:  "69",
	Option7:  "指入れ",
	Option8:  "口内発射",
	Option9:  "素股",
	Option10: "バック素股",
	Option11: "パイズリ",
	Option12: "耳かき",
	Option13: "全身マッサージ",
	Option14: "ピンクローター",
	Option15: "下着持ち帰り",
	Option16: "オナニー鑑賞",
	Option17: "バイブ",
	Option18: "即尺",
	Option19: "顔射",
	Option20: "聖水",
	Option21: "ゴックン",
	Option22: "逆AF",
	Option23: "AF",
	Option24: "3P・4P",
	Option25: "パンスト破り",
	Option26: "電マ",
	Option27: "放尿",
	Option28: "写メ",
	Option29: "コスプレ",
}

func (f OptionFlags) MarshalJSON() ([]byte, error) {
	flags := make(map[string]bool)
	for flag, name := range flagMap {
		if f&flag != 0 {
			flags[name] = true
		} else {
			flags[name] = false
		}
	}
	return json.Marshal(flags)
}

func (f *OptionFlags) UnmarshalJSON(data []byte) error {
	var flags map[string]bool
	if err := json.Unmarshal(data, &flags); err != nil {
		return err
	}

	*f = 0
	for name, value := range flags {
		for flag, flagName := range flagMap {
			if name == flagName && value {
				*f |= flag
				break
			}
		}
	}
	return nil
}

// フラグの操作メソッド
func (co *Cast) SetFlag(flag OptionFlags) {
	co.OptionFlags |= flag
}

func (co *Cast) ClearFlag(flag OptionFlags) {
	co.OptionFlags &^= flag
}

func (co *Cast) HasFlag(flag OptionFlags) bool {
	return co.OptionFlags&flag != 0
}
