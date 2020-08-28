package sheetsapi

import (
	"encoding/json"
	"fmt"
	"github.com/moooooooooose/mews/pkg/errorsutil"
	"google.golang.org/api/sheets/v4"
	"net/http"
)

const (
	GoogleSheetsAPIUrlBase = "https://sheets.googleapis.com/v4/spreadsheets"

	HeaderAuthorization = "Authorization"
)

type RequestOptions struct {
	AuthToken string
}

type Client interface {
	Get(id string, opts RequestOptions) (*sheets.Spreadsheet, error)
	Update(id string, opts RequestOptions) error
}

var _ Client = &client{}

type client struct {
	httpClient *http.Client
}

func NewClient(httpClient *http.Client) Client {
	return &client{
		httpClient: httpClient,
	}
}

func (c client) Get(id string, opts RequestOptions) (*sheets.Spreadsheet, error) {
	if id == "" {
		return nil, errorsutil.NotDefinedError("id")
	}
	if opts.AuthToken == "" {
		return nil, errorsutil.NotDefinedError("opts.AuthToken")
	}

	req, err := http.NewRequest(http.MethodGet, buildGetSpreadsheetUrlFromID(id), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build get spreadsheet request: %w", err)
	}
	setBearerHeaderFromToken(req, opts.AuthToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get spreadsheet: %w", err)
	}
	defer resp.Body.Close()

	spreadsheet := &sheets.Spreadsheet{}
	if err = json.NewDecoder(resp.Body).Decode(spreadsheet); err != nil {
		return nil, fmt.Errorf("failed to decode get spreadsheet response: %w", err)
	}

	return spreadsheet, nil
}

func (c client) Update(id string, opts RequestOptions) error {
	panic("implement me")
}

func buildGetSpreadsheetUrlFromID(id string) string {
	return fmt.Sprintf("%s/%s", GoogleSheetsAPIUrlBase, id)
}

func setBearerHeaderFromToken(request *http.Request, token string) {
	request.Header.Set(HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
}
