<script>
  import { navigate } from "svelte-routing";
</script>

<footer>
  <button class="version" on:click={() => navigate("/about")}>
    version: {import.meta.env.VITE_VERSION}
  </button>
</footer>

<style>
  footer {
    position: fixed;
    bottom: 0;
    display: flex;
    flex-direction: column;
    font-size: x-small;
    align-items: flex-start;
  }

  .version {
    color: #454545;
  }

  .version {
    border: 0;
    background-color: transparent;
  }
</style>
