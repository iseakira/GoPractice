package hsd

import (
	"reflect"
	"testing"
)

func TestStringDistance(t *testing.T){
	tests := []struct{
		name string
		ihs string
		rhs string
		want int

	}{
		{name:"ihs is longer than rhs",ihs:"foo",rhs:"fo",want:-1},
		{name:"rhs is longer than ihs",ihs:"foo",rhs:"fo",want:-1},
		{name: "No diff", ihs: "foo", rhs: "foo", want: 0},
		{name: "1 diff", ihs: "foo", rhs: "foh", want: 1},
    {name: "2 diffs", ihs: "foo", rhs: "fhh", want: 2},
    {name: "3 diffs", ihs: "foo", rhs: "hhh", want: 3},
    {name: "multibyte", ihs: "あいう", rhs: "あいえ", want: 1},

	}

	for _,tc := range tests {
		got := StringDistance(tc.ihs,tc.rhs)
		if !reflect.DeepEqual(tc.want,got){
			t.Fatalf("%s:expected:%v,got:%v",tc.name,tc.want,got)
		}

	}
}