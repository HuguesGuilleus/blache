package minecraftColor

import (
	"image/color"
	"strings"
)

// var Bloc = map[string]color.RGBA{}

// Get the color of a bloc with its name space like 'minecraft:stone'.
func GetBloc(n string) color.RGBA {
	if strings.HasPrefix(n, "minecraft:") {
		n = n[len("minecraft:"):]
	}

	if c, exist := Bloc[n]; exist {
		return c
	}
	return color.RGBA{}
}

// Add blocs with theys color to the variable Bloc.
func blocAddColor(c color.RGBA, keys []string) int {
	for _, k := range keys {
		Bloc[k] = c
	}
	return 0
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
	"birch_door":              {247, 233, 163, 255},
	"birch_fence_gate":        {247, 233, 163, 255},
	"birch_fence":             {247, 233, 163, 255},
	"birch_log":               {247, 233, 163, 255},
	"birch_planks":            {247, 233, 163, 255},
	"birch_pressure_plate":    {247, 233, 163, 255},
	"birch_sign":              {247, 233, 163, 255},
	"birch_slab":              {247, 233, 163, 255},
	"birch_stairs":            {247, 233, 163, 255},
	"birch_trapdoor":          {247, 233, 163, 255},
	"birch_wood":              {247, 233, 163, 255},
	"bone_block":              {247, 233, 163, 255},
	"chiseled_sandstone":      {247, 233, 163, 255},
	"cut_sandstone_slab":      {247, 233, 163, 255},
	"cut_sandstone":           {247, 233, 163, 255},
	"end_stone_brick_slab":    {247, 233, 163, 255},
	"end_stone_brick_stairs":  {247, 233, 163, 255},
	"end_stone_brick_wall":    {247, 233, 163, 255},
	"end_stone_bricks":        {247, 233, 163, 255},
	"end_stone":               {247, 233, 163, 255},
	"glowstone":               {247, 233, 163, 255},
	"sand":                    {247, 233, 163, 255},
	"sandstone_slab":          {247, 233, 163, 255},
	"sandstone_stairs":        {247, 233, 163, 255},
	"sandstone_wall":          {247, 233, 163, 255},
	"sandstone":               {247, 233, 163, 255},
	"scaffolding":             {247, 233, 163, 255},
	"smooth_sandstone_slab":   {247, 233, 163, 255},
	"smooth_sandstone_stairs": {247, 233, 163, 255},
	"smooth_sandstone":        {247, 233, 163, 255},
	"stripped_birch_log":      {247, 233, 163, 255},
	"stripped_birch_wood":     {247, 233, 163, 255},
	"turtle_egg":              {247, 233, 163, 255},
	// 3 WOOL
	"cobweb":        {199, 199, 199, 255},
	"mushroom_stem": {199, 199, 199, 255},
	"bed":           {199, 199, 199, 255},
	// 4 FIRE
	"fire":           {255, 0, 0, 255},
	"flowing_lava":   {255, 0, 0, 255},
	"lava":           {255, 0, 0, 255},
	"redstone_block": {255, 0, 0, 255},
	"soul_fire":      {255, 0, 0, 255},
	"tnt":            {255, 0, 0, 255},
	// 5 ICE
	"blue_ice":    {160, 160, 255, 255},
	"frosted_ice": {160, 160, 255, 255},
	"ice":         {160, 160, 255, 255},
	"packed_ice":  {160, 160, 255, 255},
	// 6 METAL
	"anvil":                         {160, 160, 255, 255},
	"brewing_stand":                 {160, 160, 255, 255},
	"chipped_anvil":                 {160, 160, 255, 255},
	"damaged_anvil":                 {160, 160, 255, 255},
	"grindstone":                    {160, 160, 255, 255},
	"heavy_weighted_pressure_plate": {160, 160, 255, 255},
	"iron_block":                    {160, 160, 255, 255},
	"iron_door":                     {160, 160, 255, 255},
	"iron_trapdoor":                 {160, 160, 255, 255},
	"lantern":                       {160, 160, 255, 255},
	"soul_lantern":                  {160, 160, 255, 255},
	// 7 PLANT
	"acacia_leaves":         {0, 124, 0, 255},
	"acacia_sapling":        {0, 124, 0, 255},
	"allium":                {0, 124, 0, 255},
	"attached_melon_stem":   {0, 124, 0, 255},
	"attached_pumpkin_stem": {0, 124, 0, 255},
	"azure_bluet":           {0, 124, 0, 255},
	"bamboo":                {0, 124, 0, 255},
	"beetroots":             {0, 124, 0, 255},
	"birch_leaves":          {0, 124, 0, 255},
	"birch_sapling":         {0, 124, 0, 255},
	"blue_orchid":           {0, 124, 0, 255},
	"brown_mushroom":        {0, 124, 0, 255},
	"cactus":                {0, 124, 0, 255},
	"carrots":               {0, 124, 0, 255},
	"cocoa":                 {0, 124, 0, 255},
	"cornflower":            {0, 124, 0, 255},
	"crimson_fungus":        {0, 124, 0, 255}, // minecraft 1.16
	"dandelion":             {0, 124, 0, 255},
	"dark_oak_leaves":       {0, 124, 0, 255},
	"dark_oak_sapling":      {0, 124, 0, 255},
	"fern":                  {0, 124, 0, 255},
	"grass":                 {0, 124, 0, 255},
	"jungle_leaves":         {0, 124, 0, 255},
	"jungle_sapling":        {0, 124, 0, 255},
	"large_fern":            {0, 124, 0, 255},
	"lilac":                 {0, 124, 0, 255},
	"lily_of_the_valley":    {0, 124, 0, 255},
	"lily_pad":              {0, 124, 0, 255},
	"melon_seeds":           {0, 124, 0, 255},
	"melon_stem":            {0, 124, 0, 255},
	"oak_leaves":            {0, 124, 0, 255},
	"oak_sapling":           {0, 124, 0, 255},
	"orange_tulip":          {0, 124, 0, 255},
	"oxeye_daisy":           {0, 124, 0, 255},
	"peony":                 {0, 124, 0, 255},
	"pink_tulip":            {0, 124, 0, 255},
	"poppy":                 {0, 124, 0, 255},
	"potatoes":              {0, 124, 0, 255},
	"pumpkin_seeds":         {0, 124, 0, 255},
	"pumpkin_stem":          {0, 124, 0, 255},
	"red_mushroom":          {0, 124, 0, 255},
	"red_tulip":             {0, 124, 0, 255},
	"rose_bush":             {0, 124, 0, 255},
	"spruce_leaves":         {0, 124, 0, 255},
	"spruce_sapling":        {0, 124, 0, 255},
	"sugar_cane":            {0, 124, 0, 255},
	"sunflower":             {0, 124, 0, 255},
	"sweet_berry_bush":      {0, 124, 0, 255},
	"tall_grass":            {0, 124, 0, 255},
	"vine":                  {0, 124, 0, 255},
	"warped_fungus":         {0, 124, 0, 255}, // minecraft 1.16
	"weeping_vines_plant":   {0, 124, 0, 255}, // minecraft 1.16
	"weeping_vines":         {0, 124, 0, 255}, // minecraft 1.16
	"wheat_seeds":           {0, 124, 0, 255},
	"wheat":                 {0, 124, 0, 255},
	"white_tulip":           {0, 124, 0, 255},
	"wither_rose":           {0, 124, 0, 255},
	// 8 SNOW
	"snow_block":               {255, 255, 255, 255},
	"snow":                     {255, 255, 255, 255},
	"white_banner":             {255, 255, 255, 255},
	"white_bed":                {255, 255, 255, 255},
	"white_carpet":             {255, 255, 255, 255},
	"white_concrete_powder":    {255, 255, 255, 255},
	"white_concrete":           {255, 255, 255, 255},
	"white_glazed_terracotta":  {255, 255, 255, 255},
	"white_shulker_box":        {255, 255, 255, 255},
	"white_stained_glass_pane": {255, 255, 255, 255},
	"white_stained_glass":      {255, 255, 255, 255},
	"white_wall_banner":        {255, 255, 255, 255},
	"white_wool":               {255, 255, 255, 255},
	// 9 CLAY
	"clay":                           {164, 168, 184, 255},
	"infested_chiseled_stone_bricks": {164, 168, 184, 255},
	"infested_cobblestone":           {164, 168, 184, 255},
	"infested_cracked_stone_bricks":  {164, 168, 184, 255},
	"infested_mossy_stone_bricks":    {164, 168, 184, 255},
	"infested_stone_bricks":          {164, 168, 184, 255},
	"infested_stone":                 {164, 168, 184, 255},
	// 10 DIRT
	"brown_mushroom_block":    {151, 109, 77, 255},
	"coarse_dirt":             {151, 109, 77, 255},
	"dirt":                    {151, 109, 77, 255},
	"farmland":                {151, 109, 77, 255},
	"granite_slab":            {151, 109, 77, 255},
	"granite_stairs":          {151, 109, 77, 255},
	"granite_wall":            {151, 109, 77, 255},
	"granite":                 {151, 109, 77, 255},
	"grass_path":              {151, 109, 77, 255},
	"jukebox":                 {151, 109, 77, 255},
	"jungle_door":             {151, 109, 77, 255},
	"jungle_fence_gate":       {151, 109, 77, 255},
	"jungle_fence":            {151, 109, 77, 255},
	"jungle_log":              {151, 109, 77, 255},
	"jungle_planks":           {151, 109, 77, 255},
	"jungle_pressure_plate":   {151, 109, 77, 255},
	"jungle_sign":             {151, 109, 77, 255},
	"jungle_slab":             {151, 109, 77, 255},
	"jungle_stairs":           {151, 109, 77, 255},
	"jungle_trapdoor":         {151, 109, 77, 255},
	"jungle_wood":             {151, 109, 77, 255},
	"polished_granite_slab":   {151, 109, 77, 255},
	"polished_granite_stairs": {151, 109, 77, 255},
	"polished_granite":        {151, 109, 77, 255},
	"stripped_jungle_log":     {151, 109, 77, 255},
	"stripped_jungle_wood":    {151, 109, 77, 255},
	// 11 STONE
	"stone": {112, 112, 112, 255},
	// ...

	// 12 WATER
	"bubble_column": {64, 64, 255, 255},
	"flowing_water": {64, 64, 255, 255},
	"kelp_plant":    {64, 64, 255, 255},
	"kelp":          {64, 64, 255, 255},
	"seagrass":      {64, 64, 255, 255},
	"tall_seagrass": {64, 64, 255, 255},
	"water":         {64, 64, 255, 255},
	// 13 WOOD
	"bamboo_sapling":     {143, 119, 72, 255},
	"barrel":             {143, 119, 72, 255},
	"bee_nest":           {143, 119, 72, 255},
	"beehive":            {143, 119, 72, 255},
	"bookshelf":          {143, 119, 72, 255},
	"cartography_table":  {143, 119, 72, 255},
	"chest":              {143, 119, 72, 255},
	"composter":          {143, 119, 72, 255},
	"crafting_table":     {143, 119, 72, 255},
	"daylight_detector":  {143, 119, 72, 255},
	"dead_bush":          {143, 119, 72, 255},
	"fletching_table":    {143, 119, 72, 255},
	"lectern":            {143, 119, 72, 255},
	"loom":               {143, 119, 72, 255},
	"note_block":         {143, 119, 72, 255},
	"oak_door":           {143, 119, 72, 255},
	"oak_fence":          {143, 119, 72, 255},
	"oak_fence_gate":     {143, 119, 72, 255},
	"oak_log":            {143, 119, 72, 255},
	"oak_planks":         {143, 119, 72, 255},
	"oak_pressure_plate": {143, 119, 72, 255},
	"oak_sign":           {143, 119, 72, 255},
	"oak_slab":           {143, 119, 72, 255},
	"oak_stairs":         {143, 119, 72, 255},
	"oak_trapdoor":       {143, 119, 72, 255},
	"oak_wood":           {143, 119, 72, 255},
	"petrified_oak_slab": {143, 119, 72, 255},
	"smithing_table":     {143, 119, 72, 255},
	"stripped_oak_log":   {143, 119, 72, 255},
	"stripped_oak_wood":  {143, 119, 72, 255},
	"trapped_chest":      {143, 119, 72, 255},
	// 14 QUARTZ
	"diorite":                 {255, 252, 245, 255},
	"diorite_slab":            {255, 252, 245, 255},
	"diorite_stairs":          {255, 252, 245, 255},
	"diorite_wall":            {255, 252, 245, 255},
	"polished_diorite":        {255, 252, 245, 255},
	"polished_diorite_slab":   {255, 252, 245, 255},
	"polished_diorite_stairs": {255, 252, 245, 255},
	"chiseled_quartz_block":   {255, 252, 245, 255},
	"quartz_block":            {255, 252, 245, 255},
	"quartz_pillar":           {255, 252, 245, 255},
	"quartz_slab":             {255, 252, 245, 255},
	"quartz_stairs":           {255, 252, 245, 255},
	"smooth_quartz":           {255, 252, 245, 255},
	"smooth_quartz_slab":      {255, 252, 245, 255},
	"smooth_quartz_stairs":    {255, 252, 245, 255},
	"sea_lantern":             {255, 252, 245, 255},
	"target":                  {255, 252, 245, 255},
	// 15 COLOR_ORANGE
	"chiseled_red_sandstone": {216, 127, 51, 255},
	"cut_red_sandstone":      {216, 127, 51, 255},
	"red_sand":               {216, 127, 51, 255},
	"red_sandstone":          {216, 127, 51, 255},
	"smooth_red_sandstone":   {216, 127, 51, 255},

	// ...

	// 29 COLOR_BLACK
	"ancient_debris":           {25, 25, 25, 255},
	"basalt":                   {25, 25, 25, 255},
	"black_banner":             {25, 25, 25, 255},
	"black_bed":                {25, 25, 25, 255},
	"black_carpet":             {25, 25, 25, 255},
	"black_concrete":           {25, 25, 25, 255},
	"black_concrete_powder":    {25, 25, 25, 255},
	"black_glazed_terracotta":  {25, 25, 25, 255},
	"black_shulker_box":        {25, 25, 25, 255},
	"black_stained_glass":      {25, 25, 25, 255},
	"black_stained_glass_pane": {25, 25, 25, 255},
	"black_wall_banner":        {25, 25, 25, 255},
	"black_wool":               {25, 25, 25, 255},
	"coal_block":               {25, 25, 25, 255},
	"crying_obsidian":          {25, 25, 25, 255},
	"dragon_egg":               {25, 25, 25, 255},
	"end_gateway":              {25, 25, 25, 255},
	"end_portal":               {25, 25, 25, 255},
	"netherite_block":          {25, 25, 25, 255},
	"obsidian":                 {25, 25, 25, 255},
	"polished_basalt":          {25, 25, 25, 255}, // minecraft 1.16 ???
	// 30 GOLD
	"bell":                          {250, 238, 77, 255},
	"gold_block":                    {250, 238, 77, 255},
	"light_weighted_pressure_plate": {250, 238, 77, 255},
	// 31 DIAMOND
	"beacon":                  {92, 219, 213, 255},
	"conduit":                 {92, 219, 213, 255},
	"dark_prismarine_slab":    {92, 219, 213, 255},
	"dark_prismarine_stairs":  {92, 219, 213, 255},
	"dark_prismarine":         {92, 219, 213, 255},
	"diamond_block":           {92, 219, 213, 255},
	"prismarine_brick_slab":   {92, 219, 213, 255},
	"prismarine_brick_stairs": {92, 219, 213, 255},
	"prismarine_bricks":       {92, 219, 213, 255},
	// 32 LAPIS
	"lapis_block": {74, 128, 255, 255},
	// 33 EMERALD
	"emerald_block": {0, 217, 58, 255},
	// 34 PODZOL
	"campfire":              {129, 86, 49, 255},
	"podzol":                {129, 86, 49, 255},
	"spruce_door":           {129, 86, 49, 255},
	"spruce_fence_gate":     {129, 86, 49, 255},
	"spruce_fence":          {129, 86, 49, 255},
	"spruce_log":            {129, 86, 49, 255},
	"spruce_planks":         {129, 86, 49, 255},
	"spruce_pressure_plate": {129, 86, 49, 255},
	"spruce_sign":           {129, 86, 49, 255},
	"spruce_slab":           {129, 86, 49, 255},
	"spruce_stairs":         {129, 86, 49, 255},
	"spruce_trapdoor":       {129, 86, 49, 255},
	"spruce_wood":           {129, 86, 49, 255},
	"stripped_spruce_log":   {129, 86, 49, 255},
	"stripped_spruce_wood":  {129, 86, 49, 255},
	// 35 NETHER
	"magma_block":             {112, 2, 0, 255},
	"nether_brick_fence":      {112, 2, 0, 255},
	"nether_brick_slab":       {112, 2, 0, 255},
	"nether_brick_stairs":     {112, 2, 0, 255},
	"nether_brick_wall":       {112, 2, 0, 255},
	"nether_bricks":           {112, 2, 0, 255},
	"nether_quartz_ore":       {112, 2, 0, 255},
	"netherrack":              {112, 2, 0, 255},
	"red_nether_brick_slab":   {112, 2, 0, 255},
	"red_nether_brick_stairs": {112, 2, 0, 255},
	"red_nether_brick_wall":   {112, 2, 0, 255},
	"red_nether_bricks":       {112, 2, 0, 255},
	// minecraft 1.16: Crimson (Roots, Planks, Slab, Pressure Plate, Fence, Trapdoor, Fence Gate, Stairs, Door, Sign)
	// 36 TERRACOTTA_WHITE
	"white_terracotta": {209, 177, 161, 255},
	// 37 TERRACOTTA_ORANGE
	"orange_terracotta": {159, 82, 36, 255},
	// 38 TERRACOTTA_MAGENTA
	"magenta_terracotta": {149, 87, 108, 255},
	// 39 TERRACOTTA_LIGHT_BLUE
	"light_blue_terracotta": {112, 108, 138, 255},
	// 40 TERRACOTTA_YELLOW
	"yellow_terracotta": {186, 133, 36, 255},
	// 41 TERRACOTTA_LIGHT_GREEN
	"lime_terracotta": {103, 117, 53, 255},
	// 42 TERRACOTTA_PINK
	"pink_terracotta": {160, 77, 78, 255},
	// 43 TERRACOTTA_GRAY
	"gray_terracotta": {57, 41, 35, 255},
	// 44 TERRACOTTA_LIGHT_GRAY
	"light_gray_terracotta": {135, 107, 98, 255},
	// 45 TERRACOTTA_CYAN
	"cyan_terracotta": {87, 92, 92, 255},
	// 46 TERRACOTTA_PURPLE
	"purple_terracotta":  {122, 73, 88, 255},
	"purple_shulker_box": {122, 73, 88, 255},
	// 47 TERRACOTTA_BLUE
	"blue_terracotta": {76, 62, 92, 255},
	// 48 TERRACOTTA_BROWN
	"brown_terracotta": {76, 50, 35, 255},
	// 49 TERRACOTTA_GREEN
	"green_terracotta": {76, 82, 42, 255},
	// 50 TERRACOTTA_RED
	"red_terracotta": {142, 60, 46, 255},
	// 51 TERRACOTTA_BLACK
	"black_terracotta": {37, 22, 16, 255},
}
