<script>
  import NoteList from "./NoteList.svelte";
  export let notes = {};
</script>

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
