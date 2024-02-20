<script>
  import { extractYTVideos } from "../../libs/renderers/extractors";

  export let body = "";
  let videoSrcs = extractYTVideos(body);
</script>

{#if videoSrcs.length < 1}
  <h3>No youtube video in the body</h3>
{:else}
  {#each videoSrcs as v}
    <div class="videoContainer">
      <iframe frameborder="0" title="YT video" src={v.embed} width="100%" />
    </div>
  {/each}
{/if}

<style>
  .videoContainer {
    position: relative;
    padding-bottom: 56.25%;
    margin-bottom: 2rem;
  }

  .videoContainer iframe {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
  }
</style>
