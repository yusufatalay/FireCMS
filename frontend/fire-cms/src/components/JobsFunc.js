import React, { useState, useEffect, Fragment } from 'react'
import { Link } from 'react-router-dom'

export default function JobsFunc(props) {
    const [jobs, setJobs] = useState([])
    const [error, setError] = useState(null)

    useEffect(() => {
        fetch("http://localhost:4000/jobs")
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
                setJobs(json.jobs)
            })
    }, [])

    if (error != null) {
        return (<div>{error.message}</div>)
    } else {
        return (
            <Fragment>
                <h2> All Jobs</h2>
                <div className='float-end' >
                    <a href='../savejobs/newjob' 
                        className='btn btn-primary ms-1'>Add New Job</a>
                </div>
                <div className='list-group'>
                    {jobs.map((j, i) => (
                        <Link key={i} className='list-group-item list-group-item-item' to={`savejob/${j["title"]}`}>{j["title"]}</Link>
                    ))}
                </div>
            </Fragment>
        )
    }
}

