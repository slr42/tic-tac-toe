package tic_tac_toe_test

import (
	"reflect"
	"testing"

	. "github.com/slr42/tic-tac-toe/utility"
)

func TestAnalyzeResult(t *testing.T) {
	type args struct {
		board  *Board
		bot    *Player
		player *Player
	}
	tests := []struct {
		name string
		args args
		want AnalyzeReport
	}{
		{
			name: "Empty board",
			args: args{
				board:  &Board{Field: [3][3]string{{"", "", ""}, {"", "", ""}, {"", "", ""}}},
				bot:    &Player{Name: "Bot", Mark: "0"},
				player: &Player{
					Name: "Player",
					Mark: "X",
					Result: Result{XCount: [3]int{}, YCount: [3]int{}, Diagonal1Count: 0, Diagonal2Count: 0},
				},
			},
			want: AnalyzeReport{
				Attacks: []Position{{X: 1, Y: 1}},
				Weights: [3][3]int{{3, 2, 3}, {2, 4, 2}, {3, 2, 3}},
			},
		},
		{
			name: "Middle end",
			args: args{
				board:        &Board{Field: [3][3]string{{"X", "0", "X"}, {"X", "0", ""}, {"0", "", ""}}},
				bot:          &Player{Name: "Bot", Mark: "0"},
				player: &Player{
					Name: "Player",
					Mark: "X",
					Result: Result{XCount: [3]int{2, 1, 0}, YCount: [3]int{2, 0, 1}, Diagonal1Count: 1, Diagonal2Count: 1},
				},
			},
			want: AnalyzeReport{
				Attacks: []Position{{2, 1}},
				Weights: [3][3]int{{0, 1, 0}, {0, 1, 0}, {1, 2, 1}},
			},
		},
		{
			name: "Anti fork",
			args: args{
				board:        &Board{Field: [3][3]string{{"0", "", ""}, {"", "", "X"}, {"X", "", ""}}},
				bot:          &Player{
					Name: "Bot",
					Mark: "0",
					Result: Result{XCount: [3]int{1, 0, 0}, YCount: [3]int{1, 0, 0}, Diagonal1Count: 1, Diagonal2Count: 0},
				},
				player: &Player{
					Name: "Player",
					Mark: "X",
					Result: Result{XCount: [3]int{0, 1, 1}, YCount: [3]int{1, 0, 1}, Diagonal1Count: 0, Diagonal2Count: 1},
				},
			},
			want: AnalyzeReport{
				Attacks: []Position{{0, 1}},
				PreemptiveDefends: []Position{{0, 2}},
			},
		},
		{
			name: "Anti fork 2",
			args: args{
				board:        &Board{Field: [3][3]string{{"X", "", ""}, {"", "", ""}, {"", "", ""}}},
				bot:          &Player{
					Name: "Bot",
					Mark: "0",
					Result: Result{XCount: [3]int{0, 0, 0}, YCount: [3]int{0, 0, 0}, Diagonal1Count: 0, Diagonal2Count: 0},
				},
				player: &Player{
					Name: "Player",
					Mark: "X",
					Result: Result{XCount: [3]int{1, 0, 0}, YCount: [3]int{1, 0, 0}, Diagonal1Count: 1, Diagonal2Count: 0},
				},
			},
			want: AnalyzeReport{
				Attacks: []Position{{0, 2}, {1, 1}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnalyzeResult(tt.args.board, tt.args.bot, tt.args.player)
			if !reflect.DeepEqual(got.Attacks, tt.want.Attacks) {
				t.Errorf("AnalyzeResult().Attacks = %v, want %v", got.Attacks, tt.want.Attacks)
			}
			if !reflect.DeepEqual(got.PreemptiveDefends, tt.want.PreemptiveDefends) {
				t.Errorf("AnalyzeResult().PreemptiveDefends = %v, want %v", got.PreemptiveDefends, tt.want.PreemptiveDefends)
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
		{
			name: "One choice",
			args: args{
				positionList: []Position{{1, 1}},
			},
			want: Position{X: 1, Y: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChooseBestPosition(tt.args.positionList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChooseBestPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPositionWeight(t *testing.T) {
	type args struct {
		board *Board
		mark  string
		x     int
		y     int
	}
	tests := []struct {
		name       string
		args       args
		wantWeight int
		wantErr    string
	}{
		{
			name: "Check opponent marked positions",
			args: args {
				board: &Board{Field: [3][3]string{{"X", "0", "X"}, {"X", "0", ""}, {"0", "", ""}}},
				mark:  "0",
				x:     1,
				y:     1,
			},
			wantWeight: 1,
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWeight, gotErr := GetPositionWeight(tt.args.board, tt.args.mark, tt.args.x, tt.args.y)
			if gotWeight != tt.wantWeight {
				t.Errorf("GetPositionWeight() gotWeight = %v, want %v", gotWeight, tt.wantWeight)
			}
			if gotErr != tt.wantErr {
				t.Errorf("GetPositionWeight() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
