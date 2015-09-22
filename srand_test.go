package srand

import "testing"

func TestGenerateRandomString(t *testing.T) {
	expectedLen := 44

	token, err := GenerateRandomString(32)
	if err != nil {
		t.Errorf("GenerateRandomString, expected token, got %v", err)
	}
	if len(token) != expectedLen {
		t.Errorf("GenerateRandomString, expected token length %v, got %v", expectedLen, len(token))
	}
}

func TestGenerateRandomInt(t *testing.T) {
	rand, err := GenerateRandomInt(100)
	if err != nil {
		t.Errorf("GenerateRandomInt, expected random int, got %v", err)
	}
	if rand <= 0 || rand > 100 {
		t.Errorf("GenerateRandomInt, expected to be in rangein [0,100), got %v", rand)
	}
}

func TestGenerateRandomIntRange(t *testing.T) {
	var min, max int64 = 42, 1337
	rand, err := GenerateRandomIntRange(42, 1337)
	if err != nil {
		t.Errorf("GenerateRandomIntRange, expected random int, got %v", err)
	}
	if rand <= min || rand > max {
		t.Errorf("GenerateRandomIntRange, expected to be in range [%v,%v), got %v", min, max, rand)
	}
	// Check min > max
	_, err = GenerateRandomIntRange(max, min)
	if err != MinMaxError {
		t.Errorf("GenerateRandomIntRange, expected error when min > max, got %v", err)
	}
}

func TestGenerateRandomFloat(t *testing.T) {
	rand, err := GenerateRandomFloat()
	if err != nil {
		t.Errorf("GenerateRandomFloat, expected random int, got %v", err)
	}
	if rand <= 0 || rand > 1.0 {
		t.Errorf("GenerateRandomFloat, expected to be in rangein [0,1), got %v", rand)
	}
}

func TestGenerateRandomFloatRange(t *testing.T) {
	min := 12.2
	max := 16.6
	rand, err := GenerateRandomFloatRange(min, max)
	if err != nil {
		t.Errorf("GenerateRandomInt, expected random int, got %v", err)
	}
	if rand <= min || rand > max {
		t.Errorf("GenerateRandomFloatInRange, expected to be in range [%v, %v), got %v", min, max, rand)
	}
	// Check min > max
	_, err = GenerateRandomFloatRange(max, min)
	if err != MinMaxError {
		t.Errorf("GenerateRandomFloatRange, expected error when min > max, got %v", err)
	}
}

func TestGenerateUUID(t *testing.T) {
	uuid, err := GenerateUUID()
	if err != nil {
		t.Errorf("GenerateUUID, expected uuid, got %v", err)
	}
	if len(uuid) != 36 {
		t.Errorf("GenerateUUID, expected length %s, got %v", 36, len(uuid))
	}
}

//
// Benchmark Tests
//

func BenchmarkGenerateRandomInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateRandomInt(100)
	}
}

func BenchmarkGenerateRandomIntRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateRandomIntRange(42, 1337)
	}
}

func BenchmarkGenerateRandomFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateRandomFloat()
	}
}

func BenchmarkGenerateRandomFloatRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateRandomFloatRange(4.2, 13.37)
	}
}
