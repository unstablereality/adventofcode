package main

import (
	"fmt"
	"strings"
)

func main() {
	idList := []string{"mvgowxqubnhaefjslkjlrptzyi",
		"pvgowlqubnhaefmslkjdrpteyi",
		"ovgowoqubnhaefmslkjnrptzyi",
		"cvgowxqubnrxefmslkjdrptzyo",
		"cvgowxqubnhaefmsokjdrprzyf",
		"cvgowxqubnhjeflslkjgrptzyi",
		"cvgowxqvbnhaefmslkhdrotzyi",
		"hvgowxqmbnharfmslkjdrptzyi",
		"cvgoaxqubqhaefmslkjdrutzyi",
		"cvxowxqdbnhaefmslkjdgptzyi",
		"cvgikxqubnhaefmslkjdrptzyz",
		"cvgnwxqubnhaqfjslkjdrptzyi",
		"cqgowxqubnhaecmslkjgrptzyi",
		"cvpowxqucnhaefmslkjdrptzyz",
		"fvuoexqubnhaefmslkjdrptzyi",
		"svgowxqubnhaefmsvkjdrttzyi",
		"cvgowxqubnhaefmblkjdfpbzyi",
		"cvkoyxqubnhaefsslkjdrptzyi",
		"bvgowxqublhaefmslkjdrptzfi",
		"xvgewxqubnhaefmslkjdrztzyi",
		"cvgowxqubzhaefmslkkrrptzyi",
		"cvgowxqubnhaefmslkudruuzyi",
		"cvgowxqubnhaefmvlkjdrptwyl",
		"cvgoyxqubnhaefmslkjvrotzyi",
		"cvgowxoubnhaewmslkjdrpbzyi",
		"cvgowxgubnhaefmslijdrptzxi",
		"lvgowxqkbnhaefmslkjdrptzqi",
		"xvgowxqubyhaefmflkjdrptzyi",
		"wvnowxgubnhaefmslkjdrptzyi",
		"cvgowxguwnhaefhslkjdrptzyi",
		"cvgowfquxnhaefmdlkjdrptzyi",
		"cvgywxqubnuaefmsldjdrpfzyi",
		"cvkowxqzbrhaefmslkjdrptzyi",
		"cviowxzubnhaefmslkjdrptqyi",
		"cvgowxqubnhaefmsozjdrptzyc",
		"cvglwxuubnhaewmslkjdrptzyi",
		"cvgowxquknhaebmsfkjdrptzyi",
		"vvgowxqubnhaesmslkjdrptzri",
		"cvgowxoubndaefmslkjdrftzyi",
		"cvgowxqubghaefmslkjdeptzyw",
		"cvgowxqubnhaetmhlkjdrpvzyi",
		"cvgowmquunhaefmslkjdrptzyt",
		"cvgooxqpbniaefmslkjdrptzyi",
		"cvgowxqubnhaeumslkjdkptiyi",
		"cvgrwxqsbnhaemmslkjdrptzyi",
		"cvrowxqubnhaefmslkjdrctcyi",
		"dvgcwxqubnhaefmslkjdrptzyq",
		"cugowxqubnhasfmmlkjdrptzyi",
		"cwgowxqobzhaefmslkjdrptzyi",
		"cvgowxquwnhaefmulkjdrptbyi",
		"nvgowxqmbnhaefmslyjdrptzyi",
		"cvgowxqubniakvmslkjdrptzyi",
		"cvyowxqubnhaefmslejdrptzyx",
		"cvgobxqubghaefeslkjdrptzyi",
		"cvgowxiubnhaebmslkjdfptzyi",
		"cvgosbqubnhaefmslkvdrptzyi",
		"cvgpwxqubnhaefvslkjdrptzyh",
		"cvgowxqubnyaefmslgjdsptzyi",
		"cvgowxqubnhaefmslkjdrprzzp",
		"cvgowxqubwhaemmslkjdrpazyi",
		"cvgowxqpbnhaemmslkjdrpczyi",
		"cvgoqxqubnhaelmslkjdrptzye",
		"cvgowxqubnhaefmslbjdrttzvi",
		"cvgowxqubnhlefmslkvurptzyi",
		"cvgowxqujngaefmslktdrptzyi",
		"cvgowxqubnhaefmsckjdcwtzyi",
		"cvcowxqubnhaetmslkjorptzyi",
		"jvnowxqubnhaefmslkjdrptzyf",
		"cygowxqkbnhaefmslejdrptzyi",
		"cvmowxqubnhaefmslkjdritzoi",
		"cvgowxqubnpaefmslkjdrpnnyi",
		"cvgowxqubnhaefmolkjdrpnzyy",
		"uvgowxoubnhaefmslkjdrptzvi",
		"cvgowxbabehaefmslkjdrptzyi",
		"cvgokxqubnhaefmsckjdrjtzyi",
		"cvgoxwqubahaefmslkjdrptzyi",
		"cvgowxqusnhaefmslijdrptyyi",
		"cvgowxqubmhaeqmslkxdrptzyi",
		"cvgouxhubnhaefmslkjdrjtzyi",
		"cvgowxqubnhaefmslrjdqptzyk",
		"cvgowxiublhaefsslkjdrptzyi",
		"cvgowxqubnxgefmslkadrptzyi",
		"ovgowxqugshaefmslkjdrptzyi",
		"cvgowxquznhaeemslsjdrptzyi",
		"cvkowxqubnhaeomslkjdeptzyi",
		"cvgvwxqubxhaefmslkjdrptzyu",
		"cvglwxqybnhaefmslkjdrptzyb",
		"cvgowxqubnlfwfmslkjdrptzyi",
		"cvaowxqubnhaefmslkjdrvtzbi",
		"cvgowxqubnrmefaslkjdrptzyi",
		"cvgowxqubnhaefmsnkjdfpwzyi",
		"cvgawxqmbnhaefmsykjdrptzyi",
		"chgowmqubnhaefmslkjdrptwyi",
		"cogowxqubnhaefmslkjxrptzri",
		"cvgohxqubnoaesmslkjdrptzyi",
		"cvdowxqubnhaofmslkjdrpvzyi",
		"vvgowrqubnhaefmslkjdrpthyi",
		"cvgowxquknhuefmslkjdoptzyi",
		"cvyowxeubnhaefmslhjdrptzyi",
		"cvglwxqubnhaefmslkjdrptdyq",
		"cvgowxqubnhaefmsikgdrptayi",
		"cvgowxqubnhaefjhlkjdrpczyi",
		"cvgzwxkubnhaefmslkjdjptzyi",
		"cxgowxqubnhaefmslkjdrptwyy",
		"cvgowxqubnhaefeslkjdrmqzyi",
		"cvgowxvubnhaefmilijdrptzyi",
		"cvgowxqzbthaeomslkjdrptzyi",
		"cvgowhqubndaefmglkjdrptzyi",
		"cvgowxvubnhaeamylkjdrptzyi",
		"cvgowiqubnhgefmslkjdrctzyi",
		"cvgowxqubchaefmslksdritzyi",
		"cvgowxqubnhaefmsnkjdreyzyi",
		"cvgowxqubihaefmslkgdrutzyi",
		"cvgowxqjbnhaeamslkjdrptzwi",
		"cvgowxzubnhaefmsxkjdrrtzyi",
		"cvgowxqubyhaetmslnjdrptzyi",
		"cvgowxquhnhaebmslkjdxptzyi",
		"cvgowxqubnhanfmslujdxptzyi",
		"cvgowxqublhnefaslkjdrptzyi",
		"cvgmwxqtbnhaefmslkjsrptzyi",
		"jvgowxqubnhaeamslkjdrpmzyi",
		"cvgowxqubhiaefmsljjdrptzyi",
		"svgowxqubnhaefmswkjdrpozyi",
		"cvgowxqebnhaeqmslkjdiptzyi",
		"cveowxqubnhayzmslkjdrptzyi",
		"cvglwxqubnhaefmxlkjdiptzyi",
		"cvgowkqubdhaefmszkjdrptzyi",
		"cvgowxkxbnhaeffslkjdrptzyi",
		"cugowxqubnnaefmslujdrptzyi",
		"cqgowxwubnhaepmslkjdrptzyi",
		"cvgowxqubnhayfmmlkjwrptzyi",
		"cvgowxquenhaefmsskxdrptzyi",
		"cvgowxqubnhiefmsrkjdtptzyi",
		"mvgowxkubnhaefmjlkjdrptzyi",
		"cvgowkquunhaefmglkjdrptzyi",
		"cvgowxqubqhaexmslgjdrptzyi",
		"jvgowxqubnhaefmslkjddptlyi",
		"cvgiwxqubnhaefmslkjdpptmyi",
		"czgowxqubntaevmslkjdrptzyi",
		"cvgotmqubnhaefmslkjdrpazyi",
		"cvgowxtubnhaefmslkqdtptzyi",
		"cvbowxqhnnhaefmslkjdrptzyi",
		"cvgowkqubshaefmstkjdrptzyi",
		"cvgowqqrbnaaefmslkjdrptzyi",
		"cvgoixqubnhaefmslkjdrpmryi",
		"cvgoxxqubnhaeimsxkjdrptzyi",
		"cvgowxqubzhaebmslkjyrptzyi",
		"cjgewxqubnhaefsslkjdrptzyi",
		"cvgowxqdbnkaefmslwjdrptzyi",
		"cvgowxqzbnhaeamslkjdrftzyi",
		"cvgoixqubnsaewmslkjdrptzyi",
		"cvgswxqubnhaxfmslkjdrptzni",
		"cvwowxmubnhgefmslkjdrptzyi",
		"cvggwxqubnhaefmslqjdbptzyi",
		"cvgzwxqjbnhaefaslkjdrptzyi",
		"cvgowzqubnharfmspkjdrptzyi",
		"cvgowxqubnhawfmslkjdeptzyb",
		"cvuowequbnhaefmslkjdrntzyi",
		"gvgowxqubnxaefmslkjdrjtzyi",
		"cvgowxqubnhmetmsldjdrptzyi",
		"cvgowxqubnhamfmsqkjdrptyyi",
		"cvgoqxqubnhaefmslkjtrpazyi",
		"cvgoexqubhhaefmslkjdrhtzyi",
		"cvgowwqubnhaeflslkjdrptzyf",
		"cvgowlpubnhaefmslkjdrptvyi",
		"cvgowxouunhaebmslkjdrptzyi",
		"cvdowhqubnhaefmslijdrptzyi",
		"cvgowxqubnkatfmslkjdrhtzyi",
		"cvgowxqpbnhxeumslkjdrptzyi",
		"cvgowxqubnhaefmsukjjrptzyn",
		"cvgowxqubnhmefmslzjdrvtzyi",
		"cvtowxqubihaefmclkjdrptzyi",
		"chgowcqubnhayfmslkjdrptzyi",
		"cvguwxqubnhaefmblkjarptzyi",
		"cvgowoqubnhaefmsikjdrytzyi",
		"cvgkwxqubnhaefmslkjdrptchi",
		"cvhowxqubnhaefmslkjdrvlzyi",
		"cvlowxfubnhaefmslkjkrptzyi",
		"cvgowxqubhhaefoslkjdrytzyi",
		"cvgowxsubqhaefmslpjdrptzyi",
		"cvgowxpubnhaefmslhjdrptzyb",
		"cvgowxqubnhrefmjlkddrptzyi",
		"cvgowxqubnhaxfmykkjdrptzyi",
		"mvgowxqubnhakfmslkjdrptnyi",
		"cwgowxqubnhaffmslkadrptzyi",
		"chgowxquwnhaefmslsjdrptzyi",
		"cvgowxqubnhaefmslkjdwpnsyi",
		"cvgawxqubnhaefmslkldyptzyi",
		"cvgowxqubnhiefmslkjdiprzyi",
		"cvgkqxqubnhaefcslkjdrptzyi",
		"cvgovoqubnhaefmslkjdrpuzyi",
		"cvgowxqubnhaefmszkjdrjtzyk",
		"cvgopxqubnhaefmslkjdqpnzyi",
		"cvgtwxqubnhaefmslkjnrptzri",
		"cvgowxqurnhaedmslfjdrptzyi",
		"cvpowxqubnhaefmswkjdrltzyi",
		"cvgowxqujnpaefmslkjdrptdyi",
		"cvgowgqubnhzifmslkjdrptzyi",
		"lvgowxqubnhaenmslkjdbptzyi",
		"ebgowxqubnhaeymslkjdrptzyi",
		"cvgowxtubqhaefmslkedrptzyi",
		"cvgowxqubshaesmslkjdrptryi",
		"cvgowxqubnhaefmflkjmrpkzyi",
		"cvgowxqubngaefmslkjdrytzgi",
		"cvgowxqubnhaefmslklhzptzyi",
		"cveowxqubnhgefmslkjdrpezyi",
		"cvgowxqubnhaeomslkjdrqtzym",
		"cvgowxqubzhaefmslwjdrptfyi",
		"cmgowxqubnhaefmsdkjdrptzui",
		"cvlowxqubnhaefmslsjdrptzwi",
		"cvhowxpubnhaefmslkjhrptzyi",
		"cveosxqurnhaefmslkjdrptzyi",
		"cvgowxqubnhaefgsdkjdrptjyi",
		"cvgvwxqubnhaefmslzjdmptzyi",
		"cviowxqubnhalfmslkjdrptzyr",
		"cvgowxqubchqefmslkjdrptzoi",
		"cvgownqubnhaefmsyktdrptzyi",
		"cvgywxqubnuaefmslkjdrpfzyi",
		"cvgobxqunnhaefmslkjdrptzbi",
		"cvgowxqubshaefgslkjdrxtzyi",
		"cvghwxqubnhaefmslkjdrbtmyi",
		"cvhowxqubnhaefmslkjdrpnzys",
		"cvgowxqubnmaefmslejdrptzyq",
		"cvmrwxqubnhaefmslkjdrpzzyi",
		"cvgowxqubshaefmslkfdrptzyu",
		"cvgowqqubnhaefmslkodrpjzyi",
		"cvgnwnquknhaefmslkjdrptzyi",
		"cvgowxquxnhacfmflkjdrptzyi",
		"ovgowxqubnhaefmslkjmrmtzyi",
		"cvgowxqubneaefmslkedrptzqi",
		"cvgowxqubphweflslkjdrptzyi",
		"cvgowxqudnhaefmplkjdrptdyi",
		"cvwowxbubnhaefmslkjurptzyi",
		"cvgowxtubnhaefmslkjdrwwzyi",
		"cvgowxqubnhkefmslajdrptzyn",
		"cvgowxqxbphaefmslkjdrptzsi",
		"cvgowxquenhaefmslmjwrptzyi",
		"zvgowdqubnhaeftslkjdrptzyi",
		"csgowxqubnhgefmslkjdrptzyy",
		"cvgolxqubahaefmslkjdrpvzyi",
		"cvgoqxquhwhaefmslkjdrptzyi",
		"cvgawxqubghaefmsrkjdrptzyi",
		"cvgozxqubnhaefmslkwdfptzyi",
		"cvgowxqubnhaefmslhjdkptzzi",
		"cvnowxqubnhaefmsqkjdrptqyi",
		"cvpowxqubnhaefmslkpdrptdyi",
		"cvgowxoubnhaermslkjdrctzyi",
		"cvgowxqubnheefmslkjdrctzyr",
		"cvgowxqunnhaqfhslkjdrptzyi",
		"cvgowxqulnhaefmslrjdrntzyi"}

	twos := 0
	threes := 0

	for _, id := range idList {
		chars := strings.Split(id, "")
		counter := make(map[string]int)
		for _, row := range chars {
			counter[row]++
		}
		two := 0
		three := 0
		for k := range counter {
			if counter[k] == 2 {
				two = 1
			}
			if counter[k] == 3 {
				three = 1
			}
		}
		if two > 0 {
			twos++
		}
		if three > 0 {
			threes++
		}
	}

	fmt.Println("Checksum:", twos*threes)
}
