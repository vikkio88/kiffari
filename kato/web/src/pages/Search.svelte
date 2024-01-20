<script>
    import Select from "svelecte";
    import { KATO_API_URL } from "../const";
    let value = "";
    let timer;

    const debounce = (v) => {
        clearTimeout(timer);
        timer = setTimeout(() => {
            value = v;
        }, 500);
    };
</script>

<div class="tagSearch">
    <input
        type="text"
        on:keyup={({ target: { value } }) => debounce(value)}
    />
</div>

<h1>{value}</h1>

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
</style>
