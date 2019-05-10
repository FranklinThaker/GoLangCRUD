import React, { Component } from 'react';
import axios from 'axios';
import './App.css'
let formdata = {
  Fname: '',
  Lname: '',
  Email: '',
  Age: '',
  Mobile: '',
}
class InsertComponent extends Component {

  constructor(props) {
    super(props);
    this.state = {
      formdata: { ...formdata },
    }
    this.handleInsertdata = this.handleInsertdata.bind(this)
    this.ChangeValue = this.ChangeValue.bind(this);
  }


  handleInsertdata(e) {

    e.preventDefault();
    const fpt = {
      Fname: this.state.formdata.Fname,
      Lname: this.state.formdata.Lname,
      Email: this.state.formdata.Email,
      Age: parseInt(this.state.formdata.Age),
      Mobile: parseInt(this.state.formdata.Mobile)
    }

    const myobj = JSON.stringify(fpt)    
      axios.post(`${process.env.REACT_APP_SERVER}/api/students/create/`, myobj, {
      headers: {
        'content-type': 'application/json'
      }
    })
      .then((response) => {
        //window.location.href = 'http://localhost:3000/'
        this.props.history.push("/")
      })
      .catch(function (error) {
        console.log(error);
      })
  }

  ChangeValue(e, currentUser, field2) {
    e.preventDefault();
    const temp = { ...this.state[currentUser] };
    temp[field2] = e.target.value;
    this.setState({ [currentUser]: temp });
  }

  render() {
    return (
      <div>
        <>
          <form>
            <h1>Insert Student Records</h1>
            <div className="question">
              <input type="text" value={this.state.formdata.Fname} onChange={(e) => this.ChangeValue(e, 'formdata', 'Fname')} required />
              <label>First Name</label>
            </div>
            <div className="question">
              <input type="text" value={this.state.formdata.Lname} onChange={(e) => this.ChangeValue(e, 'formdata', 'Lname')} required />
              <label>Last Name</label>
            </div>
            <div className="question">
              <input type="text" value={this.state.formdata.Email} onChange={(e) => this.ChangeValue(e, 'formdata', 'Email')} required />
              <label>Email Address</label>
            </div>
            <div className="question">
              <input type="text" minLength="1" maxLength="3" value={this.state.formdata.Age} onChange={(e) => this.ChangeValue(e, 'formdata', 'Age')} required />
              <label>Age</label>
            </div>
            <div className="question">
              <input type="text" value={this.state.formdata.Mobile} onChange={(e) => this.ChangeValue(e, 'formdata', 'Mobile')} required />
              <label>Mobile</label>
            </div><br />
            <input type="button" value="SUBMIT" className="btn btn-primary" onClick={(e) => this.handleInsertdata(e)}/>&nbsp;
            <input type="button" value="Home" className="btn btn-info" onClick={() => this.props.history.push('/')} />
          </form>
        </>
      </div>
    )
  }
}

export default InsertComponent;
