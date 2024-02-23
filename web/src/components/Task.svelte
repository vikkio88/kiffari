<script>
  export let task = null;
</script>

<div class="wrapper">
  <h2>{task.title}</h2>
</div>

<style>
  h2 {
    font-size: 1.3rem;
  }

  .wrapper {
    padding: 1rem 2rem;
    border: 2px white solid;
    border-radius: 10px;
    width: 60%;
    margin: 1rem 0 0 0;
  }
</style>
