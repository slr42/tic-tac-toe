package tic_tac_toe_test

import (
	"reflect"
	"testing"
	. "tic_tac_toe/utility"
)

func TestAnalyzeResult(t *testing.T) {
	type args struct {
		board        *Board
		playerResult Result
	}
	tests := []struct {
		name string
		args args
		want AnalyzeReport
	}{
		{
			name: "Empty board",
			args: args{
				board: &Board{Field: [3][3]string{{"", "", ""}, {"", "", ""}, {"", "", ""}}},
				playerResult: Result{XCount: [3]int{}, YCount: [3]int{}, Diagonal1Count: 0, Diagonal2Count: 0},
			},
			want: AnalyzeReport{
				Attacks: []Position{{X: 1, Y: 1}},
				Weights: [3][3]int{{3, 2, 3}, {2, 4, 2}, {3, 2, 3}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnalyzeResult(tt.args.board, tt.args.playerResult)
			if !reflect.DeepEqual(got.Weights, tt.want.Weights) {
				t.Errorf("AnalyzeResult().Weights = %v, want %v", got.Weights, tt.want.Weights)
			}
			if !reflect.DeepEqual(got.Attacks, tt.want.Attacks) {
				t.Errorf("AnalyzeResult().Attacks = %v, want %v", got.Attacks, tt.want.Attacks)
			}
			if !reflect.DeepEqual(got.Defends, tt.want.Defends) {
				t.Errorf("AnalyzeResult().Defends = %v, want %v", got.Defends, tt.want.Defends)
			}
		})
	}
}

func TestChooseBestPosition(t *testing.T) {
	type args struct {
		positionList []Position
	}
	tests := []struct {
		name string
		args args
		want Position
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChooseBestPosition(tt.args.positionList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChooseBestPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}
