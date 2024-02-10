<script>
  import NoteList from "../components/NoteList.svelte";
  import { catchLogout, getLatestNotes, parseOrThrow } from "../libs/api";
  import { protectedRoute } from "../libs/routes";

  protectedRoute();
  let notePromise = getLatestNotes().then(parseOrThrow).catch(catchLogout);
  let reminderNotesPromise = null;
</script>

{#if Boolean(reminderNotesPromise)}
  {#await reminderNotesPromise then reminderNotes}
    {#if reminderNotes.length > 0}
      <NoteList notes={reminderNotes} />
    {/if}
  {/await}
{/if}
<h2>Latest Notes</h2>
{#await notePromise then notes}
  {#if notes.length > 0}
    <NoteList {notes} />
  {:else}
    <h3 class="empty">No notes yet... 🤷</h3>
  {/if}
{/await}

<style>
  h3.empty {
    margin-top: 8rem;
  }
</style>
