package trs

import (
	"testing"
)

func TestTerm_IsCompatible(t *testing.T) {
	type fields struct {
		Type      TermType
		Symbol    string
		Arguments []Term
	}
	type args struct {
		rhs Term
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "test same vars",
			fields: fields{
				Type:      TermTypeVariable,
				Symbol:    "a",
				Arguments: []Term{},
			},
			args: args{
				rhs: Term{
					Type:      TermTypeVariable,
					Symbol:    "a",
					Arguments: []Term{},
				},
			},
			want: true,
		},
		{
			name: "test different vars",
			fields: fields{
				Type:      TermTypeVariable,
				Symbol:    "a",
				Arguments: []Term{},
			},
			args: args{
				rhs: Term{
					Type:      TermTypeVariable,
					Symbol:    "b",
					Arguments: []Term{},
				},
			},
			want: true,
		},
		{
			name: "test different types",
			fields: fields{
				Type:      TermTypeVariable,
				Symbol:    "a",
				Arguments: []Term{},
			},
			args: args{
				rhs: Term{
					Type:      TermTypeFunction,
					Symbol:    "a",
					Arguments: []Term{},
				},
			},
			want: false,
		},
		{
			name: "test complex 1",
			fields: fields{
				Type:      TermTypeFunction,
				Symbol:    "f",
				Arguments: []Term{{TermTypeVariable, "a", []Term{}}, {TermTypeVariable, "a", []Term{}}, {TermTypeVariable, "b", []Term{}}},
			},
			args: args{
				rhs: Term{
					Type:      TermTypeFunction,
					Symbol:    "g",
					Arguments: []Term{{TermTypeVariable, "x", []Term{}}, {TermTypeVariable, "y", []Term{}}, {TermTypeVariable, "x", []Term{}}},
				},
			},
			want: false,
		},
		{
			name: "test complex 2",
			fields: fields{
				Type:      TermTypeFunction,
				Symbol:    "f",
				Arguments: []Term{{TermTypeVariable, "a", []Term{}}, {TermTypeVariable, "a", []Term{}}, {TermTypeVariable, "b", []Term{}}},
			},
			args: args{
				rhs: Term{
					Type:      TermTypeFunction,
					Symbol:    "g",
					Arguments: []Term{{TermTypeVariable, "x", []Term{}}, {TermTypeVariable, "x", []Term{}}, {TermTypeVariable, "y", []Term{}}},
				},
			},
			want: false,
		},
		{
			name: "test complex 3",
			fields: fields{
				Type:      TermTypeFunction,
				Symbol:    "f",
				Arguments: []Term{{TermTypeVariable, "a", []Term{}}, {TermTypeVariable, "a", []Term{}}, {TermTypeVariable, "b", []Term{}}},
			},
			args: args{
				rhs: Term{
					Type:      TermTypeFunction,
					Symbol:    "f",
					Arguments: []Term{{TermTypeVariable, "x", []Term{}}, {TermTypeVariable, "x", []Term{}}, {TermTypeVariable, "y", []Term{}}},
				},
			},
			want: true,
		},
		{
			name: "test complex 4",
			fields: fields{
				Type:      TermTypeFunction,
				Symbol:    "f",
				Arguments: []Term{{TermTypeFunction, "g", []Term{{TermTypeVariable, "a", []Term{}}, {TermTypeVariable, "b", []Term{}}, {TermTypeVariable, "c", []Term{}}}}, {TermTypeVariable, "a", []Term{}}, {TermTypeVariable, "b", []Term{}}},
			},
			args: args{
				rhs: Term{
					Type:      TermTypeFunction,
					Symbol:    "f",
					Arguments: []Term{{TermTypeFunction, "g", []Term{{TermTypeVariable, "x", []Term{}}, {TermTypeVariable, "y", []Term{}}, {TermTypeVariable, "z", []Term{}}}}, {TermTypeVariable, "x", []Term{}}, {TermTypeVariable, "y", []Term{}}},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lhs := Term{
				Type:      tt.fields.Type,
				Symbol:    tt.fields.Symbol,
				Arguments: tt.fields.Arguments,
			}
			if got := lhs.IsEquival(tt.args.rhs); got != tt.want {
				t.Errorf("Term.IsCompatible() = %v, want %v", got, tt.want)
			}
		})
	}
}
