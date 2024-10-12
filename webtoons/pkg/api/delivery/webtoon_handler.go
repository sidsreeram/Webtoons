package delivery

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/webtoons/pkg/domain"
    usecase "github.com/webtoons/pkg/usecase/interfaces"
)

type WebtoonHandler struct {
    WebtoonUC usecase.WebtoonUseCase
}

func NewWebtoonHandler(webtoonUC usecase.WebtoonUseCase) *WebtoonHandler {
    return &WebtoonHandler{
        WebtoonUC: webtoonUC,
    }
}

// Define the handler methods

func (h *WebtoonHandler) GetAll(c *gin.Context) {
    web, err := h.WebtoonUC.GetAllWebtoons()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, web)
}

func (h *WebtoonHandler) Create(c *gin.Context) {
    var webtoon domain.Webtoon
    if err := c.ShouldBindJSON(&webtoon); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := h.WebtoonUC.AddWebtoon(webtoon)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, webtoon)
}

func (h *WebtoonHandler) GetByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    webtoon, err := h.WebtoonUC.GetWebtoonByID(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, webtoon)
}

func (h *WebtoonHandler) Delete(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    err = h.WebtoonUC.DeleteWebtoon(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}
