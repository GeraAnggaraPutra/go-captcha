// example of HTTP server that uses the captcha package.
package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"log"
	"net/http"

	"github.com/mojocn/base64Captcha"

	"go-captcha/captcha"
)

var (
	driver = base64Captcha.NewDriverString(
		65,
		240,
		0,
		14,
		6,
		"abcdefghkmnpqrstuvwxyz0123456789",
		&color.RGBA{0, 0, 0, 0},
		nil,
		[]string{"wqy-microhei.ttc", "ApothecaryFont.ttf"},
	)
	store = base64Captcha.DefaultMemStore
	cpt   = captcha.NewCaptchaHandler(base64Captcha.NewCaptcha(driver, store))
)

// base64Captcha create http handler
func generateCaptcha(w http.ResponseWriter, r *http.Request) {
	data, err := cpt.GenerateCaptchaHandler()
	body := map[string]interface{}{"code": 1, "data": data, "msg": "success"}
	if err != nil {
		body = map[string]interface{}{"code": 0, "msg": err.Error()}
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(body)
}

// base64Captcha verify http handler
func captchaVerify(w http.ResponseWriter, r *http.Request) {
	captchaID := r.FormValue("captcha_id")
	captchaAnswer := r.FormValue("captcha_answer")

	match := cpt.VerifyCaptchaHandler(captchaID, captchaAnswer)
	body := map[string]interface{}{"code": 1, "msg": "success"}
	if !match {
		body = map[string]interface{}{"code": 0, "msg": "invalid captcha"}
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(body)
}

// start a net/http server
func main() {
	//api for create captcha
	http.HandleFunc("/api/getCaptcha", generateCaptcha)

	//api for verify captcha
	http.HandleFunc("/api/verifyCaptcha", captchaVerify)

	fmt.Println("Server is at :8777")
	if err := http.ListenAndServe(":8777", nil); err != nil {
		log.Fatal(err)
	}
}
