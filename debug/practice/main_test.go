package practice

import "testing"

func Test_greet(t *testing.T) {
	type args struct {
		user string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Chris", args: args{user: "Chris"}, want: "Hello Chris"},
		{name: "Empty", args: args{user: ""}, want: "Hello world"},
		//{name: "Empty string", args: args{user: "Chris"}, want: "Hello Chris"},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := greet(tt.args.user); got != tt.want {
				t.Errorf("greet() = %v, want %v", got, tt.want)
			}
		})
	}
}
