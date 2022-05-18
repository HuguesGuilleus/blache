// BSD 3-Clause License in LICENSE file at the project root.
// All rights reserved.

package minecraftColor

// Ignored block when create BlocColorList.
var blocksIgnored = map[string]bool{
	"acacia_button":                 true,
	"acacia_wall_sign":              true,
	"activator_rail":                true,
	"air":                           true,
	"barrier":                       true,
	"birch_button":                  true,
	"birch_wall_sign":               true,
	"black_candle":                  true,
	"black_candle_cake":             true,
	"black_stained_glass":           true,
	"black_stained_glass_pane":      true,
	"black_wall_banner":             true,
	"blue_candle":                   true,
	"blue_candle_cake":              true,
	"blue_stained_glass":            true,
	"blue_stained_glass_pane":       true,
	"blue_wall_banner":              true,
	"brown_candle":                  true,
	"brown_candle_cake":             true,
	"brown_stained_glass":           true,
	"brown_stained_glass_pane":      true,
	"brown_wall_banner":             true,
	"cake":                          true,
	"candle":                        true,
	"candle_cake":                   true,
	"cave_air":                      true,
	"chain":                         true,
	"comparator":                    true,
	"creeper_head":                  true,
	"creeper_wall_head":             true,
	"crimson_button":                true,
	"crimson_sign":                  true,
	"crimson_wall_sign":             true,
	"cyan_candle":                   true,
	"cyan_candle_cake":              true,
	"cyan_stained_glass":            true,
	"cyan_stained_glass_pane":       true,
	"cyan_wall_banner":              true,
	"dark_oak_button":               true,
	"dark_oak_wall_sign":            true,
	"detector_rail":                 true,
	"dragon_head":                   true,
	"dragon_wall_head":              true,
	"end_rod":                       true,
	"flower_pot":                    true,
	"glass":                         true,
	"glass_pane":                    true,
	"gray_candle":                   true,
	"gray_candle_cake":              true,
	"gray_stained_glass":            true,
	"gray_stained_glass_pane":       true,
	"gray_wall_banner":              true,
	"green_candle":                  true,
	"green_candle_cake":             true,
	"green_stained_glass":           true,
	"green_stained_glass_pane":      true,
	"green_wall_banner":             true,
	"iron_bars":                     true,
	"jungle_button":                 true,
	"jungle_wall_sign":              true,
	"ladder":                        true,
	"lever":                         true,
	"light":                         true,
	"light_blue_candle":             true,
	"light_blue_candle_cake":        true,
	"light_blue_stained_glass":      true,
	"light_blue_stained_glass_pane": true,
	"light_blue_wall_banner":        true,
	"light_gray_candle":             true,
	"light_gray_candle_cake":        true,
	"light_gray_stained_glass":      true,
	"light_gray_stained_glass_pane": true,
	"light_gray_wall_banner":        true,
	"lime_candle":                   true,
	"lime_candle_cake":              true,
	"lime_stained_glass":            true,
	"lime_stained_glass_pane":       true,
	"lime_wall_banner":              true,
	"magenta_candle":                true,
	"magenta_candle_cake":           true,
	"magenta_stained_glass":         true,
	"magenta_stained_glass_pane":    true,
	"magenta_wall_banner":           true,
	"nether_portal":                 true,
	"oak_button":                    true,
	"oak_wall_sign":                 true,
	"orange_banner":                 true,
	"orange_candle":                 true,
	"orange_candle_cake":            true,
	"orange_stained_glass":          true,
	"orange_stained_glass_pane":     true,
	"orange_wall_banner":            true,
	"pink_candle":                   true,
	"pink_candle_cake":              true,
	"pink_stained_glass":            true,
	"pink_stained_glass_pane":       true,
	"pink_wall_banner":              true,
	"player_head":                   true,
	"player_wall_head":              true,
	"polished_blackstone_button":    true,
	"potted_acacia_sapling":         true,
	"potted_allium":                 true,
	"potted_azalea_bush":            true,
	"potted_azure_bluet":            true,
	"potted_bamboo":                 true,
	"potted_birch_sapling":          true,
	"potted_blue_orchid":            true,
	"potted_brown_mushroom":         true,
	"potted_cactus":                 true,
	"potted_cornflower":             true,
	"potted_crimson_fungus":         true,
	"potted_crimson_roots":          true,
	"potted_dandelion":              true,
	"potted_dark_oak_sapling":       true,
	"potted_dead_bush":              true,
	"potted_fern":                   true,
	"potted_flowering_azalea_bush":  true,
	"potted_jungle_sapling":         true,
	"potted_lily_of_the_valley":     true,
	"potted_oak_sapling":            true,
	"potted_orange_tulip":           true,
	"potted_oxeye_daisy":            true,
	"potted_pink_tulip":             true,
	"potted_poppy":                  true,
	"potted_red_mushroom":           true,
	"potted_red_tulip":              true,
	"potted_spruce_sapling":         true,
	"potted_warped_fungus":          true,
	"potted_white_tulip":            true,
	"potted_wither_rose":            true,
	"powered_rail":                  true,
	"purple_candle":                 true,
	"purple_candle_cake":            true,
	"purple_stained_glass":          true,
	"purple_stained_glass_pane":     true,
	"purple_wall_banner":            true,
	"rail":                          true,
	"red_candle":                    true,
	"red_candle_cake":               true,
	"red_stained_glass":             true,
	"red_stained_glass_pane":        true,
	"red_wall_banner":               true,
	"redstone_lamp":                 true,
	"redstone_torch":                true,
	"redstone_wall_torch":           true,
	"redstone_wire":                 true,
	"repeater":                      true,
	"skeleton_skull":                true,
	"skeleton_wall_skull":           true,
	"soul_torch":                    true,
	"soul_wall_torch":               true,
	"spruce_button":                 true,
	"spruce_wall_sign":              true,
	"stone_button":                  true,
	"structure_void":                true,
	"tinted_glass":                  true,
	"torch":                         true,
	"tripwire":                      true,
	"tripwire_hook":                 true,
	"void_air":                      true,
	"wall_torch":                    true,
	"warped_button":                 true,
	"warped_wall_sign":              true,
	"white_candle":                  true,
	"white_candle_cake":             true,
	"white_stained_glass":           true,
	"white_stained_glass_pane":      true,
	"white_wall_banner":             true,
	"wither_skeleton_skull":         true,
	"wither_skeleton_wall_skull":    true,
	"yellow_banner":                 true,
	"yellow_candle":                 true,
	"yellow_candle_cake":            true,
	"yellow_stained_glass":          true,
	"yellow_stained_glass_pane":     true,
	"yellow_wall_banner":            true,
	"zombie_head":                   true,
	"zombie_wall_head":              true,
}
