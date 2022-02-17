package hasher

import "testing"

const ITERATIONS = 30

func Test_NewToken(t *testing.T) {
	t.Parallel()
	tokens := make([]Token, ITERATIONS)
	for i := 0; i < ITERATIONS; i++ {
		tokens[i] = GenerateToken()
	}

	// Check if they are unique
	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			if tokens[i].Raw == tokens[j].Raw {
				t.Errorf("NewToken() failed to generate unique tokens")
			}
		}
	}

}

func Test_HashToken_CheckTokenHash(t *testing.T) {
	t.Parallel()
	for i := 0; i < ITERATIONS; i++ {
		// TODO: Implement

	}
}
