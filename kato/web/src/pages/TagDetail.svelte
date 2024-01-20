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
    <strong>Notes with this tag ({tag.notes.length})</strong>
    <NoteList notes={tag.notes} />
{/await}
