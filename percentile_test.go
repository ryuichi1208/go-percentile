package main

import (
	"reflect"
	"testing"
)

func Test_calc(t *testing.T) {
	type args struct {
		nums interface{}
		n    int
	}
	tests := []struct {
		name    string
		args    args
		wantNum interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNum, err := calc(tt.args.nums, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("calc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNum, tt.wantNum) {
				t.Errorf("calc() = %v, want %v", gotNum, tt.wantNum)
			}
		})
	}
}

func Test_percentileN(t *testing.T) {
	var t1, t2 []int
	var t3 []int64
	for i := 1; i < 101; i++ {
		t1 = append(t1, i)
	}
	for i := 1; i < 100001; i++ {
		t2 = append(t2, i)
	}
	for i := 1; i < 100001; i++ {
		t3 = append(t3, int64(i))
	}
	// for i := 1; i < 9223372036854775807; i++ {
	// 	t4 = append(t4, int64(i))
	// }
	type args struct {
		list interface{}
		n    int
	}
	tests := []struct {
		name     string
		args     args
		wantNums interface{}
		wantErr  bool
	}{
		{
			name:     "test1",
			args:     args{list: t1, n: 100},
			wantNums: 100,
			wantErr:  false,
		},
		{
			name:     "test1",
			args:     args{list: t1, n: 99},
			wantNums: 99,
			wantErr:  false,
		},
		{
			name:     "test3",
			args:     args{list: t2, n: 99},
			wantNums: 99000,
			wantErr:  false,
		},
		{
			name:     "test4",
			args:     args{list: t3, n: 100},
			wantNums: 100000,
			wantErr:  false,
		},
		// {
		// 	name:     "test5",
		// 	args:     args{list: t4, n: 100},
		// 	wantNums: 9223372036854775807,
		// 	wantErr:  false,
		// },
		// {
		// 	name:     "test6",
		// 	args:     args{list: t4, n: 1},
		// 	wantNums: 1,
		// 	wantErr:  false,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNums, err := percentileN(tt.args.list, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("percentileN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNums, tt.wantNums) {
				t.Errorf("percentileN() = %v, want %v", gotNums, tt.wantNums)
			}
		})
	}
}
