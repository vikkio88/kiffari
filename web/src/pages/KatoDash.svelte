<script>
  import Adder from "../components/shared/Adder.svelte";
  import { navigate } from "svelte-routing";
  import { createNote, getAllNotes, getDashNotes } from "../libs/api";
  import AllNotes from "../components/notes/AllNotes.svelte";
  import Controls from "../components/shared/Controls.svelte";
  import Footer from "../components/Footer.svelte";
  import { protectedRoute } from "../libs/routes";
  import { nowString } from "../libs/dates";
  import Spinner from "../components/shared/Spinner.svelte";
  import BtnGroup from "../components/shared/BtnGroup.svelte";
  import DashNotes from "../components/notes/DashNotes.svelte";
  protectedRoute();

  let dashNotesPromise = getDashNotes();
  // let dashNotesPromise = getAllNotes();

  function create() {
    navigate("/create-note");
  }

  async function onCreate(bodyWrapper) {
    const resp = await createNote({
      title: nowString(),
      tags: [],
      ...bodyWrapper,
    });
    if (resp.ok) {
      dashNotesPromise = getDashNotes();
      return;
    }

    //TODO: handle error
  }

  const TABS = {
    DASH: "dash",
    ALL_NOTES: "all_notes",
  };
  let activeTab = TABS.DASH;
  function setActiveTab(tab) {
    dashNotesPromise = tab === TABS.ALL_NOTES ? getAllNotes() : getDashNotes();
    activeTab = tab;
  }
</script>

<Adder
  centered
  placeholder="Note body..."
  on:added={({ detail: body }) => onCreate({ body })}
/>
<div class="wrapper">
  <BtnGroup>
    <button
      class:active={activeTab === TABS.DASH}
      on:click={() => setActiveTab(TABS.DASH)}>Dash</button
    >
    <button
      class:active={activeTab === TABS.ALL_NOTES}
      on:click={() => setActiveTab(TABS.ALL_NOTES)}>All Notes</button
    >
  </BtnGroup>
  {#await dashNotesPromise}
    <Spinner />
  {:then notes}
    {#if activeTab === TABS.DASH}
      <DashNotes {notes} />
    {:else}
      <div class="subwrapper">
        <AllNotes {notes} />
      </div>
    {/if}
  {/await}
</div>

<Footer />

<Controls>
  <button class="big-control" title="New Note" on:click={create}
    >Add Note 📝</button
  >
</Controls>

<style>
  .wrapper {
    display: grid;
    grid-template-rows: auto auto;
    padding-bottom: 2rem;
    margin-top: 1rem;
  }

  .subwrapper {
    margin-top: 1rem;
  }
</style>
