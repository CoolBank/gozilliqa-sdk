package mpt

import (
	"github.com/Zilliqa/gozilliqa-sdk/core"
	"github.com/Zilliqa/gozilliqa-sdk/util"
	"testing"
)

func TestVerify(t *testing.T) {
	key := []byte("aea3a3f1944d56741a36d05444955722e63bcea9")
	proofString := []string{
		"F851808080A0A69293FB75DAC6C34B35755C1313555FCE878A4561ACD342BFEF11E55E901F128080A0E391AE68269693AFB48CFEE8160C2FF452F3E0633E84F662E47E9722E317FDA780808080808080808080",
		"F8D180A0187829E431B80BCD24E0E088503CFF9268637F528780E5524E76EF4E6833C4C4A0EBFF5548879664503991B9243840978FD3373B267889613932B89F95CC471BEFA0FB59525B37DD9A1A972B12DFB3607F442E061C0CEB99D6256FB1EA963BB9A548A0B823D0591EF501E4C06730894881F02AA2EE15DAD3F6ACDFE5BDCB0BA6C3079FA0504B6F2652F6CC24D18F85D32E2FBD08E6E4AFE65F95C6B46FABBFF4C6AAD4EEA0D6E5581333D6E5262F6CD0273F375876C4D6BF181C193182E9838CA6B00D562C80808080808080808080",
		"F851808080A0ADFFF78854A513EC2C05EA38AFF47547920116D26A2FFB7566F4CB421408991C8080A0DC84861CF96E4472F4C5D7416079ECFBEBA65C3EA6557A6FEDE36AB87A18429F80808080808080808080",
		"F8D180A00BA1F6615371DA24BBB653524A90516D2C573FCA697AEF763AF715FE24A7A316A03D3417199D4E673733A56C276C10A4F179C440BDA2E2FBBD3625A44A6F1A9DF0A0953A0951EC60E0280A14D696B096CAF3F67E37C8973AB577A7116499F9FA52E1A01233CBC176869A3B79D07934F59954CE61320A3B04DC9FFF01A9153F9F19FFA5A0857E70EA23EF5CF62B0373C53D4372328BBEA747271DE6B0527BD9D6D47EEC71A0EF5ADA17CC7FC0C0DD573894F3277D4297D456AFAAE64751EC814056994745C780808080808080808080",
		"F851808080A0EFE48B122E4D20DF542DAFB39A23E60BB4F27C6715D172131EB19593BDD4F3E08080A014CC77D374C9E69C11BDDBC1EC4DBFA7E7EE5430C1C6658479DD110A81BE34B280808080808080808080",
		"F8D180A01C647175F09CD1D022C72457C74B7C102D1F097376A717EFA133FBC39E2F07D6A08A1920FFED18B9BDD9515201441C4563E6791322B16B3E1248A09C5A55F23D0FA0E0CB1DE45A265C15549BFCDCCB0A4527BBF0118B33CA27433F6B230C4661D5C7A0BC6852CB86AC55079E8789952B3C844FFA0ACEE260388FA728734EC7A55D98D4A0C9CBE398FE005368DE25F22CD9E0805B91DDF2D85B5CC911EBAB959DF7EF767CA0D9D4C91BF6E3B7924D1997506575904A4BE7BC04390C862A63D87555F4CD94CC80808080808080808080",
		"F851808080A0E36A38396C13ED0502C485077636DBE006B26845F688D4D9174D44374FACBE638080A083553240D5BEC6A46B28221C510AB30503D67B24F2EC622E4441600EE9F3EF7C80808080808080808080",
		"F871A0850F06D0B6E75F43F8444086D86ACB15046D1AEC8EC95B4F313626C8D209DAD98080A0C291F02DE8FF524CAFFA91E259AF4968EB8D44E8B37C47B94E99C3BC7E6AD13D808080A0B5600F3259C04A103A2B198634A09EAF01F50E14F86DB4E412E0507CFBCDA5BC808080808080808080",
		"EBA5206133663139343464353637343161333664303534343439353537323265363362636561398439393939",
	}
	var proof [][]byte

	for _, p := range proofString {
		bytes := util.DecodeHex(p)
		proof = append(proof, bytes)
	}

	db := NewFromProof(proof)

	root := util.DecodeHex("c3a31a414d193aaecfa06113341e137e12f111232f3f3b7842ec3c2ec5201200")
	value, err := Verify(key, db, root)
	if err != nil {
		t.Fatal(err.Error())
	}

	if string(value) != "9999" {
		t.Fail()
	}

	// key is contract address
	key = []byte("5050065f52bd935f9fe58937986f74373657a7fd")

	// account proof
	proofString = []string{
		"F851808080A0258AD6A826EC3C56976157B95C26A3E5E8523C554D4593F00860BC04D72A942D8080A02D58A9EA855D237C2D63613FC2D47C23EDA1A9DA361E7398B79805185579F96E80808080808080808080",
		"F887A820303530303635663532626439333566396665353839333739383666373433373336353761376664B85C080112120A100000000000000000000000000000000018002220F74E858D851B7035161C66546FC183A5B162A8EE187D10324ACB1FA8CF1391EA2A20F95D81F1E266A74B57E3BD6EC484AC9C1B2A006A23A6F3A911CE4CFE73ECD335",
		"F90151A052A5B4DEEEB4B6E2AA92857C7A02469BC551D9DBDDDA4600070FB2145C8DE8D9A08EDA68D6EAEEDE4A09E6B446AAA0F07A06C22E7830120466F05A963347CF087AA0335FB20C21D6ED59D2D8A56FC68D7CE6548EE0B3B2D7F5DB0A0E2DD396C7707EA008F2F7A3CD199E00DD4E278AA1A415CBAE1B73DB4B6F0B9D5F46C39ED3C7DA1DA0D3072B01F14552C0BEC3B77992997BC6DDB0C80B9C4EC6E3A9B609E8C7B677A2A0AEE7D93841FD1BAC793D534C501E8990AC2B3706356E44944A4C67F557283C7CA0AD772F8A49B208371AEE4D48F28DEC21390F49B0673D008886AED07016366BC8A0E25B6D16EB2DF15D1C673ACA4A5D1ACD625CAF8BA3E781BAB5ED71B6DCDB5BCEA01FFDF308CB6099DD8E1CFC0C75885E97DC197F2EB759BBD9514D4F663FFB536AA0A9E55373C48E44E3DC5290CED3A47ED538419AC8818B4549A68B952E259BA0AD80808080808080",
	}
	var proof2 [][]byte

	for _, p := range proofString {
		bytes := util.DecodeHex(p)
		proof2 = append(proof2, bytes)
	}

	db = NewFromProof(proof2)

	// StateRootHash from tx block
	root = util.DecodeHex("609c073d333b0ac7a499194e845980a8e54e76c7d1e70393e7544fa9646ad3f5")
	value, err1 := Verify(key, db, root)
	if err1 != nil {
		t.Fatal(err1.Error())
	}
	t.Log(util.EncodeHex(value))

	storageKey := core.GenerateStorageKey([]string{"zilToPolyTxHashMap", "\"0\""})
	t.Log(string(storageKey))

	// state proof
	proofString3 := []string{
		"F18F1696C546F506F6C795478486173684A04E1EA726F02730A94E371861892D2CF725010406DAF6BE9EE99500F484F286A0",
		"F83C808080808080808080CB86206E6465781683223122808080A0A3661324C976C4ABADF6D5ADA95CF2E1722A33EF90BEDF89E6D32BB197173574808080",
		"F84F882061701622302216B8442230783431663131383930376638666366353764343237636438343065356133393761383235396637373262656163333061653837376535303331353938353663326222",
		"F8718080808080A04BC50E95A1993F7F52113AFD46F720E13B9B88ECB59C5C3A30DBFB60A796FF61A0F3737A3EF468414F5F562C85218B13ACC9096333E34D014B8B8160ABDE562821A055E353BA4597779E736AFFC344B99090759987832CC9F4FA810DC9C5BED761B7808080808080808080",
		"F871A0C7F6B93AA37F85E335907E25C3303CF7717A656124AF9826C8264BC9AB3DEDCF8080A06D7499E4F652A962E5E99F6B2AE90BF4CA68B4EF98BBD7EC373E15BFF813CAD0808080808080A037E96A42801D1DAD8005E6B0EC7AE950150F7C03E20E723345D5E6BCDDF21D8D808080808080",
	}
	var proof3 [][]byte

	for _, p := range proofString3 {
		bytes := util.DecodeHex(p)
		proof3 = append(proof3, bytes)
	}

	db3 := NewFromProof(proof3)

	accountBase, _ := core.AccountBaseFromBytes(util.DecodeHex("080112120a100000000000000000000000000000000018002220f74e858d851b7035161c66546fc183a5b162a8ee187d10324acb1fa8cf1391ea2a20f95d81f1e266a74b57e3bd6ec484ac9c1b2a006a23a6f3a911ce4cfe73ecd335"))
	root3 := accountBase.StorageRoot
	value3, err2 := Verify(storageKey, db3, root3)

	if err2 != nil {
		t.Fatal(err2.Error())
	}

	t.Log(string(value3))
}

func TestVerify2(t *testing.T) {
	// key is contract address
	key := []byte("5050065f52bd935f9fe58937986f74373657a7fd")

	// account proof
	proofString := []string{
		"F851808080A098A0ACA40BC5C34789E28ECD75B0AD70DEAB48A9264CD0663F71076EF1014C5B8080A00F7C4DD7427A214DC2C7624425A934F30D06B2D5538EDC23EF4AEB5B9E8F790B80808080808080808080",
		"F887A820303530303635663532626439333566396665353839333739383666373433373336353761376664B85C080112120A100000000000000000000000000000000018002220F74E858D851B7035161C66546FC183A5B162A8EE187D10324ACB1FA8CF1391EA2A20F95D81F1E266A74B57E3BD6EC484AC9C1B2A006A23A6F3A911CE4CFE73ECD335",
		"F90151A011A5C350A90B513436E9C81D1810583791FF5EDB786A02C3EC1A0CDFC2BFBDDEA08EDA68D6EAEEDE4A09E6B446AAA0F07A06C22E7830120466F05A963347CF087AA077837D479E924BEC3353B8AA2372EFA3C5EB2ADD25CA269582A15FA50BFAA3C1A00C36669CD33ACB905ABFC65500FC528D49F4F0AA6C998DF5FC98164E653EF871A0D3072B01F14552C0BEC3B77992997BC6DDB0C80B9C4EC6E3A9B609E8C7B677A2A0AEE7D93841FD1BAC793D534C501E8990AC2B3706356E44944A4C67F557283C7CA0AD772F8A49B208371AEE4D48F28DEC21390F49B0673D008886AED07016366BC8A0028B86C2F75E4C18DB82CC6EB4C2E6F658CEA781A563CA1D4942F890CC593806A0482111D892AA258CCF6CD63F35DEFB31E6C0E754B9F9B37D716290A06BF52BDBA0EB552B368038F46CA72677668B9881DEDFECC5819B95B2954E2339CE9AA4D79F80808080808080",
	}
	var proof [][]byte

	for _, p := range proofString {
		bytes := util.DecodeHex(p)
		proof = append(proof, bytes)
	}

	db := NewFromProof(proof)

	// StateRootHash from tx block
	root := util.DecodeHex("d74a231c5368145131da1ac1cc9a606e073a4a8680628347878478289bbc78ba")
	value, err1 := Verify(key, db, root)
	if err1 != nil {
		t.Fatal(err1.Error())
	}
	t.Log(util.EncodeHex(value))

	storageKey := core.GenerateStorageKey([]string{"zilToPolyTxHashIndex"})
	t.Log(storageKey)
	key3 := []byte(storageKey)

	// state proof
	proofString3 := []string{
		"CB86206E6465781683223122",
		"F18F1696C546F506F6C795478486173684A04E1EA726F02730A94E371861892D2CF725010406DAF6BE9EE99500F484F286A0",
		"F83C808080808080808080CB86206E6465781683223122808080A0A3661324C976C4ABADF6D5ADA95CF2E1722A33EF90BEDF89E6D32BB197173574808080",
		"F8718080808080A04BC50E95A1993F7F52113AFD46F720E13B9B88ECB59C5C3A30DBFB60A796FF61A0F3737A3EF468414F5F562C85218B13ACC9096333E34D014B8B8160ABDE562821A055E353BA4597779E736AFFC344B99090759987832CC9F4FA810DC9C5BED761B7808080808080808080",
		"F871A0C7F6B93AA37F85E335907E25C3303CF7717A656124AF9826C8264BC9AB3DEDCF8080A06D7499E4F652A962E5E99F6B2AE90BF4CA68B4EF98BBD7EC373E15BFF813CAD0808080808080A037E96A42801D1DAD8005E6B0EC7AE950150F7C03E20E723345D5E6BCDDF21D8D808080808080",
	}
	var proof3 [][]byte

	for _, p := range proofString3 {
		bytes := util.DecodeHex(p)
		proof3 = append(proof3, bytes)
	}

	db3 := NewFromProof(proof3)

	accountBase, _ := core.AccountBaseFromBytes(util.DecodeHex("080112120a100000000000000000000000000000000018002220f74e858d851b7035161c66546fc183a5b162a8ee187d10324acb1fa8cf1391ea2a20f95d81f1e266a74b57e3bd6ec484ac9c1b2a006a23a6f3a911ce4cfe73ecd335"))
	root3 := accountBase.StorageRoot
	value3, err2 := Verify(key3, db3, root3)

	if err2 != nil {
		t.Fatal(err2.Error())
	}

	t.Log(string(value3))
}
