package gohistogram

// Copyright (c) 2022-07-21 by Xcod3bughunt3r.
// Dirilis di bawah lisensi MIT. Periksa file `LICENSE` untuk detailnya.

// Histogram is the interface that wraps the Add and Quantile methods.
type Histogram interface {
	// Add adds a new value, n, to the histogram. Trimming is done
	// automatically.
	Add(n float64)

	// Quantile returns an approximation.
	Quantile(n float64) (q float64)

	// String returns a string reprentation of the histogram,
	// which is useful for printing to a terminal.
	String() (str string)
}

type bin struct {
	value float64
	count float64
}
