package trs

import (
	"slices"
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

func TestTerm_Unfold(t *testing.T) {
	type fields struct {
		word      string
		variables []string
	}
	type args struct {
		trs string
		n   int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "simple test 1",
			fields: fields{
				word:      "g(g(t,t),h(t))",
				variables: []string{"t", "x", "y"},
			},
			args: args{
				trs: `g(x,y) -> g(h(x),y)
				h(t) -> f`,
				n: 2,
			},
			want: []string{"g(f(),h(t))", "g(h(h(g(t,t))),h(t))", "g(h(g(h(t),t)),h(t))", "g(h(g(t,t)),f())", "g(g(h(h(t)),t),h(t))", "g(g(f(),t),h(t))", "g(g(h(t),t),f())"},
		},
		{
			name: "test from TA 1",
			fields: fields{
				word:      "g(g(t,t),h(t))",
				variables: []string{"x", "t"},
			},
			args: args{
				trs: `g(x,x) -> h(x)
				h(t) -> f`,
				n: 2,
			},
			want: []string{"h(h(t))", "g(f(),h(t))", "g(h(t),f())"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr, err := NewTermFromString(tt.fields.word, tt.fields.variables)
			if err != nil {
				t.Error(err)
				return
			}
			trss, err := NewTermRewritingSystemFromString(tt.args.trs, tt.fields.variables)
			if err != nil {
				t.Error(err)
				return
			}

			got := tr.UnfoldDeep(trss, tt.args.n)
			if len(got) != len(tt.want) {
				t.Errorf("Term.Unfold() = %v, want %v", got, tt.want)
				return
			}

			gotTerms := make([]string, 0)
			for _, word := range got {
				gotTerms = append(gotTerms, word.String())
			}

			for _, term := range gotTerms {
				if !slices.Contains(tt.want, term) {
					t.Errorf("Term.Unfold() = %v, want %v, ", got, tt.want)
				}
			}
		})
	}
}
