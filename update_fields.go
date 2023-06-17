package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/textproto"
	"net/url"
	"strings"
	"time"

	"github.com/koddr/gosl"
)

// updateField provides update process for the given transaction ID.
func (app *App) updateField(id, fieldName string, fieldValues []string) error {
	// Set primary key (PK) to the endpoint body.
	_, body := ModifyByValue(app.Config.API.UpdateEndpoint.EndpointBody, app.Config.ColumnsOrder[0], id)

	// Check, if field has a many values.
	if len(fieldValues) > 1 {
		// If true, set many values to the field.
		_, body = ModifyByValue(body, fieldName, fieldValues)
	} else {
		// If false, set the only one value to the field.
		_, body = ModifyByValue(body, fieldName, fieldValues[0])
	}

	// Marshaling endpoint body.
	reqBody, err := gosl.Marshal(&body)
	if err != nil {
		return err
	}

	// Set endpoint from the config.
	endpoint := app.Config.API.UpdateEndpoint.EndpointName

	// If endpoint want to set primary key (PK) to it.
	if app.Config.API.UpdateEndpoint.AddPKToEndpointName {
		// Set primary key (PK) to the endpoint from the config.
		endpoint = fmt.Sprintf(app.Config.API.UpdateEndpoint.EndpointName, id)
	}

	// Build API URL.
	apiUrl := url.URL{
		Scheme: strings.ToLower(app.Config.API.BaseURLSchema),
		Host:   strings.ToLower(app.Config.API.BaseURL),
		Path:   strings.ToLower(endpoint),
	}

	// Create HTTP request to API URL with endpoint body.
	req, err := http.NewRequest(app.Config.API.UpdateEndpoint.HTTPMethod, apiUrl.String(), bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	defer req.Body.Close()

	// Set request headers.
	req.Header.Set(
		textproto.CanonicalMIMEHeaderKey("authorization"),
		gosl.Concat(app.Config.API.AuthMethod, " ", app.Config.API.Token),
	)
	req.Header.Set(
		textproto.CanonicalMIMEHeaderKey("content-type"),
		app.Config.API.UpdateEndpoint.ContentType,
	)

	// Create a new HTTP client.
	client := &http.Client{
		// Set timeout count from the config.
		Timeout: time.Second * time.Duration(app.Config.API.RequestTimeout),
	}

	// Make HTTP request by the client.
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check, if response HTTP status < 300.
	if resp.StatusCode < http.StatusMultipleChoices {
		printStyled(
			fmt.Sprintf(
				"✓ Field '%s' with values '%s' in the transaction '%s' has been successfully updated (HTTP %d)",
				fieldName, fieldValues, id, resp.StatusCode,
			),
			"margin-left",
		)
	} else {
		// Else, return error message with the raw response body.
		return fmt.Errorf("transaction '%s' returned HTTP %d", id, resp.StatusCode)
	}

	return nil
}
