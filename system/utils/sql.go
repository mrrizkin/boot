package utils

import "strings"

type (
	whereBuilder struct {
		count int
		sql   strings.Builder
		binds []interface{}
	}

	joinBuilder struct {
		sql   strings.Builder
		binds []interface{}
	}
)

func NewWhereBuilder() *whereBuilder {
	return &whereBuilder{
		count: 0,
		sql:   strings.Builder{},
		binds: make([]interface{}, 0),
	}
}

func (wb *whereBuilder) write(s string) {
	if wb.sql.Len()+len(s) > wb.sql.Cap() {
		wb.sql.Grow(len(s))
	}
	wb.sql.WriteString(s)
}

func (wb *whereBuilder) And(s string, binds ...interface{}) {
	if wb.count != 0 {
		wb.write(" AND")
	}

	wb.write(" ")
	wb.write(s)
	wb.binds = append(wb.binds, binds...)
	wb.count++
}

func (wb *whereBuilder) Or(s string, binds ...interface{}) {
	if wb.count != 0 {
		wb.write(" OR")
	}

	wb.write(" ")
	wb.write(s)
	wb.binds = append(wb.binds, binds...)
	wb.count++
}

func (wb *whereBuilder) Get() (string, []interface{}) {
	return wb.sql.String(), wb.binds
}

func NewJoinBuilder() *joinBuilder {
	return &joinBuilder{
		sql: strings.Builder{},
	}
}

func (jb *joinBuilder) write(s string) {
	if jb.sql.Len()+len(s) > jb.sql.Cap() {
		jb.sql.Grow(len(s))
	}
	jb.sql.WriteString(s)
}

func (jb *joinBuilder) InnerJoin(table string, condition string, args ...interface{}) {
	jb.write(" INNER JOIN")
	jb.write(" ")
	jb.write(table)
	jb.write(" ON ")
	jb.write(condition)
	jb.binds = append(jb.binds, args...)
}

func (jb *joinBuilder) LeftJoin(table string, condition string, args ...interface{}) {
	jb.write(" LEFT JOIN")
	jb.write(" ")
	jb.write(table)
	jb.write(" ON ")
	jb.write(condition)
	jb.binds = append(jb.binds, args...)
}

func (jb *joinBuilder) RightJoin(table string, condition string, args ...interface{}) {
	jb.write(" Right JOIN")
	jb.write(" ")
	jb.write(table)
	jb.write(" ON ")
	jb.write(condition)
	jb.binds = append(jb.binds, args...)
}

func (jb *joinBuilder) Get() (string, []interface{}) {
	return jb.sql.String(), jb.binds
}
