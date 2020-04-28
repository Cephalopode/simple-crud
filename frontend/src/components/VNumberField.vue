<template>
  <v-text-field @input="valueChanged" v-bind="$attrs"></v-text-field>
  <!-- <div>eslktjs</div> -->
</template>

<script>
export default {
  name: "v-number-field",
  methods: {
    valueChanged(val) {
      let num = parseInt(val)
      this.$emit("input", isNaN(num) ? 0 : num)
    }
  }
}
</script>