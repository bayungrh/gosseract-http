package handlers

import (
	"bayungrh/gosseract-http/lib"
	"bayungrh/gosseract-http/types"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Request struct {
  Url string
  Parse []types.ParseBox
}

func ParseImageUrl(c *gin.Context) {
  var requestParam Request
  c.BindJSON(&requestParam);

  fmt.Println(requestParam.Url)

  httpClient := http.Client{Timeout: 2 * time.Second}
  reqImg, err := httpClient.Get(requestParam.Url)

  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }

  contentType := reqImg.Header.Get("Content-Type")
  imgByte, _ := ioutil.ReadAll(reqImg.Body)
  reqImg.Body.Close()

  if contentType != "image/png" && contentType != "image/jpg" && contentType != "image/jpeg" {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": "Invalid image",
    })
  }

  data := types.Parse{Data: requestParam.Parse}

  var jsonRes []interface{}
  for _, v := range data.Data {
    p, _ := lib.ParseText(imgByte, v.Rect)
    jsonRes = append(jsonRes, types.Response{Field: v.Name, Value: p})
  }

  c.JSON(http.StatusOK, gin.H{
    "data": jsonRes,
  })
}
