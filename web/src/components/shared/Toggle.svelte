<script>
  export let checked = false;
</script>

<input type="checkbox" style="--s:15px" bind:checked />

<style>
  input {
    --s: 50px;
    height: calc(var(--s) + var(--s) / 5);
    width: auto;
    aspect-ratio: 2.25;
    border-radius: var(--s);
    margin: calc(var(--s) / 2);
    display: grid;
    cursor: pointer;
    background-color: var(--danger-color);
    box-sizing: content-box;
    overflow: hidden;
    transition: 0.3s 0.1s;
    -webkit-appearance: none;
    -moz-appearance: none;
    appearance: none;
  }
  input:before {
    content: "";
    padding: calc(var(--s) / 10);
    --_g: radial-gradient(
      circle closest-side at calc(100% - var(--s) / 2) 50%,
      #333 96%,
      #3333
    );
    background:
      var(--_g) 0 / var(--_p, var(--s)) 100% no-repeat content-box,
      #fff;
    mix-blend-mode: darken;
    filter: blur(calc(var(--s) / 12)) contrast(10);
    transition:
      0.4s,
      background-position 0.4s 0.1s,
      padding
        cubic-bezier(0, 1, 1, 2)
        0.25s 0.1s;
  }
  input:checked {
    background-color: var(--success-color);
  }
  input:checked:before {
    padding: calc(var(--s) / 10 + 0.05px) calc(var(--s) / 10);
    --_p: 100%;
    --_i: 1;
  }
</style>
