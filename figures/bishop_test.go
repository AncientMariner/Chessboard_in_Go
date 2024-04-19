package figures

import (
	"reflect"
	"testing"
)

func TestBishop_GetName(t *testing.T) {
	type fields struct {
		Figure Figure
	}
	tests := []struct {
		name   string
		fields fields
		want   rune
	}{
		{"Test get name", fields{Figure{}}, 'b'},
		{"Test get name with empty figure", fields{}, 'b'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bi := &Bishop{
				Figure: tt.fields.Figure,
			}
			if got := bi.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBishop_Handle(t *testing.T) {
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
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bishop := &Bishop{
				Figure: tt.fields.Figure,
			}
			if got := bishop.Handle(tt.args.board); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBishop_placeAttackPlacesDiagonallyBelow(t *testing.T) {
	type args struct {
		out      []rune
		position int
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{"Test position#", args{[]rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			0}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', 'x', '_', '_', '_', '_', '_', '_', '\n', '_', '_', 'x', '_', '_', '_', '_', '_', '\n'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			(&Bishop{}).placeAttackPlacesDiagonallyBelow(tt.args.out, tt.args.position)
			if !reflect.DeepEqual(tt.args.out, tt.want) {
				t.Errorf("placeAttackPlacesDiagonallyBelow() = %v, want %v", tt.args.out, tt.want)
			}
		})
	}
}
