package converter

import (
	"reflect"
	"testing"
)

func TestNewConverter(t *testing.T) {
	type args struct {
		dst string
	}
	tests := []struct {
		name    string
		args    args
		want    IConverter
		wantErr bool
	}{
		{
			name:"pngのコンバーターを生成",
			args:args{
				dst:"png",
			},
			want:converter{&pngEncoder{}},
			wantErr:false,
		},
		{
			name:"jpgのコンバーターを生成",
			args:args{
				dst:"jpg",
			},
			want:converter{&jpgEncoder{}},
			wantErr:false,
		},
		{
			name:"gifのコンバーターを生成",
			args:args{
				dst:"gif",
			},
			want:converter{&gifEncoder{}},
			wantErr:false,
		},
		{
			name:"サポート外のコンバーターを生成",
			args:args{
				dst:"bmp",
			},
			want:nil,
			wantErr:true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewConverter(tt.args.dst)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConverter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConverter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_converter_Convert(t *testing.T) {
	type fields struct {
		Encoder Encoder
	}
	type args struct {
		path     string
		fileName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:"pngからjpgの変換",
			fields:fields{
				Encoder:&jpgEncoder{},
			},
			args:args{
				path:"../testdata/src/gopher.png",
				fileName:"../testdata/dst/gopher.jpg",
			},
			wantErr:false,
		},
		{
			name:"jpgからgifの変換",
			fields:fields{
				Encoder:&gifEncoder{},
			},
			args:args{
				path:"../testdata/src/gopher.jpg",
				fileName:"../testdata/dst/gopher.gif",
			},
			wantErr:false,
		},
		{
			name:"gifからpngの変換",
			fields:fields{
				Encoder:&pngEncoder{},
			},
			args:args{
				path:"../testdata/src/gopher.gif",
				fileName:"../testdata/dst/gopher.png",
			},
			wantErr:false,
		},
		{
			name:"存在しないファイルのパスを指定",
			fields:fields{
				Encoder:&pngEncoder{},
			},
			args:args{
				path:"../testdata/src/gopher.bmp",
				fileName:"../testdata/dst/gopher.png",
			},
			wantErr:true,
		},
		{
			name:"画像ファイル以外を指定",
			fields:fields{
				Encoder:&pngEncoder{},
			},
			args:args{
				path:"../testdata/src/dummy.txt",
				fileName:"../testdata/dst/gopher.png",
			},
			wantErr:true,
		},
		{
			name:"出力先に存在しないディレクトリを指定",
			fields:fields{
				Encoder:&pngEncoder{},
			},
			args:args{
				path:"../testdata/src/gopher.gif",
				fileName:"../testdata/dummy/gopher.png",
			},
			wantErr:true,
		},
		{
			name:"出力先に存在しないディレクトリを指定",
			fields:fields{
				Encoder:&pngEncoder{},
			},
			args:args{
				path:"../testdata/src/gopher.gif",
				fileName:"../testdata/dst/gopher.jpg",
			},
			wantErr:false,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := converter{
				Encoder: tt.fields.Encoder,
			}
			if err := c.Convert(tt.args.path, tt.args.fileName); (err != nil) != tt.wantErr {
				t.Errorf("converter.Convert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
