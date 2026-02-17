# Architecture Overview

High-level view of the system architecture. Packages are grouped by responsibility layer.

```mermaid
flowchart LR
    subgraph CLI["CLI Layer"]
        main["main<br>(cmd/ascii-art)"]
    end

    subgraph Input["Input Processing"]
        flagparser["flagparser<br>CLI validation"]
        color["color<br>Color parsing"]
    end

    subgraph Core["Core Engine"]
        parser["parser<br>Banner loading"]
        renderer["renderer<br>ASCII rendering"]
    end

    subgraph Output["Output Processing"]
        coloring["coloring<br>ANSI color application"]
    end

    main -->|"validates args"| flagparser
    main -->|"parses color spec"| color
    main -->|"loads banner (embedded FS)"| parser
    main -->|"renders text"| renderer
    main -->|"applies color"| coloring

    style CLI fill:#4a90d9,color:#fff
    style Input fill:#7b68ee,color:#fff
    style Core fill:#2ecc71,color:#fff
    style Output fill:#e67e22,color:#fff
```

## Package Responsibilities

| Layer | Package | Responsibility |
|-------|---------|---------------|
| CLI | `main` | Orchestrates all packages, handles I/O |
| Input | `flagparser` | Validates CLI argument structure |
| Input | `color` | Parses color specs (named, hex, RGB) into RGB values |
| Core | `parser` | Reads banner files from provided filesystem, builds character maps |
| Core | `renderer` | Converts text to ASCII art using banner maps |
| Output | `coloring` | Applies ANSI color codes to rendered ASCII art |

## Key Design Decisions

- **Zero inter-package dependencies** — all packages depend only on the Go standard library
- **Main as orchestrator** — `main` is the only package that imports other project packages
- **Stateless packages** — all functions are pure transformations (no global state, no side effects except embedded FS in main)
- **Embedded filesystem** — banner files bundled into binary at compile time for relocatability
