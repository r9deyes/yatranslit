package yatranslit

import (
	"regexp"
	"testing"
)

func TestTranslate(t *testing.T) {
	testcases := []struct {
		name, in, expected string
	}{
		{"it_outputs_translited_slug_for_one_word", "привет", "privet"},

		{"it_outputs_translited_slug_for_two_words", "Эхо москвы", "ehkho-moskvy"},

		{"it_outputs_translited_slug_for_several_inputs1", "завтра", "zavtra"},
		{"it_outputs_translited_slug_for_several_inputs2", "сволочь", "svoloch"},
		{"it_outputs_translited_slug_for_several_inputs3", "cуперплохие", "cuperplohie"},
		{"it_outputs_translited_slug_for_several_inputs4", "жопа", "zhopa"},

		{"it_changes_whitespaces_for_dashes", "орел девятого легиона", "orel-devyatogo-legiona"},

		{"it_trims_the_input_string", " Ведущие сайты объявлений! Точный подбор для любой тематики! Эффективно! ", "vedushchie-sajty-obyavlenij-tochnyj-podbor-dlya-lyuboj-tematiki-ehffektivno"},

		// FIXME remove obscene language
		{"it_can_deal_with_h_kh_complexities", "сходил тухачевский под пихту в тюхал свой хер в хуйню", "skhodil-tuhachevskij-pod-pihtu-v-tyuhal-svoj-her-v-hujnyu"},

		{"it_forces_lowercase_on_all_data1", "Россия завершила поставку зенитных ракетных систем С-300 в Иран", "rossiya-zavershila-postavku-zenitnyh-raketnyh-sistem-s-300-v-iran"},
		{"it_forces_lowercase_on_all_data2", "Шакро Молодого назвали лидером преступного сообщества России", "shakro-molodogo-nazvali-liderom-prestupnogo-soobshchestva-rossii"},
		{"it_forces_lowercase_on_all_data3", "В России испытали поражающее ЦЕЛЬ без СНАаряДов оружие", "v-rossii-ispytali-porazhayushchee-cel-bez-snaaryadov-oruzhie"},

		{"it_changes_specialcharacters_for_nothing1", "СМИ узнали о планах России `взимать` 'плату' ~за \\пересечение /границы", "smi-uznali-o-planah-rossii-vzimat-platu-za-peresechenie-granicy"},
		{"it_changes_specialcharacters_for_nothing2", "Строка для транслитерации, по правилам Яндекса!", "stroka-dlya-transliteracii-po-pravilam-yandeksa"},
		{"it_changes_specialcharacters_for_nothing3", "Путин, Меркель и /О/лланд :!обсудили? встречу в «нормандском формате»", "putin-merkel-i-olland-obsudili-vstrechu-v-normandskom-formate"},

		{"it_translits_long_frases1",
			"Лауреатом Нобелевской премии по литературе в 2016 г. стал американский музыкант и автор-исполнитель Боб Дилан. Он получил награду «за создание новых поэтических выражений в рамках американской песенной традиции».",
			"laureatom-nobelevskoj-premii-po-literature-v-2016-g-stal-amerikanskij-muzykant-i-avtor-ispolnitel-bob-dilan-on-poluchil-nagradu-za-sozdanie-novyh-poehticheskih-vyrazhenij-v-ramkah-amerikanskoj-pesennoj-tradicii"},

		{"it_translits_long_frases2",
			"Пресс-секретарь президента России Дмитрий Песков рассказал, что Владимиру Путину сообщили о публикациях в СМИ о нападении на дочь бойца смешанных стилей Федора Емельяненко. Российская газета 13:50",
			"press-sekretar-prezidenta-rossii-dmitrij-peskov-rasskazal-chto-vladimiru-putinu-soobshchili-o-publikaciyah-v-smi-o-napadenii-na-doch-bojca-smeshannyh-stilej-fedora-emelyanenko-rossijskaya-gazeta-13-50"},

		{"it_translits_long_frases3",
			"В сервисе есть мощный инструмент для создания списков подписчиков как по простым, так и по невероятно сложным сценариям. Например, можно легко выделить тех, кто не открыл ни одно ваше письмо или тех, кто ни разу не кликнул по ссылкам. Можно найти всех подписчиков на яндексе и выделить их в отдельный список. В общем, в зависимости от того, какие данные о подписчике вы собираете, можно формировать очень точные выборки и сегменты подписчиков. И делается это все очень просто и быстро.",
			"v-servise-est-moshchnyj-instrument-dlya-sozdaniya-spiskov-podpischikov-kak-po-prostym-tak-i-po-neveroyatno-slozhnym-scenariyam-naprimer-mozhno-legko-vydelit-tekh-kto-ne-otkryl-ni-odno-vashe-pismo-ili-tekh-kto-ni-razu-ne-kliknul-po-ssylkam-mozhno-najti-vsekh-podpischikov-na-yandekse-i-vydelit-ih-v-otdelnyj-spisok-v-obshchem-v-zavisimosti-ot-togo-kakie-dannye-o-podpischike-vy-sobiraete-mozhno-formirovat-ochen-tochnye-vyborki-i-segmenty-podpischikov-i-delaetsya-ehto-vse-ochen-prosto-i-bystro"},
		{"it_can_override_when_a_string_starts_or_ends_with_special_symbol", ":Times: Путин начнет атаку на Алеппо с прибытием «Адмирала Кузнецова»", "times-putin-nachnet-ataku-na-aleppo-s-pribytiem-admirala-kuznecova"},
		{"it_outputs_translited_slug_for_word_containing_su", "сутки", "cutki"},
	}

	DT := NewTranslit()
	for _, tc := range testcases {
		transl := DT.Transform(tc.in)
		if transl != tc.expected {
			t.Errorf("Translit case %s: \ngot: %q \nexp: %q \n", tc.name, transl, tc.expected)
		}
	}
}

func TestTranslateWMaxLength(t *testing.T) {
	testcases := []struct {
		name, inText string
		expected     any
		maxLength    int
	}{
		{"it_can_limit_the_output_to_255_chars_on_get_slug_call",
			"В сервисе есть мощный инструмент для создания списков подписчиков как по простым, так и по невероятно сложным сценариям. Например, можно легко выделить тех, кто не открыл ни одно ваше письмо или тех, кто ни разу не кликнул по ссылкам. Можно найти всех подписчиков на яндексе и выделить их в отдельный список. В общем, в зависимости от того, какие данные о подписчике вы собираете, можно формировать очень точные выборки и сегменты подписчиков. И делается это все очень просто и быстро.",
			shouldMatch(`^[\w-]{1,255}$`),
			255,
		},

		{"it_can_limit_the_output_to_120_chars_on_get_slug_call",
			"В сервисе есть мощный инструмент для создания списков подписчиков как по простым, так и по невероятно сложным сценариям. Например, можно легко выделить тех, кто не открыл ни одно ваше письмо или тех, кто ни разу не кликнул по ссылкам. Можно найти всех подписчиков на яндексе и выделить их в отдельный список. В общем, в зависимости от того, какие данные о подписчике вы собираете, можно формировать очень точные выборки и сегменты подписчиков. И делается это все очень просто и быстро.",
			shouldMatch(`^[\w-]{1,120}$`),
			120,
		},

		{"it_can_limit_the_output_to_some_number_of_characters",
			"Россия завершила поставку зенитных ракетных систем С-300 в Иран",
			"rossiya-zavershila-postav",
			25,
		},

		{"it_can_limit_the_output_to_some_number_of_characters2",
			"очень длинный текст...",
			"ochen-dlin",
			10,
		},
	}

	DT := NewTranslit()
	for _, tc := range testcases {
		DT.SetMaxLength(tc.maxLength)
		transl := DT.Transform(tc.inText)
		switch tc.expected.(type) {
		case string:
			if transl != tc.expected.(string) {
				t.Errorf("Translit case %s: \ngot: %q \nexp: %q \n", tc.name, transl, tc.expected.(string))
			}
		case func(string) bool:
			match := tc.expected.(func(string) bool)
			if !match(transl) {
				t.Errorf("Translit case %s: \ngot: %q \nlenght: %d\nexpected length: %d \n", tc.name, transl, len([]rune(transl)), tc.maxLength)
			}
		}
	}
}

func shouldMatch(s string) func(string) bool {
	re, _ := regexp.Compile(s)
	return re.MatchString
}
