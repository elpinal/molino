package vm

import (
  "errors"
  "fmt"
  "reflect"
  "sort"
  "strconv"
  "strings"
)

type Env struct {
  name   string
  env    map[string]reflect.Value
  parent *Env
}

type Func func(args ...reflect.Value) (reflect.Value, error)

type Keyword string

// NewEnv create new global scope.
func NewEnv() *Env {
  return &Env{env: make(map[string]reflect.Value), parent: nil}
}

// NewEnv create new child scope.
func (e *Env) NewEnv() *Env {
  return &Env{env: make(map[string]reflect.Value), parent: e, name: e.name}
}

// Destroy delete current scope.
func (e *Env) Destroy() {
  if e.parent == nil {
    return
  }
  for k, v := range e.parent.env {
    if v.IsValid() && v.Interface() == e {
      delete(e.parent.env, k)
    }
  }
  e.parent = nil
  e.env = nil
}

func (e *Env) DefineGlobal(k string, v reflect.Value) error {
  global := e
  for global.parent != nil {
    global = global.parent
  }
  return global.Define(k, v)
}

func (e *Env) Define(k string, v reflect.Value) error {
  e.env[k] = v
  return nil
}

func (e *Env) Get(k string) (reflect.Value, error) {
  for {
    if e.parent == nil {
      v, ok := e.env[k]
      if !ok {
        return reflect.ValueOf(nil), fmt.Errorf("Undefined symbol '%s'", k)
      }
      return v, nil
    }
    if v, ok := e.env[k]; ok {
      return v, nil
    }
    e = e.parent
  }
  return reflect.ValueOf(nil), fmt.Errorf("Undefined symbol '%s'", k)
}

func Run(stmt Statement, env *Env) (reflect.Value, error) {
  v, err := Evaluate(stmt, env)
  if err != nil {
    return v, err
  }
  return v, nil
}

func Evaluate(statement Statement, env *Env) (reflect.Value, error) {
  switch stmt := statement.(type) {
  case *ExpressionStatement:
    v, err := evaluateExpr(stmt.Expr, env)
    if err != nil {
      return v, err
    }
    //return strconv.Itoa(v), nil
    return v, nil
  default:
    panic("Unknown Statement type")
  }
}

func evaluateExprDef(name string, v reflect.Value, env *Env) (reflect.Value, error) {
  env.DefineGlobal(name, v)
  return v, nil
}

func evaluateExpr(expr Expression, env *Env) (reflect.Value, error) {
  switch e := expr.(type) {
  case *NumberExpression:
    v, err := strconv.ParseInt(e.Lit, 10, 64)
    if err != nil {
      return reflect.ValueOf(0), err
    }
    return reflect.ValueOf(v), nil
  case *IdentifierExpression:
//    if v, ok := env.env[e.Lit]; ok {
//      return v, nil
//    } else {
//      return reflect.ValueOf(0), fmt.Errorf("Undefined variable: %s", e.Lit)
//    }
    return env.Get(e.Lit)
  case *BoolExpression:
    return reflect.ValueOf(e.Bool), nil
  case *NilExpression:
    return reflect.ValueOf(nil), nil
  case *StringExpression:
    return reflect.ValueOf(e.Lit), nil
  case *VectorExpression:
    a := make([]interface{}, len(e.Exprs))
    for i, expr := range e.Exprs {
      arg, err := evaluateExpr(expr, env)
      if err != nil {
        return arg, err
      }
      a[i] = arg.Interface()
    }
    return reflect.ValueOf(a), nil
  case *MapExpression:
    a := make(map[interface{}]interface{})
    for k, v := range e.Map {
      kk, err := evaluateExpr(k, env)
      if err != nil {
        return kk, err
      }
      vv, err := evaluateExpr(v, env)
      if err != nil {
        return vv, err
      }
      a[kk.Interface()] = vv.Interface()
    }
    return reflect.ValueOf(a), nil
  case *CallExpression:
    v, err := evaluateExpr(e.Expr, env)
    if err != nil {
      return reflect.ValueOf(v), err
    }
    if v.Kind() != reflect.Func {
      return reflect.ValueOf(v), errors.New(fmt.Sprint("Unknown Function: ", v.Interface()))
    }

    _, isReflect := v.Interface().(Func)

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
    rets := v.Call(args)
    ev := rets[1].Interface()
    if ev != nil {
      return reflect.ValueOf(nil), ev.(error)
    }
    var ret reflect.Value
    if !isReflect {
      ret = rets[0].Interface().(reflect.Value)
    } else {
      ret = rets[0].Interface().(reflect.Value) //.Interface().(reflect.Value)
    }

    return ret, nil
  case *DefExpression:
    v, err := evaluateExpr(e.Expr, env)
    if err != nil {
      return v, err
    }
    if v.Kind() == reflect.Interface {
      v = v.Elem()
    }
    return evaluateExprDef(e.VarName, v, env)
  case *UnaryMinusExpression:
    v, err := evaluateExpr(e.SubExpr, env)
    if err != nil {
      return reflect.ValueOf(0), err
    }
    return reflect.ValueOf((-v.Int())), nil
    //return fmt.Errorf("Error of minus %v", v), err
  case *UnaryKeywordExpression:
    var v Keyword = Keyword(e.Lit)
    /*
    v, err := evaluateExpr(e.Expr, env)
    if err != nil {
      return reflect.ValueOf(0), err
    }
    */
    return reflect.ValueOf(v), nil
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
      if s[i-1] == s[i] {
        return reflect.ValueOf(nil), errors.New("Can't have more than 2 overloads with some arity")
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
          if len(args) != len(fns[n].Args.Args) {
            var more []interface{}
            for i := len(fns[n].Args.Args); i < len(args); i++ {
              more = append(more, args[i].Interface())
            }
            newenv.env[fns[n].Args.More] = reflect.ValueOf(more)
          }
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
          rr, _ := evaluateExpr(ex, newenv)
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

  case *IfExpression:
    v, err := evaluateExpr(e.Expr, env)
    if err != nil {
      return v, err
    }
    if toBool(v) {
      t, err := evaluateExpr(e.True, env)
      if err != nil {
        return t, err
      }
      return t, nil
    } else {
      if reflect.ValueOf(e.False) != reflect.ValueOf(nil) {
        f, err := evaluateExpr(e.False, env)
        if err != nil {
          return f, err
        }
        return f, nil
      }
    }
    return v, nil
  case *BinOpExpression:
    //a := make([]interface{}, len(e.HS))
    //for i, hs := range e.HS {
    //  hsV, err := evaluateExpr(hs, env)
    //  if err != nil {
    //    return hsV, err
    //  }
    //  a[i] = hsV.Int()
    //}
    c, err := evaluateExpr(e.HS[0], env)
    if err != nil {
      return c, err
    }
    var b = c.Interface()
    for i := 1; i < len(e.HS); i++ {
      a, err := evaluateExpr(e.HS[i], env)
      if err != nil {
        return a, err
      }
      switch e.Operator {
      case '+':
        b = b.(int64) + a.Interface().(int64)
      case '-':
        b = b.(int64) - a.Interface().(int64)
      case '*':
        b = b.(int64) * a.Interface().(int64)
      case '/':
        b = b.(int64) / a.Interface().(int64)
      case '%':
        b = b.(int64) % a.Interface().(int64)
      case '=':
        if i == 1 {
          b = equal(a, c)
        } else {
          b = equal(a, c) && b.(bool)
        }
      default:
        panic("Unknown operator")
      }
    }
    return reflect.ValueOf(b), nil
  default:
    panic("Unknown Expression type")
  }
}

// toBool convert all reflect.Value-s into bool.
func toBool(v reflect.Value) bool {
  if v.Kind() == reflect.Interface {
    v = v.Elem()
  }
  switch {
  case v == reflect.ValueOf(nil):
    return false
  case v.Kind() == reflect.Bool:
    return v.Bool()
  }
  return true
}


func isNil(v reflect.Value) bool {
  if !v.IsValid() || v.Kind().String() == "unsafe.Pointer" {
    return true
  }
  if (v.Kind() == reflect.Interface || v.Kind() == reflect.Ptr) && v.IsNil() {
    return true
  }
  return false
}

func isNum(v reflect.Value) bool {
  switch v.Kind() {
  case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64:
    return true
  }
  return false
}

// equal return true when lhsV and rhsV is same value.
func equal(lhsV, rhsV reflect.Value) bool {
  if isNil(lhsV) && isNil(rhsV) {
    return true
  }
  if lhsV.Kind() == reflect.Interface || lhsV.Kind() == reflect.Ptr {
    lhsV = lhsV.Elem()
  }
  if rhsV.Kind() == reflect.Interface || rhsV.Kind() == reflect.Ptr {
    rhsV = rhsV.Elem()
  }
  if !lhsV.IsValid() || !rhsV.IsValid() {
    return true
  }
  if isNum(lhsV) && isNum(rhsV) {
    if rhsV.Type().ConvertibleTo(lhsV.Type()) {
      rhsV = rhsV.Convert(lhsV.Type())
    }
  }
  if lhsV.CanInterface() && rhsV.CanInterface() {
    return reflect.DeepEqual(lhsV.Interface(), rhsV.Interface())
  }
  return reflect.DeepEqual(lhsV, rhsV)
}

// toInt64 convert all reflect.Value-s into int64.
func toInt64(v reflect.Value) int64 {
  if v.Kind() == reflect.Interface {
    v = v.Elem()
  }
  switch v.Kind() {
  case reflect.Float32, reflect.Float64:
    return int64(v.Float())
  case reflect.Int, reflect.Int32, reflect.Int64:
    return v.Int()
  case reflect.String:
    s := v.String()
    var i int64
    var err error
    if strings.HasPrefix(s, "0x") {
      i, err = strconv.ParseInt(s, 16, 64)
    } else {
      i, err = strconv.ParseInt(s, 10, 64)
    }
    if err == nil {
      return int64(i)
    }
  }
  return 0
}
