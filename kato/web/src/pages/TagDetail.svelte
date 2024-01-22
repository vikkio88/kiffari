<script>
    import NoteList from "../components/NoteList.svelte";
    import { KATO_API_URL } from "../const";

    export let id = "";
    let tagPromise = fetch(`${KATO_API_URL}/tags/${id}`).then((resp) =>
        resp.json(),
    );
</script>

{#await tagPromise then tag}
    <div class="header">
        <span>Tag</span>
        <h1>
            {tag.label}
        </h1>
    </div>
    <div class="results">
        <h3>Notes</h3>
        <NoteList notes={tag.notes} />
    </div>
{/await}

<style>
    .header {
        border-top: 2px white solid;
        margin-bottom: 5em;
    }
    .header > h1 {
        margin: 0;
        font-size: 3rem;
    }

    .header > span {
        font-size: 2rem;
        margin: 0;
        padding: 0 1rem;
        display: inline-block;
        transform: translateY(-50%);
        background-color: #242424;
    }

    .results {
        margin-top: 5rem;
        border-top: 2px white solid;
    }

    .results > h3 {
        margin: 0;
        padding: 0 1rem;
        display: inline-block;
        transform: translateY(-50%);
        background-color: #242424;
    }
</style>
