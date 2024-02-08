<script>
  import { removeComments } from "../../libs/renderers/cleanup";
  import { parseLinks } from "../../libs/renderers/extractors";

  export let body = "";

  const links = parseLinks(body);
</script>

<div class="urls">
  {#each links as link}
    <a class="link" href={link.href} target="_blank">{link.label}</a>
  {/each}
</div>

<style>
  a.link {
    font-size: 2rem;
  }

  .urls {
    display: flex;
    flex-direction: column;
  }
</style>
