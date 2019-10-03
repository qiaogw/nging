// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/echo/param"
	
	"time"
)

type Slice_CollectorGroup []*CollectorGroup

func (s Slice_CollectorGroup) Range(fn func(m factory.Model) error ) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_CollectorGroup) RangeRaw(fn func(m *CollectorGroup) error ) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

// CollectorGroup 采集规则组
type CollectorGroup struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*CollectorGroup
	namer   func(string) string
	connID  int
	
	Id         	uint    	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Uid        	uint    	`db:"uid" bson:"uid" comment:"用户ID" json:"uid" xml:"uid"`
	Name       	string  	`db:"name" bson:"name" comment:"组名" json:"name" xml:"name"`
	Type       	string  	`db:"type" bson:"type" comment:"类型(page-页面规则组;export-导出规则组)" json:"type" xml:"type"`
	Description	string  	`db:"description" bson:"description" comment:"说明" json:"description" xml:"description"`
	Created    	uint    	`db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
}

func (this *CollectorGroup) Trans() *factory.Transaction {
	return this.trans
}

func (this *CollectorGroup) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *CollectorGroup) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *CollectorGroup) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *CollectorGroup) Objects() []*CollectorGroup {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *CollectorGroup) NewObjects() factory.Ranger {
	return &Slice_CollectorGroup{}
}

func (this *CollectorGroup) InitObjects() *[]*CollectorGroup {
	this.objects = []*CollectorGroup{}
	return &this.objects
}

func (this *CollectorGroup) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *CollectorGroup) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *CollectorGroup) Short_() string {
	return "collector_group"
}

func (this *CollectorGroup) Struct_() string {
	return "CollectorGroup"
}

func (this *CollectorGroup) Name_() string {
	if this.namer != nil {
		return WithPrefix(this.namer(this.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(this.Short_())(this))
}

func (this *CollectorGroup) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *CollectorGroup) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *CollectorGroup) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *CollectorGroup) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.InitObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *CollectorGroup) GroupBy(keyField string, inputRows ...[]*CollectorGroup) map[string][]*CollectorGroup {
	var rows []*CollectorGroup
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string][]*CollectorGroup{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*CollectorGroup{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (this *CollectorGroup) KeyBy(keyField string, inputRows ...[]*CollectorGroup) map[string]*CollectorGroup {
	var rows []*CollectorGroup
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]*CollectorGroup{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (this *CollectorGroup) AsKV(keyField string, valueField string, inputRows ...[]*CollectorGroup) map[string]interface{} {
	var rows []*CollectorGroup
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]interface{}{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (this *CollectorGroup) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.InitObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *CollectorGroup) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Type) == 0 { this.Type = "page" }
	err = DBI.EventFire("creating", this, nil)
	if err != nil {
		return
	}
	pk, err = this.Param().SetSend(this).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	if err == nil {
		err = DBI.EventFire("created", this, nil)
	}
	return
}

func (this *CollectorGroup) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	
	if len(this.Type) == 0 { this.Type = "page" }
	if err = DBI.EventFire("updating", this, mw, args...); err != nil {
		return
	}
	if err = this.Setter(mw, args...).SetSend(this).Update(); err != nil {
		return
	}
	return DBI.EventFire("updated", this, mw, args...)
}

func (this *CollectorGroup) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *CollectorGroup) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *CollectorGroup) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {
	
	if val, ok := kvset["type"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["type"] = "page" } }
	m := *this
	m.FromMap(kvset)
	if err = DBI.EventFire("updating", &m, mw, args...); err != nil {
		return
	}
	if err = this.Setter(mw, args...).SetSend(kvset).Update(); err != nil {
		return
	}
	return DBI.EventFire("updated", &m, mw, args...)
}

func (this *CollectorGroup) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func() error { 
	if len(this.Type) == 0 { this.Type = "page" }
		return DBI.EventFire("updating", this, mw, args...)
	}, func() error { this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Type) == 0 { this.Type = "page" }
		return DBI.EventFire("creating", this, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	if err == nil {
		if pk == nil {
			err = DBI.EventFire("updated", this, mw, args...)
		} else {
			err = DBI.EventFire("created", this, nil)
		}
	} 
	return 
}

func (this *CollectorGroup) Delete(mw func(db.Result) db.Result, args ...interface{})  (err error) {
	
	if err = DBI.EventFire("deleting", this, mw, args...); err != nil {
		return
	}
	if err = this.Param().SetArgs(args...).SetMiddleware(mw).Delete(); err != nil {
		return
	}
	return DBI.EventFire("deleted", this, mw, args...)
}

func (this *CollectorGroup) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *CollectorGroup) Reset() *CollectorGroup {
	this.Id = 0
	this.Uid = 0
	this.Name = ``
	this.Type = ``
	this.Description = ``
	this.Created = 0
	return this
}

func (this *CollectorGroup) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["Uid"] = this.Uid
	r["Name"] = this.Name
	r["Type"] = this.Type
	r["Description"] = this.Description
	r["Created"] = this.Created
	return r
}

func (this *CollectorGroup) FromMap(rows map[string]interface{}) {
	for key, value := range rows {
		switch key {
			case "id": this.Id = param.AsUint(value)
			case "uid": this.Uid = param.AsUint(value)
			case "name": this.Name = param.AsString(value)
			case "type": this.Type = param.AsString(value)
			case "description": this.Description = param.AsString(value)
			case "created": this.Created = param.AsUint(value)
		}
	}
}

func (this *CollectorGroup) Set(key interface{}, value ...interface{}) {
	switch k := key.(type) {
		case map[string]interface{}:
			for kk, vv := range k {
				this.Set(kk, vv)
			}
		default:
			var (
				kk string
				vv interface{}
			)
			if k, y := key.(string); y {
				kk = k
			} else {
				kk = fmt.Sprint(key)
			}
			if len(value) > 0 {
				vv = value[0]
			}
			switch kk {
				case "Id": this.Id = param.AsUint(vv)
				case "Uid": this.Uid = param.AsUint(vv)
				case "Name": this.Name = param.AsString(vv)
				case "Type": this.Type = param.AsString(vv)
				case "Description": this.Description = param.AsString(vv)
				case "Created": this.Created = param.AsUint(vv)
			}
	}
}

func (this *CollectorGroup) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = this.Id
	r["uid"] = this.Uid
	r["name"] = this.Name
	r["type"] = this.Type
	r["description"] = this.Description
	r["created"] = this.Created
	return r
}

func (this *CollectorGroup) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = this.AsRow()
	}
	return factory.BatchValidate(this.Short_(), kvset)
}

func (this *CollectorGroup) Validate(field string, value interface{}) error {
	return factory.Validate(this.Short_(), field, value)
}

