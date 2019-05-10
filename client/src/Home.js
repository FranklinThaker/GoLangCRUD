import React, { Component } from 'react';
import axios from 'axios';
import { Button } from 'reactstrap';
import { Link } from "react-router-dom";
import 'bootstrap/dist/css/bootstrap.min.css';
import debounce from './debounce'

class Home extends Component {
  constructor(props) {
    super(props);
    this.state = {
      users: [],
      searchTerm: '',
    }

    this.handleDeletedata = this.handleDeletedata.bind(this);
    this.handleSearch = this.handleSearch.bind(this);
    this.fetchAll = this.fetchAll.bind(this);
    this.setSearchTerm = this.setSearchTerm.bind(this);
  }

  setSearchTerm = debounce(x => {
    this.setState({ searchTerm: x })
    this.handleSearch();
  }, 300)

  

  handleSearch() {
    axios.get(`${process.env.REACT_APP_SERVER}/api/students/?search=` + this.state.searchTerm)
      .then(response => {
        this.setState({ users: response.data });
      })
      .catch(function (error) {
        console.log(error);
      });
  }


  componentDidMount() {
    axios.get(`${process.env.REACT_APP_SERVER}/api/students/list`)
      .then(response => {
        this.setState({ users: response.data });
      })
      .catch(function (error) {
        console.log(error);
      })
  }

  

  handleDeletedata(id) {
    axios.delete(`${process.env.REACT_APP_SERVER}/api/students/delete/?id=` + id)
      .then((response) => {
       
        
        //window.location.href='http://localhost:3000/'
        // //this.props.history.push("/");
        // let users = fpt.state.users.filter(function (user) { 
        //   return user.id !== id 
        //   });
        //   fpt.setState({ users: users });
        this.fetchAll();
      })
      .catch(function (error) {
        console.log(error);
      })
  }

  fetchAll() {
    axios.get(`${process.env.REACT_APP_SERVER}/api/students/list`)
      .then(response => {
        this.setState({ users: response.data });
      })
      .catch(function (error) {
        console.log(error);
      })
  }

  render() {
    return (
      <>
        <nav className="navbar navbar-expand-sm bg-dark navbar-dark">
          <input className="form-control mr-sm-2" type="text" placeholder="Search" onChange={(e) => this.setSearchTerm(e.target.value)} />
          <Link to={`/insert/`}><Button color="primary">Insert!</Button></Link>
        </nav>

        <div>
          <table className="table table-bordered table-dark">
            <thead>
              <tr>
                <th>Id</th>
                <th>First Name</th>
                <th>Last Name</th>
                <th>Em@il</th>
                <th>Age</th>
                <th>Mobile</th>
                <th>Edit</th>
                <th>Delete</th>
              </tr>
            </thead>
            <tbody>
              {this.state.users.map(data =>
                <tr key={data.ID}>
                  <td>{data.ID} </td>
                  <td>{data.Fname} </td>
                  <td>{data.Lname}</td>
                  <td>{data.Email}</td>
                  <td>{data.Age}</td>
                  <td>{data.Mobile}</td>
                  <td><Link to={`/edit/${data.ID}`} ><Button color="primary">Edit</Button></Link></td>
                  <td><Button color="danger" onClick={() => this.handleDeletedata(data.ID)}>Delete!</Button></td>
                </tr>
              )}
            </tbody>
          </table>
        </div>
      </>
    );
  }
}


export default Home;
