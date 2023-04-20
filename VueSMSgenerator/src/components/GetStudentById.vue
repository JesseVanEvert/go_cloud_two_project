<template>
  <div class="container mt-5">
    <div class="row">
      <div class="col">
        <caption>
          <h2 style="white-space: nowrap;">Find a Student by ID</h2>
        </caption>
        <div class="input-group mb-3">
          <input type="number" v-model="studentId1" placeholder="Enter Student ID" class="form-control">
          <div class="input-group-append">
            <button @click.prevent="getStudentById(studentId1)" class="btn btn-primary">Get Student</button>
          </div>
        </div>
        <div class="container">
          <div v-if="student" class="row">
            <div class="col">
              <h2>{{ student.fname }} {{ student.lname }}</h2>
              <p>Email: {{ student.email }}</p>
              <p>Classroom ID: {{ student.classroom_id }}</p>
            </div>
          </div>
        </div>
        <caption>
          <h2 style="white-space: nowrap;">Add a Student</h2>
        </caption>
        <div class="add-form">
          <form @submit.prevent="addStudent">
            <input type="text" v-model="newStudent.fname" placeholder="First Name" class="form-control mb-3">
            <input type="text" v-model="newStudent.lname" placeholder="Last Name" class="form-control mb-3">
            <input type="text" v-model="newStudent.email" placeholder="Email" class="form-control mb-3">
            <input type="number" v-model="newStudent.classroom_id" placeholder="Classroom ID" class="form-control mb-3">
            <button type="submit" class="btn btn-primary">Submit</button>
          </form>
        </div>
        <caption>
          <h2 style="white-space: nowrap;">Delete a Student</h2>
        </caption>
        <div class="add-form">
          <form @submit.prevent="deleteStudent">
            <input type="number" v-model="studentId" placeholder="Student ID" class="form-control mb-3">
            <button type="submit" class="btn btn-primary">Submit</button>
          </form>
        </div>


      </div>
    </div>
  </div>





  <div class="container mt-5">
    <div class="row">
      <div class="col">
        <caption>
          <h2 style="white-space: nowrap;">List of Students</h2>
        </caption>
        <div class="d-flex justify-content-center">
          <table class="table table-dark mx-auto">
            <thead>
              <tr>
                <th scope="col">First Name</th>
                <th scope="col">Last Name</th>
                <th scope="col">Email</th>
                <th scope="col">Classroom ID</th>
              </tr>
            </thead>
            <tbody>
              <tr class="table-row" v-for="student in students">
                <td class="table-cell">{{ student.fname }}</td>
                <td class="table-cell">{{ student.lname }}</td>
                <td class="table-cell">{{ student.email }}</td>
                <td class="table-cell">{{ student.classroom_id }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

  
<script>
import axios from 'axios';

export default {
  data() {
    return {
      studentId: '',
      student: null,
      students: [{
        "studentId": 1,
        "fname": "John",
        "lname": "Doe",
        "email": "john.doe@example.com",
        "classroom_id": 101
      },
      {
        "studentId": 2,
        "fname": "Jane",
        "lname": "Doe",
        "email": "jane.doe@example.com",
        "classroom_id": 102
      },
      {
        "studentId": 1,
        "fname": "John",
        "lname": "Doe",
        "email": "john.doe@example.com",
        "classroom_id": 101
      },
      {
        "studentId": 1,
        "fname": "John",
        "lname": "Doe",
        "email": "john.doe@example.com",
        "classroom_id": 101
      },
      {
        "studentId": 1,
        "fname": "John",
        "lname": "Doe",
        "email": "john.doe@example.com",
        "classroom_id": 101
      },
      {
        "studentId": 1,
        "fname": "John",
        "lname": "Doe",
        "email": "john.doe@example.com",
        "classroom_id": 101
      },
      {
        "studentId": 1,
        "fname": "John",
        "lname": "Doe",
        "email": "john.doe@example.com",
        "classroom_id": 101
      },
      {
        "studentId": 1,
        "fname": "John",
        "lname": "Doe",
        "email": "john.doe@example.com",
        "classroom_id": 101
      },
      {
        "studentId": 1,
        "fname": "John",
        "lname": "Doe",
        "email": "john.doe@example.com",
        "classroom_id": 101
      },
      {
        "studentId": 1,
        "fname": "John",
        "lname": "Doe",
        "email": "john.doe@example.com",
        "classroom_id": 101
      },
      {
        "studentId": 1,
        "fname": "John",
        "lname": "Doe",
        "email": "john.doe@example.com",
        "classroom_id": 101
      },],
      newStudent: {
        classroom_id: '',
        email: '',
        fname: '',
        lname: '',
      }
    };
  },
  mounted() {
    this.fetchStudents();
    setInterval(() => {
      this.fetchStudents();
    }, 5000); // reload every 5 seconds
  },
  methods: {
    getStudentById(id) {
      axios
        .get(`http://localhost:8000/api/student/${id}`)
        .then(response => {
          this.student = response.data;
        })
        .catch(error => {
          console.error(error);
        });
    },
    addStudent() {
      axios
        .post("http://localhost:8000/api/student", this.newStudent)
        .then(response => {
          this.students.push(response.data);
          this.newStudent = { classroom_id: '', email: '', fname: '', lname: '' };
        })
        .catch(error => {
          console.error(error);
        });
    },
    fetchStudents() {
      axios
        .get("http://localhost:8000/api/student")
        .then(response => {
          this.students = response.data;
        })
        .catch(error => {
          console.error(error);
        });
    },
    deleteStudent() {
      axios
        .put(`http://localhost:8000/api/student/delete/${this.studentId}`)
        .then(response => {
          console.log(response);
        })
        .catch(error => {
          console.error(error);
        });
    },
  }
};
</script>
  