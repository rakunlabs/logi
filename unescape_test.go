package logi

import "testing"

func Test_unescapeJSONInLine(t *testing.T) {
	tests := []struct {
		name string
		line []byte
		want []byte
	}{
		{
			name: "empty line",
			line: []byte{},
			want: []byte{},
		},
		{
			name: "plain text without JSON",
			line: []byte("Hello World"),
			want: []byte("Hello World"),
		},
		{
			name: "escaped JSON object",
			line: []byte(`"{\"key\":\"value\"}"`),
			want: []byte(`{"key":"value"}`),
		},
		{
			name: "escaped JSON array",
			line: []byte(`"[\"item1\",\"item2\"]"`),
			want: []byte(`["item1","item2"]`),
		},
		{
			name: "text with ANSI color codes",
			line: []byte("\x1b[31mRed Text\x1b[0m"),
			want: []byte("\x1b[31mRed Text\x1b[0m"),
		},
		{
			name: "ANSI codes with escaped JSON",
			line: []byte("\x1b[32m\"{\\\"status\\\":\\\"ok\\\"}\"\x1b[0m"),
			want: []byte("\x1b[32m{\"status\":\"ok\"}\x1b[0m"),
		},
		{
			name: "complex escaped JSON with nested objects",
			line: []byte(`"{\"user\":{\"name\":\"John\",\"age\":30}}"`),
			want: []byte(`{"user":{"name":"John","age":30}}`),
		},
		{
			name: "quoted string without escaped content",
			line: []byte(`"simple string"`),
			want: []byte(`"simple string"`),
		},
		{
			name: "JSON-like but not escaped",
			line: []byte(`{"key":"value"}`),
			want: []byte(`{"key":"value"}`),
		},
		{
			name: "mixed content with ANSI and escaped JSON",
			line: []byte("prefix \x1b[33m\"{\\\"data\\\":[1,2,3]}\"\x1b[0m suffix"),
			want: []byte("prefix \x1b[33m{\"data\":[1,2,3]}\x1b[0m suffix"),
		},
		{
			name: "escaped JSON with special characters",
			line: []byte(`"{\"message\":\"Hello\\nWorld\\t!\"}"`),
			want: []byte("{\"message\":\"Hello\\nWorld\\t!\"}"),
		},
		{
			name: "multiple ANSI codes",
			line: []byte("\x1b[1m\x1b[31mBold Red\x1b[0m"),
			want: []byte("\x1b[1m\x1b[31mBold Red\x1b[0m"),
		},
		{
			name: "quote at the end",
			line: []byte(`end with "`),
			want: []byte(`end with "`),
		},
		{
			name: "escaped backslash in JSON",
			line: []byte(`"{\"path\":\"C:\\\\Users\"}"`),
			want: []byte(`{"path":"C:\\Users"}`),
		},
		{
			name: "incomplete ANSI sequence",
			line: []byte("\x1b[31"),
			want: []byte("\x1b[31"),
		},
		{
			name: "escaped array with numbers",
			line: []byte(`"[1,2,3,4,5]"`),
			want: []byte(`"[1,2,3,4,5]"`),
		},
		{
			name: "complex log line with timestamp and escaped JSON",
			line: []byte(`2023-10-01 INFO "{\"request\":{\"method\":\"GET\"}}"`),
			want: []byte(`2023-10-01 INFO {"request":{"method":"GET"}}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UnescapeJSONInLine(tt.line)
			if string(got) != string(tt.want) {
				t.Errorf("unescapeJSONInLine() = %q, want %q", string(got), string(tt.want))
			}
		})
	}
}
