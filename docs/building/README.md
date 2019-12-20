# Immortal Documentation

Documentation for building and running the mud.

## Setting up New Immortals

In the player file (.plr) you must have access level > 0:

Levels: 

To edit find the following line in the .plr file.
Access Level:

## Creating instances

Use the commands to bring things to life in the MUD based on an existing vnum (unique number for a template of an object of some type).

### Creating a City

```
claim
```


```
city found
```

### Create Instances of Buildings

Useful Help: BUILD

```
build <buildingtype>
```

You will need the resources to build the building.

### Creating Instances of Objects

Useful Help: CREATE

```
create <thing>
```

## Creating New Things

### Change The World Terrain

To change the current tile you are on you can type:
```
.map terrain <terraintype>
```
Available Terrain Types:

```
Plains, Light Forest, Forest, Shady Forest, Overgrown Forest, River, Ocean, Crop,
Mountain, Road, Building, Inside, Desert Crop, Seeded Field, Sandy Field, Jungle Field, Jungle Crop, 
Trench, Tower of Souls, Desert, Oasis, Sandy Trench, Grove, Partial Jungle, Jungle, Swamp, Tundra, 
Adventure, Lake, Irrigated Plains, Irrigated Forest, Irrigated Jungle, Enchanted Saplings, Clear-Cut 
Enchanted Forest, Enchanted Copse, Enchanted Woods, Enchanted Forest, Dead Forest, Flowing Lava, 
Cooling Lava, Volcanic Mountain, Scorched Woods, Scorched Grove, Scorched Plains, Scorched Crop, 
Scorched Desert Crop, Scorched Desert, Frozen River, Thin Evergreen Forest, Evergreen Forest, 
Overgrown Evergreen Forest, Thick Evergreen Forest, Beaver Dam, Flooded Plains, Flooded Woods, 
Flooded Desert, bugs, wheat, corn, hops, cotton, barley, peaches, cherries, oranges, rice, olives, 
peas, gourds, apples, plantains, tomatoes, bananas, coconuts, grapes, peppers, potatoes, plums, 
strawberries, mangoes, blueberries, coffee, oats, watermelons, lettuce, saguaro cacti, mesquite 
trees, alligator pears, limes, firemelons, pineapples, agave, prickly pears, aloe, firefruit, 
dragonfruit, glowkra, spadebrush, dragonsteeth, bladegrass, gemfruit, magic mushrooms, pudding figs, 
giant beanstalks, axeroot, puppy plants, pickthorn, vigilant vines, rhubarb, frost ferns, chestnuts, 
zephyr cattails, eggplants, carrots, blackberries, pears, flax, popcorn, papayas
```

### How to Use the OLC Editor

The olc editor is a built in editor that lets you create or edit templates of rooms, buildings, objects, triggers, etc.

### Creating a New Building

Useful helps: BUILDING, OLC

You always start by coping an existing building.

.b copy VNUM NEWVNUM