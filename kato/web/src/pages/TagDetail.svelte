<script>
    import { KATO_API_URL } from "../const";

    export let id = "";
    let tagPromise = fetch(`${KATO_API_URL}/tags/${id}`).then((resp) =>
        resp.json(),
    );
</script>

{#await tagPromise then tag}
    <h2>{tag.label}</h2>
    <h3>{tag.notes.length} notes</h3>
{/await}
