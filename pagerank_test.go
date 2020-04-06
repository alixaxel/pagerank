package pagerank

import (
	"reflect"
	"testing"
)

func TestEmpty64(t *testing.T) {
	graph := NewGraph64()

	actual := map[uint64]float64{}
	expected := map[uint64]float64{}

	graph.Rank(0.85, 0.000001, func(node uint64, rank float64) {
		actual[node] = rank
	})

	if reflect.DeepEqual(actual, expected) != true {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestSimple64(t *testing.T) {
	graph := NewGraph64()

	graph.Link(1, 2, 1.0)
	graph.Link(1, 3, 1.0)
	graph.Link(2, 3, 1.0)
	graph.Link(2, 4, 1.0)
	graph.Link(3, 1, 1.0)

	actual := map[uint64]float64{}
	expected := map[uint64]float64{
		1: 0.32721836185043207,
		2: 0.2108699481253495,
		3: 0.3004897566512289,
		4: 0.16142193337298952,
	}

	graph.Rank(0.85, 0.000001, func(node uint64, rank float64) {
		actual[node] = rank
	})

	if reflect.DeepEqual(actual, expected) != true {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestWeighted64(t *testing.T) {
	graph := NewGraph64()

	graph.Link(1, 2, 1.0)
	graph.Link(1, 3, 2.0)
	graph.Link(2, 3, 3.0)
	graph.Link(2, 4, 4.0)
	graph.Link(3, 1, 5.0)

	actual := map[uint64]float64{}
	expected := map[uint64]float64{
		1: 0.34983779905464363,
		2: 0.1688733284604475,
		3: 0.3295121849483849,
		4: 0.15177668753652385,
	}

	graph.Rank(0.85, 0.000001, func(node uint64, rank float64) {
		actual[node] = rank
	})

	if reflect.DeepEqual(actual, expected) != true {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestDuplicates64(t *testing.T) {
	graph := NewGraph64()

	graph.Link(1, 2, 1.0)
	graph.Link(1, 3, 2.0)
	graph.Link(2, 3, 3.0)
	graph.Link(2, 4, 4.0)
	graph.Link(3, 1, 5.0)

	graph.Link(1, 2, 6.0)
	graph.Link(1, 3, 7.0)

	actual := map[uint64]float64{}
	expected := map[uint64]float64{
		1: 0.3312334209098247,
		2: 0.19655848316544225,
		3: 0.3033555769882879,
		4: 0.168852518936445,
	}

	graph.Rank(0.85, 0.000001, func(node uint64, rank float64) {
		actual[node] = rank
	})

	if reflect.DeepEqual(actual, expected) != true {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestDuplicatesAfterReset64(t *testing.T) {
	graph := NewGraph64()

	graph.Link(1, 2, 1.0)
	graph.Link(1, 3, 2.0)
	graph.Link(2, 3, 3.0)
	graph.Link(2, 4, 4.0)
	graph.Link(3, 1, 5.0)

	graph.Reset()

	graph.Link(1, 2, 6.0)
	graph.Link(1, 3, 7.0)

	actual := map[uint64]float64{}
	expected := map[uint64]float64{
		1: 0.25974019022001016,
		2: 0.3616383883769191,
		3: 0.3786214214030706,
	}

	graph.Rank(0.85, 0.000001, func(node uint64, rank float64) {
		actual[node] = rank
	})

	if reflect.DeepEqual(actual, expected) != true {
		t.Error("Expected", expected, "but got", actual)
	}
}

func BenchmarkGraph64(b *testing.B) {
	for n := 0; n < b.N; n++ {
		graph := NewGraph64()

		graph.Link(1, 2, 1.0)
		graph.Link(1, 3, 2.0)
		graph.Link(2, 3, 3.0)
		graph.Link(2, 4, 4.0)
		graph.Link(3, 1, 5.0)

		results := map[uint64]float64{}

		graph.Rank(0.85, 0.000001, func(node uint64, rank float64) {
			results[node] = rank
		})
	}
}

func TestEmpty32(t *testing.T) {
	graph := NewGraph32()

	actual := map[uint64]float32{}
	expected := map[uint64]float32{}

	graph.Rank(0.85, 0.000001, func(node uint64, rank float32) {
		actual[node] = rank
	})

	if reflect.DeepEqual(actual, expected) != true {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestSimple32(t *testing.T) {
	graph := NewGraph32()

	graph.Link(1, 2, 1.0)
	graph.Link(1, 3, 1.0)
	graph.Link(2, 3, 1.0)
	graph.Link(2, 4, 1.0)
	graph.Link(3, 1, 1.0)

	actual := map[uint64]float32{}
	expected := map[uint64]float32{
		1: 0.32721835,
		2: 0.21086994,
		3: 0.30048975,
		4: 0.16142192,
	}

	graph.Rank(0.85, 0.000001, func(node uint64, rank float32) {
		actual[node] = rank
	})

	if reflect.DeepEqual(actual, expected) != true {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestWeighted32(t *testing.T) {
	graph := NewGraph32()

	graph.Link(1, 2, 1.0)
	graph.Link(1, 3, 2.0)
	graph.Link(2, 3, 3.0)
	graph.Link(2, 4, 4.0)
	graph.Link(3, 1, 5.0)

	actual := map[uint64]float32{}
	expected := map[uint64]float32{
		1: 0.34983802,
		2: 0.16887322,
		3: 0.32951212,
		4: 0.15177679,
	}

	graph.Rank(0.85, 0.000001, func(node uint64, rank float32) {
		actual[node] = rank
	})

	if reflect.DeepEqual(actual, expected) != true {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestDuplicates32(t *testing.T) {
	graph := NewGraph32()

	graph.Link(1, 2, 1.0)
	graph.Link(1, 3, 2.0)
	graph.Link(2, 3, 3.0)
	graph.Link(2, 4, 4.0)
	graph.Link(3, 1, 5.0)

	graph.Link(1, 2, 6.0)
	graph.Link(1, 3, 7.0)

	actual := map[uint64]float32{}
	expected := map[uint64]float32{
		1: 0.33123338,
		2: 0.19655848,
		3: 0.30335557,
		4: 0.16885251,
	}

	graph.Rank(0.85, 0.000001, func(node uint64, rank float32) {
		actual[node] = rank
	})

	if reflect.DeepEqual(actual, expected) != true {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestDuplicatesAfterReset32(t *testing.T) {
	graph := NewGraph32()

	graph.Link(1, 2, 1.0)
	graph.Link(1, 3, 2.0)
	graph.Link(2, 3, 3.0)
	graph.Link(2, 4, 4.0)
	graph.Link(3, 1, 5.0)

	graph.Reset()

	graph.Link(1, 2, 6.0)
	graph.Link(1, 3, 7.0)

	actual := map[uint64]float32{}
	expected := map[uint64]float32{
		1: 0.25974017,
		2: 0.36163837,
		3: 0.3786214,
	}

	graph.Rank(0.85, 0.000001, func(node uint64, rank float32) {
		actual[node] = rank
	})

	if reflect.DeepEqual(actual, expected) != true {
		t.Error("Expected", expected, "but got", actual)
	}
}

func BenchmarkGraph32(b *testing.B) {
	for n := 0; n < b.N; n++ {
		graph := NewGraph64()

		graph.Link(1, 2, 1.0)
		graph.Link(1, 3, 2.0)
		graph.Link(2, 3, 3.0)
		graph.Link(2, 4, 4.0)
		graph.Link(3, 1, 5.0)

		results := map[uint64]float64{}

		graph.Rank(0.85, 0.000001, func(node uint64, rank float64) {
			results[node] = rank
		})
	}
}
