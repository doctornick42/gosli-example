package types

import (
	"errors"
	lib "github.com/doctornick42/gosli/lib"
)

type ASlice []A

func (r ASlice) FirstOrDefault(f func(A) bool) A {
	for _, slEl := range r {
		if f(slEl) {
			return slEl
		}
	}
	var defVal A
	return defVal
}
func (r ASlice) First(f func(A) bool) (A, error) {
	for _, slEl := range r {
		if f(slEl) {
			return slEl, nil
		}
	}
	var defVal A
	return defVal, errors.New("Not found")
}
func (r ASlice) Where(f func(A) bool) ASlice {
	res := make([]A, 0)
	for _, slEl := range r {
		if f(slEl) {
			res = append(res, slEl)
		}
	}
	return ASlice(res)
}
func (r ASlice) Select(f func(A) interface{}) []interface{} {
	res := make([]interface{}, len(r))
	for i := range r {
		res[i] = f(r[i])
	}
	return res
}
func (r ASlice) Page(number int64, perPage int64) (ASlice, error) {
	if number <= 0 {
		return nil, errors.New("Page number should start with 1")
	}
	number--
	first := number * perPage
	if first > int64(len(r)) {
		return []A{}, nil
	}
	last := first + perPage
	if last > int64(len(r)) {
		last = int64(len(r))
	}
	return ASlice(r[first:last]), nil
}
func (r ASlice) Any(f func(A) bool) bool {
	_, err := r.First(f)
	return err == nil
}
func (r ASlice) sliceToEqualers() []lib.Equaler {
	equalerSl := make([]lib.Equaler, len(r))
	for i := range r {
		equalerSl[i] = &r[i]
	}
	return equalerSl
}
func (r ASlice) Contains(el A) (bool, error) {
	equalerSl := r.sliceToEqualers()
	return lib.Contains(equalerSl, &el)
}
func (r ASlice) processSliceOperation(sl2 ASlice, f func([]lib.Equaler, []lib.Equaler) ([]lib.Equaler, error)) (ASlice, error) {
	equalerSl1 := r.sliceToEqualers()
	equalerSl2 := sl2.sliceToEqualers()
	untypedRes, err := f(equalerSl1, equalerSl2)
	if err != nil {
		return nil, err
	}
	res := make([]A, len(untypedRes))
	for i := range untypedRes {
		res[i] = *untypedRes[i].(*A)
	}
	return ASlice(res), nil
}
func (r ASlice) GetUnion(sl2 []A) (ASlice, error) {
	return r.processSliceOperation(sl2, lib.GetUnion)
}
func (r ASlice) InFirstOnly(sl2 []A) (ASlice, error) {
	return r.processSliceOperation(sl2, lib.InFirstOnly)
}
