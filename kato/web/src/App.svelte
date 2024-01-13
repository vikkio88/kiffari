<script>
  import TagSearch from "./components/TagSearch.svelte";
  import Note from "./components/Note.svelte";
  import { tags } from "./store";
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
