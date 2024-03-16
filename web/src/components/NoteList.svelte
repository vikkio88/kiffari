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
                note.updated_at,
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
        {#if compact}
          <p>{previewMd(note.body)}</p>
          <div class="controls">
            <button
              class="smaller"
              on:click={() => navigate(`/edit-note/${note.id}`)}
            >
              📝
            </button>
            <button
              class="smaller"
              on:click={() => navigate(`/notes/${note.id}`)}
            >
              ➡️
            </button>
          </div>
        {:else}
          <button
            class="smaller"
            on:click={() => navigate(`/notes/${note.id}`)}
          >
            ➡️
          </button>
        {/if}
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
    display: flex;
    flex-direction: row;
    flex: 1 1 0;
  }

  .note-item {
    border: var(--default-borders);
    border-radius: var(--border-radius);
    font-size: var(--input-font-size);
    padding: 1rem 1rem;
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

  .controls > button {
    margin-left: auto;
  }

  .compact > .controls > button {
    margin: unset;
  }

  .controls {
    display: flex;
    flex-direction: row;
    justify-content: flex-end;
  }

  h3 {
    margin: 0 auto;
    text-align: center;
  }

  .info {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 1rem;
  }

  .dates {
    display: flex;
    flex-direction: column;
  }
</style>
