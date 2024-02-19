<script>
  import { navigate } from "svelte-routing";
</script>

<button class="version" on:click={() => navigate("/about")}>
  version: {import.meta.env.VITE_VERSION}
</button>

<style>
  .version {
    color: #454545;
  }

  .version {
    border: 0;
    background-color: transparent;
  }
</style>
