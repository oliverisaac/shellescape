package shellescape

var EscapeCharList = map[byte]bool{
	0x00: true, // ''
	0x01: true, // $'\001'
	0x02: true, // $'\002'
	0x03: true, // $'\003'
	0x04: true, // $'\004'
	0x05: true, // $'\005'
	0x06: true, // $'\006'
	0x07: true, // $'\a'
	0x08: true, // $'\b'
	0x09: true, // $'\t'
	0x0A: true, // $'\n'
	0x0B: true, // $'\v'
	0x0C: true, // $'\f'
	0x0D: true, // $'\r'
	0x0E: true, // $'\016'
	0x0F: true, // $'\017'
	0x10: true, // $'\020'
	0x11: true, // $'\021'
	0x12: true, // $'\022'
	0x13: true, // $'\023'
	0x14: true, // $'\024'
	0x15: true, // $'\025'
	0x16: true, // $'\026'
	0x17: true, // $'\027'
	0x18: true, // $'\030'
	0x19: true, // $'\031'
	0x1A: true, // $'\032'
	0x1B: true, // $'\E'
	0x1C: true, // $'\034'
	0x1D: true, // $'\035'
	0x1E: true, // $'\036'
	0x1F: true, // $'\037'
	0x20: true, // \
	0x21: true, // \!
	0x22: true, // \"
	0x23: true, // \#
	0x24: true, // \$
	0x26: true, // \&
	0x27: true, // \'
	0x28: true, // \(
	0x29: true, // \)
	0x2A: true, // \*
	0x2C: true, // \,
	0x3B: true, // \;
	0x3C: true, // \<
	0x3E: true, // \>
	0x3F: true, // \?
	0x5B: true, // \[
	0x5C: true, // \\
	0x5D: true, // \]
	0x5E: true, // \^
	0x60: true, // \`
	0x7B: true, // \{
	0x7C: true, // \|
	0x7D: true, // \}
	0x7E: true, // \~
	0x7F: true, // $'\177'
}

func Escape(in string) (escaped string) {
	out := make([]byte, 0, len([]byte(in)))
	for _, ch := range []byte(in) {
		if _, ok := EscapeCharList[ch]; ok {
			out = append(out, 0x5C)
		}
		out = append(out, ch)
	}
	return string(out)
}

func EscapeArgs(args []string) (escaped []string) {
	for i, arg := range args {
		args[i] = Escape(arg)
	}
	return args
}
