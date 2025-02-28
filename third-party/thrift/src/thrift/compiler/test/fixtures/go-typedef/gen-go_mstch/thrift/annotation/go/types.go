// @generated by Thrift for [[[ program path ]]]
// This file is probably not the place you want to edit!

package go_ // [[[ program thrift source path ]]]

import (
    "fmt"

    thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift"
)


// (needed to ensure safety because of naive import list construction)
var _ = fmt.Printf
var _ = thrift.ZERO


type Name struct {
    Name string `thrift:"name,1" json:"name" db:"name"`
}
// Compile time interface enforcer
var _ thrift.Struct = &Name{}

func NewName() *Name {
    return (&Name{}).
        SetNameNonCompat("")
}

func (x *Name) GetNameNonCompat() string {
    return x.Name
}

func (x *Name) GetName() string {
    return x.Name
}

func (x *Name) SetNameNonCompat(value string) *Name {
    x.Name = value
    return x
}

func (x *Name) SetName(value string) *Name {
    x.Name = value
    return x
}

func (x *Name) writeField1(p thrift.Protocol) error {  // Name
    if err := p.WriteFieldBegin("name", thrift.STRING, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetNameNonCompat()
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Name) readField1(p thrift.Protocol) error {  // Name
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetNameNonCompat(result)
    return nil
}

func (x *Name) String() string {
    type NameAlias Name
    valueAlias := (*NameAlias)(x)
    return fmt.Sprintf("%+v", valueAlias)
}


// Deprecated: Use Name.Set* methods instead or set the fields directly.
type NameBuilder struct {
    obj *Name
}

func NewNameBuilder() *NameBuilder {
    return &NameBuilder{
        obj: NewName(),
    }
}

func (x *NameBuilder) Name(value string) *NameBuilder {
    x.obj.Name = value
    return x
}

func (x *NameBuilder) Emit() *Name {
    var objCopy Name = *x.obj
    return &objCopy
}

func (x *Name) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("Name"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *Name) Read(p thrift.Protocol) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, typ, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if typ == thrift.STOP {
            break;
        }

        switch id {
        case 1:  // name
            if err := x.readField1(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(typ); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}


type Tag struct {
    Tag string `thrift:"tag,1" json:"tag" db:"tag"`
}
// Compile time interface enforcer
var _ thrift.Struct = &Tag{}

func NewTag() *Tag {
    return (&Tag{}).
        SetTagNonCompat("")
}

func (x *Tag) GetTagNonCompat() string {
    return x.Tag
}

func (x *Tag) GetTag() string {
    return x.Tag
}

func (x *Tag) SetTagNonCompat(value string) *Tag {
    x.Tag = value
    return x
}

func (x *Tag) SetTag(value string) *Tag {
    x.Tag = value
    return x
}

func (x *Tag) writeField1(p thrift.Protocol) error {  // Tag
    if err := p.WriteFieldBegin("tag", thrift.STRING, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetTagNonCompat()
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Tag) readField1(p thrift.Protocol) error {  // Tag
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetTagNonCompat(result)
    return nil
}

func (x *Tag) String() string {
    type TagAlias Tag
    valueAlias := (*TagAlias)(x)
    return fmt.Sprintf("%+v", valueAlias)
}


// Deprecated: Use Tag.Set* methods instead or set the fields directly.
type TagBuilder struct {
    obj *Tag
}

func NewTagBuilder() *TagBuilder {
    return &TagBuilder{
        obj: NewTag(),
    }
}

func (x *TagBuilder) Tag(value string) *TagBuilder {
    x.obj.Tag = value
    return x
}

func (x *TagBuilder) Emit() *Tag {
    var objCopy Tag = *x.obj
    return &objCopy
}

func (x *Tag) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("Tag"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *Tag) Read(p thrift.Protocol) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, typ, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if typ == thrift.STOP {
            break;
        }

        switch id {
        case 1:  // tag
            if err := x.readField1(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(typ); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}


type NewType_ struct {
}
// Compile time interface enforcer
var _ thrift.Struct = &NewType_{}

func NewNewType_() *NewType_ {
    return (&NewType_{})
}

func (x *NewType_) String() string {
    type NewType_Alias NewType_
    valueAlias := (*NewType_Alias)(x)
    return fmt.Sprintf("%+v", valueAlias)
}


// Deprecated: Use NewType_.Set* methods instead or set the fields directly.
type NewType_Builder struct {
    obj *NewType_
}

func NewNewType_Builder() *NewType_Builder {
    return &NewType_Builder{
        obj: NewNewType_(),
    }
}

func (x *NewType_Builder) Emit() *NewType_ {
    var objCopy NewType_ = *x.obj
    return &objCopy
}

func (x *NewType_) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("NewType"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *NewType_) Read(p thrift.Protocol) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, typ, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if typ == thrift.STOP {
            break;
        }

        switch id {
        default:
            if err := p.Skip(typ); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

