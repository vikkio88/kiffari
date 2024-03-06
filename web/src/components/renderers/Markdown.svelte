<script>
    import SvelteMarkdown from "svelte-markdown";

    export let body = "";
    function copyText(e) {
        navigator.clipboard.writeText(e.target.innerText);
        //TODO: maybe add some sort of tooltip that tells you that the text is copie
        //<span class="tooltiptext">Tooltip text</span> like in
        //https://www.w3schools.com/css/tryit.asp?filename=trycss_tooltip_arrow_bottom
    }

    function handleCopyPaste(e) {
        const { checked } = e.currentTarget;

        console.log(checked);
        const snippets = document.querySelectorAll("code");
        if (checked) {
            snippets.forEach((s) => s.addEventListener("click", copyText));
        } else {
            snippets.forEach((s) => s.removeEventListener("click", copyText));
        }
    }
</script>

<div class="mdBody">
    <SvelteMarkdown source={body} />
</div>
<div>
    <input type="checkbox" on:change={handleCopyPaste} />
</div>

<style>
    .mdBody {
        min-height: 30vh;
    }
</style>
