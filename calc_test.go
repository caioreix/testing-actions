package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestSum(t *testing.T) {
// 	tt := []struct{
// 		name string
// 		x int
// 		y int
// 		want int
// 	}{
// 		{
// 			name: "x: 1,
// 			y: 5,Soma normal",

// 			want: 6,
// 		},
// 		{
// 			name: "Soma de numeros negativos",
// 			x: -1,
// 			y: -3,
// 			want: -4,
// 		},
// 		{
// 			name: "Soma de numeros negativos",
// 			x: -1,
// 			y: -3,
// 			want: -4,
// 		},
// 	}
// }

func TestSetArgs(t *testing.T) {
	tt := []struct {
		name    string
		command *Command
		args    []string
		want    *Command
	}{
		{
			name:    "Sucesso",
			command: &Command{},
			args:    []string{"nome", "idade", "cidade"},
			want: &Command{
				Args: []string{"nome", "idade"},
			},
		},
		{
			name:    "Se um commando com nome ... foi criado e apos isso adicionado outro commando vai retornar error",
			command: &Command{},
			args:    []string{"nome", "...descricao", "cidade"},
			want: &Command{
				Errs: []error{erroDeMultiplasChamadas},
			},
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := tc.command.SetArgs(tc.args...)
			assert.Equal(t, tc.want, got)
		})
	}
}
