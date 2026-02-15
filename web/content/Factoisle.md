---
title: "FactoIsle"
template: "article"
path: "/factoisle/"
meta:
  description: "FactoIsle is a base-building, automation, and sandbox indie game developed using the Unity Engine. It features procedural terrain generation, a chunk-based world system, flexible save system, and complex automation mechanics, including pipe and item transport systems."
  image: "/images/factoisle_pipes.png"
---

# Making FactoIsle
FactoIsle is a base-building, automation, and sandbox indie game developed using the Unity Engine. It features procedural terrain generation, a chunk-based world system, flexible save system, and complex automation mechanics, including pipe and item transport systems.

Below are some of the key technical aspects that introduced challenges during the game's development.

## Chunks
The game world measures 512x512 units (in Unity metrics). While terrain generation extends beyond this area, any terrain generated outside this boundary is underwater and devoid of entities. Loading the entire world at once is impractical due to the presence of over 100,000 entities, which would overwhelm the Unity Engine. To address this, the world is divided into 32x32 chunks.

<img
  src="/images/factoisle_chunk.png"
  alt="Chunk System"
/>

Chunks are dynamically loaded and unloaded based on the player's proximity. Additionally, a ticketing system ensures that certain chunks remain loaded, allowing machines to operate even when the player is not nearby and guaranteeing that the player's current chunk is always active.

## Saving
Upon initial loading, each chunk populates itself with entities according to its biome. Entities are composed of logic and data components. The data containers are binary serializable and can hold any necessary information for the entity.

When a chunk is unloaded, its entities are processed, and their data containers are saved alongside the entity's ID and type. This method allows the state of each chunk to be stored separately, enabling the entire game world to be saved in a single file of less than 10MB.

Example usage of data containers in an entity:
```cs
NBData tag = Data.Get<NBData>(); // Named Binary Data
if (tag.Has("itemInMouth"))
{
    int logType = tag.GetInt("itemInMouth");
    Entity entity = Library.Entity(logType);
    if (entity != null && entity is Log log)
    {
      _itemInMouth = log;
    }
}
```

## Pipes
A core objective in FactoIsle is to automate production lines by connecting machine inputs and outputs. This is achieved through the item and fluid pipe system. Pipes can be placed as blocks, and upon placement, the world triggers a tube system reload. This process connects all pipes into systems, taking into account the inputs and outputs of machines via their ports.

<img
  src="/images/factoisle_pipes.png"
  alt="Pipes"
/>

Pipes possess internal storage and tick mechanisms similar to entities, with their contents saved using the same data container system.

## Multiblocks
Machines in FactoIsle are constructed by arranging specific blocks in predefined patterns to form multiblocks.

An example (multiblock pattern of the Sawmill)
<img
  src="/images/factoisle_pattern.png"
  alt="Pattern"
/>

The multiblock pattern checker can identify all possible rotations of a pattern. The pattern checker activates whenever a block is placed or destroyed.

Example of a multiblock declaration:
```cs
public override void AssembleMultiblockPattern(MultiblockPattern pattern)
{
    pattern.AddBlock(new Vector3Int(0, 0, 0), Library.Entity<BlockPlank>());
    pattern.AddBlock(new Vector3Int(1, 0, 0), Library.Entity<BlockPlank>());
    pattern.AddBlock(new Vector3Int(2, 0, 0), Library.Entity<BlockPlank>());

    pattern.AddBlock(new Vector3Int(0, 0, 1), Library.Entity<BlockPlank>());
    pattern.AddBlock(new Vector3Int(1, 0, 1), Library.Entity<BlockPlank>());
    pattern.AddBlock(new Vector3Int(2, 0, 1), Library.Entity<BlockPlank>());

    pattern.AddBlock(new Vector3Int(1, 1, 1), Library.Entity<BlockMachineAdvanced>());
    pattern.AddBlock(new Vector3Int(1, 1, 0), Library.Entity<BlockMachineBasic>());
}
```

## Quests
Quests serve as the primary means of progression in FactoIsle, unlocking new items, liquids, and machines.

<img
  src="/images/factoisle_quests.png"
  alt="Quests"
/>

Each quest is implemented as a separate class extending the Quest base class. Quests are defined declaratively, specifying requirements, unlocks, and dependencies on other quests. They are loaded via reflection by the story system and can store custom data using entity data containers.

Example of a quest declaration (`ChopTreeQ.cs`): 
```cs
public class ChopTreeQ : Quest
{
    public override Sprite GetSprite()
    {
        return Library.Entity<OakTree>().GetSprite(); // Sprite in quest book
    }
    
    protected override void OnInit(Action<Requirement> add)
    {
        // chop down 2 trees
        add(new GatherRequirement(Library.Entity<OakTree>(), 2));
        // pickup 10 logs
        add(new PickupRequirement(Library.Entity<LogOak>(), 10)); 
    }
    
    public override List<Quest> UnlockedBy(Story story)
    {
        return new List<Quest>() { story.GetQuest<PickupStoneQ>() };
    }
    
    public override List<Entity> UnlockedEntities()
    {
        return new List<Entity>
        {
            Library.Entity<Plank>(),
        };
    }
    
    public override Chapter GetChapter(Story story)
    {
        return story.GetChapter<TutorialC>();
    }
}
```

## Thanks
Thank you for reading this overview of FactoIsle's development.
The game is now available in alpha:
<a href="https://imdonix.itch.io/factoisle">imdonix.itch.io/factoisle</a> 

If you're interested in following the game's progress, seeing updates, or looking for game development tips and tricks, please follow me on X:
<a href="https://x.com/imdonix">@imdonix</a>
Feel free to reach out with any questions or for assistance!
