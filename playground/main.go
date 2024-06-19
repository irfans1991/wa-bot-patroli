package main

import (
	"fmt"
	"time"
)

// func main() {
// 	jsonString := `{"firstname": "John", "lastname": "Doe"}`

// 	fmt.Println(jsonString)

// 	var data map[string]string

// 	err := json.Unmarshal([]byte(jsonString), &data)
// 	if err != nil {
// 		log.Fatalf("Unable to unmarshal json: %v", err)
// 	}

// 	if firstname, ok := data["firstname"]; ok {
// 		fmt.Println("Firstname:", firstname)
// 	} else {
// 		fmt.Println("Firstname key not found in the JSON data.")
// 	}
// }

// func main() {
// 	// Replace this value with your Unix timestamp in milliseconds
// 	// Replace this value with your Unix timestamp in milliseconds
// 	unixMilli := int64(1715526000000)

// 	// Convert Unix timestamp in milliseconds to time.Time
// 	timeObj := time.Unix(unixMilli/1000, (unixMilli%1000)*int64(time.Millisecond))

// 	// Format the time as "2006-Jan-02"
// 	dateString := timeObj.Format("2006-Jan-02")

// 	parts := strings.Split(dateString, "-")

// 	// Print the result
// 	fmt.Println("Unix Milliseconds:", unixMilli)
// 	fmt.Println("Formatted Date:", dateString)
// 	fmt.Println("Date Parts:", parts)

// }

// func main() {
// 	// Create an array of strings in the format "Name#address#age"
// 	dataArray := []string{
// 		"John#123 Main St#25",
// 		"Alice#456 Oak St#30",
// 		"Bob#789 Pine St#22",
// 	}

// 	// Convert the array to a single string
// 	arrayString := strings.Join(dataArray, "#")

// 	fmt.Println(arrayString)

// 	// Split the string back to an array
// 	stringArray := strings.Split(arrayString, "#")

// 	// Print the reconstructed array
// 	fmt.Println(stringArray[1])
// }

// func main() {
// 	// String dengan format "NamaAwal#NamaAkhir#Umur"
// 	inputString := "John#Doe#25"

// 	// Memisahkan string berdasarkan tanda pagar (#)
// 	parts := strings.Split(inputString, "#")

// 	// Menyimpan nilai dari dua kolom pertama ke dalam variabel yang sesuai
// 	if len(parts) >= 2 {
// 		firstName := parts[0]
// 		lastName := parts[1]

// 		// Menampilkan hasil
// 		fmt.Printf("Nama Awal: %s\n", firstName)
// 		fmt.Printf("Nama Akhir: %s\n", lastName)
// 	} else {
// 		fmt.Println("Format string tidak sesuai")
// 	}
// }

// func main() {
// 	// String dengan format "NamaAwal#NamaAkhir#Umur"
// 	inputString := "John#Doe#25"

// 	// Mengecek apakah string mengandung karakter #

// 	pagar := strings.Contains(inputString, "#")
// 	fmt.Println(pagar)
// 	if strings.Contains(inputString, "#") {
// 		// Mendapatkan posisi karakter #
// 		index := strings.Index(inputString, "#")
// 		fmt.Println(index)

// 		// Menampilkan hasil
// 		fmt.Printf("Karakter # ditemukan pada posisi: %d\n", index)
// 	} else {
// 		fmt.Println("Karakter # tidak ditemukan dalam string")
// 	}
// }

func main() {
	waktuSekarang := time.Now()
	selamat := waktuSelamat(waktuSekarang.Hour())

	fmt.Println("Selamat", selamat)
}

func waktuSelamat(jam int) string {
	switch {
	case jam >= 4 && jam < 12:
		return "pagi, semangat untuk hari yang baru!"
	case jam >= 12 && jam < 17:
		return "siang, jangan lupa istirahat sejenak!"
	case jam >= 17 && jam < 20:
		return "sore, tetap semangat sampai pulang!"
	default:
		return "malam, selamat beristirahat dan mimpi indah!"
	}
}
