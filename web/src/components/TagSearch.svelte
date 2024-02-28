<script>
  import { createEventDispatcher } from "svelte";
  import { filterTags, trendingTags } from "../libs/api";
  import TagsList from "./TagsList.svelte";
  import Spinner from "./shared/Spinner.svelte";
  const d = createEventDispatcher();

  function onTagSelected({ id = null, label }) {
    const newTag = { label };
    if (Boolean(id)) {
      newTag.id = id;
    }
    selectedTags = [...selectedTags, newTag];
    d("updatedSelection", selectedTags);
  }

  function deselect(tagLabel) {
    selectedTags = selectedTags.filter((t) => t.label != tagLabel);
    d("updatedSelection", selectedTags);
  }

  export let suggestTrending = false;
  export let selectedTags = [];
  let searchResultTags = [];

  let timer;
  let controller = null;
  let searchPromise = null;
  let searchValue = "";

  const debounce = (v, k) => {
    searchPromise = null;
    clearTimeout(timer);
    timer = setTimeout(() => {
      if (v.length <= 2) {
        searchResultTags = [];
        return;
      }

      if (controller) controller.abort();
      controller = new AbortController();

      searchPromise = filterTags(v, controller);
    }, 500);
  };

  if (suggestTrending) {
    searchPromise = trendingTags();
  }
</script>

<div class="tagSearch">
  <div class="selectedTags">
    {#each selectedTags as tag}
      <div class="tag" class:NewTag={!Boolean(tag.id)}>
        <span class="hash">#</span>{tag.label}
        <button on:click={() => deselect(tag.label)}>❌</button>
      </div>
    {/each}
  </div>
  <input
    type="text"
    width="50%"
    bind:value={searchValue}
    on:keyup={({ currentTarget: { value }, key }) => debounce(value, key)}
    placeholder="Add Tags..."
  />
  <div class="searchResultTags">
    {#if searchPromise}
      {#await searchPromise}
        <Spinner />
      {:then tags}
        {#if tags.length > 0}
          {#if !Boolean(searchValue)}
            <h3>Suggested tags:</h3>
          {/if}
          <TagsList
            tags={tags.filter(
              (t) => !selectedTags.some((st) => st.id === t.id)
            )}
            onTagClick={(tag) => () => {
              onTagSelected(tag);
              searchValue = "";
              searchPromise = null;
            }}
          />
        {:else if Boolean(searchValue)}
          <button
            on:click={() => {
              onTagSelected({ label: searchValue });
              searchValue = "";
              searchPromise = null;
            }}>🆕 {searchValue}</button
          >
        {/if}
      {/await}
    {/if}
  </div>
</div>

<style>
  .selectedTags {
    padding: 2rem;
    display: flex;
    flex-direction: row-reverse;
    align-items: center;
    justify-content: flex-start;
    overflow: auto;
  }

  .selectedTags > .tag {
    border: solid 2px #333;
    border-radius: 10px;
    padding: 0.5rem 1rem;
    margin: 0 0.5rem;
  }

  .selectedTags > .tag:hover {
    background-color: #333;
    border-color: white;
  }

  .selectedTags > .NewTag {
    border: dashed 2px #333;
  }
  .selectedTags > .tag > button {
    font-size: 0.5rem;
    margin-left: 0.5rem;
  }

  input {
    padding: 1em;
    border-radius: 10px;
  }

  .searchResultTags {
    margin-top: 1rem;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
  }

  .hash {
    font-size: 1.1rem;
    color: #646cff;
    margin-right: 0.2rem;
  }
</style>
