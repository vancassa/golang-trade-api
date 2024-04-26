Pada final project ini, kalian akan diminta untuk membuat suatu aplikasi bernama BasicTrade, yang dimana pada aplikasi ini admin (seller) dapat membuat suatu produk dengan menyimpan foto dari produk dan membuat variant yang ada pada produk tersebut. Aplikasi ini akan dilengkapi dengan proses autentikasi (login & register) serta CRUD dengan table dan alur yang dijelaskan berikut ini:

Project ini bebas dikerjakan dengan library apapun. Namun agar proses pengerjaannya lebihcepat dan mudah, disarankan untuk menggunakan framework Gin Gonic dan orm Gorm.
Berikut merupakan library/package yang wajib digunakan:

- github.com/golang-jwt/jwt/v5
- golang.org/x/crypto
- github.com/cloudinary/cloudinary-go/v2
  Dalam project ini akan memerlukan 3 table, yaitu admins, products, dan variants.
  Adanya Validasi, dan Authorization untuk dijadikan middleware.
  Semua table tersebut harus mempunyai validasi-validasi pada tiap field-field nya. Validasi boleh dibuat sendiri ataupun menggunakan package seperti Go Validator.
  Proses upload foto menggunakan Cloudinary.
  Untuk endpoint-endpoint yang berguna untuk memodifikasi data kepemilikan seperti Update atau Delete maka harus melalui proses authorization. Perlu diingat disini bahwa proses autorisasi dilakukan setelah proses autentikasi, bukan sebaliknya.
  Seluruh routing endpoint harus diikuti dengan betul.
  Untuk endpoint dengan method GET, tidak perlu menggunakan authentication (Get All & Get Data By ID)
  Final project wajib di deploy pada Railways.
  List scoring pada Final Project

Score maximal 90 point
Mengerjakan semua rules yang sudah diberikan.
Tidak perlu implementasi pagination pada API get all data (Get All Product dan Get All Variant)
Implementasi search by name pada Get All Product
Implementasi search by variant_name pada Get All Variant
Cukup menggunakan id saja, tidak menggunakan uuid di setiap table dan pada url endpoint untuk action Update, Delete, dan Get By ID
Score maximal 100 point
Mengerjakan semua rules yang sudah diberikan.
Implementasi pagination pada API get all data (Get All Product dan Get All Variant)
Implementasi search by name pada Get All Product
Implementasi search by variant_name pada Get All Variant
Menggunakan uuid di setiap table dan pada url endpoint untuk action Update, Delete, dan Get By ID
Yang harus dikumpulkan:
Source code dalam github
Sertakan url domain ketika berhasil deploy menggunakan Railways
Bisa sertakan postman collection pada github. Format collection: Base Trade API - Nama
