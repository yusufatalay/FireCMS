import React from 'react'
import { BrowserRouter as Router, Route, Link, Switch } from 'react-router-dom'
import CoursesFunc from './components/CoursesFunc'
import EditCourseFunc from './components/EditCourseFunc'
import EditGigSiteFunc from './components/EditGigSiteFunc'
import GigSitesFunc from './components/GigSitesFunc'
import JobsFunc from './components/JobsFunc'
import EditJobFunc from './components/EditJobFunc'

function AppFunc(props) {
  // put some authentication checks in here later.

  return (
    <Router>
      <div className='container'>
        <div className='row'>
          <div className='col mt-3'>
            <h1 className='mt-3'>
              FireCMS
            </h1>
          </div>
          <hr className='mb-3'></hr>
        </div>

        <div className='row'>
          <div className='col-md-2'>
            <nav>
              <ul className='list-group'>
                <li className='list-group-item'>
                  <Link to="/">Home</Link>
                </li>
                <li className='list-group-item'>
                  <Link to="/courses">Courses</Link>
                </li>
                <li className='list-group-item'>
                  <Link to="/gigsites">Gig Sites</Link>
                </li>
                <li className='list-group-item'>
                  <Link to="/jobs">Jobs</Link>
                </li>
              </ul>
            </nav>
          </div>
          <div className='col-md-10'>
            <Switch>
              <Route path="/savecourse/:courseName" component={(props) => (
                <EditCourseFunc {...props} /* add jwt control here as well *//>
              )}></Route>
              <Route path="/courses">
                <CoursesFunc />
              </Route>

              <Route path="/savegigsite/:gigsiteName" component={(props)=>(
                <EditGigSiteFunc {...props}/>
              )}></Route>

              <Route path="/gigsites">
                <GigSitesFunc/>
              </Route>
              

              <Route path="/savejob/:jobName" component={(props)=>(
                <EditJobFunc {...props}/>
              )}></Route>

              <Route path="/jobs">
                <JobsFunc/>
              </Route>

              {/* <Route path="/">
                <HomeFunc />
              </Route> */}
            </Switch>
          </div>
        </div>
      </div>
    </Router>
  )
}


export default AppFunc