package echo

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

type (
	testParams struct {
		reqParams        url.Values
		reqParamNames    []string
		reqParamValues   []string
		reqBody          []byte
		wantError        error
		wantErrorMessage string
		wantStatus       int
		wantContentType  string
		wantBodyData     []byte
	}
)

const (
	testOk   = "\u2713"
	testFail = "\u2717"
)

var (
	e = echo.New()
)

func TestRootHandler(t *testing.T) {
	res, err := execHandler(t, rootHandler, nil, nil, nil, nil)
	if err != nil {
		t.Fatalf("\t%s\tERROR: handler failed with error: %v", testFail, err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("\t%s\tERROR: got handler http status: %d, want %d", testFail, res.StatusCode, http.StatusOK)
	}
	bodyData, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("\t%s\tERROR: failed to read body data, error: %v", testFail, err)
	}
	res.Body.Close()
	wantedBody := "Hello Gopher!"
	if string(bodyData) != wantedBody {
		t.Errorf("\t%s\tERROR: got body %s, want %s", testFail, bodyData, wantedBody)
	}
}

func TestGetItemHandler(t *testing.T) {
	for _, tp := range []struct {
		testParams
		name string
	}{
		{
			name: "first_item_id",
			testParams: testParams{
				reqParamNames:   []string{"id"},
				reqParamValues:  []string{"1"},
				wantStatus:      http.StatusOK,
				wantContentType: echo.MIMEApplicationJSONCharsetUTF8,
				wantBodyData:    []byte(`{"id":1,"name":"first item"}` + "\n"),
			},
		},
		{
			name: "other_item_id",
			testParams: testParams{
				reqParamNames:   []string{"id"},
				reqParamValues:  []string{"2"},
				wantStatus:      http.StatusOK,
				wantContentType: echo.MIMEApplicationJSONCharsetUTF8,
				wantBodyData:    []byte(`{"id":2,"name":"other item"}` + "\n"),
			},
		},
		{
			name: "item_id_not_found",
			testParams: testParams{
				reqParamNames:   []string{"id"},
				reqParamValues:  []string{"42"},
				wantStatus:      http.StatusNotFound,
				wantContentType: echo.MIMEApplicationJSONCharsetUTF8,
				wantBodyData:    []byte(`""` + "\n"),
			},
		},
		{
			name: "no_item_id",
			testParams: testParams{
				wantErrorMessage: `strconv.Atoi: parsing "": invalid syntax`,
			},
		},
	} {
		t.Run(tp.name, testEchoHandler(getItemHandler, tp.testParams))
	}
}

func TestPostItemHandler(t *testing.T) {
	item2BodyBytes := func(ID int, name string) []byte {
		data, err := json.Marshal(&item{ID: ID, Name: name})
		if err != nil {
			t.Fatalf("ERROR: failed to marshal item, error: %v", err)
		}
		return data
	}
	for _, tp := range []struct {
		testParams
		name string
	}{
		{
			name: "first_item_id",
			testParams: testParams{
				reqParamNames:   []string{"id"},
				reqParamValues:  []string{"1"},
				reqBody:         item2BodyBytes(1, "first item"),
				wantStatus:      http.StatusOK,
				wantContentType: echo.MIMEApplicationJSONCharsetUTF8,
				wantBodyData:    []byte(`{"id":1,"name":"first item"}` + "\n"),
			},
		},
		{
			name: "other_item_id",
			testParams: testParams{
				reqParamNames:   []string{"id"},
				reqParamValues:  []string{"2"},
				reqBody:         item2BodyBytes(2, "other item"),
				wantStatus:      http.StatusOK,
				wantContentType: echo.MIMEApplicationJSONCharsetUTF8,
				wantBodyData:    []byte(`{"id":2,"name":"other item"}` + "\n"),
			},
		},
		{
			name: "itemid_item_mismatch",
			testParams: testParams{
				reqParamNames:   []string{"id"},
				reqParamValues:  []string{"3"},
				reqBody:         item2BodyBytes(4, "some other item"),
				wantStatus:      http.StatusBadRequest,
				wantContentType: echo.MIMEApplicationJSONCharsetUTF8,
				wantBodyData:    []byte(`""` + "\n"),
			},
		},
		{
			name: "no_item_id",
			testParams: testParams{
				wantErrorMessage: `strconv.Atoi: parsing "": invalid syntax`,
			},
		},
	} {
		t.Run(tp.name, testEchoHandler(postItemHandler, tp.testParams))
	}
}

func execHandler(t *testing.T, handler func(c echo.Context) error, params url.Values, paramNames, paramValues []string, body []byte) (*http.Response, error) {
	uri, err := url.Parse("http://test.local")
	if err != nil {
		t.Fatalf("\t%s\tERROR: Failed to parse test URL: %v", testFail, err)
	}
	if params != nil {
		uri.RawQuery = params.Encode()
	}

	var bodyReader io.Reader
	if len(body) > 0 {
		bodyReader = bytes.NewReader(body)
	}

	req := httptest.NewRequest("", uri.String(), bodyReader)
	if bodyReader != nil {
		req.Header.Add("Content-Type", echo.MIMEApplicationJSONCharsetUTF8)
	}
	res := httptest.NewRecorder()
	ctx := e.NewContext(req, res)
	ctx.SetParamNames(paramNames...)
	ctx.SetParamValues(paramValues...)
	err = handler(ctx)
	return res.Result(), err
}

func validateResponse(t *testing.T, res *http.Response, status int, ct string, body []byte) {
	if res.StatusCode != status {
		t.Errorf("\t%s\tERROR: got staus: %d, want: %d", testFail, res.StatusCode, status)
	} else {
		t.Logf("\t%s\treturned status code OK", testOk)
	}

	if res.Header.Get(echo.HeaderContentType) != ct {
		t.Errorf("\t%s\tERROR: got content type: %s, want: %s", testFail, res.Header.Get(echo.HeaderContentType), ct)
	} else {
		t.Logf("\t%s\tcontent type OK", testOk)
	}

	bd, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("\t%s\tERROR: failed to read response body, error: %v", testFail, err)
	}
	if !reflect.DeepEqual(bd, body) {
		t.Errorf("\t%s\tERROR, got body: %s, want: %s", testFail, bd, body)
	} else {
		t.Logf("\t%s\tbody data OK", testOk)
	}
}

func testEchoHandler(h func(c echo.Context) error, tp testParams) func(t *testing.T) {
	return func(t *testing.T) {
		res, err := execHandler(t, h, tp.reqParams, tp.reqParamNames, tp.reqParamValues, tp.reqBody)
		if tp.wantError != nil {
			if err != tp.wantError {
				t.Fatalf("\t%s\tERROR: got error %v from handler, want: %v", testFail, err, tp.wantError)
			} else {
				t.Logf("\t%s\thandler returned with expected error (%v)", testOk, err)
			}
			return
		}
		if tp.wantErrorMessage != "" {
			if err == nil {
				t.Fatalf("\t%s\tERROR: handler did not return an error", testFail)
			} else if err.Error() != tp.wantErrorMessage {
				t.Fatalf("\t%s\tERROR: got error: %v, want: %s", testFail, err, tp.wantErrorMessage)
			} else {
				t.Logf("\t%s\thandler returned with expected error (%v)", testOk, err)
			}
			return
		}
		if err != nil {
			t.Fatalf("\t%s\tERROR: handler returned unexpected error: %v", testFail, err)
		} else {
			validateResponse(t, res, tp.wantStatus, tp.wantContentType, tp.wantBodyData)
		}
	}
}
