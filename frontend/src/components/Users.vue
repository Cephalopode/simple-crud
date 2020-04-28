<template>
  <v-container class="fill-height" fluid style=" max-width:800px">
    <v-row align="center" justify="center">
      <v-col class="text-center">
        <v-data-table :headers="headers" :items="users" :items-per-page="5" class="elevation-1 category-table">
          <template v-slot:top>
            <v-toolbar flat color="white">
              <v-icon large>mdi-account</v-icon>
              <h1>User List</h1>
              <div class="flex-grow-1"></div>

              <v-dialog v-model="dialog" max-width="500px">
                <template v-slot:activator="{ on }">
                  <v-btn color="success" dark class="mb-2" v-on="on">Create user</v-btn>
                </template>
                <v-card>
                  <v-card-title>
                    <span style="margin-bottom: 40px;" class="headline">{{ formTitle }}</span>
                  </v-card-title>

                  <v-card-text v-if="!dialogDelete">
                    <v-container>
                      <v-row class="justify-center">
                        <v-col cols="12" sm="6" md="4">
                          <v-text-field v-model="editedItem.name" label="Name:" outlined></v-text-field>
                          <v-number-field v-model="editedItem.age" label="Age:" outlined />
                        </v-col>
                      </v-row>
                    </v-container>
                  </v-card-text>

                  <v-card-actions>
                    <div class="flex-grow-1"></div>
                    <div v-if="dialogDelete">
                      <v-btn color="blue darken-1" text @click="close">Cancel</v-btn>
                      <v-btn color="blue darken-1" text @click="remove">Delete</v-btn>
                    </div>
                    <div v-else>
                      <v-btn color="blue darken-1" text @click="close">Cancel</v-btn>
                      <v-btn color="blue darken-1" text @click="save">Save</v-btn>
                    </div>
                  </v-card-actions>
                </v-card>
              </v-dialog>
            </v-toolbar>
          </template>

          <template v-slot:item.action="{ item }">
            <v-icon medium color="success" class="mr-2" @click="editItem(item)">mdi-pencil</v-icon>
            <v-icon medium color="success" @click="deleteItem(item)">mdi-delete</v-icon>
          </template>

          <template v-slot:no-data>
            <v-btn color="primary" @click="load">Reset</v-btn>
          </template>
        </v-data-table>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import axios from "axios"
export default {
  name: "Users",
  data: () => ({
    dialogDelete: false,
    dialog: false,
    headers: [
      { text: "Id", value: "id", width: "10%" },
      { text: "Name", align: "center", sortable: true, value: "name" },
      { text: "Age", align: "center", sortable: true, value: "age" },
      { text: "Actions", value: "action", sortable: false, width: "15%" }
    ],
    users: [],
    editedItem: {
      id: 0,
      name: "",
      age: 0
    },
    defaultItem: {
      id: 0,
      name: "",
      age: 0
    }
  }),

  computed: {
    formTitle() {
      let typeName = " User"
      if (this.dialogDelete) {
        return "Delete" + typeName
      } else if (this.editedItem.id === 0) {
        return "New" + typeName
      } else {
        return "Edit" + typeName
      }
    }
  },

  watch: {
    dialog(val) {
      val || this.close()
    }
  },

  created() {
    this.load()
  },

  methods: {
    load() {
      axios.get("/api/users").then(ret => {
        this.users = ret.data
      })
    },

    editItem(item) {
      this.editedItem = Object.assign({}, item)
      this.dialogDelete = false
      this.dialog = true
    },

    deleteItem(item) {
      this.editedItem = Object.assign({}, item)
      this.dialogDelete = true
      this.dialog = true
    },

    close() {
      this.dialog = false
      setTimeout(() => {
        this.editedItem = Object.assign({}, this.defaultItem)
        this.dialogDelete = false
      }, 300)
    },

    save() {
      let item = this.editedItem
      axios.put("/api/users", item).then(ret => {
        this.load()
        this.close()
      })
    },

    remove() {
      axios
        .delete("/api/users", { data: { id: this.editedItem.id } })
        .then(ret => {
          this.load()
          this.close()
        })
    }
  }
}
</script>

