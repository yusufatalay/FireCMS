import React, { useState, useEffect, Fragment } from 'react'
import { Link } from 'react-router-dom';
import Alert from './ui-components/Alert';
import { confirmAlert } from 'react-confirm-alert'; // Import
import Input from './form-components/Input';


export default function OneCourseFunc(props) {
    const [course, setCourse] = useState({
        "Course Name": "",
        "Course Owner": "",
        "Course Picture": "",
        "Certificate": "",
        "Course (owner) Picture (dummy pics)": "",
        "Introduction Link": "",
        "Location": "",
        "Price": "",
        "Profession Name": "",
        "isSaved": "",
        "rating": 0
    })

    const [error, setError] = useState(null)
    const [errors, setErrors] = useState([])
    const [alert, setAlert] = useState({ type: "d-none", message: "" })

    const handleSubmit = (evt) => {
        evt.preventDefault();
        // do client side validation
        let errors = []
        if (course['Course Name'] === "") {
            errors.push("name")
        }
        if (course['Course Owner'] === "") {
            errors.push("owner")
        }
        if (course['Course Picture'] === "") {
            errors.push("picture")
        }
        if (course['Certificate'] === "") {
            errors.push("certificate")
        }
        if (course['Course (owner) Picture (dummy pics)'] === "") {
            errors.push("owner picture")
        }
        if (course['Introduction Link'] === "") {
            errors.push("intro")
        }
        if (course['Location'] === "") {
            errors.push("location")
        }
        if (course['Price'] === "") {
            errors.push("price")
        }
        if (course['Profession Name'] === "") {
            errors.push("profession name")
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

        fetch("http://localhost:4000/savecourse", requestOptions)
            .then(response => response.json())
            .then(data => {
                if (data.error) {
                    setAlert({ type: "alert-danger", message: data.error.message })
                } else {
                    props.history.push('/..')
                }
            })
    }

    const handleChange = (evt) => {
        let value = evt.target.value;
        let name = evt.target.name;

        setCourse((prevState) => ({
            ...prevState,
            [name]: value,
        }))
    }

    function hasError(key) {
        return errors.indexOf(key) !== -1
    }

    useEffect(() => {
        // also do auth check here

        // if no courseName given as parameter then it means that user is going to create a new course
        // like -> /saveCourse/
        const courseName = props.match.params.courseName
        // if a courseName has given then it means user is going to edit that course
        if (courseName !== "") {
            fetch("http://localhost:4000/course/" + courseName)
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
                    setCourse(js.course)
                })
        }
    }, [props.history, props.match.params.courseName])

    const confirmDelete = (e) => {
        confirmAlert({
            title: 'Delete Course ?.',
            message: 'Are you sure to DELETE this course?.',
            buttons: [
                {
                    label: 'Yes',
                    onClick: () => {
                        const myHeaders = new Headers();
                        myHeaders.append("Content-Type", "application/json")
                        // also do security check
                        fetch("http://localhost:4000/deletecourse/" + course['Course Name'], { method: "GET", headers: myHeaders })
                            .then((response) => response.json())
                            .then(data => {
                                if (data.error) {
                                    setAlert({ type: "alert-danger", message: data.error.message })
                                } else {
                                    setAlert({ type: "alert-success", message: "course deleted" })

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
                <h2>Save Course</h2>
                <Alert>alertType={alert.type} alertMessage={alert.message}</Alert>
                <hr />
                <form onSubmit={handleSubmit}>
                    <Input
                        title={'Course Name'}
                        className={hasError("name") ? "is-invalid" : ""}
                        type={'text'}
                        name={'Course Name'}
                        value={course['Course Name']}
                        handleChange={handleChange}
                        errorDiv={hasError("name") ? "text-danger" : "d-none"}
                        errorMsg={"Course Name must be provided"}
                    />
                    <Input
                        title={'Course Owne'}
                        className={hasError("owner") ? "is-invalid" : ""}
                        type={'text'}
                        name={'Course Owner'}
                        value={course['Course Owner']}
                        handleChange={handleChange}
                        errorDiv={hasError("owner") ? "text-danger" : "d-none"}
                        errorMsg={"Course Owner must be provided"}
                    />
                    <Input
                        title={'Course Picture'}
                        className={hasError("picture") ? "is-invalid" : ""}
                        type={'text'}
                        name={'Course Picture'}
                        value={course['Course Picture']}
                        handleChange={handleChange}
                        errorDiv={hasError("picture") ? "text-danger" : "d-none"}
                        errorMsg={"Course Picture Link must be provided"}
                    />
                    <Input
                        title={'Certificate'}
                        className={hasError("certificate") ? "is-invalid" : ""}
                        type={'text'}
                        name={'Certificate'}
                        value={course['Certificate']}
                        handleChange={handleChange}
                        errorDiv={hasError("certificate") ? "text-danger" : "d-none"}
                        errorMsg={"Course Certificate availability must be noted (yes/no)"}
                    />
                    <Input
                        title={'Course (owner) Picture'}
                        className={hasError("owner picture") ? "is-invalid" : ""}
                        type={'text'}
                        name={'Course (owner) Picture (dummy pics)'}
                        value={course['Course (owner) Picture (dummy pics)']}
                        handleChange={handleChange}
                        errorDiv={hasError("owner picture") ? "text-danger" : "d-none"}
                        errorMsg={"Course Owner Picture  must be provided"}
                    />
                    <Input
                        title={'Introduction Link'}
                        className={hasError("intro") ? "is-invalid" : ""}
                        type={'text'}
                        name={'Introduction Link'}
                        value={course['Introduction Link']}
                        handleChange={handleChange}
                        errorDiv={hasError("intro") ? "text-danger" : "d-none"}
                        errorMsg={"Introduction Link must be provided"}
                    />
                    <Input
                        title={'Location'}
                        className={hasError("location") ? "is-invalid" : ""}
                        type={'text'}
                        name={'Location'}
                        value={course['Location']}
                        handleChange={handleChange}
                        errorDiv={hasError("location") ? "text-danger" : "d-none"}
                        errorMsg={"Course Location must be provided"}
                    />
                    <Input
                        title={'Price'}
                        className={hasError("price") ? "is-invalid" : ""}
                        type={'text'}
                        name={'Price'}
                        value={course['Price']}
                        handleChange={handleChange}
                        errorDiv={hasError("price") ? "text-danger" : "d-none"}
                        errorMsg={"Course Price must be provided"}
                    />
                    <Input
                        title={'Profession Name'}
                        className={hasError("profession name") ? "is-invalid" : ""}
                        type={'text'}
                        name={'Profession Name'}
                        value={course['Profession Name']}
                        handleChange={handleChange}
                        errorDiv={hasError("profession name") ? "text-danger" : "d-none"}
                        errorMsg={"Profession Name must be provided"}
                    />
                    {/*  rating should't be setted in here , think another way */}
                    <hr />
                    <div className='d-flex justify-content-center'>
                        <button className='btn btn-primary'>Save</button>
                        <Link to="/" className='btn btn-warning ms-1'>Cancel</Link>
                        {
                            course['Course Name'] !== "" && (
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