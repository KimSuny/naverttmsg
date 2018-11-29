package naverttmsg

/*
https://github.com/navertalk/chatbot-api 참고.
*/

type message struct {
	Standby          bool              `json:"standby,omitempty"` //핸드오버(채팅에 사람이 직접 개입되는) 관련: https://github.com/navertalk/chatbot-api/blob/master/handover_v1.md
	Event            string            `json:"event"`             //이벤트명
	User             string            `json:"user"`              // 유저 식별값
	Partner          string            `json:"partner"`
	TextContent      *textContent      `json:"textContent,omitempty"`  //텍스트 메시지
	ImageContent     *imageContent     `json:"imageContent,omitempty"` //단수 이미지로 구성된 메시지
	CompositeContent *compositeContent `json:"compositeContent"`       //텍스트와 이미지 그리고 버튼을 포함하는 복합 메시지
	Options          *options          `json:"options,omitempty"`      //추가 속성
}

type textContent struct {
	Text       string      `json:"text"`                 //전송하고자 하는 텍스트
	Code       string      `json:"code,omitempty"`       //compositeContent 에서 버튼클릭시 전달받는 코드값, code는 채팅에 노출되지 않습니다.
	InputType  string      `json:"inputType,omitempty"`  //typing|button|sticker|inquiry|vphone 수신시에만 받는 속성으로, 사용자가 어떤 매개체로 봇에게 입력하는지를 나타내는 값입니다.
	QuickReply *quickReply `json:"quickReply,omitempty"` //퀵버튼 - 빠른응답
}

type imageContent struct {
	ImageUrl   string      `json:"imageUrl"`             //외부에서도 접근가능한 URL이여야 합니다.
	QuickReply *quickReply `json:"quickReply,omitempty"` //퀵버튼 - 빠른응답
}

type compositeContent struct {
	CompositeList []compositeList `json:"compositeList"` //composite 데이터
}

type compositeList struct {
	Title       string       `json:"title,,omitempty"`      //타이틀
	Description string       `json:"description,omitempty"` //설명
	Image       *image       `json:"image,omitempty"`       //이미지
	ElementList *elementList `json:"elementList,omitempty"` //요소 리스트
	ButtonList  []button     `json:"buttonList,omitempty"`  //버튼 요소 리스트
}

type elementList struct {
	Type string        `json:"type"` //리스트 요소의 타입. 현재는 LIST 타입만 존재, 필수
	Data []elementData `json:"data"` //리스트 요소의 데이터, 필수
}

type image struct {
	ImageURL string `json:"imageUrl"` //이미지 URL
}

//Button Object
type button struct {
	Type string      `json:"type"` //버튼 요소의 타입. TEXT, LINK, OPTION, PAY, 필수
	Data *buttonData `json:"data"` //버튼 요소의 데이터, 필수
}

//ButtonData Object (LINK 타입일땐 ButtonList 사용 안함)
type buttonData struct {
	Title      string   `json:"title"`          //버튼에 노출되는 텍스트. 유저가 버튼을 클릭하면 전송되는 텍스트, 필수
	Url        string   `json:"url"`            //PC버전 채팅창에서 버튼을 클릭하면 이동할 페이지 URL, 필수
	MobileUrl  string   `json:"mobileUrl"`      //모바일 버전 채팅창에서 버튼을 클릭하면 이동할 페이지 URL, 필수
	Code       string   `json:"code,omitempty"` //유저가 버튼을 클릭하면 전송되는 코드
	ButtonList []button `json:"buttonList"`     //버튼 클릭시 채팅창 하단에 노출되는 버튼의 목록 TEXT, LINK, PAY
}

//ElementData Object (LIST 타입)
type elementData struct {
	Title          string  `json:"title"`                    //"리스트 요소 타이틀", 필수, 최대 10자까지 노출되고 그 이상은 줄임표시(...)됩니다.
	Description    string  `json:"description,omitempty"`    //"리스트 요소 설명1", 최대 25자까지 노출되고 그 이상은 줄임표시(...)됩니다.
	SubDescription string  `json:"subDescription,omitempty"` //"리스트 요소 설명2", 최대 13자까지 노출되고 그 이상은 줄임표시(...)됩니다.
	Image          *image  `json:"image,omitempty"`          //이미지, 입력안하면 기본 이미지가 노출됩니다.
	Button         *button `json:"button,omitempty"`         //버튼은 TEXT와 LINK타입만 허용되며 title길이는 10자로 제한됩니다.
}

type quickReply struct {
	ButtonList []button `json:"buttonList"` //버튼 리스트
}

type options struct {
	Inflow       string `json:"inflow"`       //어떤방식으로 유입되었는지를 구분 button|list|none
	Referer      string `json:"referer"`      //유입페이지 URL
	From         string `json:"from"`         //from을 값을 전달 받을 수 있다.
	Friend       bool   `json:"friend"`       //false: 친구아님,  true: 친구
	Under14      bool   `json:"under14"`      //false: 만14세이상,  true: 만14세미만
	Under19      bool   `json:"under19"`      //false: 만19세이상,  true: 만19세미만
	Set          string `json:"set"`          //on: 친구추가,  off: 친구철회
	Notification bool   `json:"notification"` //push를 보낼때 사용
	Action       string `json:"action"`       //대화창에서 추가적인 행위가 필요할 때 사용할 수 있는 이벤트 typingOn|typingOff
	Control      string `json:"control"`
	TargetId     int    `json:"targetId"`
}

type nerror struct {
	Success       bool   `json:"success"`                 //API호출 성공여부, false시 resultCode에따라 대응한다.
	ResultCode    string `json:"resultCode"`              //"00"코드는 성공이며, 그외 실패시 원인이되는 코드값을 반환한다.
	ResultMessage string `json:"resultMessage,omitempty"` //resultCode를 구체적으로 설명한다.
}

type persistentMenu struct { //persistentMenu(이하 고정 메뉴)는 사용자가 대화 중에 상시로 접근할 수 있는 메뉴입니다.
	Event       string        `json:"event"`
	MenuContent []menuContent `json:"menuContent"` //menuContent는 리스트 형태로 받습니다. menuContent를 초기화하려면 빈 리스트([])를 보냅니다.
}

type menuContent struct {
	Menus []menu `json:"menus"` //메뉴 목록. 최대 4개까지 추가할 수 있습니다. menus는 null일 수 없습니다.
}

type menu struct {
	Type string   `json:"type"` //메뉴 요소의 타입. TEXT, LINK, NESTED
	Data menuData `json:"data"` //메뉴 요소의 데이터
}

type menuData struct {
	Title     string `json:"title"` //메뉴에 노출되는 텍스트. 사용자가 메뉴를 클릭하면 채팅창에 나타나는 텍스트. 최대 길이는 20자입니다.
	Code      string `json:"code"`  //사용자가 메뉴를 클릭하면 클라이언트에게 전송되는 코드.
	Url       string `json:"url"`   //url에 tel:{전화번호}값을 넣으면 모바일에서는 전화 연결이 가능
	MobileUrl string `json:"mobileUrl,omitempty"`
}
