package main

import "testing"

func TestBook_NetPriceCalculate(t *testing.T) {
	type fields struct {
		Title           string
		Author          string
		Copies          int
		Series          bool
		SeriesId        int
		PriceCents      int
		DiscountPercent int
	}
	type args struct {
		originalPrice int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			fields: fields{
				DiscountPercent: 10,
			  },
			args: args{
				originalPrice: 900,
			},
			want: 810,
		  },
		  {
			fields: fields{
				DiscountPercent: 10,
			  },
			args: args{
				originalPrice: 100,
			},
			want: 90,
		  },
		  {
			fields: fields{
				DiscountPercent: 10,
			  },
			args: args{
				originalPrice: 0,
			},
			want: 0,
		  },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Book{
				Title:           tt.fields.Title,
				Author:          tt.fields.Author,
				Copies:          tt.fields.Copies,
				Series:          tt.fields.Series,
				SeriesId:        tt.fields.SeriesId,
				PriceCents:      tt.fields.PriceCents,
				DiscountPercent: tt.fields.DiscountPercent,
			}
			if got := b.NetPriceCalculate(tt.args.originalPrice); got != tt.want {
				t.Errorf("Book.NetPriceCalculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
