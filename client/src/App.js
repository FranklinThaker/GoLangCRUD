import React from 'react';
import { BrowserRouter as Router, Route } from "react-router-dom";
import Home from './Home'
import EditComponent from './EditComponent'
import InsertComponent from './InsertComponent'

const App = () => {
    return (
        <React.Fragment>
            <Router>
                <div>
                    <Route exact path='/' component={Home} />
                    <Route exact path='/edit/:id' component={EditComponent} />
                    <Route exact path='/insert' component={InsertComponent} />
                </div>
            </Router>
        </React.Fragment>
    )
}

export default App;
