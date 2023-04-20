<template>
    

    <div class="container">
        <caption>
            <h2 style="white-space: nowrap;">Add a Lecturer?</h2>
        </caption>
        <div class="add-form">
            <form @submit.prevent="addLecturer">
                <input type="text" v-model="newLecturer.first_name" placeholder="First Name">
                <input type="text" v-model="newLecturer.last_name" placeholder="Last Name">
                <input type="text" v-model="newLecturer.email" placeholder="Email">
                <select v-model="newLecturer.classrooms">
                    <option v-for="c in classrooms" :value="c.id">{{ c.classname }}</option>
                </select>
                <button type="submit" class="btn btn-primary">Submit</button>
            </form>
            <br>
            <caption>
                <h2 style="white-space: nowrap;">Find a Lecturer by ID</h2>
            </caption>
            <div class="input-group mb-3">
                <input type="number" v-model="lecturerId" placeholder="Enter Lecturer ID" class="form-control">
                <div class="input-group-append">
                    <button @click.prevent="getLecturerById(lecturerId)" class="btn btn-primary">Get Lecturer</button>
                </div>
            </div>
            <br>
            <h2>Add Lecturer to Class</h2>
            <form @submit.prevent="addLecturer">
                <label for="lecturerID">Lecturer ID:</label>
                <input type="number" id="lecturerID" v-model="lecturerID">
                <label for="classID">Class ID:</label>
                <input type="number" id="classID" v-model="classID">
                 <button type="submit" class="btn btn-primary">Add lecturer</button>
            </form>
            <div v-if="message">{{ message }}</div>
        </div>
    </div>
    <div class="container">
        <h2 style="white-space: nowrap;">List of Lecturers</h2>
        <div class="d-flex justify-content-center">
            <table class="table table-dark mx-auto">
                <thead>
                    <tr>
                        <th scope="col">Lecturer ID</th>
                        <th scope="col">Lecturer First Name</th>
                        <th scope="col">Lecturer Last Name</th>
                        <th scope="col">Lecturer Email</th>
                        <th scope="col">Class</th>

                    </tr>
                </thead>
                <tbody>
                    <tr class="table-row" v-for="lects in lecturer">
                        <td class="table-cell">{{ lects.id }}</td>
                        <td class="table-cell">{{ lects.first_name }}</td>
                        <td class="table-cell">{{ lects.last_name }}</td>
                        <td class="table-cell">{{ lects.email }}</td>
                        <td>{{ lects.classroom.join(', ') }}</td>
                        <td class="table-cell">
                            <button class="btn btn-danger" @click="deleteLecturer(lects.id)">Delete</button>
                        </td>
                    </tr>
                </tbody>
            </table>

        </div>
    </div>
    
</template>
<style scoped>
.delete-button {
    background-color: #f44336;
    color: white;
    border: none;
    border-radius: 5px;
    padding: 8px 16px;
    font-size: 14px;
    cursor: pointer;
}
</style>
<script>
import axios from 'axios';
export default {
    data() {
        return {
            lecturer: [],
            newLecturer: {
                first_name: '',
                last_name: '',
                email: '',
                classroom: ''
            },
            classrooms: []
        }
    },
    mounted() {
        this.fetchLecturers();
        setInterval(() => {
            this.fetchLecturers();
        }, 5000); // reload every 5 seconds
    },
    methods: {

        fetchLecturers() {
            axios
                .get("http://localhost:8000/api/lecturer")
                .then(response => {
                    this.lecturer = response.data;
                })
                .catch(error => {
                    console.error(error);
                });
        },
        getStudentById(id) {
            axios
                .get(`http://localhost:8000/api/lecturer/${id}`)
                .then(response => {
                    this.lecturer = response.data;
                })
                .catch(error => {
                    console.error(error);
                });
        },
        addLecturer() {
            axios
                .post("http://localhost:8000/api/lecturer", this.newLecturer)
                .then(response => {
                    this.lecturer.push(response.data);
                    this.newLecturer = { first_name: '', last_name: '', email: '', classroom: '' };
                })
                .catch(error => {
                    console.error(error);
                });

        },
        deleteLecturer(first_name) {
            axios
                .delete("http://localhost:8000/api/lecturer/" + first_name)
                .then(response => {
                    console.log(response);
                    this.fetchLecturers();
                })
                .catch(error => {
                    console.error(error);
                });
        },
        async addLecturer() {
            try {
                const response = await axios.post('/api/add-lecturer-to-class', {
                    lecturerID: this.lecturerID,
                    classID: this.classID
                });
                this.message = response.data;
            } catch (error) {
                this.message = error.message;
            }
        }

    }
}
</script>