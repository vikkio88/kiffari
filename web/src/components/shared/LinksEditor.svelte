<script>
  export let links = [];
  function remove(link) {
    links = links.filter(
      (l) => !(l.label === link.label && link.href === l.href)
    );
  }

  $: hasEmpty =
    Array.isArray(links) && links.some((l) => l.href === "" || l.label === "");
</script>

{#if Array.isArray(links)}
  {#each links as link}
    <div>
      🔗
      <input
        type="text"
        required
        bind:value={link.label}
        placeholder="Link Label..."
      />
      <input
        type="text"
        required
        pattern="/(https?:\/\/\S+)/"
        bind:value={link.href}
        placeholder="http://some.website/"
      />
      <button type="button" class="smaller" on:click={() => remove(link)}>
        ❌
      </button>
    </div>
  {/each}
{/if}
<button
  type="button"
  disabled={hasEmpty}
  on:click={() => {
    links = [...(links ?? []), { label: "", href: "" }];
  }}
>
  Add 🔗
</button>

<style>
  div {
    display: flex;
    flex-direction: row;
    gap: 0.5rem;
  }

  input {
    padding: 0.5rem;
  }
</style>
