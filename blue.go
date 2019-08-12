package goblueoc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// API struct
type API struct {
	baseurl  string
	username string
	password string
	client   *http.Client
}

type param struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

// NewAPI func
func NewAPI(url string, username string, password string) *API {
	return &API{url, username, password, &http.Client{}}
}

// HTTPDo func
func (api *API) HTTPDo(url string, method string) (body []byte, err error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return
	}
	req.SetBasicAuth(api.username, api.password)
	response, err := api.client.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	return

}

// getParams func
func getParams(d map[string]interface{}) (res []byte, err error) {
	/*
		var res string
		var resArr []string
		for k, v := range d {
			resArr = append(resArr, fmt.Sprintf("%s=%s", k, v))
		}
		res = strings.Join(resArr, "&")
		return res
	*/
	var params map[string][]param
	params = make(map[string][]param)
	var resArr []param
	for k, v := range d {
		resArr = append(resArr, param{Name: k, Value: v})
	}
	params["parameters"] = resArr
	res, err = json.Marshal(params)
	return

}

// HTTPDoHeader func
func (api *API) HTTPDoHeader(url string, method string, header map[string]string, data []byte) (body []byte, err error) {
	/*
		data, err := getParams(d)
		if err != nil {
			return
		}
	*/
	//fmt.Println(bytes.NewBuffer(data))
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return
	}
	req.SetBasicAuth(api.username, api.password)

	for k, v := range header {
		req.Header.Add(k, v)
	}
	response, err := api.client.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	return

}

// ListPipelines func
func (api *API) ListPipelines() (data []Pipeline, err error) {
	url := api.baseurl + "/blue/rest/organizations/jenkins/pipelines/"
	body, err := api.HTTPDo(url, "GET")
	if err != nil {
		return
	}
	json.Unmarshal(body, &data)
	return
}

// GetPipeline func
func (api *API) GetPipeline(p string) (data Pipeline, err error) {
	url := fmt.Sprintf("%s/blue/rest/organizations/jenkins/pipelines/%s/", api.baseurl, p)
	body, err := api.HTTPDo(url, "GET")
	if err != nil {
		return
	}
	json.Unmarshal(body, &data)
	return
}

// ListPipelineRuns func
func (api *API) ListPipelineRuns(p string) (data []PipelineRun, err error) {
	url := fmt.Sprintf("%s/blue/rest/organizations/jenkins/pipelines/%s/runs/", api.baseurl, p)
	body, err := api.HTTPDo(url, "GET")
	if err != nil {
		return
	}
	json.Unmarshal(body, &data)
	return
}

// GetPipelineRun func
func (api *API) GetPipelineRun(p string, r string) (data PipelineRun, err error) {
	url := fmt.Sprintf("%s/blue/rest/organizations/jenkins/pipelines/%s/runs/%s/", api.baseurl, p, r)
	body, err := api.HTTPDo(url, "GET")
	if err != nil {
		return
	}
	json.Unmarshal(body, &data)
	return
}

// ListRunNodes func
func (api *API) ListRunNodes(p string, r string) (data []Node, err error) {
	url := fmt.Sprintf("%s/blue/rest/organizations/jenkins/pipelines/%s/runs/%s/nodes/", api.baseurl, p, r)
	body, err := api.HTTPDo(url, "GET")
	if err != nil {
		return
	}
	json.Unmarshal(body, &data)
	return

}

// ListNodeSteps func
func (api *API) ListNodeSteps(p string, r string, n string) (data []Step, err error) {
	url := fmt.Sprintf("%s/blue/rest/organizations/jenkins/pipelines/%s/runs/%s/nodes/%s/steps/", api.baseurl, p, r, n)
	body, err := api.HTTPDo(url, "GET")
	if err != nil {
		return
	}
	json.Unmarshal(body, &data)
	return
}

// GetStepLog func
func (api *API) GetStepLog(p string, r string, n string, s string) (data string, err error) {
	url := fmt.Sprintf("%s/blue/rest/organizations/jenkins/pipelines/%s/runs/%s/nodes/%s/steps/%s/log/", api.baseurl, p, r, n, s)
	body, err := api.HTTPDo(url, "GET")
	if err != nil {
		return
	}
	data = string(body)
	return
}

// GetRunLog func
func (api *API) GetRunLog(p string, r string) (data string, err error) {
	url := fmt.Sprintf("%s/blue/rest/organizations/jenkins/pipelines/%s/runs/%s/log/", api.baseurl, p, r)
	body, err := api.HTTPDo(url, "GET")
	if err != nil {
		return
	}
	data = string(body)
	return
}

// Run func
// Start a build
func (api *API) Run(p string, d map[string]interface{}) error {
	url := fmt.Sprintf("%s/blue/rest/organizations/jenkins/pipelines/%s/runs/", api.baseurl, p)
	//body, err := api.HTTPDo(url, "POST")
	header := map[string]string{
		"Content-Type": "application/json",
	}
	data, err := getParams(d)
	if err != nil {
		return err
	}
	_, err = api.HTTPDoHeader(url, "POST", header, data)
	if err != nil {
		return err
	}
	//data := string(body)
	//fmt.Println(data)
	return err

}

// Stop func
// Stop a build
func (api *API) Stop(p string, r string) error {
	url := fmt.Sprintf("%s/blue/rest/organizations/jenkins/pipelines/%s/runs/%s/stop", api.baseurl, p, r)
	header := map[string]string{
		"Content-Type": "application/json",
	}
	_, err := api.HTTPDoHeader(url, "PUT", header, nil)
	if err != nil {
		return err
	}
	return err

}

// ListFavorites func
func (api *API) ListFavorites(user string) (data []Favorite, err error) {
	url := api.baseurl + "/blue/rest/users/" + user + "/favorites/"
	body, err := api.HTTPDo(url, "GET")
	if err != nil {
		return
	}
	json.Unmarshal(body, &data)
	return
}

// FavoritePipeline func
func (api *API) FavoritePipeline(p string, fav string) error {
	url := fmt.Sprintf("%s/blue/rest/organizations/jenkins/pipelines/%s/favorite/", api.baseurl, p)
	header := map[string]string{
		"Content-Type": "application/json",
	}
	var d map[string]interface{}
	d = make(map[string]interface{})
	d["favorite"] = fav
	data, err := json.Marshal(d)
	if err != nil {
		return err
	}
	_, err = api.HTTPDoHeader(url, "PUT", header, data)
	if err != nil {
		return err
	}
	return err
}
