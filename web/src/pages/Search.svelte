<script>
  import TagsList from "../components/TagsList.svelte";
  import Spinner from "../components/shared/Spinner.svelte";
  import { protectedRoute } from "../libs/routes";
  import { filterTags, filterNotes } from "../libs/api";
  import NoteList from "../components/notes/NoteList.svelte";
  import BtnGroup from "../components/shared/BtnGroup.svelte";
  protectedRoute();

  const TABS = {
    TAGS: "tags",
    NOTES: "notes",
  };

  let searchBar;
  let timer;
  let results = [];
  let controller = null;
  let searchPromise = null;

  const debounce = (v) => {
    clearTimeout(timer);
    timer = setTimeout(() => {
      if (v.length <= 2) {
        results = [];
        return;
      }

      if (controller) controller.abort();
      controller = new AbortController();

      let search = activeTab === TABS.TAGS ? filterTags : filterNotes;

      searchPromise = search(v, controller);
    }, 500);
  };

  let activeTab = TABS.NOTES;
  function setActiveTab(tab) {
    results = [];
    searchPromise = null;
    searchBar.value = "";
    searchBar.focus();
    activeTab = tab;
  }
</script>

<h2>Search</h2>
<BtnGroup>
  <button
    class:active={activeTab === TABS.NOTES}
    on:click={() => setActiveTab(TABS.NOTES)}>Notes</button
  >
  <button
    class:active={activeTab === TABS.TAGS}
    on:click={() => setActiveTab(TABS.TAGS)}>Tags</button
  >
</BtnGroup>
<div class="tagSearch">
  <input
    bind:this={searchBar}
    type="text"
    class="searchBar"
    on:keyup={({ currentTarget: { value } }) => debounce(value)}
    placeholder={activeTab === TABS.TAGS ? "Tag..." : "Note title, body...."}
  />
</div>

{#if searchPromise}
  {#await searchPromise}
    <Spinner />
  {:then results}
    {#if activeTab === TABS.TAGS}
      <TagsList big tags={results} />
    {:else}
      <NoteList notes={results} />
    {/if}
  {/await}
{:else}
  <strong>Here there will be search results...</strong>
{/if}

<style>
  h2 {
    margin: 0;
  }
  .tagSearch {
    color: black;
    padding: var(--input-default-pad);
  }

  .searchBar {
    min-width: 40%;
  }
  .tagSearch input {
    padding: var(--input-default-pad);
    font-size: var(--input-medium-font-size);
    background-color: var(--dirty-white);
    color: black;
  }

  strong {
    color: var(--x-light-gray);
  }
</style>
