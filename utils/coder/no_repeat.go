package coder

//func NoRepeat() string {
//	return GenerateNoRepeat(DefaultLen, TypeBlend)
//}

//func NewNoRepeat(length int) string {
//	return GenerateNoRepeat(length, TypeBlend)
//}

//func GenerateNoRepeat(length int, codeType int) (coder string) {
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
//			coder = fmt.Sprintf("%s%c", coder, characterBlend[index])
//		case TypeDigit:
//			for index >= len(characterDigit) {
//				index = index - len(characterDigit)
//			}
//			coder = fmt.Sprintf("%s%c", coder, characterDigit[index])
//		case TypeLetterLower:
//			for index >= len(characterLetterLower) {
//				index = index - len(characterLetterLower)
//			}
//			coder = fmt.Sprintf("%s%c", coder, characterLetterLower[index])
//		case TypeLetterUpper:
//			for index >= len(characterLetterUpper) {
//				index = index - len(characterLetterUpper)
//			}
//			coder = fmt.Sprintf("%s%c", coder, characterLetterUpper[index])
//		default:
//			for index >= len(characterBlend) {
//				index = index - len(characterBlend)
//			}
//			coder = fmt.Sprintf("%s%c", coder, characterBlend[index])
//		}
//	}
//	return coder
//}
