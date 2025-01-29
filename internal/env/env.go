package env

import (

	"os"
	"strconv"
)

func GetString(key string) string {
  val:= os.Getenv(key)
  return val
}

func GetInt(key string) int{
val :=os.Getenv(key)

valAsInt,err := strconv.Atoi(val)
if err != nil{
	return 0
}
return valAsInt
}