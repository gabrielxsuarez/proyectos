package main

type SessionID struct {
	Seq       int    `json:"seq"`
	SessionID string `json:"sessionid"`
}

type Status struct {
	Seq             int    `json:"seq"`
	PendingDNS      int    `json:"pendingDNS"`
	Cmd             string `json:"cmd"`
	Pending         int    `json:"pending"`
	Hostname        string `json:"hostname"`
	Connections     int    `json:"connections"`
	Channel         string `json:"channel"`
	ReadyToConnect  bool   `json:"readyToConnect"`
}

type Connected struct {
	Seq     int    `json:"seq"`
	Cmd     string `json:"cmd"`
	Channel string `json:"channel"`
	Name    string `json:"name"`
	Network string `json:"network"`
}

type Init struct {
	Seq     int    `json:"seq"`
	Cmd     string `json:"cmd"`
	Nick    string `json:"nick"`
	Channel string `json:"channel"`
	CTime   string `json:"ctime"`
}

type NickInfo struct {
	Host string `json:"host"`
	Nick string `json:"nick"`
}

type NickList struct {
	Seq          int        `json:"seq"`
	Cmd          string     `json:"cmd"`
	LocalChannel string     `json:"localchannel"`
	Channel      string     `json:"channel"`
	Nicks        []NickInfo `json:"nicks"`
	ChannelType  string     `json:"channeltype"`
}

type Topic struct {
	Seq          int    `json:"seq"`
	Cmd          string `json:"cmd"`
	Channel      string `json:"channel"`
	Nick         string `json:"nick"`
	Topic        string `json:"topic"`
	Time         int64  `json:"time"`
	LocalChannel string `json:"localchannel"`
	ChannelType  string `json:"channeltype"`
}

type TopicWho struct {
	Seq          int    `json:"seq"`
	Cmd          string `json:"cmd"`
	LocalChannel string `json:"localchannel"`
	Channel      string `json:"channel"`
	ChannelType  string `json:"channeltype"`
	Creator      string `json:"creator"`
	Date         string `json:"date"`
}

type Message struct {
	Seq          int    `json:"seq"`
	Cmd          string `json:"cmd"`
	Nick         string `json:"nick"`
	LocalChannel string `json:"localchannel"`
	ChannelType  string `json:"channeltype"`
	Channel      string `json:"channel"`
	Msg          string `json:"msg"`
	ReplyTo      string `json:"replyTo,omitempty"`
}

type Emote struct {
	Seq          int    `json:"seq"`
	Cmd          string `json:"cmd"`
	CTime        string `json:"ctime"`
	Nick         string `json:"nick"`
	LocalChannel string `json:"localchannel"`
	ChannelType  string `json:"channeltype"`
	Channel      string `json:"channel"`
	Emote        string `json:"emote"`
}

type Typing struct {
	Seq     int    `json:"seq"`
	Cmd     string `json:"cmd"`
	Nick    string `json:"nick"`
	Channel string `json:"channel"`
	Typing  bool   `json:"typing"`
}

type Join struct {
	Seq          int    `json:"seq"`
	Cmd          string `json:"cmd"`
	User         string `json:"user"`
	Host         string `json:"host"`
	Nick         string `json:"nick"`
	LocalChannel string `json:"localchannel"`
	Channel      string `json:"channel"`
	ChannelType  string `json:"channeltype"`
	CTime        string `json:"ctime"`
}

type Part struct {
	Seq     int    `json:"seq"`
	Cmd     string `json:"cmd"`
	Nick    string `json:"nick"`
	User    string `json:"user"`
	Host    string `json:"host"`
	Quit    bool   `json:"quit"`
	Message string `json:"message"`
	Channel string `json:"channel"`
	CTime   string `json:"ctime"`
}

type ChangeNick struct {
	Seq     int    `json:"seq"`
	Cmd     string `json:"cmd"`
	Nick    string `json:"nick"`
	NewNick string `json:"newnick"`
	Channel string `json:"channel"`
	CTime   string `json:"ctime"`
}

type Away struct {
	Seq     int    `json:"seq"`
	Cmd     string `json:"cmd"`
	Nick    string `json:"nick"`
	Channel string `json:"channel"`
}

type UserInfo struct {
	Seq          int    `json:"seq"`
	Cmd          string `json:"cmd"`
	Nick         string `json:"nick"`
	LocalChannel string `json:"localchannel"`
	Channel      string `json:"channel"`
	ChannelType  string `json:"channeltype"`
	Ct           int64  `json:"ct"`
	Info         string `json:"info"`
	Cc           string `json:"cc"`
	Tz           string `json:"tz"`
}

type UserIcon struct {
	Seq          int    `json:"seq"`
	Cmd          string `json:"cmd"`
	Nick         string `json:"nick"`
	Channel      string `json:"channel"`
	Icon         string `json:"icon"`
	ChannelType  string `json:"channeltype"`
	LocalChannel string `json:"localchannel"`
}

type Avatar struct {
	Seq          int    `json:"seq"`
	Cmd          string `json:"cmd"`
	Nick         string `json:"nick"`
	LocalChannel string `json:"localchannel"`
	Channel      string `json:"channel"`
	ChannelType  string `json:"channeltype"`
	UserIcon     string `json:"userIcon"`
}

type Kick struct {
	Seq          int    `json:"seq"`
	Cmd          string `json:"cmd"`
	Nick         string `json:"nick"`
	Kicker       string `json:"kicker"`
	LocalChannel string `json:"localchannel"`
	Channel      string `json:"channel"`
	ChannelType  string `json:"channeltype"`
	CTime        string `json:"ctime"`
	Reason       string `json:"reason"`
}

type Ban struct {
	Seq         int    `json:"seq"`
	Cmd         string `json:"cmd"`
	By          string `json:"by"`
	Nick        string `json:"nick"`
	Channel     string `json:"channel"`
	ChannelType string `json:"channeltype"`
	CTime       string `json:"ctime"`
}

type Unban struct {
	Seq         int    `json:"seq"`
	Cmd         string `json:"cmd"`
	By          string `json:"by"`
	Nick        string `json:"nick"`
	Channel     string `json:"channel"`
	ChannelType string `json:"channeltype"`
	CTime       string `json:"ctime"`
}

type Settings struct {
	UserListShowIcons bool   `json:"userListShowIcons"`
	UserListSort      string `json:"userListSort"`
	Nick              string `json:"nick"`
	Channel           string `json:"channel"`
	Server            string `json:"server"`
	Port              string `json:"port"`
	SSL               bool   `json:"ssl"`
	WebircGateway     string `json:"webircGateway"`
}

type SettingsMessage struct {
	Seq          int      `json:"seq"`
	Cmd          string   `json:"cmd"`
	LocalChannel string   `json:"localchannel"`
	Channel      string   `json:"channel"`
	ChannelType  string   `json:"channeltype"`
	Settings     Settings `json:"settings"`
}

type MOTD struct {
	Seq          int      `json:"seq"`
	Cmd          string   `json:"cmd"`
	LocalChannel string   `json:"localchannel"`
	Channel      string   `json:"channel"`
	ChannelType  string   `json:"channeltype"`
	MOTD         []string `json:"motd"`
}

type Notice struct {
	Seq     int    `json:"seq"`
	Cmd     string `json:"cmd"`
	Nick    string `json:"nick"`
	Notice  string `json:"notice"`
	CTime   string `json:"ctime"`
	Channel string `json:"channel"`
}

type ServerText struct {
	Seq          int    `json:"seq"`
	Cmd          string `json:"cmd"`
	Text         string `json:"text"`
	LocalChannel string `json:"localchannel"`
	Channel      string `json:"channel"`
	ChannelType  string `json:"channeltype"`
}

type GenericMessage struct {
	Seq interface{} `json:"seq"`
	Cmd string      `json:"cmd"`
}