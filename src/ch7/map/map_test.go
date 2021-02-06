package _map

import "testing"

func TestInitMap(t *testing.T) {
	m := make(map[int]string, 10)
	t.Log(m, len(m))
}

func TestAccessMap(t *testing.T) {
	m := map[int]int{}
	if ele, ok := m[0]; ok {
		t.Logf("exists: %v", ele)
	} else {
		t.Log("not exists")
	}
}
