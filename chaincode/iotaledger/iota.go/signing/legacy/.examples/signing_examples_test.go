package signing_examples_test

import (
	"fmt"

	"github.com/iotaledger/iota.go/consts"
	"github.com/iotaledger/iota.go/signing/legacy"
	"github.com/iotaledger/iota.go/signing/utils"
	"github.com/iotaledger/iota.go/trinary"
)

// i req: seed, The seed from which to derive the subseed from.
// i req: index, The index of the subseed.
// i: spongeFunc, The optional sponge function to use.
// o: Trits, The Trits representation of the subseed.
// o: error, Returned for invalid seeds and internal errors.
func ExampleSubseed() {
	seed := "ZLNM9UHJWKTTDEZOTH9CXDEIFUJQCIACDPJIXPOWBDW9LTBHC9AQRIXTIHYLIIURLZCXNSTGNIVC9ISVB"
	subseed, err := signing.Subseed(seed, 0, sponge.NewCurlP81())
	if err != nil {
		// handle error
		return
	}
	fmt.Println(trinary.MustTritsToTrytes(subseed))
	// output: DRUVUYIMTSVUXGUCZA9BRWOGZIEJXE9LXAPOO9OVVWMCYJYJYSZALBHACPPIQBGRCNDWLLBXLOAYWLCRU
}

// i req: subseed, The subseed from which to derive the private key from.
// i req: securityLevel, The used security level.
// i: spongeFunc, The optional sponge function to use.
// o: Trits, The Trits representation of the private key.
// o: error, Returned for internal errors.
func ExampleKey() {
	subseed := "DRUVUYIMTSVUXGUCZA9BRWOGZIEJXE9LXAPOO9OVVWMCYJYJYSZALBHACPPIQBGRCNDWLLBXLOAYWLCRU"
	subseedTrits := trinary.MustTrytesToTrits(subseed)
	key, err := signing.Key(subseedTrits, consts.SecurityLevelLow, sponge.NewCurlP81())
	if err != nil {
		// handle error
		return
	}
	fmt.Println(trinary.MustTritsToTrytes(key))
	// output: HWJVLIDXSNPTMROGNJPBMF9BZBLGJECQZWWWTLAGLGAMLHRWFOXCTZDWMTQSYEVKWBESSWKEWV9AYZHFIRWZRXUZGZQN9KLDLYQCKCGRAXHNFIEJQWFFIHVXDWJIBYLA9BLKFFUQFRBWIXTHDLYSR9KJQWRZUQBFEOJPL9RDDNJCZHAWHCZXRLKYCWCUIODROQA9HDFZFJDJXR9IEILQSYLKXZ9UA9FATTDXOGQL9TOLIHOEAPXYCRKIJNQPTJDOQRKBWYIFIGBQDDMGEU99GSAEVXJ9UTWYTLXHDNRZICLBCIIEGUYEBPAHYY9RIFQUTRVNXSHVKANGTWCVTIWOO9BUPRLLDQ9ZYGKSJUNNTGAFPINPTJXYUTXCLXFGFGDZSOFQ9CTEV9YBXPKKJ9QJFOJHIWGXEMCURNVTJBMKMZ9EYTZ9O9MLZTCHIHFBFZJAT9PEAHGN9RKKCBEBHYLWNSLLHDEAPYFCVYQPXGNORCHXONSDXWVYKDZVIBQWYKSAAQNW9PUTRLGP9B9HEZGHDIHDCOKWDLEITQVM9S9BALPQRJIZFMFXT9FHGEHXJWYSKWYFXOYBIPJUCVJVEAUBULEKKGRIEODONQLU9CQLPDQBBKWIQMDK9ICOXDCYEQUUFMIFLZKUPMOL99LOCLDYEQXYGIKJMCCXN9OTJSSRVLUJNGLKYPJPCDPTYLRQOTTCVHTVGVPGCXCM9AXKPKGFPTX9KFOEWDMJAOSSJ9OKUJXADWZIRHJJQGJ9XZBBEWATROGUIFUEF9JVHU9NF9KQRHTTZIOQPYZEBDMVJZTDKKHAWFNRSWIPDJRIVGIJPUQSKEFBWIGLEFPJSSWQTDVCDOFPYDALFZHZJFKP9SVFERWAZUZSOELYWRDOWHFNXWVVFQTSHQWD9TVICIXQUCKUIPLOLSYJ9SJRMQNSJNPTEEKELNAY9WCTRQSHIHQFYDSV9FTHMUBRFJLCEP9HUBANVOSUVZNVUGAPUZXAOOFMWQWXWNIRMIALUZSSNSWKNQYUOUCPQ99XRTKX9XRVTGVZQVWDSLQPYLJNBZYYTRWAXLGIBHCAKJISJLICHZKZ99PAGHAJVGAOEVAZFOEHOMYEORCHBXSIKBFTANGBNDZTDMKGKUDAREITODOK9NITUDTRQNFWYRRAKGHXDCAHMVXWHCKZWBSAUHQDSENADYUYEGXUISMGAZOIAUPGVSYGUFAZLRASZBLXLRCCGU9YTXF9DFHRHWBE9YZHDOLLJI9OCGUWJDXWOOQZYEHDNGMAVGPJWPJDMQIWDVXRRKPMFJQACJADWXTRTNOSUNCUZD9NIRYCKJXOVQUJPHXPCJMUZAPDIPXQXJGEZWXGJSBSFEVX9REVNPEUMBDEFSXMFECLSD9AHZAXTNUHEHYRQMSBIBCMMSMBNCCNFYNJLCICPTMHATFZNNHBQPNNJJNVFLYUWZIBQUKISWCIVNKBRYUFHSZXNMXSMZGFCPQUMYUSJIKEVKLAQZLEN9VKIEVZUTPKCLJMWETXIUZJRWZGCACNUEOIQDVIZVDHKJMXJERUF9QVRCFOTXGMWMTKZRQCKY9TJD9KMAK9WUKRCWUKHSRQKIYHANCOSOUUGXUIEG9ZOXLVXBELUMBBSCVITQVCNYGOUIHXGZKQ9RMLUSESKQIWHPHEUHIZORWKWZJXPOEWMYAVJYGGHEUQOMIWVALCSNJPSCJUUGAJPNWNPR9NEFYXVNSMXBHRNWOJOKZ9KZCVPBCNVKZRBNVFGZMNBEDEQYSKWLARPJAIBWYGASNUSBCEPGKWVNSRDO9CBDTJXIBUVEPKGKYCRWSABH9MAATCWZYVFDCOUPONSSJPHNDDXVTXYJPXKPQAVHSURLGAYRQRCDAGXFNVYJPUQZS9QKYSQ9RIECWWBCEVCORSLUIEWUIOSOGKYWCQCDEKBEPZDVUBEBMYRSOTTICZGQDHYJFKAOHWKKUCRRGFLINY9XLMXGQZDGGVSXDBPENVCGMMVAVOZVDKOBAULIIYBXPAFKAVOKMQSYHZDFSHLFYSMXIBOZHF9OMMQUJTZGHGO9KOCRR9EUUVH9GFVAKFMBLCCZCPXDCMPRTIIFUZLEDDSLGGWBZQNMEEOCDEGJBBWSPITOHGISBDVMPHGCHNTC9LBKIURKNXXWEANRCF9HLGLXTCINCDDTAXUXOEHFVLEAX
}

// i req: key, The private key from which to derive the digests from.
// i: spongeFunc, The optional sponge function to use.
// o: Trits, The Trits representation of the digests.
// o: error, Returned for internal errors.
func ExampleDigests() {
	key := "HWJVLIDXSNPTMROGNJPBMF9BZBLGJECQZWWWTLAGLGAMLHRWFOXCTZDWMTQSYEVKWBESSWKEWV9AYZHFIRWZRXUZGZQN9KLDLYQCKCGRAXHNFIEJQWFFIHVXDWJIBYLA9BLKFFUQFRBWIXTHDLYSR9KJQWRZUQBFEOJPL9RDDNJCZHAWHCZXRLKYCWCUIODROQA9HDFZFJDJXR9IEILQSYLKXZ9UA9FATTDXOGQL9TOLIHOEAPXYCRKIJNQPTJDOQRKBWYIFIGBQDDMGEU99GSAEVXJ9UTWYTLXHDNRZICLBCIIEGUYEBPAHYY9RIFQUTRVNXSHVKANGTWCVTIWOO9BUPRLLDQ9ZYGKSJUNNTGAFPINPTJXYUTXCLXFGFGDZSOFQ9CTEV9YBXPKKJ9QJFOJHIWGXEMCURNVTJBMKMZ9EYTZ9O9MLZTCHIHFBFZJAT9PEAHGN9RKKCBEBHYLWNSLLHDEAPYFCVYQPXGNORCHXONSDXWVYKDZVIBQWYKSAAQNW9PUTRLGP9B9HEZGHDIHDCOKWDLEITQVM9S9BALPQRJIZFMFXT9FHGEHXJWYSKWYFXOYBIPJUCVJVEAUBULEKKGRIEODONQLU9CQLPDQBBKWIQMDK9ICOXDCYEQUUFMIFLZKUPMOL99LOCLDYEQXYGIKJMCCXN9OTJSSRVLUJNGLKYPJPCDPTYLRQOTTCVHTVGVPGCXCM9AXKPKGFPTX9KFOEWDMJAOSSJ9OKUJXADWZIRHJJQGJ9XZBBEWATROGUIFUEF9JVHU9NF9KQRHTTZIOQPYZEBDMVJZTDKKHAWFNRSWIPDJRIVGIJPUQSKEFBWIGLEFPJSSWQTDVCDOFPYDALFZHZJFKP9SVFERWAZUZSOELYWRDOWHFNXWVVFQTSHQWD9TVICIXQUCKUIPLOLSYJ9SJRMQNSJNPTEEKELNAY9WCTRQSHIHQFYDSV9FTHMUBRFJLCEP9HUBANVOSUVZNVUGAPUZXAOOFMWQWXWNIRMIALUZSSNSWKNQYUOUCPQ99XRTKX9XRVTGVZQVWDSLQPYLJNBZYYTRWAXLGIBHCAKJISJLICHZKZ99PAGHAJVGAOEVAZFOEHOMYEORCHBXSIKBFTANGBNDZTDMKGKUDAREITODOK9NITUDTRQNFWYRRAKGHXDCAHMVXWHCKZWBSAUHQDSENADYUYEGXUISMGAZOIAUPGVSYGUFAZLRASZBLXLRCCGU9YTXF9DFHRHWBE9YZHDOLLJI9OCGUWJDXWOOQZYEHDNGMAVGPJWPJDMQIWDVXRRKPMFJQACJADWXTRTNOSUNCUZD9NIRYCKJXOVQUJPHXPCJMUZAPDIPXQXJGEZWXGJSBSFEVX9REVNPEUMBDEFSXMFECLSD9AHZAXTNUHEHYRQMSBIBCMMSMBNCCNFYNJLCICPTMHATFZNNHBQPNNJJNVFLYUWZIBQUKISWCIVNKBRYUFHSZXNMXSMZGFCPQUMYUSJIKEVKLAQZLEN9VKIEVZUTPKCLJMWETXIUZJRWZGCACNUEOIQDVIZVDHKJMXJERUF9QVRCFOTXGMWMTKZRQCKY9TJD9KMAK9WUKRCWUKHSRQKIYHANCOSOUUGXUIEG9ZOXLVXBELUMBBSCVITQVCNYGOUIHXGZKQ9RMLUSESKQIWHPHEUHIZORWKWZJXPOEWMYAVJYGGHEUQOMIWVALCSNJPSCJUUGAJPNWNPR9NEFYXVNSMXBHRNWOJOKZ9KZCVPBCNVKZRBNVFGZMNBEDEQYSKWLARPJAIBWYGASNUSBCEPGKWVNSRDO9CBDTJXIBUVEPKGKYCRWSABH9MAATCWZYVFDCOUPONSSJPHNDDXVTXYJPXKPQAVHSURLGAYRQRCDAGXFNVYJPUQZS9QKYSQ9RIECWWBCEVCORSLUIEWUIOSOGKYWCQCDEKBEPZDVUBEBMYRSOTTICZGQDHYJFKAOHWKKUCRRGFLINY9XLMXGQZDGGVSXDBPENVCGMMVAVOZVDKOBAULIIYBXPAFKAVOKMQSYHZDFSHLFYSMXIBOZHF9OMMQUJTZGHGO9KOCRR9EUUVH9GFVAKFMBLCCZCPXDCMPRTIIFUZLEDDSLGGWBZQNMEEOCDEGJBBWSPITOHGISBDVMPHGCHNTC9LBKIURKNXXWEANRCF9HLGLXTCINCDDTAXUXOEHFVLEAX"
	keyTrits := trinary.MustTrytesToTrits(key)
	digests, err := signing.Digests(keyTrits, sponge.NewCurlP81())
	if err != nil {
		// handle error
		return
	}
	fmt.Println(trinary.MustTritsToTrytes(digests))
	// output: CPJIVRRPXR9GALMWKJEYDNCVQNAAMTXRDACBTJKKLQOCBSBKJWAJOPDQHBWJCVQN9WOFIRNVDHDJXNGVC
}

// i req: digests, The digests from which to derive the address from.
// i: spongeFunc, The optional sponge function to use.
// o: Trits, The Trits representation of the address.
// o: error, Returned for internal errors.
func ExampleAddress() {
	digests := "CPJIVRRPXR9GALMWKJEYDNCVQNAAMTXRDACBTJKKLQOCBSBKJWAJOPDQHBWJCVQN9WOFIRNVDHDJXNGVC"
	digestsTrits := trinary.MustTrytesToTrits(digests)
	address, err := signing.Address(digestsTrits, sponge.NewCurlP81())
	if err != nil {
		// handle error
		return
	}
	fmt.Println(trinary.MustTritsToTrytes(address))
	// output: SKCPWXDBWKCGIHILLGAXUFFXGMDNVVZYLKZUJIZDCYFDEWDNOIMZEITSQLEOLKBXV9WXOKFKBXIZSGSYL
}

// i req: normalizedBundleHashFragment, The normalized bundle hash.
// i req: keyFragment, The fragment of the private key.
// i req: startOffset, The start offset for the signature chunk.
// i: spongeFunc, The optional sponge function to use.
// o: Trits, The Trits representation of the signed message signature fragment.
// o: error, Returned for internal errors.
func ExampleSignatureFragment() {
	hashSignature := "B9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9SB9YB9YB9YB9YB9YB9YB9YB9YB9E"
	hashSignatureTrits := trinary.MustTrytesToTrits(hashSignature)

	// the private key
	key := "HWJVLIDXSNPTMROGNJPBMF9BZBLGJECQZWWWTLAGLGAMLHRWFOXCTZDWMTQSYEVKWBESSWKEWV9AYZHFIRWZRXUZGZQN9KLDLYQCKCGRAXHNFIEJQWFFIHVXDWJIBYLA9BLKFFUQFRBWIXTHDLYSR9KJQWRZUQBFEOJPL9RDDNJCZHAWHCZXRLKYCWCUIODROQA9HDFZFJDJXR9IEILQSYLKXZ9UA9FATTDXOGQL9TOLIHOEAPXYCRKIJNQPTJDOQRKBWYIFIGBQDDMGEU99GSAEVXJ9UTWYTLXHDNRZICLBCIIEGUYEBPAHYY9RIFQUTRVNXSHVKANGTWCVTIWOO9BUPRLLDQ9ZYGKSJUNNTGAFPINPTJXYUTXCLXFGFGDZSOFQ9CTEV9YBXPKKJ9QJFOJHIWGXEMCURNVTJBMKMZ9EYTZ9O9MLZTCHIHFBFZJAT9PEAHGN9RKKCBEBHYLWNSLLHDEAPYFCVYQPXGNORCHXONSDXWVYKDZVIBQWYKSAAQNW9PUTRLGP9B9HEZGHDIHDCOKWDLEITQVM9S9BALPQRJIZFMFXT9FHGEHXJWYSKWYFXOYBIPJUCVJVEAUBULEKKGRIEODONQLU9CQLPDQBBKWIQMDK9ICOXDCYEQUUFMIFLZKUPMOL99LOCLDYEQXYGIKJMCCXN9OTJSSRVLUJNGLKYPJPCDPTYLRQOTTCVHTVGVPGCXCM9AXKPKGFPTX9KFOEWDMJAOSSJ9OKUJXADWZIRHJJQGJ9XZBBEWATROGUIFUEF9JVHU9NF9KQRHTTZIOQPYZEBDMVJZTDKKHAWFNRSWIPDJRIVGIJPUQSKEFBWIGLEFPJSSWQTDVCDOFPYDALFZHZJFKP9SVFERWAZUZSOELYWRDOWHFNXWVVFQTSHQWD9TVICIXQUCKUIPLOLSYJ9SJRMQNSJNPTEEKELNAY9WCTRQSHIHQFYDSV9FTHMUBRFJLCEP9HUBANVOSUVZNVUGAPUZXAOOFMWQWXWNIRMIALUZSSNSWKNQYUOUCPQ99XRTKX9XRVTGVZQVWDSLQPYLJNBZYYTRWAXLGIBHCAKJISJLICHZKZ99PAGHAJVGAOEVAZFOEHOMYEORCHBXSIKBFTANGBNDZTDMKGKUDAREITODOK9NITUDTRQNFWYRRAKGHXDCAHMVXWHCKZWBSAUHQDSENADYUYEGXUISMGAZOIAUPGVSYGUFAZLRASZBLXLRCCGU9YTXF9DFHRHWBE9YZHDOLLJI9OCGUWJDXWOOQZYEHDNGMAVGPJWPJDMQIWDVXRRKPMFJQACJADWXTRTNOSUNCUZD9NIRYCKJXOVQUJPHXPCJMUZAPDIPXQXJGEZWXGJSBSFEVX9REVNPEUMBDEFSXMFECLSD9AHZAXTNUHEHYRQMSBIBCMMSMBNCCNFYNJLCICPTMHATFZNNHBQPNNJJNVFLYUWZIBQUKISWCIVNKBRYUFHSZXNMXSMZGFCPQUMYUSJIKEVKLAQZLEN9VKIEVZUTPKCLJMWETXIUZJRWZGCACNUEOIQDVIZVDHKJMXJERUF9QVRCFOTXGMWMTKZRQCKY9TJD9KMAK9WUKRCWUKHSRQKIYHANCOSOUUGXUIEG9ZOXLVXBELUMBBSCVITQVCNYGOUIHXGZKQ9RMLUSESKQIWHPHEUHIZORWKWZJXPOEWMYAVJYGGHEUQOMIWVALCSNJPSCJUUGAJPNWNPR9NEFYXVNSMXBHRNWOJOKZ9KZCVPBCNVKZRBNVFGZMNBEDEQYSKWLARPJAIBWYGASNUSBCEPGKWVNSRDO9CBDTJXIBUVEPKGKYCRWSABH9MAATCWZYVFDCOUPONSSJPHNDDXVTXYJPXKPQAVHSURLGAYRQRCDAGXFNVYJPUQZS9QKYSQ9RIECWWBCEVCORSLUIEWUIOSOGKYWCQCDEKBEPZDVUBEBMYRSOTTICZGQDHYJFKAOHWKKUCRRGFLINY9XLMXGQZDGGVSXDBPENVCGMMVAVOZVDKOBAULIIYBXPAFKAVOKMQSYHZDFSHLFYSMXIBOZHF9OMMQUJTZGHGO9KOCRR9EUUVH9GFVAKFMBLCCZCPXDCMPRTIIFUZLEDDSLGGWBZQNMEEOCDEGJBBWSPITOHGISBDVMPHGCHNTC9LBKIURKNXXWEANRCF9HLGLXTCINCDDTAXUXOEHFVLEAX"
	keyTrits := trinary.MustTrytesToTrits(key)

	sigFragment, err := signing.SignatureFragment(hashSignatureTrits, keyTrits, 0, sponge.NewCurlP81())
	if err != nil {
		// handle error
		return
	}
	fmt.Println(trinary.MustTritsToTrytes(sigFragment))
	// output:
	// HGHPPLSOXAXGXYZTPVVMM9WALFW9WCMYETKMXJPZTXUXJHSXECKKSCRGHHIWVUMTUFGVPXJJHGMHBZVRACVLRKGQSRHLWLDHGP9RHGHPRGNVCBWPLHCXAJESPJTGGRWZUEIZRUREXCJ9XJBGSZZVQLMCQJISOHKKOFLSGCDRZNUMHATOHYKBHNLCPU9EOJHEBYQZVULFTN9GWPOMIYVPGR9TNYXFUTED9IUPTORUIMYMUBOWUVDCRMCGZWMKSKPHSTCIRPIBNJJGBRMDSQKPSB9LXMRQNKPQHGXMAAFVRUNS9STMFDYQZQMIXXHCSPOCKQJZDOOQRVXYGVELDPCS9ZVSQBIEXPXMIOOCSTEJJJJZUWAMGHCZRONRWLKZQ9UNWEWSUOXVNUJJBZTBFWFFZTWZHDGSTKTPTWKVLVBND9AHWEHFAORKTOVTCHQDUJHRSKRG99BWIXJGEZVWSYSEBHPVFZZTCCDPDALNGOKXJVZUWRXVAPNQSWQHZKRVMIQJBEQKLJAQVLKJUWBYPMEOIQDYDQEX9OOSRBXSQKJRPWEHXLFSEXEXAUSXYEQRFIZNNMDQSYIQUJAPND9OFKT9AWGQPKYNSSXRSHZNHVOCLWJE9E9DODOUFVLZQLJPNRQZFOFBNMQNFSGASMJYOPPSEGZWEDOSPVUJLPKWM99DCIFWIARHAQDIQEVOIZQNFEUPUBCQNGAVDUGLYEXJNZBDJKJMHQICYILUWCJVRNUOIWBOEHJBQVOFIDWYOLFKCSQEYWYZL9YXDCAULSMBSFHBZWWCDXZAYMPOCRWYYCFFGHH9AVWNUYNEEWWLROPRFDUDOBWAFRVFDVR9KLZGHHBFJBWQVHRMVPTTTXGGRLKEJGIOVTSATGNUTVHKISCYEECQPOQNYXKMCGJCVSXYNDPKQIVRPHUCHXTPFTYSMXVJZPHGWFKIURGWMRSMZFEOZBGQBZHURDSKZVWDSZNPCNLCVCFZAHVITYTJXOUVINWVDUZRSSTMTIYVQ9VXBHNBELRUIRYTSHPBFSPJBWDJNDZK9TKLNAXVNHDMTWBNNEBDF9XTXZDBGGLXMLW9YNNGC9YTVCPKNF9VAAMUENNJU9DLILQNRWMGTVIEZRNAQWYFLWDWZGNPDFCTCOVBEQNWVLLZZXJPIRBQ9VMQGYYNJEXCX9GNFDREEIGLTJMSBLHV9OSACLJKRGAVBEOFMQPQNNYQPBBYLCLTSOEBPZCJNQ9QHCWCRTLGZLJBPJSMHWQZMOKNDGZTBNRTMTQLGAODZSCQCFWGXYDDEHFLUVGMPZYTVHJYHSFSCAONZNVPPXMRBMSWGK9NHQKAUUYMIQHQSSFQDIDVKEHXBWITMCNTAUWJGEXJSGHFIXYMNEAERTRSLEJOGMAMWLYWRTKALYPBNJCPYGDKBCFPNFGZLKGPBLSJLDYDIYTHBRD9PKUFILVXBPMQIVGGTUPEDBQXWNMWLZFLEZQWWAWGCZBADQHEBMEOBDQRRDNQ9ZQCEMNQABFRTFZUIEL9MJXKYWTRIREUQUBOM9NANGBMUGMDNHNIIZPORLQGAK9WAYNRHQRTHRREQWBWGUCWBIHBLFZLFXQMURPDQQLUGBITEJNYL9WJXDPYBEREVUEORHGQRYO9AYBRW9PDPPBQDNMOPJNSUZQLRMJE9YFCSSEHJBRSDDPSDMOF9HBMSYJFTN9X9PHLHQPURKTQYMFDBDTNDNGFXIEDJZBUFCBZ9MAZTYPREKTS99DMLKJPHZE9Z9MOPD9SDEUTJOTIOANGKXLLPB9AIFBSGRFLYIKPLSLKSWDTXAEGKMAFIWLVANHOIBEMLGQWAWNLOFIRCHYZEBOSMZL9EDM9ECSYWGTHIZAFWOQWROMBXMOEIKSZJO9GEDEHSKYCMQICHNGTFOEAGZWIJYCHNHEFZQZQPETINBFFJPMGZYW9HSUFZLQFDMWPECWJHRECAWFZZBSQSGBASBCAWQPELHYFKV9WUWJOISBRULLDLMWLXGSYETLKXTQPVPIWUEYCGKLUOTDMVXJYGHHBMYARWNDFPMGFVOPCAUW9GHFPJZPS9ESPCTFOWYHQWIYAJGCCOJWLQVQZPSIXJFJ9TE9W9RFXYOCKJ9NTGWBTNQLAYTSQASRWAZLFVBMBEGIDSESPQSUIOOVCGVKMVWHIJJVFPMWRIVVSJPRWGKSIQHE
}

// i req: normalizedBundleHashFragment, The fragment of the normalized bundle hash.
// i req: signatureFragment, The signature fragment corresponding to the bundle hash fragment.
// i req: startOffset, The start offset for the signature chunk.
// i: spongeFunc, The optional sponge function to use.
// o: Trits, The Trits representation of the digest.
// o: error, Returned for internal errors.
func ExampleDigest() {
	hashSignature := "B9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9SB9YB9YB9YB9YB9YB9YB9YB9YB9E"
	hashSignatureTrits := trinary.MustTrytesToTrits(hashSignature)

	// the signature fragment
	sigFragment := "HGHPPLSOXAXGXYZTPVVMM9WALFW9WCMYETKMXJPZTXUXJHSXECKKSCRGHHIWVUMTUFGVPXJJHGMHBZVRACVLRKGQSRHLWLDHGP9RHGHPRGNVCBWPLHCXAJESPJTGGRWZUEIZRUREXCJ9XJBGSZZVQLMCQJISOHKKOFLSGCDRZNUMHATOHYKBHNLCPU9EOJHEBYQZVULFTN9GWPOMIYVPGR9TNYXFUTED9IUPTORUIMYMUBOWUVDCRMCGZWMKSKPHSTCIRPIBNJJGBRMDSQKPSB9LXMRQNKPQHGXMAAFVRUNS9STMFDYQZQMIXXHCSPOCKQJZDOOQRVXYGVELDPCS9ZVSQBIEXPXMIOOCSTEJJJJZUWAMGHCZRONRWLKZQ9UNWEWSUOXVNUJJBZTBFWFFZTWZHDGSTKTPTWKVLVBND9AHWEHFAORKTOVTCHQDUJHRSKRG99BWIXJGEZVWSYSEBHPVFZZTCCDPDALNGOKXJVZUWRXVAPNQSWQHZKRVMIQJBEQKLJAQVLKJUWBYPMEOIQDYDQEX9OOSRBXSQKJRPWEHXLFSEXEXAUSXYEQRFIZNNMDQSYIQUJAPND9OFKT9AWGQPKYNSSXRSHZNHVOCLWJE9E9DODOUFVLZQLJPNRQZFOFBNMQNFSGASMJYOPPSEGZWEDOSPVUJLPKWM99DCIFWIARHAQDIQEVOIZQNFEUPUBCQNGAVDUGLYEXJNZBDJKJMHQICYILUWCJVRNUOIWBOEHJBQVOFIDWYOLFKCSQEYWYZL9YXDCAULSMBSFHBZWWCDXZAYMPOCRWYYCFFGHH9AVWNUYNEEWWLROPRFDUDOBWAFRVFDVR9KLZGHHBFJBWQVHRMVPTTTXGGRLKEJGIOVTSATGNUTVHKISCYEECQPOQNYXKMCGJCVSXYNDPKQIVRPHUCHXTPFTYSMXVJZPHGWFKIURGWMRSMZFEOZBGQBZHURDSKZVWDSZNPCNLCVCFZAHVITYTJXOUVINWVDUZRSSTMTIYVQ9VXBHNBELRUIRYTSHPBFSPJBWDJNDZK9TKLNAXVNHDMTWBNNEBDF9XTXZDBGGLXMLW9YNNGC9YTVCPKNF9VAAMUENNJU9DLILQNRWMGTVIEZRNAQWYFLWDWZGNPDFCTCOVBEQNWVLLZZXJPIRBQ9VMQGYYNJEXCX9GNFDREEIGLTJMSBLHV9OSACLJKRGAVBEOFMQPQNNYQPBBYLCLTSOEBPZCJNQ9QHCWCRTLGZLJBPJSMHWQZMOKNDGZTBNRTMTQLGAODZSCQCFWGXYDDEHFLUVGMPZYTVHJYHSFSCAONZNVPPXMRBMSWGK9NHQKAUUYMIQHQSSFQDIDVKEHXBWITMCNTAUWJGEXJSGHFIXYMNEAERTRSLEJOGMAMWLYWRTKALYPBNJCPYGDKBCFPNFGZLKGPBLSJLDYDIYTHBRD9PKUFILVXBPMQIVGGTUPEDBQXWNMWLZFLEZQWWAWGCZBADQHEBMEOBDQRRDNQ9ZQCEMNQABFRTFZUIEL9MJXKYWTRIREUQUBOM9NANGBMUGMDNHNIIZPORLQGAK9WAYNRHQRTHRREQWBWGUCWBIHBLFZLFXQMURPDQQLUGBITEJNYL9WJXDPYBEREVUEORHGQRYO9AYBRW9PDPPBQDNMOPJNSUZQLRMJE9YFCSSEHJBRSDDPSDMOF9HBMSYJFTN9X9PHLHQPURKTQYMFDBDTNDNGFXIEDJZBUFCBZ9MAZTYPREKTS99DMLKJPHZE9Z9MOPD9SDEUTJOTIOANGKXLLPB9AIFBSGRFLYIKPLSLKSWDTXAEGKMAFIWLVANHOIBEMLGQWAWNLOFIRCHYZEBOSMZL9EDM9ECSYWGTHIZAFWOQWROMBXMOEIKSZJO9GEDEHSKYCMQICHNGTFOEAGZWIJYCHNHEFZQZQPETINBFFJPMGZYW9HSUFZLQFDMWPECWJHRECAWFZZBSQSGBASBCAWQPELHYFKV9WUWJOISBRULLDLMWLXGSYETLKXTQPVPIWUEYCGKLUOTDMVXJYGHHBMYARWNDFPMGFVOPCAUW9GHFPJZPS9ESPCTFOWYHQWIYAJGCCOJWLQVQZPSIXJFJ9TE9W9RFXYOCKJ9NTGWBTNQLAYTSQASRWAZLFVBMBEGIDSESPQSUIOOVCGVKMVWHIJJVFPMWRIVVSJPRWGKSIQHE"
	sigFragmentTrits := trinary.MustTrytesToTrits(sigFragment)

	digest, err := signing.Digest(hashSignatureTrits, sigFragmentTrits, 0, sponge.NewCurlP81())
	if err != nil {
		// handle error
		return
	}
	fmt.Println(trinary.MustTritsToTrytes(digest))
	// output:
	// CPJIVRRPXR9GALMWKJEYDNCVQNAAMTXRDACBTJKKLQOCBSBKJWAJOPDQHBWJCVQN9WOFIRNVDHDJXNGVC
}

// i req: expectedAddress, The address to validate against to check whether the signatures are valid.
// i req: fragments, The signed signature fragments.
// i req: bundleHash, The hash of the bundle.
// i: spongeFunc, The optional sponge function to use.
// o: bool, Whether the signatures are valid.
// o: error, Returned for internal errors.
func ExampleValidateSignatures() {
	hashSignature := "B9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9SB9YB9YB9YB9YB9YB9YB9YB9YB9E"
	address := "SKCPWXDBWKCGIHILLGAXUFFXGMDNVVZYLKZUJIZDCYFDEWDNOIMZEITSQLEOLKBXV9WXOKFKBXIZSGSYL"

	// the signature fragment
	sigFragment := "HGHPPLSOXAXGXYZTPVVMM9WALFW9WCMYETKMXJPZTXUXJHSXECKKSCRGHHIWVUMTUFGVPXJJHGMHBZVRACVLRKGQSRHLWLDHGP9RHGHPRGNVCBWPLHCXAJESPJTGGRWZUEIZRUREXCJ9XJBGSZZVQLMCQJISOHKKOFLSGCDRZNUMHATOHYKBHNLCPU9EOJHEBYQZVULFTN9GWPOMIYVPGR9TNYXFUTED9IUPTORUIMYMUBOWUVDCRMCGZWMKSKPHSTCIRPIBNJJGBRMDSQKPSB9LXMRQNKPQHGXMAAFVRUNS9STMFDYQZQMIXXHCSPOCKQJZDOOQRVXYGVELDPCS9ZVSQBIEXPXMIOOCSTEJJJJZUWAMGHCZRONRWLKZQ9UNWEWSUOXVNUJJBZTBFWFFZTWZHDGSTKTPTWKVLVBND9AHWEHFAORKTOVTCHQDUJHRSKRG99BWIXJGEZVWSYSEBHPVFZZTCCDPDALNGOKXJVZUWRXVAPNQSWQHZKRVMIQJBEQKLJAQVLKJUWBYPMEOIQDYDQEX9OOSRBXSQKJRPWEHXLFSEXEXAUSXYEQRFIZNNMDQSYIQUJAPND9OFKT9AWGQPKYNSSXRSHZNHVOCLWJE9E9DODOUFVLZQLJPNRQZFOFBNMQNFSGASMJYOPPSEGZWEDOSPVUJLPKWM99DCIFWIARHAQDIQEVOIZQNFEUPUBCQNGAVDUGLYEXJNZBDJKJMHQICYILUWCJVRNUOIWBOEHJBQVOFIDWYOLFKCSQEYWYZL9YXDCAULSMBSFHBZWWCDXZAYMPOCRWYYCFFGHH9AVWNUYNEEWWLROPRFDUDOBWAFRVFDVR9KLZGHHBFJBWQVHRMVPTTTXGGRLKEJGIOVTSATGNUTVHKISCYEECQPOQNYXKMCGJCVSXYNDPKQIVRPHUCHXTPFTYSMXVJZPHGWFKIURGWMRSMZFEOZBGQBZHURDSKZVWDSZNPCNLCVCFZAHVITYTJXOUVINWVDUZRSSTMTIYVQ9VXBHNBELRUIRYTSHPBFSPJBWDJNDZK9TKLNAXVNHDMTWBNNEBDF9XTXZDBGGLXMLW9YNNGC9YTVCPKNF9VAAMUENNJU9DLILQNRWMGTVIEZRNAQWYFLWDWZGNPDFCTCOVBEQNWVLLZZXJPIRBQ9VMQGYYNJEXCX9GNFDREEIGLTJMSBLHV9OSACLJKRGAVBEOFMQPQNNYQPBBYLCLTSOEBPZCJNQ9QHCWCRTLGZLJBPJSMHWQZMOKNDGZTBNRTMTQLGAODZSCQCFWGXYDDEHFLUVGMPZYTVHJYHSFSCAONZNVPPXMRBMSWGK9NHQKAUUYMIQHQSSFQDIDVKEHXBWITMCNTAUWJGEXJSGHFIXYMNEAERTRSLEJOGMAMWLYWRTKALYPBNJCPYGDKBCFPNFGZLKGPBLSJLDYDIYTHBRD9PKUFILVXBPMQIVGGTUPEDBQXWNMWLZFLEZQWWAWGCZBADQHEBMEOBDQRRDNQ9ZQCEMNQABFRTFZUIEL9MJXKYWTRIREUQUBOM9NANGBMUGMDNHNIIZPORLQGAK9WAYNRHQRTHRREQWBWGUCWBIHBLFZLFXQMURPDQQLUGBITEJNYL9WJXDPYBEREVUEORHGQRYO9AYBRW9PDPPBQDNMOPJNSUZQLRMJE9YFCSSEHJBRSDDPSDMOF9HBMSYJFTN9X9PHLHQPURKTQYMFDBDTNDNGFXIEDJZBUFCBZ9MAZTYPREKTS99DMLKJPHZE9Z9MOPD9SDEUTJOTIOANGKXLLPB9AIFBSGRFLYIKPLSLKSWDTXAEGKMAFIWLVANHOIBEMLGQWAWNLOFIRCHYZEBOSMZL9EDM9ECSYWGTHIZAFWOQWROMBXMOEIKSZJO9GEDEHSKYCMQICHNGTFOEAGZWIJYCHNHEFZQZQPETINBFFJPMGZYW9HSUFZLQFDMWPECWJHRECAWFZZBSQSGBASBCAWQPELHYFKV9WUWJOISBRULLDLMWLXGSYETLKXTQPVPIWUEYCGKLUOTDMVXJYGHHBMYARWNDFPMGFVOPCAUW9GHFPJZPS9ESPCTFOWYHQWIYAJGCCOJWLQVQZPSIXJFJ9TE9W9RFXYOCKJ9NTGWBTNQLAYTSQASRWAZLFVBMBEGIDSESPQSUIOOVCGVKMVWHIJJVFPMWRIVVSJPRWGKSIQHE"

	valid, err := signing.ValidateSignatures(address, sigFragment, hashSignature, sponge.NewCurlP81())
	if err != nil {
		// handle error
		return
	}
	fmt.Println(valid)
	// output: true
}
