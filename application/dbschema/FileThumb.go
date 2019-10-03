// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/echo/param"
	
)

type Slice_FileThumb []*FileThumb

func (s Slice_FileThumb) Range(fn func(m factory.Model) error ) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_FileThumb) RangeRaw(fn func(m *FileThumb) error ) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

// FileThumb 图片文件缩略图
type FileThumb struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*FileThumb
	namer   func(string) string
	connID  int
	
	Id        	uint64  	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"主键" json:"id" xml:"id"`
	FileId    	uint64  	`db:"file_id" bson:"file_id" comment:"文件ID" json:"file_id" xml:"file_id"`
	Size      	uint64  	`db:"size" bson:"size" comment:"文件大小" json:"size" xml:"size"`
	Width     	uint    	`db:"width" bson:"width" comment:"宽度(像素)" json:"width" xml:"width"`
	Height    	uint    	`db:"height" bson:"height" comment:"高度(像素)" json:"height" xml:"height"`
	Dpi       	uint    	`db:"dpi" bson:"dpi" comment:"分辨率" json:"dpi" xml:"dpi"`
	SaveName  	string  	`db:"save_name" bson:"save_name" comment:"保存名称" json:"save_name" xml:"save_name"`
	SavePath  	string  	`db:"save_path" bson:"save_path" comment:"保存路径" json:"save_path" xml:"save_path"`
	ViewUrl   	string  	`db:"view_url" bson:"view_url" comment:"访问网址" json:"view_url" xml:"view_url"`
	UsedTimes 	uint    	`db:"used_times" bson:"used_times" comment:"被使用的次数" json:"used_times" xml:"used_times"`
	Md5       	string  	`db:"md5" bson:"md5" comment:"缩略图文件MD5值" json:"md5" xml:"md5"`
}

func (this *FileThumb) Trans() *factory.Transaction {
	return this.trans
}

func (this *FileThumb) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *FileThumb) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *FileThumb) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *FileThumb) Objects() []*FileThumb {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *FileThumb) NewObjects() factory.Ranger {
	return &Slice_FileThumb{}
}

func (this *FileThumb) InitObjects() *[]*FileThumb {
	this.objects = []*FileThumb{}
	return &this.objects
}

func (this *FileThumb) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *FileThumb) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *FileThumb) Short_() string {
	return "file_thumb"
}

func (this *FileThumb) Struct_() string {
	return "FileThumb"
}

func (this *FileThumb) Name_() string {
	if this.namer != nil {
		return WithPrefix(this.namer(this.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(this.Short_())(this))
}

func (this *FileThumb) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *FileThumb) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *FileThumb) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *FileThumb) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.InitObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *FileThumb) GroupBy(keyField string, inputRows ...[]*FileThumb) map[string][]*FileThumb {
	var rows []*FileThumb
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string][]*FileThumb{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*FileThumb{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (this *FileThumb) KeyBy(keyField string, inputRows ...[]*FileThumb) map[string]*FileThumb {
	var rows []*FileThumb
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]*FileThumb{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (this *FileThumb) AsKV(keyField string, valueField string, inputRows ...[]*FileThumb) map[string]interface{} {
	var rows []*FileThumb
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

func (this *FileThumb) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.InitObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *FileThumb) Add() (pk interface{}, err error) {
	this.Id = 0
	err = DBI.EventFire("creating", this, nil)
	if err != nil {
		return
	}
	pk, err = this.Param().SetSend(this).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint64(v)
		}
	}
	if err == nil {
		err = DBI.EventFire("created", this, nil)
	}
	return
}

func (this *FileThumb) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	
	if err = DBI.EventFire("updating", this, mw, args...); err != nil {
		return
	}
	if err = this.Setter(mw, args...).SetSend(this).Update(); err != nil {
		return
	}
	return DBI.EventFire("updated", this, mw, args...)
}

func (this *FileThumb) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *FileThumb) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *FileThumb) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {
	
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

func (this *FileThumb) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func() error { 
		return DBI.EventFire("updating", this, mw, args...)
	}, func() error { this.Id = 0
		return DBI.EventFire("creating", this, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint64(v)
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

func (this *FileThumb) Delete(mw func(db.Result) db.Result, args ...interface{})  (err error) {
	
	if err = DBI.EventFire("deleting", this, mw, args...); err != nil {
		return
	}
	if err = this.Param().SetArgs(args...).SetMiddleware(mw).Delete(); err != nil {
		return
	}
	return DBI.EventFire("deleted", this, mw, args...)
}

func (this *FileThumb) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *FileThumb) Reset() *FileThumb {
	this.Id = 0
	this.FileId = 0
	this.Size = 0
	this.Width = 0
	this.Height = 0
	this.Dpi = 0
	this.SaveName = ``
	this.SavePath = ``
	this.ViewUrl = ``
	this.UsedTimes = 0
	this.Md5 = ``
	return this
}

func (this *FileThumb) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["FileId"] = this.FileId
	r["Size"] = this.Size
	r["Width"] = this.Width
	r["Height"] = this.Height
	r["Dpi"] = this.Dpi
	r["SaveName"] = this.SaveName
	r["SavePath"] = this.SavePath
	r["ViewUrl"] = this.ViewUrl
	r["UsedTimes"] = this.UsedTimes
	r["Md5"] = this.Md5
	return r
}

func (this *FileThumb) FromMap(rows map[string]interface{}) {
	for key, value := range rows {
		switch key {
			case "id": this.Id = param.AsUint64(value)
			case "file_id": this.FileId = param.AsUint64(value)
			case "size": this.Size = param.AsUint64(value)
			case "width": this.Width = param.AsUint(value)
			case "height": this.Height = param.AsUint(value)
			case "dpi": this.Dpi = param.AsUint(value)
			case "save_name": this.SaveName = param.AsString(value)
			case "save_path": this.SavePath = param.AsString(value)
			case "view_url": this.ViewUrl = param.AsString(value)
			case "used_times": this.UsedTimes = param.AsUint(value)
			case "md5": this.Md5 = param.AsString(value)
		}
	}
}

func (this *FileThumb) Set(key interface{}, value ...interface{}) {
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
				case "Id": this.Id = param.AsUint64(vv)
				case "FileId": this.FileId = param.AsUint64(vv)
				case "Size": this.Size = param.AsUint64(vv)
				case "Width": this.Width = param.AsUint(vv)
				case "Height": this.Height = param.AsUint(vv)
				case "Dpi": this.Dpi = param.AsUint(vv)
				case "SaveName": this.SaveName = param.AsString(vv)
				case "SavePath": this.SavePath = param.AsString(vv)
				case "ViewUrl": this.ViewUrl = param.AsString(vv)
				case "UsedTimes": this.UsedTimes = param.AsUint(vv)
				case "Md5": this.Md5 = param.AsString(vv)
			}
	}
}

func (this *FileThumb) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = this.Id
	r["file_id"] = this.FileId
	r["size"] = this.Size
	r["width"] = this.Width
	r["height"] = this.Height
	r["dpi"] = this.Dpi
	r["save_name"] = this.SaveName
	r["save_path"] = this.SavePath
	r["view_url"] = this.ViewUrl
	r["used_times"] = this.UsedTimes
	r["md5"] = this.Md5
	return r
}

func (this *FileThumb) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = this.AsRow()
	}
	return factory.BatchValidate(this.Short_(), kvset)
}

func (this *FileThumb) Validate(field string, value interface{}) error {
	return factory.Validate(this.Short_(), field, value)
}

