<script>
    import SvelteMarkdown from "svelte-markdown";
    import TagSearch from "./TagSearch.svelte";

    let showPreview = false;
    let text = `a new note`;

    let selectedTags = []

    function handleKeydown(e) {
        if (event.key !== 'Tab') return;

        event.preventDefault();

		const { selectionStart, selectionEnd, value } = this;
        
        text = value.substring(0, selectionStart) +"  " + value.substring(selectionEnd);
        this.selectionStart = this.selectionEnd = selectionStart + 1;
    }

    function onTagSelected(e) {
        selectedTags = e.detail.tags
    }

    function onSave() {
        console.log({content: text, tags: selectedTags})
    }

</script>

<div class="editor">
    <TagSearch on:added_tag={onTagSelected} />
    <button on:click={() => showPreview = !showPreview}>Toggle Preview</button>
    <div class="note">
        {#if !showPreview}
        <div>
        <textarea bind:value={text} rows="10" cols="50" on:keydown={handleKeydown}/>
        </div>
        {:else}
        <div class="preview">
            <SvelteMarkdown source={text} />
        </div>
        {/if}
    </div>
    <div class="controls">
        <button on:click={onSave}>Save</button>
    </div>
</div>

<style>
    .editor {
        padding: 2em;
    }
    .note textarea {
        font-size:16px;
    }

    .preview {
        display: block;
    }

    .controls {
        position: absolute;
        bottom: 0;
        right: 0;
        padding: 2em;
    }
</style>
