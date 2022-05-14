<template>
  <Nav />
  <div class="container mt-5">
    <section class="row">
      <section class="border-end col-md-3 mb-5 me-5">
        <Sidevar />
      </section>
      <section class="col-md-8">
        <router-view />
      </section>
    </section>
  </div>
  <footer class="bg-secondary text-center text-light p-5 mt-5">
    <small>Copyright©2022 conforto All Rights Reserved.</small>
  </footer>
</template>

<script>
import Nav from "@/components/Nav"
import Sidevar from "@/components/Sidevar"
export default {
  components: {
    Nav,
    Sidevar
  }
}
</script>