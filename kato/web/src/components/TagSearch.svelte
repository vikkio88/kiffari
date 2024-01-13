<script>
    import Select from "svelte-select";
    import { tags } from "../store";
    let filterText = "";

    let value = null;

    function handleFilter(e) {
        if (
            value?.find((i) =>
                i.label.toLowerCase().includes(filterText.toLowerCase()),
            )
        )
            return;
        if (e.detail.length === 0 && filterText.length > 0) {
            const prev = $tags.filter((i) => !i.created);
            $tags = [
                ...prev,
                { value: filterText, label: filterText, created: true },
            ];
        }
    }

    function handleChange(e) {
        console.log(e);
        $tags = $tags.map((i) => {
            if (i.created) {
                //CREATE
            }
            delete i.created;
            return i;
        });
    }
</script>

<div class="tagSearch">
    <Select
        on:change={handleChange}
        multiple
        on:filter={handleFilter}
        bind:filterText
        bind:value
        items={$tags}
        on:create
    >
        <div slot="item" let:item>
            {item.created ? "Add new: " : ""}
            {item.label}
        </div>
    </Select>
</div>

<style>
    .tagSearch {
        width: 100%;
        color: black;
    }
</style>
