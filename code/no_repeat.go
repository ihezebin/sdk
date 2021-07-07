package code

//func NoRepeat() string {
//	return GenerateNoRepeat(DefaultLen, TypeBlend)
//}

//func NewNoRepeat(length int) string {
//	return GenerateNoRepeat(length, TypeBlend)
//}

//func GenerateNoRepeat(length int, codeType int) (code string) {
//	temp := time.Now().Unix()
//	expend := len(fmt.Sprint(temp)) / length
//	if expend < 2 {
//		expend = 2
//	}
//	remainder := math.Pow10(expend)
//	for i := 0; i < length; i++ {
//		index := int(temp % int64(remainder))
//		temp = temp / int64(remainder)
//		switch codeType {
//		case TypeBlend:
//			for index >= len(characterBlend) {
//				index = index - len(characterBlend)
//			}
//			code = fmt.Sprintf("%s%c", code, characterBlend[index])
//		case TypeDigit:
//			for index >= len(characterDigit) {
//				index = index - len(characterDigit)
//			}
//			code = fmt.Sprintf("%s%c", code, characterDigit[index])
//		case TypeLetterLower:
//			for index >= len(characterLetterLower) {
//				index = index - len(characterLetterLower)
//			}
//			code = fmt.Sprintf("%s%c", code, characterLetterLower[index])
//		case TypeLetterUpper:
//			for index >= len(characterLetterUpper) {
//				index = index - len(characterLetterUpper)
//			}
//			code = fmt.Sprintf("%s%c", code, characterLetterUpper[index])
//		default:
//			for index >= len(characterBlend) {
//				index = index - len(characterBlend)
//			}
//			code = fmt.Sprintf("%s%c", code, characterBlend[index])
//		}
//	}
//	return code
//}
