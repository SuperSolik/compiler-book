package object

type Environment struct {
	store  map[string]Object
	parent *Environment
}

func NewEnvironment(parent *Environment) *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, parent: parent}
}

func (env *Environment) Get(name string) (Object, bool) {
	// NOTE: this implements scopes:
	//    lookup the variable in the envs/ctxs/scopes up until we don't have anywhere to look in
	obj, ok := env.store[name]
	if !ok && env.parent != nil {
		return env.parent.Get(name)
	}
	return obj, ok
}

func (env *Environment) Set(name string, val Object) Object {
	env.store[name] = val
	return val
}
