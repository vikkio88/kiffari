<script>
  import { mdConfigParser } from "../libs/renderers/mdConfigParser";
  import Code from "./renderers/Code.svelte";
  import Markdown from "./renderers/Markdown.svelte";
  import Link from "./renderers/Link.svelte";
  import Youtube from "./renderers/Youtube.svelte";

  export let body = "";
  let config = mdConfigParser(body);

  const pluginMapper = {
    Code: Code,
    Link: Link,
    Youtube: Youtube,
    default: Markdown,
  };
</script>

<svelte:component
  this={pluginMapper[config.Plugin] ?? pluginMapper.default}
  {body}
/>
