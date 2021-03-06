# API -> GetTrytes()
GetTrytes fetches the transaction trytes given a list of transaction hashes.
> **Important note:** This API is currently in Beta and is subject to change. Use of these APIs in production applications is not supported.


## Input

| Parameter       | Type | Required or Optional | Description |
|:---------------|:--------|:--------| :--------|
| hashes | ...Hash | Required | The hashes of the transactions of which to get the Trytes of.  |




## Output

| Return type     | Description |
|:---------------|:--------|
| []Trytes | The Trytes of the requested transactions. |
| error | Returned for internal errors. |




## Example

```go
func ExampleGetTrytes() 
	trytes, err := iotaAPI.GetTrytes("CLXCQVSDAOHWLGKVLNUKKJOOANL9OVGEHSNGRQFLOZJUSJSSXBGJDROUHALTSNUPMTSAVFF9IQEEA9999")
	if err != nil {
		// handle error
		return
	}
	fmt.Println(trytes[0])
	// output: ZYVOTKSGSVLVAHEBG9TQJVBASNIFUXGJVZ9YIQFFJAXLRQDRKA9RXFIDVHONTBJPZ9UCKPGZQQN9BDAKCJCLHR9ZNQTHRYR9GVZXT9JJFLIORBTTXEOWKGKWUVREIFGF9NVYSCUKFCSMAERHGCYDNYFX9SYBMDVXJ9QE9MCZLDFBOOASQELBVCQKGKDBEYGFOOJKNGZ9CICUBHZMCGJGGLVXIKEDDWUCWDCMPMIMNVBXYYNQISWNRGZRKSEFHB9USXTIEWOYXDPDFAKFEMS99NTDZRSAJSBSOCKWVVJIOHHENIGLNNSPFUAPNCDULETAQNNDOF9FGFYTOSMFYSKLLPLBDRASEDUJOBBEETIIDD9BJWF9MLWZNXN9XZSCJZHQWFVYPBGMNKDHIXZHLFIEXL9FPUUX9KARDTNGVDZNBPCBNTSGKUUFAJDZGPGTWRKXEYRBYFPAZPKZ9XZLVUKTYUPKXUBCSTHJPABBSZJNJDVYMLRDGBGQWCANSZIJCQAPL9LBVHTHHFXFUDZKKFV9O9IWKYPK9JQVLGXVQPUCFFFTEHLDEIROIQWKEKVMPAUTAHVUQXHXRIDQNGYVFLPZHOBAQFIGEIIZSRWLETZVWVLOYF9JTEVPEMJHMXAUHQSEOOBOMJYAEGBSASSF9BHWCOCUOYTKH99RISXDKIACLXNEFWKYJD9U9KPCFSMFGYDMLFKJNIMTKJPJGZLKUZVHLYTWACYTCLPSHXARNMMGEZQLVAYJOSJIVNXJRSPRFHVPIHULYBXEWHIUYOGZNWCRIUSZRYGGZALGDZNZXA99BD9IQEWEBXALRAVOUKDAMBENLTTERRNJEEPYLG9V9MABZOWWFTMFMRWWUQIFROCCUBOKABIIPDCRBP9QEHZRHRASPBEALUBPIRTCMYRIKZTTCLYCLVOJEZEFLYRXABVPZWCKWJJJEQLWCDBMLSQTEJRPJITMJHBAAKD9SSKQHTFZWEOAVDERPLOTOK9EXJBXSZABZUCMCTEJY9P9CZKYLMXRZBLX9PRKVRUMVURTVTUDJUEWPDDBCVCCPIFAUOXOK9AAMFDSPKEGOMHSVESUUEFIPSGQTEUWDHTAPRTOGEPVNFWIHMSJVDEKCOKV9PWDQNFAYXQJHVABMVVNNISXPQYQHSIDFYOMDAHRAPCZRQOQONXSRDDUYKHF9WKBWXXSQWKZKWP9NEEKNXPOQLETRUWXMMRZ9SHTO9IAHKUOMXI9EMUEMNAKXZETHMXGMIIPZQY9WXYFHBKMLVHRSEPZVNIHPUNPXQCKAKWYTOAASXNCAHGCOTUGSHYURUJXPRBF9RMHQB9AKBUEZNTYUFPIQQKYBLMDKPMVV9CRUQGOIWCWJPNUYFYZ9FO9DKWDEKD9ELAEFCXQUOOZGPVRXBGNAHQSMUZRUENYKFNWWUATSASTKKP9HNA9LCJEBKFUOAMBKKDGSP9WAWJWAEUTF9MIEBF9MFOCYZYAGSN9VDOWUYD9SHIYBVSG9CJQJSLCVMSAUJFZCCLEDHI9IXGDNILDWAJRRAAJNAGCSARCRQBXKARTSZCALDSAXXZCXSEGWWAOYCMQU9L9KYJCHSJJAEBXJZAORMXISBMKJYEZEMGKVMB9YZSAXMEQUHHABSXROERMLBHKDAHOSVRGJWDDZTDKRAJUWSBSNVASMFVHNDSYBEPZOSCMZQNQRVLVUJKKCESMJGMKYNEDYNDENMSZRGPYZFXFQLKVTYTUZJSUYZBBPGKIUDJFCXCSYXJJFLVEYUDBVISZUNINDTXUJXNCYSXHLGNIWMZEZPAVSU9YDBMSVIGKRHAMS9UPLIZANLFUESVTRFVWP9SPRPD9F9ILJIXMFRWKZYCZUIDNPNGFUJLOEQWFUJROZTOILCC9NYZAMKNZKIFQBIEJZCYPOVDGUBPSETFRXVTIDKCTY9TBBRUPFMXMNHY9REKOYVAAZFFGC9D9BPSVRYKNRVUMLNPVYS9FNSPGJMYLSPNYSKJSLEIJPJSPZJDRDYUEWKRJBOOZKAODYUQEBZRRLEOGUDTSIELWKCDNTVQJCGWIRZCNJXYSSDG9ELNRNGDBAVCFFWSAJWZKMRBWVCOFROYTDNKYJBRYUR9KDMPU9JDATLLIFXULLZOAIKWZOZYWIKSYCNUHAZKCRDIKCJQDWADJTYJYRKXES9UKNDFOFUGZHOAAETRKPXQNEK9GWM9ILSODEOZEFDDROCNKYQLWBDHWAEQJIGMSOJSETHNAMZOWDIVVMYPOPSFJRZYMDNRDBSCZ99999999999999999999999GC9999999999999999999999999CBDAMZD99999999999I999999999RVBBAEIONQBWLFIOLFQTEETLYCULDNPK9LIUAE9EIDUU9VDVL9OD9LHLEPYFEDJPFHOSTZCBILVKCRTYXOOGSSQFKDLHWUNQGIWJSYAL9AUMKINSYUKTMZBZKXVETCJLPX9D9KWPBZYUYFTFWOFFLIVKLCFDA9999FAPCHJMEDCUAIYOILWUTGFNTI9UWKTQKNMUEFQCNICFZNUSKPVIXVKSEIPSGEXHIKMWYCORQXYGD99999GC9999999999999999999999999SSNZBWLLE999999999MMMMMMMMMQFAVPOJVPZDLNJUHF9HSFWJUQDD
}

```
