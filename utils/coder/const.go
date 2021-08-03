package coder

import (
	"math/rand"
	"time"
)

const (
	TemplateBlend       = "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	TemplateDigit       = "1234567890"
	TemplateLetterLower = "abcdefghijklmnopqrstuvwxyz"
	TemplateLetterUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	TemplateZnCn        = "的一了是我不在人们有来他这上着个地到大里说去子得也和那要下看天时过出小么起你都把好还多没为又可家学只以主会样年想能生同老中从自面前" +
		"头到它后然走很像见两用她国动进成回什边作对开而已些现山民候经发工向事命给长水几义三声于高正妈手知理眼志点心战二问但身方实吃做叫当住听" +
		"革打呢真党全才四已所敌之最光产情路分总条白话东席次亲如被花口放儿常西气五第使写军吧文运在果怎定许快明行因别飞外树物活部门无往船望新带" +
		"队先力完间却站代员机更九您每风级跟笑啊孩万少直意夜比阶连车重便斗马哪化太指变社似士者干石满决百原拿群究各六本思解立河爸村八难早论吗根" +
		"共让相研今其书坐接应关信觉死步反处记将千找争领或师结块跑谁草越字加脚紧爱等习阵怕月青半火法题建赶位唱海七女任件感准张团屋爷离色脸片科" +
		"倒睛利世病刚且由送切星晚表够整认响雪流未场该并底深刻平伟忙提确近亮轻讲农古黑告界拉名呀土清阳照办史改历转画造嘴此治北必服雨穿父内识验" +
		"传业菜爬睡兴"

	TypeBlend       = 0
	TypeDigit       = 1
	TypeLetterLower = 2
	TypeLetterUpper = 3

	DefaultLen = 6
)

type CodeType = int
type Template = string

func init() {
	rand.Seed(time.Now().UnixNano())
}
