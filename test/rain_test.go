package rain

import "testing"

func TestSpeak(t *testing.T){
	result := SpeakRain();
	if result == 1 {
		t.Logf("Success good")
	}else{
		t.Errorf("Oh sorry failed")
	}
}

func TestMath(t *testing.T) {
    t.Run("Addition", func(t *testing.T) {
        if 2+2 != 4 {
            t.Errorf("expected 4, got something else")
        }
    })

    t.Run("Multiplication", func(t *testing.T) {
        if 3*3 != 9 {
            t.Errorf("expected 9, got something else")
        }
    })
}