## NBT

```nbf
byte // signed
int16 // signed big endian
int32
int64

type := byte

float32
float64: 8 bytes / 64 bits, signed, big endian, IEEE 754-2008, binary64

string := int16 utf8content .

tree := tagType string tagContent .

tag := 0 ; TAG_End fin de TAG_Compound ou TAG_List vide
	| 1 string byte           ; TAG_Byte
	| 2 string int16          ; TAG_Short
	| 3 string int32          ; TAG_Int
	| 4 string int64          ; TAG_Long
	| 5 string float32        ; TAG_Float
	| 6 string float64        ; TAG_Double
	| 8 string content:string ; TAG_String

	|  7 string size:int32 size*byte  ; TAG_Byte_Array
	| 11 string size:int32 size*int32 ; TAG_Int_Array
	| 12 string size:int32 size*int64 ; TAG_Long_Array

	| 9 string type:byte size:int32 size*element ; TAG_List

	| 10 string {tag} tag_end ; TAG_Compound
	.
	// tag_type + string + value
```
