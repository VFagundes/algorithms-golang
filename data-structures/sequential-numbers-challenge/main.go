package main

import "fmt"

// split in 8 digits, each group should follow and
// sort the order and the first must be this 11607598 91016305
// 11607598 - 11,607,598
// 11607598 - 11,607,598
var sequentialNumbers string = "1160759861497197987708526328106841360184846531193256198195371330994933652379375761182412725897098630868958517955135143393527988240497195600352417656878439179118928481356923324744742231281257376815044901847344222435563156573094201946142182365231956077456598952242778447482862998509646351400428395929883429819069862881551322892581755310072932233716106245000487642738056582438082450292846585707385259807782566202555791573468830967454828090285194930780284856229788935303186109107577106685113275197037631918192751190379756448503698747188593872861454522434082645450350474806433946295651100122475473510216758785374585394837967574477240250759767644382339565792007832632600447704308317225880234787500175036217120650091528073607008418108040338235385819028002171564139593853854700189812370909311371411296162198967873679789160380593295920426172741759156265003508699612429565953025227740256999077412970160544063372036104943283815989576078004329714119563516108878355966431589117020750823224735746779538223976641977317007399891729491096375"

func main() {
	// ar := make([]string, len(sequentialNumbers)/8)
	// j := 8
	// k := 0
	// for i := 0; i+8 <= len(sequentialNumbers); {
	// 	v := fmt.Sprintf("%s,%s,%s,\n", sequentialNumbers[i:i+2], sequentialNumbers[i+2:i+5], sequentialNumbers[i+5:i+8])
	// 	ar[k] = fmt.Sprintf("%s - %s", sequentialNumbers[i:j], v)
	// 	i += 8
	// 	j += 8
	// 	k++
	// }
	// fmt.Println(ar)
	// fmt.Println(len(ar))
	sequentialNumberFn(sequentialNumbers)
}

func sequentialNumberFn(seqNumbers string) {
	j := len(seqNumbers)
	for i := 0; i < j; i += 8 {
		s := fmt.Sprintf("%s - %s,%s,%s", seqNumbers[i:i+8], seqNumbers[i:i+2], seqNumbers[i+2:i+5], seqNumbers[i+5:i+8])
		fmt.Println(s)
	}

}

// print('Create a function that is able to translate an input number from "0" to "100000" (as string) to its text form:')
// # print('- "1" → "one"')
// # print('- "10" → "ten"')
// # print('- "99" → "ninety nine"')
// # print('- "100000" → "one hundred thousand"')
// numbers_dict ={
// '0':'zero',
// '1':'one',
// '2':'two',
// '3':'three',
// '4':'four',
// '5':'five',
// '6':'six',
// '7':'seven',
// '8':'eight',
// '9':'nine',
// '10':'ten',
// '11':'eleven',
// '12':'twelve',
// '13':'thirteen',
// '14':'fourteen',
// '15':'fifteen',
// '16':'sixteen',
// '17':'seventeen',
// '18':'eighteen',
// '19':'nineteen',
// '20':'twenty',
// '30':'thirty',
// '40':'forty',
// '50':'fifty',
// '60':'sixty',
// '70':'seventy',
// '80':'eighty',
// '90':'ninety',
// '100': 'one hundred',
// '1000':'thousand'

// }
// def mapNumbers(s):
//   if numbers_dict.get(s,None):
//     return numbers_dict.get(s,None)
//   v = int(s)
//   if v > 20
//   # if len(s)==1:
//   #    return numbers_dict.get(s,None)
//   # elif len(s)==3:
//   #   v= int(s)

//   #   result= '{c} hundred and {final} '.format(c=numbers_dict[s[0]],final=numbers_dict[s[1:]])
//   # elif len(s)==4:
//   #   result='thousand '

//   return result

// # print(mapNumbers('120'))
// print(mapNumbers('100'))
