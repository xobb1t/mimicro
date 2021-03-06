package response

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUnmarshallingOnlyBody(t *testing.T) {
	str := `{"body":"OK"}`

	var response Response
	err := json.Unmarshal([]byte(str), &response)
	assert.Nil(t, err)

	assert.Equal(t, "OK", response.Body)
	assert.Equal(t, "text/plain", response.ContentType)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestUnmarshallingAllFields(t *testing.T) {
	str := `{"body":"{}", "content_type":"application/json", "status_code": 201}`

	var response Response
	err := json.Unmarshal([]byte(str), &response)
	assert.Nil(t, err)

	assert.Equal(t, "{}", response.Body)
	assert.Equal(t, "application/json", response.ContentType)
	assert.Equal(t, http.StatusCreated, response.StatusCode)
}

func TestWriteResponse(t *testing.T) {
	w := httptest.NewRecorder()
	var response = Response{"{\"a\":1}", "application/json", http.StatusCreated}

	response.WriteResponse(w)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.Equal(t, "{\"a\":1}", string(body))
}
