package hangman

import "testing"

func validState(t *testing.T, expectedState, actualState GameStateType) bool {
	if expectedState != actualState {
		t.Errorf("state should be '%v', got=%v", expectedState, actualState)
		return false
	}
	return true
}
