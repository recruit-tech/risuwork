package fixture

// Industries 業種とサフィックスの対応リスト
var Industries = []struct {
	ID       string
	Name     string
	Suffixes []string
}{
	{"I01", "ITサービス", []string{"テクノロジー", "システム", "ソリューションズ", "ネットワーク", "ソフトウェア"}},
	{"I02", "食品加工業", []string{"フーズ", "食品", "食料品", "プロダクツ", "キッチン"}},
	{"I03", "自動車製造業", []string{"モータース", "自動車", "カー", "エンジン", "車両"}},
	{"I04", "銀行", []string{"銀行", "ファイナンス", "金融", "キャピタル", "マネー"}},
	{"I05", "医療機関", []string{"病院", "クリニック", "ヘルスケア", "メディカル", "ケア"}},
	{"I06", "建設業", []string{"建設", "工務店", "デベロッパー", "エンジニアリング", "アーキテクト"}},
	{"I07", "旅行・観光業", []string{"トラベル", "ツーリズム", "観光", "ツアーズ", "エクスペディション"}},
	{"I08", "運輸・物流業", []string{"トランスポート", "物流", "デリバリー", "運輸", "カーゴ"}},
	{"I09", "小売業", []string{"マーケット", "ストア", "ショップ", "リテール", "バザール"}},
	{"I10", "不動産業", []string{"リアルティ", "エステート", "リース", "ディベロップメント", "プロパティ"}},
	{"I11", "教育機関", []string{"アカデミー", "学校", "カレッジ", "大学", "インスティテューション"}},
	{"I12", "エネルギー産業", []string{"エナジー", "電力", "ガス", "エレクトリック", "パワー"}},
	{"I13", "メディア・広告業", []string{"メディア", "広告", "プロモーション", "マーケティング", "パブリシティ"}},
	{"I14", "スポーツ・レジャー産業", []string{"スポーツ", "レジャー", "フィットネス", "アクティビティ", "エンターテイメント"}},
	{"I15", "保険業", []string{"保険", "ライフ", "ヘルス", "アシュアランス", "カバレッジ"}},
	{"I16", "製薬業", []string{"製薬", "ドラッグ", "バイオテック", "メディシン", "ファーマシューティカル"}},
	{"I17", "農業", []string{"農業", "ファーム", "栽培", "プランテーション", "アグリカルチャー"}},
	{"I18", "ファッション業", []string{"ファッション", "アパレル", "デザイン", "クチュール", "ウェア"}},
	{"I19", "ゲーム開発", []string{"ゲーム", "インタラクティブ", "スタジオ", "ソフトウェア", "エンターテイメント"}},
	{"I20", "航空業", []string{"エアライン", "航空", "エアウェイズ", "フライト", "エア"}},
	{"I21", "グルメ業界", []string{"レストラン", "カフェ", "ダイナー", "ビストロ", "キュイジーヌ"}},
	{"I22", "テクノロジー・ハードウェア", []string{"ハードウェア", "デバイス", "エレクトロニクス", "ガジェット", "インストゥルメンツ"}},
	{"I23", "環境・エコロジー", []string{"環境", "エコ", "サステナビリティ", "グリーン", "エコロジー"}},
	{"I24", "音楽産業", []string{"ミュージック", "レコーディング", "プロダクション", "サウンド", "レーベル"}},
	{"I25", "高等教育", []string{"大学", "カレッジ", "インスティテューション", "スクール", "アカデミー"}},
	{"I26", "鉄鋼業", []string{"鉄鋼", "メタル", "鋳造", "鍛造", "スチール"}},
	{"I27", "酒造業", []string{"酒造", "ワイナリー", "ディスティラリー", "リカー", "アルコール"}},
	{"I28", "インテリアデザイン", []string{"インテリア", "デザイン", "装飾", "ファーニッシング", "デコレーション"}},
	{"I29", "化学産業", []string{"化学", "ケミカル", "ラボラトリー", "シンセシス", "マテリアル"}},
	{"I30", "電子機器製造業", []string{"エレクトロニクス", "デバイス", "ガジェット", "インストゥルメンツ", "エレクトリック"}},
	{"I31", "船舶業", []string{"シッピング", "海運", "クルーズ", "フリート", "ナビゲーション"}},
	{"I32", "通信業", []string{"コミュニケーションズ", "ネットワーク", "コネクティビティ", "ワイヤレス", "テレコム"}},
	{"I33", "レストラン業", []string{"レストラン", "ダイニング", "ビストロ", "カフェ", "グリル"}},
	{"I34", "美容業", []string{"ビューティー", "コスメティック", "スキンケア", "サロン", "エステティック"}},
	{"I35", "食品小売業", []string{"グロサリー", "マーケット", "スーパーマーケット", "フードストア", "プロビジョン"}},
	{"I36", "ホテル業", []string{"ホテル", "リゾート", "旅館", "ホスピタリティ", "ロッジ"}},
	{"I37", "舞台芸術", []string{"シアター", "ドラマ", "パフォーマンス", "プレイハウス", "シアトリカル"}},
	{"I38", "プロパティ管理", []string{"プロパティ", "不動産", "エステート", "リース", "リノベーション"}},
	{"I39", "自然資源産業", []string{"リソース", "鉱業", "エクストラクション", "エナジー", "ナチュラル"}},
	{"I40", "ウェブ開発", []string{"ウェブ", "デジタル", "インターネット", "オンライン", "テクノロジー"}},
	{"I41", "設計業", []string{"設計", "デザイン", "エンジニアリング", "アーキテクチャ", "コンサルティング"}},
	{"I42", "金属加工業", []string{"メタル", "ファブリケーション", "フォージング", "キャスティング", "メタラジー"}},
	{"I43", "資産管理", []string{"アセット", "インベストメント", "キャピタル", "ウェルスマネジメント", "ファイナンス"}},
	{"I44", "ブロードキャスト業", []string{"ブロードキャスト", "ラジオ", "テレビ", "ネットワーク", "チャンネル"}},
	{"I45", "芸術・文化機関", []string{"アート", "カルチャー", "ミュージアム", "ギャラリー", "エキシビション"}},
	{"I46", "インフラストラクチャー", []string{"インフラ", "建設", "エンジニアリング", "ユーティリティ", "トランスポート"}},
	{"I47", "ソフトウェア開発", []string{"ソフトウェア", "プログラミング", "アプリケーション", "システム", "ソリューションズ"}},
	{"I48", "資源回収・リサイクル", []string{"リサイクル", "リカバリー", "廃棄物管理", "エコ", "サステナビリティ"}},
	{"I49", "ファーム・畜産業", []string{"農場", "アグリカルチャー", "牧場", "ライヴストック", "ハーベスト"}},
	{"I50", "芸能業", []string{"エンターテイメント", "ショービジネス", "パフォーマンス", "プロダクション", "タレント"}},
}

// CompanyPrefixes 会社名のプレフィックスとサフィックスのリスト
var CompanyPrefixes = []string{
	"未来", "大洋", "日新", "中和", "光輝", "旭日", "昇龍", "鳳凰", "福寿", "瑞光",
	"新星", "桜花", "富士山", "翠雲", "青空", "銀河", "流星", "風林火山", "大地",
	"千里", "海神", "春風", "花鳥", "月光", "緑葉", "若葉", "飛鳥", "太陽", "星空",
	"黄昏", "朝日", "夕日", "陽炎", "静寂", "風薫", "秋風", "雪月", "雷鳴", "炎天",
	"春暁", "夏陽", "秋麗", "冬凛", "青嵐", "紅葉", "陽春", "白雲", "碧空", "青雲",
	"海風", "山岳", "流雲", "河川", "湖畔", "山頂", "花園", "森羅", "雲海", "雪原",

	"ひかり", "さくら", "みずほ", "ふじ", "はるか", "ひまわり", "ひとみ", "かがやき", "さつき", "ゆめ",
	"なでしこ", "あさひ", "きらり", "ことぶき", "つばさ", "きぼう", "あおば", "ふじかげ", "いずみ", "きらめき",
	"はつひ", "しんせい", "ときわ", "あかね", "いちご", "しらゆき", "ほほえみ", "あおい", "はなみずき", "はやぶさ",
	"つくし", "ひなた", "いぶき", "さざんか", "しんじゅ", "ほしぞら", "はるかぜ", "ふくろう", "あかつき", "ふじはな",
	"たけなみ", "あいびき", "いなほ", "まつば", "もみじ", "みどり", "そよかぜ", "かぐや", "つゆくさ", "ひいらぎ",
	"ふたば", "ささゆり", "はるみ", "しらさぎ", "ほたる", "くれない", "あおぞら", "うみねこ", "こまち", "あおば",

	"フューチャー", "オーシャン", "ニューサン", "チュウワ", "シャイニング", "キョクジツ", "ライジングドラゴン", "ホウオウ", "フクジュ", "ズイコウ",
	"ニュースター", "ミライ", "サクラ", "フジサン", "スイウン", "アオゾラ", "ギンガ", "リュウセイ", "フウリンカザン", "ダイチ",
	"センリ", "カイジン", "ハルカゼ", "カチョウ", "ゲッコウ", "リョクヨウ", "ワカバ", "トビウオ", "タイヨウ", "ホシゾラ",
	"タソガレ", "アサヒ", "ユウヒ", "ヨウエン", "セイジャク", "カゼカオル", "アキカゼ", "セツゲツ", "ライメイ", "エンテン",
	"シュンギョウ", "カヨウ", "アキウララ", "フユリン", "セイラン", "コウヨウ", "ハルシュン", "シラクモ", "アオゾラ", "セイウン",
	"ウミカゼ", "サンガク", "ナガレグモ", "カセン", "コハン", "サントウ", "ハナゾノ", "シンラ", "ウンカイ", "セツゲン",
}

// LastNames 名字リスト
var LastNames = []string{
	"佐藤", "鈴木", "高橋", "田中", "伊藤", "渡辺", "山本", "中村", "小林", "加藤",
	"吉田", "山田", "佐々木", "山口", "斎藤", "松本", "井上", "木村", "林", "清水",
	"山崎", "森", "池田", "橋本", "阿部", "石川", "山下", "中島", "前田", "藤田",
	"後藤", "小川", "岡田", "村上", "石井", "近藤", "長谷川", "藤井", "杉山", "松田",
	"中川", "原田", "小野", "田村", "竹内", "金子", "和田", "中野", "工藤", "上田",
	"武田", "野口", "平野", "西村", "菅原", "新井", "堀", "川口", "矢野", "黒田",
	"岩崎", "古川", "西田", "坂本", "本田", "松井", "小山", "北村", "坂口", "大塚",
	"川崎", "菊地", "松浦", "宮崎", "永井", "藤原", "大久保", "森田", "関口", "島田",
	"安藤", "河野", "菊池", "岡崎", "小田", "青木", "高木", "久保", "岡村", "上野",
	"三浦", "福田", "浅野", "井口", "星野", "岩田", "平井", "石田", "増田", "木下",
	"伊東", "津田", "小泉", "原", "福島", "川上", "宮田", "吉岡", "片山", "黒木",
}

// FirstNames 名前リスト
var FirstNames = []string{
	"一郎", "二郎", "三郎", "四郎", "五郎", "六郎", "七郎", "八郎", "九郎", "十郎",
	"花子", "太郎", "次郎", "三郎", "四郎", "五郎", "六郎", "七郎", "八郎", "九郎",
	"一美", "二美", "三美", "四美", "五美", "六美", "七美", "八美", "九美", "十美",
	"美子", "健二", "太一", "和子", "浩二", "春男", "陽子", "奈緒", "愛子", "啓介",
	"正樹", "綾子", "直人", "香織", "悠斗", "涼子", "拓真", "絵里", "大樹", "志保",
	"真一", "桃子", "亮", "舞", "祐介", "智子", "康介", "友美", "祐樹", "亜美",
	"雄太", "由紀", "竜也", "恵", "洋平", "奈々", "優斗", "茜", "大輔", "華",
	"和也", "実", "健", "愛", "亮介", "陽菜", "雅也", "佳奈", "悠", "沙織",
	"晴人", "奈央", "翔太", "真由", "悠真", "里奈", "悠斗", "美奈", "直樹", "亜紀",
	"俊介", "沙羅", "和真", "瑞希", "翔", "智", "拓也", "美香", "隼人", "麻美",
}

// Passwords 愚直にやるとbcryptが遅いので固定で何個か用意
var Passwords = []struct {
	Raw  string
	Hash string
}{
	{"password", "$2a$10$ltDs.d1HXDkFAH.Ufip6w.k.ZvPbvaPZ89VgOazDe3fXPMy7mW/Fq"},
	{"bT3PrSC6", "$2a$10$2ROKZVtHu/Qx1.I2GLsl1uG7qro4MWr0Djjr.gfoTbfGqYyp3ZMHe"},
	{"L5jM4y6s", "$2a$10$Bfutjk9BeVT1tZAE1x0ofO4sHlNWquEtNnD12BTToZ.V0QapFgwS."},
}

// TagsList 検索タグリスト
var TagsList = []string{
	"交通費支給", "研修あり", "資格取得支援制度", "残業なし", "完全週休二日制", "託児所あり", "寮・社宅・住宅手当あり",
	"産休・育休制度実績あり", "長期休暇あり", "職種変更なし", "フレックスタイム制", "リモートワーク可", "定期健康診断", "社員旅行あり",
	"ストックオプション", "ボーナスあり", "定年制度なし", "住宅補助", "有給休暇", "カジュアルな服装", "メンター制度", "入社祝い金",
	"社員表彰制度", "自己啓発支援", "退職金制度", "昼食補助制度", "サークル活動支援制度",
}

var JobDescriptions = []struct {
	Title       string
	Description string
}{
	{
		Title: "ソフトウェアエンジニア",
		Description: "最新のテクノロジーを使用して、ユーザーにとって直感的かつ効果的なソフトウェアソリューションを設計、開発、テストします。" +
			"この役割では、エンドユーザーにとっての使いやすさと効率性を最大限に引き出すためのソフトウェアを作成します。プログラミングスキル（Java, Python, C++など）が必須であり、" +
			"問題解決能力とチームワークが重要です。具体的には、以下の業務を行います:\n" +
			"- 要件分析と仕様設計\n" +
			"- コーディングとコードレビュー\n" +
			"- 協力してエラーを特定し、修正する\n" +
			"- 継続的な改善と新機能の提案\n" +
			"さらに、ソフトウェアの品質を確保するためのテストフレームワークの開発や、システムのスケーラビリティとパフォーマンスを向上させるための最適化を行います。" +
			"敏捷な環境での経験があり、最新の技術トレンドに興味がある方を歓迎します。柔軟な思考と新しい技術を学ぶ意欲が求められます。" +
			"ソフトウェア開発プロセス全体に関与し、製品の成功に貢献することが期待されます。",
	},
	{
		Title: "プロジェクトマネージャー",
		Description: "プロジェクトの計画、実行、監視、制御、および終了を管理し、チームを導きます。優れたリーダーシップとコミュニケーション能力、スケジュール管理やリソース配分の経験が必要です。" +
			"このポジションでは、複雑なプロジェクトを成功に導くためのリーダーシップスキルが必要です。主な責任は以下の通りです:\n" +
			"- プロジェクト計画の策定とスケジュール管理\n" +
			"- リソースの最適配分と管理\n" +
			"- リスク管理と問題解決\n" +
			"- チームメンバーとの円滑なコミュニケーション\n" +
			"- 進捗状況の監視と報告\n" +
			"プロジェクトの成功には、ステークホルダーとの効果的なコミュニケーションが不可欠です。また、PMP資格保持者を優遇し、プロジェクト管理ツール（JIRA, MS Projectなど）の使用経験がある方を歓迎します。" +
			"複数のプロジェクトを同時に管理し、期限内に成果を出す能力が求められます。",
	},
	{
		Title: "データサイエンティスト",
		Description: "データの収集、分析、解釈を行い、ビジネスの意思決定を支援するためのインサイトを提供します。統計解析や機械学習の知識、PythonやRのスキルが必要です。" +
			"具体的な業務内容は以下の通りです:\n" +
			"- データの収集とクリーニング\n" +
			"- 統計モデルの構築と評価\n" +
			"- データの可視化とレポート作成\n" +
			"- ビジネス上の課題に対するデータドリブンな解決策の提案\n" +
			"大規模データセットの処理経験があり、機械学習アルゴリズムの実装と最適化に熟練した方を求めます。" +
			"データの傾向を把握し、経営陣に対して明確な提案を行うことが求められます。分析結果を元にした戦略的な意思決定を支援し、ビジネスの成長に貢献します。" +
			"また、ビジネスインテリジェンスツール（Tableau, Power BIなど）の使用経験がある方を歓迎します。",
	},
	{
		Title: "マーケター",
		Description: "市場調査、キャンペーンの立案・実施、顧客関係の構築を通じて、ブランドの認知度と売上を向上させます。デジタルマーケティングツールの使用経験、分析能力、クリエイティブな思考が求められます。" +
			"具体的な職務内容は以下の通りです:\n" +
			"- 市場トレンドと競合分析\n" +
			"- マーケティング戦略の策定と実行\n" +
			"- デジタル広告キャンペーンの管理（Google Ads, Facebook Adsなど）\n" +
			"- ソーシャルメディアの運営とコンテンツ作成\n" +
			"- マーケティング効果の測定とレポート作成\n" +
			"また、SEO/SEMの知識があり、分析ツール（Google Analyticsなど）の使用経験がある方を優遇します。創造的なアプローチと分析力を持つ方を求めます。" +
			"さらに、ブランドの声を確立し、ターゲットオーディエンスに効果的にメッセージを伝える能力が求められます。" +
			"マーケティングキャンペーンのROIを最適化し、データに基づく戦略的な意思決定を行うことが期待されます。",
	},
	{
		Title: "カスタマーサポート",
		Description: "顧客からの問い合わせに対応し、問題を迅速に解決して、優れたカスタマーサービスを提供します。優れたコミュニケーションスキルと問題解決能力が必要です。" +
			"具体的な業務内容は以下の通りです:\n" +
			"- 電話、メール、チャットによる顧客対応\n" +
			"- 製品やサービスに関する情報提供とサポート\n" +
			"- 顧客の問題の特定と解決\n" +
			"- 顧客からのフィードバックの収集と報告\n" +
			"- 顧客満足度の向上に向けた施策の実施\n" +
			"フレンドリーでプロフェッショナルな対応ができる方を求めます。また、CRMツールの使用経験がある方を歓迎します。チームワークと顧客志向の姿勢が重要です。" +
			"さらに、製品知識を深め、顧客に対して的確なアドバイスを提供する能力が求められます。" +
			"顧客のフィードバックを元にした製品改善の提案や、サポートプロセスの効率化にも貢献することが期待されます。",
	},
	{
		Title: "人事マネージャー",
		Description: "採用、トレーニング、評価、福利厚生の管理を通じて、組織の人材戦略を推進します。人事管理の経験、効果的なコミュニケーションスキル、リーダーシップが必要です。" +
			"具体的な業務内容は以下の通りです:\n" +
			"- 採用プロセスの管理と候補者の選定\n" +
			"- 新入社員のトレーニングとオンボーディング\n" +
			"- 従業員評価とパフォーマンス管理\n" +
			"- 福利厚生プログラムの設計と実施\n" +
			"- 労働法規の遵守と従業員関係の管理\n" +
			"労働法規の知識を有し、戦略的な人材管理ができる方を優遇します。また、人事関連の資格（SHRM, HRCIなど）を保持している方を歓迎します。" +
			"さらに、従業員のモチベーションを高めるためのプログラムやポリシーを開発し、企業文化の向上に貢献することが期待されます。" +
			"チームと協力して、人材開発の戦略を立案し、組織の長期的な目標を達成するためのリソースを確保します。",
	},
	{
		Title: "営業",
		Description: "新規顧客の開拓、既存顧客との関係強化、製品やサービスの販売を通じて、売上目標を達成します。営業経験、優れたコミュニケーション能力、プレゼンテーションスキルが必要です。" +
			"具体的な職務内容は以下の通りです:\n" +
			"- 新規顧客のリサーチとアプローチ\n" +
			"- 既存顧客との関係維持と強化\n" +
			"- 製品やサービスの提案とデモンストレーション\n" +
			"- 売上目標の設定と達成に向けた戦略立案\n" +
			"- 顧客のフィードバックの収集と改善提案\n" +
			"目標達成への意欲が高く、顧客志向の営業戦略を実行できる方を求めます。また、CRMツールの使用経験がある方を歓迎します。" +
			"さらに、競合分析を行い、差別化された価値提案を顧客に提供する能力が求められます。" +
			"営業活動の結果を分析し、戦略的な意思決定をサポートするための報告書を作成します。",
	},
	{
		Title: "グラフィックデザイナー",
		Description: "デザインコンセプトの作成、ビジュアルコンテンツの制作、ブランドのビジュアルアイデンティティの維持を担当します。デザインツール（Photoshop, Illustratorなど）のスキル、クリエイティブな思考、チームワークが求められます。" +
			"具体的な業務内容は以下の通りです:\n" +
			"- デザインブリーフの理解とコンセプト作成\n" +
			"- 広告、ウェブサイト、印刷物のビジュアルデザイン\n" +
			"- ブランドガイドラインの維持と適用\n" +
			"- デザインのフィードバックと修正\n" +
			"- マーケティングチームとの協力とプロジェクト管理\n" +
			"また、WebデザインやUI/UXデザインの経験がある方を優遇します。創造的なアイデアとデザインに対する情熱が求められます。" +
			"さらに、複数のプロジェクトを同時に管理し、期限内に高品質なデザインを提供する能力が必要です。" +
			"最新のデザイントレンドを追求し、ブランドの一貫性を保ちながら、革新的なビジュアルソリューションを提供することが期待されます。",
	},
	{
		Title: "ネットワークエンジニア",
		Description: "ネットワークの設計、実装、管理、およびトラブルシューティングを行い、システムのパフォーマンスを最適化します。ネットワークセキュリティの知識、問題解決能力、チームとの協力が重要です。" +
			"具体的な業務内容は以下の通りです:\n" +
			"- ネットワークインフラの設計と導入\n" +
			"- ネットワークの監視とトラブルシューティング\n" +
			"- セキュリティ対策の実施と管理\n" +
			"- ネットワークのパフォーマンス最適化\n" +
			"- 新技術の導入と既存システムの改善\n" +
			"CCNAやCCNP資格保持者を優遇し、ネットワーク関連の経験が豊富な方を歓迎します。高度な技術知識と問題解決能力が求められます。" +
			"さらに、大規模ネットワークの運用経験があり、複雑な問題を迅速に解決できる能力が必要です。" +
			"最新のネットワーク技術を導入し、システムの信頼性とパフォーマンスを向上させるための継続的な改善に貢献します。",
	},
	{
		Title: "プランナー",
		Description: "業務プロセスの分析、改善提案の作成、システム要件の定義を行い、ビジネス効率を向上させます。分析スキル、データの解釈能力、優れたコミュニケーションスキルが必要です。" +
			"具体的な業務内容は以下の通りです:\n" +
			"- 現行業務プロセスの分析と改善提案\n" +
			"- システム要件の定義と文書化\n" +
			"- データ分析と業績評価の報告\n" +
			"- 部門横断的なプロジェクトの管理\n" +
			"- 経営陣への戦略的提案\n" +
			"MBA保持者を歓迎し、ビジネスインテリジェンスツールの使用経験がある方を優遇します。戦略的思考と実行力が求められます。" +
			"さらに、ビジネスプロセスの最適化と効率化に貢献し、組織全体のパフォーマンス向上をサポートします。" +
			"データに基づく洞察を提供し、経営陣と緊密に連携して、長期的なビジネス目標を達成するための戦略を策定します。",
	},
}
