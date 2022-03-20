import React, { useState, useEffect, Fragment } from 'react'
import { Link } from 'react-router-dom';
import Alert from './ui-components/Alert';
import { confirmAlert } from 'react-confirm-alert'; // Import
import 'react-confirm-alert/src/react-confirm-alert.css'; // Import css
import Input from './form-components/Input';
import TextArea from './form-components/TextArea';

export default function EditGigSiteFunc(props) {
    const [gigsite, setGigSite] = useState({
        "BizModel": "",
        "Comments-(2nd phase)": "",
        "Definition / About": "",
        "Gig Company Logo(dummy pics)": "",
        "Gig Company Name": "",
        "Landing Page": "",
        "Profession Name": "",
        "Rating-(2nd phase)": "",
        "Remote": "",
        "Requirements": "",
        "Tips": "",
        "Where": "",
        "isSaved": "",
    })

    const [error, setError] = useState(null)
    const [errors, setErrors] = useState([])
    const [alert, setAlert] = useState({ type: "d-none", message: "" })

    const handleSubmit = (evt) => {
        evt.preventDefault();
        // do client side validation
        let errors = []
        if (gigsite['BizModel'] === "") {
            errors.push("bizmodel")
        }
        if (gigsite['Definition / About'] === "") {
            errors.push("about")
        }
        if (gigsite['Gig Company Logo(dummy pics)'] === "") {
            errors.push("logo")
        }
        if (gigsite['Gig Company Name'] === "") {
            errors.push("company name")
        }
        if (gigsite['Landing Page'] === "") {
            errors.push("landing page")
        }
        if (gigsite['Profession Name'] === "") {
            errors.push("profession name")
        }
        if (gigsite['Remote'] === "") {
            errors.push("remote")
        }
        if (gigsite['Requirements'] === "") {
            errors.push("requirements")
        }
        if (gigsite['Tips'] === "") {
            errors.push("tips")
        }
        if (gigsite['Where'] === "") {
            errors.push("where")
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

        fetch("http://localhost:4000/savegigsite", requestOptions)
            .then(response => response.json())
            .then(data => {
                if (data.error) {
                    setAlert({ type: "alert-danger", message: data.error.message })
                } else {
                    props.history.push('/..')
                }
            })
        setGigSite(payload)
    }

    const handleChange = (evt) => {
        let value = evt.target.value;
        let name = evt.target.name;

        setGigSite((prevState) => ({
            ...prevState,
            [name]: value,
        }))
    }

    function hasError(key) {
        return errors.indexOf(key) !== -1
    }

    useEffect(() => {
        // also do auth check here

        // if no gigsiteName given as parameter then it means that user is going to create a new gigsite
        // like -> /saveCourse/
        const gigsiteName = props.match.params.gigsiteName
        // user will be directed to /savegigsite/newgigsite when he/she tries to save a new gigsite
        if (gigsiteName === "newgigsite") {
            setGigSite({
                "BizModel": "",
                "Comments-(2nd phase)": "",
                "Definition / About": "",
                "Gig Company Logo(dummy pics)": "",
                "Gig Company Name": "",
                "Landing Page": "",
                "Profession Name": "",
                "Rating-(2nd phase)": "",
                "Remote": "",
                "Requirements": "",
                "Tips": "",
                "Where": "",
                "isSaved": "",
            })
        } else if (gigsiteName !== "") {
            // if a gigsiteName has given then it means user is going to edit that gigsite
            fetch("http://localhost:4000/gigsite/" + gigsiteName)
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
                    setGigSite(js.gigsite)
                })
        }
    }, [props.history, props.match.params.gigsiteName])

    const confirmDelete = (e) => {
        console.error(gigsite)
        confirmAlert({
            title: 'Delete Gig Site ?.',
            message: 'Are you sure to DELETE this gigsite?.',
            buttons: [
                {
                    label: 'Yes',
                    onClick: () => {
                        const myHeaders = new Headers();
                        myHeaders.append("Content-Type", "application/json")
                        // also do security check
                        fetch("http://localhost:4000/deletegigsite/" + gigsite["Gig Company Name"] + "%20" +gigsite["Profession Name"], { method: "GET", headers: myHeaders })
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
                        title={'Gig Company Name'}
                        className={hasError("company name") ? "is-invalid" : ""}
                        type={'text'}
                        name={'Gig Company Name'}
                        value={gigsite['Gig Company Name']}
                        handleChange={handleChange}
                        errorDiv={hasError("company name") ? "text-danger" : "d-none"}
                        errorMsg={"Company Name must be provided"}
                    />
                    <Input
                        title={'Gig Company Logo'}
                        className={hasError("logo") ? "is-invalid" : ""}
                        type={'text'}
                        name={'Gig Company Logo'}
                        value={gigsite['Gig Company Logo(dummy pics)']}
                        handleChange={handleChange}
                        errorDiv={hasError("logo") ? "text-danger" : "d-none"}
                        errorMsg={"Company Logo URL must be provided"}
                    />
                    <Input
                        title={'Landing Page'}
                        className={hasError("landing page") ? "is-invalid" : ""}
                        type={'text'}
                        name={'Course Picture'}
                        value={gigsite['Landing Page']}
                        handleChange={handleChange}
                        errorDiv={hasError("landing page") ? "text-danger" : "d-none"}
                        errorMsg={"Landing Page Link must be provided"}
                    />
                    <Input
                        title={'Profession Name'}
                        className={hasError("profession name") ? "is-invalid" : ""}
                        type={'text'}
                        name={'Profession Name'}
                        value={gigsite['Profession Name']}
                        handleChange={handleChange}
                        errorDiv={hasError("profession name") ? "text-danger" : "d-none"}
                        errorMsg={"Profession Name must be provided"}
                    />
                    <Input
                        title={'Is Remote'}
                        className={hasError("remote") ? "is-invalid" : ""}
                        type={'text'}
                        name={'Remote'}
                        value={gigsite['Remote']}
                        handleChange={handleChange}
                        errorDiv={hasError("remote") ? "text-danger" : "d-none"}
                        errorMsg={"Plaese state that if this gig is remote or not (yes/no)"}
                    />
                    <Input
                        title={'Business Model'}
                        className={hasError("bizmodel") ? "is-invalid" : ""}
                        type={'text'}
                        name={'BizModel'}
                        value={gigsite['BizModel']}
                        handleChange={handleChange}
                        errorDiv={hasError("bizmodel") ? "text-danger" : "d-none"}
                        errorMsg={"Please state the business model of this gig"}
                    />
                    <Input
                        title={'Location'}
                        className={hasError("where") ? "is-invalid" : ""}
                        type={'text'}
                        name={'Where'}
                        value={gigsite['Where']}
                        handleChange={handleChange}
                        errorDiv={hasError("where") ? "text-danger" : "d-none"}
                        errorMsg={"Gig Site Location must be provided"}
                    />
                     <TextArea
                        title={'Definition'}
                        name={'Definition'}
                        rows={"4"}
                        value={gigsite['Definition / About']}
                        handleChange={handleChange}
                    />
                     <TextArea
                        title={'Requirements'}
                        name={'Requirements'}
                        rows={"4"}
                        value={gigsite['Requirements']}
                        handleChange={handleChange}
                    />
                     <TextArea
                        title={'Tips'}
                        name={'Tips'}
                        rows={"4"}
                        value={gigsite['Tips']}
                        handleChange={handleChange}
                    />
                   
                    <hr />
                    <div className='d-flex justify-content-center'>
                        <button className='btn btn-primary'>Save</button>
                        <Link to="/" className='btn btn-warning ms-1'>Cancel</Link>
                        {
                            gigsite['Course Name'] !== "" && (
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