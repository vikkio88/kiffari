<script>
  let version = import.meta.env.VITE_VERSION;
</script>

<div class="wrapper">
  <img src="/favicon.svg" alt="A Kato rolling around" />
  <div>
    made with ♥ by <a href="https://vikkio.me" target="_blank">vikkio</a>
  </div>

  <div>
    Source available on <a
      href="https://github.com/vikkio88/kiffari/tree/main/kato"
      target="_blank"
      rel="noopener noreferrer">Github</a
    >
  </div>

  <h3>
    Version: <a
      href={version === "dev"
        ? "/"
        : `https://github.com/vikkio88/kiffari/commit/${version}`}
      target="_blank"
      rel="noopener noreferrer">{version}</a
    >
  </h3>
</div>

<style>
  .wrapper {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    min-height: 80vh;
  }

  .wrapper img {
    width: 50px;
    animation: jumping 1s ease-out infinite;
    margin-bottom: 3rem;
  }

  @keyframes jumping {
    0% {
        transform: translateY(10%) rotate(0);
    }

    25% {
        transform: translateY(0%) rotate(0);

    }

    50% {
        transform: translateY(-50%) rotate(0);

    }

    70% {
        transform: translateY(-55%) rotate(-10deg);
    }

    70% {
        transform: translateY(-55%) rotate(10deg);
    }

    80% {
        transform: translateY(-50%) rotate(0);
    }

    100% {
        transform: translateY(0%) rotate(0);
    }
  }
</style>
