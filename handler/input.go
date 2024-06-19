package handler

import (
	"context"
	"encoding/json"
	"log"
	"whatsapp-bot/domain"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var (
	sheetKey domain.Key = domain.Key{
		Type:                       "service_account",
		Project_id:                 "sturdy-rampart-414507",
		PrivateKeyId:               "f93d36b069b800113fffd57438e2e8157b864f6e",
		PrivateKey:                 "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDWs3maOaxWFrBX\nP80ySsEg1XI3mIzl2+FFtpLY6Jo9SwhoCXCqeEySm/7DZvKrQvMOf2lmlrRT/vSz\n/Alvjl7YmEdQRR35MfCwMqoDGpGzQzpa2htENu/CNoBFT0KuRI7RoR6wYD0r2Yiv\nYdx1PJFgxzl7Fw9jZGpqTM0zEg8wDFWc4PUFhrQ/pjrQeNpAKEmAHpc1CpgXf5T4\nxeDfIept0vDzvgOayLLaiebEHgqX2A76ZyMZMKQ0tojVn0ewxXFOIxur96uC7PrE\nwI4nmB+ZPUTWFjWcckFt7nPBkgpZnqZxWc12rzD8O+We2ssdJ5Lu6e0x573sXBM8\nzyZYn2oNAgMBAAECggEAFfxE7IURX7NtFVIkOg0rn6gbzROjAfq1I8VhXjLbcIFZ\neAXXcq1k8fJlcS7+lIkDiK2FZoryJlPKTQ/3C2Rh0/4r901Ml3MEuZZGkJDM8LVm\nqV0CQ8b4YEPieLZ4Fo5Qrai5EsiYGwBoncUApUSxgl6ERVWELuJef0okDSHBSmMY\n6gpJ2Bs00RCA/RttWMYCGhX7Ap7le8HeBhdo4haedK4g7AEzOE7r5wx3n29GVXYP\nWIdkPM5HGXTr0hoXqkuVzYE7sz/SR8bArCkSSaV3Elg5NhuWVo5bbKHGfDCR1pgG\nh4OLorxG5HjnUvRJi6wgbl/irMdHHZbbBkIJgNbw5QKBgQDyu2qwQ9qG16Le9Fak\neZsQveTSBArUqcakiP10J51OEDKFILgV5QmtSuVgDmvt3GtS2BvHWAQnSATmJbYE\nzCH5SA9yAj3jMULvUadqSqpUOylZWdsstzFBF4JfSexdujSstBOwkR+4bjqko+pZ\nALJLGGk8En9mz4wMA8/n4HjyawKBgQDib9D8JiXnvaOvkVnZJg53Wm0mU8ECOwz7\n4dLqd1/4Ne6i5wuakdhXxH/Xi6mTh6UpIjpkaAGTtUQ4o7UpjJZXv56B7yuXRzX8\niacG/ztbTQPgkBMKVvEMIA5ZlwBY2N4WmHW2uuqKVyiDoJXeB1ZqkfERrxGF9hhe\nTm7UxUPjZwKBgBGKeatDPAeCjcKO5bP9FbegAWkr4Zx8qF47iBVcx4Fcqkn0k1kB\ntDFRuEG4yRPnWdoiw98j/SjZi4EWOJEgL1RUYMlcYByK8stYm0CTJLvxiotqcuxI\nUMbjbAh2zx321tekK8gWtLQfQx2iusyzwC6tFO3CaHKFjRlRLTUUg6kBAoGAZIsZ\n1JxrwFaoU9DKgG0xhzWTuaz+yqrFvSIssQeIiYIJEDRfJqcT1QTfbyoIhgV5BKa6\nHtM/wAKA5FVsn6JZL2VaUH+Ob4GjKxtEGwtBo+yyiAfxLGomZiflt0ohUVIaaxYP\nTq/4bKz+xFONuCSx7mgXDq8ZrM31TAqvk7JzzokCgYAf38RFSi4K7zchae6jTF+R\npDcP+A6xuPLkdC7tLXdUMzpHGi9GPu8K+C4S4HySaDEWPuaucYq7GmnMIAna/Y53\ny2Qezpd4sZ3qmxeyZVtLiQeSZ7LmZNkR+RvExfUwg+tLqM0GIPctG+yK03uYsEpg\neay1pU4Iy0vGrfgMswGRzg==\n-----END PRIVATE KEY-----\n",
		ClientEmail:                "sheet-service-account@sturdy-rampart-414507.iam.gserviceaccount.com",
		ClientId:                   "113724890331456851600",
		AuthUri:                    "https://accounts.google.com/o/oauth2/auth",
		TokenUri:                   "https://oauth2.googleapis.com/token",
		AuthProvider_x509_cert_url: "https://www.googleapis.com/oauth2/v1/certs",
		Client_x509_cert_url:       "https://www.googleapis.com/robot/v1/metadata/x509/sheet-service-account%40sturdy-rampart-414507.iam.gserviceaccount.com",
	}
)

func getSheetConfig() *sheets.Service {
	credential, err := json.Marshal(sheetKey)
	if err != nil {
		log.Fatalf("Failed to get key: %v", err)
	}

	srv, err := sheets.NewService(context.Background(), option.WithCredentialsJSON(credential))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	return srv
}
