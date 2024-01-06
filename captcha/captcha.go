package captcha

import (
	"github.com/mojocn/base64Captcha"
)

type CaptchaPayload struct {
	Image string `json:"image"`
	ID    string `json:"captcha_id"`
}

type CaptchaHandler struct {
	Captcha *base64Captcha.Captcha
}

func NewCaptchaHandler(captcha *base64Captcha.Captcha) *CaptchaHandler {
	return &CaptchaHandler{Captcha: captcha}
}

func (h *CaptchaHandler) GenerateCaptchaHandler() (res CaptchaPayload, err error) {
	id, b64s, _, err := h.Captcha.Generate()
	if err != nil {
		return
	}

	res = CaptchaPayload{
		Image: b64s,
		ID:    id,
	}

	return
}

func (h *CaptchaHandler) VerifyCaptchaHandler(id, answer string) bool {
	match := h.Captcha.Verify(id, answer, true)

	if match {
		return true
	} else {
		return false
	}
}
