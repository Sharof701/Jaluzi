package handler

import (
	"jaluzi/models"
	"jaluzi/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Product godoc
// @ID create_product
// @Router /jaluzi/api/v1/product [POST]
// @Summary Create Product
// @Description Create Product
// @Tags Product
// @Accept json
// @Product json
// @Param Product body models.ProductCreate true "CreateProductRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateProduct(c *gin.Context) {
	var (
		productCreate models.ProductCreate
	)

	err := c.ShouldBindJSON(&productCreate)
	if err != nil {
		h.logger.Error(err.Error() + "  :  " + "error Product Should Bind Json!")
		c.JSON(http.StatusBadRequest, "Please, Enter Valid Data!")
		return
	}

	resp, err := h.storage.Product().Create(c.Request.Context(), &productCreate)
	if err != nil {
		h.logger.Error(err.Error() + "  :  " + "error Product.Create")
		c.JSON(http.StatusInternalServerError, "Server Error!")
		return
	}

	h.logger.Info("Create Product Successfully!!")
	c.JSON(http.StatusCreated, resp)
}

// GetByID Product godoc
// @ID get_by_id_product
// @Router /jaluzi/api/v1/product/{id} [GET]
// @Summary Get By ID Product
// @Description Get By ID Product
// @Tags Product
// @Accept json
// @Product json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdProduct(c *gin.Context) {
	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.logger.Error("is valid uuid!")
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	request, err := h.storage.Product().GetByID(c.Request.Context(), &models.ProductPrimaryKey{Id: id})
	if err != nil {
		h.logger.Error(err.Error() + "  :  " + "storage.Product.GetByID!")
		c.JSON(http.StatusInternalServerError, "Server Error!")
		return
	}

	h.logger.Info("GetByID Product Response!")
	c.JSON(http.StatusOK, request)
}

// GetList Product godoc
// @ID get_list_product
// @Router /jaluzi/api/v1/product [GET]
// @Summary Get List Product
// @Description Get List Product
// @Tags Product
// @Accept json
// @Product json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListProduct(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.logger.Error(err.Error() + "  :  " + "GetListProduct INVALID OFFSET!")
		c.JSON(http.StatusBadRequest, "INVALID OFFSET")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.logger.Error(err.Error() + "  :  " + "GetListProduct INVALID LIMIT!")
		c.JSON(http.StatusBadRequest, "INVALID LIMIT")
		return
	}

	resp, err := h.storage.Product().GetList(c.Request.Context(), &models.ProductGetListRequest{
		Offset: offset,
		Limit:  limit,
	})

	if err != nil && err.Error() != "no rows in result set" {
		h.logger.Error(err.Error() + "  :  " + "storage.Product.GetList!")
		c.JSON(http.StatusInternalServerError, "Server Error!")
		return
	}

	h.logger.Info("GetListProduct Response!")
	c.JSON(http.StatusOK, resp)
}

// Delete Product godoc
// @ID delete_product
// @Router /jaluzi/api/v1/product/{id} [DELETE]
// @Summary Delete Product
// @Description Delete Product
// @Tags Product
// @Accept json
// @Product json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteProduct(c *gin.Context) {
	var id = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.logger.Error("is not valid uuid!")
		c.JSON(http.StatusBadRequest, "invalid id!")
		return
	}

	err := h.storage.Product().Delete(c.Request.Context(), &models.ProductPrimaryKey{Id: id})
	if err != nil {
		h.logger.Error(err.Error() + "  :  " + "storage.Product.Delete!")
		c.JSON(http.StatusInternalServerError, "Unable to delete data, please try again later!")
		return
	}

	h.logger.Info("Product Deleted Successfully!")
	c.JSON(http.StatusNoContent, nil)
}

// Update Product godoc
// @ID update_product
// @Router /jaluzi/api/v1/product/{id} [PUT]
// @Summary Update Product
// @Description Update Product
// @Tags Product
// @Accept json
// @Product json
// @Param id path string true "id"
// @Param Product body models.ProductUpdate true "UpdateProductRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateProduct(c *gin.Context) {
	var (
		id            = c.Param("id")
		productUpdate models.ProductUpdate
	)

	if !helper.IsValidUUID(id) {
		h.logger.Error("is invalid uuid!")
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&productUpdate)
	if err != nil {
		h.logger.Error(err.Error() + "  :  " + "error Product Should Bind Json!")
		c.JSON(http.StatusBadRequest, "Please, Enter Valid Data!")
		return
	}

	productUpdate.Id = id
	rowsAffected, err := h.storage.Product().Update(c.Request.Context(), &productUpdate)
	if err != nil {
		h.logger.Error(err.Error() + "  :  " + "storage.Product.Update!")
		c.JSON(http.StatusInternalServerError, "Server Error!")
		return
	}

	if rowsAffected <= 0 {
		h.logger.Error("storage.Product.Update!")
		c.JSON(http.StatusBadRequest, "Unable to update data. Please try again later!")
		return
	}

	resp, err := h.storage.Product().GetByID(c.Request.Context(), &models.ProductPrimaryKey{Id: productUpdate.Id})
	if err != nil {
		h.logger.Error(err.Error() + "  :  " + "storage.Product.GetByID!")
		c.JSON(http.StatusInternalServerError, "Server Error!")
		return
	}

	h.logger.Info("Update Product Successfully!")
	c.JSON(http.StatusAccepted, resp)
}
