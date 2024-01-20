<script>
    import NoteList from "../components/NoteList.svelte";
    import { KATO_API_URL } from "../const";

    export let id = "";
    let tagPromise = fetch(`${KATO_API_URL}/tags/${id}`).then((resp) =>
        resp.json(),
    );
</script>

{#await tagPromise then tag}
    <h2>{tag.label}</h2>
    <h3 class="noteHead">Notes with this tag ({tag.notes.length})</h3>
    <NoteList notes={tag.notes} />
{/await}

<style>
    .noteHead {
        margin-top: 5em;
        margin-bottom: 0;
    }
</style>
