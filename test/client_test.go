package test

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"io"
	"testing"
	"time"
)

type ProposalType uint8

const (
	ProposalType_ADD     ProposalType = 1
	ProposalType_KICKOUT ProposalType = 2
	ProposalType_QUIT    ProposalType = 3
)

func Test_ProposalType(t *testing.T) {
	printType(uint8(ProposalType_ADD))
}

func printType(ptype uint8) {
	fmt.Println(ptype)
}

type Person struct {
	Name    string
	Age     int
	Context string
	Backup  string
	Score   map[string]int
}

func Test_json(t *testing.T) {
	scoreMap := make(map[string]int)
	scoreMap["science"] = 19
	scoreMap["math"] = 20
	scoreMap["art"] = 40
	p := Person{"ross", 11, "ctx", "back", scoreMap}
	p1Json, _ := json.Marshal(p)
	fmt.Println(string(p1Json))

	now := time.Now().UTC()

	nowUtcString := now.Format("2006-01-02T15:04:05.000")
	fmt.Println("nowUtcString:" + nowUtcString)

	time1, err := time.Parse("2006-01-02T15:04:05.000", nowUtcString)
	if err != nil {
		t.Fatalf("error: %v", err)
	} else {
		fmt.Println(time1)
	}
	time2, err := time.ParseInLocation("2006-01-02T15:04:05.000", nowUtcString, time.UTC)
	if err != nil {
		t.Fatalf("error: %v", err)
	} else {
		fmt.Println(time2)
	}

	// --
	s1 := "[{\"Name\":\"ross\",\"Age\":11},{\"Name\":\"Joey\",\"Age\":21}]"
	p1 := make([]Person, 0)
	err = json.Unmarshal([]byte(s1), &p1)
	if err != nil {
		t.Fatalf("error: %v", err)
	} else {
		t.Log(p1)
	}

	s2 := "[{\"Name\":\"ross\",\"Age\":11}]"
	p2 := make([]Person, 0)
	err = json.Unmarshal([]byte(s2), &p2)
	if err != nil {
		t.Fatalf("error: %v", err)
	} else {
		t.Log(p2)
	}

}

func Test_json_int(t *testing.T) {
	v := int(15)
	data, _ := json.Marshal(v)
	var vv int
	json.Unmarshal(data, &vv)
	fmt.Println(vv)
}

func Test_json_string(t *testing.T) {
	v := "hello"
	data, _ := json.Marshal(v)
	var vv string
	json.Unmarshal(data, &vv)
	fmt.Println(vv)
}

func Test_json_sort(t *testing.T) {
	scoreMap := make(map[string]int)
	scoreMap["science"] = 19
	scoreMap["math"] = 20
	scoreMap["art"] = 40
	p := Person{"ross", 11, "ctx", "back", scoreMap}

	data, _ := json.Marshal(p)
	m := make(map[string]interface{})
	json.Unmarshal(data, &m)

	dataMap, _ := json.Marshal(m)
	fmt.Println(string(dataMap))
}

func Test_sha3(t *testing.T) {
	field := "test"
	salt := "salt"
	w := sha3.New256()
	io.WriteString(w, field+salt)

	hashstr := hex.EncodeToString(w.Sum(nil))

	fmt.Println(hashstr)
}

func Test_sign(t *testing.T) {
	text := "this is a test"
	hash := getSHA3(text)
	key, _ := crypto.GenerateKey()
	pubKey := key.PublicKey
	signed, err := key.Sign(rand.Reader, []byte(hash), nil)
	if err != nil {
		t.Fatalf("error1: %v", err)
	} else {
		t.Log(hex.EncodeToString(signed))
	}
	isV := ecdsa.VerifyASN1(&pubKey, []byte(hash), signed)
	t.Logf("ecdsa.VerifyASN1:%t", isV)

	r, s, err := ecdsa.Sign(rand.Reader, key, []byte(hash))
	if err != nil {
		t.Fatalf("error2: %v", err)
	}

	rt, _ := r.MarshalText()
	st, _ := s.MarshalText()

	t.Logf("rt.len:%d", len(rt))
	t.Logf("st.len:%d", len(st))

	signStr := string(rt) + string(st)
	signature := hex.EncodeToString([]byte(signStr))
	t.Log(signature)

	t.Log(hex.EncodeToString(append(rt, st...)))

	isf := ecdsa.Verify(&pubKey, []byte(hash), r, s)
	t.Logf("ecdsa.Verify:%t", isf)

	hashKeccak256 := crypto.Keccak256([]byte(hash))
	sig, _ := crypto.Sign(hashKeccak256, key)
	t.Logf("sig:%s", hex.EncodeToString(sig))

	b := crypto.VerifySignature(crypto.FromECDSAPub(&pubKey), hashKeccak256, sig[:len(sig)-1])
	t.Logf("crypto.VerifySignature:%t", b)
}

func getSHA3(data string) string {
	w := sha3.New256()
	io.WriteString(w, data)
	return hex.EncodeToString(w.Sum(nil))
}
