package types

type Rect []int

type ParseBox struct {
  Name string
  Rect Rect
}
type Parse struct {
  Data []ParseBox
}

type Response struct {
  Field string `json:"field"`
  Value string `json:"value"`
}
