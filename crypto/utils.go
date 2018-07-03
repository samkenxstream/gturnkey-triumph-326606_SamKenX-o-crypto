// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package crypto

import (
	"encoding/hex"
	"fmt"
	"github.com/ArkEcosystem/go-crypto/crypto/base58"
	"log"
	"strings"
)

func Byte2Hex(data byte) string {
	return fmt.Sprintf("%x", data)
}

func Hex2Byte(data []byte) string {
	return strings.ToLower(fmt.Sprintf("%X", data))
}

func HexEncode(data []byte) string {
	return hex.EncodeToString(data)
}

func HexDecode(data string) []byte {
	result, err := hex.DecodeString(data)

	if err != nil {
		log.Fatal(err.Error())
	}

	return result
}

func Base58Encode(data []byte) string {
	return base58.Encode(data)
}

func Base58Decode(data string) []byte {
	result, err := base58.Decode(data)

	if err != nil {
		log.Fatal(err.Error())
	}

	return result
}
