package parser

import (
	"reflect"
	"testing"
)

func Test_parseLinkAndTextFromHtml(t *testing.T) {
	type args struct {
		file string
	}
	file1 := "ex1.html"
	want1 := []Links{{"/other-page", "A link to another page"}}
	file2 := "ex2.html"
	want2 := []Links{{"https://www.twitter.com/joncalhoun", "Check me out on twitter"}, {"https://github.com/gophercises", "Gophercises is on Github!"}}
	file3 := "ex3.html"
	want3 := []Links{{"#", "Login"}, {"/lost", "Lost? Need help?"}, {"https://twitter.com/marcusolsson", "@marcusolsson"}}
	file4 := "ex4.html"
	want4 := []Links{{"/dog-cat", "dog cat"}}
	file5 := "ex5.html"
	want5 := []Links{{"/dog", ""}}
	file6 := "ex6.html"
	want6 := []Links{{"/dog", "text inside dog link"}}
	file7 := "ex7.html"
	want7 := []Links{{"/dog", "Something in a span Text not in a span Bold text!"}}
	file8 := "ex8.html"
	want8 := []Links{{"#", "Something here nested dog link"}}
	tests := []struct {
		name string
		args args
		want []Links
	}{
		{"test1", args{file1}, want1},
		{"test2", args{file2}, want2},
		{"test3", args{file3}, want3},
		{"test4", args{file4}, want4},
		{"test5", args{file5}, want5},
		{"test6", args{file6}, want6},
		{"test7", args{file7}, want7},
		{"test8", args{file8}, want8}, // why tc8 fails: https://www.w3.org/TR/html401/struct/links.html#h-12.2.2
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseLinkAndTextFromHtml(tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseLinkAndTextFromHtml() = %v, want %v", got, tt.want)
			}
		})
	}
}
