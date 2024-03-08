<script>
  export let name = "toggle";
  export let checked = false;
  export let onChange = () => {};
</script>

<input {name} type="checkbox" id="switch" bind:checked on:change={onChange} />
<label for="switch">Toggle</label>

<style>
  input[type="checkbox"] {
    height: 0;
    width: 0;
    visibility: hidden;
  }

  label {
    cursor: pointer;
    text-indent: -9999px;
    width: 40px;
    height: 25px;
    background: var(--danger-color);
    display: block;
    border-radius: 100px;
    position: relative;
  }

  label:after {
    content: "";
    position: absolute;
    top: 2.5px;
    left: 2.5px;
    width: 18px;
    height: 19px;
    background: var(--light-gray);
    border-radius: 90px;
    transition: 0.3s;
  }

  input:checked + label {
    background: var(--success-color);
  }

  input:checked + label:after {
    left: calc(100% - 2px);
    transform: translateX(-100%);
  }

  label:active:after {
    width: 50px;
  }
</style>
