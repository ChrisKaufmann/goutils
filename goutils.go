package goutils

import (
	"crypto/sha1"
	"database/sql"
	"encoding/base64"
	"fmt"
	"html/template"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//puts path vars right into variables passed as params, until it runs out
//ex: pathVars(r,"/entry/",&id,&todo,&val) // populates id, todo, and val
func PathVars(r *http.Request, root string, vals ...*string) {
	a := strings.Split(r.URL.Path[len(root):], "/")
	for i := range vals {
		if len(a) > i {
			*vals[i] = a[i]
		} else {
			*vals[i] = ""
		}
	}
}
func Sth(db *sql.DB, s string) (a *sql.Stmt, err error) {
	a, err = db.Prepare(s)
	if err != nil {
		err.Error()
		fmt.Println(err)
	}
	return a, err
}
func Evenodd(i int) string {
	if i%2 == 0 {
		return "even"
	}
	return "odd"
}
func Tostr(i interface{}) string {
	switch t := i.(type) {
	case string:
		return i.(string)
	case int:
		return strconv.Itoa(i.(int))
	case bool:
		return strconv.FormatBool(i.(bool))
	case template.HTML:
		return string(i.(template.HTML))
	case []byte:
		return string(i.([]byte))
	default:
		t = t
	}
	return i.(string)
}
func Toint(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func Shuffle(data interface{}) {
	switch t := interface{}(data).(type) {
	case *[]int:
		a := data.(*[]int)
		data = Shufflei(*a)
	case *[]string:
		a := data.(*[]string)
		data = Shuffles(*a)
	default:
		fmt.Printf("unexpected type %T", t)
	}
}
func Shuffles(slc []string) []string {
	rand.Seed(time.Now().UnixNano())
	N := len(slc)
	for i := 0; i < N; i++ {
		r := i + rand.Intn(N-i)
		slc[r], slc[i] = slc[i], slc[r]
	}
	return slc
}
func Shufflei(slc []int) []int {
	rand.Seed(time.Now().UnixNano())
	N := len(slc)
	for i := 0; i < N; i++ {
		// choose index uniformly in [i, N-1]
		r := i + rand.Intn(N-i)
		slc[r], slc[i] = slc[i], slc[r]
	}
	return slc
}
func GetHash(s string) string {
	h := sha1.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%s", base64.URLEncoding.EncodeToString(h.Sum(nil)))
}
func RandomString(l int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(RandInt(65, 90))
	}
	return string(bytes)
}
func RandInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}
func Ef(err error) {
	if err != nil {
		panic(err.Error())
	}
}
func Ew(err error) {
	if err != nil {
		print("WARN: " + err.Error())
	}
}
