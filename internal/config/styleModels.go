package config

type Styles struct {
	Text             string `mapstructure:"text"`
	SubtitleText     string `mapstructure:"subtitle_text"`
	SubtitleBg       string `mapstructure:"subtitle_bg"`
	UnfocusedBorder  string `mapstructure:"unfocused_border"`
	FocusedBorder    string `mapstructure:"focused_border"`
	HelpKeyBg        string `mapstructure:"help_key_bg"`
	HelpKeyText      string `mapstructure:"help_key_text"`
	HelpDescText     string `mapstructure:"help_desc_text"`
	MenuSelectedBg   string `mapstructure:"menu_selected_bg"`
	MenuSelectedText string `mapstructure:"menu_selected_text"`
	ErrorText        string `mapstructure:"error_text"`
	ErrorBg          string `mapstructure:"error_bg"`
	PopupBorder      string `mapstructure:"popup_border"`
}
