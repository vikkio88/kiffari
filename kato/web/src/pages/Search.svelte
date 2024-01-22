<script>
    import Select from "svelecte";
    import { KATO_API_URL } from "../const";
    import TagsList from "../components/TagsList.svelte";
    let timer;
    let tags = [];
    let controller = null;

    const debounce = (v) => {
        clearTimeout(timer);
        timer = setTimeout(() => {
            if (v.length <= 2) {
                tags = [];
                return;
            }

            if (controller) controller.abort();
            controller = new AbortController();

            fetch(`${KATO_API_URL}/tags?q=${v}`, {
                method: "get",
                signal: controller.signal,
            })
                .then((data) => data.json())
                .then((t) => (tags = t));
        }, 500);
    };
</script>

<h2>Search Tag</h2>
<div class="tagSearch">
    <input type="text" on:keyup={({ currentTarget: { value } }) => debounce(value)} />
</div>

<div class="result">
    <TagsList big {tags} />
</div>

<style>
    .tagSearch {
        color: black;
        padding: 1em;
    }

    .tagSearch input {
        padding: 1em;
        font-size: 18px;
        background-color: #d3d3d3;
        color: black;
    }

    .result {
        display: flex;
        justify-content: center;
    }
</style>
