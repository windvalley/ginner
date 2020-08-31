package auth

type _RSA struct {
	Private string
	Public  string
}

type keySecret struct {
	MD5  string
	AES  string
	Hmac string
	JWT  string
	RSA  _RSA
}

type userInfo struct {
	keySecret
	AppName string
	Roles   []string
}

// NOTE: The users data should be stored in database in real production.
var userInfos = map[string]userInfo{
	"keyid_3rqjdjfde33derljl": {
		keySecret: keySecret{
			MD5:  "fjadoifjadjfqjowerqfdafafdjafl",
			AES:  "707c8d56d87a5650ae6492e67be6ffc4", // length must be 16
			Hmac: "b4984088af5b2dd6236b1aa5d51aa3c4",
			JWT:  "f2c40107cae3ee4e16270150b513dba0",
			RSA: _RSA{
				"auth/rsa/private.pem",
				"auth/rsa/public.pem",
			},
		},
		AppName: "robot",
		Roles:   []string{"access"},
	},
}
