// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package entities

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonB4ad3f7dDecodeQuanTokoinTestEntities(in *jlexer.Lexer, out *UserEntityList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(UserEntityList, 0, 1)
			} else {
				*out = UserEntityList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 UserEntity
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonB4ad3f7dEncodeQuanTokoinTestEntities(out *jwriter.Writer, in UserEntityList) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v UserEntityList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB4ad3f7dEncodeQuanTokoinTestEntities(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserEntityList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB4ad3f7dEncodeQuanTokoinTestEntities(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserEntityList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB4ad3f7dDecodeQuanTokoinTestEntities(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserEntityList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB4ad3f7dDecodeQuanTokoinTestEntities(l, v)
}
func easyjsonB4ad3f7dDecodeQuanTokoinTestEntities1(in *jlexer.Lexer, out *UserEntity) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "_id":
			out.Id = int(in.Int())
		case "url":
			out.Url = string(in.String())
		case "external_id":
			out.ExternalId = string(in.String())
		case "name":
			out.Name = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonB4ad3f7dEncodeQuanTokoinTestEntities1(out *jwriter.Writer, in UserEntity) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"_id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"url\":"
		out.RawString(prefix)
		out.String(string(in.Url))
	}
	{
		const prefix string = ",\"external_id\":"
		out.RawString(prefix)
		out.String(string(in.ExternalId))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserEntity) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB4ad3f7dEncodeQuanTokoinTestEntities1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserEntity) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB4ad3f7dEncodeQuanTokoinTestEntities1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserEntity) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB4ad3f7dDecodeQuanTokoinTestEntities1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserEntity) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB4ad3f7dDecodeQuanTokoinTestEntities1(l, v)
}