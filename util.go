package main

func getString(m map[string]interface{}, k string) string {
  if v, ok := m[k]; ok {
    if s, ok := v.(string); ok {
      return s
    }
  }
  return ""
}

