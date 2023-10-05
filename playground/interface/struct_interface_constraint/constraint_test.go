package struct_interface_constraint

import "testing"

func TestI_Hello(t *testing.T) {
	type fields struct {
		Constraint Constraint
		Name       string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
		{"A", fields{Name: "A"}},
		{"B", fields{Name: "B"}},
		{"Empty", fields{Name: ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &I{
				Constraint: tt.fields.Constraint,
				Name:       tt.fields.Name,
			}
			i.Hello()
		})
	}
}

func TestI_Print(t *testing.T) {
	type fields struct {
		Constraint Constraint
		Name       string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &I{
				Constraint: tt.fields.Constraint,
				Name:       tt.fields.Name,
			}
			if got := i.Print(); got != tt.want {
				t.Errorf("Print() = %v, want %v", got, tt.want)
			}
		})
	}
}
