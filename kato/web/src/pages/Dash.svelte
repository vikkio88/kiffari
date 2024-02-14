<script>
  import { navigate } from "svelte-routing";
  import NoteList from "../components/NoteList.svelte";
  import {
    catchLogout,
    getLatestNotes,
    getReminderNotes,
    parseOrThrow,
  } from "../libs/api";
  import { protectedRoute } from "../libs/routes";

  protectedRoute();
  let notePromise = getLatestNotes().then(parseOrThrow).catch(catchLogout);
  let reminderNotesPromise = getReminderNotes()
    .then(parseOrThrow)
    .then((reminders) => {
      if (reminders.length > 0) {
        return reminders;
      }
      reminderNotesPromise = null;
    })
    .catch(catchLogout);

  function create() {
    navigate("/create-note");
  }
</script>

<div class="wrapper">
  <div class="subwrapper">
    <h2>Latest Notes</h2>
    {#await notePromise then notes}
      {#if notes.length > 0}
        <NoteList {notes} compact />
      {:else}
        <h3 class="empty">No notes yet... 🤷</h3>
      {/if}
    {/await}
  </div>

  {#if Boolean(reminderNotesPromise)}
    <div id="reminders" class="subwrapper hidden">
      <h2>Reminders</h2>
      {#await reminderNotesPromise then reminderNotes}
        {#if reminderNotes.length > 0}
          <NoteList notes={reminderNotes} compact />
        {/if}
      {/await}
    </div>
  {/if}
</div>

<button class="add" title="New Note" on:click={create}>📝</button>

<style>
  h3.empty {
    margin-top: 8rem;
  }

  .wrapper {
    display: flex;
    flex-direction: row;
    justify-content: space-around;
    align-items: flex-start;
  }

  .add {
    font-size: 1.2rem;
    /* border-radius: 50rem; */
    padding: 0.5rem 1rem;
    position: absolute;
    right: 0;
    margin-right: 1rem;
  }
</style>
