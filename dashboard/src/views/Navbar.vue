<template>
    <nav class="navbar navbar-expand-lg navbar-light bg-light">
        <div class="container-fluid">
            <a class="navbar-brand">
                <img alt="Gaul logo"
                    src="../assets/gaul.webp"
                    class="align-self-center"
                    style="max-width: 55px">
                Ground Station |
             </a>
            <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse" id="navbarNav">
            <ul class="navbar-nav">
                <li class="nav-item" v-for="(link, idx) in links" :key="idx" cols="1">
                <router-link :to="link.path" class="nav-link active" aria-current="page">{{ link.name }}</router-link>
                </li>
            </ul>
            </div>
        </div>
    </nav>
</template>

<script>
export default {
  name: "Navbar",
data() {
    return {
      links: [
        { path: '/', name: 'Monitor' },
        { path: '/tracking', name: 'Tracking' }
      ]
    }
  }
}
</script>
