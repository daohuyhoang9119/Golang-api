package service

import (
	"Golang-crud/data/request"
	"Golang-crud/data/response"
	"context"
)

type BookService interface {
	Create(ctx context.Context, request request.BookCreateRequest)
	Update(ctx context.Context, request request.BookUpdateRequest)
	Delete(ctx context.Context, bookId int)
	FindById(ctx context.Context, bookId int) response.BookResponse
	FindAll(ctx context.Context) []response.BookResponse
}