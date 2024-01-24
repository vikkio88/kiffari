<script>
  import { navigate } from "svelte-routing";
  export let notes = [];
</script>

<ul class="list">
  {#each notes as note}
    <li class="note-item">
      <span class="date">
        {new Date(note.created_at).toLocaleString()}
      </span>
      <h2>{note.title}</h2>
      <button on:click={() => navigate(`/notes/${note.id}`)}>➡️</button>
    </li>
  {/each}
</ul>

<style>
  .list {
    padding: 0 1rem;
    display: flex;
    flex-direction: column;
  }
  .note-item {
    border: 2px white solid;
    border-radius: 10px;
    margin: 1rem 0;
    font-size: 16px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    padding: 1rem 2rem;
  }

  .note-item:hover {
    background-color: #3a3a3a;
  }

  .note-item > h2 {
    /* flex: 1; */
    text-align: center;
  }

  .note-item > button {
    margin-left: auto;
  }

  h2 {
    margin: 0 auto;
    text-align: center;
  }
</style>
