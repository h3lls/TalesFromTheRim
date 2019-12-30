# Immortal Documentation

Documentation for building and running the mud.

## Setting up New Immortals

In the player file (.plr) you must have access level > 0:

Levels: 

To edit find the following line in the .plr file.
Access Level:

## Special Commands

You can see a list of Immortal / God commands by typing:

```
wizhelp
```

## Creating instances

Use the commands to bring things to life in the MUD based on an existing vnum (unique number for a template of an object of some type).

### Creating a City

To create a city, first claim an area of land. After it has been claimed you can then found a city.
```
claim
```

Founding a city will have you enter the name of the city and it will start out small.
```
city found
```

You can then upgrade your city with:

```
city upgrade
```

The available city points for upgrade will depend on the tech in that city.

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

### Create Instances of Buildings

_Useful Help_: BUILD, DESIGNATE, CREATE

```
build <buildingtype>
```

Available Building Types:
```
 baths, cabinetofcuriosities, courtyard, fountain, greathall, portal
 portcullis, secretcache, sorcerytower, trapperspost, hightemple, library, mint
 pyramid, estuarygarden, greatwall, oasisgarden, raftgarden, rivergate
 shadygarden, terracedgarden, swampcomplex, swamphut, swampmanor, swampplatform
 swampwalk, barracks, gatehouse, guardtower, mountainoutpost, tradingpost
 trainingyard, wall, cemetery, henge, megalith, park, shrine, statue, temple
 tomb, burrowhaven, cavehaven, docks, househaven, pueblohaven, shipyard
 alchemist, apiary, bakery, bridge, carpenter, glassblower, oilmaker
 pigeonpost, steps, tailor, tavern, warehouse, caveshelter, jungleshelter
 rockshelter, treeshelter, largehouse, minecomplex, stonehouse, treecomplex
 treehouse, undergroundcomplex, astralnexus, cabin, cannery, claypit
 cliffdwelling, drainage, familygravesite, fence, forge, foundry, gate, granary
 grave, gravelpit, house, hut, lumberyard, mill, mine, mudhut, potter, pueblo
 quarry, smokehouse, stable, tannery, utilityshed, well
 ```

You will need the resources to build the building.

**Example:**
```
> .map terrain plains
This room is now a Plains.
> create 10 block
You create a huge stone block (x10)!
> build wall east
You start to build a wall!
You heave a huge stone blcok into place.
```

#### Renaming / Desc a Building

The customize command can be used to give a building a new name and a different description. This is a great idea since the initial building will be based on the template.

#### Opening the building for the public

You can use the publicize command to make a building public.

#### Adding Rooms

After a building is created you can use the designate to add new rooms to the base building.  This can be used to add a variety of different rooms. The first room must be across from the entrance. After that you can build in any cardinal direction.

**Example:**
```
> .map terrain plains
This room is now a Plains.
> build barracks south
The animals are herded out of the building...
You start to build a barracks!
You run out of resources and stop working.
> create 10 block
You create a huge stone block (x20)!
> create 40 lumber
You create a stack of lumber (x40)!
> create 10 nails
You create a pouch of nails (x10)!
> build
You start building.
> designate north store room
You designate the room north as a Store Room.
```
Available Rooms:
```
Bedroom, Dining Room, Hallway, Kitchen, Sitting Room, Study, Store Room, Closet, Larder
```

### Creating Instances of Objects for Testing

_Useful Help_: LOAD

First find the object with:
```
.o find <thing>
```
After that choose the vnum and load the object:
```
load obj <vnum>
```

This will create an instance of an object and place it in your inventory.

**Example:**
```
> .o find knife

Matching objects:
[ 2215] a copper knife
[ 9031] the barkeep's knife
[10853] a stone skinning knife

> load obj 10853
You create a stone skinning knife.
```

## Creating New Things

### How to Use the OLC Editor

The olc editor is a built in editor that lets you create or edit templates of rooms, buildings, objects, triggers, etc.

### Creating a New Building

_Useful Help_: BUILDING, OLC

You always start by coping an existing building.

.b copy VNUM NEWVNUM