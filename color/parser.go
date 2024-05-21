package color

import (
	"strings"
)

type Style struct {
	Property string
	Value    string
}

var styleMap = map[string]Style{
	"text-blue-50": {"color", "#eff6ff"}, "text-blue-100": {"color", "#dbeafe"}, "text-blue-200": {"color", "#bfdbfe"}, "text-blue-300": {"color", "#93c5fd"}, "text-blue-400": {"color", "#60a5fa"}, "text-blue-500": {"color", "#3b82f6"}, "text-blue-600": {"color", "#2563eb"}, "text-blue-700": {"color", "#1d4ed8"}, "text-blue-800": {"color", "#1e40af"}, "text-blue-900": {"color", "#1e3a8a"}, "text-blue-950": {"color", "#1e2a5b"},
	"text-red-50": {"color", "#fef2f2"}, "text-red-100": {"color", "#fee2e2"}, "text-red-200": {"color", "#fecaca"}, "text-red-300": {"color", "#fca5a5"}, "text-red-400": {"color", "#f87171"}, "text-red-500": {"color", "#ef4444"}, "text-red-600": {"color", "#dc2626"}, "text-red-700": {"color", "#b91c1c"}, "text-red-800": {"color", "#991b1b"}, "text-red-900": {"color", "#7f1d1d"},
    "text-green-50": {"color", "#ecfdf5"}, "text-green-100": {"color", "#d1fae5"}, "text-green-200": {"color", "#a7f3d0"}, "text-green-300": {"color", "#6ee7b7"}, "text-green-400": {"color", "#34d399"}, "text-green-500": {"color", "#10b981"}, "text-green-600": {"color", "#059669"}, "text-green-700": {"color", "#047857"}, "text-green-800": {"color", "#065f46"}, "text-green-900": {"color", "#064e3b"},
	"text-yellow-50": {"color", "#fef9c3"}, "text-yellow-100": {"color", "#fef08a"}, "text-yellow-200": {"color", "#fde047"}, "text-yellow-300": {"color", "#facc15"}, "text-yellow-400": {"color", "#eab308"}, "text-yellow-500": {"color", "#ca8a04"}, "text-yellow-600": {"color", "#a16207"}, "text-yellow-700": {"color", "#854d0e"}, "text-yellow-800": {"color", "#713f12"}, "text-yellow-900": {"color", "#5b3412"},
    "text-cyan-50": {"color", "#ecfeff"}, "text-cyan-100": {"color", "#cffafe"}, "text-cyan-200": {"color", "#a5f3fc"}, "text-cyan-300": {"color", "#67e8f9"}, "text-cyan-400": {"color", "#22d3ee"}, "text-cyan-500": {"color", "#06b6d4"}, "text-cyan-600": {"color", "#0891b2"}, "text-cyan-700": {"color", "#0e7490"}, "text-cyan-800": {"color", "#155e75"}, "text-cyan-900": {"color", "#164e63"},
    "text-purple-50": {"color", "#f5f3ff"}, "text-purple-100": {"color", "#ede9fe"}, "text-purple-200": {"color", "#ddd6fe"}, "text-purple-300": {"color", "#c4b5fd"}, "text-purple-400": {"color", "#a78bfa"}, "text-purple-500": {"color", "#8b5cf6"}, "text-purple-600": {"color", "#7c3aed"}, "text-purple-700": {"color", "#6d28d9"}, "text-purple-800": {"color", "#5b21b6"}, "text-purple-900": {"color", "#4c1d95"},
    "text-black-50": {"color", "#f9fafb"}, "text-black-100": {"color", "#f3f4f6"}, "text-black-200": {"color", "#e5e7eb"}, "text-black-300": {"color", "#d1d5db"}, "text-black-400": {"color", "#9ca3af"}, "text-black-500": {"color", "#6b7280"}, "text-black-600": {"color", "#4b5563"}, "text-black-700": {"color", "#374151"}, "text-black-800": {"color", "#1f2937"}, "text-black-900": {"color", "#111827"},
    "text-white-50": {"color", "#fdfdfd"}, "text-white-100": {"color", "#fbfbfb"}, "text-white-200": {"color", "#f8f8f8"}, "text-white-300": {"color", "#f5f5f5"}, "text-white-400": {"color", "#f2f2f2"}, "text-white-500": {"color", "#efefef"}, "text-white-600": {"color", "#ececec"}, "text-white-700": {"color", "#e9e9e9"}, "text-white-800": {"color", "#e6e6e6"}, "text-white-900": {"color", "#e3e3e3"},
}

func ParseClass(class string) Style {
	if style, ok := styleMap[class]; ok {
		return style
	}
	return Style{}
}

func ParseClasses(classes string) map[string]string {
	styles := make(map[string]string)
	for _, class := range strings.Fields(classes) {
		style := ParseClass(class)
		if style.Property != "" {
			styles[style.Property] = style.Value
		}
	}
	return styles
}

