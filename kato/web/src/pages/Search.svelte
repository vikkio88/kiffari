<script>
  import TagsList from "../components/TagsList.svelte";
  import Spinner from "../components/shared/Spinner.svelte";
  import { protectedRoute } from "../libs/routes";
  import { filterTags, parseOrThrow, catchLogout } from "../libs/api";
  protectedRoute();

  let timer;
  let tags = [];
  let controller = null;
  let searchPromise = null;

  const debounce = (v) => {
    clearTimeout(timer);
    timer = setTimeout(() => {
      if (v.length <= 2) {
        tags = [];
        return;
      }

      if (controller) controller.abort();
      controller = new AbortController();

      searchPromise = filterTags(v, controller)
        .then(parseOrThrow)
        .catch(catchLogout);
    }, 500);
  };
</script>

<h2>Search</h2>
<div class="tagSearch">
  <input
    type="text"
    size="30"
    on:keyup={({ currentTarget: { value } }) => debounce(value)}
    placeholder="Tag..."
  />
</div>

<div class="result">
  {#if searchPromise}
    {#await searchPromise}
      <Spinner />
    {:then tags}
      <TagsList big {tags} />
    {/await}
  {:else}
    <strong>Here there will be tags...</strong>
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
</style>
