<script>
  export let text = "";
  // TODO: use this component elsewhere
</script>

<div>
  <h3>{text}</h3>
</div>

<style>
  h3 {
    margin: 0;
    padding: 0 1rem;
    display: inline-block;
    transform: translateY(-50%);
    background-color: var(--main-bg-color);
    color: var(--x-light-gray);
  }
  div {
    border-top: 2px var(--x-light-gray) solid;
  }
</style>
