package imageprocessing

import (
	"image"
	"reflect"
	"testing"
)

func Test_processorPipeline_AddAction(t *testing.T) {
	type fields struct {
		imageProcesses []ImageAction
	}
	type args struct {
		processor ImageAction
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantPipeline *processorPipeline
	}{
		{
			name: "adding nil action does nothing",
			fields: fields{
				imageProcesses: []ImageAction{},
			},
			args: args{
				processor: nil,
			},
			wantPipeline: &processorPipeline{imageProcesses: []ImageAction{}},
		},
		{
			name: "adding valid action appends",
			fields: fields{
				imageProcesses: []ImageAction{},
			},
			args: args{
				processor: BuildMockProcessorPipeline(nil),
			},
			wantPipeline: &processorPipeline{
				imageProcesses: []ImageAction{BuildMockProcessorPipeline(nil)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &processorPipeline{
				imageProcesses: tt.fields.imageProcesses,
			}
			got.AddAction(tt.args.processor)
			if !reflect.DeepEqual(got, tt.wantPipeline) {
				t.Errorf("AddAction() got = %v, want %v", got, tt.wantPipeline)
			}
		})
	}
}

func Test_processorPipeline_Transform(t *testing.T) {
	type fields struct {
		imageProcesses []ImageAction
	}
	type args struct {
		image image.Image
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    image.Image
		wantErr bool
	}{
		{
			name:    "nil image results in error",
			fields:  fields{},
			args:    args{},
			wantErr: true,
		},
		{
			name: "failed transform results in error",
			fields: fields{
				imageProcesses: []ImageAction{
					BuildMockImageAction(func(mock *ImageActionMock) {
						mock.TransformFunc = func(in1 image.Image) (i image.Image, e error) {
							return nil, testError
						}
					}),
				},
			},
			args: args{
				image: BuildMockImage(),
			},
			wantErr: true,
		},
		{
			name: "successful transform completes",
			fields: fields{
				imageProcesses: []ImageAction{
					BuildMockImageAction(func(mock *ImageActionMock) {
						mock.TransformFunc = func(in1 image.Image) (i image.Image, e error) {
							return in1, nil
						}
					}),
				},
			},
			args: args{
				image: BuildMockImage(),
			},
			want: BuildMockImage(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := processorPipeline{
				imageProcesses: tt.fields.imageProcesses,
			}
			got, err := p.Transform(tt.args.image)
			if (err != nil) != tt.wantErr {
				t.Errorf("Transform() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transform() got = %v, want %v", got, tt.want)
			}
		})
	}
}
