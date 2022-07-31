# Gohistogram - Histograms in Go

![build status](https://circleci.com/gh/VividCortex/gohistogram.png?circle-token=d37ec652ea117165cd1b342400a801438f575209)

#### This package provides [Streaming Approximate Histograms](https://hackillyournet.id/blog/2022/07/21/Gohistogram/index.html) for efficient quantile approximations.

#### Histogram dalam paket ini didasarkan pada algoritma yang ditemukan di Ben-Haim & Yom-Tov's *A Streaming Parallel Decision Tree Algorithm* ([PDF](http://jmlr.org/papers/volume11/ben-haim10a/ben-haim10a.pdf)).

#### Bins histogram tidak memiliki ukuran preset. Sebagai nilai-nilai mengalir ke histogram, tempat sampah ditambahkan secara dinamis dan digabung.

#### Implementasi lain dapat ditemukan di proyek Apache Hive lihat ([NumericHistogram](http://hive.apache.org/docs/r0.11.0/api/org/apache/hadoop/hive/ql/udf/generic/NumericHistogram.html)).

#### An example:

![histogram](http://i.imgur.com/5OplaRs.png)

#### Metode yang akurat untuk menghitung kuantil (seperti persentil) memerlukan data yang akan diurutkan. Streaming histogram memungkinkan untuk memperkirakan kuantoba tanpa menyortir nilai (atau bahkan secara individual menyimpan) nilai.

#### Numerichistogram adalah implementasi yang lebih mendasar dari histogram streaming. WeightedHistogram mengimplementasikan nilai-nilai BIN sebagai rata-rata bergerak yang tertimbang secara eksponensial. Ukuran sampah maksimum dilewatkan sebagai argumen untuk metode konstruktor. Ukuran tempat sampah yang lebih besar menghasilkan perkiraan yang lebih akurat dengan biaya peningkatan pemanfaatan dan kinerja memori.

#### A picture of kittens:

![stack of kittens](http://i.imgur.com/QxRTWAE.jpg)

## Getting started
### Menggunakan dalam kode Anda sendiri

    $ go get github.com/Xcod3bughunt3r/Gohistogram
    
```go
import "github.com/Xcod3bughunt3r/Gohistogram"
```

### Menjalankan tes dan membuat modifikasi

#### Dapatkan kode ke ruang kerja Anda:

    $ cd $GOPATH
    $ git clone git@github.com:Xcod3bughunt3r/Gohistogram.git ./src/github.com/Xcod3bughunt3r/Gohistogram

#### Anda dapat menjalankan tes sekarang:

    $ cd src/github.com/Xcod3bughunt3r/Gohistogram
    $ go test .

## API Documentation
#### Dokumentasi sumber penuh dapat ditemukan di sini [godoc](http://godoc.org/github.com/Xcod3bughunt3r/Gohistogram).

## [Contributing Xcod3bughunt3r](https://github.com/Xcod3bughunt3r/Gohistogram/blob/master/Xcod3bughunt3r.md)

#### Kami hanya menerima permintaan tarik untuk perbaikan atau perbaikan kecil. Ini termasuk:

* Small bug fixes
* Typos
* Documentation or comments

#### Silakan buka masalah untuk membahas fitur-fitur baru. Permintaan tarik untuk fitur baru akan ditolak, jadi kami merekomendasikan Forking repositori dan membuat perubahan dalam garpu Anda untuk kasus penggunaan Anda.

## License
#### Copyright (c) 2022-07-21 by Xcod3bughunt3r.
#### Dirilis di bawah lisensi MIT. Periksa file [LICENSE](https://github.com/Xcod3bughunt3r/Gohistogram/blob/master/LICENSE) untuk detailnya.
