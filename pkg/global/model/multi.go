package model

import (
	"bytes"
	"strconv"
)

//批量插入sql 值int型
func MultiInsertIntSql(tableName string, field []string, value [][]int) string {
	var buffer bytes.Buffer
	buffer.WriteString("INSERT INTO ")
	buffer.WriteString(tableName)
	buffer.WriteString("(")
	fieldLen := len(field) - 1
	for i, v := range field {
		buffer.WriteString(v)
		if i < fieldLen {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString(")")
	buffer.WriteString(" VALUES ")
	valueLen := len(value) - 1
	for i, v := range value {
		buffer.WriteString("(")
		l := len(v) - 1
		for ii, vv := range v {
			buffer.WriteString(strconv.Itoa(vv))
			if ii < l {
				buffer.WriteString(",")
			}
		}
		buffer.WriteString(")")
		if i < valueLen {
			buffer.WriteString(",")
		}
	}

	return buffer.String()
}
