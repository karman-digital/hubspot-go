package hsvalidate

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"hash"
	"math"
	"net/url"
	"strconv"
	"time"
)

func ValidateWebhookSignature(secret []byte, host string, urlPath string, timestamp string, method string, signature string, body []byte) error {
	checkSum := ""
	if method == "POST" {
		if isJSON(body) {
			JSONString, err := toCompactJSONString(body)
			if err != nil {
				return err
			}
			checkSum = fmt.Sprintf("%shttps://%s%s%s%s", method, host, urlPath, JSONString, timestamp)
		} else {
			formDataString, err := toFormDataString(body)
			if err != nil {
				return err
			}
			checkSum = fmt.Sprintf("%shttps://%s%s%s%s", method, host, urlPath, formDataString, timestamp)
		}
	} else {
		checkSum = fmt.Sprintf("%shttps://%s%s%s", method, host, urlPath, timestamp)
	}
	hash := encrypt(secret, checkSum)
	if base64.StdEncoding.EncodeToString(hash.Sum(nil)) != signature {
		return fmt.Errorf("signatures mismatched sent: %s, generated: %s", signature, base64.StdEncoding.EncodeToString(hash.Sum(nil)))
	}
	return nil
}

func encrypt(secret []byte, input string) hash.Hash {
	hash := hmac.New(sha256.New, secret)
	hash.Write([]byte(input))
	return hash
}

func toCompactJSONString(input []byte) (string, error) {
	buffer := new(bytes.Buffer)
	if err := json.Compact(buffer, input); err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func isJSON(input []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(input, &js) == nil
}

func toFormDataString(input []byte) (string, error) {
	data, err := url.ParseQuery(string(input))
	if err != nil {
		return "", err
	}
	fmt.Println(data.Encode())
	return data.Encode(), nil
}

func ValidateTimeStamp(timestamp string) error {
	timestampInt, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return errors.New("could not parse timestamp")
	}
	timeUnix := time.Unix(int64(math.Round(float64(timestampInt/1000))), 0)
	fiveMinutesAgo := time.Now().Add(-5 * time.Minute)
	if timeUnix.Before(fiveMinutesAgo) {
		return errors.New("timestamp more than 5 minutes ago")
	} else if timeUnix.After(time.Now()) {
		return errors.New("timestamp is in the future")
	}
	return nil
}
