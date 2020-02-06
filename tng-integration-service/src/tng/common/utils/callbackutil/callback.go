package callbackutil

import (
	"encoding/base64"
	"encoding/hex"
	"net/url"

	"tng/common/utils/cfgutil"
	"tng/common/utils/cryptoutil"
)

// BuildCallbackURL encode and gen he callback URL
func BuildCallbackURL(ourURL, inputURL, module string) (string, error) {
	key := []byte(cfgutil.Load("CALLBACK_KEY"))
	nonce, err := hex.DecodeString(cfgutil.Load("CALLBACK_NONCE"))
	if err != nil {
		return "", err
	}
	encryptedURL, err := cryptoutil.AESGCMEncrypt(key, nonce, []byte(inputURL))
	if err != nil {
		return "", err
	}

	u, err := url.Parse(ourURL)
	if err != nil {
		return "", err
	}
	q := u.Query()
	q.Set("module", module)
	q.Set("embedded_data", base64.StdEncoding.EncodeToString(encryptedURL))
	u.RawQuery = q.Encode()
	return u.String(), nil
}
