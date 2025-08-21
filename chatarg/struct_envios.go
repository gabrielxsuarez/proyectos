package main

type ClientInfo struct {
	Seq        int    `json:"seq"`
	Cmd        string `json:"cmd"`
	LocalTime  int64  `json:"localtime"`
	TzOffset   int    `json:"tzoffset"`
	ClientInfo string `json:"clientinfo"`
}

type EmbedInfo struct {
	Seq      int    `json:"seq"`
	Channel  string `json:"channel"`
	Cmd      string `json:"cmd"`
	Referrer string `json:"referrer"`
}

type ConnectIRC struct {
	Seq          int    `json:"seq"`
	Cmd          string `json:"cmd"`
	Channel      string `json:"channel"`
	Data         string `json:"data"`
	Nick         string `json:"nick"`
	Pass         string `json:"pass"`
	AuthMethod   string `json:"authmethod"`
	JoinChannels string `json:"joinchannels"`
	Charset      string `json:"charset"`
}

type SendText struct {
	Seq     int    `json:"seq"`
	Cmd     string `json:"cmd"`
	Chan    string `json:"chan"`
	Data    string `json:"data"`
	Channel string `json:"channel"`
	Reply   string `json:"reply,omitempty"`
}

type KeepAlive struct {
	Seq int    `json:"seq"`
	Cmd string `json:"cmd"`
	M   string `json:"m"`
}

type ListChannel struct {
	Seq          int    `json:"seq"`
	Cmd          string `json:"cmd"`
	LocalChannel string `json:"localchannel"`
	Channel      string `json:"channel"`
	ChannelType  string `json:"channeltype"`
}

type SimpleMessage struct {
	Seq     int    `json:"seq"`
	Cmd     string `json:"cmd"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}