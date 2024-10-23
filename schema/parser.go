package schema

import (
	"strconv"
	"unsafe"

	. "github.com/zeromake/gencode/parser"
)

func MakeGrammar() *Grammar {
	Letter := Set("\\p{L}")
	Digit := Set("\\p{Nd}")
	Tick := Set("`")
	Any := Set("^`")

	WS := Ignore(Mult(0, 0, Set("\t\n\f\r ")))
	RWS := Ignore(Mult(1, 0, Set("\t\n\f\r ")))
	NL := Ignore(And(Mult(0, 0, Set("\t\f\r ")), Mult(1, 0, And(Lit("\n"), WS))))

	gIdentifier := And(Letter, Mult(0, 0, Or(Letter, Digit)))
	gIdentifier.Node(func(m Match) (Match, error) {
		return String(m), nil
	})

	gIntField := And(Optional(Tag("Var", Lit("v"))), Optional(Tag("Unsigned", Lit("u"))), Lit("int"), Optional(Tag("Bits", Or(Lit("8"), Lit("16"), Lit("32"), Lit("64")))))
	gIntField.Node(func(m Match) (Match, error) {
		bitsTag := GetTag(m, "Bits")
		var temp int
		var bits int64 = int64(unsafe.Sizeof(&temp)) * 8
		var err error
		var skipBits bool = false
		if bitsTag != nil {
			bits, err = strconv.ParseInt(String(bitsTag), 10, 64)
			if err != nil {
				return nil, err
			}
		} else {
			skipBits = true
		}
		signed := GetTag(m, "Unsigned") == nil
		if GetTag(m, "Var") == nil {
			return &IntType{
				Signed: signed,
				Bits:   int(bits),
				SkipBits: skipBits,
			}, nil
		}
		return &IntType{
			Signed: signed,
			Bits:   int(bits),
			VarInt: true,
			SkipBits: skipBits,
		}, nil
	})

	gFloatField := And(Lit("float"), Tag("Bits", Or(Lit("32"), Lit("64"))))
	gFloatField.Node(func(m Match) (Match, error) {
		bits, err := strconv.ParseInt(String(GetTag(m, "Bits")), 10, 64)
		if err != nil {
			return nil, err
		}
		return &FloatType{
			Bits: int(bits),
		}, nil
	})

	gByteField := And(Lit("byte"))
	gByteField.Node(func(m Match) (Match, error) {
		s := &ByteType{}
		return s, nil
	})

	gBoolField := And(Lit("bool"))
	gBoolField.Node(func(m Match) (Match, error) {
		s := &BoolType{}
		return s, nil
	})

	gStringField := And(Lit("string"))
	gStringField.Node(func(m Match) (Match, error) {
		s := &StringType{}
		return s, nil
	})

	gTimeField := And(Lit("time"))
	gTimeField.Node(func(m Match) (Match, error) {
		s := &TimeType{}
		return s, nil
	})

	gDeferField := And(gIdentifier)
	gDeferField.Node(func(m Match) (Match, error) {
		s := &DeferType{
			Defer: String(m),
		}
		return s, nil
	})

	gType := &Grammar{}

	gUnion := And(Lit("union"), Optional(And(RWS, Tag("Interface", gIdentifier))), WS, Lit("{"), WS,
		Mult(0, 0, And(WS, Tag("Defer", gType), NL)),
		Lit("}"),
	)
	gUnion.Node(func(m Match) (Match, error) {
		u := &UnionType{
			Interface: String(GetTag(m, "Interface")),
		}
		for _, v := range GetTags(m, "Defer") {
			u.Types = append(u.Types, v.(Type))
		}
		return u, nil
	})

	gSlice := And(Lit("[]"), Require(Tag("SubType", gType)))
	gSlice.Node(func(m Match) (Match, error) {
		return &SliceType{
			SubType: GetTag(m, "SubType").(Type),
		}, nil
	})

	gArray := And(Lit("["), Tag("Count", Mult(1, 0, Digit)), Lit("]"), Require(Tag("SubType", gType)))
	gArray.Node(func(m Match) (Match, error) {
		count, err := strconv.ParseUint(String(GetTag(m, "Count")), 10, 64)
		if err != nil {
			return nil, err
		}
		return &ArrayType{
			SubType: GetTag(m, "SubType").(Type),
			Count:   count,
		}, nil
	})

	gPointer := And(Lit("*"), Require(Tag("SubType", gType)))
	gPointer.Node(func(m Match) (Match, error) {
		return &PointerType{
			SubType: GetTag(m, "SubType").(Type),
		}, nil
	})

	gType.Set(Or(gSlice, gArray, gPointer, gIntField, gByteField, gBoolField, gStringField, gTimeField, gFloatField, gUnion, gDeferField))

	gStructTag := And(Tick, Mult(0, 0, Or(Any)), Tick)
	gStructTag.Node(func(m Match) (Match, error) {
		return String(m), nil
	})

	gField := And(Tag("Name", gIdentifier), Require(RWS, Tag("Type", gType)), Optional(And(RWS, Tag("Tag", gStructTag))), Require(NL))
	gField.Node(func(m Match) (Match, error) {
		// Optional struct tag
		var tg string
		if t, ok := GetTag(m, "Tag").(string); ok {
			tg = t
		}

		f := &Field{
			Name: GetTag(m, "Name").(string),
			Tag:  tg,
			Type: GetTag(m, "Type").(Type),
		}
		return TagMatch("Field", f), nil
	})

	gStruct := And(Lit("struct"), Require(RWS, Tag("Name", gIdentifier), Optional(And(RWS, Tag("Framed", Lit("framed")))), WS, Lit("{"), WS,
		Mult(0, 0, gField),
		Lit("}"), WS,
	))
	gStruct.Node(func(m Match) (Match, error) {
		s := &Struct{
			Name: GetTag(m, "Name").(string),
		}
		for _, v := range GetTags(m, "Field") {
			s.Fields = append(s.Fields, v.(*Field))
		}
		if GetTag(m, "Framed") != nil {
			s.Framed = true
		}
		return TagMatch("Struct", s), nil
	})

	gSchema := And(WS, Mult(0, 0, gStruct), WS)
	gSchema.Node(func(m Match) (Match, error) {
		s := &Schema{}
		for _, v := range GetTags(m, "Struct") {
			s.Structs = append(s.Structs, v.(*Struct))
		}
		return s, nil
	})
	return gSchema
}
