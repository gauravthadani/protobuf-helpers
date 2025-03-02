package equals

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func Equals(this proto.Message, that proto.Message, ignoreDeprecated bool) bool {
	if !ignoreDeprecated {
		return proto.Equal(this, that)
	}

	if proto.Equal(this, that) {
		return true
	}

	return areProtosEqual(this, that)
}

func arePrimitivesEqual(fd protoreflect.FieldDescriptor, v protoreflect.Value, yReflect protoreflect.Message) bool {
	return v.Equal(yReflect.Get(fd))
}

func areProtosEqual(this proto.Message, that proto.Message) bool {
	xReflect := this.ProtoReflect()
	yReflect := that.ProtoReflect()

	isEqual := false

	xReflect.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if isDeprecated(fd) {
			isEqual = true
			return isEqual
		}

		kind := fd.Kind()
		if kind == protoreflect.MessageKind {
			if fd.IsMap() {
				isEqual = equalMaps(v.Map(), yReflect.Get(fd).Map())
				return isEqual
			} else {
				xVal := v.Message().Interface()
				yVal := yReflect.Get(fd).Message().Interface()
				isEqual = areProtosEqual(xVal, yVal)
				return isEqual
			}
		}
		isEqual = arePrimitivesEqual(fd, v, yReflect)
		return isEqual
	})
	return isEqual
}

func equalMaps(v protoreflect.Map, yVal protoreflect.Map) (result bool) {
	v.Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
		result = areProtosEqual(value.Message().Interface(), yVal.Get(key).Message().Interface())
		return result
	})
	return
}

func isDeprecated(fd protoreflect.FieldDescriptor) bool {
	fo := fd.Options().(*descriptorpb.FieldOptions)
	return fo != nil && *fo.Deprecated
}
