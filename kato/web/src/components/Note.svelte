<script>
    import SvelteMarkdown from "svelte-markdown";

    let value = `# This is a header
  
  This is a paragraph.
  
  * This is a list
  * With two items
    1. And a sublist
    2. That is ordered
      * With another
      * Sublist inside`;
</script>

<div class="note">
    <textarea bind:value />
    <div class="preview">
        <SvelteMarkdown source={value} />
    </div>
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
