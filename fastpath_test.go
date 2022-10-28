package main

import "testing"

func Test_FastPathCheckVersion1Solution1(t *testing.T) {
	var pairs = [][]string{
		{"SFO", "EWR"},
	}

	src, dst := calculateDistance1(pairs)

	if src != "SFO" {
		t.Error("src path is incorrect")
		return
	}

	if dst != "EWR" {
		t.Error("src path is incorrect")
		return
	}
}

func Test_FastPathCheckVersion2Solution1(t *testing.T) {
	var pairs = [][]string{
		{"ATL", "EWR"},
		{"SFO", "ATL"},
	}

	src, dst := calculateDistance1(pairs)

	if src != "SFO" {
		t.Error("src path is incorrect")
		return
	}

	if dst != "EWR" {
		t.Error("src path is incorrect")
		return
	}
}

func Test_FastPathCheckVersion3Solution1(t *testing.T) {
	var pairs = [][]string{
		{"IND", "EWR"},
		{"SFO", "ATL"},
		{"GSO", "IND"},
		{"ATL", "GSO"},
	}

	src, dst := calculateDistance1(pairs)

	if src != "SFO" {
		t.Error("src path is incorrect")
		return
	}

	if dst != "EWR" {
		t.Error("src path is incorrect")
		return
	}
}

func Test_FastPathCheckVersion1Solution2(t *testing.T) {
	var pairs = [][]string{
		{"SFO", "EWR"},
	}

	src, dst := calculateDistance2(pairs)

	if src != "SFO" {
		t.Error("src path is incorrect")
		return
	}

	if dst != "EWR" {
		t.Error("src path is incorrect")
		return
	}
}

func Test_FastPathCheckVersion2Solution2(t *testing.T) {
	var pairs = [][]string{
		{"ATL", "EWR"},
		{"SFO", "ATL"},
	}

	src, dst := calculateDistance2(pairs)

	if src != "SFO" {
		t.Error("src path is incorrect")
		return
	}

	if dst != "EWR" {
		t.Error("src path is incorrect")
		return
	}
}

func Test_FastPathCheckVersion3Solution2(t *testing.T) {
	var pairs = [][]string{
		{"IND", "EWR"},
		{"SFO", "ATL"},
		{"GSO", "IND"},
		{"ATL", "GSO"},
	}

	src, dst := calculateDistance2(pairs)

	if src != "SFO" {
		t.Error("src path is incorrect")
		return
	}

	if dst != "EWR" {
		t.Error("src path is incorrect")
		return
	}
}

func BenchmarkFastPathSolution1(b *testing.B) {
	var pairs = [][]string{
		{"IND", "EWR"},
		{"SFO", "ATL"},
		{"GSO", "IND"},
		{"ATL", "GSO"},
	}

	for i := 0; i < b.N; i++ {
		calculateDistance1(pairs)
	}
}

func BenchmarkFastPathSolution2(b *testing.B) {
	var pairs = [][]string{
		{"IND", "EWR"},
		{"SFO", "ATL"},
		{"GSO", "IND"},
		{"ATL", "GSO"},
	}

	for i := 0; i < b.N; i++ {
		calculateDistance2(pairs)
	}
}
