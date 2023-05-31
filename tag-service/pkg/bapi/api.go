package bapi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"io/ioutil"
	"net/http"
)

type API struct {
	URL    string
	Client *http.Client
}

type AccessToken struct {
	Token string `json:"token"`
}

const (
	APP_KEY    = "eddycjy"
	APP_SECRET = "go-programming-tour-book"
)

func (a *API) getAccessToken(ctx context.Context) (string, error) {
	body, err := a.httpGet(ctx, fmt.Sprintf("%s?app_key=%s&app_secret=%s", "auth", APP_KEY, APP_SECRET))
	if err != nil {
		return "", err
	}

	var accessToken AccessToken
	err = json.Unmarshal(body, &accessToken)
	if err != nil {
		return "", err
	}

	return accessToken.Token, nil
}

func (a *API) httpGet(ctx context.Context, path string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", a.URL, a.Client)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	span, _ := opentracing.StartSpanFromContext(
		ctx, "HTTP GET: "+a.URL,
		opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
	)
	span.SetTag("url", url)
	_ = opentracing.GlobalTracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)

	req = req.WithContext(context.Background())
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	defer span.Finish()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func NewAPI(url string) *API {
	return &API{URL: url}
}

func (a *API) GetTagList(ctx context.Context, name string) ([]byte, error) {
	//token, err := a.getAccessToken(ctx)
	//if err != nil {
	//	return nil, err
	//}

	body, err := a.httpGet(ctx, fmt.Sprintf(
		"%s?token=%s",
		"api/v1/tags",
		//token,
		name,
	))
	if err != nil {
		return nil, err
	}

	return body, nil
}
