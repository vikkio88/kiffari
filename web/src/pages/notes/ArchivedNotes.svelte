<script>
    import NoteList from "../../components/NoteList.svelte";
    import { getArchivedNotes } from "../../libs/api";
    import { protectedRoute } from "../../libs/routes";
    protectedRoute();
    let notePromise = getArchivedNotes();
</script>

<h2>Archived Notes</h2>
{#await notePromise then notes}
    {#if notes.length > 0}
        <NoteList {notes} />
    {:else}
        <h3 class="empty">No archived notes... 🤷</h3>
    {/if}
{/await}

