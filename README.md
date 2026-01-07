This repository is a simple example that reproduces https://github.com/dweymouth/supersonic/issues/560

The app is just a slider that is updated by a goroutine. I have not tested whether other widgets would produce the same result.

If I run the app and put it on a workspace that isn't displayed, then it stops 

# Sample output

```
bct ❯ ./supersonic
2026/01/07 14:18:31 PlatformError: Wayland: Focusing a window requires user interaction
2026/01/07 14:18:32 tick 1
2026/01/07 14:18:32 tick 1 - start update
2026/01/07 14:18:32 tick 1 - finish update
2026/01/07 14:18:32 tick 2
2026/01/07 14:18:32 tick 2 - start update
2026/01/07 14:18:32 tick 2 - finish update
... snip repetitive output ...
2026/01/07 14:18:39 tick 15
2026/01/07 14:18:39 tick 16
2026/01/07 14:18:40 tick 17
2026/01/07 14:18:40 tick 18
2026/01/07 14:18:41 tick 19
2026/01/07 14:18:41 tick 20
2026/01/07 14:18:42 tick 21
2026/01/07 14:18:42 tick 22
2026/01/07 14:18:43 tick 23
2026/01/07 14:18:43 tick 24
... snip repetitive output ...
2026/01/07 14:18:59 tick 55
2026/01/07 14:18:59 tick 56
2026/01/07 14:19:00 tick 57
2026/01/07 14:19:00 tick 58
2026/01/07 14:19:01 tick 58 - start update
2026/01/07 14:19:01 tick 58 - finish update
2026/01/07 14:19:01 tick 58 - start update
2026/01/07 14:19:01 tick 58 - finish update
2026/01/07 14:19:01 tick 58 - start update
2026/01/07 14:19:01 tick 58 - finish update
2026/01/07 14:19:01 tick 58 - start update
2026/01/07 14:19:01 tick 58 - finish update
2026/01/07 14:19:01 tick 58 - start update
2026/01/07 14:19:01 tick 58 - finish update
2026/01/07 14:19:01 tick 58 - start update
2026/01/07 14:19:01 tick 58 - finish update
2026/01/07 14:19:01 tick 58 - start update
2026/01/07 14:19:01 tick 58 - finish update
2026/01/07 14:19:01 tick 58 - start update
2026/01/07 14:19:01 tick 58 - finish update
2026/01/07 14:19:01 tick 58 - start update
2026/01/07 14:19:01 tick 58 - finish update
```

The first few ticks are processed normally. When the window is hidden, the `fyne.Do` updates stop being processed ("start update"/"finish update" are not printed). When the window is redisplayed, the `fyne.Do` updates suddenly all get processed.

At this point the window is still usable. Later, the slider freezes in place totally. Oddly, the `Printf` messages look completely normal at this point.

# Vitalstatistix

```
OS: NixOS 25.11
Graphics Platform: Wayland
Compositor: Hyprland
```
