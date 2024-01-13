<script>
  import { onMount } from "svelte";
  import TagSearch from "./components/TagSearch.svelte";
  import Note from "./components/Note.svelte";
  import { tags } from "./store";

  onMount(() => {
    fetch("http://localhost:5111/tags")
      .then((data) => data.json())
      .then((t) => $tags = t);
  });
</script>

<main>
  <h1>Kato Web</h1>

  <div class="card">
    <ul>
      {#each $tags as tag}
        <li>{tag.label}</li>
      {/each}
    </ul>
    <Note />
    <TagSearch />
  </div>
</main>

<style>
  main {
    width: 100%;
  }
</style>
