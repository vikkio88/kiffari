<script>
  import { navigate } from "svelte-routing";
  import DateSeverity from "./shared/DateSeverity.svelte";
  import { formatRelativeNow } from "../libs/dates";
  export let notes = [];
  export let compact = false;
</script>

{#if notes.length < 1}
  <h3>No notes... 🤷</h3>
{:else}
  <ul class="list">
    {#each notes as note}
      <li class="note-item" class:archived={note.archived} class:compact>
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
        <h2>{note.title}</h2>

        <button on:click={() => navigate(`/notes/${note.id}`)}>➡️</button>
      </li>
    {/each}
  </ul>
{/if}

<style>
  .list {
    padding: 0 1rem;
    display: flex;
    flex-direction: column;
    list-style: none;
  }

  .note-item {
    border: 2px white solid;
    border-radius: 10px;
    margin: 1rem 0;
    font-size: 16px;
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
    background-color: #3a3a3a;
  }

  .note-item > h2 {
    flex: 1;
    text-align: center;
    margin: auto;
  }

  .note-item > button {
    margin-left: auto;
  }

  .compact > button {
    margin: unset;
  }

  h2 {
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
