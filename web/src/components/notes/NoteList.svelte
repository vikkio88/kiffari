<script>
  import NoteItem from "./NoteItem.svelte";
  import NoteRow from "./NoteRow.svelte";

  export let notes = [];
  export let compact = false;
</script>

{#if notes.length < 1}
  <h3>No notes... 🤷</h3>
{:else}
  <div class="list" class:compact>
    {#each notes as note}
      {#if compact}
        <NoteItem {note} />
      {:else}
        <NoteRow {note} />
      {/if}
    {/each}
  </div>
{/if}

<style>
  .list {
    padding: 0 1rem;
    display: grid;
    grid-row: auto;
    gap: 1rem;
  }

  .list.compact {
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    align-items: stretch;
    justify-content: center;
  }

  .compact {
    flex-direction: column;
  }

  h3 {
    margin: 0 auto;
    text-align: center;
  }
</style>
