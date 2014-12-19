package vm

import (
  "errors"
  "fmt"
  "reflect"
  "sort"
  "strconv"
  "strings"
)

type Func func(args ...reflect.Value) (reflect.Value, error)

func (e *Env) Get(k string) (reflect.Value, error) {
//    fmt.Println("0", intern(k))
    /*
  if sym := intern(k); sym.ns != "" {
    fmt.Println("1", sym)
    v := getmapping(sym)
    return reflect.ValueOf(v), nil
  }
  */
  if sym := intern(k); sym.ns == "" {
    v, isexist := CURRENT_NS.root.(Namespace).getmapping(sym)
    if isexist {
      return reflect.ValueOf(v), nil
    }
  }
  return reflect.ValueOf(nil), fmt.Errorf("Undefined symbol '%s'", k)
}

func evaluateExpr(expr Expression, env *Env) (reflect.Value, error) {
  switch e := expr.(type) {
  case *IdentifierExpression:
  a, err := env.Get(e.Lit)
  if err != nil {
    if e.Lit == "in-ns" {
      return reflect.ValueOf(IN_NAMESPACE), nil
    } else {
      x := analyzeSymbol(intern(e.Lit))
      fmt.Println("8013:", x)
      return reflect.ValueOf(x), nil
      //      return reflect.ValueOf(nil), err
    }
  } else {
    return a, nil
  }
  case *CallExpression:
    v, err := evaluateExpr(e.Expr, env)
    if err != nil {
      return reflect.ValueOf(v), err
    }
//    fmt.Println(v.Interface())
    if m, isVar := v.Interface().(Var); isVar {
      fmt.Printf("8012: %#v\n", v)
      v = reflect.ValueOf(m.root)
    }
    if v.Kind() != reflect.Func {
      fmt.Printf("8011: %#v\n", v)
      fmt.Printf("not func: %#v\n" , e.Expr)
      return reflect.ValueOf(v), errors.New(fmt.Sprint("Unknown Function: ", /*v.Interface(),*/ v, e.Expr, v.Kind()))
    }

    _, isReflect := v.Interface().(Func)

    if !v.Type().IsVariadic() {
      if v.Type().NumIn() != len(e.Args) {
        return reflect.ValueOf(nil), errors.New("Wrong number of args (" + fmt.Sprint(len(e.Args),v.Type().NumIn() ) + ")")
      }
    }
    args := []reflect.Value{}
    for _, expr := range e.Args {
      arg, err := evaluateExpr(expr, env)
      if err != nil {
        return arg, err
      }
      if !isReflect {
        args = append(args, reflect.ValueOf(arg))
      } else {
        args = append(args, reflect.ValueOf(arg))
      }
    }
//    fmt.Printf("%v\n", v)
    rets := v.Call(args)
//    fmt.Printf("005 %v\n", rets[0].Interface())
    ev := rets[1].Interface()
    if ev != nil {
      return reflect.ValueOf(nil), ev.(error)
    }
    //var ret reflect.Value
    ret, isValue := rets[0].Interface().(reflect.Value)
    if !isValue {
      ret = rets[0]
    }
    /*
    if !isReflect {
      ret = rets[0].Interface().(reflect.Value)
    } else {
      ret = rets[0].Interface().(reflect.Value) //.Interface().(reflect.Value)
    }
    */

    return ret, nil
  case *DefExpression:
    /*
    v, err := evaluateExpr(e.Expr, env)
    if err != nil {
      return v, err
    }
    if v.Kind() == reflect.Interface {
      v = v.Elem()
    }
    */
//    intern(e.VarName)
    o, is := lookupVar(intern(e.VarName), true)
    if !is {
      return reflect.ValueOf(nil), errors.New("def error:")
    }
    fmt.Printf("7011: %#v\n", reflect.ValueOf(o))
    CURRENT_NS.root.(Namespace).intern(intern(e.VarName))
    return reflect.ValueOf(e.VarName), nil
  case *FnExpression:
    //a := make([]interface{}, len(e.Fns))
    //for i, fn := range e.Fns {
    //  for _, expr := range fn.exprs {
    //    arg, err := evaluateExpr(expr, env)
    //    if err != nil {
    //      return arg, err
    //    }
    //    a[i] = arg.Interface()
    //  }
    //}
    //return reflect.ValueOf(a), nil

    s := make([]int, len(e.Fns))
    var v = -1
    for i, fn := range e.Fns {
      if fn.Args.Vararg {
        if v == len(fn.Args.Args) {
          return reflect.ValueOf(nil), errors.New("Can't have more than 1 variadic overload")
        }
        if v == -1 || v > len(fn.Args.Args) {
          v = len(fn.Args.Args)
        }
      } else {
        s[i] = len(fn.Args.Args)
      }
    }
    sort.Ints(s)
    for i := 1; i < len(s); i++ {
      if s[i-1] > v && v != -1 {
        return reflect.ValueOf(nil), errors.New("Can't have fixed arity function with more params than variadic function")
      }
      if v == -1 {
        if s[i-1] == s[i] {
          return reflect.ValueOf(nil), errors.New("Can't have more than 2 overloads with some arity ")
        }
      }
    }

    f := func(fns []Fn, env *Env) Func {
      return func(args ...reflect.Value) (reflect.Value, error) {

        var n int = -1
        for i, fn := range fns {
          if fn.Args.Vararg {
            if len(args) >= len(fn.Args.Args) {
              n = i
              break
            }
          } else {
            if len(args) == len(fn.Args.Args) {
              n = i
              break
            }
          }
        }
        if n == -1 {
          return reflect.ValueOf(nil), errors.New("Wrong number of args (" + fmt.Sprint(len(args)) + ")")
        }

        newenv := env.NewEnv()
        if fns[n].Args.Vararg {
          var more []interface{}
          if len(args) != len(fns[n].Args.Args) {
            for i := len(fns[n].Args.Args); i < len(args); i++ {
              if args[i] != reflect.ValueOf(nil) {
                more = append(more, args[i].Interface())
              } else {
                more = append(more, nil)
              }
            }
          }
          /*
          if len(args) == 0 {
            more = more
          }
          */
          newenv.env[fns[n].Args.More] = reflect.ValueOf(more)
          for i := 0; i < len(fns[n].Args.Args); i++ {
            newenv.env[fns[n].Args.Args[i]] = args[i]
          }
        }
        for i, arg := range fns[n].Args.Args {
          newenv.env[arg] = args[i]
        }
        //a := make([]interface{}, len(fns[n].stmts))
        var a reflect.Value
        for _, ex := range fns[n].Exprs {
          rr, err := evaluateExpr(ex, newenv)
          if err != nil {
            return a, err
          }
          //a[i] = rr.Interface()
          a = rr
        }
        //if err == ReturnError {
        //  err = nil
        //  rr = rr.Interface().(reflect.Value)
        //}
        //return rr, err
        return a, nil
      }
    }(e.Fns, env)

    //a[i] = f
    //}

    //env.env[e.Name] = f
    return reflect.ValueOf(f), nil

  case *ConstantExpression:
//    fmt.Printf("%#v\n", e.Expr)
    switch ee := e.Expr.(type) {
    case *IdentifierExpression:
//      fmt.Println("007" ,ee.Lit)
      return reflect.ValueOf(intern(ee.Lit)), nil
    case *NumberExpression, *StringExpression, *NilExpression, *UnaryKeywordExpression: //, *UnaryMinusExpression:
      v, err := evaluateExpr(ee, env)
      if err != nil {
        return v, err
      }
      return v, nil
//    case *CallExpression:  // as list
    default:
      fmt.Printf("Warn number 9011: %#v\n", ee)
    }
    return reflect.ValueOf(nil), nil
  default:
    panic("Unknown Expression type")
  }
}

func lookupVar(sym Symbol, internNew bool) (Var, bool) {
  var v Var
  fmt.Printf("001: %#v\n", sym)
  if sym.ns != "" {
    ns, isexist := namespaceFor(CURRENT_NS.root.(Namespace), sym)
    if !isexist {
      return v, isexist
    }
    var name Symbol = intern(sym.name)
    if internNew && ns.name == CURRENT_NS.root.(Namespace).name {
      v = CURRENT_NS.root.(Namespace).intern(name)
    } else {
      var interned bool
      v, interned = ns.findInternedVar(name)
      if !interned {
        return v, interned
      }
    }
  } else {
    o, isMapped := CURRENT_NS.root.(Namespace).mappings[sym]
    if !isMapped {
      if internNew {
        v = CURRENT_NS.root.(Namespace).intern(intern(sym.name))
      }
    } else {
      v = o
    }
  }
  return v, true
}

func namespaceFor(inns Namespace, sym Symbol) (Namespace, bool) {
  //note, presumes non-nil sym.ns
  // first check against currentNS' aliases...
  var nsSym Symbol = intern(sym.ns)
//  var ns Namespace = inns.lookupAlias(nsSym)
//  if(ns == null) {
    // ...otherwise check the Namespaces map.
    ns, isexist := findNamespace(nsSym)
//  }
  return ns, isexist
}

func analyzeSymbol(sym Symbol) Var {
  return resolveIn(CURRENT_NS.root.(Namespace), sym)
}

func resolveIn(n Namespace, sym Symbol) Var {
  if sym.ns != "" {
    ns, isexist := namespaceFor(n, sym)
    if !isexist {
      panic("No such namespace: " + sym.ns)
    }
    v, interned := ns.findInternedVar(intern(sym.name))
    if !interned {
      panic("No such var: " + sym.ns + "/" + sym.name)
    }
    return v
  }
  o, _ := n.getmapping(sym)
  for _, s := range namespaces {
    for k, _ := range s.mappings {
      fmt.Println(k)
    }
  }
  return o
}
