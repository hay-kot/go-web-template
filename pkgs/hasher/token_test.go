package hasher

import "testing"

const ITERATIONS = 30

func Test_NewToken(t *testing.T) {
	t.Parallel()
	tokens := make([]string, ITERATIONS)
	for i := 0; i < ITERATIONS; i++ {
		tokens[i] = NewToken()
	}

	// Check if they are unique
	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			if tokens[i] == tokens[j] {
				t.Errorf("NewToken() failed to generate unique tokens")
			}
		}
	}

}

func Test_HashToken_CheckTokenHash(t *testing.T) {
	t.Parallel()
	for i := 0; i < ITERATIONS; i++ {
		token := NewToken()
		hash, err := HashToken(token)
		if err != nil {
			t.Errorf("HashToken() failed to hash token=%v", token)
		}
		if !CheckTokenHash(token, hash) {
			t.Errorf("CheckTokenHash() failed to validate token=%v against hash=%v", token, hash)
		}
		if token == hash {
			t.Errorf("HashToken() failed to generate different token and hash")
		}
	}
}
