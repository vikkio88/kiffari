<script>
  import TagSearch from './components/TagSearch.svelte';
  import Note from './components/Note.svelte';
</script>

<main>
  <h1>Kato Web</h1>

  <div class="card">
    <Note />
    <TagSearch />
  </div>

</main>

<style>
  main {
    width: 100%;
  }
</style>
