import React, { useState, useEffect, Fragment } from 'react'
import { Link } from 'react-router-dom'

export default function CoursesFunc(props) {
    const [courses, setCourses] = useState([])
    const [error, setError] = useState(null)

    useEffect(() => {
        fetch("http://localhost:4000/courses")
            .then((response) => {

                if (response.status !== 200) {
                    let err = Error;
                    err.message = "Invalid response code : " + response.status
                    setError(err)
                } else {
                    return response.json()
                }
            })
            .then((json) => {
                setCourses(json.courses)
            })
    }, [])

    if (error != null) {
        return (<div>{error.message}</div>)
    } else {
        return (
            <Fragment>
                <h2> All Courses</h2>
                <div className='float-end' >
                    <a href='../savecourse/newcourse' 
                        className='btn btn-primary ms-1'>Add New Course</a>
                </div>
                <div className='list-group'>
                    {courses.map((c, i) => (
                        <Link key={i} className='list-group-item list-group-item-item' to={`savecourse/${c["Course Name"]}`}>{c["Course Name"]}</Link>
                    ))}
                </div>
            </Fragment>
        )
    }
}

