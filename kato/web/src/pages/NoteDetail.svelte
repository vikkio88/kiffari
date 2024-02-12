<script>
  import TagsList from "../components/TagsList.svelte";
  import Controls from "../components/shared/Controls.svelte";
  import { navigate } from "svelte-routing";
  import ConfirmButton from "../components/shared/ConfirmButton.svelte";
  import BodyRenderer from "../components/BodyRenderer.svelte";
  import { protectedRoute } from "../libs/routes";
  import {
    getNoteDetails,
    deleteNote,
    parseOrThrow,
    catchLogout,
    archiveNote,
    unArchiveNote,
  } from "../libs/api";
  import DateSeverity from "../components/shared/DateSeverity.svelte";
  protectedRoute();

  export let id = "";
  let notePromise = getNoteDetails(id).then(parseOrThrow).catch(catchLogout);

  async function onDelete() {
    const resp = await deleteNote(id);
    if (resp.ok) {
      navigate("/", { replace: true });
    }
  }

  async function onArchiveToggle(isArchived) {
    const resp = await (isArchived ? unArchiveNote(id) : archiveNote(id));
    if (resp.ok) {
      window.location.reload();
    }
  }

  function getDate({ created_at, updated_at }) {
    if (created_at != updated_at) {
      return `updated: ${new Date(updated_at).toLocaleString()}`;
    }

    return new Date(created_at).toLocaleString();
  }
</script>

{#await notePromise then note}
  <div class="note">
    {#if note.archived}
      <h3 class="archived">Archived</h3>
    {/if}
    <h2>{note.title}</h2>
    <div class="dates">
      {getDate(note)}
      <DateSeverity date={note.due_date} />
    </div>
    <div class="body">
      <BodyRenderer body={note.body} />
    </div>
    <div class="tags">
      <h3>Tags</h3>
      <div class="tagList">
        <TagsList tags={note.tags} />
      </div>
    </div>
  </div>
  <Controls>
    <ConfirmButton onConfirmed={() => onArchiveToggle(note.archived)}>
      {#if !note.archived}
        Archive 🗄️
      {:else}
        Un-Archive 🔄
      {/if}
    </ConfirmButton>
    {#if !note.archived}
      <button on:click={() => navigate(`/edit-note/${id}`)}>Edit 📝</button>
    {/if}
    <ConfirmButton onConfirmed={onDelete}>Delete 🗑️</ConfirmButton>
  </Controls>
{/await}

<style>
  .note {
    border-radius: 10px;
    padding: 0 2rem;
  }

  h3.archived {
    color: #a3a3a3;
    border-bottom: dashed 1px #a3a3a3;
  }

  .note h2 {
    margin: 1rem 0 0 0;
  }

  .note div.dates {
    font-size: smaller;
    color: #a3a3a3;
  }

  .body {
    border-radius: 10px;
    text-align: left;
    padding: 3rem;
    margin: 1em 0;
  }

  .tags {
    border-top: 2px #a3a3a3 solid;
  }

  .tagList {
    margin: 1rem 0;
    display: flex;
    flex-direction: row;
    padding-bottom: 0.5rem;
  }

  .tags h3 {
    margin: 0;
    padding: 0 1rem;
    display: inline-block;
    transform: translateY(-50%);
    background-color: #242424;
    color: #a3a3a3;
  }
</style>
