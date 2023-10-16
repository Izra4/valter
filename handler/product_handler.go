package handler

import (
	"Valter/service"
	"Valter/utility"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{productService}
}

func (ph *ProductHandler) Dummy() {
	_, err := ph.productService.Dummy()
	if err != nil {
		return
	}
	_, err = ph.productService.Dummy2()
	if err != nil {
		return
	}
}

type allProducts struct {
	ID          uint32    `json:"id"`
	Createdat   time.Time `json:"createdat"`
	Updatedat   time.Time `json:"updatedat"`
	Deletedat   time.Time `json:"deletedat"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
}

func (ph *ProductHandler) GetAllProducts(c *gin.Context) {
	data, err := ph.productService.GetAllProducts(c)
	if err != nil {
		utility.HttpInternalErrorResponse(c, "Failed to get data", err)
		return
	}

	client := utility.SupabaseClient()
	var fixedResult []allProducts

	for _, datas := range data {
		link := client.GetPublicUrl("product_image", datas.Link)
		fixedResult = append(fixedResult, allProducts{
			ID:          datas.ID,
			Createdat:   datas.Createdat.Time,
			Updatedat:   datas.Updatedat.Time,
			Deletedat:   datas.Deletedat.Time,
			Name:        datas.Name,
			Description: datas.Description,
			Link:        link.SignedURL,
		})
	}

	utility.HttpSuccessResponse(c, "Success to get data", fixedResult)
}

type features struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type productById struct {
	ID          uint32    `json:"id"`
	Createdat   time.Time `json:"createdat"`
	Updatedat   time.Time `json:"updatedat"`
	Deletedat   time.Time `json:"deletedat"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	BookPict    string    `json:"book_pict"`
	Feat        []features
}

func (ph *ProductHandler) GetProductById(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	data, err := ph.productService.GetProductsById(c, uint32(id))
	if err != nil {
		utility.HttpInternalErrorResponse(c, "Failed to get data", err)
		return
	}
	client := utility.SupabaseClient()
	link := client.GetPublicUrl("product_image", data.Link)
	fmt.Println("=========================================")
	log.Println(link)
	fmt.Println("=========================================")
	bookLink := client.GetPublicUrl("product_image", data.BookPict)
	feats, err := ph.productService.GetFeatures(data.ID)
	if err != nil {
		utility.HttpInternalErrorResponse(c, "Failed to get features", err)
		return
	}

	var fixedFeats []features
	for _, datas := range feats {
		fixedFeats = append(fixedFeats, features{
			Title: datas.Title,
			Desc:  datas.Description,
		})
	}

	fixedResult := productById{
		ID:          data.ID,
		Createdat:   data.Createdat.Time,
		Updatedat:   data.Updatedat.Time,
		Deletedat:   data.Deletedat.Time,
		Name:        data.Name,
		Description: data.Description,
		Link:        link.SignedURL,
		BookPict:    bookLink.SignedURL,
		Feat:        fixedFeats,
	}
	utility.HttpSuccessResponse(c, "Success to get data", fixedResult)
}
