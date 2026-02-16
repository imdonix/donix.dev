---
title: "Mad Tycoon"
template: "article"
path: "/mad-tycoon/"
date: 2024-05-10
meta:
  description: "Mad Tycoon is a 2D real-time park management simulation game developed using Java, exploring its gameplay, technical architecture, and development challenges."
  image: "/images/madtycoon-gameplay.png"
---

# Mad Tycoon: 2D Park Management Simulation

"Mad Tycoon" is an 2D real-time simulation tycoon game that emerged from the ELTE IK Software Technology course. Developed by a team comprising of 3, this project aimed to deliver a comprehensive park management experience, challenging players to build, manage, and optimize their own entertainment empires.

<div style="position: relative; padding-bottom: 56.25%; height: 0; overflow: hidden; max-width: 100%; background: #000;">
  <iframe src="https://www.youtube.com/embed/Qn-w_SbsGDc" style="position: absolute; top: 0; left: 0; width: 100%; height: 100%; border:0;" allowfullscreen loading="lazy"></iframe>
</div>

## Core Gameplay and Features

Mad Tycoon players in the world of park management. The primary objective revolves around constructing and strategically placing various attractions, and infrastructure to draw in visitors and keep them happy. 

## Technical Details

Mad Tycoon was built with **Java**, using **JDK 8** as its development kit, and **Apache Ant** for its build process.

### Game Entity Hierarchy

The game employs a well-defined inheritance hierarchy to structure its in-game entities, promoting modularity and extensibility. At its foundation is the `GameObject` class, which provides core functionalities common to all objects within the game world. Building upon this, the `Entity` class introduces properties and behaviors specific to active, dynamic elements, such as visitors and workers. Further specialization occurs in subclasses for specific types of entities like `Visitor` or various `Building` types.

Here’s a simplified view of the `objects` package directory structure:

```
src/hu/elte/madtycoon/objects/
├── buildings/
│   ├── decoration/
│   ├── games/
│   ├── Entrance.java
│   ├── Road.java
│   └── Shop.java
├── emotes/
├── entities/
│   ├── Cleaner.java
│   ├── RepairMan.java
│   ├── ShopAssistant.java
│   └── Visitor.java
├── Building.java
├── Decoration.java
├── Entity.java
├── Game.java
├── GameObject.java
└── Worker.java
```

### Real-time Simulation Architecture
The "real-time simulation" aspect is where much of the technical complexity lies. A typical architecture would involve:

- **Game Loop:** A continuous cycle that updates game state, handles user input, and renders the scene. This loop needs to be carefully managed to ensure consistent timing and responsiveness.
- **Event System:** A mechanism for handling various in-game events, such as a ride breaking down, a visitor getting hungry, or a new research item being unlocked.
- **Entity Management:** Representing and updating thousands of individual entities (visitors, rides, staff) and their interactions. This often benefits from an organized system like an Entity-Component-System (ECS) or a well-defined object-oriented hierarchy.
- **Pathfinding:** For visitors navigating the park, an efficient pathfinding algorithm (e.g., A* or Dijkstra's) would be essential to simulate realistic movement and avoid congestion.
- **AI for Visitors and Staff:** Simple AI behaviors would be needed for visitors to make decisions (which ride to go on, where to eat) and for staff to perform their duties (cleaning, maintenance).

## Under the Hood

To further illustrate the technical underpinnings of Mad Tycoon, let's explore some key code snippets from codebase. These examples highlight the real-time simulation, player interaction, and fundamental architectural decisions.

### The Game Loop: Driving Real-time Simulation

The heart of any real-time simulation game is its game loop. In Mad Tycoon, this is managed within the `Engine.java` class, responsible for updating the game state, handling rendering, and processing user interface elements continuously.

```java
// hu/elte/madtycoon/core/Engine.java
private void loop(ActionEvent event)
{
    // Reinitialize the render buffer for the current frame
    renderBuffer = new SpriteRenderBuffer(RENDER_BASE_CAPACITY);

    // Calculate delta time for frame-rate independent updates
    float delta = time() * timeScale;
    time += delta;

    // Update the game world's logic
    world.update(delta);

    // Render all game objects in the world
    world.render(renderBuffer);

    // Update the various UI elements
    hud.updateGUI();
    builder.updateGUI();

    // Request a repaint of the game canvas
    canvas.repaint();
}
```

**Explanation:** This snippet reveals the core `loop` method, triggered by a Swing `Timer`. Each iteration calculates the time elapsed since the last frame (`deltaTime`), ensuring that game updates are frame-rate independent. It then orchestrates the crucial steps of a game cycle: updating the `world` (where all game logic and object states are processed), rendering all visual elements through a `renderBuffer`, refreshing the HUD and builder tools, and finally repainting the display. This is a classic pattern for managing a dynamic, real-time game environment.

## Conclusion

Mad Tycoon provided the development experience in game design, real-time system architecture, and collaborative software engineering, culminating in a robust 2D park management simulation. The project is delivered a playable game but also served as a rich learning experience.

For those interested in seeing the game in action you can download it from its github page: [github.com/imdonix/mad-tycoon -- Release 1.0.0](https://github.com/imdonix/mad-tycoon/releases/tag/v1.0)
