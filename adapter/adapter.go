package adapter

/*
#cgo CFLAGS: -I"${SRCDIR}/../ext/drafter/src/" -I"${SRCDIR}/../ext/drafter/ext/snowcrash/src/"
#cgo darwin LDFLAGS: -L"${SRCDIR}/../ext/drafter/build/out/Release/" -ldrafter -lsnowcrash -lmarkdownparser -lsundown -lc++
#cgo linux LDFLAGS: -L"${SRCDIR}/../ext/drafter/build/out/Release/" -ldrafter -lsnowcrash -lmarkdownparser -lsundown -lstdc++
#include <stdlib.h>
#include <stdio.h>
#include "drafter.h"
*/
import "C"
import (
	"io"
	"strings"
	"unsafe"

	"github.com/subosito/drafter-go"
)

func parseOptions(n drafter.Options) *C.drafter_parse_options {
	options := C.drafter_init_parse_options()

	if n.NameRequired {
		C.drafter_set_name_required(options)
	}

	if n.SkipBody {
		C.drafter_set_skip_gen_bodies(options)
	}

	if n.SkipBodySchema {
		C.drafter_set_skip_gen_body_schemas(options)
	}

	return options
}

func serializeOptions(n drafter.Options) *C.drafter_serialize_options {
	options := C.drafter_init_serialize_options()

	if n.Format == drafter.JSON {
		C.drafter_set_format(options, C.DRAFTER_SERIALIZE_JSON)
	}

	if n.SourceMaps {
		C.drafter_set_sourcemaps_included(options)
	}

	return options
}

func Parse(r io.Reader, n drafter.Options) ([]byte, error) {
	s, err := readString(r)
	if err != nil {
		return nil, err
	}

	cOptions := parseOptions(n)
	cSource := C.CString(s)
	cResult := &C.drafter_result{}

	code := C.drafter_parse_blueprint(cSource, &cResult, cOptions)
	if code != C.DRAFTER_OK {
		return nil, errMap(int(code))
	}

	C.free(unsafe.Pointer(cSource))
	C.drafter_free_parse_options(cOptions)

	return serialize(cResult, n), nil
}

func Check(r io.Reader, n drafter.Options) ([]byte, error) {
	s, err := readString(r)
	if err != nil {
		return nil, err
	}

	cOptions := parseOptions(n)
	cSource := C.CString(s)
	cResult := &C.drafter_result{}

	code := C.drafter_check_blueprint(cSource, &cResult, cOptions)
	if code != C.DRAFTER_OK {
		return nil, errMap(int(code))
	}

	C.free(unsafe.Pointer(cSource))
	C.drafter_free_parse_options(cOptions)

	return serialize(cResult, n), nil
}

func Version() string {
	return drafter.Version
}

func serialize(r *C.drafter_result, n drafter.Options) []byte {
	cOptions := serializeOptions(n)
	cResult := C.drafter_serialize(r, cOptions)
	results := C.GoString(cResult)

	C.free(unsafe.Pointer(cResult))
	C.drafter_free_serialize_options(cOptions)

	return []byte(results)
}

func errMap(code int) error {
	switch code {
	case C.DRAFTER_EINVALID_INPUT:
		return drafter.ErrInvalidInput
	case C.DRAFTER_EINVALID_OUTPUT:
		return drafter.ErrInvalidOutput
	default:
		return drafter.ErrUnknown
	}
}

func readString(r io.Reader) (string, error) {
	buf := new(strings.Builder)

	if _, err := io.Copy(buf, r); err != nil {
		return "", err
	}

	return buf.String(), nil
}
