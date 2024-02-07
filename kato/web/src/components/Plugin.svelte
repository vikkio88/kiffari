<script>
  import { mdConfigParser } from "../libs/mdConfigParser";

  export let config = "";
  let plugins = parseConfig(config);

  function parseConfig(config) {
    return mdConfigParser(config);
  }
</script>

<pre>
    {JSON.stringify(plugins, null, 2)}
</pre>
