<script>
    import { navigate } from "svelte-routing";
    import DateSeverity from "../shared/DateSeverity.svelte";
    import { formatRelativeNow } from "../../libs/dates";
    export let note = {};
</script>

<div class="note-item" class:archived={note.archived}>
    <div class="info padded">
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
    <div class="content">
        <h3>{note.title}</h3>
    </div>
    <button class="smaller" on:click={() => navigate(`/notes/${note.id}`)}>
        ➡️
    </button>
</div>

<style>
    .note-item {
        border: var(--default-borders);
        border-radius: var(--border-radius);
        font-size: var(--input-font-size);
        padding: 1rem;
        display: flex;
        align-items: center;
        justify-content: space-between;
    }

    .archived {
        border: 2px white dashed;
    }
    .note-item:hover {
        background-color: var(--div-item-hover-color);
    }

    .note-item h3 {
        flex: 1;
        text-align: center;
        margin: auto;
        margin-top: 0;
        margin-bottom: 0;
    }
    .content {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: flex-start;
    }
    .padded {
        padding: 1rem;
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
