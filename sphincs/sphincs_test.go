package sphincs

import (
	"testing"
	"encoding/hex"
	"crypto/rand"
	"fmt"
	"../parameters"
	"../hypertree"
)

func TestSha256n256fRobust(t *testing.T) {
	params := SphincsParams(*parameters.MakeSphincsPlusSHA256256fRobust(false))
	
	skprf, _ := hex.DecodeString("47616c696e736f676120737562646973636f6964656120697320612072617265")

	pk := new(SPHINCS_PK)
	pk.PKseed = make([]byte, params.N)
	sk := new(SPHINCS_SK)
	sk.PKseed = make([]byte, params.N)
	sk.SKseed = make([]byte, params.N)
	sk.SKprf = skprf

	htParams := hypertree.HTParams(params)
	root := htParams.Ht_PKgen(sk.SKseed, sk.PKseed)

	pk.PKroot = root
	sk.PKroot = root

	text := "Galinsoga subdiscoidea is a rare"
	bytesToSign := []byte(text)

	signature := params.Spx_sign(bytesToSign, sk)

	if(!params.Spx_verify(bytesToSign, signature, pk)) {
		t.Errorf("Verification failed, but was expected to succeed")
	}

	/* fmt.Println("Signature")
	fmt.Print(hex.EncodeToString(signature.R)) // R is now correct!!!
	for i := 0; i < params.K; i++ {
		fmt.Print(hex.EncodeToString(signature.SIG_FORS.GetSK(i)))
		fmt.Print(hex.EncodeToString(signature.SIG_FORS.GetAUTH(i)))
	}

	for _, xmssSig := range signature.SIG_HT.XMSSSignatures {
		fmt.Print(hex.EncodeToString(xmssSig.GetWOTSSig()))
		fmt.Print(hex.EncodeToString(xmssSig.GetXMSSAUTH()))
	}
	
	fmt.Println("")

	t.Errorf("Verification failed, but was expected to succeed") */
}


func TestSignAndVerify(t *testing.T) {
	params := SphincsParams(*parameters.MakeSphincsPlusSHA256256fRobust(false))
	for i := 0; i < 5; i++ {

		message := make([]byte, params.N)
		rand.Read(message)

		sk, pk := params.Spx_keygen()
		signature := params.Spx_sign(message, sk)

		if(!params.Spx_verify(message, signature, pk)) {
			t.Errorf("Verification failed, but was expected to succeed")
		}
	}
	
}

func BenchmarkSphincsPlus(b *testing.B) {
	cases := []struct {
		Param SphincsParams
		SphincsVariant string
	} {
		{Param: SphincsParams(*parameters.MakeSphincsPlusSHA256256fRobust(false)), SphincsVariant: "SHA256256f-Robust"},
		{Param: SphincsParams(*parameters.MakeSphincsPlusSHA256256sRobust(false)), SphincsVariant: "SHA256256s-Robust"},
		{Param: SphincsParams(*parameters.MakeSphincsPlusSHA256256fSimple(false)), SphincsVariant: "SHA256256f-Simple"},
		{Param: SphincsParams(*parameters.MakeSphincsPlusSHA256256sSimple(false)), SphincsVariant: "SHA256256s-Simple"},

		{Param: SphincsParams(*parameters.MakeSphincsPlusSHA256192fRobust(false)), SphincsVariant: "SHA256192f-Robust"},
		{Param: SphincsParams(*parameters.MakeSphincsPlusSHA256192sRobust(false)), SphincsVariant: "SHA256192s-Robust"},
		{Param: SphincsParams(*parameters.MakeSphincsPlusSHA256192fSimple(false)), SphincsVariant: "SHA256192f-Simple"},
		{Param: SphincsParams(*parameters.MakeSphincsPlusSHA256192sSimple(false)), SphincsVariant: "SHA256192s-Simple"},

		{Param: SphincsParams(*parameters.MakeSphincsPlusSHA256128fRobust(false)), SphincsVariant: "SHA256128f-Robust"},
		{Param: SphincsParams(*parameters.MakeSphincsPlusSHA256128sRobust(false)), SphincsVariant: "SHA256128s-Robust"},
		{Param: SphincsParams(*parameters.MakeSphincsPlusSHA256128fSimple(false)), SphincsVariant: "SHA256128f-Simple"},
		{Param: SphincsParams(*parameters.MakeSphincsPlusSHA256128sSimple(false)), SphincsVariant: "SHA256128s-Simple"},

		{Param: SphincsParams(*parameters.MakeSphincsPlusSHAKE256256fRobust(false)), SphincsVariant: "SHAKE256256f-Robust"},
		{Param: SphincsParams(*parameters.MakeSphincsPlusSHAKE256256sRobust(false)), SphincsVariant: "SHAKE256256s-Robust"},
		{Param: SphincsParams(*parameters.MakeSphincsPlusSHAKE256256fSimple(false)), SphincsVariant: "SHAKE256256f-Simple"},
		{Param: SphincsParams(*parameters.MakeSphincsPlusSHAKE256256sSimple(false)), SphincsVariant: "SHAKE256256s-Simple"},

		{Param: SphincsParams(*parameters.MakeSphincsPlusSHAKE256192fRobust(false)), SphincsVariant: "SHAKE256192f-Robust"},
		{Param: SphincsParams(*parameters.MakeSphincsPlusSHAKE256192sRobust(false)), SphincsVariant: "SHAKE256192s-Robust"},
		{Param: SphincsParams(*parameters.MakeSphincsPlusSHAKE256192fSimple(false)), SphincsVariant: "SHAKE256192f-Simple"},
		{Param: SphincsParams(*parameters.MakeSphincsPlusSHAKE256192sSimple(false)), SphincsVariant: "SHAKE256192s-Simple"},

		{Param: SphincsParams(*parameters.MakeSphincsPlusSHAKE256128fRobust(false)), SphincsVariant: "SHAKE256128f-Robust"},
		{Param: SphincsParams(*parameters.MakeSphincsPlusSHAKE256128sRobust(false)), SphincsVariant: "SHAKE256128s-Robust"},
		{Param: SphincsParams(*parameters.MakeSphincsPlusSHAKE256128fSimple(false)), SphincsVariant: "SHAKE256128f-Simple"},
		{Param: SphincsParams(*parameters.MakeSphincsPlusSHAKE256128sSimple(false)), SphincsVariant: "SHAKE256128s-Simple"},

	}

	for _, paramVal := range cases {
		b.Run(fmt.Sprintf("Keygen %s", paramVal.SphincsVariant), func(b *testing.B) { benchmarkKeygen(b, paramVal.Param) })
		b.Run(fmt.Sprintf("Sign %s", paramVal.SphincsVariant), func(b *testing.B) { benchmarkSign(b, paramVal.Param) })
		b.Run(fmt.Sprintf("Verify %s", paramVal.SphincsVariant), func(b *testing.B) { benchmarkVerify(b, paramVal.Param) })
	}
	
}

func benchmarkKeygen(b *testing.B, params SphincsParams) {
	for i := 0; i < b.N; i++ {	
		params.Spx_keygen()
	}
}

func benchmarkSign(b *testing.B, params SphincsParams) {
	message := make([]byte, 32)
	rand.Read(message)
	sk, _ := params.Spx_keygen()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {	
		params.Spx_sign(message, sk)
	}
	
}

func benchmarkVerify(b *testing.B, params SphincsParams) {
	message := make([]byte, 32)
	rand.Read(message)
	sk, pk := params.Spx_keygen()
	sig := params.Spx_sign(message, sk)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {	
		params.Spx_verify(message, sig, pk)
	}
}

