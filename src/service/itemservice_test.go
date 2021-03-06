package main

import "log"

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

/*var (
	setOfItems = []Item{

		{
			Id: "1",
			Description: "first product",
			Price: 2.0,
		},
		{
			Id: "2",
			Description: "second product",
			Price: 34.0,
		},
		{
			Id: "3",
			Description: "third product",
			Price: 332.0,
		},
		{
			Id: "4",
			Description: "forth product",
			Price: 22.0,
		},
	}

	postItems = ItemContainer{Items:setOfItems}
)

//Testing service to check GET /catalog/product/{id}
func Test_GET_item_returns_404_when_it_does_not_exist(t *testing.T){

	testingServer := getServer(NewItemService(NewMemDb()))
	defer testingServer.Close()

	itemToBeAdded := createItemDto()

	//GETting URL
	url := getURLToBeTested(testingServer.URL, itemToBeAdded.Id);

	res, err := httpGet("lagigliaivan@gmail.com", url)
	if !isHTTPStatus(http.StatusNotFound, res, err){
		deb("GET", url, res.StatusCode, http.StatusOK)
		t.FailNow()
	}
}

func Test_GET_item_returns_200_when_it_exists(t *testing.T){

	itemToBeAdded := createItemDto()
	service := NewItemService(NewMemDb())

	//Adding ITEM without calling RESTapi. Calling a service function directly
	service.addUpdateItem(itemToBeAdded)

	router := NewRouter();
	service.ConfigureRouter(router)

	testingServer := httptest.NewServer(router)
	defer testingServer.Close()

	//GETting URL
	url := getURLToBeTested(testingServer.URL, itemToBeAdded.Id);
	log.Printf("url:%s\n", url)
	res, err := httpGet("lagigliaivan@gmail.com", url)
	if !isHTTPStatus(http.StatusOK, res, err){
		deb("GET", url, res.StatusCode, http.StatusOK)
		t.FailNow()
	}
}

//Testing server to check POST /catalog/product/{id}
func Test_POST_item_returns_201_when_it_is_successfully_created (t *testing.T) {

	testingServer := getServer(NewItemService(NewMemDb()))

	defer testingServer.Close()

	itemToBeAdded := createItemDto()
	items := NewItemContainer()
	items.Add(itemToBeAdded)

	url := getURLToBeTested(testingServer.URL);
	res, err := httpPost("121aseda2123123", strings.TrimSuffix(url, "/"), items)

	if !isHTTPStatus(http.StatusCreated, res, err){
		deb(http.MethodPost, url, res.StatusCode, http.StatusCreated)
		t.FailNow()
	}
}

func Test_POST_GET_returns_the_same_item_after_it_is_created(t *testing.T){

	itemToBeAdded := createItemDto()

	server := getServer(NewItemService(NewMemDb()))
	defer server.Close()

	//POST Item
	url := getURLToBeTested(server.URL);


	items := NewItemContainer()
	items.Add(itemToBeAdded)
	res, err := httpPost("abafadfaf9a9fa0fa", strings.TrimSuffix(url, "/"), items)

	if !isHTTPStatus(http.StatusCreated, res, err){
		deb(http.MethodPost, url, res.StatusCode, http.StatusCreated)
		t.FailNow()
	}

	//GET Item

	url = getURLToBeTested(server.URL, itemToBeAdded.Id);

	res, err = httpGet("abafadfaf9a9fa0fa", url)

	if !isHTTPStatus(http.StatusOK, res, err){
		deb(http.MethodGet, url, res.StatusCode, http.StatusOK)
		t.FailNow()
	}

	if !reflect.DeepEqual(itemToBeAdded, createItemFromJson(res.Body)) {
		log.Printf("Error when GETting item to contrast it with the saved one")
		t.FailNow()
	}

}

func Test_PUT_item_returns_200_when_it_is_successfully_updated (t *testing.T) {

	//POST Item
	itemToBeAdded := createItemDto()

	testingServer := getServer(NewItemService(NewMemDb()))
	defer testingServer.Close()
	url := getURLToBeTested(testingServer.URL);

	items := NewItemContainer()
	items.Add(itemToBeAdded)

	res, err := httpPost("lagigliaiv@gmail.com.ar", strings.TrimSuffix(url, "/"), items)

	if !isHTTPStatus(http.StatusCreated, res, err){
		deb(http.MethodPut, url, res.StatusCode, http.StatusCreated)
		t.FailNow()
	}

	//PUT Item
	url = getURLToBeTested(testingServer.URL, itemToBeAdded.Id);

	itemToBeAdded.Description = "Description updated"
	itemToBeAdded.Price = float32(21)

	res, err = httpPut("lagigliaiv@gmail.com.ar", url, itemToBeAdded)

	if !isHTTPStatus(http.StatusOK, res, err){
		deb(http.MethodPut, url, res.StatusCode, http.StatusOK)
		t.FailNow()
	}

	//GET Item
	res, err = httpGet("lagigliaiv@gmail.com.ar", url)

	if !isHTTPStatus(http.StatusOK, res, err){
		deb(http.MethodGet, url, res.StatusCode, http.StatusOK)
		t.FailNow()
	}

	if !reflect.DeepEqual(itemToBeAdded, createItemFromJson(res.Body)) {
		log.Printf("Error when GETting item to contrast it with the saved one")
		t.FailNow()
	}
}

func Test_POST_item_returns_400_when_body_is_sent_without_item_id(t *testing.T){

	itemToBeAdded := createItemDto()
	itemToBeAdded.Id = ""

	server := getServer(NewItemService(NewMemDb()))
	defer server.Close()

	//POST Item
	url := getURLToBeTested(server.URL);

	items := NewItemContainer()
	items.Add(itemToBeAdded)

	res, err := httpPost("lagigliaiv@gmail.com.ar", strings.TrimSuffix(url, "/"), items)

	if !isHTTPStatus(http.StatusBadRequest, res, err){
		deb(http.MethodPost, url, res.StatusCode, http.StatusBadRequest)
		t.FailNow()
	}
}

func Test_GET_items_returns_a_list_of_items(t *testing.T){

	server := getServer(NewItemService(NewMemDb()))
	defer server.Close()

	if httpPOST("lagigliaiv@gmail.com.ar", *server) != nil {
		t.FailNow();
	}

	url := getURLToBeTested(server.URL);

	defer server.Close()
	log.Printf("server.URL:%s", url)

	res, err := httpGet("lagigliaiv@gmail.com.ar", url)

	if err != nil {
		log.Printf("ERROR")
		t.FailNow()
	}

	items := NewItemContainer()

	body, err := ioutil.ReadAll(res.Body)

	if err := json.Unmarshal(body, &items); err != nil {

		log.Printf("Error when reading response %s", err)
		t.FailNow()
	}

	if len(items.GetItems()) != len(setOfItems){
		log.Printf("Error: Expected items quantity is different from the received one")
		t.FailNow()

	}

	itemsFound := 0

	for _, i := range items.GetItems() {

		for _, x := range setOfItems {
			if x.Id == i.Id && reflect.DeepEqual(x, i){
				itemsFound++
				break
			}
		}
	}

	if itemsFound != len(setOfItems) {
		log.Printf("Error: Some items are missing")
		t.FailNow()
	}

}

//Testing functions that are not exposed as REST services.
func Test_returns_an_error_when_item_does_NOT_exist (t *testing.T) {

	service := NewItemService(NewMemDb());
	item := service.getItem("1021")

	if item.Id == "1021" {
		t.FailNow()
	}
}

func Test_returns_an_item_just_saved (t *testing.T) {

	itemToBeAdded := createItemDto()

	service := NewItemService(NewMemDb());
	service.addUpdateItem(itemToBeAdded)

	item := service.getItem(itemToBeAdded.Id)

	if item.Id != itemToBeAdded.Id || item.Description != itemToBeAdded.Description || item.Price != itemToBeAdded.Price {
		t.FailNow()
	}
}

func Test_returns_an_empty_item_if_it_does_not_exist(t *testing.T){

	service := NewItemService(NewMemDb());
	item := service.getItem("non_existing_item")

	if item.Id != "" {
		t.FailNow()
	}
}

func Test_returns_no_error_when_adding_an_item(t *testing.T){

	service := NewItemService(NewMemDb());
	item := createItemDto()
	err := service.addUpdateItem(item)

	if err != 0 {
		t.FailNow()
	}
}

//Tests auxiliary functions

func deb(method string, url string, expectedStatusCode int, receivedStatusCode int){

	var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Lshortfile)
	logger.Printf("%s URL: %s StatusCode %d different from what was expected %d", method, url, expectedStatusCode, receivedStatusCode)
	fmt.Print(&buf)
}

func createItemDto() Item {

	id := "12345"
	price := float32(10.1)
	descr := "milk 100 cm3"
	category := "Mercaderia"

	return Item{id, descr, price, category }
}
*/
/*
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


func httpPOST(user string, server httptest.Server) error{

	//POST Items for later being retrieved.
	url := getURLToBeTested(server.URL);

	res, err := httpPost(user, strings.TrimSuffix(url, "/"), postItems)

	if !isHTTPStatus(http.StatusCreated, res, err){
		deb(http.MethodPost, url, res.StatusCode, http.StatusCreated)
		return err
	}

	return nil
}

func httpPut(user, url string, item Stringifiable) (resp * http.Response, err error) {

	bodyAsString := item.ToJsonString()
	log.Printf("body: %s", bodyAsString)
	body := strings.NewReader(bodyAsString)

	req, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		log.Printf("Error when creating PUT request %d.", err)
		return nil, err
	}
	req.Header.Add(HEADER, user)
	resp, err = http.DefaultClient.Do(req)
	return resp, err
}
func createItemFromJson(itemAsJson io.ReadCloser) Item {

	item := new(Item)
	response, err := ioutil.ReadAll(itemAsJson)

	if err != nil {
		log.Printf("Error when reading Json from response")
	}

	if err := json.Unmarshal(response, item); err != nil {
		log.Printf("Error when unmarshaling Json to item.Item")
	}

	return *item
}
*/
