<script>
    import { Link } from "svelte-routing";
    export let notes = [];
</script>

<ul class="list">
    {#each notes as note}
        <li class="note-item">
            <Link to={`notes/${note.id}`}>{note.title}</Link>
            <span class="date">
                {new Date(note.created_at).toLocaleDateString()}
            </span>
        </li>
    {/each}
</ul>

<style>
    .list {
        display: flex;
        flex-direction: row;
    }
    .note-item {
        border: 2px white solid;
        border-radius: 10px;
        padding: 0.5em;
        margin: 1em;
        display: flex;
        flex-direction: column;
        font-size: 16px;
        flex: 1 1 0;
    }

    span.date {
        font-size: x-small;
    }
</style>
