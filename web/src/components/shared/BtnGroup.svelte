<div class="btnGroup">
  <slot />
</div>

<style global>
  div.btnGroup {
    flex-direction: row;
    align-items: center;
    justify-content: space-around;
  }

  div.btnGroup :global(> button) {
    border-radius: 0px;
    margin: 0;
  }
  div.btnGroup :global(> button.active) {
    background-color: black;
    color: var(--link-color);
  }
  div.btnGroup :global(> button.active:focus),
  div.btnGroup :global(> button.active:focus-visible) {
    outline: none;
  }
</style>
