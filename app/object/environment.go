package object

type Environment struct {
	store map[string]Object
	outer *Environment
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok {
		if e.outer != nil {
			return e.outer.Get(name)
		}
	}
	return obj, ok
}

/* Return the exact environment level where a variable exists */
func (e *Environment) GetEnv(name string) *Environment {
	_, ok := e.store[name]
	if !ok {
		if e.outer != nil {
			return e.outer.GetEnv(name)
		}
	}
	return e
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func (e *Environment) PopStack() *Environment {
	return e.outer
}
