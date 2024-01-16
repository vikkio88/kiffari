<script>
    import Select from "svelecte";
    import { createEventDispatcher } from "svelte";

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
        placeholder="Search for tags..."
        fetch="http://localhost:5111/tags?q=[query]"
    />
</div>

<style>
    .tagSearch {
        color: black;
        padding: 1em;
    }
</style>
