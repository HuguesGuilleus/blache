package biome

import (
	"image/color"
)

// Convert hex color to RGBA color.
func toRGBA(i int) color.RGBA {
	return color.RGBA{
		R: uint8(i / 0x10000),
		G: uint8(i / 0x100),
		B: uint8(i),
		A: 0xFF,
	}
}

var (
	BlackRGBA     = toRGBA(0)
	// https://minecraft-el.gamepedia.com/Biome/ID
	Color = [256]color.RGBA{
		toRGBA(0x000070), // Ocean
		toRGBA(0x8DB360), // Plains
		toRGBA(0xFA9418), // Desert
		toRGBA(0x606060), // Extreme Hills
		toRGBA(0x056621), // Forest
		toRGBA(0x0B6659), // Taiga
		toRGBA(0x07F9B2), // Swampland
		toRGBA(0x0000FF), // River
		toRGBA(0xFF0000), // Hell
		toRGBA(0x8080FF), // The End (Sky)
		toRGBA(0x9090A0), // FrozenOcean
		toRGBA(0xA0A0FF), // FrozenRiver
		toRGBA(0xFFFFFF), // Ice Plains
		toRGBA(0xA0A0A0), // Ice Mountains
		toRGBA(0xFF00FF), // MushroomIsland
		toRGBA(0xA000FF), // MushroomIslandShore
		toRGBA(0xFADE55), // Beach
		toRGBA(0xD25F12), // DesertHills
		toRGBA(0x22551C), // ForestHills
		toRGBA(0x163933), // TaigaHills
		toRGBA(0x72789A), // Extreme Hills Edge
		toRGBA(0x537B09), // Jungle
		toRGBA(0x2C4205), // JungleHills
		toRGBA(0x628B17), // JungleEdge
		toRGBA(0x000030), // Deep Ocean
		toRGBA(0xA2A284), // Stone Beach
		toRGBA(0xFAF0C0), // Cold Beach
		toRGBA(0x307444), // Birch Forest
		toRGBA(0x1F5F32), // Birch Forest Hills
		toRGBA(0x40511A), // Roofed Forest
		toRGBA(0x31554A), // Cold Taiga
		toRGBA(0x243F36), // Cold Taiga Hills
		toRGBA(0x596651), // Mega Taiga
		toRGBA(0x545F3E), // Mega Taiga Hills
		toRGBA(0x507050), // Extreme Hills+
		toRGBA(0xBDB25F), // Savanna
		toRGBA(0xA79D64), // Savanna Plateau
		toRGBA(0xD94515), // Mesa
		toRGBA(0xB09765), // Mesa Plateau F
		toRGBA(0xCA8C65), // Mesa Plateau
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,        // void
		BlackRGBA,        // Plains M
		toRGBA(0xB5DB88), // Sunflower Plains
		toRGBA(0xFFBC40), // Desert M
		toRGBA(0x888888), // Extreme Hills M
		toRGBA(0x6A7425), // Flower Forest
		toRGBA(0x596651), // Taiga M
		toRGBA(0x2FFFDA), // Swampland M
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		toRGBA(0xB4DCDC), // Ice Plains Spikes
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		toRGBA(0x7BA331), // Jungle M
		BlackRGBA,
		toRGBA(0x8AB33F), // JungleEdge M
		BlackRGBA,
		BlackRGBA,
		BlackRGBA,
		toRGBA(0x589C6C), // Birch Forest M
		toRGBA(0x47875A), // Birch Forest Hills M
		toRGBA(0x687942), // Roofed Forest M
		toRGBA(0x597D72), // Cold Taiga M
		BlackRGBA,
		toRGBA(0x6B5F4C), // Mega Spruce Taiga
		toRGBA(0x6D7766), // Redwood Taiga Hills M
		toRGBA(0x789878), // Extreme Hills+ M
		toRGBA(0xE5DA87), // Savanna M
		toRGBA(0xCFC58C), // Savanna Plateau M
		toRGBA(0xFF6D3D), // Mesa (Bryce)
		toRGBA(0xD8BF8D), // Mesa Plateau F M
		toRGBA(0xF2B48D), // Mesa Plateau M
	}
)
