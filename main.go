package twofa

import "github.com/pquerna/otp/totp"

// two-factor authentication

// Generate 2FA, you should input a id about account name, it will output 2FA secret and qrcode url.
// 生成2FA
func GenerateTwoFA(id string) (string, string) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "NyaLog",
		AccountName: id,
	})
	if err != nil {
		return "", ""
	}
	return key.Secret(), key.URL()
}

// Validate 2FA, you should input the 2FA code and secret, it will output a number is it right or not
// 验证2FA
func Validate2FA(code string, secret string) int {
	validate := totp.Validate(code, secret)
	if validate {
		return 0
	}
	return 1
}
