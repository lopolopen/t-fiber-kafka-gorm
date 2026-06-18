package x

import "testing"

func c() {}

type s struct{}

func Test_componentName(t *testing.T) {
	type args struct {
		component any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "string",
			args: args{
				component: "C",
			},
			want: "C",
		},
		{
			name: "func",
			args: args{
				component: c,
			},
			want: "gitlab.loopopen.com/pigou/admin-api/internal/pkg/x.c",
		},
		{
			name: "struct",
			args: args{
				component: s{},
			},
			want: "gitlab.loopopen.com/pigou/admin-api/internal/pkg/x.s",
		},
		{
			name: "struct",
			args: args{
				component: &s{},
			},
			want: "gitlab.loopopen.com/pigou/admin-api/internal/pkg/x.s",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := componentName(tt.args.component); got != tt.want {
				t.Errorf("componentName() = %v, want %v", got, tt.want)
			}
		})
	}
}
