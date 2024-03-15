<script>
  import { navigate } from "svelte-routing";
  import DateSeverity from "./shared/DateSeverity.svelte";
  import { formatRelativeNow } from "../libs/dates";
  import { previewMd } from "../libs/renderers/cleanup";
  export let notes = [];
  export let compact = false;
</script>

{#if notes.length < 1}
  <h3>No notes... 🤷</h3>
{:else}
  <div class="list" class:compact>
    {#each notes as note}
      <div class="note-item" class:archived={note.archived} class:compact>
        <div class="info">
          <div class="dates">
            <DateSeverity date={note.due_date} />
            <div
              title={`last updated: ${new Date(
                note.updated_at
              ).toLocaleString()}`}
            >
              {formatRelativeNow(note.updated_at)}
            </div>
          </div>
          {#if note.archived}
            <span class="crs-pointer" title="Archived">🗄️</span>
          {/if}
        </div>
        <h3>{note.title}</h3>
        <p>{previewMd(note.body)}</p>
        <button on:click={() => navigate(`/notes/${note.id}`)}>➡️</button>
      </div>
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
    display: grid;
    grid-template-columns: auto auto;
  }

  .note-item {
    border: var(--default-borders);
    border-radius: var(--border-radius);
    font-size: var(--input-font-size);
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    padding: 1rem 2rem;
  }

  .compact {
    flex-direction: column;
  }

  .archived {
    border: 2px white dashed;
  }
  .note-item:hover {
    background-color: var(--div-item-hover-color);
  }

  .note-item > h3 {
    flex: 1;
    text-align: center;
    margin: auto;
    margin-top: 0;
    margin-bottom: 0;
  }

  .note-item > button {
    margin-left: auto;
  }

  .compact > button {
    margin: unset;
  }

  h3 {
    margin: 0 auto;
    text-align: center;
  }

  .info {
    display: flex;
    gap: 1rem;
  }

  .dates {
    display: flex;
    flex-direction: column;
  }
</style>
