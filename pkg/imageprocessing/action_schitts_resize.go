package imageprocessing

import (
	"errors"
	"fmt"
	"github.com/moooooooooose/mews/pkg/sheetsapi"
	"google.golang.org/api/sheets/v4"
	"image"
	"net/http"
)

type actionGoogleSheetResize struct {
	SheetId string
	AuthToken string
}

var _ ImageAction = actionGoogleSheetResize{}

func NewGoogleSheetResize(sheetId, authToken string) ImageAction {
	return &actionGoogleSheetResize{
		SheetId: sheetId,
		AuthToken: authToken,
	}
}

func (a actionGoogleSheetResize) Transform(image image.Image) (image.Image, error) {
	spreadsheet, err := getSheetConfiguration(a.SheetId, a.AuthToken)
	if err != nil {
		return nil, fmt.Errorf("error getting sheet information %w", err)
	}

	if len(spreadsheet.Sheets) == 0 {
		return nil, errors.New("error reading sheets meta data, no sheets founds")
	}

	sheetColumns := spreadsheet.Sheets[0].Properties.GridProperties.ColumnCount
	sheetRows := spreadsheet.Sheets[0].Properties.GridProperties.RowCount

	imageBounds := image.Bounds()
	if imageBounds.Dx() > imageBounds.Dy() {
		return resize(image, uint(sheetColumns), 0)
	}
	return resize(image, 0, uint(sheetRows))
}

func getSheetConfiguration(sheetId, authToken string) (*sheets.Spreadsheet, error) {
	sheetsClient := sheetsapi.NewClient(http.DefaultClient)
	return sheetsClient.Get(sheetId, sheetsapi.RequestOptions{AuthToken: authToken})
}



