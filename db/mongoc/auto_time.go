package mongoc

// Deprecated
//func (model *Base) handleAutoTimeUpdate(update interface{}) (interface{}, error) {
//	if update == nil {
//		return nil, errors.New("the document can not be nil")
//	}
//	v := reflect.ValueOf(update)
//	if v.Kind() != reflect.Map {
//		return nil, errors.New(fmt.Sprintf("the update %+v is not a map", v.Interface()))
//	}
//	// 校验map是否为空
//	if len(v.MapKeys()) == 0 {
//		return nil, errors.New("the update can not have no field")
//	}
//	if model.Client().Config().AutoTime {
//		v.SetMapIndex(reflect.ValueOf("$set"), reflect.ValueOf(bson.M{autoTimeKeyUpdate: timer.Now()}))
//	}
//	doc := make(map[string]interface{})
//	for _, k := range v.MapKeys() {
//		doc[k.String()] = v.MapIndex(k).Interface()
//	}
//	return doc, nil
//}

// Deprecated
//func (model *Base) handleAutoTimeInsert(documents ...interface{}) ([]interface{}, error) {
//	docs := make([]interface{}, 0, len(documents))
//	for _, document := range documents {
//		if document == nil {
//			return nil, errors.New("the document can not be nil")
//		}
//		v := reflect.ValueOf(document)
//		for v.Kind() == reflect.Ptr {
//			v = v.Elem()
//		}
//		if v.Kind() == reflect.Struct {
//			m, err := mapper.Struct2Map(v.Interface())
//			if err != nil {
//				return nil, err
//			}
//			v = reflect.ValueOf(m)
//		}
//		if v.Kind() != reflect.Map {
//			return nil, errors.New(fmt.Sprintf("the document %+v is not a map or struct", v.Interface()))
//		}
//		// 校验map是否为空
//		if len(v.MapKeys()) == 0 {
//			return nil, errors.New("the document can not have no field")
//		}
//		if model.Client().Config().AutoTime {
//			// 生成时间
//			now := timer.Now()
//			v.SetMapIndex(reflect.ValueOf(autoTimeKeyCreate), reflect.ValueOf(now))
//			v.SetMapIndex(reflect.ValueOf(autoTimeKeyUpdate), reflect.ValueOf(now))
//		}
//		// 转换成map, 并删除_id, 采用mongo自动生成的
//		doc := make(map[string]interface{})
//		for _, k := range v.MapKeys() {
//			doc[k.String()] = v.MapIndex(k).Interface()
//		}
//		delete(doc, "_id")
//		docs = append(docs, doc)
//	}
//	return docs, nil
//}

var (
	autoTimeKeyCreate = "create_time"
	autoTimeKeyUpdate = "update_time"
)

//func SetAutoTimeKeyCreate(key string) {
//	autoTimeKeyCreate = key
//}
//
//func SetAutoTimeKeyUpdate(key string) {
//	autoTimeKeyUpdate = key
//}
