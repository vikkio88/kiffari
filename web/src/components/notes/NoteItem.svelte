<script>
    import { navigate } from "svelte-routing";
    import DateSeverity from "../shared/DateSeverity.svelte";
    import { formatRelativeNow } from "../../libs/dates";
    import { previewMd } from "../../libs/renderers/cleanup";
    import { getNoteItemConfig } from "./utils";
    import { catchLogout, del, pinNote, unpinNote } from "../../libs/api";
    import ConfirmButton from "../shared/ConfirmButton.svelte";

    export let allowPin = true;
    export let note = {};
    export let pinned = note?.pinned || false;

    const config = getNoteItemConfig(note);
    async function onPinToggle() {
        const toggle = pinned ? unpinNote : pinNote;
        const resp = await toggle(note.id);
        if (resp.status == 401) {
            catchLogout();
            return;
        }

        window.location.reload();
    }

    async function onDelete() {
        const resp = await del("notes", note.id);
        if (resp.status == 401) {
            catchLogout();
            return;
        }

        window.location.reload();
    }
</script>

<div class="note-item" class:archived={note.archived}>
    {#if allowPin && !note.archived && !Boolean(note.due_date)}
        <div class="top">
            <button
                class="smaller"
                on:click={onPinToggle}
                title={`${pinned ? "Un-Pin" : "Pin"}`}
            >
                {pinned ? "📍" : "📌"}
            </button>
        </div>
    {/if}
    <div class="info padded">
        {#if config.dates}
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
        {/if}
        {#if note.archived}
            <span class="crs-pointer" title="Archived">🗄️</span>
        {/if}
    </div>
    <div class="padded">
        <h3>{note.title}</h3>
        {#if config.preview}
            <p>{previewMd(note.body)}</p>
        {/if}
        {#if config.icon}
            <p class="plugin" title={`Plugin: ${config.title}`}>{config.icon}</p>
        {/if}
    </div>
    <div class="controls">
        {#if config.editBtn}
            <button
                class="smaller"
                on:click={() => navigate(`/edit-note/${note.id}`)}
            >
                📝
            </button>
        {/if}
        {#if config.doneBtn}
            <ConfirmButton
                classes="smaller"
                title="Mark as Done"
                confirmLabel=""
                onConfirmed={onDelete}
            >
                ✅
            </ConfirmButton>
        {/if}
        {#if config.viewBtn}
            <button
                class="smaller"
                on:click={() => navigate(`/notes/${note.id}`)}
            >
                ➡️
            </button>
        {/if}
    </div>
</div>

<style>
    .note-item {
        border: var(--default-borders);
        border-radius: var(--border-radius);
        font-size: var(--input-font-size);
        flex-direction: column;
        padding: 0;
    }

    .top {
        margin-bottom: 0.5rem;
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
    .padded {
        padding: 1rem;
    }

    .controls > button {
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

    p.plugin {
        background-color: var(--fade-gray);
    }
</style>
