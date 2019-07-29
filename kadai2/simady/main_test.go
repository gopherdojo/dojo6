package main

import (
	"reflect"
	"testing"
)

func Test_getConversionTargetFiles(t *testing.T) {
	type args struct {
		dir       string
		targetExt string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "単体ディレクトリの検索",
			args: args{
				dir:       "testdata/root/sub",
				targetExt: "png",
			},
			want: []string{
				"testdata/root/sub/sub_dummy.png",
			},
			wantErr: false,
		},
		{
			name: "再帰的な検索",
			args: args{
				dir:       "testdata/root",
				targetExt: "jpg",
			},
			want: []string{
				"testdata/root/root_dummy.jpg",
				"testdata/root/sub/sub_dummy.jpg",
			},
			wantErr: false,
		},
		{
			name: "一致対象なしの検索",
			args: args{
				dir:       "testdata/root",
				targetExt: "bmp",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "存在しないディレクトリの検索",
			args: args{
				dir:       "notexists",
				targetExt: "png",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getConversionTargetFiles(tt.args.dir, tt.args.targetExt)
			if (err != nil) != tt.wantErr {
				t.Errorf("getConversionTargetFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConversionTargetFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSupported(t *testing.T) {
	type args struct {
		ext string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "pngはサポート対象の拡張子",
			args: args{
				ext: "png",
			},
			want: true,
		},
		{
			name: "jpgはサポート対象の拡張子",
			args: args{
				ext: "jpg",
			},
			want: true,
		},
		{
			name: "gifはサポート対象の拡張子",
			args: args{
				ext: "gif",
			},
			want: true,
		},
		{
			name: "bmpはサポート対象外の拡張子",
			args: args{
				ext: "bmp",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSupported(tt.args.ext); got != tt.want {
				t.Errorf("isSupported() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "main関数の実行",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_imgConvMain(t *testing.T) {
	type args struct {
		in  string
		out string
		src string
		dst string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "pngからjpgに変換",
			args: args{
				in:  "./testdata/src",
				out: "./testdata/dst",
				src: "png",
				dst: "jpg",
			},
			wantErr: false,
		},
		{
			name: "jpgからgifに変換",
			args: args{
				in:  "./testdata/src",
				out: "./testdata/dst",
				src: "jpg",
				dst: "gif",
			},
			wantErr: false,
		},
		{
			name: "gifからpngに変換",
			args: args{
				in:  "./testdata/src",
				out: "./testdata/dst",
				src: "gif",
				dst: "png",
			},
			wantErr: false,
		},
		{
			name: "変換元がサポート対象外",
			args: args{
				in:  "./testdata/src",
				out: "./testdata/dst",
				src: "bmp",
				dst: "png",
			},
			wantErr: false,
		},
		{
			name: "変換先がサポート対象外",
			args: args{
				in:  "./testdata/src",
				out: "./testdata/dst",
				src: "png",
				dst: "bmp",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := imgConvMain(tt.args.in, tt.args.out, tt.args.src, tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("imgConvMain() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
