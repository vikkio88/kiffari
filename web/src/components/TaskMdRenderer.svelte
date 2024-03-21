<script>
  import { MARKDOWN_PLUGINS as PLUGINS } from "../const";
  import { mdConfigParser } from "../libs/renderers/mdConfigParser";
  import Markdown from "./renderers/TaskMarkdown.svelte";
  import Todos from "./renderers/TaskTodos.svelte";

  export let task = {};
  let config = mdConfigParser(task?.description ?? {});

  const pluginMapper = {
    [PLUGINS.TODO]: Todos,
    [PLUGINS.NONE]: Markdown,
  };
</script>

<svelte:component
  this={pluginMapper[config.plugin] ?? pluginMapper[PLUGINS.NONE]}
  {task}
/>
