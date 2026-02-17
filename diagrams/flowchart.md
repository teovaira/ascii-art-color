# Program Flowchart

Execution flow from CLI input to ASCII art output. The program has two modes: **Normal** (text only) and **Color** (with ANSI coloring).

```mermaid
flowchart TD
    A["CLI Arguments<br>os.Args"] --> B{"hasColorFlag?<br>--color= prefix"}

    B -->|No| C["ParseArgs()<br>text, banner"]
    B -->|Yes| D["flagparser.ParseArgs()<br>validate syntax"]

    C --> E["GetBannerPath()<br>banner file path"]
    D --> F["extractColorArgs()<br>colorSpec, substring,<br>text, banner"]

    E --> E2["GetBannerFS()<br>embedded filesystem"]
    E2 --> G["parser.LoadBanner(fsys, path)<br>Banner map"]
    F --> H["color.Parse()<br>RGB struct"]

    H --> I["GetBannerPath()<br>banner file path"]
    I --> I2["GetBannerFS()<br>embedded filesystem"]
    I2 --> J["parser.LoadBanner(fsys, path)<br>Banner map"]
    J --> K["color.ANSI()<br>ANSI escape code"]

    G --> L["renderer.ASCII()<br>ASCII art string"]
    L --> M["fmt.Print()<br>stdout"]

    K --> N["For each line in text"]
    N --> O["renderer.ASCII()<br>ASCII art lines"]
    O --> P["parser.CharWidths()<br>character widths"]
    P --> Q["coloring.ApplyColor()<br>colored ASCII art"]
    Q --> R{"More lines?"}
    R -->|Yes| N
    R -->|No| S["fmt.Print()<br>stdout"]

    style B fill:#f39c12,color:#fff
    style R fill:#f39c12,color:#fff
    style M fill:#2ecc71,color:#fff
    style S fill:#2ecc71,color:#fff
```

## Mode Comparison

| Aspect | Normal Mode | Color Mode |
|--------|------------|------------|
| Validation | `ParseArgs()` | `flagparser.ParseArgs()` |
| Color parsing | — | `color.Parse()` → `color.ANSI()` |
| Rendering | Single call | Per-line loop |
| Post-processing | — | `CharWidths()` + `ApplyColor()` |
