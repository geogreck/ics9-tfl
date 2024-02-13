package trs

import (
	"reflect"
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
			want: []string{"g(f,h(t))", "g(h(h(g(t,t))),h(t))", "g(h(g(h(t),t)),h(t))", "g(h(g(t,t)),f)", "g(g(h(h(t)),t),h(t))", "g(g(f,t),h(t))", "g(g(h(t),t),f)"},
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
			want: []string{"h(h(t))", "g(f,h(t))", "g(h(t),f)"},
		},
		{
			name: "test test",
			fields: fields{
				word:      "g(g(t,t),h(t))",
				variables: []string{"x", "t"},
			},
			args: args{
				trs: `g(x,x) -> h(x)
				h(t) -> f`,
				n: 1,
			},
			want: []string{"g(h(t),h(t))", "g(g(t,t),f)"},
		},
		{
			name: "test test test",
			fields: fields{
				word:      "g(g(t,t),h(t))",
				variables: []string{"x", "t"},
			},
			args: args{
				trs: `g(x,x) -> h(x)
				h(t) -> f(t)
				f(t) -> d(t,t,t)`,
				n: 4,
			},
			want: []string{"d(h(t),h(t),h(t))", "f(f(t))", "h(d(t,t,t))", "g(d(t,t,t),f(t))", "h(f(t))", "g(f(t),d(t,t,t))"},
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

func TestTerm_BindArguments(t *testing.T) {
	type args struct {
		lhsWord   string
		variables []string
		rhsWord   string
	}
	tests := []struct {
		name      string
		args      args
		wantWords map[string]string
		wantErr   bool
	}{
		{
			name: "basic test ok",
			args: args{
				lhsWord:   "g(t, t)",
				rhsWord:   "g(x, x)",
				variables: []string{"x", "t"},
			},
			wantWords: map[string]string{
				"t": "x",
			},
			wantErr: false,
		},
		{
			name: "basic test fail",
			args: args{
				lhsWord:   "g(t, t)",
				rhsWord:   "g(x, y)",
				variables: []string{"x", "t"},
			},
			wantWords: map[string]string{},
			wantErr:   true,
		},
		{
			name: "basic test ok",
			args: args{
				lhsWord:   "g(t, t)",
				rhsWord:   "g(x, x)",
				variables: []string{"x", "t"},
			},
			wantWords: map[string]string{
				"t": "x",
			},
			wantErr: false,
		},
		{
			name: "mid test ok",
			args: args{
				lhsWord:   "g(x, y)",
				rhsWord:   "g(g(a, b), c)",
				variables: []string{"x", "y", "a", "b", "c"},
			},
			wantWords: map[string]string{
				"x": "g(a,b)",
				"y": "c",
			},
			wantErr: false,
		},
		{
			name: "mid test fail",
			args: args{
				lhsWord:   "g(g(a, b), c)",
				rhsWord:   "g(x, y)",
				variables: []string{"x", "y", "a", "b", "c"},
			},
			wantWords: map[string]string{
				"x": "g(a,b)",
				"y": "c",
			},
			wantErr: true,
		},
		{
			name: "big test ok",
			args: args{
				lhsWord:   "f(a,b,b)",
				rhsWord:   "f(g(a,g(b)),h(a),h(a))",
				variables: []string{"x", "y", "a", "b", "c"},
			},
			wantWords: map[string]string{
				"a": "g(a,g(b))",
				"b": "h(a)",
			},
			wantErr: false,
		},
		{
			name: "big test ok",
			args: args{
				lhsWord:   "f(a,b,c)",
				rhsWord:   "f(g(a,g(b)),h(a),h(a))",
				variables: []string{"x", "y", "a", "b", "c"},
			},
			wantWords: map[string]string{
				"a": "g(a,g(b))",
				"b": "h(a)",
				"c": "h(a)",
			},
			wantErr: false,
		},
		{
			name: "ta test 2",
			args: args{
				lhsWord:   "g(f(z))",
				rhsWord:   " g(f(z))",
				variables: []string{"x", "t"},
			},
			wantWords: map[string]string{
				"z": "z",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lhs, err := NewTermFromString(tt.args.lhsWord, tt.args.variables)
			if err != nil {
				t.Error(err)
				return
			}
			rhs, err := NewTermFromString(tt.args.rhsWord, tt.args.variables)
			if err != nil {
				t.Error(err)
				return
			}

			got, err := lhs.BindArguments(rhs)
			if (err != nil) != tt.wantErr {
				t.Errorf("Term.BindArguments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			gotWords := make(map[string]string)
			for key, value := range got {
				gotWords[key] = value.String()
			}

			if !tt.wantErr && len(got) != len(tt.wantWords) {
				t.Errorf("Term.BindArguments() = %v, want %v", got, tt.wantWords)
				return
			}

			for key, value := range gotWords {
				if val, ok := tt.wantWords[key]; !tt.wantErr && (!ok || val != value) {
					t.Errorf("Term.BindArguments() = %v, want %v", got, tt.wantWords)
					return
				}
			}
		})
	}
}

func TestTerm_ApplyArgsBindings(t *testing.T) {
	type args struct {
		rhs       string
		bindings  map[string]string
		variables []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ta test 1",
			args: args{
				rhs: "g(f(z))",
				bindings: map[string]string{
					"z": "z",
				},
				variables: []string{},
			},
			want: "g(f(z))",
		},
		{
			name: "test 2",
			args: args{
				rhs: "g(f(z))",
				bindings: map[string]string{
					"z": "g(z, z)",
				},
				variables: []string{},
			},
			want: "g(f(g(z,z)))",
		},
		{
			name: "test 3",
			args: args{
				rhs: "g(f(z))",
				bindings: map[string]string{
					"z": "g(g(z))",
				},
				variables: []string{"x", "z"},
			},
			want: "g(f(g(g(z))))",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr, err := NewTermFromString(tt.args.rhs, tt.args.variables)
			if err != nil {
				t.Error(err)
				return
			}
			tb := make(map[string]Term)
			for key, value := range tt.args.bindings {
				trr, err := NewTermFromString(value, tt.args.variables)
				if err != nil {
					t.Error(err)
					return
				}
				tb[key] = trr
			}
			tw, err := NewTermFromString(tt.want, tt.args.variables)
			if err != nil {
				t.Error(err)
				return
			}
			if got := tr.ApplyArgsBindings(tb); !reflect.DeepEqual(got, tw) {
				t.Errorf("Term.ApplyArgsBindings() = %v, want %v", got, tt.want)
			}
		})
	}
}
