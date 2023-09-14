package server

func CreateServerListPingResponse() string {
	return `{"version":{"name": "Koperwiek 1.20.1","protocol": 763},
		"players": {"max":` + "420," + `"online": ` + "1" + `,"sample":[]},
		"description": {
			"text": "This is definitely an motd\n         koperwiek testing"
		},"enforcesSecureChat": true,"previewsChat": true}`
}
