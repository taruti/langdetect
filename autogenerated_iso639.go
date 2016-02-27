package langdetect

var (
	Aa = Language{[2]byte{'a', 'a'}, 0, 0}
	Ab = Language{[2]byte{'a', 'b'}, 1, 0}
	Ae = Language{[2]byte{'a', 'e'}, 2, 0}
	Af = Language{[2]byte{'a', 'f'}, 3, 0}
	Ak = Language{[2]byte{'a', 'k'}, 4, 0}
	Am = Language{[2]byte{'a', 'm'}, 5, 0}
	An = Language{[2]byte{'a', 'n'}, 6, 0}
	Ar = Language{[2]byte{'a', 'r'}, 7, 0}
	As = Language{[2]byte{'a', 's'}, 8, 0}
	Av = Language{[2]byte{'a', 'v'}, 9, 0}
	Ay = Language{[2]byte{'a', 'y'}, 10, 0}
	Az = Language{[2]byte{'a', 'z'}, 11, 0}
	Ba = Language{[2]byte{'b', 'a'}, 12, 0}
	Be = Language{[2]byte{'b', 'e'}, 13, 0}
	Bg = Language{[2]byte{'b', 'g'}, 14, 0}
	Bh = Language{[2]byte{'b', 'h'}, 15, 0}
	Bi = Language{[2]byte{'b', 'i'}, 16, 0}
	Bm = Language{[2]byte{'b', 'm'}, 17, 0}
	Bn = Language{[2]byte{'b', 'n'}, 18, 0}
	Bo = Language{[2]byte{'b', 'o'}, 19, 0}
	Br = Language{[2]byte{'b', 'r'}, 20, 0}
	Bs = Language{[2]byte{'b', 's'}, 21, 0}
	Ca = Language{[2]byte{'c', 'a'}, 22, 0}
	Ce = Language{[2]byte{'c', 'e'}, 23, 0}
	Ch = Language{[2]byte{'c', 'h'}, 24, 0}
	Co = Language{[2]byte{'c', 'o'}, 25, 0}
	Cr = Language{[2]byte{'c', 'r'}, 26, 0}
	Cs = Language{[2]byte{'c', 's'}, 27, 0}
	Cu = Language{[2]byte{'c', 'u'}, 28, 0}
	Cv = Language{[2]byte{'c', 'v'}, 29, 0}
	Cy = Language{[2]byte{'c', 'y'}, 30, 0}
	Da = Language{[2]byte{'d', 'a'}, 31, 0}
	De = Language{[2]byte{'d', 'e'}, 32, 0}
	Dv = Language{[2]byte{'d', 'v'}, 33, 0}
	Dz = Language{[2]byte{'d', 'z'}, 34, 0}
	Ee = Language{[2]byte{'e', 'e'}, 35, 0}
	El = Language{[2]byte{'e', 'l'}, 36, 0}
	En = Language{[2]byte{'e', 'n'}, 37, 0}
	Eo = Language{[2]byte{'e', 'o'}, 38, 0}
	Es = Language{[2]byte{'e', 's'}, 39, 0}
	Et = Language{[2]byte{'e', 't'}, 40, 0}
	Eu = Language{[2]byte{'e', 'u'}, 41, 0}
	Fa = Language{[2]byte{'f', 'a'}, 42, 0}
	Ff = Language{[2]byte{'f', 'f'}, 43, 0}
	Fi = Language{[2]byte{'f', 'i'}, 44, 0}
	Fj = Language{[2]byte{'f', 'j'}, 45, 0}
	Fo = Language{[2]byte{'f', 'o'}, 46, 0}
	Fr = Language{[2]byte{'f', 'r'}, 47, 0}
	Fy = Language{[2]byte{'f', 'y'}, 48, 0}
	Ga = Language{[2]byte{'g', 'a'}, 49, 0}
	Gd = Language{[2]byte{'g', 'd'}, 50, 0}
	Gl = Language{[2]byte{'g', 'l'}, 51, 0}
	Gn = Language{[2]byte{'g', 'n'}, 52, 0}
	Gu = Language{[2]byte{'g', 'u'}, 53, 0}
	Gv = Language{[2]byte{'g', 'v'}, 54, 0}
	Ha = Language{[2]byte{'h', 'a'}, 55, 0}
	He = Language{[2]byte{'h', 'e'}, 56, 0}
	Hi = Language{[2]byte{'h', 'i'}, 57, 0}
	Ho = Language{[2]byte{'h', 'o'}, 58, 0}
	Hr = Language{[2]byte{'h', 'r'}, 59, 0}
	Ht = Language{[2]byte{'h', 't'}, 60, 0}
	Hu = Language{[2]byte{'h', 'u'}, 61, 0}
	Hy = Language{[2]byte{'h', 'y'}, 62, 0}
	Hz = Language{[2]byte{'h', 'z'}, 63, 0}
	Ia = Language{[2]byte{'i', 'a'}, 64, 0}
	Id = Language{[2]byte{'i', 'd'}, 65, 0}
	Ie = Language{[2]byte{'i', 'e'}, 66, 0}
	Ig = Language{[2]byte{'i', 'g'}, 67, 0}
	Ii = Language{[2]byte{'i', 'i'}, 68, 0}
	Ik = Language{[2]byte{'i', 'k'}, 69, 0}
	Io = Language{[2]byte{'i', 'o'}, 70, 0}
	Is = Language{[2]byte{'i', 's'}, 71, 0}
	It = Language{[2]byte{'i', 't'}, 72, 0}
	Iu = Language{[2]byte{'i', 'u'}, 73, 0}
	Ja = Language{[2]byte{'j', 'a'}, 74, 0}
	Jv = Language{[2]byte{'j', 'v'}, 75, 0}
	Ka = Language{[2]byte{'k', 'a'}, 76, 0}
	Kg = Language{[2]byte{'k', 'g'}, 77, 0}
	Ki = Language{[2]byte{'k', 'i'}, 78, 0}
	Kj = Language{[2]byte{'k', 'j'}, 79, 0}
	Kk = Language{[2]byte{'k', 'k'}, 80, 0}
	Kl = Language{[2]byte{'k', 'l'}, 81, 0}
	Km = Language{[2]byte{'k', 'm'}, 82, 0}
	Kn = Language{[2]byte{'k', 'n'}, 83, 0}
	Ko = Language{[2]byte{'k', 'o'}, 84, 0}
	Kr = Language{[2]byte{'k', 'r'}, 85, 0}
	Ks = Language{[2]byte{'k', 's'}, 86, 0}
	Ku = Language{[2]byte{'k', 'u'}, 87, 0}
	Kv = Language{[2]byte{'k', 'v'}, 88, 0}
	Kw = Language{[2]byte{'k', 'w'}, 89, 0}
	Ky = Language{[2]byte{'k', 'y'}, 90, 0}
	La = Language{[2]byte{'l', 'a'}, 91, 0}
	Lb = Language{[2]byte{'l', 'b'}, 92, 0}
	Lg = Language{[2]byte{'l', 'g'}, 93, 0}
	Li = Language{[2]byte{'l', 'i'}, 94, 0}
	Ln = Language{[2]byte{'l', 'n'}, 95, 0}
	Lo = Language{[2]byte{'l', 'o'}, 96, 0}
	Lt = Language{[2]byte{'l', 't'}, 97, 0}
	Lu = Language{[2]byte{'l', 'u'}, 98, 0}
	Lv = Language{[2]byte{'l', 'v'}, 99, 0}
	Mg = Language{[2]byte{'m', 'g'}, 100, 0}
	Mh = Language{[2]byte{'m', 'h'}, 101, 0}
	Mi = Language{[2]byte{'m', 'i'}, 102, 0}
	Mk = Language{[2]byte{'m', 'k'}, 103, 0}
	Ml = Language{[2]byte{'m', 'l'}, 104, 0}
	Mn = Language{[2]byte{'m', 'n'}, 105, 0}
	Mr = Language{[2]byte{'m', 'r'}, 106, 0}
	Ms = Language{[2]byte{'m', 's'}, 107, 0}
	Mt = Language{[2]byte{'m', 't'}, 108, 0}
	My = Language{[2]byte{'m', 'y'}, 109, 0}
	Na = Language{[2]byte{'n', 'a'}, 110, 0}
	Nb = Language{[2]byte{'n', 'b'}, 111, 0}
	Nd = Language{[2]byte{'n', 'd'}, 112, 0}
	Ne = Language{[2]byte{'n', 'e'}, 113, 0}
	Ng = Language{[2]byte{'n', 'g'}, 114, 0}
	Nl = Language{[2]byte{'n', 'l'}, 115, 0}
	Nn = Language{[2]byte{'n', 'n'}, 116, 0}
	No = Language{[2]byte{'n', 'o'}, 117, 0}
	Nr = Language{[2]byte{'n', 'r'}, 118, 0}
	Nv = Language{[2]byte{'n', 'v'}, 119, 0}
	Ny = Language{[2]byte{'n', 'y'}, 120, 0}
	Oc = Language{[2]byte{'o', 'c'}, 121, 0}
	Oj = Language{[2]byte{'o', 'j'}, 122, 0}
	Om = Language{[2]byte{'o', 'm'}, 123, 0}
	Or = Language{[2]byte{'o', 'r'}, 124, 0}
	Os = Language{[2]byte{'o', 's'}, 125, 0}
	Pa = Language{[2]byte{'p', 'a'}, 126, 0}
	Pi = Language{[2]byte{'p', 'i'}, 127, 0}
	Pl = Language{[2]byte{'p', 'l'}, 128, 0}
	Ps = Language{[2]byte{'p', 's'}, 129, 0}
	Pt = Language{[2]byte{'p', 't'}, 130, 0}
	Qu = Language{[2]byte{'q', 'u'}, 131, 0}
	Rm = Language{[2]byte{'r', 'm'}, 132, 0}
	Rn = Language{[2]byte{'r', 'n'}, 133, 0}
	Ro = Language{[2]byte{'r', 'o'}, 134, 0}
	Ru = Language{[2]byte{'r', 'u'}, 135, 0}
	Rw = Language{[2]byte{'r', 'w'}, 136, 0}
	Sa = Language{[2]byte{'s', 'a'}, 137, 0}
	Sc = Language{[2]byte{'s', 'c'}, 138, 0}
	Sd = Language{[2]byte{'s', 'd'}, 139, 0}
	Se = Language{[2]byte{'s', 'e'}, 140, 0}
	Sg = Language{[2]byte{'s', 'g'}, 141, 0}
	Si = Language{[2]byte{'s', 'i'}, 142, 0}
	Sk = Language{[2]byte{'s', 'k'}, 143, 0}
	Sl = Language{[2]byte{'s', 'l'}, 144, 0}
	Sm = Language{[2]byte{'s', 'm'}, 145, 0}
	Sn = Language{[2]byte{'s', 'n'}, 146, 0}
	So = Language{[2]byte{'s', 'o'}, 147, 0}
	Sq = Language{[2]byte{'s', 'q'}, 148, 0}
	Sr = Language{[2]byte{'s', 'r'}, 149, 0}
	Ss = Language{[2]byte{'s', 's'}, 150, 0}
	St = Language{[2]byte{'s', 't'}, 151, 0}
	Su = Language{[2]byte{'s', 'u'}, 152, 0}
	Sv = Language{[2]byte{'s', 'v'}, 153, 0}
	Sw = Language{[2]byte{'s', 'w'}, 154, 0}
	Ta = Language{[2]byte{'t', 'a'}, 155, 0}
	Te = Language{[2]byte{'t', 'e'}, 156, 0}
	Tg = Language{[2]byte{'t', 'g'}, 157, 0}
	Th = Language{[2]byte{'t', 'h'}, 158, 0}
	Ti = Language{[2]byte{'t', 'i'}, 159, 0}
	Tk = Language{[2]byte{'t', 'k'}, 160, 0}
	Tl = Language{[2]byte{'t', 'l'}, 161, 0}
	Tn = Language{[2]byte{'t', 'n'}, 162, 0}
	To = Language{[2]byte{'t', 'o'}, 163, 0}
	Tr = Language{[2]byte{'t', 'r'}, 164, 0}
	Ts = Language{[2]byte{'t', 's'}, 165, 0}
	Tt = Language{[2]byte{'t', 't'}, 166, 0}
	Tw = Language{[2]byte{'t', 'w'}, 167, 0}
	Ty = Language{[2]byte{'t', 'y'}, 168, 0}
	Ug = Language{[2]byte{'u', 'g'}, 169, 0}
	Uk = Language{[2]byte{'u', 'k'}, 170, 0}
	Ur = Language{[2]byte{'u', 'r'}, 171, 0}
	Uz = Language{[2]byte{'u', 'z'}, 172, 0}
	Ve = Language{[2]byte{'v', 'e'}, 173, 0}
	Vi = Language{[2]byte{'v', 'i'}, 174, 0}
	Vo = Language{[2]byte{'v', 'o'}, 175, 0}
	Wa = Language{[2]byte{'w', 'a'}, 176, 0}
	Wo = Language{[2]byte{'w', 'o'}, 177, 0}
	Xh = Language{[2]byte{'x', 'h'}, 178, 0}
	Yi = Language{[2]byte{'y', 'i'}, 179, 0}
	Yo = Language{[2]byte{'y', 'o'}, 180, 0}
	Za = Language{[2]byte{'z', 'a'}, 181, 0}
	Zh = Language{[2]byte{'z', 'h'}, 182, 0}
	Zu = Language{[2]byte{'z', 'u'}, 183, 0}
)
var lcodesString = "aaabaeafakamanarasavayazbabebgbhbibmbnbobrbscacechcocrcscucvcydadedvdzeeeleneoeseteufafffifjfofrfygagdglgngugvhahehihohrhthuhyhziaidieigiiikioisitiujajvkakgkikjkkklkmknkokrkskukvkwkylalblglilnloltlulvmgmhmimkmlmnmrmsmtmynanbndnengnlnnnonrnvnyocojomorospapiplpsptqurmrnrorurwsascsdsesgsiskslsmsnsosqsrssstsusvswtatetgthtitktltntotrtstttwtyugukuruzvevivowawoxhyiyozazhzu"
