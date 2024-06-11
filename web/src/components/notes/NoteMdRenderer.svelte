<script>
  import { mdConfigParser } from "../../libs/renderers/mdConfigParser";
  import Code from "../renderers/Code.svelte";
  import Markdown from "../renderers/Markdown.svelte";
  import Link from "../renderers/Link.svelte";
  import Youtube from "../renderers/Youtube.svelte";
  import Todos from "../renderers/NoteTodos.svelte";
  import Tracker from "../renderers/Tracker.svelte";
  import { MARKDOWN_PLUGINS as PLUGINS } from "../../const";

  export let note = {};
  let config = mdConfigParser(note?.body ?? {});

  const pluginMapper = {
    [PLUGINS.CODE]: Code,
    [PLUGINS.LINK]: Link,
    [PLUGINS.YOUTUBE]: Youtube,
    [PLUGINS.TODO]: Todos,
    [PLUGINS.TRACKER]: Tracker,
    [PLUGINS.NONE]: Markdown,
  };
</script>

<svelte:component
  this={pluginMapper[config.plugin] ?? pluginMapper[PLUGINS.NONE]}
  {note}
/>
