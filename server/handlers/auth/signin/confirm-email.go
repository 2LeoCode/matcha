package signin

import (
	"matcha/components/auth/signin"
	"matcha/components/lib"
	"matcha/database"
	"matcha/handlers/auth"
	"net/http"

	"goji.io/pat"
)

func GetConfirmEmail(res http.ResponseWriter, req *http.Request) {
	if auth.GetOtpEntry(pat.Param(req, "user")) == nil {
		http.Redirect(res, req, "/not-found", http.StatusNotFound)
	} else {
		lib.Page("Confirm email", signin.ConfirmEmail(false)).Render(req.Context(), res)
	}
}

func PostConfirmEmail(res http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		res.WriteHeader(http.StatusBadRequest)
	} else {
		user := pat.Param(req, "user")
		inputOtp := req.Form.Get("otp")

		entry := auth.GetOtpEntry(user)
		switch {
		case entry == nil:
			res.WriteHeader(http.StatusBadRequest)
		case entry.Second != inputOtp:
			signin.ConfirmEmail(true).Render(req.Context(), res)
		default:
			http.Redirect(res, req, "/auth/signin/success", http.StatusSeeOther)
		}
	}
}
