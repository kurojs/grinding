package remove_k_digits

import "testing"

func Test_removeKdigits(t *testing.T) {
	type args struct {
		num string
		k   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				num: "10",
				k:   1,
			},
			want: "0",
		},
		{
			name: "case 2",
			args: args{
				num: "1432219",
				k:   3,
			},
			want: "1219",
		},
		{
			name: "case 3",
			args: args{
				num: "10200",
				k:   1,
			},
			want: "200",
		},
		{
			name: "case 4",
			args: args{
				num: "100",
				k:   1,
			},
			want: "0",
		},
		{
			name: "case 5",
			args: args{
				num: "112",
				k:   1,
			},
			want: "11",
		},
		{
			name: "case 6",
			args: args{
				num: "1234567890",
				k:   9,
			},
			want: "0",
		},
		{
			name: "case 7",
			args: args{
				num: "1432219",
				k:   3,
			},
			want: "1219",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeKdigits(tt.args.num, tt.args.k); got != tt.want {
				t.Errorf("removeKdigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
