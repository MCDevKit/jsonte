package functions

import (
	"github.com/MCDevKit/jsonte/jsonte/types"
)

// Typed invokers let the most common builtin signatures be called without reflect.Value.Call,
// which allocates the argument/result slices and boxes every value. Argument types are
// already validated by checkParams, so the type assertions below cannot fail.

type (
	tN  = *types.JsonNumber
	tS  = *types.JsonString
	tB  = *types.JsonBool
	tA  = *types.JsonArray
	tO  = *types.JsonObject
	tL  = *types.JsonLambda
	tT  = types.JsonType
	tV  = *types.Semver
	tP  = *types.JsonPath
	tNl = *types.JsonNull
)

// boxResult converts a typed result to types.JsonType. A nil interface result panics, exactly
// like the `call[0].Interface().(types.JsonType)` assertion on the reflection path.
func boxResult[R types.JsonType](r R) types.JsonType {
	v := types.JsonType(r)
	if v == nil {
		panic("interface conversion: interface is nil, not types.JsonType")
	}
	return v
}

func inv0[R types.JsonType](f func() R) invoker {
	return func(a []types.JsonType) (types.JsonType, error) {
		return boxResult(f()), nil
	}
}

func inv0E[R types.JsonType](f func() (R, error)) invoker {
	return func(a []types.JsonType) (types.JsonType, error) {
		r, err := f()
		if err != nil {
			return nil, err
		}
		return boxResult(r), nil
	}
}

func inv1[A any, R types.JsonType](f func(A) R) invoker {
	return func(a []types.JsonType) (types.JsonType, error) {
		return boxResult(f(a[0].(A))), nil
	}
}

func inv1E[A any, R types.JsonType](f func(A) (R, error)) invoker {
	return func(a []types.JsonType) (types.JsonType, error) {
		r, err := f(a[0].(A))
		if err != nil {
			return nil, err
		}
		return boxResult(r), nil
	}
}

func inv2[A, B any, R types.JsonType](f func(A, B) R) invoker {
	return func(a []types.JsonType) (types.JsonType, error) {
		return boxResult(f(a[0].(A), a[1].(B))), nil
	}
}

func inv2E[A, B any, R types.JsonType](f func(A, B) (R, error)) invoker {
	return func(a []types.JsonType) (types.JsonType, error) {
		r, err := f(a[0].(A), a[1].(B))
		if err != nil {
			return nil, err
		}
		return boxResult(r), nil
	}
}

func inv3[A, B, C any, R types.JsonType](f func(A, B, C) R) invoker {
	return func(a []types.JsonType) (types.JsonType, error) {
		return boxResult(f(a[0].(A), a[1].(B), a[2].(C))), nil
	}
}

func inv3E[A, B, C any, R types.JsonType](f func(A, B, C) (R, error)) invoker {
	return func(a []types.JsonType) (types.JsonType, error) {
		r, err := f(a[0].(A), a[1].(B), a[2].(C))
		if err != nil {
			return nil, err
		}
		return boxResult(r), nil
	}
}

// makeInvoker returns a typed invoker for the given function body, or nil if the signature
// is not one of the covered ones (the caller then falls back to reflection).
func makeInvoker(body interface{}) invoker {
	switch f := body.(type) {
	// no arguments
	case func() tN:
		return inv0(f)
	case func() tA:
		return inv0(f)
	case func() (tS, error):
		return inv0E(f)

	// one argument
	case func(tN) tN:
		return inv1(f)
	case func(tN) tS:
		return inv1(f)
	case func(tS) tS:
		return inv1(f)
	case func(tS) tN:
		return inv1(f)
	case func(tS) tA:
		return inv1(f)
	case func(tT) tB:
		return inv1(f)
	case func(tT) tS:
		return inv1(f)
	case func(tT) tN:
		return inv1(f)
	case func(tA) tA:
		return inv1(f)
	case func(tA) tN:
		return inv1(f)
	case func(tA) tS:
		return inv1(f)
	case func(tA) tT:
		return inv1(f)
	case func(tO) tA:
		return inv1(f)
	case func(tV) tV:
		return inv1(f)
	case func(tP) tP:
		return inv1(f)
	case func(tS) (tN, error):
		return inv1E(f)
	case func(tS) (tA, error):
		return inv1E(f)
	case func(tS) (tO, error):
		return inv1E(f)
	case func(tS) (tT, error):
		return inv1E(f)
	case func(tS) (tS, error):
		return inv1E(f)
	case func(tS) (tB, error):
		return inv1E(f)
	case func(tS) (tV, error):
		return inv1E(f)
	case func(tS) (tP, error):
		return inv1E(f)
	case func(tA) (tT, error):
		return inv1E(f)
	case func(tA) (tA, error):
		return inv1E(f)
	case func(tA) (tS, error):
		return inv1E(f)
	case func(tA) (tV, error):
		return inv1E(f)
	case func(tT) (tNl, error):
		return inv1E(f)

	// two arguments
	case func(tN, tN) tN:
		return inv2(f)
	case func(tN, tN) tS:
		return inv2(f)
	case func(tN, tN) tA:
		return inv2(f)
	case func(tS, tS) tB:
		return inv2(f)
	case func(tS, tS) tN:
		return inv2(f)
	case func(tS, tS) tS:
		return inv2(f)
	case func(tS, tS) tA:
		return inv2(f)
	case func(tS, tN) tS:
		return inv2(f)
	case func(tA, tT) tN:
		return inv2(f)
	case func(tA, tT) tB:
		return inv2(f)
	case func(tA, tS) tS:
		return inv2(f)
	case func(tA, tN) tA:
		return inv2(f)
	case func(tA, tL) (tT, error):
		return inv2E(f)
	case func(tA, tL) (tA, error):
		return inv2E(f)
	case func(tA, tL) (tB, error):
		return inv2E(f)
	case func(tA, tL) (tN, error):
		return inv2E(f)
	case func(tA, tN) (tN, error):
		return inv2E(f)
	case func(tA, tN) (tT, error):
		return inv2E(f)
	case func(tN, tN) (tN, error):
		return inv2E(f)
	case func(tS, tS) (tA, error):
		return inv2E(f)
	case func(tS, tB) (tT, error):
		return inv2E(f)
	case func(tS, tV) (tA, error):
		return inv2E(f)
	case func(tS, tV) (tT, error):
		return inv2E(f)
	case func(tS, tN) (tO, error):
		return inv2E(f)
	case func(tS, tO) (tS, error):
		return inv2E(f)

	// three arguments
	case func(tN, tN, tN) tN:
		return inv3(f)
	case func(tS, tS, tS) tS:
		return inv3(f)
	case func(tS, tN, tN) tS:
		return inv3(f)
	case func(tA, tN, tN) tA:
		return inv3(f)
	case func(tO, tS, tS) tA:
		return inv3(f)
	case func(tN, tN, tN) (tV, error):
		return inv3E(f)
	case func(tS, tS, tS) (tS, error):
		return inv3E(f)
	case func(tS, tV, tB) (tT, error):
		return inv3E(f)
	case func(tA, tN, tL) (tN, error):
		return inv3E(f)
	case func(tA, tL, tT) (tT, error):
		return inv3E(f)
	}
	return nil
}
