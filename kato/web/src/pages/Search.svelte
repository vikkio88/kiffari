<script>
  import TagsList from "../components/TagsList.svelte";
  import Spinner from "../components/shared/Spinner.svelte";
  import { protectedRoute } from "../libs/routes";
  import {
    filterTags,
    parseOrThrow,
    catchLogout,
    filterNotes,
  } from "../libs/api";
  import NoteList from "../components/NoteList.svelte";
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

      searchPromise = search(v, controller)
        .then(parseOrThrow)
        .catch(catchLogout);
    }, 500);
  };

  let activeTab = TABS.TAGS;
  function setActiveTab(tab) {
    results = [];
    searchPromise = null;
    searchBar.value = "";
    searchBar.focus();
    activeTab = tab;
  }
</script>

<h2>Search</h2>
<div class="tabs">
  <button
    class:active={activeTab === TABS.TAGS}
    on:click={() => setActiveTab(TABS.TAGS)}>Tags</button
  >
  <button
    class:active={activeTab === TABS.NOTES}
    on:click={() => setActiveTab(TABS.NOTES)}>Notes</button
  >
</div>
<div class="tagSearch">
  <input
    bind:this={searchBar}
    type="text"
    class="searchBar"
    on:keyup={({ currentTarget: { value } }) => debounce(value)}
    placeholder={activeTab === TABS.TAGS ? "Tag..." : "Note title, body...."}
  />
</div>

<div class="result">
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
</div>

<style>
  h2 {
    margin: 0;
  }
  .tagSearch {
    color: black;
    padding: 1em;
  }

  .searchBar {
    min-width: 40%;
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

  strong {
    color: #a3a3a3;
  }

  div.tabs {
    flex-direction: row;
    align-items: center;
    justify-content: space-around;
  }

  div.tabs > button {
    border-radius: 0px;
    margin: 0;
  }
  div.tabs > button.active {
    background-color: black;
    color: #646cff;
  }
  div.tabs > button.active:focus,
  div.tabs > button.active:focus-visible {
    outline: none;
  }
</style>
