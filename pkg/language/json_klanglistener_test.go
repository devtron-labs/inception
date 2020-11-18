package language

import (
	"github.com/devtron-labs/inception/pkg/language/parser"
	"testing"
)

func TestKlangListener_handleJsonSelect(t *testing.T) {
	type fields struct {
		input  string
		values map[string]valHolder
	}
	type args struct {
		ctx *parser.AssignmentContext
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Json Select",
			fields: fields{
				input: `
x = {"name":{"first":"abc","last":"def"}};
y = jsonSelect(x, "name.last");
`,
				values: map[string]valHolder{
					"y": {
						dataType: "STRING",
						name:     "y",
						value:    "def",
					},
				},
			},
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := setup(tt.fields.input)
			if d, ok := r.values["y"]; ok {
				m := map[string]valHolder{
					"y": d,
				}
				if diff := compare(tt.fields.values, m); !diff {
					t.Errorf("expected %+v, found %+v\n", tt.fields.values, r.Values())
				}
			} else {
				t.Errorf("expected %+v, found %+v\n", tt.fields.values, r.Values())
			}
		})
	}
}

func TestKlangListener_handleJsonEdit(t *testing.T) {
	type fields struct {
		input  string
		values map[string]valHolder
	}
	type args struct {
		ctx *parser.AssignmentContext
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Json Edit",
			fields: fields{
				input: `
x = {"name":{"first":"abc","last":"def"}};
jsonEdit(x, "name.first", "xyz");
`,
				values: map[string]valHolder{
					"x": {
						dataType: "STRING",
						name:     "x",
						value:    "{\"name\":{\"first\":\"xyz\",\"last\":\"def\"}}",
					},
				},
			},
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := setup(tt.fields.input)
			if d, ok := r.values["x"]; ok {
				m := map[string]valHolder{
					"x": d,
				}
				if diff := compare(tt.fields.values, m); !diff {
					t.Errorf("expected %+v, found %+v\n", tt.fields.values, r.Values())
				}
			} else {
				t.Errorf("expected %+v, found %+v\n", tt.fields.values, r.Values())
			}
		})
	}
}

