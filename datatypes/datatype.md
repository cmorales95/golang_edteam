`uint`
* uint8 unsigned 8-bit integers (0 to 255)
* uint16 unsigned 16-bit integers (0 to 65535)
* uint32 unsigned 32-bit integers (0 to 4294967295)
* uint64 unsigned 64-bit integers (0 to 18446744073709551615)

***solo para valores positivos, sino indicamos ningun unit
tomara por defecto el unit64***

`int`
* int8 signed 8-bit integers (-128 to 127)
* int16 signed 16-bit integers (-32768 to 32768)
* int32 signed 32-bit integers (-2147483648 to 2147483648)
* int64 signed 32-bit integers (-9223372036854775808 to 9223372036854775808)

`byte` // alias for uint8
 
 `rune` // alias for int32
        // represents a unicode code point
        
`float32`, `float64`