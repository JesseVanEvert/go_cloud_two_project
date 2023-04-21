<template>
  <div class="container mt-5">
    <!-- List of Messages -->
    <div class="row">
      <div class="col">
        <caption>
          <h2 style="white-space: nowrap;">List of Messages</h2>
        </caption>
        <div class="d-flex justify-content-center" v-if="messages.length > 0">
          <table class="table table-dark mx-auto">
            <thead>
              <tr>
                <th scope="col">ID</th>
                <th scope="col">Lecturer Email</th>
                <th scope="col">Recipient</th>
                <th scope="col">Message</th>
              </tr>
            </thead>
            <tbody>
              <tr class="table-row" v-for="message in messages" :key="message.messageID">
                <td class="table-cell">{{ message.messageID }}</td>
                <td class="table-cell">{{ message.lecturerEmail }}</td>
                <td class="table-cell">{{ message.to }}</td>
                <td class="table-cell">{{ message.content }}</td>
              </tr>
            </tbody>
          </table>
        </div>
        <div v-else class="no-message">No messages found.</div>
      </div>
    </div>
    </div> 

  <!-- MessageItem -->
  <div class="container mt-5">
    <div class="row">
      <div class="col">
        <caption>
          <h2 style="white-space: nowrap;">Find a Message by ID</h2>
        </caption>
        <form>
          <div class="input-group mb-3">
            <input v-model="messageID" type="text" class="form-control" placeholder="Enter message ID">
          </div>
          <button @click="fetchMessage" class="btn btn-primary">Fetch Message</button>
        </form>
        <table v-if="message" class="table">
          <tbody>
            <tr>
              <td>ID:</td>
              <td>{{ message.messageID }}</td>
            </tr>
            <tr>
              <td>Lecturer Email:</td>
              <td>{{ message.lecturerEmail }}</td>
            </tr>
            <tr>
              <td>Recipient:</td>
              <td>{{ message.to }}</td>
            </tr>
            <tr>
              <td>Content:</td>
              <td>{{ message.content }}</td>
            </tr>
          </tbody>
        </table>
        <!-- MessageWithLecturerMail -->
        <caption>
        <h2 style="white-space: nowrap;">Find a Message by LecturerMail</h2>
      </caption>
        <div class="form-group">
          <input type="text" v-bind:value="lecturerEmail" v-on:input="$emit('update:lecturerEmail', $event.target.value)">

        </div>
        <div class="input-group mb-3">
          <button @click="fetchMessagesByLecturerEmail" class="btn btn-primary">Get Messages</button>
        </div>
        
        <table v-if="messages.length" class="table">
          <thead class="thead-dark">
            <tr>
              <th>Message ID</th>
              <th>Lecturer Email</th>
              <th>To Email</th>
              <th>Content</th>
            </tr>
          </thead>
          
        </table>
      </div>
    </div>
    </div>

    
 
</template>
  
<script>
import axios from 'axios';
export default {
  data() {
    return {
      lecturerEmail: "",
      messageID: '',
      message: JSON.parse(localStorage.getItem('message')) || null,
      error: null,
      messages: [{
    "messageID": 1,
    "lecturerEmail": "john@example.com",
    "to": "jane@example.com",
    "content": "Hi Jane, can we schedule a meeting for next week?"
  },
  {
    "messageID": 2,
    "lecturerEmail": "jane@example.com",
    "to": "john@example.com",
    "content": "Sure, how about Tuesday at 2 PM?"
  },
  {
    "messageID": 3,
    "lecturerEmail": "jane@example.com",
    "to": "james@example.com",
    "content": "Hi James, I wanted to follow up on our conversation from last week."
  }
  ,
  {
    "messageID": 2,
    "lecturerEmail": "jane@example.com",
    "to": "john@example.com",
    "content": "Sure, how about Tuesday at 2 PM?"
  },
  
  {
    "messageID": 2,
    "lecturerEmail": "jane@example.com",
    "to": "john@example.com",
    "content": "Sure, how about Tuesday at 2 PM?"
  },
  
  {
    "messageID": 2,
    "lecturerEmail": "jane@example.com",
    "to": "john@example.com",
    "content": "Sure, how about Tuesday at 2 PM?"
  },
  
  {
    "messageID": 2,
    "lecturerEmail": "jane@example.com",
    "to": "john@example.com",
    "content": "Sure, how about Tuesday at 2 PM?"
  },
  
  {
    "messageID": 2,
    "lecturerEmail": "jane@example.com",
    "to": "john@example.com",
    "content": "Sure, how about Tuesday at 2 PM?"
  },
  {
    "messageID": 2,
    "lecturerEmail": "jane@example.com",
    "to": "john@example.com",
    "content": "Sure, how about Tuesday at 2 PM?"
  },
  {
    "messageID": 2,
    "lecturerEmail": "jane@example.com",
    "to": "john@example.com",
    "content": "Sure, how about Tuesday at 2 PM?"
  },
  {
    "messageID": 2,
    "lecturerEmail": "jane@example.com",
    "to": "john@example.com",
    "content": "Sure, how about Tuesday at 2 PM?"
  },
  {
    "messageID": 2,
    "lecturerEmail": "jane@example.com",
    "to": "john@example.com",
    "content": "Sure, how about Tuesday at 2 PM?"
  },
  {
    "messageID": 2,
    "lecturerEmail": "jane@example.com",
    "to": "john@example.com",
    "content": "Sure, how about Tuesday at 2 PM?"
  },
]
    };
  },
  mounted() {
    this.fetchMessages();
    setInterval(() => {
      this.fetchMessages();
    }, 5000); // reload every 5 seconds
  },
  created() {
    axios
      .get('/api/message/lecturer/' + this.lecturerEmail)
      .then(response => {
        this.messages = response.data;
      })
      .catch(error => {
        console.error(error);
      });
  },
  props: ['lecturerEmail'],

  methods: {
    fetchMessages() {
      axios
        .get("http://localhost:8080/messages")
        .then(response => {
          this.messages = response.data;
        })
        .catch(error => {
          console.error(error);
        });
    },
    async fetchMessage() {
      try {
        const response = await axios.get(`http://localhost:8080/messages/${this.messageID}`);
        this.message = response.data;
        localStorage.setItem('message', JSON.stringify(this.message));
      } catch (err) {
        this.error = 'Error fetching message';
      }
    },
    async fetchMessagesByLecturerEmail() {
      try {
        const response = await axios.get(`http://localhost:8080/messages/lecturer/${this.lecturerEmail}`);
        this.messages = response.data;
      } catch (error) {
        console.error(error);
      }
    }
  }
};
</script>