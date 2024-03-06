<script>
  export let selected = null;
  export let options = [];
</script>

<select on:change={(e) => (selected = e.currentTarget.value)}>
  {#each options as option}
    <option
      selected={(option.value ?? option) === selected}
      value={option.value ?? option}
    >
      {option.label ?? option}
    </option>
  {/each}
</select>


<style>
    select {
        padding: .8rem;
    }
</style>