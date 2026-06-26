package confx

import "testing"

func Test_smartKey(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal",
			args: args{
				input: "mysql_dsn",
			},
			want: "mysql.dsn",
		},
		{
			name: "with_underscope",
			args: args{
				input: "orm_log__level",
			},
			want: "orm.log_level",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smartKey(tt.args.input); got != tt.want {
				t.Errorf("smartKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
