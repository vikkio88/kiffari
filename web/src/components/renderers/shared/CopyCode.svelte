<script>
  import { hasCodeSnippets } from "../../../libs/renderers/helpers";
  import Toggle from "../../shared/Toggle.svelte";
  export let body = "";
  let hasCode = hasCodeSnippets(body);

  function copyText(e) {
    navigator.clipboard.writeText(e.target.innerText);
    const span = document.createElement("span");
    span.classList.add("tip");
    span.innerText = "📋 Copied!";
    e.target.parentElement.appendChild(span);
    setTimeout(() => {
      e.target.parentElement.removeChild(span);
    }, 800);
  }

  function handleCopyPaste(e) {
    const { checked } = e.currentTarget;
    const snippets = document.querySelectorAll("code");
    if (checked) {
      snippets.forEach((s) => {
        s.classList.add("light-bg");
        s.parentElement.classList.add("p-r");
        s.parentElement.classList.add("crs-pointer");
        s.addEventListener("click", copyText);
      });
    } else {
      snippets.forEach((s) => {
        s.classList.remove("light-bg");
        s.parentElement.classList.remove("p-r");
        s.parentElement.classList.remove("crs-pointer");
        s.removeEventListener("click", copyText);
      });
    }
  }
  let checked = false;
</script>

<div class="hidden row" class:hasCode>
  <Toggle name="copypaste" onChange={handleCopyPaste} />
  <label for="copypaste">📋 Copy <i>Code</i></label>
</div>

<style>
  .hasCode {
    visibility: visible;
  }

  .row {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    gap: .5rem;
  }
</style>
