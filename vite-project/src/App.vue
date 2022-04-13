<script setup>
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
import { ref } from "vue"
import CourseItem from "./components/AvailableCourse.vue"
import TakenCourse from "./components/TakenCourse.vue"
// import { Addition, Substitution, GetCourses } from "./components/axios.js"
// import axios from "./components/axios.js"
const StudentID = ref("")
const TotalCourses = ref([
    { id: "1", course: "lmaoo", units: 4 },
    { id: "2", course: "lmaoo", units: 4 },
    { id: "3", course: "lmaoo", units: 4 },
    { id: "4", course: "lmaoo", units: 4 },
    { id: "5", course: "lmaoo", units: 4 },
    { id: "6", course: "lmaoo", units: 4 },
])

const TakenCourses = ref([
    { id: "11", course: "lmaoo", units: 4 },
    { id: "12", course: "lmaoo", units: 4 },
    { id: "13", course: "lmaoo", units: 4 },
    { id: "14", course: "lmaoo", units: 4 },
])

const TotalUnits = ref(0)
const oldId = ref(-1)

const onSubstitute = (oldId, newID) => {
    const oldIndex = TotalCourses.value.findIndex(
        (course) => course.id === oldId
    )
    const newIndex = TotalCourses.value.findIndex(
        (course) => course.id === newID
    )
    const oldCourse = TotalCourses.value[oldIndex]
    const newCourse = TotalCourses.value[newIndex]
    TotalCourses.value[oldIndex] = newCourse
    TotalCourses.value[newIndex] = oldCourse

    TotalUnits.value -= oldCourse.units
    TotalUnits.value += newCourse.units

    oldId.value = -1
}

const onTakeCourse = (id) => {
    // check if oldId and newID are not -1
    GetSeats(id).then(seats => {

    if (if seats < 100 && oldId.value !== -1) {
        Substitution(oldId.value, id, StudentID.value).then((res) => {
            if (res.data.success) {
                onSubstitute(oldId.value, id)
            }
        })
    } else {
        Addition(id, StudentID.value).then((res) => {
            if (res.data.success) {
                let Course = TotalCourses.value.filter((x) => x.id === id)
                TotalCourses.value = TotalCourses.value.filter(
                    (x) => x.id !== id
                )
                TakenCourses.value.push(Course[0])
                TotalUnits.value += Course[0].units
            }
        })
    }
    })
}

const onReplaceCourse = (id) => {
    oldId.value = id
};
</script>

<template>
    <div class="no-scroll gap-4 flex flex-col overflow-auto max-h-full">
        <p class="text-2xl font-semibold mb-6 mt-[20vh]">Available Courses</p>

        <CourseItem
            v-for="course in TotalCourses"
            :course="course.course"
            :units="course.units"
            :id="course.id"
            :choose="onTakeCourse"
        />
    </div>
    <div class="no-scroll gap-4 flex flex-col overflow-auto max-h-full">
        <p class="text-2xl font-semibold mb-6 mt-[20vh]">
            Preference Order {{ ` (Units used: ${TotalUnits.value})` }}
        </p>
        <TakenCourse
            v-for="course in TakenCourses"
            :course="course.course"
            :id="course.id"
            :replace="onReplaceCourse"
            :units="course.units"
            :selected="oldId.value === course.id"
        />
    </div>
</template>

<style>
@import url("https://fonts.googleapis.com/css2?family=Montserrat&display=swap");
@tailwind base;
@tailwind components;
@tailwind utilities;

#app {
    font-family: "Montserrat", sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    display: flex;
    flex-direction: row;
    justify-content: space-evenly;
    align-items: center;
    height: 100vh;
    width: 100vw;
}

.no-scroll {
    -ms-overflow-style: none; /* IE and Edge */
    scrollbar-width: none; /* Firefox */
}

.no-scroll::-webkit-scrollbar {
    display: none;
}
</style>
