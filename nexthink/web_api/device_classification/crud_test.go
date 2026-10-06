package device_classification

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/device_classification/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func load[T any](t *testing.T, name string) *T {
	t.Helper()
	var value T
	require.NoError(t, json.Unmarshal(mocks.Fixture(name), &value))
	return &value
}

type contract struct {
	name, verb, path                 string
	body, multipart, download, empty bool
	call                             func(*Service) (any, *interfaces.Response, error)
}

func contracts(t *testing.T) []contract {
	ctx := context.Background()
	return []contract{{name: "GetOrganization", verb: "GET", path: "/apigateway/entity-manager/api/v1/entities", body: false, multipart: false, download: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetOrganization(ctx) }},
		{name: "CreateOrganization", verb: "POST", path: "/apigateway/entity-manager/api/v1/entities", body: true, multipart: true, download: false, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			r, e := s.CreateOrganization(ctx, load[RulesetUpload](t, "CreateOrganization_request"))
			return nil, r, e
		}},
		{name: "UpdateOrganization", verb: "PUT", path: "/apigateway/entity-manager/api/v1/entities", body: true, multipart: true, download: false, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			r, e := s.UpdateOrganization(ctx, load[RulesetUpload](t, "UpdateOrganization_request"))
			return nil, r, e
		}},
		{name: "DownloadOrganization", verb: "GET", path: "/apigateway/entity-manager/api/v1/entities/download", body: false, multipart: false, download: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.DownloadOrganization(ctx) }},
		{name: "GetLocationType", verb: "GET", path: "/apigateway/entity-manager/api/v1/location-type", body: false, multipart: false, download: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetLocationType(ctx) }},
		{name: "CreateLocationType", verb: "POST", path: "/apigateway/entity-manager/api/v1/location-type", body: true, multipart: true, download: false, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			r, e := s.CreateLocationType(ctx, load[RulesetUpload](t, "CreateLocationType_request"))
			return nil, r, e
		}},
		{name: "UpdateLocationType", verb: "PUT", path: "/apigateway/entity-manager/api/v1/location-type", body: true, multipart: true, download: false, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			r, e := s.UpdateLocationType(ctx, load[RulesetUpload](t, "UpdateLocationType_request"))
			return nil, r, e
		}},
		{name: "DownloadLocationType", verb: "GET", path: "/apigateway/entity-manager/api/v1/location-type/download", body: false, multipart: false, download: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.DownloadLocationType(ctx) }},
		{name: "GetVPNEgress", verb: "GET", path: "/apigateway/entity-manager/api/v1/vpn-egress", body: false, multipart: false, download: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetVPNEgress(ctx) }},
		{name: "CreateVPNEgress", verb: "POST", path: "/apigateway/entity-manager/api/v1/vpn-egress", body: true, multipart: true, download: false, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			r, e := s.CreateVPNEgress(ctx, load[RulesetUpload](t, "CreateVPNEgress_request"))
			return nil, r, e
		}},
		{name: "UpdateVPNEgress", verb: "PUT", path: "/apigateway/entity-manager/api/v1/vpn-egress", body: true, multipart: true, download: false, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			r, e := s.UpdateVPNEgress(ctx, load[RulesetUpload](t, "UpdateVPNEgress_request"))
			return nil, r, e
		}},
		{name: "DownloadVPNEgress", verb: "GET", path: "/apigateway/entity-manager/api/v1/vpn-egress/download", body: false, multipart: false, download: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.DownloadVPNEgress(ctx) }},
		{name: "DeleteVPNEgress", verb: "DELETE", path: "/apigateway/entity-manager/api/v1/vpn-egress", body: false, multipart: false, download: false, empty: true, call: func(s *Service) (any, *interfaces.Response, error) { r, e := s.DeleteVPNEgress(ctx); return nil, r, e }},
		{name: "GetGeoIP", verb: "GET", path: "/apigateway/entity-manager/api/v1/configuration/geoip", body: false, multipart: false, download: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetGeoIP(ctx) }},
		{name: "UpdateGeoIP", verb: "PUT", path: "/apigateway/entity-manager/api/v1/configuration/geoip", body: true, multipart: false, download: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.UpdateGeoIP(ctx, load[GeoIPConfiguration](t, "UpdateGeoIP_request"))
		}}}
}
func TestContracts(t *testing.T) {
	for _, tt := range contracts(t) {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				if tt.body {
					if tt.multipart {
						checkMultipart(t, r)
					} else {
						body, err := io.ReadAll(r.Body)
						require.NoError(t, err)
						assert.JSONEq(t, string(mocks.Fixture(tt.name+"_request")), string(body))
						assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
					}
				}
				status, body, contentType := 200, "", "application/json"
				if tt.empty {
					status = 204
				} else if tt.download {
					require.NoError(t, json.Unmarshal(mocks.Fixture(tt.name+"_success"), &body))
					contentType = "text/csv"
				} else {
					body = string(mocks.Fixture(tt.name + "_success"))
				}
				response := httpmock.NewStringResponse(status, body)
				response.Header.Set("Content-Type", contentType)
				response.Header.Set("X-Request-ID", "fixture-request")
				return response, nil
			})
			result, response, err := tt.call(NewService(transport))
			require.NoError(t, err)
			require.NotNil(t, response)
			assert.Equal(t, "fixture-request", response.Headers.Get("X-Request-ID"))
			assert.Equal(t, 1, mock.GetTotalCallCount())
			if tt.empty {
				assert.Nil(t, result)
				assert.Equal(t, 204, response.StatusCode)
			} else if tt.download {
				var expected string
				require.NoError(t, json.Unmarshal(mocks.Fixture(tt.name+"_success"), &expected))
				assert.Equal(t, []byte(expected), result)
			} else {
				actual, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture(tt.name+"_success")), string(actual))
			}
		})
	}
}
func TestHTTPErrors(t *testing.T) {
	for _, tt := range contracts(t) {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, httpmock.NewStringResponder(404, string(mocks.Fixture("error"))))
			_, response, err := tt.call(NewService(transport))
			require.Error(t, err)
			require.NotNil(t, response)
			assert.Equal(t, 404, response.StatusCode)
			assert.JSONEq(t, string(mocks.Fixture("error")), string(response.Body))
		})
	}
}

func checkMultipart(t *testing.T, r *http.Request) {
	t.Helper()
	assert.Equal(t, "SDK ruleset", r.Header.Get("x-nxt-entities-name"))
	assert.Equal(t, "Q2Fmw6kgbGFi", r.Header.Get("x-nxt-entities-description"))
	reader, err := r.MultipartReader()
	require.NoError(t, err)
	part, err := reader.NextPart()
	require.NoError(t, err)
	assert.Equal(t, "file", part.FormName())
	assert.Equal(t, "fixture.csv", part.FileName())
	assert.Equal(t, "text/csv", part.Header.Get("Content-Type"))
	body, err := io.ReadAll(part)
	require.NoError(t, err)
	assert.Equal(t, "source,value\r\nfixture,example\r\n", string(body))
	_, err = reader.NextPart()
	require.ErrorIs(t, err, io.EOF)
}
func TestMetadataOnlyUpdate(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	mock.RegisterResponder("PUT", testutil.BaseURL+Endpoint+"/entities", func(r *http.Request) (*http.Response, error) {
		reader, err := r.MultipartReader()
		require.NoError(t, err)
		_, err = reader.NextPart()
		require.ErrorIs(t, err, io.EOF)
		return httpmock.NewStringResponse(204, ""), nil
	})
	_, err := NewService(transport).UpdateOrganization(context.Background(), &RulesetUpload{Name: "SDK ruleset", Description: "metadata only"})
	require.NoError(t, err)
}
func TestUploadValidation(t *testing.T) {
	require.Error(t, validateUpload(nil, true))
	require.Error(t, validateUpload(&RulesetUpload{Name: "test"}, true))
	require.Error(t, validateUpload(&RulesetUpload{Name: "test", Filename: "test.csv", CSV: []byte{0xff}}, true))
	require.Error(t, validateUpload(&RulesetUpload{Name: "bad\r\nheader"}, false))
	require.NoError(t, validateUpload(&RulesetUpload{Name: "test"}, false))
	request := load[RulesetUpload](t, "CreateOrganization_request")
	require.NoError(t, validateUpload(request, true))
}

func TestRulesetPreservesNullAndUnknownFields(t *testing.T) {
	raw := `{"name":"fixture","description":null,"filename":"f.csv","lastModified":1791295200,"new_field":{"enabled":false}}`
	var value Ruleset
	require.NoError(t, json.Unmarshal([]byte(raw), &value))
	assert.Equal(t, "fixture", value.Name)
	encoded, err := json.Marshal(value)
	require.NoError(t, err)
	assert.JSONEq(t, raw, string(encoded))
	value.Name = "updated"
	encoded, err = json.Marshal(value)
	require.NoError(t, err)
	assert.Contains(t, string(encoded), `"updated"`)
}
func TestUnicodeFilename(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	mock.RegisterResponder("POST", testutil.BaseURL+Endpoint+"/entities", func(r *http.Request) (*http.Response, error) {
		reader, err := r.MultipartReader()
		require.NoError(t, err)
		part, err := reader.NextPart()
		require.NoError(t, err)
		assert.Equal(t, "département.csv", part.FileName())
		return httpmock.NewStringResponse(204, ""), nil
	})
	_, err := NewService(transport).CreateOrganization(context.Background(), &RulesetUpload{Name: "test", Filename: "département.csv", CSV: []byte("column\nvalue\n")})
	require.NoError(t, err)
}
func TestNilGeoIPRequestDoesNotSend(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	_, _, err := NewService(transport).UpdateGeoIP(context.Background(), nil)
	require.Error(t, err)
	assert.Equal(t, 0, mock.GetTotalCallCount())
}

func TestMalformedJSON(t *testing.T) {
	for _, tt := range contracts(t) {
		if tt.empty || tt.download {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, func(r *http.Request) (*http.Response, error) {
				response := httpmock.NewStringResponse(200, "{broken")
				response.Header.Set("Content-Type", "application/json")
				return response, nil
			})
			_, response, err := tt.call(NewService(transport))
			require.Error(t, err)
			require.NotNil(t, response)
			assert.Equal(t, 200, response.StatusCode)
		})
	}
}
