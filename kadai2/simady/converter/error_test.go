package converter

import "testing"

func TestImgconvError_Error(t *testing.T) {
	type fields struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:"独自エラーの生成",
			fields:fields{
				message:"error.",
			},
			want:"error.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := ImgconvError{
				Message: tt.fields.message,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("ImgconvError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
