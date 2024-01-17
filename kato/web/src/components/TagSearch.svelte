<script>
    import Select from "svelecte";
    import { createEventDispatcher } from "svelte";
    import { KATO_API_URL } from "../const";

    const d = createEventDispatcher();

    function onChange({ detail: selected }) {
        d("added_tag", { tags: selected });
    }
</script>

<div class="tagSearch">
    <Select
        name="selection"
        on:change={onChange}
        multiple
        required
        creatable
        placeholder="Tags..."
        fetch={`${KATO_API_URL}/tags?q=[query]`}
    />
</div>

<style>
    .tagSearch {
        color: black;
        padding: 1em;
    }
</style>
