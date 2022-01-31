package api

import (
	"Products/pkg/data"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
)

type productAPI struct {
	data *data.ProductData
}

type ProductApi interface {
	ShowParametersWithSpecificGroup(writer http.ResponseWriter, request *http.Request)
	ShowParametersWithoutSpecificGroup(writer http.ResponseWriter, request *http.Request)
	ShowDataAboutProductsWithSpecificGroup(writer http.ResponseWriter, request *http.Request)
	ShowDataAboutProductsAndAllParametersWithValue(writer http.ResponseWriter, request *http.Request)
	DeleteProductWithSpecificParameters(writer http.ResponseWriter, request *http.Request)
	MoveGroupOfParametersFromOneGroupToAnother(writer http.ResponseWriter, request *http.Request)
}

func InitConnectionToServer(r *mux.Router, data data.ProductData) {
	api := &productAPI{data: &data}

	r.HandleFunc("/", start)
	r.HandleFunc("/input", api.OpenInputPage).Methods("GET")

	{
		r.HandleFunc("/products_with_parameters", api.ShowParametersWithSpecificGroup).Methods("GET")
		r.HandleFunc("/products_without_parameters", api.ShowParametersWithoutSpecificGroup).Methods("GET")
		r.HandleFunc("/products_with_specific_group", api.ShowDataAboutProductsWithSpecificGroup).Methods("GET")
		r.HandleFunc("/products_all", api.ShowDataAboutProductsAndAllParametersWithValue).Methods("GET")
		r.HandleFunc("/products_delete", api.DeleteProductWithSpecificParameters).Methods("GET")
		r.HandleFunc("/products_put", api.MoveGroupOfParametersFromOneGroupToAnother).Methods("GET")
	}
}

func start(writer http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFiles("C:\\Users\\Dzmitry_Putskou\\go\\src\\Products\\pkg\\api\\templates\\main_page.html")
	if err != nil {
		logrus.Error(err)
	}
	tmpl.Execute(writer, "")
}

type inputPage struct {
	URL string
	Action string
	Product []data.Products
}

type mainPage struct {
	ProductWithParameters []data.Parameters
	ProductWithoutParameters []data.Products
	ProductWithParameterGroup []data.Products
	AllProduct []data.Products
	Err string
}

func (p productAPI) OpenInputPage(writer http.ResponseWriter, request *http.Request) {
	action := request.FormValue("submit")
	var actionResponse inputPage
	switch action {
		case "1":
			actionResponse = inputPage{
				URL:    "/products_with_parameters/",
				Action: "1",
			}
		case "2":
			actionResponse = inputPage{
				URL:    "/products_without_parameters/",
				Action: "2",
			}
		case "3":
			actionResponse = inputPage{
				URL:    "/products_with_specific_group/",
				Action: "3",
			}
		case "4":
			actionResponse = inputPage{
				URL:    "/products_all/",
				Action: "4",
			}
		case "5":
			actionResponse = inputPage{
				URL:    "/products_delete/",
				Action: "5",
			}
		case "6":
			product, err := p.data.ShowGroupOfParamAndGroupOfProduct()
			if err != nil {
				logrus.Error("Don't have opportunity to show group of param and group of product")
			}
			actionResponse = inputPage{
				URL:    "/products_put/",
				Action: "6",
				Product: product,
			}
	}
	tmpl, err := template.ParseFiles("C:\\Users\\Dzmitry_Putskou\\go\\src\\Products\\pkg\\api\\templates\\input_page.html")
	if err != nil {
		logrus.Error(err)
	}
	tmpl.Execute(writer, actionResponse)
}

func (p productAPI) ShowParametersWithSpecificGroup(writer http.ResponseWriter, request *http.Request) {
	var mainPageValue mainPage
	action := request.FormValue("title")
	tmpl, err := template.ParseFiles("C:\\Users\\Dzmitry_Putskou\\go\\src\\Products\\pkg\\api\\templates\\main_page.html")
	if err != nil {
		logrus.Error(err)
	}
	parameters, err := p.data.ShowParametersWithSpecificGroup(action)
	if err != nil {
		tmpl.Execute(writer, mainPage{Err: SendErrorFromAPI(err, action)})
		return
	}
	mainPageValue = mainPage{ProductWithParameters: parameters}
	tmpl.Execute(writer, mainPageValue)
}

func (p productAPI) ShowParametersWithoutSpecificGroup(writer http.ResponseWriter, request *http.Request) {
	var mainPageValue mainPage
	action := request.FormValue("title")
	tmpl, err := template.ParseFiles("C:\\Users\\Dzmitry_Putskou\\go\\src\\Products\\pkg\\api\\templates\\main_page.html")
	if err != nil {
		logrus.Error(err)
	}
	product, err := p.data.ShowParametersWithoutSpecificGroup(action)
	if err != nil {
		tmpl.Execute(writer, mainPage{Err: SendErrorFromAPI(err, action)})
		return
	}
	mainPageValue = mainPage{ProductWithoutParameters: product}
	tmpl.Execute(writer, mainPageValue)
}

func (p productAPI) ShowDataAboutProductsWithSpecificGroup(writer http.ResponseWriter, request *http.Request) {
	var mainPageValue mainPage
	action := request.FormValue("title")
	tmpl, err := template.ParseFiles("C:\\Users\\Dzmitry_Putskou\\go\\src\\Products\\pkg\\api\\templates\\main_page.html")
	if err != nil {
		logrus.Error(err)
	}
	product, err := p.data.ShowProductWithSpecificProductGroups(action)
	if err != nil {
		tmpl.Execute(writer, mainPage{Err: SendErrorFromAPI(err, action)})
		return
	}
	mainPageValue = mainPage{ProductWithParameterGroup: product}
	tmpl.Execute(writer, mainPageValue)
}

func (p productAPI) ShowDataAboutProductsAndAllParametersWithValue(writer http.ResponseWriter, request *http.Request) {
	var mainPageValue mainPage
	tmpl, err := template.ParseFiles("C:\\Users\\Dzmitry_Putskou\\go\\src\\Products\\pkg\\api\\templates\\main_page.html")
	if err != nil {
		logrus.Error(err)
	}
	product, err := p.data.ShowProduct()
	if err != nil {
		tmpl.Execute(writer, mainPage{Err: "got an error when tried to get parameters"})
		return
	}
	mainPageValue = mainPage{AllProduct: product}
	tmpl.Execute(writer, mainPageValue)
}

func (p productAPI) DeleteProductWithSpecificParameters(writer http.ResponseWriter, request *http.Request) {
	var mainPageValue mainPage
	action := request.FormValue("title")
	tmpl, err := template.ParseFiles("C:\\Users\\Dzmitry_Putskou\\go\\src\\Products\\pkg\\api\\templates\\main_page.html")
	if err != nil {
		logrus.Error(err)
	}
	err = p.data.DeleteDataWithSpecialParameters(action)
	if err != nil {
		tmpl.Execute(writer, mainPage{Err: SendErrorFromAPI(err, action)})
		return
	}
	tmpl.Execute(writer, mainPageValue)
}

func (p productAPI) MoveGroupOfParametersFromOneGroupToAnother(writer http.ResponseWriter, request *http.Request) {
	var mainPageValue mainPage
	tmpl, err := template.ParseFiles("C:\\Users\\Dzmitry_Putskou\\go\\src\\Products\\pkg\\api\\templates\\main_page.html")
	if err != nil {
		logrus.Error(err)
	}
	paramOfGroupParam := request.FormValue("paramOfGroupParam")
	paramOfGroupProduct2 := request.FormValue("paramOfGroupProduct2")
	err = p.data.ChangeDataIntoDb(paramOfGroupParam, paramOfGroupProduct2)
	if err != nil {
		tmpl.Execute(writer, mainPage{Err: SendErrorFromAPI(err, fmt.Sprint(paramOfGroupParam, paramOfGroupProduct2))})
		return
	}
	tmpl.Execute(writer, mainPageValue)
}