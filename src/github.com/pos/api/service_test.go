package main

import (
	"log"
	"testing"
	"github.com/pos/infrastructure"
	"net/http"
	"net/http/httptest"
	"strings"
	"bytes"
	"fmt"
	"github.com/pos/dto"
	"io/ioutil"
	"encoding/json"
//	"reflect"
	"io"
)

//DB MOCK
var db infrastructure.Mem_DB

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

//Testing service to check GET /catalog/productestingServer/{id}
func Test_returns_404_when_item_does_not_exist(t *testing.T){

	service := NewService(infrastructure.NewMemDb())
	testingServer := httptest.NewServer(http.HandlerFunc(service.HandleGetItem))
	defer testingServer.Close()

	itemToBeAdded := createItemDto()

	//GETting URL
	url := getURLToBeTested(testingServer.URL, itemToBeAdded.Id);

	res, err := httpGet(url)
	if !isHTTPStatus(http.StatusNotFound, res, err){
		debug("GET", url, res.StatusCode, http.StatusOK)
		t.FailNow()
	}
}

func Test_returns_200_when_item_exists(t *testing.T){

	itemToBeAdded := createItemDto()

	service := NewService(infrastructure.NewMemDb())

	//Adding ITEM without calling RESTapi. Calling a service function directly
	service.AddUpdateItem(itemToBeAdded)

	service.GetRequestParameters = returnItemIdFromURL(itemToBeAdded.Id)

	testingServer := httptest.NewServer(http.HandlerFunc(service.HandleGetItem))
	defer testingServer.Close()

	//GETting URL
	url := getURLToBeTested(testingServer.URL, itemToBeAdded.Id);

	res, err := httpGet(url)
	if !isHTTPStatus(http.StatusOK, res, err){
		debug("GET", url, res.StatusCode, http.StatusOK)
		t.FailNow()
	}
}

//TEStestingServer to check POST /catalog/productestingServer/{id} resultestingServer
func Test_returns_201_when_item_is_created (t *testing.T) {

	itemToBeAdded := createItemDto()
	service := NewService(infrastructure.NewMemDb())

	testingServer := httptest.NewServer(http.HandlerFunc(service.HandlePostItem))
	defer testingServer.Close()
	url := getURLToBeTested(testingServer.URL);

	res, err := httpPost(url, itemToBeAdded)

	if !isHTTPStatus(http.StatusCreated, res, err){
		debug(http.MethodPost, url, res.StatusCode, http.StatusCreated)
		t.FailNow()
	}
}

func Test_returns_200_when_item_is_updated (t *testing.T) {

	//POST Item
	itemToBeAdded := createItemDto()
	service := NewService(infrastructure.NewMemDb())
	service.GetRequestParameters = returnItemIdFromURL(itemToBeAdded.Id)

	testingServer := httptest.NewServer(http.HandlerFunc(service.HandlePostItem))
	defer testingServer.Close()
	url := getURLToBeTested(testingServer.URL);


	res, err := httpPost(url, itemToBeAdded)

	if !isHTTPStatus(http.StatusCreated, res, err){
		debug(http.MethodPut, url, res.StatusCode, http.StatusCreated)
		t.FailNow()
	}

	//PUT Item
	testingServer = httptest.NewServer(http.HandlerFunc(service.HandlePutItem))
	defer testingServer.Close()


	url = getURLToBeTested(testingServer.URL, itemToBeAdded.Id);

	itemToBeAdded.Desc = "Description updated"
	itemToBeAdded.Price = float32(21)

	res, err = httpPut(url, itemToBeAdded)

	if !isHTTPStatus(http.StatusOK, res, err){
		debug(http.MethodPut, url, res.StatusCode, http.StatusOK)
		t.FailNow()
	}

	//GET Item
	testingServerGET := httptest.NewServer(http.HandlerFunc(service.HandleGetItem))
	defer testingServerGET.Close()

	res, err = httpGet(url)

	if !isHTTPStatus(http.StatusOK, res, err){
		debug(http.MethodGet, url, res.StatusCode, http.StatusOK)
		t.FailNow()
	}

	if !areItemsEquals(itemToBeAdded, createItemFromJson(res.Body)){
		log.Printf("Error when GETting item to contrast it with the saved one")
		t.FailNow()
	}


}
func Test_returns_400_when_post_is_sent_without_itemid(t *testing.T){

	itemToBeAdded := createItemDto()
	itemToBeAdded.Id = ""

	service := NewService(infrastructure.NewMemDb())

	testingServerPUT := httptest.NewServer(http.HandlerFunc(service.HandlePostItem))
	defer testingServerPUT.Close()

	//POST Item
	url := getURLToBeTested(testingServerPUT.URL);

	res, err := httpPost(url, itemToBeAdded)

	if !isHTTPStatus(http.StatusBadRequest, res, err){
		debug(http.MethodPost, url, res.StatusCode, http.StatusBadRequest)
		t.FailNow()
	}
}
func Test_returns_the_same_item_after_it_is_created(t *testing.T){

	itemToBeAdded := createItemDto()

	service := NewService(infrastructure.NewMemDb())
	service.GetRequestParameters = returnItemIdFromURL(itemToBeAdded.Id)

	testingServerPUT := httptest.NewServer(http.HandlerFunc(service.HandlePostItem))
	defer testingServerPUT.Close()

	//POST Item
	url := getURLToBeTested(testingServerPUT.URL);

	res, err := httpPost(url, itemToBeAdded)

	if !isHTTPStatus(http.StatusCreated, res, err){
		debug(http.MethodPost, url, res.StatusCode, http.StatusCreated)
		t.FailNow()
	}

	//GET Item
	testingServerGET := httptest.NewServer(http.HandlerFunc(service.HandleGetItem))
	defer testingServerGET.Close()
	url = getURLToBeTested(testingServerGET.URL, itemToBeAdded.Id);

	res, err = httpGet(url)

	if !isHTTPStatus(http.StatusOK, res, err){
		debug(http.MethodGet, url, res.StatusCode, http.StatusOK)
		t.FailNow()
	}

	if !areItemsEquals(itemToBeAdded, createItemFromJson(res.Body)){
		log.Printf("Error when GETting item to contrast it with the saved one")
		t.FailNow()
	}
}

func Test_returns_an_error_when_item_does_NOT_exist (t *testing.T) {

	service := NewService(infrastructure.NewMemDb());
	item := service.GetItem("1021")

	if item.Id == "1021" {
		t.FailNow()
	}
}

func Test_returns_an_item_just_saved (t *testing.T) {

	itemToBeAdded := createItemDto()

	service := NewService(infrastructure.NewMemDb());
	service.AddUpdateItem(itemToBeAdded)

	item := service.GetItem(itemToBeAdded.Id)

	if item.Id != itemToBeAdded.Id || item.Desc != itemToBeAdded.Desc || item.Price != itemToBeAdded.Price {
		t.FailNow()
	}
}

func Test_returns_an_empty_item_if_it_does_not_exist(t *testing.T){

	service := NewService(infrastructure.NewMemDb());
	item := service.GetItem("non_existing_item")

	if item.Id != "" {
		t.FailNow()
	}
}

func Test_returns_no_error_when_adding_an_item(t *testing.T){

	service := NewService(infrastructure.NewMemDb());
	item := createItemDto()
	err := service.AddUpdateItem(item)

	if err != 0 {
		t.FailNow()
	}
}

//Test auxiliary functions
func debug(method string, url string, expectedStatusCode int, receivedStatusCode int){

	var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Lshortfile)
	logger.Printf("%s URL: %s StatusCode %d different from what was expected %d", method, url, expectedStatusCode, receivedStatusCode)
	fmt.Print(&buf)
}

func createItemDto() dto.Item {

	id := "12345"
	price := float32(10.1)
	descr := "milk 100 cm3"

	return dto.Item{id, descr, price}
}

func getRequestBody(item dto.Item) string {

	body := "{\"id\":\"" + item.Id + "\", \"description\":\"" + item.Desc + "\", \"price\":" + fmt.Sprintf("%.1f", item.Price) + "}"
	return body
}

//This function is a callback to return the item id from URL. When running normal, this ID is taken from the URL but
//this process is resolved by the multiplexer (gorilla.mux)
func returnItemIdFromURL(itemId string) func (r *http.Request) map[string]string {

	if itemId == ""{
		log.Printf("ItemId to be return is empty.")
	}

	f := func (r *http.Request) map[string]string {
		req_vars := make(map[string]string)
		req_vars["id"]= itemId
		log.Printf("Returning itemid:%s from a testing callback.", req_vars["id"])
		return req_vars
	}

	return f
}

//API to be tested
func getURLToBeTested(base_url string, params ... string) string {

	var p string

	for _, v := range params {
		p = v + "/"
	}

	catalog_api := "/catalog/products/"

	p = strings.TrimSuffix(p, "/")

	return base_url + catalog_api + p;
}

func httpPut(url string, item dto.Item) (resp * http.Response, err error) {

	bodyAsString := "{\"description\":\"" + item.Desc + "\", \"price\": 23}"
	log.Printf("body: %s", bodyAsString)
	body := strings.NewReader(bodyAsString)

	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		log.Printf("Error when creating PUT request %d.", err)
		return nil, err
	}

	resp, err = http.DefaultClient.Do(req)
	return resp, err
}

func httpGet(url string) (*http.Response, error){
	return http.Get(url)
}

func httpPost(url string, item dto.Item) (*http.Response, error){

	bodyAsString := getRequestBody(item)
	body := strings.NewReader(bodyAsString)
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		log.Printf("Error when creating POST request %d.", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Length", fmt.Sprintf("%d", len(bodyAsString)))
	resp, err := http.DefaultClient.Do(req)

	return resp, err
}

func isHTTPStatus(httpStatus int, res *http.Response, err error ) bool {
	return !( (err != nil) || (res.StatusCode != httpStatus) )
}

func areItemsEquals(item, item2 dto.Item) bool{

	return !((item.Id != item2.Id) || (item.Desc != item2.Desc) || (item.Price != item2.Price))
}

func createItemFromJson(itemAsJson io.ReadCloser) dto.Item {

	item := new(dto.Item)
	response, err := ioutil.ReadAll(itemAsJson)

	if err != nil {
		log.Printf("Error when reading Json from response")
	}

	if err := json.Unmarshal(response, item); err != nil {
		log.Printf("Error when unmarshaling Json to dto.Item")
	}

	return *item
}