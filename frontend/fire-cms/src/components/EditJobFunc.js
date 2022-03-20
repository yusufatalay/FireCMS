import React, { useState, useEffect, Fragment } from 'react'
import { Link } from 'react-router-dom';
import Alert from './ui-components/Alert';
import { confirmAlert } from 'react-confirm-alert'; // Import
import 'react-confirm-alert/src/react-confirm-alert.css'; // Import css
import Input from './form-components/Input';


export default function EditJobFunc(props) {
    const [job, setJob] = useState({
        "Earning": "",
        "Good At": [],
        "Like": [],
        "Picture": "",
        "title": "",
    })

    const [error, setError] = useState(null)
    const [errors, setErrors] = useState([])
    const [alert, setAlert] = useState({ type: "d-none", message: "" })

    const handleSubmit = (evt) => {
        evt.preventDefault();
        // do client side validation
        let errors = []
        if (job['title'] === "") {
            errors.push("title")
        }
        if (job['Earning'] === "") {
            errors.push("earning")
        }
        if (job['Good At'].length === 0) {
            errors.push("good at")
        }
        if (job['Like'].length === 0) {
            errors.push("like")
        }
        if (job['Picture'] === "") {
            errors.push("picture")
        }
        
        setErrors(errors)
        if (errors.length > 0) {
            return false
        }

        const data = new FormData(evt.target)
        const payload = Object.fromEntries(data.entries())
        const myHeaders = new Headers()
        myHeaders.append("Content-Type", "application/json")
        // handle auth header too
        const requestOptions = {
            method: 'POST',
            headers: myHeaders,
            body: JSON.stringify(payload)
        }

        fetch("http://localhost:4000/savejob", requestOptions)
            .then(response => response.json())
            .then(data => {
                if (data.error) {
                    setAlert({ type: "alert-danger", message: data.error.message })
                } else {
                    props.history.push('/..')
                }
            })
        setJob(payload)
    }

    const handleChange = (evt) => {
        let value = evt.target.value;
        let name = evt.target.name;

        setJob((prevState) => ({
            ...prevState,
            [name]: value,
        }))
    }

    function hasError(key) {
        return errors.indexOf(key) !== -1
    }

    useEffect(() => {
        // also do auth check here

        const jobName = props.match.params.jobName
        if (jobName === "newjob") {
            setJob({
                "Earning": "",
                "Good At": [],
                "Like": [],
                "Picture": "",
                "title": "",
            })
        } else if (jobName !== "") {
            fetch("http://localhost:4000/job/" + jobName)
                .then((response) => {
                    if (response.status !== 200) {
                        let err = Error
                        err.message = "Invalid response code : " + response.status

                        setError(err)
                    } else {
                        setError(null)
                    }
                    return response.json()
                })
                .then((js) => {
                    setJob(js.job)
                })
        }
    }, [props.history, props.match.params.jobName])

    const confirmDelete = (e) => {
        console.error(job)
        confirmAlert({
            title: 'Delete Job ?.',
            message: 'Are you sure to DELETE this job?.',
            buttons: [
                {
                    label: 'Yes',
                    onClick: () => {
                        const myHeaders = new Headers();
                        myHeaders.append("Content-Type", "application/json")
                        // also do security check
                        fetch("http://localhost:4000/deletejob/" + job['title'], { method: "GET", headers: myHeaders })
                            .then((response) => response.json())
                            .then(data => {
                                if (data.error) {
                                    setAlert({ type: "alert-danger", message: data.error.message })
                                } else {
                                    setAlert({ type: "alert-success", message: "job deleted" })

                                    props.history.push(
                                        { pathname: "/", }
                                    )
                                }
                            })
                    }
                },
                {
                    label: 'No',
                    onClick: () => { }
                }
            ]
        });
    }
    if (error) {
        return (<div>{error.message}</div>)
    } else {
        return (
            <Fragment>
                <h2>Save Job</h2>
                <Alert>alertType={alert.type} alertMessage={alert.message}</Alert>
                <hr />
                <form onSubmit={handleSubmit}>
                    <Input
                        title={'Job Title'}
                        className={hasError("title") ? "is-invalid" : ""}
                        type={'text'}
                        name={'title'}
                        value={job['title']}
                        handleChange={handleChange}
                        errorDiv={hasError("title") ? "text-danger" : "d-none"}
                        errorMsg={"Job Title must be provided"}
                    />
                    <Input
                        title={'Picture URL'}
                        className={hasError("picture") ? "is-invalid" : ""}
                        type={'text'}
                        name={'Picture'}
                        value={job['Picture']}
                        handleChange={handleChange}
                        errorDiv={hasError("picture") ? "text-danger" : "d-none"}
                        errorMsg={"Please provide a representative picture of this job"}
                    />
                    <Input
                        title={'Earning Rate'}
                        className={hasError("earning") ? "is-invalid" : ""}
                        type={'text'}
                        name={'Earning'}
                        value={job['Earning']}
                        handleChange={handleChange}
                        errorDiv={hasError("earning") ? "text-danger" : "d-none"}
                        errorMsg={"Job's Earning Rate must be provided"}
                    />
                    <Input
                        title={'Like'}
                        className={hasError("like") ? "is-invalid" : ""}
                        type={'text'}
                        name={'Like'}
                        isArray={true}
                        value={job['Like']}
                        handleChange={handleChange}
                        errorDiv={hasError("like") ? "text-danger" : "d-none"}
                        errorMsg={"Required Likes should be provided (comma seperated)"}
                    />
                    <Input
                        title={'Good Ats'}
                        className={hasError("good at") ? "is-invalid" : ""}
                        type={'text'}
                        name={'Good At'}
                        isArray={true}
                        value={job['Good At']}
                        handleChange={handleChange}
                        errorDiv={hasError("good at") ? "text-danger" : "d-none"}
                        errorMsg={"Required Good Ats should be provided (comma seperated)"}
                    />
                    
                    <hr />
                    <div className='d-flex justify-content-center'>
                        <button className='btn btn-primary'>Save</button>
                        <Link to="/" className='btn btn-warning ms-1'>Cancel</Link>
                        {
                            job['title'] !== "" && (
                                <a href='#!' onClick={confirmDelete}
                                    className='btn btn-danger ms-1'>Delete</a>
                            )
                        }
                    </div>
                </form>
            </Fragment>
        )
    }

}