package lang

type PersistentHashMap struct {
	count    int
	hasNil   bool
	nilValue interface{}
}

type TransientHashMap struct {
	count    int
	hasNil   bool
	nilValue interface{}
	leafFlag Box
}

func (h PersistentHashMap) createWithCheck(init []interface{}) PersistentHashMap {
	var ret ITransientMap = PersistentHashMap{}.asTransient()
	for i := 0; i < len(init); i += 2 {
		ret = ret.assoc(init[i], init[i + 1])
	}
	return ret.persistent().(PersistentHashMap)
}

func (h PersistentHashMap) assoc(key, val interface{}) IPersistentMap {
	//var ret ITransientMap = PersistentHashMap{}.asTransient()
	//
	return PersistentHashMap{}
}

func (h PersistentHashMap) asTransient() TransientHashMap {
	return TransientHashMap{count: h.count}
}


func (t TransientHashMap) assoc(key, val interface{}) ITransientMap {
	return t.doAssoc(key, val)
}

func (t TransientHashMap) doAssoc(key, val interface{}) ITransientMap {
	if key == nil {
		if t.nilValue != val {
			t.nilValue = val
		}
		if !t.hasNil {
			t.count++
			t.hasNil = true
		}
		return t
	}
	t.leafFlag.val = nil
	//
	if t.leafFlag.val != nil {
		t.count++
	}
	return t
}

func (t TransientHashMap) persistent() IPersistentMap {
	return t.doPersistent()
}

func (t TransientHashMap) doPersistent() IPersistentMap {
	return PersistentHashMap{count: t.count}
}
