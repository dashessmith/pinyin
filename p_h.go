package pingyin

// 拼音 --> 汉字, init 中会拆分去重排序, 转换成 "a" 的形式
var PH = map[string][]string{
	"a":      {"吖", "呵", "啊", "嗄", "腌", "錒", "锕", "阿"},
	"ai":     {"伌僾凒叆哀哎唉啀嗌嗳嘊噯埃堨塧壒娭娾嫒嬡嵦愛懓懝挨捱敱敳昹暧曖欬欸毐溰溾濭焥爱瑷璦癌皑皚皧瞹矮砹硋碍磑礙絠艾蔼薆藹諰譪譺賹躷銰鎄鑀锿隑隘霭靄靉餲馤鱫鴱"},
	"an":     {"侒俺偣儑厂厈咹唵啽垵埯堓媕安屵屽岸峖广庵按揞晻暗案桉氨洝犴玵痷盦盫罯胺腤荌菴萻葊蓭裺誝諳谙豻貋銨錌铵闇隌雸鞌鞍韽馣鮟鵪鶕鹌黯"},
	"ang":    {"仰卬岇昂昻枊盎肮醠骯"},
	"ao":     {"傲凹厫嗷嗸坳垇墺墽奡奥奧媪媼嫯岙岰嶅嶴廒慠懊扷抝拗摮擙敖柪梎泑滶澚澳熝熬爊獒獓璈眑磝翱翶翺聱艹芺蔜螯袄襖謷謸軪遨郩鏊鏕鏖镺隞隩驁骜鰲鳌鷔鼇"},
	"ba":     {"丷仈伯八叐叭吧哱哵坝垻壩夿妭岜峇巴弝扒扷把抜拔捌朳柭欛灞炦爸犮玐疤癹皅矲笆粑罢罷羓耙胈芭茇菝蚆覇詙豝跁跋軷釟鈀钯霸靶颰魃魞鮁鮊鲃鲅鲌鼥"},
	"bai":    {"伯佰呗唄庍拜拝挀捭掰摆擘擺敗柏栢猈瓸白百稗竡粨粺絔薭襬贁败韛"},
	"ban":    {"伴办半坂坢姅岅怑扮扳拌搫搬攽斑斒昄朌板柈湴版班瓣瓪瘢癍秚粄絆绊肦舨般蝂螌褩跘辦辬鈑鉡钣闆阪靽頒颁魬鳻"},
	"bang":   {"傍垹塝帮幇幚幫彭挷捠搒旁梆棒棓榜浜牓玤硥磅稖綁縍绑膀艕蒡蚄蚌蛖蜯謗谤邦邫鎊镑鞤髈"},
	"bao":    {"佨保儤刨剝剥勹勽包嚗堡堢報媬孢宝宲寚寳寶忁怉报抱暴曓曝枹瀑炮煲爆犦犳珤窇笣緥胞苞菢葆蕔薄藵虣蚫袌裦褒褓襃豹賲趵鉋鑤铇闁雹靌飹飽饱駂骲髱鮑鲍鳵鴇鸨齙龅"},
	"bei":    {"俻倍偝偹備僃勃北卑呗唄备孛庳悖悲惫愂憊揹昁杮杯柸桮梖棑棓椑焙牬犕狈狽珼琲盃碑碚禙箄糒背臂苝蓓藣蛽被褙誖諀貝贝軰輩辈邶郥鄁鉳鋇鐾钡陂鞁鞞鞴骳鵯鹎"},
	"ben":    {"倴喯坋坌夯夲奔奙捹撪本栟桳楍泍渀漰炃犇畚笨翉苯蟦賁贲逩錛锛"},
	"beng":   {"伻嗙嘣埄埲堋塴奟崩嵭揼泵琣琫甏甭甮痭祊絣綳繃绷菶蚌跰蹦迸逬鏰镚閍鞛"},
	"bi":     {"佊佖佛俾偪匕吡咇哔啚嗶坒堛壁夶奰妣妼娝婢嬖嬶屄崥币幣幤庇廦弊弻弼彃彼必怭怶愊愎拂敝斃旇朇朼柀柲梐楅榌比毕毖毙毞毴沘泌湢滗滭潷濞煏熚狴獘獙珌璧畀畁畐畢疕疪痹痺皀皕睤睥碧祕禆秕秘稫笓笔筆筚箅箆篦篳粃粊紕紴綼縪繴纰罼翍聛肶胇腷臂舭芘苾荜荸萆萞蓖蓽蔽薜蜌螕袐裨襅襞襣觱詖诐豍貏貱賁贔贲赑跸蹕躃躄辟逼避邲鄙鄨鄪鉍鎞鏎鐴铋閇閈閉閟闬闭陛鞞鞸韠飶饆馝駜驆髀髲魓鮅鰏鲾鵖鶝鷝鷩鼊鼻"},
	"bian":   {"便匾卞变変峅弁徧忭惼扁抃揙昪汳汴炞煸牑猵獱玣甂疺砭碥稨窆笾箯籓籩糄編緶缏编臱艑苄萹藊蝙褊覍覵諚變貶贬辡辧辨辩辫辮辯边辺遍邉邊釆鍽閞鞕鞭鯾鯿鳊鴘鶣"},
	"biao":   {"俵僄儦墂婊嫑幖彪摽杓标標檦淲滮瀌灬熛爂猋瘭磦穮篻脿膔膘臕蔈藨表裱褾諘謤贆錶鏢鑣镖镳颩颮颷飆飇飈飊飑飙飚驫骉骠髟鳔麃"},
	"bie":    {"別别彆徶憋柭柲瘪癟莂虌蛂蟞襒蹩鱉鳖鼈龞"},
	"bin":    {"份傧儐宾彬摈擯斌梹椕槟檳殡殯汃滨濒濱濵瀕玢瑸璸砏繽缤膑臏虨豩豳賓賔邠鑌镔霦頻顮频髌髕髩鬂鬓鬢"},
	"bing":   {"丙並仌仒併倂偋傡兵冫冰垪寎屏屛并幷庰怲抦掤摒昞昺枋柄栟栤梹棅槟檳氷炳琕病眪禀秉稟窉竝絣苪蛃誁邴鈵鉼鋲陃靐鞆鞞鞸餅餠饼鮩"},
	"bo":     {"亳仢伯佛侼僠僰剝剥勃博卜哱啵噃嚗壆孹嶓帛彴愽懪拨挬搏撥播擗擘柏柭桲榑檗檘欂殕泊波浡淿渤溊煿牔狛猼玻瓝瓟癶癷皪盋砵碆磻礡礴秡穛箔簙簸簿糪紴缽肑胉脖膊舶艊苩菠萡葧蔔薄蘗蚾袯袰袹襏襮誖謈譒趵跛踣蹳郣鈸鉑鉢鋍鎛鑮钵钹铂镈餑餺饽馎馛馞駁駮驋驳髆髉鮊鱍鲌鵓鹁"},
	"bu":     {"不佈勏卜卟吥咘哺埔埗埠堡婄峬布庯廍怖悑抪捕捗捬晡柨步歨歩瓿箁篰簿荹蔀补補誧踄轐逋部郶醭鈈鈽錻钚钸餔餢鯆鳪鵏鸔"},
	"ca":     {"嚓囃拆擦攃礤礸遪"},
	"cai":    {"倸偲啋埰婇寀彩才扐採材棌毝猜睬綵縩纔菜蔡裁財财跴踩采"},
	"can":    {"傪儏参參叄叅喰嘇囋嬠孱惨惭慘慙慚憯掺摲摻朁残殘湌澯灿燦爘璨穇篸粲薒蚕蝅蠶蠺謲鏒飡餐驂骖黪黲"},
	"cang":   {"仓仺伧倉傖凔匨嵢欌沧滄濸獊玱瑲篬臧舱艙苍蒼藏螥賶鑶鶬鸧"},
	"cao":    {"傮嘈屮嶆愺慅慒懆撡操曹曺槽漕糙肏艚艸草蓸螬褿襙鄵鏪騲鼜"},
	"ce":     {"侧側冊册厕厠夨廁恻惻憡拺敇栅测測畟笧策筞筴箣簎粣荝萗萴蓛"},
	"cen":    {"参參叄叅埁岑嵾汵涔穇笒篸"},
	"ceng":   {"噌层層嶒曽曾橧竲蹭驓"},
	"cha":    {"仛侘偛刹剎叉喳嗏嚓垞奼姹察岔嵖差扠扱挿捈插揷搽摖杈查査梌楂槎檫汊猹疀碴秅紁臿艖芆苴茬茶衩詧詫诧蹅銟鍤鎈鑔锸镲靫餷馇"},
	"chai":   {"侪儕喍囆差拆柴犲瘥祡肞芆茝虿蠆袃訍豺釵钗"},
	"chan":   {"丳产僝儃儳冁刬剗剷劖单啴單嘽嚵囅壥婵嬋孱崭嵼嶃嶄巉幝幨廛忏懴懺掺搀摌摲摻攙斺旵梴棎欃毚浐湹滻潹潺澶瀍瀺灛煘燀獑產産硟磛禅禪簅緂緾繟繵纏纒缠羼艬苂蒇蕆蝉螹蟐蟬蟾裧襜覘觇誗諂譂讇讒谄谗躔辿鄽酁醦鋋鋓鏟鑱铲镵閳闡阐韂顫颤饞馋骣"},
	"chang":  {"仧仩伥倀倘倡偿僘償兏厂厰唱嘗嚐场場塲娼嫦尝常廠徜怅悵惝敞昌昶暢椙氅淌淐焻猖玚琩瑒瑺瓺甞畅畼肠腸膓苌菖萇裮裳誯鋹錩锠長镸长閶闛阊韔鬯鯧鱨鲳鲿鼚龦"},
	"chao":   {"仦仯剿劋勦吵嘲巐巢巣弨怊抄摷晁朝槱樔欩漅潮炒焣焯煼牊眧窲窼粆綽绰罺耖觘訬謿超轈鄛鈔钞麨鼂鼌"},
	"che":    {"伡俥偖勶呫唓坼奲尺屮彻徹扯掣揊摰撤撦澈烢烲焎爡瞮砗硨硩聅莗蛼車车迠頙"},
	"chen":   {"伧偁傖儭嗔嚫塵墋夦宸尘忱愖抻捵敐敶晨曟栕桭梣棽榇樄櫬沈沉煁爯琛疢瘎瘨瞋硶碜磣称稱穪綝肜臣茞莀莐蔯薼螴衬襯訦諃諶謓讖谌谶賝贂趁趂趻跈踸軙辰迧郴醦鈂鍖陈陳霃鷐麎齓齔龀"},
	"cheng":  {"丞乗乘侱偁僜净凈呈城埕埩堘塍塖娍宬峸庱徎悜惩憆憕懲成承抢挰掁揨搶摚撐撑撜晟晿朾枨柽棖棦椉橕橖橙檉檙氶泟洆浧浾淨溗澂澄瀓爯牚珵珹琤畻盛睈瞠矃碀秤称程稱穪窚竀筬絾緽罉脭荿蛏蟶裎誠诚赪赬逞郕酲鋮鎗鏿鐣鐺铖铛阷靗頳饓騁騬骋鯎"},
	"chi":    {"伬佁侈侙俿傺勅勑匙卶叱叺吃呎哧啻喫嗤噄噭坻垑墀奓妛媸尺弛彨彲彳徲恜恥慗慸憏懘扡抶拸持捇摛攡敕斥杘柅柢樆欼歭歯殦池沶湁漦灻炽烾熾狋瓻痴痸瘈瘛癡眙眵瞝竾笞筂箈箎篪粚絺翄翅翤翨耛耻肔胣胵腟茌荎菭蚇蚩蚳螭袲袳裭褫訵誃誺謘豉貾赤赿趩跢踟迟迡遅遟遫遲邌鉓鉹銐鍉雴飭饬馳驰魑鴟鵄鶒鶗鶙鷘鸱麶黐齒齝齿"},
	"chong":  {"充冲嘃埫宠寵崇崈徸忡憃憧揰摏沖浺涌漴爞珫痋种緟罿翀舂艟茺虫蝩蟲衝褈蹖重銃铳隀"},
	"chou":   {"丑丒仇俦偢儔吜嚋婤嬦帱幬怞惆愁懤抽掫搊杻杽栦梼椆檮殠燽牰犨犫畤畴疇瘳皗盩瞅矁稠筹篘籌紬絒綢绸臭臰菗薵裯詶讎讐诪跾踌躊遚酧酬醔醜醻雔雠雦魗鮘"},
	"chu":    {"亍俶傗储儊儲処出刍初助厨嘼埱处媰岀幮廚怵慉憷搐摢摴敊斶杵椘楚楮榋樗橱橻檚櫉櫖櫥欪歜泏滀滁濋炪犓珿璴畜矗础礎禇竌竐篨絀绌耝耡臅芻菆蒢蒭藸處蜍蟵触觸諔豖豠貙趎蹰躇躕鄐鉏鋤锄閦除雏雛鶵黜齣齭齼"},
	"chua":   {"欻"},
	"chuai":  {"啜揣搋欼膗膪踹"},
	"chuan":  {"串传傳僢剶喘圌堾巛川暷椯椽歂氚汌猭玔瑏穿篅膞舛舡舩船荈賗踳輲遄釧钏鶨"},
	"chuang": {"傸刅创刱剏剙創喠噇幢床怆愴摐朣橦牀牎牕疮瘡磢窓窗窻闖闯"},
	"chui":   {"倕吹圌垂埀惙捶搥桘棰椎槌炊硾箠腄菙錘鎚锤陲顀龡"},
	"chun":   {"偆唇媋惷旾春暙朐杶椿槆橁櫄浱淳湻滣漘犉瑃睶箺純纯肫脣膞芚莼萅萶蒓蓴蝽蠢賰輴醇醕錞陙鯙鰆鶉鶞鹑"},
	"chuo":   {"吷哾啜嚽娕娖婥婼孎惙戳拺擉歠涰淖磭箹綽繛绰趠踀踔輟辍辵辶逴酫鋜鏃鑡餟齪龊"},
	"ci":     {"伺佌佽偨兹刺刾呰呲啙垐嬨差庛慈朿柌栨次此泚濨玼珁瓷甆疵皉磁礠祠粢糍絘縒茦茨莿薋蛓螆蠀詞词賜赐趀跐辝辞辤辭雌飺餈骴髊鮆鴜鶿鷀鹚齹"},
	"cong":   {"丛从匆叢囪囱婃孮従徔徖從忩怱悤悰愡慒憁暰枞棇楤樅樬樷欉淙漎漗潀潈潨灇焧熜爜琮瑽璁瞛碂篵繱聡聦聪聰苁茐葱蓯蔥藂蟌誴謥賨賩錝鍯鏓鏦騘驄骢"},
	"cou":    {"傶凑楱湊腠輳辏"},
	"cu":     {"促卒噈媨徂怚憱橻殂猝瘄瘯皻簇粗縬脨蔟觕誎趋趗趣趨踧踿蹙蹴蹵醋顣麁麄麆麤鼀"},
	"cuan":   {"巑撺攒攛攢櫕欑殩汆濽灒熶爨穳窜窽窾竄篡簒蹿躥鋑鑹镩"},
	"cui":    {"伜倅催凗啐啛墔崒崔嵟忰悴慛摧椊榱槯毳淬漼濢焠熣琗璀疩瘁皠磪竁粋粹紣綷縗缞翆翠脃脆脺膬膵臎萃衰襊趡鏙隹顇"},
	"cun":    {"侟刌吋壿存寸忖拵村澊皴竴籿膥邨"},
	"cuo":    {"剉剒厝夎嵯嵳庴挫措搓撮瑳痤瘥睉矬磋縒脞莝莡蒫蓌蔖虘諎蹉逪遳酂醝銼錯鎈锉错鹺鹾"},
	"da":     {"亣剳匒呾咑哒嗒噠垯墶大妲怛惮憚打搭撘汏沓溚炟燵畗畣瘩眔笚笪答繨羍耷荅荙薘蟽褡詚跶达迏迖迚逹達鎉鎝鐽靼鞑韃龖龘"},
	"dai":    {"代侢傣叇呆呔垈埭大岱帒带帯帶廗待怠懛戴曃柋棣歹殆毒瀻獃玳瑇甙箉簤紿緿绐艜蚮蝳袋襶詒诒貣貸贷跢蹛軑軚軩轪迨逮釱霴靆駘鮘鴏黛黱"},
	"dan":    {"丹亶伔但僤儋刐勯匰单単呾唌啖啗啿單嘾噉嚪妉媅帎弹弾彈惔惮愖憚憺抌担掸撢撣擔旦柦殚殫氮沊泹淡澹燀狚玬瓭甔疍疸瘅癉癚皽眈石砃禫窞箪簞紞繵耼耽聃聸胆腅膻膽萏蓞蛋蜑衴褝襌觛訑誕诞賧贉赕躭郸鄲酖醈霮頕餤饏馾駳髧鴠黕黮黵"},
	"dang":   {"偒儅党凼噹圵垱壋婸宕崵嵣当愓挡擋攩档檔欓氹潒澢灙珰璗璫瓽當盪瞊砀碭礑筜簜簹艡荡菪蕩蘯蟷裆襠譡讜谠趤逿鐺铛闣雼黨"},
	"dao":    {"倒刀刂到叨嘄噵壔导導岛島嶋嶌嶹帱幬忉悼捣捯搗擣朷梼椡槝檤檮氘焘燾瓙盗盜祷禂禱稲稻箌絩纛翢翿舠艔菿螩衜衟裯蹈軇辺道釖陦隝隯魛鱽"},
	"de":     {"嘚地底得徳德恴悳惪旳棏淂的蚮鍀锝"},
	"dei":    {"嘚得"},
	"den":    {"扥扽"},
	"deng":   {"僜凳噔墱嬁嶝憕戥櫈澄瀓灯燈璒登瞪磴竳等簦艠覴豋蹬邓鄧鐙镫隥"},
	"di":     {"仾低俤偙僀厎呧哋唙啇啲嘀嚁地坔坘坻埅埊埞堤墆墑墬奃娣媂嫡岻嵽嶳帝底廸弟弤彽怟慸扚抵拞掋提揥摕敌敵旳杕枑枤柢梊梑棣楴樀氐浟涤渧滌滴焍牴狄玓珶甋疐的眱睇砥碮碲磾祶禘秪笛第篴籴糴締缔羝翟聜肑腣苐苖荻菂菧蒂蓧蔋蔐蔕藡蝃蝭螮袛覿觌觝詆諟諦诋谛豴趆踶蹢軧迪递逓遆遞遰邸釱鉪鏑镝阺隄靮鞮頔馰骶髢鯳鸐"},
	"dia":    {"嗲"},
	"dian":   {"佃佔傎典厧嚸坫垫墊壂奌奠婝婰嵮巅巓巔店惦扂掂攧敁敟椣槇槙橂橝殿淀滇澱点玷琔电甸痁痶瘨癜癫癲碘磹窴簟蒧蕇蜔跕踮蹎鈿钿阽電靛顚顛颠驔點齻"},
	"diao":   {"伄凋刁刟叼吊奝屌弔弴彫扚掉椆汈琱瘹盄瞗碉窎窵竨簓絩蓧藋虭蛁訋誂調调貂釣鈟銚銱鋽錭鑃钓铞铫雕雿魡鮉鯛鲷鳥鳭鵰鼦"},
	"die":    {"佚叠咥哋啑喋垤堞峌崼嵽幉怢恎惵戜挕揲昳曡柣楪槢殜氎泆爹牃牒瓞畳疂疉疊眣眰碟絰绖耊耋胅臷艓苵蜨蝶螲褋褺詄諜谍趃跌跕跮蹀軼迭镻鞢鰈鲽鴩"},
	"ding":   {"丁仃叮啶奵定嵿帄忊掟椗濎玎町甼疔盯矴碇碠磸耵聢腚艼萣薡虰蝊訂订酊釘鋌錠鐤钉铤锭靪頂顁顶飣饤鼎鼑"},
	"diu":    {"丟丢銩铥颩"},
	"dong":   {"东侗倲働冬冻凍动動勭咚垌埬墥姛娻嬞岽峒崠崬徚恫懂戙挏揰昸東栋棟氡氭洞涷湩烔燑狪硐笗箽絧胨胴腖苳菄董蕫蝀諌迵霘駧鮗鯟鶇鶫鸫鼕"},
	"dou":    {"乧侸兜兠吋吺唗唞抖斗斣枓梪橷毭浢渎瀆狵痘瞗窦竇篼脰荳蔸蚪读豆逗郖都酘鈄鋀钭閗闘阧陡餖饾鬥鬦鬪鬬鬭"},
	"du":     {"凟剢匵厾嘟堵妒妬姤嬻度斁晵暏杜椟櫝殬殰毒涜渎渡瀆牍牘犊犢独獨琽瓄皾督睹秺竺笃篤肚芏荰螙蠧蠹裻覩読讀讟读豄賭贕赌都醏錖鍍鑟镀闍阇靯韇韣韥騳髑黩黷"},
	"duan":   {"偳剬塅媏断斷椴段毈煅瑖短碫端簖籪緞缎耑腶葮褍躖鍛鍴锻"},
	"dui":    {"兊兌兑啍垖堆塠对対對怼憝憞懟敓敚敦杸濧濻瀢瀩痽碓磓祋綐薱譈譵轛追鐓鐜镦队陮隊頧鴭"},
	"dun":    {"伅吨噸囤坉墩墪墫庉忳惇憞撉撴敦楯橔沌潡炖燉犜獤盹盾砘碷礅蜳趸踲蹲蹾躉逇遁遯鈍钝頓顿驐"},
	"duo":    {"亸仛凙刴剁剟剫咄哆哚喥嚉嚲垛垜埵堕墮墯多夛夺奪奲媠尮崜度惰憜挅挆掇敓敚敠敪朵朶杕杝枤柁柂柮桗椯毲沰沲畓痥硾綞缍舵袳裰趓跢跥跺踱躱躲軃鈬鐸铎陊陏隋隓隳飿饳馱駄鬌鵽"},
	"e":      {"俄偔僫匎卾厄吪呃呝咢咹哦噁噩囮垩堊堨堮妸妿姶娥娾娿婀屙岋峉峨峩崿廅恶悪惡愕戹扼搕搤搹擜敋枙櫮歞歺洝涐湂珴琧痷痾皒睋砈砐砨硆磀礘腭苊莪萼蕚蚅蛾蝁覨訛誐諤譌讍讹谔豟軛軶轭迗遌遏遻鄂鈋鈪鋨鍔鑩锇锷閼阏阨阸阿頋頞頟額顎颚额餓餩饿騀魤鰐鰪鱷鳄鵈鵝鵞鶚鹅鹗齃齶"},
	"ei":     {"欸誒诶"},
	"en":     {"奀峎恩摁煾蒽"},
	"er":     {"二佴侕儿児兒刵咡唲尒尓尔峏弍弐杒栭栮梕樲毦洏洱爾珥粫而耏耳聏胹荋薾衈袻誀貮貳贰趰輀轜迩邇鉺铒陑隭餌饵駬髵髶鮞鲕鴯鸸"},
	"fa":     {"乏伐佱傠冹发坺垡墢姂峜彂栰橃沷法浌灋珐琺疺発發瞂砝笩筏罚罰罸茷蕟藅醗醱鍅閥阀髪髮"},
	"fan":    {"仮凡凢凣勫匥反噃墦奿婏嬎嬏帆幡忛憣払拚旙旛杋柉桳梵棥樊橎氾汎泛滼瀪瀿烦煩燔犯犿璠畈畨番盕矾礬笲笵範籓籵緐繁繙羳翻膰舤舧范蕃薠藩蘩蠜襎訉販贩蹯軓軬輽轓返釩鐇鐢钒颿飜飯飰饭鱕鷭"},
	"fang":   {"仿倣匚坊埅堏妨彷房放方旊昉昘枋汸淓牥瓬眆眪祊紡纺肪舫芳蚄訪访趽邡鈁錺钫防髣魴鲂鴋鶭"},
	"fei":    {"俷剕匪厞吠啡奜妃婓婔屝废廃廢悱扉斐昲暃曊朏杮柹棐榧橃橨櫠沸淝渄濷犻狒猆疿痱癈砩笰篚緋绯翡肥肺胇胏胐腓芾菲萉蕜蜚蜰蟦裶誹诽費费鐨镄陫霏靅非靟飛飝飞餥馡騑騛鯡鲱鼣"},
	"fen":    {"份偾僨分吩坆坋坟墳奋奮妢岎帉幩弅忿愤憤昐朆朌枌梤棻棼橨氛汾濆瀵炃焚燌燓獖玢瞓砏秎竕粉粪糞紛纷羒羵翂膹芬蒶蕡蚠蚡衯訜豮豶躮轒酚鈖鐼隫雰餴饙馚馩魵鱝鲼鳻黂黺鼖鼢"},
	"feng":   {"丰仹俸偑僼冯凤凨凬凮唪埄堸堼夆奉妦寷封峯峰崶摓枫桻楓檒沣沨浲渢湗溄漨灃炐烽焨煈熢犎猦琒疯瘋盽砜碸篈綘縫缝艂莑葑蘴蜂蠭覂諷讽豐賵赗逢鄷酆鋒鎽鏠锋霻靊風飌风馮鳯鳳鴌麷"},
	"fiao":   {"覅"},
	"fo":     {"仏仸佛坲梻"},
	"fou":    {"不否妚殕炰紑缶缹缻雬鴀"},
	"fu":     {"乀乶付伏伕佛俌俘俛俯偩傅冨凫刜副匐呋呒咈咐哹嘸坿垘垺复夫妇妋姇娐婏婦媍嬔孚孵宓富尃岪峊巿帗幅幞府弗弣彿復怤怫懯払扶抙抚拂拊捊捬撫敷斧旉服枎枹柎柫柭栿桴棴椨椱榑氟泭洑浮涪滏澓炥烰焤父玞玸琈璷甫甶畉畐畗癁盙砆砩祓祔福禣秿稃稪竎符笰筟箁箙簠粰糐紨紱紼絥綍綒緮縛纀绂绋缚罘罦翇肤胕脯腐腑腹膚艀艴芙芣芾苻茀茯荂荴莩菔萯葍蓲蕧虙蚥蚨蚹蛗蜅蜉蝜蝠蝮衭袚袝袱複褔襆襥覄覆訃詂諨讣豧負賦賻负赋赙赴趺跗踾輔輹輻辅辐邞郙郛鄜酜釜釡鈇鉘鉜鍑鍢阜阝附韍韨頫颫颰馥駙驸髴鬴鮄鮒鮲鰒鲋鳆鳧鳬鳺鴔鵩鶝麩麬麱麸黻黼"},
	"ga":     {"伽咖嘎嘠噶夹夾尕尜尬旮朒玍軋釓錷钆魀"},
	"gai":    {"丐乢侅匃匄垓姟峐忋戤摡改晐杚概槩槪溉漑瓂畡盖祴絠絯芥荄葢蓋該该豥賅賌赅郂鈣钙陔隑"},
	"gan":    {"乾仠佄倝凎凲咁坩尲尴尶尷干幹忓感擀攼敢旰杆柑桿榦橄檊汵泔淦漧澉澸灨玕玵甘疳皯盰矸秆稈竿筸篢簳粓紺绀肝芉苷衦詌贑贛赣赶趕迀酐骭魐鱤鳡鳱龫"},
	"gang":   {"冈冮刚剛堈堽岗岡崗戅戆戇扛掆摃杠棡槓港焵焹牨犅犺疘矼碙筻綱纲缸罁罡肛釭鋼鎠钢"},
	"gao":    {"勂吿告夰峼搞杲槀槁槔槹橰檺櫜獋皋皐睪睾祮祰禞稁稾稿筶篙糕縞缟羔羙膏臯菒藁藳誥诰郜鋯鎬锆镐韟餻高髙鷎鷱鼛"},
	"ge":     {"个亇仡佫佮個割匌各合吤呄咯哥哿嗝嗰嘅圪塥彁愅戈戓戨扢挌搁搿擱敋格槅櫊歌滆滒牫犵猲獦疙盖硌笴箇紇纥肐胳膈臈臵舸茖葛蓋虼蛒蛤袼裓觡諽謌輵轕鉻鎑鎘鎶铬镉閣閤阁隔革鞈鞷韐韚颌饹騔骼鬲鮥鮯鲄鴚鴿鸽"},
	"gei":    {"給给"},
	"gen":    {"亘亙哏揯搄根艮茛跟"},
	"geng":   {"刯哽埂堩峺庚挭更梗椩浭焿畊絙絚綆緪縆绠羮羹耕耿莄菮賡赓郠颈骾鯁鲠鶊鹒"},
	"gong":   {"侊供公共功匑匔厷唝嗊塨宫宮工巩幊廾弓恭愩慐拱拲攻杛栱汞澒熕珙碽篢糼糿紅红羾肱蚣觥觵貢贡躬躳輁銾鞏髸龏龔龚龷"},
	"gou":    {"佝傋冓勾句呴啂坸垢够夠姤媾岣彀拘搆撀构枸構沟溝煹狗玽痀笱篝簼糓緱缑耇耈耉芶苟蚼袧褠覯觏訽詬诟豿購购遘鈎鉤钩雊鞲韝"},
	"gu":     {"估傦僱凅古呱咕唂唃啒嗗嘏固堌夃姑嫴孤尳崓崮巬巭怘愲扢抇故杚柧梏棝榖榾橭櫎毂汩沽泒淈濲瀔牯牿痼皷皼盬瞽磆祻稒穀笟箍箛篐糓縎罛罟羖股脵臌苽菇菰蓇薣蛄蛊蠱觚詁诂谷賈贾軱軲轂轱辜逧酤鈲鈷錮钴锢雇顧顾餶馉骨鮕鯝鲴鴣鶻鸪鹘鼓鼔"},
	"gua":    {"冎刮剐剮劀卦叧呙呱咼啩坬寡括挂掛栝桰歄煱瓜絓緺罣罫胍袿褂詿諣诖趏踻銛銽铦颪颳騧鴰鸹"},
	"guai":   {"乖叏夬怪恠拐枴柺箉罫"},
	"guan":   {"丱倌关冠卝官悹悺惯慣掼摜斡棺樌欟毌泴涫淉潅灌爟琯璭瓘痯瘝癏盥矔矜矝礶祼窤筦管綸纶罆罐舘莞蒄覌観觀观貫贯躀輨遦錧鏆鑵閞関闗關雚館馆鰥鱞鱹鳏鳤鸛鹳"},
	"guang":  {"俇僙光咣垙姯广広廣挄撗桄洸潢灮炗炚炛烡犷獷珖硄胱臦臩茪輄逛銧黆"},
	"gui":    {"亀佹傀刽刿劊劌匦匭匮匱厬圭垝妫姽媯嫢嬀宄嶡嶲巂帰庋庪廆归恑摫撌攰攱昋晷朹柜桂桧椝椢槣槶槻槼樻檜櫃櫰櫷歸氿湀溎炅猤珪瑰璝瓌癐癸皈眭瞆瞡瞶硅硊祪禬窐筀簂簋胿茥蓕蘬蛫螝蟡袿襘規规觤詭诡貴贵趹跪軌軓轨邽郌鐀鑎閨闺陒雟鞼騩鬶鬹鬼鮭鱖鱥鲑鳜鳺龜龟"},
	"gun":    {"丨惃棍滚滾睔硍磙緄緷绲蓘蔉衮袞裷謴輥辊鮌鯀鲧"},
	"guo":    {"呙咼啯喐嘓囗囯囶囻国圀國埚堝墎崞帼幗彉彍惈慖懖掴摑敋果椁楇槨櫎涡淉渦漍濄猓瘑矌簂粿綶聒聝腂腘膕菓蔮虢蜾蝈蟈裹褁輠过過郭鈛鍋鐹锅餜馃馘"},
	"ha":     {"哈奤紦虾蛤蝦鉿铪"},
	"hai":    {"亥咍咳嗐嗨嚡塰妎孩害拸氦海烸猲絯胲还還郂酼醢餀饚駭駴骇骸"},
	"han":    {"丆仠佄傼兯函凾厂厈含哻唅喊圅垾娢嫨寒屽崡嵅忓悍憨憾扞捍撖撼攼旰旱晗晘暵梒椷歛汉汗汵浛浫涆涵漢澏瀚焊焓熯爳犴猂琀甝皔睅笒筨罕翰肣莟菡蔊虷蚶蛿蜬蜭螒譀谽豃輚邗邯酣釬銲鋎鋡閈闞闬雗韓韩頇頷顄顸颔馠馯駻魽鳱鶾鼾"},
	"hang":   {"吭垳夯妔巷忼斻杭桁沆笐筕絎绗航苀蚢行裄貥迒邟頏颃魧"},
	"hao":    {"乚侴傐儫号呺哠嗥嘷噑嚆嚎壕好恏悎昊昦晧暠暤暭曍椃毜毫浩淏滈滜澔濠灏灝獆獋獔皋皓皜皞皡皥秏竓籇耗聕茠蒿薃薅薧藃號蚝蠔諕譹豪貉郝鄗鎬镐顥颢鰝"},
	"he":     {"何劾合吓呙呵咊和咼哬啝喛喝嗃嗬噈嚇垎壑姀寉峆惒抲敆曷朅柇核楁欱毼河涸渮澕焃煂熆熇燺爀狢猲癋皬盇盉盍盒盖碋礉禾篕籺粭紇纥翮翯苛荷菏萂蓋螛蠚袔褐覈訶訸詥謞诃貈貉賀贺赫郃釛鉌鑉閡閤闔阂阖隺靍靎靏鞨頜颌餄饸魺鲄鶡鶮鶴鸖鹖鹤麧齕龁龢"},
	"hei":    {"嗨嘿潶黑黒"},
	"hen":    {"佷哏噷很恨拫狠痕詪鞎"},
	"heng":   {"亨哼啈姮恆恒悙撗桁横橫涥烆珩胻脝蘅衡鑅鴴鵆鸻"},
	"hong":   {"仜厷叿吰吽呍哄唝嗊嚝垬妅娂宏宖屸巆弘彋愩慐揈撔晎汯沗泓洪浤浲渱渹溄潂澒灴烘焢玒玜瓨硔硡竑竤篊粠紅紘紭綋红纮羾翃翝耾苰荭葒葓蕻薨虹訇訌讧谹谼谾軣輷轟轰鈜鉷銾鋐鍧閎閧闀闂闳霐霟鞃鬨魟鴻鸿黉黌"},
	"hou":    {"侯候厚后吼吽呴喉垕堠帿後洉犼猴瘊睺矦篌糇翭翵茩葔豞逅郈鄇銗鍭餱骺鮜鯸鱟鲎鲘齁"},
	"hu":     {"乎乕乯互俿冱匢匫呼和唬唿喖嗀嗃嘑嘝嚛囫垀壶壷壺姱婟媩嫭嫮寣岵帍幠弖弧忽怘怙恗惚戏戯戱戲戶户戸戽扈抇护搰摢斛昈昒曶枑核楛楜槲槴歑汻沍沪泘浒淲淴湖滬滸滹濩瀫烀焀煳熩狐猢琥瑚瓠瓡瓳礐祜笏箶簄粐糊絗綔縎縠胡膴芐芔芴苸萀葫蔛蔰虍虎虖虝蝴螜衚觳觷謼護许豰軤轷鄠醐錿鍙鍸隺雐雽頀頶餬鬍魱鯱鰗鱯鳠鳸鵠鶘鶦鶮鶻鸌鹄鹕鹱"},
	"hua":    {"划劃化华吪呚哗嘩埖夻姡婲婳嫿嬅崋搳摦撶杹枠桦椛槬樺滑澅猾画畫畵砉硴磆糀繣舙花芲華蒊蕐蘤螖觟話誮諙諣譁譮话豁釪釫鋘錵鏵铧驊骅鷨黊"},
	"huai":   {"佪咶坏壊壞徊怀懐懷槐櫰淮瀤耲蘹蘾褢褱踝龨"},
	"huan":   {"唤喚喛嚾圜垸堚奂奐嬛孉宦寏寰峘嵈幻患愌懁懽换換擐攌桓梙槵欢歓歡洹浣涣渙漶澣澴烉焕煥犿狟獂獾环瑍環瓛痪瘓睆瞏糫絙綄緩繯缓缳羦肒荁萈萑蒝藧讙豢豲貆貛轘还逭還郇酄鍰鐶锾镮闤阛雈雚驩鬟鯇鯶鰀鲩鴅鵍鹮"},
	"huang":  {"偟兤凰喤堭塃墴奛媓宺崲巟幌徨怳恍惶愰慌揘晃晄曂朚楻榥櫎汻湟滉潢炾煌熀熿爌獚瑝璜癀皇皝皩磺穔篁簧縨肓艎荒葟蝗蟥衁詤諻謊谎趪遑鍠鎤鐄锽隍韹餭騜鰉鱑鳇鷬黃黄"},
	"hui":    {"会佪僡儶匯卉咴哕喙嘒噅噕噦嚖囘回囬圚堕墮婎媈嬒孈寭屶屷幑廆廻廽彗彙彚徻徽恚恛恢恵悔惠慧憓懳拻挥揮撝晖晦暉暳會桧椲楎槥橞檅檓檜櫘殨毀毁毇汇泋洃洄浍湏溃滙潓潰澮濊瀈灰灳烠烣烩煇燬燴獩珲琿璤璯痐瘣睢睳瞺禈秽穢篲絵繢繪绘缋翙翚翬翽芔茴荟蔧蕙薈薉藱蘳虺蚘蛔蛕蜖螝蟪袆褘詯詼誨諱譓譭譿讳诙诲豗賄贿輝辉迴逥銊鏸鐬闠阓隓隳靧頮顪颒餯鮰鰴麾"},
	"hun":    {"俒倱圂婚婫尡忶惛慁掍昏昬棍棔殙浑涽混渾湣湷溷焄焝珲琿眃睧睯睴荤葷蔒觨諢诨轋閽阍餛馄魂鼲"},
	"huo":    {"伙佸俰剨劐吙和咟喐嚄嚯嚿夥奯姡彟彠惑或捇掝擭攉旤曤檴沎活湱漷濊濩瀖火灬煷獲癨眓矆矐礊祸禍秮秳秴穫窢篧耠耯臛艧获蒦藿蠖謋豁貨货趏邩鈥鍃鑊钬锪镬閄雘霍靃韄騞"},
	"ji":     {"丌丮乁乩亟亼亽伋伎佶偈偮僟其兾冀几击刉刏剂剞剤劑勣卙即卽及叝叽吉咭哜唧唶喞嗘嘰嚌圾坖垍基堲塈塉墼奇妀妓姞姬嫉季寂寄尐屐屰岌嵆嵇嵴嶯己帺幾庴廭彐彑彶徛忌忣急悸惎愱憿懠懻戟戢技挤掎揤摖撃撠撽擊擠攲敧旡既旣暨暩曁朞期机极枅梞棘楖楫極槉槣樭機橶檕檝檵櫅櫭殛毄汥汲泲洎济淁済湒漃漈潗激濈濟瀱焏犄犱玑璣璾畟畸畿疾痵瘠癠癪皀皍瞉矶磯磼祭禝禨积稘稩稷稽穄穊積穖穧笄笈筓箕箿簊簎籍系紀紒級給継綨緁緝績繋繫繼级纪给继绩缉罽羁羇羈耤耭肌脊膌臮艥芨芰苙茍茤荠葪蒺蓟蔇蕀蕺薊薺藉蘎蘮蘻虀虮蝍螏蟣裚襀襋覉覊覬觊觙觭計記誋諅諔譏譤计讥记诘谻賫賷赍趌趞跡跻跽踑踖蹐蹟躋躤躸輯轚辑迹郆銈銡錤鍓鏶鐖鑇鑙际際隮集雞雧霁霵霽鞊鞿韲飢饑饥驥骥髻鬾魕魝魢魥鮆鯚鯽鰶鰿鱀鱭鱾鲚鲫鳮鵋鶏鶺鷄鷑鸄鸡鹡麂齌齎齏齐齑龮"},
	"jia":    {"乫价伽佳假傢價加叚哿唊嘉圿埉夹夾婽嫁宊家岬幏徍徦忦恝戛戞扴抸拁拮挟挾擖斚斝架枷梜椵榎榢槚檟毠泇浃浹犌猰猳玾珈甲痂瘕稼笳糘耞胛脥腵荚莢葭蛱蛺袈袷裌豭貑賈贾跏迦郏郟鉀鉫鋏鎵钾铗镓頬頰颉颊駕驾鴐鴶鵊麚"},
	"jian":   {"件侟俭俴倹健傔僣僭儉兼冿减剑剣剪剱劍劎劒劔囏囝坚堅堑堿墹奸姦姧寋尖帴幵建弿彅徤惤戋戔戩戬拣挸捡揀揃揵搛撿擶攕旔暕朁枧柬栫梘检検椷椾楗榗槛樫橏檢檻櫼歼殱殲毽洊浅涧渐減湔湕溅漸澗濺瀐瀳瀸瀽煎熞熸牋牮犍猏玪珔瑊瑐监監睑睷瞯瞷瞼硷碊碱磵礀礆礛礷稴笕笺筧简箋箭篯簡籈籛糋絸緘縑繝繭缄缣翦聻肩腱臶舰艦艰艱茧荐菅菺葌蒹蔪蕑蕳薦藆虃螹蠒袸裥襇襉襺見覵覸见詃諓諫謇謭譛譼譾谏谫豜豣賎賤贱趝趼跈践踐踺蹇轞釰釼鉴鋄鋑鋻鍳鍵鎫鏩鐗鐧鐱鑑鑒鑬鑯鑳锏键閒間间雃靬鞬鞯韀韉餞餰饯馢鬋鰎鰔鰜鰹鲣鳒鳽鵑鵳鶼鹣鹸鹻鹼麉"},
	"jiang":  {"傋僵勥匞匠塂壃夅奖奨奬姜将將嵹弜弶強强彊摪摾桨槳橿櫤殭江洚浆港滰漿犟獎畕畺疅疆礓糡糨絳繮绛缰翞耩膙茳葁蒋蔣薑螀螿講謽讲豇酱醤醬降韁顜鱂鳉"},
	"jiao":   {"交佼侥僥僬儌剿劋勦叫呌喬嘂嘄嘐嘦噍噭嚼姣娇嬌嬓孂峤峧嶕嶠嶣徺徼恔悎憍憢憿挍挢捁搅摎摷撟撹攪敎教敫敽敿斍斠晈暞曒校椒樔浇湫湬滘漖潐澆灂灚烄焦煍燋燞狡獥珓璬癄皎皦皭矫矯礁穚窌窖笅筊簥絞繳纐绞缴胶脚腳膠膲臫艽茭茮蕉藠虠蛟蟜蟭覐覚覺觉角訆譑譥賋趭跤踋較轇轎轿较郊酵醮釂釥鉸鐎铰餃饺驕骄鮫鱎鲛鵁鷦鷮鹪龣"},
	"jie":    {"丯介价借倢偈偼傑價刦刧刼劫劼卩卪吤哜唶啑喈喼嗟嚌堦堺她妎姐婕媎媘媫嫅孑家尐屆届岊岕崨嵑嵥嶰嶱嶻巀幯庎徣悈戒截扢拮捷接掲掶揭搩擑擮擳斺昅杢杰桀桔桝椄楐楬楶楷榤檞毑洁洯湝滐潔煯犗狤玠琾瑎界畍疌疖疥痎癤皆睫砎碣礍秸稭竭節紒結絜结羯耤脻节芥莭菨蓵藉蚧蛣蛶蜐蝔蠘蠞蠽街衱衸袓袷袺裓褯解觧訐詰誡誱謯讦诘诫趌跲踕迼鉣鍻鎅镼阶階鞂鞊頡颉飷骱魪鮚鲒鶛鶡"},
	"jin":    {"仅今仐伒侭僅僸儘凚劤劲勁卺厪唫噤嚍埐堇堻墐壗妗婜嫤嬐嬧寖尽巹巾廑惍慬搢斤晉晋枃榗槿歏殣津浕浸溍漌濅濜烬煡燼珒琎琻瑨瑾璡璶盡矜矝砛祲禁竻笒筋紟紧緊縉缙臸荕荩菫菳蓳藎衿襟覲觐觔謹谨賮贐赆近进進金釒釿錦钅锦靳饉馑馸鹶黅齽"},
	"jing":   {"丼井京亰俓倞傹儆兢净凈刭剄劤劲勁坓坕坙境妌婙婛婧宑巠幜弪弳径徑惊憬憼擏敬旌旍景晶暻曔桱梷橸殌殑汫汬泾浄涇淨澋瀞猄獍璄璟璥痉痙睛秔稉穽竞竟竧竫競竸箐粇粳精経經綡经聙肼胫脛腈茎荆荊莖菁葏蟼誩警踁迳逕鋞鏡镜阱靓靖静靚靜頚頸颈驚鯨鲸鵛鶁鶄麖麠鼱"},
	"jiong":  {"侰僒冂冋冏囧坰埛扃扄泂浻澃炅炯烱煚煛熲燛窘絅綗臦臩蘏蘔褧迥逈銄顈颎駉駫"},
	"jiu":    {"丩久乆九乣倃僦剹勼匓匛匶厩咎啾奺就廄廏廐慦捄揂揪揫摎救旧朻杦柩柾桕樛欍殧氿汣灸牞玖疚稵究糺糾紤纠臼舅舊舏萛赳酒镹阄韭韮鬏鬮鯦鳩鷲鸠鹫麔齨"},
	"ju":     {"且举乬伡侷俥俱倨倶僪具凥剧劇勮匊句咀啹埧埾壉姖娵婅婮寠局居屦屨岠岨崌巈巨巪弆忂怇怐怚惧愳懅懼抅拒拘拠挙挶据掬揟據擧昛枸柜桔梮椇椈椐榉榘橘檋櫸欅歫毩毱沮泃泦洰涺淗渠湨澽炬烥焗焣犋犑狊狙珇琚疽眗瞿矩砠秬窭窶筥簴粔粷繘罝耟聚聥腒舉艍苣苴莒菊蒟蒩蓻蘜虡蚷蜛螶袓裾襷詎諊讵豦貗趄趜跔跙距跼踘踞踽蹫躆躹車輂车遽邭郹醵鉅鋦鋸鐻钜锔锯閰陱雎鞠鞫颶飓駏駒駶驧驹鮈鮔鴡鵙鵴鶋鶪鼰鼳齟龃"},
	"juan":   {"倦劵勌勬卷呟圈圏埍奆奍姢娟婘嶲巂巻帣弮悁惓慻捐捲朘桊梋棬椦涓淃焆狷獧瓹眷睊睠絭絹縳绢罥羂脧腃臇菤蔨蠲裐讂鄄鋑鋗錈鎸鐫锩镌闂隽雋鞙飬餋鹃龹"},
	"jue":    {"亅倔傕决刔劂匷厥叕吷噊噘噱嚼埆壆妜孒孓屩屫崛崫嶡嶥弡彏憠憰戄抉挗捔掘撅撧攫斍桷橛橜櫭欮氒決泬灍焳熦燋爑爝爴爵獗玃玦玨珏瑴璚疦瘚矍矞矡砄穱絕絶繘绝臄芵蕝蕨虳蚗蟨蟩袦覐覚覺觉角觖觮觼訣誳譎诀谲貜赽趉趹蹶蹷躩鈌鐍鐝钁镢駃髉鳜鴂鴃鷢"},
	"jun":    {"俊儁军君呁均埈姰寯峻懏捃攈攟晙桾棞汮浚濬焌燇珺畯皲皸皹碅竣筠箘箟莙菌葰蚐蜠袀覠賐軍郡鈞銁銞鍕钧陖隽雋頵餕馂駿骏鮶鲪鵔鵕鵘麇麏麕龜龟"},
	"ka":     {"佧卡咔咖咯喀垰胩裃鉲"},
	"kai":    {"凯凱剀剴勓嘅垲塏奒岂嵦开忾恺愒愷愾慨揩暟楷欬欯濭炌烗蒈豈輆鍇鎎鎧鐦铠锎锴開闓闿颽"},
	"kan":    {"侃偘冚凵刊勘坎埳堪墈崁嵁嵌惂戡扻栞槛檻欿歁看瞰矙砍磡竷莰衎輡轗闞阚顑鬫龕龛"},
	"kang":   {"亢伉匟囥嫝嵻康忼慷扛抗槺漮炕犺砊穅粇糠躿邟鈧鏮钪閌闶鱇"},
	"kao":    {"丂尻嵪拷攷栲洘烤犒考薧銬铐靠髛鮳鯌鲓"},
	"ke":     {"克刻剋勀勊匼可呵咳嗑坷堁壳娔客尅岢峇嵑嵙嶱恪悈愘愙揢搕敤柯棵榼樖歁殻殼氪渇渴溘濭炣牁牱犐珂疴痾瞌砢硞碦磆磕礊礚科稞窠窼緙缂翗艐苛萪薖蚵蝌衉袔課课趷軻轲醘鈳鉿錁钶铪锞閜頦顆颏颗騍骒髁"},
	"kei":    {"剋尅"},
	"ken":    {"啃垦墾恳懇掯珢硍肎肯肻裉褃豤貇錹龂龈"},
	"keng":   {"劥吭坈坑妔挳揁摼殸牼硁硜硻誙銵鍞鏗铿阬"},
	"kong":   {"倥埪孔崆恐悾控椌涳矼硿空箜羫躻錓鞚鵼"},
	"kou":    {"佝冦剾劶口叩宼寇彄怐扣抠摳敂滱眍眗瞉瞘窛竘筘簆芤蔲蔻釦鷇"},
	"ku":     {"俈刳哭喾嚳圐堀库庫廤扝挎捁朏枯桍楛焅狜瘔矻硞秙窋窟絝绔胐苦袴裤褲跍郀酷骷鮬龲"},
	"kua":    {"侉咵垮夸姱恗挎晇楇絓胯舿誇趶跨銙骻"},
	"kuai":   {"会侩儈凷哙噲圦块塊墤巜廥快擓旝會欳浍澮狯獪璯筷糩脍膾蒉蒯郐鄶駃鱠鲙"},
	"kuan":   {"宽寛寬梡欵款歀窽窾臗鑧髋髖"},
	"kuang":  {"丱儣况劻匡匩卝哐圹壙夼岲忹恇懬懭抂旷昿曠框況洭爌狂狅眖眶矌矿砿硄礦穬筐筺絋絖纊纩誆誑诓诳貺贶軖軠軦軭迋邝邼鄺鉱鋛鑛鵟黋"},
	"kui":    {"亏傀刲匮匱喟喹嘳夔奎媿嬇尯岿巋巙悝愦愧憒戣揆晆暌楏楑櫆殨溃潰煃犪盔睽瞆瞶磈窥窺篑簣籄聧聩聭聵胿膭葵蒉蕢藈蘷虁虧蝰謉跬蹞躨逵鄈鍨鍷闚隗頍頯顝餽饋馈馗騤骙魁"},
	"kun":    {"困坤堃堒壸壼婫崐崑悃捆昆晜梱涃潉焜熴猑琨瑻睏硱祵稇稛綑菎蜫裈裍裩褌貇醌錕锟閫閸阃騉髠髡髨鯤鲲鵾鶤鹍"},
	"kuo":    {"噋廓扩拡括挄擴漷濶秮秳筈萿葀蛞銛銽铦闊阔霩鞟鞹韕頢髺鬠"},
	"la":     {"剌啦喇嚹垃拉揦揧搚擸攋旯柆楋櫴溂爉瓎瘌砬磖翋腊臈臘菈落藞蜡蝋蝲蠟辢辣邋鑞镴鞡鬎鯻"},
	"lai":    {"來俫倈唻娕婡崃崍庲徕徠攋来梾棶櫴涞淶濑瀨瀬猍琜疠癘癞癩睐睞筙箂籁籟莱萊藾襰誺賚賴赉赖逨郲錸铼頼顂騋鯠鵣鶆麳"},
	"lan":    {"儖兰厱啉嚂囒坔壈壏婪嬾孄孏岚嵐幱惏懒懢懶拦揽擥攔攬斓斕暕栏榄欄欖欗浨滥漤澜濫瀾灆灠灡烂燗燣燷爁爛爤爦璼瓓篮籃籣糷繿纜缆罱葻蓝藍蘫蘭褴襕襤襴襽覧覽览譋讕谰躝醂鑭钄镧闌阑韊顲"},
	"lang":   {"勆哴啷埌塱嫏崀廊斏朖朗朤桹榔樃欴浪烺狼琅瑯硠稂筤艆莨蒗蓈蓢蜋螂誏躴郎郒郞鋃鎯锒閬阆駺"},
	"lao":    {"乐佬僗劳労勞咾哰唠嗠嘮姥嫪崂嶗恅憥憦捞撈朥栳樂橯浶涝潦澇烙牢狫珯痨癆硓磱窂簩粩絡络老耂耢耮荖落蛯蟧躼軂轑酪醪銠鐒铑铹顟髝鮱"},
	"le":     {"乐了仂勒叻嘞忇扐楽樂氻泐牞玏砳竻簕肋艻阞韷餎饹鰳鱳鳓"},
	"lei":    {"傫儡儽勒厽嘞垒塁壘壨嫘擂攂樏檑櫐櫑欙泪洡涙淚灅瓃畾瘣癗矋磊磥礌礧礨祱禷类累絫縲纇纍纝缧罍羸耒肋蔂蕌蕾藟蘱蘲蘽虆蠝誄讄诔轠酹銇錑鐳鑘鑸镭雷靁頛頪類颣鱩鸓鼺"},
	"leng":   {"倰冷唥堎塄崚愣棱楞睖碐稜薐踜"},
	"li":     {"丽例俐俚俪傈儮儷凓刕利剓剺劙力励勵历厉厘厤厯厲叓叕吏呖哩唎唳喱嚟嚦囄囇坜塛壢娌娳婯嫠孋孷屴岦峛峢峲巁廲悝悡悧悷慄戾扐扚搮擽攊攦攭斄暦曆曞朸李杝枥柂栃栎栗栛梨梸棃棙樆檪櫔櫟櫪欐欚歴歷沥沴沵浬浰涖溧漓澧濿瀝灕爄爏犁犂犛犡狸猁珕理琍瑮璃瓅瓈瓑瓥疠疬痢癘癧皪盠盭睙睝砅砬砺砾磿礪礫礰礼禮禲离秝穲立竰笠筣篥篱籬粒粝粴糎糲綟縭纅纚缡罹脷艃苈苙茘荔荲莅莉菞蒚蒞蓠蔾藜藶蘺蚸蛎蛠蜊蜧蝷蟍蟸蠇蠡蠣蠫裏裡褵觻詈謧讈豊貍赲跞躒轢轣轹迣逦邌邐郦酈醨醴里釐鉝銐鋫鋰錅鎘鏫鑗锂镉隶隷隸離雳靂靋驪骊鬁鬲鯉鯏鯬鱧鱱鱺鲡鲤鳢鳨鴗鵹鷅鸝鹂麗麜黎黧"},
	"lia":    {"俩倆"},
	"lian":   {"亷僆劆匲匳嗹噒堜奁奩媡嫾嬚帘廉怜恋慩憐戀摙敛斂梿楝槏槤檶櫣歛殓殮涟湅溓漣潋澰濂濓瀮瀲炼煉熑燫琏瑓璉磏稴簾籢籨練縺纞练羷翴联聨聫聮聯脸臁臉莲莶萰蓮蔹薕薟蘝蘞螊蠊裢裣褳襝覝謰蹥连連鄻醶錬鍊鎌鏈鐮链镰鬑鰊鰱鲢"},
	"liang":  {"両两亮俍俩倆倞兩凉哴唡啢喨墚悢掚晾梁椋樑涼湸粮粱糧緉脼良莨蜽裲諒谅踉輌輛輬辆辌量鍄靓駺魉魎"},
	"liao":   {"了僚嘹嫽寥寮尞尥尦屪嵺嶚嶛廖廫憀憭摎撂撩敹料暸曢樛橑漻潦炓燎爎爒獠璙疗療瞭窌窷竂簝繆繚缪缭聊膋膫蓼藔蟟蟧豂賿蹘蹽辽遼鄝釕鐐钌镣镽飂飉髎鷯鹩"},
	"lie":    {"儠冽列劣劦劽咧哷埒埓姴巁巤忚挒挘捩擸栵棙毟洌浖烈烮煭燤爄爉犣猎猟獦獵綟聗脟茢蛚裂趔躐迾颲鬛鬣鮤鱲鴷"},
	"lin":    {"临亃僯冧凛凜厸吝啉壣崊嶙廩廪恡悋惏懍懔拎撛斴晽暽林橉檁檩淋潾澟瀶焛燐獜琳璘甐疄痳癛癝瞵碄磷箖粦粼綝繗翷膦臨菻蔺藺賃赁蹸躏躙躪轔轥辚遴邻鄰鏻閵阾隣霖驎鱗鳞麐麟"},
	"ling":   {"令伶凌刢另呤囹坽夌姈婈孁岭岺崚嶺彾掕昤朎柃棂棱櫺欞泠淩澪灵炩燯爧狑玲琌瓴皊砱祾秢稜竛笭紷綾绫羐羚翎聆舲苓菱蓤蔆蕶蘦蛉衑袊裬詅跉軨輘酃醽鈴錂铃閝陵零霊霗霛霝靈領领駖魿鯪鲮鴒鸰鹷麢齡齢龄龗"},
	"liu":    {"六刘劉嚠塯媹嬼嵧廇懰摎斿旈旒柳栁桞桺榴橊橮流浏溜漻澑瀏熘熮珋琉瑠瑬璢畄留畱疁瘤癅硫碌磂磟窌綹绺罶羀翏蒥蓅藰蟉裗蹓遛鋶鎏鎦鏐鐂锍镏镠雡霤飀飂飅飗餾馏駠駵騮驑骝鬸鰡鶹鷚鹠鹨麍"},
	"long":   {"儱咙哢嚨垄垅壟壠尨屸嶐巃巄弄徿拢攏昽曨朧栊梇槞櫳泷湰滝漋瀧爖珑瓏癃眬矓砻硦礱礲窿竉竜笼篢篭籠聋聾胧茏蘢蠪蠬襱豅贚躘鏧鑨陇隆隴霳靇驡鸗龍龒龓龙"},
	"lou":    {"偻僂剅喽嘍塿娄婁寠屚嵝嶁廔慺搂摟楼樓溇漊漏熡甊瘘瘺瘻瞜篓簍耧耬艛蒌蔞蝼螻謱貗軁遱鏤镂陋露鞻髅髏鷜"},
	"lu":     {"侓僇六剹勎勠卢卤噜嚕嚧圥坴垆塶塷壚娽峍庐廘廬彔录戮掳摝撸擄擼攎曥枦栌椂樐樚橹櫓櫨氇氌泸淕淥渌滷漉潞澛瀂瀘炉熝爐獹玈琭璐璷瓐甪盝盧睩矑硉硵碌磠祿禄稑穋穞箓簏簬簵簶籙籚粶緑纑绿罏胪膔臚舮舻艣艪艫芦菉蓼蓾蔍蕗蘆虂虏虜螰蠦觻謢賂赂趢路踛蹗輅轆轤轳辂辘逯醁鈩錄録錴鏀鏕鏴鐪鑥鑪镥陆陸露顱颅馿騄騼髗魯魲鯥鱸鲁鲈鴼鵦鵱鷺鸕鸬鹭鹵鹿麓黸"},
	"luan":   {"乱乿亂卵圝圞奱娈孌孪孿峦巒挛攣曫栾欒滦灓灤癴癵羉脔臠臡薍虊覶釠銮鑾鵉鸞鸾龻"},
	"lun":    {"仑伦侖倫囵圇埨婨崘崙惀抡掄棆沦淪溣碖磮稐綸纶耣腀芲菕蜦論论踚輪轮錀陯鯩"},
	"luo":    {"倮儸剆咯啰囉囖寽峈捋捰摞擽攊攎攞攭曪椤欏泺洛洜渃漯濼烙犖猡玀珞瘰癳皪砢硌硦磱笿箩籮絡纙络罗羅脶腡臝荦萝落蓏蘿蛒螺蠃裸覶覼詻跞躒躶逻邏鏍鑼锣镙雒頱饠駱騾驘骆骡鮥鵅鸁"},
	"lv":     {"侣侶偻僂儢勴卛吕呂哷垏屡屢履嵂律慮慺挔捋捛旅梠榈櫖櫚氀氯滤濾焒爈率瞜祣稆穭箻絽綠緑縷繂绿缕膂膐膟膢葎藘虑褛褸郘鋁鑢铝閭闾驢驴鷜"},
	"lve":    {"圙寽掠略畧稤鋝鋢锊"},
	"ma":     {"么亇傌吗唛嗎嘛嘜妈媽嫲孖庅抹摩杩榪溤犘犸獁玛瑪痲痳睰码碼祃禡罵菻蔴蚂螞蟆蟇遤鎷閁馬駡马骂鬕鰢鷌麻麽"},
	"mai":    {"买佅劢勱卖嘪埋売脈脉荬蕒薶衇買賣迈邁霡霢霾鷶麥麦"},
	"man":    {"僈埋墁姏娨嫚屘幔悗慢慲摱曼槾樠満满滿漫澫澷熳獌瞒瞞矕絻縵缦蔄蔓蘰蛮螨蟎蠻襔謾谩鄤鏋鏝镘鞔顢颟饅馒鬗鬘鰻鳗"},
	"mang":   {"厖吂哤壾娏尨庬忙恾朚杗杧氓汒浝漭牤牻狵痝盲盳硥硭笀芒茫茻莽莾蘉蛖蟒蠎邙釯鋩铓駹"},
	"mao":    {"乮冃冇冐冒卯堥夘媢峁嵍帽愗懋戼旄昴暓枆柕楙毛毷氂泖渵牦犛猫瑁皃眊瞀矛秏笷緢罞耄芼茂茅茆萺蓩蝐蝥蟊袤覒貌貓貿贸軞鄚鄮酕鉚鉾錨铆锚髦髳鶜"},
	"me":     {"么嚒嚜嚰庅濹麼麽"},
	"mei":    {"凂嚜坆堳塺妹媄媒媚媺嬍寐嵄嵋徾抺挴攗旀昧枚栂梅楣楳槑櫗殙毎每沒没沬浼渼湄湈煝煤燘猸玫珻瑂痗眉眛睂睸祙禖穈篃糜美羙脄脢腜苺莓葿蝞袂跊躾郿酶鋂鎂鎇镁镅霉韎鬽魅鶥鹛黣黴"},
	"men":    {"亹们們怋悗悶惛懑懣扪捫斖暪椚汶焖燜玧璊瞒瞞穈菛虋鍆钔門閅门闷鞔"},
	"meng":   {"儚冡勐夢夣孟尨幪庬懜懞懵掹擝曚朦梦橗檬氋氓溕濛猛獴甍甿癦盟瞢矇矒礞艋艨莔萌蒙蕄虻蜢蝱蟒蠓鄳鄸鋂錳锰雺霥霿靀顭饛髳鯍鯭鸏鹲黽黾鼆"},
	"mi":     {"侎冖冞冪咪嘧塓孊宓宻密峚幂幎幦弥弭彌怽戂摵擟攠敉枈榓樒檷櫁汨沕沵泌洣淧渳滵漞濔濗瀰灖熐爢猕獼瓕眯瞇瞴祕祢禰秘簚米粎糜糸縻羃羋脒芈葞蒾蔝蔤藌蘪蘼蜜袮覓覔覛觅詸謎謐谜谧迷醚醾醿釄銤镾靡鸍麊麋麛麿鼏"},
	"mian":   {"丏偭免冕勉勔喕娩婂媔嬵宀愐杣棉檰櫋汅沔渑湎澠牑眄眠睌矈矊矏糆絻綿緜緬绵缅腼臱芇莬葂蝒面靣靦鮸麪麫麵麺黽黾"},
	"miao":   {"仯劰厸喵妙媌嫹庙庿廟描杪淼渺玅眇瞄秒竗篎緢緲繆缈缪苗藐邈鶓鹋"},
	"mie":    {"乜吀咩哶孭幭懱搣櫗滅瀎灭烕眜礣篾簚蔑薎蠛衊覕鑖鱴鴓"},
	"min":    {"僶冺刡勄呡垊姄岷崏忞忟怋悯惽愍慜憫抿捪敃敏敯旻旼暋民泯渂湏湣潣玟珉琘琝瑉痻皿盿砇碈笢笽簢緍緡缗罠苠蠠賯鈱錉鍲閔閩闵闽鰵鳘鴖黽黾"},
	"ming":   {"佲冥凕名命姳嫇慏掵明暝朙榠洺溟猽眀眳瞑茗蓂螟覭詺鄍酩銘铭鳴鸣"},
	"miu":    {"繆缪謬谬"},
	"mo":     {"万么冒劘嗼嘿嚜嚤圽塻墨妺嫫嫼嬤嬷寞尛帓帞庅懡戂抹摩摸摹擵攠昩末枺模橅歾歿殁沒没沫漠瀎爅狢獏瘼皌眜眽眿瞐瞙砞磨礳秣糢絈絔縸纆耱脉膜茉莈莫蓦藦蘑蛨蟆蟔袹謨謩谟貃貈貉貊貘銆鏌镆陌靺饃饝馍驀骳髍魔魩魹麼麽默黙"},
	"mou":    {"侔劺厶呣哞堥恈敄某桙牟眸繆缪蛑謀谋踎鍪鞪鴾麰"},
	"mu":     {"亩仫凩募坶墓姆姥娒峔幕幙慔慕拇暮暯木朷楘模樢母毣毪氁沐炑牟牡牧牳狇畂畆畒畝畞畮目睦砪穆縸繆缪胟艒苜莫莯蚞踇鉧鉬钼雮霂鞪"},
	"n":      {"唔嗯"},
	"na":     {"乸内南吶呐哪嗱妠娜抐拏拿挐捺搻秅笚笝納纳肭蒘蒳衲袦誽豽貀軜那郍鈉鎿钠镎雫靹魶"},
	"nai":    {"乃佴倷奈奶妳嬭孻廼搱摨柰氖渿疓耏耐能腉艿萘螚褦迺釢錼鼐"},
	"nan":    {"侽南喃囝囡妠娚婻嫨戁抩揇暔枏柟楠湳煵男畘腩莮萳蝻諵赧遖难難"},
	"nang":   {"乪儾嚢囊囔憹搑擃攮曩欜涳瀼灢蠰饟饢馕齉"},
	"nao":    {"匘呶垴堖夒婥嫐孬峱嶩巎怓恼悩惱憹挠摎撓桡橈淖猱獶獿瑙硇碙碯脑脳腦臑蛲蟯詉譊鐃铙閙闹鬧"},
	"ne":     {"吶呐呢哪抐疒眲訥讷"},
	"nei":    {"內内哪娞氝氞浽脮腇那錗餒馁鮾鯘"},
	"nen":    {"媆嫩嫰恁"},
	"neng":   {"能螚"},
	"ni":     {"伱伲你倪儗儞匿呢坭埿堄妮妳婗嫟嬺孴尼屔屰怩惄愵懝抐抳拟掜擬旎昵晲暱柅棿檷氼泥淣溺濔狔猊疑眤睨秜籾縌聣聻胒腻膩臡苨薿蚭蛪蜺觬誽貎跜輗逆郳鈮铌隬霓馜鯢鲵麑齯"},
	"nian":   {"卄哖唸埝姩年廿念悥惗拈捻撚撵攆榐涊淰溓焾碾秊秥簐粘艌蔫蹍蹨躎輦辇辗鮎鯰鲇鲶鵇黏"},
	"niang":  {"娘嬢孃酿醸釀"},
	"niao":   {"嫋嬝嬲尿樢溺脲茑蔦袅裊褭鳥鸟"},
	"nie":    {"乜啮嗫噛嚙囁囐囓圼孼孽峊嵲嶭巕帇捏捻掜揑摄摰攝敜枿槷櫱涅湼痆篞籋糱糵聂聶臬臲苶菍蘖蠥讘踂踗踙蹑躡銸鋷錜鎳鑈鑷钀镊镍闑陧隉顳颞齧"},
	"nin":    {"囜恁您拰脌"},
	"ning":   {"佞侫倿儜凝咛嚀嬣宁寍寕寗寜寧拧擰柠橣檸泞澝濘狞獰甯矃聍聹苧薴鑏鬡鸋"},
	"niu":    {"妞孧忸怓扭抝拗杻汼沑炄牛牜狃紐纽莥鈕钮靵"},
	"nong":   {"侬儂农哝噥弄挊挵檂欁浓濃燶癑禯秾穠繷脓膿蕽襛譨農辳醲鬞齈"},
	"nou":    {"搙槈檽獳羺耨鎒鐞"},
	"nu":     {"伮傉努奴孥帑弩怒搙砮笯胬駑驽"},
	"nuan":   {"奻暖渜湪煖煗餪"},
	"nun":    {"黁"},
	"nuo":    {"傩儺喏娜愞懦懧挪掿搙搦搻梛榒橠毭稬穤糑糥糯耎袲袳諾诺蹃逽鍩锘难難"},
	"nv":     {"女恧朒沑籹聏衂衄釹钕"},
	"nve":    {"婩疟瘧硸虐"},
	"o":      {"哦喔噢嚄筽"},
	"ou":     {"偶区區吘吽呕嘔塸怄慪櫙欧歐殴毆沤漚熰瓯甌禺耦腢膒蕅藕藲謳讴鏂鴎鷗鸥"},
	"pa":     {"啪妑帊帕怕扒掱杷汃潖爬琶皅筢耙舥苩葩袙趴钯"},
	"pai":    {"俳哌廹徘拍排棑沠派渒湃牌犤猅箄簰簲蒎輫迫鎃"},
	"pan":    {"丬冸判叛坢媻幋拌拚搫攀柈槃沜泮洀湴溿潘瀊炍爿牉畔畨番盘盤盻盼眅眫磐磻籓縏聁肨胖膰萠蒰蟠袢襻詊跘蹒蹣鋬鎜鑻闆鞶頖鵥"},
	"pang":   {"乓厐嗙夆嫎尨庞庬彷徬旁汸沗滂炐眫磅篣耪肨胖胮膖舽螃覫趽逄雱霶鰟鳑龎龐"},
	"pao":    {"刨匏咆嚗垉奅庖抛拋泡炮炰烰爮狍瓟疱皰砲礟礮穮脬萢藨袌袍褜跁跑軳鉋铇靤鞄麃麅麭"},
	"pei":    {"伂佩俖呸培妃妚姵婄媐嶏帔怌抷攈斾旆柭柸棑毰沛浿淠犻珮琣肧胚蓜衃裴裵賠赔轡辔配醅錇锫阫陪霈馷"},
	"pen":    {"呠喯喷噴歕湓濆瓫盆翸葐"},
	"peng":   {"亨倗傰剻匉嘭堋塜塳弸彭怦恲憉抨挷捀捧掽搒摓朋梈棚椖椪槰樥泙淎淜漨漰澎烹熢痭皏砰硑硼碰磞稝竼篣篷絣纄胓膨芃苹莑蓬蘕蟚蟛袶踫軯輣錋鑝閛韸韼駍騯髼鬅鬔鵬鹏"},
	"pi":     {"丕仳伓伾僻副劈匹吡否啤噼噽嚊嚭圮坯埤壀妚媲嫓屁岯崥嶏庀怶悂憵批披抷揊擗旇朇枇枈榌毗毘淠潎澼濞炋焷狉狓猈琵甓疈疋疲痞癖皮睥砒磇礔礕禆秛秠稫笓篺紕纰罢罴罷羆翍耚肶脴脾腗膍芘苉蚍蚽蚾蜱螕螷蠯被裨諀譬豼豾貔辟邳郫釽鈚鈹鉟銔銢錃錍鎞铍闢阰陂陴隦霹鞞駓髬魮魾鮍鲏鴄鵧鷿鸊鼙"},
	"pian":   {"便偏囨媥扁楄楩片犏猵獱瑸璸篇緶缏翩胼腁覑諞谝貵賆跰蹁鍂駢騈騗騙骈骗骿魸鶣"},
	"piao":   {"僄剽勡嘌嫖彯徱慓旚朴殍漂犥瓢皫瞟票篻縹缥翲莩薸螵醥闝顠飃飄飘驃骠魒鰾"},
	"pie":    {"丿嫳撆撇暼氕潎瞥苤覕鐅"},
	"pin":    {"品嚬姘娦嫔嬪拚拼榀汖牝玭琕矉砏礗穦聘薲蘋蠙貧贫顰频颦馪驞"},
	"ping":   {"乒俜冯凭凴呯坪塀娉屏屛岼帡帲幈平庰慿憑枰檘泙洴涄淜焩玶瓶甁甹砯硑竮箳簈缾聠胓艵苹荓萍蓱蘋蚲蛢評评軿輧郱鉼頩馮鮃鲆"},
	"po":     {"叵哱嘙嚩坡奤婆尀岥岶廹搫敀昢朴櫇泊泺泼洦湐溌潑濼烞猼珀皤破砶笸箥粕繁翍蒪蔢迫鄱酦醗醱釙鉕鏺钋钷陂頗颇駊髲魄"},
	"pou":    {"剖咅哣垺娝婄抔抙捊掊棓犃箁裒"},
	"pu":     {"仆僕匍噗圃圑圤埔堡墣扑抪捗撲擈攴攵普暜暴曝朴柨樸檏氆浦溥潽濮瀑炇烳獛璞痡瞨秿穙箁纀舖舗莆菐菩葡蒱蒲諩譜谱贌蹼酺鋪鏷鐠铺镤镨陠鯆"},
	"qi":     {"七丌乞亓亝亟企伎俟倛傶僛其凄剘启呇呚呮咠唘唭啐啓啔啟喰嘁噐器圻埼墄夡奇契妻娸婍屺岂岐岓崎嵜帺弃忔忯恓悽愒愭慼慽憇憩懠戚扱捿掑揭摖攲敧斉斊旂旗晵暣期杞枝柒栔栖桤桼梩棄棊棋棨棲榿槭檱櫀欫欹欺歧气気氣汔汽沏泣洓淇淒湆湇滊漆濝炁焏猉玂玘琦琪璂甈畁畦疧盀盵矵砌碁碕碛碶磜磧磩礘祁祇祈祺禥稘稽竐竒簯簱籏粸紪綥綦綨綮綺緀緝绮缉罊耆肵脐臍艩芑芞芪荠萁萋萕葺蕲薺藄藒蘄蚑蚔蚚蛴蜝蜞螇螧蟿蠐袳褀褄觭訖諆諬讫豈起趞跂踦軝迄迉邔郪鄿釮錡鏚鐑锜闙陭隑霋頎颀騎騏騹骐骑鬐鬿魌鯕鰭鲯鳍鵸鶀鶈麒麡齊齐"},
	"qia":    {"冾匼卡咭圶峠帢恰愘抲拤掐揢擖殎洽硈磍胢葜袷跒酠鞐髂"},
	"qian":   {"乹乾亁仟仱伣佥俔倩偂傔僉儙凵刋前千厱唊嗛圱圲堑塹墘壍奷婜媊嬱孅孯岍岒嵌嵰忴悓悭愆慊慳扦扲拑拪挳掅掔掮揵搴摼撁撍攐攑攓朁杄棈椠榩槏槧橬檶櫏欠欦欿歉歬汘汧浅淺潛潜濳灊炶煔牵牽犍玪皘磏竏签箝篏篟簽籖籤粁綪縴繾纤缱羬肷膁臤芊芡茜茾荨葥蒨蔳蕁虔蚈蚙蜸褰諐謙譴谦谴谸軡輤迁遣遷釺鈆鈐鉗鉛銭錢鎆鏲鑓钎钤钱钳铅阡韆顅騚騝騫骞鬜鬝鰬鳽鵮鹐黔黚"},
	"qiang":  {"呛唴啌嗆嗴墏墙墻嫱嬙嶈廧強强彊戕戗戧抢搶摤摪斨枪椌槍樯檣溬漒炝熗牄牆猐獇玱瑲疆矼篬繈繦羌羗羟羥羫羻腔艢蔃蔷薔蘠蜣襁謒跄蹌蹡鎗鏘鏹锖锵镪鶬"},
	"qiao":   {"丂乔侨俏偢僑僺劁勪喬喿嘺噭墝墧墽壳嫶峤峭嵪嶠巧帩幓幧悄愀憔摮撬撽敲桥槗樵橇橋橾櫵殻殼毃潐煍燆燋犞癄睄瞧硗硚碻磝磽礄礉窍竅箾繑繰缲翘翹荍荞菬蕎藮誚譙诮谯趫趬跷踍蹺蹻躈郻鄡鄥釥鍫鍬鏒鐈鐰锹陗雀鞒鞘鞩鞽韒頝顦骹髚髜"},
	"qie":    {"且伽切匧厒嗛契妾帹怯悏惬愜慊挈朅栔椄洯淁猰疌癿砌稧穕窃竊笡箧篋籡聺苆茄蛪趄跙踥郄鍥锲鯜"},
	"qin":    {"亲侵兓勤吢吣唚嗪噙坅埐媇嫀寑寖寝寢寴嵚嵰嶔嶜庈廑忴慬懃懄扲抋捦揿搇撳擒斳昑梫檎欽沁溱澿瀙珡琴琹瘽矜矝禽秦笉綅耹肣芩芹菣菦菳藽蚙螓螼蠄衾覃親誛赾鈙鈫鋟钦锓雂靲顉駸骎鮼鳹鵭"},
	"qing":   {"亲倾傾儬凊剠勍卿啨圊埥夝寈庆庼廎情慶掅擎擏晴暒棾樈檠檾櫦殑殸氢氫氰淸清渹漀濪甠硘碃磬箐綪綮罄苘葝蜻親請謦请軽輕轻郬錆鑋靑青靘頃顷鯖鲭黥"},
	"qiong":  {"儝卭嬛宆惸憌桏橩焪焭煢熍琁琼璚瓊瓗睘瞏穷穹窮竆笻筇舼茕藑藭蛩蛬赹跫邛銎"},
	"qiu":    {"丘丠仇俅厹叴唒囚坵塸媝崷巯巰恘恷扏搝朹梂楸橚殏毬求汓泅浗渞湫湬湭煪牫犰玌球璆皳盚秋秌穐篍糗紌絿緧肍艽芁莍萩蓲蘒虬虯蚯蛷蝤蝵蟗蠤裘觓觩訄訅賕赇趥逎逑遒邱酋醔釚釻銶鞦鞧頄鮂鯄鰌鰍鰽鳅鶖鹙鼽龜龝龟"},
	"qu":     {"伹佉佢刞劬匤区區厺去取呿坥娶屈岖岴嶇忂憈戵抾敺斪曲朐欋欔欪氍浀淭渠灈焌煀爠璖璩癯瞿磲祛竘竬筁籧粬紶絇翑翵耝胊胠臞苣菃葋蕖蘧蚼蛆蛐蝺螶蟝蠷蠼衐衢袪覰覷覻觑詓詘誳诎趋趍趣趨躣躯軀軥鑺閴闃阒阹駆駈驅驱髷魼鰸鱋鴝鶌鸜鸲麮麯麴麹黢鼁鼩齲龋"},
	"quan":   {"佺全券劝勧勸啳圈圏埢奍姾婘孉峑巏弮恮悛惓拳捲搼权棬椦楾槫権權泉洤湶灥烇牶牷犈犬犭犮狋瑔畎痊矔硂筌箞絟綣縓绻荃葲虇蜷蠸觠詮诠跧踡輇辁醛銓鐉铨闎韏顴颧駩騡鬈鰁鳈齤"},
	"que":    {"傕却卻埆塙墧崅悫愨慤搉敠敪棤榷汋炔燩琷瘸皵硞确碏確碻礐礭缺缼蒛蚗趞闋闕阕阙雀鵲鹊"},
	"qun":    {"囷夋宭峮帬歏箘羣群裙裠踆逡麇"},
	"ran":    {"冄冉呥嘫姌媣染橪然熯燃珃繎肰苒蒅蚦蚺衻袇袡髥髯"},
	"rang":   {"儴勷嚷壌壤懹攘瀼爙獽瓤禳穣穰纕蘘譲讓让躟鬤"},
	"rao":    {"娆嬈扰擾桡橈穘繞绕荛蕘襓遶隢饒饶"},
	"re":     {"喏惹捼渃热熱若"},
	"ren":    {"人亻仁仞仭任刃刄壬妊姙屻岃忈忍忎扨朲杒栠栣梕棯牣祍秂秹稔紉紝絍綛纫纴肕腍芢荏荵菍衽袵訒認认讱躵軔轫釰釼鈓銋靭靱韌韧飪餁饪魜鵀"},
	"reng":   {"仍扔礽芿辸陾"},
	"ri":     {"囸日氜衵釰釼鈤馹驲"},
	"rong":   {"傇傛冗坈媶嫆嬫宂容峵嵘嵤嶸戎搈搑曧栄榕榮榵槦毧氄溶瀜烿熔爃狨瑢穁穃絨縙绒羢肜茙茸荣蓉蝾融螎蠑褣軵鎔镕駥"},
	"rou":    {"厹媃宍揉柔楺渘煣瑈瓇禸粈糅肉脜腬葇蝚譳蹂輮鍒鞣韖騥鰇鶔"},
	"ru":     {"乳侞儒入吺嗕嚅女如媷嬬孺嶿帤扖挐擩曘杁桇檽汝洳渪溽濡獳筎縟繻缛肗茹蒘蓐蕠薷蝡蠕袽褥襦辱邚鄏醹銣铷顬颥鱬鳰鴑鴽"},
	"ruan":   {"偄堧壖媆撋朊檽瑌瓀碝礝緛耎腝軟輭软阮"},
	"rui":    {"兊兌兑叡壡惢抐枘桵橤汭瑞甤睿笍緌繠芮苼蕊蕋蕤蘂蘃蚋蜹銳鋭鏸锐"},
	"run":    {"橍润潤瞤膶閏閠闰"},
	"ruo":    {"偌叒婼嵶弱挼捼楉渃焫爇箬篛若蒻鄀鰙鰯鶸"},
	"sa":     {"仨卅摋撒攃櫒泧洒潵灑纚脎萨蕯薩訯躠鈒鏾钑隡靸鞈颯飒馺"},
	"sai":    {"僿嗮嘥噻塞思愢揌毢毸簺腮賽赛顋鰓鳃"},
	"san":    {"三仐伞俕傘厁叁参參叄叅壭帴弎散毵毶毿潵犙糁糝糣糤繖鏒鏾閐饊馓鬖"},
	"sang":   {"丧喪嗓搡桑桒槡磉褬鎟顙颡"},
	"sao":    {"埽嫂慅懆扫掃掻搔梢氉溞瘙矂繅繰缫缲臊螦颾騒騷骚髞鰠鱢鳋"},
	"se":     {"啬嗇塞廧懎拺摵擌栜槭歮歰洓涩渋澀澁濇濏瀒琗瑟璱瘷穑穡穯粣繬色薔譅轖銫鎍鏼铯閪闟雭飋"},
	"sen":    {"森椮槮襂"},
	"seng":   {"僧鬙"},
	"sha":    {"乷倽傻儍刹剎厦唦唼啑啥喢嗄帹廈挱挲摋杀杉桬榝樧歃歰殺毮沙濈煞猀痧砂硰箑粆紗繺纱翜翣莎菨萐蔱裟鎩铩閯閷霎髿魦鯊鯋鲨"},
	"shai":   {"摋攦晒曬筛篩簁簛籭諰"},
	"shan":   {"傓僐儃删刪剡剼单善單嘇圸埏墠墡姍姗嬗山嶦幓彡扇挻掞掸掺搧摻撣擅敾晱杉柵栅椫樿檆歚汕潬潸澘灗炶煔煽熌狦珊疝痁睒磰禅禪穇笘縿繕缮羴羶脠膳膻舢芟苫蟮蟺衫覢訕謆譱讪贍赡赸跚軕邖鄯釤銏鐥钐閃閄閊闪陕陝饍騸骟鯅鱓鱔鳝鳣"},
	"shang":  {"丄上伤傷商垧埫墒尙尚恦愓慯扄晌樉殇殤汤湯滳漡熵禓緔绱蔏螪裳觞觴謪賞贘赏鋿鏛鑜鞝鬺"},
	"shao":   {"佋劭勺卲召哨娋少弰捎旓杓柖梢潲烧焼燒玿睄稍筲紹綤绍艄芍苕莦萷蕱袑輎邵鞘韶髾鮹"},
	"she":    {"佘厍厙奓奢射弽慑慴懾折拾挕捨揲摂摄攝檨欇歙泏涉涻渉滠灄猞畬畭畲磼社舌舍舎蔎虵蛇蛥蠂設设賒賖赊赦輋闍阇韘騇麝"},
	"shei":   {"誰"},
	"shen":   {"什伸侁侺信兟参參叄叅吲呻哂堔妽姺娠婶嫀嬸审宷審屾峷幓弞愼慎扟抌搷敒昚曋曑柛棯棽椹榊氠沈涁淰深渖渗滲瀋燊珅甚甡甧申瘆瘮眒眘瞫矤矧砷神祳穼籶籸糁糂糝糣紳绅罙罧肾胂脤腎莘葚葠蓡蔘薓蜃蜄裑覾訠訷詵諗讅诜谂谉谌身邥鉮鋠頣駪魫鯓鯵鰰鰺鲹鵢黮"},
	"sheng":  {"乗乘偗剩剰勝升呏圣墭声娍嵊憴斘昇晟晠曻枡椉榺橳殅殸泩渑渻湦澠焺牲狌珄琞生甥甸盛省眚竔笙縄繉繩绳聖聲胜苼蕂譝貹賸鉎鍟阩陞陹鵿鼪"},
	"shi":    {"世丗乨乭亊事什仕似使侍兘冟势勢匙十卋厔史呞呩咶啫嗜嘘噓噬埘埶塒士失奭始姼嬕实実室宩宲寔實尸屍屎峕峙崻崼嵵市师師式弑弒徥忕忯恀恃惿戺拭拾揓施时旹是昰時枾柹柿栅栻楴榁榯檡殖氏泽浉湜湤湿溡溮溼澤澨濕炻烒煶狧狮狶獅瑡畤痑眂眎眡睗矢石硕碩示礻祏秲竍笶筮箷簭籂絁耆舐舓莳葹蒒蒔蓍虱蚀蝕蝨螫褷襫襹視视觢試詩誓諟諡謚識识试诗谥豕貰贳軾轼辻适逝遈適遾邿酾醳釃釈释釋釶鈰鉂鉃鉇鉐鉽銴鍦铈食飠飾餙餝饣饰馶駛驶鮖鯴鰘鰣鰤鲥鲺鳲鳾鶳鸤鼫鼭"},
	"shou":   {"兽収受售嘼垨壽夀守寿手扌授收敊涭狩獣獸痩瘦綬绶膄艏鏉首龵"},
	"shu":    {"书侸俞倏倐儵兪叔咰塾墅姝婌孰尌尗属屬庶庻忬怷怸恕悆戍抒捈捒掓摅攄数數暏暑曙書朮术朱束杸枢树梳樞樹橾殊殳毹毺沭淑漱潄潻澍濖瀭焂熟琡瑹璹疋疎疏癙秫稌竖竪籔糬紓絉綀緰纾署腧舒荗菽蒁蒣蔬薥薯藷虪蜀蠴術裋襡襩豎贖赎跾踈軗輸输述鄃鉥錰鏣陎隃霔鮛鱪鱰鵨鶐黍鼠鼡龧"},
	"shua":   {"刷唰耍誜"},
	"shuai":  {"卛帅帥摔率甩縗缞蟀衰"},
	"shuan":  {"拴栓槫涮絟腨閂闩"},
	"shuang": {"双塽孀孇慡欆泷漺瀧灀爽礵縔艭鏯雙霜騻驦骦鷞鸘鹴"},
	"shui":   {"娷帨挩捝水氵氺涗涚睡祱稅税脽裞說説誰说谁閖"},
	"shun":   {"吮楯橓眴瞚瞤瞬舜蕣順顺鬊"},
	"shuo":   {"妁愬搠数數朔槊欶洬烁爍獡矟硕碩箾蒴說説说銏鎙鑠铄"},
	"si":     {"丝亖伺似佀価俟俬偲傂儩兕凘厮厶司咝嗣嘶噝四姒娰媤孠寺巳廝思恖愢撕斯杫枱柶梩楒榹死汜泀泗泤洍洠涘澌瀃燍牭磃祀禗禠禩私竢笥糸糹絲緦纟缌罳耛耜肂肆蕬蕼虒蛳蜤螄蟖蟴覗貄釲鈶鈻鉰銉銯鋖鍶鐁锶颸飔食飤飴飼饲駟騃騦驷鯣鷥鸶鼶龱"},
	"song":   {"倯傱凇吅娀宋崧嵩嵷庺忪怂悚愡愯慫憽捒揔摗松枀枩柗梥棇楤檧淞漎濍硹竦耸聳菘訟誦讼诵送鎹頌颂餸駷鬆"},
	"sou":    {"傁凁叜叟嗖嗽嗾廀廋捒捜搜摉摗撨擞擻棷櫢欶溲獀瘶瞍籔艘蒐蓃薮藪螋鄋醙鎪鏉锼颼颾飕餿馊騪"},
	"su":     {"俗傃僳嗉囌圱埣塐塑夙嫊宿愫愬憟摵梀棴榡樎樕橚櫯殐泝洬涑溯溸潚潥玊珟璛甦碿稡稣穌窣簌粛粟素縤縮缩肃肅膆苏莤蔌藗蘇蘓觫訴謖诉谡趚蹜速遡遬酥鋉餗驌骕鯂鱐鷫鹔"},
	"suan":   {"匴狻痠祘笇筭算篹蒜酸"},
	"sui":    {"亗倠哸嗺埣夊娞嬘尿岁嵗旞檖歲歳毸浽滖澻濉瀡煫熣燧璲瓍眭睟睢砕碎祟禭穂穗穟綏縗繀繐繸绥缞膸芕荽荾虽襚誶譢谇賥遀遂邃鏸鐆鐩隋随隧隨雖鞖韢髄髓"},
	"sun":    {"喰孙孫扻损損搎摌榫槂狲猻笋筍箰簨荪蓀蕵薞鎨隼飧飱鶽"},
	"suo":    {"乺些傞唆唢嗍嗦嗩娑嫅惢所挱挲摍暛桫梭溑溹獕琐琑瑣璅睃簑簔索縒縮缩羧莎莏蓑蜶褨趖逤鎍鎖鎻鏁锁鮻"},
	"ta":     {"他侤傝嗒嚃嚺塌塔墖她它崉拓挞搨撻榙榻橽毾沓涾溚溻漯澾濌牠狧獭獺祂禢粏褟誻譶趿踏蹋蹹躂躢遝遢鉈錔鎉鎑铊闒闟闥闧闼阘鞈鞜鞳鮙鰨鳎"},
	"tai":    {"儓冭台囼坮大太夳奤嬯孡忕忲态態抬擡斄旲枱檯汏汰汱泰溙炱炲燤珆箈籉肽胎臺舦苔菭薹跆邰酞鈦钛颱駘骀鮐鲐"},
	"tan":    {"倓傝僋儃叹啴嗿嘆嘽嘾坍坛坦埮墰墵壇壜婒弹弾彈忐怹惔憛憳憻探摊擹攤昙暺曇榃橝檀歎毯湠滩潬潭澹灘炭燂璮痑痰瘫癱碳緂繵罈罎舑舕菼藫袒襢覃談譚譠谈谭貚貪贪郯醈醓醰鉭錟钽锬顃"},
	"tang":   {"伖倘偒傏傥儻劏唐啺嘡坣堂塘嵣帑愓戃搪摥曭棠榶樘橖汤淌湯溏漟灙烫煻燙爣瑭矘磄禟篖糃糖糛羰耥膅膛蓎薚蝪螗螳赯趟踼蹚躺鄌醣鎕鎲鏜鐋钂铴镋镗闛隚鞺餹饄鶶鼞"},
	"tao":    {"匋叨咷啕夲夵套嫍幍弢慆抭掏搯桃梼槄檮洮涛涭淘滔濤瑫祹絛綯縚縧绦绹萄蜪裪討詜謟讨跳轁迯逃醄鋾陶鞀鞉鞱韜韬飸饀饕駣騊鼗"},
	"te":     {"忑忒慝特犆脦蟘鋱铽"},
	"teng":   {"儯幐滕漛熥疼痋籐籘縢腾膯藤虅螣誊謄邆霯駦騰驣鰧鼟"},
	"ti":     {"体俶倜偍剃剔厗啑啼嗁嚏嚔奃媂媞屉屜屟崹弟徥徲悌悐惕惖惿戻挮掦提揥擿替朑桋梯棣歒殢涕渧漽珶瑅瓋眣睼碮磃禔禵稊笹籊綈緹绨缇罤苐荑蕛薙裼褅褆謕趧趯踢蹄蹏躰軆逖逷遆醍銻鍗锑題题騠骵體髰鬀鬄鮷鯷鳀鵜鶗鶙鷈鷉鷤鹈"},
	"tian":   {"佃倎兲呑唺塡填天奵婖屇忝恬悿掭搷晪栝殄沺淟添湉琠瑱璳甛甜田甸畋畑痶盷睓碵磌窴紾緂胋腆舔舚菾覥觍賟酟銛銽錪钿铦闐阗靔靝餂鴫鷆鷏黇"},
	"tiao":   {"佻啁嬥宨岧岹庣恌挑斢旫晀朓朷条條樤眺祒祧窕窱笤粜糶絩聎脁芀苕萔蓚蓧蓨蜩覜誂調调趒跳迢銚鋚鎥铫鞗髫鯈鰷鲦齠龆"},
	"tie":    {"僣呫帖怗聑萜蛈貼贴跕鉄銕鐡鐵铁飻餮驖"},
	"ting":   {"亭侱侹停厅厛听圢娗婷嵉庁庍庭廰廳廷忊挺桯梃楟榳汀涏渟濎烃烴烶珽町甼筳綎耓聤聴聼聽脡艇莛葶蜓蝏誔諪邒鋌铤閮霆鞓頲颋鼮"},
	"tong":   {"仝佟侗偅僮勭同哃嗵囲垌峂峒峝庝彤恸慟憅捅晍曈朣桐桶樋橦氃浵潼炵烔燑犝狪獞痌痛眮瞳砼硧秱穜童筒筩粡絧統綂统膧茼蓪蚒衕詷赨通酮鉖鉵銅铜餇鮦鲖"},
	"tou":    {"亠偷偸埱头妵婾媮愉投敨斢紏緰蘣透鋀鍮頭飳骰黈"},
	"tu":     {"兎兔凃凸吐唋啚図图圖圗土圡堍堗塗墿宊屠峹嵞嶀庩廜徒怢悇捈捸揬摕梌檡汢涂涋湥潳瑹痜瘏禿秃稌突筡腯荼莵菟葖蒤趃跿迌途酴釷鈯鋀鋵鍎钍馟駼鵌鵚鵵鶟鷋鷵鼵"},
	"tuan":   {"剸团団圕團塼墥彖慱抟揣摶槫檲湍湪漙煓猯畽疃篿糰褖貒鏄鷒鷻"},
	"tui":    {"侻俀僓墤娧尵弚弟忒推橔煺穨聉腿蓷藬蘈蛻蜕螁褪讉蹆蹪退隤頹頺頽颓駾骽魋"},
	"tun":    {"吞呑啍噋囤坉屯庉忳憞旽暾朜氽涒焞臀臋芚蛌褪豘豚軘霕飩饨魨鲀黗"},
	"tuo":    {"乇仛佗侂侻咃咜唾啴坨堶妥媠嫷岮嶞庹彵托扡拓拕拖挩捝撱杔杝柁柝椭楕槖橐橢毤毻汑沰沱沲涶狏砣砤碢箨籜紽脫脱莌萚蘀袉袥託詑讬跅跎酡鉈鋖铊阤陀陁飥饦馱馲駄駝駞騨驒驝驮驼魠鮀鮵鰖鴕鵎鸵鼉鼍鼧"},
	"wa":     {"佤凹劸呙咓咼哇唲啘嗗嗢娃娲婠媧屲帓徍挖搲攨洼溛漥瓦瓸瓾畖砙穵窊窐窪聉腽膃蛙袜襪邷韈韎韤鼃"},
	"wai":    {"呙咼喎外崴歪瀤竵顡"},
	"wan":    {"万丸倇刓剜卍卐唍埦塆壪夗夘妧婉婠完宛岏帵弯彎忨惋惌抏挽捖捥掔晚晥晩晼杤杬梚椀槾毌汍涴湾潫澫灣烷玩琓琬畹皖盌睕瞣碗笂箢紈絻綄綩綰纨绾翫脕脘腕芄莞莬菀萖萬蔓薍蚖蜿蟃豌貦贃贎踠輐輓鋄鋔鋺錽鎫頑顽"},
	"wang":   {"亡亾仼兦妄尣尩尪尫彺往徃徍忘忹惘抂旺暀朚望朢枉棢汪瀇王盳網网罒罓罔罖莣菵蚟蛧蝄誷輞辋迋迬魍"},
	"wei":    {"为亹伟伪位倭偉偎偽僞儰卫危厃叞味唩唯喂喡喴囗囲围圍圩墛壝委威娓媁媙媦寪尉尾屗峗峞崣崴嵔嵬嶉嶶巍帏帷幃廆徫微惟愄愇慰懀捤捼揋揻撝撱斖暐未桅梶椲椳楲欈沩洈洧浘涠渨渭湋溈溦潍潙潿濰濻瀢炜為烓煒煟煨熭燰爲犚犩猚猥猬玮琟瑋瓗畏痏痿癐癓瞆矀硊硙碨磈磑維緭緯縅纬维罻胃腲艉芛苇苿荱菋萎葦葨葳蒍蓶蔚蔿薇薳藯蘶蜲蜼蝛蝟螱衛衞褽覣覹詴諉謂讆讏诿谓踓躗躛軎轊违逶違鄬醀鍏鍡鏏闈闱隇隈隗隹霨霺韋韑韙韡韦韪頠颹餧餵饖骩骪骫魏鮇鮠鮪鰃鰄鲔鳂鳚"},
	"wen":    {"伆刎吻呅呡問塭妏彣忞忟抆揾搵文昷桽榅榲歾殟汶渂温溫炆玟珳瑥璺瘒瘟稳穏穩紊紋絻緼縕繧纹缊聞肳脗芠莬蚉蚊螡蟁豱輼轀辒鎾閺閿闅闦问闻阌雯鞰韫顐饂馼駇魰鰛鰮鳁鳼鴍鼤"},
	"weng":   {"勜嗡塕奣嵡攚暡滃瓮甕瞈罋翁聬蓊蕹螉鎓鶲鹟齆"},
	"wo":     {"仴倭偓卧呙咼唩喔婐婑媉幄我挝捰捾握撾斡枂楃沃涡涴涹渥渦濄濣焥猧瓁瞃硪窝窩緺肟腛臒臥莴萵薶蜗蝸踒馧齷龌龏"},
	"wu":     {"乄乌五亡仵伍侮俉倵僫儛兀剭务務勿午卼吳吴吾呉呜唔啎嗚圬坞堥塢墲奦妩娒娪娬婺嫵寤屋屼岉峿嵍嵨巫庑廡弙忢忤怃恶悞悟悮惡憮戊扜扝扤捂摀敄於无旿晤杅杇杌柮梧橆歍武毋汙汚污沕洖洿浯渞溩潕烏焐無熃熓物牾玝珷珸瑦璑甒痦瞴矹碔祦窏窹箼粅膴舞芜芴茣莁蕪蘁蜈螐蟱誈誣誤譕诬误趶躌迕逜邬郚鄔釫鋈鋘鋙鎢钨铻阢陚隖雺雾霚霧靰騖骛鯃鰞鴮鵐鵡鶩鷡鹀鹉鹜鼯鼿齀"},
	"xi":     {"习係俙傒僖兮凞匸卌卥厀吸呬咥咭唏唽喜喺嘻噏嚱囍墍壐夕奚娭媳嬆嬉屃屓屖屣屭嵠嶍嶲巂巇希席徆徙徯忚忥忾怬怸恄恓息悉悕惁惜愾慀憘憙戏戯戱戲扢扱扸摡昔晞晰晳暿曦杫析枲栖桸椞椺榽槢樨橀橲檄欪欯欷歖歙氥汐洒洗浠淅渓溪滊漇漝潝潟澙烯焁焈焟焬煕熂熄熈熙熹熺熻燨爔牺犀犔犠犧狶獡玺琋璽瓕瘜皙盻睎瞦矖矽硒碏磎磶礂禊禧稀稧穸窸粞糦系細綌緆縘縰繥繫纚细绤羛羲習翕翖肸肹腊膝舃舄舾莃菥葈葸蒠蒵蓆蓰蕮薂虩蜥螅螇蟋蟢蠵衋袭裼襲西覡覤觋觹觽觿誒諰謑謵譆诶谿豀豨豯貕赥赩趇趘蹊蹝躧邜郄郋郗郤鄎酅醯釐釳釸鈢鉨鉩銑錫鎴鏭鐊鑴钑铣锡閪闟阋隙隟隰隵雟雭霫霼飁餏餼饎饩饻騱騽驨鬩鯑鰓鰼鱚鳃鳛鵗鸂黖鼷"},
	"xia":    {"丅下乤侠俠傄匣厦吓呷唬嚇圷埉夏夓峡峽廈徦懗捾敮暇柙梺炠烚煆狎狭狹珨瑕疜疨瘕睱瞎硖硤碬磍祫笚筪縖罅翈舝舺芐蕸虲虾蝦螛諕谺赮轄辖遐鍜鎋鎼鏬閕閜陜陿霞颬騢魻鰕鶷黠"},
	"xian":   {"仙仚伭佡僊僩僲僴先冼县咁咞咸哯唌啣嗛嘕垷埳塪壏奾妶姭姺娊娴娹婱嫌嫺嫻嬐孅宪尟尠屳岘峴崄嶮幰廯弦彡忺憪憲憸懢挦掀掺搚搟摻撊撏攇攕显晛暹杴枮槏橌橺櫶欦毨氙洗涀涎湺溓澖瀗灦烍燹狝猃献獫獮獻玁现玹珗現甉痫癇癎盷県睍瞯硍礥祆禒秈稴筅箲籼粯糮絃絤綅綖綫線縣繊纎纖纤线缐羡羨羬胘腺臔臤臽舷苋苮莧莶薟藓藖蘚蚬蚿蛝蜆衔衘褼襳見见誢誸諴譣豏賢贒贤赻跣跹蹮躚軐輱酰醎醶銑銛銜銽鋧錎鍁鍌鏾鑦铣铦锨閑閒闲限陥险陷険險霰韅韯韱顕顯餡馅馦鮮鱻鲜鶱鷳鷴鷼鹇鹹麙麲鼸"},
	"xiang":  {"乡享亯佭像儴勨勷厢向响啍嚮夅姠嶑巷庠廂忀想晑曏栙楿橡欀湘潒珦瓖瓨相祥稥箱絴緗纕缃缿羊羏翔膷芗萫葙薌蚃蟓蠁衖襄襐詳详象跭郷鄉鄊鄕銄銗鐌鑲镶閧闂降響項项飨餉饗饟饷香驤骧鮝鯗鱌鱜鱶鲞麘"},
	"xiao":   {"侾俏俲傚削効呺咲哓哮啸嗃嘋嘐嘨嘯嘵嚣嚻囂姣婋孝宯宵小崤庨彇恔憢揱撨效敩斅斆晓暁曉枭枵校梟櫹歊歗殽毊洨消涍淆潇瀟灱灲烋焇熇熽猇獟獢痚痟皛皢睄硝硣穘窙笑筱筿箫箾篠簘簫綃绡翛肖膮莦萧蕭薂藃虈虓蛸蟂蟏蟰蠨訤詨誟誵謏踃逍郩銷销霄颵驍骁骹髇髐魈鴞鴵鷍鸮"},
	"xie":    {"些亵伳偕偞偰僁儶写冩劦勰协協卨卸叶嗋噧嚡垥塮夑奊契娎媟孈寫屑屟屧峫嶰廨徢恊愶慀懈拹挟挾揳搚携撷擕擷攜斜旪暬枻栧梋械楔榍榝榭樧歇泄泻洩渫滊澥瀉瀣灺炧炨烲焎熁燮燲爕猲獦獬瑎碿祄禼糏紲絏絜絬綊緤緳縀繲纈绁缬缷翓胁脅脇脋膎薢薤藛蝎蝢蟹蠍蠏血衺褉褻襭觟解觧諧謝讗谐谢躞躠邂邪靾鞋鞢鞵韘韰颉鮭鲑齂齘齛齥龤"},
	"xin":    {"伈伩信俽噷噺囟妡嬜孞廞心忄忻惞愖新昕杺枔欣歆炘焮盺礥脪舋芯莘薪衅襑訢訫軐辛邤釁鈊鋅鐔鑫锌镡阠顖馨馫馸"},
	"xing":   {"侀倖兴刑哘型垶姓娙婞嬹巠幸形性悻惺擤星曐杏洐涬滎煋狌猩瑆皨省睲硎箵篂緈腥臖興荇荥莕葕蛵行觪觲謃邢郉醒鈃鉶銒鋞钘铏陉陘餳饧騂骍鮏鯹"},
	"xiong":  {"兄兇凶匂匈哅夐忷恟敻汹洶焸焽熊熋胷胸芎訩詗詾讻诇賯雄"},
	"xiu":    {"休俢修咻嗅嚊宿岫峀庥朽樇溴滫潃烌珛琇璓秀糔綉繍繡绣羞脙脩臭臹苬螑袖褎褏貅銝銹鎀鏅鏥鏽锈飍饈馐髤髹鮴鱃鵂鸺齅"},
	"xu":     {"休伵侐俆偦冔勖勗卹叙吁呴喐喣嘘嘼噓圩垿墟壻妶姁婿媭嬃序徐怴恓恤惐慉戌掝揟敍敘旭旴昫晇暊朂朐栩楈槒欨欰欻歔歘殈汿沀洫浒淢湏湑溆滀滸漵潊烅烼煦燸獝珝珬畜疞盢盨盱眗瞁矞砉禑稰稸糈絮綇続緒緖緰縃續绪续聓聟胥芧蒣蓄蓲蓿蕦藇藚虗虚虛蚼蛡蝑裇訏許訹詡諝諿譃许诩谞賉邪鄦酗醑銊鑐需須頊须顼驉鬚魆魖鱮"},
	"xuan":   {"伭儇券县吅咺喛喧塇夐妶媗嫙嬛宣弲怰悬愃愋懁懸揎敻旋昍昡晅暄暅暶梋楥楦檈泫渲漩炫烜煊玄玹琁琄瑄璇璿痃癣癬眩眴睻矎碹禤箮絢縣縼繏绚翧翾萱萲蓒蔙蕿藼蘐蜁蝖蠉衒袨諠諼譞讂谖贙軒轩选選鉉鋗鍹鏇铉镟鞙颴駨駽鰚"},
	"xue":    {"乴削吷噱坹壆学學岤峃嶨彐怴敩斈桖樰泬泶滈澩瀥燢狘疦疶瞲穴膤艝茓蒆薛血袕觷謔谑趐踅轌辥辪雪靴鞾鱈鳕鷽鸴"},
	"xun":    {"伨侚偱勋勛勲勳卂咰噀噚嚑坃埙塤壎壦奞姰孙孫寻尋峋巡巺巽廵徇循恂愻揗攳旬曛杊栒桪樳殉殾毥汛洵浔浚潠潯濬灥焄煇熏燅燖燻爋爓狥獯珣璕畃眴矄稄窨紃纁臐荀荨蕁蕈薫薰蘍蟫蟳訊訓訙詢训讯询迅迿逊遜郇鄩醺鑂顨馴駨驯鱏鱘鲟鶽"},
	"ya":     {"丫亚亜亞伢俹冴劜厊压厑厓吖呀哑唖啞圔圠圧垭埡堐壓姶娅婭孲岈崕崖庌庘押拁挜掗揠枒桠椏椻氩氬涯漄潝牙犽猚猰玡琊瓛疋疨痖瘂睚砑碣磍稏窫笌聐芽蕥蚜衙襾覀訝讶軋轧迓邪釾錏鐚铔雅鴉鴨鵶鸦鸭齖齾"},
	"yan":    {"严乵俨偃偐傿儼兖兗剡剦匽厃厌厣厭厳厴咽唁唌啱喦喭噞嚈嚥嚴囐埏堰塩墕壛壧夵奄妍妟姲姸娫娮嫣嬊嬐嬮嬿孍宴岩崦嵃嵒嵓嶖巌巖巗巘巚延弇彥彦恹愝懕懨戭扊抁掞掩揅揜敥昖晏晻暥曕曣曮梴棪椻椼楌樮橪檐檿櫩欕歅殗殷沇沿洝淊淹渰渷湮溎滟演漹灎灔灧灩炎炏烟烻焉焑焔焰焱煙熖燄燕爓牪狿猏猒珚琂琰甗盐眼研砚硏硯硽碞礹窴筵篶簷綖縯罨羬胭腌臙艳艶艷芫莚菸萒蔅虤蜒蝘衍裺褗覎觃觾言訁訮詽諺讌讞讠谚谳豓豔豣贋贗赝躽軅遃郔郾鄢酀酓酽醃醶醼釅铅閆閹閻閼闫阉阎隁隒雁顏顔顩颜餍饜騐験騴驗驠验鬳魇魘鰋鳫鳱鴈鴳鶠鷃鷰鹽麣黡黤黫黬黭黰黶鼴鼹齞齴龑"},
	"yang":   {"仰佒佯傟养劷勜卬咉坱垟央姎婸岟崵崸徉怏恙愓慃懩扬抰揚攁敭旸昜暘杨柍样楊楧様樣殃氜氧氱泱洋漾瀁炀炴烊煬玚珜瑒疡痒瘍癢眏眻礢禓秧紻羊羏羕胦蛘蝆詇諹详軮輰鉠鍈鍚钖阦阳陽雵霷鞅颺飏飬養駚鰑鴦鴹鸉鸯"},
	"yao":    {"么乐仸侥倄偠傜僥匋吆咬喓嗂垚堯夭妖姚婹媱宎尧尭岆峣崾嶢嶤幺徭徺怮恌愮抭揺搖摇摿撽暚曜杳枖柼楆榚榣樂殀殽溔滧烄烑熎燿爻狕猺獟珧瑤瑶疟瘧皐眑矅磘祅穾窅窈窑窔窕窯窰筄箹約繇纅约耀肴腰舀艞苭药葯葽蓔薬藥蘨袎要覞訞詏謠謡讑谣軺轺遙遥邀邎銚鎐钥铫闄隃靿顤颻飖餆餚騕鰩鱙鳐鴁鴢鷂鷕鹞鼼齩"},
	"ye":     {"业也亪亱倻僷冶叶吔咽喝嘢噎埜堨墷壄夜射峫嶪嶫抴拽捓捙掖揲揶擖擛擨擪擫晔暍曄曅曗曳曵枼枽椰楪業歋殗洂液漜潱澲烨焆煠燁爗爷爺瑘璍瓛痷皣瞱瞸礏窫緤耶腋葉蠮謁谒邪邺鄴野釾鋣鍱鎁鎑鐷铘靥靨頁页餣饁馌驜鵺鸈"},
	"yi":     {"一乁乂义乊乙乚乛亄亦亿仡以仪伇伊伿佁佚佾侇依俋倚偯儀億儗兿冝刈劓劮勚勩匇匜医叕吚呓呭呹咦咿唈嗌噫囈圛圯坄垼埶埸墿壱壹夁夷奕妷姨媐嫕嫛嬄嬑嬟宐宜宧寱寲射尾屹峄峓崺嶧嶬嶷已巸帟帠幆庡廙异弈弋弌弬彛彜彝彞彵役忆忔怈怡怿恞悒悘悥意憶懌懝懿扅扆扡抑拸挹掜揖撎攲攺敡敧敼斁旑旖易昳晹暆曀曎杙杝枍枻柂栘栧栺桋棭椅椬椸榏槸檍檥檹欥欭歝殔殪殹毅毉沂沶泄泆洟洢浂浥浳渏湙溢漪潩澺瀷炈焬焲煕熠熤熪熼燚燡燱狋狏猗獈玴珆瑿瓵畩異疑疙疫痍痬瘗瘞瘱癔益眙睪瞖矣礒祎禕秇移稦穓竩笖箷篒簃籎縊繄繶繹绎缢羛羠義羿翊翌翳翼耴肄肊胰膉臆舣艗艤艺芅苅苡苢荑萓萟蓺薏藙藝蘙虉蚁蛇蛜蛡蛦蛾蜴螔螘螠蟻衣衤衪袘袣裔裛裿褹襼觺訲訳詍詒詣誃誼謻譩譯議讉讛议译诒诣谊豙豛豷貖貤貽贀贻跇跠軼輢轙轶辷迆迤迱迻逘逸遗遺邑郼鄓酏醫醳醷釔釴釶鈘鈠鉯銥鎰鏔鐿钇铱镒镱阣隿霬頉頤頥顊顗颐食飴饐饴駅驛驿骮鮧鮨鯣鳦鴺鶂鶃鶍鷁鷊鷖鷧鷾鸃鹝鹢鹥黓黟黳齮齸"},
	"yin":    {"乑伒侌冘凐印吟吲唫喑噖噾嚚囙因圁圻垔垠垽堙堷夤姻婣婬寅尹峾崟崯嶾廕廴引愔慇慭憖憗懚斦朄栶梀檃檭檼櫽欭歅殥殷氤泿洇洕淫淾湚湮溵滛濥濦烎犾狺猌珢璌瘖瘾癊癮硍碒磤禋秵窨筃粌絪緸縯胤苂茚茵荫荶蒑蔩蔭蘟蚓螾蟫裀訔訚訡誾諲讔赺趛輑鄞酳鈏鈝銀銦铟银闉阥阴陰陻隂隐隠隱霒霠霪靷鞇音韾飮飲饮駰骃鮣鷣齗齦龂龈"},
	"ying":   {"偀僌啨営嘤噟嚶塋夃婴媖媵嫈嬰嬴孆孾嵤巆巊应廮影応愥應摬撄攍攖旲映景暎朠柍桜桯梬楧楹樱櫻櫿浧渶溁溋滎滢潁潆濙濚濴瀅瀛瀠瀯瀴灐灜焸焽煐熒營爃珱瑛瑩璎瓔甇甖瘿癭盁盈眏矨硬碤礯禜穎籝籯緓縈纓绬缨罂罃罌耺膡膺英茔荥荧莹莺萤营萦萾蓥藀蘡蛍蝇蝧蝿螢蠅蠳褮覮謍譍譻賏贏赢軈迎郢鎣鐛鑍锳霙鞕韺頴颍颕颖鱦鴬鶑鶧鶯鷪鷹鸎鸚鹦鹰"},
	"yo":     {"哟唷喲"},
	"yong":   {"佣俑傭勇勈咏喁嗈噰埇塎墉壅嫞嵱庸廱彮怺恿悀惥愑愹慂慵拥揘擁柡栐槦永泳涌湧滽澭灉牅牗用甬痈癕癰砽硧臃苚蒏蛹詠踊踴邕郺鄘醟銿鏞镛雍雝顒颙饔鯒鰫鱅鲬鳙鷛"},
	"you":    {"丣亴优佑佦侑偤優冘卣又友右呦哊唀嚘囿妋姷孧宥尢尤峟峳幼幽庮忧怞怣怮悠憂懮揂攸斿有柚栯梄梎楢槱櫌櫾汼沋油泑浟游湵滺瀀牖牗牰犹狖猶猷獶獿由甴疣痏祐禉秞繇纋羑羪耰聈肬苃莜莠莤莸蕕蚰蚴蜏蝣訧誘诱貁輏輶迶逌逰遊邮郵鄾酉酭釉鈾銪铀铕駀魷鮋鱿鲉麀黝鼬"},
	"yu":     {"丂与乻予于亐伃伛余俁俞俣俼偊傴僪儥兪匬叞吁吾唹喅喩喻噳圄圉圫域堉堣堬奥奧妤妪娛娪娯娱婾媀媮嫗嬩宇寓寙尉屿峪崳嵎嵛嶎嶼庽庾彧御忬悆惐愈愉愚慾懙戫扜扝扵挧捓捥揄敔斔斞於旟昱杅栯桙棛棜棫楀楡楰榆櫲欎欝欤欥欲歈歟歶毓汩浴淢淤淯渔渝湡滪漁潏澚澞澦灪灹焴煜熨燏燠爩牏狱狳獄玉玗玙琙琟瑀瑜璵畬痏瘀瘉瘐癒盂盓睮矞砡硢硲礇礖礜祤禦禹禺秗稢稶穥穻窬窳竽箊篽籅籞籲粖粥紆緎纡罭羭羽翑聿肀育腴臾舁舆與艅艈芋芌茟茰萭萮萸蒮蓣蓹蕍蕷薁藇蘌蘛虞虶蜟蜮蝓螸衧袬裕褕覦觎誉語諛諭謣譽语谀谕豫貐踰軉輍輿轝込迂迃逳逾遇遹邘邪郁鄅酑醧釪鈺銉鋊錥鍝鐭钰閾阈陓隃隅隩雓雤雨雩霱預頨预飫餘饇饫馀馭騟驈驭骬髃鬰鬱鬻魊魚魣鮽鯲鰅鱊鱼鳿鴥鴧鴪鵒鷠鷸鸆鸒鹆鹬麌齬齵龉"},
	"yuan":   {"傆元円冤剈原厡厵员員喛噮囦园圆圎園圓圜垣垸塬夗妧妴媛媴嫄嬽宛寃弲怨悁惌愿掾援杬棩楥榞榬橼櫞沅涴淵渁渆渊渕湲源溒灁爰猨猿瑗盶眢禐笎箢緣縁缘羱肙芫苑茒葾蒝蒬薗薳蚖蜎蜵蝝蝯螈衏袁裫裷褑褤謜貟贠轅辕远逺遠邍邧酛鈨鋺鎱院願駌騵魭鳶鴛鵷鶢鶰鸢鸳鹓黿鼋鼘鼝"},
	"yue":    {"乐兊兌兑刖哕哾啘噦妜嬳岄岳嶽恱悅悦戉扚抈捳擽曰曱月枂栎楽樂樾櫟泧瀹爚玥矆矱礿禴箹篗籆籥籰粤粵約约臒蘥蚎蚏說説说越趯跀跃躍軏鈅鉞鑰钥钺閱閲阅髺鸑鸙黦龠龯"},
	"yun":    {"云伝傊允勻匀员員喗囩夽奫妘孕恽惲愠愪慍抎抣昀晕暈枟榅橒殒殞氲氳沄涒涢溳澐煇煴煾熅熉熨狁玧畇眃磒秐筠筼篔紜緷緼縕縜纭缊耘耺腪芸荺菀蒀蒕蒷蕓蕰蕴薀藴蘊蝹褞貟賱贇贠赟运運郓郧鄆鄖酝醖醞鈗鋆阭陨隕雲霣鞰韗韞韫韵韻餫馧馻齫齳"},
	"za":     {"偺匝咂咋咱喒噈囃囋囐帀扎拶杂桚沞沯砸紥紮臜臢襍迊鉔雑雜雥韴魳"},
	"zai":    {"仔傤儎再哉在宰崽扗抂栽洅渽溨災灾烖甾畠睵縡菑賳載载酨"},
	"zan":    {"偺儧儹兂咱喒噆囋寁拶揝撍攅攒攢昝暂暫桚橵濽灒瓉瓒瓚禶穳篸簪簮糌襸讃讚賛贊赞趱趲蹔鄼酂酇錾鏨鏩鐕鐟饡鵤"},
	"zang":   {"匨塟奘弉牂羘脏臓臟臧葬蔵藏賍賘贓贜赃銺駔驵髒"},
	"zao":    {"傮凿唕唣喿噪慥早枣栆梍棗澡灶煰燥璅璪皁皂竃竈簉糟艁薻藻蚤譟趮蹧躁造遭醩鑿"},
	"ze":     {"仄侧側则則咋唶啧啫嘖嫧崱帻幘庂択择捑擇昃昗樍歵汄沢泎泽溭滜澤灂皟睪瞔矠礋稄笮箦簀耫舴蔶蠌襗諎謮責賾责赜迮飵鸅齚齰"},
	"zei":    {"戝蠈賊贼鰂鱡鲗"},
	"zen":    {"怎撍譖谮"},
	"zeng":   {"囎増增憎曽曾橧熷璔甑矰磳竲綜縡繒综缯罾譄贈赠鄫鋥锃鬵鱛"},
	"zha":    {"乍偧剳劄厏吒咋咤哳喋喥喳奓宱怍扎抯拃挓揸搾摣札柞柤查査栅楂榨樝渣溠潳灹炸煠牐甴痄皶皻眨砟箚紥紮耫苲蚱蚻觰詐諎譇譗诈軋醡鍘铡閘闸霅飵鮓鮺鲊鲝齄齇"},
	"zhai":   {"亝侧债側債厇厏哜嚌夈宅寨抧择捚摘擇擿斋斎榸檡瘵砦祭窄粂翟責责鉙齊齋齐"},
	"zhan":   {"佔偡僝占噡嫸展岾崭嵁嶃嶄嶘嶦怗惉战戦戰搌斩斬旃旜枬栈栴桟棧榐橏欃毡氈氊沾湛琖皽盏盞瞻碊站粘綻绽菚薝蘸虥虦蛅覱詀詹譧譫讝谵趈輾轏辗邅醆閚霑颤颭飐飦饘驏驙魙鱣鳣鳽鸇鹯黵龪"},
	"zhang":  {"丈仉仗仧傽兏墇嫜嶂帐帳幛幥张弡張彰慞扙掌暲杖樟涨涱漲漳獐璋痮瘬瘴瞕礃章粀粻胀脹蔁蟑賬账遧鄣鏱長长障鞝餦騿鱆麞"},
	"zhao":   {"佋兆召啁啅嘲垗妱巶找招旐昭曌朝枛棹櫂沼濯炤照燳爪爫狣瑵皽着瞾笊箌罀罩羄肁肇肈菬詔诏赵趙釗釽鉊鍣钊駋鮡鳭"},
	"zhe":    {"乇乽厇哲啠喆嗻嚞埑嫬悊慹扸折搩摺晢晣杔柘棏樜歽浙淛着矺砓磔籷粍者聑著蔗虴蛰蜇螫蟄蟅袩褶襵詟謫謺讁讋谪赭踷輒輙轍辄辙这這遮銸鍺锗馲鮿鷓鹧"},
	"zhei":   {"这"},
	"zhen":   {"侦侲偵圳塦填姫嫃寊屒帧帪幀弫抮挋振揕搸斟昣朕枕枮栚桢桭椹楨榐榛槙樼殝浈湞溱潧澵獉珍珎瑧瑱甄甽畛疹眕眞真眹砧碪祯禎禛稹竧箴籈紖紾絼縝縥纼缜聄胗臻萙葴蒖蓁薽袗裖覙診誫诊貞賑贞赈軫轃轸辴遉酙針鉁鋴錱鍼鎭鎮针镇阵陣震駗鬒鱵鴆鸩黰"},
	"zheng":  {"丁争佂凧埥埩塣姃媜峥崝崢幁征徎徰徴徵怔愸憕抍拯挣掙掟揁撜政整晸正氶炡烝爭狰猙症癥眐睁睜筝箏篜糽綪聇脀蒸証諍證证诤踭郑鄭鉦錚鏳钲铮靕鬇鴊"},
	"zhi":    {"之乿侄俧倁値值偫傂儨凪制剬劕劧卮厔只吱呮咫址坁坧垁埴執墆墌夂妷姪娡嬂寘峙崻巵帋帙帜幟庢庤廌彘徏徝徴徵志忮恉慹憄懥懫戠执扺扻抧拓挃指挚捗掷搘搱摕摨摭摯擲擳擿支斦旘旨晊智杫枝枳柣栀栉栺桎梔梽植椥楖榰槯樴櫍櫛止歭殖氏汁汥汦沚治泜洔洷淔淽滍滞滯漐潌潪瀄炙熫犆狾猘璏瓆瓡畤疐疷疻痓痔痣瘈直眰知砋礩祇祉祑祗祬禃禔秇秓秖秩秲秷稙稚稺穉窒筫紙紩絷絺綕緻縶織纸织置翐聀职聜職肢胑胝脂膣膱至致臷臸芖芝芷茋菭藢蘵蚔蛭蜘螲蟙衹衼袟袠製襧覟觗觯觶訨誌识豑豒豸貭質贄质贽趾跖跱踬踯蹠躑躓軄軹輊轵轾迣郅郦酯釞銍銴鋕鑕铚锧阯陟隲隻雉馶馽駤騭騺驇骘鯯鳷鴙鴲鶨鷙鸷黹鼅"},
	"zhong":  {"中仲伀众偅冢刣喠堹塚夂妐妕媑尰幒彸徸忠忪斔柊歱汷泈炂煄狆瘇盅眾祌种種穜筗籦終緟终肿腫舯茽蔠蚛蚣蜙螤螽衆衳衶衷諥踵蹱重鈡銿鍾鐘钟锺鴤鼨"},
	"zhou":   {"伷侜僽冑周呪咒咮啁喌噣嚋妯婤宙州帚徟掫昼晝晭椆洲淍炿烐珘甃疛皱皺盩睭矪碡箒籀籒籕粙粥紂縐繇纣绉翢肘胄胕舟荮菷葤薵詋謅譸诌诪调賙赒軸輈輖轴辀週郮酎銂霌駎駲騆驟骤鯞鵃鸼"},
	"zhu":    {"丶主乼伫佇住侏劚助劯嗻嘱囑坾墸壴孎宁宔尌属屬嵀帾拀拄敱斀斸曯朮术朱杼柱柷株槠樦橥櫡櫧櫫欘殶注洙渚潴澍濐瀦灟炢炷烛煑煮燝燭爥猪珠疰瘃眝瞩矚砫硃磩祝祩秼窋竚竹竺笁笜筑筯箸築篫篴簗紵紸絑纻罜羜翥舳芧苎苧茱茿莇著蓫蕏藸蛀蛛蝫蠋蠩蠾袾褚註詝誅諸诛诸豬貯贮跓跦躅軴迬逐逫邾鉒銖鋳鑄钃铢铸陼飳馵駐駯驻鮢鯺鱁鴸鸀麆麈鼄"},
	"zhua":   {"抓挝撾檛爪爫簻膼髽"},
	"zhuai":  {"拽睉跩转"},
	"zhuan":  {"专传傳僎僝剸叀啭囀堟塼嫥孨専專恮撰漙灷瑑瑼甎砖磗磚竱篆篹篿簨籑縳耑腞膞蒃蟤襈諯譔賺贃赚転轉转鄟顓颛饌馔鱄"},
	"zhuang": {"僮壮壯壵奘妆妝娤幢庄庒戅戆戇撞桩梉樁湷漴焋状狀粧糚荘莊装裝"},
	"zhui":   {"坠墜奞娷惴揣椎沝甀畷硾礈笍綴縋缀缒腏膇諈贅赘追醊錐錣鑆锥隊隹騅骓鵻"},
	"zhun":   {"准凖啍圫埻宒屯忳旽淳準稕窀綧肫衠訰諄谆迍"},
	"zhuo":   {"丵倬剢劅卓叕啄啅噣圴墌妰娺彴拙捉捔撯擆擢斀斫斮斱斲斵晫桌梲棁棳棹椓槕櫡汋浊浞涿淖濁濯灂灼炪烵焯燋犳琢琸着硺禚穛穱窡窧篧籗籱繳缴罬聉茁著蓔藋蠗蠿諁諑謶诼趠酌鋜鐯鐲镯鵫鷟龺"},
	"zi":     {"仔倳兹剚吇吱呰呲咨啙嗞嗭姉姊姕姿子孖字孜孳孶崰嵫恣杍栥梓椔榟橴沝洓淄渍湽滋滓漬澬牸玆璾甾畠眥眦矷禌秄秭秶稵笫籽粢紎紫緇緕纃缁耔胏胔胾自芓茈茊茡茲荢菑葘蓻薋虸觜訾訿諮谘貲資赀资赼趑趦輜輺辎鄑釨鈭錙鍿鎡锱镃頾頿髭鯔鰦鲻鶅鼒齍齐齜龇"},
	"zong":   {"从倊倧偬傯堫宗嵏嵕嵸從总惣惾捴揔搃摠昮朡枞棕椶樅潀潈潨焧熜熧燪猔猣疭瘲碂磫稯粽糉糭綜緃総緫緵縂縦縱總繌纵综翪腙葼蓗蝬豵踨踪蹤鍐鑁騌騣骔鬃鬉鬷鯮鯼"},
	"zou":    {"奏媰掫揍搊棷棸楱箃緅菆諏诹走赱邹郰鄒鄹陬騶驺鯐鯫鲰黀齱齺龰"},
	"zu":     {"俎倅卆卒哫唨崒崪怚族柤椊爼珇祖租稡箤組组菹葅蒩詛诅足踤鎐鎺鏃镞阻靻"},
	"zuan":   {"劗揝攥篹籫繤纂纉纘缵賺躜躦鉆鑚鑽钻"},
	"zui":    {"冣厜咀嗺嘴噿嶊嶵晬最朘栬槜樶檇檌欈濢璻睟祽稡絊纗罪蕞蟕觜辠酔酻醉鋷錊"},
	"zun":    {"僔噂墫尊嶟拵捘撙栫樽瀳繜罇袸譐遵銌鐏鱒鳟鶎鷷"},
	"zuo":    {"佐作侳做咗唑嘬坐岝岞左座怍捽撮昨柞柮琢祚秨稓穝笮筰糳繓胙苲莋葃葄蓙袏諎酢鈼阼飵"},
}
