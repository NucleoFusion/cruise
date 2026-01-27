package config

type Styles struct {
	Text             string `mapstructure:"text" toml:"text"`
	SubtitleText     string `mapstructure:"subtitle_text" toml:"subtitle_text"`
	SubtitleBg       string `mapstructure:"subtitle_bg" toml:"subtitle_bg"`
	UnfocusedBorder  string `mapstructure:"unfocused_border" toml:"unfocused_border"`
	FocusedBorder    string `mapstructure:"focused_border" toml:"focused_border"`
	HelpKeyBg        string `mapstructure:"help_key_bg" toml:"help_key_bg"`
	HelpKeyText      string `mapstructure:"help_key_text" toml:"help_key_text"`
	HelpDescText     string `mapstructure:"help_desc_text" toml:"help_desc_text"`
	MenuSelectedBg   string `mapstructure:"menu_selected_bg" toml:"menu_selected_bg"`
	MenuSelectedText string `mapstructure:"menu_selected_text" toml:"menu_selected_text"`
	ErrorText        string `mapstructure:"error_text" toml:"error_text"`
	ErrorBg          string `mapstructure:"error_bg" toml:"error_bg"`
	PopupBorder      string `mapstructure:"popup_border" toml:"popup_border"`
	PlaceholderText  string `mapstructure:"placeholder_text" toml:"placeholder_text"`
	MsgText          string `mapstructure:"msg_text" toml:"msg_text"`
}
