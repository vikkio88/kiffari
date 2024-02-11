<script>
  import { navigate } from "svelte-routing";
  import TimeAgo from "javascript-time-ago";
  const timeAgo = new TimeAgo("en-GB");
  export let notes = [];

  function dueSeverityClass(dueDate) {
    const now = new Date();
    const date = new Date(dueDate);

    const diff = (date.getTime() - now.getTime()) / 1000;

    if (diff <= 24 * 3600) {
      return "due-now";
    }

    if (diff <= 7 * 24 * 3600) {
      return "due-soon";
    }

    return "due-later";
  }
</script>

<ul class="list">
  {#each notes as note}
    <li class="note-item" class:archived={note.archived}>
      <div class="info">
        <div class="dates">
          {#if Boolean(note.due_date)}
            <div
              class={`due ${dueSeverityClass(note.due_date)}`}
              title={`due on: ${new Date(note.due_date).toLocaleString()}`}
            >
              {`⏰: ${new Date(note.due_date).toLocaleString()}`}
            </div>
          {/if}
          <div
            title={`last updated: ${new Date(
              note.updated_at,
            ).toLocaleString()}`}
          >
            {timeAgo.format(new Date(note.updated_at), "twitter-now")}
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

<style>
  .list {
    padding: 0 1rem;
    display: flex;
    flex-direction: column;
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
  .due-now {
    color: #db3a34;
  }

  .due-soon {
    color: #ed9b40;
  }

  .due-later {
    color: #3bceac;
  }
</style>
