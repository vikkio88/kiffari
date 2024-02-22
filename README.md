# kiffari
note and tasks manager

This project entails 2 different subprojects, [Kato](#kato) and [Kiffari](#kiffari)

## Kato
from _sicilian_ _Catu(o)_ = Bucket.

A selfhostable webapp to save notes and download them as markdown files (compatible with [Obsidian.md](https://obsidian.md/)).

## How it looks
[![Watch the video](https://img.youtube.com/vi/2p7wZoG-bLw/default.jpg)](https://youtu.be/2p7wZoG-bLw)

## Roadmap
- [x] Create/Edit Notes and Reminders with md format.
- [x] Add `due date` to create reminders.
- [x] Export as simple `.md` files.
- [x] Create set of plugins to load different styles of renderers
  - [x] Link Plugin to just show list of links
  - [x] Youtube/Videoplayer Plugin to showcase video
  - [ ] Code Plugin to show code snippets
- [x] Search by tag
- [x] Search by title/body free text
- [ ] Create/Explore PWA with offline/online sync.

### How to
Run Watch both BE/FE
(make sure you `cd web && nvm use`)
then:

```
make -j4 watch-be watch-fe
```

## Kiffari
from _sicilian_ _(Aju) Chiffari_ = I've got stuff to do.

A project management kanban-ish system to handle tasks.
