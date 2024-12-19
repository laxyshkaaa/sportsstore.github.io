package handling2

import (
	"html/template"
	"net/http"
)

type LoginHandler struct{}

func (h LoginHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Обработка POST-запроса (например, проверка логина и пароля)
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Логика аутентификации, например, сравнение с базой данных
		if username == "admin" && password == "password" {
			// Пример успешной аутентификации
			http.Redirect(w, r, "/welcome", http.StatusSeeOther)
		} else {
			// Неверный логин/пароль
			renderTemplate(w, "login.html", "Invalid username or password.")
		}
	} else {
		// Обработка GET-запроса (отображение формы)
		renderTemplate(w, "login.html", "")
	}
}

type RegisterHandler struct{}

func (h RegisterHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Обработка POST-запроса регистрации
		password := r.FormValue("password")
		confirmPassword := r.FormValue("confirm_password")

		if password != confirmPassword {
			// Если пароли не совпадают
			renderTemplate(w, "register.html", "Passwords do not match.")
			return
		}

		// Логика для создания нового пользователя (например, сохранение в базе данных)
		// Здесь просто перенаправим на страницу входа
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		// Обработка GET-запроса (отображение формы регистрации)
		renderTemplate(w, "register.html", "")
	}
}

// Функция для рендеринга HTML-шаблонов
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplParsed, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmplParsed.Execute(w, data)
}
