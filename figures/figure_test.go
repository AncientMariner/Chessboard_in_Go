package figures

import (
	"reflect"
	"testing"
)

func TestFigure_GetNext(t *testing.T) {
	type fields struct {
		next FigureBehaviour
	}
	tests := []struct {
		name   string
		fields fields
		want   FigureBehaviour
	}{
		{"Test get next nil", fields{nil}, nil},
		{"Test get next nil", fields{&King{Figure{next: nil}}}, &King{Figure{next: nil}}},
		{"Test get next nil", fields{&Queen{Figure{next: nil}}}, &Queen{Figure{next: nil}}},
		{"Test get next nil", fields{&Bishop{Figure{next: nil}}}, &Bishop{Figure{next: nil}}},
		{"Test get next nil", fields{&Rook{Figure{next: nil}}}, &Rook{Figure{next: nil}}},
		{"Test get next nil", fields{&Knight{Figure{next: nil}}}, &Knight{Figure{next: nil}}},
		{"Test get next nil", fields{&Knight{Figure{next: &King{Figure{next: &Queen{}}}}}}, &Knight{Figure{next: &King{Figure{next: &Queen{}}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Figure{
				next: tt.fields.next,
			}
			if got := f.GetNext(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNext() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestFigure_GetNextGetNext(t *testing.T) {
	type fields struct {
		next FigureBehaviour
	}
	tests := []struct {
		name   string
		fields fields
		want   FigureBehaviour
	}{
		{"Test get next nil", fields{&Knight{Figure{next: &King{Figure{next: &Queen{}}}}}}, &King{Figure{next: &Queen{}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Figure{
				next: tt.fields.next,
			}
			if got := f.GetNext().GetNext(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFigure_SetNext(t *testing.T) {
	type fields struct {
		next FigureBehaviour
	}
	type args struct {
		next FigureBehaviour
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   FigureBehaviour
	}{
		{"Test set next", fields{&Knight{Figure{next: nil}}}, args{&Queen{Figure{next: nil}}}, &Queen{Figure{next: nil}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Figure{
				next: tt.fields.next,
			}
			if got := f.SetNext(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetNext() = %v, want %v", got, tt.want)
			}
		})
	}
}
