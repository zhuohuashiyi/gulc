package gulc

import (
	"testing"
)


func TestSubstrImplHash(t *testing.T) {
	t.Log(SubstrImplHash("bbcbbcbbbcbbcb", "bbcbbcb"))
	t.Log(SubstrImplHash("cabcacabbacabccbabbbacabcacabbacabccbabbbcbcabcacabbacabccbabbbcabcacabbacabccbabbbaacabcaca", "cabcacabbacabccbabbb"))	
}


