package auth

type RSA struct {
	Private string
	Public  string
}

type keySecret struct {
	MD5  string
	AES  string
	Hmac string
	RSA
}

type userInfo struct {
	keySecret
	AppName string
	Roles   []string
}

// NOTE: The users data should be stored in database in real production.
var UserInfos = map[string]userInfo{
	"keyid_3rqjdjfde33derljl": userInfo{
		keySecret: keySecret{
			MD5:  "fjadoifjadjfqjowerqfdafafdjafl",
			AES:  "707c8d56d87a5650ae6492e67be6ffc4", // length must be 16
			Hmac: "707c8d56d87a5650ae6492e67be6ffc4",
			RSA: RSA{
				"auth/rsa/private.pem",
				"auth/rsa/public.pem",
			},
		},
		AppName: "robot",
		Roles:   []string{"access"},
	},
}
