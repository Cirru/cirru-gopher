
package interpreter

func (env *scope) array(xs sequence) (ret unitype) {
  ret.Type = uniArray
  list := &map[unitype]unitype{}
  for i, item := range xs {
    index := float64(i)
    (*list)[uni(index)] = env.getValue(item)
  }
  ret.Value = list
  return
}
