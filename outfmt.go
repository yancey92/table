package lksctl

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

const TAG = "head"

type bd struct {
	H   string // table horizontal "─"
	V   string // table vertical "│"
	HV  string // "┼"
	VR  string // "├"
	VL  string // "┤"
	VD  string // "┬"
	VU  string // "┴"
	LUA string // left up angle（左上角）
	LDA string // left down angle（左下角）
	RUA string // right up angle
	RDA string // right down angle
}

/*
(1)simple e.g:
NAME        AGE   GENDER   HOBBY
Zhang san   23    Man      Basketball
                           Ping pang ball
Li si       22    WoMan    Swimming
Wang wu     21    Man      Basketball

(2)complex e.g:
┌───────────┬─────┬────────┬─────────────────┐
│ NAME      │ AGE │ GENDER │ HOBBY           │
├───────────┼─────┼────────┼─────────────────┤
│ Zhang san │ 23  │ Man    │ Basketball      │
│           │     │        │ Ping pang ball  │
├───────────┼─────┼────────┼─────────────────┤
│ Li si     │ 22  │ WoMan  │ Swimming        │
├───────────┼─────┼────────┼─────────────────┤
│ Wang wu   │ 21  │ Man    │                 │
└───────────┴─────┴────────┴─────────────────┘
*/
var m = map[string]bd{
	"simple": {},
	"ascii":  {"-", "|", "+", "+", "+", "+", "+", "+", "+", "+", "+"},
	"box":    {"─", "│", "┼", "├", "┤", "┬", "┴", "┌", "└", "┐", "┘"},
}

/* The data stored in the cells of a table, e.g:
┌────────────────┐
│ Basketball     │
│ Ping pang ball │
└────────────────┘
 */
type TableCell struct {
	Items []string
}

/*
 * Param "obj" must be a Slice interface
 * head []TableCell：  	store head of table
 * models [][]TableCell: 	store rows of table
 */
func CreateTableCell(obj interface{}) (models [][]TableCell, err error) {
	refValue := reflect.ValueOf(obj)
	if refValue.Kind() != reflect.Slice {
		err= errors.New(fmt.Sprintf("param \"%T\" should be on slice value", obj))
		return
	}

	valueSlice := make([]interface{}, refValue.Len())
	for i := 0; i < refValue.Len(); i++ {
		valueSlice[i] = refValue.Index(i).Interface()
	}

	head := make([]TableCell, 0)
	models = append(models, head)

	// Iterate over each object
	for i, element := range valueSlice {

		rv := reflect.ValueOf(element)
		rt := reflect.TypeOf(element)
		// If it is a pointer, it gets the element to which it points
		if rv.Kind() == reflect.Ptr {
			rv = rv.Elem()
			rt = rt.Elem()
		}

		if rv.Kind() != reflect.Struct {
			err= errors.New(fmt.Sprintf("items of slice \"%T\" should be on struct value", models))
			return
		}

		var model []TableCell
		for n := 0; n < rv.NumField(); n++ {
			// PkgPath is a package path that qualifies lowercase (not exported) field names.
			// It is empty for uppercase (export) field names.
			if rt.Field(n).PkgPath != "" {
				continue
			}

			field := rt.Field(n)
			tag := field.Tag.Get(TAG)
			if tag == `` {
				tag = field.Name
			}
			if i == 0 {
				head = append(head, TableCell{[]string{tag}})
			}

			tc := TableCell{}
			if field.Type.Kind() == reflect.Slice {
				for j := 0; j < rv.FieldByName(field.Name).Len(); j++ {
					tc.Items = append(tc.Items, fmt.Sprintf("%v", rv.FieldByName(field.Name).Index(j).Interface()))
				}
			} else {
				tc.Items = append(tc.Items, fmt.Sprintf("%v", rv.FieldByName(field.Name)))
			}

			model = append(model, tc)
		}
		models = append(models, model)
	}
	models[0] = head
	return
}

// Gets the maximum character length in each column
func columnMaxLen(models [][]TableCell) []int {
	m := make([]int, len(models[0]))

	for _, model := range models {

		for i, f := range model {

			for _, item := range f.Items {
				if len(item) > m[i] {
					m[i] = len(item)
				}
			}

		}
	}
	return m
}

/* Gets the maximum items in each model.
e.g: The maximum number of rows for the following instance is 2
┌───────────┬─────┬────────┬─────────────────┐
│ Zhang san │ 23  │ Man    │ Basketball      │
│           │     │        │ Ping pang ball  │
└───────────┴─────┴────────┴─────────────────┘
*/
func modelMaxWide(models [][]TableCell) []int {
	m := make([]int, len(models))
	for i, model := range models {

		for _, f := range model {

			if len(f.Items) > m[i] {
				m[i] = len(f.Items)
			}

		}
	}
	return m
}

// Get a row from model ([]TableCell) ，store data in []interface{}
func strSliceToInterfaceSlice(tc []TableCell, i int) []interface{} {
	interfaceSlice := make([]interface{}, 0)

	for _, v := range tc {
		if len(v.Items) > i {
			interfaceSlice = append(interfaceSlice, v.Items[i])
		} else {
			interfaceSlice = append(interfaceSlice, "")
		}
	}

	return interfaceSlice
}

func PrintTableSimple(models [][]TableCell) {
	ml := columnMaxLen(models)
	mw := modelMaxWide(models)

	format := ""
	for i, v := range ml {
		format += "%-" + strconv.Itoa(v) + "v"
		if i != len(ml)-1 {
			format += "   "
		} else {
			format += "\n"
		}
	}

	for i, model := range models {
		for n := 0; n < mw[i]; n++ {
			row := strSliceToInterfaceSlice(model, n)
			fmt.Printf(format, row...)
		}

	}
}

func PrintTableAscii(models [][]TableCell) {
	printComplexTable(models, m["ascii"])
}

func PrintTableBox(models [][]TableCell) {
	printComplexTable(models, m["box"])
}

func printComplexTable(models [][]TableCell, t bd) {
	columnNum := len(models[0])
	ml := columnMaxLen(models)
	mw := modelMaxWide(models)

	fmtLineTop := t.LUA + ""   // e.g: ┌────┬──────────────────────┬─────┬─────────┬─────────────┐
	fmtLineMiddle := t.VR + "" // e.g: ├────┼──────────────────────┼─────┼─────────┼─────────────┤
	fmtLineDown := t.LDA + ""  // e.g: └────┴──────────────────────┴─────┴─────────┴─────────────┘
	formatRow := t.V + ""

	for i, v := range ml {
		formatRow += " %-" + strconv.Itoa(v) + "v "
		fmtLineTop += t.H + "%-" + strconv.Itoa(v) + "v" + t.H
		fmtLineMiddle += t.H + "%-" + strconv.Itoa(v) + "v" + t.H
		fmtLineDown += t.H + "%-" + strconv.Itoa(v) + "v" + t.H
		if i != len(ml)-1 {
			formatRow += t.V
			fmtLineTop += t.VD
			fmtLineMiddle += t.HV
			fmtLineDown += t.VU
		} else {
			formatRow += " " + t.V + "\n"
			fmtLineTop += t.H + t.RUA + "\n"
			fmtLineMiddle += t.H + t.VL + "\n"
			fmtLineDown += t.H + t.RDA + "\n"
		}
	}

	hs := make([]interface{}, columnNum)
	for i, v := range ml {
		h := ""
		for j := 0; j < v; j++ {
			h += t.H
		}
		hs[i] = h
	}

	fmt.Printf(fmtLineTop, hs...)
	for i, model := range models {
		for n := 0; n < mw[i]; n++ {
			row := strSliceToInterfaceSlice(model, n)
			fmt.Printf(formatRow, row...)
		}
		if i < len(models)-1 {
			fmt.Printf(fmtLineMiddle, hs...)
		}
	}
	fmt.Printf(fmtLineDown, hs...)
}
