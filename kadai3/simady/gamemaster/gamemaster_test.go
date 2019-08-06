package gamemaster

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"testing"
	"time"
)

var (
	allCorrectReader  *os.File
	halfCorrectReader *os.File
	timeUpReader      *os.File
)

func setup(t *testing.T) {
	t.Helper()
	var err error

	allCorrectReader, err = os.Open("../testdata/all_correct.txt")
	if err != nil {
		t.Fatalf("Failed to open file. err = %v", err)
	}

	halfCorrectReader, err = os.Open("../testdata/half_correct.txt")
	if err != nil {
		t.Fatalf("Failed to open file. err = %v", err)
	}

	timeUpReader, err = os.Open("../testdata/time_up.txt")
	if err != nil {
		t.Fatalf("Failed to open file. err = %v", err)
	}
}

func TestNew(t *testing.T) {
	setup(t)
	type args struct {
		reader    io.Reader
		writer    io.Writer
		timeLimit time.Duration
		problems  []string
	}
	tests := []struct {
		name string
		args args
		want GameMaster
	}{
		{
			name: "GameMasterの生成",
			args: args{
				reader:    allCorrectReader,
				writer:    &bytes.Buffer{},
				timeLimit: 10,
				problems:  []string{"apple", "bake", "cup", "dog", "egg", "fight", "green", "hoge", "idea", "japan"},
			},
			want: &gameMaster{
				r:         allCorrectReader,
				w:         &bytes.Buffer{},
				timeLimit: 10,
				problems:  []string{"apple", "bake", "cup", "dog", "egg", "fight", "green", "hoge", "idea", "japan"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.reader, tt.args.writer, tt.args.timeLimit, tt.args.problems); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gameMaster_Play(t *testing.T) {
	setup(t)
	type fields struct {
		r         io.Reader
		w         io.Writer
		timeLimit time.Duration
		problems  []string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "一連のゲーム処理",
			fields: fields{
				r:         allCorrectReader,
				w:         &bytes.Buffer{},
				timeLimit: 10,
				problems:  []string{"apple", "bake", "cup", "dog", "egg", "fight", "green", "hoge", "idea", "japan"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gm := &gameMaster{
				r:         tt.fields.r,
				w:         tt.fields.w,
				timeLimit: tt.fields.timeLimit,
				problems:  tt.fields.problems,
			}
			// 他のテストでパターンは網羅しているので正常終了のみ確認
			gm.Play()
		})
	}
}

func Test_gameMaster_displayRule(t *testing.T) {
	setup(t)
	type fields struct {
		r         io.Reader
		w         io.Writer
		timeLimit time.Duration
		problems  []string
	}
	tests := []struct {
		name        string
		fields      fields
		wantMessage string
	}{
		{
			name: "ルールの表示",
			fields: fields{
				r:         allCorrectReader,
				w:         &bytes.Buffer{},
				timeLimit: 10,
				problems:  []string{"apple", "bake", "cup", "dog", "egg", "fight", "green", "hoge", "idea", "japan"},
			},
			wantMessage: "【タイピングゲーム】画面に表示される英単語をできるだけ多く入力しましょう！\n制限時間は10秒です。\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gm := &gameMaster{
				r:         tt.fields.r,
				w:         tt.fields.w,
				timeLimit: tt.fields.timeLimit,
				problems:  tt.fields.problems,
			}
			gm.displayRule()
			if gotWriter := tt.fields.w.(*bytes.Buffer).String(); gotWriter != tt.wantMessage {
				t.Errorf("Message output by displayRule() = %v, want %v", gotWriter, tt.wantMessage)
			}
		})
	}
}

func Test_gameMaster_game(t *testing.T) {
	setup(t)
	type fields struct {
		r         io.Reader
		w         io.Writer
		timeLimit time.Duration
		problems  []string
	}
	tests := []struct {
		name                 string
		fields               fields
		wantMessage          string
		wantAnswerNum        int
		wantCorrectAnswerNum int
	}{
		{
			name: "全問正解",
			fields: fields{
				r:         allCorrectReader,
				w:         &bytes.Buffer{},
				timeLimit: 10,
				problems:  []string{"apple", "bake", "cup", "dog", "egg", "fight", "green", "hoge", "idea", "japan"},
			},
			wantMessage: "1問目: apple\n>正解！ 現在の正答率:1/1\n" +
				"2問目: bake\n>正解！ 現在の正答率:2/2\n" +
				"3問目: cup\n>正解！ 現在の正答率:3/3\n" +
				"4問目: dog\n>正解！ 現在の正答率:4/4\n" +
				"5問目: egg\n>正解！ 現在の正答率:5/5\n" +
				"6問目: fight\n>正解！ 現在の正答率:6/6\n" +
				"7問目: green\n>正解！ 現在の正答率:7/7\n" +
				"8問目: hoge\n>正解！ 現在の正答率:8/8\n" +
				"9問目: idea\n>正解！ 現在の正答率:9/9\n" +
				"10問目: japan\n>正解！ 現在の正答率:10/10\n",
			wantAnswerNum:        10,
			wantCorrectAnswerNum: 10,
		},
		{
			name: "半分正解",
			fields: fields{
				r:         halfCorrectReader,
				w:         &bytes.Buffer{},
				timeLimit: 10,
				problems:  []string{"apple", "bake", "cup", "dog", "egg", "fight", "green", "hoge", "idea", "japan"},
			},
			wantMessage: "1問目: apple\n>正解！ 現在の正答率:1/1\n" +
				"2問目: bake\n>不正解... 現在の正答率:1/2\n" +
				"3問目: cup\n>不正解... 現在の正答率:1/3\n" +
				"4問目: dog\n>正解！ 現在の正答率:2/4\n" +
				"5問目: egg\n>正解！ 現在の正答率:3/5\n" +
				"6問目: fight\n>正解！ 現在の正答率:4/6\n" +
				"7問目: green\n>不正解... 現在の正答率:4/7\n" +
				"8問目: hoge\n>不正解... 現在の正答率:4/8\n" +
				"9問目: idea\n>正解！ 現在の正答率:5/9\n" +
				"10問目: japan\n>不正解... 現在の正答率:5/10\n",
			wantAnswerNum:        10,
			wantCorrectAnswerNum: 5,
		},
		{
			name: "読み込み完了",
			fields: fields{
				r:         timeUpReader,
				w:         &bytes.Buffer{},
				timeLimit: 10,
				problems:  []string{"apple", "bake", "cup", "dog", "egg", "fight", "green", "hoge", "idea", "japan"},
			},
			wantMessage: "1問目: apple\n>正解！ 現在の正答率:1/1\n" +
				"2問目: bake\n>正解！ 現在の正答率:2/2\n" +
				"3問目: cup\n>正解！ 現在の正答率:3/3\n" +
				"4問目: dog\n>不正解... 現在の正答率:3/4\n" +
				"5問目: egg\n>正解！ 現在の正答率:4/5\n" +
				"6問目: fight\n>正解！ 現在の正答率:5/6\n" +
				"7問目: green\n>正解！ 現在の正答率:6/7\n" +
				"8問目: hoge\n>",
			wantAnswerNum:        7,
			wantCorrectAnswerNum: 6,
		},
		{
			name: "タイムアップ",
			fields: fields{
				r:         timeUpReader,
				w:         &bytes.Buffer{},
				timeLimit: 0,
				problems:  []string{"apple", "bake", "cup", "dog", "egg", "fight", "green", "hoge", "idea", "japan"},
			},
			wantMessage:          "1問目: apple\n>",
			wantAnswerNum:        0,
			wantCorrectAnswerNum: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gm := &gameMaster{
				r:         tt.fields.r,
				w:         tt.fields.w,
				timeLimit: tt.fields.timeLimit,
				problems:  tt.fields.problems,
			}
			gm.game()

			if gotWriter := tt.fields.w.(*bytes.Buffer).String(); gotWriter != tt.wantMessage {
				t.Errorf("Message output by game() = %v, want %v", gotWriter, tt.wantMessage)
			}
			if got := gm.answerNum; got != tt.wantAnswerNum {
				t.Errorf("gm.answerNum = %v, want %v", got, tt.wantAnswerNum)
			}
			if got := gm.correctAnswerNum; got != tt.wantCorrectAnswerNum {
				t.Errorf("gm.correctAnswerNum = %v, want %v", got, tt.wantCorrectAnswerNum)
			}
		})
	}
}

func Test_gameMaster_displayResult(t *testing.T) {
	setup(t)
	type fields struct {
		r                io.Reader
		w                io.Writer
		timeLimit        time.Duration
		problems         []string
		answerNum        int
		correctAnswerNum int
	}
	tests := []struct {
		name        string
		fields      fields
		wantMessage string
	}{
		{
			name: "全問回答",
			fields: fields{
				r:                allCorrectReader,
				w:                &bytes.Buffer{},
				timeLimit:        10,
				problems:         []string{"apple", "bake", "cup", "dog", "egg", "fight", "green", "hoge", "idea", "japan"},
				answerNum:        10,
				correctAnswerNum: 9,
			},
			wantMessage: "\n全問回答しました！\n***\n10問中9問正解\n***\nお疲れ様でした！\n",
		},
		{
			name: "タイムアップ",
			fields: fields{
				r:                allCorrectReader,
				w:                &bytes.Buffer{},
				timeLimit:        10,
				problems:         []string{"apple", "bake", "cup", "dog", "egg", "fight", "green", "hoge", "idea", "japan"},
				answerNum:        8,
				correctAnswerNum: 3,
			},
			wantMessage: "\nタイムアップ！\n***\n8問中3問正解\n***\nお疲れ様でした！\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gm := &gameMaster{
				r:                tt.fields.r,
				w:                tt.fields.w,
				timeLimit:        tt.fields.timeLimit,
				problems:         tt.fields.problems,
				answerNum:        tt.fields.answerNum,
				correctAnswerNum: tt.fields.correctAnswerNum,
			}
			gm.displayResult()
			if gotWriter := tt.fields.w.(*bytes.Buffer).String(); gotWriter != tt.wantMessage {
				t.Errorf("Message output by displayResult() = %v, want %v", gotWriter, tt.wantMessage)
			}
		})
	}
}

func Test_gameMaster_input(t *testing.T) {
	setup(t)
	type fields struct {
		r io.Reader
	}
	tests := []struct {
		name              string
		fields            fields
		wantReceivedItems []string
	}{
		{
			name: "入力チャネルの取得",
			fields: fields{
				r: allCorrectReader,
			},
			wantReceivedItems: []string{"apple", "bake", "cup", "dog", "egg", "fight", "green", "hoge", "idea", "japan"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gm := &gameMaster{
				r: tt.fields.r,
			}
			got := gm.input()
			var i int
			for m := range got {
				if m != tt.wantReceivedItems[i] {
					t.Errorf("gameMaster.input() = %v, want %v", m, tt.wantReceivedItems[i])
				}
				i++
			}
		})
	}
}
