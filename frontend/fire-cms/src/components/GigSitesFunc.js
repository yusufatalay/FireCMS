import React, { useState, useEffect, Fragment } from 'react'
import { Link } from 'react-router-dom'

export default function GigSitesFunc(props) {
    const [gigsites, setGigSites] = useState([])
    const [error, setError] = useState(null)

    useEffect(() => {
        fetch("http://localhost:4000/gigsites")
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
                setGigSites(json.gigsites)
            })
    }, [])

    if (error != null) {
        return (<div>{error.message}</div>)
    } else {
        return (
            <Fragment>
                <h2> All Gig Sites</h2>
                <div className='float-end' >
                    <a href='../savegigsite/newgigsite' // add onclick
                        className='btn btn-primary ms-1'>Add New Gig Site</a>
                </div>
                <div className='list-group'>
                    {gigsites.map((g, i) => (
                        <Link key={i} className='list-group-item list-group-item-item' to={`savegigsite/${g["Gig Company Name"]}%20${g["Profession Name"]}`}>{g["Gig Company Name"]} {g["Profession Name"]}</Link>
                    ))}
                </div>
            </Fragment>
        )
    }
}

