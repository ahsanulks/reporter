# reporter

## masalah
terdapat data nasabah dimana kita harus melakukan proses dalam 1 waktu dan multi thread
1. menghitung average balanced dari setiap nasabah
2. memberikan benefit sesuai kriteria yang telah ditetapkan
3. menambahkan promo sesuai dengan kriteria yang telah ditetapkan
dalam 3 proses tersebut terdapat value pada field balanced yang krusial, karena value ini berfungsi sebagai
1. value reference untuk mendapatkan sebuah benefit
2. value yang berubah karena mendapatkan benefit ataupun promo
3. value yang digunakan untuk menghitung average balanced
dan kita melakukan pemrosesan data tersebut secara concurrent. sehingga jika kita melakukan perubahan field balance secara tidak berurutan sesuai proses yang ada maka value dari average balance serta benefit yang diberikan bisa berbeda dari apa yang kita harapkan

## solusi
1. Melakukan proses average balanced, menambahkan benefit, serta menambahkan promo secara sequential. karena proses ini tidak boleh terbalik
2. Melakukan perhitungan pada setiap prosesnya secara concurrent
3. Melindungi value field balanced menggunakan mutual exclusion object (mutex) agar tidak terjadi race condition saat kita mengakses atau merubah value tersebut
