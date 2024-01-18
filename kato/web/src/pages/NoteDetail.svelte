<script>
  import { KATO_API_URL } from "../const";
  import SvelteMarkdown from "svelte-markdown";
  import TagsList from "../components/TagsList.svelte";

  export let id = "";
  let notePromise = fetch(`${KATO_API_URL}/notes/${id}`).then((resp) =>
    resp.json()
  );
</script>

{#await notePromise then note}
  <SvelteMarkdown source={note.body} />
  <TagsList tags={note.tags} />
{/await}
