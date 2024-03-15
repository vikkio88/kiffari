<script>
  import { navigate } from "svelte-routing";
  import Spinner from "../components/shared/Spinner.svelte";

  import { KATO_API_URL, LOGIN_TOKEN_KEY } from "../const";
  import { userToken } from "../store";
  import { onMount } from "svelte";
  import ErrorToast from "../components/shared/ErrorToast.svelte";
  import Version from "../components/shared/Version.svelte";
  import { appConfig } from "../store";

  let error = null;
  let loginPromise = null;
  let passwordInput;
  let value;

  function login() {
    fetch(`${KATO_API_URL}/login`, {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ passkey: value }),
    })
      .then((resp) => {
        if (resp.status != 200) {
          throw Error(`${resp.status}`);
        }
        return resp.json();
      })
      .then(({ token }) => {
        $userToken = token;
        window.localStorage.setItem(LOGIN_TOKEN_KEY, token);
        navigate("/", { replace: true });
      })
      .catch((err) => {
        value = "";
        error = `Could not login(#'${err}')`;
      });
  }

  onMount(() => {
    if (Boolean(window.localStorage.getItem(LOGIN_TOKEN_KEY))) {
      navigate("/", { replace: true });
    }
  });
</script>

{#if $appConfig.kiffari}
  <h1 title="Kiffari & Kato">🎫 🪣</h1>
{:else}
  <h1>Kato 🪣</h1>
{/if}

{#if !loginPromise}
  <form class="wrapper" on:submit|preventDefault|stopPropagation={login}>
    <input
      bind:this={passwordInput}
      type="password"
      placeholder="Insert your Passkey..."
      bind:value
    />
    <button disabled={!Boolean(value) || value.length < 6}>Login</button>
  </form>
{:else}
  <Spinner />
{/if}

<h3><Version /></h3>
<ErrorToast
  {error}
  onDismiss={() => {
    error = null;
    passwordInput.focus();
  }}
/>

<style>
  .wrapper {
    margin-top: 3rem;
    display: flex;
    flex-direction: column;
    align-items: center;
  }
  .wrapper > input {
    padding: var(--input-default-pad);
    font-size: var(--input-font-size);
    border-radius: var(--default-borders-radius);
    margin-bottom: 1rem;
  }

  .wrapper > button {
    font-size: var(--input-medium-font-size);
    padding: var(--input-default-pad);
  }

  h1 {
    margin: 0;
    padding: 1rem;
  }
</style>
