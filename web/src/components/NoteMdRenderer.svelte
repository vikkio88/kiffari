<script>
  import { mdConfigParser } from "../libs/renderers/mdConfigParser";
  import Code from "./renderers/Code.svelte";
  import Markdown from "./renderers/Markdown.svelte";
  import Link from "./renderers/Link.svelte";
  import Youtube from "./renderers/Youtube.svelte";
  import Todos from "./renderers/NoteTodos.svelte";

  export let note = {};
  let config = mdConfigParser(note?.body ?? {});

  const pluginMapper = {
    code: Code,
    link: Link,
    youtube: Youtube,
    todo: Todos,
    default: Markdown,
  };
</script>

<svelte:component
  this={pluginMapper[config.plugin] ?? pluginMapper.default}
  {note}
/>
