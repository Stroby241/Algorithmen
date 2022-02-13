package main

import "fmt"

func main() {
	input := "Pv ystuxum xfc xbt mtytu jdn eppq. Txfjeyot, iplamjnbfo Hwüwlxvymdi, tla njs hcf ev ocft fyntdiwümtfme bbtu! Yiu cbo, hpu clx.\nXbt hus efc Mdimüdmfmtlna?\nJds bbct mctifc hpdi ycdiu lotqszvjfse, xbifc oocfocohu duhfo hcf ev pm hftnbbgge bbtu.\nRlvß wpx Mdiftwifm"

	findCipherLength(input)

	bruteForceText(input, 4, 6)

	sloveText(input, 5, 20)

	input = "LILKJDXBZEBPKTJVTFGBFQAWFYQAKMFRLGJRPSNTFAOXEQWDYERLZVYCLHAQOUPXUEOJSABLALTRWTGMIROTJDYPHJXBUHKLLEKIPUHKUOCTQPIQZXCCLUNQZFOTSCMJGPBPUHTRYJQPQYBLGBZCWMORKTYOCVBJBBPGXDDBUMKUPXUEPTUMPAOTWPEIOAQEYNMYXTQPBGJAVV"

	findCipherLength(input)

	sloveText(input, 10, 5)

	fmt.Println("Done")
}

// blubb
// xyhtrklboi
