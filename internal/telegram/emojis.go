package telegram

import "math/rand"

var emojis = []string{
	"🤔", "👹", "👺", "🤡", "💩",
	"👾", "🤖", "😽", "👌🏻", "🕸",
	"🤌🏻", "🤙🏻", "👁", "👀", "🦷",
	"🥷🏻", "💃", "🐥", "🐈", "🦦",
	"🐉", "🀄️", "🐲", "🌝", "🌚",
	"🥕", "🥨", "🥩", "📡", "🚬",
	"🪤", "🧱", "💔", "💭", "👉🏻👈🏻"}

func emoji() string {
	return emojis[rand.Int()%len(emojis)]
}
