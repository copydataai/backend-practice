package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type RequestReview struct {
	message string
}

type RequestProduct struct {
	Name          string `json:"name"`
	Detail        string `json:"detail"`
	Trademark     string `json:"trademark"`
	SKU           string
	Manufacturing string `json:"manufacturing"`
}

func Query(sql string) string {
	fmt.Print(sql)
	return sql
}

func productType(name, detail string) string {
	sqlStatment := fmt.Sprintf("INSERT INTO product_type(name, detail) VALUES(%s, %s);", name, detail)
	rows := Query(sqlStatment)
	return rows
}

func product(name, detail, trademark, manufacturing, sku string) string {
	id := productType(name, detail)
	sqlStatment := fmt.Sprintf("INSERT INTO products(detail, trademark, manufacturing, sku, created_at) VALUES(%d, %s, %s %s, %s)", id, trademark, manufacturing, sku, time.Now().String())
	rows := Query(sqlStatment)
	return rows
}

func reviews(title, content string, author, product uint32) string {
	sqlStatment := fmt.Sprintf("INSERT INTO reviews(title, content, user, product created_at) VALUES(%s, %s, %d, %d, %s)", title, content, author, product, time.Now().String())
	rows := Query(sqlStatment)
	return rows
}

func main() {
	lambda.Start()
}
