package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	// Aqui atribuimos um valor ao contexto usando o conceito de chave-valor
	ctx = context.WithValue(ctx, "token", "senha")
	bookHotel(ctx, "Hotel")
}

// por convenção, a variavel de contexto sempre é a primeira na assinatura de uma função
func bookHotel(ctx context.Context, name string) {
	// Aqui resgatamos o valor setado
	token := ctx.Value("token")
	fmt.Println(token)
}
