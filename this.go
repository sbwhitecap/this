package this

import (
	"strings"
	"fmt"
)

var (
	s = "Gur Mra bs Clguba, ol Gvz Crgref\n" +
	    "\n" +
	    "Ornhgvshy vf orggre guna htyl.\n" +
	    "Rkcyvpvg vf orggre guna vzcyvpvg.\n" +
	    "Fvzcyr vf orggre guna pbzcyrk.\n" +
	    "Pbzcyrk vf orggre guna pbzcyvpngrq.\n" +
	    "Syng vf orggre guna arfgrq.\n" +
	    "Fcnefr vf orggre guna qrafr.\n" +
	    "Ernqnovyvgl pbhagf.\n" +
	    "Fcrpvny pnfrf nera'g fcrpvny rabhtu gb oernx gur ehyrf.\n" +
	    "Nygubhtu cenpgvpnyvgl orngf chevgl.\n" +
	    "Reebef fubhyq arire cnff fvyragyl.\n" +
	    "Hayrff rkcyvpvgyl fvyraprq.\n" +
	    "Va gur snpr bs nzovthvgl, ershfr gur grzcgngvba gb thrff.\n" +
	    "Gurer fubhyq or bar-- naq cersrenoyl bayl bar --boivbhf jnl gb qb vg.\n" +
	    "Nygubhtu gung jnl znl abg or boivbhf ng svefg hayrff lbh'er Qhgpu.\n" +
	    "Abj vf orggre guna arire.\n" +
	    "Nygubhtu arire vf bsgra orggre guna *evtug* abj.\n" +
	    "Vs gur vzcyrzragngvba vf uneq gb rkcynva, vg'f n onq vqrn.\n" +
	    "Vs gur vzcyrzragngvba vf rnfl gb rkcynva, vg znl or n tbbq vqrn.\n" +
	    "Anzrfcnprf ner bar ubaxvat terng vqrn -- yrg'f qb zber bs gubfr!"
)

func this() string {
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <='Z':
			return 'A' + (r-'A'+13) % 26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13) % 26
		}
		return r
	}

	return strings.Map(rot13, s)
}

func init() {
	fmt.Println(this())
}
