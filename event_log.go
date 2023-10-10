package w3utils

import (
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

// UnpackLogDataIntoMap uses contract's ABI to parse the data bytes into map of key-value pair
func UnpackLogDataIntoMap(
	contractABI abi.ABI,
	eventKey string,
	logData []byte,
) (
	map[string]interface{},
	error,
) {
	unpacked := make(map[string]interface{})
	if err := contractABI.UnpackIntoMap(unpacked, eventKey, logData); err != nil {
		err := errors.Wrapf(
			err, "failed to unpack eventKey %s with log data %s",
			eventKey, common.Bytes2Hex(logData),
		)

		return nil, errors.Wrap(ErrUnpackFailed, err.Error())
	}

	return unpacked, nil
}

func ExtractFieldFromUnpacked[T any](unpacked map[string]interface{}, field string) (T, error) {
	var t T

	v, found := unpacked[field]
	if !found {
		return t, errors.Wrapf(ErrFieldNotFound, "field %s not found in unpacked", field)
	}

	value, ok := v.(T)
	if !ok {
		return t, errors.Wrapf(ErrTypeAssertion, "field %s is not %s", field, reflect.TypeOf(t).String())
	}

	return value, nil
}
