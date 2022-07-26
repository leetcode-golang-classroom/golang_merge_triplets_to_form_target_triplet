package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	triplets := [][]int{{2, 5, 3}, {2, 3, 4}, {1, 2, 5}, {5, 2, 3}}
	target := []int{5, 5, 5}
	for idx := 0; idx < b.N; idx++ {
		mergeTriplets(triplets, target)
	}
}
func Test_mergeTriplets(t *testing.T) {
	type args struct {
		triplets [][]int
		target   []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "triplets = [[2,5,3],[1,8,4],[1,7,5]], target = [2,7,5]",
			args: args{triplets: [][]int{{2, 5, 3}, {1, 8, 4}, {1, 7, 5}}, target: []int{2, 7, 5}},
			want: true,
		},
		{
			name: "triplets = [[3,4,5],[4,5,6]], target = [3,2,5]",
			args: args{triplets: [][]int{{3, 4, 5}, {4, 5, 6}}, target: []int{3, 2, 5}},
			want: false,
		},
		{
			name: "triplets = [[2,5,3],[2,3,4],[1,2,5],[5,2,3]], target = [5,5,5]",
			args: args{triplets: [][]int{{2, 5, 3}, {2, 3, 4}, {1, 2, 5}, {5, 2, 3}}, target: []int{5, 5, 5}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTriplets(tt.args.triplets, tt.args.target); got != tt.want {
				t.Errorf("mergeTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
