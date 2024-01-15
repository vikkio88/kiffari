<script>
    import Select from "svelecte";
    function onChange({ detail: selected }) {
        console.log("change", selected);
    }

    function onCreate({ detail: created }) {
        console.log("created", created);
    }
</script>

<div class="tagSearch">
    <Select
        name="selection"
        on:change={onChange}
        multiple
        required
        creatable
        on:createoption={onCreate}
        placeholder="Search for color"
        fetch="http://localhost:5111/tags?q=[query]"
    />
</div>

<style>
    .tagSearch {
        width: 100%;
        color: black;
    }
</style>
