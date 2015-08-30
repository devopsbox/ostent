package params

import (
	"fmt"
	"html/template"
	"math"
	"time"

	"github.com/ostrost/ostent/flags"
)

func (f ParamsFuncs) HrefT(p *Params, num *Num) (template.HTMLAttr, error) {
	href, err := p.EncodeT(num)
	return SprintfAttr(" href=%q", href), err
}

func (f ParamsFuncs) LessD(p *Params, d *Delay, bclass string) (ALink, error) {
	return f.LinkD(p, d, bclass, f.DelayLess(*d, p.MinDelay.Duration), "-")
}

func (f ParamsFuncs) MoreD(p *Params, d *Delay, bclass string) (ALink, error) {
	return f.LinkD(p, d, bclass, f.DelayMore(*d, p.MinDelay.Duration), "+")
}

func (f ParamsFuncs) LessN(p *Params, num *Num, bclass string) (ALink, error) {
	return f.LinkN(p, num, bclass, f.Pow2Less(num.Absolute), "-")
}

func (f ParamsFuncs) MoreN(p *Params, num *Num, bclass string) (ALink, error) {
	return f.LinkN(p, num, bclass, f.Pow2More(num.Absolute), "+")
}

func (f ParamsFuncs) ZeroN(p *Params, num *Num, bclass string) (ALink, error) {
	return f.LinkN(p, num, bclass, 0, "")
}

func (f ParamsFuncs) Vlink(p *Params, num *Num, absolute int, text, alignClass string) (VLink, error) {
	// f is unused
	vl := VLink{LinkText: text, LinkClass: "state"}
	negative := new(bool) // EncodeN will use .Negative being false by default
	if num.Absolute == absolute {
		vl.CaretClass = "caret"
		vl.LinkClass += " current"
		if (num.Alpha && !num.Negative) || (!num.Alpha && num.Negative) {
			vl.LinkClass += " dropup"
		}
		*negative = !num.Negative
	}
	qs, err := p.EncodeN(num, absolute, negative)
	if err != nil {
		return VLink{}, err
	}
	vl.LinkHref = qs
	vl.AlignClass = alignClass
	return vl, nil
}

func (f ParamsFuncs) DelayMore(d Delay, step time.Duration) time.Duration {
	// f is unused
	const s = time.Second
	const m = time.Second * 60
	var table = map[time.Duration]time.Duration{
		s:      2 * s,
		2 * s:  5 * s,
		5 * s:  10 * s,
		10 * s: 30 * s,
		30 * s: m,
		m:      2 * m,
		2 * m:  5 * m,
		5 * m:  10 * m,
		10 * m: 30 * m,
		30 * m: 60 * m,
	}
	if more, ok := table[d.D]; ok {
		return more
	}
	return d.D + step
}

func (f ParamsFuncs) DelayLess(d Delay, step time.Duration) time.Duration {
	// f is unused
	const s = time.Second
	const m = time.Second * 60
	var table = map[time.Duration]time.Duration{
		s:      s,
		2 * s:  s,
		5 * s:  2 * s,
		10 * s: 5 * s,
		30 * s: 10 * s,
		m:      30 * s,
		2 * m:  m,
		5 * m:  2 * m,
		10 * m: 5 * m,
		30 * m: 10 * m,
		60 * m: 30 * m,
	}
	if less, ok := table[d.D]; ok {
		return less
	}
	return d.D - step
}

func (f ParamsFuncs) LinkD(p *Params, d *Delay, bclass string, set time.Duration, badge string) (ALink, error) {
	// f is unused
	href, err := p.EncodeD(d, set)
	if err != nil {
		return ALink{}, err
	}
	var eclass string
	if badge == "-" && d.D == p.MinDelay.Duration {
		eclass = " disabled"
	}
	return ALink{
		Href:       href,
		Text:       flags.DurationString(set),
		Badge:      badge,
		Class:      bclass + " " + eclass,
		ExtraClass: eclass,
	}, nil
}

func (f ParamsFuncs) Pow2Less(v int) int {
	// f is unused
	switch v {
	case 0:
		return 0
	case 1:
		return 0
	case 2:
		return 1
	}
	g := math.Log2(float64(v))
	n := math.Floor(g)
	if n == g {
		n--
	}
	return int(math.Pow(2, n))
}

func (f ParamsFuncs) Pow2More(v int) int {
	// f is unused
	switch v {
	case 0:
		return 1
	case 1:
		return 2
	case 2:
		return 4
	}
	if v <= 32768 { // up to 65536
		v = int(math.Pow(2, 1+math.Floor(math.Log2(float64(v)))))
	}
	return v
}

func (f ParamsFuncs) LinkN(p *Params, num *Num, bclass string, absolute int, badge string) (ALink, error) {
	// f is unused
	href, err := p.EncodeN(num, absolute, nil)
	if err != nil {
		return ALink{}, err
	}
	var eclass string
	if badge == "" && num.Absolute == 0 { // "0" case && param is 0
		eclass = " disabled active"
	}
	if badge == "+" && num.Absolute >= num.Limit && absolute > num.Limit {
		eclass = " disabled"
	}
	if badge == "-" && absolute == 0 {
		eclass = " disabled"
	}
	return ALink{
		Href:       href,
		Text:       fmt.Sprintf("%d", absolute),
		Badge:      badge,
		Class:      bclass + " " + eclass,
		ExtraClass: eclass,
	}, nil
}

type ParamsFuncs struct{}

func SprintfAttr(format string, args ...interface{}) template.HTMLAttr {
	return template.HTMLAttr(fmt.Sprintf(format, args...))
}