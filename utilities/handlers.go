package utilities

import (
	"net/http"
	"strconv"
	"text/template"
)

func FormHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/form.html")
	t.Execute(w, nil)
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	num1, err1 := strconv.Atoi(r.FormValue("num1"))
	num2, err2 := strconv.Atoi(r.FormValue("num2"))
	operation := r.FormValue("operation")

	if err1 != nil || err2 != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	var result int
	switch operation {
	case "add":
		result = num1 + num2
	case "sub":
		result = num1 - num2
	case "mul":
		result = num1 * num2
	case "div":
		if num2 == 0 {
			http.Error(w, "Cannot divide by zero", http.StatusBadRequest)
			return
		}
		result = num1 / num2
	case "mod":
		if num2 == 0 {
			http.Error(w, "Cannot mod by zero", http.StatusBadRequest)
			return
		}
		result = num1 % num2
	default:
		http.Error(w, "Invalid operation", http.StatusBadRequest)
		return
	}

	data := struct {
		Result int
	}{
		Result: result,
	}

	t, _ := template.ParseFiles("templates/result.html")
	t.Execute(w, data)
}
