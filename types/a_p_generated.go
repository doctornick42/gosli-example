package types

import (
	"errors"
	lib "github.com/doctornick42/gosli/lib"
)

type APSlice []*A

func (r APSlice) FirstOrDefault(f func(*A) bool) *A {
	for _, slEl := range r {
		if f(slEl) {
			return slEl
		}
	}
	var defVal *A
	return defVal
}
func (r APSlice) First(f func(*A) bool) (*A, error) {
	for _, slEl := range r {
		if f(slEl) {
			return slEl, nil
		}
	}
	var defVal *A
	return defVal, errors.New("Not found")
}
func (r APSlice) Where(f func(*A) bool) APSlice {
	res := make([]*A, 0)
	for _, slEl := range r {
		if f(slEl) {
			res = append(res, slEl)
		}
	}
	return APSlice(res)
}
func (r APSlice) Select(f func(*A) interface{}) []interface{} {
	res := make([]interface{}, len(r))
	for i := range r {
		res[i] = f(r[i])
	}
	return res
}
func (r APSlice) Page(number int64, perPage int64) (APSlice, error) {
	if number <= 0 {
		return nil, errors.New("Page number should start with 1")
	}
	number--
	first := number * perPage
	if first > int64(len(r)) {
		return []*A{}, nil
	}
	last := first + perPage
	if last > int64(len(r)) {
		last = int64(len(r))
	}
	return APSlice(r[first:last]), nil
}
func (r APSlice) Any(f func(*A) bool) bool {
	_, err := r.First(f)
	return err == nil
}
func (r APSlice) sliceToEqualers() []lib.Equaler {
	equalerSl := make([]lib.Equaler, len(r))
	for i := range r {
		equalerSl[i] = r[i]
	}
	return equalerSl
}
func (r APSlice) Contains(el *A) (bool, error) {
	equalerSl := r.sliceToEqualers()
	return lib.Contains(equalerSl, el)
}
func (r APSlice) processSliceOperation(sl2 APSlice, f func([]lib.Equaler, []lib.Equaler) ([]lib.Equaler, error)) (APSlice, error) {
	equalerSl1 := r.sliceToEqualers()
	equalerSl2 := sl2.sliceToEqualers()
	untypedRes, err := f(equalerSl1, equalerSl2)
	if err != nil {
		return nil, err
	}
	res := make([]*A, len(untypedRes))
	for i := range untypedRes {
		res[i] = untypedRes[i].(*A)
	}
	return APSlice(res), nil
}
func (r APSlice) GetUnion(sl2 []*A) (APSlice, error) {
	return r.processSliceOperation(sl2, lib.GetUnion)
}
func (r APSlice) InFirstOnly(sl2 []*A) (APSlice, error) {
	return r.processSliceOperation(sl2, lib.InFirstOnly)
}
func (r *A) Equal(another lib.Equaler) (bool, error) {
	anotherCasted, ok := another.(*A)
	if !ok {
		return false, errors.New("Types mismatch")
	}
	return r.equal(anotherCasted), nil
}
