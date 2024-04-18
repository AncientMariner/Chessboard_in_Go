package figures

import (
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
		{"Test get name", fields{Figure{}}, 'k'},
		{"Test get name with empty figure", fields{}, 'k'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (&King{
				tt.fields.Figure,
			}).GetName(); got != tt.want {
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
		want   int
	}{
		{"Test handle board size 8", fields{Figure{next: nil}}, args{board: "____\n"}, 4},
		{"Test handle board size 8", fields{Figure{next: nil}}, args{board: "________\n"}, 8},
		{"Test handle board size 64", fields{Figure{next: nil}}, args{board: "________\n________\n________\n________\n________\n________\n________\n________\n"}, 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			king := &King{
				Figure: tt.fields.Figure,
			}
			if got := king.Handle(tt.args.board); got.Size() != tt.want {
				t.Errorf("Handle() size = %v, want %v", got, tt.want)
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
		want []rune
	}{
		{"Test horizontally 1 on right", args{
			out:      []rune{'_', '_', '_', '_'},
			position: 0,
		}, []rune{'_', 'x', '_', '_'}},
		{"Test horizontally 1 on left", args{
			out:      []rune{'_', '_', '_', '_'},
			position: 3,
		}, []rune{'_', '_', 'x', '_'}},
		{"Test horizontally left and right", args{
			out:      []rune{'_', '_', '_', '_'},
			position: 1,
		}, []rune{'x', '_', 'x', '_'}},
		{"Test horizontally left and right another field", args{
			out:      []rune{'_', '_', '_', '_'},
			position: 2,
		}, []rune{'_', 'x', '_', 'x'}},
		{"Test horizontally left and right another field", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 7,
		}, []rune{'_', '_', '_', '_', '_', '_', 'x', '_', '\n'}},
		{"Test horizontally non existing case", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 8,
		}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test horizontally left and right another field", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 9,
		}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', 'x', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test horizontally", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 16,
		}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', 'x', '_', '\n'}},
		{"Test horizontally non existing case", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 17,
		}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test horizontally non existing case", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 18,
		}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test horizontally non existing case", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 25,
		}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', 'x', '_', '\n'}},
		{"Test horizontally non existing case", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 26,
		}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			placeAttackPlacesHorizontally(tt.args.out, tt.args.position)
			if !reflect.DeepEqual(tt.args.out, tt.want) {
				t.Errorf("placeAttackPlacesHorizontally() = %v, want %v", tt.args.out, tt.want)
			}
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
		want []rune
	}{
		{"Test vertically 1 below", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 0,
		}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', 'x', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test vertically 1 below another one", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 1,
		}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', 'x', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test vertically 1 below another one", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 7,
		}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', 'x', '\n'}},
		{"Test vertically non existing case", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 8,
		}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test vertically one above and below", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 9,
		}, []rune{'x', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', 'x', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test vertically one above and below", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 16,
		}, []rune{'_', '_', '_', '_', '_', '_', '_', 'x', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', 'x', '\n'}},
		{"Test vertically non existing case", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 17,
		}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test vertically one above", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 18,
		}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', 'x', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			placeAttackPlacesVertically(tt.args.out, tt.args.position)
			if !reflect.DeepEqual(tt.args.out, tt.want) {
				t.Errorf("placeAttackPlacesVertically() = %v, want %v", tt.args.out, tt.want)
			}
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
		want []rune
	}{
		{"Test diag 1 above", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 0,
		}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test diag 1 above", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 1,
		}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test diag 1 above", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 7,
		}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test diag 1 above non existing case", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 8,
		}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test diag above", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 9,
		}, []rune{'_', 'x', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test diag above", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 10,
		}, []rune{'x', '_', 'x', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test diag above", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 15,
		}, []rune{'_', '_', '_', '_', '_', 'x', '_', 'x', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test diag above", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 16,
		}, []rune{'_', '_', '_', '_', '_', '_', 'x', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test diag above non existing case", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 17,
		}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test diag above", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 18,
		}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', 'x', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			placeDiagonallyAbove(tt.args.out, tt.args.position)
			if !reflect.DeepEqual(tt.args.out, tt.want) {
				t.Errorf("placeDiagonallyAbove() = %v, want %v", tt.args.out, tt.want)
			}
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
		want []rune
	}{
		{"Test diag below", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 0,
		}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', 'x', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test diag below", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 1,
		}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', 'x', '_', 'x', '_', '_', '_', '_', '_', '\n'}},
		{"Test diag below", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 7,
		}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', 'x', '_', '\n'}},
		{"Test diag below non existing case", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 8,
		}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test diag below", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 9,
		}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test diag below", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 16,
		}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test diag below non existing case", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 17,
		}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test diag below", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 18,
		}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', 'x', '_', '_', '_', '_', '_', '_', '\n'}},
		{"Test diag below", args{
			out:      []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n'},
			position: 19,
		}, []rune{'_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', '_', '_', '_', '_', '_', '_', '_', '_', '\n', 'x', '_', 'x', '_', '_', '_', '_', '_', '\n'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			placeDiagonallyBelow(tt.args.out, tt.args.position)
			if !reflect.DeepEqual(tt.args.out, tt.want) {
				t.Errorf("placeDiagonallyBelow() = %v, want %v", tt.args.out, tt.want)
			}
		})
	}
}
