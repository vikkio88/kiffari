<script>
  import Adder from "../components/shared/Adder.svelte";
  import { navigate } from "svelte-routing";
  import { createNote, getDashNotes } from "../libs/api";
  import NoteList from "../components/notes/NoteList.svelte";
  import Controls from "../components/shared/Controls.svelte";
  import Footer from "../components/Footer.svelte";
  import { protectedRoute } from "../libs/routes";
  import { nowString } from "../libs/dates";
  import Spinner from "../components/shared/Spinner.svelte";
  protectedRoute();

  let dashNotesPromise = getDashNotes();

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
      dashNotesPromise = getDashNotes();
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
  {#await dashNotesPromise}
    <Spinner />
  {:then notes}
    {#if Array.isArray(notes.pinned) && notes.pinned.length > 0}
      <div class="subwrapper">
        <h2>📍 Notes</h2>
        <NoteList notes={notes.pinned} compact />
      </div>
    {/if}

    {#if Array.isArray(notes.reminders) && notes.reminders.length > 0}
      <div class="subwrapper">
        <h2>⏰ Notes</h2>
        <NoteList notes={notes.reminders} compact />
      </div>
    {/if}
    {#if Array.isArray(notes.latest) && notes.latest.length > 0}
      <div class="subwrapper">
        <h2>🆕 Notes</h2>
        <NoteList notes={notes.latest} compact />
      </div>
    {/if}
  {/await}
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
