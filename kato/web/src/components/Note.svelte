<script>
    import SvelteMarkdown from "svelte-markdown";
    import TagSearch from "./TagSearch.svelte";

    let value = `
  This is a paragraph.

  * This is a list
  * With two items
    1. And a sublist
    2. That is ordered
      * With another`;
</script>

<div class="editor">
    <div class="note">
        <textarea bind:value />
        <div class="preview">
            <SvelteMarkdown source={value} />
        </div>
    </div>
    <TagSearch />
</div>

<style>
    .note {
        width: 100%;
        display: flex;
        flex-direction: row;
    }

    .preview {
        display: block;
    }
</style>
