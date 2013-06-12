package langdetect

import (
	//"fmt"
	//"github.com/patrick-higgins/summstat"
	"math"
	"sort"
)

type trigram uint32

var latinTrigrams = []langTrigram{
/*
	{Cs, trigram_cs},
	{Da, trigram_da},
	{De, trigram_de},
	{En, trigram_en},
	{Es, trigram_es},
	{Et, trigram_et},
	{Fi, trigram_fi},
	{Fr, trigram_fr},
	{Hu, trigram_hu},
	{It, trigram_it},
	{Lt, trigram_lt},
	{Lv, trigram_lv},
	{Nl, trigram_nl},
	{Pl, trigram_pl},
	{Pt, trigram_pt},
	{Ro, trigram_ro},
	{Sk, trigram_sk},
	{Sl, trigram_sl},
	{Sv, trigram_sv},
*/
}

type langTrigram struct {
	l Language
	t trigramfreqs
}

func detectLatinByteTrigram(bs []byte) Language {
	vs := trigramfreqs(calcLatinByteTrigram(bs))
	if len(vs) > 500 {
		vs = vs[:499]
	}
	sort.Sort(vs)

	var best Language
	var bestFit = math.Inf(0)
	for _, lt := range latinTrigrams {
		c := trigramfreqs(vs).Diff(lt.t)
		if c < bestFit {
			bestFit = c
			best = lt.l
		}
	}
	return best
}
func calcLatinByteTrigram(bs []byte) trigramfreqpop {
	// Calculate trigrams
	freq := map[trigram]uint32{}
	var acc trigram
	var count uint32
	for _, cur := range bs {
		switch {
		case cur < 'A':
			cur = 0
		case cur <= 'Z':
			cur |= 32
		case cur < 'a':
			cur = 0
		case cur < 'z':
		case cur < 0x80:
			cur = 0
		default:
		}
		if cur == 0 && acc&0xFF == 0 {
			continue
		}
		acc = ((acc & 0xFFFF) << 8) | trigram(cur)
		freq[acc]++
		count++
	}
	//	acc[0], acc[1], acc[2] = acc[1], acc[2], 0
	//	freq[acc]++
	//	acc[0], acc[1], acc[2] = acc[1], acc[2], 0
	//	freq[acc]++
	// Dump
	var vs trigramfreqs
	var cutoff uint32
	switch {
	case len(freq) > 50000:
		cutoff = 3
	case len(freq) > 5000:
		cutoff = 2
	case len(freq) > 500:
		cutoff = 1
	}
	for k, v := range freq {
		if v > cutoff {
			//			fmt.Println(v, calc_normfreq16(v, count), calc_normfreq16(v, count).Float64())
			vs = append(vs, trigramfreq{k, calc_normfreq16(v, count)})
		}
	}
	sort.Sort(trigramfreqpop(vs))
	/*
		sts := summstat.NewStats()
		for _, v := range vs {
			sts.AddSample(summstat.Sample(v.freq.Float64()))
		}
			fmt.Println(count, " ntrigrams ", len(vs), " min ", sts.Min(), " max ", sts.Max(), " mean ", sts.Mean(), " median ", sts.Median(),
				" 100-diff ", vs.Diff(vs[:100]),
				" 500-diff ", vs.Diff(vs[:500]),
				" 999-diff ", vs.Diff(vs[:999]),
			)
	*/
	return trigramfreqpop(vs)
}

func calc_normfreq16(k uint32, n uint32) normfreq16 {
	return normfreq16((uint64(k) << 16) / uint64(n))
}
func (n normfreq16) Float64() float64 {
	return float64(n) / float64(0xFFFF)
}

type normfreq16 uint16
type trigramfreq struct {
	val  trigram
	freq normfreq16
}

type trigramfreqpop trigramfreqs

func (t trigramfreqpop) Len() int      { return len(t) }
func (t trigramfreqpop) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t trigramfreqpop) Less(i, j int) bool {
	return t[i].freq > t[j].freq
}

type trigramfreqs []trigramfreq

func (t trigramfreqs) Len() int { return len(t) }
func (t trigramfreqs) Less(i, j int) bool {
	return t[i].val < t[j].val
}
func (t trigramfreqs) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

/* This is slow...
func (a trigramfreqs) DiffOld(b trigramfreqs) float64 {
	am := map[trigram]float64{}
	bm := map[trigram]float64{}
	//var atotal, btotal float64
	for _, v := range a {
		am[v.val] = v.freq.Float64()
		//		atotal += float64(v.freq)
	}
	for _, v := range b {
		bm[v.val] = v.freq.Float64()
		//		btotal += float64(v.freq)
	}
	var diff float64
	for k, av := range am {
		bv, ok := bm[k]
		if ok {
			delete(bm, k)
		}
		//	x := (av / atotal) - (bv / btotal)
		x := av - bv
		diff += x * x
	}
	for _, bv := range bm {
		x := bv // / btotal
		diff += x * x
	}
	return math.Sqrt(diff)
}
*/
func (a trigramfreqs) Diff(b trigramfreqs) float64 {
	var i, j int
	var diff, x float64
	if len(a) > 0 && len(b) > 0 {
		av := a[0]
		bv := b[0]
		for {
			//		fmt.Printf("%v %v i=%d j=%d => %d\n", av, bv, i, j, trigramcmp(av.val, bv.val))
			switch {
			case av.val == bv.val:
				x = av.freq.Float64() - bv.freq.Float64()
				i++
				j++
				if i == len(a) || j == len(b) {
					goto here
				}
				av = a[i]
				bv = b[j]
			case av.val > bv.val:
				x = bv.freq.Float64()
				j++
				if j == len(b) {
					goto here
				}
				bv = b[j]
			default:
				x = av.freq.Float64()
				i++
				if i == len(a) {
					goto here
				}
				av = a[i]
			}
			diff += x * x
		}
	}
here:
	for ; i < len(a); i++ {
		x := a[i].freq.Float64()
		diff += x * x
	}
	for ; j < len(b); j++ {
		x := b[j].freq.Float64()
		diff += x * x
	}
	return math.Sqrt(diff)
}
