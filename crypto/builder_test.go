// This file is part of Ark Go Crypto.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildTransferWithPassphrase(t *testing.T) {
	transaction := BuildTransfer(
		&Transaction{
			Amount: FlexToshi(133380000000),
			Nonce: 5,
			RecipientId: "AXoXnFi4z1Z6aFvjEYkDVCtBGW2PaRiM25",
			VendorField: "This is a transaction from Go",
		},
		"This is a top secret passphrase",
		"",
	)

	assert := assert.New(t)

	assert.True(transaction.Verify())
}

func TestBuildTransferWithSecondPassphrase(t *testing.T) {
	secondPassPhrase := "This is a top secret second passphrase"

	transaction := BuildTransfer(
		&Transaction{
			Amount: FlexToshi(133380000000),
			Nonce: 5,
			RecipientId: "AXoXnFi4z1Z6aFvjEYkDVCtBGW2PaRiM25",
			VendorField: "This is a transaction from Go",
		},
		"This is a top secret passphrase",
		secondPassPhrase,
	)

	secondPublicKey, _ := PublicKeyFromPassphrase(secondPassPhrase)

	assert := assert.New(t)

	assert.True(transaction.Verify())
	assert.True(transaction.SecondVerify(secondPublicKey))
}

func TestBuildSecondSignatureRegistration(t *testing.T) {
	transaction := BuildSecondSignatureRegistration(
		&Transaction{
			Nonce: 5,
		},
		"This is a top secret passphrase",
		"This is a top secret second passphrase",
	)

	assert := assert.New(t)

	assert.True(transaction.Verify())
}

func TestBuildDelegateRegistrationWithPassphrase(t *testing.T) {
	transaction := BuildDelegateRegistration(
		&Transaction{
			Asset: &TransactionAsset{
				Delegate: &DelegateAsset{
					Username: "polopolo",
				},
			},
			Nonce: 5,
		},
		"lumber desk thought industry island man slow vendor pact fragile enact season",
		"",
	)

	assert := assert.New(t)

	assert.True(transaction.Verify())
}

func TestBuildDelegateRegistrationWithSecondPassphrase(t *testing.T) {
	secondPassPhrase := "This is a top secret second passphrase"

	transaction := BuildDelegateRegistration(
		&Transaction{
			Asset: &TransactionAsset{
				Delegate: &DelegateAsset{
					Username: "polopolo",
				},
			},
			Nonce: 5,
		},
		"This is a top secret passphrase",
		secondPassPhrase,
	)

	assert := assert.New(t)

	assert.True(transaction.Verify())

	secondPublicKey, _ := PublicKeyFromPassphrase(secondPassPhrase)
	assert.True(transaction.SecondVerify(secondPublicKey))
}

func TestBuildVoteWithPassphrase(t *testing.T) {
	transaction := BuildVote(
		&Transaction{
			Asset: &TransactionAsset{
				Votes: []string{ "+034151a3ec46b5670a682b0a63394f863587d1bc97483b1b6c70eb58e7f0aed192" },
			},
			Nonce: 5,
		},
		"This is a top secret passphrase",
		"",
	)

	assert := assert.New(t)

	assert.True(transaction.Verify())
}

func TestBuildVoteWithSecondPassphrase(t *testing.T) {
	secondPassPhrase := "This is a top secret second passphrase"

	transaction := BuildVote(
		&Transaction{
			Asset: &TransactionAsset{
				Votes: []string{ "+034151a3ec46b5670a682b0a63394f863587d1bc97483b1b6c70eb58e7f0aed192" },
			},
			Nonce: 5,
		},
		"This is a top secret passphrase",
		secondPassPhrase,
	)

	assert := assert.New(t)

	assert.True(transaction.Verify())

	secondPublicKey, _ := PublicKeyFromPassphrase(secondPassPhrase)

	assert.True(transaction.SecondVerify(secondPublicKey))
}

func TestBuildMultiSignatureRegistrationWithPassphrase(t *testing.T) {
	transaction := BuildMultiSignatureRegistration(
		&Transaction{
			Asset: &TransactionAsset{
				MultiSignature: &MultiSignatureRegistrationAsset{
					Min: 2,
					PublicKeys: []string{
						"03a02b9d5fdd1307c2ee4652ba54d492d1fd11a7d1bb3f3a44c4a05e79f19de933",
						"03b02b9d5fdd1307c2ee4652ba54d492d1fd11a7d1bb3f3a44c4a05e79f19de933",
						"03c02b9d5fdd1307c2ee4652ba54d492d1fd11a7d1bb3f3a44c4a05e79f19de933",
					},
				},
			},
			Nonce: 5,
		},
		"This is a top secret passphrase",
		"",
	)

	assert := assert.New(t)

	assert.True(transaction.Verify())
}
