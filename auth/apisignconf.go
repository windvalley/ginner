package auth

type keySecret struct {
	MD5 string
	AES string
	RSA string
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
			MD5: "fjadoifjadjfqjowerqfdafafdjafl",
			AES: "707c8d56d87a5650ae6492e67be6ffc4", // length must be 16
			RSA: "auth/rsa/id_rsa.pub",
		},
		AppName: "robot",
		Roles:   []string{"access"},
	},
}
