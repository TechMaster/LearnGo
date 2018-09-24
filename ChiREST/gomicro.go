package main

import (
	"context"
	"log"

	proto "github.com/TechMaster/LearnGo/ChiREST/proto/service-blog-generate"
	"github.com/micro/go-micro"
)

type Blog struct{}

func (blog *Blog) GetAllArticle(ctx context.Context, req *proto.EmptyRequest, rsp *proto.GetAllArticleResponse) error {
	articles := []*proto.Article {
		{	Id:     "1",
			UserId: 1,
			Title:  "Multi-layered 5th generation workforce",
			Slug:   "radical" },
		{	Id:     "2",
			UserId: 6,
			Title:  "Future-proofed 6th generation secured line",
			Slug:   "systematic"},
		{
			Id:     "3",
			UserId: 4,
			Title:  "Automated bandwidth-monitored matrix",
			Slug:   "Implemented"},
		{
			Id:     "4",
			UserId: 3,
			Title:  "Balanced human-resource help-desk",
			Slug:   "frame"},
		{
			Id:     "5",
			UserId: 7,
			Title:  "Secured leading edge circuit",
			Slug:   "systemic"},
		{
			Id:     "6",
			UserId: 2,
			Title:  "Public-key dynamic matrix",
			Slug:   "tertiary"},
		{
			Id:     "7",
			UserId: 8,
			Title:  "Triple-buffered attitude-oriented access",
			Slug:   "high-level"},
		{
			Id:     "8",
			UserId: 2,
			Title:  "Up-sized content-based installation",
			Slug:   "Polarised"},
		{
			Id:     "9",
			UserId: 8,
			Title:  "Monitored even-keeled infrastructure",
			Slug:   "asynchronous"},
		{
			Id:     "10",
			UserId: 9,
			Title:  "Open-source object-oriented application",
			Slug:   "contextually-based"},
	}

	rsp.Articles = articles
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.blog"),
	)

	service.Init()

	proto.RegisterBlogHandler(service.Server(), new(Blog))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}