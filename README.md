# learngo

## Another way using condition


Untuk seleksi kondisi dengan jumlah kondisi lebih dari satu, bisa gunakan else if.

```bash
{{if pipeline}}
    a
{{else if pipeline}}
    b
{{else}}
    c
{{end}}
```
Untuk seleksi kondisi yang kondisinya adalah bersumber dari variabel bertipe bool, maka langsung saja tulis tanpa menggunakan eq. Jika kondisi yang diinginkan adalah kebalikan dari nilai variabel, maka gunakan ne. Contohnya bisa dilihat pada kode berikut.

```bash
{{if .IsTrue}}
    <p>true</p>
{{end}}

{{isTrue := true}}

{{if isTrue}}
    <p>true</p>
{{end}}

{{if eq isTrue}}
    <p>true</p>
{{end}}

{{if ne isTrue}}
    <p>not true (false)</p>
{{end}}
```