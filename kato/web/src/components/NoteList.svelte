<script>
    import { navigate } from "svelte-routing";
    export let notes = [];
</script>

<ul class="list">
    {#each notes as note}
        <li class="note-item">
            <span class="date">
                {new Date(note.created_at).toLocaleString()}
            </span>
            <h2>{note.title}</h2>
            <button on:click={() => navigate(`/notes/${note.id}`)}>➡️</button>
        </li>
    {/each}
</ul>

<style>
    .list {
        width: 80%;
        margin: 0 auto;
    }
    .note-item {
        border: 2px white solid;
        border-radius: 10px;
        padding: 2rem;
        margin: 2rem 0;
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: space-around;
        font-size: 16px;
    }

    .note-item:hover {
        background-color: #3a3a3a;
    }

    h2 {
        margin: 0 auto;
        text-align: center;
    }
</style>
