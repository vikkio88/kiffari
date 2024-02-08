<script>
  import { link } from "svelte-routing";
  import { removeComments } from "../../libs/renderers/cleanup";

  export let body = "";

  const links = parseLinks(body);

  function parseLinks(body) {
    const cleanBody = removeComments(body);
    const regex = /(?:^|\n)(\S*)(?:\s+((?:https?:\/\/)?\S+))?(?=\n|$)/sg;

    const links = [];
    const matches = cleanBody.matchAll(regex);

    for (const match of matches) {
      const label = match[1] || match[2];
      const href = match[2];
      links.push({ href, label });
    }
    console.log(JSON.stringify(links));
    return links;
  }
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
