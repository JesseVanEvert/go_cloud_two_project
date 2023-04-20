<template>
  <form @submit.prevent="sendMessage">
    <div>
      <label for="sender">From:</label>
      <input type="text" id="sender" v-model="sender">
    </div>
    <div>
      <label for="recipients">To:</label>
      <input type="text" id="recipients" v-model="recipients">
    </div>
    <div>
      <label for="message">Message:</label>
      <textarea id="message" v-model="message"></textarea>
    </div>
    <button type="submit">Send Message</button>
  </form>
</template>

  
  <script>
import axios from 'axios'

export default {
  data() {
    return {
      sender: '',
      recipients: '',
      message: ''
    }
  },
  methods: {
    sendMessage() {
      const payload = {
        action: 'message',
        message: {
          from: this.sender,
          to: this.recipients.split(','),
          message: this.message
        }
      }
      axios.post('/send-message', payload)
        .then(response => {
          alert(response.data.message)
        })
        .catch(error => {
          console.error(error)
          alert('Error sending message')
        })
    }
  }
}
</script>

  