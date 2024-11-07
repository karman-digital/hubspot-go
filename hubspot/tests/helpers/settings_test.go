package adapter_test

import (
	"testing"

	hshelpers "github.com/karman-digital/hubspot/hubspot/adapter/helpers"
	hubspotmodels "github.com/karman-digital/hubspot/hubspot/api/models"
	"github.com/stretchr/testify/assert"
)

func TestConvertBodyBytesToAccountDataPayload(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		want    hubspotmodels.AccountDataPayload
		wantErr bool
	}{
		{
			name:  "Valid input",
			input: []byte(`{"actionType":"CREATE","portalId":"62515","userId":"12345","userEmail":"karman-digital","appId":"12345"}`),
			want: hubspotmodels.AccountDataPayload{
				ActionType: "CREATE",
				PortalId:   "62515",
				UserId:     "12345",
				UserEmail:  "karman-digital",
				AppId:      "12345",
			},
		},
		{
			name:    "Invalid JSON",
			input:   []byte(`invalid`),
			wantErr: true,
		},
		{
			name:    "Empty JSON",
			input:   []byte(``),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hshelpers.ConvertBodyBytesToAccountDataPayload(tt.input)

			if tt.wantErr {
				assert.Error(t, err, "ConvertBodyBytesToAccountDataPayload() should return an error")
			} else {
				assert.NoError(t, err, "ConvertBodyBytesToAccountDataPayload() should not return an error")
				assert.Equal(t, tt.want, got, "ConvertBodyBytesToAccountDataPayload() returned unexpected value")
			}
		})
	}
}
