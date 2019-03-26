package validator

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"net"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

var timeType = reflect.TypeOf(time.Time{})

func (t *KValidator) IsURLEncoded() bool {
	return uRLEncodedRegex.MatchString(t.data.String())
}

func (t *KValidator) IsHTMLEncoded() bool {
	return hTMLEncodedRegex.MatchString(t.data.String())
}

func (t *KValidator) IsHTML() bool {
	return hTMLRegex.MatchString(t.data.String())
}

// IsMAC is the validation function for validating if the field's value is a valid MAC address.
func (t *KValidator) IsMAC() bool {

	_, err := net.ParseMAC(t.data.String())

	return err == nil
}

// IsCIDRv4 is the validation function for validating if the field's value is a valid v4 CIDR address.
func (t *KValidator) IsCIDRv4() bool {

	ip, _, err := net.ParseCIDR(t.data.String())

	return err == nil && ip.To4() != nil
}

// IsCIDRv6 is the validation function for validating if the field's value is a valid v6 CIDR address.
func (t *KValidator) IsCIDRv6() bool {

	ip, _, err := net.ParseCIDR(t.data.String())

	return err == nil && ip.To4() == nil
}

// IsCIDR is the validation function for validating if the field's value is a valid v4 or v6 CIDR address.
func (t *KValidator) IsCIDR() bool {

	_, _, err := net.ParseCIDR(t.data.String())

	return err == nil
}

// IsIPv4 is the validation function for validating if a value is a valid v4 IP address.
func (t *KValidator) IsIPv4() bool {

	ip := net.ParseIP(t.data.String())

	return ip != nil && ip.To4() != nil
}

// IsIPv6 is the validation function for validating if the field's value is a valid v6 IP address.
func (t *KValidator) IsIPv6() bool {

	ip := net.ParseIP(t.data.String())

	return ip != nil && ip.To4() == nil
}

// IsIP is the validation function for validating if the field's value is a valid v4 or v6 IP address.
func (t *KValidator) IsIP() bool {

	ip := net.ParseIP(t.data.String())

	return ip != nil
}

// IsSSN is the validation function for validating if the field's value is a valid SSN.
func (t *KValidator) IsSSN() bool {

	if t.data.Len() != 11 {
		return false
	}

	return sSNRegex.MatchString(t.data.String())
}

// IsLongitude is the validation function for validating if the field's value is a valid longitude coordinate.
func (t *KValidator) isLongitude() bool {
	field := t.data
	var v string
	switch field.Kind() {
	case reflect.String:
		v = field.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v = strconv.FormatInt(field.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v = strconv.FormatUint(field.Uint(), 10)
	case reflect.Float32:
		v = strconv.FormatFloat(field.Float(), 'f', -1, 32)
	case reflect.Float64:
		v = strconv.FormatFloat(field.Float(), 'f', -1, 64)
	default:
		panic(fmt.Sprintf("Bad field type %T", field.Interface()))
	}

	return longitudeRegex.MatchString(v)
}

// IsLatitude is the validation function for validating if the field's value is a valid latitude coordinate.
func (t *KValidator) isLatitude() bool {
	field := t.data

	var v string
	switch field.Kind() {
	case reflect.String:
		v = field.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v = strconv.FormatInt(field.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v = strconv.FormatUint(field.Uint(), 10)
	case reflect.Float32:
		v = strconv.FormatFloat(field.Float(), 'f', -1, 32)
	case reflect.Float64:
		v = strconv.FormatFloat(field.Float(), 'f', -1, 64)
	default:
		panic(fmt.Sprintf("Bad field type %T", field.Interface()))
	}

	return latitudeRegex.MatchString(v)
}

// IsDataURI is the validation function for validating if the field's value is a valid data URI.
func (t *KValidator) isDataURI() bool {

	uri := strings.SplitN(t.data.String(), ",", 2)

	if len(uri) != 2 {
		return false
	}

	if !dataURIRegex.MatchString(uri[0]) {
		return false
	}

	return base64Regex.MatchString(uri[1])
}

// HasMultiByteCharacter is the validation function for validating if the field's value has a multi byte character.
func (t *KValidator) HasMultiByteCharacter() bool {

	field := t.data

	if field.Len() == 0 {
		return true
	}

	return multibyteRegex.MatchString(field.String())
}

// IsPrintableASCII is the validation function for validating if the field's value is a valid printable ASCII character.
func (t *KValidator) IsPrintableASCII() bool {
	return printableASCIIRegex.MatchString(t.data.String())
}

// IsASCII is the validation function for validating if the field's value is a valid ASCII character.
func (t *KValidator) IsASCII() bool {
	return aSCIIRegex.MatchString(t.data.String())
}

// IsUUID5 is the validation function for validating if the field's value is a valid v5 UUID.
func (t *KValidator) IsUUID5() bool {
	return uUID5Regex.MatchString(t.data.String())
}

// IsUUID4 is the validation function for validating if the field's value is a valid v4 UUID.
func (t *KValidator) IsUUID4() bool {
	return uUID4Regex.MatchString(t.data.String())
}

// IsUUID3 is the validation function for validating if the field's value is a valid v3 UUID.
func (t *KValidator) IsUUID3() bool {
	return uUID3Regex.MatchString(t.data.String())
}

// IsUUID is the validation function for validating if the field's value is a valid UUID of any version.
func (t *KValidator) IsUUID() bool {
	return uUIDRegex.MatchString(t.data.String())
}

// IsUUID5RFC4122 is the validation function for validating if the field's value is a valid RFC4122 v5 UUID.
func (t *KValidator) IsUUID5RFC4122() bool {
	return uUID5RFC4122Regex.MatchString(t.data.String())
}

// IsUUID4RFC4122 is the validation function for validating if the field's value is a valid RFC4122 v4 UUID.
func (t *KValidator) IsUUID4RFC4122() bool {
	return uUID4RFC4122Regex.MatchString(t.data.String())
}

// IsUUID3RFC4122 is the validation function for validating if the field's value is a valid RFC4122 v3 UUID.
func (t *KValidator) IsUUID3RFC4122() bool {
	return uUID3RFC4122Regex.MatchString(t.data.String())
}

// IsUUIDRFC4122 is the validation function for validating if the field's value is a valid RFC4122 UUID of any version.
func (t *KValidator) IsUUIDRFC4122() bool {
	return uUIDRFC4122Regex.MatchString(t.data.String())
}

// IsISBN is the validation function for validating if the field's value is a valid v10 or v13 ISBN.
func (t *KValidator) isISBN() bool {
	return t.IsISBN10() || t.IsISBN13()
}

// IsISBN13 is the validation function for validating if the field's value is a valid v13 ISBN.
func (t *KValidator) IsISBN13() bool {

	s := strings.Replace(strings.Replace(t.data.String(), "-", "", 4), " ", "", 4)

	if !iSBN13Regex.MatchString(s) {
		return false
	}

	var checksum int32
	var i int32

	factor := []int32{1, 3}

	for i = 0; i < 12; i++ {
		checksum += factor[i%2] * int32(s[i]-'0')
	}

	return (int32(s[12] - '0'))-((10-(checksum%10))%10) == 0
}

// IsISBN10 is the validation function for validating if the field's value is a valid v10 ISBN.
func (t *KValidator) IsISBN10() bool {

	s := strings.Replace(strings.Replace(t.data.String(), "-", "", 3), " ", "", 3)

	if !iSBN10Regex.MatchString(s) {
		return false
	}

	var checksum int32
	var i int32

	for i = 0; i < 9; i++ {
		checksum += (i + 1) * int32(s[i]-'0')
	}

	if s[9] == 'X' {
		checksum += 10 * 10
	} else {
		checksum += 10 * int32(s[9]-'0')
	}

	return checksum%11 == 0
}

// IsEthereumAddress is the validation function for validating if the field's value is a valid ethereum address based currently only on the format
func (t *KValidator) IsEthereumAddress() bool {
	address := t.data.String()

	if !ethAddressRegex.MatchString(address) {
		return false
	}

	if ethaddressRegexUpper.MatchString(address) || ethAddressRegexLower.MatchString(address) {
		return true
	}

	// checksum validation is blocked by https://github.com/golang/crypto/pull/28

	return true
}

// IsBitcoinAddress is the validation function for validating if the field's value is a valid btc address
func (t *KValidator) IsBitcoinAddress() bool {
	address := t.data.String()

	if !btcAddressRegex.MatchString(address) {
		return false
	}

	alphabet := []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

	decode := [25]byte{}

	for _, n := range []byte(address) {
		d := bytes.IndexByte(alphabet, n)

		for i := 24; i >= 0; i-- {
			d += 58 * int(decode[i])
			decode[i] = byte(d % 256)
			d /= 256
		}
	}

	h := sha256.New()
	_, _ = h.Write(decode[:21])
	d := h.Sum([]byte{})
	h = sha256.New()
	_, _ = h.Write(d)

	validchecksum := [4]byte{}
	computedchecksum := [4]byte{}

	copy(computedchecksum[:], h.Sum(d[:0]))
	copy(validchecksum[:], decode[21:])

	return validchecksum == computedchecksum
}

// IsBitcoinBech32Address is the validation function for validating if the field's value is a valid bech32 btc address
func (t *KValidator) IsBitcoinBech32Address() bool {
	address := t.data.String()

	if !btcLowerAddressRegexBech32.MatchString(address) && !btcUpperAddressRegexBech32.MatchString(address) {
		return false
	}

	am := len(address) % 8

	if am == 0 || am == 3 || am == 5 {
		return false
	}

	address = strings.ToLower(address)

	alphabet := "qpzry9x8gf2tvdw0s3jn54khce6mua7l"

	hr := []int{3, 3, 0, 2, 3} // the human readable part will always be bc
	addr := address[3:]
	dp := make([]int, 0, len(addr))

	for _, c := range addr {
		dp = append(dp, strings.IndexRune(alphabet, c))
	}

	ver := dp[0]

	if ver < 0 || ver > 16 {
		return false
	}

	if ver == 0 {
		if len(address) != 42 && len(address) != 62 {
			return false
		}
	}

	values := append(hr, dp...)

	GEN := []int{0x3b6a57b2, 0x26508e6d, 0x1ea119fa, 0x3d4233dd, 0x2a1462b3}

	p := 1

	for _, v := range values {
		b := p >> 25
		p = (p&0x1ffffff)<<5 ^ v

		for i := 0; i < 5; i++ {
			if (b>>uint(i))&1 == 1 {
				p ^= GEN[i]
			}
		}
	}

	if p != 1 {
		return false
	}

	b := uint(0)
	acc := 0
	mv := (1 << 5) - 1
	var sw []int

	for _, v := range dp[1 : len(dp)-6] {
		acc = (acc << 5) | v
		b += 5
		for b >= 8 {
			b -= 8
			sw = append(sw, (acc>>b)&mv)
		}
	}

	if len(sw) < 2 || len(sw) > 40 {
		return false
	}

	return true
}

// ExcludesRune is the validation function for validating that the field's value does not contain the rune specified within the param.
func (t *KValidator) excludesRune() bool {
	return !t.ContainsRune()
}

// ExcludesAll is the validation function for validating that the field's value does not contain any of the characters specified within the param.
func (t *KValidator) ExcludesAll() bool {
	return !t.ContainsAny()
}

// Excludes is the validation function for validating that the field's value does not contain the text specified within the param.
func (t *KValidator) Excludes() bool {
	return !t.Contains()
}

// ContainsRune is the validation function for validating that the field's value contains the rune specified within the param.
func (t *KValidator) ContainsRune() bool {

	r, _ := utf8.DecodeRuneInString(t.data.String())

	return strings.ContainsRune(t.data.String(), r)
}

// ContainsAny is the validation function for validating that the field's value contains any of the characters specified within the param.
func (t *KValidator) ContainsAny() bool {
	return strings.ContainsAny(t.data.String(), t.data.String())
}

// Contains is the validation function for validating that the field's value contains the text specified within the param.
func (t *KValidator) Contains() bool {
	return strings.Contains(t.data.String(), t.data.String())
}

// FieldContains is the validation function for validating if the current field's value contains the field specified by the param's value.
func (t *KValidator) FieldContains(s string) bool {
	field := t.data
	return strings.Contains(field.String(), s)
}

// FieldExcludes is the validation function for validating if the current field's value excludes the field specified by the param's value.
func (t *KValidator) FieldExcludes(s string) bool {
	field := t.data
	return !strings.Contains(field.String(), s)
}

// IsNe is the validation function for validating that the field's value does not equal the provided param value.
func (t *KValidator) IsNe() bool {
	return !t.IsEq()
}

// IsEq is the validation function for validating if the current field's value is equal to the param's value.
func (t *KValidator) IsEq() bool {

	field := t.data
	param := field.String()

	switch field.Kind() {

	case reflect.String:
		return field.String() == param

	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)

		return int64(field.Len()) == p

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p := asInt(param)

		return field.Int() == p

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p := asUint(param)

		return field.Uint() == p

	case reflect.Float32, reflect.Float64:
		p := asFloat(param)

		return field.Float() == p
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

// IsBase64 is the validation function for validating if the current field's value is a valid base 64.
func (t *KValidator) IsBase64() bool {
	return base64Regex.MatchString(t.data.String())
}

// IsBase64URL is the validation function for validating if the current field's value is a valid base64 URL safe string.
func (t *KValidator) IsBase64URL() bool {
	return base64URLRegex.MatchString(t.data.String())
}

// IsURI is the validation function for validating if the current field's value is a valid URI.
func (t *KValidator) IsURI() bool {

	switch t.data.Kind() {

	case reflect.String:

		s := t.data.String()

		// checks needed as of Go 1.6 because of change https://github.com/golang/go/commit/617c93ce740c3c3cc28cdd1a0d712be183d0b328#diff-6c2d018290e298803c0c9419d8739885L195
		// emulate browser and strip the '#' suffix prior to validation. see issue-#237
		if i := strings.Index(s, "#"); i > -1 {
			s = s[:i]
		}

		if len(s) == 0 {
			return false
		}

		_, err := url.ParseRequestURI(s)

		return err == nil
	}

	panic(fmt.Sprintf("Bad field type %T", t.data.Interface()))
}

// IsURL is the validation function for validating if the current field's value is a valid URL.
func (t *KValidator) IsURL() bool {

	field := t.data

	switch field.Kind() {

	case reflect.String:

		var i int
		s := field.String()

		// checks needed as of Go 1.6 because of change https://github.com/golang/go/commit/617c93ce740c3c3cc28cdd1a0d712be183d0b328#diff-6c2d018290e298803c0c9419d8739885L195
		// emulate browser and strip the '#' suffix prior to validation. see issue-#237
		if i = strings.Index(s, "#"); i > -1 {
			s = s[:i]
		}

		if len(s) == 0 {
			return false
		}

		url, err := url.ParseRequestURI(s)

		if err != nil || url.Scheme == "" {
			return false
		}

		return err == nil
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

// IsFile is the validation function for validating if the current field's value is a valid file path.
func (t *KValidator) IsFile() bool {
	field := t.data

	switch field.Kind() {
	case reflect.String:
		fileInfo, err := os.Stat(field.String())
		if err != nil {
			return false
		}

		return !fileInfo.IsDir()
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

// IsEmail is the validation function for validating if the current field's value is a valid email address.
func (t *KValidator) IsEmail() bool {
	return emailRegex.MatchString(t.data.String())
}

// IsHSLA is the validation function for validating if the current field's value is a valid HSLA color.
func (t *KValidator) IsHSLA() bool {
	return hslaRegex.MatchString(t.data.String())
}

// IsHSL is the validation function for validating if the current field's value is a valid HSL color.
func (t *KValidator) IsHSL() bool {
	return hslRegex.MatchString(t.data.String())
}

// IsRGBA is the validation function for validating if the current field's value is a valid RGBA color.
func (t *KValidator) IsRGBA() bool {
	return rgbaRegex.MatchString(t.data.String())
}

// IsRGB is the validation function for validating if the current field's value is a valid RGB color.
func (t *KValidator) IsRGB() bool {
	return rgbRegex.MatchString(t.data.String())
}

// IsHEXColor is the validation function for validating if the current field's value is a valid HEX color.
func (t *KValidator) IsHEXColor() bool {
	return hexcolorRegex.MatchString(t.data.String())
}

// IsHexadecimal is the validation function for validating if the current field's value is a valid hexadecimal.
func (t *KValidator) IsHexadecimal() bool {
	return hexadecimalRegex.MatchString(t.data.String())
}

// HasValue is the validation function for validating if the current field's value is not the default static value.
func (t *KValidator) Required() bool {

	switch t.data.Kind() {
	case reflect.Slice, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Chan, reflect.Func:
		return !t.data.IsNil()
	default:
		if t.data.Interface() != nil {
			return true
		}

		return t.data.IsValid() && t.data.Interface() != reflect.Zero(t.data.Type()).Interface()
	}
}

// IsGte is the validation function for validating if the current field's value is greater than or equal to the param's value.
func (t *KValidator) IsGte() bool {

	field := t.data
	param := t.data.String()

	switch field.Kind() {

	case reflect.String:
		p := asInt(param)

		return int64(utf8.RuneCountInString(field.String())) >= p

	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)

		return int64(field.Len()) >= p

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p := asInt(param)

		return field.Int() >= p

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p := asUint(param)

		return field.Uint() >= p

	case reflect.Float32, reflect.Float64:
		p := asFloat(param)

		return field.Float() >= p

	case reflect.Struct:

		if field.Type() == timeType {

			now := time.Now().UTC()
			t := field.Interface().(time.Time)

			return t.After(now) || t.Equal(now)
		}
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

// IsGt is the validation function for validating if the current field's value is greater than the param's value.
func (t *KValidator) IsGt() bool {

	field := t.data
	param := t.data.String()

	switch field.Kind() {

	case reflect.String:
		p := asInt(param)

		return int64(utf8.RuneCountInString(field.String())) > p

	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)

		return int64(field.Len()) > p

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p := asInt(param)

		return field.Int() > p

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p := asUint(param)

		return field.Uint() > p

	case reflect.Float32, reflect.Float64:
		p := asFloat(param)

		return field.Float() > p
	case reflect.Struct:

		if field.Type() == timeType {

			return field.Interface().(time.Time).After(time.Now().UTC())
		}
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

// HasLengthOf is the validation function for validating if the current field's value is equal to the param's value.
func (t *KValidator) HasLengthOf() bool {

	field := t.data
	param := t.data.String()

	switch field.Kind() {

	case reflect.String:
		p := asInt(param)

		return int64(utf8.RuneCountInString(field.String())) == p

	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)

		return int64(field.Len()) == p

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p := asInt(param)

		return field.Int() == p

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p := asUint(param)

		return field.Uint() == p

	case reflect.Float32, reflect.Float64:
		p := asFloat(param)

		return field.Float() == p
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

// HasMinOf is the validation function for validating if the current field's value is greater than or equal to the param's value.
func (t *KValidator) HasMinOf() bool {
	return t.IsGte()
}

// IsLte is the validation function for validating if the current field's value is less than or equal to the param's value.
func (t *KValidator) IsLte() bool {

	field := t.data
	param := t.data.String()

	switch field.Kind() {

	case reflect.String:
		p := asInt(param)

		return int64(utf8.RuneCountInString(field.String())) <= p

	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)

		return int64(field.Len()) <= p

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p := asInt(param)

		return field.Int() <= p

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p := asUint(param)

		return field.Uint() <= p

	case reflect.Float32, reflect.Float64:
		p := asFloat(param)

		return field.Float() <= p

	case reflect.Struct:

		if field.Type() == timeType {

			now := time.Now().UTC()
			t := field.Interface().(time.Time)

			return t.Before(now) || t.Equal(now)
		}
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

// IsLt is the validation function for validating if the current field's value is less than the param's value.
func (t *KValidator) IsLt() bool {

	field := t.data
	param := t.data.String()

	switch field.Kind() {

	case reflect.String:
		p := asInt(param)

		return int64(utf8.RuneCountInString(field.String())) < p

	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)

		return int64(field.Len()) < p

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p := asInt(param)

		return field.Int() < p

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p := asUint(param)

		return field.Uint() < p

	case reflect.Float32, reflect.Float64:
		p := asFloat(param)

		return field.Float() < p

	case reflect.Struct:

		if field.Type() == timeType {

			return field.Interface().(time.Time).Before(time.Now().UTC())
		}
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

// HasMaxOf is the validation function for validating if the current field's value is less than or equal to the param's value.
func (t *KValidator) HasMaxOf() bool {
	return t.IsLte()
}

// IsTCP4AddrResolvable is the validation function for validating if the field's value is a resolvable tcp4 address.
func (t *KValidator) IsTCP4AddrResolvable() bool {

	if !t.IsIP4Addr() {
		return false
	}

	_, err := net.ResolveTCPAddr("tcp4", t.data.String())
	return err == nil
}

// IsTCP6AddrResolvable is the validation function for validating if the field's value is a resolvable tcp6 address.
func (t *KValidator) IsTCP6AddrResolvable() bool {

	if !t.IsIP6Addr() {
		return false
	}

	_, err := net.ResolveTCPAddr("tcp6", t.data.String())

	return err == nil
}

// IsTCPAddrResolvable is the validation function for validating if the field's value is a resolvable tcp address.
func (t *KValidator) IsTCPAddrResolvable() bool {

	if !t.IsIP4Addr() && !t.IsIP6Addr() {
		return false
	}

	_, err := net.ResolveTCPAddr("tcp", t.data.String())

	return err == nil
}

// IsUDP4AddrResolvable is the validation function for validating if the field's value is a resolvable udp4 address.
func (t *KValidator) IsUDP4AddrResolvable() bool {

	if !t.IsIP4Addr() {
		return false
	}

	_, err := net.ResolveUDPAddr("udp4", t.data.String())

	return err == nil
}

// IsUDP6AddrResolvable is the validation function for validating if the field's value is a resolvable udp6 address.
func (t *KValidator) IsUDP6AddrResolvable() bool {

	if !t.IsIP6Addr() {
		return false
	}

	_, err := net.ResolveUDPAddr("udp6", t.data.String())

	return err == nil
}

// IsUDPAddrResolvable is the validation function for validating if the field's value is a resolvable udp address.
func (t *KValidator) IsUDPAddrResolvable() bool {

	if !t.IsIP4Addr() && !t.IsIP6Addr() {
		return false
	}

	_, err := net.ResolveUDPAddr("udp", t.data.String())

	return err == nil
}

// IsIP4AddrResolvable is the validation function for validating if the field's value is a resolvable ip4 address.
func (t *KValidator) IsIP4AddrResolvable() bool {

	if !t.IsIPv4() {
		return false
	}

	_, err := net.ResolveIPAddr("ip4", t.data.String())

	return err == nil
}

// IsIP6AddrResolvable is the validation function for validating if the field's value is a resolvable ip6 address.
func (t *KValidator) IsIP6AddrResolvable() bool {

	if !t.IsIPv6() {
		return false
	}

	_, err := net.ResolveIPAddr("ip6", t.data.String())

	return err == nil
}

// IsIPAddrResolvable is the validation function for validating if the field's value is a resolvable ip address.
func (t *KValidator) IsIPAddrResolvable() bool {

	if !t.IsIP() {
		return false
	}

	_, err := net.ResolveIPAddr("ip", t.data.String())

	return err == nil
}

// IsUnixAddrResolvable is the validation function for validating if the field's value is a resolvable unix address.
func (t *KValidator) IsUnixAddrResolvable() bool {

	_, err := net.ResolveUnixAddr("unix", t.data.String())

	return err == nil
}

func (t *KValidator) IsIP4Addr() bool {

	val := t.data.String()

	if idx := strings.LastIndex(val, ":"); idx != -1 {
		val = val[0:idx]
	}

	ip := net.ParseIP(val)

	return ip != nil && ip.To4() != nil
}

func (t *KValidator) IsIP6Addr() bool {

	val := t.data.String()

	if idx := strings.LastIndex(val, ":"); idx != -1 {
		if idx != 0 && val[idx-1:idx] == "]" {
			val = val[1 : idx-1]
		}
	}

	ip := net.ParseIP(val)

	return ip != nil && ip.To4() == nil
}

func (t *KValidator) IsHostnameRFC952() bool {
	return hostnameRegexRFC952.MatchString(t.data.String())
}

func (t *KValidator) IsHostnameRFC1123() bool {
	return hostnameRegexRFC1123.MatchString(t.data.String())
}

func (t *KValidator) IsFQDN() bool {
	val := t.data.String()

	if val == "" {
		return false
	}

	if val[len(val)-1] == '.' {
		val = val[0 : len(val)-1]
	}

	return strings.ContainsAny(val, ".") &&
		hostnameRegexRFC952.MatchString(val)
}

// IsDir is the validation function for validating if the current field's value is a valid directory.
func (t *KValidator) IsDir() bool {
	if t.data.Kind() == reflect.String {
		fileInfo, err := os.Stat(t.data.String())
		if err != nil {
			return false
		}

		return fileInfo.IsDir()
	}

	panic(fmt.Sprintf("Bad field type %T", t.data.Interface()))
}
