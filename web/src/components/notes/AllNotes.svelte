<script>
  import NoteItem from "./NoteItem.svelte";

  export let notes = [];
</script>

<div class="list">
  {#each notes as note}
    <NoteItem {note} allowPin={false} />
  {/each}
</div>

<style>
    .list {
        max-width: 90vw;
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        gap: 1rem;
    }
</style>
