package figures

import (
	"github.com/hashicorp/go-set/v2"
	"reflect"
	"testing"
)

func TestRook_GetName(t *testing.T) {
	type fields struct {
		Figure Figure
	}
	tests := []struct {
		name   string
		fields fields
		want   rune
	}{
		{"Test get name", fields{Figure{}}, 'r'},
		{"Test get name with empty figure", fields{}, 'r'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ro := &Rook{
				Figure: tt.fields.Figure,
			}
			if got := ro.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRook_Handle(t *testing.T) {
	type fields struct {
		Figure Figure
	}
	type args struct {
		request string
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
			rook := &Rook{
				Figure: tt.fields.Figure,
			}
			if got := rook.Handle(tt.args.request); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRook_placeAttackPlacesHorizontally(t *testing.T) {
	type args struct {
		out      []rune
		position int
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{"Test horizontal placement", args{[]rune{'_', '_', '_', '_'},
			0}, []rune{'_', 'x', 'x', 'x'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			(&Rook{}).placeAttackPlacesHorizontally(tt.args.out, tt.args.position)
			if !reflect.DeepEqual(tt.args.out, tt.want) {
				t.Errorf("placeAttackPlacesHorizontally() = %v, want %v", tt.args.out, tt.want)
			}
		})
	}
}
