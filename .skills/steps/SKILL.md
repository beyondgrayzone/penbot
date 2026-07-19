---
title: Steps
description: A component for displaying sequential step-by-step instructions.
---

# Steps

**Import:** `import { Steps, Step } from "@penbot/core";`

## Usage

Display a sequential series of steps.

```svelte
<Steps>
  <Step>Install the package</Step>
  
  You can install via `npm` or `pnpm`.
  
  ```bash
  npm install @penbot/ui
  ```
  
  <Step>Start your engines</Step>
  
  Run `npm run dev` to start.
  
  ```bash
  npm run dev
  ```
</Steps>
```

> **Note:** Leave a blank line between components and markdown content in `.md` files.
