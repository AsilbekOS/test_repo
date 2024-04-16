// 1) 12 chi uy vazifada ValidateUser funksiya yaratilgan edi.
// 2) Bugungi test_repo repository yaratilgan ichida o'zgarishlar kiritilishi kerak
// 3) ValidateUser dasturri test_repo ni ichiga joylang
// 4) user package yaratib uning ichida Validate digan funksiyani e'lon qiling
// 6) main package da user.Validate qilingan holda chaqirinsin
// 7) dastur ishlashini hosil qiling (go run yoki go build o'rqali)
// 8) barchasi ishlagan taqdirda, hamma o'zgarishlarni git add, git commit va git push commandar o'rqali bajaring
// 9) natijada local codelaringiz github ga joylangan bolishi kerak
// 10) Repository linkini jonating

package user

import (
	"fmt"
	"regexp"
)

type User struct {
	Name  string
	Age   int
	Email string
}

func ValidateUser(name string, age int, email string) []error {
	var errors []error

	if len(name) < 6 {
		errors = append(errors, fmt.Errorf("name: uzunligi 6 ta belgidan kam bo`lmasligi kerak"))
	}

	if age <= 0 || age >= 120 {
		errors = append(errors, fmt.Errorf("age: 0 dan 100 gacha bo'lishi kerak"))
	}

	if email == "" {
		errors = append(errors, fmt.Errorf("email: bo'sh bo'lishi mumkin emas"))
	} else {
		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
		if !emailRegex.MatchString(email) {
			errors = append(errors, fmt.Errorf("email: to'g'ri formatda bo'lishi kerak (masalan, example@domain.com)"))
		}
	}

	return errors
}
