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
    background-color: #242424;
    color: #a3a3a3;
  }
  div {
    border-top: 2px #a3a3a3 solid;
  }
</style>
