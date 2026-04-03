package seeder

import (
	"bcc-geazy/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedInformasi(db *gorm.DB) {
	info := entity.Informasi{
		Id:    uuid.New(),
		Judul: "5 Cara Mudah Membiasakan Gaya Hidup Sehat pada Anak",
		Ringkasan: `Membangun kebiasaan sehat pada anak memang butuh kesabaran ekstra, apalagi di tengah kesibukan Ayah dan Bunda sehari-hari. Namun, kebiasaan yang ditanamkan sejak dini akan menjadi fondasi kuat bagi tumbuh kembang mereka di masa depan.

Yuk, terapkan 5 langkah sederhana ini untuk memulai gaya hidup sehat di rumah!

1. Jadilah Role Model yang Baik
Anak adalah peniru yang ulung. Sebelum menyuruh si kecil makan sayur atau rajin berolahraga, pastikan Ayah dan Bunda sudah mempraktikkannya lebih dulu. Menunjukkan kebiasaan sehat secara langsung jauh lebih efektif daripada sekadar menasihati.

2. Sediakan Menu Sehat di Meja Makan
Tahukah Bunda? Selera makan anak saat dewasa sangat dipengaruhi oleh apa yang sering mereka makan saat kecil. Biasakan menyajikan makanan bergizi seimbang di rumah. Jika anak terlalu sering disuguhi makanan cepat saji (fast food), mereka akan terus mencari makanan tersebut hingga besar nanti.

3. Sempatkan Aktif Bergerak Bersama
Tubuh anak dirancang untuk terus aktif! Mereka membutuhkan setidaknya satu jam aktivitas fisik setiap harinya. Cobalah luangkan waktu di sore hari atau akhir pekan untuk bersepeda, jalan santai, atau sekadar bermain bola bersama di halaman. Selain sehat, ini juga mempererat bonding keluarga.

4. Hindari Menjadikan Junk Food Sebagai Hadiah
Sering menjanjikan es krim atau cokelat saat anak mendapat nilai bagus? Sebaiknya kebiasaan ini mulai dikurangi, ya. Menggunakan makanan manis atau junk food sebagai reward akan membuat anak berpikir bahwa makanan tidak sehat adalah sesuatu yang "spesial". Ganti hadiahnya dengan kegiatan seru, seperti jalan-jalan ke taman atau membacakan buku dongeng baru.

5. Ciptakan Zona Bebas Gadget di Kamar Tidur
Tidur yang cukup dan berkualitas sangat penting untuk menjaga berat badan ideal dan daya ingat anak di sekolah. Agar waktu istirahatnya tidak terganggu, pastikan kamar tidur bebas dari TV, video game, maupun smartphone. Biarkan kamar menjadi tempat yang tenang khusus untuk beristirahat.`,
	}

	var existing entity.Informasi
	err := db.Where("judul = ?", info.Judul).First(&existing).Error

	if err == gorm.ErrRecordNotFound {
		db.Create(&info)
	}
}
