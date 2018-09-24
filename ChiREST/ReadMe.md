# Demo REST Microservice using Go-Chi

3rd parties libraries used:
1. [Go Chi](https://github.com/go-chi/chi)
2. [sJSON](https://github.com/tidwall/sjson)

Thúy em viết cho anh một go micro service trả về danh sách các Article.

Xem kiểu Articles tại [models/article.go](models/article.go)

Dữ liệu mẫu ở đây:

```go
var articles = []*models.Article{
	{	ID:     "1",
		UserID: 1,
		Title:  "Multi-layered 5th generation workforce",
		Slug:   "radical" },
	{	ID:     "2",
		UserID: 6,
		Title:  "Future-proofed 6th generation secured line",
		Slug:   "systematic"},
	{
		ID:     "3",
		UserID: 4,
		Title:  "Automated bandwidth-monitored matrix",
		Slug:   "Implemented"},
	{
		ID:     "4",
		UserID: 3,
		Title:  "Balanced human-resource help-desk",
		Slug:   "frame"},
	{
		ID:     "5",
		UserID: 7,
		Title:  "Secured leading edge circuit",
		Slug:   "systemic"},
	{
		ID:     "6",
		UserID: 2,
		Title:  "Public-key dynamic matrix",
		Slug:   "tertiary"},
	{
		ID:     "7",
		UserID: 8,
		Title:  "Triple-buffered attitude-oriented access",
		Slug:   "high-level"},
	{
		ID:     "8",
		UserID: 2,
		Title:  "Up-sized content-based installation",
		Slug:   "Polarised"},
	{
		ID:     "9",
		UserID: 8,
		Title:  "Monitored even-keeled infrastructure",
		Slug:   "asynchronous"},
	{
		ID:     "10",
		UserID: 9,
		Title:  "Open-source object-oriented application",
		Slug:   "contextually-based"},
}
```