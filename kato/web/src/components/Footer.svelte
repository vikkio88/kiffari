<script>
  import Version from "./shared/Version.svelte";
</script>

<footer>
  <Version />
</footer>

<style>
  footer {
    position: fixed;
    bottom: 0;
    left: 0;
    display: flex;
    flex-direction: column;
    font-size: x-small;
    align-items: flex-start;
  }
</style>
