const axios = require("axios")

const instance = axios.create({
    baseURL: "https://localhost:8080",
    timeout: 1000,
})

export const Addition = async () => {}
export const Substitution = async () => {}
export const GetCourses = async () => {
    const res = await axios.post("/read/student", {})
    return res.data
}
