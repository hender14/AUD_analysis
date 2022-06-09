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
  <Footer />
</template>

<script>
import Nav from "@/components/Nav"
import Sidevar from "@/components/Sidevar"
import Footer from "@/components/Footer"
export default {
  components: {
    Nav,
    Sidevar,
    Footer
  }
}
</script>