
TASK : 
Pada task akhir VIX Full Stack Developer ini kalian diarahkan untuk membentuk API berdasarkan kasus yang telah diberikan. Pada kasus ini, kalian diinstruksikan untuk membuat API untuk mengupload dan menghapus gambar. API yang kalian bentuk adalah POST, GET, PUT, dan DELETE.

1. Pendahuluan
			a.	Project ini bertujuan untuk meningkatkan engagement user pada m-banking untuk meningkatkan aspek memiliki user pada aplikasi tersebut.
			b.	Project ini bertujuan untuk membentuk personalize user yang dapat memungkinkan user bisa mengupload dan menghapus photo

2.	Fungsi
			•	User : Register, Login
			•	Photo(Protected using middleware) : getPhoto, getPhotoById, uploadPhoto, updatePhoto, deletePhoto 


3.	Depedencies
			•	Gorm : 
			Sebuah ORM (Object-Relational-Mapping) untuk bahasa pemrograman golang. ORM membantu developer untuk berinteraksi dengan database atau suatu package yang dapat menghubungkan code dengan database. Gorm ini 				saya pakai untuk Query, Relation, CRUD, Mapping
			•	JWT : 
			Library golang yang digunakan untuk  mengubah data menjadi token JSON yang aman dan terenskripsi, Memastikan bahwa token JSON valid
			•	Bcrypt : 
			Package yang disediakan golang untuk generating hash dari password dan memverifikasi password terhadap hash yang tersimpan.
			•	Mux :
			Package yang disediakan golang sebagai router untuk HTTP server. Biasanya package ini digunakan untuk memproses request HTTP seperti CRUD dalam API. Mux juga dapat memungkinkan developer untuk 										mendefinisikan handler berbeda untuk menangani request HTTP berdasarkan kebutuhan.



4.	Struktur Dokument 
			•	App : Sebagai Folder Utama aplikasi golang yang berisi Subfolder
			•	Config : Menyimpan secret key untuk JWT dan struktur data untuk payload JWT
			•	Controllers : berisi logic database yaitu models dan query
			•	Helper : berisi fungsi untuk mengirimkan response JSON dalam aplikasi golang yang menangani request HTTP
			•	Middlewares: berisi fungsi untuk authentication JWT 
			•	Models : berisi struct Photo,user dan ConnectDatabase
			•	Go.mod : package library yang diinstal 
			•	Main.go : berisi route https server atau endpoint 

5.	Configuration
		Project ini memerlukan beberapa Configuration untuk dapat dijalankan.
			•	Configuration Database 
					Saya menggunakan database MySQL dengan 
					Host : localhost
					Port:3306
					Username : root
					Password : root123
					Dbase_name : rakamin_project
			•	Configuration Server
					Saya menggunakan port 8080 untuk menjalankan API
			•	Configuration JWT
					Configuration secret key dan juga struktur data untuk payload JWT



 API DOC

https://documenter.getpostman.com/view/13370431/2s9YysBM9i
