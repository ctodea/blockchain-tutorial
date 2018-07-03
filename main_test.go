package main

import (
	"fmt"
	"testing"
)

func TestCalculateHash(t *testing.T) {
	block1 := Block{
		Index:     2,
		Timestamp: "2018-07-02 16:13:33.6772407 +0100 BST m=+221.676074771",
		BPM:       100,
		Hash:      "3a423a8c52007d5b9ff6e3d2557e2502acf6a500f2d421fc0ceb1fe1d4f81408",
		PrevHash:  "1d45e83715c025e41b0778c7e23c9da13b36ab3f3f5933abd18927b25ac0e0bf",
	}

	block2 := Block{
		Index:     3,
		Timestamp: "2018-07-02 16:17:05.781945867 +0100 BST m=+433.780779935",
		BPM:       70,
		Hash:      "77bd3dc259e9aaab751ccf9417a3dbc66ae436f33190d444de389b85b1b49bd0",
		PrevHash:  "3a423a8c52007d5b9ff6e3d2557e2502acf6a500f2d421fc0ceb1fe1d4f81408",
	}

	// var tests = []struct {
	// 	input	string
	// 	output	string
	// }{
	// 	{"foo", "b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c"},
	// 	{"bar", "7d865e959b2466918c9863afca942d0fb89d7c9ac0c99bafc3749504ded97730"}
	// }
}

// func TestIsBlockValid(t *testing.T) {
// 	hashBlock1 :=
// 	block1 := Block{2, time.Now.String()}

// 	genesisBlock := Block{rand.Intn(100), time.Now().String(), rand.Intn(100), , ""}

// }

func main() {
	block1 := Block{3, "2018-07-02 16:00:21.715124691 +0100 BST m=+0.001360449", 100, "", "aec070645fe53ee3b3763059376134f058cc337247c978add178b6ccdfb0019f"}
	fmt.Println(calculateHash())
}
