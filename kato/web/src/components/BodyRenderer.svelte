<script>
  import { mdConfigParser } from "../libs/renderers/mdConfigParser";
  import Code from "./renderers/Code.svelte";
  import Markdown from "./renderers/Markdown.svelte";
  import Link from "./renderers/Link.svelte";

  export let body = "";
  let config = mdConfigParser(body);

  const pluginMapper = {
    Code: Code,
    Link: Link,
    default: Markdown,
  };
</script>

<svelte:component
  this={pluginMapper[config.Plugin] ?? pluginMapper.default}
  {body}
/>
