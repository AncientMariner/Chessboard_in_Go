package figures

import (
	"github.com/hashicorp/go-set/v2"
	"reflect"
	"testing"
)

func TestKing_GetName(t *testing.T) {
	type fields struct {
		Figure Figure
	}
	tests := []struct {
		name   string
		fields fields
		want   rune
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ki := &King{
				Figure: tt.fields.Figure,
			}
			if got := ki.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKing_Handle(t *testing.T) {
	type fields struct {
		Figure Figure
	}
	type args struct {
		board string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *set.HashSet[*FigurePosition, string]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			king := &King{
				Figure: tt.fields.Figure,
			}
			if got := king.Handle(tt.args.board); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_placeAttackPlacesHorizontally(t *testing.T) {
	type args struct {
		out      []rune
		position int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			placeAttackPlacesHorizontally(tt.args.out, tt.args.position)
		})
	}
}

func Test_placeAttackPlacesVertically(t *testing.T) {
	type args struct {
		out      []rune
		position int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			placeAttackPlacesVertically(tt.args.out, tt.args.position)
		})
	}
}

func Test_placeDiagonallyAbove(t *testing.T) {
	type args struct {
		out      []rune
		position int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			placeDiagonallyAbove(tt.args.out, tt.args.position)
		})
	}
}

func Test_placeDiagonallyBelow(t *testing.T) {
	type args struct {
		out      []rune
		position int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			placeDiagonallyBelow(tt.args.out, tt.args.position)
		})
	}
}
