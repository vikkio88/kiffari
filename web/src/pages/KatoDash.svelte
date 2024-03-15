<script>
  import Adder from "../components/shared/Adder.svelte";
  import { navigate } from "svelte-routing";
  import {
    catchLogout,
    createNote,
    getLatestNotes,
    getReminderNotes,
    parseOrThrow,
  } from "../libs/api";
  import NoteList from "../components/NoteList.svelte";
  import Controls from "../components/shared/Controls.svelte";
  import Footer from "../components/Footer.svelte";
  import { protectedRoute } from "../libs/routes";
  import { nowString } from "../libs/dates";
  protectedRoute();

  let notePromise = getLatestNotes();
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

  async function onCreate(bodyWrapper) {
    const resp = await createNote({
      title: nowString(),
      tags: [],
      ...bodyWrapper,
    });
    if (resp.ok) {
      notePromise = getLatestNotes();
      return;
    }

    //TODO: handle error
  }
</script>

<Adder
  centered
  placeholder="Note body..."
  on:added={({ detail: body }) => onCreate({ body })}
/>
<div class="wrapper">
  {#if Boolean(reminderNotesPromise)}
    <div id="reminders" class="subwrapper">
      <h2>Reminders</h2>
      {#await reminderNotesPromise then reminderNotes}
        {#if reminderNotes.length > 0}
          <NoteList notes={reminderNotes} compact />
        {/if}
      {/await}
    </div>
  {/if}

  <div class="subwrapper">
    <h2>Latest Notes</h2>
    {#await notePromise then notes}
      <NoteList {notes} compact />
    {/await}
  </div>
</div>

<Footer />

<Controls>
  <button class="big-control" title="New Note" on:click={create}
    >Add Note 📝</button
  >
</Controls>

<style>
  .wrapper {
    display: grid;
    grid-template-rows: auto auto;
    padding-bottom: 2rem;
  }
</style>
