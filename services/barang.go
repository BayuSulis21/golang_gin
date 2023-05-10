package services

import (
	"fmt"
	"golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// narik barang menggunakan param
var DbConn *gorm.DB

func SearchBarangHandler(c *gin.Context) {

	kategori := c.Param("kategori")
	//supplier := c.Param("supplier")

	var stringQuery string = fmt.Sprintf(`SELECT * FROM barang WHERE kategori like %s`,
		"'%"+kategori+"%'")
	var Listbarang []models.Barang

	respDB := DbConn.Raw(stringQuery).Scan(&Listbarang)
	fmt.Println(respDB.Error)

	c.JSON(http.StatusOK, models.SearchBarangJsonResponse{
		Code:     "200",
		Message:  "success",
		Kategori: kategori,
		ListData: Listbarang,
	})

}

//narik barang menggunakan query

func SearchBarangQueryHandler(c *gin.Context) {
	// kategori := c.Query("kategori")
	// supplier := c.Query("supplier")

	// DataBarang := map[string]interface{}{
	// 	"nama":  "gelas",
	// 	"stock": "10",
	// }
	// c.JSON(http.StatusOK, models.SearchBarangJsonResponse{
	// 	Code:     "200",
	// 	Message:  "success",
	// 	Kategori: kategori,
	// 	Supplier: supplier,
	// 	Barang:   DataBarang,
	// })
}

//narik barang menggunakann json

func SearchBarangJsonHandler(c *gin.Context) {
	var paramContoh models.ParamBody
	c.ShouldBindJSON(&paramContoh)

	// DataBarang := map[string]interface{}{
	// 	"nama":  "gelas",
	// 	"stock": "10",
	// }
	// c.JSON(http.StatusOK, models.SearchBarangJsonResponse{
	// 	Code:     "200",
	// 	Message:  "success",
	// 	Kategori: paramContoh.Kategori,
	// 	ListData: Listbarang,
	// })
}

// narik barang menggunakan header

func SearchBarangHeaderHandler(c *gin.Context) {
	// kategori := c.Request.Header.Get("kategori")
	// supplier := c.Request.Header.Get("supplier")

	// DataBarang := map[string]interface{}{
	// 	"nama":  "gelas",
	// 	"stock": "10",
	// }
	// c.JSON(http.StatusOK, models.SearchBarangJsonResponse{
	// 	Code:     "200",
	// 	Message:  "success",
	// 	Kategori: kategori,
	// 	ListData: Listbarang,
	// })
}
