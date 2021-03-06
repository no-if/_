package people

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	even     []int
	modulus  []int
	province []int
	code     []string

	name_0   []string
	name_1_0 []string
	name_1_1 []string
)

func init() {
	rand.Seed(time.Now().UnixNano())

	even = []int{0, 2, 4, 6, 8}
	modulus = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	province = []int{11, 12, 13, 14, 15, 21, 22, 23, 31, 32, 33, 34, 35, 36, 37, 41, 42, 43, 44, 45, 46, 51, 52, 53, 54, 61, 62, 63, 64, 65}
	code = []string{"1", "0", "x", "9", "8", "7", "6", "5", "4", "3", "2"}

	name_0 = []string{"李", "王", "张", "刘", "陈", "杨", "赵", "黄", "周", "吴", "徐", "孙", "胡", "朱", "高", "林", "何", "郭", "马"}
	name_1_0 = []string{"秀", "娟", "英", "华", "慧", "美", "娜", "静", "惠", "珠", "翠", "雅", "芝", "玉", "萍", "红", "娥", "玲", "芬", "芳", "燕", "彩", "菊", "兰", "凤", "洁", "梅", "琳", "云", "莲", "雪", "霞", "莺", "媛", "艳", "瑞", "凡", "佳", "嘉", "琼", "勤", "珍", "贞", "莉", "桂", "娣", "叶", "璧", "璐", "娅", "琦", "晶", "妍", "茜", "秋", "珊", "莎", "锦", "黛", "青", "倩", "婷", "姣", "婉", "娴", "瑾", "颖", "露", "瑶", "怡", "婵", "雁", "蓓", "纨", "仪", "荷", "丹", "蓉", "眉", "琴", "蕊", "薇", "菁", "梦", "岚", "苑", "筠", "柔", "霭", "凝", "晓", "欢", "霄", "枫", "芸", "菲", "欣", "滢", "伊", "亚", "舒", "影", "荔", "枝", "思", "丽", "秀", "飘", "育", "馥", "琦", "晶", "妍", "茜", "秋", "珊", "莎", "锦", "黛", "青", "倩", "婷", "宁", "蓓", "苑", "婕", "馨", "瑗", "琰", "韵", "融", "园", "艺", "咏", "卿", "澜", "毓", "悦", "昭", "冰", "琬", "茗", "洋", "萌", "冬", "卉", "芹", "蝶", "晴", "萱"}
	name_1_1 = []string{"源", "伟", "刚", "勇", "毅", "俊", "峰", "强", "军", "平", "保", "东", "文", "辉", "嘉", "明", "永", "健", "轩", "凡", "志", "涵", "兴", "良", "海", "山", "仁", "波", "宁", "贵", "福", "生", "龙", "元", "全", "国", "胜", "学", "祥", "才", "发", "武", "新", "利", "清", "飞", "彬", "富", "顺", "信", "阳", "杰", "涛", "昌", "成", "康", "星", "光", "天", "达", "安", "岩", "中", "茂", "进", "林", "有", "坚", "和", "彪", "博", "诚", "先", "敬", "震", "振", "壮", "会", "思", "群", "豪", "心", "邦", "承", "乐", "绍", "功", "松", "善", "厚", "庆", "磊", "民", "友", "裕", "河", "哲", "江", "超", "浩", "亮", "政", "谦", "亨", "奇", "固", "之", "轮", "翰", "朗", "伯", "宏", "言", "若", "鸣", "朋", "斌", "梁", "栋", "维", "启", "克", "伦", "翔", "旭", "鹏", "泽", "晨", "辰", "士", "以", "建", "家", "致", "树", "炎", "德", "行", "时", "泰", "盛", "聪", "昊", "喆", "宇", "立", "楠", "航", "渊", "宁"}
}

func New(sex, min, max int) (name, id string) {
	return Name(sex), Id(sex, min, max)
}

func Name(sex int) (name string) {
	name = name_0[rand.Intn(len(name_0))]
	if sex == 0 {
		return name + name_1_0[rand.Intn(len(name_1_0))]
	}
	return name + name_1_1[rand.Intn(len(name_1_1))]
}

func Id(sex, min, max int) (id string) {
	var id_tab [17]int
	p_01 := province[rand.Intn(len(province))]
	id_tab[0] = p_01 / 10
	id_tab[1] = p_01 % 10
	id_tab[3] = 1
	id_tab[5] = 1

	var age, year, month, day int
	age = rand.Intn(max-min+1) + min //20-30
	year = time.Now().Year() - age
	month = rand.Intn(12) + 1
	day = rand.Intn(28) + 1

	id_tab[6] = year / 1000
	id_tab[7] = year / 100 % 10
	id_tab[8] = year / 10 % 10
	id_tab[9] = year % 10

	id_tab[10] = month / 10
	id_tab[11] = month % 10

	id_tab[12] = day / 10
	id_tab[13] = day % 10

	id_tab[14] = rand.Intn(10)
	id_tab[15] = rand.Intn(10)

	id_tab[16] = func(sex int) (i int) {
		i = even[rand.Intn(5)]
		if sex == 0 {
			if i%2 == 0 {
				return i
			} else {
				i = i + 1
			}
		} else {
			if i%2 == 1 {
				return i
			} else {
				i = i + 1
			}
		}
		return i
	}(sex)

	var sum int
	for i := 0; i < 17; i++ {
		sum = sum + id_tab[i]*modulus[i]
		id = fmt.Sprintf("%s%d", id, id_tab[i])
	}
	return id + code[sum%11]
}
