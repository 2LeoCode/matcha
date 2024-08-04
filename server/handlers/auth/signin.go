package auth

import (
	"context"
	"fmt"
	"matcha/components/auth"
	"matcha/components/lib"
	"matcha/utils"
	"net/http"
	"regexp"
	"slices"
	"sync/atomic"
	"time"
)

var OtpEntries atomic.Value = func() atomic.Value {
	var entries atomic.Value
	entries.Store([]*utils.Pair[string, string]{})
	return entries
}()

func AddOtpEntry(username, otp string) {
	OtpEntries.Store(
		append(
			OtpEntries.Load().([]*utils.Pair[string, string]),
			utils.NewPair(username, otp),
		),
	)
}

func GetOtpEntry(username string) *utils.Pair[string, string] {
	entries := OtpEntries.Load().([]*utils.Pair[string, string])
	idx := slices.IndexFunc(
		entries,
		func(entry *utils.Pair[string, string]) bool {
			return entry.First == username
		},
	)
	if idx == -1 {
		return nil
	}
	return entries[idx]
}

func RemoveOtpEntry(username string) bool {
	entries := OtpEntries.Load().([]*utils.Pair[string, string])
	idx := slices.IndexFunc(
		entries,
		func(entry *utils.Pair[string, string]) bool {
			return entry.First == username
		},
	)
	if idx == -1 {
		return false
	}
	entries[idx] = entries[len(entries)-1]
	OtpEntries.Store(entries[:len(entries)-1])
	return true
}

func GetSignin(res http.ResponseWriter, req *http.Request) {
	lib.Page("Sign-in", auth.Signin(false, false, false, false)).Render(req.Context(), res)
}

func PostSignin(res http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	// email := req.Form.Get("email")
	username := req.Form.Get("username")
	password := req.Form.Get("password")
	confirmPassword := req.Form.Get("confirm-password")

	usernameCheck := regexp.MustCompile(`^[A-Za-z0-9_]{3,16}$`)
	passwordCheck := regexp.MustCompile(`^[ -~]{3,16}$`)

	invalidUsername := false
	invalidPassword := false
	passwordDontMatch := false
	if !usernameCheck.MatchString(username) {
		invalidUsername = true
	}
	if !passwordCheck.MatchString(password) {
		invalidPassword = true
	}
	if password != confirmPassword {
		passwordDontMatch = true
	}

	if invalidUsername || invalidPassword || passwordDontMatch {
		auth.Signin(false, invalidUsername, invalidPassword, passwordDontMatch)
	} else {

		otp := "000000" // generate otp

		// try to send email
		if false { // if we fail to send email
			auth.Signin(true, false, false, false).Render(req.Context(), res)
		} else {
			AddOtpEntry(username, otp)
			go func() {
				time.Sleep(5 * time.Minute)
				RemoveOtpEntry(username)
			}()
			http.Redirect(res, req, fmt.Sprintf("/auth/signin/confirm-email/%s", username), http.StatusNetworkAuthenticationRequired)
		}

	}
}
