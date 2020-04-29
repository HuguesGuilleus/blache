package minecraftColor

import (
	"image/color"
	"regexp"
	"strings"
)

var (
	getEnd = regexp.MustCompile(`\w+_$`)
)

// Get the color of a bloc with its name space like 'minecraft:stone'.
func GetBloc(n string) color.RGBA {
	if strings.HasPrefix(n, "minecraft:") {
		n = n[len("minecraft:"):]
	}

	if c, exist := Bloc[n]; exist {
		return c
	}
	if c, exist := Bloc[getEnd.ReplaceAllString(n, "")]; exist {
		return c
	}
	return color.RGBA{}
}

// The bloc like "oak_log" are stored in "log" key.
// Source:
// 	https://minecraft.gamepedia.com/Map_item_format#Map_colors
// 	https://minecraft.gamepedia.com/index.php?title=Java_Edition_data_values/Blocks
var Bloc = map[string]color.RGBA{
	// 1 GRASS
	"grass_block": {127, 178, 56, 255},
	"slime_block": {127, 178, 56, 255},
	// 2 SAND
	"sand":                   {247, 233, 163, 255},
	"planks":                 {247, 233, 163, 255},
	"log":                    {247, 233, 163, 255},
	"stem":                   {247, 233, 163, 255},
	"sign":                   {247, 233, 163, 255},
	"wood":                   {247, 233, 163, 255},
	"hyphae":                 {247, 233, 163, 255},
	"trapdoor":               {247, 233, 163, 255},
	"glowstone":              {247, 233, 163, 255},
	"fence":                  {247, 233, 163, 255},
	"sandstone":              {247, 233, 163, 255},
	"cut_sandstone":          {247, 233, 163, 255},
	"chiseled_sandstone":     {247, 233, 163, 255},
	"oak_fence_gate":         {247, 233, 163, 255},
	"spruce_fence_gate":      {247, 233, 163, 255},
	"birch_fence_gate":       {247, 233, 163, 255},
	"jungle_fence_gate":      {247, 233, 163, 255},
	"acacia_fence_gate":      {247, 233, 163, 255},
	"dark_oak_fence_gate":    {247, 233, 163, 255},
	"crimson_fence_gate":     {247, 233, 163, 255},
	"warped_fence_gate":      {247, 233, 163, 255},
	"scaffolding":            {247, 233, 163, 255},
	"bone_block":             {247, 233, 163, 255},
	"turtle_egg":             {247, 233, 163, 255},
	"end_stone":              {247, 233, 163, 255},
	"end_stone_brick_slab":   {247, 233, 163, 255},
	"end_stone_brick_stairs": {247, 233, 163, 255},
	"end_stone_brick_wall":   {247, 233, 163, 255},
	"end_stone_bricks":       {247, 233, 163, 255},
	// 3
	// "":                   {,,, 255},

	// 8 SNOW
	"snow":       {255, 255, 255, 255},
	"snow_block": {255, 255, 255, 255},

	// 10 DIRT
	"dirt": {151, 109, 77, 255},
	// 11 STONE
	"stone": {112, 112, 112, 255},

	// 12 WATER
	"kelp":          {64, 64, 255, 255},
	"kelp_plant":    {64, 64, 255, 255},
	"seagrass":      {64, 64, 255, 255},
	"tall_seagrass": {64, 64, 255, 255},
	"water":         {64, 64, 255, 255},
	"flowing_water": {64, 64, 255, 255},
	"bubble_column": {64, 64, 255, 255},
	// 13 WOOD
	// 15 COLOR_ORANGE
	"red_sand":               {216, 127, 51, 255},
	"red_sandstone":          {216, 127, 51, 255},
	"cut_red_sandstone":      {216, 127, 51, 255},
	"chiseled_red_sandstone": {216, 127, 51, 255},
	"smooth_red_sandstone":   {216, 127, 51, 255},
}

// Dépend du matériaux
// https://minecraft.gamepedia.com/Slab#ID
// https://minecraft.gamepedia.com/Pressure_Plate#ID
// https://minecraft.gamepedia.com/Stairs#ID
// Trapdoor
// Door
