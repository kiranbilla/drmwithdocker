/*
	Opendrm, an open source implementation of industry-grade DRM
	(Digital Rights Management) or Key System.
	Copyright (C) 2018  wilkk

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

/*
	This file can be used to generate encryption&decryption key. You can use the
	GenKeyBySeed method to spawn&respawn the same key without storing the relationship
	between KID and Key in db.
*/

package key

import (
	"crypto/sha256"
	"log"
	"math/rand"
	"os/exec"
	"strings"
	"time"
)

// At least 30 bytes
var defaultKeySeed = []byte("b1cc1aa664122baca692107d4ba5d6d21ef9787ee82f8020ec93adcc25d44b8f")

type KeyGenerator struct {
	// Key seed used for GenKeyBySeed.
	// Set to nil if not used.
	seed []byte
}

func NewKeyGenerator(seed []byte) *KeyGenerator {
	return &KeyGenerator{
		seed: seed,
	}
}

func (this *KeyGenerator) GenKeyBySeed(kid string) []byte {
	return generateKeyAndKidBySeed(kid, this.seed)
}

func (this *KeyGenerator) GenKeyByDefaultSeed(uuid string) []byte {
	return generateKeyAndKidBySeed(uuid, defaultKeySeed)
}

func (this *KeyGenerator) GenRandKey() ([]byte, string) {
	return generateRandKeyAndKid()
}

// Based on https://docs.microsoft.com/en-us/playready/specifications/playready-key-seed
// Ck(KID) = f(KID, KeySeed)
func generateKeyAndKidBySeed(uuid string, seed []byte) []byte {
	drm_aes_keysize_128 := 16
	contentKey := make([]byte, drm_aes_keysize_128)

	// Truncate seed to 30 bytes
	truncKeySeed := seed[:30]
	// Get the keyId bytes.
	keyIdBytes := []byte(uuid)

	// Calculate the SHA of truncKeySeed and keyIdBytes.
	shaA := sha256.New()
	shaA.Write(truncKeySeed)
	shaA.Write(keyIdBytes)
	outputA := shaA.Sum(nil)

	// Calculate the SHA of truncKeySeed, keyIdBytes, and truncKeySeed.
	shaB := sha256.New()
	shaB.Write(truncKeySeed)
	shaB.Write(keyIdBytes)
	shaB.Write(truncKeySeed)
	outputB := shaB.Sum(nil)

	// Calculate the SHA of truncKeySeed, keyIdBytes, truncKeySeed again, and keyIdBytes again.
	shaC := sha256.New()
	shaC.Write(truncKeySeed)
	shaC.Write(keyIdBytes)
	shaC.Write(truncKeySeed)
	shaC.Write(keyIdBytes)
	outputC := shaB.Sum(nil)

	for i := 0; i < drm_aes_keysize_128; i++ {
		contentKey[i] = outputA[i] ^ outputA[i+drm_aes_keysize_128] ^
			outputB[i] ^ outputB[i+drm_aes_keysize_128] ^
			outputC[i] ^ outputC[i+drm_aes_keysize_128]
	}

	return contentKey
}

func GenerateUUID() string {
	kidBytes, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
		return ""
	} else {
		return strings.TrimSuffix(string(kidBytes), "\n")
	}
}

// If error happens, kid is empty string.
func generateRandKeyAndKid() ([]byte, string) {
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	key := make([]byte, 16)
	rnd.Read(key)

	kid := GenerateUUID()

	return key, kid
}
