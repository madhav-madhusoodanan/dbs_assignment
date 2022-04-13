const axios = require("axios")

const instance = axios.create({
    baseURL: "https://localhost:8080",
    timeout: 1000,
})

export const Addition = async (studentID, courseID) => {
    const res = await instance.post("/edit/add", {
        student_id: studentID,
        course_id: courseID,
    })
    return res.data.status
}
export const Substitution = async (studentID, oldID, newID) => {
    const res = await instance.post("/edit/substi", {
        id: studentID,
        old_course: oldID,
        new_course: newID,
    })
    return res.data.status
}
export const GetSeats = async (courseID) => {
    const res = await instance.post("/read/course", {
        course_id: courseID,
    })
    return res.data.count
}
export const GetCourses = async (studentID) => {
    const res = await instance.post("/read/student", {
        id: studentID,
    })
    return res.data.data
}
