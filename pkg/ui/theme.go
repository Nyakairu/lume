package ui

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/charmbracelet/lipgloss"
	"github.com/Tyooughtul/lume/pkg/scanner"
)

// Theme defines a complete color theme
type Theme struct {
	Name        string `json:"name"`
	Description string `json:"description"`

	// Core colors
	Primary   string `json:"primary"`   // primary (titles, selected items)
	Secondary string `json:"secondary"` // secondary (success, idle state)
	Accent    string `json:"accent"`    // accent (special highlights)
	Danger    string `json:"danger"`    // danger / alert
	Warning   string `json:"warning"`   // warning / attention
	Success   string `json:"success"`   // success / complete

	// Neutral colors
	Background string `json:"background"` // background
	Foreground string `json:"foreground"` // foreground / text
	Gray       string `json:"gray"`       // medium gray
	LightGray  string `json:"light_gray"` // light gray (secondary text)
	Dim        string `json:"dim"`        // dim gray (dividers)

	// Interactive states
	SelectedBg string `json:"selected_bg"` // selected background
	SelectedFg string `json:"selected_fg"` // selected foreground
	Border     string `json:"border"`      // border color
}

// Lipgloss colors
func (t *Theme) PrimaryColor() lipgloss.Color   { return lipgloss.Color(t.Primary) }
func (t *Theme) SecondaryColor() lipgloss.Color { return lipgloss.Color(t.Secondary) }
func (t *Theme) AccentColor() lipgloss.Color    { return lipgloss.Color(t.Accent) }
func (t *Theme) DangerColor() lipgloss.Color    { return lipgloss.Color(t.Danger) }
func (t *Theme) WarningColor() lipgloss.Color   { return lipgloss.Color(t.Warning) }
func (t *Theme) SuccessColor() lipgloss.Color   { return lipgloss.Color(t.Success) }
func (t *Theme) ForegroundColor() lipgloss.Color { return lipgloss.Color(t.Foreground) }
func (t *Theme) GrayColor() lipgloss.Color      { return lipgloss.Color(t.Gray) }
func (t *Theme) LightGrayColor() lipgloss.Color { return lipgloss.Color(t.LightGray) }
func (t *Theme) DimColor() lipgloss.Color       { return lipgloss.Color(t.Dim) }
func (t *Theme) SelectedBgColor() lipgloss.Color { return lipgloss.Color(t.SelectedBg) }
func (t *Theme) SelectedFgColor() lipgloss.Color { return lipgloss.Color(t.SelectedFg) }
func (t *Theme) BorderColor() lipgloss.Color    { return lipgloss.Color(t.Border) }

// PresetThemes contains all built-in themes
var PresetThemes = map[string]Theme{
	"modern": {
		Name:        "modern",
		Description: "Modern Cyber (default)",
		Primary:     "#00d4ff", // neon cyan
		Secondary:   "#00ff88", // neon green
		Accent:      "#ff00ff", // neon purple
		Danger:      "#ff3366", // neon red
		Warning:     "#ffcc00", // neon yellow
		Success:     "#00ff88", // neon green
		Foreground:  "#ffffff",
		Gray:        "#6b7280",
		LightGray:   "#9ca3af",
		Dim:         "#4e4e4e",
		SelectedBg:  "#0a3d62",
		SelectedFg:  "#ffffff",
		Border:      "#00d4ff",
	},
	"retro": {
		Name:        "retro",
		Description: "Retro Terminal",
		Primary:     "#33ff33", // matrix green
		Secondary:   "#00ff00", // bright green
		Accent:      "#ffff00", // amber yellow
		Danger:      "#ff3333", // dark red
		Warning:     "#ffaa00", // amber
		Success:     "#33ff33", // matrix green
		Foreground:  "#33ff33", // matrix green
		Gray:        "#228822",
		LightGray:   "#44aa44",
		Dim:         "#115511",
		SelectedBg:  "#004400",
		SelectedFg:  "#33ff33",
		Border:      "#228822",
	},
	"amber": {
		Name:        "amber",
		Description: "Amber Monitor",
		Primary:     "#ffb000", // amber
		Secondary:   "#ffcc00", // bright amber
		Accent:      "#ffdd44", // light amber
		Danger:      "#ff6600", // orange red
		Warning:     "#ffaa00", // orange amber
		Success:     "#cc9900", // dark amber
		Foreground:  "#ffb000",
		Gray:        "#996600",
		LightGray:   "#cc8800",
		Dim:         "#553300",
		SelectedBg:  "#442200",
		SelectedFg:  "#ffb000",
		Border:      "#996600",
	},
	"ocean": {
		Name:        "ocean",
		Description: "Deep Ocean",
		Primary:     "#4fc3f7", // sky blue
		Secondary:   "#80cbc4", // teal
		Accent:      "#ff80ab", // coral pink
		Danger:      "#ef5350", // coral red
		Warning:     "#ffca28", // gold
		Success:     "#66bb6a", // sea green
		Foreground:  "#e3f2fd",
		Gray:        "#78909c",
		LightGray:   "#b0bec5",
		Dim:         "#37474f",
		SelectedBg:  "#1565c0",
		SelectedFg:  "#e3f2fd",
		Border:      "#4fc3f7",
	},
	"highcontrast": {
		Name:        "highcontrast",
		Description: "High Contrast (Accessibility)",
		Primary:     "#ffffff", // pure white
		Secondary:   "#00ff00", // pure green
		Accent:      "#ffff00", // pure yellow
		Danger:      "#ff0000", // pure red
		Warning:     "#ffff00", // pure yellow
		Success:     "#00ff00", // pure green
		Foreground:  "#ffffff",
		Gray:        "#888888",
		LightGray:   "#cccccc",
		Dim:         "#666666",
		SelectedBg:  "#ffffff",
		SelectedFg:  "#000000",
		Border:      "#ffffff",
	},
	"dracula": {
		Name:        "dracula",
		Description: "Dracula",
		Primary:     "#bd93f9", // purple
		Secondary:   "#50fa7b", // green
		Accent:      "#ff79c6", // pink
		Danger:      "#ff5555", // red
		Warning:     "#f1fa8c", // yellow
		Success:     "#50fa7b", // green
		Foreground:  "#f8f8f2",
		Gray:        "#6272a4",
		LightGray:   "#8be9fd",
		Dim:         "#44475a",
		SelectedBg:  "#44475a",
		SelectedFg:  "#f8f8f2",
		Border:      "#6272a4",
	},
	"solarized": {
		Name:        "solarized",
		Description: "Solarized Dark",
		Primary:     "#268bd2", // blue
		Secondary:   "#2aa198", // cyan
		Accent:      "#d33682", // magenta
		Danger:      "#dc322f", // red
		Warning:     "#b58900", // yellow
		Success:     "#859900", // green
		Foreground:  "#839496",
		Gray:        "#586e75",
		LightGray:   "#93a1a1",
		Dim:         "#073642",
		SelectedBg:  "#073642",
		SelectedFg:  "#eee8d5",
		Border:      "#586e75",
	},
	"monokai": {
		Name:        "monokai",
		Description: "Monokai",
		Primary:     "#66d9ef", // cyan
		Secondary:   "#a6e22e", // green
		Accent:      "#f92672", // pink
		Danger:      "#f92672", // red
		Warning:     "#e6db74", // yellow
		Success:     "#a6e22e", // green
		Foreground:  "#f8f8f2",
		Gray:        "#75715e",
		LightGray:   "#ae81ff",
		Dim:         "#49483e",
		SelectedBg:  "#49483e",
		SelectedFg:  "#f8f8f2",
		Border:      "#75715e",
	},
}

// ThemeManager manages theme configuration
type ThemeManager struct {
	CurrentTheme Theme
	AllThemes    map[string]Theme
	ConfigPath   string
}

// NewThemeManager creates a theme manager
func NewThemeManager() *ThemeManager {
	tm := &ThemeManager{
		AllThemes:  make(map[string]Theme),
		ConfigPath: getThemeConfigPath(),
	}

	// Load preset themes
	for name, theme := range PresetThemes {
		tm.AllThemes[name] = theme
	}

	// Load user custom themes
	tm.loadUserThemes()

	// Set default theme or use saved theme
	tm.CurrentTheme = tm.AllThemes["modern"]
	tm.loadCurrentTheme()

	return tm
}

// GetThemeNames returns all available theme names (sorted for stable switch order)
func (tm *ThemeManager) GetThemeNames() []string {
	names := make([]string, 0, len(tm.AllThemes))
	for name := range tm.AllThemes {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// SetTheme switches to the specified theme
func (tm *ThemeManager) SetTheme(name string) error {
	if theme, ok := tm.AllThemes[name]; ok {
		tm.CurrentTheme = theme
		tm.saveCurrentTheme()
		// Update global color variables
		tm.applyTheme()
		return nil
	}
	return fmt.Errorf("theme '%s' not found", name)
}

// NextTheme cycles to the next theme
func (tm *ThemeManager) NextTheme() string {
	names := tm.GetThemeNames()
	if len(names) == 0 {
		return ""
	}

	// Find current theme index
	currentIdx := 0
	for i, name := range names {
		if name == tm.CurrentTheme.Name {
			currentIdx = i
			break
		}
	}

	// Next theme
	nextIdx := (currentIdx + 1) % len(names)
	nextName := names[nextIdx]
	tm.SetTheme(nextName)
	return nextName
}

// Apply current theme to global variables
func (tm *ThemeManager) applyTheme() {
	t := &tm.CurrentTheme
	PrimaryColor = t.PrimaryColor()
	SecondaryColor = t.SecondaryColor()
	AccentColor = t.AccentColor()
	DangerColor = t.DangerColor()
	WarningColor = t.WarningColor()
	SuccessColor = t.SuccessColor()
	GrayColor = t.GrayColor()
	LightGrayColor = t.LightGrayColor()
	DimColor = t.DimColor()
	WhiteColor = t.ForegroundColor()
	BgSelected = t.SelectedBgColor()

	// Update styles
	TitleStyle = TitleStyle.Foreground(PrimaryColor)
	SubtitleStyle = SubtitleStyle.Foreground(LightGrayColor)
	HelpStyle = HelpStyle.Foreground(GrayColor)
	DimStyle = DimStyle.Foreground(DimColor)
	AccentStyle = AccentStyle.Foreground(AccentColor)
	WarningStyle = WarningStyle.Foreground(WarningColor)
	ErrorStyle = ErrorStyle.Foreground(DangerColor)
	SuccessStyle = SuccessStyle.Foreground(SuccessColor)
	InfoBoxStyle = InfoBoxStyle.BorderForeground(GrayColor)
	SelectedScanItemStyle = SelectedScanItemStyle.Background(BgSelected).Foreground(WhiteColor)

	// Update risk styles
	RiskLowStyle = RiskLowStyle.Foreground(SuccessColor)
	RiskMediumStyle = RiskMediumStyle.Foreground(WarningColor)
	RiskHighStyle = RiskHighStyle.Foreground(DangerColor)
}

// getThemeConfigPath returns the config file path
func getThemeConfigPath() string {
	home := scanner.GetRealHomeDir()
	if home == "" {
		return ""
	}
	return filepath.Join(home, ".config", "lume", "theme.json")
}

// saveCurrentTheme persists theme selection to disk
func (tm *ThemeManager) saveCurrentTheme() {
	if tm.ConfigPath == "" {
		return
	}

	data := map[string]string{
		"current_theme": tm.CurrentTheme.Name,
	}

	// Ensure directory exists
	dir := filepath.Dir(tm.ConfigPath)
	os.MkdirAll(dir, 0755)

	jsonData, _ := json.MarshalIndent(data, "", "  ")
	os.WriteFile(tm.ConfigPath, jsonData, 0644)
}

// loadCurrentTheme loads the saved theme setting
func (tm *ThemeManager) loadCurrentTheme() {
	if tm.ConfigPath == "" {
		return
	}

	data, err := os.ReadFile(tm.ConfigPath)
	if err != nil {
		return
	}

	var config map[string]string
	if err := json.Unmarshal(data, &config); err != nil {
		return
	}

	if themeName, ok := config["current_theme"]; ok {
		if _, exists := tm.AllThemes[themeName]; exists {
			tm.SetTheme(themeName)
		}
	}
}

// loadUserThemes loads custom themes from user config directory
func (tm *ThemeManager) loadUserThemes() {
	home := scanner.GetRealHomeDir()
	if home == "" {
		return
	}

	// Load custom themes from ~/.config/lume/themes/
	themesDir := filepath.Join(home, ".config", "lume", "themes")
	files, err := os.ReadDir(themesDir)
	if err != nil {
		return
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) != ".json" {
			continue
		}

		path := filepath.Join(themesDir, file.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}

		var theme Theme
		if err := json.Unmarshal(data, &theme); err != nil {
			continue
		}

		// Use filename as theme name
		name := file.Name()[:len(file.Name())-5] // strip .json
		if theme.Name == "" {
			theme.Name = name
		}
		tm.AllThemes[name] = theme
	}
}

// Global theme manager instance
var GlobalThemeManager *ThemeManager

// InitThemeManager initializes the global theme manager
func InitThemeManager() {
	GlobalThemeManager = NewThemeManager()
	GlobalThemeManager.applyTheme()
}
